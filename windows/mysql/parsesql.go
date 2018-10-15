package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func ParseFile(name string) (s []string) {
	f, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	return Parsesql(f)
}

func Parsesql(f io.Reader) (s []string) {
	r := bufio.NewReaderSize(f, 10*1024*1024)
	// r := bufio.NewReader(f)
	var delimiter = ";"
	var buffer bytes.Buffer
	line, isPrefix, err := r.ReadLine()
	// 移除 UTF8 BOM
	line = bytes.Trim(line, "\xef\xbb\xbf")

	// for err == nil && !isPrefix {
	for err == nil {
		// 移除行尾注释
		// if i := bytes.Index(line, []byte("--")); i > 0 {
		// 	line = line[:i]
		// }
		// 移除行首行尾空白
		line = bytes.TrimSpace(line)
		switch {
		// 1. 被截断继续读取
		case true == isPrefix:
			buffer.Write(line)
		// 2. 忽略注释
		case bytes.HasPrefix(line, []byte("--")):
			if ";" != delimiter {
				buffer.Write(line)
				buffer.WriteByte('\n')
			}
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
			// 找到delimiter结束语句
			buffer.Write(bytes.TrimSuffix(bytes.TrimSuffix(line, []byte(delimiter)), []byte(";")))
			content := buffer.String()
			// ExecDBS(w, &content, dbs)
			s = append(s, content)
			buffer.Reset()
		default:
			buffer.Write(line)
			buffer.WriteByte('\n')
		}

		// 3. read next line
		line, isPrefix, err = r.ReadLine()
	}
	// 5. 保证结尾不含 delimiter 的文本也会正确解析到
	if err == io.EOF {
		content := buffer.String()
		// 解决 Error 1605: Qurey was empty
		if strings.TrimSpace(content) != "" {
			s = append(s, content)
		}
	}

	// if isPrefix {
	// 	fmt.Println("buffer size to small")
	// 	return
	// }
	if err != io.EOF {
		fmt.Println(err)
		return s
	}
	return s
}
