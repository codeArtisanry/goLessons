// Copyright 2015 Aller Media AS.  All rights reserved.
// License: GPL3
package main

import (
	"fmt"
	"github.com/go-zoo/bone"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	// Host and port to bind the server too.
	host string

	// Base directory for where we will store all our dump files.
	workingDirectory string

	// Max allowed number of dumps to run at the same time
	workers uint

	dumpQueue = make(chan *DumpRequest, workers)

	// Indivdual databases may have different connection info
	dbInfo = make(map[string]DatabaseInfo)
)

type DatabaseInfo struct {
	Host, Username, Password string
}

type Dumper interface {
	Dump(*DumpRequest, io.Writer) error
}

type DumpRequest struct {
	ID   string
	Name string
	// Tables listed here will not have their data dumped, but the create statement will be generated
	TableIgnore []string
	// Limit the number of rows to this value
	Limit int

	// Once a job has finished, information about it will be sent on this channel
	Finished chan DumpReady
	// Notify server when the file has been served to the client
	Sent chan bool
}

type DumpReady struct {
	Job   *DumpRequest
	File  string
	Error error
}

// Track current active jobs.
type ActiveJobs struct {
	sync.RWMutex
	jobs map[string][]*DumpRequest
}

func (a ActiveJobs) Add(job *DumpRequest) bool {
	a.Lock()
	var r bool
	if _, ok := a.jobs[job.ID]; !ok {
		a.jobs[job.ID] = make([]*DumpRequest, 0)
		r = true
	}
	a.jobs[job.ID] = append(a.jobs[job.ID], job)
	a.Unlock()
	return r
}

func (a ActiveJobs) Get(job DumpReady) []*DumpRequest {
	a.RLock()
	defer a.RUnlock()
	if j, ok := a.jobs[job.Job.ID]; ok {
		return j
	}
	log.Println("Unknown job", job.Job.ID)
	return nil
}

func (a ActiveJobs) Delete(job DumpReady) {
	a.Lock()
	if _, ok := a.jobs[job.Job.ID]; ok {
		delete(a.jobs, job.Job.ID)
	}
	a.Unlock()
}

func flagParse() {
	host = os.Getenv("DBSYNC_HOST")
	if host == "" {
		log.Fatalln("Missing DBSYNC_HOST (ie: localhost:8080)")
	}
	workingDirectory := os.Getenv("DBSYNC_WORK_DIR")
	if workingDirectory == "" {
		log.Fatalln("Missing DBSYNC_WORK_DIR (ie: /tmp)")
	}

	workersStr := os.Getenv("DBSYNC_WORKERS")
	if workersStr == "" {
		log.Fatalln("Missing DBSYNC_WORKERS (ie: 5)")
	}
	if w, err := strconv.ParseUint(workersStr, 0, 0); err != nil {
		log.Fatalln("Invalid type passed to DBSYNC_WORKERS")
	} else {
		workers = uint(w)
	}

	dbHost := os.Getenv("DBSYNC_DB_HOST")
	if dbHost == "" {
		log.Fatalln("Missing DBSYNC_DB_HOST")
	}
	dbUsername := os.Getenv("DBSYNC_DB_USERNAME")
	if dbUsername == "" {
		log.Fatalln("Missing DBSYNC_DB_USERNAME")
	}
	dbPassword := os.Getenv("DBSYNC_DB_PASSWORD")
	if dbPassword == "" {
		log.Fatalln("Missing DBSYNC_DB_PASSWORD")
	}

	dbInfo["__default"] = DatabaseInfo{
		Host:     dbHost,
		Username: dbUsername,
		Password: dbPassword,
	}

	otherDatabase := os.Getenv("DBSYNC_DB_DATABASE_1")
	if otherDatabase != "" {
		dbHost := os.Getenv("DBSYNC_DB_HOST_1")
		if dbHost == "" {
			log.Fatalln("Missing DBSYNC_DB_HOST_1")
		}
		dbUsername := os.Getenv("DBSYNC_DB_USERNAME_1")
		if dbUsername == "" {
			log.Fatalln("Missing DBSYNC_DB_USERNAME_1")
		}
		dbPassword := os.Getenv("DBSYNC_DB_PASSWORD_1")
		if dbPassword == "" {
			log.Fatalln("Missing DBSYNC_DB_PASSWORD_1")
		}

		for _, v := range strings.Split(otherDatabase, ",") {
			dbInfo[v] = DatabaseInfo{
				Host:     dbHost,
				Username: dbUsername,
				Password: dbPassword,
			}
		}
	}
}

