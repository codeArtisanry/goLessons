package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var hit int64
var wg sync.WaitGroup

func main() {
	// wg.Add(2)
	for i := 0; i < 10; i++ {
		t := int64(0)
		i := i
		ch := make(chan int)
		ch1 := make(chan int)
		go func() {
			// atomic.AddInt64(&t, int64(i))
			atomic.AddInt64(&t, int64(10))
			// fmt.Printf("t%d %d addr: %p\n", i, t, &t)

			ch1 <- 1
		}()
		go func() {
			<-ch1
			atomic.AddInt64(&t, int64(i))
			atomic.AddInt64(&t, int64(1))
			// fmt.Printf("t%d %d addr: %p\n", i, t, &t)

			ch <- 1
		}()
		go func() {
			<-ch
			fmt.Printf("t%d %d addr: %p\n", i, t, &t)
		}()
	}
	time.Sleep(2 * time.Second)
	// go incrementor("foo")
	// go incrementor("bar")
	// wg.Wait()
	fmt.Println("Hit times:", hit)
}

func incrementor(s string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&hit, 2)
		fmt.Println(s, i, "Hit times:", atomic.LoadInt64(&hit))
	}
	wg.Done()
}
