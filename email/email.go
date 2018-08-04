package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/jordan-wright/email"
)

//return GoString's buffer slice(enable modify string)
func StringBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// convert b to string without copy
func BytesString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// var c = "日志归类"

// access.log-20180331.gz
var attachment = fmt.Sprint("/var/log/nginx/access.log-", time.Now().Format("20060102"), ".gz")

func main() {
	QQMail()
}

func QQMail() {
	click, ipnum := ParseNginx(attachment)
	fmt.Println("clik:", click, "ip:", ipnum)
	var content = bytes.NewBufferString("本统计基于真实的服务器访问日志.\n")
	content.WriteString("昨日点击量:")
	content.WriteString(strconv.Itoa(click))
	content.WriteString("\n昨日独立IP数:")
	content.WriteString(strconv.Itoa(ipnum))
	e := email.NewEmail()
	e.From = "hangruan <lishoulei@hangruan.cn>"
	e.To = []string{"admin@pytool.com", "xuxy@hangruan.cn", "15863904966@163.com"}
	e.Subject = "一带一路" + time.Now().Format("2006年01月02日") + "访问日志"
	e.Text = StringBytes(content.String())
	// e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	e.AttachFile(attachment)
	e.Send("smtp.exmail.qq.com:587",
		smtp.PlainAuth("", "lishoulei@hangruan.cn", "Passwd", "smtp.exmail.qq.com"))
	// print(time.Date())
}

func ParseNginx(fname string) (click, ipnum int) {
	// f := "access.log-20180721.gz"
	var r *bufio.Reader
	fi, err := os.Open(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], fname,
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
			fmt.Println("Done reading file", err)
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
	// var ipnum, click int
	for _, v := range ips {
		click += v
		ipnum++
	}
	// fmt.Println(click, ipnum)
	return click, ipnum
}
