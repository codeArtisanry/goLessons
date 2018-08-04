package typeof

import (
	"fmt"
	"reflect"
	"testing"
)

func TestType(t *testing.T) {

	type MyInt int
	var x MyInt = 7
	v := reflect.TypeOf(x) //typeof.MyInt
	fmt.Println(v)
}
