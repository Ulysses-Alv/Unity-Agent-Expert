# Web Player settings

Use **Player** settings to know how Unity [builds](webgl-building.md) and displays your final Web application. For a description of the general **Player** settings, refer to [Player](class-PlayerSettings.md#general) settings.

**Note**: The [Facebook Instant Games platform](instant-games.md) shares these settings.

To access Web **Player settings**:

* From the Unity main menu, go to **Edit** > **Project Settings** > **Player**. The **Player** settings window appears.
* Select the **Web** tab to view the Web Player settings.

![Web Player settings](../uploads/Main/class-PlayerSettingsWebGL.png)

Web Player settings

You can find documentation for the properties in the following sections:

* [Resolution and Presentation](#Resolution)
* [Splash Image](#Splash)
* [Other Settings](#Other)
* [Publishing Settings](#Publishing)

**Note:** Although the **Icon** panel appears on the Web **Player** settings, there are no icon settings available because Web games don’t use icons.

For more information about Web **Publishing Settings**, refer to [Web Building and Running](webgl-building.md).

## Resolution and Presentation

Use the **Resolution and Presentation** section to customize the aspects of the screen’s appearance in the [Resolution](#resolution) section.

![Resolution section for the Web Player platform](../uploads/Main/WebGLResolutionandPresentationWindow.png)

Resolution section for the Web Player platform

### Resolution

You can customize the screen mode and default size of the Web canvas element by editing the following options:

| **Setting** | **Function** |
| --- | --- |
| **Default Canvas Width** | Set the width of the Web canvas element. |
| **Default Canvas Height** | Set the height of the Web canvas element. |
| **Run In Background** | Enable this option to allow your content to continue to run when the canvas or the browser window loses focus. |

### Web template

Select a template to use for your Web Project:

* The **Default** page is a simple white page with a loading bar on a grey canvas.
* The **Minimal** page has only the necessary boilerplate code to run the Web content.
* The **PWA** page has a **Progressive Web App** including a web manifest file and service worker code.

You can use your own template to run your game in an environment similar to the finished game using the instructions in [Web templates](webgl-templates.md).

## Splash Image

Use the **Virtual Reality Splash Image** setting to select a custom splash image for [XR](XR.md) displays. For information on common Splash Screen settings, refer to [Splash Screen](class-PlayerSettingsSplashScreen.md).

![Splash screen settings for virtual reality.](../uploads/Main/PlayerSetPCSplash.png)

Splash screen settings for virtual reality.

## Other Settings

This section allows you to customize a range of options organized into the following groups:

* [Rendering](#Rendering)
* [Configuration](#Configuration)
* [Shader Variant Loading](#ShaderVariantLoading)
* [Optimization](#Optimization)
* [Stack Trace](#Logging)
* [Legacy](#Legacy)

### Rendering

Use these settings to customize how Unity renders your game for the Web platform.

![Rendering Player settings for the Web platform](../uploads/Main/PlayerSetWebGLOther-Rendering.png)

Rendering Player settings for the Web platform

| **Setting** | **Function** |
| --- | --- |
| **Color Space** | Choose which color space to use for rendering. For more information, refer to [Color spaces](color-spaces-landing.md).  * **Gamma**: Gamma color space is typically used for calculating lighting on older hardware restricted to 8 bits per channel for the framebuffer format. Even though monitors today are digital, they might still take a gamma-encoded signal as input. * **Linear**: Linear color space rendering gives more precise results. When you select to work in linear color space, the Editor defaults to using [sRGB](https://en.wikipedia.org/wiki/SRGB) sampling. If your [Textures](Textures.md) are in linear color space, you need to work in linear color space and deactivate sRGB sampling for each Texture. |
| **Auto Graphics API** | Enables the default Web graphics API (WebGL 2.0). By default, this option is enabled.   Disable this option to display the **Graphics API** list where you can manually pick and reorder the graphics APIs. |
| **Graphics API** | Prioritize the first API on this list as the default web graphics API for your project. If a browser doesn’t support that API or there’s another issue, Unity tries the next API on the list (if one exists) as a fallback.  Click and drag entries on the list to reorder them. For more information about graphics APIs for Web, refer to [Web graphics APIs](webgl-graphics.md).  This property is available only when **Auto Graphics API** is disabled. |
| **Static Batching** | Enable this option to use [static batching](static-batching). |
| **Dynamic Batching** | Use [Dynamic Batching](DrawCallBatching.md) on your build (enabled by default).  **Note:** Dynamic batching has no effect when a [Scriptable Render Pipeline](scriptable-render-pipeline-introduction.md) is active, so this setting is only visible if the **Scriptable Render Pipeline Asset** [Graphics](class-GraphicsSettings.md#SRLoop) setting is blank. |
| **Graphics Jobs (Experimental)** | Enable this option to instruct Unity to offload graphics tasks (render loops) to worker threads running on other CPU cores. This is intended to reduce the time spent in `Camera.Render` on the main thread, which is often a bottleneck.   **Note:** This feature is experimental. It might not deliver a performance improvement for your project, and might introduce new crashes. |
| **Texture compression format** | Choose **DXT**, **ETC2**, or **ASTC** to set the texture compression format for the Web platform. For information on how to pick the right format, refer to [Texture compression format overview](texture-compression-formats.md). To create builds for desktop and mobile browsers from a script, refer to [Texture Compression in Web](webgl-texture-compression.md). |
| **Lightmap Encoding** | Defines the encoding scheme and compression format of the lightmaps.  You can choose from **Low Quality**, **Normal Quality**, or **High Quality** |
| **HDR Cubemap Encoding** | Defines the encoding scheme and compression format of the HDR Cubemaps.  You can choose from **Low Quality**, **Normal Quality**, or **High Quality**. For more information, refer to [Lightmaps: Technical information](Lightmaps-TechnicalInformation.md). |
| **Lightmap Streaming** | Enable this option to use [Mipmap Streaming](TextureStreaming.md) for lightmaps. Unity applies this setting to all lightmaps when it generates them.  **Note:** To use this setting, you must enable the [Texture Mipmap Streaming Quality](class-QualitySettings.md#texStream) setting. |
| **Streaming Priority** | Set the priority for all lightmaps in the [Mipmap Streaming system](TextureStreaming.md). Unity applies this setting to all lightmaps when it generates them. Positive numbers give higher priority. Valid values range from `-128` to `127`. |
| **Frame Timing Stats** | Enable this property to gather CPU and GPU frame timing data using [FrameTimingManager](frame-timing-manager.md) API. If you disable this property, [Dynamic Resolution](DynamicResolution-landing.md) camera setting can’t use this data to dynamically adjust the resolution to reduce GPU workload. |
| **Virtual Texturing (Experimental)** | Reduce GPU memory usage and texture loading times if your Scene has many high resolution textures. For more information, refer to [Virtual Texturing](svt-streaming-virtual-texturing.md). **Note**: The Unity Editor requires a restart for this property to take effect. |
| **Shader Precision Model** | Select the default precision of samplers and the definition of `half` used in shaders. For more information, refer to [ShaderPrecisionModel](../ScriptReference/ShaderPrecisionModel.md). |
| **360 Stereo Capture** | Indicate whether Unity can capture stereoscopic 360 images and videos. When enabled, Unity compiles additional shader variants to support 360 capture (currently only on Windows/OSX). The `enable_360_capture` keyword is added during the [`RenderToCubemap`](../ScriptReference/Camera.RenderToCubemap.md) call, but isn’t triggered outside of this function. |

### Configuration

![Configuration settings for the Web platform](../uploads/Main/PlayerSetWebGLOther-Config.png)

Configuration settings for the Web platform

| **Property** | **Description** |
| --- | --- |
| **Scripting Backend** | Choose the scripting backend you want to use. The scripting backend determines how Unity compiles and executes C# code in your Project.  * **Mono**: Compiles C# code into .NET Common Intermediate Language (CIL) and executes that CIL using a Common Language Runtime. For more information, refer to [Mono](scripting-backends-mono.md). * **IL2CPP**: Compiles C# code into CIL, converts the CIL to C++ and then compiles that C++ into native machine code, which executes directly at runtime. For more information, refer to [IL2CPP](scripting-backends-il2cpp.md). |
| **API Compatibility Level** | Choose which .NET APIs you can use in your project. This setting can affect compatibility with third-party libraries. However, it has no effect on Editor-specific code (code in an Editor directory, or within an Editor-specific Assembly Definition).  **Tip:** If you’re having problems with a third-party assembly, you can try the suggestion in the [API Compatibility Level](#apiComp) section.    * **.Net Framework**: Compatible with the .NET Framework 4 (which includes everything in the .NET Standard 2.0 profile plus additional APIs). Choose this option when using libraries that access APIs not included in .NET Standard 2.0. Produces larger builds and any additional APIs available aren’t necessarily supported on all platforms. For more information, refer to [Referencing additional class library assemblies](dotnet-profile-assemblies.md). * **.Net Standard 2.1**: Produces smaller builds and has full cross-platform support. |
| **Editor Assemblies Compatibility Level** | Select which .NET APIs to use in your Editor assemblies.  * **.NET Framework**: Compatible with the .NET Framework 4 (which includes everything in the .NET Standard 2.1 profile plus additional APIs). Choose this option when using libraries that access APIs not included in .NET Standard 2.1. Produces larger builds and any additional APIs available aren’t necessarily supported on all platforms. For more information, refer to [Referencing additional class library assemblies](dotnet-profile-assemblies.md). * **.NET Standard**: Compatible with [.NET Standard 2.1](https://docs.microsoft.com/en-us/dotnet/standard/net-standard). Produces smaller builds and has full cross-platform support. |
| **IL2CPP Code Generation** | Defines how Unity manages IL2CPP code generation.  **Note**: To use this, set **Scripting Backend** to **IL2CPP**. |
| **C++ Compiler Configuration** | Choose the C++ compiler configuration used when compiling IL2CPP generated code.  * **Debug**: The debug configuration turns off all optimizations, which makes the code quicker to build but slower to run. * **Release**: The release configuration enables optimizations so that the compiled code runs faster, the binary size is reduced, but compilation takes longer. * **Master**: Master configuration enables all possible optimizations, squeezing every bit of performance possible. For instance, on platforms that use the MSVC++ compiler, this option enables link-time code generation. Compiling code using this configuration can take significantly longer than it does using the Release configuration. The recommended best practice is to build the shipping version of your game using the Master configuration if the increase in build time is acceptable. |
| **Use incremental GC** | Uses the incremental garbage collector, which spreads garbage collection over several frames to reduce garbage collection-related spikes in frame duration. For more information, refer to [Automatic Memory Management](performance-managed-memory.md). |
| **Allow downloads over HTTP** | Indicate whether to allow downloading content over HTTP. The default option is **Not allowed** due to the recommended protocol being HTTPS, which is more secure.  * **Not Allowed**: Never allow downloads over HTTP. * **Allowed in Development Builds**: Only allow downloads over HTTP in development builds. * **Always Allowed**: Allow downloads over HTTP in development and release builds. |
| **Active Input Handling** | Choose how to handle input from users.  * **Input Manager (Old)**: Use the original [Input](class-InputManager.md) settings. * **Input System Package (New)**: Uses the new [input system package](https://docs.unity3d.com/Packages/com.unity.inputsystem@latest). * **Both**: Use both systems. |

#### Shader Variant Loading

Use these settings to control how much memory **shaders** use at runtime.

| **Property** | **Description** |
| --- | --- |
| **Default chunk size (MB)** | Sets the maximum size of compressed shader variant data chunks Unity stores in your built application for all platforms. The default is `16`. For more information, refer to [Shader loading](shader-loading.md). |
| **Default chunk count** | Sets the default limit on how many decompressed chunks Unity keeps in memory on all platforms. The default is `0`, which means there’s no limit. |
| **Override** | Enables overriding **Default chunk size** and **Default chunk count** for this build target. |
| **Chunk size (MB)** | Overrides the value of **Default chunk size (MB)** on this build target. |
| **Chunk count** | Overrides the value of **Default chunk count** on this build target. |

#### API Compatibility Level

You can choose your mono API compatibility level for all targets. Sometimes a third-party .NET library uses functionality that’s outside of your .NET compatibility level. To understand what is going on in such cases, and how to best fix it, try following these suggestions:

1. Install [ILSpy](https://www.microsoft.com/en-us/p/ilspy/9mxfbkfvsq13?cid=msft_web_chart&activetab=pivot:overviewtab) for Windows.
2. Drag the .NET assemblies for the API compatibility level that you’re having issues with into ILSpy. You can find the problematic files under `Frameworks/Mono/lib/mono/YOURSUBSET/`.
3. Drag in your third-party assembly.
4. Right-click your third-party assembly and select **Analyze**.
5. In the analysis report, inspect the **Depends on** section. The report highlights anything that the third-party assembly depends on, but that’s not available in the .NET compatibility level of your choice in red.

### Script Compilation

![Script compilation settings for the Web platform](../uploads/Main/ProjectSettingsScriptCompilation.png)

Script compilation settings for the Web platform

| **Property** | **Description** |
| --- | --- |
| **Scripting Define Symbols** | Sets custom compilation flags.    For more details, refer to [Platform dependent compilation](platform-dependent-compilation.md). |
| **Additional Compiler Arguments** | Adds entries to this list to pass additional arguments to the Roslyn compiler. Use one new entry for each additional argument. To create a new entry, click **Add** (**+**). To remove an entry, click **Remove** (**-**).  When you have added all desired arguments, click **Apply** to include your additional arguments in future compilations. Click **Revert** to reset this list to the most recent applied state. |
| **Suppress Common Warnings** | Indicates whether to display the C# warnings [CS0169](https://docs.microsoft.com/en-us/dotnet/csharp/misc/cs0169) and [CS0649](https://docs.microsoft.com/en-us/dotnet/csharp/misc/CS0649). |
| **Allow ‘unsafe’ Code** | Enables support for compiling [‘unsafe’ C# code](https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/keywords/unsafe) in a pre-defined assembly (for example, `Assembly-CSharp.dll`).  For Assembly Definition Files (`.asmdef`), click on one of your `.asmdef` files and enable the option in the Inspector window that appears. |
| **Use Deterministic Compilation** | Indicates whether to prevent compilation with the -deterministic C# flag. With this setting enabled, compiled assemblies are byte-for-byte the same each time they’re compiled.  For more information, refer to [C# Compiler Options that control code generation](https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/compiler-options/deterministic-compiler-option). |

### Optimization

![Optimization settings for the Web platform](../uploads/Main/PlayerSetWebGLOther-Optimization.png)

Optimization settings for the Web platform

| **Property** | **Description** |
| --- | --- |
| **Prebake Collision Meshes** | Adds collision data to [Meshes](class-Mesh.md) at build time. |
| **Preloaded Assets** | Sets an array of Assets for the player to load on startup.  To add new Assets, increase the value of the **Size** property and then set a reference to the Asset to load in the new **Element** box that appears. |
| **Strip Engine Code** | Enable this option if you want the Unity Linker tool to remove code for Unity Engine features that your Project doesn’t use. This setting is only available with the [IL2CPP scripting backend](IL2CPP). Most apps do not use every available DLL. This option strips out DLLs that your app doesn’t use to reduce the size of the built Player. If your app is using one or more classes that would normally be stripped out under your current settings, Unity displays a debug message when you try to build the app. |
| **Managed Stripping Level** | Choose how aggressively Unity strips unused managed (C#) code. When Unity builds your app, the Unity Linker process can strip unused code from the managed DLLs your Project uses. Stripping code can make the resulting executable smaller, but can sometimes remove code that’s in use.   For more information about these options and bytecode stripping with IL2CPP, refer to [ManagedStrippingLevel](../ScriptReference/ManagedStrippingLevel.md).    * **Minimal**: Use this to strip class libraries, UnityEngine, Windows Runtime assemblies, and copy all other assemblies. * **Low**: Remove unreachable managed code to reduce build size and Mono/IL2CPP build times. * **Medium**: Run UnityLinker to reduce code size beyond what **Low** can achieve. You might need to support a custom link.xml file, and some reflection code paths might not behave the same. * **High**: UnityLinker will strip as much code as possible. This will further reduce code size beyond what Medium can achieve but managed code debugging of some methods might no longer work. You might need to support a custom link.xml file, and some reflection code paths might not behave the same. |
| **Enable Internal Profiler (Deprecated)** | This feature is deprecated and will be retired in a future version of Unity. Use the [Profiler window](ProfilerWindow.md) instead (menu: **Window > Analytics > Profiler**).  The Profiler collects application performance data and prints a report to the console. The report contains the number of milliseconds each Unity subsystem takes to execute on each frame, averaged across 30 frames. |
| **Vertex Compression** | Sets vertex compression per channel. This affects all the meshes in your project.  Typically, Vertex Compression is used to reduce the size of mesh data in memory, reduce file size, and improve GPU performance.   For more information on how to configure vertex compression and limitations of this setting, refer to [Compressing mesh data](compressing-mesh-data-optimization.md). |
| **Optimize Mesh Data** | Enable this option to strip unused vertex attributes from the mesh used in a build. This option reduces the amount of data in the mesh, which can help reduce build size, loading times, and runtime memory usage.    **Warning:** If you have this setting enabled, don’t change material or shader settings at runtime.    For more information, refer to [PlayerSettings.stripUnusedMeshComponents](../ScriptReference/PlayerSettings-stripUnusedMeshComponents.md). |
| **Texture Mipmap Stripping** | Enables mipmap stripping for all platforms. It strips unused mipmap levels from Textures at build time.  Unity determines unused mipmap levels by comparing the mipmap level against the quality settings for the current platform. If a mipmap level is excluded from every quality setting for the current platform, then Unity strips those mipmap levels from the build at build time. If `QualitySettings.globalTextureMipmapLimit` is set to a mipmap level that has been stripped, Unity will set the value to the closest mipmap level that hasn’t been stripped. |

### Stack Trace

Select the logging settings for the Web platform.

![Logging settings for the Web platform](../uploads/Main/PlayerSetPCOther-Logging.png)

Logging settings for the Web platform

Select your preferred stack trace method by enabling the option that corresponds to each Log Type (**Error**, **Assert**, **Warning**, **Log**, and **Exception**) based on the type of logging you require. For more information, refer to [stack trace logging](stack-trace.md).

| **Property** | **Description** |
| --- | --- |
| **None** | No logs are ever recorded. |
| **ScriptOnly** | Logs only when running **scripts**. |
| **Full** | Logs all the time. |

### Legacy

Enable the **Clamp BlendShapes (Deprecated)** option to clamp the range of blend shape weights in [Skinned Mesh Renderers](class-SkinnedMeshRenderer.md).

![Legacy settings for the Web platform](../uploads/Main/PlayerSetPCOther-Legacy.png)

Legacy settings for the Web platform

## Publishing settings

Use the Publishing Settings to configure how Unity builds your Web platform application. For example, you can choose to enable the browser cache to store its files in your build.

![Publishing settings for the Web platform](../uploads/Main/WebGLBuilding-PublishingSettings.png)

Publishing settings for the Web platform

| **Property** | **Description** |
| --- | --- |
| **Compression Format** | Choose the **compression** format to use for release build files. The options are: Gzip, Brotli, and Disabled (none). This option doesn’t affect **development builds**. |
| **Name Files As Hashes** | Enable this option to use an MD5 hash of the uncompressed file contents as a file name for each file in the build. |
| **Data caching** | Enable this option to automatically cache your contents Asset data on the user’s machine so it doesn’t have to be re-downloaded on subsequent runs (unless the contents have changed).  Caching is implemented using the IndexedDB API provided by the browser. Some browsers might implement restrictions around this, such as asking the user for permission to cache data over a specific size. |
| **Debug Symbols** | Select from the available options to specify how debug symbols are added in your build.  * **Off**: Select this option if you don’t want any debug symbols to be added in your build. * **External**: Select this option to store the debug symbols in a separate file that you can download from the server when an error occurs. It’s recommended to choose this option for release builds. * **Embedded**: Select this option to embed the debug symbols in a WASM file. This option helps preserve debug symbols and perform demangling (displaying the original function names) of the stack trace when an error occurs. By default, demangling support is available for this option. |
| **Decompression Fallback** | Include decompression fallback code for build files in the loader. Use this option if you’re unable to configure server response headers according to the selected compression method. |
| **Power Preference** | Set a preference for which GPU to use when rendering on a multi-GPU device. Note that the browser might ignore this preference.  * **Default**: Let the browser choose the GPU configuration. * **Low Performance**: Request that the browser use the integrated GPU to prioritize power savings. * **High Performance**: Request that the browser use the external GPU to prioritize rendering performance. |

### WebAssembly Language Features

Use this section to customize the WebAssembly language features for your Web application.

| **Property** | **Description** |
| --- | --- |
| **Enable Exceptions** | Choose how to handle unexpected code behavior (generally considered errors) at runtime.  * **None**: Any exception thrown causes your content to stop with an error. Select this option if you don’t need any exception support. It gives the best performance and smallest builds. * **Explicitly Thrown Exceptions Only** (default): Capture exceptions that are explicitly specified from a `throw` statement in your scripts and make sure `finally` blocks are called. This option makes the generated JavaScript code from your scripts longer and slower. This might only be an issue if scripts are the main bottleneck in your project. * **Full Without Stacktrace**: Capture exceptions that are explicitly specified from `throw` statements in your scripts, Null References, and Out of Bounds Array accesses. * **Full With Stacktrace**:Same as **Full Without Stacktrace**, but also captures stack traces. Unity generates these exceptions by embedding checks for them in the code, so this option decreases performance and increases browser memory usage. Only use this option for debugging, and always test in a 64-bit browser. |
| **Enable Native C/C++ Multithreading** | Enable this option to use the native Unity C/C++ engine code that targets WebAssembly/`SharedArrayBuffer` multithreading (experimental). If this setting is enabled, the server configuration needs to set Cross-Origin Opener Policy (COOP), Cross-Origin Embedded Policy (COEP), and Cross-Origin Resource Policy (CORP) headers. For examples of how to set these headers, refer to [Server configuration code samples](webgl-server-configuration-code-samples.md). Note that not all browsers support `SharedArrayBuffer`.  This setting doesn’t enable multithreading C# code, as that requires further advances to the WebAssembly language standard. It’s recommended to disable this option and use it only for evaluating future web features. |
| **Enable WebAssembly 2023** | If enabled, the generated WebAssembly code targets WebAssembly 2023, which is a Unity-coined name for a selection of newer WebAssembly language features. These features include: sign-extension opcodes, non-trapping fp-to-int instructions, bulk memory, JS BigInt integration, `WebAssembly.Table`, native WebAssembly exceptions, and SIMD (requires Chrome ≥ 91 (May 2021), Firefox ≥ 89 (June 2021) or Safari ≥ 16.4 (March 2023)). If disabled, targets the original WebAssembly **MVP** **feature set**. |
| **Use WebAssembly.Table** | Enable this option if you want the Web build to target the `WebAssembly.Table` language feature for faster JS-Wasm interop and build times. When this option is disabled, the Web build targets the older deprecated Emscripten `-sDYNCALLS` model for backwards compatibility with older Unity Web JS plug-ins. It’s recommended to enable this option for new projects that don’t use any old incompatible JavaScript plug-ins, and to disable it if you’re using `.jslib` files that rely on the old `dynCall()` mechanism.   **Note:** `WebAssembly.Table` isn’t backwards compatible. When targeting `WebAssembly.Table`, use the Emscripten runtime function `getWasmTableEntry(functionPtr)` to convert a function pointer to a callable JS function, or the Emscripten code template `{{{ makeDynCall('sig', 'variableName') }}}(...args);` to make a function pointer call. The earlier Emscripten runtime function family `dynCall_sig(ptr, ...args);` is no longer supported with `WebAssembly.Table`. For code examples, refer to the Emscripten documentation. |
| **Enable BigInt** | Enable this option to target the `WebAssembly.BigInt` language feature and use the `BigInt` type in WebAssembly, which results in faster build times and slightly smaller code size. When this option is disabled, the `BigInt` type isn’t available in WebAssembly. The generated WebAssembly code relies on the BigInt ABI for function signatures containing 64-bit variables. Disable this option if you want to target old browsers that don’t support the Wasm BigInt feature. It’s recommended to enable this option for new projects, fast build iteration times, and to disable it if targeting backward compatibility with older browsers is important.   **Note:** The Wasm BigInt feature requires at least Chrome 85 (Aug 25, 2020), Firefox 78 (Jun 30, 2020), Safari 14.5 (Apr 26, 2021), or newer. |
| **Initial Memory Size** | The initial size of the WASM heap memory in megabytes (MB). By default, this is set to 32 MB. If **Memory Growth Mode** is set to **None**, then this is also the maximum size of WASM heap memory. |
| **Memory Growth Mode** | Choose the growth mode for the WASM heap memory from the following options. The recommended option is **Geometric**.  * **None**: The WASM heap memory has a fixed size configured in **Initial Memory Size**. * **Linear**: The WASM heap memory increases by a fixed amount configured by **Linear Memory Growth Step**. * **Geometric**: The WASM heap memory increases relative to the current heap size depending on the factor configured in **Geometric Memory Growth Step** and **Geometric Memory Growth Cap**. |
| **Maximum Memory Size** | The maximum size of the WASM heap memory in MB. By default, this is set to 2048 MB, which is the recommended setting. You can enter a memory size up to 4096 MB, but there are known Firefox and Chrome bugs for builds over 2048 MB. This option is only available for the **Memory Growth Mode** **Linear** or **Geometric**. |
| **Linear Memory Growth Step** | Advanced tuning option to control the WASM heap growth step in MB. By default, this is set to 16 MB. A growth step of 16 MB indicates that the heap is increased by 16 MB each time it needs to grow. Only available if **Memory Growth Mode** is set to **Linear**. |
| **Geometric Memory Growth Step** | Advanced tuning option to control the WASM heap growth factor relative to the current heap size. By default, this is set to 0.2. A growth factor of 0.2 means that the size of the heap is increased by 0.2 \* currentHeapSize each time the heap needs to grow. Only available if **Memory Growth Mode** is set to **Geometric**. |
| **Geometric Memory Growth Cap (MB)** | Advanced tuning option to control the upper limit for a heap growth step in MB. By default, this is set to 96 MB. A growth cap of 96MB means that the size of the heap is increased by at most 96 MB. Only available if **Memory Growth Mode** is set to **Geometric**. |
| **Enable Submodule Stripping Compatibility** | If enabled, produces a build that can be used with the submodule stripping functionality of the Web Stripping Tool package. If disabled, produces a build that is not intended to be submodule stripped. It is best to leave this option disabled whenever not using the Web Stripping Tool package, or generated code size can be adversely affected. |

### Show Diagnostic Overlay setting

To help optimize Web builds and diagnose potential problems, you can view diagnostics information (currently limited to memory usage) by enabling this setting. Once enabled, an icon appears on the build that displays an overlay with data about the build. It’s available for both Development and Release builds.

1. To view the diagnostics information, enable the **Show Diagnostics Overlay** option in the Player settings window (**File > Build Profiles > Player Settings > Publishing Settings**).

   On desktop, the Diagnostics icon appears on the footer of the Web canvas:

   ![Diagnostics button on the footer of the canvas](../uploads/Main/webgl-diagnostic-desktop.png)

   Diagnostics button on the footer of the canvas

   On a mobile device, the Diagnostics icon appears on the bottom-right of the screen:

   ![An Android phone displaying the diagnostics button](../uploads/Main/webgl-diagnostic-mobile.png)

   An Android phone displaying the diagnostics button
2. Click the **Diagnostics** icon. An overlay appears showing the JavaScript memory, which is further broken down to display WASM heap memory usage:

![The Diagnostics overlay shows the JavaScript memory distribution and the WASM heap memory usage](../uploads/Main/webgl-diagnostic-overlay.png)

The Diagnostics overlay shows the JavaScript memory distribution and the WASM heap memory usage

The following diagnostics appear on the overlay screen:

| **Property** | **Function** |
| --- | --- |
| **Total JS Memory** | The current size of the JavaScript (JS) heap, including unused memory not allocated to any JS objects in megabytes.  * **Used JS Memory**:Memory in use by JS objects in megabytes. |
| **Total WASM heap memory** | Linear memory representing the entire heap of the C/C++ Unity engine that’s compiled with Emscripten, including unallocated memory in megabytes.  * **Used WASM heap**: The space of the WASM heap that’s allocated in megabytes. |
| **Page Load Time to First Frame** | The total time from the beginning of page load until the first application frame rendering is complete in milliseconds.  * **Page Load Time**: The time it takes from page load to get to the beginning of rendering the first frame, including downloads, compilation, parsing, and the main application in milliseconds. * **Code download time**: The time it takes for the build to download the code file in milliseconds. * **Load time of asset file(.data)**: The time it takes for the build to download the .data file binary, in milliseconds. * **WebAssembly startup time**: The time it takes between loading the JavaScript framework and reaching Unity’s C++ main(). This is close to compilation time of WASM, in milliseconds. * **Game startup time**: The time it takes to finish executing Unity’s C++ main() to the first frame of main, which typically contains the loading of the first game scene in milliseconds. |
| **Average FPS (10 s)** | The average of last 10 frames per second.  * **Current frames per second**: The number of frames rendered on screen per second. |
| **Number of Frame Stalls** | The number of rendered frames that took unusually long to complete compared to their previous frames. |

### Important note about JS Memory

The JS Memory information is obtained using the [performance.memory API](https://developer.mozilla.org/en-US/docs/Web/API/Performance/memory), which is currently only supported on Chrome or Edge. There are no other APIs available that return this information for Safari or Firefox.

**Note:** The `performance.memory` API isn’t supported on iOS devices.

On browsers where this API isn’t supported, a message showing N/A appears.

![A N/A message appears on browsers that dont support performance.memory API](../uploads/Main/webgl-wasm-memory.png)

A N/A message appears on browsers that don’t support performance.memory API

## Additional resources:

* [Texture Compression in Web](webgl-texture-compression.md)
* [Web templates](webgl-templates.md)
* [Splash Screen](class-PlayerSettingsSplashScreen.md)
* [Texture compression format overview](texture-compression-formats.md)

PlayerSettingsWebGL