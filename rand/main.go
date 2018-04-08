package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Shuffle mixes randomly the input slice
func Shuffle(a []int) {
	rand.Seed(time.Now().UnixNano())

	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}
func main() {
	// 相同的种子会生成相同的随机值，
	// 所以如果不是每次都重置种子的话每次都会生成相同的序列
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int())     // int随机值，返回值为int
	fmt.Println(rand.Intn(100)) // [0,100)的随机值，返回值为int

	// 注意该函数只返回int32表示范围内的非负数，位数为31，因此该函数叫做Int31
	for index := 0; index < 100; index++ {
		fmt.Printf("0x%x", rand.Int()&0x80000000000000)       // 31位int随机值，返回值为int32
		fmt.Printf("\t0x%x\n", rand.Int63()&0x80000000000000) // 31位int随机值，返回值为int32
	}
	fmt.Println(rand.Int31n(100)) // [0,100)的随机值，返回值为int32
	fmt.Println(rand.Float32())   // 32位float随机值，返回值为float32
	fmt.Println(rand.Float64())   // 64位float随机值，返回值为float64

	// fmt.Println(rand.Shuffle()) // 64位float随机值，返回值为float64

	// 如果要产生负数到正数的随机值，只需要将生成的随机数减去相应数值即可
	fmt.Println(rand.Intn(100) - 50) // [-50, 50)的随机值

	//
	// Rand对象
	//

	r := rand.New(rand.NewSource(time.Now().Unix()))

	fmt.Println(r.Int())       // int随机值，返回值为int
	fmt.Println(r.Intn(100))   // [0,100)的随机值，返回值为int
	fmt.Println(r.Int31())     // 31位int随机值，返回值为int32
	fmt.Println(r.Int31n(100)) // [0,100)的随机值，返回值为int32
	fmt.Println(r.Float32())   // 32位float随机值，返回值为float32
	fmt.Println(r.Float64())   // 64位float随机值，返回值为float64

	// 如果要产生负数到正数的随机值，只需要将生成的随机数减去相应数值即可
	fmt.Println(r.Intn(100) - 50) // [-50, 50)的随机值
	n := 26
	i := 0
	for i < n {
		x := rand.Perm(3)
		fmt.Println(x)
		i += 1
	}
}
