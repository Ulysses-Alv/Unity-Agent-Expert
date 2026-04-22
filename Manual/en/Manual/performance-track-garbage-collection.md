# Tracking garbage collection allocations

Unity has the following tools to keep track of memory allocations:

* [The CPU Usage Profiler module](profiler-cpu.md): Provides details of the garbage allocation per frame.
* [The Memory Profiler module](ProfilerMemory.md): Provides high-level memory usage frame by frame.
* [The Memory Profiler package](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest): A separate Unity package which provides detailed information about memory usage at specific points in time during in your application.

## Tracking allocations with the CPU Usage module

To track garbage collection allocations with the [CPU Usage Profiler module](profiler-cpu.md), perform the following steps:

1. Open the [Profiler window](ProfilerWindow.md) (**Window** > **Analysis** > **Profiler**).
2. [Collect some performance data](profiling-collect-data-introduction.md).
3. Select a frame in the CPU Usage Profiler module.
4. In the module details pane, open the [Hierarchy view](ProfilerCPU.md), and inspect the **GC.Alloc** column.

The **GC.Alloc** column displays the number of bytes that Unity allocated on the managed heap for the selected frame and thread.

Allocations can occur on all threads so you might want to use the thread selection drop down to check out other threads than the main thread. You can also use the [Timeline view](ProfilerCPU.md) and keep an eye out for the `GC.Alloc` samples, which are colored bright magenta.

**Tip:** Use the [Call Stacks mode](profiler-cpu-navigating.md#call-stacks) to enable the full call stack traces for `GC.Alloc` samples. Call stacks give you precise details of where the garbage collector made the allocation without needing to use [Deep Profiling](profiler-deep-profiling.md) mode, which might impact performance.

## Tracking allocations with the Memory module

You can use the [Memory Profiler module](profiler-memory-introduction.md) to track garbage collection memory allocations, but it only provides a high-level overview of where Unity allocated memory. For detailed information on memory usage in your application, use the [Memory Profiler package](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest).

To track garbage collection memory allocations with the Memory Profiler module, perform the following steps:

1. Open the [Profiler window](ProfilerWindow.md) (**Window** > **Analysis** > **Profiler**).
2. [Collect some performance data](profiling-collect-data-introduction.md).
3. Select a frame in the Memory Profiler module.
4. In the module details pane, inspect the **GC allocated in frame** statistic.

The **Managed Heap** statistic displays the amount of memory that the garbage collector managed, and it includes memory that Unity might have allocated and reused in subsequent frames. This means that the sum of the `GC.Alloc` samples over all frames doesn’t total how much the managed memory grew in that time.

The **GC allocated in frame** statistic diplays the amount of memory that was allocated in this frame. This amount might be higher than the CPU Usage module displays in the **GC.Alloc** column of the Hierarchy view. This is because the `GC.Alloc` statistic includes allocations made across all threads and also within [Editor only samples](profiler-play-edit-samples.md), or as part of the `EditorLoop`. The Hierarchy view only displays one thread at a time and hides the allocations made in Editor only samples by default and collapses the code running as part of the `EditorLoop` unless the [profiler targets the Editor](profiling-edit-mode.md) and not Play mode.

**Important:** To get the most accurate information, [profile your application](profiler-profiling-applications.md) on a **development build** on the target platform or device you want to build to. The Unity Editor works in a different way to a build, and this affects the profiling data. For example, the `GetComponent` method always allocates memory when it’s executed in the Editor, but not in a built project.

## Additional resources

* [Managed memory introduction](performance-managed-memory-introduction.md)
* [Garbage collector introduction](performance-garbage-collector.md)
* [Collect performance data](profiler-collect-data)
* [Memory Profiler package](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest)