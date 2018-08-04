package meddler

import (
	"database/sql"
	"fmt"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/russross/meddler"
)

var once sync.Once
var db *sql.DB
var when = time.Date(2013, 6, 23, 15, 30, 12, 0, time.UTC)

type Person struct {
	ID        int64  `meddler:"id,pk"`
	Name      string `meddler:"name"`
	private   int
	Email     string
	Ephemeral int        `meddler:"-"`
	Age       int        `meddler:",zeroisnull"`
	Opened    time.Time  `meddler:"opened,utctime"`
	Closed    time.Time  `meddler:"closed,utctimez"`
	Updated   *time.Time `meddler:"updated,localtime"`
	Height    *int       `meddler:"height"`
}

type HalfPerson struct {
	ID        int64 `meddler:"id,pk"`
	private   int
	Ephemeral int        `meddler:"-"`
	Age       int        `meddler:",zeroisnull"`
	Closed    time.Time  `meddler:"closed,utctimez"`
	Updated   *time.Time `meddler:"updated,localtime"`
}

type UintPerson struct {
	ID        uint64 `meddler:"id,pk"`
	Name      string `meddler:"name"`
	private   int
	Email     string
	Ephemeral int        `meddler:"-"`
	Age       int        `meddler:",zeroisnull"`
	Opened    time.Time  `meddler:"opened,utctime"`
	Closed    time.Time  `meddler:"closed,utctimez"`
	Updated   *time.Time `meddler:"updated,localtime"`
	Height    *int       `meddler:"height"`
}

const schema1 = `create table if not exists person (
	id integer primary key,
	name text not null,
	Email text not null,
	Age integer,
	opened datetime not null,
	closed datetime,
	updated datetime,
	height integer
)`

const schema2 = `create table if not exists item  (
	id integer primary key,
	stuff text not null,
	stuffz blob not null
)`

var aliceHeight int = 65
var alice = &Person{
	Name:      "Alice",
	Email:     "alice@alice.com",
	Ephemeral: 12,
	Age:       32,
	Opened:    when.Local(),
	Closed:    when,
	Updated:   &when,
	Height:    &aliceHeight,
}

var bob = &Person{
	Name:   "Bob",
	Email:  "bob@bob.com",
	Opened: when,
}

func setup() {
	var err error

	// create the database
	// db, err = sql.Open("sqlite3", ":memory:")
	db, err = sql.Open("sqlite3", "persion.db")
	if err != nil {
		panic("error creating test database: " + err.Error())
	}

	// create the tables
	if _, err = db.Exec(schema1); err != nil {
		panic("error creating person table: " + err.Error())
	}
	if _, err = db.Exec(schema2); err != nil {
		panic("error creating item table: " + err.Error())
	}
}

func personEqual(t *testing.T, elt *Person, ref *Person) {
	if elt == nil {
		t.Errorf("Person %s is nil", ref.Name)
		return
	}
	if elt.ID != ref.ID {
		t.Errorf("Person %s ID is %v", ref.Name, elt.ID)
	}
	if elt.Name != ref.Name {
		t.Errorf("Person %s Name is %v", ref.Name, elt.Name)
	}
	if elt.private != ref.private {
		t.Errorf("Person %s private is %v", ref.Name, elt.private)
	}
	if elt.Email != ref.Email {
		t.Errorf("Person %s Email is %v", ref.Name, elt.Email)
	}
	if elt.Ephemeral != ref.Ephemeral {
		t.Errorf("Person %d Ephemeral is %d", ref.Ephemeral, elt.Ephemeral)
	}
	if elt.Age != ref.Age {
		t.Errorf("Person %s Age is %v", ref.Name, elt.Age)
	}
	if !elt.Opened.Equal(ref.Opened) {
		t.Errorf("Person %s Opened is %v", ref.Name, elt.Opened)
	}
	if !elt.Closed.Equal(ref.Closed) {
		t.Errorf("Person %s Closed is %v", ref.Name, elt.Closed)
	}
	if (elt.Updated == nil) != (ref.Updated == nil) {
		t.Errorf("Person %s Updated == nil is %v", ref.Name, elt.Updated == nil)
	} else if elt.Updated != nil && !elt.Updated.Equal(*ref.Updated) {
		t.Errorf("Person %s Updated is %v", ref.Name, *elt.Updated)
	}
	if elt.Updated != nil {
		zone, _ := elt.Updated.Zone()
		local, _ := when.Local().Zone()
		if zone != local {
			t.Errorf("Person %s Updated in time zone %v, expected %v", ref.Name, zone, local)
		}
	}
	if (elt.Height == nil) != (ref.Height == nil) {
		t.Errorf("Person %s Height == nil is %v", ref.Name, elt.Height == nil)
	} else if elt.Height != nil && *elt.Height != *ref.Height {
		t.Errorf("Person %s Height is %v", ref.Name, *elt.Height)
	}
}

