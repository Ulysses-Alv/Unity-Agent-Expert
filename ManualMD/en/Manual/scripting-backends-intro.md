# Introduction to scripting back ends

Scripting back end is the Unity term for the runtime technology that compiles and executes your C# **scripts**. It determines how your code is turned into executable instructions and what runtime manages it on target platforms.

The choice of a scripting back end can impact various aspects of your project, including:

* Build times and iteration speed.
* Runtime performance and memory.
* Platform support and requirements.
* Support for some .NET features such as reflection and dynamic code generation.
* Debugging and profiling experience.

Unity supports the following scripting back ends:

* [Mono](scripting-backends-mono.md): An open-source implementation of the .NET Framework based on the ECMA standards for C# and the Common Language Infrastructure (CLI). Mono uses just-in-time (JIT) compilation to convert your C# code into machine code at runtime.
* [IL2CPP](scripting-backends-il2cpp.md): A Unity-developed scripting back end that converts your C# code and assemblies into C++ code, which is then compiled using a platform-native C++ compiler to produce a binary file for the target platform. IL2CPP uses ahead-of-time (AOT) compilation to convert your code before runtime.

## Platform support

Mono is supported on the following platforms:

* Windows x86/x64/Arm64
* OSX x64/Arm64
* Linux x64
* Android Armv7

IL2CPP is supported on all platforms.

Any platform not mentioned in the previous list of Mono-supporting platforms is IL2CPP-only. On platforms that support both back ends, Mono is the default.

**Note**: Scripting back end support can change between releases. For the latest information, refer to the [platform-specific Player settings](PlatformSpecific.md) for your target platform and Unity version.

## Changing the scripting back end

You can change the scripting back end Unity uses to build your application in one of two ways:

* Through the ****Player Settings**** menu in the Editor. Perform the following steps to change the scripting back end through the **Player Settings** menu:

  1. Go to **Edit** > **Project Settings**.
  2. Click on the **Player Settings** button to open the [**Player**](class-PlayerSettings.md) settings for the current platform in the [**Inspector**](UsingTheInspector.md).
  3. Navigate to the **Configuration** section heading under the **Other Settings** sub-menu.
  4. Click on the ****Scripting Backend**** dropdown menu, then select **Mono** or **IL2CPP**.

  You can also open the **Player Settings** menu from the **Build Profiles** window; go to **File** > **Build Profiles** and click on the **Player Settings** tab.
* Through the Editor scripting API. Use the [`PlayerSettings.SetScriptingBackend`](../ScriptReference/PlayerSettings.SetScriptingBackend.md) property to change the scripting back end that Unity uses.

![The Configuration section of the Player settings](../uploads/Main/backend-mono.png)

The **Configuration** section of the **Player** settings

## Scripting backends comparison

The C#/.NET API surface available for Mono and IL2CPP is nearly identical. Both scripting back ends use the same base class library and both support the available [API compatibility levels](dotnet-profile-support.md). However, some restrictions specific to IL2CPP and certain AOT platforms apply. For more information, refer to [IL2CPP restrictions](scripting-restrictions.md).

The following table provides a side‑by‑side comparison of Mono and IL2CPP scripting back ends across key features:

| Feature | Mono | IL2CPP |
| --- | --- | --- |
| Execution model | Just-in-time (JIT). | Ahead-of-time (AOT): C# to IL to C++ to native code |
| Platforms | Desktop and Android. | Widest support and required for iOS and most consoles. |
| Cross‑compilation toolchain | Minimal; JIT compiles managed assemblies as needed. | Requires C++ toolchain (Clang/MSVC), platform SDKs; longer build/link times. For more information, refer to [Linux IL2CPP cross-compiler](linux-il2cpp-crosscompiler.md). |
| Build time | Faster iteration, smaller linking steps. | Slower (C++ codegen, compiling, and linking), larger generated codebases. |
| Startup time | Typically slower with JIT (warmup, JIT cost). | Typically faster (fully native, no JIT warmup). |
| Memory usage | Smaller code size with JIT | Larger native binaries; additional generated C++ and metadata |
| Exception handling | Managed exceptions via Mono runtime. | Managed semantics mapped to native via IL2CPP runtime support. |
| Reflection | Full reflection works, dynamic codegen supported. | Reflection works for existing metadata, but dynamic codegen has limits. `System.Reflection.Emit` and `dynamic` methods not supported. |
| Generics | Full support, JIT can specialize at runtime. | AOT uses ahead‑of‑time generic sharing and IL2CPP might need additional AOT hints to avoid missing method or type errors at runtime. For more information, refer to [Generic types and methods](scripting-restrictions.md#Generics). |
| Code stripping/linker | Managed linker, easier to keep symbols due to JIT. | More aggressive stripping, must annotate or configure to preserve reflected or dynamically loaded members. Use a `link.xml` file or [`[Preserve]`](managed-code-stripping-preserving.md) attribute to keep types and methods alive. For more information, refer to [Managed code stripping](managed-code-stripping.md). |
| Interop (`P/Invoke`) | Standard `P/Invoke`, runtime resolves at load. | Standard `P/Invoke`, names must match because AOT/native binary linkage is stricter. For more information, refer to [Calling managed methods from native code](scripting-restrictions.md#Interop). |
| Debugging (C#) | Fast domain reloads, edit‑and‑continue on platforms allowing JIT, informative managed stack traces by default. | Supported managed debugging, stack traces can be optionally enhanced. For more information, refer to [IL2CPP managed stack traces](il2cpp-managed-stack-traces.md). |

On platforms where you have a choice between IL2CPP and Mono, prefer IL2CPP if Player builds need faster startup, stricter platform compliance, and more predictable performance. Mono offers faster iteration in Editor and desktop development, faster build times, and is the only option if you need to use APIs [not supported by IL2CPP](scripting-restrictions.md).

**Tip**: Whichever scripting back end you choose, you can supplement it with the Burst compiler to compile select parts of your code to highly optimized native code. For more information, refer to [Burst compilation](script-compilation-burst.md).

## Additional resources

* [Mono scripting back end](scripting-backends-mono.md)
* [IL2CPP scripting back end](scripting-backens-il2cpp)
* [Scripting restrictions](scripting-restrictions.md)