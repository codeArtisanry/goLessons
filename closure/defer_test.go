package main

import (
	"fmt"
	"testing"
)

func f(i int) func() int {
	return func() int {
		i++
		return i
	}
}

func TestClosure(t *testing.T) {
	m1 := f(2)
	fmt.Println(m1()) // 指针指向 i, i = 2, 输出 3
	fmt.Println(m1()) // 指针指向 i，i = 3，输出 4

	m2 := f(2)
	fmt.Println(m2()) // 指针指向 另外一个 i，i = 2，输出 3
}
