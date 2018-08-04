package letcode

func removeDuplicates(nums []int) int {
	var m = make(map[int]int)
	for _, v := range nums {
		m[v] = v
	}
	return len(m)
}
