// A wrapper around the linux syscall sysinfo(2).
package main

import "github.com/rinetd/go-learning/sysinfo/sysinfo"

func main() {
	// println(sysinfo.New().ToString())
	// println(sysinfo.New().String())
	sysinfo.New().Print()
}
