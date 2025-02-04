package main

import (
	"fmt"
	"net"
	"os"
)

func GetIntranetIp() {
	addrs, err := net.InterfaceAddrs()
	fmt.Println("ip:", addrs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("ip:", ipnet.IP.String())
			}

		}
	}
}

func main() {
	GetIntranetIp()
}
