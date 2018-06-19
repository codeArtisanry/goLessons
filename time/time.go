package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println("Hello World!")

	//返回现在时间
	tNow := time.Now()
	//时间转化为string，layout必须为 "2006-01-02 15:04:05"
	timeNow := tNow.Format("01-02 15:04 2006")
	fmt.Println("tNow(time format):", tNow)
	fmt.Println("tNow(string format):", timeNow)

	// //string转化为时间，layout必须为 "2006-01-02 15:04:05"
	// t, _ := time.Parse("2006-01-02 15:04:05", "2014-06-15 08:37:18")
	// fmt.Println("t(time format)", t)

	// //某个时间点 前后判断
	// trueOrFalse := t.After(tNow)
	// if trueOrFalse == true {
	// 	fmt.Println("t（2014-06-15 08:37:18）在tNow之后!")
	// } else {
	// 	fmt.Println("t（2014-06-15 08:37:18）在tNow之前!")
	// }
	// fmt.Println()
}
