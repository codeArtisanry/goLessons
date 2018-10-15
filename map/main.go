package main

import (
	"fmt"
	"sync"

	"math/rand"
)

func main1() {
	const N = 6
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			m[rand.Int()] = rand.Int()
		}()
	}
	wg.Wait()
	println(len(m))

}

// mutex 实现并发
var counter = struct {
	sync.RWMutex
	m map[string]int
}{m: make(map[string]int)}

func read() {

	counter.RLock()
	n := counter.m["some_key"]
	counter.RUnlock()
	fmt.Println("some_key:", n)
}

func write() {

	counter.Lock()
	counter.m["some_key"]++
	counter.Unlock()
}
