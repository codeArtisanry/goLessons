package main

import (
	"math/rand"
	"sync"
)

// sync.Map的实现有几个优化点，这里先列出来，我们后面慢慢分析。

//     空间换时间。 通过冗余的两个数据结构(read、dirty),实现加锁对性能的影响。
//     使用只读数据(read)，避免读写冲突。
//     动态调整，miss次数多了之后，将dirty数据提升为read。
//     double-checking。
//     延迟删除。 删除一个键值只是打标记，只有在提升dirty的时候才清理删除的数据。
//     优先从read读取、更新、删除，因为对read的读取不需要锁。

func main() {

	// var syncm sync.Map
	// fmt.Printf("%#v\n\n", syncm)

	const N = 6
	var m sync.Map
	wg := &sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			m.Store(rand.Int(), rand.Int())

		}()
	}
	wg.Wait()
	m.Range()
	// syncm.Store("abc", "1")
	// res, err := syncm.Load("abc")
	// fmt.Println(res, err)
}
