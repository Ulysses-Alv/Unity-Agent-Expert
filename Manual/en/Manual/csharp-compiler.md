# C# compiler and language version reference

This version of the Unity Editor uses the following C# compiler and language version:

* **C# compiler**: [Roslyn](https://github.com/dotnet/roslyn)
* **C# language version**: [C# 9.0](https://docs.microsoft.com/en-us/dotnet/csharp/whats-new/csharp-9)

The Editor passes a default set of options to the C# compiler. To pass additional options in your project, refer to [Conditional compilation in Unity](platform-dependent-compilation.md).

## Garbage collection

Unity uses the [Boehm-Demers-Weiser garbage collector](https://www.hboehm.info/gc/) for both the Mono and **IL2CPP** [scripting backends](scripting-backends.md) and incremental mode by default. For more information on the available garbage collection modes, their meaning, and how to switch between them, refer to [Garbage collection modes](performance-incremental-garbage-collection.md).

## Unsupported features

### C# 9.0

* Suppress emitting localsinit flag
* Covariant return types
* Module Initializers
* Extensible calling conventions for unmanaged function pointers
* Init only setters

If you try to use unsupported features in your project, compilation generates errors.

### Record support

C# 9 init and record support comes with a few caveats.

* The type `System.Runtime.CompilerServices.IsExternalInit` is required for full record support as it uses init only setters, but is only available in .NET 5 and later (which Unity doesn’t support). Users can work around this issue by declaring the `System.Runtime.CompilerServices.IsExternalInit` type in their own projects.
* You shouldn’t use C# records in serialized types because Unity’s serialization system doesn’t support C# records.

### Unmanaged function pointer support

Unity supports unmanaged functions pointers as introduced in C# 9, but it doesn’t support extensible calling conventions. The following example code provides more detailed information about how to correctly use unmanaged function pointers.

The following example targets Windows platforms and requires **Allow ‘unsafe’ Code** to be enabled in the [Player Settings](class-PlayerSettings.md). To enable it, go to: **Project Settings** > **Player**. Expand the **Other Settings** panel, navigate to the **Script Compilation** section. For more information on C#’s `unsafe` context, refer to [Microsoft’s unsafe (C# Reference) documentation](https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/keywords/unsafe) or [Microsoft’s Unsafe code, pointer types, and function pointers documentation](https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/unsafe-code).

```
using System;
using System.Runtime.InteropServices;
using UnityEngine;

public class UnmanagedFunctionPointers : MonoBehaviour
{
  [DllImport("kernel32.dll")]
  static extern IntPtr LoadLibrary(string lpLibFileName);
  
  [DllImport("kernel32.dll")]
  static extern IntPtr GetProcAddress(IntPtr hModule, string lpProcName);
  
  // You must enable "Allow 'unsafe' code" in the Player Settings
  unsafe void Start()
  {
#if UNITY_EDITOR_WIN || UNITY_STANDALONE_WIN
    // This example is only valid on Windows
    
    // Get a pointer to an unmanaged function
    IntPtr kernel32 = LoadLibrary("kernel32.dll");
    IntPtr getCurrentThreadId = GetProcAddress(kernel32, "GetCurrentThreadId");

    // The unmanaged calling convention
    delegate* unmanaged[Stdcall]<UInt32> getCurrentThreadIdUnmanagedStdcall = (delegate* unmanaged[Stdcall]<UInt32>)getCurrentThreadId;
    Debug.Log(getCurrentThreadIdUnmanagedStdcall());
#endif
  }
}
```

## Additional resources

* [API compatibility levels for .NET](dotnet-profile-support.md).
* [Platform dependent compilation](platform-dependent-compilation.md).