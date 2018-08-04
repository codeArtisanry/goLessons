package main

import (
	"fmt"
	"testing"
)

func Test_FuncMap(t *testing.T) {
	array := make(map[int]func() int)
	array[func() int { return 10 }()] = func() int {
		return 12
	}
	fmt.Println(array)
}

// func Test_FuncMapErr(t *testing.T) {
// 	array := make(map[func() int]int)  //编译错误
// 	array[func() int { return 12 }] = 10
// 	fmt.Println(array)
// }
