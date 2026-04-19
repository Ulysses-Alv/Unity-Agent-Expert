# Global Illumination Profiler module reference

The **Global Illumination** Profiler module displays statistics about how much CPU time the [Enlighten Realtime Global Illumination](realtime-gi-using-enlighten.md) subsystem uses across all worker threads. To control the number of **Enlighten** worker threads included, you can pass the Unity Editor a [command line argument](EditorCommandLineArguments.md).

To open the Profiler window, go to menu: **Window** > **Analysis** > **Profiler**.

![Global Illumination Profiler module](../uploads/Main/gi-chart.png)

Global Illumination Profiler module

## Chart categories

The Global Illumination Profiler module’s chart tracks the time the global illumination subsystem spent on all worker threads. The timings are divided into the following categories:

| **Chart** | **Description** |
| --- | --- |
| **Light Probe** | The time spent updating [Light Probes](LightProbes.md). |
| **Setup** | The time spent in the setup stage. |
| **Environment** | The time spent processing environment lighting. |
| **Input Lighting** | The time spent processing input lighting. |
| **Systems** | The time spent updating systems. |
| **Solve Tasks** | The time spent running [radiosity solver tasks](class-MeshRenderer.md). |
| **Dynamic Objects** | Time spent updating dynamic **GameObjects**. |
| **Other Commands** | Time spent updating other commands. |
| **Blocked Command Write** | Time spent in a blocked state, waiting for the command buffer. |

To change the order of the categories in the chart, you can drag and drop them in the chart’s legend. You can also click a category’s colored legend to toggle its display.

## Module details pane

When you select the Global Illumination module, the details pane displays a breakdown of where the application spent time in the selected frame. The data available is as follows:

| **Data** | **Description** |
| --- | --- |
| **Total CPU Time** | Total Enlighten Global Illumination CPU time across all threads in ms. |
| **Probe Update Time** | Time spent updating light probes in ms. |
| **Setup Time** | Time spent in the setup stage in ms. |
| **Environment Time** | Time spent processing environment lighting in ms. |
| **Input Lighting Time** | Time spent processing input lighting in ms. |
| **Systems Time** | Time spent updating systems in ms. |
| **Solve Tasks Time** | Time spent running radiosity solver tasks in ms. |
| **Dynamic Objects Time** | Time spent updating dynamic GameObjects in ms. |
| **Time Between Updates** | Time between Global Illumination updates in ms. |
| **Other Commands Time** | Time spent processing other commands in ms. |
| **Blocked Command Write Time** | Time spent in blocked state, waiting for command buffer in ms. |
| **Blocked Buffer Writes** | Number of writes to the command buffer that were blocking. |
| **Total Light Probes** | Total number of Light Probes in the **scene**. |
| **Solved Light Probes** | Number of solved Light Probes since the last update. |
| **Probe Sets** | Number of Light Probe sets in the scene. |
| **Systems** | Number of Enlighten Global Illumination systems in the scene. |
| **Pending Material GPU Renders** | Number of Albedo/Emission renders queued for rendering on the GPU. |
| **Pending Material Updates** | Number of material updates waiting to get processed. |

## Additional resources

* [Profiler window introduction](ProfilerWindow.md)
* [Profiling your application](profiler-profiling-applications.md)