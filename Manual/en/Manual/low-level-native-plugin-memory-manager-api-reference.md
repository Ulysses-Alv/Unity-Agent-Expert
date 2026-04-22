# IUnityMemoryManager API reference

This pages provides the API reference for the `IUnityMemoryManager` interface.

## CreateAllocator

### Declaration

`UnityAllocator* (UNITY_INTERFACE_API * CreateAllocator)(const char* areaName, const char* objectName);`

### Parameters

| **Parameter** | **Description** |
| --- | --- |
| **const char\* areaName** | The name for the broad category for this allocator. |
| **const char\* objectName** | The name for this specific allocator. |

### Description

Creates a new Allocator object which can allocate blocks of memory.

## DestroyAllocator

### Declaration

`void(UNITY_INTERFACE_API * DestroyAllocator)(UnityAllocator * allocator);`

### Parameters

| **Parameter** | **Description** |
| --- | --- |
| **UnityAllocator \* allocator** | The allocator to delete. |

### Description

Deletes an existing Allocator object.

## Allocate

### Declaration

`void* (UNITY_INTERFACE_API * Allocate)(UnityAllocator * allocator, size_t size, size_t align, const char* file, int32_t line);`

### Parameters

| **Parameter** | **Description** |
| --- | --- |
| **UnityAllocator \* allocator** | The allocator to use for allocation. |
| **size\_t size** | How much memory to allocate, in bytes. |
| **size\_t align** | The alignment of the memory address for the resulting pointer. |
| **const char\* file** | The path to the source file where the call to make this allocation came from. Use the predefined macro *FILE* here. |
| **int32\_t line** | The line number in the source file where the call to make this allocation came from. Use the predefined macro *LINE* here. |

### Description

Allocates a block of memory using an existing allocator. This method returns a pointer to the newly allocated memory.

## Deallocate

### Declaration

`void(UNITY_INTERFACE_API * Deallocate)(UnityAllocator * allocator, void* ptr, const char* file, int32_t line);`

### Parameters

| **Parameter** | **Description** |
| --- | --- |
| **UnityAllocator \* allocator** | The allocator to use for deallocation. |
| **void\* ptr** | The pointer to the memory to be deallocated. |
| **const char\* file** | The path to the source file where the call to make this deallocation came from. Use the predefined macro *FILE* here. |
| **int32\_t line** | The line number in the source file where the call to make this deallocation came from. Use the predefined macro *LINE* here. |

### Description

Deallocates the memory that the specified pointer points to. This doesn’t set the pointer to NULL.

## Reallocate

### Declaration

`void* (UNITY_INTERFACE_API * Reallocate)(UnityAllocator * allocator, void* ptr, size_t size, size_t align, const char* file, int32_t line);`

### Parameters

| **Parameter** | **Description** |
| --- | --- |
| **UnityAllocator \* allocator** | The allocator to use for the reallocation operation. |
| **void\* ptr** | The pointer to the memory to be deallocated. |
| **size\_t size** | How much memory to allocate, in bytes. |
| **size\_t align** | The alignment of the memory address for the resulting pointer. |
| **const char\* file** | The path to the source file where the call to make this reallocation came from. Use the predefined macro *FILE* here. |
| **int32\_t line** | The line number in the source file where the call to make this reallocation came from. Use the predefined macro *LINE* here. |

### Description

Reallocates an existing pointer to point to a different block of memory.

## Implementation example

Below is an example implementation of the `IUnityMemoryManager` interface.

```
#include "IUnityInterface.h"
#include "IUnityMemoryManager.h"
#include <cstdint>

static IUnityMemoryManager* s_MemoryManager = NULL;
static UnityAllocator* s_Alloc = NULL;

extern "C" void UNITY_INTERFACE_EXPORT UNITY_INTERFACE_API UnityPluginLoad(IUnityInterfaces * unityInterfaces)
{
    s_MemoryManager = unityInterfaces->Get<IUnityMemoryManager>();
    if (s_MemoryManager  == NULL)
    return;

    // Create an allocator. This allows you to see the allocation root in the profiler when taking snapshots. Under plug-ins-native - Plugin Backend Allocator
   // All memory allocated here also goes under kMemNativePlugin
    s_Alloc = s_MemoryManager->CreateAllocator("plug-ins-native", "Plugin Backend Allocator");
}

extern "C" void UNITY_INTERFACE_EXPORT UNITY_INTERFACE_API UnityPluginUnload()
{
    //Free allocator
    s_MemoryManager->DestroyAllocator(s_Alloc);
    s_Alloc = NULL;
    s_MemoryManager = NULL;
}

void DoMemoryOperations()
{  
    // Allocate 1KB memory
    void* mem = s_MemoryManager->Allocate(s_Alloc, 1 * 1024, 16, __FILE__, __LINE__);
     // Reallocate the same pointer with 2KB
    mem = s_MemManager->Reallocate(s_Alloc, mem, 2 * 1024, 16, __FILE__, __LINE__);
    // Delete allocated memory
    s_MemoryManager->Deallocate(s_Alloc, mem, __FILE__, __LINE__);
}
```

## Additional resources

* [Native plug-in API for shader compiler](low-level-native-plugin-shader-compiler-access.md)
* [Native plug-in API for profiling](LowLevelNativePluginProfiler.md)
* [Native plug-in API for logging](low-level-native-plugin-logging.md)