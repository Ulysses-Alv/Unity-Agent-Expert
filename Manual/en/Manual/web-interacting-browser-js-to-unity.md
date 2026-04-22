# Call JavaScript functions from Unity C# scripts

You can use functions from your JavaScript **plug-ins** in your Unity C# code. It can be useful to use JavaScript code in Unity because you might need to communicate with other elements on your web page or Web APIs.

To learn about the file types and how to set up a JavaScript plug-in for interaction with Unity **scripts**, refer to [Set up your JavaScript plug-in](web-interacting-browser-js.md). To learn how to interact with C/C++/C# plug-ins instead, refer to [Call C/C++/C# functions from Unity C# scripts](web-interacting-browsers-c-to-unity.md).

| Topic | Description |
| --- | --- |
| [Pass different variables from JavaScript to Unity](#pass-different-variables) | Tips on how to pass different variables between Unity and JavaScript. |
| [Execute build code in its own scope](#make-code-visible) | Tips on how to improve visibility of code between Unity and JavaScript. |
| [Example Unity C# code that calls JavaScript functions](#example-unity-code) | This example code shows how to call JavaScript functions in Unity. |
| [Use Unity plug-ins as a reference](#plugin-ref) | Shows a list of plug-ins included with Unity you can use as a reference. |

## Pass different variables from JavaScript to Unity

To integrate JavaScript with Unity, you need efficient communication between the two. Use the following tips on how to pass various types of data from JavaScript to Unity.

### Numeric types

You can pass simple numeric types to JavaScript in function parameters without converting them.

For example, this function in JavaScript:
`lang-js
AddNumbers: function (x, y) {
return x + y;
},`

accepts the integer values from C# without conversions:

```
int result = AddNumbers(5, 7);
```

You can pass other data types as a pointer in the Emscripten heap. The Emscripten heap is just a large array in JavaScript memory.

### Strings

To convert a string to a JavaScript string, use the `UTF8ToString` helper function.

```
var stringMessage = UTF8ToString("Hello World");
```

To return a string value, call `_malloc` to allocate some memory and the `stringToUTF8` helper function to write a JavaScript string to the memory. If the string is a return value, then the **IL2CPP** runtime automatically frees up the memory for you.

```
 var returnStr = "Hello World";
    var bufferSize = lengthBytesUTF8(returnStr) + 1;
    var buffer = _malloc(bufferSize);
    stringToUTF8(returnStr, buffer, bufferSize);
```

### Arrays

For arrays of primitive types, Emscripten provides different `ArrayBufferViews` into its heap for different sizes of integer, unsigned integer or floating point representations of memory: `HEAP8`, `HEAPU8`, `HEAP16`, `HEAPU16`, `HEAP32`, `HEAPU32`, `HEAPF32`, `HEAPF64`.

The following function loops through the `HEAPF32` array and outputs the values at each index.

```
PrintFloatArray: function (array, size) {
    for(var i = 0; i < size; i++)
    console.log(HEAPF32[(array >> 2) + i]);
  },
```

### Textures

To access a texture in **WebGL**, Emscripten provides the `GL.textures` array which maps native texture IDs from Unity to WebGL texture objects. You can call WebGL functions on Emscripten’s WebGL context, `GLctx`.

For example, the following function binds a WebGL texture from the GL texture array to a 2D texture.

```
 BindWebGLTexture: function (texture) {
    GLctx.bindTexture(GLctx.TEXTURE_2D, GL.textures[texture]);
  },
```

## Execute build code in its own scope

It is recommended that you execute all build code in its own scope. If you have your code in its own scope, you can embed your content on a page without causing conflicts with the embedding page code, and you can embed more than one build on the same page.

### Code in .jslib plug-ins

If you have all your JavaScript code in the form of `.jslib` plug-ins inside your project, then this JavaScript code will run inside the same scope as the compiled build and your code should work the same way as in previous versions of Unity. Some of the objects and functions directly visible from the JavaScript plug-in code include:
\* `Module`
\* `SendMessage`
\* `HEAP8`
\* `ccall`

### Call JavaScript functions from global scope

If you want to call the internal JavaScript functions from the global scope of the embedded page, in your Web template `index.html`, you must use the `unityInstance` variable.

Use `unityInstance` after the Unity engine instantiation succeeds, for example:

```
  var MyGameInstance = null;
  script.onload = () => {
    createUnityInstance(canvas, config, (progress) => { /*...*/ }).then((unityInstance) => {
      MyGameInstance = unityInstance;
```

Then, you can use `MyGameInstance.SendMessage()` to send a message to the build, or use `MyGameInstance.Module` to access the build module object.

## Example Unity C# code that calls JavaScript functions

The following JavaScript code creates a function named `Hello`. For this example, use this code in your JavaScript plug-in so you can call it in your Unity C# scripts.

```
mergeInto(LibraryManager.library, {

  Hello: function () {
    window.alert("Hello, world!");
  },
});
```

Then, use this C# code in your Unity project to call the `Hello` function from your JavaScript code:

```
using UnityEngine;
using System.Runtime.InteropServices;

public class NewBehaviourScript : MonoBehaviour {

    [DllImport("__Internal")]
    private static extern void Hello();

    void Start() {
        Hello();
    }
}
```

For a larger code example with varied function types, refer to [Code examples: Call JavaScript and C/C++/C# functions in Unity](web-interacting-code-example.md).

## Use Unity plug-ins as a reference

There are several plug-ins in the Unity installation folder that you can use as reference:

* `PlaybackEngines/WebGLSupport/BuildTools/lib`
* `PlaybackEngines/WebGLSupport/BuildTools/Emscripten/src/library*`

## Additional resources

* [Interaction with browser scripting](webgl-interactingwithbrowserscripting.md)
* [Set up your JavaScript plug-in](web-interacting-browser-js.md)
* [Call Unity C# script functions from JavaScript](web-interacting-browser-unity-to-js.md)
* [Call C/C++/C# functions from Unity C# scripts](web-interacting-browsers-c-to-unity.md)
* [Create callbacks between Unity C#, JavaScript, and C/C++/C# code](web-interacting-browser-example.md)