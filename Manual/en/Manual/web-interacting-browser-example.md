# Create callbacks between Unity C#, JavaScript, and C/C++/C# code

Callbacks are an important part of web development because they let your Unity project communicate better with your JavaScript browser, JavaScript **plug-in**, and C/C++/C# plug-ins. You might want to call functions from your C/C++/C# or JavaScript code in your Unity C# code, and callbacks are an effective way to do this.

## Use the makeDynCall function to make callbacks

Use `makeDynCall` to call C, C++, or C# callbacks from JavaScript code. For example, the following JavaScript code calls the function `callback`:

```
{{{ makeDynCall('vii', 'callback') }}}(1, 2)
```

* `'vii'` is the signature of the parameters. For example, the callback function is a `void` type (v), and the `(1,2)` are both integers (ii). This signature tells the underlying Emscripten toolchain what data types the function returns and accepts as parameters. For a list of all supported type characters, refer to the [dyncall library documentation](https://dyncall.org/docs/manual/manualse4.html).
* `'callback'` is the name of the callback function.
* `(1, 2)` are the arguments to pass into the function.

The following code is another example of code that does a callback with the function “onresponse” and passes in a float (f) and multiple integer type values:

```
{{{ makeDynCall('vfii', 'onresponse') }}}(2.8, arg, 0);
```

## Use callbacks in your scripts

The following **scripts** show an example of how to use callback functions to communicate between JavaScript, C++, and C# code.

### Javascript plug-in code

The following code is an example of a `.jslib` file that schedules a timeout and invokes a callback.

```
mergeInto(LibraryManager.library, {
    JsSetTimeout: function (message, timeout, callback) {
        // Create copy of message because it might be deleted before callback is run
        var stringMessage = UTF8ToString(message);
        var buffer = stringToNewUTF8(stringMessage);
        setTimeout(function () {
            {{{ makeDynCall('vi', 'callback') }}} (buffer);
            _free(buffer);
        }, timeout);
    }
});
```

### C++ plug-in code

The following code is an example of a C++ file that:
\* calls the external JavaScript function (`JsSetTimeout`)
\* contains the callback function the JavaScript calls when a timeout happens.

```
// CallbackExample.cpp
#include <stdio.h>

typedef void (*OnTimeOutCallback)(char* message);

extern "C"
{
    void JsSetTimeout(char* message, int timeout, OnTimeOutCallback callback);
}

void CCallback(char* message)
{
    printf("C callback received \"%s\"\n", message);
}

extern "C" void SetCTimeout()
{
    JsSetTimeout("Hello World", 1000, CCallback);
}
```

### Unity C# code

The following code is an example of C# Unity code that receives the messages from the C++ and JavaScript callbacks.

```
using AOT;
using System;
using System.Text;
using System.Runtime.InteropServices;
using UnityEngine;

public class CallbackExample : MonoBehaviour
{
    void Awake()
    {
        SetCSharpTimeout();
        SetCTimeout();
    }

    void SetCSharpTimeout()
    {
        JsSetTimeout("Hello World", 500, CSSharpCallback);
    }

    [DllImport("__Internal")]
    public static extern void SetCTimeout();    
    
    [DllImport("__Internal")]
    public static extern void JsSetTimeout(string message, int timeout, Action<string> action);

    [MonoPInvokeCallback(typeof(Action<string>))]
    public static void CSSharpCallback(string message)
    {
       Debug.Log($"C# callback received \"{message}\"");
    }
}
```

## Additional resources

* [Interaction with browser scripting](webgl-interactingwithbrowserscripting.md)
* [Set up your JavaScript plug-in](web-interacting-browser-js.md)
* [Call JavaScript functions from Unity C# scripts](web-interacting-browser-js-to-unity.md)
* [Call Unity C# script functions from JavaScript](web-interacting-browser-unity-to-js.md)
* [Call C/C++/C# functions from Unity C# scripts](web-interacting-browsers-c-to-unity.md)
* [Replace deprecated browser interaction code](web-interacting-browser-deprecated.md)