package main

import (
	"fmt"
	"path"
	"strings"
)

func GetDirName(uri string) string {

	es := strings.Split(uri, "/")
	path := es[:len(es)-1]
	return strings.Join(path, "/")
}

func main() {
	//Path操作
	// fmt.Println("Path操作-----------------")
	// fmt.Println(path.Base("http://www.baidu.com/file/aa.jpg")) //aa.jpg
	// fmt.Println(path.Clean("c:\\file//abc///aa.jpg"))          //c:\file/abc/aa.jpg
	// fmt.Println(os.Getwd())                                    //D:\Projects\GoPath\source\demo\syntax\path <nil>
	// fmt.Println(path.Dir("http://www.baidu.com/aa/aaa.jpg"))   //http:/www.baidu.com/aa
	fmt.Println(path.Dir("c:/a/b/c"))       //c:/a/b/c
	fmt.Println(path.Dir("c:/a/b/c/"))      //c:/a/b/c
	fmt.Println(path.Dir("c:/a/b/c/d.txt")) //c:/a/b/c
	// fmt.Println(path.Dir("c:\\a/b.txt"))                       //c:\a
	// fmt.Println(path.Ext("c:\\a/b.txt"))                       //.txt
	// fmt.Println(path.IsAbs("c:/wind/aa/bb/b.txt"))             //false
	// fmt.Println(path.Join("c:", "aa", "bb", "cc.txt"))         //c:/aa/bb/cc.txt
	// isMatch, err := path.Match("c:/windows/*/", "c:/windows/system/")
	// fmt.Println(isMatch, err) //true <nil>
	// fmt.Println(path.Join("/s/d/de/", "file.txt"))

}
