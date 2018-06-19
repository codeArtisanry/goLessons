package main

import (
	"fmt"
	"net"
	"os"
)

/**
 * 将域名解析Ip地址
 * 获得域名对应的所有Ip地址
 */
func Main_net() {

	domain := "www.baidu.com"
	ipAddr, err := net.ResolveIPAddr("ip", domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}
	fmt.Fprintf(os.Stdout, "%s IP: %s Network: %s Zone: %s\n", ipAddr.String(), ipAddr.IP, ipAddr.Network(), ipAddr.Zone)

	// 百度，虽然只有一个域名，但实际上，他对应电信，网通，联通等又有多个IP地址
	ns, err := net.LookupHost(domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}

	for _, n := range ns {
		fmt.Fprintf(os.Stdout, "%s\n", n) // 115.239.210.26    115.239.210.27 这2个地址打开都是百度
	}
}
