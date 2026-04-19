# iOS Player settings

Player settings for the iOS Player let you customize how Unity builds and displays your final application when targeting the iOS platform. You can access them by navigating to **Edit > Project Settings > Player** and then selecting the **iOS** tab. For a description of the general Player settings, refer to [Player Settings](class-PlayerSettings.md).

![Player settings sections for iOS.](../uploads/Main/ios-player-settings.png)

Player settings sections for iOS.

You can find documentation for the properties in the following sections:

* [Icon](#icon)
* [Resolution and Presentation](#resolution)
* [Splash Image](#splash)
* [Debugging and crash reporting](#debugging)
* [Other Settings](#other)

## Icon

![Icon settings for iOS.](../uploads/Main/ios-icon-settings.png)

Icon settings for iOS.

Use the **Icon** settings to customize the branding for your application on the Apple Store.

You can upload and assign the following types of custom icons for each iOS device:

| **Setting** | **Description** |
| --- | --- |
| **Application icons** | The main device icons for your application. |
| **Spotlight icons** | Icons that appears in spotlight search results. |
| **Settings icons** | Icons that appears in the device’s main settings page. |
| **Notification icons** | Icons that display in notifications sent from your application. |
| **Marketing icons** | Icons to use in the Apple App Store. |

Icon files can be any image file [supported by Unity](texture-choose-format-by-platform.md), in the form of a [Texture2D](../ScriptReference/Texture2D.md) asset. If an icon Texture is missing, Unity scales the icon Texture with the nearest size, with a preference towards using larger Textures.

## Resolution and Presentation

Use the following **Resolution and Presentation** settings to customize aspects of the screen’s appearance.

* [Resolution](#Scaling)
* [Miscellaneous](#Miscellaneous)
* [Orientation](#Orientation)
* [Allowed Orientations for Auto Rotation](#supportedOris)
* [Multitasking Support](#Multitasking)
* [Status Bar](#Status)

### Resolution

| **Property** | **Description** |
| --- | --- |
| **Resolution Scaling Mode** | Use the **Resolution Scaling Mode** setting to set the scaling to be equal to or below the native screen resolution.    * **Disabled**:Scaling doesn’t apply and the application renders to its native screen resolution. * **Fixed DPI**: When you select this value, the **Target DPI** property appears. |
| **Target DPI** | Use the **Target DPI** property to set the target DPI of the game screen. Unity downscales the game screen to match this setting if the device’s native screen DPI is higher than this value. This helps you optimize performance and battery life. You can also use it to target a specific DPI setting.  Unity calculates the scale as min(Target DPI \* Factor / Screen DPI, 1). The factor uses the **Resolution Scaling Fixed DPI Factor** value in [Quality](class-QualitySettings.md) settings.    **Note**: This setting is visible only when **Resolution Scaling Mode** is **Fixed DPI**. |
| **Reset resolution on window resize** | Set the screen resolution to the new window size when the window size changes. If you set **Resolution Scaling Mode** to **Fixed DPI**, the resolution recalculates based on the **Fixed DPI** property. |

### Miscellaneous

| **Property** | **Description** |
| --- | --- |
| **Show Loading Indicator** | Specify how the loading indicator appears.    * **Don’t Show**: The loading indicator doesn’t appear. * **Medium**: The default loading indicator appears. * **Large**: A large loading indicator appears. |

### Orientation

| **Property** | **Description** |
| --- | --- |
| **Default Orientation** | Specify the orientation of the application window within the device’s screen.  **Note**: Unity shares the value you set for this setting between Android, iOS, and UWP.    * **Portrait**: Portrait orientation where the bottom of the application aligns with the bottom of the screen. * **Portrait Upside Down**: Portrait orientation where the bottom of the application aligns with the top of the screen. * **Landscape Right**: Landscape orientation where the right side of the application aligns with the top of the screen. * **Landscape Left**: Landscape orientation where the left side of the application aligns with the top of the screen. * **Auto Rotation**: The screen can rotate to any of the orientations you specify in the [Allowed Orientations for Auto Rotation](#supportedOris) section. |
| **Use Animated Autorotation (Deprecated)** | Enable an animated transition when the user rotates the device. Otherwise, screen rotation changes instantly.    **Notes**:  * This property is visible only when you set the **Default Orientation** to **Auto Rotation**. * On iOS 16 and later versions, this property has no effect as autorotation is always animated. * This property is deprecated and will be removed in a future release. |

### Allowed Orientations for Auto Rotation

Use the **Allowed Orientations for Auto Rotation** section to specify which orientations the application supports when you set **Default Orientation** to **Auto Rotation**. This is useful, for example, to lock the application to landscape orientation but allow the user to change between landscape left and landscape right.

**Note**: This section is visible only when you set **Default Orientation** in the Orientation section to **Auto Rotation**.

| **Property** | **Description** |
| --- | --- |
| **Portrait** | Allow portrait orientation with the bottom of the application aligned with the bottom of the screen. |
| **Portrait Upside Down** | Allow portrait orientation with the bottom of the application aligned with the top of the screen. |
| **Landscape Right** | Allow landscape orientation with the right side of the application aligned with the top of the screen. |
| **Landscape Left** | Allow landscape orientation with the left side of the application aligned with the top of the screen. |

### Multitasking Support

Enable the **Requires Fullscreen** option if your application needs to run in full screen mode.

### Status Bar

Use the **Status Bar** section to specify parameters related to the iOS status bar.

| **Property** | **Description** |
| --- | --- |
| **Status Bar Hidden** | Hide the status bar when your application launches. |
| **Status Bar Style** | Define the style of the status bar when the application launches. |

### Additional Resolution and Presentation settings

| **Property** | **Description** |
| --- | --- |
| **Disable Depth and Stencil** | Disable the depth and stencil buffers. |
| **Render Over Native UI** | Force your application to render on top of the native UI. For this setting to take effect, set your Camera’s **Clear Flags** to use a solid color with an alpha value lower than 1. |

## Splash image

Use the **Virtual Reality Splash Image** setting to select a custom splash image for [Virtual Reality](XR.md) displays. For information on common Splash Screen settings, refer to [Splash Screen](class-PlayerSettingsSplashScreen.md).

### Launch Screens

![Launch Screen settings for iOS.](../uploads/Main/ios-launch-screen.png)

Launch Screen settings for iOS.

A **launch screen** is a Storyboard file from which iOS creates a splash screen on the device. It has the following limitations:

* You can’t display different content based on device orientation for iPads.
* All iPhones support landscape launch screens, but certain iOS versions display the **Landscape Left** launch screen instead of the **Landscape Right** one due to a bug.
* Before iOS 13, it wasn’t possible to override the root view controller for a storyboard, so the launch screen supports all orientations and ignores some UI settings like status bar behavior. Starting with iOS 13, launch screen is fully compliant with your app’s settings for orientation and status bar appearance.

These limitations also apply to custom storyboards.

Select the **Launch screen type** for iPhone and iPad in the **iPhone Launch Screen** and **iPad Launch Screen** sections:

| **Property** | **Description** |
| --- | --- |
| **Launch screen type** | Select the **Launch screen type** to use for either iPhone and iPad.    * **Default**: When you select this option, Unity adds a solid blue-black image that’s the same color as the background for Unity’s default splash screen. The image is used for both portrait and landscape orientations, and displays using aspect-fill mode. * **None**: Apple no longer supports Launch Images. This option exists for backwards compatibility reasons. If selected, your project will use Unity’s default splash screen instead. * **Image and background (relative size)**: Displays an image in the center of the screen and fills the remaining area with a solid color. * **Image and background (constant size)**: Displays an image in the center of the screen and fills the area with a solid color. This has the same options as **Image and background (relative size)**, except that you specify image size in points rather than as a percentage. * **Custom Storyboard**: Use a [**Custom Storyboard**](#storyboards). |
| **Portrait Image** | Select an image for portrait orientation.  **Note** : This property is visible only when you set the **Launch screen type** to **Image and background (relative size)** or **Image and background (constant size)**. |
| **Landscape Image** | Select an image for landscape orientation.  **Note** : This property is visible only when you set the **Launch screen type** to **Image and background (relative size)** or **Image and background (constant size)**. |
| **Background Color** | Select a background color to fill any screen space that the image doesn’t cover.  **Note** : This property is visible only when you set the **Launch screen type** to **Image and background (relative size)** or **Image and background (constant size)**. |
| **Fill percentage** | Specify image size as a percentage of the screen size. Unity calculates image size based on the length or width of the screen, whichever one is smaller (vertical for landscape displays, horizontal for portrait ones).  **Note** : This property is visible only when you set the **Launch screen type** to **Image and background (relative size)**. |
| **Size in Points** | Specify an image size in points.  **Tip:** Points and pixels aren’t the same. A pixel (px) is a single dot in your image. A point (pt) is a unit of length that measures 1/72 of an inch. Pixel sizes are relative to the screen size and resolution, while point sizes are the size on any screen. For more information, refer to the [W3C guidance](https://www.w3.org/Style/Examples/007/units.en.html) on size units.  **Note** : This property is visible only when you set the **Launch screen type** to **Image and background (constant size)**. |

### Custom storyboards

When you select the **Custom Storyboard** option in the **Launch screen type** dropdown for either iPhone or iPad, the **Custom Storyboard** button appears.

Click the **Custom Storyboard** button to select a storyboard to display when your game starts up on the device. For your storyboard to appear here, you must first create the storyboard in Xcode and copy it to your Project. For more information, refer to [storyboards, scenes, and connections](https://help.apple.com/xcode/mac/10.2/#/dev62c993289).

If you don’t use custom storyboards, you can use Unity’s built-in launch screens to implement splash images.

## Debugging and crash reporting

Use the debugging and crash reporting settings to gather data about your app’s performance and troubleshoot crashes.

### Debugging

| **Property** | **Description** |
| --- | --- |
| **Enable Internal Profiler (Deprecated)** | This feature is deprecated and will be removed in a future version of Unity. Use the [Profiler window](ProfilerWindow.md) instead.  The Profiler collects application performance data and prints a report to the console. The report has the number of milliseconds each Unity subsystem takes to execute on each frame, averaged across 30 frames. |

### Crash Reporting

| **Property** | **Description** |
| --- | --- |
| **On .Net UnhandledException (Deprecated)** | Select the action Unity takes when a .NET unhandled exception occurs.   **Note**: This property is deprecated and will be removed in a future release.    * **Crash**: Generates a crash report if the application crashes which users can submit to iTunes and developers can use to troubleshoot the crash. * **Silent Exit**: The application exits with no errors and doesn’t generate a crash report. |
| **Log Obj-C Uncaught Exceptions** | Print Objective-C Uncaught Exception information to the console. |
| **Enable CrashReport API** | Enable a custom crash reporter to capture crashes. You can use scripts to access crash logs via the [CrashReport](../ScriptReference/CrashReport.md) API. |

## Other settings

You can configure a range of Player settings in the **Other Settings** section. These options are organized into the following groups:

* [Rendering](#Rendering)
* [Identification](#Identification)
* [Configuration](#Configuration)
* [Shader Settings](#ShaderSettings)
* [Shader Variant Loading](#ShaderVariantLoading)
* [Script Compilation](#ScriptCompilation)
* [Optimization](#Optimization)
* [Stack Trace](#Logging)
* [Legacy](#Legacy)

### Rendering

Use these settings to customize how Unity renders your application for the iOS platform.

| **Property** | **Description** |
| --- | --- |
| **Color Space** | Choose which color space to use for rendering. For more information, refer to [Color spaces](color-spaces-landing.md).    * **Gamma**: Gamma color space is typically used for calculating lighting on older hardware restricted to 8 bits per channel for the framebuffer format. Even though monitors today are digital, they might still take a gamma-encoded signal as input. * **Linear**: Linear color space rendering gives more precise results. When you select to work in linear color space, the Editor defaults to using [sRGB](https://en.wikipedia.org/wiki/SRGB) sampling. If your [Textures](Textures.md) are in linear color space, you need to work in linear color space and deactivate sRGB sampling for each Texture. |
| **Color Gamut** | You can add or remove [color gamuts](https://en.wikipedia.org/wiki/Gamut) to use for rendering. Click the plus (+) icon to view a list of available gamuts. A color gamut defines a possible range of colors available for a given device (such as a monitor or screen). The sRGB gamut is the default (and required) gamut. |
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
| **Texture compression format** | Set ASTC as the default texture compression format for iOS. For more information, refer to [texture compression format overview](texture-choose-format-by-platform.md#ios-and-tvos). |
| **Normal Map Encoding** | Choose **XYZ** or **DXT5nm-style** to set the normal map encoding. This setting affects the encoding scheme and compression format used for normal maps. **DXT5nm-style** normal maps are of higher quality, but more expensive to decode in shaders. |
| **Lightmap Encoding** | Defines the encoding scheme and compression format of the lightmaps.  You can choose from **Low Quality**, **Normal Quality**, or **High Quality** |
| **HDR Cubemap Encoding** | Defines the encoding scheme and compression format of the HDR Cubemaps.  You can choose from **Low Quality**, **Normal Quality**, or **High Quality**. For more information, refer to [Lightmaps: Technical information](Lightmaps-TechnicalInformation.md). |
| **Lightmap Streaming** | Enable this option to use [Mipmap Streaming](TextureStreaming.md) for lightmaps. Unity applies this setting to all lightmaps when it generates them.  **Note:** To use this setting, you must enable the [Texture Mipmap Streaming Quality](class-QualitySettings.md#texStream) setting. |
| **Streaming Priority** | Set the priority for all lightmaps in the [Mipmap Streaming system](TextureStreaming.md). Unity applies this setting to all lightmaps when it generates them. Positive numbers give higher priority. Valid values range from `-128` to `127`. |
| **Frame Timing Stats** | Enable this property to gather CPU and GPU frame timing data using [FrameTimingManager](frame-timing-manager.md) API. If you disable this property, [Dynamic Resolution](DynamicResolution-landing.md) camera setting can’t use this data to dynamically adjust the resolution to reduce GPU workload. |
| **Allow HDR Display Output** | Enable HDR mode output when the application runs. This only works on displays that support this feature. If the display doesn’t support HDR mode, the game runs in standard mode. |
| **Use HDR Display Output** | Checks if the main display supports HDR, and if it does, swaps to HDR output when the application launches.  **Note**: This option is available only when **Allow HDR Display Output** is active. |
| **Swap Chain Bit Depth** | Select the number of bits in each color channel for swap chain buffers. For more information on bit depth, refer to [D3DHDRDisplayBitDepth](../ScriptReference/D3DHDRDisplayBitDepth.md).  **Note**: This setting is visible only when **Use HDR Display Output** is enabled.    * **Bit Depth 10**: Unity uses the R10G10B10A2 buffer format and Rec2020 primaries with ST2084 PQ encoding. * **Bit Depth 16**: Unity uses the R16G16B16A16 buffer format and Rec709 primaries with linear color (no encoding). |
| **Virtual Texturing (Experimental)** | Reduce GPU memory usage and texture loading times if your Scene has many high resolution textures. For more information, refer to [Virtual Texturing](svt-streaming-virtual-texturing.md). **Note**: The Unity Editor requires a restart for this property to take effect. |
| **360 Stereo Capture** | Indicate whether Unity can capture stereoscopic 360 images and videos. When enabled, Unity compiles additional shader variants to support 360 capture (currently only on Windows/OSX). The `enable_360_capture` keyword is added during the [`RenderToCubemap`](../ScriptReference/Camera.RenderToCubemap.md) call, but isn’t triggered outside of this function. |
| **Load/Store Action Debug Mode** | Highlights undefined pixels that might cause rendering problems on mobile platforms. This affects the Unity Editor Game view, and your built application if you select **Development Build** in the **Platform Settings** section of the **Build Profiles** window. For more information, refer to [LoadStoreActionDebugModeSettings](../ScriptReference/Rendering.LoadStoreActionDebugModeSettings.md). |

### Identification

Enter identifying information for your app.

| **Property** | **Description** |
| --- | --- |
| **Override Default Bundle Identifier** | Indicates whether you can manually set the bundle identifier.  **Note:** This setting affects macOS, iOS, tvOS, and Android. |
| **Bundle Identifier** | Enter the bundle identifier of the application you are building. The basic structure of the identifier is `com.CompanyName.ProductName`. This structure might vary based on which country you live in, so always default to the string Apple provides to you for your Developer Account. Your **ProductName** is available from your provisioning certificates.  The Bundle Identifier must follow the convention `com.YourCompanyName.YourProductName` and must contain only alphanumeric and hyphen characters. For more information, refer to [CFBundleIdentifier](https://developer.apple.com/library/content/documentation/General/Reference/InfoPlistKeyReference/Articles/CoreFoundationKeys.html#/apple_ref/doc/uid/20001431-102070). **Important**: Unity automatically replaces any invalid characters you type with a hyphen.  To set this property, enable **Override Default Bundle Identifier**. |
| **Version** | Enter the release-version-number string for the bundle (for example, 4.3.6).  This value appears as `CFBundleShortVersionString` in the associated `info.plist` file. For more information, refer to [CFBundleShortVersionString](https://developer.apple.com/library/archive/documentation/General/Reference/InfoPlistKeyReference/Articles/CoreFoundationKeys.html#//apple_ref/doc/uid/20001431-102364). |
| **Build** | Enter the build number for this version of your app.  This value appears as `CFBundleVersion` in the associated `info.plist` file. For more information, refer to [CFBundleVersion](https://developer.apple.com/library/archive/documentation/General/Reference/InfoPlistKeyReference/Articles/CoreFoundationKeys.html#//apple_ref/doc/uid/20001431-102364). |
| **Signing Team ID** | Enter your Apple Developer Team ID. You can find this on the Apple Developer website under [Xcode Help](https://help.apple.com/xcode/mac/current/#/devaf282080a?sub=dev8877b4398). This sets the Team ID for the generated Xcode project, allowing developers to use **Build and Run** functionality. You must enter an Apple Developer Team ID for automatic signing of your app.  For more information, refer to [Creating Your Team Provisioning Profile](https://help.apple.com/xcode/mac/current/#/dev60b6fbbc7). |
| **Automatically Sign** | Enable this option to instruct Xcode to automatically sign your build. |

### Configuration

| **Property** | **Description** |
| --- | --- |
| **Scripting Backend** | Choose the scripting backend you want to use. The scripting backend determines how Unity compiles and executes C# code in your Project.    * **IL2CPP**: Compiles C# code into CIL, converts the CIL to C++ and then compiles that C++ into native machine code, which executes directly at runtime. For more information, refer to [IL2CPP](scripting-backends-il2cpp.md). |
| **API Compatibility Level** | Choose which .NET APIs you can use in your project. This setting can affect compatibility with third-party libraries. However, it has no effect on Editor-specific code (code in an Editor directory, or within an Editor-specific Assembly Definition).  **Tip:** If you’re having problems with a third-party assembly, you can try the suggestion in the [API Compatibility Level](#apiComp) section.    * **.Net Framework**: Compatible with the .NET Framework 4 (which includes everything in the .NET Standard 2.0 profile plus additional APIs). Choose this option when using libraries that access APIs not included in .NET Standard 2.0. Produces larger builds and any additional APIs available aren’t necessarily supported on all platforms. For more information, refer to [Referencing additional class library assemblies](dotnet-profile-assemblies.md). * **.Net Standard 2.1**: Produces smaller builds and has full cross-platform support. |
| **Editor Assemblies Compatibility Level** | Select which .NET APIs to use in your Editor assemblies.    * **.NET Framework**: Compatible with the .NET Framework 4 (which includes everything in the .NET Standard 2.1 profile plus additional APIs). Choose this option when using libraries that access APIs not included in .NET Standard 2.1. Produces larger builds and any additional APIs available aren’t necessarily supported on all platforms. For more information, refer to [Referencing additional class library assemblies](dotnet-profile-assemblies.md). * **.NET Standard**: Compatible with [.NET Standard 2.1](https://docs.microsoft.com/en-us/dotnet/standard/net-standard). Produces smaller builds and has full cross-platform support. |
| **IL2CPP Code Generation** | Defines how Unity manages IL2CPP code generation.  **Note**: To use this, set **Scripting Backend** to **IL2CPP**. |
| **C++ Compiler Configuration** | Choose the C++ compiler configuration used when compiling IL2CPP generated code.    * **Debug**: The debug configuration turns off all optimizations, which makes the code quicker to build but slower to run. * **Release**: The release configuration enables optimizations so that the compiled code runs faster, the binary size is reduced, but compilation takes longer. * **Master**: Master configuration enables all possible optimizations, squeezing every bit of performance possible. For instance, on platforms that use the MSVC++ compiler, this option enables link-time code generation. Compiling code using this configuration can take significantly longer than it does using the Release configuration. The recommended best practice is to build the shipping version of your game using the Master configuration if the increase in build time is acceptable. |
| **IL2CPP Stacktrace Information** | Choose the information to include in a stack trace. For further details on the information types, refer to [Managed stack traces with IL2CPP](il2cpp-managed-stack-traces.md).    * **Method Name**: Include each managed method in the stack trace. * **Method Name, File Name, and Line Number**: Include each managed method with file and line number information in the stack trace.    **Note**: Using this option can increase both the build time and final size of the built program. |
| **Use incremental GC** | Uses the incremental garbage collector, which spreads garbage collection over several frames to reduce garbage collection-related spikes in frame duration. For more information, refer to [Automatic Memory Management](performance-managed-memory.md). |
| **Allow downloads over HTTP** | Indicate whether to allow downloading content over HTTP. The default option is **Not allowed** due to the recommended protocol being HTTPS, which is more secure.    * **Not Allowed**: Never allow downloads over HTTP. * **Allowed in Development Builds**: Only allow downloads over HTTP in development builds. * **Always Allowed**: Allow downloads over HTTP in development and release builds. |
| **Camera Usage Description** | Enter the reason for accessing the camera on the device. |
| **Microphone Usage Description** | Enter the reason for accessing the microphone on the device. |
| **Location Usage Description** | Enter the reason for accessing the location of the device. |
| **Use On-Demand Resources** | Use on-demand resources and enable **Variant map for app slicing**. |
| **Accelerometer Frequency** | Choose how often to sample the accelerometer. Use a lower frequency for better performance. If your app doesn’t use the accelerometer, select **Disabled**. |
| **Mute Other Audio Sources** | Indicates whether your Unity application stops audio from applications running in the background. Otherwise, Audio from background applications continues to play alongside your Unity application. |
| **Prepare iOS for Recording** | Initializes the microphone recording APIs. This lowers recording latency, but it also re-routes iPhone audio output via earphones. |
| **Force iOS Speakers when Recording** | Send the phone’s audio output through the internal speakers, even when headphones are plugged in and recording. |
| **Requires Persistent WiFi** | Enable this option to require a Wi-Fi connection. This maintains the active Wi-Fi connection while the app is running. |
| **Supported URL schemes** | A list of [supported URL schemes](https://developer.apple.com/documentation/xcode/defining-a-custom-url-scheme-for-your-app/).  To add new schemes, increase the value of the **Size** property, then set a reference to the URL to open in the new **Element** box that appears. |
| **Target Device** | Select which devices the app targets. |
| **Target SDK** | Select which SDK the app targets.  **Note:** Be sure to select the correct SDK. For example, if you select the Device SDK in Unity and target the Simulator in Xcode, your build will fail. |
| **Target minimum iOS Version** | Select the minimum version of iOS that the app works on. |
| **Enable ProMotion Support** | Use high frequency refresh rates (120 Hz) on ProMotion displays. This might cause your app to use more battery power. |
| **Requires ARKit support** | Restrict the app to iPhone 6s/iOS 11 and later devices when you publish it to the App Store. |
| **Automatically add capabilities** | Generate an `entitlements.plist` file and add capabilities for Game Center if detected in your project. Disable this setting if you intend to sign your app with an enterprise certificate or use a wildcard bundle identifier. |
| **Defer system gestures on edges** | Select one or more edges that users must swipe twice to enact system gestures. |
| **Hide home button on iPhone X** | Hide the home button on iPhone X devices when the app is running. |
| **Render Extra Frame on Pause** | Issue an additional frame after the frame when the app pauses. This lets your app display graphics that display the paused state when the app goes into the background. |
| **Enable Custom Background Behaviors** | Choose what the app can do when the user sends the app to the background. For an example, refer to the [BackgroundFetch](https://github.com/Unity-Technologies/iOSNativeCodeSamples/tree/2019-dev/NativeIntegration/BackgroundTasks/BackgroundFetch) GitHub project. |
| **Variant map for app slicing** | Expand the Variant map for app slicing section to find the list of variant names configured in scripting. For more information on variants, refer to [App slicing](ios-app-slicing.md).  **Note**: Enable **Use on-demand resources** to view this property. |
| **Architecture** | Choose which architecture to target. Only **ARM64** is available for iOS. |
| **Simulator Architecture** | Choose the architecture to build for the Simulator.   This property is visible only when you set the **Target SDK** to **Simulator SDK**.    * **ARM64**: Supports 64-bit ARM architecture. * **X86\_64**: Supports 64-bit Intel/AMD architecture. * **Universal**: Supports all architectures including 64-bit ARM and Intel/AMD. |
| **Active Input Handling** | Choose how to handle input from users.    * **Input Manager (Old)**: Use the original [Input](class-InputManager.md) settings. * **Input System Package (New)**: Uses the new [input system package](https://docs.unity3d.com/Packages/com.unity.inputsystem@latest). * **Both**: Use both systems. |

### API Compatibility Level

You can choose your mono API compatibility level for all targets. Sometimes a third-party .NET library uses functionality that is outside of your .NET compatibility level. If you’re on Windows, you can use the third-party software ILSpy to understand what’s happening and how to fix it. Follow these steps:

1. Install [ILSpy](https://www.microsoft.com/en-us/p/ilspy/9mxfbkfvsq13?cid=msft_web_chart&activetab=pivot:overviewtab) for Windows.
2. Drag the .NET assemblies for the API compatibility level that you are having issues with into ILSpy. You can find these under `Frameworks/Mono/lib/mono/YOURSUBSET/`.
3. Drag in your third-party assembly.
4. Right-click your third-party assembly and select **Analyze**.
5. In the analysis report, inspect the **Depends on** section. Anything that the third-party assembly depends on, but that’s not available in the .NET compatibility level of your choice, is highlighted in red.

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

### Script Compilation

![Script Compilation settings for iOS.](../uploads/Main/ios-script-compilation.png)

Script Compilation settings for iOS.

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
| **AOT Compilation Options** | Additional options for ahead-of-time (AOT) compilation. This helps [optimize the size of the built iOS player](iphone-playerSizeOptimization.md). |
| **Strip Engine Code** | Enable this option if you want the Unity Linker tool to remove code for Unity Engine features that your Project doesn’t use. This setting is only available with the [IL2CPP scripting backend](IL2CPP). Most apps do not use every available DLL. This option strips out DLLs that your app doesn’t use to reduce the size of the built Player. If your app is using one or more classes that would normally be stripped out under your current settings, Unity displays a debug message when you try to build the app. |
| **Managed Stripping Level** | Choose how aggressively Unity strips unused managed (C#) code. When Unity builds your app, the Unity Linker process can strip unused code from the managed DLLs your Project uses. Stripping code can make the resulting executable smaller, but can sometimes remove code that’s in use.   For more information about these options and bytecode stripping with IL2CPP, refer to [ManagedStrippingLevel](../ScriptReference/ManagedStrippingLevel.md).    * **Minimal**: Use this to strip class libraries, UnityEngine, Windows Runtime assemblies, and copy all other assemblies. * **Low**: Remove unreachable managed code to reduce build size and Mono/IL2CPP build times. * **Medium**: Run UnityLinker to reduce code size beyond what **Low** can achieve. You might need to support a custom link.xml file, and some reflection code paths might not behave the same. * **High**: UnityLinker will strip as much code as possible. This will further reduce code size beyond what Medium can achieve but managed code debugging of some methods might no longer work. You might need to support a custom link.xml file, and some reflection code paths might not behave the same. |
  
| **Script Call Optimization** | Choose how to disable exception handling for a speed boost at runtime. For more information, refer to [iOS Optimization](iphone-iOS-Optimization.md). |
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

PlayerSettingsiOS