package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	numProcsPtr := flag.Int("gomaxprocs", 0, "GOMAXPROCS")
	flag.Parse()
	numProcs := *numProcsPtr
	returnedProcs := runtime.GOMAXPROCS(numProcs)
	if numProcs == 0 {
		numProcs = returnedProcs
	}

	fmt.Print("starting benchmark w/ GOMAXPROCS =", numProcs, "\n")

	var queries uint64 = 0
	var numThreads int = 32

	var wg sync.WaitGroup
	wg.Add(numThreads)

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/mysql_benchmark?charset=utf8&autocommit=false")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	startTime := time.Now()

	for threadNum := 1; threadNum <= numThreads; threadNum++ {
		go func() {
			defer wg.Done()
			for i := 1; i <= 100000; i++ {
				rows, err := db.Query("SELECT 1") // WHERE number = 13
				if err != nil {
					panic(err.Error()) // proper error handling instead of panic in your app
				}
				defer rows.Close()

				for rows.Next() {
					atomic.AddUint64(&queries, 1)
					/*
						var name string
						if err := rows.Scan(&name); err != nil {
							log.Fatal(err)
						}*/
				}
			}
		}()
	}

	fmt.Printf("Waiting for queries to complete...\n")
	wg.Wait()

	totalTime := time.Now().Sub(startTime)
	fmt.Printf("Time taken: %s\n", totalTime)
	fmt.Printf("Queries run: %d\n", queries)
	fmt.Printf("Throughput: %d\n", float64(queries)/totalTime.Seconds())
	log.Printf("done!")
}