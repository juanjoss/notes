package main

import (
	"unsafe"
)

type badStruct struct {
	SomeBool       bool
	SomeInt64      int64
	SomeOtherBool  bool
	SomeStr        string
	SomeOtherInt64 int64
}

type goodStruct struct {
	SomeInt64      int64
	SomeOtherInt64 int64
	SomeStr        string
	SomeBool       bool
	SomeOtherBool  bool
}

func main() {
	var bs badStruct
	var gs goodStruct

	println(unsafe.Sizeof(bs))
	println(unsafe.Sizeof(gs))
}
