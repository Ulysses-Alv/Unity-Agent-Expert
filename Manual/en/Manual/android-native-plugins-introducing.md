# Introducing native plug-ins for Android

You can use [native plug-ins](plug-ins-native.md) in Android applications. There are different types of native **plug-in**, and the project’s [scripting backend](scripting-backends.md) determines which plug-in types Unity supports. The possible plug-in types are:

* **Shared library**: Plug-ins packaged in a shared library (`.so`).
* **Static library**: Plug-ins packaged in a static library (`.a`).
* **C/C++ source files**: C/C++ source files that Unity compiles along with [IL2CPP](./scripting-backends-il2cpp.md) generated files. This includes all C/C++ source files with extensions `.c`, `.cc`, `.cpp` and `.h`.

## Scripting backend compatibility

The following table shows which scripting backends support the different types of native plug-ins.

| **Scripting backend** | **Shared library** | **Static library** | **C/C++ source files** |
| --- | --- | --- | --- |
| **IL2CPP** | Yes | Yes | Yes |
| **Mono** | Yes | No | No |

**Notes**:

* IL2CPP scripting backend doesn’t support direct use of .NET dynamic linked libraries (.dll) at runtime. These libraries must be converted to C++ code during the build process.
* Integration with other .NET runtime platforms, such as Xamarin isn’t supported.

## Additional resources

* [Create a native plug-in for Android](android-native-plugins-create.md)