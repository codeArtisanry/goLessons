package unsafe

import (
	"reflect"
	"unsafe"
)

func ToString(buf []byte) string {
	str := *(*reflect.StringHeader)(unsafe.Pointer(&buf))
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{Data: str.Data, Len: str.Len}))
}

func ToBytes(v string) []byte {
	sli := *(*reflect.StringHeader)(unsafe.Pointer(&v))
	arr := (*[1 << 30]byte)(unsafe.Pointer(sli.Data))
	return arr[:sli.Len]
}

func I2String(buf []byte) string {
	return *(*string)(unsafe.Pointer(&buf))
}

func I2Bytes(v string) []byte {
	return *(*[]byte)(unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&v))))
}
