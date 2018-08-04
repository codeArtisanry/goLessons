package main

import (
	"encoding/json"
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

type Param map[string]interface{}

type Show struct {
	Param
}

func TestMap(t *testing.T) {
	s := new(Show)
	s.Param["RMB"] = 10000 //map 没有初始化
}

type student struct {
	Name string
}

func zhoujielun(v interface{}) {
	switch msg := v.(type) {
	case *student, student:
		// fmt.Println(msg.Name)
		fmt.Println(msg)
	}
}
func TestInterface(t *testing.T) {
	var msg student
	msg.Name = "liming"
	zhoujielun(msg)
}

type People struct {
	name string `json:"name"`
}

func TestJson(t *testing.T) {

	js := `{
		"name":"11"
	}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p) //未导出字段输出为空 people:  {}
}

type People1 struct {
	Name string
}

func (p *People1) String() string {
	// return fmt.Sprintf("print: %v", p)
	return "str"
}

func TestMethod(t *testing.T) {
	p := People1{}
	// var p = new(People1)
	p.String() // error Sprintf format %v with arg p causes recursive String method call
}

// panic: send on closed channel
func TestChan(t *testing.T) {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i //panic: send on closed channel
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}

var value int32

func SetValue(delta int32) {
	for {
		v := value
		if atomic.CompareAndSwapInt32(&value, v, (v + delta)) {
			break
		}
	}
}
func TestAtomic(t *testing.T) {
	SetValue(32)
}

func TestChan2(t *testing.T) {

	abc := make(chan int, 1000)
	for i := 0; i < 10; i++ {
		abc <- i
	}
	go func() {
		for {
			a := <-abc
			fmt.Println("a: ", a)
		}
	}()
	close(abc)
	fmt.Println("close")
	time.Sleep(time.Second * 100)
}

type Student struct {
	name string
}

// cannot assign to struct field m["people"].name in map
func TestMap2(t *testing.T) {
	// m := map[string]Student{"people": {"zhoujielun"}}
	// m["people"].name = "wuyanzu"
}

type query func(string) string

func exec(name string, vs ...query) string {
	ch := make(chan string)
	fn := func(i int) {
		ch <- vs[i](name)
	}
	for i, _ := range vs {
		go fn(i)
	}
	return <-ch
}

func TestExe(t *testing.T) {

	ret := exec("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	})
	fmt.Println(ret)
}
