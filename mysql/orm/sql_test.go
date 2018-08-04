package bench

// go test -bench=. -benchmem
import (
	"database/sql"
	"fmt"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var sqlDB *sql.DB

func BenchmarkSQLInsert(b *testing.B) {
	var err error
	sqlDB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	for x := 0; x < b.N; x++ {
		i++
		var u user
		u.ID = i
		u.Username = fmt.Sprintf("ABCDEFG%d", i)
		u.Password = "HELLO, WORLD!"
		stmt, err := sqlDB.Prepare("INSERT INTO BrenchmarkTests (ID, Username, Password) VALUES (?, ?, ?)")
		if err != nil {
			panic(err)
		}
		_, err = stmt.Exec(u.ID, u.Username, u.Password)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkSQLSelect1(b *testing.B) {
	limit := strconv.Itoa(1)
	for x := 0; x < b.N; x++ {
		sqlSelect(limit)
	}
}

func BenchmarkSQLSelect10(b *testing.B) {
	limit := strconv.Itoa(10)
	for x := 0; x < b.N; x++ {
		sqlSelect(limit)
	}
}

func BenchmarkSQLSelect100(b *testing.B) {
	limit := strconv.Itoa(100)
	for x := 0; x < b.N; x++ {
		sqlSelect(limit)
	}
}

func BenchmarkSQLSelect1000(b *testing.B) {
	limit := strconv.Itoa(1000)
	for x := 0; x < b.N; x++ {
		sqlSelect(limit)
	}
}

func BenchmarkSQLSelect10000(b *testing.B) {
	limit := strconv.Itoa(10000)
	for x := 0; x < b.N; x++ {
		sqlSelect(limit)
	}
}
func sqlSelect(limit string) {
	var us []user
	rows, err := sqlDB.Query("SELECT * FROM BrenchmarkTests LIMIT " + limit)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var u user
		rows.Scan(&u.ID, &u.Username, &u.Password)
		us = append(us, u)
	}
}
