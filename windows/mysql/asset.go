package main

import (
	"fmt"

	"github.com/rinetd/go-learning/windows/mysql/asset"
)

//go:generate go-bindata -ignore=.+\.go  -o=asset/gen.go -pkg=asset asset/...
var Asset_files = []string{
	"asset/20180803_all_inner.sql",
}

func ddl() {

	for _, file := range Asset_files {

		bytes, err := asset.Asset(file) // 根据地址获取对应内容
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(bytes)
	}
}
