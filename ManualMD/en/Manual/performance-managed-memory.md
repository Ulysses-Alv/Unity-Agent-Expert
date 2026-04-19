# Managed memory

Unity’s managed memory system is a C# scripting environment based on the Mono or **IL2CPP** Virtual Machines (VMs). The benefit of the managed memory system is that it manages the release of memory, so you don’t need to manually request the release of memory through your code. It makes use of a garbage collector to automatically free memory allocations.

| **Topic** | **Description** |
| --- | --- |
| **[Managed memory introduction](performance-managed-memory-introduction.md)** | Automatically manage the release and allocation of memory in your application. |
| **[Garbage collector introduction](performance-garbage-collector.md)** | Reclaim unused memory in your application. |
| **[Garbage collection modes](performance-incremental-garbage-collection.md)** | Overview of the different ways that the garbage collector runs. |
| **[Configuring garbage collection](performance-disabling-garbage-collection.md)** | Set up the garbage collector in your project. |
| **[Tracking garbage collection allocations](performance-track-garbage-collection.md)** | Monitor when your application performed garbage collection. |
| **[Avoid C# reflection overhead](performance-gc-avoid-reflection.md)** | Reduce garbage collector overhead by avoiding use of C# reflection at runtime. |

## Additional resources

* [Memory Profiler package](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest)
* [Memory in Unity introduction](performance-memory-overview.md)
* [Memory Profiler module](profiler-memory.md)