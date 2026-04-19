# Unmanaged C# memory

The C# unmanaged memory layer allows you to access the native memory layer to fine-tune memory allocations, with the convenience of writing C# code.

You can use the `Unity.Collections` namespace (including [`NativeArray`](../ScriptReference/Unity.Collections.NativeArray_1.md)) in the Unity core API, and the data structures in the [Unity Collections package](https://docs.unity3d.com/Packages/com.unity.collections@latest) to access C# unmanaged memory. If you use Unity’s [job system](job-system.md), or [Burst](http://docs.unity3d.com/Packages/com.unity.burst@latest), you must use C# unmanaged memory.

## Native container memory allocators overview

Unity’s native code libraries track native memory usage through a system of memory labels, areas, and roots that are managed by Unity’s native memory manager. The [Profiler](Profiler.md) uses the native memory manager to monitor native memory usage. The asset garbage collector also tracks types that inherit from `UnityEngine.Object` so that on calls to `Resources.UnloadUnusedAssets` it can clean up memory assigned to assets that are no longer referenced.

Unity’s native memory manager uses memory labels to categorize memory usage and to choose which [`Allocator`](../ScriptReference/Unity.Collections.Allocator.md) allocates the memory. Which allocator the memory manager uses for each memory label varies between platforms. You can also create your own named memory labels with the [`MemoryLabel` API](../ScriptReference/Unity.Collections.MemoryLabel.md). **Important:** The `MemoryLabel` API isn’t compatible with the Collections package.

The memory manager uses areas to categorize memory usage for profiling purposes. It uses roots to track related memory allocations under one root. For example, the native memory for each type that inherits from `UnityEngine.Object` has a root in the area called `Object`. When the object is deleted, for example via [`Resources.UnloadUnusedAssets`](../ScriptReference/Resources.UnloadUnusedAssets.md), all allocations associated with that object’s root are also freed.

You can use the [Memory Profiler package](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest/) to check memory usage ordered by memory labels, areas, and roots. This includes any custom labels you’ve defined using MemoryLabel. You can also discover the underlying allocator that allocated the memory through region names.

## Allocator types

You can use certain native [Allocator types](../ScriptReference/Unity.Collections.Allocator.md) with native container types in the `Unity.Collections` namespace. The allocator types available in this way are:

* [`Temp`](../ScriptReference/Unity.Collections.Allocator.Temp.md): A temporary allocator type for short-lived allocations.
* [`TempJob`](../ScriptReference/Unity.Collections.Allocator.TempJob.md): A temporary allocator type for longer-lived allocations passed from job to job.
* [`Persistent`](../ScriptReference/Unity.Collections.Allocator.Persistent.md): An allocator without any temporal restrictions.
* [`AudioKernel`](../ScriptReference/Unity.Collections.Allocator.AudioKernel.md): An allocator for DSP audio kernel allocations.
* [`Domain`](../ScriptReference/Unity.Collections.Allocator.Domain.md): An allocator that lasts the lifetime of the domain.

To use these types, pass an [`Allocator`](../ScriptReference/Unity.Collections.Allocator.md) enum value to select the native allocator type’s memory label to use. Alternatively, for low-level allocations where you need custom categorization, you can create and pass a [`MemoryLabel`](../ScriptReference/Unity.Collections.MemoryLabel.md) instance directly to allocation functions. If a chosen `Temp` allocator is full, Unity chooses a fallback allocator instead (usually `TempJob`, or Temp Overflow).

### Temp

`Temp` is Unity’s temporary allocator type. It’s designed for small short lived allocations that can’t cross a frame boundary, such as temporary allocations made within a single job. The `Temp` type is a stack type allocator for small amounts of memory which Unity allocates from the main memory pool.

There are multiple `Temp` allocators, and each thread that needs an allocator has one. Unity recycles `Temp` allocators each frame ensuring a clean memory usage pattern. However, if memory spills over into the fallback allocators, their memory might get fragmented.

You can use [`Profiler.GetTempAllocatorSize`](../ScriptReference/Profiling.Profiler.GetTempAllocatorSize.md) and [`Profiler.SetTempAllocatorRequestedSize`](../ScriptReference/Profiling.Profiler.SetTempAllocatorRequestedSize.md) to get and set the size for `Temp` allocators.

The sizes for the `Temp` allocator default to:

* Editor: 16 MB main thread, 256 KB worker threads.
* Players: 4 MB main thread, 256 KB worker threads.

If these sizes are exceeded, allocations on the main thread fall back to the `TempJob` allocator, and on other threads they fall back to the Temp Overflow allocator. For more information on temporary native allocations, refer to the [Thread storage local (TLS) allocator](performance-native-allocators.md#tls) documentation.

### TempJob

This type uses a linear style allocator which Unity stores in a `TempJob` area that it allocates from main memory. It’s designed for larger temporary allocations that are passed from job to job, and those that might carry data between frames. If the allocations aren’t disposed of within 4 frames of their creation, a warning is displayed in the Unity Editor.

**Note**: If the memory available is exhausted, Unity makes the overspill allocations in main memory (malloc). Most platforms can have up to 64 MB of `TempJob` memory, but on smaller memory systems this limit can be as low as 16 MB.

If this allocator is full, allocations fall back to the Temp Job Overflow Allocator. For more information on temporary job allocations, refer to the [Thread safe linear allocator](performance-native-allocators.md#thread-safe-linear) documentation.

### Persistent

This allocator allocates memory with no time restrictions, is only limited by the amount of free memory on a device, and therefore has no fallback. It can exist for as long as you need it, but this also means that only the [`DisposeSentinel`](../ScriptReference/Unity.Collections.LowLevel.Unsafe.DisposeSentinel.md) class warns you if you don’t free the allocated memory again, and then only if your handle to the memory was garbage collected. Allocating persistent memory is slower than allocating temporary memory.

## Additional resources

* [Collections package](https://docs.unity3d.com/Packages/com.unity.collections@latest)
* [Memory in Unity introduction](performance-memory-overview.md)