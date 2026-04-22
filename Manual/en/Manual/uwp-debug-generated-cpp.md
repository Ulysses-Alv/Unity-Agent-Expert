# Debug generated C++ code

You can debug generated C++ code for your Universal Windows Platform (UWP) application using Visual Studio.

## Understand class and method naming in generated C++ code

### IL2CPP classes

IL2CPP classes follow the format of `<ClassName>_t#number`, where:

* `<ClassName>` is the class name
* The optional `#number` is a unique type number

Example **IL2CPP** classes:

* `String_t` - [String class](https://learn.microsoft.com/en-us/dotnet/api/system.string)
* `Object_t` - [Object class](https://learn.microsoft.com/en-us/dotnet/api/system.object)
* `Type_t` - [Type class](https://learn.microsoft.com/en-us/dotnet/api/system.type)
* `StringBuilder_t26` - [StringBuilder class](https://learn.microsoft.com/en-us/dotnet/api/system.text.stringbuilder)

### IL2CPP methods

IL2CPP methods follow the format of `<ClassName>_<MethodName>_m#number`, where:

* `<ClassName>` is the class name of the method’s declaring type
* `<MethodName>` is the method name
* `#number` is a unique method number

Example IL2CPP methods:

* `ConfigurationSection_DoDeserializeSection_m1275` - [DeserializationSection method of the ConfigurationSection class](https://learn.microsoft.com/en-us/dotnet/api/system.configuration.configurationsection.deserializesection)
* `String_Format_m4102` - [Format method of the String class](https://learn.microsoft.com/en-us/dotnet/api/system.string.format)
* `Mathf_Sqrt_m289` - [Sqrt method of the Mathf class](https://learn.microsoft.com/en-us/dotnet/api/system.mathf.sqrt)

### IL2CPP static field structures

Static fields structures follow the format of `<ClassName>_t#number_StaticFields`, where the first part of the structure name is the same as the declaring type.

Example static fields structures:

* `StringBuilder_t26_StaticFields`
* `Thing_t24_StaticFields`

### C++ comments

Preceding each class or method definition, C++ automatically generates a comment stating the full class or method name.

Example C++ comment:

```
    // System.Text.StringBuilder
    struct StringBuilder_t26  : public Object_t
    {
        // System.Int32 System.Text.StringBuilder::_length
        int32_t length_1;
        // System.Int32 System.Text.StringBuilder::_maxCapacity
        int32_t maxCapacity_2;
    };
```

## Observe variable values

You can debug your generated C++ code by observing the values of your variables [using the Visual Studio debugger](https://learn.microsoft.com/en-us/visualstudio/gamedev/unity/get-started/using-visual-studio-tools-for-unity?pivots=windows#unity-debugging).

You can [set breakpoints in Visual Studio](managed-code-debugging.md#BreakpointsVS) where you want the debugger to stop so you can observe your variables. Visual Studio allows you to observe your variable values by either mousing over the variable or [using a Watch window](https://learn.microsoft.com/en-us/visualstudio/debugger/watch-and-quickwatch-windows?).

## Observe a static field

In IL2CPP, Unity stores static fields in an Il2CppClass instance. To observe the values of the static field, you need to:

1. Find the [pointer](https://learn.microsoft.com/en-us/cpp/cpp/pointers-cpp) to the Il2CppClass structure of that type in your code.
     
   **Note:** These pointers are in scope of the methods that use them, but after observing it once, it will remain at the same memory address during the application run.
2. Retrieve the value of the `static_fields` field from that pointer. This is a pointer to a memory block containing static fields for that particular type.
3. Cast the value to the corresponding static field structure.
4. Observe the value in [the Visual Studio debugger](https://learn.microsoft.com/en-us/visualstudio/gamedev/unity/get-started/using-visual-studio-tools-for-unity?pivots=windows#unity-debugging).

## Investigate exceptions

IL2CPP uses native C++ exceptions to implement .NET exceptions.

To investigate the exceptions in your code, you can:

* Inspect the exception objects in the [watch window](https://learn.microsoft.com/en-us/visualstudio/debugger/watch-and-quickwatch-windows?)
* [Enable the debugger to break](https://learn.microsoft.com/en-us/visualstudio/debugger/managing-exceptions-with-the-debugger) on C++ exceptions

## Additional resources

* [Microsoft documentation on how to Debug C++ code using Visual Studio](https://learn.microsoft.com/en-us/visualstudio/debugger/getting-started-with-the-debugger-cpp?)
* [Microsoft documentation on Unity debugging with Visual Studio](https://learn.microsoft.com/en-us/visualstudio/gamedev/unity/get-started/using-visual-studio-tools-for-unity?pivots=windows#unity-debugging)
* [IL2CPP overview](./scripting-backends-il2cpp.md)
* [IL2CPP scripting backend for UWP](uwp-il2cpp-scripting.md)