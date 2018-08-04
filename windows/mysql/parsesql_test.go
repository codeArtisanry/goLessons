package main

import (
	"fmt"
	"os"
	"testing"
)

func TestParsesql(t *testing.T) {
	f, err := os.Open("a.sql")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	s := Parsesql(f)
	fmt.Println(s, len(s))
}

func TestParseFile(t *testing.T) {
	fmt.Println(ParseFile("a.sql"))
}
