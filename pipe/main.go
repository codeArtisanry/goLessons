package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

//bind函数主要是用来为pipe函数整合用的，通过将闭包将函数签名变成pipe所需的样子
//返回一个函数闭包，将一个函数字面量app和字符串slice 传入其中
func bind(app func(in io.Reader, out io.Writer, args []string), args []string) func(in io.Reader, out io.Writer) {
	return func(in io.Reader, out io.Writer) {
		app(in, out, args)
	}
}

//将两个函数插入到管道的中间，调用者只需调用pipe返回的函数字面量，并传入管道的首尾两端，即可实现管道
//返回一个新的函数闭包
func pipe(app1 func(in io.Reader, out io.Writer), app2 func(in io.Reader, out io.Writer)) func(in io.Reader, out io.Writer) {
	return func(in io.Reader, out io.Writer) {
		pr, pw := io.Pipe()
		defer pw.Close()
		go func() {
			defer pr.Close()
			app2(pr, out)
		}()
		app1(in, pw)
	}
}

//读取args slice的每个字符串，将其作为文件名，读取文件,并在文件的每一行首部加上行号，写入到out中
//此处in没有使用到，主要是为了保证管道定义的一致性
func app1(in io.Reader, out io.Writer, args []string) {
	for _, v := range args {
		//fmt.Println(v)
		file, err := os.Open(v)
		if err != nil {
			continue
		}
		defer file.Close()
		buf := bufio.NewReader(file)
		for i := 1; ; i++ {
			line, err := buf.ReadBytes('\n')
			if err != nil {
				break
			}
			linenum := strconv.Itoa(i)
			nline := []byte(linenum + " ")
			nline = append(nline, line...)
			fmt.Print(nline)
			out.Write(nline)
		}
	}
}

//app2 主要是将字节流转化为大写,中文可能会有点问题，不过主要是演示用，重在理解思想
//read from in, convert byte to Upper ,write the result to out
func toUpper(in io.Reader, out io.Writer) {
	rd := bufio.NewReader(in)
	p := make([]byte, 10)
	for {
		n, _ := rd.Read(p)
		if n == 0 {
			break
		}
		t := bytes.ToUpper(p[:n])
		out.Write(t)
	}
}

func main() {
	args := os.Args[1:]
	for _, v := range args {
		fmt.Println(v)
	}
	p := pipe(bind(app1, args), toUpper)
	p(os.Stdin, os.Stdout)
}
