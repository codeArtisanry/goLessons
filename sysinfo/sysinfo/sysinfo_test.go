// A wrapper around the linux syscall sysinfo(2).
package sysinfo

import (
	"fmt"
	t "testing"
)

func TestNew(t *t.T) {
	fmt.Println(New())
	fmt.Println()
	fmt.Println(New().ToString())
	fmt.Println()
}

func ExampleNew() {
	si := New()
	fmt.Printf("%v\n", si.Loads)
	fmt.Println(si.ToString())
}

/*
amd64
BenchmarkNew		 1000000	      2391 ns/op
BenchmarkString		  200000	     11041 ns/op
BenchmarkToString	  200000	     11206 ns/op

arm
BenchmarkNew		  200000	      8431 ns/op
BenchmarkString		   10000	    199392 ns/op
BenchmarkToString	   10000	    203040 ns/op
*/
func BenchmarkNew(b *t.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}

func BenchmarkString(b *t.B) {
	for i := 0; i < b.N; i++ {
		New().String()
	}
}

func BenchmarkToString(b *t.B) {
	for i := 0; i < b.N; i++ {
		New().ToString()
	}
}
