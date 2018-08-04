package main

import "fmt"

type Matrix struct {
}

func Inverse(a Matrix) Matrix {
	fmt.Println("a")
	return a
}
func Product(a ...Matrix) Matrix {
	return a[0]
}
func InverseFuture(a Matrix) chan Matrix {
	future := make(chan Matrix)
	go func() {
		future <- Inverse(a)
	}()
	return future
}

func InverseProduct(a Matrix, b Matrix) Matrix {
	a_inv_future := InverseFuture(a) // start as a goroutine
	b_inv_future := InverseFuture(b) // start as a goroutine
	a_inv := <-a_inv_future
	b_inv := <-b_inv_future
	return Product(a_inv, b_inv)
}
