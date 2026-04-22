# Native memory

The Unity engine’s internal C/C++ core has its own memory management system, which is referred to as native memory. In most situations, you can’t directly access or modify this memory type, but you can customize some of the memory allocators.

| **Topic** | **Description** |
| --- | --- |
| **[Native memory introduction](performance-native-memory-introduction.md)** | Understand how Unity’s internal processes use memory. |
| **[Native memory allocators](performance-native-allocators.md)** | Understand the different types of allocators Unity uses to allocate memory. |
| **[Customizing native memory allocators](memory-allocator-customization.md)** | Customize the allocator types with the Editor or command line. |
| **[Native memory allocators reference](performance-native-memory-allocator-reference.md)** | Reference for native memory allocator settings |
| **[Native memory allocator examples](performance-native-memory-allocator-examples.md)** | Refer to these native memory allocator examples for common issues. |

## Additional resources

* [Memory in Unity introduction](performance-memory-overview.md)
* [Managed memory](performance-managed-memory.md)