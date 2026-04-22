# IL2CPP runtime code checks

You can use the `[Il2CppSetOption]` C# attribute and its `Option` parameter to control which safety checks the **IL2CPP** compiler includes in the C++ code it generates.

The `[Il2CppSetOption]` attribute is not part of the standard Unity Editor and Engine public APIs but its source is shipped separately as part of your Unity installation. To use the `[Il2CppSetOption]` attribute:

1. In the directory where your Unity version is installed, navigate to the `Data\il2cpp` directory on Windows, or the `Contents/Frameworks/il2cpp` directory on macOS.
2. Find the `Il2CppSetOptionAttribute.cs` source file.
3. Copy the source file into your project’s `Assets` folder.

The supported options for the attribute are as follows:

| **Property** | **Description** | **Default** |
| --- | --- | --- |
| **Null checks** | It’s recommended to keep this option enabled. When disabled, IL2CPP generates C++ without null checks and won’t throw managed `NullReferenceException` exceptions. This can improve runtime performance, but attempts to access null values won’t be prevented, which can lead to unintended behavior including crashes. | Enabled |
| **Array bounds checks** | It’s recommended to keep this option enabled. When disabled, IL2CPP generates C++ without array bounds checks and won’t throw managed `IndexOutOfRangeException` exceptions. This might improve runtime performance but allows code to access an array with invalid indices, which can lead to incorrect behavior, including reading from or writing to arbitrary memory locations. In most cases, these memory accesses occur without any immediate side effects, and can corrupt the state of the application with no obvious warning signs, making the errors very hard to debug. | Enabled |
| **Divide by zero checks** | Keep this option disabled unless you need to run divide by zero checks. When enabled, IL2CPP generates C++ that contains divide by zero checks for integer division and throws managed `DivideByZeroException` exceptions as necessary. If this option is disabled, IL2CPP doesn’t emit the divide by zero checks on integer division into the generated C++ code.   These checks have an impact on performance at runtime. | Disabled |

The following example shows how to use the `[Il2CppSetOption]` attribute:

```
[Il2CppSetOption(Option.NullChecks, false)]
public static string MethodWithNullChecksDisabled()
{
    var tmp = new object();
    return tmp.ToString();
}
```

You can apply `[Il2CppSetOption]` to assemblies, types, methods, and properties. Unity uses the attribute from the most local scope.

```
[Il2CppSetOption(Option.NullChecks, false)]
public class TypeWithNullChecksDisabled
{
    public static string AnyMethod()
    {
        // Unity doesn’t perform null checks in this method.
        var tmp = new object();
        return tmp.ToString();
    }

    [Il2CppSetOption(Option.NullChecks, true)]
    public static string MethodWithNullChecksEnabled()
    {
        // Unity performs null checks in this method.
        var tmp = new object();
        return tmp.ToString();
    }
}

public class SomeType
{
    [Il2CppSetOption(Option.NullChecks, false)]
    public string PropertyWithNullChecksDisabled
    {
        get
        {
            // Unity doesn't perform null checks here.
            var tmp = new object();
            return tmp.ToString();
        }
        set
        {
            // Unity doesn't perform null checks here.
            value.ToString();
        }
    }

    public string PropertyWithNullChecksDisabledOnGetterOnly
    {
        [Il2CppSetOption(Option.NullChecks, false)]
        get
        {
            // Unity doesn’t perform null checks here.
            var tmp = new object();
            return tmp.ToString();
        }
        set
        {
            // Unity performs null checks here.
            value.ToString();
        }
    }
}
```

## Additional resources

* [IL2CPP managed stack trace reference](il2cpp-managed-stack-traces.md)
* [IL2CPP additional arguments reference](handling-il2cpp-additional-args.md)