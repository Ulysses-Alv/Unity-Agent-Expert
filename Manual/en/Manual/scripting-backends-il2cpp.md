# IL2CPP scripting back end

The **IL2CPP** (Intermediate Language To C++) scripting back end is an alternative to Mono that compiles code ahead-of-time (AOT) and supports a wider range of platforms.

| **Topic** | **Description** |
| --- | --- |
| **[Introduction to IL2CPP](il2cpp-introduction.md)** | Understand IL2CPP’s role in Unity script compilation and how it works. |
| **[IL2CPP runtime code checks](handling-il2cpp-additional-args.md)** | Configure IL2CPP’s generation of C++ to enable or disable the inclusion of runtime safety features such as null reference and out of bounds checks. |
| **[IL2CPP managed stack traces](il2cpp-managed-stack-trace)** | Configure IL2CPP compiler settings to produce more informative managed stack traces for easier debugging. |
| **[Additional IL2CPP compiler arguments](handling-il2cpp-additional-args.md)** | Supply custom additional arguments to the IL2CPP compiler from code. |
| **[Linux IL2CPP cross-compiler](linux-il2cpp-crosscompiler.md)** | Use the Linux IL2CPP cross-compiler to build Linux IL2CPP Players on any Standalone platform without needing to use the Linux Unity Editor or rely on Mono. |
| **[IL2CPP limitations](scripting-restrictions.md)** | Understand the restrictions that apply on some AOT platforms when using IL2CPP as your scripting back end. |

## Additional resources

* [Handling IL2CPP additional arguments](handling-il2cpp-additional-args.md)
* [Linux IL2CPP cross-compiler](linux-il2cpp-crosscompiler.md)
* [Managed stack traces](il2cpp-managed-stack-traces.md)