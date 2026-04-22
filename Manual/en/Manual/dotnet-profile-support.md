# API compatibility levels for .NET

Unity supports two different .NET API compatability levels that your project code can compile against. The compatability level you select determines which .NET class libraries are available to your C# code. The available levels are:

* **.NET Standard 2.1**: [.NET Standard](https://learn.microsoft.com/en-us/dotnet/standard/net-standard?tabs=net-standard-1-0) 2.1, as published by the .NET Foundation. This is the default API compatibility Level.
* **.NET Framework**: .NET Framework 4.8, as published by Microsoft, plus additional APIs from .NET Standard 2.1.

To change the .NET profile:

* In the Unity Editor, go to **Edit** > **Project Settings**. In the **Player** tab, navigate to the **Other Settings** tab and in the **Configuration** section, set **Api Compatibility Level** to the desired setting.
* From code, use [`PlayerSettings.SetAPICompatibilityLevel`](../ScriptReference/PlayerSettings.SetApiCompatibilityLevel.md).

## Choosing an API compatibility level

In general, prefer .NET Standard over .NET Framework for all new projects for the following reasons:

* .NET Standard is a smaller API surface and so has a smaller implementation. This reduces the size of your final executable file.
* .NET Standard has better cross-platform support, so your code is more likely to work across all platforms. Although some platforms don’t fully support the NET Standard, the .NET Framework option is less suitable for cross-platform compatibility. The .NET Framework profile includes all APIs in the .NET Standard profile plus additional APIs, some of which might work on few or no platforms.
* All .NET runtimes support .NET Standard, so your code works across more runtime environments (for example, .NET Framework, .NET Core, Xamarin, Unity) when you use .NET Standard.
* .NET Standard moves more errors to compile time. A number of APIs in .NET Framework are available at compile time, but have implementations on some platforms that throw an exception at runtime.

The .NET Framework API compatibility level can be useful if, for example, you need to provide support for an older existing application.

## Managed plug-ins

[Managed plug-ins](plug-ins-managed.md) are .NET assemblies that are managed outside of Unity and compiled into dynamically linked libraries (DLLs). You can use managed **plug-ins** with either of the available API compatibility levels. However, the level of support varies depending on the API level the managed plug-ins were compiled against. Because the .NET Framework API surface is larger than .NET Standard, if your project targets .NET Standard then it might not support all plug-ins compiled against .NET Framework. The following table indicates the configurations that Unity supports:

| **Managed plug-in compilation target** | **.NET Standard 2.1** | **.NET Framework 4.x** |
| --- | --- | --- |
| **.NET Standard (any version)** | Supported | Supported |
| **.NET Framework (any version)** | Limited support | Supported |
| **.NET Core (any version)** | Not Supported | Not Supported |

### Third-party .NET libraries

Only use third-party .NET libraries that have been extensively tested on a wide range of Unity configurations and platforms.

The [Mono](scripting-backends-mono.md) scripting backend’s just-in-time (JIT) compilation enables dynamic C#/.NET Intermediate Language (IL) code generation during the runtime of your application. The [IL2CPP](scripting-backends-il2cpp.md) scripting backend’s ahead-of-time (AOT) compilation doesn’t support dynamic code generation.

This is important to consider when you use third-party libraries, because they might have different code paths for JIT and **AOT compilation**, or they might use code paths that rely on dynamically generated code. For more information on how to generate code at runtime, refer to Microsoft’s [ModuleBuilder](https://docs.microsoft.com/en-us/dotnet/api/system.reflection.emit.modulebuilder?view=netcore-3.1) documentation.

The performance characteristics of JIT and AOT code paths in third-party libraries might be significantly different. AOT generally reduces startup times and is suited to larger applications for this reason but increases the binary file size to accommodate the compiled code. AOT also takes longer to build during development.

JIT adjusts at runtime based on the platform it’s running on, which can improve runtime performance at the cost of a potentially longer application startup time. As such, it’s recommended to profile your application in both the Editor, and on your target platform. For more information, refer to [Profiler overview](Profiler.md).

When you review a third-party library, consider the following areas:

* Compatibility: Third-party libraries might not be compatible with some Unity platforms and scripting backends.
* Performance: Third-party libraries might have very different performance characteristics in Unity compared to other .NET runtimes.
* AOT binary size: Third-party libraries might increase AOT binary size significantly because of the number of dependencies the library uses.

## .NET System libraries

Unity supports many platforms and might use different [scripting backends](scripting-backends.md) depending on the platform. The libraries in the .NET [`System`](https://learn.microsoft.com/en-us/dotnet/api/system?view=net-9.0) namespace require platform-specific implementations to work correctly in some cases. While Unity tries to support as much of the .NET ecosystem as possible, there are some parts of the .NET `System` libraries that Unity explicitly doesn’t support:

* Unity makes no performance or allocation guarantees for the .NET `System` libraries across Unity versions.
* Generally, Unity doesn’t fix any performance regressions in the .NET `System` libraries.
* Unity doesn’t support the [`System.Drawing`](https://learn.microsoft.com/en-us/dotnet/api/system.drawing?view=net-8.0) library and it isn’t guaranteed to work on all platforms.

**Tip**: Profile the usage of your .NET System libraries on all target platforms because their performance characteristics might vary depending on the scripting backends, .NET versions, and API compatibility level you use.

## Transport Layer Security (TLS) 1.2

The [`UnityWebRequest`](../ScriptReference/Networking.UnityWebRequest.md) API and all .NET Framework Web APIs fully support TLS 1.2 on all platforms except Web. The Web platform uses the security settings from the browser the application runs in and the web server instead. The platform-specific local certificate store automatically verifies TLS certificates if available. If access to the certificate store isn’t possible, Unity uses an embedded root certificate store.

## Additional resources

* [Add class library references to .NET Framework](dotnet-profile-assemblies.md)