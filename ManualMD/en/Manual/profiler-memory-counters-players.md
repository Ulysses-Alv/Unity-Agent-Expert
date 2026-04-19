# Accessing memory counters in players

You can use the [ProfilerRecorder API](../ScriptReference/Unity.Profiling.ProfilerRecorder.md) to access the Memory Profiler module’s counters in a player.

The following example contains a simple script that collects **Total Reserved Memory**, **GC Reserved Memory** and **System Used Memory** metrics, and displays those as a [`GUI.TextArea`](../ScriptReference/GUI.TextArea.md). The Memory Profiler module information belongs to the [`ProfilerCategory.Memory`](../ScriptReference/Unity.Profiling.ProfilerCategory.Memory.md) **Profiler category**.

```
using System.Text;
using Unity.Profiling;
using UnityEngine;

public class MemoryStatsScript : MonoBehaviour
{
    string statsText;
    ProfilerRecorder totalReservedMemoryRecorder;
    ProfilerRecorder gcReservedMemoryRecorder;
    ProfilerRecorder systemUsedMemoryRecorder;

    void OnEnable()
    {
        totalReservedMemoryRecorder = ProfilerRecorder.StartNew(ProfilerCategory.Memory, "Total Reserved Memory");
        gcReservedMemoryRecorder = ProfilerRecorder.StartNew(ProfilerCategory.Memory, "GC Reserved Memory");
        systemUsedMemoryRecorder = ProfilerRecorder.StartNew(ProfilerCategory.Memory, "System Used Memory");
    }

    void OnDisable()
    {
        totalReservedMemoryRecorder.Dispose();
        gcReservedMemoryRecorder.Dispose();
        systemUsedMemoryRecorder.Dispose();
    }

    void Update()
    {
        var sb = new StringBuilder(500);
        if (totalReservedMemoryRecorder.Valid)
            sb.AppendLine($"Total Reserved Memory: {totalReservedMemoryRecorder.LastValue}");
        if (gcReservedMemoryRecorder.Valid)
            sb.AppendLine($"GC Reserved Memory: {gcReservedMemoryRecorder.LastValue}");
        if (systemUsedMemoryRecorder.Valid)
            sb.AppendLine($"System Used Memory: {systemUsedMemoryRecorder.LastValue}");
        statsText = sb.ToString();
    }

    void OnGUI()
    {
        GUI.TextArea(new Rect(10, 30, 250, 50), statsText);
    }
}
```

The following screenshot shows the result of adding the script to the [Tanks! tutorial project](https://assetstore.unity.com/packages/essentials/tutorial-projects/tanks-tutorial-46209).

![Tanks! tutorial with the overlaid memory information](../uploads/Main/profiler-tanks-memory-overlay.png)

Tanks! tutorial with the overlaid memory information

This information is available in release players, as are the other [high level counters available in the Memory Profiler module](ProfilerMemory.md). If you want to view the selected memory counters in a custom module in the Profiler window, use the [Profiler module editor](profiler-module-editor.md) to configure the chart.

For more information on adding profiler information to your code, refer to [Adding profiling information to your code](profiler-adding-information-code.md).

## Additional resources

* [Profiler counters reference](profiler-counters-reference.md)
* [Adding profiling information to your code](profiler-adding-information-code.md)
* [Visualizing profiler counters](profiler-creating-custom-counters.md)
* [Memory Profiler module reference](ProfilerMemory.md)