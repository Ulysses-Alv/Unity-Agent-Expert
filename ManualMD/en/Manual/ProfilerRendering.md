# Rendering Profiler module reference

The Rendering Profiler displays rendering statistics and information about what the CPU and GPU do to render the **Scene**. You can use these statistics to measure the resource intensity of different areas of the Scene, which is useful for optimization.

To open the Profiler window, go to menu: **Window** > **Analysis** > **Profiler**.

The chart displays the number of Batches, SetPass Calls, Triangles and Vertices your application rendered. The lower pane displays more rendering statistics, which match the ones shown in the GameView [Rendering Statistics](RenderingStatistics.md) window.

![The Rendering Profiler module](../uploads/Main/RenderProfiler.png)

The Rendering Profiler module

The Rendering Profiler module’s chart is divided into the following categories:

| **Chart** | **Description** |
| --- | --- |
| **Batches Count** | The number of [batches](DrawCallBatching.md) Unity processed during a frame. |
| **SetPass Calls Count** | The number of times Unity switched which **shader** pass it used to render **GameObjects** during a frame. A shader might contain several shader passes and each pass renders GameObjects in the scene differently. |
| **Triangles Count** | The number of [triangles](../ScriptReference/Mesh-triangle.md) Unity processed during a frame. |
| **Vertices Count** | The number of [vertices](../ScriptReference/Mesh-vertices.md) Unity processed during a frame. |

## Module details pane

When you click on the Rendering Profiler module, the details pane in the lower half of the window displays detailed rendering statistics. These statistics are similar to the statistics shown in the [Rendering Statistics](RenderingStatistics.md) window.

In the top left of the details pane, select Open Frame Debugger to open the [Frame Debugger](FrameDebugger.md), which gives you information on individual drawcalls that rendered the frame.

These statistics are also available via the [ProfilerRecorder API](../ScriptReference/Unity.Profiling.ProfilerRecorder.md) and in the [Profiler Module Editor](profiler-module-editor.md) so you can add them to a custom Profiler module.

| **Statistic** | **Description** | **Accessible in Release Players** |
| --- | --- | --- |
| **SetPass Calls** | The number of times Unity switched which shader pass it used to render GameObjects during a frame. A shader might contain several shader passes and each pass renders GameObjects in the scene differently. | Yes |
| **Draw Calls** | The total number of draw calls Unity issued during a frame. Unity issues draw calls when it renders GameObjects to the screen. This number includes non batched draw calls and dynamic and static batched draw calls. | Yes |
| **Batches** | The total number of [batches](DrawCallBatching.md) Unity processed during a frame. This number includes both static and dynamic batches. | Yes |
| **Triangles** | The number of [triangles](../ScriptReference/Mesh-triangles.md) Unity processed during a frame. | Yes |
| **Vertices** | The number of [vertices](../ScriptReference/Mesh-vertices.md) Unity processed during the frame. | Yes |
| **(Dynamic Batching)** | This section contains statistics on [dynamic batching](DrawCallBatching.md). | No |
| Batched Draw Calls | The number of draw calls Unity combined into dynamic batches. | No |
| Batches | The number of dynamic batches Unity processed during the frame. | No |
| Triangles | The number of triangles in the GameObjects included in the dynamic batches. | No |
| Vertices | The number of vertices in the GameObjects included in the dynamic batches. | No |
| Time | The time Unity spent creating dynamic batching structures. | No |
| **(Static Batching)** | This section contains statistics on [static batching](DrawCallBatching.md). | No |
| Batched Draw Calls | The number of draw calls Unity combined into static batches. | No |
| Batches | The number of static batches Unity processed during a frame. | No |
| Triangles | The number of triangles in the GameObjects included in the static batches. | No |
| Vertices | The number of vertices in the GameObjects included in the static batches. | No |
| **(Instancing)** | This section contains statistics on [GPU instancing](GPUInstancing.md). | No |
| Batched Draw Calls | The number of draw calls Unity combined into instance batches. | No |
| Batches | The number of batches Unity processed to render [instanced](GPUInstancing.md) GameObjects during a frame. | No |
| Triangles | The number of triangles in the instanced GameObjects. | No |
| Vertices | The number of vertices in the instanced GameObjects. | No |
| **Used Textures** | The number of [textures](Textures.md) Unity used during the frame and the amount of memory the textures used. | No |
| ****Render Textures**** | The number of [RenderTextures](class-RenderTexture.md) Unity used during the frame and the amount of memory the RenderTextures used. | Yes |
| **Render Textures Changes** | The number of times Unity set one or multiple RenderTextures as render targets during the frame. | Yes |
| **Used Buffers** | The total number of GPU buffers and memory they used. This includes vertex, index and compute buffers and all internal buffers required for rendering. | Yes |
| **Vertex Buffer Upload In Frame** | The amount of geometry that the CPU uploaded to the GPU in the frame. This represents the vertex/normal/texcoord data. There might already be some geometry on the GPU. This statistic only includes geometry that Unity transfers in a frame. | Yes |
| **Index Buffer Upload In Frame Count** | The amount of geometry that the CPU uploaded to the GPU in the frame. This represents the triangle indices data. There might already be some geometry on the GPU. This statistic only includes geometry that Unity transfers in a frame. | Yes |
| **Shadow Casters Count** | The number of GameObjects that cast shadows in a frame. If a GameObject casts multiple shadows (because multiple lights light it), it has one entry per shadow it casts. | Yes |

