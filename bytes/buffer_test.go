package bytes

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

// bytes.NewBuffer 实现了很多基本的接口，可以通过 bytes 包学习接口的实现
func TestIO(b *testing.T) {
	buf := bytes.NewBuffer([]byte("Hello World!"))
	bu := make([]byte, buf.Len())
	// var bb = bytes.Buffer{}

	n, err := buf.Read(bu)
	fmt.Printf("%s   %v\n", bu[:n], err)
	// Hello World!   <nil>

	buf.WriteString("ABCDEFG\n")
	buf.WriteTo(os.Stdout)
	// ABCDEFG

	n, err = buf.Write(bu)
	fmt.Printf("%d   %s   %v\n", n, buf.String(), err)
	// 12   Hello World!   <nil>

	c, err := buf.ReadByte()
	fmt.Printf("%c   %s   %v\n", c, buf.String(), err)
	// H   ello World!   <nil>

	c, err = buf.ReadByte()
	fmt.Printf("%c   %s   %v\n", c, buf.String(), err)
	// e   llo World!   <nil>

	err = buf.UnreadByte()
	fmt.Printf("%s   %v\n", buf.String(), err)
	// ello World!   <nil>

	err = buf.UnreadByte()
	fmt.Printf("%s   %v\n", buf.String(), err)
	// ello World!   bytes.Buffer: UnreadByte: previous operation was not a read
}
