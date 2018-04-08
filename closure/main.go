package main

import (
	"fmt"
	"time"
)

// CostTime 采用defer CostTime()()实现无侵入调试
//1. 闭包复制的是原对象指针,这就很容易解释延迟引用现象
func CostTime() func() {
	// x := 100
	// fmt.Printf("2. x (%p) = %d\n", &x, x) //x (0x2101ef018) = 100
	start := time.Now()
	return func() {
		fmt.Println("CostTime:", time.Since(start))
		// fmt.Printf("1. x (%p) = %d\n", &x, x) //x (0x2101ef018) = 100
	}
}

// 2.因为我们有对闭包的支持，我们也能引用一些定义在函数作用域当中的数据
func test2() {
	x := 5
	fn := func() {
		fmt.Println("x =", x)
		x++
	}
	fn() // x=5
	x++
	fn() // x=7
}

const debug = true

func main() {

	if debug == true {
		defer CostTime()()
	}
	// x := closure(10)
	// y := x(1)
	// fmt.Println(y)

	// b := closure1()
	// b() //这里输出11
	// b() //这里输出12
	var a = 10
	p := &a
	*p++
	println(&a)
	println(p, *p)
	closure0()
	test2()
}

func closure0() {
	var data = []int{1, 2, 35, 4, 5, 66}
	var pivot = 10
	var c = 0
	var done chan int = make(chan int)
	go func() {

		for data[c] < pivot {
			c++
		}
		c++
		done <- c
	}()
	end := <-done
	println(end)
	go func() {
		for data[c] < pivot {
			c++
		}
		done <- c
	}()
	end = <-done
	println(end)
}

// 闭包里传递的都是变量的引用而非值的拷贝。
// 可以发现输出的x的地址都是一样的。
// x:0xc420014178
// x:0xc420014178
// 11
func closure(x int) func(i int) int {
	fmt.Printf("x:%p\n", &x)
	return func(y int) int {
		fmt.Printf("x:%p\n", &x)
		return x + y
	}
}

// closure x: 11
// closure x: 12
func closure1() func() {
	var x = 10
	return func() {
		x++
		fmt.Println("closure x:", x)
	}
}

//

func adder() func(int) int {
	sum := 100
	return func(x int) int {
		sum += x
		return sum
	}
}

func addmain() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-1*i))
	}
}
