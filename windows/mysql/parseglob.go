package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

// var dbnames = []string{
// 	"mengshan",
// 	"mengyin",
// 	"pingyi",
// 	"shizhi",
// 	"lanshan",
// 	"luozhuang",
// 	"hedong",
// 	"jingkaiqu",
// 	"gaoxinqu",
// 	"lingang",
// 	"tancheng",
// 	"yinan",
// 	"yishui",
// 	"feixian",
// 	"junan",
// 	"lanling",
// 	"linshu",
// }

var mysqlMap map[string]*sql.DB

func NewMySQL(username, password, host, dbname string, parseTime bool, loc string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=%v&loc=%s",
		username, password, host, dbname, parseTime, loc)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	return db, db.Ping()
}

func InitDB() {

	mysqlMap = make(map[string]*sql.DB)
	Log.Info().Msg("正在初始化")

	var err error
	fmt.Println("1. 正在初始化!", *dbhost)

	for _, dbname := range dbnames {
		mysqlMap[dbname], err = NewMySQL(*dbuser, *dbpassword, *dbhost+":"+*dbport, dbname, true, "Local")
		if err != nil {
			fmt.Println(dbname, err)
			// ioutil.WriteFile("初始化失败", []byte(time.Now().String()), 0755)
			Log.Info().Msg("初始化失败")
		}
	}
	mysqlMap["bi"], err = NewMySQL(*dbuser, *dbpassword, *dbhost+":"+*dbport, "lzkp_bi", true, "Local")
	if err != nil {
		fmt.Println("lzkp_bi", err)
		// ioutil.WriteFile("初始化失败", []byte(time.Now().String()), 0755)
		Log.Info().Msg("lzkp_bi初始化失败")
	}
	Log.Info().Msg("初始化完成")

	fmt.Println("2. 初始化完成!")
}

// func ParseGlob(filename string) (s []string) {
// 	// files, _ := filepath.Glob("*.sql")
// 	files, _ := filepath.Glob(filename)
// 	for _, f := range files {
// 		cc := ParseFile(f)
// 		dbs := strings.Split(f, "_")
// 		for _, fild := range dbs {
// 			if fild == "all" {
// 				dbs = dbnames
// 				break
// 			}
// 		}
// 		for _, db := range dbs {
// 			for _, sdb := range dbnames {
// 				if sdb == db {
// 					s = append(s, sdb)
// 					for _, c := range cc {
// 						// mysqlMap[db].Exec(c)
// 						fmt.Println(c)
// 					}
// 					fmt.Println("Run on db")
// 				}
// 			}
// 		}
// 	}
// 	return s
// }
