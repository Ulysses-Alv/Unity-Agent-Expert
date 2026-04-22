# Profiler window reference

The Profiler window has the following areas:

* [Captures list](#captures-list): Contains a list of saved Profiler sessions in the `ProfilerCaptures` folder of your project.
* [Profiler modules](#profiler-modules): Add and remove modules to the window which collect different types of data related to your application. Each module has its own chart in the top half of the window.
* [Profiler controls](#profiler-controls): At the top of the window there are controls which you can use to stop, start, and navigate profiling sessions
* Module details pane: The bottom half of the Profiler window displays detailed information about the selected frame, depending on the selected module.

![Profiler window with a frame in the CPU Usage Profiler module selected.](../uploads/Main/profiler-cpu-module.png)

Profiler window with a frame in the CPU Usage Profiler module selected.

## Captures list

When you [save a Profiler capture](#profiler-controls), Unity saves it as a `.data` file to the `ProfilerCaptures` folder of your project. You can also save files in the `.raw` format with the [`-profiler-log-file` command](profiler-command-line-arguments.md). The **Captures List** displays all the Profiler sessions in the `ProfilerCaptures` folder. You can change the save location in the [Profiler Preferences window](profiler-preferences-reference.md).

| **Operation** | **Description** |
| --- | --- |
| **Import** | Select a `.data` file, or a [`.raw` file](profiler-command-line-arguments.md) to import into your project. Unity saves it into the `ProfilerCaptures` folder. |
| **Context menu** | Right-click on a capture to edit its details:  * **Delete**: Permanently deletes the §`.data` file from the `ProfilerCaptures` folder. * **Rename**: Rename the file. * **Open Folder**: Opens the containing folder for the capture. * **Target Frame Time:** Select a target frame time from a default list, or add your own custom frame time to benchmark the data in the capture. The target frame time is used in the [Highights module](ProfilerHighlights.md) |

**Tip:** You can use the [`SetScreenshotCaptureFrameInterval`](../ScriptReference/Profiling.Profiler.SetScreenshotCaptureFrameInterval.md) API to set the frequency at which Unity captures the screenshot for the thumbnail in the list.

## Default Profiler modules

Use the dropdown to add or remove modules to the Profiler. Profiler modules display performance data over time on a frame-by-frame basis in the charts in the top half of the window.

When you select a profiler module its chart displays in the top half of the profiler window. The bottom half of the window contains a module details panel which displays information related to the Profiler module selected. This area is blank when you open the Profiler for the first time, and fills with information when you start profiling your application.

The Profiler includes the following Profiler modules by default. For more information about how to use Profiler modules, refer to [Profiler modules introduction](profiler-modules-introduction.md).

| **Module** | **Description** |
| --- | --- |
| **Highlights** | Displays information on whether your application is meeting its target frame rate and if its performance is bound by the CPU or the GPU. For more information, refer to the [Highlights Profiler module](ProfilerHighlights.md). |
| **CPU Usage** | Displays an overview of where your application spends the most time, in areas such as physics, **scripts**, animation, and garbage collection. This module contains broad profiling information about your application, and you can use it to decide which further modules to use to investigate more specific issues in your application. This module is always active, even if you close it. For more information, refer to [CPU Usage Profiler module](ProfilerCPU.md). |
| **GPU Usage** | Displays information related to graphics processing. By default this module isn’t active, because it has a high overhead. For more information, refer to [GPU Usage Profiler module](ProfilerGPU.md). |
| **Rendering** | Displays information on how Unity renders graphics in your application. For more information, refer to [Rendering Profiler module](ProfilerRendering.md). |
| **Memory** | Displays information on how Unity allocates memory in your application. This module is useful to investigate how scripting allocations lead to [garbage collection](performance-garbage-collector.md), or how your application’s asset memory usage trends over time. For more information, refer to [Memory Profiler module](ProfilerMemory.md). |
| **Audio** | Displays information related to the audio in your application, such as how much CPU usage the audio system requires, and how much memory Unity allocates to it. For more information, refer to [Audio Profiler module](ProfilerAudio.md). |
| **Video** | Displays information related to [video](VideoPlayer.md) in your application. For more information, refer [Video Profiler module](profiler-video-profiler-module.md). |
| **Physics** | Displays information about the physics in your application that the physics system has processed. For more information, refer to [Physics Profiler module](ProfilerPhysics.md). |
| **Physics (2D)** | Displays information about where the physics system has processed 2D physics in your application. For more information, refer to [2D Physics Profiler module](2d-physics/physics-profiler/physics-2d-profiler-module-reference.md). |
| **UI (Canvas)** | Displays information about how uGUI handles UI batching for your application, including why and how uGUI batches items. For more information, refer to [UI (Canvas) and UI Details (Canvas) Profiler module](https://docs.unity3d.com/Packages/com.unity.ugui@2.0/manual/ProfilerUI.html). |
| **UI Details (Canvas)** | This module’s chart adds data about batch and vertices count, and markers which include information about user input events that trigger uGUI’s UI changes. For more information, refer to [UI (Canvas) and UI Details (Canvas) Profiler module](https://docs.unity3d.com/Packages/com.unity.ugui@2.0/manual/ProfilerUI.html). |
| **Realtime GI** | Displays information on how much CPU resource Unity spends on the **Global Illumination** lighting subsystem in your application. For more information, refer to [Global Illumination Profiler window](ProfilerGI.md). |
| **Virtual Texturing** | Displays statistics about [Streaming Virtual Texturing](svt-streaming-virtual-texturing.md) in your application. For more information, refer to [Virtual Texturing Profiler module](profiler-virtual-texturing-module.md). |
| **File Access** | Displays information about file accesses in your application. For more information, refer to [File Access Profiler module](profiler-file-access-module.md). |
| **Asset Loading** | Displays information about how your application loads assets. For more information, refer to [Asset Loading Profiler module](profiler-asset-loading-module.md). |
| **Profiler module editor (⚙)** | Open the [Profiler module editor](profiler-module-editor.md) to customize the Profiler modules in the list. |
| **Restore defaults** | Select Restore Defaults to remove any custom Profiler modules and reorder the module list to its default ordering. |

## Profiler controls

The Profiler controls are in the **toolbar** at the top of the Profiler window. Use these to start or stop recording profiler data, and to navigate through profiled frames.

| **Property** | **Description** |
| --- | --- |
| **Target Selection dropdown** | Use the Target Selection dropdown to select a platform or player for the Profiler to collect data for:   * **Search**: Use the search bar to find a player. You can search by its name or device category, for example `Remote`. When you search by category, the result displays all devices in that category. * **Player Name**: Displays any running players that Unity detects through the network or by direct connection. You can identify these players by their name and the platform that’s running the player, for example `iPhonePlayer (My iPhone)`. You can override this name in the [Profiler Preferences](profiler-preferences-reference.md) window. * **Play Mode (default)**: Profile your application in Play mode. This is the default mode, but you can change this behavior in the Preferences window. For more information, refer to [Collect performance data in Play mode](profiling-play-mode.md). * **Edit Mode**: Profile the Unity Editor and display the resources that the Editor is currently using. For more information, refer to [Collect performance data about the Unity Editor](profiling-edit-mode.md). * **Direct Connection**: Displays any devices directly connected to your computer. Select **Enter IP** to manually enter the IP address of the device you want to profile your application on. For more information, refer to [Collect performance data on a target platform](profiling-target-device.md). |
| **Record (⏺)** | Enable this setting to record profiling information for the [active modules](profiler-modules-activate.md) when you run your application. If Record is disabled, the Profiler doesn’t collect any data when you run your application. |
| **Previous frame (back arrow)** | Navigate one frame back. |
| **Next frame (forward arrow)** | Navigate one frame forward. |
| **Current frame (⏭)** | When you select the Current Frame button, the frame indicator line jumps to the last recorded frame, and the Profiler enters Current Frame mode. While the Profiler collects data in this mode, it stays on the current frame and displays the data it collects in real-time. Select the button again to exit Current Frame mode. |
| **Frame number** | Indicates the selected frame’s number. The number on the left is the current frame selected, and the number on the right is the total number of frames combined that the Profiler collected during your entire profiling session. |
| **Clear** | Erase all data from the Profiler window. |
| **Clear on Play** | Enable this setting to erase all data from the Profiler window next time you click Play in the Player window, or when you connect to a new target device. |
| **Deep Profile** | Enable this setting to profile all C# methods. When you enable this setting, Unity adds instrumentation to all mono calls, which then allows for a more detailed investigation of your scripts. For more information, refer to [Deep Profiling](profiler-deep-profiling.md). |
| **Call Stacks** | Select samples to record full call stacks for scripting memory allocations. Frames that the Profiler records when you enable this option have information about the selected samples on the full call stack that lead to a managed scripting allocation, even when the [Deep Profiling setting](profiler-deep-profiling.md) isn’t active.   You can select the following values, and have more than one selection active at once:   * **GC.Alloc**: Enables call stacks for all managed memory allocations. Unity emits `GC.Alloc` samples any time it makes a [managed memory allocation](performance-managed-memory.md). * **UnsafeUtility.Malloc (Persistent)**: Enables call stacks for [native memory allocations](performance-native-allocators.md) that are made with [`Allocator.Persistent`](https://docs.unity3d.com/Packages/com.unity.collections@latest?subfolder=/manual/allocator-overview.html). While the Temp and TempJop allocators are faster, they only allow for short lived allocations and to keep them as fast as possible, Unity doesn’t instrument them with call stacks. * **JobHandle.Complete**: Enables call stacks for jobs completed with [`JobHandle.Complete`](../ScriptReference/Unity.Jobs.JobHandle.Complete.md). Whenever a job uses `Complete` it results in idle time on the thread, so to optimize the code, you can enable call stacks to discover where in your code the sync point happened. |
| **Load (square and arrow)** | Load saved Profiler data into the Profiler window. You can also load binary profile data that the Player has written out to file via the [Profiler.logFile](../ScriptReference/Profiling.Profiler-logFile.md) API.  Hold down the Shift button and click the Load button to append the file contents to the current profile frames. |
| **Save (💾)** | Save the Profiler data into a .data file in your Project folder. |

## More (⋮) menu settings

The More menu contains the following settings:

| **Setting** | **Description** |
| --- | --- |
| **Color Blind Mode** | Enable this setting to make the Profiler use higher contrast colors in its graphs. This enhances visibility for users with red-green color blindness (such as deuteranopia, protanopia, or tritanopia). |
| **Show stats for “current” frame** | By default, when you select the Current Frame button and enter Current Frame mode, the frame indicator line doesn’t have annotations with the stats for the current frame. This is because the stats annotations might make it difficult to view data in real-time. To display the annotations, enable this setting. |
| **Preferences** | Open the [Preferences window](profiler-preferences-reference.md) to adjust Profiler-specific properties. |

## Additional resources

* [Navigating the Profiler window](profiler-window-navigating.md)
* [Profiler Preferences reference](profiler-preferences-reference.md)
* [Profiling your application](profiler-profiling-applications.md)