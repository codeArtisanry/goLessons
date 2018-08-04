package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

// Load for loading data for testing
func MysqlLoad(file, dbname string) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Errorf("Fail to connect database. %s", err.Error())
	}
	defer db.Close()
	db.Ping()

	_, err = db.Exec("USE " + dbname)
	if err != nil {
		fmt.Printf("[db.EXEC] switch databases %s -> %s : \n", dbname, err.Error())
		return
	}

	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Errorf("readfile error")
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[EXEC RECOVER] error ", err)
		}
	}()

	queries := strings.Split(string(content), ";")
	// fmt.Println(queries)
	for _, query := range queries {
		if strings.TrimSpace(query) != "" {

			// fmt.Println(query)
			_, err := db.Exec(query)
			if err != nil {
				fmt.Errorf("[db.EXEC] %s -> %s : ", query, err.Error())
			}
			// a, _ := res.RowsAffected()
			// fmt.Println(a)
		}
	}
	fmt.Println("import ", file, " finish !")
}

var names = map[string]bool{
	"mengshan":  true,
	"mengyin":   true,
	"pingyi":    true,
	"shizhi":    true,
	"lanshan":   true,
	"luozhuang": true,
	"hedong":    true,
	"jingkaiqu": true,
	"gaoxinqu":  true,
	"lingang":   true,
	"tancheng":  true,
	"yinan":     true,
	"yishui":    true,
	"feixian":   true,
	"junan":     true,
	"lanling":   true,
	"linshu":    true,
}

func DBExecAll() {
	files, _ := filepath.Glob("*.sql")
	for _, f := range files {
		db := strings.Split(f, "_")
		if names[db[0]] {
			fmt.Println(f)
			MysqlLoad(f, db[0])
		} else {
			for k, _ := range names {
				fmt.Println(k)
				MysqlLoad(f, k)
			}
		}
	}
}
