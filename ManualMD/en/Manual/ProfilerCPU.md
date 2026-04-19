# CPU Usage Profiler module reference

The CPU Usage Profiler module is divided into charts in the top half of the [Profiler window](ProfilerWindow.md), and has timeline and hierarchy views in the details pane in the bottom half of the window.

![Profiler window with a frame in the CPU Usage Profiler module selected. The Timeline view is selected in the details pane.](../uploads/Main/profiler-cpu-module.png)

Profiler window with a frame in the CPU Usage Profiler module selected. The Timeline view is selected in the details pane.

## Chart categories

The CPU Usage Profiler module’s chart tracks the time spent on the application’s main thread.

| **Chart** | **Description** |
| --- | --- |
| **Rendering** | How much time your application spends on rendering graphics. |
| ****Scripts**** | How much time your application spends on running scripts. |
| **Physics** | How much time your application spends on the physics system. |
| **Animation** | How much time your application spends on animating skinned **mesh** Renderers, **GameObjects**, and other components in your application. This also includes the time spent on calculations for systems the Animation and **Animator components** use. |
| **GarbageCollector** | How much time your application spends running the [managed garbage collector](performance-garbage-collector.md). |
| ****VSync**** | How much time your application spends per frame waiting for the `targetFrameRate` or the next `VBlank` to sync with. To determine the VSync value, Unity uses the [`QualitySettings.vSyncCount`](../ScriptReference/QualitySettings-vSyncCount.md) value, the target frame rate, or the default or enforced maximum VSync value of the platform your application is running on. For more information about VSync, refer to the [Profiler markers documentation](profiler-markers.md). |
| ****Global Illumination**** | How much time your application spends on lighting. |
| **UI** | How much time your application spends on displaying its UI. |
| **Others** | How much time your application spends on code that doesn’t fall into any of the other categories. This includes areas like the entire EditorLoop, or the profiling overhead when you [profile in Play mode](profiling-play-mode.md). For more information refer to [Play mode and Editor samples](profiler-play-edit-samples.md). |

## Module details pane toolbar

When you select the CPU Usage module, the details pane displays a breakdown of where the application spent time in the selected frame. Use the **toolbar** at the top of the pane to navigate the views.

