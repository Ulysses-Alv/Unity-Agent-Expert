# Optimizing draw calls in URP

Techniques for speeding up rendering by reducing the number of drawing commands the CPU sends to the GPU, in the Universal **Render Pipeline** (URP).

For more information about techniques for optimizing draw calls, refer to [choose a method for optimizing draw calls](optimizing-draw-calls-choose-method.md).

| **Page** | **Description** |
| --- | --- |
| [Scriptable Render Pipeline Batcher](SRPBatcher-landing.md) | Resources for using the Scriptable Render Pipeline (SRP) Batcher to reduce the number of render state changes between draw calls. |
| [GPU Resident Drawer](urp/reduce-rendering-work-on-cpu.md) | Automatically use the `BatchRendererGroup` API to use instancing and reduce the number of draw calls. |
| [BatchRendererGroup API](batch-renderer-group.md) | Resources for using the `BatchRendererGroup` API to reduce the number of batches in the SRP Batcher. |

## Additional resources

* [Reduce rendering work on the CPU](urp/reduce-rendering-work-on-cpu.md)
* [Optimize rendering lots of objects](reduce-draw-calls-landing.md)
* [Graphics rendering: Getting the best performance with Unity 6](https://www.youtube.com/watch?v=Oc6T4hh5gaI)
* [Performance tips and tricks from a Unity consultant](https://www.youtube.com/watch?v=CmD8MVGkDxQ)