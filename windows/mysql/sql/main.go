package main

import "fmt"

func main() {
	bytes, err := Asset("02_gaoxinqu_jingkaiqu_inner_%E6%B7%BB%E5%8A%A0%E6%96%B0%E7%BC%96%E5%88%B6%E7%B1%BB%E5%9E%8B.sql") // 根据地址获取对应内容
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
}
