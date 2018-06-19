package main

import (
	"fmt"
	"net"
)

func Main_dial()) {
	fmt.Print("client start \n")

	var addr net.TCPAddr
	// addr.IP = net.ParseIP("0.0.0.0")
	addr.IP = net.ParseIP("127.0.0.1")
	addr.Port = 12345

	Conn, e := net.DialTCP("tcp", nil, &addr)
	fmt.Println("Conn,e:", Conn, e)

	buf := make([]byte, 512)
	Conn.Read(buf)

	// fin.Read(buf)

	fmt.Printf("%q\n", buf)
	var MAX int
	MAX = len(buf)
	fmt.Println(MAX)

	for i := 0; i < MAX; i++ {
		if i == 0 {
			a := buf[i]
			fmt.Printf("%q\n", a)
		} else {
			if i%2 == 0 {
				a := buf[i]
				fmt.Printf("%q\n", a)
			}
		}
	}
	fmt.Println("buf:", buf)

	fmt.Print("client end \n")
	var message []byte
	Conn.Write(message)
}
