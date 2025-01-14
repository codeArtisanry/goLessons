package main

import "fmt"
import "github.com/axgle/mahonia"

func ConvertToString(src string, srcCode string, tagCode string) string {
    srcCoder := mahonia.NewDecoder(srcCode)
    srcResult := srcCoder.ConvertString(src)
    tagCoder := mahonia.NewDecoder(tagCode)
    _, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
    result := string(cdata)
    return result
}
func main() {

    str := "乱码的字符串变量"
    str = ConvertToString(str, "gbk", "utf-8")
    fmt.Println(str)

}
