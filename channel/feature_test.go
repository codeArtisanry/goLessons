package main

import (
	"fmt"
	"testing"
)

// channel 的接收方式有三种:
// 1. <- ch
// 2. for c := range ch
// 3. select{}
// -   给一个 nil channel 发送数据，造成永远阻塞
// -   从一个 nil channel 接收数据，造成永远阻塞
// -   给一个已经关闭的 channel 发送数据，引起 panic
// -   从一个已经关闭的 channel 接收数据，立即返回一个零值

// 1. 给一个 nil channel 发送数据，造成永远阻塞
func TestSendToNilChan(t *testing.T) {
	var c chan string
	c <- "let's get started" // deadlock
	// fatal error: all goroutines are asleep - deadlock!
	// goroutine 1 [chan send (nil chan)]:

}

//  2. 从一个 nil channel 接收数据，造成永远阻塞
func TestReciveFromNilChan(t *testing.T) {
	var c chan string
	fmt.Println(<-c) // deadlock
	//fatal error: all goroutines are asleep - deadlock!
	//goroutine 1 [chan receive (nil chan)]:
}

// 3.给一个已经关闭的 channel 发送数据，引起 panic
func TestSendToCloseChan(t *testing.T) {
	var c = make(chan int, 100)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j //panic: send on closed channel
			}
			close(c)
		}()
	}
	for i := range c {
		fmt.Println(i)
	}
}

// 4.从一个已经关闭的 channel 接收数据，立即返回一个零值
func TestReciveFromCloseChan(t *testing.T) {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	for i := 0; i < 4; i++ {
		fmt.Printf("%d ", <-c) // prints 1 2 3 0
	}
}
