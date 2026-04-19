# Editor

Use the **Editor** settings to apply global settings for working in the Unity Editor. To access the Editor settings, go to **Edit** > **Project Settings** then select the **Editor** category.

## Unity Remote

Unity Remote is deprecated. For most use cases, [Device Simulator package](https://docs.unity3d.com/Packages/com.unity.device-simulator@latest) replaces Unity Remote.

| **Property** | **Description** |
| --- | --- |
| **Device** | Choose the device type you want to use for Unity Remote testing.  [Unity Remote](UnityRemote5.md) is a downloadable app designed to help with Android, iOS and tvOS development. |
| **Compression** | Choose the type of image compression to use when transmitting the game screen to the device via Unity Remote. Choose from the following options:  * **JPEG**: Usually gives higher compression and performance, but the graphics quality is a little lower. This is the default option. * **PNG**: Gives a more accurate representation of the game display, but can result in lower performance. |
| **Resolution** | Choose the resolution the game runs at on Unity Remote:  * **Downsize**: Display the game at a lower resolution. This results in better performance, but lower graphical accuracy. This is the default option. * **Normal**: Display the game at normal resolution. This results in better graphical accuracy, but lower performance. |
| **Joystick Source** | Choose the connection source for the joysticks you’re using:  * **Remote**: Use joysticks that connect to a device running Unity Remote. This is the default option. * **Local**: Use joysticks that connect to the computer running the Editor. |

## Asset Serialization

| **Property** | **Description** |
| --- | --- |
| **Mode** | Choose which format to use for storing serialized Assets. This is set to **Force Text** by default.  Unity uses [serialization](script-serialization.md) to load and save assets and AssetBundles to and from your computer’s hard drive. To help with version control merges, Unity can store scene files in a [text-based format](TextSceneFormat.md). If you aren’t merging scenes, Unity can store scenes in a more space-efficient binary format, or allow both text and binary scene files to exist at the same time. Choose from the following formats:  * **Mixed**: Assets in Binary mode remain in Binary mode, and Assets in Text mode remain in Text mode. Unity uses Binary mode by default for new assets. * **Force Binary**: Convert all assets to Binary mode, including new assets. * **Force Text**: Convert all assets to Text mode, including new assets. This is the default option. |
| **Reduce **version control** noise** | Forces the Editor to write references and similar YAML structures on one line, which reduces version control noise. When Unity reaches a line length of 80 characters it attempts to split the YAML data over multiple lines. |

## Scene Handling

| **Property** | **Description** |
| --- | --- |
| **Force GC on Scene Loads** | Force a managed heap garbage collection and unload unused assets after loading each scene in the Editor in single mode or exiting [prefab editing mode](EditingInPrefabMode.md). In complex projects, this can slow down performance. Disable this if you do not require class finalizers to run at the end of scene loading. Note that regular dynamic GC and unused asset unloading still occur when memory usage is high, but without this option, class finalizers might not run immediately after each scene load. |

## Default Behavior Mode

| **Property** | **Description** |
| --- | --- |
| **Mode** | Choose the default [2D or 3D development](2DAnd3DModeSettings.md) mode. Unity sets up the certain default behaviors according to the mode you choose to make development easier:  * **3D**: Set Unity up for 3D development. This is the default option. * **2D**: Set Unity up for 2D development. |

## Asset Pipeline

| **Property** | **Description** |
| --- | --- |
| **Remove unused Artifacts on Restart** | Enable to remove unused artifact files in the Library folder and remove their entries in the Asset Database. This setting is a form of garbage collection. Use this setting to disable the Asset Database garbage collection, so that previous artifact revisions which are no longer used are still preserved after you restart the Editor. This is useful if you need to debug unexpected import results. To learn more, refer to [Import Activity window](ImportActivityWindow.md). |
| **Parallel Import** | Use multiple processes to import assets simultaneously. By default, Unity imports assets one after another sequentially on the main Editor process. Parallel importing can be faster than the default sequential method of importing. To learn more, refer to [Importing assets simultaneously](ParallelImport.md) |
| **Desired worker count** | Set the number of import worker processes that the import pipeline considers the optimal number to run in parallel. |
| **Standby Import Worker Count** | Set the minimum number of worker processes to keep, even if they’re idle. If there are more worker processes than this, Unity shuts down import workers that have been idle for some time, to free up system resources. This property allows you to manage how Unity balances system resources when some processes are idle, compared with the time it takes to start up new import worker processes. You might see an improvement in import performance by increasing this value if you are frequently iterating on model, animation or texture work, and are therefore frequently importing batches of models or image files. |
| **Idle Import Worker Shutdown Delay** | Set the amount of time in seconds to wait before shutting down an idle worker. |
| **Exactly Match Referenced Clip Names** | Controls how referenced clips are matched with models animated in Legacy mode. When enabled, model names and referenced clip names must exactly match. When disabled, only the start of the model name must match the referenced clip name. |

## Cache Server (project specific)

Contains settings for [Unity Accelerator](UnityAccelerator.md).

| **Setting** | **Description** |
| --- | --- |
| **Mode** | Choose whether to use the cache server. Choose from the following options:  * **Use global settings:** Uses the settings in the **Preferences** window (**Settings > Asset Pipeline > Unity Accelerator**). * **Enabled**: Enable the cache server. * **Disabled:** Disable the cache server. |
| **IP Address** | Input the **accelerator**’s IP address and port. |
| **Check Connection** | Select **Check Connection** to test the connectivity of the accelerator. |
| **Namespace prefix** | Set a custom namespace for the server. |
| **Download** | Enable downloads from the cache server. |
| **Upload** | Enable uploads from the cache server. |
| **TLS/SSL** | Enable encryption on the cache server. |
| **Authentication (Using Unity ID)** | Enable authentication for the cache server using Unity ID. |
| **Content Validation** | Select the level of content validation: Disabled, Upload Only, Enabled, Required. |
| **Download Batch Size** | Set the size of download. |

## Prefab Mode

Contains settings related to [prefab editing mode](EditingInPrefabMode.md).

| **Property** | **Description** |
| --- | --- |
| **Allow Auto Save** | Enable an **Auto Save** toggle in [prefab editing mode](EditingInPrefabMode.md). When **Allow Auto Save** is disabled, any edits you make aren’t automatically saved and the toggle doesn’t display. For more information, refer to [Automatically save edits to the prefab asset](EditingInPrefabMode.md#automatically-save-edits-to-the-prefab-asset). |
| **Regular Environment** | Assign a scene as an editing environment in **prefab** editing mode for regular prefabs (those with a regular [Transform](class-Transform.md) component). This allows you to edit a prefab against a backdrop of your choosing rather than an empty scene. For more information, refer to [Editing prefab assets](EditingInPrefabMode.md). |
| **UI Environment** | Assign a scene as an editing environment in prefab editing mode for UI prefabs (those with a [Rect Transform](https://docs.unity3d.com/Packages/com.unity.ugui@2.0/manual/class-RectTransform.html) component). This allows you to edit a prefab against a backdrop of your choosing rather than an empty scene. For more information, refer to [Editing prefab assets](EditingInPrefabMode.md). |

## Graphics

| **Property** | **Description** |
| --- | --- |
| **Use legacy **Light Probe** sample counts** | Use fixed Light Probe sample counts for baking with the Progressive **Lightmapper**. The sample counts are: 64 direct samples, 2048 indirect samples, and 2048 environment samples. |

## Sprite Atlas

| **Property** | **Description** |
| --- | --- |
| **Mode** | Choose a mode to configure the [Sprite Atlas](sprite/atlas/atlas-landing.md) tool:  * **Disabled**: Do not pack Sprite Atlases. This is the default setting. * **Sprite Atlas V1 - Enabled For Builds**: Pack Sprite Atlases for builds only. * **Sprite Atlas V1 - Always Enabled**: Pack Sprite Atlases for builds and before entering Play mode. * **Sprite Atlas V2 - Enabled**: Pack Sprite Atlases for both builds and before entering Play mode. * **Sprite Atlas V2 - Enabled For Builds**: Pack Sprite Atlases for builds only. |

## C# Project Generation

| **Property** | **Description** |
| --- | --- |
| **Additional extensions to include** | Include a list of additional file types to add to the C# project. Separate each file type with a semicolon. By default, this field contains `txt;xml;fnt;cd;asmdef;rsp;asmref`. |
| **Root namespace** | Fill in the namespace to use for the C# project `RootNamespace` property. Refer to [Common MSBuild Project Properties](https://docs.microsoft.com/en-us/visualstudio/msbuild/common-msbuild-project-properties?view=vs-2022) for more information. This field is blank by default. |

## Texture Compressors

| **Property** | **Description** |
| --- | --- |
| **BC7 Compressor** | Select the compressor to use for BC7 format texture **compression**. |
| **ETC Compressor** | Select the compressors to use for different compression qualities of ETC Textures. The compression tools available are [etcpak](https://bitbucket.org/wolfpld/etcpak/wiki/Home), [ETCPACK](https://github.com/Ericsson/ETCPACK) and [Etc2Comp](https://github.com/google/etc2comp). These are all third-party compressor libraries. Choose from the following:  * **Default**: Use the default configuration for Unity. This sets the following properties: **Fast**: etcpack, **Normal**: ETCPACK Fast, **Best**: Etc2Comp Best. * **Custom**: Customize the ETC Texture compression configuration. When you choose this option, the **Fast**, **Normal**, and **Best** properties are enabled. This maps to the **Compressor Quality** setting in the [Texture Importer](class-TextureImporter.md) for the supported platforms. |
| **Fast** | Define the compression quality tool to use for Fast compression. This property is modifiable only when **ETC Compressor** is set to **Custom**. |
| **Normal** | Define the compression quality tool to use for Normal compression. This property is modifiable only when **ETC Compressor** is set to **Custom**. |
| **Best** | Define the compression quality tool to use for Best compression. This property is modifiable only when **ETC Compressor** is set to **Custom**. |

## Line Endings For New Scripts

| **Property** | **Description** |
| --- | --- |
| **Mode** | Choose the file line endings to apply to new scripts created within the Editor. Note that configuring these settings does not convert existing scripts. Choose from the following:  * **OS Native**: Apply line endings based on the operating system the Editor is running on. * **Unix**: Apply line endings based on the Unix operating system. * **Windows**: Apply line endings based on the Windows operating system. |

## Texture Streaming Settings

If you enable Mipmap Streaming in the Editor for one mode but not the other, entering and exiting Play mode takes slightly longer. Enabling Mipmap Streaming for both modes prevents Unity from unloading and reloading mipmap data, and increases the speed of entering and exiting Play Mode.

| **Property** | **Description** |
| --- | --- |
| **Enable Mipmap Streaming in Play Mode** | Enable [Mipmap Streaming](TextureStreaming.md) in Play mode.   The **Texture Mipmap Streaming** setting in the [Quality](class-QualitySettings.md) project settings must also be enabled.  Play mode might take longer to open and close if you enable this setting and disable **Enable Mipmap Streaming in Edit Mode**. |
| **Enable Mipmap Streaming in Edit Mode** | Enable Mipmap Streaming in Edit Mode.   The **Texture Mipmap Streaming** setting in the [Quality](class-QualitySettings.md) project settings must also be enabled.   Play mode might take longer to open and close if you enable this setting but disable **Enable Mipmap Streaming in Play Mode**. |
| **Load texture data on demand** | Load CPU-side texture data from disk asynchronously on demand to avoid some stalls and reduce CPU memory usage. If you enable mipmap streaming, this setting requires more processing time on the CPU, and might cause textures to temporarily appear at a lower resolution while Unity loads a higher resolution mipmap level from disk. |

## Shader Compilation

| **Property** | **Description** |
| --- | --- |
| **Asynchronous **Shader** Compilation** | Enable this checkbox to make Unity compile shaders in the background. While compiling, the Unity engine uses a placeholder Shader to render geometry in the Editor. When Shader compilation has finished, the engine swaps your Shader Variant back in. This means the Editor runs seamlessly, without having to wait for the Unity engine to compile every single Shader variant before rendering. For more information, refer to [Asynchronous Shader compilation](AsynchronousShaderCompilation.md). |

## Enter Play Mode Settings

| **Property** | **Description** |
| --- | --- |
| **When entering play mode** | Select which reload options to start when you enter Play mode. [Domain reloading](domain-reloading.md) is when the Editor resets the scripting state before it starts Play mode. [Scene reloading](scene-reloading.md) is when the Editor destroys all scene GameObjects and reloads the scene from disk before Play mode starts. To enter Play mode faster, you can disable [scene](scene-reloading.md) or [domain](domain-reloading.md) reloads. To learn more, refer to [Details of disabling Domain and Scene Reload](configurable-enter-play-mode-details.md). Choose from the following options:  * **Reload Domain and Scene**: Reload both the [domain](domain-reloading.md) and [scene](scene-reloading.md) when you enter Play mode. **Reload Domain and Scene** is enabled by default. * **Reload Scene Only**: Reload the [scene](scene-reloading.md) and do not reload the [domain](domain-reloading.md) when you enter Play mode. * **Reload Domain Only**: Reload the [domain](domain-reloading.md) and do not reload the [scene](scene-reloading.md) when you enter Play mode. * **Do not reload Domain or Scene**: Do not reload either the [scene](scene-reloading.md) or [domain](domain-reloading.md) when you enter Play mode. |

## Numbering Scheme

| **Property** | **Description** |
| --- | --- |
| **Game Object Naming** | Naming scheme for duplicated **GameObjects**. Duplicated or instantiated GameObjects are named by appending successive numbers to the original object name. |
| **Game Object Digits** | Sets the amount of digits to use for duplicated GameObject numbers. |
| **Space Before Number in Asset Names** | Controls whether to insert a space before a number in duplicated Asset names. |

## Inspector

| **Property** | **Description** |
| --- | --- |
| **Use IMGUI Default **Inspector**** | Revert to using IMGUI to generate Default Inspectors where no custom Inspector or Editor is defined. |

## Additional resources

* [Project Settings](comp-ManagerGroup.md)
* [Unity Remote](UnityRemote5.md)
* [Script serialization](script-serialization.md)
* [2D and 3D mode settings](2DAnd3DModeSettings.md)
* [Importing assets simultaneously](ParallelImport.md)
* [Preferences](Preferences.md)
* [Transforms](class-Transform.md)
* [Cookies](Cookies.md)
* [Progressive Lightmapper](progressive-lightmapper.md)
* [The Mipmap Streaming system](TextureStreaming.md)
* [Asynchronous Shader compilation](AsynchronousShaderCompilation.md)
* [Sprite Atlas V2](sprite/atlas/v2/v2-landing)

EditorManager