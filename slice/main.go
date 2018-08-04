package main

import (
	"fmt"
)

// type slice struct {
//     array unsafe.Pointer
//     len   int
//     cap   int
// }

// 这里函数必须返回[]int,因为append函数超过slice容量时将生成新的数组对象，不会影响原来slice
func dataappend(data []int) []int {
	fmt.Printf("2.参数 slice地址: %p \t参数 array地址: %p\n", &data, data)
	return append(data, data...)
}
func main() {
	var x [1024]int

	s := make([]int, 1024)
	fmt.Printf("addr %p \n", &x)
	fmt.Printf("addr %p \n", &s)

	// slice 初始化的6种方式
	// 空切片和 nil 切片的区别在于，空切片指向的地址不是nil 0
	var slice1 []*int         //nil切片 addr: 0 len: 0 cap:0          地址指针:固定
	slice2 := []*int{}        //空切片 addr: 0x547fa8 len: 0 cap:0    地址指针:固定
	slice3 := make([]*int, 0) //空切片 addr: 0x547fa8 len: 0 cap:0    地址指针:固定
	pslice4 := new([](*int))  //切片 addr: 0xc42000a080 len: 0 cap:0

	slice5 := make([]int, 1) //切片 addr: 0xc420014210 len: 1 cap:1
	slice6 := []int{1}       //切片 addr: 0xc420014218 len: 1 cap:1

	//
	fmt.Printf("slice1:%p %v %v\n", slice1, len(slice1), cap(slice1))      //slice1:0x0 0 0
	fmt.Printf("slice2:%p %v %v\n", slice2, len(slice2), cap(slice2))      //slice2:0x547fa8 0 0
	fmt.Printf("slice3:%p %v %v\n", slice3, len(slice3), cap(slice3))      //slice3:0x547fa8 0 0
	fmt.Printf("slice4:%p %v %v\n", pslice4, len(*pslice4), cap(*pslice4)) //slice4:0xc42009e080 0 0
	fmt.Printf("slice5:%p %v %v\n", slice5, len(slice5), cap(slice5))      //slice5:0xc420092010 1 1
	fmt.Printf("slice6:%p %v %v\n", slice6, len(slice6), cap(slice6))      //slice6:0xc420092018 1 1

	fmt.Printf("slice1:%p %p \n", &slice1, slice1)   //slice1:0xc42000a080 0x0
	fmt.Printf("slice2:%p %p \n", &slice2, slice2)   //slice2:0xc42000a0a0 0x547fa8
	fmt.Printf("slice3:%p %p \n", &slice3, slice3)   //slice3:0xc42000a0c0 0x547fa8
	fmt.Printf("slice4:%p %p \n", pslice4, *pslice4) //slice4:0xc42000a0e0 0x0
	fmt.Printf("slice5:%p %p \n", &slice5, slice5)   //slice5:0xc42000a100 0xc420014210
	fmt.Printf("slice6:%p %p \n", &slice6, slice6)   //slice6:0xc42000a120 0xc420014218

	var Data = make([]int, 5, 6)
	// slice中的坑之
	fmt.Printf("1.全局 slice地址: %p \t全局 array地址: %p\n", &Data, Data)
	d := dataappend(Data)
	fmt.Printf("3.返回 slice地址: %p \t返回 array地址: %p\n\n", &d, d)

	Data = make([]int, 5, 10)
	fmt.Printf("1.全局 slice地址: %p \t全局 array地址: %p\n", &Data, Data)
	d = dataappend(Data)
	fmt.Printf("3.返回 slice地址: %p \t返回 array地址: %p\n\n", &d, d)

	fmt.Println("结论1:参数传递的是silce的拷贝,数据是引用地址")
	fmt.Println("结论2:局部变量对slice的增加超过cap(slice)时，必须返回新slice")
	fmt.Println("结论3:对副本的修改不影响原slice")
}
