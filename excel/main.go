package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/xuri/excelize"
)

const (
	upload_path string = "./upload/"
)

func main() {
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/uploads", uploadHandle)
	err := http.ListenAndServe(":8780", nil)
	if err != nil {
		fmt.Println("服务器启动失败")
		return
	}
	fmt.Println("服务器启动成功")
}
func indexHandle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<html><head><title>xlsx文件导入</title></head><body><form action='/uploads' method=\"post\" enctype=\"multipart/form-data\"><label>上传EXCEL</label><input type=\"file\" name='file'  /><br/><label><input type=\"submit\" value=\"上传EXCEL\"/></label></form></body></html>")
}
func uploadHandle(w http.ResponseWriter, r *http.Request) {
	//获取文件内容 要这样获取
	file, head, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	//当期时间格式化
	filename := time.Now().Format("20060102150405")
	//获取文件的后缀
	fileSuffix := path.Ext(head.Filename)

	filePath := upload_path + filename + fileSuffix
	//创建文件
	fW, err := os.Create(filePath)
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	defer fW.Close()
	_, err = io.Copy(fW, file)
	if err != nil {
		fmt.Println("文件保存失败")
		return
	}
	if fileSuffix == ".xlsx" {
		fileXlsx(filePath)
	}
	//跳转到首页
	http.Redirect(w, r, "/", http.StatusFound)

}
func fileXlsx(filePath string) {
	xlsx, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rows := xlsx.GetRows("sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}

	}

}
