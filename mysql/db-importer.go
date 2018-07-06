package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func mustExec(db *sql.DB, str string) {
	_, err := db.Exec(str)
	if err != nil {
		panic(err)
	}
}

// return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?interpolateParams=%t&autocommit=true&charset=utf8mb4,utf8,latin1", this.User, this.Password, hostname, this.Key.Port, databaseName, interpolateParams)
func main() {
	var (
		admin    = os.Getenv("ADMINID")
		adminpw  = os.Getenv("ADMINPASSWORD")
		user     = os.Getenv("USERID")
		pw       = os.Getenv("USERPASSWORD")
		database = os.Getenv("DB")
		host     = os.Getenv("HOST")
	)

	buf, err := ioutil.ReadFile("schema.sql")
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%s:%s@tcp(%s)/?multiStatements=true", admin, adminpw, host)
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	mustExec(db, "CREATE DATABASE "+database)
	mustExec(db, fmt.Sprintf("DROP USER '%s'@'%%'", user))
	mustExec(db, fmt.Sprintf("CREATE USER '%s'@'%%' IDENTIFIED BY '%s'", user, pw))
	mustExec(db, fmt.Sprintf("GRANT ALL PRIVILEGES ON `%s`.* TO '%s'@'%%'", database, user))
	mustExec(db, "USE "+database)
	mustExec(db, string(buf))
}
