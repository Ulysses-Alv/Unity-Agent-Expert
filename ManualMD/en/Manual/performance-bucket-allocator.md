# Bucket allocator example

## Prerequisites

The examples use the memory usage reports that are written to the log when you close the Player or Unity Editor. To create these reports, use the `-log-memory-performance-stats` command line argument. To find your project’s log files, follow the instructions on [the log files page](log-files.md).

## Example configuration

The bucket allocator reserves blocks of memory for allocations. Each block is divided into subsections of 16 KB. The following example configuration demonstrates the process of reserving blocks for allocations:

![Shared Bucket Allocator for Windows, macOS, and Linux Player](../uploads/Main/Shared_Bucket.png)

Shared Bucket Allocator for Windows, macOS, and Linux Player

In this setup, the total block size (**Bucket Allocator Block Size**) is 4 MB, and the granularity of allocations (**Bucket Allocator Granularity**) is 16 bytes. The first allocation size it creates buckets for is 16 bytes, the second is 32 bytes (2 \* 16), then 48 bytes, 64 bytes, 80 bytes, 96 bytes, 112 bytes, and 128 bytes, for a total of eight buckets (**Bucket Allocator BucketCount**).

Each subsection contains a different number of buckets. To calculate the number of buckets in a subsection, divide the subsection size (16 KB) by the granularity size. For example:

* When the allocation granularity is 64 bytes, 256 buckets fit in a subsection.
* When the allocation granularity is 16 bytes, 1,024 buckets fit in a subsection.

## Development and release build comparison

Bucket allocators produce different usage reports for a [development build and a release build](building-introduction.md) because in a **development build** each allocation has an additional 40 byte header. The following diagram demonstrates the difference between development and release builds for 16 byte and 64 byte allocations:

![Development and Release builds comparison](../uploads/Main/memory_allocation.png)

Development and Release builds comparison

The header is the reason the allocator reports being full after allocating only 2 MB of its 4 MB:

```
[ALLOC_BUCKET]
      Large Block size 4.0 MB
      Used Block count 1
      Peak Allocated bytes 2.0 MB
      Failed Allocations. Bucket layout:
        16B: 64 Subsections = 18724 buckets. Failed count: 3889
        32B: 17 Subsections = 3868 buckets. Failed count: 169583
        48B: 31 Subsections = 5771 buckets. Failed count: 39674
        64B: 28 Subsections = 4411 buckets. Failed count: 9981
        80B: 17 Subsections = 2321 buckets. Failed count: 14299
        96B: 6 Subsections = 722 buckets. Failed count: 9384
        112B: 44 Subsections = 4742 buckets. Failed count: 5909
        128B: 49 Subsections = 4778 buckets. Failed count: 8715
```

In a release build for the same project, the allocator block size is enough:

```
[ALLOC_BUCKET]
      Large Block size 4.0 MB
      Used Block count 1
      Peak Allocated bytes 3.3 MB
```

If the bucket allocator is full, the allocation falls back to another allocator. The usage report displays usage statistics, including how many allocations failed. If the report displays a fail count that increases linearly, it’s likely that the failed allocations happen when calculating the frames, not the load. Fallback allocations aren’t a problem for a **scene** load, but they can impact performance if they happen when calculating frames.

To prevent these fallback allocations, increase the block size, and limit the new block size to match the frames’ peak usage, rather than the scene load peak usage. This prevents the block from becoming so large that it reserves a lot of memory which is then not available at runtime.

**Tip**: [The Profiler](Profiler.md) allocators share an instance of a bucket allocator. You can customize this shared instance in the [Profiler Shared Bucket Allocator setting](memory-allocator-customization.md).

## Additional resources

* [Native memory allocators](performance-native-allocators.md)
* [Native memory allocator examples](performance-native-memory-allocator-examples.md)