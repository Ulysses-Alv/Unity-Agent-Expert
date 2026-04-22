# Introduction to native plug-in API

In addition to the public C# APIs for writing managed code, Unity provides a smaller native interface, which you can use to access Unity Editor and Engine functionality from your native **plug-ins**.

## Plugin API folder

The native interfaces are a set of C or C++ compatible header (.h) files incuded as part of an Editor installation. The files are in the `PluginAPI` folder, which you can find in the following locations:

* **Windows**: `<UnityInstallPath>\Editor\Data\PluginAPI`
* **macOS**: Right-click on the Unity application, and select Show Package Contents. The headers are in `Contents\PluginAPI`

Each header file contains additional documentation as code comments within it. For more information on getting started with implementing the native interfaces, refer to the code comments in the main header file `IUnityInterface.h`.

## Interface compatibility

All Unity **native plug-in** API header files are compatible with plug-ins written in C++, but only some are compatible with plug-ins written in C.

Files that aren’t C-compatible report the error `"This file cannot be compiled in a C environment"` if you attempt to compile them as C. You’ll also find the corresponding check in the header file source:

```
#ifndef __cplusplus
#error "This file cannot be compiled in a C environment"
#endif
```

## Interface registry

To handle main Unity events, a plug-in must export `UnityPluginLoad` and `UnityPluginUnload` functions. `IUnityInterfaces` enables the plug-in to access these functions, which you can find in `IUnityInterface.h` in the plug-in API.

The following example uses `IUnityInterfaces` to load the `IUnityGraphics` interface into a pointer. This is a standard method you can repeat to load other interfaces from the native plug-in API:

```
#include "IUnityInterface.h"
#include "IUnityGraphics.h"
// Unity plugin load event
extern "C" void UNITY_INTERFACE_EXPORT UNITY_INTERFACE_API
    UnityPluginLoad(IUnityInterfaces* unityInterfaces)
{
    IUnityGraphics* graphics = unityInterfaces->Get<IUnityGraphics>();
}
```

## Additional resources

* [Native plug-in API for graphics and rendering](low-level-native-plugin-rendering-extensions.md)
* [Native plug-in API for shader compiler](low-level-native-plugin-shader-compiler-access.md)
* [Native plug-in API for memory manager](low-level-native-plugin-memory-manager-api.md)
* [Native plug-in API for profiling](LowLevelNativePluginProfiler.md)
* [Native plug-in API for logging](low-level-native-plugin-logging.md)