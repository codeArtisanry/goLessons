package main

import (
	"log"
	"time"

	"upper.io/db.v3/sqlite"
)

// --' example.sql

// DROP TABLE IF EXISTS "birthday";

// CREATE TABLE "birthday" (
//   "name" varchar(50) DEFAULT NULL,
//   "born" varchar(12) DEFAULT NULL
// );
var settings = sqlite.ConnectionURL{
	Database: `example.db`, // Path to database file.
}

type Birthday struct {
	// Maps the "Name" property to the "name" column
	// of the "birthday" table.
	Name string `db:"name"`
	// Maps the "Born" property to the "born" column
	// of the "birthday" table.
	Born time.Time `db:"born"`
}

func main() {

	// Attemping to open the "example.db" database file.
	sess, err := sqlite.Open(settings)
	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}
	defer sess.Close() // Remember to close the database session.

	// Pointing to the "birthday" table.
	birthdayCollection := sess.Collection("birthday")

	// Attempt to remove existing rows (if any).
	err = birthdayCollection.Truncate()
	if err != nil {
		log.Fatalf("Truncate(): %q\n", err)
	}

	// Inserting some rows into the "birthday" table.
	birthdayCollection.Insert(Birthday{
		Name: "Hayao Miyazaki",
		Born: time.Date(1941, time.January, 5, 0, 0, 0, 0, time.Local),
	})

	birthdayCollection.Insert(Birthday{
		Name: "Nobuo Uematsu",
		Born: time.Date(1959, time.March, 21, 0, 0, 0, 0, time.Local),
	})

	birthdayCollection.Insert(Birthday{
		Name: "Hironobu Sakaguchi",
		Born: time.Date(1962, time.November, 25, 0, 0, 0, 0, time.Local),
	})

	// // Let's query for the results we've just inserted.
	// res := birthdayCollection.Find()

	// // Query all results and fill the birthday variable with them.
	// var birthdays []Birthday

	// err = res.All(&birthdays)
	// if err != nil {
	// 	log.Fatalf("res.All(): %q\n", err)
	// }

	// // Printing to stdout.
	// for _, birthday := range birthdays {
	// 	fmt.Printf("%s was born in %s.\n",
	// 		birthday.Name,
	// 		birthday.Born.Format("January 2, 2006"),
	// 	)
	// }
}
