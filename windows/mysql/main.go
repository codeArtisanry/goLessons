package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	fmt.Println("3. 正在执行!")
	ParseGlob("sql/*.sql")
	fmt.Println("4. 执行成功!")
	ioutil.WriteFile("执行成功", []byte(time.Now().String()), 0755)
}
