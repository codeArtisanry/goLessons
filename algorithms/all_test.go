package algorithms

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

//initTest init two arrays
func initTest(size int) ([]int, []int) {
	array := make([]int, size)
	a1 := make([]int, size)
	a2 := make([]int, size)
	//rand.Read(a1)
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(size)
	}

	copy(a2, array)
	copy(a1, array)

	return a1, a2
}

func initBenchmark(size int) []int {

	array := make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(size)
	}

	return array
}

func testArray(t *testing.T, size int, a1, a2 []int) int {
	for i := 0; i < size; i++ {
		if a1[i] != a2[i] {
			fmt.Println("built-in:\n", a1)
			fmt.Println("mergesort:\n", a2)
			t.Error("pos:", i, "-built-in result != mergesort result")
			return -1
		}
	}
	return 0
}

func mergeSortGeneral(t *testing.T, size int) int {
	a1, a2 := initTest(size)
	sort.Ints(a1)
	a2 = MergeSort(a2)

	return testArray(t, size, a1, a2)
}

func TestMergeSort(t *testing.T) {
	var size = 100000

	for i := 1; i <= size; i *= 10 {
		if mergeSortGeneral(t, i) == -1 {
			break
		}
	}
}

func mergeSortNoAllocGeneral(t *testing.T, size int) int {
	a1, a2 := initTest(size)
	sort.Ints(a1)
	MergeSortNoAlloc(a2)

	return testArray(t, size, a1, a2)
}

func TestMergeSortFast(t *testing.T) {
	var size = 100000

	for i := 1; i <= size; i *= 10 {
		if mergeSortNoAllocGeneral(t, i) == -1 {
			break
		}
	}
}

func quickSortGeneral(t *testing.T, size int) int {
	a1, a2 := initTest(size)
	sort.Ints(a1)
	QuickSort(a2)

	return testArray(t, size, a1, a2)
}

// func TestQuickSort(t *testing.T){
// 	var size = 1000000

// 	for i := 1; i <= size; i *= 10 {
// 		if  quickSortGeneral(t, i) == -1 {
// 			break
// 		}
// 	}
// }

func quickSortParallelGeneral(t *testing.T, size int) int {
	a1, a2 := initTest(size)
	sort.Ints(a1)
	// QuickSortParallel(a2)

	return testArray(t, size, a1, a2)
}

func TestQuickSortParallel(t *testing.T) {
	var size = 1000000

	for i := 1; i <= size; i *= 10 {
		if quickSortParallelGeneral(t, i) == -1 {
			break
		}
	}
}

func bubbleSortGeneral(t *testing.T, size int) int {
	a1, a2 := initTest(size)
	sort.Ints(a1)
	BubbleSort(a2)

	return testArray(t, size, a1, a2)
}

func TestBubbleSort(t *testing.T) {
	var size = 100000

	for i := 1; i <= size; i *= 10 {
		if bubbleSortGeneral(t, i) == -1 {
			break
		}
	}
}

func selectionSortGeneral(t *testing.T, size int) int {
	a1, a2 := initTest(size)
	sort.Ints(a1)
	SelectionSort(a2)

	return testArray(t, size, a1, a2)
}

func TestSelectionSort(t *testing.T) {
	var size = 100000

	for i := 1; i <= size; i *= 10 {
		if selectionSortGeneral(t, i) == -1 {
			break
		}
	}
}

func BenchmarkBuiltin1k(b *testing.B) {

	var size = 1000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	sort.Ints(array)
}

func BenchmarkMergeSort1k(b *testing.B) {

	var size = 1000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	MergeSort(array)
}

func BenchmarkMergeSortNoAlloc1k(b *testing.B) {

	var size = 1000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	MergeSortNoAlloc(array)
}

func BenchmarkQuickSort1k(b *testing.B) {

	var size = 1000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	QuickSort(array)
}

func BenchmarkQuickSortParallel1k(b *testing.B) {

	var size = 1000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	HoareSort(array)
}

func BenchmarkBubbleSort1k(b *testing.B) {

	var size = 1000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	BubbleSort(array)
}

func BenchmarkSelectionSort1k(b *testing.B) {

	var size = 1000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	SelectionSort(array)
}

func BenchmarkBuiltin10k(b *testing.B) {

	var size = 10000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	sort.Ints(array)
}

func BenchmarkMergeSort10k(b *testing.B) {

	var size = 10000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	MergeSort(array)
}

func BenchmarkMergeSortNoAlloc10k(b *testing.B) {

	var size = 10000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	MergeSortNoAlloc(array)
}

func BenchmarkQuickSort10k(b *testing.B) {

	var size = 10000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	QuickSort(array)
}

func BenchmarkQuickSortParallel10k(b *testing.B) {

	var size = 10000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	HoareSort(array)
}

func BenchmarkBubbleSort10k(b *testing.B) {

	var size = 10000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	BubbleSort(array)
}

func BenchmarkSelectionSort10k(b *testing.B) {

	var size = 10000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	SelectionSort(array)
}

func BenchmarkBuiltin1M(b *testing.B) {

	var size = 1000000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	sort.Ints(array)
}

func BenchmarkMergeSort1M(b *testing.B) {

	var size = 1000000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	MergeSort(array)
}

func BenchmarkMergeSortNoAlloc1M(b *testing.B) {

	var size = 1000000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	MergeSortNoAlloc(array)
}

func BenchmarkQuickSort1M(b *testing.B) {

	var size = 1000000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	QuickSort(array)
}

func BenchmarkQuickSortParallel1M(b *testing.B) {

	var size = 1000000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	HoareSort(array)
}

func BenchmarkBuiltin20M(b *testing.B) {

	var size = 20000000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	sort.Ints(array)
}

func BenchmarkMergeSort20M(b *testing.B) {

	var size = 20000000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	MergeSort(array)
}

func BenchmarkMergeSortNoAlloc20M(b *testing.B) {

	var size = 20000000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	MergeSortNoAlloc(array)
}

func BenchmarkQuickSort20M(b *testing.B) {

	var size = 20000000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	QuickSort(array)
}

func BenchmarkQuickSortParallel20M(b *testing.B) {

	var size = 20000000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	HoareSort(array)
}

func BenchmarkBuiltin100M(b *testing.B) {

	var size = 100000000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	sort.Ints(array)
}

func BenchmarkQuickSort100M(b *testing.B) {

	var size = 100000000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	QuickSort(array)
}

func BenchmarkQuickSortParallel100M(b *testing.B) {

	var size = 100000000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()

	HoareSort(array)
}
