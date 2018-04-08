package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nim4/DBShield/dbshield/logger"
)

func signals() <-chan bool {
	quit := make(chan bool)

	go func() {
		signals := make(chan os.Signal)
		defer close(signals)

		signal.Notify(signals, syscall.SIGQUIT, syscall.SIGTERM, os.Interrupt)
		// defer signalStop(signals)

		<-signals
		quit <- true
	}()

	return quit
}

func TrapSignals() {
	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGUSR1)

		for sig := range sigchan {
			switch sig {
			case syscall.SIGTERM, syscall.SIGINT:
				os.Exit(0)
			case syscall.SIGQUIT:
				os.Exit(1)
			case syscall.SIGHUP:
				println("syscall.SIGUSR1")
			case syscall.SIGUSR1:
				println("syscall.SIGUSR1")
			}
		}
	}()
}
func signal1() {
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	n := <-signals

	if n == os.Interrupt {
		fmt.Println("输出：Ctrl-C SIGINT")
	}
	if n == os.Kill {
		fmt.Println("输出：kill -9 SIGKILL")
	}
	if n == syscall.SIGUSR1 {
		fmt.Println("signal user1")
	}
	if n == syscall.SIGUSR2 {
		fmt.Println("signal user2")
	}
}

// 捕捉信号
func catchSignal() {

	c := make(chan os.Signal)
	// todo 配置热更新, windows 不支持 syscall.SIGUSR1, syscall.SIGUSR2
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	for s := range c {
		logger.SetLevel()
		logger.Info("收到信号 -- ", s)
		switch s {
		case os.Interrupt:

			logger.Info("收到终端断开信号, 忽略后继续运行～～")
		case syscall.SIGHUP:
			logger.Info("收到终端断开信号, 忽略后继续运行～～")
		case syscall.SIGINT, syscall.SIGTERM:
			// shutdown()
			logger.Info("收到终端断开信号, 忽略后继续运行～～")

		}
	}
}
func main() {
	// TrapSignals()
	// signal1()
	catchSignal()
}

// 测试：
// ctrl + c
// kill pid
// kill -USR1 pid
// kill -USR2 pid
