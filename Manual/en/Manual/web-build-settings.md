# Web build settings reference

Use the Web build settings to configure how Unity builds your application for the Web platform.

You can update your Web platform build settings from the [Build Profiles](BuildSettings.md) window.

**Note**: The [Facebook Instant Games platform](instant-games.md) shares these settings but uses different templates customized for device orientations.

## Platform settings

| **Property** | **Description** |
| --- | --- |
| **Client Browser Type** | Select the browser client that you want your application to launch at runtime. For example, if you choose System Default, then your application launches the default browser, and removes the **Path to Client Browser** setting.  * **System Default**: This is the default setting. Select this option to launch the application in the default browser. * **Microsoft Edge**: Select this option to launch the application in Edge browser. * **Apple Safari**: Select this option to launch the application in Safari. * **Mozilla Firefox**: Select this option to launch the application in Firefox. * **Google Chrome**: Select this option to launch the application in Chrome. * **Chromium**: Select this option to launch the application in a Chromium browser. |
| **Path to Client Browser** | Specify the path to the browser client that you want your application to launch at runtime.  **Note:** This option is only visible if the Client Browser Type isn’t set to the default browser |
| **Texture Compression** | The texture compression format to use for the build. For more information, refer to [Web texture compression](webgl-texture-compression.md).  * **Use Player Settings**: This is the default selection. It uses the texture compression format you set in the [Player settings](class-PlayerSettingsWebGL.md) window. * **ETC2**: Uses ETC2 format, which is widely supported on mobile devices. * **ASTC**: Uses ASTC format, which is widely supported on mobile devices. * **DXT**: Uses DXT format, which is widely supported on desktop devices. |
| **Development Build** | Include scripting debug symbols and the [Profiler](Profiler.md) in your build. Use this setting when you want to test your application. When you select this option, Unity sets the `DEVELOPMENT_BUILD` scripting define symbol. Your build then includes preprocessor directives that set `DEVELOPMENT_BUILD` as a condition.  For more information, refer to [Platform dependent compilation](platform-dependent-compilation.md). |
| **Code Optimization** | Select the optimization mode to use for compiling the Web code.  * **Shorter Build Time**: This is the default setting. Select this option to generate Web code that takes the least amount of time to build. * **Runtime Speed**: Select this option to generate Web code that’s optimized for runtime performance, at the expense of taking a longer time to build. * **Runtime Speed with LTO**: This is the same as Runtime Speed, but with an additional Link Time Optimizations (LTO) stage (sometimes called Whole Program Optimization). Enable LTO when compiling shipping builds for deployment to end users. * **Disk Size**: Select this option to favor optimizations that minimize build size, at the expense of taking a longer time to build. It’s recommended to select this option when targeting mobile web browsers that might have limits on the size of WebAssembly files that can be loaded. Smaller disk sizes generally result in shorter page startup times. * **Disk Size with LTO**: This is the same as Disk Size but enables an additional Link Time Optimizations (LTO) stage. This is also known as Whole Program Optimization. Select this option when compiling shipping builds for deployment to end users. |
| **Autoconnect Profiler** | Automatically connect the Unity Profiler to your build. For more information, refer to [Profiler](Profiler.md).  **Note**: This option is available only if you select **Development Build**. |
| **Deep Profiling** | Allow the Profiler to process all your script code and record every function call, returning detailed profiling data. For more information, refer to [Deep Profiling](ProfilerWindow.md#deep-profiling).   This property is available only if you enable **Development Build**.   **Note**: Enabling **Deep Profiling** might slow down script execution. |

## Build options

The following build options are available only for the Web platform. For all other build options, refer to [Build Profiles window reference](build-profiles-reference.md#build-options).

| **Property** | **Description** |
| --- | --- |
| **Publish to Play** | Opens a window where you can upload your current build or a different one to [Unity Play](https://play.unity.com/en).  After the upload completes, the Unity Play website launches in your default browser, where you can finish the publishing process. You can choose to upload the build as a new game or replace an existing game.  **Notes**:  * Unity Play logs you in with your Engine account. * Web builds must be under 1 GB in size to upload to Unity Play. |

## Additional resources

* [Build and distribute a Web application](webgl-building-distribution.md)
* [Build your Web application](webgl-building.md)
* [Deploy a Web application](webgl-deploying.md)
* [BuildOptions API reference](../ScriptReference/BuildOptions.md)