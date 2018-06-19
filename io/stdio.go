package main

import (
	"io"
	"log"
	"os"
)

func main() {
	in := os.Stdin
	out := os.Stdout
	if _, err := io.Copy(out, in); err != nil {
		log.Fatal(err)
	}
}
