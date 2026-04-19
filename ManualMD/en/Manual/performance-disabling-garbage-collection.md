# Configuring garbage collection

By default, Unity carries out [garbage collection incrementally](performance-incremental-garbage-collection.md). If you want more control over when and how garbage collection happens, you have the following options:

* [Disable incremental garbage collection](#disable-incremental-garbage-collection)
* [Use the APIs in the `GarbageCollector` class](#garbagecollector-api) to manually control the garbage collector, or disable it completely.
* Use [`System.GC.Collect`](https://learn.microsoft.com/en-us/dotnet/api/system.gc.collect?view=net-8.0) to perform a full, blocking garbage collection.

## Disabling incremental garbage collection

To disable incremental garbage collection:

1. Open the Project Settings window (**Edit** > **Project Settings**)
2. Navigate to the Configuration settings (**Player** > **Other Settings** > **Configuration**)
3. Disable the **Use Incremental GC** checkbox.

Incremental garbage collection is useful in most Unity projects, especially if the project has garbage collection spikes. However, incremental garbage collection adds write barriers to any calls that change references, so you might want to disable it if your project doesn’t trigger garbage collection in performance critical sections. Disabling incremental garbage collection in this situation can improve the performance of your scripting code in your project.

[Use the Profiler](profiler-introduction.md) to verify that your application performs as you expect. If you analyze a Profiler capture in isolation, it can be difficult to find out how incremental garbage collection affects performance. It’s best practice to profile the same performance critical section twice: once with incremental garbage collection enabled, and once with incremental garbage collection disabled. You can then compare both Profiler captures with the [Profile Analyzer package](https://docs.unity3d.com/Packages/com.unity.performance.profile-analyzer@latest). For CPU bound projects, the difference can be as much as 1 ms per frame.

For more information on disabling garbage collection, refer to [Garbage collection modes](performance-incremental-garbage-collection.md).

## Using the GarbageCollector API

You can use the APIs in the [`GarbageCollector`](../ScriptReference/Scripting.GarbageCollector.md) class to manually control the garbage collector, or disable it completely. You can use the following APIs to control the garbage collector:

* [`GarbageCollector.Mode.Disabled`](../ScriptReference/Scripting.GarbageCollector.Mode.Disabled.md): Setting [GarbageCollector.GCMode](../ScriptReference/Scripting.GarbageCollector.GCMode.md) to this fully disables the garbage collector. Using [`System.GC.Collect`](https://learn.microsoft.com/en-us/dotnet/api/system.gc.collect?view=net-8.0) in this mode has no effect.
* [`GarbageCollector.Mode.Manual`](../ScriptReference/Scripting.GarbageCollector.Mode.Manual.md): Setting [GarbageCollector.GCMode](../ScriptReference/Scripting.GarbageCollector.GCMode.md) to this fully disables automatic invocations of the garbage collector, but you can still use [`System.GC.Collect`](https://learn.microsoft.com/en-us/dotnet/api/system.gc.collect?view=net-8.0) to run a full collection.
* [`GarbageCollector.CollectIncremental`](../ScriptReference/Scripting.GarbageCollector.CollectIncremental.md): Runs the garbage collector [incrementally](performance-incremental-garbage-collection.md#incremental-garbage-collection).

### Disabling the garbage collector

To disable the garbage collector completely, set `GarbageCollector.GCMode` to [`Disabled`](../ScriptReference/Scripting.GarbageCollector.Mode.Disabled.md). When you disable the garbage collector, Unity doesn’t perform any garbage collection. Calling [`System.GC.Collect`](https://learn.microsoft.com/en-us/dotnet/api/system.gc.collect?view=net-8.0) has no effect and doesn’t start a collection.

Disabling the garbage collector prevents CPU spikes, but the memory usage of your application never decreases, because the garbage collector doesn’t collect objects that no longer have any references.

Disabling the garbage collector requires careful memory management. If you don’t manage memory carefully, the managed heap continuously expands until your application runs out of memory, and the operating system shuts it down.

Ideally, you should allocate all memory before you disable the garbage collector and avoid additional allocations while it is disabled.

### Manually run the garbage collector

To disable automatic garbage collection and manually choose when to run it, set `GarbageCollector.GCMode` to [`Manual`](../ScriptReference/Scripting.GarbageCollector.Mode.Manual.md).

This disables automatic invocations of the garbage collector, but still allows you to manually perform garbage collection. Manual collection gives you control over when collections happen, so you can fine-tune the smoothness of your content or your memory usage. Call either [`System.GC.Collect`](https://learn.microsoft.com/en-us/dotnet/api/system.gc.collect?view=net-8.0) for a full, blocking collection, or [`GarbageCollector.CollectIncremental`](../ScriptReference/Scripting.GarbageCollector.CollectIncremental.md) to perform incremental garbage collection.

## Additional resources

* [Garbage collection modes](performance-incremental-garbage-collection.md)
* [`GarbageCollector` API reference](../ScriptReference/Scripting.GarbageCollector.md)
* [Microsoft’s `System.GC.Collect` API reference](https://learn.microsoft.com/en-us/dotnet/api/system.gc.collect?view=net-8.0)