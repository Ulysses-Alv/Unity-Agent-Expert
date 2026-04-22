# Visualizing profiler counters

Once you’ve [added Profiler counters to your code](profiler-add-counters-code.md), you can visualize them in the Profiler window, or in a build of your application.

## Visualizing custom counters in the Profiler window

To view the data that `ProfilerCounter` or `ProfilerCounterValue` generates in the Profiler window, you can use a [custom Profiler module](profiler-customizing.md). This might help to visually recognize relationships with other system metrics and identify performance issues quickly.

### Adding metrics to your code

The following example assumes that the GameManager class handles high level logic and knows about enemies. To update the value of the counter, in the `Update` or `LateUpdate` method (depending on when the logic is performed with spawning or destroying enemies), use the `Sample` method to push the enemy’s count value to the Profiler:

```
using UnityEngine;
using Unity.Profiling;

class GameManager : MonoBehaviour
{
    Enemy[] m_Enemies;

    void Update()
    {
        GameStats.EnemyCount.Sample(m_Enemies.Length);
    }
}
```

To pass the bullet count to the Profiler, this example assumes that there’s a shell script that manages the bullet lifecycle. It then increases the `GameStats.BulletCount` value in `Awake` and decreases it in `OnDestroy` to give accurate information about the current bullet flow in the game.

```
using UnityEngine;
using Unity.Profiling;

public class Shell : MonoBehaviour
{
    void Awake()
    {
        GameStats.BulletCount.Value += 1;
    }
    void OnDestroy()
    {
        GameStats.BulletCount.Value -= 1;
    }
}
```

**Note:** Unity compiles out both `ProfilerCounter` and `ProfilerCounterValue` in non-development builds.

### Add a custom module to the Profiler window

Use the [Profiler Module Editor](profiler-module-editor.md) to select built-in or newly added counters for the visualization. To open the Profiler Module Editor:

1. Open the Profiler Window (**Window > Analysis > Profiler**).
2. Select the Profiler Module dropdown in the top left of the window. Click the gear icon, and the Profiler Module Editor opens in a new window.
3. Select **Add** and then choose the counters you want to visualize from the **User** list on the right hand side of the window.

**Important:** If you don’t have any data loaded into the Profiler window, then any counters you’ve created don’t appear in the Available Counters pane when you load the Profiler Module Editor. To view custom counters, [capture or load some data](profiler-collect-data) that has your emitted counters in with the Profiler, and reopen the Profiler Module Editor.

You can then view the data in the Profiler window alongside other counters. For more information, refer to [Creating custom Profiler modules](profiler-creating-custom-modules.md).

![Module with custom counters in Profiler Window.](../uploads/Main/profiler-window-custom-module.png)

Module with custom counters in Profiler Window.

**Note:** Counters declared as static are dynamically initialized in the C# code when a type is initialized and might not be available until they’re actually initialized and used. This applies to both Edit and Play modes. If your counters don’t appear in the Profiler Module Editor, [record some data with the Profiler](profiling-target-device.md) first for a few frames.

## Visualizing Profiler counters in a built player

When you run your project in a built player, you don’t have access to the Profiler window. However, you can display counters as UI elements in a built player. This means you can include profiling tools in a build of your application.

The following image displays counters in the top left of the **scene** using custom UI in a built player:

![Tanks game with custom information overlaid](../uploads/Main/Tanks-Runtime-Stats.png)

Tanks game with custom information overlaid

**Note:** Not all **Profiler counters** are available in Release Players (non-development builds). Use [`ProfilerRecorder.Valid`](../ScriptReference/Unity.Profiling.ProfilerRecorder.Valid.md) to determine if the data is available and that the Profiler can record it. Alternatively, you can use [`ProfilerRecorderHandle.GetAvailable`](../ScriptReference/Unity.Profiling.LowLevel.Unsafe.ProfilerRecorderHandle.GetAvailable.md) to enumerate all available Profiler statistics.

