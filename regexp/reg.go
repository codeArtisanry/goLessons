package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("a(x*)b(y)")
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-")[2])
	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))
}
