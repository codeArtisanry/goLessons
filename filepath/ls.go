package filepath

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// 利用filepath的Glob方法 读取文件
func LsByGlob() {
	fmt.Println("--------systemc------------")
	files, _ := filepath.Glob("sql/*.sql")
	for _, f := range files {
		fmt.Println(f)
	}
	fmt.Println(files) // contains a list of all files in the current directory
}

// 利用ioutil的ReadDir方法 读取文件

func LsByDir() {
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
