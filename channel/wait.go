package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	fmt.Println("goroutine监控中...", &stop)
	GoWait(stop)
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	stop <- true
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}
func GoWait(stop chan bool) {
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...", stop)
				time.Sleep(2 * time.Second)
			}
		}
	}()
}
