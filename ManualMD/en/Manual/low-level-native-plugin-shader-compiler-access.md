# Native plug-in API for shader compiler

Use the Unity low-level **shader** compiler APIs to inject different variants into a shader. This is an event-driven approach in which the **plug-in** receives callbacks when certain built-in events happen.

The shader compiler access extension definition exposed by Unity is to be found in `IUnityShaderCompilerAccess.h` header file, located in the [PluginAPI folder](native-plugin-interface-introduction.md).

**Note**: These extensions are currently supported only on D3D11. `IUnityShaderCompilerAccess.h` can only compile as C++. Attempting to compile it as C produces an [error](native-plugin-interface-introduction.md#compatibility).

## Shader compiler access extension API

To use the rendering extension, plug-ins must export `UnityShaderCompilerExtEvent`. For more information on implementing the API, refer to the code comments in the `IUnityShaderCompilerAccess.h` header file.

A plug-in receives a callback via `UnityShaderCompilerExtEvent` whenever Unity triggers one of the built-in events. The callbacks can also be added to CommandBuffers via `CommandBuffer.IssuePluginEventAndData` or `CommandBuffer.IssuePluginCustomBlit` commands from **scripts**.

In addition to the basic script interface, [native plug-ins](plug-ins.md) in Unity can receive callbacks when certain events happen. This is mostly used to implement low-level rendering in your plug-in and enable it to work with Unity’s multithreaded rendering.

## Shader compiler access configuration interface

Unity provides the (`IUnityShaderCompilerExtPluginConfigure`) interface for configuring shader compiler access. A plug-in can use this interface to reserve its own keyword(s) and configure shader program and GPU program compiler masks. This determines which shader types or GPU programs to invoke the plug-in for.

## Additional resources

* [Native plug-in API for memory manager](low-level-native-plugin-memory-manager-api.md)
* [Native plug-in API for profiling](LowLevelNativePluginProfiler.md)
* [Native plug-in API for logging](low-level-native-plugin-logging.md)