// Copyright 2015 Aller Media AS.  All rights reserved.
// License: GPL3
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

type MySQLDump struct {
	Host     string
	Username string
	Password string
	Database string
}

func (m MySQLDump) Dump(job *DumpRequest, w io.Writer) error {

	dumpDirectory, err := ioutil.TempDir(workingDirectory, "ddb")
	if err != nil {
		return err
	}
	defer func() {
		if err := os.RemoveAll(dumpDirectory); err != nil {
			log.Fatalln(err)
		}
	}()

	if err := os.Chmod(dumpDirectory, 0777); err != nil {
		return err
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", m.Username, m.Password, m.Host, job.Name))
	if err != nil {
		return err
	}
	defer db.Close()

	row, err := db.Query("show tables")
	if err != nil {
		return err
	}

	tables := make(chan string)
	statements := make(chan string)
	var wg sync.WaitGroup

	// Get list of all tables in the database and queue them for processing
	go func() {
		for row.Next() {
			var table string
			err := row.Scan(&table)
			if err != nil {
				log.Println(err)
			} else {
				tables <- table
			}
		}
		close(tables)
	}()

	// Run X dump queries simultaneously.  5 seemed like a good number, maybe 4 or 6 or 8 is better?
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
		tableWorker:
			for table := range tables {
				// Build a list of statements to allow us to recreate our table
				var tn, c string
				err := db.QueryRow(fmt.Sprintf("show create table %s", table)).Scan(&tn, &c)
				if err != nil {
					log.Fatalln(err)
				}
				statements <- fmt.Sprintf("drop table if exists %s", table)
				statements <- c

				// We still want the create statement for every table, just not the data
				for _, v := range job.TableIgnore {
					r, _ := regexp.Compile(fmt.Sprintf("^%s", v))
					if r.MatchString(table) {
						log.Println("Ignored table", table)
						goto tableWorker
					}
				}

				var limit string
				if job.Limit != 0 {
					var t, n, kn, seq, ca, co, sp, p, nu, it, com, ic interface{}
					var cn string
					// If no primary key found, then we just take X rows otherwise we take the last x rows order by primary key
					if err := db.QueryRow(fmt.Sprintf("show keys from %s where Key_name = 'PRIMARY'", table)).Scan(&t, &n, &kn, &seq, &cn, &co, &ca, &sp, &p, &nu, &it, &com, &ic); err != nil {
						limit = fmt.Sprintf("limit %d", job.Limit)
					} else {
						//limit = fmt.Sprintf("where %s > ((select max(%s) from %s)-%s+1) order by %s desc", cn, cn, table, job.Limit, cn)
						limit = fmt.Sprintf("order by %s desc limit %d", cn, job.Limit)
					}
				}

				if _, err := db.Exec(fmt.Sprintf("select * into outfile '%s/%s.txt' fields terminated by ',' optionally enclosed by '\"' lines terminated by '\n' from %s %s", dumpDirectory, table, table, limit)); err != nil {
					log.Fatalln(err)
				}
			}
			wg.Done()
		}()
	}

	go func() {
		sw, err := os.Create(filepath.Join(dumpDirectory, string(filepath.Separator), "statements.sql"))
		if err != nil {
			log.Fatalln(err)
		}

		for statement := range statements {
			fmt.Fprintln(sw, statement+";")
		}
		sw.Close()
	}()

	wg.Wait()
	close(statements)

	return TarAndGz(dumpDirectory, w)
}
