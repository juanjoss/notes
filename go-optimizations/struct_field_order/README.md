# Struct Field Ordering

> *Order struct fields by type in increasing order according to its size.*

This is because of how memory is filled, generally in chunks of 32-64 bits according to the system architecture. In the case of `badStruct` the computer would read `SomeBool` (1 byte) first in a 64 bit chunk, and the remaining is padded, which results in a inneficient memory usage.

That's because placing structs according to its 

```go
package main

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

	println(unsafe.Sizeof(bs)) // 48 bytes
	println(unsafe.Sizeof(gs)) // 40 bytes
}
```

```bash
go run main.go
```