# Optimize for untethered XR devices in URP

Understand the ways you can optimize your URP project for **XR** devices.

This page describes the optimization techniques for universal **render pipeline** (URP) projects that target untethered extended reality (XR) devices. An untethered XR device is an XR device that aren’t connected to a desktop with a wire.

Most untethered XR devices use tile-based GPUs. The guidelines on this page help you use this hardware architecture more efficiently and avoid using rendering techniques that are less efficient on those devices. To learn more about how tiled-based GPUs work, refer to the additional resources section.

To optimize your XR project, consider the following:

* [Use Vulkan API](#vulkan)
* [Use OpenXR](#openxr)
* [Use the render graph system](#render-graph)
* [Use Forward rendering](#forward-rendering)
* [Use on-tile post-processing](#post-processing)
* [Avoid geometry shaders](#geometry-shaders)
* [Use MSAA for anti-aliasing](#msaa)
* [Disable depth priming](#depth-priming)
* [Disable Opaque texture and Depth texture properties](#textures)
* [Disable SSAO](#ssao)
* [Disable HDR](#hdr)
* [Implement resolution scaling](#resolution-scaling)

The following sections describe these methods in more detail.

## Use Vulkan API

Vulkan API is more stable and provides better performance compared to OpenGL ES API in URP projects targeting XR platforms. Many new OpenXR features are only available on Vulkan.

Refer to [Configure graphics APIs](configure-graphicsAPIs.md#override-default-gfxapi) for information on how to change the graphics API to Vulkan.

## Use OpenXR Plugin

Use the [OpenXR Plugin](https://docs.unity3d.com/Packages/com.unity.xr.openxr@1.16) in projects that target XR platforms.

Enable the following settings in your OpenXR project:

* [Multi-view \ Single pass rendering](https://docs.unity3d.com/Packages/com.unity.xr.openxr@1.15/manual/project-configuration.html#set-the-render-mode)
* [Foveated rendering](https://docs.unity3d.com/Packages/com.unity.xr.openxr@1.15/manual/features/foveatedrendering.html)
* [Multiview render regions](https://docs.unity3d.com/Packages/com.unity.xr.openxr@1.15/manual/features/multiviewrenderregions.html)

**Tip:** Consider whether [URP Application SpaceWarp](https://docs.unity3d.com/Packages/com.unity.xr.openxr@1.16/manual/features/spacewarp.html) is appropriate for your project. Application SpaceWarp can improve frame rate significantly, but might cause visible artifacts for some content.

## Use render graph system

In Unity 6.0 and newer, new URP projects use the render graph system. Refer to [Benefits of the render graph system](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@17.0/manual/render-graph-benefits.html) to understand the benefits of render graph.

## Use Forward rendering

In URP, Deferred rendering generates several render targets for the G-buffer. Unity performs multiple [graphics memory loads](https://developer.qualcomm.com/software/snapdragon-profiler/app-notes/avoid-gmem-loads) (Qualcomm documentation) to access those render targets, which is slow on tile-based GPUs.

Refer to [Deferred rendering implementation details](urp/rendering/deferred-rendering-path-landing.md) for more information on the implementation of the G-buffer in URP.

The ****Rendering Path**** settings is in the [**Rendering**](urp/urp-universal-renderer.md#rendering) section of the [Universal Renderer Asset](urp/urp-universal-renderer.md).

## Use on-tile post processing (Unity 6.3 and newer)

On XR devices, Unity recommends that you use on-tile **post-processing**. Post-processing techniques without on-tile post-processing enabled are expensive on untethered XR devices because post-processing adds intermediate textures.

In Unity 6.3 and newer, you can use on-tile post-processing to enable certain URP post-processing effects on XR devices. Refer to [On-tile post-processing (URP)](xr-graphics-on-tile-post-processing.md) to learn more about on-tile post-processing and how to enable it in your XR project.

In versions of Unity older than Unity 6.3, Unity recommends that you disable post-processing techniques because on-tile post-processing isn’t available.

## Avoid geometry shaders

Avoid using geometry **shaders** on platforms with tile-based GPUs. Some devices don’t support geometry shaders.

The generation of additional primitives and vertices breaks the tiled GPU flow because the division of primitives after the binning pass becomes invalid.

## Use MSAA for anti-aliasing

Tile-based GPUs can store more samples in the same tile. This makes Multi-sample Anti-aliasing (MSAA) efficient on mobile and untethered XR platforms. 2X MSAA value provides a good balance between visual quality and performance.

You can change the MSAA settings in the **Quality** section of the URP Asset. For more information on MSAA, refer to [Anti-aliasing in URP](urp/anti-aliasing.md).

## Disable depth priming

Disable depth priming on XR platforms. XR devices have two views, which increases the performance impact from performing the depth pre-pass.

For untethered XR devices there are no benefits of performing the depth priming. You can obtain similar results using hardware optimization features, such as Low-Resolution-Z (LRZ) or Hidden Surface Removal (HSR).

For information on how to configure or disable depth priming, refer to the [Depth Priming Mode](urp/urp-universal-renderer.md#rendering) property description.

## Disable Opaque texture and Depth texture properties

Disable the **Opaque Texture** and **Depth Texture** properties to improve performance. Enabling those options causes extra texture copy operations, which requires extra GMEM loads.

Refer to [Rendering](urp/universalrp-asset.md#rendering) for more information on these opaque and depth texture properties..

To identify if your current configuration is using intermediate textures, use the [Render Graph Viewer](urp/render-graph-view.md) window.

## Disable SSAO

Screen Space **Ambient Occlusion** (SSAO) might have poor performance on mobile and untethered XR devices.

SSAO in URP requires the depth priming pass, two blur passes to reduce the noise, and an additional **blit** pass to clean the image. Such passes require several GMEM loads which have a high performance impact on tile-based GPUs.

Refer to [SSAO](urp/post-processing-ssao.md) for more information.

## Disable HDR

High Definition Rendering (HDR) has higher precision and improves graphics fidelity, but requires more bits per **pixel** to process. Disable **HDR** to reduce memory bandwidth and improve performance.

Most untethered XR devices don’t support HDR rendering.

To disable HDR:

1. Select a URP Asset.
2. In the ****Inspector**** window, in the **Quality** section, clear the **HDR** property.

**Note**: some features, such as HDR light estimation require HDR. Check the requirements for the features your project uses to ensure HDR isn’t required.

## Implement resolution scaling

Unity recommends that you implement resolution scaling to reduce GPU workload and maintain a stable frame rate.

To learn more, refer to [Resolution control in XR projects](xr-graphics-resolution-scaling.md).

## Additional resources

* [Tile-based rendering in XR](xr-graphics-on-tile-rendering.md)
* [Tile-Based Rendering](https://developer.arm.com/documentation/102662/0100/Overview) (Arm)