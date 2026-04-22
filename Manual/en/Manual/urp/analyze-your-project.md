# Analyze your project in URP

You can use the [Unity Profiler](https://docs.unity3d.com/Manual/Profiler.html) to get data on the performance of your project in areas such as the CPU and memory.

## Profiler markers

The following table lists markers that appear in the Unity Profiler for a URP frame and have a significant effect on performance.

The table doesn’t include a marker if it’s deep in the Profiler hierarchy, or the label already describes what URP does.

| **Marker** | **Description** |
| --- | --- |
| **Inl\_UniversalRenderPipeline. RenderSingleCameraInternal** | URP builds a list of rendering commands in the [`ScriptableRenderContext`](https://docs.unity3d.com/ScriptReference/Rendering.ScriptableRenderContext.html), for a single camera. URP only records rendering commands in this marker, but doesn’t yet execute them. The marker includes the camera name, for example **Main Camera**. This marker has the following sub-markers:  * **Inl\_ScriptableRenderer.Setup**: URP prepares for rendering, for example preparing render textures for the camera and shadow maps. * **CullScriptable**: URP generates a list of GameObjects and lights to render, and culls (excludes) any that are outside the camera’s view. The time this takes depends on the number of GameObjects and lights in your scene. |
| **Inl\_ScriptableRenderContext.Submit** | URP submits the list of commands in the `ScriptableRenderContext` to the graphics API. This marker might appear more than once if URP submits commands more than once per frame, or you call [`ScriptableRenderContext.Submit`](https://docs.unity3d.com/ScriptReference/Rendering.ScriptableRenderContext.Submit.html) in your own code. This marker has the following sub-markers:  * **MainLightShadow**: URP renders a [shadow map](../shadow-mapping.md) for the main Directional Light. * **AdditionalLightsShadow**: URP renders shadow maps for other lights. * **UberPostProcess**: URP renders [post-processing effects](EffectList.md) you enable. This marker contains separate markers for some post-processing effects. * **RenderLoop.DrawSRPBatcher**: URP uses the [Scriptable Render Pipeline Batcher](../SRPBatcher.md) to render one or more batches of objects. |
| **CopyColor** | URP copies the color buffer from one **render texture** to another. You can disable **Opaque Texture** in the [URP Asset](universalrp-asset.md), so that URP only copies the color buffer if it needs to. |
| **CopyDepth** | URP copies the **depth buffer** from one render texture to another. You can disable **Depth Texture** in the [URP Asset](universalrp-asset.md) unless you need the depth texture (for example, if you use a shader that uses scene depth). |
| **FinalBlit** | URP copies a render texture to the current **camera** render target. |

## Use a GPU profiler to analyze your project

You can use a platform GPU profiler such as [Xcode](https://docs.unity3d.com/Manual/XcodeFrameDebuggerIntegration.html) to get data on the performance of the GPU during rendering. You can also use a profiler such as [RenderDoc](https://docs.unity3d.com/Manual/RenderDocIntegration.html), but it might provide less accurate performance data.

Data from a GPU profiler includes URP markers for rendering events, such as different render passes.

## Use other tools to analyze your project

You can also use the following tools to analyze the performance of your project:

* [Scene view View Options](https://docs.unity3d.com/Manual/ViewModes.html)
* [Rendering Debugger](features/rendering-debugger.md)
* [Frame Debugger](https://docs.unity3d.com/Manual/frame-debugger-window.html)
* [Graphics performance and profiling reference](../profiling-landing.md)