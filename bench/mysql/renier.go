import (
	"fmt"
	"testing"

	reiner "gopkg.in/TeaMeow/Reiner.v1"
)

var reinerDB *reiner.Wrapper

func BenchmarkReinerInsert(b *testing.B) {
	for x := 0; x < b.N; x++ {
		i++
		var u user
		u.ID = i
		u.Username = fmt.Sprintf("ABCDEFG%d", i)
		u.Password = "HELLO, WORLD!"
		err := reinerDB.Table("BrenchmarkTests").Insert(map[string]interface{}{
			"ID":       u.ID,
			"Username": u.Username,
			"Password": u.Password,
		})
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkReinerSelect1(b *testing.B) {
	for x := 0; x < b.N; x++ {
		reinerSelect(1)
	}
}

func BenchmarkReinerSelect10(b *testing.B) {
	for x := 0; x < b.N; x++ {
		reinerSelect(10)
	}
}

func BenchmarkReinerSelect100(b *testing.B) {
	for x := 0; x < b.N; x++ {
		reinerSelect(100)
	}
}

func BenchmarkReinerSelect1000(b *testing.B) {
	for x := 0; x < b.N; x++ {
		reinerSelect(1000)
	}
}

func BenchmarkReinerSelect10000(b *testing.B) {
	for x := 0; x < b.N; x++ {
		reinerSelect(10000)
	}
}

func reinerSelect(limit int) {
	var us []user
	err := reinerDB.Table("BrenchmarkTests").Bind(&us).Limit(limit).Get()
	if err != nil {
		panic(err)
	}
}
