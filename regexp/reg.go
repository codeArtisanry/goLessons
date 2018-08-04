package main

import (
	"log"

	"github.com/moovweb/rubex"
)

func main() {
	rxp := rubex.MustCompile("?:=/,")

	result := rxp.FindString("a me my")
	str := rxp.ReplaceAllString(`asdf\,dasdf\,sa,a\\,sdf`, "\n")
	log.Println(str)
	if result != "" {
		// FOUND A STRING!! YAY! Must be "a" in this instance
		log.Println("find")
	} else {
		// no good
	}
}
