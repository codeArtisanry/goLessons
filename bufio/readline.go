package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadLine(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()
	for err == nil && !isPrefix {
		s := string(line)
		fmt.Println(s)
		line, isPrefix, err = r.ReadLine()
	}
	if isPrefix {
		fmt.Println("buffer size to small")
		return
	}
	if err != io.EOF {
		fmt.Println(err)
		return
	}
}

func ReadString(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := bufio.NewReader(f)
	line, err := r.ReadString('\n')
	for err == nil {
		fmt.Print(line)
		line, err = r.ReadString('\n')
	}
	if err != io.EOF {
		fmt.Println(err)
		return
	}
}

func main() {
	filename := `testfile`
	ReadLine(filename)
	ReadString(filename)
}
