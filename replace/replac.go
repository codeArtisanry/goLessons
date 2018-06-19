package main

import (
	"regexp"
)

func main() {
	var str0 = `aa,bb\,cc`
	re, _ := regexp.Compile("(?<!a),")
	str1 := re.ReplaceAllString(str0, "n")

	log.Println(str1)
	// 输出 aa,bb\,cc

	// // log.Println(str1)
	// b := "#! /bin/bash\n" + str1

	// print(b)
	// ioutil.WriteFile("b.txt", []byte(b), 755)

	// // fd.WriteString(fmt.Sprint())
	// // return nil
}
