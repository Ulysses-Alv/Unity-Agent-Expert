# Frame timing manager introduction

You can use the [`FrameTimingManager` class](../ScriptReference/FrameTimingManager.md) to capture high level timing data about individual frame performance in an application. You can then use this data to assess whether your application meets performance targets.

You can use the `FrameTimingManager` class in the following situations:

* You need to monitor performance at a frame-by-frame level.
* Your application uses the [Dynamic Resolution](DynamicResolution-landing.md) feature.
* Your application uses [Adaptive Performance](adaptive-performance/adaptive-performance.md).

Frame timings don’t replace data from the [Profiler](Profiler.md). After you record high level metrics of your application with `FrameTimingManager`, you can use the Profiler to then investigate specific details.

**Note:** `FrameTimingManager` might decrease GPU performance when it records data on OpenGL ES, so it can’t produce an accurate measurement of how your application performs.

To use the data you collect from `FrameTimingManager` refer to [Get frame timing data](frame-timing-manager-get-timing-data.md).

## FrameTimingManager data availability

`FrameTimingManager` provides results with a set delay of four frames. This is for the following reasons:

* CPU timing results aren’t immediately available at the end of each frame.
* Unity reads GPU timing results with three frames delay.

**Note:** The four frame delay doesn’t guarantee accurate timing results, because the GPU might not have any available resources to return the results, or might fail to return them correctly.

For a detailed description of available timings, refer to [Get frame timing data](frame-timing-manager-get-timing-data.md).

`FrameTimingManger` changes how it produces a `FrameTimeComplete` timestamp under the following circumstances:

* If the GPU supports GPU timestamps, the GPU provides a `FrameTimeComplete` timestamp.
* If the GPU doesn’t support GPU timestamps but returns a [GPU time](ProfilerHighlights.md#gpu-time), `FrameTimingManager` calculates a value for `gpuFrameTime`. The value is the sum of the reported GPU time and the `FirstSubmitTimestamp` values.
* If the GPU doesn’t support GPU timestamps and doesn’t return GPU time, `FrameTimingManager` sets the value of `PresentTimestamp` as the value of `FrameTimeComplete`.

## Platform support

`FrameTimingManager` is supported on the following platforms and graphics APIs:

| **Platform** | **Graphics** | **Supported** |
| --- | --- | --- |
| **Android** | OpenGL ES | Yes |
| **Android** | Vulkan | Yes |
| **iOS** | Metal | Yes |
| **Linux** | OpenGL | Partial: [GPU Frame Time](profiler-counters-reference.md#rendering) measurement is unsupported. |
| **Linux** | Vulkan | Yes |
| **macOS** | Metal | Yes |
| **tvOS** | Metal | Yes |
| ****WebGL**** | WebGL | Partial: [GPU Frame Time](profiler-counters-reference.md#rendering) measurement is unsupported. |
| **Windows** | DirectX 11 | Yes |
| **Windows** | DirectX 12 | Yes |
| **Windows** | OpenGL | Yes |
| **Windows** | Vulkan | Yes |
| ****XR**** | OpenGL ES | Partial: [CPU Render Thread Frame Time and GPU Frame Time](profiler-counters-reference.md#rendering) measurement is unsupported. |
| **XR** | Vulkan | Partial: [CPU Render Thread Frame Time and GPU Frame Time](profiler-counters-reference.md#rendering) measurement is unsupported. |

### Compatibility with tile-based deferred rendering GPUs

For GPUs that use tile-based deferred rendering architecture, such as Metal GPUs in Apple devices, the reported GPU time might be larger than the reported frame time.

This can happen when the GPU is under heavy load, or when the GPU pipeline is full. In these cases, the GPU might defer execution of some rendering phases. Because `FrameTimingManager` measures the time between the beginning and end of the frame rendering, any gaps between phases increase the reported GPU time.

In the following example, no GPU resources are available, because the GPU passes a job from the vertex queue to the fragment queue. The GPU’s graphics API therefore defers the execution of the next phase. When this happens, the GPU time measurement includes phase work time and any gap in between. The result is that `FrameTimingManager` reports a higher GPU time measurement than expected.

![Diagram showing how the discrepancy in reported GPU time can happen in the Metal API](../uploads/Main/frame-timing-manager-deferred-rendering-diagram.png)

Diagram showing how the discrepancy in reported GPU time can happen in the Metal API

## Additional resources

* [`FrameTimingManager` API](../ScriptReference/FrameTimingManager.md)
* [Profiler overview](Profiler.md)
* [Dynamic resolution](DynamicResolution-landing.md)
* [`FrameTiming` API](../ScriptReference/FrameTiming.md)