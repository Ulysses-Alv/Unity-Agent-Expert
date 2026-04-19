# Request a temporary performance boost

Request a temporary CPU or GPU boost to handle short, performance-intensive tasks.

Before using a boost, make sure you understand the performance trade-offs. Refer to [Temporary performance boosts](temporary-performance-boosts.md) for background information.

## Request a boost at runtime

**Important**: Boosting the CPU and GPU simultaneously requires a lot of power. Avoid boosting both if possible. You might overtax the device, which can lead to unintended behaviors.

To request a boost at runtime, set [`CpuPerformanceBoost`](../../ScriptReference/AdaptivePerformance.IDevicePerformanceControl.CpuPerformanceBoost.md) or [`GpuPerformanceBoost`](../../ScriptReference/AdaptivePerformance.IDevicePerformanceControl.GpuPerformanceBoost.md) to `true`. For example:

```
using UnityEngine;
using UnityEngine.AdaptivePerformance;

public class PerformanceBoostRequester : MonoBehaviour
{
    // Declare 'ap' as a private class-level field to hold the IAdaptivePerformance instance.
    private IAdaptivePerformance ap;

    void Start()
    {
        // Get the Adaptive Performance instance.
        ap = Holder.Instance;
    }
    public void BoostCPU()
{
    // Ensure Adaptive Performance is active before requesting a boost.
    if (ap == null || !ap.Active)
        return;
    // Request a CPU performance boost.
    ap.DevicePerformanceControl.CpuPerformanceBoost = true;
}
}
```

If the CPU or GPU doesn’t require additional resources, or if it’s idle, then requesting a boost doesn’t have an effect.

For a practical example of using temporary performance boosts, refer to the [Adaptive Performance samples](https://docs.unity3d.com/Packages/com.unity.adaptiveperformance@latest?subfolder=/manual/samples-guide.html#boost-sample).

## Disable boosts on startup (optional)

By default, Unity enables boosts for the CPU and the GPU on startup because fast startup time is a common requirement.

To disable boosts on startup:

1. Go to **Edit** > **Project Settings** > **Adaptive Performance**.
2. In the **Runtime Settings** section, disable **Boost mode on startup**.

## Additional resources

* [Temporary performance boosts](temporary-performance-boosts.md)
* [CPU and GPU performance control](cpu-gpu-performance-control.md)
* [Apply performance optimizations](performance-optimization-strategies.md)