| **Property** | **Description** |
| --- | --- |
| **Timing data view dropdown** | Choose to display the timing data as either a timeline or a hierarchical table. To change the display, use the dropdown in the details pane (set to **Timeline** by default). The following views are available:   * **Timeline**: Displays a breakdown of the timings for a particular frame alongside a time axis of the frame’s length.  This is the only view that you can use to view timings on all threads at once and at the times within the frame at which they happened, so that you can correlate timings across threads. For example, [job system](job-system.md) worker threads starting up after a system on the main thread schedules them. For more information, refer to [Timeline view](profiler-cpu-introduction.md#timeline-view). * **Hierarchy**: Groups the timing data by its internal hierarchical structure. This view displays the elements that your application called in a descending list format, ordered by the time spent by default. You can also order the information by the amount of scripting memory allocated, or the number of calls. * **Raw Hierarchy**: Displays the timing data in a hierarchical structure that’s similar to the call stacks where the timing happened. Unity lists each call stack separately in this mode instead of merging them, as it does in Hierarchy view. * **Inverted Hierarchy**: Groups samples by [profiler marker](profiler-markers.md) and displays them with inverted sample stacks. The first level of the hierarchy shows an item for each profiler marker.  Expand an item in the tree to display the markers that contain this one in their sample stack. This option helps to reveal larger performance issues caused by lots of instances of small performance impacts. These kinds of issues can be harder to spot in the Timeline or non-inverted hierarchy views. |
| **Live setting** | Displays data about the selected frame in real time in the module details pane when you record new data in Play mode or the Editor. By default, this setting is disabled, and the module details pane is blank when you record data.   **Note:** This setting increases the overhead of the `EditorLoop` when the Profiler window is repainted. |
| **Thread selector** (Hierarchy views only) | Use the Thread dropdown to select a specific thread, like the Main Thread or Render Thread to inspect in these views. |
| **Details dropdown** (Hierarchy views only) | Display information about where your application calls and uses profiled functions. Choose from the following options:   * **No Details**: Doesn’t display any extra information * **Related data**: Displays a list of `UnityEngine.Objects` that use a [`Begin`](../ScriptReference/Unity.Profiling.ProfilerMarker.Begin.md) overload which are associated with the Profiler sample. * **Calls**: Displays where the selected sample is called from and what other functions it calls to. |

## More (⋮) menu settings

The More (⋮) menu contains the following settings:

| **Setting** | **Description** |
| --- | --- |
| **Show Full Scripting Method Names** | Displays the fully qualified names for all scripting methods (`Assembly::Class::MethodName`). |
| **Show Flow events (Timeline view only)** | Displays the relationship between systems, jobs, and threads. When you enable this setting, the Profiler adds white event markers to the Profiler samples that schedule jobs, or wait on scheduled jobs to complete. It also darkens unrelated samples to make it easier to visualize the sample you select.  When you select a sample, the Profiler connects the relevant flow event markers together with lines. A thicker line highlights the particular flow line you select. For example, if a `begin` sample points to two other `next` samples, when you click one of the `next` samples, the Profiler draws a thicker line to it.  The Profiler adds the following arrow types to the samples:   * **Down arrow (▿)**: Indicates the beginning of a flow, and that this sample scheduled some work. * **Right arrow (▹)**: Indicates the next item in a flow, and that a different sample scheduled this. * **Up arrow (▵)**: Indicates the end of a flow, and that the work ended or synchronized on this sample. |
| **Collapse EditorOnly samples (Hierarchy views only)** | Displays `EditorOnly` samples in the view. By default, Unity collapses all `EditorOnly` samples in Hierarchy views. EditorOnly samples are samples in the player loop that only happen because of Editor-only safety checks. When the samples are collapsed, their `GC.Alloc` value doesn’t contribute to the `GC.Alloc` value of their enclosing sample. For more information, refer to [Play mode and Editor samples](profiler-play-edit-samples.md). |

## Timeline tooltips

When you select a sample in the Timeline view, the tooltip on the selected sample provides further details, such as the number of instances and the total time of this sample across all threads. You can select the text within the tooltip and copy it and use the buttons to interact with the sample further.

| **Tooltip** | **Description** |
| --- | --- |
| **Copy** | Copy the call stack and the entire contents of the tooltip to your clipboard. |
| **Show** | Select this dropdown to view the sample in different contexts:   * **Hierarchy**: Change the selected sample’s view to [Hierarchy](profiler-cpu-introduction.md#views) view. * **Raw Hierarchy**: Change the selected sample’s view [Raw Hierarchy](profiler-cpu-introduction.md#views) view. * **Inverted Hierarchy**: Change the selected sample’s view [Inverted Hierarchy](profiler-cpu-introduction.md#views) view. * **Full Details for Call Stacks**: Enable this property to display the full list of method pointer addresses of the call stack.  Unity records call stacks as a list of method pointer addresses, which it uses to display the method name, file path, and line number of the stack. Whenever only the pointer address is present, Unity ignores it to conserve screen space for the actionable items that have further information available. * **Selected Sample Stack**: View the details of the sample stack. Unity opens this information in a separate window. You can then copy the sample stack information to your clipboard.  The sample stack differs from a method’s call stack because Unity doesn’t tie every sample to a specific method, nor does it record every call as a sample. If you select a sample in a different frame and there isn’t a sample with the same sample stack in the displayed frame, this window shows both the sample stack of the original selection, and the approximate selection for this frame. |

## Hierarchy columns

The hierarchy views display the following detailed information for each item:

| **Column** | **Description** |
| --- | --- |
| **Total** | The total amount of time Unity spent in a particular sample, as a percentage of the total frame time. |
| **Self** | The total amount of time Unity spent in a particular sample as a percentage of the total frame time, excluding the time from sub samples. |
| **Calls** | The number of calls made to this sample in this frame. In the Raw Hierarchy view the values in this column are always `1` because the Profiler doesn’t merge the hierarchy of samples. |
| **GC Alloc** | How much scripting heap memory Unity has allocated in the current frame. The garbage collector manages the scripting heap memory.  For more details about the managed heap refer to the documentation on [Understanding Automatic Memory Management](performance-managed-memory.md). |
| **Time ms** | The total amount of time Unity spent in a particular sample, in milliseconds.  **Note:** If your application uses the [job system](job-system.md) or multithreaded rendering, this information might be misleading, because it only contains the time Unity spent on the selected thread. To change the thread, select the **Thread** dropdown at the top of the Hierarchy pane. |
| **Self ms** | The total amount of time Unity spent in a particular sample, in milliseconds, excluding the time Unity spends calling sub functions. |
| **Warning (⚠)** | Displays how many times the application has triggered a warning during the current frame. For more information, refer to [Performance warnings](profiler-markers.md). |

## Additonal resources

* [Profiler window reference](ProfilerWindow.md)
* [Understanding Automatic Memory Management](performance-managed-memory.md)
* [Play mode and Editor samples](profiler-play-edit-samples.md)
* [Collect performance data in Play mode](profiling-play-mode.md)