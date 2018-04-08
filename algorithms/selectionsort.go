package algorithms


func SelectionSort(data []int) {

	var n int = len(data)
	var p_min int
	for i := 0; i < n-1; i++ {

		p_min = i

		for j := i; j < n; j++ {
			if data[j] < data[p_min] {
				p_min = j
			}
		}
		
		if p_min != i {
			data[i], data[p_min] = data[p_min], data[i]
		}
	}
}
