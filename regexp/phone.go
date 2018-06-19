package main

import "fmt"
import "regexp"

const (
	regphone = "^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0,5-9]))\\d{8}$"
)

func validate(mobileNum string) bool {
	reg := regexp.MustCompile(regphone)
	return reg.MatchString(mobileNum)
}

var p = "15966638724"

func Main_phone() {
	fmt.Println(validate(p))

}
