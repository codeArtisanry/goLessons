package main

import (
	"fmt"
	"net/http"
	"time"
)

var content = `
<!DOCTYPE html>
<html>
<head>
	<meta http-equiv="content-type" content="text/html; charset=utf-8" />
	<title>Multipart Test</title>
</head>
<body>
<div style="margin-left:auto;margin-right:auto;max-width:670px">
<hr><div style="float:right;margin-right:200px;">
<form action="/backup" method="POST">
<input type="submit" value="数据备份"/>  
</form>
</div>  
<p style="clear:both;">
<p>使用说明:</p>
<p>1. 本系统具有较大的破坏性,非专业人士,请勿操作,除非明确知道自己要干什么!!!</p>
<p>2. 因业务需要,只有本系统执行时会忽略掉错误的语句,继续往下执行.</p>

<form  action="/" enctype="multipart/form-data" method="post">
<fileset>
	<input type="checkbox" name="dbnames" value="shizhi" > 市值
	<input type="checkbox" name="dbnames" value="lanshan" checked> 兰山
	<input type="checkbox" name="dbnames" value="luozhang" checked> 罗庄
	<input type="checkbox" name="dbnames" value="hedong" checked> 河东
	<input type="checkbox" name="dbnames" value="gaoxinqu" checked> 高新区
	<input type="checkbox" name="dbnames" value="jingkaiqu" checked> 经开区
	<input type="checkbox" name="dbnames" value="lingang" checked> 临港
	<input type="checkbox" name="dbnames" value="yinan" checked> 沂南
	<input type="checkbox" name="dbnames" value="tancheng" checked> 郯城
	<input type="checkbox" name="dbnames" value="yishui" checked> 沂水
	<input type="checkbox" name="dbnames" value="lanling" checked> 兰陵
	<input type="checkbox" name="dbnames" value="feixian" checked> 费县
	<input type="checkbox" name="dbnames" value="pingyi" checked> 平邑
	<input type="checkbox" name="dbnames" value="junan" checked> 莒南
	<input type="checkbox" name="dbnames" value="mengyin" checked> 蒙阴
	<input type="file" name="file" >
	<input type="file" name="files" multiple />
	<input type="file" name="dirs" webkitdirectory/>
	<textarea style="height: 150px; width: 600px;" name="content"  placeholder="">
	</textarea>
	
	<input type="submit" />
</fileset>
</form>
</body>
</html>
`

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(content))
	if http.MethodPost == r.Method {

		r.FormFile("file")
		if err := r.ParseMultipartForm(100000000000); nil != err {
			fmt.Println("Parse MultipartForm failed")
			return
		}

		for name, fileHeaders := range r.MultipartForm.File {
			fmt.Println(name)
			for _, fileHeader := range fileHeaders {
				fmt.Println(fileHeader)
			}

		}

		// for name, fileHeaders := range r.MultipartForm.File["file[]"] {
		// 	fmt.Println(name, fileHeaders)

		// }
		// for name, fileHeaders := range r.MultipartForm.File["files"] {
		// 	fmt.Println(name, fileHeaders)
		// }

	}

}

func main() {

	http.HandleFunc("/", handler)
	s := &http.Server{
		Addr:           ":12345",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe() // listen and serve on 0.0.0.0:8080
	// go http.ListenAndServe(":12345", nil)
	// http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
	for {
	}
}
