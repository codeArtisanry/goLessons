package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// 获取大小的借口
type Sizer interface {
	Size() int64
}

var html = `
<form enctype="multipart/form-data" action="/hello" method="POST">
    Send this file: <input name="uploadfile" type="file" />
    <input type="submit" value="Send File" />
</form>
`

// UploadServer world, the web server
func UploadServer(w http.ResponseWriter, r *http.Request) {
	println("hello")
	if "GET" == r.Method {
		file, _, err := r.FormFile("uploadfile")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer file.Close()
		f, err := os.Create("filenametosaveas")
		defer f.Close()
		io.Copy(f, file)
		fmt.Fprintf(w, "上传文件的大小为: %d", file.(Sizer).Size())
		return
	}
	if "POST" == r.Method {
		// 根据字段名获取表单文件
		formFile, header, err := r.FormFile("uploadfile")
		if err != nil {
			log.Printf("Get form file failed: %s\n", err)
			return
		}
		defer formFile.Close()

		// 创建保存文件
		destFile, err := os.Create("." + r.URL.Path + "/" + header.Filename)
		if err != nil {
			log.Printf("Create failed: %s\n", err)
			return
		}
		defer destFile.Close()

		// 读取表单文件，写入保存文件
		_, err = io.Copy(destFile, formFile)
		if err != nil {
			log.Printf("Write file failed: %s\n", err)
			return
		}
	}
	// 上传页面
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)
	io.WriteString(w, html)
}

func main() {
	http.Handle("/", http.RedirectHandler("/hello", http.StatusTemporaryRedirect))
	http.HandleFunc("/hello", UploadServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
