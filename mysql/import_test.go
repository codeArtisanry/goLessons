package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

// CreateTestDatabase will create a test-database and return the db, db name, and cleanup function
func CreateTestDatabase(t *testing.T, dbHost, dbPort, dbUser, dbPassword, dbName string) (*sql.DB, string, func()) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, dbErr := sql.Open("mysql", connectionString)
	if dbErr != nil {
		t.Fatalf("Fail to connect database. %s", dbErr.Error())
	}

	rand.Seed(time.Now().UnixNano())
	testDBName := "test" + strconv.FormatInt(rand.Int63(), 10)

	_, err := db.Exec("CREATE DATABASE " + testDBName)
	if err != nil {
		t.Fatalf("Fail to create database %s. %s", testDBName, err.Error())
	}

	testConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbUser, dbPassword, dbHost, dbPort, testDBName)
	testDB, dbErr := sql.Open("mysql", testConnectionString)
	if dbErr != nil {
		t.Fatalf("Fail to connect database. %s", dbErr.Error())
	}

	return testDB, testDBName, func() {
		_, err := db.Exec("DROP DATABASE " + testDBName)
		if err != nil {
			t.Fatalf("Fail to drop database %s. %s", testDBName, err.Error())
		}
	}
}

// LoadFixtures for loading data for testing
func LoadFixtures(t *testing.T, db *sql.DB, fixtureName string) {
	content, err := ioutil.ReadFile(fmt.Sprintf("./testdata/%s.sql", fixtureName))
	ok(t, err)

	queries := strings.Split(string(content), ";")
	for _, query := range queries {
		if strings.TrimSpace(query) != "" {
			_, err := db.Exec(query)
			ok(t, err)
		}
	}
}

// LoadSchema for loading the schema of the database. Filepath can be a relative path from the test file
func LoadSchema(t *testing.T, db *sql.DB, filepath string) {
	content, err := ioutil.ReadFile(filepath)
	ok(t, err)

	queries := strings.Split(string(content), ";")
	for _, query := range queries {
		if strings.TrimSpace(query) != "" {
			_, err := db.Exec(query)
			ok(t, err)
		}
	}
}
