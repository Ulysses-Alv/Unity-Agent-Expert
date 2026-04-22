# Introduction to IL2CPP

IL2CPP (Intermediate Language To C++) is Unity’s custom, ahead-of-time (AOT) scripting back end, which was originally developed for platforms that don’t support [Mono](scripting-backends-mono.md) and just-in-time (JIT) compilation, such as iOS. **IL2CPP** converts IL ([Intermediate Language](https://learn.microsoft.com/en-us/dotnet/standard/managed-code#intermediate-language--execution)) to C++, and compiles it into platform-specific native code, which is then packaged with your application in the standard format for the target platform, such as **APK**/AAB, iOS app bundle, or Windows executable and DLLs.

IL2CPP is supported on all platforms and can offer several benefits over the Mono scripting back end, including improved performance and shorter startup times. However, the need to include machine code in built applications generally increases both the build time and the size of the final built application.

IL2CPP supports the debugging of managed code in the same way as the Mono. For more information, refer to [Debugging C# code in Unity](managed-code-debugging.md).

## Building a project with IL2CPP

To build a project with IL2CPP, you must have Il2CPP installed as part of your Unity installation. You can select IL2CPP as an optional module when you first install a version of Unity, or add IL2CPP support to an existing installation through the Unity Hub. For more information, refer to [Installing the Unity Hub](https://docs.unity3d.com/hub/manual/InstallHub.html) and [Add modules to the Unity Editor](https://docs.unity3d.com/hub/manual/AddModules.html).

IL2CPP also requires some systems native to the target platform to generate the C++ code. This means that cross-compilation is generally not supported and to build an IL2CPP Player for a particular target platform you must build from an Editor running on the same platform. For example, to build a macOS Player with IL2CPP, you must build on macOS. The exception to this is Linux, where building Linux Players on other desktop platforms is supported. For more information, refer to [Linux IL2CPP cross-compiler](linux-il2cpp-crosscompiler.md).

You can change the scripting back end Unity uses to build your application in one of two ways:

* Through the **Player Settings** menu in the Editor. Perform the following steps to change the scripting back end through the **Player Settings** menu:

  1. Go to **Edit** > **Project Settings**.
  2. Click on the **Player Settings** button to open the [**Player**](class-PlayerSettings.md) settings for the current platform in the [**Inspector**](UsingTheInspector.md).
  3. Navigate to the **Configuration** section heading under the **Other Settings** sub-menu.
  4. Click on the ****Scripting Backend**** dropdown menu, then select **IL2CPP**.

  You can also open the **Player Settings** menu from the **Build Profiles** window; go to **File** > **Build Profiles** and click on the **Player Settings** tab.
* Through the Editor scripting API. Use the [`PlayerSettings.SetScriptingBackend`](../ScriptReference/PlayerSettings.SetScriptingBackend.md) property to change the scripting back end that Unity uses.

![The Configuration section of the Player settings](../uploads/Main/IL2CPP-build-settings.png)

The **Configuration** section of the **Player** settings

For more information about system requirements for desktop platforms, including IL2CPP requirements for individual platforms, refer to the Desktop section of [System Requirements for Unity](https://docs.unity3d.com/Manual/system-requirements.html#desktop).

## How IL2CPP works

When you build with IL2CPP, Unity automatically performs the following steps:

1. The Roslyn C# compiler compiles your application’s C# code and any required package code to .NET DLLs (managed assemblies).
2. Unity applies [managed code stripping](managed-code-stripping.md). This step can significantly reduce the size of a built application.
3. The IL2CPP back end converts all managed assemblies into standard C++ code.
4. The C++ compiler compiles the generated C++ code and the runtime part of IL2CPP with a native platform compiler.
5. Unity creates either an executable file or a DLL, depending on the platform you target.

## Optimizing IL2CPP build times

Project build times with IL2CPP can be significantly longer than with Mono. However, you can do several things to reduce build time:

* Exclude your Unity project and target build folders from anti-malware software scans before you build your project.
* Store your project and target build folder on the fastest disk drive you have available. Converting IL code to C++ and compiling it involves a large number of read/write operations, so a faster storage device speeds up this process.
* Change the IL2CPP Code Generation option in the Player Settings. To change how IL2CPP generates code, open the [Player Settings](class-PlayerSettings.md) and configure the IL2CPP Code Generation option. By default, the **Optimize for runtime speed** option is enabled, which produces more machine code that reduces the impact of IL2CPP at runtime. To reduce build times, you can set this option to **Optimize for code size and build time**. This method produces and includes less machine code in the binary executable and so can reduce performance at runtime, but also significantly reduces build times and binary size.

## Optimizing runtime performance with Burst

In suitable projects, you can improve the runtime performance of your project by using the Burst compiler alongside IL2CPP to compile compatible sections of your code to highly optimized machine code. For more information, refer to [Burst compilation](script-compilation-burst.md).

## Additional resources

* 📚 **Documentation**: [IL2CPP compiler options reference](il2cpp-options-reference)
* 📚 **Documentation**: [Linux IL2CPP cross-compiler](linux-il2cpp-crosscompiler.md)
* 📚 **Documentation**: [Reducing the file size of a build](ReducingFilesize.md)
* 👥 **Community**: [An introduction to IL2CPP internals.](https://blog.unity.com/technology/an-introduction-to-ilcpp-internals)