## Availability in Players

The Rendering module **Profiler counters** are also available in Players. Use the [ProfilerRecorder API](../ScriptReference/Unity.Profiling.ProfilerRecorder.md) to access Render Profiler module information in Players. High level counters are also available in Release Player.

The following example contains a simple script that collects `SetPass Calls Count`, `Draw Calls Count` and `Vertices Count` metrics and displays those as TextArea.

```
using System.Text;
using Unity.Profiling;
using UnityEngine;
 
public class RenderStatsScript : MonoBehaviour
{
    string statsText;
    ProfilerRecorder setPassCallsRecorder;
    ProfilerRecorder drawCallsRecorder;
    ProfilerRecorder verticesRecorder;

    void OnEnable()
    {
        setPassCallsRecorder = ProfilerRecorder.StartNew(ProfilerCategory.Render, "SetPass Calls Count");
        drawCallsRecorder = ProfilerRecorder.StartNew(ProfilerCategory.Render, "Draw Calls Count");
        verticesRecorder = ProfilerRecorder.StartNew(ProfilerCategory.Render, "Vertices Count");
    }
 
    void OnDisable()
    {
        setPassCallsRecorder.Dispose();
        drawCallsRecorder.Dispose();
        verticesRecorder.Dispose();
    }

    void Update()
    {
        var sb = new StringBuilder(500);
        if (setPassCallsRecorder.Valid)
            sb.AppendLine($"SetPass Calls: {setPassCallsRecorder.LastValue}");
        if (drawCallsRecorder.Valid)
            sb.AppendLine($"Draw Calls: {drawCallsRecorder.LastValue}");
        if (verticesRecorder.Valid)
            sb.AppendLine($"Vertices: {verticesRecorder.LastValue}");
        statsText = sb.ToString();
    }
 
    void OnGUI()
    {
        GUI.TextArea(new Rect(10, 30, 250, 50), statsText);
    }
}
```

The Rendering Profiler module information belongs to the [ProfilerCategory.Render](../ScriptReference/Unity.Profiling.ProfilerCategory.Render.md) **Profiler category**.

If you want to highlight the selected Rendering counters in a custom module, use the [Module Editor](profiler-module-editor.md) to configure the chart and detailed view.

## Additional resources

* [Profiler window introduction](ProfilerWindow.md)
* [Profiling your application](profiler-profiling-applications.md)