package main

import "fmt"

type person struct {
	name string
}

func main() {
	var db map[string]person
	dbm := make(map[string]person, 10)
	dbm1 := map[string]person{}
	fmt.Printf("%p %+T\n", &db, db)
	fmt.Printf("%p %p", &dbm, &dbm1)

}
