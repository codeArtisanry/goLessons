package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// 使用waitgroup
// sync包中的Waitgroup结构，是Go语言为我们提供的多个goroutine之间同步的好刀。下面是官方文档对它的描述：

// A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for.
// Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.

// 通常情况下，我们像下面这样使用waitgroup:

//     创建一个Waitgroup的实例，假设此处我们叫它wg
//     在每个goroutine启动的时候，调用wg.Add(1)，这个操作可以在goroutine启动之前调用，也可以在goroutine里面调用。当然，也可以在创建n个goroutine前调用wg.Add(n)
//     当每个goroutine完成任务后，调用wg.Done()
//     在等待所有goroutine的地方调用wg.Wait()，它在所有执行了wg.Add(1)的goroutine都调用完wg.Done()前阻塞，当所有goroutine都调用完wg.Done()之后它会返回。

// 那么，如果我们的goroutine是一匹不知疲倦的牛，一直孜孜不倦地工作的话，如何在主流程中告知并等待它退出呢？像下面这样做：

type Service struct {
	// Other things

	ch chan bool
	wg *sync.WaitGroup
}

func NewService() *Service {
	s := &Service{
		// Init Other things
		ch: make(chan bool),
		// wg: &sync.WaitGroup{},
		wg: new(sync.WaitGroup),
	}
	return s
}

func (s *Service) Stop() {
	close(s.ch)
	s.wg.Wait()
}

func (s *Service) Serve() {
	s.wg.Add(1)
	defer s.wg.Done()

	for {
		select {
		case <-s.ch:
			fmt.Println("stopping...")
			return
		default:
		}
		s.wg.Add(1)
		go s.anotherServer()
	}
}
func (s *Service) anotherServer() {
	defer s.wg.Done()
	for {
		select {
		case <-s.ch:
			fmt.Println("stopping...")
			return
		default:
		}

		// Do something
	}
}

func main() {

	service := NewService()
	go service.Serve()

	// Handle SIGINT and SIGTERM.
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println(<-ch)

	// Stop the service gracefully.
	service.Stop()

	//########
	done := make(chan bool, 1)
	signals := make(chan os.Signal, 1)
	// os/signal包中的两个方法：1. 一个是signal.Notify(ch)方法用来监听收到的信号；2. s := <-ch  3.一个是 signal.Stop(ch)方法用来取消监听。
	// 第一个参数表示接收信号的channel, 第二个及后面的参数表示设置要监听的信号，如果不设置表示监听所有的信号。
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signals
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

}
