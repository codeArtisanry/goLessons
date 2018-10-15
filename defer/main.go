package main

import (
	"fmt"
)

func closure() func(i int) {
	return func(i int) {
		println("1. closure:", i)
		return
	}
}
func test() {
	x, y := 10, 20
	defer func(i int) {
		println("defer: x=", i, "y=", y) // y 闭包引用
	}(x) // x=10 被复制
	x += 10
	y += 100
	println("x =", x, "y =", y) //defer: 10 120
}
func main() {
	var fs = [4]func(){}
	// data := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < 4; i++ {
		defer closure()(i)
		defer fmt.Println("2. defer i= ", i) //这是一个i作为参数传进去的输出，因为i是int型，所以遵循一个规则值拷贝的传递，还有defer是倒序执行的，所以先后输出3,2,1,0，跟下面的defer交替执行4次

		defer func() {
			fmt.Println("3. defer_func_closure i = ", i) //3. defer_func_closure i =  4 i 取引用一直为4
			// fmt.Println("3. data:", data)
		}()

		defer func(i int) {
			fmt.Println("4. defer_func_closure i = ", i)
		}(i) //执行完下面的代码后，到了该defer了，这也是一个匿名函数，同样的也没有参数，也没有定义i，所以这也是个闭包，用的也是外面的i，所以先输出4，接着执行上面的defer，这样反复执行4次
		{
			var i = i
			defer func(i int) {
				fmt.Println("5. defer_func_closure i = ", i)
			}(i) //执行完下面的代码后，到了该defer了，这也是一个匿名函数，同样的也没有参数，也没有定义i，所以这也是个闭包，用的也是外面的i，所以先输出4，接着执行上面的defer，这样反复执行4次

		}
		fs[i] = func() {
			fmt.Println("6. closure i = ", i)
		} //把相应的4个匿名函数存到function类型的slice里，因为这是个匿名函数，又没有参数，且也没有定义i，所以i就是外层函数的地址引用，就是for循环的i的地址，执行完for后i的值为4，所以输出4个4
		fs[i]() //立即执行
		// 5. closure i =  0
		// 5. closure i =  1
		// 5. closure i =  2
		// 5. closure i =  3
	}
	println("end")
	// test()
}
