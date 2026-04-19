# Thread-safe linear allocator example

## Prerequisites

The examples use the memory usage reports that Unity writes to the log when you close the Player or Unity Editor. To create these reports, use the `-log-memory-performance-stats` command line argument. To find your project’s log files, follow the instructions on [the log files page](log-files.md).

## Allocator overview

This allocator allocates blocks of memory, then linearly allocates memory within those blocks. Unity holds the available blocks in a pool. When one block is full, the allocator fetches a new block from the pool. When the allocator no longer needs the memory in a block, it clears the block, and the block returns to the pool of available blocks.

You can customize the block size. The allocator allocates up to 64 blocks, as needed.

![Default value for Fast Thread Shared Temporary Allocators for the Editor](../uploads/Main/Fast_Thread.png)

Default value for Fast Thread Shared Temporary Allocators for the Editor

## Example usage report

If all blocks are in use, or an allocation is too big for a block, the allocation falls back to the main heap allocator, which is much slower than the job allocator. A few overflow allocations are fine: 1 to 10 in a frame, or a few hundred, especially during load. If the overflow count grows with every frame, you can increase the block size to avoid fallback allocations. However, if you increase the block size too much (for example, to match peak use in events such as scene loading), you might leave a lot of memory unavailable during play.

For example:

```
[ALLOC_TEMP_JOB_4_FRAMES (JobTemp)]
  Initial Block Size 0.5 MB
  Used Block Count 64
  Overflow Count (too large) 0
  Overflow Count (full) 50408
```

In this example usage report, the 0.5 MB block size was too small to accommodate the job memory that the application needed, and the full allocator caused many allocations to overflow.

To check whether your build’s frame overflow is enough, run it for a short time and then for a longer time. If the overflow count remains steady, the overflow is a high watermark that occurs during load. If the overflow count increases with a longer run, the build is processing a per-frame overflow. In both cases, you can increase the block size to reduce the overflow, but the overflow is less critical during load than per frame.

## Additional resources

* [Native memory allocators](performance-native-allocators.md)
* [Native memory allocator examples](performance-native-memory-allocator-examples.md)