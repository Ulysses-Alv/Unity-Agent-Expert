# Embedded Linux Player settings

Embedded Linux Player settings lets you customize how Unity builds and displays your final application. To access the Player settings window for Embedded Linux, go to **Edit** > **Project Settings** > **Player** and select the **Embedded Linux** tab. You can use the [PlayerSettings API](../ScriptReference/PlayerSettings.md) to control most of the settings available in this window. For a description of the general Player settings, refer to [Player Settings](class-PlayerSettings.md).

## General settings

The Player settings differ between the platform modules that you have installed. Each platform has its own Player settings which you must set for each version of your application you want to build. To navigate between them, click the tabs with the appropriate platform operating system icon.

![The main Player settings window for Embedded Linux](../uploads/Main/embedded-linux-player-settings.png)

The main Player settings window for Embedded Linux

You can find documentation for the properties in the following sections:

* [Icon](#icon)
* [Resolution and Presentation](#resolution)
* [Splash Image](#splash)
* [Other Settings](#other)

## Icon

Enable the **Override for Embedded Linux** setting to assign a custom icon for your game.

## Resolution and Presentation

Use the Resolution and Presentation section to customize aspects of the screen’s appearance in the Resolution section.

### Resolution

This section allows you to customize the screen mode and default size.

| **Property** | **Description** |
| --- | --- |
| **Fullscreen Mode** | Choose the full-screen mode. This defines the default window mode at startup:  * **Fullscreen Window**: The application window fills the full-screen native resolution of the device. To fill the full-screen window, Unity scales the application contents. To match the aspect ratio of the output device, Unity might add black bars to the rendered output so the content doesn’t stretch. This process is called [letterboxing](https://en.wikipedia.org/wiki/Letterboxing_\(filming\)). In this mode, the navigation bar is always hidden. * **Windowed** : The application uses a standard, non-full-screen, movable window. The size of the window depends on the application’s resolution. In this mode, the window is resizable by default. To disable this, disable Resizable Window. |
| **Default Is Native Resolution** | Enable this option to make the game use the default resolution used on the target machine. This option is not available if **Fullscreen Mode** is set to **Windowed**. |
| **Use 32-bit Display Buffer** | Embedded Linux only supports 32-bit color buffers. |
| **Disable Depth and Stencil** | This feature is not supported on Embedded Linux. |
| **Render Over Native UI** | Enable this option only if you want to force your app to render on top of the native iOS or Android UI. For this setting to take effect, set your **Camera**’s Clear flags to use a solid color with an alpha value lower than 1. |

## Splash Image

Use the **Virtual Reality** Splash Image setting to select a custom splash image for Virtual Reality displays. For information on common Splash Image settings, refer to [Splash Image Player settings](class-PlayerSettingsSplashScreen.md).

## Other Settings

This section allows you to customize a range of options organized into the following groups:

* [Rendering](#Renderin)
* [Configuration](#Configur)

### Rendering

Use these settings to customize how Unity renders your game for the Embedded Linux platform.

![Rendering settings for Embedded Linux](../uploads/Main/embedded-linux-other-rendering.png)

Rendering settings for Embedded Linux

| **Property** | **Description** |
| --- | --- |
| **Color Space** | Choose which color space to use for rendering. For more information, refer to [Color spaces](color-spaces-landing.md).    * **Gamma**: Gamma color space is typically used for calculating lighting on older hardware restricted to 8 bits per channel for the framebuffer format. Even though monitors today are digital, they might still take a gamma-encoded signal as input. * **Linear**: Linear color space rendering gives more precise results. When you select to work in linear color space, the Editor defaults to using [sRGB](https://en.wikipedia.org/wiki/SRGB) sampling. If your [Textures](Textures.md) are in linear color space, you need to work in linear color space and deactivate sRGB sampling for each Texture. |
| **Force SRGB blit** | Enable this to use Force SRGB blit for Linear color space. If your graphics drivers don’t support the linear color space, selecting this option will convert Linear color space to Gamma (SRGB) color space. |
| **Auto Graphics API** | Enable this option to use the best Graphics API on the device the application is running on. Disable it to add and remove supported Graphics APIs. |
| **Multithreaded Rendering** | Enable this option to use multithreaded rendering. |
| **Static Batching** | Enable this option to use Static batching. |
| **Static Batching Threshold** | Controls the maximum vertex threshold used when batching. For more information, refer to [Sprite Batch Vertex Threshold](../ScriptReference/PlayerSettings-spriteBatchVertexThreshold.md). |
| **GPU Compute Skinning** | Enable this option to enable GPU compute skinning, freeing up CPU resources. |
| **Graphics Jobs** | Enable this option to instruct Unity to offload graphics tasks (render loops) to worker threads running on other CPU cores. This is intended to reduce the time spent in `Camera.Render` on the main thread, which is often a bottleneck. |
| **Texture compression format** | Choose the texture compression format to use for textures in your Project. The options are ETC, ETC2, ASTC, and DXT. For more information on each of these compression formats, refer to [TextureImporterOverride](texture-choose-format-by-platform.md). |
| **Normal Map Encoding** | Choose Normal Quality or High Quality to set the lightmap encoding. This setting affects the encoding scheme and compression format of the lightmaps. |
| **Lightmap Streaming** | Enable this option to load only the lightmap mipmaps as needed to render the current game Cameras. This value applies to the lightmap textures as they are generated. Note: To use this setting, you must enable the Texture Streaming Quality setting.  * **Streaming Priority**: Set the lightmap mip map streaming priority to resolve resource conflicts. These values are applied to the light map textures as they’re generated. Positive numbers give higher priority. Valid values range from –128 to 127. |
| **Frame Timing Stats** | Enable this option to gather CPU/GPU frame timing statistics. |
| **OpenGL: Profiler GPU Recorders** | Enable profiler recorders when rendering with OpenGL. |
| **Virtual Texturing (Experimental)** | Indicates whether to enable [Virtual Texturing](svt-streaming-virtual-texturing.md). **Note:** The Unity Editor requires a restart for this setting to take effect. |
| **Load/Store Action Debug Mode** | Highlights undefined pixels that might cause rendering problems on mobile platforms. This affects the Unity Editor Game view, and your built application if you select **Development Build** in the **Platform Settings** of the **Build Profiles** window. Refer to [LoadStoreActionDebugModeSettings](../ScriptReference/Rendering.LoadStoreActionDebugModeSettings.md) for more information. |
| **360 Stereo Capture** | Indicates whether Unity can capture stereoscopic 360 images and videos. If you enable this setting, Unity compiles additional shader variants to support 360 capture (currently only on Windows/macOS). When enabled, enable\_360\_capture keyword is added during the Stereo RenderCubemap call. Note that this keyword isn’t triggered outside the Stereo RenderCubemap function. For more information, refer to [Stereo 360 Image and Video Capture](https://blog.unity.com/technology/stereo-360-image-and-video-capture?_ga=2.228952412.1984266774.1645442174-1855761588.1624871061). |

### Vulkan settings

![Vulkan Player settings for Embedded Linux](../uploads/Main/embedded-linux-vulkan.png)

Vulkan Player settings for Embedded Linux

| **Property** | **Description** |
| --- | --- |
| **SRGB Write Mode** | Enable this option to allow `Graphics.SetSRGBWrite()` renderer to toggle the sRGB write mode during runtime. That is, if you want to temporarily turn off Linear-to-sRGB write color conversion, you can use this property to achieve that. Enabling this has a negative impact on performance on mobile tile-based GPUs; therefore, do NOT enable this for mobile. |
| **Number of swapchain buffers** | Set this option to 2 for double-buffering, or 3 for triple-buffering to use with Vulkan renderer. This setting may help with latency on some platforms, but in most cases you should not change this from the default value of 3. Double-buffering might have a negative impact on performance. Do not use this setting on Android. |
| **Get swapchain image late as possible** | If enabled, Vulkan delays acquiring the backbuffer until after it renders the frame to an offscreen image. Vulkan uses a staging image to achieve this. Enabling this setting causes an extra **blit** when presenting the backbuffer. This setting, in combination with double-buffering, can improve performance. However, it also can cause performance issues because the additional blit takes up bandwidth. |
| **Recycle command buffers** | Indicates whether to recycle or free after Unity executes them. |

### Configuration

Use this section to specify configuration settings for the Embedded Linux platform.

![Configuration settings for Embedded Linux](../uploads/Main/embedded-linux-other-config.png)

Configuration settings for Embedded Linux

| **Property** | **Description** |
| --- | --- |
| **Scripting Backend** | Determines how Unity compiles and executes C# code in your application. The default scripting back end for Embedded Linux is [IL2CPP](scripting-backends-il2cpp.md). |
| **API Compatibility Level** | Choose which .NET APIs you can use in your project. This setting can affect compatibility with third-party libraries. However, it has no effect on Editor-specific code (code in an Editor directory, or within an Editor-specific Assembly Definition).  **Tip:** If you’re having problems with a third-party assembly, you can try the suggestion in the [API Compatibility Level](#apiComp) section.    * **.Net Framework**: Compatible with the .NET Framework 4 (which includes everything in the .NET Standard 2.0 profile plus additional APIs). Choose this option when using libraries that access APIs not included in .NET Standard 2.0. Produces larger builds and any additional APIs available aren’t necessarily supported on all platforms. For more information, refer to [Referencing additional class library assemblies](dotnet-profile-assemblies.md). * **.Net Standard 2.1**: Produces smaller builds and has full cross-platform support. |
| **IL2CPP Code Generation** | Defines how Unity manages IL2CPP code generation. |
| **C++ Compiler Configuration** | Choose the C++ compiler configuration used when compiling IL2CPP generated code.    * **Debug**: The debug configuration turns off all optimizations, which makes the code quicker to build but slower to run. * **Release**: The release configuration enables optimizations so that the compiled code runs faster, the binary size is reduced, but compilation takes longer. * **Master**: Master configuration enables all possible optimizations, squeezing every bit of performance possible. For instance, on platforms that use the MSVC++ compiler, this option enables link-time code generation. Compiling code using this configuration can take significantly longer than it does using the Release configuration. The recommended best practice is to build the shipping version of your game using the Master configuration if the increase in build time is acceptable. |
| **IL2CPP Stacktrace Information** | Choose the information to include in a stack trace. For further details on the information types, refer to [Managed stack traces with IL2CPP](il2cpp-managed-stack-traces.md).    * **Method Name**: Include each managed method in the stack trace. * **Method Name, File Name, and Line Number**: Include each managed method with file and line number information in the stack trace.    **Note**: Using this option can increase both the build time and final size of the built program. |
| **Use incremental GC** | Enable this to use the incremental garbage collector, which spreads garbage collection over several frames to reduce gc-related spikes in frame duration. |
| **Assembly Version Validation** | This is an Editor setting that doesn’t apply in runtime. |
| **Player Data path** | Enter the directory path on the system where you want to save the `.config` and log files. You can also change this from the command line of the player by adding the following arguments: `-platform-hmi-player-data-path <pathname>`. |
| **Enable Game Controllers** | When selected, this setting enables game controllers. You can disable this if you don’t need game controllers, to help reduce the player startup time. |
| **CPU Configuration** | Set the target CPU configuration for the player runtime. The default number of cores is 0, but you can increase it by entering a number. The options for each CPU are: **Disabled**, **High Performance**, and **Low Performance**. |
| **Loading image** | Use this setting to select a custom splash image for the loading screen. |
| **Active Input Handling** | Choose how to handle input from users.  * **Input Manager (Old)**: Use the original [Input](class-InputManager.md) settings. * **Input System Package (New)**: Uses the new [input system package](https://docs.unity3d.com/Packages/com.unity.inputsystem@latest). * **Both**: Use both systems. |

### Shader Settings and Shader Variant Loading Settings

Use these settings to control how much memory **shaders** use at runtime.

![Shader settings for Embedded Linux](../uploads/Main/embedded-linux-shader-settings.png)

Shader settings for Embedded Linux

| **Property** | **Description** |
| --- | --- |
| **Shader precision model** | Select the default precision of samplers used in shaders. For more information, refer to [Writing HLSL shader programs](writing-shader-writing-shader-programs-hlsl.md). |
| **Strict shader variant matching** | Use the error shader if a shader variant is missing and display an error in the console. |
| **Keep Loaded Shaders Alive** | When enabled, you can’t unload a shader. For more information, refer to [Shader loading](shader-loading.md). |
| **Default chunk size (MB)** | Sets the maximum size of compressed shader variant data chunks Unity stores in your built application for all platforms. The default is 16. For more information, refer to [Shader loading](shader-loading.md). |
| **Default chunk count** | Sets the default limit on how many decompressed chunks Unity keeps in memory on all platforms. The default is 0, which means there’s no limit. |
| **Override** | Enable this to override Default chunk size and Default chunk count for this build target. |

### Script Compilation

![Script Compilation settings for Embedded Linux](../uploads/Main/embedded-linux-script-compilation.png)

Script Compilation settings for Embedded Linux

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
| **Managed Stripping Level** | Choose how aggressively Unity strips unused managed (C#) code. When Unity builds your app, the Unity Linker process can strip unused code from the managed DLLs your Project uses. Stripping code can make the resulting executable smaller, but can sometimes remove code that’s in use.   For more information about these options and bytecode stripping with IL2CPP, refer to [ManagedStrippingLevel](../ScriptReference/ManagedStrippingLevel.md).  * **Minimal**: Use this to strip class libraries, UnityEngine, Windows Runtime assemblies, and copy all other assemblies. * **Low**: Remove unreachable managed code to reduce build size and Mono/IL2CPP build times. * **Medium**: Run UnityLinker to reduce code size beyond what **Low** can achieve. You might need to support a custom link.xml file, and some reflection code paths might not behave the same. * **High**: UnityLinker will strip as much code as possible. This will further reduce code size beyond what Medium can achieve but managed code debugging of some methods might no longer work. You might need to support a custom link.xml file, and some reflection code paths might not behave the same. |
  
| **Vertex Compression** | Sets vertex compression per channel. This affects all the meshes in your project.  Typically, Vertex Compression is used to reduce the size of mesh data in memory, reduce file size, and improve GPU performance.   For more information on how to configure vertex compression and limitations of this setting, refer to [Compressing mesh data](compressing-mesh-data-optimization.md). |
| **Optimize Mesh Data** | Enable this option to strip unused vertex attributes from the mesh used in a build. This option reduces the amount of data in the mesh, which can help reduce build size, loading times, and runtime memory usage.    **Warning:** If you have this setting enabled, don’t change material or shader settings at runtime.    For more information, refer to [PlayerSettings.stripUnusedMeshComponents](../ScriptReference/PlayerSettings-stripUnusedMeshComponents.md). |
| **Texture Mipmap Stripping** | Enables mipmap stripping for all platforms. It strips unused mipmap levels from Textures at build time.  Unity determines unused mipmap levels by comparing the mipmap level against the quality settings for the current platform. If a mipmap level is excluded from every quality setting for the current platform, then Unity strips those mipmap levels from the build at build time. If `QualitySettings.globalTextureMipmapLimit` is set to a mipmap level that has been stripped, Unity will set the value to the closest mipmap level that hasn’t been stripped. |

### Logging

Select what type of logging you want to allow for Embedded Linux builds.

* Select your preferred logging method from the available options.
* Check a box that corresponds to each Log Type (Error, Assert, Warning, Log, and Exception) based on the type of logging you require. For example:

  + **ScriptOnly -** Logs only when running **scripts**.
  + **Full** **-** Logs all the time.
  + **None** **-** No logs are ever recorded.

### Legacy

This section covers the legacy player settings.

| Property | Description |
| --- | --- |
| **Clamp BlendShapes (Deprecated)** | Enable the option to clamp the range of blend shape weights in [SkinnedMeshRenderers](class-SkinnedMeshRenderer.md). |
| **Upload Cleared Texture Data** | This is a legacy feature and currently not needed because it uses up the bandwidth. By default, this is enabled for debugging purposes. Enabling this setting clears the initial data and automatically uploads the texture from script to the video memory. |

### Additional resources:

* [Player settings](class-PlayerSettings.md)
* [IL2CPP overview](scripting-backends-il2cpp.md)