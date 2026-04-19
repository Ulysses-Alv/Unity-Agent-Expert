# Introduction to debugging Unity in Windows

You can use various tools and techniques to debug Unity applications on Windows and identify and resolve runtime issues.

This guide covers practical debugging types, such as live debugging and forensic debugging. It also covers vital concepts like symbols and managed exceptions, which support and enhance your debugging process.

* [Unity symbols](#symbols)
* [Built-in debugging tools](#native-debugging)
* [C# managed debugging](#c-managed-debugging)
* [Live debugging](#live-debugging)
* [Forensic debugging](#forensic-debugging)
* [The UnityMixedCallstack plug-in](#plug-in)

## Unity symbols

Symbols provide debugging information including function names, variable names, and source code line numbers. This information helps a debugger (like WinDbg or Visual Studio) map low-level memory addresses back to readable, human-friendly data.

Unity’s Symbol Store (`http://symbolserver.unity3d.com/`) provides debugging symbols specific to Unity applications and development. You can configure Windows Debugger (WinDbg) or Visual Studio (2019 and later) to automatically download and resolve these symbols to help you debug Unity applications.

For instructions on how to set up WinDbg or Visual Studio to resolve Unity symbols, refer to [Configure your debugging tool to debug Unity in Windows](WindowsDebugging-instructions.md).

**Note**: By design, the Unity Symbol Store URL isn’t a browsable web page. If you visit the link directly in a browser, it doesn’t display any content. The symbols are accessible only via debugging tools.

## Built-in debugging tools

In Unity development, built-in debugging refers to the process of debugging unmanaged code, which is code written in low-level languages such as C, C++, or Assembly. This type of debugging is usually needed for Unity’s engine code, built-in **plug-ins**, or external libraries.

On Windows, debugging tools use Program Database (PDB) files (with the .pdb extension). These files store symbol information, such as function names and line numbers, for associated binary files like executables (.exe) or libraries (.dll).

Use this debugging type to analyze low-level issues like memory corruption, access violations, or performance-related issues.

To effectively use built-in debugging tools for Unity, you must configure them to work with Unity symbols, which provide additional visibility into Unity-specific unmanaged code. For instructions, refer to [Configure your debugging tool to debug Unity in Windows](WindowsDebugging-instructions.md).

## C# managed debugging

Managed debugging refers to debugging high-level code that runs within a managed runtime environment. In Unity, this primarily involves debugging the logic of C# **scripts** written to extend game behavior using MonoBehaviour classes and other Unity APIs.

You can use debugging tools such as Visual Studio to debug your Unity managed code, but you can also use the Unity **console window**. The Console window displays errors, warnings, and console messages if you enter Play mode in the Unity Editor.

## Live debugging

Live debugging is the process in which you attach a debugger to a process that’s already running, or to a process where an exception is caught. For the debugger to spot the issue, you must set up your debugging tool to use Unity symbols in the build. For instructions, refer to [Set up Visual Studio to debug Unity](WindowsDebugging-instructions.md#vs-setup).

For instructions on how to use live debugging in Unity, refer to [Set up live debugging for Unity](WindowsDebugging-live.md).

**Note:** If your game executable has the same name as your game name, the debugger might struggle to find the correct .pdb file. This is especially true if the debugger doesn’t have access to the renamed executable.

## Forensic debugging

When applications terminate due to a crash, they sometimes produce a crash dump file. Forensic debugging (or post-mortem debugging) is when you analyze this file to investigate the cause of the crash.

By default, Windows is configured to send application crash reports to Microsoft’s Dr. Watson or Windows Error Reporting services. However, if you have Visual Studio or WinDbg installed, Microsoft provides an option to [debug the crashes](https://msdn.microsoft.com/en-us/library/windows/desktop/bb204634(v=vs.85).aspx).

Windows provides facilities to investigate crash dump files (.dmp or .mdmp). Depending on the crash dump, you might either find stack information or the entire process memory. The crash dump file contains information that helps determine the cause of the crash. It typically includes a stack trace, provided the stack is intact and valid for analysis.

To investigate a dump file, you can load it in Visual Studio or WinDbg. While Visual Studio is easier to use, WinDbg provides additional features, which makes it a preferred debugging tool.

To learn how to make your application create a dump file, refer to [Set up forensic debugging for Unity](WindowsDebugging-forensic.md).

## The UnityMixedCallstack plug-in

When you run Visual Studio, you can use the [UnityMixedCallstack](https://github.com/Unity-Technologies/UnityMixedCallstack) extension for Visual Studio Code by Unity. This extension translates unreadable Mono-managed stack frames into readable ones, which simplifies built-in code debugging.

### Managed Exceptions in Unity

Managed exceptions occur when errors happen within managed code, for example, C# code that executes under Mono or the .NET runtime. The managed runtime’s exception-handling mechanism handles these exceptions.

For example, a NullReferenceException occurs when code attempts to access a member on a null object reference.

### Example: Debugging a null reference exception

The following samples show an example of what a null reference with a built-in debugger attached can look like.

This example shows a Unity C# script that triggers a null reference exception:

```
    private string name = null;

    void Start()
    {
        Debug.Log(name.Length);
    }
```

When this code runs with a built-in debugger (such as Visual Studio), the output window displays:

```
Exception thrown at 0x00000142725F6438 in Unity.exe: 0xC0000005: Access violation reading location 0x0000000000000010.
```

If you use Visual Studio with the `UnityMixedCallstack` plug-in enabled, the managed stack frames become readable, as seen in this call stack:

```
    [Assembly-CSharp.dll] SpinMe:Start ()   Unknown
    [mscorlib.dll] (wrapper runtime-invoke) object:runtime_invoke_void__this__ (object,intptr,intptr,intptr)    Unknown
    mono-2.0-bdwgc.dll!mono_jit_runtime_invoke(_MonoMethod * method, void * obj, void * * params, _MonoObject * * exc, _MonoError * error) Line 3445    C
    mono-2.0-bdwgc.dll!do_runtime_invoke(_MonoMethod * method, void * obj, void * * params, _MonoObject * * exc, _MonoError * error) Line 3068  C
    [Inline Frame] mono-2.0-bdwgc.dll!mono_runtime_try_invoke(_MonoMethod *) Line 3175  C
    mono-2.0-bdwgc.dll!mono_runtime_invoke(_MonoMethod * method, void * obj, void * * params, _MonoObject * * exc) Line 3115
```

The `SpinMe:Start()` appears in the call stack because of the mixed call stack plug-in. Without the plug-in, the method name is unreadable.

#### Resolve unreadable managed frames

Managed frames are stack trace lines that lack detailed information, such as file names or line numbers. Instead, they display `Unknown`.

To retrieve more detailed information, you can use the built-in `mono_pmip` function. This function accepts the address of a stack frame and returns a char\* with information. You can invoke `mono_pmip` in the Visual Studio’s **Immediate** window (**Debug** > **Windows** > **Immediate**) for debugging. For more information about this function, refer to [Debugging With Visual Studio on Windows](https://www.mono-project.com/docs/debug+profile/debug/#debugging-with-visual-studio-on-windows) (Mono).

```
> `?(char*){,,mono-2.0-bdwgc.dll}mono_pmip((void*)0x1b45558c)`  
0x26a296c0 " Tiles:OnPostRender () + 0x1e4 (1B4553A8 1B4555DC) [065C6BD0 - Unity Child Domain]"`
```

**Note:** This works only where `mono-2.0-bdwgc.dll` symbols load fully.

## Additional resources

* [Unity debugging guide for Windows](WindowsDebugging.md)
* [Configure your debugging tool to debug Unity in Windows](WindowsDebugging-instructions.md)
* [Set up live debugging for Unity](WindowsDebugging-live.md)
* [Set up forensic debugging for Unity](WindowsDebugging-forensic.md)