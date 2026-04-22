# Native memory introduction

The Unity engine’s internal C/C++ core has its own memory management system, called native memory. In most situations, you can’t directly access or modify this memory type.

Native memory isn’t automatically managed or subject to [garbage collection](performance-garbage-collector.md). This means it’s difficult to profile and handle in a way that doesn’t cause memory leaks.

Unity uses different memory allocator types, which all use different algorithms to organize memory. Unity’s memory manager uses these allocator types in different areas to organize the memory in your application effectively.

To get greater control over how Unity allocates native memory, you can adjust the size of each allocation for each area. You can adjust the size of allocations either through the [Player settings](class-PlayerSettings.md) window in the Unity Editor, or through the command line. For more information, refer to [Customizing allocators](memory-allocator-customization.md).

## Additional resources

* [Memory in Unity introduction](performance-memory-overview.md)
* [Native allocators introduction](performance-native-allocators.md)
* [Customizing allocators](memory-allocator-customization.md)