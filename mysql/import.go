package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

// CreateTestDatabase will create a test-database and return the db, db name, and cleanup function
func CreateTestDatabase(dbHost, dbPort, dbUser, dbPassword, dbName string) (*sql.DB, string, func()) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, dbErr := sql.Open("mysql", connectionString)
	if dbErr != nil {
		fmt.Errorf("Fail to connect database. %s", dbErr.Error())
	}

	rand.Seed(time.Now().UnixNano())
	testDBName := "test" + strconv.FormatInt(rand.Int63(), 10)

	_, err := db.Exec("CREATE DATABASE " + testDBName)
	if err != nil {
		fmt.Errorf("Fail to create database %s. %s", testDBName, err.Error())
	}

	testConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbUser, dbPassword, dbHost, dbPort, testDBName)
	testDB, dbErr := sql.Open("mysql", testConnectionString)
	if dbErr != nil {
		fmt.Errorf("Fail to connect database. %s", dbErr.Error())
	}

	return testDB, testDBName, func() {
		_, err := db.Exec("DROP DATABASE " + testDBName)
		if err != nil {
			fmt.Errorf("Fail to drop database %s. %s", testDBName, err.Error())
		}
	}
}

// LoadFixtures for loading data for testing
func LoadFixtures(db *sql.DB, fixtureName string) {
	content, err := ioutil.ReadFile(fmt.Sprintf("%s.sql", fixtureName))
	if err != nil {
		fmt.Errorf("readfile error")
	}
	queries := strings.Split(string(content), ";")
	for _, query := range queries {
		if strings.TrimSpace(query) != "" {
			_, err := db.Exec(query)
			if err != nil {
				fmt.Errorf("readfile error")
			}
		}
	}
}