func TestInsertAliceBob(t *testing.T) {
	once.Do(setup)
	// insert Alice as row #1
	alice.ID = 0
	if err := meddler.Insert(db, "person", alice); err != nil {
		t.Errorf("Error inserting Alice: %v", err)
	}
	if alice.ID != 1 {
		t.Errorf("Alice ID is %d, expecting 1", alice.ID)
	}

	// insert Bob as row #2
	bob.ID = 0
	if err := meddler.Insert(db, "person", bob); err != nil {
		t.Errorf("Error inserting Bob: %v", err)
	}
	if bob.ID != 2 {
		t.Errorf("Bob ID is %d, expecting 2", bob.ID)
	}
}

func TestColumns(t *testing.T) {
	once.Do(setup)

	p := new(Person)
	names, err := meddler.Columns(p, true)
	if err != nil {
		t.Errorf("Error getting Columns: %v", err)
	}
	fmt.Println(names)
	expected := []string{"id", "name", "Email", "Age", "opened", "closed", "updated", "height"}
	sort.Strings(expected)

	if len(names) != len(expected) {
		t.Errorf("Expected %d columns, got %d", len(expected), len(names))
	}
	sort.Strings(names)
	for i := 0; i < len(expected); i++ {
		if expected[i] != names[i] {
			t.Errorf("Expected %s at position %d, got %s", expected[i], i, names[i])
		}
	}
}

func TestColumnsQuoted(t *testing.T) {
	once.Do(setup)

	p := new(Person)
	names, err := meddler.ColumnsQuoted(p, true)
	if err != nil {
		t.Errorf("Error getting ColumnsQuoted: %v", err)
	}
	fmt.Println(names)
	lst := []string{"id", "name", "Email", "Age", "opened", "closed", "updated", "height"}
	sort.Strings(lst)

	expected := strings.Join(lst, ",")

	if len(names) != len(expected) {
		t.Errorf("Length mismatch: expected %d, got %d", len(expected), len(names))
	}

	fields := strings.Split(names, ",")
	sort.Strings(fields)
	names = strings.Join(fields, ",")

	if expected != names {
		t.Errorf("Mismatch: expected %s, got %s", expected, names)
	}
}

func TestPrimaryKey(t *testing.T) {
	p := new(Person)
	p.ID = 56
	name, val, err := meddler.PrimaryKey(p)
	if err != nil {
		t.Errorf("Error getting PrimaryKey: %v", err)
	}
	if name != "id" {
		t.Errorf("Expected pk name to be id, found %s", name)
	}
	if val != 56 {
		t.Errorf("Expected pk value to be 56, found %d", val)
	}
	fmt.Println(name, val)
	p2 := new(UintPerson)
	p2.ID = 56
	name, val, err = meddler.PrimaryKey(p2)
	if err != nil {
		t.Errorf("Error getting PrimaryKey: %v", err)
	}
	if name != "id" {
		t.Errorf("Expected pk name to be id, found %s", name)
	}
	if val != 56 {
		t.Errorf("Expected pk value to be 56, found %d", val)
	}
}

func TestSetPrimaryKey(t *testing.T) {
	p := new(Person)
	err := meddler.SetPrimaryKey(p, 14)
	if err != nil {
		t.Errorf("Error in SetPrimaryKey: %v", err)
	}
	if p.ID != 14 {
		t.Errorf("Expected id to be 14, found %d", p.ID)
	}

	p2 := new(Person)
	err = meddler.SetPrimaryKey(p2, 14)
	if err != nil {
		t.Errorf("Error in SetPrimaryKey: %v", err)
	}
	if p2.ID != 14 {
		t.Errorf("Expected id to be 14, found %d", p2.ID)
	}
}

