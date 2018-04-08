package unsafe

import (
	"bytes"
	"testing"
)

// ----------------------------------------------------------

func Test(t *testing.T) {

	str := "Hello, world"
	b := []byte(str)

	b2 := ToBytes(str)
	t.Logf("str %p", &str)
	t.Logf("b %p", &b[0])
	t.Logf("b2 %T", &b2)
	if !bytes.Equal(b, b2) {
		t.Fatal("ToBytes failed:", string(b2))
	}

	str2 := ToString(b)
	if str != str2 {
		t.Fatal("ToString failed:", str2)
	}
}
