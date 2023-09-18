# Inlining

Inlining is the process of replacing a function call with the content of that function. This reduces the number of instructions the processor needs to execute (calling functions has a cost). 

- Pro: Reduces function calling overhead.
- Con: Increases binary size.

Example:

```go
package main

func avg(sumFunc func(...float64) float64, nums ...float64) float64 {
	return sumFunc(nums...) / float64(len(nums))
}

func sum(nums ...float64) float64 {
	var total float64

	for i := range nums {
		total += nums[i]
	}

	return total
}
```

```go
func BenchmarkAvg(b *testing.B) {
	nums := []float64{1, 2, 3, 4, 5}

	for i := 0; i < b.N; i++ {
		avg(sum, nums...)
	}
}
```

Benchmark results for this inital version:

```bash
go test -bench=BenchmarkAvg
```

    BenchmarkAvg-8  446951605   2.651 ns/op     0 B/op      0 allocs/op

Version without the compiler's inlining optimization:

```go
//go:noinline
func sumWithoutInlining(nums ...float64) float64 {
	var total float64

	for i := range nums {
		total += nums[i]
	}

	return total
}
```

```go
func BenchmarkAvgWithoutInlining(b *testing.B) {
	nums := []float64{1, 2, 3, 4, 5}

	for i := 0; i < b.N; i++ {
		avg(sumWithoutInlining, nums...)
	}
}
```

```bash
go test -bench=BenchmarkAvgWithoutInlining
```

    BenchmarkAvgWithoutInlining-8   214392748   5.470 ns/op     0 B/op      0 allocs/op

Inlining improved almost 2x the speed of the program. So it's always a good idea writing functions that can be inlined by the compiler. In general inlining works best with ***small functions***, but to understand if a function it's being inlined you need to ask the compiler.

```bash
go build -gcflags '-m=2' .
```

    ./main.go:3:6: can inline sum with cost 16 as: func(...float64) float64 { total = <nil>; for loop; return total }
    ./main.go:13:6: can inline avg with cost 24 as: func(...float64) float64 { return sum(nums...) / float64(len(nums)) }
    ./main.go:14:12: inlining call to sum
    ./main.go:18:6: cannot inline sumWithoutInlining: marked go:noinline
    ./main.go:28:6: can inline avgWithoutInlining with cost 65 as: func(...float64) float64 { return sumWithoutInlining(nums...) / float64(len(nums)) }

The functions `avg` and `sum` were inlined, but `sumWithoutInlining` wasn't, since it was forced to not be inlined. But if the function changes to:

```go
func sumWithHighInliningCost(nums ...float64) float64 {
	var total float64
	numsLen := len(nums) - 1

	total += nums[numsLen]
	total -= nums[numsLen]
	total += nums[numsLen]
	total -= nums[numsLen]
	total += nums[numsLen]
	numsLen--

	total += nums[numsLen]
	total -= nums[numsLen]
	total += nums[numsLen]
	total -= nums[numsLen]
	total += nums[numsLen]
	numsLen--

	total += nums[numsLen]
	total -= nums[numsLen]
	total += nums[numsLen]
	total -= nums[numsLen]
	total += nums[numsLen]
	numsLen--

	total += nums[numsLen]
	total -= nums[numsLen]
	total += nums[numsLen]
	total -= nums[numsLen]
	total += nums[numsLen]
	numsLen--

	total += nums[numsLen]
	total -= nums[numsLen]
	total += nums[numsLen]
	total -= nums[numsLen]
	total += nums[numsLen]
	numsLen--

	return total
}
```

```bash
go build -gcflags '-m=2' .
```

    ./main.go:3:6: can inline avg with cost 65 as: func(func(...float64) float64, ...float64) float64 { return sumFunc(nums...) / float64(len(nums)) }
    ./main.go:7:6: can inline sum with cost 16 as: func(...float64) float64 { total = <nil>; for loop; return total }
    ./main.go:18:6: cannot inline sumWithoutInlining: marked go:noinline
    ./main.go:28:6: cannot inline sumWithHighInliningCost: function too complex: cost 154 exceeds budget 80

We'll get a high inlining cost. This version of the function (although useless) shows how a big function can produce a big inlining cost, a thus, making a program slower.

Version with high inlining cost:

```go
func BenchmarkAvgWithHighInliningCost(b *testing.B) {
	nums := []float64{1, 2, 3, 4, 5}

	for i := 0; i < b.N; i++ {
		avg(sumWithHighInliningCost, nums...)
	}
}
```

```bash
go test -bench=BenchmarkAvgWithHighInliningCost
```

    BenchmarkAvgWithHighInliningCost-8  148438526   8.037 ns/op     0 B/op      0 allocs/op

Even worse than the non-inlined forced version (probably because this version does more operations inside it).