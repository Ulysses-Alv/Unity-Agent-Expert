# Garbage collection modes

Unity’s [garbage collector](performance-garbage-collector.md) has the following modes:

* [Incremental garbage collection](#incremental-garbage-collection): Enabled by default. This mode spreads out the process of garbage collection over multiple frames.
* [Non-incremental garbage collection mode](#non-incremental-garbage-collection): If you disable the **Incremental GC Player Setting** (**Project Settings** > **Player** > **Configuration**), the garbage collector stops running your application to inspect and process objects on the heap. This mode is also known as stop-the-world garbage collection.
* [Manual garbage collection](#manual-garbage-collection): Use the [`GarbageCollector.GCMode`](../ScriptReference/Scripting.GarbageCollector.GCMode.md) API to take full control of when Unity runs the garbage collector.

For information on how to use each of the garbage collection modes, refer to [Configuring garbage collection](performance-disabling-garbage-collection.md).

## Incremental garbage collection

Incremental garbage collection spreads out the process of garbage collection over multiple frames. This is the default garbage collection behavior in Unity.

Unity’s garbage collector uses the [Boehm-Demers-Weiser garbage collector](https://www.hboehm.info/gc/). By default, Unity uses it in incremental mode, which means that the garbage collector splits up its workload over multiple frames and makes shorter interruptions to your application’s execution.

In [non-incremental mode](#non-incremental-garbage-collection), Unity stops the main CPU thread ([stop-the-world garbage collection](https://en.wikipedia.org/wiki/Tracing_garbage_collection#Stop-the-world_vs._incremental_vs._concurrent)) in one long interruption to process all objects on the managed heap.

Incremental mode doesn’t make garbage collection faster, but because it distributes the workload over multiple frames, performance spikes related to garbage collection are reduced. These interruptions are called **GC spikes** because they appear as large spikes in the [Profiler window’s](ProfilerWindow.md) frame time graph.

**Important:** The web platform doesn’t support incremental garbage collection. For more information, refer to [Memory in web garbage collection considerations](webgl-memory.md#garbagecollection).

### Incremental garbage collection example

The following screenshots from the Profiler illustrate how incremental garbage collection reduces frame rate problems:

![Profiling session with Incremental GC enabled](../uploads/Main/profiler-incremental-gc.png)

Profiling session with Incremental GC enabled

![Profiling session with Incremental GC disabled](../uploads/Main/profiler-incremental-gc-disabled.png)

Profiling session with Incremental GC disabled

In the top profiling session, the [Incremental GC setting](performance-disabling-garbage-collection.md) is enabled. The application has a consistent 60 **fps** frame rate, because the garbage collector distributes the garbage collection operation over several frames. The garbage collector uses a small time slice of each frame, indicated by the darker green fringe just above the yellow VSync trace.

The bottom profiling session has the **Incremental GC** setting disabled, and there’s a GC spike visible. The GC spike interrupts the consistent 60 fps frame rate, and pushes the frame in which garbage collection happens over the 16 ms limit required to maintain 60 fps.

If your application uses [VSync](../ScriptReference/QualitySettings-vSyncCount.md) or [`Application.targetFrameRate`](../ScriptReference/Application-targetFrameRate.md), Unity adjusts the time it allocates to garbage collection based on the remaining available frame time. This way, Unity can run the garbage collector in the time it spends waiting, and can carry out garbage collection with a minimal performance impact.

**Note:** If you set the **VSync Count** to anything other than **Don’t Sync** (in your project’s [Quality settings](class-QualitySettings.md) or with the [`Application.VSync`](../ScriptReference/QualitySettings-vSyncCount.md) property), or you enable the [`Application.targetFrameRate`](../ScriptReference/Application-targetFrameRate.md) property, Unity automatically uses any idle time left at the end of a given frame for incremental garbage collection.

To get more precise control over incremental garbage collection behavior, you can use the [`Scripting.GarbageCollector`](../ScriptReference/Scripting.GarbageCollector.md) class. For example, if you don’t want to use VSync or a target frame rate, you can calculate the amount of time available before the end of a frame and provide that time to the garbage collector.

For more information, refer to [Configuring garbage collection](performance-disabling-garbage-collection.md).

## Non-incremental garbage collection mode

Incremental garbage collection might be problematic for your application, because when the garbage collector divides its work in this mode, it also divides the marking phase. The marking phase is the phase in which the garbage collector scans all managed objects to determine which objects are still in use, and which objects it can clean up.

Dividing up the marking phase works well when most of the references between objects don’t change between slices of work. However, when an object reference changes, the garbage collector must scan those objects again in the next iteration. This means that too many changes might overwhelm the incremental garbage collector and create a situation where the marking phase never finishes because it always has more work to do. If this happens, the garbage collector falls back to doing a full, non-incremental collection.

When Unity uses incremental garbage collection, it generates additional code (known as write barriers) to inform the garbage collector when it needs to scan an object whenever a reference changes. This adds some overhead when [changing references](performance-managed-memory-introduction.md#managed-heap), which has a performance impact in managed code.

Therefore, there are some situations where you might want to disable incremental garbage collection to improve the performance of your application. When you [disable incremental mode](performance-disabling-garbage-collection.md), the garbage collector must examine the entire heap when it performs a collection pass.

This is known as **stop-the-world** garbage collection, because whenever the garbage collector runs, it stops the main CPU thread. It only resumes execution once it has processed all objects on the managed heap, which might lead to GC spikes affecting the performance of your application. This delay might last for hundreds of milliseconds, depending on how many allocations the garbage collector needs to process, and the platform that your application is running on.

The garbage collector is also non-compacting, which means that Unity doesn’t redistribute any objects in memory to close the gaps between objects.

Non-incremental garbage collection is problematic for real-time applications such as games, because it’s difficult for your application to sustain the consistent frame rate that smooth animation requires when the garbage collector suspends your application’s execution.

To disable incremental garbage collection, refer to [Configuring garbage collection](performance-disabling-garbage-collection.md).

## Manual garbage collection

You can use the [`GarbageCollector.GCMode`](../ScriptReference/Scripting.GarbageCollector.GCMode.md) API to disable garbage collection at runtime. This prevents CPU spikes, but the memory usage of your application never decreases, because the garbage collector doesn’t collect objects that no longer have any references.

Disabling the garbage collector requires careful memory management, or the managed heap continuously expands until your application runs out of memory, and the operating system shuts it down.

It’s best practice to disable garbage collection during short, performance-critical parts of your application, when you’re able to calculate and control how much memory to allocate. After the performance-critical section is over, re-enable the garbage collector. [Profile your project](profiling-collect-data-introduction.md) to ensure that you don’t trigger additional managed allocations which might cause the managed heap to get too big.

It’s best practice to only disable the garbage collector for long-lived allocations. For example, you might want to allocate all required memory for a level of your game before it loads, and then disable the garbage collector to avoid performance overhead during the level. After the level is completed and all memory is released, you can then enable the garbage collector again and use `System.GC.Collect` to reclaim memory before loading the next level.

For more details on how to enable and disable garbage collection at runtime, refer to [Configuring garbage collection](performance-disabling-garbage-collection.md).

## Additional resources

* [Configuring garbage collection](performance-disabling-garbage-collection.md)
* [Garbage collector introduction](performance-garbage-collector.md)
* [`GarbageCollector.GCMode` API reference](../ScriptReference/Scripting.GarbageCollector.GCMode.md)
* [`Scripting.GarbageCollector` API reference](../ScriptReference/Scripting.GarbageCollector.md)