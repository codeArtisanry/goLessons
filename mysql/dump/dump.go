package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func GetDB() *sql.DB {
	db, err := sql.Open("mysql", "lzkp:yqhtfjzm@tcp(192.168.5.100:3306)/shizhi")
	// defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	return db
}
func main() {
	db := GetDB()
	Dump(db)
}

//  Loads the column data of the table
func GetColums(db *sql.DB) error {
	log.Debug(fmt.Sprintf("Get '%s' table columns", t.TableName))
	rows, err := db.Query(fmt.Sprintf(GET_ONE_ROW_FMT, t.TableName))
	if err != nil {
		log.Error(fmt.Sprintf("Error getting '%s' table columns", t.TableName))
		return err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return err
	}

	// Store the colume names in the list
	vals := make([]string, len(cols))
	for i, col := range cols {
		vals[i] = col
	}
	t.Columns = vals

	return err
}

func Dump(db *sql.DB) error {
	row, err := db.Query("show tables")
	if err != nil {
		return err
	}

	tables := make(chan string)

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

	var wg sync.WaitGroup
	statements := make(chan string)
	// Run X dump queries simultaneously.  5 seemed like a good number, maybe 4 or 6 or 8 is better?
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for table := range tables {
				// Build a list of statements to allow us to recreate our table

				var tn, c string
				err := db.QueryRow(fmt.Sprintf("show create table %s", table)).Scan(&tn, &c)
				if err != nil {
					log.Fatalln(err)
				}
				statements <- fmt.Sprintf("drop table if exists %s", table)
				statements <- c

				// if _, err := db.Exec(fmt.Sprintf("select * into outfile './%s.txt' fields terminated by ',' optionally enclosed by '\"' lines terminated by '\n' from %s", table, table)); err != nil {
				// 	log.Fatalln(err)
				// }
			}
			wg.Done()
		}()
	}

	go func() {
		sw, err := os.Create("statements.sql")
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
	return nil
}
