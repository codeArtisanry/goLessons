package app

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

//测试获得当前时间, 并以各种形式显示
//时间戳是全世界唯一的计时数据. 同一个时间戳在不同的时区得到不同的字符串时间.
func TestGetCurrentTime(t *testing.T) {
	//定义时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		location = time.Local
	}

	t1 := time.Now().Unix()
	fmt.Println("当前时间戳: " + strconv.FormatInt(t1, 10))
	//把时间戳在location时区显示成字符串形式
	fmt.Println("当前时间的字符串格式: " + time.Now().In(location).Format("2006-01-02 15:04:05"))

	//字符串转时间戳 把字符串时间认为是在location时区
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2006-01-03 15:04:05", location)
	fmt.Println("2006-01-03 15:04:05 的时间戳是: " + strconv.FormatInt(t2.Unix(), 10))
}
