# Play mode and Editor profile samples

When you profile the Unity Editor or Play mode, Unity groups their relevant timings into two sample groups that have the [following markers](profiler-markers.md):

* `PlayerLoop`: Timings related to Play mode
* `EditorLoop`: Timings related to the Editor

## Play mode samples

When you profile Play mode, the Profiler only collects timing samples that happened inside the `PlayerLoop`. This group helps to reduce the amount of misleading measurements in the [CPU](ProfilerCPU.md) and [GPU Profiler modules](ProfilerGPU.md).

Unity categorizes any `EditorLoop` samples as **Others** in the [CPU Profiler module charts](ProfilerCPU.md), meaning that `EditorLoop` samples are the biggest contributors to the **Others** category. If you want to find out what the Editor does in this time and get a detailed breakdown of what contributes to the **Others** category, [profile the Editor process](profiling-edit-mode.md).

**Important:** If you use [Deep Profiling](profiler-deep-profiling.md), and target Play mode, it has a performance impact on every call that happens in both `PlayerLoop` and `EditorLoop`. This is because Deep Profiling hooks into the beginning and end of any scripting method call on domain reload and it doesn’t detect which parts are never called from the `PlayerLoop`. The method calls that happen in the `EditorLoop` don’t have the full overhead of creating a sample, but they still check if they emit one, which causes a smaller overhead.

## Editor samples

When you [profile the Editor process](profiling-edit-mode.md), all the samples that were previously hidden under the `EditorLoop` marker contribute to their respective categories. This means that the information in the [CPU Profiler module’s detail pane](ProfilerCPU.md) and its charts changes significantly.

There are certain [profiler markers](profiler-markers.md) that only appear when you [profile in the Unity Editor](profiling-edit-mode.md). These markers don’t appear in Player-related activity, and only relate to Editor activity. Editor-only markers include:

* Security checks like the GetComponentNullErrorWrapper, which helps to identify a null component usage.
* CheckConsistency, which validates object setup.
* CheckAllowDestructionRecursive, which is a destruction check.
* Prefab-related activities.

By default, in the [CPU Profiler module’s **Hierarchy** view](ProfilerCPU.md), sample stacks with Editor-only markers are collapsed and are named `EditorOnly [SampleName]`. While these sample stacks or their child samples might cause managed allocations that can trigger garbage collection, they don’t contribute to the `GC.Alloc` value of their parent sample if they’re collapsed.

To change the default behavior, in the top right of the module details pane, select the context menu and disable the [Collapse EditorOnly Samples](ProfilerCPU.md) setting. When you do this, you can expand the sample and contribute its `GC.Alloc` value to the enclosing marker.

This option doesn’t affect the **Timeline** view, and the samples and their children appear expanded. You can usually ignore the samples with these markers, because they relate to Editor-only activity and don’t have any impact on reducing managed allocations. However, they can be useful to investigate if they have a significant performance impact on Play mode.

## Additional resources

* [Collect performance data in Play mode](profiling-play-mode.md)
* [Collect performance data about the Unity Editor](profiling-edit-mode.md)
* [CPU Usage Profiler module](ProfilerCPU.md)
* [Profiler markers reference](profiler-markers.md)