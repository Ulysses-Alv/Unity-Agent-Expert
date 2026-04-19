# Android Player settings

This page details the Player Settings for the Android platform. For a description of the general Player settings, refer to [Player Settings](class-PlayerSettings.md).

![Player Settings for Android](../uploads/Main/class-PlayerSettingsAndroid.png)

Player Settings for Android

Documentation for the properties is grouped according to their respective sections in the Player UI:

* [Icon](#Icon)
* [Resolution and Presentation](#Resolution)
* [Splash Image](#Splash)
* [Other Settings](#Other)
* [Publishing Settings](#Publishing)

## Icon

Use the **Icon** section to specify icons to represent your application on the device.

| **Setting** | **Description** |
| --- | --- |
| **Adaptive** | Specify textures of various sizes to represent your application on devices running Android 8.0 (API level 26) or higher. For more information, refer to the Android documentation on [Adaptive icons](https://developer.android.com/develop/ui/views/launch/icon_design_adaptive). |
| **Round** | Specify textures of various sizes to represent your application on devices running Android 7.1 (API level 25) or higher.   **Note**: **Round** icons are deprecated and will be removed in a future version of Unity. Use **Adaptive** icons instead. To display a round icon for your application on devices running API 26 and later, use an **Adaptive** icon. |
| **Legacy** | Specify textures of various sizes to represent your application on devices running versions earlier than Android 7.1 (API level 25).   **Note**: **Legacy** icons are deprecated and will be removed in a future version of Unity. Use **Adaptive** icons instead. |

## Resolution and Presentation

Use the following settings in **Resolution and Presentation** section to customize aspects of the screen’s appearance.

* [Resolution](#Resolution)
* [Resolution Scaling](#Scaling)
* [Supported Aspect Ratio](#Aspect)
* [Orientation](#Orientation)
* [Allowed Orientations for Auto Rotation](#supportedOris)
* [Miscellaneous](#OtherRes)

### Resolution

Use the **Resolution** section to customize the screen mode and default size.

| **Property** | **Description** |
| --- | --- |
| **Run Without Focus** | Enable this property to keep your application running even when it’s not in focus but is visible, such as on devices supporting split-screen mode. When disabled, the application pauses but remains visible. |
| **Fullscreen Mode** | Specifies how the window appears on the device. This sets the default window mode at startup.     * **Fullscreen Window**: The application window fills the full-screen native resolution of the device. To fill the full-screen window, Unity scales the application contents. **Resolution Scaling Mode** controls how Unity scales the contents. * **Windowed**: The application uses a standard, non-full-screen, movable window. The size of the window depends on the application’s resolution. In this mode, the application is resizable by default. To prevent application resizing, disable [Resizeable Activity](#resizeable). |
| **Resizeable Activity** | Indicates whether the user can resize the application. This setting enables multi-window capabilities in your application on Android phones and tablets. For more information, refer to [Google’s developer documentation](https://developer.android.com/guide/topics/ui/multi-window). |
| **Minimum Window Width** | The minimum width of the application window in pixels. |
| **Minimum Window Height** | The minimum height of the application window in pixels. |
| **Default Window Width** | The default width of the application window in pixels. This option is only available if the **Fullscreen Mode** is set to **Windowed**. |
| **Default Window Height** | The default height of the application window in pixels. This option is only available if the **Fullscreen Mode** is set to **Windowed**. |
| **Hide Navigation Bar** | Indicates whether to hide the navigation bar that appears at the bottom of the screen.  **Note:** Disable this setting to prevent your application from automatically entering full-screen mode. |
| **Render outside safe area** | Enable this option to allow the application to use all available screen space to render, including areas of the display that are cut out (notched). For more information, refer to Android’s [display cutout support](https://developer.android.com/guide/topics/display-cutout) documentation. The behavior of this setting varies depending on the Android version, as outlined in the following table.     * **Android 15 and newer**: This setting has no effect as the application uses the entire screen space by default. For more information, refer to [Android documentation](https://developer.android.com/reference/android/view/WindowManager.LayoutParams?_gl=1*1x8u5ar*_up*MQ..*_ga*MjQ0MTQ3Nzk4LjE3MjE5OTAzMDE.*_ga_6HH9YJMN9M*MTcyMTk5MDMwMC4xLjAuMTcyMTk5MDMwMC4wLjAuMA..#LAYOUT_IN_DISPLAY_CUTOUT_MODE_ALWAYS). * **Android 11 to 14**: When enabled, the application uses the entire screen space available. When disabled, the application uses the space around the cutout only if it’s within the system bar (top or bottom of the device), otherwise the application displays black bars over the cutout area. For more information on when the application uses the space around the cutout, refer to [Android documentation](https://developer.android.com/reference/android/view/WindowManager.LayoutParams?_gl=1*1x8u5ar*_up*MQ..*_ga*MjQ0MTQ3Nzk4LjE3MjE5OTAzMDE.*_ga_6HH9YJMN9M*MTcyMTk5MDMwMC4xLjAuMTcyMTk5MDMwMC4wLjAuMA..#LAYOUT_IN_DISPLAY_CUTOUT_MODE_DEFAULT). * **Android 10 and older**: When enabled, the application uses the screen space around the cutout on the shorter edge (such as the top or bottom) of the device. However, the application blocks the cutout on the longer edge (such as the sides) of the device with black bars. For more information on how the application uses the space around cutouts on shorter edges, refer to [Android documentation](https://developer.android.com/reference/android/view/WindowManager.LayoutParams?_gl=1*1x8u5ar*_up*MQ..*_ga*MjQ0MTQ3Nzk4LjE3MjE5OTAzMDE.*_ga_6HH9YJMN9M*MTcyMTk5MDMwMC4xLjAuMTcyMTk5MDMwMC4wLjAuMA..#LAYOUT_IN_DISPLAY_CUTOUT_MODE_SHORT_EDGES). When disabled, the application uses the space around the cutout only if it’s within the system bar, otherwise the application displays black bars over the cutout area. For more information on when the application uses the space around the cutout, refer to [Android documentation](https://developer.android.com/reference/android/view/WindowManager.LayoutParams?_gl=1*1x8u5ar*_up*MQ..*_ga*MjQ0MTQ3Nzk4LjE3MjE5OTAzMDE.*_ga_6HH9YJMN9M*MTcyMTk5MDMwMC4xLjAuMTcyMTk5MDMwMC4wLjAuMA..#LAYOUT_IN_DISPLAY_CUTOUT_MODE_DEFAULT). |
| **Optimized Frame Pacing** | Enable this option to allow Unity to evenly distribute frames for less variance in frame rate and create a smoother experience. |

### Resolution Scaling

Use the **Resolution Scaling** section to customize settings relating to screen resolution scaling. Using a resolution lower than the device’s native resolution can improve performance and battery life.

| **Property** | **Description** |
| --- | --- |
| **Resolution Scaling Mode** | Specifies whether and how the application scales its resolution. You can set the scaling to be equal to or lower than the native screen resolution. Using a lower resolution can improve performance and battery life.     * **Disabled**: Doesn’t apply resolution scaling and the application renders to the device’s native screen resolution. * **FixedDPI**: Applies resolution scaling using a target DPI. Use this to optimize performance and battery life or target a specific DPI setting. * **Letterboxed**: Adds black bars to the rendered output so the content doesn’t stretch. This process is called [letterboxing](https://en.wikipedia.org/wiki/Letterboxing_(filming)). |
| **Target DPI** | The resolution of the application. If the device’s native screen DPI is higher than this value, Unity downscales the application’s resolution to match this setting. Unity calculates the scale using: **min(Target DPI \* Factor / Screen DPI, 1)** where **Factor** is the **Resolution Scaling Fixed DPI Factor** from [Quality settings](class-QualitySettings.md).  **Note:** This option only appears when you set **Resolution Scaling Mode** to **Fixed DPI**. |
| **Reset resolution on window resize** | Indicates whether to set the screen resolution to the new native window size when the native window size changes. If you set **Resolution Scaling Mode** to **Fixed DPI**, Unity recalculates the resolution based on **Fixed DPI** property. |
| **Blit Type** | Controls whether to use a blit to render the final image to the screen. Using a blit is compatible with most devices but is usually slower than not using a blit.     * **Always**: Unity renders to an offscreen buffer and then uses a blit to copy the contents of the buffer to the device’s framebuffer. This is compatible with most devices but is usually slower than not using blit. * **Never**: Unity renders to the framebuffer provided by the device’s operating system. If this fails, the application prints a one-time warning to the device log. This is usually faster than using blit, but it isn’t compatible with all devices. * **Auto**: Unity renders to the framebuffer provided by the device’s operating system if possible. If this fails, Unity prints a warning to the device console and uses a blit to render the final image to the screen. |

### Supported Aspect Ratio

Use the **Supported **Aspect Ratio**** section to customize settings relating to which device aspect ratios to support.

| **Property** | **Description** |
| --- | --- |
| **Aspect Ratio Mode** | Specifies the largest aspect ratio the application supports. If the device’s aspect ratio is greater than this aspect ratio, Unity uses this aspect ratio for the application and adds black bars so the application doesn’t stretch.     * **Legacy Wide Screen (1.86)**: The application supports aspect ratios up to Android’s legacy wide-screen aspect ratio. * **Native Aspect Ratio**: The application supports aspect ratios up to Android’s native aspect ratio. * **Custom**: The application supports aspect ratios up to the aspect ratio you set in **Maximum**. |
| **Minimum** | The custom minimum aspect ratio. This value must be less than **Maximum**.  **Note**: This setting only appears when you set **Aspect Ratio Mode** to **Custom**. |
| **Maximum** | The custom maximum aspect ratio. This value must be greater than **Minimum**.  **Note**: This setting only appears when you set **Aspect Ratio Mode** to **Custom**. |

### Miscellaneous

| **Property** | **Description** |
| --- | --- |
| **Show Loading Indicator** | Specifies if and how the loading indicator appears.    * **Don’t Show**: The loading indicator doesn’t appear. * **Large**: A large loading indicator appears. * **Inversed Large**: A large loading indicator appears with inversed color. * **Small**: A small loading indicator appears. * **Inversed Small**: A small loading indicator appears with inversed color. |
| **Display Options** | Defines Android display settings that determine how your application renders on the Android devices at runtime.    * **None**: Disables any configured display settings. When you select this option, automatic rendering is disabled and you must manually handle rendering on the secondary screen. For example, by creating a [Unity-as-a-library(UaaL)](UnityasaLibrary-Android.md) based application. * **Detect Presentation Display**: Enables automatic rendering of the camera output on a secondary screen. Automatic rendering is supported only on screens that the Android system identifies as a [presentation](https://developer.android.com/reference/android/app/Presentation) display and when the [Target Display](class-Camera.md) is set to `2` in the Unity Editor. The Unity Player renders the camera output in full screen, even when the application is running in the background. This option is selected by default. |

### Orientation

| **Property** | **Description** |
| --- | --- |
| **Default Orientation** | Specify the orientation of the application window within the device’s screen.  **Note**: Unity shares the value you set for this setting between Android, iOS, and UWP.    * **Portrait**: Portrait orientation where the bottom of the application aligns with the bottom of the screen. * **Portrait Upside Down**: Portrait orientation where the bottom of the application aligns with the top of the screen. * **Landscape Right**: Landscape orientation where the right side of the application aligns with the top of the screen. * **Landscape Left**: Landscape orientation where the left side of the application aligns with the top of the screen. * **Auto Rotation**: The screen can rotate to any of the orientations you specify in the [Allowed Orientations for Auto Rotation](#supportedOris) section. |
| **Auto Rotation Behavior** | Specify how the application window adjusts its orientation based on the device’s rotation sensor and orientation settings, when the **Default Orientation** is set to **Auto Rotation**.  **Note**: This property is visible only when you set the **Default Orientation** to **Auto Rotation**. By default, this property is set to **User**.    * **User**: The application window adjusts its orientation according to the device’s orientation settings. If the user locks the device’s auto rotate orientation setting, the application window doesn’t follow the preferences set in the **Allowed Orientations for Auto Rotation** section. The application window only rotates as per the set preferences when the user turns off the device’s auto rotate orientation setting. * **Sensor**: The application window adjusts its orientation according to the device’s rotation sensor regardless of the device’s orientation settings. However, the application window only rotates as per the preferences set in the **Allowed Orientations for Auto Rotation** section. |

### Allowed Orientations for Auto Rotation

Use the **Allowed Orientations for Auto Rotation** section to specify which orientations the application supports when you set **Default Orientation** to **Auto Rotation**. This is useful, for example, to lock the application to landscape orientation but allow the user to change between landscape left and landscape right.

**Note**: This section is visible only when you set **Default Orientation** in the Orientation section to **Auto Rotation**.

| **Property** | **Description** |
| --- | --- |
| **Portrait** | Allow portrait orientation with the bottom of the application aligned with the bottom of the screen. |
| **Portrait Upside Down** | Allow portrait orientation with the bottom of the application aligned with the top of the screen. |
| **Landscape Right** | Allow landscape orientation with the right side of the application aligned with the top of the screen. |
| **Landscape Left** | Allow landscape orientation with the left side of the application aligned with the top of the screen. |

### Additional Resolution and Presentation settings

| **Property** | **Description** |
| --- | --- |
| **Use 32-bit Display Buffer** | Indicates whether the display buffer holds 32-bit color values instead of 16-bit color values. Enable this setting if you see banding, or need alpha values in [post-processing effects](PostProcessingOverview.md). Some post-processing effects require this because they create [Render Textures](class-RenderTexture.md) in the same format as the display buffer. |
| **Disable Depth and Stencil** | Disable the depth and stencil buffers. |
| **Render Over Native UI** | Force your application to render on top of the native UI. For this setting to take effect, set your Camera’s **Clear Flags** to use a solid color with an alpha value lower than 1.  **Note:** When the application is in split-screen mode, the native UI behind the application window isn’t rendered. |

## Splash Image

Use the **Virtual Reality Splash Image** setting to select a custom splash image for [Virtual Reality](XR.md) displays. For information on common Splash Screen settings, refer to [Splash Screen](class-PlayerSettingsSplashScreen.md).

![Splash screen settings for virtual reality](../uploads/Main/PlayerSetAndroidSplash.png)

Splash screen settings for virtual reality

Below the common Splash Screen settings, you can set up an Android-specific **Static Splash Image**.

![Splash screen settings for Android](../uploads/Main/PlayerSetAndroidSplashStory.png)

Splash screen settings for Android

| **Property** | **Description** |
| --- | --- |
| **Image** | Specifies the texture that the application uses for the Android splash screen. The standard size for the splash screen image is 320x480. |
| **Scaling** | Specifies how to scale the splash image to fit the device’s screen.    * **Center (only scale down)**: Draws the image at its native size unless it’s too large, in which case Unity scales the image down to fit. * **Scale to Fit (letter-boxed)**: Scales the image so that the longer dimension fits the screen size exactly. Unity fills in the empty space around the sides in the shorter dimension in black. * **Scale to Fill (cropped)**: Scales the image so that the shorter dimension fits the screen size exactly. Unity crops the image in the longer dimension. |

## Other Settings

This section allows you to customize a range of options organized into the following groups:

* [Rendering](#Rendering)
* [Vulkan Settings](#Vulkan)
* [Identification](#Identification)
* [Configuration](#Configuration)
* [Android Application Configuration](#AndroidApplicationConfiguration)
* [Shader Settings](#ShaderSettings)
* [Shader Variant Loading Settings](#ShaderVariantLoading)
* [Script Compilation](#ScriptCompilation)
* [Optimization](#Optimization)
* [Stack Trace](#Logging)
* [Legacy](#Legacy)

### Rendering

Use these settings to customize how Unity renders your game for the Android platform.

| **Property** | **Description** |
| --- | --- |
| **Color Space** | Choose which color space to use for rendering. For more information, refer to [Color spaces](color-spaces-landing.md).    * **Gamma**: Gamma color space is typically used for calculating lighting on older hardware restricted to 8 bits per channel for the framebuffer format. Even though monitors today are digital, they might still take a gamma-encoded signal as input. * **Linear**: Linear color space rendering gives more precise results. When you select to work in linear color space, the Editor defaults to using [sRGB](https://en.wikipedia.org/wiki/SRGB) sampling. If your [Textures](Textures.md) are in linear color space, you need to work in linear color space and deactivate sRGB sampling for each Texture. |
| **Auto Graphics API** | Disable this option to manually pick and reorder the graphics APIs. By default, this option is enabled, and Unity tries to use Vulkan. If the device doesn’t support Vulkan, Unity falls back to GLES3.2, GLES3.1 or GLES3.0.    **Important**: Unity adds the GLES3.1/GLES3.1+AEP/GLES3.2 requirement to your Android App Manifest only if you enable **Auto Graphics API** or if **Graphics APIs** includes **OpenGLES3**. |
| **Graphics APIs** | Specifies which graphics APIs the application supports. Unity tries to use the graphics API at the top of the list first and, if the device doesn’t support it, Unity falls back to the next graphics APIs in the list.    * **Vulkan:** Adds Vulkan as an API to target. * **OpenGLES3:** Adds OpenGLES 3 as an API to target.  **Note**: This property is visible only when Auto Graphics API is set to false. |
| **Require ES3.1** | Indicates whether to require that the minimum OpenGL ES 3 minor version is 3.1. This property is visible only if you enable **Auto Graphics API** or if **Graphics APIs** includes **OpenGLES3**. |
| **Require ES3.1+AEP** | Indicates whether to require that the minimum OpenGL ES 3 minor version is 3.1+AEP.  This property is visible only if you enable **Auto Graphics API** or if **Graphics APIs** includes **OpenGLES3**. |
| **Require ES3.2** | Indicates whether to require that the minimum OpenGL ES 3 minor version is 3.2. This property is visible only if you enable **Auto Graphics API** or if **Graphics APIs** includes **OpenGLES3**. |
| **Color Gamut** | You can add or remove [color gamuts](https://en.wikipedia.org/wiki/Gamut) to use for rendering. Click the plus (+) icon to view a list of available gamuts. A color gamut defines a possible range of colors available for a given device (such as a monitor or screen). The sRGB gamut is the default (and required) gamut. |
| **Multithreaded Rendering** | Enable this option to move graphics API calls from Unity’s main thread to a separate worker thread. This can help to improve performance in applications that have high CPU usage on the main thread. |
| **Static Batching** | Enable this option to use [static batching](static-batching). |
| **Dynamic Batching** | Use [Dynamic Batching](DrawCallBatching.md) on your build (enabled by default).  **Note:** Dynamic batching has no effect when a [Scriptable Render Pipeline](scriptable-render-pipeline-introduction.md) is active, so this setting is only visible if the **Scriptable Render Pipeline Asset** [Graphics](class-GraphicsSettings.md#SRLoop) setting is blank. |
| **Sprite Batching Threshold** | Controls the maximum vertex threshold used when batching. |
| **GPU Skinning** | Calculate mesh skinning and blend shapes on the GPU via shaders to free up CPU resources and improve performance.    * **CPU**: Select this option to perform mesh skinning and blend shapes calculation on the CPU. * **GPU**: Select this option to perform mesh skinning and blend shapes calculation on the GPU. * **GPU (Batched)**: Select this option to use batching and reordering to perform mesh skinning and blend shapes calculation on the GPU. Batching reduces the number of dispatch calls to the GPU and can improve performance. |
| **Graphics Jobs** | Offload graphics tasks (render loops) to worker threads running on other CPU cores. This option reduces the time spent in `Camera.Render` on the main thread, which can be a bottleneck. |
| **Graphics Jobs Mode** | Specify the graphics jobs mode to use in your application. For information about jobs in Unity, refer to [Job system overview](job-system-overview.md).    * **Native:** The main thread writes Unity graphics commands for worker threads. The worker threads write the commands to a task executor which the render thread consumes. * **Legacy:** The main thread writes Unity graphics commands for worker threads. The worker threads write commands directly to the render thread. The render thread reads the Unity graphics commands and converts them to native graphics commands. * **Split:** The main thread writes Unity graphics commands for worker threads. The render thread reads the Unity graphics commands and converts them to native graphics commands. The render thread then starts worker threads to write native graphics commands.  **Note**: This option is available only when **Graphics Jobs** is enabled and **Graphics APIs** is set to Vulkan. |
| **Texture compression formats** | Choose between ASTC, ETC2 and ETC (ETC1 for RGB, ETC2 for RGBA), DXTC and DXT + RGTC(BC4, BC5). Refer to [texture compression format overview](texture-choose-format-by-platform.md#android) for more information on how to pick the right format. It’s possible to add multiple texture compression formats to this list for [texture compression targeting](android-distribution-google-play.md#texture-compression-targeting) support, however you can’t have both DXT and DXT + RGTC(BC4, BC5) in this list at the same time. If you enable texture compression targeting, Unity also enables, [Split Application Binary](#splitapplicationbinary). Refer to [Texture compression settings](android-requirements-and-compatibility.md#texture-compression) for more details on how this interacts with the texture compression setting in the [build settings](android-build-settings.md).  **Note**: If you [export](android-export-process.md) or build your application as an APK, Unity only uses the first texture compression format in this list. |
| **Normal Map Encoding** | Choose **XYZ** or **DXT5nm-style** to set the normal map encoding. This setting affects the encoding scheme and compression format used for normal maps. **DXT5nm-style** normal maps are of higher quality, but more expensive to decode in shaders. |
| **Lightmap Encoding** | Defines the encoding scheme and compression format of the lightmaps.  You can choose from **Low Quality**, **Normal Quality**, or **High Quality** |
| **HDR Cubemap Encoding** | Defines the encoding scheme and compression format of the HDR Cubemaps.  You can choose from **Low Quality**, **Normal Quality**, or **High Quality**. For more information, refer to [Lightmaps: Technical information](Lightmaps-TechnicalInformation.md). |
| **Lightmap Streaming** | Enable this option to use [Mipmap Streaming](TextureStreaming.md) for lightmaps. Unity applies this setting to all lightmaps when it generates them.  **Note:** To use this setting, you must enable the [Texture Mipmap Streaming Quality](class-QualitySettings.md#texStream) setting. |
| **Streaming Priority** | Set the priority for all lightmaps in the [Mipmap Streaming system](TextureStreaming.md). Unity applies this setting to all lightmaps when it generates them. Positive numbers give higher priority. Valid values range from `-128` to `127`. |
| **Frame Timing Stats** | Enable this property to gather CPU and GPU frame timing data using [FrameTimingManager](frame-timing-manager.md) API. If you disable this property, [Dynamic Resolution](DynamicResolution-landing.md) camera setting can’t use this data to dynamically adjust the resolution to reduce GPU workload. |
| **OpenGL: Profiler GPU Recorders** | Indicates whether to enable profiler recorders when rendering with OpenGL. This option is always enabled for other rendering APIs but is optional on OpenGL due to potential compatibility issues with Frame Timing Stats and the GPU Profiler. |
| **Allow HDR Display Output** | Enable HDR mode output when the application runs. This only works on displays that support this feature. If the display doesn’t support HDR mode, the game runs in standard mode. |
| **Use HDR Display Output** | Checks if the main display supports HDR, and if it does, swaps to HDR output when the application launches.  **Note**: This option is available only when **Allow HDR Display Output** is active. |
| **Virtual Texturing (Experimental)** | Indicates whether to enable [Virtual Texturing](svt-streaming-virtual-texturing.md).  **Note**: Virtual Texturing isn’t compatible with Android. |
| **360 Stereo Capture** | Indicate whether Unity can capture stereoscopic 360 images and videos.  **Note**: 360 stereoscopic capturing isn’t compatible with Android. |
| **Load/Store Action Debug Mode** | Highlights undefined pixels that might cause rendering problems on mobile platforms. This affects the Unity Editor Game view, and your built application if you select **Development Build** in the **Platform Settings** section of the **Build Profiles** window. For more information, refer to [LoadStoreActionDebugModeSettings](../ScriptReference/Rendering.LoadStoreActionDebugModeSettings.md). |
| **Editor Only** | Indicates whether the Load/Store Action Debug Mode only runs in the Unity Editor.  **Note**:This property is visible only when Load/Store Action Debug Mode is set to true. |

### Vulkan Settings

| **Property** | **Description** |
| --- | --- |
| **SRGB Write Mode** | Enable this option to allow `Graphics.SetSRGBWrite()` renderer to toggle the sRGB write mode during runtime. That is, if you want to temporarily turn off Linear-to-sRGB write color conversion, you can use this property to achieve that. Enabling this has a negative impact on performance on mobile tile-based GPUs; therefore, do NOT enable this for mobile. |
| **Number of swapchain buffers** | Set this option to 2 for double-buffering, or 3 for triple-buffering to use with Vulkan renderer. This setting may help with latency on some platforms, but in most cases you should not change this from the default value of 3. Double-buffering might have a negative impact on performance. Do not use this setting on Android. |
| **Get swapchain image late as possible** | If enabled, Vulkan delays acquiring the backbuffer until after it renders the frame to an offscreen image. Vulkan uses a staging image to achieve this. Enabling this setting causes an extra blit when presenting the backbuffer. This setting, in combination with double-buffering, can improve performance. However, it also can cause performance issues because the additional blit takes up bandwidth. |
| **Recycle command buffers** | Indicates whether to recycle or free [CommandBuffers](../ScriptReference/Rendering.CommandBuffer.md) after Unity executes them. |
| **Apply display rotation during rendering** | Enable this to perform all rendering in the native orientation of the display. This has a performance benefit on many devices. For more information, see documentation on [Framebuffer orientation](vulkan-swapchain-pre-rotation.md). |
| **Deny Filter List** | Specify criteria to identify which Android devices should not use Vulkan API when running a Unity application. For more information, refer to the documentation on [Configure Vulkan API usage](allow-deny-vulkan-usage.md). For information on the parameters used to define the rejection criteria, refer to [Vulkan Device Filtering Asset reference](vulkan-device-filter-list-asset-reference.md).  **Important**: **Deny Filter List** property is obsolete and is no longer available by default in the new projects. This property is available in the Player settings only if you upgrade an existing project that previously used these properties. It is recommended to use the **Vulkan Device Filtering Asset** which replaces the **Deny Filter list**. This asset provides an option to import values defined in the **Deny Filter list** of the existing projects. For more information, refer to [Import legacy Allow and Deny Filter List values](import-legacy-filter-list-values.md). |
| **Allow Filter List** | Specify criteria to identify which Android devices should always use the Vulkan graphics API when running a Unity application. For more information, refer to the documentation on [Configure Vulkan API usage](allow-deny-vulkan-usage.md). For information on the parameters used to define the criteria, refer to [Vulkan Device Filtering Asset reference](vulkan-device-filter-list-asset-reference.md).   **Important**: **Allow Filter List** property is obsolete and is no longer available by default in the new projects. This property is available in the Player settings only if you upgrade an existing project that previously used these properties. It is recommended to use the **Vulkan Device Filtering Asset** which replaces the **Allow Filter list**. This asset provides an option to import values defined in the **Allow Filter list** of the existing projects. For more information, refer to [Import legacy Allow and Deny Filter List values](import-legacy-filter-list-values.md). |
| **Vulkan Device Filtering Asset** | Specifies criteria for Android devices using filter lists to define the following:    * which devices should or should not use the Vulkan graphics API * which devices should use specific graphics jobs modes.  For more information on how to create the **Vulkan Device Filtering asset** and set the criteria using filter lists, refer to [Vulkan API and graphics jobs mode configuration](vulkanapi-graphics-jobs-configuration.md). For information on the parameters used in the filter lists, refer to [Vulkan Device Filtering Asset reference](vulkan-device-filter-list-asset-reference.md).   **Notes**:   * If no asset is selected, the graphics jobs mode set using the [Graphics Jobs Mode](#GraphicsJobsMode) Player setting applies to all the Android devices. * If no asset is selected or if no criteria is set in the **Allow and Deny Filters** of the asset, Unity uses the criteria set in the legacy **Allow and Deny Filter lists** Player settings. * Graphics jobs modes can be selected only for applications that use Vulkan API. |

### Identification

![Identification settings for the Android platform ](../uploads/Main/android-identification.png)

Identification settings for the Android platform

| **Property** | **Description** |
| --- | --- |
| **Override Default Package Name** | Indicates whether to override the default package name for your application. Enabling this property also enables the following property.    * **Package Name:** Set the application ID, which uniquely identifies your app on the device and in Google Play Store. The application ID must follow the convention `com.YourCompanyName.YourProductName` and must contain only alphanumeric and underscore characters. Each segment must start with an alphabetical character. For more information, refer to [Set the application ID](https://developer.android.com/studio/build/configure-app-module#set-application-id). **Important**: Unity automatically removes any invalid characters you type.  **Note:** This setting affects macOS, iOS, tvOS, and Android. |
| **Version** | Enter the build version number of the bundle, which identifies an iteration (released or unreleased) of the bundle. The version is specified in the common format of a string containing numbers separated by dots (For example, 4.3.2). (Shared between iOS and Android.) |
| **Bundle Version Code** | An internal version number. This number is used only to determine whether one version is more recent than another, with higher numbers indicating more recent versions. This isn’t the version number shown to users; that number is set by the `versionName` attribute. The value must be set as an integer, such as “100”. You can define it however you want, as long as each successive version has a higher number.   For example, it could be a build number. Or you could translate a version number in “x.y” format to an integer by encoding the “x” and “y” separately in the lower and upper 16 bits. Or you could simply increase the number by one each time a new version is released.   Keep this number under 100000 if **Split APKs by target architecture** is enabled. Each APK must have a unique version code so Unity adds 100000 to the number for ARMv7, and 200000 for ARM64. |
| **Minimum API Level** | Minimum Android version (API level) required to run the application. |
| **Target API Level** | Target Android version (API level) against which to compile the application. |

### Configuration

| **Property** | **Description** |
| --- | --- |
| **Scripting Backend** | Choose the scripting backend you want to use. The scripting backend determines how Unity compiles and executes C# code in your Project.    * **Mono**: Compiles C# code into .NET Common Intermediate Language (CIL) and executes that CIL using a Common Language Runtime. For more information, refer to [Mono](scripting-backends-mono.md). * **IL2CPP**: Compiles C# code into CIL, converts the CIL to C++ and then compiles that C++ into native machine code, which executes directly at runtime. For more information, refer to [IL2CPP](scripting-backends-il2cpp.md). |
| **API Compatibility Level** | Choose which .NET APIs you can use in your project. This setting can affect compatibility with third-party libraries. However, it has no effect on Editor-specific code (code in an Editor directory, or within an Editor-specific Assembly Definition).  **Tip:** If you’re having problems with a third-party assembly, you can try the suggestion in the [API Compatibility Level](#apiComp) section.    * **.Net Framework**: Compatible with the .NET Framework 4 (which includes everything in the .NET Standard 2.0 profile plus additional APIs). Choose this option when using libraries that access APIs not included in .NET Standard 2.0. Produces larger builds and any additional APIs available aren’t necessarily supported on all platforms. For more information, refer to [Referencing additional class library assemblies](dotnet-profile-assemblies.md). * **.Net Standard 2.1**: Produces smaller builds and has full cross-platform support. |
| **Editor Assemblies Compatibility Level** | Select which .NET APIs to use in your Editor assemblies.    * **.NET Framework**: Compatible with the .NET Framework 4 (which includes everything in the .NET Standard 2.1 profile plus additional APIs). Choose this option when using libraries that access APIs not included in .NET Standard 2.1. Produces larger builds and any additional APIs available aren’t necessarily supported on all platforms. For more information, refer to [Referencing additional class library assemblies](dotnet-profile-assemblies.md). * **.NET Standard**: Compatible with [.NET Standard 2.1](https://docs.microsoft.com/en-us/dotnet/standard/net-standard). Produces smaller builds and has full cross-platform support. |
| **IL2CPP Code Generation** | Defines how Unity manages IL2CPP code generation.  **Note**: To use this, set **Scripting Backend** to **IL2CPP**. |
| **C++ Compiler Configuration** | Choose the C++ compiler configuration used when compiling IL2CPP generated code.    * **Debug**: The debug configuration turns off all optimizations, which makes the code quicker to build but slower to run. * **Release**: The release configuration enables optimizations so that the compiled code runs faster, the binary size is reduced, but compilation takes longer. * **Master**: Master configuration enables all possible optimizations, squeezing every bit of performance possible. For instance, on platforms that use the MSVC++ compiler, this option enables link-time code generation. Compiling code using this configuration can take significantly longer than it does using the Release configuration. The recommended best practice is to build the shipping version of your game using the Master configuration if the increase in build time is acceptable. |
| **IL2CPP Stacktrace Information** | Choose the information to include in a stack trace. For further details on the information types, refer to [Managed stack traces with IL2CPP](il2cpp-managed-stack-traces.md).    * **Method Name**: Include each managed method in the stack trace. * **Method Name, File Name, and Line Number**: Include each managed method with file and line number information in the stack trace.    **Note**: Using this option can increase both the build time and final size of the built program. |
| **Use incremental GC** | Uses the incremental garbage collector, which spreads garbage collection over several frames to reduce garbage collection-related spikes in frame duration. For more information, refer to [Automatic Memory Management](performance-managed-memory.md). |
| **Allow downloads over HTTP** | Indicate whether to allow downloading content over HTTP. The default option is **Not allowed** due to the recommended protocol being HTTPS, which is more secure.    * **Not Allowed**: Never allow downloads over HTTP. * **Allowed in Development Builds**: Only allow downloads over HTTP in development builds. * **Always Allowed**: Allow downloads over HTTP in development and release builds. |
| **Mute Other Audio Sources** | Indicates whether your Unity application stops audio from applications running in the background. Otherwise, Audio from background applications continues to play alongside your Unity application. |

### Android Application Configuration

| **Property** | **Description** |
| --- | --- |
| **Target Architectures** | Specifies which architecture to target.    * **ARMv7**: Enable support for ARMv7 architecture. * **ARM64**: Enable support for ARM64 architecture. **Note**: This property is enabled only when you set **Scripting Backend** to **IL2CPP**. * **x86–64 (Magic Leap 2)**: Enable support for x86–64 architecture. **Note**: This property is enabled only when you set **Scripting Backend** to **IL2CPP**. Support for **x86–64 (Magic Leap)** target architecture is limited to existing projects only. * **Enable Armv9 Security Features for Arm64**: Enable Pointer Authentication (PAuth, PAC) and Branch Target Identification (BTI) for ARM64 builds. **Note**: This property is enabled only when you set **ARM64** as the target architecture. |
| **Split APKs by target architecture** | Enable this option to create a separate APK for each CPU architecture selected in **Target Architectures**. This makes download size smaller for Google Play Store users. This is primarily a Google Play store feature and may not work in other stores. For more details, refer to [Multiple APK Support](https://developer.android.com/google/play/publishing/multiple-apks.html). |
| **Install Location** | Specifies application install location on the device (for detailed information, refer to [Android Developer documentation on install locations](http://developer.android.com/guide/appendix/install-location.html).)    * **Automatic**: Let the operating system decide. User will be able to move the app back and forth. * **Prefer External**: Install the application to external storage (SD card) if possible. The operating system doesn’t guarantee it; if not possible, the app will be installed to internal memory. * **Force Internal**: Force the application to be installed to internal memory. The user will be unable to move the app to external storage. |
| **Storage Location** | Specifies storage location on the device (for detailed information, refer to Android Developer documentation on [storage](https://developer.android.com/reference/android/content/Context#getExternalFilesDir(java.lang.String)) or [cache](https://developer.android.com/reference/android/content/Context#getExternalCacheDir()) locations)    * **Prefer External**: Use external storage (SD card) if present. Permission-wise, this should always be possible since Android [Build.VERSION\_CODES.KITKAT](https://developer.android.com/reference/android/os/Build.VERSION_CODES#KITKAT) version. Otherwise, the internal storage will be used. * **Force Internal**: Force the use of the internal storage. Use this if the external storage can be unmounted while the application is running. |
| **Internet Access** | Choose whether to always add the networking (`INTERNET`) permission to the [Android App Manifest](android-manifest.md), even if you aren’t using any networking APIs. Set to **Require** by default for development builds.    * **Auto**: Only add the internet access permission if you are using a networking API. * **Require**: Always add the internet access permission. |
| **Force UnityWebRequest via HTTP/1.1** | Enable this option to force [UnityWebRequest](web-request.md) to use the HTTP/1.1 protocol to communicate with web servers instead of the HTTP/2 protocol. By default, this option is disabled and UnityWebRequest uses HTTP/2. |
| **Write Permission** | Choose whether to enable write access to the external storage (such as the SD card) and add a corresponding permission to the Android App Manifest. Set to External(SDCard) by default for development builds.    * **Internal**: Only grant write permission to internal storage. * **External(SDCard)**: Enable write permission to external storage. |
| **Filter Touches When Obscured** | Enable this option to discard touches received when another visible window is covering the Unity application. This is to prevent tapjacking. |
| **Sustained Performance Mode** | Enable this option to set a predictable and consistent level of device performance over longer periods of time, without thermal throttling. Overall performance might be lower when this setting is enabled. Based on the [Android Sustained Performance API](https://developer.android.com/about/versions/nougat/android-7.0.html#sustained_performance_api). |
| **Maximum Java Heap Size** | Set the maximum Java heap size to user for building (in megabytes). Defaults to 4096. |
| **Low Accuracy Location** | Enable this option to use low accuracy values with Android location APIs instead. |
| **Predictive Back Gesture Support** | Enable this option to use Android’s [OnBackInvokedCallback](https://developer.android.com/guide/navigation/predictive-back-gesture) for handling back events on Android 13 and above. |
| **App Category** | Specify the type of application you’re building. For the existing Unity projects with [androidIsGame](../ScriptReference/PlayerSettings.Android-androidIsGame.md) set to false, this property is disabled by default. When enabled, select a category from the dropdown that best matches the primary functionality of your application. For more information on the available categories, refer to the [Android documentation](https://developer.android.com/guide/topics/manifest/application-element#appCategory).   **Notes**:  * If the required category isn’t currently listed in the dropdown, select **Other** to manually enter a category name. The entered name must be a valid [appCategory](https://developer.android.com/guide/topics/manifest/application-element#appCategory) defined by Android. Providing an invalid name causes the project build to fail. * For new Unity projects, this property is enabled and set to **Game** by default. Categorizing the application as a game ensures that your application is exempted from Android 16 behavior changes. For more information, refer to the documentation on [Android 16 behavior changes for large screen devices](android-large-screen-and-foldable-support.md#Android16). * If your project uses [Adaptive Performance](adaptive-performance/adaptive-performance.md) or [GameMode API](../ScriptReference/Android.AndroidGame.GameMode.md), ensure to select the **App Category** as **Game.** * If you enable the property and don’t select a category, the project build fails. |
| **Category Name** | Enter a valid application category name that matches one of the [appCategory](https://developer.android.com/guide/topics/manifest/application-element#appCategory) options defined by Android. A custom category name or one not defined by Android isn’t supported.   **Note**: This property is visible only when you set the **App Category** to **Other**. |
| **Android TV Compatibility** | Enable this option to mark the application as Android TV compatible. |
| **Android Gamepad Support Level** | Choose the level of support your application offers for a gamepad.  **Note**: This property is visible only when **Android TV Compatibility** is set to true.    * **Works with D-Pad**: The application is fully operational with a D-pad. No gamepad is needed. * **Supports Gamepad**: The application works with a gamepad, but doesn’t require it. * **Requires Gamepad**: The application requires a gamepad to use. |
| **Warn about App Bundle size** | Enable this option to receive a warning when the size of the [Android App Bundle](https://developer.android.com/platform/technology/app-bundle) exceeds a certain threshold. This option is selected by default and you can only configure it if you select Build Type as **Release** and enable the **Build App Bundle (Google Play)** option in the [build settings](android-build-settings.md).    * **App Bundle size threshold**: Enter a size in MB. When your App Bundle exceeds this size, Unity displays a warning and suggests to perform full size validation for the generated release App Bundle according to [Google Play maximum size limits](https://support.google.com/googleplay/android-developer/answer/9859372#size_limits). |
| **Warn about Symbols size** | Enable this option to receive a warning when the size of the symbols zip package or embedded symbols inside [Android App Bundle](https://developer.android.com/platform/technology/app-bundle) exceeds a certain threshold. This option is selected by default and you can only configure it if you set the **Debug Symbols to Public/Debugging** in the [build settings](android-build-settings.md).    * **Symbols size threshold**: Enter a size in MB. When your symbols package exceeds this size, Unity displays a warning. |
| **Application Entry Point** | The application entry points to generate for your application. For more information, refer to [Android application entry points](android-application-entries.md).    * **Activity**: Generate an [Activity application entry point](android-application-entries-activity.md). * **GameActivity**: Generate a [GameActivity application entry point](android-application-entries-game-activity.md). |
| **Active Input Handling** | Choose how to handle input from users.    * **Input Manager (Old)**: Use the original [Input](class-InputManager.md) settings. * **Input System Package (New)**: Uses the new [input system package](https://docs.unity3d.com/Packages/com.unity.inputsystem@latest). * **Both**: Use both systems. This option isn’t supported on Android. |

#### API Compatibility Level

You can choose your mono API compatibility level for all targets. Sometimes a third-party .NET library uses functionality that’s outside of your .NET compatibility level. To understand what’s going on in such cases, and how to best fix it, try following these suggestions:

1. Install [ILSpy](https://www.microsoft.com/en-us/p/ilspy/9mxfbkfvsq13?cid=msft_web_chart&activetab=pivot:overviewtab) for Windows.
2. Drag the .NET assemblies for the API compatibility level that you are having issues with into ILSpy. You can find these under `Frameworks/Mono/lib/mono/YOURSUBSET/`.
3. Drag in your third-party assembly.
4. Right-click your third-party assembly and select **Analyze**.
5. In the analysis report, inspect the **Depends on** section. The report highlights anything that the third-party assembly depends on, but that’s not available in the .NET compatibility level of your choice in red.

### Shader Settings

| **Property** | **Description** |
| --- | --- |
| **Shader Precision Model** | Select the default precision shaders use. For more information, refer to [Use 16-bit precision in shaders](SL-Use16BitPrecisionInShaders.md).    * **Platform default**: Use lower precision on mobile platforms, and full precision on other platforms. * **Unified**: Use lower precision if the platform supports it. |
| **Strict shader variant matching** | Enable this option to use the error shader for rendering if a shader variant is missing in the Player build and display an error in the console. The error specifies the shader, subshader index, pass, and keywords used for shader variant search |
| **Keep Loaded Shaders Alive** | Keep all loaded shaders alive and prevent unloading. |

### Shader Variant Loading Settings

Use these settings to control how much memory **shaders** use at runtime.

| **Property** | **Description** |
| --- | --- |
| **Default chunk size (MB)** | Sets the maximum size of compressed shader variant data chunks Unity stores in your built application for all platforms. The default is `16`. For more information, refer to [Shader loading](shader-loading.md). |
| **Default chunk count** | Sets the default limit on how many decompressed chunks Unity keeps in memory on all platforms. The default is `0`, which means there’s no limit. |
| **Override** | Enables overriding **Default chunk size** and **Default chunk count** for this build target. |
| **Chunk size (MB)** | Overrides the value of **Default chunk size (MB)** on this build target. |
| **Chunk count** | Overrides the value of **Default chunk count** on this build target. |

### Script Compilation

![Script compilation settings for the Android platform](../uploads/Main/ProjectSettingsScriptCompilation.png)

Script compilation settings for the Android platform

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
| **Strip Engine Code** | Enable this option if you want the Unity Linker tool to remove code for Unity Engine features that your Project doesn’t use. This setting is only available with the [IL2CPP scripting backend](IL2CPP). Most apps do not use every available DLL. This option strips out DLLs that your app doesn’t use to reduce the size of the built Player. If your app is using one or more classes that would normally be stripped out under your current settings, Unity displays a debug message when you try to build the app. |
| **Managed Stripping Level** | Choose how aggressively Unity strips unused managed (C#) code. When Unity builds your app, the Unity Linker process can strip unused code from the managed DLLs your Project uses. Stripping code can make the resulting executable smaller, but can sometimes remove code that’s in use.   For more information about these options and bytecode stripping with IL2CPP, refer to [ManagedStrippingLevel](../ScriptReference/ManagedStrippingLevel.md).    * **Minimal**: Use this to strip class libraries, UnityEngine, Windows Runtime assemblies, and copy all other assemblies. * **Low**: Remove unreachable managed code to reduce build size and Mono/IL2CPP build times. * **Medium**: Run UnityLinker to reduce code size beyond what **Low** can achieve. You might need to support a custom link.xml file, and some reflection code paths might not behave the same. * **High**: UnityLinker will strip as much code as possible. This will further reduce code size beyond what Medium can achieve but managed code debugging of some methods might no longer work. You might need to support a custom link.xml file, and some reflection code paths might not behave the same. |
| **Enable Internal Profiler (Deprecated)** | This feature is deprecated and will be retired in a future version of Unity. Use the [Profiler window](ProfilerWindow.md) instead (menu: **Window > Analytics > Profiler**).  The Profiler collects application performance data and prints a report to the console. The report contains the number of milliseconds each Unity subsystem takes to execute on each frame, averaged across 30 frames. |
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

For more information, refer to [stack trace logging](stack-trace.md).

### Legacy

Enable the **Clamp BlendShapes (Deprecated)** option to clamp the range of blend shape weights in [SkinnedMeshRenderers](class-SkinnedMeshRenderer.md).

## Publishing Settings

Use the following Publishing Settings to configure how Unity builds your Android application.

* [Project Keystore](#projectkeystore)
* [Build](#build)
* [Minify](#minify)
* [Split Application Binary](#splitapplicationbinary)
* [Report Dependencies in App Bundle](#reportDependencies)

![Android publishing panel](../uploads/Main/Android-publishing-page.png)

Android publishing panel

**Note**: For security reasons, Unity doesn’t save your Keystore or Project Key passwords.

Use the Keystore Manager window to create, configure, and load your keys and keystores.
You can load existing keystores and keys from either the Keystore Manager or the main Android Publishing panel. If you choose to load these from inside the Keystore Manager, Unity automatically fills the Project Keystore and Project Key fields.
For further information see documentation on the [Keystore Manager](android-keystore-manager.md).

### Project Keystore

A keystore is a container that holds signing keys for application security. For details, refer to Android developer documentation: [Android keystore system.](https://developer.android.com/training/articles/keystore)

Use the **Project Keystore** settings to choose which keystore to use for the open project. When you load a keystore, Unity loads all the keys in that keystore.
To load and use an existing keystore in your open project:

1. Enable **Custom Keystore**.
2. Open the **Select** dropdown, select **Browse**, and choose a keystore from your file system.
3. Enter your **Keystore password**.

If you don’t have an existing keystore, leave **Custom Keystore** disabled.

Unity uses a debug keystore to sign your application.
A debug keystore is a working keystore. It allows you to sign the application and to test it locally. However, the app store will decline apps signed in this way. This is because the app store is unable to verify the validity and ownership of the application using a debug keystore.

When you load a keystore, Unity loads all the keys in that keystore. Use the **Key alias** dropdown to choose one key from that keystore to use as the active key for the open project.

| **Property** | **Description** |
| --- | --- |
| **Custom Keystore** | Enable **Custom Keystore** to load and use an existing Keystore. |
| **Select** | When **Custom Keystore** is enabled, use this to select the keystore you want to use. The keystores below the partition in the **Select** dropdown are stored in a predefined dedicated location. For more details, refer to [Choose the keystore location](android-keystore-create.md#choose-the-keystore-location). |
| **Keystore path** | You don’t need to enter your keystore path. Unity provides this based on the keystore you choose. |
| **Keystore password** | Enter your keystore password to load your chosen keystore. |
| **Key alias** | Select the key you want to use for the open project. |
| **Key password** | Enter your key password. |

### Build

By default, Unity builds your application with the manifest files, **Gradle** templates and Proguard files provided with the Unity installation. Use the **Build** section of the Android Publishing Settings to change these.

To use a custom manifest file, Gradle template or Proguard file:

1. Enable the appropriate checkbox. Unity creates a default file in your project, and the file location appears below the checkbox.
2. Open the new file and make your changes.
3. Save your changes to this file. Unity automatically uses the changes saved to this file next time it builds the application.

The settings in the Build section only apply to the build process for the current project.

| **Property** | **Description** |
| --- | --- |
| **Custom Main Manifest** | Customizable version of the Android `LibraryManifest.xml` file. This file contains important metadata about your Android application. For more information about the responsibilities of the Main/Unity Library Manifest refer to [Unity Library Manifest](android-library-manifest.md). |
| **Custom Launcher Manifest** | Customizable version of the Android `LauncherManifest.xml` file. This file contains important metadata about your Android application’s launcher. For more information about the responsibilities of the Unity Launcher Manifest refer to [Unity Launcher Manifest](android-launcher-manifest.md). |
| **Custom Main Gradle Template** | Customizable version of the `mainTemplate.gradle` file. This file contains information on how to build your Android application as a library. For more information refer to the documentation on [Gradle project files](android-gradle-overview.md#gradle-project-files). |
| **Custom Launcher Gradle Template** | Customizable version of the `launcherTemplate.gradle` file. This file contains instructions on how to build your Android application. For more information refer to the documentation on [Gradle project files](android-gradle-overview.md#gradle-project-files). |
| **Custom Base Gradle Template** | Customizable version of the `baseProjectTemplate.gradle` file. This file contains configuration that’s shared between all other templates and Gradle projects. For more information, refer to the documentation on [Gradle project files](android-gradle-overview.md#gradle-project-files). |
| **Custom Gradle Properties Template** | Customizable version of the `gradle.properties` file. This file contains configuration settings for the Gradle build environment. This includes:    * The JVM (Java Virtual Machine) memory configuration. * A property to allow Gradle to build using multiple JVMs. * A property for choosing the tool to do the minification. * A property to not compress native libs when building an app bundle. |
| **Custom Gradle Settings Template** | Customizable version of the `settingsTemplate.gradle` file. This file contains declaration of artifact repositories to resolve external dependencies required for your application. |
| **Custom Proguard File** | Customizable version of the `proguard.txt` file. This file contains configuration settings for the minification process. If minification removes some Java code which should be kept, add a rule to keep that code in this file. For more information refer to the documentation on [Minification](#minify). |

### Minify

Minification is a process which shrinks, obfuscates, and optimizes the code in your application. It can reduce the code size and make the code harder to disassemble. However, minification can strip out code that’s actually needed, so ensure to configure this process correctly. Unity uses Proguard for minification. Use the **Minify** settings to define when and how Unity should apply minification to your build. For information on Proguard, refer to [Proguard manual](https://www.guardsquare.com/manual/configuration/usage).

In most cases, it’s good practice to only apply minification to release builds, and not debug builds. This is because minification takes time, and can make the builds slower. It can also make debugging more complicated due to the optimization that the code undergoes.

**Note**: If the Java code is referenced via reflection, the minification process cannot infer that the code is in use and strips that code. You can use [Custom Proguard File](#customproguard) to retain the code that you require.

The settings in the Minify section only apply to the build process for the current project.

| **Property** | **Description** |
| --- | --- |
| **Release** | Instruct Unity to minify your application’s code in release builds. |
| **Debug** | Instruct Unity to minify your application’s code in debug builds. |

### Split Application Binary

Enable the **Split Application Binary** option to optimize the initial install size of your output application. The Google Play Store requires this if you want to publish applications larger than 100 MB.

When enabled, Unity optimizes your output application differently based on its format.

* ****APK****: Unity splits your output application into a main APK and an expansion file (OBB)
* **AAB**: Unity generates a single AAB file that includes a [base module](https://developer.android.com/guide/app-bundle/configure-base) and asset packs

For more information, refer to [Introduction to asset splitting](android-asset-splitting.md).

### Report Dependencies in App Bundle

Enable this option to include information about dependencies in App Bundle if you’re planning to build your application as Android App Bundle (AAB) for Google Play. For more information, refer to [Report App Dependencies](android-distribution-google-play.md#reportappdependencies) section.

**Note**: You can only interact with this setting if you enable **Build App Bundle (Google Play)** option in [Android build settings](android-build-settings.md).

PlayerSettingsAndroid