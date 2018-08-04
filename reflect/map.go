package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	s1 := `{"Name":"test","Age":1}`
	var i interface{}
	json.Unmarshal([]byte(s1), &i)

	var person Person
	convert(i, &person)
	fmt.Println(person.Name, person.Age)
}

func convert(Map interface{}, pointer interface{}) {
	// reflect.Ptr类型 *main.Person
	pointertype := reflect.TypeOf(pointer)
	// reflect.Value类型
	pointervalue := reflect.ValueOf(pointer)
	// struct类型  main.Person
	structType := pointertype.Elem()
	// 将interface{}类型的map转换为  map[string]interface{}
	m := Map.(map[string]interface{})
	// 遍历结构体字段
	for i := 0; i < structType.NumField(); i++ {
		// 获取指定字段的反射值
		f := pointervalue.Elem().Field(i)
		//获取struct的指定字段
		stf := structType.Field(i)
		// 获取tag
		name := strings.Split(stf.Tag.Get("json"), ",")[0]
		// 判断是否为忽略字段
		if name == "-" {
			continue
		}
		// 判断是否为空，若为空则使用字段本身的名称获取value值
		if name == "" {
			name = stf.Name
		}
		//获取value值
		v, ok := m[name]
		if !ok {
			continue
		}

		//获取指定字段的类型
		kind := pointervalue.Elem().Field(i).Kind()
		// 若字段为指针类型
		if kind == reflect.Ptr {
			// 获取对应字段的kind
			kind = f.Type().Elem().Kind()
		}
		// 设置对应字段的值
		switch kind {
		case reflect.Int:
			res, _ := strconv.ParseInt(fmt.Sprint(v), 10, 64)
			pointervalue.Elem().Field(i).SetInt(res)

		case reflect.String:
			pointervalue.Elem().Field(i).SetString(fmt.Sprint(v))
		}
	}
}
