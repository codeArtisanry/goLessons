package main

import (
	"bufio"
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"os"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

var db *sql.DB

func init() {
	var DSN = `lzkp:yqhtfjzm@tcp(192.168.5.100:3306)/shizhi`
	var err error
	db, err = sql.Open("mysql", DSN)
	if err != nil {
		fmt.Println("\nconnection mysql error")
		return
	}
}
func main() {
	ReadLine("shizhi.sql", DO)
}
func DO(s string) {
	fmt.Println(s)
}
func RunDB(s string) {

	_, err := db.Exec(s)
	if err != nil {
		fmt.Printf("[db.EXEC] %s -> %s : \n", s, err)
	}
	// a, _ := res.RowsAffected()
	// fmt.Println(a)
}
func ReadLine(filename string, do func(s string)) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := bufio.NewReaderSize(f, 10*1024*1024)
	// r := bufio.NewReader(f)
	var delimiter = ";"
	var buffer bytes.Buffer
	line, isPrefix, err := r.ReadLine()
	// for err == nil && !isPrefix {
	for err == nil {
		line = bytes.TrimSpace(line)
		switch {
		// 1. 被截断继续读取
		case true == isPrefix:
			buffer.Write(line)
		// 2. 忽略注释
		case bytes.HasPrefix(line, []byte("--")):
			// buffer.Reset()
		// 3. 重设delimiter
		case bytes.HasPrefix(bytes.ToLower(line), []byte("delimiter")):
			ds := bytes.Fields(line)
			if len(ds) > 1 {
				delimiter = string(ds[1])
			}
			// println()
			// fmt.Println("this line delimiter is", delimiter)
			// buffer.Reset()
		// 4. 根据delimiter处理完整的语句
		case bytes.HasSuffix(line, []byte(delimiter)):
			buffer.Write(bytes.TrimSuffix(bytes.TrimSuffix(line, []byte(delimiter)), []byte(";")))
			do(buffer.String())
			buffer.Reset()
		default:
			buffer.Write(line)
			buffer.WriteByte('\n')
		}

		// 3. read next line
		line, isPrefix, err = r.ReadLine()
	}
	// if isPrefix {
	// 	fmt.Println("buffer size to small")
	// 	return
	// }
	if err != io.EOF {
		fmt.Println(err)
		return
	}

	return
}
