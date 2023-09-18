# Go Optimizations

## Topics

- Escape analysis and memory sharing ([here](./escape_analysis/README.md))
- Inlining ([here](./inlining/README.md))
- Gargabe Collector ([here](./garbage_collector/README.md))
- Struct Field Ordering ([here](./struct_field_order/README.md))

## WIP

### GOGC, GOMEMLIMIT and Pacing

The GC Percentage value is used to set when new collections will occur. It can be changed with the env `GOGC` and its default value is `100`.

It's not a good idea to change `GOGC` to improve performance by reducing latency produced by GC runs. This is beacause the GC can be triggered before the goal limit set by `GOGC` is reached, a pacing algorithm is used to determine when a collection will start, by keeping track the stress of the running application. So the same application running with two different `GOGC` values could have the same latency produced by GC runs when high loads are present. 

***To solve latency issues related to GC runs the best solution is to reduce allocations.***

Nontheless, in some situations can be useful to set `GOGC` when there's already a good understanding about the program resource usage. 

## Remainders

### Profiling

Memory Profiling

```bash
go test -bench . -benchmem -memprofile memprof.out -gcflags -m=1
```

CPU Profiling

```bash
go test -bench . -benchmem -cpuprofile cpuprof.out
```

Visualize Profiling Data

```bash
go tool pprof -http :8080 prof.out
```