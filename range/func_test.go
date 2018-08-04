package main

import (
	"fmt"
	"testing"
)

func Test_func(t *testing.T) {
	var msgs []func()
	array := []string{
		"1", "2", "3", "4",
	}
	for _, e := range array {
		msgs = append(msgs, func() {
			fmt.Println(e) //e作为临时变量,每次循环都被复用
		})
	}
	for _, v := range msgs {
		v()
	}
}

func Test_func2(t *testing.T) {
	var msgs []func()
	array := []string{
		"1", "2", "3", "4",
	}
	for _, e := range array {
		elem := e //每次循环都,
		msgs = append(msgs, func() {
			fmt.Println(elem) //e作为临时变量,每次循环都被复用
		})
	}
	for _, v := range msgs {
		v()
	}
}
