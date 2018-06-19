package main

import (
	"fmt"
	"time"
)

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() (int, int) {
	a, b := 0, 1
	return func() (int, int) {
		a, b = b, a+b
		// time.Sleep(1 * time.Second)
		return a, b
	}
}

func main() {
	f := fib()
	// Function calls are evaluated left-to-right.
	// fmt.Println(f(), f(), f(), f(), f())
	for i := 0; i < 18; i++ {
		a, b := f()
		fmt.Println(i, a, b)
	}
	time.Sleep(10 * time.Second)
}
