# Choose a method for optimizing draw calls

Use draw call optimization if the CPU sends too many draw calls to the GPU, which is also known as being CPU-bound. For more information about draw calls, refer to [Introduction to optimizing draw calls](optimizing-draw-calls.md).

## Check the number of draw calls

To check if your **scene** sends too many draw calls, do any of the following:

* Open the [Rendering Statistics window](RenderingStatistics.md) and check the **SetPass calls** value.
* Open the [Profiler](Profiler.md) and select the **Rendering** section to display the number of draw calls.
* Check the number of draw calls in the [Frame Debugger](FrameDebugger-landing.md).

## Choose optimization methods

The supported and recommended methods for optimizing draw calls depend on whether you use the Universal **Render Pipeline** (URP), the High Definition Render Pipeline (HDRP), or the Built-In Render Pipeline (BIRP).

| **Feature** | **Type** | **Multithreaded** | **Recommendation in URP** | **Recommendation in HDRP** | **Recommendation in the Built-In Render Pipeline** |
| --- | --- | --- | --- | --- | --- |
| **Scriptable Render Pipeline (SRP) Batcher** | Reduces render state updates | No. | Enable. Refer to [SRP Batcher in URP](SRPBatcher.md) | Enable. Refer to [SRP Batcher in HDRP](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.3/manual/SRPBatcher-landing.html) | Not supported. |
| **GPU Resident Drawer** | Uses GPU instancing. | Yes. | Enable. Refer to [GPU Resident Drawer in URP](urp/gpu-resident-drawer.md) | Enable. Refer to [GPU Resident Drawer in HDRP](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.3/manual/gpu-resident-drawer.html) | Not supported. |
| [**`BatchRendererGroup` (BRG) API**](batch-renderer-group.md) | Uses GPU instancing. | Yes. | Use the GPU Resident Drawer instead, except for advanced use cases. | Use the GPU Resident Drawer instead, except for advanced use cases. | Not supported. |
| **GPU Instancing** | Uses GPU instancing. | No. | Disable, to avoid extra **shader** variants. | Disable, to avoid extra shader variants. | Enable if you have many instances or lights. Refer to [GPU Instancing](GPUInstancing-landing.md). |
| **Batching** | Combines meshes | No. | Disable. **Static batching** isn’t compatible with the BRG API or GPU Resident Drawer. | Disable. Static batching isn’t compatible with the BRG API or GPU Resident Drawer. | Enable static batching. **Dynamic batching** is no longer recommended. Refer to [Batching](DrawCallBatching-landing.md). |

If you have many instances or lights, you can also use [optimize mesh rendering using level of detail (LOD)](lod-landing.md).

To get the best results from any draw call optimization method, do the following:

* Use the same materials for different **GameObjects** as much as possible.
* If you want to render the same **mesh** with different properties, for example different colors, use [Material Variants](materialvariant-landingpage.md) instead of multiple materials.
* Avoid using the `MaterialPropertyBlock` API in URP and HDRP.

## Use multiple optimization methods

You can use multiple draw call optimization methods in the same scene, but each GameObject uses only one of the following methods, in this priority order:

1. SRP Batcher, and static batching if you enable it.
2. GPU instancing, including the GPU Resident Drawer or the `BatchRendererGroup` API.
3. Dynamic batching.

## Additional resources

* [Shader variants](shader-variants.md)
* [Fantasy Kingdom in Unity 6](https://unity.com/demos/fantasy-kingdom), which uses the GPU Resident Drawer.
* [SRP Batcher: Speed up your rendering](https://unity.com/blog/engine-platform/srp-batcher-speed-up-your-rendering)
* [BatchRendererGroup sample: Achieve high frame rate even on budget devices](https://unity.com/blog/engine-platform/batchrenderergroup-sample-high-frame-rate-on-budget-devices)
* [Megacity Metro](https://unity.com/demos/megacity-competitive-action-sample), which uses the BatchRendererGroup API to render a large number of objects.