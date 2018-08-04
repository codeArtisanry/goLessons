package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		println(e)
	}
}
func write() {
	// 1. ioutil
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)
	// 2. os
	f, err := os.Create("/tmp/dat2")
	check(err)
	defer f.Close()
	d2 := []byte{115, 111, 109, 101, 10} //some
	f.Write(d2)
	f.WriteString("writes\n")

	f.Sync()

	// 3. buffer
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush()
}
