# Dual thread allocator example

## Prerequisites

The examples use the memory usage reports that Unity writes to the log when you close the Player or Unity Editor. To create these reports, use the `-log-memory-performance-stats` command line argument. To find your project’s log files, follow the instructions on [the log files page](log-files.md).

## Customize block size

You can customize the block sizes of the two dynamic heap allocators:

![Main Allocator, with a custom value for Shared Thread Block Size](../uploads/Main/Main_Allocator.png)

Main Allocator, with a custom value for Shared Thread Block Size

## Usage report example

The usage report contains information for all three parts of the allocator. For example:

```
[ALLOC_DEFAULT] Dual Thread Allocator
  Peak main deferred allocation count 135
    [ALLOC_BUCKET]
      Large Block size 4.0 MB
      Used Block count 1
      Peak Allocated bytes 3.3 MB
    [ALLOC_DEFAULT_MAIN]
      Peak usage frame count: [16.0 MB-32.0 MB]: 8283 frames, [32.0 MB-64.0 MB]: 1 frames
      Requested Block Size 16.0 MB
      Peak Block count 2
      Peak Allocated memory 53.3 MB
      Peak Large allocation bytes 40.2 MB
    [ALLOC_DEFAULT_THREAD]
      Peak usage frame count: [64.0 MB-128.0 MB]: 8284 frames
      Requested Block Size 16.0 MB
      Peak Block count 2
      Peak Allocated memory 78.3 MB
      Peak Large allocation bytes 47.3 MB
```

**Note**: The **Peak main deferred allocation count** is the number of items in a deletion queue. The main thread must delete any allocation it made. If another thread deletes an allocation, Unity adds that allocation to a queue. The allocation waits in the queue for the main thread to delete it. It’s then counted as a deferred allocation.

## Additional resources

* [Native memory allocators](performance-native-allocators.md)
* [Native memory allocator examples](performance-native-memory-allocator-examples.md)