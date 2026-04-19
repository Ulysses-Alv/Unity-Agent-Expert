# Memory in Unity introduction

To ensure your application runs with no performance issues, it’s important to understand how Unity uses and allocates memory.

Unity uses the following memory management layers to handle memory in your application:

* **[Managed memory](#managed-memory)**: A controlled memory layer that uses a managed heap and a [garbage collector](performance-garbage-collector.md) to automatically allocate and assign memory.
* **[C# unmanaged memory](#unmanaged-memory)**: A layer of memory management that you can use with the `Unity.Collections` namespace, the [Collections package](https://docs.unity3d.com/Packages/com.unity.collections@latest), and via [`UnsafeUtility.Malloc`](../ScriptReference/Unity.Collections.LowLevel.Unsafe.UnsafeUtility.Malloc.md) and [`UnsafeUtility.Free`](../ScriptReference/Unity.Collections.LowLevel.Unsafe.UnsafeUtility.Free.md). This memory type is called unmanaged because it doesn’t use a garbage collector to track and manage when memory stops being used, but requires your code to explicitly manage and release the memory itself, by calling `Dispose` on the collection or `Free` on memory allocated with `UnsafeUtility.Malloc`.
* **[Native memory](#native-memory)**: C++ memory that Unity uses to run the engine. It contains memory related to areas such as memory used for the assets in your project, and managers for Unity’s different native subsystems like the rendering or animation systems. In most situations, you can’t access this memory through your C# code, but because it’s usually the biggest chunk of your application’s memory footprint, it’s useful to be aware of it if you want to optimize your application’s memory usage.

**Note**: You can use the [Memory Profiler](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest) package to analyze your memory usage.

## Managed memory

[Mono and IL2CPP’s](scripting-backends.md) scripting virtual machines (VM) implement the managed memory system, which is sometimes referred to as the scripting memory system. The VM offer a controlled memory environment divided into the following different types:

* **The managed heap**: A section of memory that the VM automatically controls with a [garbage collector](performance-garbage-collector.md).
* **The scripting stack:** This is built up and unwound as your application steps into and out of any code scopes.
* **Native VM memory:** Contains memory related to Unity’s scripting layer.

Using managed memory in Unity is the easiest way to manage the memory in your application, but it has some disadvantages. The garbage collector is convenient to use, but it’s also unpredictable in how it releases and allocates memory, which might lead to performance issues such as stuttering, which happens when the garbage collector has to stop to release and allocate memory. To work around this unpredictability, you can use the [C# unmanaged memory layer](#unmanaged-memory).

For more information on how managed memory works, refer to the section on [Managed memory](performance-managed-memory.md).

## C# unmanaged memory

The **C# unmanaged memory layer** allows you to access the native memory layer to fine-tune memory allocations, with the convenience of writing C# code.

You can use the `Unity.Collections` namespace (including [`NativeArray`](../ScriptReference/Unity.Collections.NativeArray_1.md)) in the Unity core API, the data structures in the [Unity Collections package](https://docs.unity3d.com/Packages/com.unity.collections@latest), and [`UnsafeUtility.Malloc`](../ScriptReference/Unity.Collections.LowLevel.Unsafe.UnsafeUtility.Malloc.md) and [`UnsafeUtility.Free`](../ScriptReference/Unity.Collections.LowLevel.Unsafe.UnsafeUtility.Free.md) to access C# unmanaged memory.

If you use Unity’s [job system](job-system.md), or [Burst](http://docs.unity3d.com/Packages/com.unity.burst@latest), you must use C# unmanaged memory. Additionally, to be able to use Burst you can only use the containers in the `Unity.Collections` namespace or types that fulfill the C# [unmanaged constraint](https://learn.microsoft.com/en-us/dotnet/csharp/language-reference/builtin-types/unmanaged-types).

Unity’s native code, and sometimes its native allocators, allocate the memory that’s accessible to C# through these APIs. Because of this, tools like the [Memory Profiler](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest) group C# unmanaged memory in with the native memory.

## Native memory

The Unity engine’s internal C/C++ core has its own memory management system, called native memory. In most situations, you can’t directly access or modify this memory type.

Unity stores the **scenes** in your project, assets, graphics APIs, graphics drivers, subsystems, **plug-in** buffers, and allocations inside native memory. This means that you can indirectly access the native memory via Unity’s C# API and manipulate your application’s data in a safe way, and get the benefits of the native and efficient code that’s in the internal C/C++ core.

Most of the time, you won’t need to interact with Unity’s native memory, but you can check how it affects the performance of your application whenever you use the [Profiler](ProfilerWindow.md), through [Profiler markers](profiler-markers.md). You can also adjust the performance of your application by changing the settings on your assets and via configurable settings for Unity’s native memory allocators. For more information, refer to [Native memory allocators](performance-native-allocators.md).

## Additional resources

* [Managed memory](performance-managed-memory.md)
* [Native memory](performance-native-memory.md)
* [C# unmanaged memory](performance-unmanaged-memory.md)