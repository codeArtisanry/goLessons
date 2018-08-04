package main

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDatabase(t *testing.T) {

	db, err := sql.Open("mysql", "root:000000@tcp(192.168.5.100:3306)/")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var dbname string
	db.QueryRow("SHOW DATABASES where `database` NOT LIKE 'info%'").Scan(&dbname)
	println(dbname)

	// 2. get all databases
	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println(rows.Columns()) //[Database] nil
	var dbNames []string
	for rows.Next() {
		err := rows.Scan(&dbname)
		if err != nil {
			log.Fatal(err)
		}
		dbNames = append(dbNames, dbname)
		fmt.Println(dbname)
	}
}
