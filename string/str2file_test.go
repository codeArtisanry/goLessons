package string

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	// _ "github.com/joho/godotenv/autoload"
)

func TestStrtofile(t *testing.T) {

	var str = []string{"======",
		"fd_time",
		"=====",
		"s",
		"\n"}
	// fd, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	str1 := strings.Join(str, "\n")
	// log.Println(str1)
	b := "#! /bin/bash\n" + str1

	// print(b)
	// fmt.Println(str)
	ioutil.WriteFile("b.txt", []byte(b), 0755)
	os.Remove("b.txt")
	fmt.Println(os.Stat("/usr/bin/zip"))
	// fd.WriteString(fmt.Sprint())
	// return nil
}
