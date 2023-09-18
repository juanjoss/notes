package main

import "testing"

func BenchmarkAvg(b *testing.B) {
	nums := []float64{1, 2, 3, 4, 5}

	for i := 0; i < b.N; i++ {
		avg(sum, nums...)
	}
}

func BenchmarkAvgWithoutInlining(b *testing.B) {
	nums := []float64{1, 2, 3, 4, 5}

	for i := 0; i < b.N; i++ {
		avg(sumWithoutInlining, nums...)
	}
}

func BenchmarkAvgWithHighInliningCost(b *testing.B) {
	nums := []float64{1, 2, 3, 4, 5}

	for i := 0; i < b.N; i++ {
		avg(sumWithHighInliningCost, nums...)
	}
}
