# Track timing and thermal data

Retrieve real-time device information using the Adaptive Performance API, including frame timing, performance bottlenecks, CPU core configuration, and thermal warnings.

## Track frame timing

Adaptive Performance tracks the average GPU, CPU, and overall frame times, and updates them every frame.

To access the latest timing data, use the [`IPerformanceStatus.FrameTiming`](../../ScriptReference/AdaptivePerformance.IPerformanceStatus.FrameTiming.md) property, which contains the following information:

* **Overall frame time**: The time difference between frames. Use this value to calculate the current frame rate of the application.
* **CPU time**: The time the CPU is actually executing Unity’s main thread and the render thread. CPU time doesn’t include the times when the operating system might be blocking Unity, or when Unity needs to wait for the GPU to catch up with rendering.
* **GPU time**: The time the GPU is actively processing data to render a frame. GPU time doesn’t include the time when the GPU has to wait for Unity to provide data to render.

## Identify performance bottlenecks

Adaptive Performance uses the target frame rate configured in [`Application.targetFrameRate`](../../ScriptReference/Application-targetFrameRate.md) and [`QualitySettings`](../../ScriptReference/QualitySettings.md), along with the information that [`FrameTiming`](../../ScriptReference/FrameTiming.md) provides, to calculate what is limiting the application’s frame rate. If the application isn’t performing at the desired target frame rate, it might be bound by either CPU or GPU processing.

To get a notification whenever the current performance bottleneck of the application changes, subscribe with a delegate function to the [`IPerformanceStatus.PerformanceBottleneckChangeEvent`](../../ScriptReference/AdaptivePerformance.IPerformanceStatus.PerformanceBottleneckChangeEvent.md) event.

Use the information about the current performance bottleneck to make targeted adjustments to the game content at runtime. For example, in a GPU-bound application, lowering the rendering resolution often improves the frame rate significantly, but the same change might not make a big difference for a CPU-bound application.

## Get cluster information

CPU cores come in different sizes and configurations on heterogeneous systems. In certain scenarios, it’s important to know about the cluster in more detail.

Use the [`clusterInfo`](../../ScriptReference/AdaptivePerformance.ClusterInfo.md) struct to get information about the count of cores. For example:

```
  var clusterInfo = Holder.Instance.PerformanceStatus.PerformanceMetrics.ClusterInfo;
  Debug.Log($"Cluster Info = Big Cores: {clusterInfo.BigCore} Medium Cores: {clusterInfo.MediumCore} Little Cores: {clusterInfo.LittleCore}");
```

A cluster is unlikely to change during runtime, so you can cache the data.

For a practical example of using `clusterInfo`, refer to the [Adaptive Performance samples](https://docs.unity3d.com/Packages/com.unity.adaptiveperformance@latest?subfolder=/manual/samples-guide.html#cluster-info-sample).

## Monitor performance mode

Adaptive Performance tracks the user-selected performance mode for the application if available.

To access the current performance mode, use the [`IPerformanceModeStatus.PerformanceMode`](../../ScriptReference/AdaptivePerformance.IPerformanceModeStatus.PerformanceMode.md) property.

To get a notification whenever the current performance mode of the application changes, subscribe with a delegate function to the [`IPerformanceModeStatus.PerformanceModeEvent`](../../ScriptReference/AdaptivePerformance.IPerformanceModeStatus.PerformanceModeEvent.md) event. For example:

```
using UnityEngine;
using UnityEngine.AdaptivePerformance;
public class PerformanceModeAttacher : MonoBehaviour
{
    void Start()
    {
        Holder.Instance.PerformanceModeStatus.PerformanceModeEvent+= (PerformanceMode performanceMode) =>
        {
            Debug.Log("Performance Mode Changed: " + performanceMode);
        };
    }
}
```

## Track thermal and power states

The Adaptive Performance API gives you access to the current thermal warning level of the device via [`ThermalMetrics.WarningLevel`](../../ScriptReference/AdaptivePerformance.ThermalMetrics.WarningLevel.md) and a more detailed temperature level via [`ThermalMetrics.TemperatureLevel`](../../ScriptReference/AdaptivePerformance.ThermalMetrics.TemperatureLevel.md). The application can make modifications based on these values to stop the operating system from throttling.

The following example shows the implementation of a Unity component that uses Adaptive Performance feedback to adjust the global **LOD** bias:

```
using UnityEngine;
using UnityEngine.AdaptivePerformance;

public class AdaptiveLOD : MonoBehaviour
{
    // Declare 'ap' as a private class-level field to hold the IAdaptivePerformance instance.
    private IAdaptivePerformance ap = null;

    void Start() {
        // Get the Adaptive Performance instance.
        ap = Holder.Instance;
        if (!ap.Active)
            return;

        QualitySettings.lodBias = 1.0f;
        ap.ThermalStatus.ThermalEvent += OnThermalEvent;
    }

    void OnThermalEvent(ThermalMetrics ev) {
        // Adjust LOD bias based on the current thermal warning level.
        switch (ev.WarningLevel) {
            case WarningLevel.NoWarning:
                QualitySettings.lodBias = 1;
                break;
            case WarningLevel.ThrottlingImminent:
                if (ev.temperatureLevel > 0.8f)
                    QualitySettings.lodBias = 0.75f;
                else
                    QualitySettings.lodBias = 1.0f;
                break;
            case WarningLevel.Throttling:
                QualitySettings.lodBias = 0.5f;
                break;
        }
    }
}
```

## Additional resources

* [CPU and GPU performance control](cpu-gpu-performance-control.md)
* [Temporary performance boosts](temporary-performance-boosts.md)
* [Adaptive Performance Profiler integration](profiler-integration.md)
* [Scalers](scalers.md)