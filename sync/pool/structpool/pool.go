package main

import (
	"sync"
)

type Something struct {
	Name string
}

var Somethingpool = sync.Pool{
	New: func() interface{} {
		return &Something{}
	},
}

// GetBuffer returns a buffer from the pool.
func GetSomthing() (buf *Something) {
	return Somethingpool.Get().(*Something)
}

// PutBuffer returns a buffer to the pool.
// The buffer is reset before it is put back into circulation.
func PutSomething(buf *Something) {
	// buf.Reset()
	Somethingpool.Put(buf)
}

func main() {
	// s := pool.Get().(*Something)
	// defer pool.Put(s)
	s := GetSomthing()
	defer PutSomething(s)
	s.Name = "hello"
	// use the object
}
