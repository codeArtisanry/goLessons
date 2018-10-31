package main

import (
	"fmt"
	"net"
)

func iplocal() {
	// host, _ := os.Hostname()
	addrs, _ := net.LookupIP("192.168.5.100")
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			fmt.Println("IPv4: ", ipv4)
		}
	}
}

func main() {
	addrs, _ := net.InterfaceAddrs()
	fmt.Printf("%v\n", addrs)
	for _, addr := range addrs {
		fmt.Println("IPv4: ", addr)
	}
	println()
	iplocal()
}
