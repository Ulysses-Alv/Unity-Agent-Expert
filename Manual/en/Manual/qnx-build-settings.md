# QNX build settings reference

Use the QNX build settings to configure and build your application for the QNX platform. The QNX build settings are included in the **Platform Settings** section of the [Build Profiles window](BuildSettings.md).

| **Property** | **Description** |
| --- | --- |
| **QNX Version** | The target OS version for QNX is automatically detected depending on the QNX environment. For example, Neutrino RTOS 7.0 or Neutrino RTOS 7.1. |
| **Architecture** | Choose the CPU architecture to build for the QNX platform:   * **x64**: 64-bit CPU * **arm64**: 64-bit ARM CPU |
| **Remote Device** | Deploy and launch the app to a connected device on a remote device based on the connection and authentication information you provide.  * **Address (required)**: The address of the remote device. You can also specify a port by adding a colon after the address (for example, 10.110.91.3:2121). * **Username (optional)**: The user name on the remote device. * **Exports**: Specify the additional exports, if they’re required to launch the device remotely. Multiple environment variables are separated by a space. For example, `ENV_VAR1=1 ENV_VAR2=2`. * **Install Path**: The installation path for the application. If not set, the default value `~tmp/unity/test` will be used. |
| **Development Build** | Include scripting debug symbols and the [Profiler](Profiler.md) in your build. Use this setting when you want to test your application. When you select this option, Unity sets the `DEVELOPMENT_BUILD` scripting define symbol. Your build then includes preprocessor directives that set `DEVELOPMENT_BUILD` as a condition.  For more information, refer to [Platform dependent compilation](platform-dependent-compilation.md). |
| **Autoconnect Profiler** | Automatically connect the Unity Profiler to your build. For more information, refer to [Profiler](Profiler.md).  **Note**: This option is available only if you select **Development Build**. |
| **Deep Profiling** | Allow the Profiler to process all your script code and record every function call, returning detailed profiling data. For more information, refer to [Deep Profiling](ProfilerWindow.md#deep-profiling).   This property is available only if you enable **Development Build**.   **Note**: Enabling **Deep Profiling** might slow down script execution. |
| **Script Debugging** | Attach script debuggers to the Player remotely.   This property is available only if you enable **Development Build**. |
| **Wait for Managed Debugger** | Make the Player wait for a debugger to be attached before it executes any script code.  This property is visible only if you enable **Script Debugging**. |
| **Compression Method** | Specifies the method Unity uses to compress the data in your Project when it builds the Player. This includes [Assets](assets-supported-types.md), [Scenes](CreatingScenes.md), [Player settings](class-PlayerSettings.md), and [GI data](GICache.md).  * **Default**: On Windows, Mac, Linux Standalone, and iOS, there’s no default compression. On Android, the default compression is ZIP, which gives slightly better compression results than LZ4HC. However, ZIP data is slower to decompress. * **LZ4**: A fast compression format that is useful for development builds. For more information, refer to [BuildOptions.CompressWithLz4](../ScriptReference/BuildOptions.CompressWithLz4.md). * **LZ4HC**: A high compression variant of LZ4 that is slower to build but produces better results for release builds. For more information, refer to [BuildOptions.CompressWithLz4HC](../ScriptReference/BuildOptions.CompressWithLz4HC.md). |

## Additional resources

* [Build Profiles](BuildSettings.md)
* [Build and deliver for QNX](qnx-build-and-deliver.md)
* [Project Settings](comp-ManagerGroup.md)