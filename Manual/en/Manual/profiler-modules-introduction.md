# Profiler modules introduction

Collect specific performance data about your application with Profiler modules.

The top area of the [Profiler window](ProfilerWindow.md) contains Profiler modules that profile specific areas of your application. When you profile your application, Unity displays the data related to each module in corresponding charts.

The Profiler includes [default Profiler modules](#profiler-modules). When you install a package, it might include its own module, for example, the [Entities Profiler modules](https://docs.unity3d.com/Packages/com.unity.entities@latest?subfolder=/manual/profiler-modules-entities.html) or the [Adressables profiler module](https://docs.unity3d.com/Packages/com.unity.addressables@latest?subfolder=/manual/load-profiler-module.html). You can also [create a custom Profiler module](#custom-profiler-modules).

![Profiler window with a frame in the CPU Usage Profiler module selected. The Timeline view is selected in the details panel.](../uploads/Main/profiler-cpu-module.png)

Profiler window with a frame in the CPU Usage Profiler module selected. The Timeline view is selected in the details panel.

## Module types

The [CPU Usage module](ProfilerCPU.md) provides an overview of how much time your application spends on each frame. The other modules collect data which you can use to inspect specific areas or to monitor the vitals of your application, such as memory consumption, rendering, or audio statistics.

Each module has its own chart. When you select a module, the details panel in the bottom section of the Profiler window displays detailed data that the module collects. You can then use this data to identify areas of improvement in your application.

## Default Profiler modules

The Profiler window includes the following Profiler modules by default. To learn about other Profiler properties and settings, refer to [Profiler window reference](ProfilerWindow.md).

| **Module** | **Description** |
| --- | --- |
| **Highlights** | Displays information on whether your application is meeting its target frame rate and if its performance is bound by the CPU or the GPU. For more information, refer to the [Highlights Profiler module](ProfilerHighlights.md). |
| **CPU Usage** | Displays an overview of where your application spends the most time, in areas such as physics, **scripts**, animation, and garbage collection. This module contains broad profiling information about your application, and you can use it to decide which further modules to use to investigate more specific issues in your application. This module is always active, even if you close it. For more information, refer to [CPU Usage Profiler module](ProfilerCPU.md). |
| **GPU Usage** | Displays information related to graphics processing. By default this module isn’t active, because it has a high overhead. For more information, refer to [GPU Usage Profiler module](ProfilerGPU.md). |
| **Rendering** | Displays information on how Unity renders graphics in your application. For more information, refer to [Rendering Profiler module](ProfilerRendering.md). |
| **Memory** | Displays information on how Unity allocates memory in your application. This module is useful to investigate how scripting allocations lead to [garbage collection](performance-garbage-collector.md), or how your application’s asset memory usage trends over time. For more information, refer to [Memory Profiler module](ProfilerMemory.md). |
| **Audio** | Displays information related to the audio in your application, such as how much CPU usage the audio system requires, and how much memory Unity allocates to it. For more information, refer to [Audio Profiler module](ProfilerAudio.md). |
| **Video** | Displays information related to [video](VideoPlayer.md) in your application. For more information, refer to [Video Profiler module](profiler-video-profiler-module.md). |
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

## Custom Profiler modules

You can add your own Profiler modules to the Profiler window to capture and visualize specific performance data in your application. You can either use the [Profiler Module Editor](profiler-module-editor.md) or use scripts to automatically create and populate modules.

The following image of a customized Profiler window contains:

* **A**: A [custom Profiler module](profiler-creating-custom-modules.md) named **Tank Effects**.
* **B**: A [custom module details panel](profiler-customizing-details-view.md) that visualizes the data in the Tank Effects profiler module.
* **C**: [Custom Profiler counters](profiler-adding-information-code-intro.md) which track **particle** data.

![A custom Profiler module with custom game data displayed](../uploads/Main/Profiler_Tank_details.png)

A custom Profiler module with custom game data displayed

For more information, refer to [Customizing Profiler modules](profiler-customizing.md).

## Additional resources

* [Using the Profiler window](ProfilerWindow.md)
* [Profiling your application](profiler-profiling-applications.md)
* [Profiler window reference](ProfilerWindow.md)
* [Customizing Profiler modules](profiler-customizing.md)