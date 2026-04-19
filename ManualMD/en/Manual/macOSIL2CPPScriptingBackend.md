# Use IL2CPP with macOS

Describes the use of **plug-ins** for macOS applications using Intermediate Language To C++ (./scripting-backends-il2cpp).

IL2CPP is a fully supported scripting back end that you can use as an alternative to Mono when building projects for macOS Player.

When you use IL2CPP to build a project, Unity converts Intermediate Language (IL) code from **scripts** and assemblies to C++ before creating a native binary. Refer to [IL2CPP](./scripting-backends-il2cpp.md) for more information.

## C++ source code plug-ins for IL2CPP

You can add C++ (.cpp) code files directly into a Unity Project when using the [IL2CPP](./scripting-backends-il2cpp.md) scripting back end. These C++ files act as plug-ins within the **Inspector**. If you configure the C++ files to be compatible with macOS Player, Unity compiles them together with C++ code that gets generated from managed assemblies. Refer to [Import and configure plug-ins](plug-in-inspector.md) for more information on plug-in configuration.

Generated C++ code links the functions together, removing the need for a separate Dynamic Link Library (DLL). Instead of using the DLL name, use the `"__Internal"` keyword to make the C++ linker responsible for resolving functions rather than loading them at runtime. For example:

```
[DllImport("__Internal")]
private static extern int
CountLettersInString([MarshalAs(UnmanagedType.LPWStr)]string str);
```

You can define this kind of function in NativeFunctions.cpp as follows:

```
extern "C" __declspec(dllexport) int __stdcall CountLettersInString(wchar_t* str)
{
    int length = 0;
    while (*str++ != nullptr)
        length++;
    return length;
}
```

When the linker resolves the function call, errors in the function declaration on the managed side result in a linker error rather than a runtime error. This means that no dynamic loading needs to happen during runtime, with the function called directly from C#. This decreases the performance overhead of a `P/Invoke` call.

## Additional resources

* [Import and configure plug-ins](plug-in-inspector.md)