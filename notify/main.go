package main

import (
	"log"
	"time"

	"github.com/rinetd/go-learning/notify/notify"
)

func main() {
	// producer of "my_event"
	go func() {
		for {
			time.Sleep(time.Duration(1) * time.Second)
			notify.Post("my_event", time.Now().Unix())
		}
	}()
	go func() {
		for {
			time.Sleep(time.Duration(2) * time.Second)
			notify.Post("my_event", "string")
		}
	}()
	// observer of "my_event" (normally some independent component that
	// needs to be notified when "my_event" occurs)
	myEventChan := make(chan interface{})
	myEventChan1 := make(chan interface{})
	notify.Start("my_event", myEventChan)
	// notify.Start("my_event", myEventChan)
	go func() {
		for {
			data := <-myEventChan
			log.Printf("MY_EVENT: %#v", data)
		}
	}()
	go func() {
		for {
			data := <-myEventChan1
			log.Printf("MY_EVENT1: %#v", data)
		}
	}()
	for {
	}
}
