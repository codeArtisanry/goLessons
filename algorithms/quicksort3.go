package algorithms

import "sync"

func QuickSort3(src []int, first, last int, wg *sync.WaitGroup) {
	defer wg.Done()
	flag := first
	left := first
	right := last
	if first >= last {
		return
	}
	for first < last {
		for first < last {
			if src[last] >= src[flag] {
				last--
				continue
			} else {
				tmp := src[last]
				src[last] = src[flag]
				src[flag] = tmp
				flag = last
				break
			}
		}
		for first < last {
			if src[first] <= src[flag] {
				first++
				continue
			} else {
				tmp := src[first]
				src[first] = src[flag]
				src[flag] = tmp
				flag = first
				break
			}
		}
	}
	wg.Add(2)
	go QuickSort3(src, left, flag-1, wg)
	go QuickSort3(src, flag+1, right, wg)
}
