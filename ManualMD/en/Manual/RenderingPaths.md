# Introduction to rendering paths in the Built-In Render Pipeline

Unity’s Built-In **Render Pipeline** supports different rendering paths. A **rendering path** is a series of operations related to lighting and shading. Different rendering paths have different capabilities and performance characteristics. Deciding on which rendering path is most suitable for your Project depends on the type of Project, and on the target hardware.

You can choose the rendering path that your Project uses in the [Graphics](class-GraphicsSettings.md) window, and you can override that path for each [Camera](class-Camera.md).

If the GPU on the device running your Project does not support the rendering path that you have selected, Unity automatically uses a lower fidelity rendering path. For example, on a GPU that does not handle **Deferred Shading**, Unity uses **Forward Rendering**.

* [Forward](RenderTech-ForwardRendering.md) is the default rendering path in the Built-in Render Pipeline. It is a general-purpose rendering path.
* [Deferred](RenderTech-DeferredShading.md) is the rendering path with the most lighting and shadow fidelity in the Built-in Render Pipeline.
* [Legacy Vertex Lit](RenderTech-VertexLit.md) is the rendering path with the lowest lighting fidelity and no support for real-time shadows. It is a subset of the Forward rendering path.

## Rendering Path comparison

|  | ***Deferred*** | ***Forward*** | ***Vertex Lit*** |
| --- | --- | --- | --- |
| **Features** |  |  |  |
| Per-pixel lighting (normal maps, light cookies) | Yes | Yes | - |
| Real-time shadows | Yes | With caveats | - |
| **Reflection Probes** | Yes | Yes | - |
| Depth & Normal Buffers | Yes | Additional render passes | - |
| Soft **Particles** | Yes | - | - |
| Semi-transparent objects | - | Yes | Yes |
| Anti-Aliasing | - | Yes | Yes |
| Light **Culling Masks** | Limited | Yes | Yes |
| Lighting Fidelity | All per-pixel | Some per-pixel | All per-vertex |
| **Performance** |  |  |  |
| Cost of a per-pixel Light | Number of **pixels** it illuminates | Number of pixels \* Number of objects it illuminates | - |
| Number of times objects are normally rendered | 1 | Number of per-pixel lights | 1 |
| Overhead for simple **Scenes** | High | None | None |
| **Platform Support** |  |  |  |
| PC (Windows/Mac) | **Shader** Model 3.0+ & MRT | All | All |
| Mobile (iOS/Android) | OpenGL ES 3.0 & MRT, Metal | All | All |
| Consoles | XB1, PS4 | All | - |