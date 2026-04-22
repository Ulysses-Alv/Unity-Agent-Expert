# Dynamic heap allocator example

## Prerequisites

The examples use the memory usage reports that Unity writes to the log when you close the Player or Unity Editor. To create these reports, use the `-log-memory-performance-stats` command line argument. To find your project’s log files, follow the instructions on [the log files page](log-files.md).

## Example usage report

An example usage report for the dynamic heap allocator is as follows:

```
[ALLOC_DEFAULT_MAIN]
Peak usage frame count: [16.0 MB-32.0 MB]: 497 frames, [32.0 MB-64.0 MB]: 1 frames
Requested Block Size 16.0 MB
Peak Block count 2
Peak Allocated memory 54.2 MB
Peak Large allocation bytes 40.2 MB
```

In this example, the [Two Level Segregated Fit (TLSF) block size](performance-native-allocators.md#dynamic-heap) is 16 MB, and Unity has allocated two blocks. The peak usage of the allocator was 54.2 MB. Of those 52.4 MB, 40.2 MB weren’t allocated in the TLSF block, and instead fell back to virtual memory. Most frames had 16 to 32 MB of allocated memory, while one frame, which was likely the loading frame, peaked at 32 to 64 MB of memory.

If you increased the block size the large allocation stays in the dynamic heap rather than fall back into virtual memory. However, that block size might lead to memory waste, because the blocks might not be fully used.

The type tree and file cache allocators use dynamic heap allocation. To save the memory blocks they might otherwise use under this algorithm, you can set the type tree block size and file cache block size to 0. Allocations that would have used typetree and cache fall back to the main allocator instead.

However, reducing the block size to 0 has a risk of increased native memory fragmentation. Refer to [Customize allocators](memory-allocator-customization.md) for how to set these block sizes.

## Additional resources

* [Native memory allocators](performance-native-allocators.md)
* [Native memory allocator examples](performance-native-memory-allocator-examples.md)