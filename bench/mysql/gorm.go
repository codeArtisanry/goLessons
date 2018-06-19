package bench

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
)

var gormDB *gorm.DB

func BenchmarkGormInsert(b *testing.B) {
	var err error
	gormDB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	for x := 0; x < b.N; x++ {
		i++
		var u user
		u.ID = i
		u.Username = fmt.Sprintf("ABCDEFG%d", i)
		u.Password = "HELLO, WORLD!"
		gormDB.Table("BrenchmarkTests").Create(u)
		if gormDB.Error != nil {
			panic(err)
		}
	}
}

func BenchmarkGormSelect1(b *testing.B) {
	for x := 0; x < b.N; x++ {
		gormSelect(1)
	}
}

func BenchmarkGormSelect10(b *testing.B) {
	for x := 0; x < b.N; x++ {
		gormSelect(10)
	}
}

func BenchmarkGormSelect100(b *testing.B) {
	for x := 0; x < b.N; x++ {
		gormSelect(100)
	}
}

func BenchmarkGormSelect1000(b *testing.B) {
	for x := 0; x < b.N; x++ {
		gormSelect(1000)
	}
}

func BenchmarkGormSelect10000(b *testing.B) {
	for x := 0; x < b.N; x++ {
		gormSelect(10000)
	}
}

func gormSelect(limit int) {
	var us []user
	gormDB.Table("BrenchmarkTests").Limit(limit).Find(&us)
	if gormDB.Error != nil {
		panic(gormDB.Error)
	}
}