### Getting counter values in players

Profiler counters give you an insight into important game or engine system metrics. If you have a continuous integration setup or want to visualize key performance metrics in your application during a test play through you can use the [`ProfilerRecorder`](../ScriptReference/Unity.Profiling.ProfilerRecorder.md) API to get custom and built-in Profiler counter values.

For example, the following script displays the frame time, Mono/**IL2CPP** heap size, and total memory that the application uses.

```
using System.Collections.Generic;
using System.Text;
using Unity.Profiling;
using UnityEngine;

public class StatsScript : MonoBehaviour
{
    string statsText;
    ProfilerRecorder systemMemoryRecorder;
    ProfilerRecorder gcMemoryRecorder;
    ProfilerRecorder mainThreadTimeRecorder;

    double GetRecorderFrameAverage(ProfilerRecorder recorder)
    {
        var samplesCount = recorder.Capacity;
        if (samplesCount == 0)
            return 0;

        double r = 0;
        unsafe
        {
            var samples = stackalloc ProfilerRecorderSample[samplesCount];
            recorder.CopyTo(samples, samplesCount);
            for (var i = 0; i < samplesCount; ++i)
                r += samples[i].Value;
            r /= samplesCount;
        }

        return r;
    }

    void OnEnable()
    {
        systemMemoryRecorder = ProfilerRecorder.StartNew(ProfilerCategory.Memory, "System Used Memory");
        gcMemoryRecorder = ProfilerRecorder.StartNew(ProfilerCategory.Memory, "GC Reserved Memory");
        mainThreadTimeRecorder = ProfilerRecorder.StartNew(ProfilerCategory.Internal, "Main Thread", 15);
    }

    void OnDisable()
    {
        systemMemoryRecorder.Dispose();
        gcMemoryRecorder.Dispose();
        mainThreadTimeRecorder.Dispose();
    }

    void Update()
    {
        var sb = new StringBuilder(500);
        sb.AppendLine($"Frame Time: {GetRecorderFrameAverage(mainThreadTimeRecorder) * (1e-6f):F1} ms");
        sb.AppendLine($"GC Memory: {gcMemoryRecorder.LastValue / (1024 * 1024)} MB");
        sb.AppendLine($"System Memory: {systemMemoryRecorder.LastValue / (1024 * 1024)} MB");
        statsText = sb.ToString();
    }

    void OnGUI()
    {
        GUI.TextArea(new Rect(10, 30, 250, 50), statsText);
    }
}
```

**Important:** Use `ProfilerRecorder.Dispose` to free unmanaged resources associated with the `ProfilerRecorder`.

## Getting counter values from Profiler frame data in the Editor

To get Profiler counter values when processing Profiler frame data in the Editor, use the [`FrameDataView`](../ScriptReference/Profiling.FrameDataView.md) API.

You can use the [`FrameDataView.GetCounterValueAsInt`](../ScriptReference/Profiling.FrameDataView.GetCounterValueAsInt.md), [`FrameDataView.GetCounterValueAsLong`](../ScriptReference/Profiling.FrameDataView.GetCounterValueAsLong.md), [`FrameDataView.GetCounterValueAsFloat`](../ScriptReference/Profiling.FrameDataView.GetCounterValueAsFloat.md) and [`FrameDataView.GetCounterValueAsDouble`](../ScriptReference/Profiling.FrameDataView.GetCounterValueAsDouble.md) to get a frame value of the specific counter, like so:

```
using UnityEditor.Profiling;

class Example
{
    static int ExtractMyCounterValue(FrameDataView frameData, string counterName)
    {
        var counterMarkerId = frameData.GetMarkerId(counterName);
        return frameData.GetCounterValueAsInt(counterMarkerId);
    }
}
```

## Additional resources

* [Adding profiler counters to your code](profiler-add-counters-code.md)
* [Profiler Module Editor reference](profiler-module-editor.md)