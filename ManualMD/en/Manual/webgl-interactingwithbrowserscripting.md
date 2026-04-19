# Interaction with browser scripting

When you build content for the web, you might need to communicate with other elements on your web page or use Web APIs to implement functionality that Unity doesn’t expose by default.

In both cases, you need to directly interface with the browser’s JavaScript engine. Unity Web provides different methods to handle these interactions.

| Topic | Description |
| --- | --- |
| [Code examples: Call JavaScript and C/C++/C# functions in Unity](web-interacting-code-example.md) | Code examples that show interactions between Unity, JavaScript, and C-based code. |
| [Set up your JavaScript plug-in](web-interacting-browser-js.md) | Create a JavaScript **plug-in** that your Unity project can interact with. |
| [Customize error handling](web-interacting-browser-error-handling.md) | Customize how Unity handles errors in your Web builds. |
| [Call JavaScript functions from Unity C# scripts](web-interacting-browser-js-to-unity.md) | Call functions from your JavaScript plug-in or browser in your Unity project. |
| [Call Unity C# script functions from JavaScript](web-interacting-browser-unity-to-js.md) | Call functions from your Unity project in your JavaScript plug-in or browser. |
| [Call C/C++/C# functions from Unity C# scripts](web-interacting-browsers-c-to-unity.md) | Call functions from your C or C++ code in your Unity project. |
| [Compile a static library as a Unity plug-in](web-interacting-browsers-library.md) | Call functions from a static library. |
| [Create callbacks between Unity C#, JavaScript, and C/C++/C# code](web-interacting-browser-example.md) | Learn how to use callbacks to communicate between your plug-in, browser, and Unity project. |
| [JavaScript interface in Unity Web builds](web-js-interface.md) | Learn about some useful functions you can use in the JavaScript interface. |
| [Replace deprecated browser interaction code](web-interacting-browser-deprecated.md) | Replace any deprecated code with the updated code. |

## Additional resources

* [Web native plug-ins for Emscripten](webgl-native-plugins-with-emscripten.md)