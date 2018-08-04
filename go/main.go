package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	data := []byte("sqwerjohovaoshqohro234hlcvnz")
	wait(data, len(data))
	// wait(data, len(data))
	for {
	}
}
func wait(data []byte, n int) {
	var pivot = data[0]
	var left = 0
	var right = n - 1
	var end int32
	//partition

	fmt.Println(data)
	for left <= right {
		atomic.StoreInt32(&end, 0)
		// for data[left] < pivot {
		// 	left++
		// }
		go func() {
			for data[left] < pivot {
				left++
			}
			atomic.StoreInt32(&end, 1)
		}()
		for data[right] > pivot {
			right--
		}
		for {
			if 1 == atomic.LoadInt32(&end) {
				break
			}
		}
		// println(left)
		if left <= right {
			// swap(&data[left], &data[right])
			data[left], data[right] = data[right], data[left]
			left++
			right--
		}
		fmt.Println(data)
	}

}
