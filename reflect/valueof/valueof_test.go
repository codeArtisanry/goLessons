package valueof

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValue(test *testing.T) {

	type MyInt int
	var x MyInt = 7
	ty := reflect.TypeOf(x)
	fmt.Println(ty) //typeof.MyInt
	var xf float64 = 3.4
	fmt.Println("value:", reflect.ValueOf(xf))

	v := reflect.ValueOf(x)
	fmt.Println(v.Type()) //valueof.MyInt  ,
	fmt.Println(v.Kind()) //int            ,Kind 方法 会返回底层数据的类型 reflect.Int
	fmt.Println(v)        //7
}
