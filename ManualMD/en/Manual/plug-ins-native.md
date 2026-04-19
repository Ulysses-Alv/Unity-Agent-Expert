# Native plug-ins

Unity supports native **plug-ins**, which are libraries of native code you can write in languages such as C, C++, and Objective-C. Plug-ins allow the code you write in C# to call functions from these libraries. This feature allows Unity to integrate with middleware libraries or existing C/C++ code.

A **native plug-in** provides a simple native interface, which the C# script then exposes to your other **scripts**. Unity exposes a set of native APIs that allow your plug-ins to respond to low-level Engine events in systems such as rendering, profiling, and logging. For more information, refer to [Native plug-in APIs](native-plugin-interface.md).

For an example of a native plug-in, refer to [Native Renderer Plugin](https://github.com/Unity-Technologies/NativeRenderingPlugin).

## Using a native plug-in

To use a native plug-in:

1. Write functions in a C-based language to access the features you need.
2. Compile them into a library.
3. In Unity, create a C# script that calls functions in the native library.

You build native plug-ins with native code compilers on the target platform. Because plug-in functions use a C-based call interface, you must declare the functions with C linkage to avoid name mangling issues.

For platform-specific information on native plug-ins, refer to the relevant section under [Platform development](PlatformSpecific.md).

## Example

A simple native library with a single function might have code that looks like this:

`float ExamplePluginFunction () { return 5.0F; }`

To access this code from within Unity, use the following C# script:

```
using UnityEngine;
using System.Runtime.InteropServices;

class ExampleScript : MonoBehaviour {
    #if UNITY_IPHONE
    // On iOS plugins are statically linked into
    // the executable, so we have to use __Internal as the
    // library name.
    [DllImport ("__Internal")]
    #else
    // Other platforms load plugins dynamically, so pass the
    // name of the plugin's dynamic library.
    [DllImport ("PluginName")]   
    #endif
    private static extern float ExamplePluginFunction ();

    void Awake () {
        // Calls the ExamplePluginFunction inside the plugin
        // And prints 5 to the console
        print (ExamplePluginFunction ());
       }
    }
```

## Additional resources

* [Low-level native plug-in interface](native-plugin-interface.md)
* [Mono Interop with native libraries](https://www.mono-project.com/docs/advanced/pinvoke/)
* [P-invoke documentation on MSDN](https://docs.microsoft.com/en-us/dotnet/framework/interop/marshaling-data-with-platform-invoke?redirectedfrom=MSDN)