package main

import (
	"bufio"
	"bytes"
	"compress/bzip2"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func genericReader(filename string) (io.Reader, *os.File, error) {
	if filename == "" {
		return bufio.NewReader(os.Stdin), nil, nil
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	if strings.HasSuffix(filename, "bz2") {
		return bufio.NewReader(bzip2.NewReader(bufio.NewReader(file))), file, err
	}

	if strings.HasSuffix(filename, "gz") {
		reader, err := gzip.NewReader(bufio.NewReader(file))
		if err != nil {
			return nil, nil, err
		}
		return bufio.NewReader(reader), file, err
	}
	return bufio.NewReader(file), file, err
}

func main() {
	f := "access.log-20180721.gz"
	var r *bufio.Reader
	fi, err := os.Open(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], f,
			err)
		os.Exit(1)
	}
	fz, err := gzip.NewReader(fi)
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}
	var ips = make(map[string]int)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			// os.Exit(0)
			break
		}
		// parser nginx
		ip := strings.Fields(line)
		// fmt.Println(ip[0])

		// 查找键值是否存在
		if _, ok := ips[ip[0]]; ok {
			// fmt.Println(v)
			ips[ip[0]] = ips[ip[0]] + 1
		} else {
			ips[ip[0]] = 1
		}

	}
	var ipnum, click int
	for _, v := range ips {
		click += v
		ipnum++
	}
	fmt.Println(click, ipnum)

	var content = bytes.NewBufferString("本统计基于真实的服务器访问日志,可保证数据真实有效\n")
	content.WriteString("点击量:")
	content.WriteString(strconv.Itoa(click))
	content.WriteString("\n独立IP数")
	content.WriteString(strconv.Itoa(ipnum))
	fmt.Println(content.String())

}
