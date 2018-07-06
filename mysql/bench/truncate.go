package models

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

const (
	testDBConfigStr = "host=127.0.0.1 port=5432 dbname=pocket_dev user=pocket password=pocket sslmode=disable"
)

func newEmptyDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", testDBConfigStr)
	if err != nil {
		t.Fatalf("cannot open db connection; err= %v", err)
	}

	truncate(t, db)

	return db
}

func truncate(t *testing.T, db *sql.DB) {
	const sqlstr = `
TRUNCATE TABLE users CASCADE
`

	if _, err := db.Exec(sqlstr); err != nil {
		t.Fatalf("cannot truncate db; err= %v", err)
	}
}
