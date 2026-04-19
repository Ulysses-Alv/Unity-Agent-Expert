# Building plug-ins for desktop platforms

Plug-ins for desktop platforms are libraries of native code you can write in C, C++ and Objective C. This page describes **plug-ins** for macOS, Windows, and Linux. For more information, refer to [Native plug-ins](plug-ins-native.md).

## macOS plug-ins

You can deploy macOS plug-ins as bundles or, if you are using the ****IL2CPP**** **scripting backend**, loose C++ files, which you can invoke with `[DllImport("__Internal")]` syntax. For further information on loose C++ plug-ins refer to [C++ source code plugins for IL2CPP](macOSIL2CPPScriptingBackend.md).

To create the bundle project with Xcode:

1. Open Xcode.
2. Select **File** > **New** > **Project** > **macOS** > **Framework & Library** > **Bundle**.

For more information about working with Xcode, refer to [Apple’s documentation on Xcode](https://developer.apple.com/documentation/xcode).

### Requirements

* You can build your plug-in as a universal binary that’s compatible with 64-bit architectures. Alternatively, you can provide separate dylib files.
* If you are using C++ (`.cpp`) or Objective-C (`.mm`) to implement the plug-in, declare the functions with C linkage to avoid name mangling issues:

```
extern "C"
{
  float ExamplePluginFunction ();
}
```

## Windows plug-ins

Plug-ins on Windows are either `.dll` files with exported functions, or loose C++ files if you are using IL2CPP. You can use most languages and development environments that can create `.dll` files to create plug-ins. You must declare any C++ functions with C linkage to avoid name mangling issues.

## Linux plug-ins

Plug-ins on Linux are `.so` files with exported functions. Although these libraries are usually in C or C++, you can use any language. You must declare any C++ functions with C linkage to avoid name mangling issues.

When you build the native plug-ins for Linux, if the built library has dependencies on another native plug-in, you must specify the `rpath` for that library while compiling it.

Add the linker flag `-Wl, -rpath=$ORIGIN` to specify the runtime search path. The linker flag instructs the loader to find its dependencies in the current directory of the library besides searching the system search path. You can add other linker flags along with `-Wl, -rpath=$ORIGIN`, however, Unity doesn’t control them. For example, `/usr/bin/g++ -o binary.c.o -Wl,-rpath=$ORIGIN`.

Alternatively, you can set `LD_LIBRARY_PATH=dependency path` in the environment to instruct the loader to search that path for dependencies. Linux doesn’t automatically search the current directory for dependencies. Make sure to set the correct dependency search path because incorrect path causes missing library errors in the Unity Editor.

## Managing plug-ins inside Unity

In Unity, the **Plugin **Inspector**** manages your plug-ins. To access the **Plugin Inspector**, select a plug-in file in the ****Project window****. For Standalone platforms you can choose the CPU architecture with which the library is compatible. For cross-platform plug-ins you must include the `.bundle` file (for macOS), the `.dll` file(for Windows), and the `.so` file (for Linux). Unity automatically picks the right plug-in for the target platform and includes it with the player. For further information, refer to [Import and configure plug-ins](plug-in-inspector.md).

![Plugin Inspector](../uploads/Main/plugin-inspector.png)

Plugin Inspector

## Invoking your plug-in from a C# script

Place your built plug-in in the **Assets** folder or the appropriate architecture-specific sub-directory in your Unity Project. Unity then finds it by name when you invoke it from a C# script. For example: `[DllImport ("PluginName")] private static extern float ExamplePluginFunction ();`

**Note**: Don’t include the library prefix or file extension in the`PluginName` value. For example, if the actual name of the plug-in file is `PluginName.dll` on Windows or `libPluginName` on Linux, the value should be `PluginName` in both cases.

## Example plug-ins

You can download and use these projects to learn how to implement plug-ins in Unity.

* [Simplest Plugin Example](https://github.com/Unity-Technologies/DesktopSamples/tree/master/SimplestPluginExample): This project implements basic operations (for example, print a number, print a string, add two floats and add two integers). This project includes Windows, macOS and Linux project files.
* [Native Renderer Plugin](https://github.com/Unity-Technologies/NativeRenderingPlugin): This is a low-level rendering plug-in that renders a rotating triangle from C++ code after all regular rendering is done and fills a procedural texture from C++ code, using Texture.GetNativeTexturePtr to access it. This project includes Windows, UWP, macOS, Web and Android files.

## Additional resources

* [Managed plug-ins](plug-ins-managed.md)
* [Native plug-ins](plug-ins-native.md)
* [Native plug-in interface](native-plugin-interface.md)