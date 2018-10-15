package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/axgle/mahonia"
)

func main() {
	//生成日志文件
	t := time.Now()
	filepath := "./log_" + strings.Replace(t.String()[:19], ":", "_", 3) + ".txt"
	// file, err := os.OpenFile(filepath, os.O_CREATE, 0666)
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatal("failed!")
	}
	defer file.Close()
	// s := readfile()
	// file.WriteString("aaaa")
	wFile := bufio.NewWriter(file)
	wFile.WriteString(readfile())
	println(wFile.Buffered())
	wFile.Flush()
}

func readfile() string {
	f, err := os.Open("test.txt")
	if err != nil {
		return err.Error()
	}
	defer f.Close()
	buf := make([]byte, 1024)

	decoder := mahonia.NewDecoder("gb18030")
	if decoder == nil {
		fmt.Println("编码不存在!")
		return "编码不存在!"
	}
	var str string = ""
	for {
		n, _ := f.Read(buf)
		if 0 == n {
			break
		}
		//解码为UTF-8
		fmt.Println(string(buf[:n]))
		str += decoder.ConvertString(string(buf[:n]))
		fmt.Println(str)
	}

	return str
}
