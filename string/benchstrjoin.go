package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

func benchmarkStringFunction(n int, index int) (d time.Duration) {
	v := "ni shuo wo shi bu shi tai wu liao le a? ni shuo wo shi bu shi tai wu liao le a?ni shuo wo shi bu shi tai wu liao le a?ni shuo wo shi bu shi tai wu liao le a?ni shuo wo shi bu shi tai wu liao le a?ni shuo wo shi bu shi tai wu liao le a?ni shuo wo shi bu shi tai wu liao le a?ni shuo wo shi bu shi tai wu liao le a?ni shuo wo shi bu shi tai wu liao le a?ni shuo wo shi bu shi tai wu liao le a?ni shuo wo shi bu shi tai wu liao le a?ni shuo wo shi bu shi tai wu liao le a?"
	var s string
	var buf bytes.Buffer
	var builder strings.Builder
	builder.Grow(100000)
	t0 := time.Now()
	for i := 0; i < n; i++ {
		switch index {
		case 0: // string +
			s = s + "[" + v + "]"
		case 1: // fmt.Sprintf
			s = fmt.Sprintf("%s[%s]", s, v)
		case 2: // strings.Join
			s = strings.Join([]string{s, "[", v, "]"}, "")
		case 3: // temporary bytes.Buffer
			b := bytes.Buffer{}
			b.WriteString("[")
			b.WriteString(v)
			b.WriteString("]")
			s = b.String()
		case 4: // stable bytes.Buffer
			buf.WriteString("[")
			buf.WriteString(v)
			buf.WriteString("]")
		case 5: // strings.builder
			builder.WriteString("[")
			builder.WriteString(v)
			builder.WriteString("]")
			// s = builder.String()
		}

		if i == n-1 {
			if index == 4 { // for stable bytes.Buffer
				s = buf.String()
			}
			fmt.Println(len(s)) // consume s to avoid compiler optimization
		}
	}
	t1 := time.Now()
	d = t1.Sub(t0)
	fmt.Printf("time of way(%d)=%v\n", index, d)
	return d
}

func main() {
	k := 6
	d := [6]time.Duration{}
	for i := 0; i < k; i++ {
		d[i] = benchmarkStringFunction(5000, i)
	}

	for i := 0; i < k-1; i++ {
		fmt.Printf("way %d is %6.1f times of way %d\n", i, float32(d[i])/float32(d[k-1]), k-1)
	}
}

// 4710000
// time of way(0)=9.232134919s
// 4710000
// time of way(1)=5.203677988s
// 4710000
// time of way(2)=13.641946846s
// 471
// time of way(3)=6.64121ms
// 4710000
// time of way(4)=9.274626ms
// 471
// time of way(5)=3.579785ms
// way 0 is 2579.0 times of way 5
// way 1 is 1453.6 times of way 5
// way 2 is 3810.8 times of way 5
// way 3 is    1.9 times of way 5
// way 4 is    2.6 times of way 5
