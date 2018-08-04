package main

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
	fmt.Println(v)
}

type student struct {
	Name string
	Age  int
}

func TestRangeFalse(t *testing.T) {
	var stus []student
	stus = []student{
		{Name: "one", Age: 18},
		{Name: "two", Age: 19},
	}
	data := make(map[int]student)
	for i, v := range stus {
		//这里每次都是取到临时变量
		data[i] = v //应该改为：data[i] = &stus[i]
	}
	for i, v := range data {
		fmt.Printf("key=%d, value=%v \n", i, v)
	}
}
func TestRangeTrue(t *testing.T) {
	var stus []student
	stus = []student{
		{Name: "one", Age: 18},
		{Name: "two", Age: 19},
	}
	data := make(map[int]*student)
	for i, v := range stus {
		e := v
		data[i] = &e //应该改为：data[i] = &stus[i]
	}
	for i, v := range data {
		fmt.Printf("key=%d, value=%v \n", i, v)
	}
}
