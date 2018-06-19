package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

const filename = "test.txt"

func checkFileIsExist(filename string) bool {
	var exist bool = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func main() {
	var f *os.File

	// //1.使用 io.WriteString 写入文件
	// if checkFileIsExist(filename) {  //如果文件存在
	//     f, _ = os.OpenFile(filename, os.O_APPEND, 0666)  //打开文件
	// } else {
	//     f, _ = os.Create(filename)  //创建文件
	// }
	// io.WriteString(f, "testststst") //写入文件(字符串)

	//2.使用 ioutil.WriteFile 写入文件
	var d = []byte("testststst1\n" + "sdf")
	var buf = bytes.Buffer{}
	buf.WriteString("d")
	buf.Bytes()
	// strings.Join()
	err := ioutil.WriteFile(filename, d, 0666) //写入文件(字节数组)
	if err != nil {
		fmt.Println(err)
	}

	//3.使用 File(Write,WriteString) 写入文件
	f.WriteString("writesn") //写入文件(字节数组)
	f.Sync()

	//4.使用 bufio.NewWriter 写入文件
	w := bufio.NewWriter(f) //创建新的 Writer 对象
	w.WriteString("bufferedn")
	w.Flush()
	f.Close()
}
