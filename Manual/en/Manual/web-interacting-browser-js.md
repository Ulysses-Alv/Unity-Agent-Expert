# Set up your JavaScript plug-in

You can call functions from your JavaScript plug-ins in your Unity projects. Unity supports two JavaScript **plug-in** file types that let you add JavaScript code to your Unity project:

* .jslib
* .jspre

**Note**: Unity currently supports [ECMAScript 5 (ES5) syntax](https://www.w3schools.com/js/js_es5.asp) in .jslib and .jspre files. [ES6 syntax](https://www.w3schools.com/js/js_es6.asp) isn’t yet supported.

If you want to call functions from C++ plug-ins instead, refer to [Call C/C++/C# functions from Unity C# scripts](web-interacting-browsers-c-to-unity.md).

## Import your JavaScript files into your Unity project

The recommended way to use browser JavaScript in your project is to add your JavaScript sources (.jspre and .jslib files) to your project, and then use those functions or libraries directly in your Unity C# script code.

Place your JavaScript plug-in files in any folder, such as `Assets/JSPlugins`.

## Call functions from the .jslib file type

You can call functions from your .jslib plug-in files in your Unity C# or native **scripts**.

The .jslib file type uses the `--js-library` Emscripten module. For more information, refer to the Emscripten documentation about the [–js-library Emscripten option](https://emscripten.org/docs/tools_reference/emcc.html#emcc-js-library).

The following code shows an example of a .jslib plug-in file with the ideal syntax that defines some functions (`Hello` and `HelloString`).

```
mergeInto(LibraryManager.library, {

  Hello: function () {
    window.alert("Hello, world!");
  },

  HelloString: function (str) {
    window.alert(UTF8ToString(str));
  },
});
```

You can then call these functions in your Unity C# code:

```
using UnityEngine;
using System.Runtime.InteropServices;

public class NewBehaviourScript : MonoBehaviour {

    [DllImport("__Internal")]
    private static extern void Hello();

    [DllImport("__Internal")]
    private static extern void HelloString(string str);

    void Start() {
        Hello();
        HelloString("This is a string.");
    }
}
```

For a full example of code interactions between Unity C# and JavaScript functions, refer to [Code examples: Call JavaScript and C/C++/C# functions in Unity](web-interacting-code-example.md).

For more information about interactions between Unity C# and JavaScript, refer to [Call JavaScript functions from Unity C# scripts](web-interacting-browser-js-to-unity.md#example-unity-code).

## Include JavaScript libraries with the .jspre file type

Use the .jspre plug-in file type to include existing JavaScript libraries in your JavaScript code. You can’t use Unity code to interact with the .jspre files, but you can use them in the .jslib code.

The .jspre file type uses the `--pre-js` Emscripten option. For more information, refer to the Emscripten documentation about the [–pre-js Emscripten option](https://emscripten.org/docs/tools_reference/emcc.html#emcc-pre-js).

During the build process, Emscripten creates the `*.framework.js` file and copies the contents of the .jspre file into the start of the `*.framework.js` file. This process is useful because all the code is combined into one file so it’s easier to manage the files and the code benefits from Emscripten optimizations.

## Additional resources

* [Interaction with browser scripting](webgl-interactingwithbrowserscripting.md)
* [Call JavaScript functions from Unity C# scripts](web-interacting-browser-js-to-unity.md)
* [Call Unity C# script functions from JavaScript](web-interacting-browser-unity-to-js.md)
* [Call C/C++/C# functions from Unity C# scripts](web-interacting-browsers-c-to-unity.md)
* [Create callbacks between Unity C#, JavaScript, and C/C++/C# code](web-interacting-browser-example.md)