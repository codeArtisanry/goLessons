package bench

import (
	"fmt"
	"testing"

	"github.com/gocraft/dbr"
)

var dbrDB *dbr.Session

func BenchmarkDbrInsert(b *testing.B) {
	conn, err := dbr.Open("mysql", dsn, nil)
	dbrDB = conn.NewSession(nil)
	if err != nil {
		panic(err)
	}
	for x := 0; x < b.N; x++ {
		i++
		var u user
		u.ID = i
		u.Username = fmt.Sprintf("ABCDEFG%d", i)
		u.Password = "HELLO, WORLD!"
		_, err := dbrDB.InsertInto("BrenchmarkTests").
			Columns("ID", "Username", "Password").
			Values(u.ID, u.Username, u.Password).
			Exec()
		if err != nil {
			panic(err)
		}
	}
}
func BenchmarkDbrSelect1(b *testing.B) {
	for x := 0; x < b.N; x++ {
		dbrSelect(1)
	}
}

func BenchmarkDbrSelect10(b *testing.B) {
	for x := 0; x < b.N; x++ {
		dbrSelect(10)
	}
}

func BenchmarkDbrSelect100(b *testing.B) {
	for x := 0; x < b.N; x++ {
		dbrSelect(100)
	}
}

func BenchmarkDbrSelect1000(b *testing.B) {
	for x := 0; x < b.N; x++ {
		dbrSelect(1000)
	}
}

func BenchmarkDbrSelect10000(b *testing.B) {
	for x := 0; x < b.N; x++ {
		dbrSelect(10000)
	}
}

func dbrSelect(limit uint64) {
	us := []user{}
	_, err := dbrDB.Select("*").From("BrenchmarkTests").Limit(limit).Load(&us)
	if err != nil {
		panic(err)
	}
}
