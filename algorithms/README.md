#Benchmark
```
BenchmarkBuiltin1k-8               	   10000	    236781 ns/op	       0 B/op	       0 allocs/op
BenchmarkMergeSort1k-8             	   10000	    241577 ns/op	  194758 B/op	     999 allocs/op
BenchmarkMergeSortNoAlloc1k-8      	   10000	    243649 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickSort1k-8             	   20000	     87298 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickSortParallel1k-8     	   50000	     69501 ns/op	       0 B/op	       0 allocs/op
BenchmarkBubbleSort1k-8            	     100	 146537762 ns/op	       0 B/op	       0 allocs/op
BenchmarkSelectionSort1k-8         	     100	  97326609 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltin10k-8              	    1000	   2502735 ns/op	       0 B/op	       0 allocs/op
BenchmarkMergeSort10k-8            	     500	   2435212 ns/op	 1867577 B/op	    9999 allocs/op
BenchmarkMergeSortNoAlloc10k-8     	    1000	   2437066 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickSort10k-8            	    2000	    853220 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickSortParallel10k-8    	    5000	    474438 ns/op	       0 B/op	       0 allocs/op
BenchmarkBubbleSort10k-8           	      10	1406513097 ns/op	       0 B/op	       0 allocs/op
BenchmarkSelectionSort10k-8        	      10	 980868682 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltin1M-8               	       5	 230500588 ns/op	       6 B/op	       0 allocs/op
BenchmarkMergeSort1M-8             	       5	 269519997 ns/op	186757734 B/op	  999999 allocs/op
BenchmarkMergeSortNoAlloc1M-8      	       5	 229108418 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickSort1M-8             	      20	  85437935 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickSortParallel1M-8     	      30	  45751602 ns/op	      51 B/op	       0 allocs/op
BenchmarkBuiltin20M-8              	       1	4953628956 ns/op	      32 B/op	       1 allocs/op
BenchmarkMergeSort20M-8            	       1	3851057895 ns/op	4055167056 B/op	20000000 allocs/op
BenchmarkMergeSortNoAlloc20M-8     	       1	5285590601 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickSort20M-8            	       1	1794865996 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickSortParallel20M-8    	       2	 830047703 ns/op	     936 B/op	       6 allocs/op
BenchmarkBuiltin100M-8             	       1	27012136518 ns/op	      32 B/op	       1 allocs/op
BenchmarkQuickSort100M-8           	       1	9062693372 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickSortParallel100M-8   	       1	5563502291 ns/op	    1792 B/op	      12 allocs/op
PASS
ok  	_/root/go-first-steps/algorithms	197.134s

```
