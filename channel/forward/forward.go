package main

import "fmt"

func increased(source <-chan int) chan int {
	increased := make(chan int)
	go func() {
		for {
			increased <- (<-source + 2)
		}
	}()
	return increased
}

func main() {
	first := make(chan int)
	var result = first
	fmt.Println(result)
	for i := 1; i <= 20; i++ {
		result = increased(result)
		// fmt.Println(result)
	}
	fmt.Println("start")
	first <- 1
	fmt.Printf("Result is %d", <-result)
}
