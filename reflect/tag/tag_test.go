package typeof

import (
	"fmt"
	"reflect"
	"testing"
)

func TestType(t *testing.T) {
	type Home struct {
		user int `json:"i"`
		name int `json:"j"`
	}

	home := new(Home)
	home.user = 5
	rcvr := reflect.ValueOf(home)
	typ := reflect.Indirect(rcvr).Type()
	fmt.Println(typ.Kind().String())
	x := typ.NumField()
	for i := 0; i < x; i++ {
		json := typ.Field(i).Tag.Get("json")
		fmt.Println(json)
	}

}
