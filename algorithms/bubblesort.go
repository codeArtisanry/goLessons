package algorithms

func BubbleSort(data []int) {

	var n int = len(data)
	
	for i := n-1 ; i > 0; i-- {
		for j := 0; j < i; j++ {
			if data[j] > data[j + 1] {
				data[j], data[j + 1] = data[j + 1], data[j]
			}
		}
	}	
}