# Unity Editor command line arguments reference

You can launch the Unity Editor and Player with additional command line arguments to control startup behavior, configure settings, or automate workflows.

To use these arguments, append them to the Unity executable path when launching from a terminal or command prompt. For example: `unity.exe -batchmode -quit -executeMethod MyClass.MyMethod`.

For more information, refer to [Run Unity from a command-line interface](command-line-run-unity.md).

* [Configuration arguments](#configuration-arguments)
* [Batch mode arguments](#batchmode)
* [Package Manager arguments](#package-manager-arguments)
* [Build arguments](#build-arguments)
* [Unity Accelerator cache server arguments](#accelerator)
* [Debugging arguments](#debugging)
* [Graphics API arguments](#gfx-api-arguments)
* [License arguments](#license-arguments)
* [Metal arguments (macOS only)](#metal-arguments-macos-only)
* [Profiler arguments](#profiler-arguments)
* [Unity Editor special command line arguments](#unity-editor-special-command-line-arguments)
* [Unity Hub special command line arguments](#unity-hub-special-command-line-arguments)
* [Additional resources](#additional-resources)

## Configuration arguments

You can run the Unity Editor and build Unity applications with additional commands and information on startup. This page lists the command line arguments you can use to launch and configure Unity Editor and Unity Player instances.

| **Command** | **Description** |
| --- | --- |
| `-createProject <pathname>` | Create an empty project at the given path. |
| `-consistencyCheck` | Performs an [importer consistency check](ImporterConsistency.md) on Editor startup. By default Unity performs a `local` check, which forces a reimport of all assets locally. It also checks if the import differs from the previous import.  **Note:** Consistency checking mode disables uploads to [Unity Accelerator](UnityAccelerator.md) and is intended to be used to test that importer output is deterministic. For more information, refer to the description for `-consistencyCheckSourceMode`. |
| `-consistencyCheckSourceMode string` | Sets the source to check against when Unity compares the asset imports. Choose from the following values:  * `local`: Forces a reimport of all assets locally and checks if the import differs from the previous import. * `cacheserver`: If you have enabled [Unity Accelerator](UnityAccelerator.md), this option gets the [metadata for the assets](AssetMetadata.md) from the cache server and compares if the results locally match what is on the cache server. This option only reimports those assets which have a presence on the cache server. If the consistency check succeeds then the original import result is kept, and no upload is required because the result is consistent and already cached. If the check fails, then an Editor log message is logged detailing which imports failed. The inconsistent result is not uploaded, to prevent cache poisoning.  For more information, refer to [Check the consistency of an importer](ImporterConsistency.md). |
| `-disable-assembly-updater <assembly1 assembly2>` | Specify a space-separated list of assembly names as parameters for Unity to ignore on automatic updates.  The space-separated list of assembly names is optional: pass the command line options without any assembly names to ignore all assemblies, as in example 1.  **Example 1** `unity.exe -disable-assembly-updater`  **Example 2** `unity.exe -disable-assembly-updater A1.dll subfolder/A2.dll`  Example 2 has two assembly names, one with a pathname. Example 2 ignores `A1.dll`, no matter what folder it is stored in, and ignores `A2.dll` only if it is stored under `subfolder` folder:  If you list an assembly in the `-disable-assembly-updater` command line parameter (or if you don’t specify assemblies), Unity logs the following message to [Editor.log](log-files.md):  `[Assembly Updater] warning: Ignoring assembly [assembly_path] as requested by command line parameter.”).`  Use this to avoid unnecessary [API Updater](APIUpdater.md) overhead when you import assemblies.  This argument is useful if you want to import assemblies that access a Unity API which doesn’t need updating. It’s also useful when you import assemblies that don’t access any Unity APIs (for example, if you have built some or all your source code outside of Unity, and you want to import the resulting assemblies into your Unity project).  **Notes**:  * If you disable the update of any assembly that needs updating, you might get errors at compile time, runtime, or both, that are hard to track. * This argument only disables the AssemblyUpdater component of the [API Updater](APIUpdater.md). The ScriptUpdater component of the API Updater still runs for the code in your project, including packages, unless you run the Editor in batch mode without the `-accept-apiupdate` argument. For more information, refer to [Batch mode arguments](#batchmode). |
| `-disable-gpu-skinning` | Disable Graphics Processing Unit (GPU) **skinning** at startup. |
| `-disable-playback-engines <PlaybackEngineNames>` | Disable specified playback engines (build modules).   This argument is useful for preventing Unity from generating platform-specific metadata for untargeted platforms, removing the need to uninstall build modules for the same outcome.   This argument doesn’t have a limit to the number of passed parameters. Parameters must match the name of the build module’s install directory and are case-sensitive.  **Example** `unity.exe -disable-playback-engines AndroidPlayer iOSSupport` |
| `-executeMethod <ClassName.MethodName>` or `-executeMethod <NamespaceName.ClassName.MethodName>` | Execute the static method as soon as Unity opens the project, and after the optional Asset server update is complete. You can use this for tasks such as continuous integration, performing Unit Tests, making builds or preparing data. If `-activeBuildProfile` is provided, Unity compiles scripts with the profile’s scripting defines before running `-executeMethod`.  To return an error from the command line process, either throw an exception which causes Unity to exit with return code 1, or call [`EditorApplication.Exit`](../ScriptReference/EditorApplication.Exit.md) with a non-zero return code.  To pass parameters, add them to the command line and retrieve them inside the function using `System.Environment.GetCommandLineArgs`. To use `-executeMethod`, you need to place the enclosing script in an Editor folder. The method you execute must be defined as static. |
| `-job-worker-count <N>` | Specify the maximum thread count for the Unity JobQueue [Job Worker Count](../ScriptReference/Unity.Jobs.LowLevel.Unsafe.JobsUtility.JobWorkerMaximumCount.md). |
| `-gc-helper-count <N>` | Specify the number of asset garbage collector helper threads. The default value is the number of cores available. |
| `-logFile <pathname>` | Specify where Unity writes the Editor or operating system’s standalone log file. To output to the console, specify `-` for the path name. On Windows, use `-logfile` to direct the output to `stdout`, which by default isn’t the console. Also refer to `-upmLogFile`. |
| `-noUpm` | Disables the Unity Package Manager. |
| `-openfile <path>` | Open the project from a path to a scene or package file. Alternatively, you can use the `-projectPath` argument |
| `-password <password>` | Enter a password into the log-in form during the activation of the Unity Editor. |
| `-projectPath <pathname>` | Open the project at the given path, which can be absolute or relative to the current working directory. If the pathname contains spaces, enclose it in quotes. |
| `-quit` | Quit the Unity Editor after other commands have finished executing. This can hide some error messages, but they still appear in the Editor’s [log file](log-files.md). The `-quit` argument has the following limitations:  * If the Editor is connected to the Accelerator Cache Server, then the additional argument `-cacheServerWaitForUploadCompletion` is required to prevent Unity quitting before pending cache uploads can complete. * If the Editor is running tests with the `-runTests` argument, then `-quit` causes the Editor to quit immediately, before in-progress tests have chance to complete. * If the Editor is running asynchronous code, then `-quit` can cause the application to hang and become unresponsive. |
| `-quitTimeout` | When you use the `-quit` argument the Editor waits for pending Async tasks to finish with a default timeout of 300 seconds. Use this parameter to set a different timeout in seconds. |
| `-releaseCodeOptimization` | Enables release code optimization mode, overriding the current default code optimization mode for the session. |
| `-setDefaultPlatformTextureFormat` (Android only) | Set the default texture ****compression**** to the desired format before you import a texture or build the project. This is so you don’t have to import the texture again with the format you want. The available formats are dxt, atc, etc, etc2, and astc. This argument is ignored when [texture compression targeting](android-distribution-google-play.md#texture-compression-targeting) is enabled. |
| `-overrideMaxTextureSize` | Overrides the maximum texture size (in pixels) that Unity applies when importing assets. This setting affects all textures in your project, and overrides any **Max Size** configured for individual textures in the [texture import settings](class-TextureImporter.md). The value is an integer that represents a maximum limit on either the height or width of the texture, whichever is longer. **Aspect ratio** is always maintained when resizing. For example, if a texture is 2048×1024 **pixels** but `-overrideMaxTextureSize` is set to 512, Unity imports the texture at a size of 512×256 pixels. For more details, refer to [EditorUserBuildSettings.overrideMaxTextureSize](../ScriptReference/EditorUserBuildSettings-overrideMaxTextureSize.md). |
| `-overrideTextureCompression` | Overrides the **texture compression** settings that Unity uses when it imports assets. This is particularly useful during local development, to speed up texture importing or build target switching. For details, refer to [EditorUserBuildSettings.overrideTextureCompression](../ScriptReference/EditorUserBuildSettings-overrideTextureCompression.md). |
| `-silent-crashes` | Prevent Unity from displaying the dialog that appears when a Standalone Player crashes. This argument is useful when you want to run the Player in automated builds or tests, where you don’t want a dialog prompt to obstruct automation. |
| `-upmLogFile <pathname>` | Specify the path and file name where Unity writes the log file for the UPM background application, which runs independently of the Editor. You can specify `-upmLogFile` and `-logFile` in the same command. If you include both arguments, you can specify independent paths for the two log files. |
| `-username <email_address>` | Enter the email address associated with the Unity ID you use to log in to Unity products and services.  For more information on using this option when activating a Unity license through the command line, refer to [Manage your license through the command line](ManagingYourUnityLicense.md).  For more information on using this option when exporting and signing a UPM package you created, refer to [Export and sign your UPM package](cus-export.md#export-cli). |
| `-vcsMode <mode>` | Set the version control mode. Available modes are `"Visible Meta Files"`, `"Hidden Meta Files"`, `Perforce`, and `PlasticSCM`. You can use additional flags to fill out the configuration fields for the given version control mode. These flags are based on the [`Provider.GetActiveConfigFields`](../ScriptReference/VersionControl.Provider.GetActiveConfigFields.md) method. For example, you can use the `-vcPerforceUsername`, `-vcPerforcePassword`, `-vcPerforceWorkspace` and `-vcPerforceServer` to set the Perforce username, workspace and server fields.  **Note**: `<mode>` arguments that contain spaces must be wrapped in double quotes (“).| |`-vcsModeSession <mode>`|Set the version control mode for this session. Available modes are `"Visible Meta Files"`, `"Hidden Meta Files"`, `Perforce`, and `PlasticSCM`. You can use additional flags to fill out the configuration fields for the given version control mode. These flags are based on the [Provider.GetActiveConfigFields](../ScriptReference/VersionControl.Provider.GetActiveConfigFields.md) method. For example, you can use the `-vcPerforceUsername`, `-vcPerforcePassword`, `-vcPerforceWorkspace` and `-vcPerforceServer` to set the Perforce username, workspace and server fields.  **Note**: `<mode>` arguments that contain spaces must be wrapped in double quotes (”). |
| `-version` | Print the version number of the Unity Editor in the command line, without launching the Editor. |
| `-timestamps` | Prefixes every [Editor.log](log-files.md) message with the current timestamp and thread ID. |

## Batch mode arguments

Use the following arguments to configure Unity’s batch mode. Batch mode enables Unity to run predefined tasks without user input, which makes batch mode useful for automated tasks like testing.

**Important:** Unless subject to separate Commercial Terms with Unity, using Unity in batch mode is subject to Unity’s Terms of Service, including applicable Additional Terms. For more information, refer to the latest version of the [Unity Terms of Service](https://unity.com/legal/terms-of-service) and the [Additional Terms, including the Unity Editor Software Terms](https://unity.com/legal/additional-terms).

| **Command** | **Description** |
| --- | --- |
| `-accept-apiupdate` | Use this command line option to run [API Updater](APIUpdater.md) when Unity is launched in batch mode.  Example: `unity.exe -accept-apiupdate -batchmode [other params]`  The API Updater doesn’t run if you omit this command line argument when you launch Unity in batch mode. This might lead to compiler errors. |
| `-batchmode` | Run Unity in batch mode. In batch mode, Unity runs command line arguments without the need for human interaction. It also suppresses dialog windows that require human interaction (such as the **Save Scene** window). Always run Unity in batch mode when using command line arguments, because it allows automation to run without interruption.  When an exception occurs during execution of the script code, the Asset server updates fail, or other operations fail, Unity immediately exits with return code **1**.  In batch mode, Unity sends a minimal version of its log output to the console. However, the [Log Files](log-files.md) still contain the full log information. You can’t open a project in batch mode while the Editor has the same project open; only a single instance of Unity can run at a time.  To check whether the Editor or Standalone Player is running in batch mode, use the [Application.isBatchMode](../ScriptReference/Application-isBatchMode.md) operator.  If the project hasn’t yet been imported when using `-batchmode`, the target platform is the default one. To force a different platform, use the `-buildTarget` option. |
| `‑ignorecompilererrors` | When you use this argument, Unity continues to start your application even if there are compilation errors. |
| `-nographics` | When you run this in batch mode, Unity doesn’t initialize the graphics device. You can then run automated workflows on machines that don’t have a GPU. Automated workflows only work when you have a window in focus, otherwise you can’t send simulated input commands. `-nographics` doesn’t allow you to bake GI, because Enlighten Realtime Global Illumination requires a GPU for Meta Pass rendering (See the [Meta Pass section of the Lightmapping and Shaders page](MetaPass.md) for more information).   **Note**: Output logs are turned off in this mode. To enable the creation of output logs, specify a file location using the command `-logFile`. |

## Package Manager arguments

Use the following arguments to manage packages from the command line.

| **Command** | **Description** |
| --- | --- |
| `-exportPackage <exportAssetPath1 exportAssetPath2 ... exportAssetPathN exportFileName>` | Export an asset package (`.unitypackage`) from the specified path or paths within a project’s `Assets` folder. In this example, `exportAssetPath` refers to folders relative to the root of the Unity project, and `exportFileName` is the package name. This option exports only whole folders at a time. Use this command with the `-projectPath` argument. |
| `-importPackage <pathname>` | Import the given [asset package](AssetPackagesCreate.md). Unity doesn’t display any import dialog. |
| `-upmPack` | Use this command in batch mode to export and sign a UPM package you created. Upon successful completion, the command creates a tarball file at the specified location. The following example shows the mandatory parameters for this command:  `-batchmode -username <email_address> -password <your_password> -upmPack <package_folder_path> <output_path> -cloudOrganization <your_org_id>`  For more information, refer to [Export and sign your UPM package](cus-export.md#export-cli).   **Note:** When you use the `-upmPack` argument, you must also use the `-batchmode` argument. |

## Build arguments

Use the following arguments to build Players for various platforms from the command line. For more information about building Players from the command line, refer to [Create a build from the command line](build-command-line.md).

| **Command** | **Description** |
| --- | --- |
| `-activeBuildProfile <pathname>` | Set the [build profile](build-profiles.md) saved at the given path as an active build profile (for example, `-activeBuildProfile "Assets/Settings/Build Profiles/WindowsDemo.asset"`).    **Notes:**  * `-activeBuildProfile <pathname>` applies custom scripting defines found in the **Build data** section of the active build profile before compilation. These scripting defines are added to the current set of defines, and don’t override or delete the existing scripting defines in your project. * The path to the build profile asset must be relative to the root folder of the project. * When specified together with `-executeMethod`, Unity applies the profile’s scripting defines and triggers script compilation before invoking the method. The profile’s defines are added to the current project defines and are effective during the subsequent `-executeMethod` execution. This enables CI/CD workflows such as switching target platforms in a single batch session. |
| `-build <pathName>` | Build a Player from `activeBuildProfile <pathname>` to the specified target path. For some platforms, the path must end with a file extension specific to the target platform. For example, `-build path/to/your/build.exe`. For information on which platforms require the extensions, refer to [Platform build path reference](build-path-requirements.md). |
| `-buildLinux64Player <pathname>` | Build a 64-bit standalone Linux player (for example, `-buildLinux64Player path/to/your/build`). |
| `-buildLinuxHeadlessSimulation <pathname>` | Build a 64-bit Linux headless simulation player (for example, `-buildLinuxHeadlessSimulation path/to/your/build`). |
| `-buildOSXUniversalPlayer <pathname>` | Build a 64-bit standalone macOS player (for example, `-buildOSXUniversalPlayer path/to/your/build.app`). |
| `-buildTarget <name>` | Select an active build target to launch the Editor in. The options available depend on which build targets you have enabled in the Editor. These options correspond to the options available in API enumeration [`BuildTarget`](../ScriptReference/BuildTarget.md) as follows:  * `android` * `cloudrendering` * `ios` * `linux64` * `osxuniversal` * `tvos` * `visionos` * `windowsstoreapps` * `webgl` * `win` * `win64`  If you’re a developer with access to closed platforms, you might have other `-buildTarget` options available. For details, refer to the documentation for your target platform. |
| `-buildWindowsPlayer <pathname>` | Build a 32-bit standalone Windows player (for example, `-buildWindowsPlayer path/to/your/build.exe`). |
| `-buildWindows64Player <pathname>` | Build a 64-bit standalone Windows player (for example, `-buildWindows64Player path/to/your/build.exe`). |
| `-standaloneBuildSubtarget <name>` | Select an active build subtarget for the Standalone platforms before loading a project:  * Player * Server |

## Unity Accelerator cache server arguments

Use the following arguments to configure Unity’s use of the [Unity Accelerator](UnityAccelerator.md) cache server. These arguments override any corresponding configuration stored in the Editor Preferences.

| **Command** | **Description** |
| --- | --- |
| `-EnableCacheServer` | Enable Unity Accelerator. You must also use `-cacheServerEndpoint <host:port>` to specify the address of the Unity Accelerator instance. |
| `-cacheServerEndpoint <host:port>` | Specifies the endpoint address of the Unity Accelerator instance.  Example: `-cacheServerEndpoint 127.0.0.1:10080` The default port is `10080` but can be changed during installation. |
| `-cacheServerNamespacePrefix <prefix>` | Set the namespace prefix for the Unity Accelerator instance. Use this argument with a namespace name to partition data on the cache server, for example to isolate cache data for different projects or Unity versions.  Example: `-cacheServerNamespacePrefix MyProject_Unity23LTS`  A prefix that contains spaces must be wrapped in double quotes (`"`). The supplied name is used as the prefix for up to 3 namespaces in the cache which separate data by type - metadata, artifact data and shader cache data. |
| `-cacheServerEnableDownload true` `-cacheServerEnableDownload false` | Enable or disable downloading from the Unity Accelerator instance. |
| `-cacheServerEnableUpload true` `-cacheServerEnableUpload false` | Enable or disable uploading to the Unity Accelerator instance.  Example: `-cacheServerEnableUpload true` |
| `-cacheServerWaitForConnection <ms>` | Specify the time in milliseconds that Unity will wait for a connection to the Unity Accelerator instance while loading a project, before starting the initial project refresh.  Example: `-cacheServerWaitForConnection 5000` |
| `-cacheServerWaitForUploadCompletion` | Prevent Unity from closing until any pending Unity **Accelerator** uploads are complete. This is recommended when Unity Accelerator is used in combination with the command line argument `-quit`, which closes the Unity Editor after other scripted commands have finished executing. |
| `-cacheServerDownloadBatchSize <batchSize>` | Set the number of artifacts in individual cache server download batches. Download batches are an optimization used to reduce the number of download operations by grouping individual downloads into larger batches.  Example: `-cacheServerDownloadBatchSize 256` Default value is `128`. |
| `-cacheServerUploadExistingImports` | Upload any existing imports that haven’t been previously uploaded when Unity begins to upload new imports to the Unity Accelerator instance. Imports that have been previously uploaded aren’t re-uploaded. |
| `-cacheServerUploadAllRevisions` | Upload any existing imports including all previous revisions known to the Asset Database that haven’t previously been uploaded when Unity begins to upload new imports to the Unity Accelerator instance. Import revisions that have been previously uploaded, aren’t re-uploaded. |
| `-cacheServerUploadExistingShaderCache` | Upload any existing **shader** cache imports that haven’t previously been uploaded when Unity begins to upload new imports to the cache server. Shader imports that have been previously uploaded, aren’t re-uploaded. |

## Debugging arguments

| **Command** | **Description** |
| --- | --- |
| `-disableManagedDebugger` | Disables the debugger listen socket. |
| `-diag-debug-shader-compiler` | Unity launches only one instance of the **Shader** Compiler, and forces its timeout to be one hour. Useful for debugging Shader Compiler issues. |
| `-debugCodeOptimization` | Enables debug code optimization mode, overriding the current default code optimization mode for the session. |
| `-enableCodeCoverage` | Enables code coverage and allows access to the [Coverage API](../ScriptReference/TestTools.Coverage.md). |
| `-force-d3d12-debug` | Enables the DX12 validation layer. This is useful for working on **XR** **plug-ins** or **native plug-ins**. |
| `-force-d3d12-debug-gbv` | Enables the DX12 GPU-based validation. This is useful for working on XR plug-ins or native plug-ins. |
| `-force-vulkan-layers` | Enables the Vulkan validation layer. This is useful for working on XR plug-ins or native plug-ins. |
| `-stackTraceLogType` | Allow detailed debugging. All settings allow **None**, **Script Only** and **Full** to be selected (for example, `-stackTraceLogType Full`). |
| `-log-memory-performance-stats` | Adds detailed memory and performance reports to the main Unity Editor log file when closing the editor. |
| `-wait-for-managed-debugger` | Makes the Editor wait for a managed debugger to attach before launching. Generates a dialog that displays the Editor’s process ID. For more information on managed debugging, refer to [Debug C# code in Unity](managed-code-debugging.md#DebugInEditor). |
| `-wait-for-native-debugger` | Makes the Editor wait for a native debugger to attach before launching. Generates a dialog that displays the Editor’s process ID. |

## Graphics API arguments

Use the following arguments to force the Unity Editor to use a specific graphics API.

| **Command** | **Description** |
| --- | --- |
| `-force-clamped` | Use this with `-force-glcoreXY` to prevent Unity from checking for additional OpenGL extensions, allowing it to run between platforms with the same code paths. This approach helps to determine if an issue is platform-specific. For example, a driver bug. |
| `-force-d3d11` (Windows only) | Make the Editor use Direct3D 11 for rendering. Normally the graphics API depends on ****Player Settings**** (typically defaults to D3D11). |
| `-force-d3d12` (Windows only) | Make the Editor use Direct3D 12 for rendering. Normally the graphics API depends on **Player Settings**. |
| `-force-device-index` | Make the Editor use a particular GPU device by passing it the index of that GPU. For example, use this to force the Editor to use the integrated GPU instead of the dedicated one. Mapping between GPU indices and physical GPUs is determined by the order in which the GPUs are enumerated by your system. This enumeration is typically done by the GPU driver or the underlying framework being used (for example Metal, Direct3D or Vulkan). Normally, index `0` will be the GPU connected to your primary monitor. On desktop computers, this is usually the dedicated GPU, while on laptops it’s usually the integrated GPU. Therefore, on desktop computers `-force-device-index=0` is usually the dedicated graphics card. |
| `-force-glcore` | Make the Editor use OpenGL 3/4 core profile for rendering. The Editor tries to use the best OpenGL version available and all OpenGL extensions exposed by the OpenGL drivers. If the platform isn’t supported, the editor uses Direct3D. |
| `-force-glcoreXY` | Similar to`-force-glcore`, but requests a specific OpenGL context version. Accepted values for XY: 32, 33, 40, 41, 42, 43, 44, or 45. If the platform doesn’t support a specific version of OpenGL, Unity will fallback to a supported version. |
| `-force-gles` (Windows only) | Make the Editor use OpenGL for Embedded Systems for rendering. The Editor tries to use the best OpenGL ES version available, and all OpenGL ES extensions exposed by the OpenGL drivers. |
| `-force-glesXY` (Windows only) | Similar to `-force-gles`, but requests a specific OpenGL ES context version. Accepted values for XY: 30, 31, 31aep, or 32. If the platform doesn’t support a specific version of OpenGL ES, Unity will fallback to a supported version. If the platform doesn’t support OpenGL ES, Unity will fallback to another graphics API. |
| `-force-opengl` | Make the Editor use the legacy OpenGL back end. |
| `-force-vulkan` | Make the Editor use Vulkan for rendering. Normally the graphics API depends on **Player Settings**. |

## License arguments

Use the following arguments to process a Unity license or run the Unity Editor with different license options.

| **Command** | **Description** |
| --- | --- |
| `-createManualActivationFile` | Step one of a [three-step process](ManualActivationGuide.md) to manually activate a Unity license. For more information, refer to [Manual license activation](ManualActivationGuide.md). |
| `-manualLicenseFile <yourulffile>` | Step three of a [three-step process](ManualActivationGuide.md) to manually activate a Unity license. For more information, refer to [Manual license activation](ManualActivationGuide.md). |
| `-returnlicense` | Return the currently active serial-based or named user license. Don’t use `-returnlicense` to return a floating license. For more information, refer to [Returning your license](ManagingYourUnityLicense.md#license-return-cli). For more information about licensing models, refer to [Licensing overview](LicenseOverview.md). |
| `-serial <serial-key>` | Activate your paid Unity license. Include the `<serial-key>` value to activate a serial-based license, or to activate a named user license and its serial key simultaneously. Exclude the `<serial-key>` value to activate a named user license without activating its serial key. For more information, refer to [Activate a license from the command line](ManagingYourUnityLicense.md#license-activation-cli). For more information about licensing models, refer to [Licensing overview](LicenseOverview.md).  **Note**: When you use the `-serial` argument, you must also use the `-batchmode` argument. It’s also good practice to specify the `-quit` argument. |

## Metal arguments (macOS only)

Use the following arguments to configure Unity’s use of the [Metal](Metal.md) graphics API for Apple devices.

| **Command** | **Details:** |
| --- | --- |
| `-force-low-power-device` (macOS only) | If you use Metal, make the Editor use a low power device. |
| `-force-metal` | Make the Editor use Metal as the default graphics API (macOS only). |
| `-enable-metal-capture` | Enable the ability to take Metal captures from inside the Editor. |

## Profiler arguments

Use the following arguments to configure Unity’s use of the [Profiler](Profiler.md).

| **Command** | **Description** |
| --- | --- |
| `-deepprofiling` | Enable Deep Profiling option for the [CPU profiler](ProfilerWindow.md). |
| `-profiler-enable` | Profile the start-up of a Player or the Editor. When you use this argument with a Player, it has the same effect as building the Player with the **Autoconnect Profiler** option enabled in [Build Profiles](build-profiles.md).  When you use this argument with the Editor, it starts collecting and displaying Profiler information in the Profiler window on start-up of the Editor. |
| `-profiler-log-file <Path/To/Log/File.raw>` | This argument sets up the [Unity Profiler](Profiler.md) to stream the profile data to a .raw file on startup. It works for both Players and the Editor. |
| `-profiler-capture-frame-count <NumberOfFrames>` | This argument sets how many frames the Profiler should capture in a profile when streaming to a .raw file on start-up. It only works on Players. |
| `-profiler-maxusedmemory` | By default, `maxUsedMemory` for the Unity Profiler is 16MB for Players and 256MB for the Editor. You can use this argument to set the `maxUsedMemory` parameter to a custom size at start-up (for example, `-profiler-maxusedmemory 16777216`). The size is set in bytes. |

## Unity Editor special command line arguments

Only use these under special circumstances, or when directed by Unity Support.

| **Command** | **Description** |
| --- | --- |
| `-enableIncompatibleAssetDowngrade` | Use this when you have Assets made by a newer, incompatible version of Unity, that you want to downgrade to work with your current version of Unity. When you enable this, Unity presents you with a dialog asking for confirmation of this downgrade if you attempt to open a project that would require it.  Note: This procedure is unsupported and highly risky, and should only be used as a last resort. |
| `-giCustomCacheLocation` | Sets a custom location for the [Global Illumination (GI) cache](GICache.md) folder. The value must be an absolute path, which must be enclosed in quotes if it contains any spaces. If `-giCustomCacheLocation` is set, then for the current Editor session its value overrides any custom **Cache Folder Location** previously defined in the [GI Cache preferences](Preferences.md#gi-cache). This value doesn’t persist after the Editor quits. |
| `-rebuildLibrary` | This procedure deletes the `Library` project subdirectory, which triggers a full reimport of all assets. Use this when Unity reports a database corruption error during project initialization. |

## Unity Hub special command line arguments

Only use these under special circumstances, or when directed by Unity Support. These arguments are used in versions of the Unity Editor that don’t use the Unity Hub by default.

| **Command** | **Description** |
| --- | --- |
| `-useHub` | Launch the Unity Editor with Hub integration. |
| `-hubIPC` | Enable communication and interaction between the Unity Editor and the Unity Hub. |
| `-hubSessionId` | Specify a unique identifier for a Unity Hub session. |

In all supported versions of the Unity Editor, these arguments are already set by default.

## Additional resources

* [Create a build from the command line](build-command-line.md)
* [Android Player command-line arguments](android-custom-activity-command-line.md)
* [Use Unity Accelerator on the command line](accelerator-command-line.md)
* [Build for Embedded Linux from the command line](embedded-linux-build-command-line.md)
* [Build for QNX from the command line](qnx-build-command-line.md)
* [Test Framework command-line arguments reference](test-framework/reference-command-line.md)
* [Profiler command-line arguments reference](profiler-command-line-arguments.md)
* [Run Project Auditor from the command line](project-auditor/run-from-command-line)