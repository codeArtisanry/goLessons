package bench

import (
	"database/sql"
	"fmt"
	"strconv"
	"testing"

	"gopkg.in/TeaMeow/Reiner.v1"
	// The MySQL driver.
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/gocraft/dbr"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	// "github.com/jinzhu/gorm"
	// "github.com/jmoiron/sqlx"
)

var (
	// dsn = "bench:bench@tcp(118.190.83.129)/bench?charset=utf8"
	dsn = "bench:bench@tcp(localhost)/bench?charset=utf8"
	i   = 0
)

var (
	reinerDB *reiner.Wrapper
	sqlDB    *sql.DB
	dbrDB    *dbr.Session
	sqlxDB   *sqlx.DB
	gormDB   *gorm.DB
	xormDB   *xorm.Engine
)

type user struct {
	ID       int    `xorm:"ID" db:"ID"`
	Username string `xorm:"Username" db:"Username"`
	Password string `xorm:"Password" db:"Password"`
}

func init() {
	var err error
	reinerDB, err = reiner.New(dsn)
	if err != nil {
		panic(err)
	}
	migration := reinerDB.Migration()
	err = migration.Drop("BrenchmarkTests")
	if err != nil {
		panic(err)
	}
	err = migration.
		Table("BrenchmarkTests").
		Column("ID").Int(10).Unsigned().AutoIncrement().Primary().
		Column("Username").Varchar(255).
		Column("Password").Varchar(255).
		Create()
	if err != nil {
		panic(err)
	}
}


func BenchmarkCreate(b *testing.B){
	initsql = `	
	CREATE TABLE `BrenchmarkTests` (
		`ID` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
		`Username` varchar(255) NOT NULL,
		`Password` varchar(255) NOT NULL,
		PRIMARY KEY (`ID`)
	)
	`
start := time.Now()
db, e := sql.Open("mysql", dns)

if _, err := db.Exec(initsql); err != nil {
	// err(err)
	print("exec sql error")
}

}
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
