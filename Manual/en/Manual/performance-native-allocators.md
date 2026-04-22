# Native memory allocators

Unity provides several allocator types, where each type represents a different relative prioritization of the following key factors:

* The performance impact of creating, accessing, and freeing memory.
* The overhead of tracking memory usage.
* Thread access safety.

If an application has a lot of spare memory, it can use faster, memory-heavy allocators when it loads **scenes** and frames. However, the more an application’s memory is limited, the more efficiently it must use the memory available, which often means using slower allocators.

To help you get the best performance for different projects, you can [customize Unity’s allocators](memory-allocator-customization.md) to fit the size and requirements of each application.

## Native memory allocator types

Unity uses different allocator types, which all use different algorithms to organize memory. Unity’s memory manager uses these allocator types in different areas to organize the memory in your application effectively. You can adjust the size of each allocation for each area, either through the [Player settings](class-PlayerSettings.md) window in the Unity Editor, or through the command line.

The different types of allocators are as follows:

* [**Dual thread:**](#dual-thread) A wrapper that combines two dynamic heap allocators, and a bucket allocator for small allocations. The dual thread allocator redirects allocations based on size and thread ID.
* [**Dynamic heap**:](#dynamic-heap) A main heap allocator, which uses the [Two Level Segregated Fit (TLSF) algorithm](http://www.gii.upv.es/tlsf/).
* [**Bucket:**](#bucket) A fixed size lock-free allocator for small allocations before they go on the main heap.
* [**Thread local storage (TLS) stack:**](#tls) A last-in, first-out (LIFO) stack allocator, for temporary allocations on threads.
* [**Threadsafe linear:**](#thread-safe-linear) A first-in, first-out (FIFO) round-robin allocator for worker threads.

## Allocator usage

Each allocator type has a different algorithm for fitting allocations into blocks of memory, and is therefore useful for different types of allocations. The important difference between allocations is usually persistence, or allocation lifespan, which determines where to put an allocation. For example, a long-lived (persistent) allocation goes to the heap and bucket allocators, while short-lived allocations go to the thread-safe linear and thread local storage allocators.

When an allocation is made somewhere in code, Unity does the following:

* Checks if the code is currently running on the main thread:
  + If the code is running on the main thread, it uses a main thread dynamic heap allocator.
  + If the code isn’t running on the main thread, the allocation blocks to get access to a threaded allocator.
* Checks the size of the allocator:
  + If the size of the allocation is small enough to fit into the bucket allocator, and there’s space for it, it allocates the memory in the bucket allocator.
  + If the allocation size is too big for the bucket allocator, or there’s no room, it uses the dynamic heap allocator.

## Dual thread allocator

The dual thread allocator is used for longer lived allocations and those that need to be passed between threads, such as graphics allocations that need passing from the main thread to the render thread.

The allocator contains two dynamic heap allocators: a lock-free one for the main thread, and one for all other threads, which locks on allocation and deallocation. It also contains a bucket allocator for small, long-lived allocations.

Unity creates several dual thread allocators by default, including the Main and Gfx dual thread allocators. File cache and Profiler allocations are always [dynamic heap allocations](#dynamic-heap).

For an example of dual thread allocator usage refer to [Dual thread allocator example](performance-dual-thread-allocator.md).

## Dynamic heap allocator

The dynamic heap allocator is the main heap allocator. It uses a [Two Level Segregated Fit (TLSF) algorithm](http://www.gii.upv.es/tlsf/) to manage its memory. It’s used for large, long-lived allocations and for allocations that can’t fit in the bucket allocator. The allocations are divided into memory blocks of a given size, which are requested as needed. An allocation must be smaller than half a block. An allocation of half a block or more is too large for the dynamic heap allocator and in such cases Unity uses the virtual memory API to make the allocation instead.

For an example of the dynamic heap allocator usage, refer to [Dynamic heap allocator example](performance-dynamic-heap-allocator.md).

## Bucket allocator

The bucket allocator is a fast lock-free allocator that performs small allocations. The dual thread allocator uses the bucket allocator for small, long-lived allocations. These allocations prevent small ones from fragmenting the dynamic heap allocators.

The bucket allocator reserves blocks of memory for allocations. Each block is divided into subsections of 16 KB. This isn’t configurable, and doesn’t appear in the user interface. Each subsection is divided into allocations. The allocation size is a multiple of a configured fixed size, called granularity. The bucket allocator has a maximum number of blocks of a given size.

**Important:** The bucket allocator behaves differently between [development and release builds](building-introduction.md), due to an additional 40 byte header in the **development build**.

For an example of the bucket allocator usage, refer to [Bucket allocator example](performance-bucket-allocator.md).

## Thread local storage stack allocator

The Thread local storage (TLS) stack allocator is for per thread short-lived allocations. It’s the fastest allocator. It consists of last-in-first-out (LIFO) memory stacks that set each stack pointer to 0 to clear memory. The TLS allocator is mostly used for [`Allocator.Temp`](../ScriptReference/Unity.Collections.Allocator.Temp.md) allocations, which can also come from user-written code.

The default block size for the thread local storage (TLS) stack allocator is 4 MB for platforms and 16 MB for the Unity Editor. You can [customize these values](memory-allocator-customization.md). If the allocator usage exceeds the configured block size, Unity increases the block size. The limit for this increase is twice the original size.

The TLS stack allocator’s memory allocations ideally live for less than a frame and have a similar lifespan as their encompassing scope.

Because TLS is a stack allocator, interleaved allocation and deallocation patterns (for example, allocate a, allocate b, deallocate a, allocate c, deallocate b, etc.) can fragment it and stop the allocator from resetting the stack point down-stack. This might cause it to fill up entirely while theoretically having free space available.

If a thread’s stack allocator is full, allocations fall back to the [threadsafe linear allocator](#thread-safe-linear).

For an example of the TLS stack allocator usage, refer to [Thread local storage stack allocator example](performance-tls-stack-allocator.md).

## Thread-safe linear allocator

The thread-safe linear allocator is a fallback for the TLS stack allocator, and it allocates up to 16 blocks of memory. It uses a first-in-first-out (FIFO) algorithm for lock-free allocations of work buffers for jobs. When complete, the jobs dispose of these buffers. It’s mostly used for [`Allocator.TempJob`](../ScriptReference/Unity.Collections.Allocator.TempJob.md) allocations, which have a 4 frame lifetime limit, and they can also come from user-written code.

This allocator allocates blocks of memory, then linearly allocates memory within those blocks. Available blocks are held in a pool. When one block is full, the allocator fetches a new block from the pool. When the allocator no longer needs the memory in a block, it clears the block, and the block returns to the pool of available blocks. It’s important to clear allocations quickly to make blocks available again, so a job should not stay allocated for more than a few frames.

If all blocks of memory are full, the allocations fall back to the main [dual thread allocator](#dual-thread), which is a much slower allocator.

For an example of thread-safe linear allocator usage refer to [Thread safe linear allocator example](performance-threadsafe-linear-allocator.md).

## Additional resources

* [Native memory allocator examples](performance-native-memory-allocator-examples.md)
* [Customizing native memory allocators](memory-allocator-customization.md)
* [Native memory allocators reference](performance-native-memory-allocator-reference.md)