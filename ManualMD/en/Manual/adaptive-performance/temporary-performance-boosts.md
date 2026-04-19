# Temporary performance boosts

Request a temporary CPU or GPU boost to increase device performance for short periods.

Temporary performance boosts in Adaptive Performance allow you to raise CPU or GPU frequencies for a limited time. This can help your application complete demanding tasks more quickly and avoid visible performance drops. Boosts increases energy use and heat, so use them sparingly and only when necessary.

To use performance boosts, refer to [Request a temporary performance boost](request-performance-boost.md). For a practical example of using temporary performance boosts, refer to the [Adaptive Performance samples](https://docs.unity3d.com/Packages/com.unity.adaptiveperformance@latest?subfolder=/manual/samples-guide.html#boost-sample).

## How performance boosts work

A performance boost temporarily raises both the minimum and maximum frequencies of the CPU or GPU to highest configured setting, providing more resources and enabling the CPU or GPU to do more work.

The downside is a higher energy consumption and heat production. Using a boost can bring the device into a throttled state. Profile your application and look for additional optimizations before considering using a boost. Only use performance boosts for short bursts.

## When to use performance boosts

You might use performance boosts to do the following:

* Decrease load times if a task is CPU- or GPU-bound
* Avoid visible hitches when you expect a spike in CPU or GPU work (such as network processing, compiling shaders, or generating content)
* Preemptively ramp up performance before high-load moments like:
  + Loading a scene
  + Switching **scene** content (for example, when entering a boss fight)
  + Spawning many objects
  + Showing advertisements
  + During bursts of network traffic

**Note**: When using performance boosts, make sure [`IDevicePerformanceControl.PerformanceControlMode`](../../ScriptReference/AdaptivePerformance.IDevicePerformanceControl.PerformanceControlMode.md) is set to `PerformanceControlMode.System`. Refer to [CPU and GPU performance control](cpu-gpu-performance-control.md) for more information.

## Example of requesting a CPU performance boost

In the following sample case, the big and medium cores of the CPU are overloaded. Requesting a CPU boost provides additional resources, so the program can execute everything in time.

![A graph from the Unity Profiler showing CPU core frequencies and loads before and after a temporary performance boost. Big and medium cores accelerate while tiny cores are on the same level as before.](../../uploads/Main/ap/boost-mode.png)

A graph from the Unity Profiler showing CPU core frequencies and loads before and after a temporary performance boost. Big and medium cores accelerate while tiny cores are on the same level as before.

## Additional resources

* [Request a temporary performance boost](request-performance-boost.md)
* [Apply performance optimizations](performance-optimization-strategies.md)
* [Provider settings reference](provider-settings-reference.md)
* [Device Simulator reference](device-simulator-reference.md)