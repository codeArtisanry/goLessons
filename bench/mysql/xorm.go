package bench

import (
	"fmt"
	"testing"

	"github.com/go-xorm/xorm"
)

var xormDB *xorm.Engine

func BenchmarkXormInsert(b *testing.B) {
	var err error
	xormDB, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic(err)
	}
	for x := 0; x < b.N; x++ {
		i++
		var u user
		u.ID = i
		u.Username = fmt.Sprintf("ABCDEFG%d", i)
		u.Password = "HELLO, WORLD!"
		_, err := xormDB.Table("BrenchmarkTests").Insert(u)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkXormSelect1(b *testing.B) {
	for x := 0; x < b.N; x++ {
		xormSelect(1)
	}
}

func BenchmarkXormSelect10(b *testing.B) {
	for x := 0; x < b.N; x++ {
		xormSelect(10)
	}
}

func BenchmarkXormSelect100(b *testing.B) {
	for x := 0; x < b.N; x++ {
		xormSelect(100)
	}
}

func BenchmarkXormSelect1000(b *testing.B) {
	for x := 0; x < b.N; x++ {
		xormSelect(1000)
	}
}

func BenchmarkXormSelect10000(b *testing.B) {
	for x := 0; x < b.N; x++ {
		xormSelect(10000)
	}
}

func xormSelect(limit int) {
	u := new(user)
	us := []user{}
	rows, err := xormDB.Table("BrenchmarkTests").Limit(limit).Rows(u)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err := rows.Scan(u)
		if err != nil {
			panic(err)
		}
		us = append(us, *u)
	}
}
