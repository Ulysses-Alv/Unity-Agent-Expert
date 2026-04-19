# Frame timing API counter reference

The following table describes the purpose of the counters that [the `FrameTiming` API](../ScriptReference/FrameTiming.md) provides.

| **FrameTiming API field** | ****Profiler counter**** | **Description** |
| --- | --- | --- |
| [`cpuFrameTime`](../ScriptReference/FrameTiming-cpuFrameTime.md) | **CPU Total Frame Time** | The total CPU frame time, in nanoseconds. Unity calculates this as the time between the ends of two frames, including any overheads or time spent waiting between frames. |
| [`cpuMainThreadFrameTime`](../ScriptReference/FrameTiming-cpuMainThreadFrameTime.md) | **CPU Main Thread Frame Time** | Active main thread work time in nanoseconds. This is the **PlayerLoop** time without the time the main thread was waiting for the VSync or render thread to finish its work, calculated as follows:   `cpuMainThreadFrameTime = PlayerLoop - GfxDevice.WaitForRenderThread - Gfx.WaitForPresentOnGfxThread - WaitForTargetFPS` |
| [`cpuMainThreadPresentWaitTime`](../ScriptReference/FrameTiming-cpuMainThreadPresentWaitTime.md) | N/A | The CPU time in nanoseconds spent waiting for the frame presentation or VSync during the frame, calculated as follows:   `cpuMainThreadPresentWaitTime = Gfx.WaitForPresentOnGfxThread + WaitForTargetFPS`. |
| [`cpuRenderThreadFrameTime`](../ScriptReference/FrameTiming-cpuRenderThreadFrameTime.md) | **CPU Render Thread Frame Time** | Active Render Thread work time in nanoseconds without Unity calling the `Present()` function, calculated as follows:   `cpuRenderThreadFrameTime = RenderLoop - Gfx.PresentFrame`. |
| [`gpuFrameTime`](../ScriptReference/FrameTiming-gpuFrameTime.md) | **GPU Frame Time** | The time difference in nanoseconds between the beginning and the end of the GPU rendering a single frame. |

The following diagram visualizes the relationships between **FrameTiming** and [the Profiler](Profiler.md) frame data.

![Diagram showing relation of FrameTiming API to Profiler data](../uploads/Main/frame-timing-manager-times.drawio.svg)

Diagram showing relation of FrameTiming API to Profiler data

## Additional resources

* [Get frame timing data](frame-timing-manager-get-timing-data.md)
* [Record frame timing data](frame-timing-manager-record-timing-data.md)