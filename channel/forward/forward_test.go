package main

import (
	"fmt"
	"testing"
)

type bike struct {
	name  string
	转速    int64
	speed float32
}

func increased(source <-chan int) chan int {
	increased := make(chan int)
	go func() {
		for {
			increased <- (<-source + 2)
		}
	}()
	return increased
}

var bikes = bike{
	name: "nihao",
	转速:   25,
}

func TestCh(t *testing.T) {
	fmt.Println(bikes)
	first := make(chan int)
	var result = first
	fmt.Println(result)
	for i := 1; i <= 20; i++ {
		result = increased(result)
		fmt.Println(result)
	}
	fmt.Println("start")
	first <- 1
	fmt.Printf("Result is %d\n", <-result)
}
