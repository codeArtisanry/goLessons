package main

import (
	"fmt"
	"io"
	"os"
)

var path = "/tmp/test.txt"

func main() {
	createFile()
	writeFile()
	readFile()
	deleteFile()
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
func createFile() {
	// 1. detect if file exists
	var _, err = os.Stat(path)

	// 2. create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("==> done creating file", path)
}

func writeFile() {
	// open file using READ & WRITE permission
	var f, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer f.Close()

	// write some text line-by-line to file
	_, err = f.WriteString("halo\n")
	if isError(err) {
		return
	}
	_, err = f.WriteString("mari belajar golang\n")
	if isError(err) {
		return
	}

	// save changes
	err = f.Sync()
	if isError(err) {
		return
	}

	fmt.Println("==> done writing to file")
}

func readFile() {
	// re-open file
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}

	fmt.Println("==> done reading from file")
	fmt.Println(string(text))
}

func deleteFile() {
	// delete file
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("==> done deleting file")
}
