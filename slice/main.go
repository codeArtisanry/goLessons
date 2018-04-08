package main

import "fmt"

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
	var Data = make([]int, 5, 6)

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
