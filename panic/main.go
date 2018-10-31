package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

var ERR_CREATE_RESOURCE1_FAILED = errors.New("err")

func createResource1() error {
	return ERR_CREATE_RESOURCE1_FAILED
}
func destroyResource1() {
	fmt.Println("destroyResource1")
}
func deferDemo() error {
	defer func() {
		print("deff")
	}()
	err := createResource1()
	if err != nil {
		return ERR_CREATE_RESOURCE1_FAILED
	}
	defer func() {
		if err != nil {
			destroyResource1()
		}
	}()
	return nil
}

func funcA() error {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("panic recover! p: %v", p)
			debug.PrintStack()
		}
	}()
	return funcB()
}

func funcB() error {
	// simulation
	panic("foo")
	return errors.New("success")
}

func test() {
	err := funcA()
	if err == nil {
		fmt.Printf("err is nil\\n")
	} else {
		fmt.Printf("err is %v\\n", err)
	}
}
func main() {
	deferDemo()
	test()
}
