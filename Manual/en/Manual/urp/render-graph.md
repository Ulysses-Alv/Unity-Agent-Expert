# Render graph system in URP

The render graph system is a set of APIs you use to create a [Scriptable Render Pass](renderer-features/intro-to-scriptable-render-passes.md).

| Page | Description |
| --- | --- |
| [Introduction to the render graph system](render-graph-introduction.md) | What the render graph system is, and how it optimizes rendering. |
| [Write a render pass using the render graph system](render-graph-write-render-pass.md) | Write a Scriptable Render Pass using the render graph APIs. |
| [Blit using the render graph system](render-graph-blit.md) | **Blit** from one texture to another in the render graph system with the `AddBlitPass` API. |
| [Textures in the render graph system](working-with-textures.md) | Access and use textures in your render passes. |
| [Frame data in the render graph system](render-graph-frame-data.md) | Get the textures URP creates for the current frame and use them in your render passes. |
| [Draw objects in the render graph system](render-graph-draw-objects-in-a-pass.md) | Draw objects in the render graph system using the `RendererList` API. |
| [Compute shaders in the render graph system](render-graph-compute-shader.md) | Create a render pass that runs a compute **shader**. |
| [Analyze a render graph](render-graph-view.md) | Check a render graph using the Render Graph Viewer, Rendering Debugger, or Frame Debugger. |
| [Optimize a render graph](render-graph-optimize.md) | To optimize a render graph, merge or reduce the number of render passes. |
| [Use the CommandBuffer interface in a render graph](render-graph-unsafe-pass.md) | Use `CommandBuffer` interface APIs such as `SetRenderTarget` in render graph system render passes. |
| [Render Graph Viewer window reference](render-graph-viewer-reference.md) | Reference for the **Render Graph Viewer** window. |

## Additional resources

* [Inject a render pass using a Scriptable Renderer Feature](customize/inject-render-pass-via-script.md)
* [Frame Debugger](https://docs.unity3d.com/2023.3/Documentation/Manual/frame-debugger-window.html)