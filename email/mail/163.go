package mail

// const (
// 	SMTPHOST = "smtp.163.com"
// 	SMTPPORT = "25"
// 	USERNAME = "pansifan0525@163.com"
// 	PASSWORD = "psmortal0525"
// 	IDENTITY = ""
// )

// func Send163(addr, msg string) error {
// 	auth := smtp.PlainAuth(IDENTITY, USERNAME, PASSWORD, SMTPHOST)
// 	var m = []byte("To: " + addr + "\r\n" +
// 		"From: " + "pansifan0525@163.com" + "\r\n" +
// 		"Subject: " + "账号激活" + "\r\n" +
// 		"Content-Type: text/html; charset=UTF-8" + "\r\n" +
// 		"\r\n" + msg +
// 		"\r\n")
// 	to := []string{"xupengzysq@163.com"}

// 	err := smtp.SendMail(SMTPHOST+":25", auth, "pansifan0525@163.com", []string{addr}, []byte(m))

// 	if err != nil {
// 		fmt.Println("邮件发送失败!", err.Error())
// 		return err
// 	}
// 	fmt.Println("successfully")
// 	return nil
// }
