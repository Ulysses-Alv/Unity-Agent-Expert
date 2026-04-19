# Video Profiler module reference

The Video Profiler module displays information about what resources the video in your application is using, such as memory, buffering, and number of [video clips](class-VideoClip.md). You can use this data to determine how efficiently your application plays back and buffers videos on your selected platforms. You can also use the CPU Usage Profiler module to assess where Unity spends time on video. For more information, refer to [CPU Usage Profiler module reference](ProfilerCPU.md).

To open the Profiler window, go to menu: **Window** > **Analysis** > **Profiler**.

![The Video Profiler module](../uploads/Main/video-profiler-module.png)

The Video Profiler module

## Chart categories

The Video Profiler module’s chart contains the following categories:

| **Chart** | **Description** |
| --- | --- |
| **Total Video Sources** | The total number of video sources in your **Scene**. |
| **Playing Video Sources** | The number of video sources playing in your Scene. |
| **Pre-buffered frames** | The total number of pre-buffered frames. |
| **Total Video Memory** | The amount of system memory the video in your application is using. |

To change the order of the categories in the chart, you can drag them in the chart’s legend. You can also click a category’s colored legend to toggle its display. For more information on how to use the Profiler window, refer to [Getting started with the Profiler window](ProfilerWindow.md).

## Module details pane

When you select a frame in the Video Profiler module, the module details pane displays further detailed information on the video playback in your scene. The following information is available:

| **Detail** | **Description** |
| --- | --- |
| **Total Video Sources** | The number of video sources in your scene. |
| **Playing Video Sources** | The number of video sources playing in your Scene. |
| **Paused Video Sources** | The number of video sources that are paused. |
| **Software Video Playback** | The number of videos playing that the platform doesn’t natively support. |
| **Pre-buffered frames** | The total number of pre-buffered frames. |
| **Pre-buffered frame limit** | The pre-buffered frame limit. Unity buffers up to 16 frames per clip. |
| **Total frames dropped** | Number of frames that Unity had to skip to maintain real time. This might happen when your application runs slowly and can’t produce frames fast enough to play in real time. |
| **Video Clip Count** | The number of video clips in your scene. |
| **Total Video Memory** | The amount of system memory the video in your application is using. |

## Additional resources

* [Profiler window introduction](ProfilerWindow.md)
* [Profiling your application](profiler-profiling-applications.md)