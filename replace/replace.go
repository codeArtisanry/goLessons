package main

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"
)

const (
	stated  = false
	escaped = true
	escape  = "\\"
)

func Replace(input, delimiter, new string) string {
	parsed := ""
	state := stated
	for _, c := range input {
		c := string(c)
		if state == stated {
			if c == delimiter {
				parsed += new
			} else if c == escape {
				state = escaped
			} else {
				parsed += c
			}
		} else {
			if c == delimiter {

				parsed += c

			} else {
				parsed += escape
				parsed += c
			}
			state = stated
		}
	}

	return parsed
}

func main() {

	s := "adfb"
	r := []rune(s)
	r = append(r, 46)

	fmt.Println(r[len(r)-1], len(r), string(r))

	var str = `asdfaf,qwerq\\\,wef,a你'好sd\\,我asdfa,,123`
	out := Replace(str, ",", "\\n")
	fmt.Println(out)
	// var str = []string{"======",
	// 	"fd_time",
	// 	"=====",
	// 	"s",
	// 	"\n"}
	// // fd, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	// str1 := strings.Join(str, "\n")
	// // log.Println(str1)
	// b := "#! /bin/bash\n" + str1

	// print(b)
	// ioutil.WriteFile("b.txt", []byte(b), 755)

	// fd.WriteString(fmt.Sprint())
	// return nil
}
