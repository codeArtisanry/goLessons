package algorithms

import (
	"math/rand"
	"runtime"
	"sort"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	var a = make(sort.IntSlice, 20000000)
	var b = make(sort.IntSlice, 20000000)
	var c = make(sort.IntSlice, 20000000)
	var d = make(sort.IntSlice, 20000000)
	var e = make(sort.IntSlice, 20000000)
	var f = make(sort.IntSlice, 20000000)

	rand.Seed(time.Now().Unix())
	a = sort.IntSlice(rand.Perm(10000000))
	a = append(a, a...)
	copy(b, a)
	copy(c, a)
	copy(d, a)
	copy(e, a)
	copy(f, a)

	t.Logf("%p\t%p", a, b)

	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "sort1",
			args: args{data: rand.Perm(1)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.name)
			// t.Log("before sort:", tt.args.data)
			var s time.Time
			var end time.Duration
			// s = time.Now()
			// QuickSort([]int(a))
			// end = time.Since(s)
			// if sort.IsSorted(a) {
			// 	// t.Log("after sort:", tt.args.data)
			// 	t.Log(end)

			// } else {
			// 	t.Error("就是不通过")
			// }
			// s = time.Now()
			// QuickSortParallel([]int(b))
			// end = time.Since(s)
			// if sort.IsSorted(b) {
			// 	// t.Log("after sort:", tt.args.data)
			// 	t.Log(end)

			// } else {
			// 	t.Error("就是不通过")
			// }

			// s = time.Now()
			// sort.Sort(c)
			// end = time.Since(s)
			// if sort.IsSorted(c) {
			// 	// t.Log("after sort:", tt.args.data)
			// 	t.Log(end)

			// } else {
			// 	t.Error("就是不通过")
			// }
			s = time.Now()
			QuickSortP([]int(b), runtime.NumCPU()*8)
			end = time.Since(s)
			if sort.IsSorted(b) {
				// t.Log("after sort:", tt.args.data)
				t.Log(end)

			} else {
				t.Error("就是不通过")
			}
			s = time.Now()
			QuickSortP([]int(c), runtime.NumCPU()*10)
			end = time.Since(s)
			if sort.IsSorted(c) {
				// t.Log("after sort:", tt.args.data)
				t.Log(end)

			} else {
				t.Error("就是不通过")
			}
			s = time.Now()
			QuickSortP([]int(d), runtime.NumCPU()*12)
			end = time.Since(s)
			if sort.IsSorted(d) {
				// t.Log("after sort:", tt.args.data)
				t.Log(end)

			} else {
				t.Error("就是不通过")
			}
			s = time.Now()
			QuickSortP([]int(e), runtime.NumCPU()*14)
			end = time.Since(s)
			if sort.IsSorted(e) {
				// t.Log("after sort:", tt.args.data)
				t.Log(end)

			} else {
				t.Error("就是不通过")
			}
			s = time.Now()
			QuickSortP([]int(f), runtime.NumCPU()*16)
			end = time.Since(s)
			if sort.IsSorted(f) {
				// t.Log("after sort:", tt.args.data)
				t.Log(end)

			} else {
				t.Error("就是不通过")
			}
			// t.Error("就是不通过")
		})
	}
}
