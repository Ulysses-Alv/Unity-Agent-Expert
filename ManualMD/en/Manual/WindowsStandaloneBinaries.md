# Windows build settings reference

Use these settings to configure how Unity builds your application for Windows platform. You can update your Windows build settings from the [Build Profiles](BuildSettings.md) window.

| **Property** | **Description** |
| --- | --- |
| **Architecture** | Select the CPU to build for (only applies to Build And Run).  * **Intel 64-bit**: Use Intel/AMD 64-bit CPU architecture. * **Intel 32-bit**: Use 32-bit Intel CPU architecture. * **ARM 64-bit**: Use 64-bit ARM CPU architecture. |
| **Build and Run on** | Select the target device or transport to deploy and launch the app during **Build And Run**.  * **Local Machine**: Deploys and launches the app on the local PC. * **Remote Device (via Device Portal)**: Deploys and launches the app to a connected device over the Device Portal transport. |
| **Copy PDB files** | Enable this setting to include Microsoft program database (PDB) files in the built Player. PDB files contain debugging information for your application, but might increase the size of your Player. |
| **Create Visual Studio Solution** | Enable this setting to generate Visual Studio Solution files for your Project, so you can build your final executable in Visual Studio. |
| **Development Build** | Include scripting debug symbols and the [Profiler](Profiler.md) in your build. Use this setting when you want to test your application. When you select this option, Unity sets the `DEVELOPMENT_BUILD` scripting define symbol. Your build then includes preprocessor directives that set `DEVELOPMENT_BUILD` as a condition.  For more information, refer to [Platform dependent compilation](platform-dependent-compilation.md). |
| **Autoconnect Profiler** | Automatically connect the Unity Profiler to your build. For more information, refer to [Profiler](Profiler.md).  **Note**: This option is available only if you select **Development Build**. |
| **Deep Profiling** | Allow the Profiler to process all your script code and record every function call, returning detailed profiling data. For more information, refer to [Deep Profiling](ProfilerWindow.md#deep-profiling).   This property is available only if you enable **Development Build**.   **Note**: Enabling **Deep Profiling** might slow down script execution. |
| **Script Debugging** | Attach script debuggers to the Player remotely.   This property is available only if you enable **Development Build**. |
| **Wait for Managed Debugger** | Make the Player wait for a debugger to be attached before it executes any script code.  This property is visible only if you enable **Script Debugging**. |
| **Compression Method** | Specifies the method Unity uses to compress the data in your Project when it builds the Player. This includes [Assets](assets-supported-types.md), [Scenes](CreatingScenes.md), [Player settings](class-PlayerSettings.md), and [GI data](GICache.md).  * **Default**: On Windows, Mac, Linux Standalone, and iOS, there’s no default compression. On Android, the default compression is ZIP, which gives slightly better compression results than LZ4HC. However, ZIP data is slower to decompress. * **LZ4**: A fast compression format that is useful for development builds. For more information, refer to [BuildOptions.CompressWithLz4](../ScriptReference/BuildOptions.CompressWithLz4.md). * **LZ4HC**: A high compression variant of LZ4 that is slower to build but produces better results for release builds. For more information, refer to [BuildOptions.CompressWithLz4HC](../ScriptReference/BuildOptions.CompressWithLz4HC.md). |

## Windows Player build binaries

When you build a Unity title on the Windows platform, Unity produces the following files, where `ProjectName` is the name of your application:

* `ProjectName.exe` - The project executable file, that’s your project application. This contains the program entry point which calls into the Unity engine when launched. The path to the source code for `ProjectName.exe` is in the `WindowsPlayer` folder: `Editor\Data\PlaybackEngines\WindowsStandaloneSupport\Source\WindowsPlayer`.
* `UnityPlayer.dll` - The DLL file that contains all the native Unity engine code. It’s signed with the Unity Technologies certificate, which lets you verify that no malicious entities have tampered with your engine.
* `*.pdb files` - Symbol files for debugging. Unity copies these to the build directory if you enable **Copy PDB files** in the **Build Profiles** window.
* `WinPixEventRuntime.dll` - This DLL enables [Introducing PIX on Windows (beta)](https://blogs.msdn.microsoft.com/pix/2017/01/17/introducing-pix-on-windows-beta/) support. Unity only creates this file if you enable the ****Development Build**** option in the **Build Profiles** window.
* `ProjectName_Data folder` - This folder contains all the data needed to run your application.

## Modify and rebuild your application

To modify your application or distribute the built code (for example, if you want to sign it or modify the built executable metadata), you must rebuild your project and place the output in your built game directory.

Modifying the built executable metadata, such as the publisher or company name is necessary to ensure that your application is correctly identified by the system. For example, various operating system tools and dialogs, such as Windows firewall use the executable metadata to display your application’s information, such as the publisher or company name. The Unity Player executable included with the Unity installation is pre-built and might not reflect details specific your application. In this case, you can modify the executable metadata to include your application’s information.

To modify the metadata, follow these steps:

1. Enable the **Create Visual Studio Solution** setting to generate Visual Studio Solution files for your project and build your project in the Unity Editor.
2. Open the build folder and locate a subfolder with the same name as your project. In the subfolder, find the `WindowsPlayerVersion.rc` file and open it with a text editor.
3. Remove the lines of code that include conditional compilation directive `#if defined(BUILDING_UNITY_PLAYER_DLL)` and the corresponding `#endif`.
4. Modify any metadata, such as `CompanyName`, `FileVersion`, `FileDescription`, or `ProductName` as required. For more information on these fields, refer to the documentation on [StringFileInfo BLOCK Statement](https://learn.microsoft.com/en-us/windows/win32/menurc/stringfileinfo-block).
5. Open the Visual Studio Solution from the build folder using Visual Studio and rebuild your project.

**Note**: To build your application outside of Unity, you need [Visual Studio](https://www.visualstudio.com/) 2019 or 2022 with Desktop development with C++ workload installed.

## Additional resources

* [Visual Studio project generation for Windows](VisualStudioprojectgenerationWindows.md)
* [BuildOptions API reference](../ScriptReference/BuildOptions.md)