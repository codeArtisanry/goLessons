package algorithms

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
}
func GoSort(data []int) {
	sort.Sort(sort.IntSlice(data))
}
func QuickSort_Test() {
	var d = []int{2, 7, 10, 8, 3, 5, 9, 6, 9, 10, 12}
	fmt.Println(d)
	QuickSort(d)
	fmt.Println(d)

	var num int = 10000
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

// QuickSort
func QuickSort(data []int) {
	var n int = len(data)
	var middle int = (n - 1) / 2
	var pivot int = data[middle]

	var left int = 0
	var right int = n - 1
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

// HoareSort
func HoareSort(data []int) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg = new(sync.WaitGroup)
	// wg.Add(1)
	QuickSortParallel(data, wg, runtime.NumCPU()<<10)
	wg.Wait()
}

func QuickSortParallel(data []int, wg *sync.WaitGroup, threads int) {
	var n int = len(data)
	defer wg.Done()

	if threads <= 1 || n < 1000 {
		QuickSort(data)
	} else {

		var middle int = (n - 1) / 2
		var pivot int = data[middle]

		var left int = 0
		var right int = n - 1

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

		wgN := new(sync.WaitGroup)
		if right > 0 {
			wgN.Add(1)
			go QuickSortParallel(data[:right+1], wgN, threads/2)
		}
		if left < n {
			wgN.Add(1)
			go QuickSortParallel(data[left:], wgN, threads/2)
		}
		wgN.Wait()
	}
}
