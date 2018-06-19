package bench

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/jmoiron/sqlx"
)

var sqlxDB *sqlx.DB

func BenchmarkSQLxInsert(b *testing.B) {
	var err error
	sqlxDB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	for x := 0; x < b.N; x++ {
		i++
		var u user
		u.ID = i
		u.Username = fmt.Sprintf("ABCDEFG%d", i)
		u.Password = "HELLO, WORLD!"
		stmt, err := sqlxDB.Prepare("INSERT INTO BrenchmarkTests (ID, Username, Password) VALUES (?, ?, ?)")
		if err != nil {
			panic(err)
		}
		_, err = stmt.Exec(u.ID, u.Username, u.Password)
		if err != nil {
			panic(err)
		}
	}
}
func BenchmarkSQLxSelect1(b *testing.B) {
	limit := strconv.Itoa(1)
	for x := 0; x < b.N; x++ {
		sqlxSelect(limit)
	}
}

func BenchmarkSQLxSelect10(b *testing.B) {
	limit := strconv.Itoa(10)
	for x := 0; x < b.N; x++ {
		sqlxSelect(limit)
	}
}

func BenchmarkSQLxSelect100(b *testing.B) {
	limit := strconv.Itoa(100)
	for x := 0; x < b.N; x++ {
		sqlxSelect(limit)
	}
}

func BenchmarkSQLxSelect1000(b *testing.B) {
	limit := strconv.Itoa(1000)
	for x := 0; x < b.N; x++ {
		sqlxSelect(limit)
	}
}

func BenchmarkSQLxSelect10000(b *testing.B) {
	limit := strconv.Itoa(1000)
	for x := 0; x < b.N; x++ {
		sqlxSelect(limit)
	}
}

func sqlxSelect(limit string) {
	us := []user{}
	err := sqlxDB.Select(&us, "SELECT * FROM BrenchmarkTests LIMIT "+limit)
	if err != nil {
		panic(err)
	}
}
