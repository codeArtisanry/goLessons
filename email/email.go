package main

import (
	"fmt"
	"net/smtp"
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

var c = "日志归类"

// access.log-20180331.gz
var file = fmt.Sprint("/var/log/nginx/access.log-", time.Now().Format("20060102"), ".gz")

func main() {
	e := email.NewEmail()
	e.From = "hangruan <lishoulei@hangruan.cn>"
	e.To = []string{"rinetd@163.com", "15863904966@163.com"}
	e.Subject = "一带一路访问日志"
	e.Text = StringBytes(c)
	// e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	e.AttachFile(file)
	e.Send("smtp.exmail.qq.com:587", smtp.PlainAuth("", "lishoulei@hangruan.cn", "Sdlylshl871016", "smtp.exmail.qq.com"))
	// print(time.Date())
}
