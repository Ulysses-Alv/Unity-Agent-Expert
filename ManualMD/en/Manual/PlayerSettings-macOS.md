# macOS Player settings reference

This page details the **Player** settings specific to macOS. For a description of the general **Player** settings, refer to [Player Settings](class-PlayerSettings.md#general).

You can find documentation for the properties in the following sections:

* [Icon](#Icon)
* [Resolution and Presentation](#Resolution)
* [Splash Image](#Splash)
* [Other Settings](#Other)

## Icon

Activate the **Override for Windows, Mac, Linux** setting to assign a custom icon for your desktop game. You can upload different sizes of the icon to fit each of the squares provided.

## Resolution and presentation

Use the **Resolution and Presentation** section to customize aspects of the screen’s appearance.

### Resolution

This section allows you to customize the screen mode and default size.

| **Property** | **Description** |
| --- | --- |
| **Run In Background** | Enable this option to allow the application to run in the background when it loses focus. When disabled, the application pauses when it loses focus. |
| **Fullscreen Mode** | Choose the full-screen mode. This defines the default window mode at startup.    * **Fullscreen Window**: Set your app window to the full-screen native display resolution, covering the whole screen. This mode is also known as borderless full-screen. Unity renders the app content at the resolution set by a script, or the native display resolution if none is set and scales it to fill the window. When scaling, Unity adds black bars to the rendered output to match the display aspect ratio to prevent content stretching. This process is called [letterboxing](https://en.wikipedia.org/wiki/Letterboxing_(filming)). The OS overlay UI displays on top of the full-screen window (such as IME input windows). All platforms support this mode. * **Exclusive Fullscreen (Windows only)**: Set your app to maintain sole full-screen use of a display. Unlike **Fullscreen Window**, this mode changes the OS resolution of the display to match the app’s chosen resolution. This option is only supported on Windows. * **Maximized Window (Windows and Mac only)**: Set the app window to the operating system’s definition of **maximized**. On Windows, it is a full-screen window with a title bar. On macOS, it is a full-screen window with a hidden menu bar and dock. **Fullscreen Window** is the default setting for other platforms. * **Windowed**: Set your app to a standard, non-full-screen movable window, the size of which is dependent on the app resolution. In this mode, the window is resizable by default. Use the [Resizable Window](#resizable) setting to disable this. All desktop platforms support this full-screen mode. |
| **Default Is Native Resolution** | Enable this option to make the game use the default resolution used on the target machine.  **Note**: This property isn’t visible if you set **Fullscreen Mode** to **Windowed**. |
| **Default Screen Width** | Set the default width of the game screen in pixels.   **Note**: This property is visible only if you set **Fullscreen Mode** to **Windowed**. |
| **Default Screen Height** | Set the default height of the game screen in pixels.   **Note**: This property is visible only if you set **Fullscreen Mode** to **Windowed**. |
| **Mac Retina Support** | Enable this option to enable support for high DPI (Retina) screens on a Mac. Unity enables this by default. This enhances Projects on a Retina display, but it’s somewhat resource-intensive when active. |

### Standalone Player Options

Use this section to specify the settings to customize the screen. For example, you can set options for users to resize the screen and specify the number of instances that can run concurrently.

| **Property** | **Description** |
| --- | --- |
| **Use Player Log** | Activate this option to write a log file with debugging information. **Warning:** If you plan to submit your application to the Mac App Store, leave this option deactivated. For more information, refer to [Identification](#AppStore). |
| **Resizable Window** | Activate this option to allow resizing of the desktop player window. **Note:** If you deactivate this option, your application can’t use the **Windowed** [Fullscreen Mode](#fullscreen). |
| **Visible in Background** | Activate this option to display the application in the background when using **Windowed** [Fullscreen Mode](#fullscreen). This option isn’t supported on macOS. |
| **Allow Fullscreen Switch** | Activate this option to allow default OS full-screen key presses to toggle between full-screen and windowed modes. |
| **Force Single Instance** | Activate this option to restrict desktop players to a single concurrent running instance. |
| **Use DXGI flip model swap chain for D3D11** | Using the flip model ensures the best performance. This setting affects the D3D11 graphics API. Deactivate this option to fall back to the Windows 7-style BitBlt model. For more information, refer to [PlayerSettings.useFlipModelSwapchain](../ScriptReference/PlayerSettings-useFlipModelSwapchain.md). |

## Splash Image

Use the **Virtual Reality Splash Image** setting to select a custom splash image for [Virtual Reality](XR.md) displays. For information on common Splash Screen settings, refer to [Splash Screen](class-PlayerSettingsSplashScreen.md).

## Other Settings

* [Rendering](#Rendering)
* [Identification](#AppStore)
* [Configuration](#Configuration)
* [Shader Settings](#ShaderSettings)
* [Shader Variant Loading Settings](#ShaderVariantLoading)
* [Script Compilation](#ScriptCompilation)
* [Optimization](#Optimization)
* [Stack Trace](#stack-trace)
* [Legacy](#Legacy)
* [Capture Logs](#capture-logs)

### Rendering

Use these settings to customize how Unity renders your game for desktop platforms.

| **Property** | **Description** |
| --- | --- |
| **Color Space** | Choose which color space to use for rendering. For more information, refer to [Color spaces](color-spaces-landing.md).    * **Gamma**: Gamma color space is typically used for calculating lighting on older hardware restricted to 8 bits per channel for the framebuffer format. Even though monitors today are digital, they might still take a gamma-encoded signal as input. * **Linear**: Linear color space rendering gives more precise results. When you select to work in linear color space, the Editor defaults to using [sRGB](https://en.wikipedia.org/wiki/SRGB) sampling. If your [Textures](Textures.md) are in linear color space, you need to work in linear color space and deactivate sRGB sampling for each Texture. |
| **MSAA Fallback** | Select the multi sample antialiasing fallback strategy to upgrade or downgrade the sample count if the sample count requested by the user is not supported by the device.    * **Upgrade**: The sample count reduces to the nearest supported lower sample count. * **Downgrade**: The sample count increases to the next higher sample count. If that sample count is not supported, then it reduces to the nearest supported lower sample count. |
| **Auto Graphics API for Windows** | Enable this option to use the recommended Graphics API on the Windows machine the game is running on. Disable it to add and remove supported Graphics APIs. |
| **Auto Graphics API for Mac** | Enable this option to use the recommended Graphics API on the Mac the game is running on. Disable it to add and remove supported Graphics APIs. Windows doesn’t support this property. |
| **Auto Graphics API for Linux** | Enable this option to use the recommended Graphics API on the Linux machine it runs on. Disable it to add and remove supported Graphics APIs. |
| **Metal API Validation** | Enable this option to use in-Editor tool to debug Shader issues. On a draw call, the tool checks a shaders expected textures and buffers against the attached and confirms its compatibility. The results of the checks are available in the console.   **Metal API Validation** doesn’t enable Apple’s [Metal API validation](https://developer.apple.com/documentation/xcode/validating-your-apps-metal-api-usage), but you can enable it by setting **Run in Xcode as** to **Debug** in the **Build Profiles** window for [macOS](macosbuildsettings.md) and [iOS](BuildSettingsiOS.md).  **Note**: This setting only appears when running the Editor on macOS. |
| **Metal Write-Only Backbuffer** | Allow improved performance in non-default device orientation. This sets the `frameBufferOnly` flag on the back buffer, which prevents readback from the back buffer and enables some driver optimization. |
| **Force hard shadows on Metal** | Enable this option to force Unity to use point sampling for shadows on Metal. This reduces shadow quality, which improves performance. |
| **Memoryless Depth** | Choose when to use [memoryless render textures](../ScriptReference/RenderTextureMemoryless.Depth.md). Memoryless render textures are temporarily stored in the on-tile memory when rendered, not in CPU or GPU memory. This reduces memory usage of your app but you can’t read or write to these render textures.  **Note**: Memoryless render textures are only supported on iOS, tvOS 10.0+, Metal, and Vulkan. Render textures are read/write protected and stored in CPU or GPU memory on other platforms.    * **Unused**: Never use memoryless framebuffer depth. * **Forced**: Always use memoryless framebuffer depth. * **Automatic**: Let Unity decide when to use memoryless framebuffer depth. |
| **Multithreaded Rendering** | Enable this option to move graphics API calls from Unity’s main thread to a separate worker thread. This can help to improve performance in applications that have high CPU usage on the main thread. |
| **Static Batching** | Enable this option to use [static batching](static-batching). |
| **Dynamic Batching** | Use [Dynamic Batching](DrawCallBatching.md) on your build (enabled by default).  **Note:** Dynamic batching has no effect when a [Scriptable Render Pipeline](scriptable-render-pipeline-introduction.md) is active, so this setting is only visible if the **Scriptable Render Pipeline Asset** [Graphics](class-GraphicsSettings.md#SRLoop) setting is blank. |
| **Sprite Batching Threshold** | Controls the maximum vertex threshold used when batching. |
| **GPU Skinning** | Calculate mesh skinning and blend shapes on the GPU via shaders to free up CPU resources and improve performance.    * **CPU**: Select this option to perform mesh skinning and blend shapes calculation on the CPU. * **GPU**: Select this option to perform mesh skinning and blend shapes calculation on the GPU. * **GPU (Batched)**: Select this option to use batching and reordering to perform mesh skinning and blend shapes calculation on the GPU. Batching reduces the number of dispatch calls to the GPU and can improve performance. |
| **Graphics Jobs (Experimental)** | Enable this option to instruct Unity to offload graphics tasks (render loops) to worker threads running on other CPU cores. This is intended to reduce the time spent in `Camera.Render` on the main thread, which is often a bottleneck.   **Note:** This feature is experimental. It might not deliver a performance improvement for your project, and might introduce new crashes. |
| **Lightmap Encoding** | Defines the encoding scheme and compression format of the lightmaps.  You can choose from **Low Quality**, **Normal Quality**, or **High Quality** |
| **HDR Cubemap Encoding** | Defines the encoding scheme and compression format of the HDR Cubemaps.  You can choose from **Low Quality**, **Normal Quality**, or **High Quality**. For more information, refer to [Lightmaps: Technical information](Lightmaps-TechnicalInformation.md). |
| **Lightmap Streaming** | Enable this option to use [Mipmap Streaming](TextureStreaming.md) for lightmaps. Unity applies this setting to all lightmaps when it generates them.  **Note:** To use this setting, you must enable the [Texture Mipmap Streaming Quality](class-QualitySettings.md#texStream) setting. |
| **Streaming Priority** | Set the priority for all lightmaps in the [Mipmap Streaming system](TextureStreaming.md). Unity applies this setting to all lightmaps when it generates them. Positive numbers give higher priority. Valid values range from `-128` to `127`. |
| **Frame Timing Stats** | Enable this property to gather CPU and GPU frame timing data using [FrameTimingManager](frame-timing-manager.md) API. If you disable this property, [Dynamic Resolution](DynamicResolution-landing.md) camera setting can’t use this data to dynamically adjust the resolution to reduce GPU workload. |
| **OpenGL: Profiler GPU Recorders** | Enable profiler recorders when rendering with OpenGL. |
| **Allow HDR Display Output** | Enable HDR mode output when the application runs. This only works on displays that support this feature. If the display doesn’t support HDR mode, the game runs in standard mode. |
| **Use HDR Display Output** | Checks if the main display supports HDR, and if it does, swaps to HDR output when the application launches.  **Note**: This option is available only when **Allow HDR Display Output** is active. |
| **Swap Chain Bit Depth** | Select the number of bits in each color channel for swap chain buffers. For more information on bit depth, refer to [D3DHDRDisplayBitDepth](../ScriptReference/D3DHDRDisplayBitDepth.md).  **Note**: This setting is visible only when **Use HDR Display Output** is enabled.    * **Bit Depth 10**: Unity uses the R10G10B10A2 buffer format and Rec2020 primaries with ST2084 PQ encoding. * **Bit Depth 16**: Unity uses the R16G16B16A16 buffer format and Rec709 primaries with linear color (no encoding). |
| **Virtual Texturing (Experimental)** | Reduce GPU memory usage and texture loading times if your Scene has many high resolution textures. For more information, refer to [Virtual Texturing](svt-streaming-virtual-texturing.md). **Note**: The Unity Editor requires a restart for this property to take effect. |
| **360 Stereo Capture** | Indicate whether Unity can capture stereoscopic 360 images and videos. When enabled, Unity compiles additional shader variants to support 360 capture (currently only on Windows/OSX). The `enable_360_capture` keyword is added during the [`RenderToCubemap`](../ScriptReference/Camera.RenderToCubemap.md) call, but isn’t triggered outside of this function. |
| **Load/Store Action Debug Mode** | Highlights undefined pixels that might cause rendering problems on mobile platforms. This affects the Unity Editor Game view, and your built application if you select **Development Build** in the **Platform Settings** section of the **Build Profiles** window. For more information, refer to [LoadStoreActionDebugModeSettings](../ScriptReference/Rendering.LoadStoreActionDebugModeSettings.md). |

### Identification

| **Property** | **Description** |
| --- | --- |
| **Override Default Bundle Identifier** | Indicates whether you can manually set the bundle identifier.  **Note:** This setting affects macOS, iOS, tvOS, and Android. |
| **Bundle Identifier** | Enter the Bundle Identifier of your application. This appears as `CFBundleIdentifier` in the associated `info.plist` file. The Bundle Identifier must follow the convention `com.YourCompanyName.YourProductName` and must contain only alphanumeric and hyphen characters. For more information, refer to [CFBundleIdentifier](https://developer.apple.com/library/content/documentation/General/Reference/InfoPlistKeyReference/Articles/CoreFoundationKeys.html#/apple_ref/doc/uid/20001431-102070).  **Important**: Unity automatically replaces any invalid characters you type with a hyphen. |
| **Build** | Enter the build number for this version of your app. This appears as `CFBundleVersion` in the associated `info.plist` file. For more information, refer to [CFBundleVersion](https://developer.apple.com/library/archive/documentation/General/Reference/InfoPlistKeyReference/Articles/CoreFoundationKeys.html#//apple_ref/doc/uid/20001431-102364). |
| **Category** | Enter the string corresponding to the app’s type. The App Store uses this string to select the appropriate categorization for the app. By default, this is `public.app-category.games`. For more information, refer to [LSApplicationCategoryType](https://developer.apple.com/library/archive/documentation/General/Reference/InfoPlistKeyReference/Articles/LaunchServicesKeys.html#//apple_ref/doc/uid/TP40009250-SW8). |
| **Mac App Store Validation** | Activate this so that your app only runs when it has a valid receipt from the Mac App Store. This prevents people from running the game on a different device. Only deactivate this setting if you have implemented your own receipt validation. |

### Publishing to the Mac App Store

The [Use Player Log](#UsePlayerLog) property creates a log file with debugging information, helping to investigate any problems with your game. Deactivate this when publishing games for Apple’s Mac App Store, as Apple can reject your submission if activated. For more information, refer to [Log Files](log-files.md).
The **Use Mac App Store Validation** property activates receipt validation for the Mac App Store. If activated, your game only runs when it has a valid receipt from the Mac App Store. Use this when submitting games to Apple for publishing on the App Store. This prevents people from running the game on a different computer.

**Note**: This feature doesn’t implement any strong copy protection. In particular, any potential crack for one Unity game can work for any other Unity content. For this reason, it’s recommended that you implement your own receipt validation code on top of this, using Unity’s **plug-in** feature. Because Apple requires plug-in validation to initially happen before showing the screen setup dialog, it’s recommended to activate this property to avoid Apple rejecting your submission.

### Configuration

| **Property** | **Description** |
| --- | --- |
| **Scripting Backend** | Choose the scripting backend you want to use. The scripting backend determines how Unity compiles and executes C# code in your Project.    * **IL2CPP**: Compiles C# code into CIL, converts the CIL to C++ and then compiles that C++ into native machine code, which executes directly at runtime. For more information, refer to [IL2CPP](scripting-backends-il2cpp.md). * **Mono**: Compiles C# code into .NET Common Intermediate Language (CIL) and executes that CIL using a Common Language Runtime. For more information, refer to [Mono](scripting-backends-mono.md). |
| **API Compatibility Level** | Choose which .NET APIs you can use in your project. This setting can affect compatibility with third-party libraries. However, it has no effect on Editor-specific code (code in an Editor directory, or within an Editor-specific Assembly Definition).  **Tip:** If you’re having problems with a third-party assembly, you can try the suggestion in the [API Compatibility Level](#apiComp) section.    * **.Net Framework**: Compatible with the .NET Framework 4 (which includes everything in the .NET Standard 2.0 profile plus additional APIs). Choose this option when using libraries that access APIs not included in .NET Standard 2.0. Produces larger builds and any additional APIs available aren’t necessarily supported on all platforms. For more information, refer to [Referencing additional class library assemblies](dotnet-profile-assemblies.md). * **.Net Standard 2.1**: Produces smaller builds and has full cross-platform support. |
| **Editor Assemblies Compatibility Level** | Select which .NET APIs to use in your Editor assemblies.    * **.NET Framework**: Compatible with the .NET Framework 4 (which includes everything in the .NET Standard 2.1 profile plus additional APIs). Choose this option when using libraries that access APIs not included in .NET Standard 2.1. Produces larger builds and any additional APIs available aren’t necessarily supported on all platforms. For more information, refer to [Referencing additional class library assemblies](dotnet-profile-assemblies.md). * **.NET Standard**: Compatible with [.NET Standard 2.1](https://docs.microsoft.com/en-us/dotnet/standard/net-standard). Produces smaller builds and has full cross-platform support. |
| **IL2CPP Code Generation** | Defines how Unity manages IL2CPP code generation.  **Note**: To use this, set **Scripting Backend** to **IL2CPP**. |
| **C++ Compiler Configuration** | Choose the C++ compiler configuration used when compiling IL2CPP generated code.    * **Debug**: The debug configuration turns off all optimizations, which makes the code quicker to build but slower to run. * **Release**: The release configuration enables optimizations so that the compiled code runs faster, the binary size is reduced, but compilation takes longer. * **Master**: Master configuration enables all possible optimizations, squeezing every bit of performance possible. For instance, on platforms that use the MSVC++ compiler, this option enables link-time code generation. Compiling code using this configuration can take significantly longer than it does using the Release configuration. The recommended best practice is to build the shipping version of your game using the Master configuration if the increase in build time is acceptable. |
| **IL2CPP Stacktrace Information** | Choose the information to include in a stack trace. For further details on the information types, refer to [Managed stack traces with IL2CPP](il2cpp-managed-stack-traces.md).    * **Method Name**: Include each managed method in the stack trace. * **Method Name, File Name, and Line Number**: Include each managed method with file and line number information in the stack trace.    **Note**: Using this option can increase both the build time and final size of the built program. |
| **Use incremental GC** | Uses the incremental garbage collector, which spreads garbage collection over several frames to reduce garbage collection-related spikes in frame duration. For more information, refer to [Automatic Memory Management](performance-managed-memory.md). |
| **Allow downloads over HTTP** | Indicate whether to allow downloading content over HTTP. The default option is **Not allowed** due to the recommended protocol being HTTPS, which is more secure.    * **Not Allowed**: Never allow downloads over HTTP. * **Allowed in Development Builds**: Only allow downloads over HTTP in development builds. * **Always Allowed**: Allow downloads over HTTP in development and release builds. |
| **Active Input Handling** | Choose how to handle input from users.    * **Input Manager (Old)**: Use the original [Input](class-InputManager.md) settings. * **Input System Package (New)**: Uses the new [input system package](https://docs.unity3d.com/Packages/com.unity.inputsystem@latest). * **Both**: Use both systems. |
| **Target minimum macOS Version** | Enter the minimum version of macOS that the application will run on. |
| **Camera Usage Description** | Enter the reason for accessing the camera on the device. |
| **Microphone Usage Description** | Enter the reason for accessing the microphone on the device. |
| **Bluetooth Usage Description** | Enter the reason for accessing the device’s Bluetooth connection. |
| **Supported URL schemes** | A list of [supported URL schemes](https://developer.apple.com/documentation/xcode/defining-a-custom-url-scheme-for-your-app/). To add new schemes, increase the value of the **Size** property then set a reference to the Asset to load in the **Element** field. |

#### API compatibility level

You can choose your mono API compatibility level for all targets. Sometimes a third-party .NET library uses functionality that’s outside of your .NET compatibility level. To understand what’s going on in such cases, and how to best fix it, try following these suggestions:

1. Install [ILSpy](https://www.microsoft.com/en-us/p/ilspy/9mxfbkfvsq13?cid=msft_web_chart&activetab=pivot:overviewtab) for Windows.
2. Drag the .NET assemblies for the API compatibility level that you are having issues with into ILSpy. You can find these under `Frameworks/Mono/lib/mono/YOURSUBSET/`.
3. Drag in your third-party assembly.
4. Right-click your third-party assembly and select **Analyze**.
5. In the analysis report, inspect the **Depends on** section. The report highlights anything that the third-party assembly depends on, but that’s not available in the .NET compatibility level of your choice in red.

### Shader Settings

| **Property** | **Description** | |
| --- | --- | --- |
| **Shader Precision Model** | Select the default precision shaders use. For more information, refer to [Use 16-bit precision in shaders](SL-Use16BitPrecisionInShaders.md).    * **Platform default**: Use lower precision on mobile platforms, and full precision on other platforms. * **Unified**: Use lower precision if the platform supports it. |
| **Strict shader variant matching** | Enable this option to use the error shader for rendering if a shader variant is missing in the Player build and display an error in the console. The error specifies the shader, subshader index, pass, and keywords used for shader variant search |
| **Keep Loaded Shaders Alive** | Keep all loaded shaders alive and prevent unloading. |

### Shader Variant Loading

Use these settings to control how much memory **shaders** use at runtime.

| **Property** | **Description** |
| --- | --- |
| **Default chunk size (MB)** | Sets the maximum size of compressed shader variant data chunks Unity stores in your built application for all platforms. The default is `16`. For more information, refer to [Shader loading](shader-loading.md). |
| **Default chunk count** | Sets the default limit on how many decompressed chunks Unity keeps in memory on all platforms. The default is `0`, which means there’s no limit. |
| **Override** | Enables overriding **Default chunk size** and **Default chunk count** for this build target. |
| **Chunk size (MB)** | Overrides the value of **Default chunk size (MB)** on this build target. |
| **Chunk count** | Overrides the value of **Default chunk count** on this build target. |

### Script compilation

| **Property** | **Description** |
| --- | --- |
| **Scripting Define Symbols** | Sets custom compilation flags.    For more details, refer to [Platform dependent compilation](platform-dependent-compilation.md). |
| **Additional Compiler Arguments** | Adds entries to this list to pass additional arguments to the Roslyn compiler. Use one new entry for each additional argument. To create a new entry, click **Add** (**+**). To remove an entry, click **Remove** (**-**).  When you have added all desired arguments, click **Apply** to include your additional arguments in future compilations. Click **Revert** to reset this list to the most recent applied state. |
| **Suppress Common Warnings** | Indicates whether to display the C# warnings [CS0169](https://docs.microsoft.com/en-us/dotnet/csharp/misc/cs0169) and [CS0649](https://docs.microsoft.com/en-us/dotnet/csharp/misc/CS0649). |
| **Allow ‘unsafe’ Code** | Enables support for compiling [‘unsafe’ C# code](https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/keywords/unsafe) in a pre-defined assembly (for example, `Assembly-CSharp.dll`).  For Assembly Definition Files (`.asmdef`), click on one of your `.asmdef` files and enable the option in the Inspector window that appears. |
| **Use Deterministic Compilation** | Indicates whether to prevent compilation with the -deterministic C# flag. With this setting enabled, compiled assemblies are byte-for-byte the same each time they’re compiled.  For more information, refer to [C# Compiler Options that control code generation](https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/compiler-options/deterministic-compiler-option). |

### Optimization

| **Property** | **Description** |
| --- | --- |
| **Prebake Collision Meshes** | Adds collision data to [Meshes](class-Mesh.md) at build time. |
| **Preloaded Assets** | Sets an array of Assets for the player to load on startup.  To add new Assets, increase the value of the **Size** property and then set a reference to the Asset to load in the new **Element** box that appears. |
| **Managed Stripping Level** | Choose how aggressively Unity strips unused managed (C#) code. When Unity builds your app, the Unity Linker process can strip unused code from the managed DLLs your Project uses. Stripping code can make the resulting executable smaller, but can sometimes remove code that’s in use.   For more information about these options and bytecode stripping with IL2CPP, refer to [ManagedStrippingLevel](../ScriptReference/ManagedStrippingLevel.md).    * **Minimal**: Use this to strip class libraries, UnityEngine, Windows Runtime assemblies, and copy all other assemblies. * **Low**: Remove unreachable managed code to reduce build size and Mono/IL2CPP build times. * **Medium**: Run UnityLinker to reduce code size beyond what **Low** can achieve. You might need to support a custom link.xml file, and some reflection code paths might not behave the same. * **High**: UnityLinker will strip as much code as possible. This will further reduce code size beyond what Medium can achieve but managed code debugging of some methods might no longer work. You might need to support a custom link.xml file, and some reflection code paths might not behave the same. |
  
| **Vertex Compression** | Sets vertex compression per channel. This affects all the meshes in your project.  Typically, Vertex Compression is used to reduce the size of mesh data in memory, reduce file size, and improve GPU performance.   For more information on how to configure vertex compression and limitations of this setting, refer to [Compressing mesh data](compressing-mesh-data-optimization.md). |
| **Optimize Mesh Data** | Enable this option to strip unused vertex attributes from the mesh used in a build. This option reduces the amount of data in the mesh, which can help reduce build size, loading times, and runtime memory usage.    **Warning:** If you have this setting enabled, don’t change material or shader settings at runtime.    For more information, refer to [PlayerSettings.stripUnusedMeshComponents](../ScriptReference/PlayerSettings-stripUnusedMeshComponents.md). |
| **Texture Mipmap Stripping** | Enables mipmap stripping for all platforms. It strips unused mipmap levels from Textures at build time.  Unity determines unused mipmap levels by comparing the mipmap level against the quality settings for the current platform. If a mipmap level is excluded from every quality setting for the current platform, then Unity strips those mipmap levels from the build at build time. If `QualitySettings.globalTextureMipmapLimit` is set to a mipmap level that has been stripped, Unity will set the value to the closest mipmap level that hasn’t been stripped. |

### Stack Trace

Select your preferred stack trace method by enabling the option that corresponds to each Log Type (**Error**, **Assert**, **Warning**, **Log**, and **Exception**) based on the type of logging you require. For more information, refer to [stack trace logging](stack-trace.md).

| **Property** | **Description** |
| --- | --- |
| **None** | No logs are ever recorded. |
| **ScriptOnly** | Logs only when running **scripts**. |
| **Full** | Logs all the time. |

### Legacy

| **Property** | **Description** |
| --- | --- |
| **Clamp BlendShapes (Deprecated)** | Activate the option to clamp the range of blend shape weights in [SkinnedMeshRenderers](class-SkinnedMeshRenderer.md). |

### Capture Logs

| **Property** | **Description** |
| --- | --- |
| **Capture Startup Logs** | Enable this option to generate debug information in the log files during the startup of your application. |