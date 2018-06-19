package main

import "fmt"

func main() {
	s := []int{1}
	s = append(s, 2)
	s = append(s, 3)
	x := append(s, 4)
	y := append(x, 5)
	fmt.Println(s,x,y)
}
