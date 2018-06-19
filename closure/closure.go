package main

import "fmt"

// 闭包和匿名函数经常一起使用，可以使用闭包来访问函数中的局部变量（被访问操作的变量为指针指向关系，操作的是同一个局部变量）
func closure(x int) (func(), func(int)) {
	fmt.Printf("初始值x为: %d, 内存地址: %p\n", x, &x)
	f1 := func() {
		x = x + 5
		fmt.Printf("x+5: %d, 内存地址: %p\n", x, &x)
	}
	f2 := func(y int) {
		x = x + y
		fmt.Printf("x+%d: %d, 内存地址: %p\n", y, x, &x)
	}
	return f1, f2
}
func main() {
	// 初始值x为: 10, 内存地址: 0xc420014210
	// x+5: 15, 内存地址: 0xc420014210
	// x+10: 25, 内存地址: 0xc420014210
	// x+20: 45, 内存地址: 0xc420014210
	func1, func2 := closure(10)
	func1()
	func2(10)
	func2(20)
}