func main() {
	flagParse()

	mux := bone.New()
	mux.NotFound(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("knick-knack paddywhack give a dog a bone, this old man is no longer home"))
	})

	go databaseDumper(workers)

	mux.Get("/database/:name", http.HandlerFunc(dumpHandler))

	log.Printf("Listening on %s\n", host)
	log.Fatal(http.ListenAndServe(host, mux))
}

// Work delegator
func databaseDumper(workers uint) {
	workerQueue := make(chan *DumpRequest, workers)
	finishedJobs := make(chan DumpReady, workers)

	mysqlDump := MySQLDump{}

	// Set up workers to process the work queue
	for i := uint(0); i < workers; i++ {
		go func() {
			for {
				select {
				case job := <-workerQueue:
					// Main workhorse for each worker routine

					file, err := ioutil.TempFile(workingDirectory, "dbd")
					if err != nil {
						log.Fatalln(err)
					}

					// Hacked in possibility to use mulitple database servers for different databases
					if d, ok := dbInfo[job.Name]; ok {
						mysqlDump.Host = d.Host
						mysqlDump.Username = d.Username
						mysqlDump.Password = d.Password
					} else {
						mysqlDump.Host = dbInfo["__default"].Host
						mysqlDump.Username = dbInfo["__default"].Username
						mysqlDump.Password = dbInfo["__default"].Password
					}

					err = mysqlDump.Dump(job, file)

					// We send the location to the file as we need to open the file and stream for each waiting client.  Sending from the same handle causes file corruption.
					// If something has gone wrong then an error will also be sent.
					finishedJobs <- DumpReady{
						Job:   job,
						File:  file.Name(),
						Error: err,
					}
					file.Close()
				}
			}
		}()
	}

	// Maintain a list of all running jobs. Any duplicate jobs we receive while a job is running will not be processed but instead put on a notification list.
	activeJobs := ActiveJobs{
		jobs: make(map[string][]*DumpRequest),
	}

	for {
		select {
		// Add all incoming dump requests to a job worker queue. If two exact matching requests come through, then only one will be actually processed.
		// The other will be placed in a notification queue so it will also get a copy of the output from the first request.
		case job := <-dumpQueue:
			if activeJobs.Add(job) {
				log.Printf("Queueing dump for %s (%#v)", job.ID, job)
				workerQueue <- job
			}
		// Process finished jobs sending notification with relevant information to relevant callers.
		case job := <-finishedJobs:
			log.Printf("Finished dump for %s", job.Job.ID)
			var wg sync.WaitGroup
			waiting := activeJobs.Get(job)
			for _, w := range waiting {
				wg.Add(1)
				go func(d *DumpRequest) {
					d.Finished <- job
					<-d.Sent
					wg.Done()
				}(w)
			}
			// Wait for all files to be sent before cleaning up.
			// Probably should add a timeout for misbehaving clients
			wg.Wait()
			activeJobs.Delete(job)

			err := os.Remove(job.File)
			if err != nil {
				log.Println(err)
			}
			log.Printf("Cleanup complete for %s", job.Job.ID)
		}
	}
}

func dumpHandler(w http.ResponseWriter, r *http.Request) {
	database := bone.GetValue(r, "name")
	ignore := r.FormValue("ignore")
	// ignore=cache*,voting*
	tableIgnore := make([]string, 0)
	if ignore != "" {
		tableIgnore = strings.Split(ignore, ",")
	}
	// limit=5000
	limitStr := r.FormValue("limit")
	var limit int64
	if limitStr != "" {
		var err error
		limit, err = strconv.ParseInt(limitStr, 0, 0)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`Invalid limit variable supplied`))
			log.Println(err)
			return
		}
	}

	log.Printf("Dump of %s requested from %s", database, r.RemoteAddr)

	dumpReady := make(chan DumpReady, 1)
	dumpSent := make(chan bool, 1)
	dr := &DumpRequest{
		ID:          database,
		Name:        database,
		TableIgnore: tableIgnore,
		Limit:       int(limit),
		Finished:    dumpReady,
		Sent:        dumpSent,
	}
	dumpQueue <- dr

	dur := <-dumpReady
	if dur.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprint(dur.Error)))
		log.Println(dur.Error)
	} else {
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s-db-dump.tar.gz", database))
		w.Header().Set("Content-Type", "application/x-gzip")
		w.Header().Set("Content-Transfer-Encoding", "binary")

		http.ServeFile(w, r, dur.File)

		log.Printf("Dump of %s sent to %s", database, r.RemoteAddr)
	}

	//w.(http.Flusher).Flush()

	// Notify that the file has been sent to the client
	dumpSent <- true
}
