package psort

import (
	"fmt"
	"math/rand"
	"runtime"
	"sort"
	"sync"
	"time"
)

func swap(a, b *int) {
	*a, *b = *b, *a
}
func analysis(data []int, fn func([]int)) {
	x := time.Now()
	fn(data)
	e := time.Since(x)
	if sort.IsSorted(sort.IntSlice(data)) {
		fmt.Println(e)
	} else {
		fmt.Println("erro")
	}
	// time.Sleep(time.Second * 6)
	// fmt.Println(data)
}

// GoSort is inner sort for go
func GoSort(data []int) {
	sort.Sort(sort.IntSlice(data))
}
func main() {
	var num = 10000
	rand.Seed(time.Now().UnixNano())
	var data = make([][]int, 2)
	for k := range data {
		data[k] = make([]int, num)
	}
	a := rand.Perm(num)
	// for k := range data {
	// 	copy(data[k], a)
	// }
	println("start")
	analysis(a, HoareSort)
	// analysis(data[0], HoareSort)
	// analysis(data[1], QuickSort)
	// analysis(data[2], GoSort)
	// x := time.Now()
	// HoareSort(b)
	// fmt.Println(time.Since(x))
	// 1.569865083s

	// x = time.Now()
	// QuickSort(a)
	// fmt.Println(time.Since(x))
	// 8.595509892s

	// x = time.Now()
	// sort.Sort(sort.IntSlice(c))
	// fmt.Println(time.Since(x))
	// 20.921417413s
}

// QuickSort for single
func QuickSort(data []int) {

	var n = len(data)
	var middle = (n - 1) / 2
	var pivot = data[middle]

	var left = 0
	var right = n - 1

	//partition
	for left <= right {
		for data[left] < pivot {
			left++
		}
		for data[right] > pivot {
			right--
		}
		if left <= right {
			// swap(&data[left], &data[right])
			data[left], data[right] = data[right], data[left]
			left++
			right--
		}
	}

	if right > 0 {
		QuickSort(data[:right+1])
	}
	if left < n {
		QuickSort(data[left:])
	}
}

// HoareSort is faster sort
func HoareSort(data []int) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg = new(sync.WaitGroup)
	wg.Add(1)
	quickSortParallel(data, wg, runtime.NumCPU()<<10)
	wg.Wait()
}

func quickSortParallel(data []int, wg *sync.WaitGroup, threads int) {
	var n = len(data)
	defer wg.Done()

	if threads <= 1 || n < 1000 {
		QuickSort(data)
	} else {

		var middle = (n - 1) / 2
		var pivot = data[middle]

		var left = 0
		var right = n - 1

		//partition
		for left <= right {
			// for data[left] < pivot {
			// 	left++
			// }

			for data[left] < pivot {
				left++
			}
			for data[right] > pivot {
				right--
			}
			// println(left)
			if left <= right {
				// swap(&data[left], &data[right])
				data[left], data[right] = data[right], data[left]
				left++
				right--
			}
		}

		// var ch1 = make(chan int)
		// var ch2 = make(chan int)
		// var ch11 = make(chan int)
		// var ch22 = make(chan int)
		// //partition

		// // for data[left] < pivot {
		// // 	left++
		// // }

		// go func() {
		// 	for range ch1 {
		// 		for data[left] > pivot {
		// 			left++
		// 		}
		// 		// println("after:", left, right, data[left], data[right])
		// 		// println("chanel 1")
		// 		ch11 <- 1
		// 	}
		// }()
		// go func() {
		// 	i := 1
		// 	for range ch2 {

		// 		for data[right] > pivot {
		// 			right--
		// 		}
		// 		i++
		// 		// print(i, " ")
		// 		// println("chanel 2")

		// 		// println(i, ">>>>>:", left, right, data[left], data[right])
		// 		ch22 <- 1
		// 	}
		// }()
		// for left <= right {
		// 	println("go....")

		// 	ch2 <- 1
		// 	ch1 <- 1
		// 	println("gogo....")

		// 	<-ch11
		// 	<-ch22
		// 	println("gogogo....")
		// 	println(left, right, data[left], data[right])

		// 	// 	// swap(&data[left], &data[right])
		// 	if data[left] > data[right] {

		// 		data[left], data[right] = data[right], data[left]

		// 	}

		// 	left++
		// 	right--

		// }
		// // println("pivot:", pivot)

		// close(ch1)
		// close(ch2)
		// close(ch11)
		// close(ch22)
		// for {
		// }
		// // println(left)
		// // if left <= right {
		// // 	// swap(&data[left], &data[right])
		// // 	data[left], data[right] = data[right], data[left]
		// // 	left++
		// // 	right--
		// // }

		wgN := new(sync.WaitGroup)
		if right > 0 {
			wgN.Add(1)
			go quickSortParallel(data[:right+1], wgN, threads/2)
		}
		if left < n {
			wgN.Add(1)
			go quickSortParallel(data[left:], wgN, threads/2)
		}
		wgN.Wait()
	}
}
