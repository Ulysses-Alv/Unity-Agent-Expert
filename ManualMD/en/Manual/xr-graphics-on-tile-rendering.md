# Tile-based rendering in XR

Understand on-tile rendering in **XR**.

Tile-based GPUs are common in mobile and untethered XR devices. Performance on devices with a tile-based GPU is improved when your project uses features that are optimized for tile-based architecture.

Rendering off-tile results in additional render passes and requires intermediate data to be written to and read from external memory. Graphical loads are inefficient and increase memory bandwidth usage and power consumption, leading to reduced performance and shorter battery life. To maximize performance on untethered XR devices, minimize off-tile operations and use features optimized for tile-based rendering.

Refer to the following sections to learn more about tile-based rendering in XR, and methods to ensure rendering remains on-tile.

## About tile-based GPUs

Tile-based GPUs divide the screen into small regions (tiles), which are rendered on-chip in the tile buffer. When a tile is rendered, its data is written to the framebuffer in the main memory. Tile-based GPUs enable faster memory and improved power-efficiency on mobile and untethered XR devices.

## Tips to ensure tile-based rendering

If your project targets an untethered XR device and enables [On-tile post-processing](xr-graphics-on-tile-post-processing.md), you must ensure your project only renders on-tile.

Any feature that creates intermediate textures results in an additional final **blit** pass and breaks tile-based rendering, resulting in app and device performance issues.

Only [memoryless intermediate textures](../ScriptReference/RenderTexture-memorylessMode.md) are supported on-tile. To use memoryless intermediate textures, your project must use the Vulkan or Metal API. To understand how to configure your API, refer to [Configure graphics APIs](configure-graphicsAPIs.md#override-default-gfxapi).

Features that create non-memoryless intermediate textures result in an additional final blit pass and are unable to stay on-tile, which breaks the render pass. Non-memoryless intermediate textures introduce performance issues from moving data from on-tile to off-tile memory.

You should avoid the following features on untethered XR devices as they introduce off-tile rendering:

* [Intermediate textures](urp/urp-universal-renderer.md)
* [Depth](urp/universalrp-asset.md) and [Opaque](urp/universalrp-asset.md) textures.
* [Upscaling filter](urp/universalrp-asset.md)
* URP built-in Post processing
* [Renderer feature](urp/urp-renderer-feature.md) that require intermediate texture:
  + [Decal](urp/renderer-feature-decal-landing.md)
  + [FullScreenPass](urp/renderer-features/renderer-feature-full-screen-pass.md) if **FetchColorBuffer** is `true`
  + [SSAO](urp/post-processing-ssao-landing.md)

### Use Render Graph Viewer to check render passes

To ensure your project remains on-tile, configure your [URP asset](urp/universalrp-asset.md) and [URP renderer](urp/urp-renderer-feature.md) to render directly to the eye texture, which means that your project uses a single merged render pass. You can use the [Render Graph Viewer](urp/render-graph-viewer-reference.md) to [Analyze a render graph in URP](urp/render-graph-view.md) and [Merge render passes](urp/render-graph-optimize.md#merge-render-passes.md).

#### Render Graph Viewer example

The following image demonstrates the optimal pipeline configuration with on-tile **post-processing** and how to interpret the Render Graph Viewer:

![Render Graph Viewer example of merged render passes that will render on-tile.](../uploads/Main/xr-graphics-on-tile-render-graph.png)

Render Graph Viewer example of merged render passes that will render on-tile.

The blue line (A) represents merged render passes.

**cameraTargetAttachment** (B) and **cameraDepthAttachment (C)** are memoryless on-tile resources. The on-tile post-processing features uses frame buffer fetch (denoted by an **F** symbol) (D) to fetch the **cameraTargetAttachment** and write to the **Backbuffer color** (E).

## Additional resources

* [On-tile post processing](xr-graphics-on-tile-post-processing.md)
* [Optimize for untethered XR devices in URP](xr-untethered-device-optimization.md)
* [Mobile GPUs and tiled rendering]([https://developers.meta.com/horizon/documentation/unity/gpu-tiled](https://developers.meta.com/horizon/documentation/unity/gpu-tiled/?locale=en_GB)) (Meta developer documentation)