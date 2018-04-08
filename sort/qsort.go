package main

import "fmt"

// threads 线程标识创建线程的个数
func quicksort(nums []int, ch chan int, level int, threads int) {
	level = level * 2
	if len(nums) == 1 {
		ch <- nums[0]
		close(ch)
		return
	} //ch<-nums[0] 表示将nums[0] 数据写到ch通道中
	if len(nums) == 0 {
		close(ch)
		return
	}

	less := make([]int, 0) //
	greater := make([]int, 0)
	left := nums[0] //快速排序的轴
	nums = nums[1:]

	//从左向右扫描数据 大于轴的放到greater里小于的放到less中
	for _, num_data := range nums {
		switch {
		case num_data <= left:
			less = append(less, num_data)
		case num_data > left:
			greater = append(greater, num_data)
		}
	}

	left_ch := make(chan int, len(less))
	right_ch := make(chan int, len(greater))

	if level <= threads {
		go quicksort(less, left_ch, level, threads) //分任务
		go quicksort(greater, right_ch, level, threads)
	} else {
		quicksort(less, left_ch, level, threads)
		quicksort(greater, right_ch, level, threads)
	}

	//合并数据
	for i := range left_ch {
		ch <- i
	}
	ch <- left
	for i := range right_ch {
		ch <- i
	}
	close(ch)
	return
}

func main() {
	x := []int{3, 1, 4, 1, 5, 9, 2, 6}
	ch := make(chan int)
	go quicksort(x, ch, 0, 0) // 0 0 表示不限制线程个数
	for v := range ch {
		fmt.Println(v)
	}
}
