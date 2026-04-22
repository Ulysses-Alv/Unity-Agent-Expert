# Native plug-in API for memory management

The `IUnityMemoryManager` memory manager API allows you to use Unity’s memory management and memory profiling in native **plug-ins** written in C or C++.

This API enables you to:

* Access Unity’s memory manager through a memory allocator.
* Track your plug-in’s memory use through Unity’s Memory Profiler package.

These features make it easier to manage and profile your plug-in’s memory allocations when compared to the equivalent C++ memory management methods.

The plug-in API is provided by the `IUnityMemoryManager` interface, which is declared in the `IUnityMemoryManager.h` header file, located in the [PluginAPI folder](native-plugin-interface-introduction.md#plugin-api-folder).

For more information, refer to the documentation provided as code comments in the header file.

You should be familiar with the following concepts to use this API effectively:

* [C++ Pointers](https://learn.microsoft.com/en-us/cpp/cpp/pointers-cpp?view=msvc-170)
* [Memory in Unity](performance-memory-overview.md)
* [The Memory Profiler package](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest)
* [Memory allocator customization](memory-allocator-customization.md)
* [Predefined macros](https://learn.microsoft.com/en-us/cpp/preprocessor/predefined-macros?view=msvc-170)

## Track memory use in Unity

To track your plug-in’s memory usage, use the [Memory Profiler package](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest) to take a snapshot, then open the snapshot in the [All Of Memory tab](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest/index.html?subfolder=/manual/all-memory-tab.html). When you use the `IUnityMemoryManager` to allocate memory, the Memory Profiler displays the plug-in’s memory allocations under the area and object name you assigned when you created each allocator.

The following screenshot shows the Memory Profiler package window displaying memory used by a native plug-in with memory allocated with the `IUnityMemoryManager` API. In this example, the method **CreateAllocator** was called, with “MyNativePlugin” as the areaName parameter, and “MyPluginAllocator” as the objectName parameter. For more information, refer to [IUnityMemoryManager API reference](low-level-native-plugin-memory-manager-api-reference.md).

![The Memory Profiler package window displaying the memory used by a user-defined allocator named Plugin Backend Allocator.](../uploads/Main/native-plugin-memory-snapshot.png)

The Memory Profiler package window displaying the memory used by a user-defined allocator named Plugin Backend Allocator.

For more information, refer to [Snapshots](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest/index.html?subfolder=/manual/snapshots.html).

## Memory management limitations

This API enables you to use Unity’s memory management system when you develop native plug-ins. This has major benefits as described above, but there are still limitations. Unity’s memory management system:

* Isn’t automatically managed; you need to allocate and deallocate the memory yourself.
* Isn’t tracked and cleaned up by a garbage collector.

Since memory in native C++ isn’t managed, you need to keep track of any memory requirements your application has. This includes choosing the correct amount of memory to allocate and making sure you deallocate it when it’s no longer needed.

The `IUnityMemoryManager` API impacts performance because each allocation requires a virtual call. To minimize this performance impact, use the API to allocate larger blocks of memory less frequently. To handle smaller and more frequent allocations, use this API to allocate a single larger block, then write your own code to manage the memory within this block. Don’t use this API for frequent small allocations.

## Additional resources

* [Native Plug-in interface](native-plugin-interface.md)
* [Native plug-ins](plug-ins-native.md)
* [Memory in Unity introduction](performance-memory-overview.md)
* [Customizing native memory allocators](memory-allocator-customization.md)