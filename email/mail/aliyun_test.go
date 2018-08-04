package mail

import (
	"fmt"
	"net/smtp"
	"strings"
	"testing"
)

//     email.host=smtp.bitbao.net
// #为email.user
// email.frommail=shigang.xu@bitbao.net
// email.port=25
// email.message=test
// #必须与mail.frommail相同
// email.user=shigang.xu@bitbao.net
// email.passwd=111

// user和from mail用户必须是相同的。
//阿里云企业邮箱
func EmailAli(title string, content *string, toWho string) error {
	toWho += ";"
	host := "smtp.mxhichina.com:25"
	to := strings.Split(toWho, ";")
	content_type := "Content-Type: text/html; charset=UTF-8"
	msg := []byte("To: " + toWho + "\r\nFrom: admin@pytool.com \r\nSubject:" + title + "\r\n" + content_type + "\r\n\r\n" + *content)
	err := smtp.SendMail(
		host,
		smtp.PlainAuth("", "admin@pytool.com", "password", "smtp.mxhichina.com"),
		"admin@pytool.com",
		to,
		[]byte(msg))
	return err
}

func SendEmail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}
func TestSendToMail(t *testing.T) {

	user := "admin@pytool.com"
	password := "password"
	host := "smtp.mxhichina.com:25"
	to := "lishoulei@hangruan.cn"

	subject := "标题"

	body := `<html>
		<body>
		<h3>呵呵</h3>
		使用Golang发送邮件1111使用Golang发送邮件1111使用Golang发送邮件1111使用Golang发送邮件1111使用Golang发送邮件1111
		<img src="https://www.baidu.com/img/baidu_jgylogo3.gif">
		</body>
		</html>`
	fmt.Println("send email")
	err := SendEmail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}
