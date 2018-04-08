package main

import "fmt"

const ptrSize = 4 << (^uintptr(0) >> 63) // unsafe.Sizeof(uintptr(0)) but an ideal const

func main() {
	uintptrNone := ^uintptr(0)
	ptr := ^uintptr(0) >> 63
	fmt.Printf("%64b\n", uintptrNone)
	fmt.Printf("%08b\n", uintptrNone)
	fmt.Printf("%b\n", uintptrNone)
	fmt.Println(ptr)
	fmt.Println(ptrSize)
	fmt.Println(4 << 0)
}
