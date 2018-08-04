package main

import (
	"fmt"
	"strings"
)

func main() {
	ss := []string{
		"A",
		"B",
		"C",
	}

	var b strings.Builder
	for _, s := range ss {
		fmt.Fprint(&b, s)
	}

	print(b.String())
}
func Grow() {
	var b strings.Builder
	b.Grow(32) //在已知大小的前提下 预先分配大小
	for i, p := range []int{2, 3, 5, 7, 11, 13} {
		fmt.Fprintf(&b, "%d:%d, ", i+1, p)
	}
	s := b.String()   // no copying
	s = s[:b.Len()-2] // no copying (removes trailing ", ")
	fmt.Println(s)
}
