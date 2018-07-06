package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

// var (
// 	// dbHost     = "15.14.12.152"
// 	dbHost     = "192.168.5.100"
// 	dbPort     = "3306"
// 	dbUser     = "root"
// 	dbPassword = "toor"
// 	dbName     = "shizhi"
// )

var db *sql.DB

func main() {

	var src = flag.String("s", "shizhi.sql", "需要导入的sql")
	var dbHost = flag.String("h", "192.168.5.100", "数据库HOST")
	var dbName = flag.String("n", "shizhi_test", "数据库名")
	var dbPassword = flag.String("p", "toor", "数据库密码")
	var dbUser = flag.String("u", "root", "数据库用户")
	var dbPort = flag.String("P", "3306", "数据库端口")

	flag.Parse()
	fmt.Printf("program's name:%s  -s=%s \n", os.Args[0], *src)
	var conf = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", *dbUser, *dbPassword, *dbHost, *dbPort, *dbName)

	fmt.Println(conf)
	var err error
	db, err = sql.Open("mysql", conf)

	if err != nil {
		fmt.Errorf("Fail to connect database. %s", err.Error())
	}
	// defer db.Close()
	db.Ping()
	db.Stats()
	// db.Exec("Create DATABASE if not exist" + *dbName)

	Load(db, *src)
}

// Load for loading data for testing
func Load(db *sql.DB, fixtureName string) {
	content, err := ioutil.ReadFile(fixtureName)
	if err != nil {
		fmt.Errorf("readfile error")
	}
	queries := strings.Split(string(content), ";")
	fmt.Println(queries)
	for _, query := range queries {
		if strings.TrimSpace(query) != "" {
			_, err := db.Exec(query)
			if err != nil {
				fmt.Errorf("readfile error")
			}
		}
	}
}
