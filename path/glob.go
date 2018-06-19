package main

import (
	"fmt"
	"path/filepath"
)

func Main_glob() {
	//Path操作
	a := "nihao hub  "
	fmt.Println("Path操作-----------------")
	fmt.Println(filepath.Glob("/home/*/***/*.md")) //aa.jpg
	for _, v := range a {
		fmt.Println("%v", v)
	}
}
