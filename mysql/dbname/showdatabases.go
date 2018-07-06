package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:000000@tcp(192.168.5.100:3306)/")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var dbname string
	db.QueryRow("SHOW DATABASES where `database` NOT LIKE 'info%'").Scan(&dbname)
	println(dbname)
	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var dbNames string
	for rows.Next() {
		err := rows.Scan(&dbNames)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dbNames)
	}
}
