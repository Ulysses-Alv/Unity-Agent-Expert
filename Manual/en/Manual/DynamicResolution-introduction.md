# Introduction to Dynamic Resolution

Dynamic resolution is a **Camera** setting that allows you to dynamically scale individual render targets, to reduce workload on the GPU. For example, you can gradually scale down the resolution to maintain a consistent frame rate. If scaled gradually, **dynamic resolution** can be almost unnoticeable.

## Render pipeline compatibility

Dynamic resolution support depends on which [render pipeline](render-pipelines.md) your project uses.

| **Feature** | **Universal Render Pipeline (URP)** | **High Definition Render Pipeline (HDRP)** | **Built-in Render Pipeline** |
| --- | --- | --- | --- |
| **Dynamic resolution** | Yes (1) | Yes (2) | Yes (1) |

**Notes**:

1. The [Built-in Render Pipeline](built-in-render-pipeline.md), and the [Universal Render Pipeline (URP)](universal-render-pipeline.md) both support dynamic resolution as described in this document.
2. The [High Definition Render Pipeline (HDRP)](high-definition-render-pipeline.md) supports dynamic resolution, but you enable and use it in a different way. For information on dynamic resolution in HDRP, see [Dynamic resoluton in HDRP](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest?subfolder=/manual/Dynamic-Resolution.html).

## Supported platforms

Unity supports dynamic resolution on iOS, macOS and tvOS (Metal only), Android (Vulkan only), Windows Standalone (DirectX 12 only), and UWP (DirectX 12 only).

## Impact on render targets

With dynamic resolution, Unity does not re-allocate render targets. Conceptually, Unity scales the render target; however, in reality, Unity uses aliasing, and the scaled-down render target only uses a small portion of the original render target. Unity allocates the render targets at their full resolution, and then the dynamic resolution system scales them down and back up again, using a portion of the original target instead of re-allocating a new target.

## Additional resources

* [FrameTimingManager API reference](../ScriptReference/FrameTimingManager.md)
* [FrameTimingManager User Manual documentation](frame-timing-manager.md)
* [ScalableBufferManager API reference](../ScriptReference/ScalableBufferManager.md)
* [RenderTexture API reference](../ScriptReference/RenderTexture.md)