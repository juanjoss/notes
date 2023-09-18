# Garbage Collector

The Go's garbage collector uses a mark and sweep algorithm. This is how the algorithm works at a high level:

- Marking
   
   1. Setup Phase:
      1. Stop all running goroutines.
      2. Establish a write barrier.
   
   2. Marking Phase:
   		The GC runs concurrently alongside other goroutines to mark objects in heap memory.
	
   3. Termination Phase:
      1. Stop all running goroutines.
      2. Remove the write barrier.
   
- Sweeping
	
   The sweeping occurs when a goroutine requests to allocate new memory, by removing marked objects.

## Example

See [main.go](./main.go) for a brief example on the garbage collector.

The example shows how the garbage collector cleans `arr`, after it was created, filled and its reference was lost. By gathering runtime memory stats, we can see the memory increase before and after `arr` was created and filled, and how after the GC run memory was freed.