func TestValues(t *testing.T) {
	alice.ID = 15
	lst, err := meddler.Values(alice, true)
	if err != nil {
		t.Errorf("Values error: %v", err)
	}

	if lst[0] != int64(15) {
		t.Errorf("expected 15, got %v", lst[0])
	}
	if lst[1] != "Alice" {
		t.Errorf("Expected Alice, got %v", lst[1])
	}
	if lst[2] != "alice@alice.com" {
		t.Errorf("Expected alice@alice.com, got %v", lst[2])
	}
	if lst[3] != 32 {
		t.Errorf("Expected 32, got %v", lst[3])
	}
	if lst[4] != when.UTC() {
		t.Errorf("Expected %v, got %v", when.UTC(), lst[4])
	}
	if lst[5] != when.UTC() {
		t.Errorf("Expected %v, got %v", when.UTC(), lst[5])
	}
	if lst[6] != when.UTC() {
		t.Errorf("Expected %v, got %v", when.UTC(), lst[6])
	}
	if *(lst[7].(*int)) != aliceHeight {
		t.Errorf("Expected %d, got %v", aliceHeight, lst[7])
	}

	lst, err = meddler.Values(alice, false)
	if err != nil {
		t.Errorf("Values error: %v", err)
	}
	if lst[0] != "Alice" {
		t.Errorf("Expected Alice, got %v", lst[0])
	}
}

func TestPlaceholders(t *testing.T) {
	lst, err := meddler.MySQL.Placeholders(alice, true)
	if err != nil {
		t.Errorf("Error in Placeholders: %v", err)
	}
	if len(lst) != 8 {
		t.Errorf("expected 8 items, found %d", len(lst))
	}
	fmt.Println(lst)
	for _, elt := range lst {
		if elt != meddler.MySQL.Placeholder {
			t.Errorf("expected %s, found %s", meddler.MySQL.Placeholder, elt)
		}
	}

	lst, err = meddler.PostgreSQL.Placeholders(alice, false)
	if err != nil {
		t.Errorf("Error in Placeholders: %v", err)
	}
	if len(lst) != 7 {
		t.Errorf("expected 7 items, found %d", len(lst))
	}
	for i, elt := range lst {
		expected := fmt.Sprintf("$%d", i+1)
		if expected != elt {
			t.Errorf("expected %s, found %s", expected, elt)
		}
	}
}

func TestPlaceholdersString(t *testing.T) {
	s, err := meddler.SQLite.PlaceholdersString(alice, false)
	if err != nil {
		t.Errorf("Error in PlaceholdersString: %v", err)
	}
	expected := "?,?,?,?,?,?,?"
	if s != expected {
		t.Errorf("expected %s, found %s", expected, s)
	}

	s, err = meddler.PostgreSQL.PlaceholdersString(alice, true)
	if err != nil {
		t.Errorf("Error in PlaceholdersString: %v", err)
	}
	expected = "$1,$2,$3,$4,$5,$6,$7,$8"
	if s != expected {
		t.Errorf("expected %s, found %s", expected, s)
	}
}

func TestScanRow(t *testing.T) {
	once.Do(setup)
	// insertAliceBob(t)

	rows, err := db.Query("select * from person where id in (1,2) order by id")
	if err != nil {
		t.Errorf("DB error on query: %v", err)
		return
	}

	alice := new(Person)
	if err = meddler.Scan(rows, alice); err != nil {
		t.Errorf("Scan error on Alice: %v", err)
		return
	}

	bob := new(Person)
	bob.Age = 50
	bob.Closed = time.Now()
	bob.private = 14
	bob.Ephemeral = 16
	if err = meddler.ScanRow(rows, bob); err != nil {
		t.Errorf("ScanRow error on Bob: %v", err)
		return
	}

	height := 65
	personEqual(t, alice, &Person{1, "Alice", 0, "alice@alice.com", 0, 32, when, when, &when, &height})
	personEqual(t, bob, &Person{2, "Bob", 14, "bob@bob.com", 16, 0, when, time.Time{}, nil, nil})
	// db.Exec("delete from person")
}

func TestScanAll(t *testing.T) {
	once.Do(setup)
	// insertAliceBob(t)

	rows, err := db.Query("select * from person order by id")
	if err != nil {
		t.Errorf("DB error on query: %v", err)
		return
	}

	var lst []*Person
	if err = meddler.ScanAll(rows, &lst); err != nil {
		t.Errorf("ScanAll error: %v", err)
		return
	}

	if len(lst) != 2 {
		t.Errorf("ScanAll found %d rows, expected 2", len(lst))
		return
	}

	height := 65
	personEqual(t, lst[0], &Person{1, "Alice", 0, "alice@alice.com", 0, 32, when, when, &when, &height})
	personEqual(t, lst[1], &Person{2, "Bob", 0, "bob@bob.com", 0, 0, when, time.Time{}, nil, nil})
	db.Exec("delete from person")
}

func TestThrowAway(t *testing.T) {
	once.Do(setup)
	// insertAliceBob(t)

	meddler.Debug = false
	hp := new(HalfPerson)
	err := meddler.QueryRow(db, hp, "select * from person where id = 1")
	if err != nil {
		t.Errorf("QueryRow error: %v", err)
	}
	meddler.Debug = true
	db.Exec("delete from person")
}
