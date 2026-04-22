# Read frame timing data with ProfilerRecorder

You can use the [`ProfilerRecorder`](../ScriptReference/Unity.Profiling.ProfilerRecorder.md) API to get custom and built-in **Profiler counter** values, and to read `FrameTimingManager` values.

The benefit of this is that when you use the `ProfilerRecorder` API, `FrameTimingManager` only records values when you attach a recorder to a specific counter. This behavior enables you to control which counters collect data and reduces the impact that the `FrameTimingManager` has on performance.

The following example tracks only the CPU Main Thread Frame Time variable with the `ProfilerRecorder` API:

```
using Unity.Profiling;
using UnityEngine;

public class ExampleScript : MonoBehaviour
{
    ProfilerRecorder mainThreadTimeRecorder;

    void OnEnable()
    {
        mainThreadTimeRecorder = ProfilerRecorder.StartNew(ProfilerCategory.Internal, "CPU Main Thread Frame Time");
    }

    void OnDisable()
    {
        mainThreadTimeRecorder.Dispose();
    }

    void Update()
    {
        var frameTime = mainThreadTimeRecorder.LastValue;

        // Your code logic here
    }
}
```

## Additional resources

* [Get frame timing data](frame-timing-manager-get-timing-data.md)
* [Adding profiling information to your code](profiler-adding-information-code.md)