package main

import (
	"database/sql"
	"fmt"
)

//reconnect tries to reconnect untill sucess
func reconnect(db *sql.DB, constr string) {
	go func() {
		for {
			//fmt.Println("....checking connection")
			if err := db.Ping(); err != nil {
				//fmt.Println("connection lost, reconnecting...")
				initDB(constr)
				fmt.Println("reconnected")
			}
		}
	}()
}
