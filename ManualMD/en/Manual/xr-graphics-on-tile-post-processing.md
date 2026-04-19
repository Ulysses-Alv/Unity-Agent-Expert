# On-tile post-processing (URP)

Learn about on-tile **post-processing** for untethered **XR** devices.

On-tile post-processing leverages tile-based GPUs to apply post-processing effects while the rendered image data is in the GPU’s tile-based memory. On-tile post-processing enables you to use certain URP post-processing techniques on untethered XR devices, which otherwise you should avoid. On-tile post-processing can result in significant performance gains, and reduce device power and memory consumption.

Refer to the following sections to learn about on-tile post processing.

## Introduction to on-tile post-processing

On non-tile-based GPUs, post-processing effects are applied after the entire frame is rendered and stored in the framebuffer. On untethered XR devices, Unity recommends that you don’t enable post-processing techniques without on-tile post-processing. Without on-tile post-processing enabled, post-processing adds intermediate textures which you should avoid on XR devices because they can decrease performance of your application.

With on-tile post-processing, post-processing effects are applied per-tile, before the tile is moved to the framebuffer in the main memory. This reduces memory usage, reduces power consumption, and improves performance, because the tile is only stored in its processed state. On-tile post processing enables you to use [supported post-processing techniques](#supported-techniques) in your XR project, and improve the graphics in your application.

The URP on-tile post-processing feature is available as a Renderer Feature. To learn more about renderer features, refer to [Add a Renderer Feature to a URP Renderer](urp/urp-renderer-feature.md).

## Requirements

To use on-tile post-processing your project must meet the following requirements:

* Unity 6.3 or newer.
* Use the [Universal Render Pipeline](urp/urp-introduction.md).
* Disable integrated [URP post-processing](urp/integration-with-post-processing.md).
* Use the Vulkan graphics API.

To use on-tile post-processing, rendering must stay on-tile. Rendering off-tile will result in compilation errors when on-tile post-processing is enabled. For tips to ensure your project renders to on-tile memory, refer to [Tile-based rendering in XR](xr-graphics-on-tile-rendering.md).

## Supported platforms

On-tile post-processing is recommended for untethered XR devices such as Meta Quest 2, Meta Quest 3, Android XR devices, and other [OpenXR runtimes](https://docs.unity3d.com/Packages/com.unity.xr.openxr@1.16/manual/index.html#runtimes).

On-tile post-processing is also supported on Play mode for [Android](android-BuildProcess.md) and [Meta Quest](xr-meta-quest-build-profile.md) build targets. For example [XR Mock HMD](https://docs.unity3d.com/Packages/com.unity.xr.mock-hmd@1.4/manual/index.html), and [Meta Quest link](https://docs.unity3d.com/Packages/com.unity.xr.meta-openxr@2.1/manual/get-started/link.html).

**Note**: If you enable on-tile post-processing on unsupported platforms, URP uses a fallback mechanism to use a limited version of the post-processing feature. This fallback mechanism ensures correct rendering but doesn’t apply post-processing on-tile.

## Supported post-processing techniques

The following [URP post-processing techniques](urp/integration-with-post-processing.md) support on-tile post-processing:

* Color grading
* Vignette
* Tonemapping
* Dithering
* Film Grain

## Enable on-tile post-processing

**Tip**: Before you enable on-tile post-processing on device, Unity recommends that you [Use the Render Graph Viewer to check render passes](xr-graphics-on-tile-rendering.md#use-render-graph-viewer).

To enable on-tile post-processing in your project:

1. [Open your URP Renderer Data](urp/urp-universal-renderer.md) in the ****Inspector**** window.
2. Ensure **Post-processing** is disabled.
3. Under **Renderer Features**, enable **On Tile Post Processing (Untethered XR)**.

![On-tile post-processing highlighted in the Inspector window.](../uploads/Main/xr-graphics-on-tile-post-processing.png)

On-tile post-processing highlighted in the Inspector window.

When you enable on-tile post-processing, any [supported post-processing](#supported-techniques) techniques you enable in your project will be applied on-tile on the target device. Enabling on-tile post-processing doesn’t enable any post-processing technique automatically, and you must enable each technique.

**Note**: When you enable on-tile post-processing, Unity will display a warning that post-processing is disabled on the current URP renderer. You can safely ignore this message.

## Additional resources

* [Introduction to render pipelines](render-pipelines-overview.md)
* [On-tile rendering in XR](xr-graphics-on-tile-rendering.md)