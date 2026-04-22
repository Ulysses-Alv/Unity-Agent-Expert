# Debug and troubleshoot Web builds

Visual Studio doesn’t support debugging Unity Web content. Use the following tips to get your build information.

## Debug your build in a browser’s JavaScript console

The Unity Web platform doesn’t have access to your file system, so it doesn’t write a log file like other platforms. However, it does write all logging information such as `Debug.Log`, `Console.WriteLine` or Unity’s internal logging to the browser’s JavaScript console.

To open the JavaScript console:

| **OS** | **Browser** | **Instructions** |
| --- | --- | --- |
| Windows | Firefox | Press Ctrl-Shift-K. |
| Windows | Chrome | Press Ctrl-Shift-J. |
| Windows | Microsoft Edge | Press F12. |
| Windows | Internet Explorer | Press F12. |
| Mac | Firefox | Press Command-Option-K. |
| Mac | Chrome | Press Command-Option-J. |
| Mac | Safari | 1. Go to **Preferences** > **Advanced** > **Develop**. 2. Press Command-Option-C. |

## Create a development build to debug

You might want to make a **development build** in Unity to debug your code. To make a development build:

1. Open the [Build Profiles window](PublishingBuilds).
2. Enable **Development Build**.

Development builds allow you to connect the profiler. Unity doesn’t [minify](https://en.wikipedia.org/wiki/Minification_%28programming%29) the code, so the emitted JavaScript code still contains human-readable, [C++-mangled](https://en.wikipedia.org/wiki/Name_mangling#Name_mangling_in_C.2B.2B), function names.

The browser uses these to display stack traces if you run into a browser error, when using `Debug.LogError`, or when an exception occurs and exception support is disabled. Unlike the managed stack traces that occur when you have full exception support, these stack traces have mangled names, and contain managed code and the internal Unity Engine code.

## Exception support

Web has different levels of exception support, but by default, Unity Web only supports explicitly thrown exceptions. For more information, refer to [Web Player settings](class-PlayerSettingsWebGL.md#wasm-language-features). You can enable **Full** exception support, which emits additional checks in the IL2CPP-generated code, to catch access to null references and out-of-bounds array elements in your managed code. These additional checks significantly impact performance and increase code size and load times, so you must only use it for debugging.

Full exception support also emits function names to generate stack traces for your managed code. For this reason, stack traces appear in the console for uncaught exceptions and for `Debug.Log` statements. Use `System.Environment.StackTrace` to get a stack trace string.

## Troubleshooting

### Problem: The build runs out of memory

This is a common problem, especially on 32-bit browsers. For more information on Web memory issues and how to fix them, refer to the documentation on [Memory in Unity Web](webgl-memory.md).

### Error message: Incorrect header check

The browser console log usually prints this error due to incorrect server configuration. For more information on how to deploy a release build, refer to documentation on [Deploying compressed builds](webgl-deploying.md).

### Error message: Decompressing this format (1) isn’t supported on this platform

The browser console log prints this error when the content tries to load an AssetBundle compressed using LZMA, which Unity Web doesn’t support. Re-compress the AssetBundle using LZ4 **compression** to solve this problem. For more information on compression for Web, refer to documentation on [Web building](webgl-building.md), particularly the **AssetBundles** section.

## Additional resources:

* [Web Development](webgl-develop.md)
* [Web performance considerations](webgl-performance.md)
* [Build and distribute a Web application](webgl-building-distribution.md)