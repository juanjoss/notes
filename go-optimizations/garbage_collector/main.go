package main

import (
	"fmt"
	"runtime"
)

type memStat struct {
	NumGC   uint32
	Allocs  uint64
	Objects uint64
	Sys     uint64
}

func main() {
	var stat memStat
	updateMemStats(&stat)

	var arr [][]int

	for i := 0; i < 4; i++ {
		vec := make([]int, 0, 25000)
		arr = append(arr, vec)
	}

	arr = nil
	updateMemStats(&stat)

	runtime.GC()
	updateMemStats(&stat)
}

func updateMemStats(memStat *memStat) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	memStat.Allocs = (m.HeapAlloc)
	memStat.Objects = (m.HeapObjects)
	memStat.Sys = (m.Sys)
	memStat.NumGC = m.NumGC

	fmt.Println(memStat)
}
