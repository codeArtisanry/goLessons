package main

import (
	"fmt"
	"regexp"
	"testing"
)

const nginxErrorRegExp = "^(?P<time>[\\d+/ :]+) \\[(?P<severity>.+)\\] .*?: (?P<message>.+), client: (?P<client>.+), server: (?P<server>.+), request: \"(?P<method>\\S+) (?P<path>\\S+) (?P<version>.+?)\", host: \"(?P<host>.+)\"$"

func Testmain(t *testing.T) {
	re := regexp.MustCompile("a(x*)b(y)")
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-")[2])
	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))
}

var skipRe = regexp.MustCompile(`\[(?i:ci *skip|skip *ci)\]`)

func TestSkip(t *testing.T) {

	skipMatch := skipRe.FindString("build.Message [CI skip]")
	fmt.Println(skipMatch)
}
