package psort

import (
	"math/rand"
	"os"
	"runtime"
	"sort"
	"strconv"
	"sync"
	"testing"
	"time"
)

var l = 10000

func gendata() {
	rand.Seed(time.Now().UnixNano())
	var data = rand.Perm(l)
	f, _ := os.Create("data")
	defer f.Close()
	// buf := new(bytes.Buffer)
	// binary.Write(buf, binary.LittleEndian, data)
	// fmt.Println(buf.Bytes())
	// f.Write(buf.Bytes())
	for i := range data {
		f.WriteString(strconv.Itoa(data[i]))
		f.WriteString("\t")
	}
	// f.Sync()
}

func TestQuickSort(t *testing.T) {

	// var l = 10000000
	rand.Seed(time.Now().UnixNano())
	var data = rand.Perm(l)
	if sort.IsSorted(sort.IntSlice(data)) {
		// time.Sleep(2 * time.Second)
		t.Error("sorted")
	} else {
		t.Log("unsorted")
	}
	QuickSort(data)
	if sort.IsSorted(sort.IntSlice(data)) {
		// time.Sleep(2 * time.Second)
		t.Log("sorted")
	} else {
		t.Error("unsorted")
	}

}
func Test_quickSortParallel(t *testing.T) {

	// var l = 10000000
	rand.Seed(time.Now().UnixNano())
	var data = rand.Perm(l)
	gendata()
	type args struct {
		data []int
		n    int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "string(len)",
			// TODO: Add test cases.
			args: args{
				n:    l,
				data: data,
			},
		},
	}
	for _, tt := range tests {
		if sort.IsSorted(sort.IntSlice(data)) {
			// time.Sleep(2 * time.Second)
			t.Error("sorted")
		} else {
			t.Log("unsorted")
		}
		// t.Log(data)
		t.Run(tt.name, func(t *testing.T) {
			runtime.GOMAXPROCS(runtime.NumCPU())
			var wg = new(sync.WaitGroup)
			wg.Add(1)
			quickSortParallel(tt.args.data, wg, runtime.NumCPU()<<10)
			wg.Wait()
		})

		if sort.IsSorted(sort.IntSlice(data)) {
			// time.Sleep(2 * time.Second)
			t.Log("sorted")
		} else {
			t.Error("unsorted")
		}
		// t.Log(data)
	}
}

func Test_quickSortParallel3(t *testing.T) {

	// var l = 10000000
	rand.Seed(time.Now().UnixNano())
	var data = rand.Perm(l)
	gendata()
	type args struct {
		data []int
		n    int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "string(len)",
			// TODO: Add test cases.
			args: args{
				n:    l,
				data: data,
			},
		},
	}
	for _, tt := range tests {
		if sort.IsSorted(sort.IntSlice(data)) {
			// time.Sleep(2 * time.Second)
			t.Error("sorted")
		} else {
			t.Log("unsorted")
		}
		// t.Log(data)
		t.Run(tt.name, func(t *testing.T) {
			runtime.GOMAXPROCS(runtime.NumCPU())
			var wg = new(sync.WaitGroup)
			wg.Add(1)
			quickSortParallel3(tt.args.data, wg, runtime.NumCPU()<<10)
			wg.Wait()
		})

		if sort.IsSorted(sort.IntSlice(data)) {
			// time.Sleep(2 * time.Second)
			t.Log("sorted")
		} else {
			t.Error("unsorted")
		}
		// t.Log(data)
	}
}
