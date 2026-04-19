# Multiview Render Regions

Understand the Multiview Render Regions feature for **XR** devices.

The Multiview Render Regions feature is an optimization technique to prevent processing on areas of the screen not visible by the user (the nasal region of a VR headset).

## About Multiview Render Regions

Multiview Render Regions is a Vulkan-specific feature that sets multiple scissors, **viewports**, and render areas on the headset to reduce rendering load. On **VR** devices, Multiview Render Regions sets specific values per eye so that areas outside of the user’s view aren’t processed. On Head-Mounted Devices (HMD), this is the nasal region of the headset.

![With Multiview Render Regions enabled (right), the nasal region hidden from the user’s view isn’t rendered to reduce unnecessary processing. Theres no visible difference in the headset (bottom).](../uploads/Main/xr-multiview-render-regions.png)

With Multiview Render Regions enabled (right), the nasal region hidden from the user’s view isn’t rendered to reduce unnecessary processing. There’s no visible difference in the headset (bottom).

Multiview Render Regions incorporates two Vulkan extensions ([Multiview per view viewports](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_multiview_per_view_viewports.html), and [Multiview per view render areas](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_multiview_per_view_render_areas.html)) (Khronos documentation). When you enable the Multiview Render Regions feature, you opt into using these extensions.

## Prerequisites

To use Multiview Render Regions, your project must meet the following requirements:

* Unity 6.1 or newer.
* Use the Vulkan API. (Refer to [Configure graphics APIs](configure-graphicsAPIs.md) to understand how to set your graphics API to Vulkan.)
* [All Passes](https://docs.unity3d.com/Packages/com.unity.xr.openxr@1.15/manual/features/multiviewrenderregions.html) mode is only available in Unity 6.2 or newer with OpenXR 1.15 or newer.
* In Unity 6.3, your project must use the [render graph](urp/render-graph.md) system. (Refer to [Enable render graph](#enable-render-graph) for more information.)

### Supported plug-ins

The following **plug-ins** support Multiview Render Regions:

* [Oculus](https://docs.unity3d.com/Packages/com.unity.xr.oculus@latest) 4.6 and newer
* [OpenXR](https://docs.unity3d.com/Packages/com.unity.xr.openxr@1.15/manual/features/multiviewrenderregions.html) 1.14 and newer

## Configure project settings

Configure the following settings in the **XR Plug-in Management** section of your **Project Settings** (menu: **Edit** > **Project Settings** > **XR Plug-in Management**):

* **Stereo Rendering Mode** set to **Multiview**.
* Enable **Symmetric Projection (Vulkan)**.

### Post-processing and intermediate textures

In Unity 6.2 and newer with [OpenXR 1.15](https://docs.unity3d.com/Packages/com.unity.xr.openxr@1.15/manual/features/multiviewrenderregions.html), you can choose which render passes Unity applies Multiview Render Regions to. To learn about the available options for **Multiview Render Regions Optimizations** in OpenXR 1.15, refer to the OpenXR [Multiview Render Regions](https://docs.unity3d.com/Packages/com.unity.xr.openxr@1.15/manual/features/multiviewrenderregions.html) documentation.

If you select **Final Pass** in OpenXR 1.15, or enable **Optimize Multiview Render Regions** in OpenXR 1.14 or Oculus, Multiview Render Regions is applied to the final render pass only. This means Multiview Render Regions is only applied to render passes that output to eye textures. As such, performance gains aren’t expected when using intermediate textures or **post-processing** with this feature enabled. If you use intermediate textures or post-processing in your project, profile your project to verify whether Multiview Render Regions will lead to performance gains before you enable it.

**Note**: Unity recommends that you avoid post-processing if your project targets standalone XR devices. Refer to [Optimize for untethered XR devices in URP](urp/xr-untethered-device-optimization) for more information about graphics optimization in XR.

## Enable render graph

In Unity 6.3 and newer, Multiview Render Regions requires that your project uses [render graph](urp/render-graph.md) for both **Final Pass** and **All Passes** modes. Render graph helps prevent some artifacts that occurred when you select the **All Passes** option in previous versions of Unity.

The render graph [MultiviewRenderRegionsCompatible](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@17.3/api/UnityEngine.Rendering.RenderGraphModule.ExtendedFeatureFlags.html) flag marks [compatible passes](#passes) to enable Multiview Render Regions only on compatible passes. This prevents rendering artifacts observed in earlier versions of Unity, where Multiview Render Regions is applied to passes that aren’t compatible.

To ensure your project uses render graph, disable [Compatibility Mode (Render Graph Disabled)](urp/compatibility-mode) in your **Graphics** settings.

### Manage render pass compatibility

In Unity 6.3 and newer with render graph, Unity enables Multiview Render Regions only on passes where it increases performance and doesn’t result in rendering artifacts. Compatible passes are marked with the [MultiviewRenderRegionsCompatible](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@17.3/api/UnityEngine.Rendering.RenderGraphModule.ExtendedFeatureFlags.html) flag.

The following passes work well with Multiview Render Regions and don’t cause any rendering artifacts:

* [CopyDepthPass](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@17.3/api/UnityEngine.Rendering.Universal.Internal.CopyDepthPass.html)
* DecalScreenSpaceRenderPass
* [DepthNormalOnlyPass](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@17.3/api/UnityEngine.Rendering.Universal.Internal.DepthNormalOnlyPass.html)
* [DepthOnlyPass](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@17.3/api/UnityEngine.Rendering.Universal.Internal.DepthOnlyPass.html)
* [DrawObjectsPass](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@17.3/api/UnityEngine.Rendering.Universal.Internal.DrawObjectsPass.html)
* [DrawSkyboxPass](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@17.3/api/UnityEngine.Rendering.Universal.DrawSkyboxPass.html)
* [FinalBlitPass](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@17.3/api/UnityEngine.Rendering.Universal.Internal.FinalBlitPass.html)
* MotionVectorRenderPass
* PostProcessPassRenderGraph
* [RenderObjectsPass](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@17.3/api/UnityEngine.Rendering.Universal.RenderObjectsPass.html)
* [TemporalAA](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@17.3/api/UnityEngine.Rendering.Universal.TemporalAA.html)
* [XRDepthMotionPass](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@17.3/api/UnityEngine.Rendering.Universal.XRDepthMotionPass.html)
* [XROcclusionMeshPass](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@17.3/api/UnityEngine.Rendering.Universal.XROcclusionMeshPass.html)

To add additional render graph passes to use with the Multiview Render Regions feature, use the render graph [SetExtendedFeatureFlags(ExtendedFeatureFlags)](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@17.3/api/UnityEngine.Rendering.RenderGraphModule.IRasterRenderGraphBuilder.html#UnityEngine_Rendering_RenderGraphModule_IRasterRenderGraphBuilder_SetExtendedFeatureFlags_UnityEngine_Rendering_RenderGraphModule_ExtendedFeatureFlags_) method to set the [MultiviewRenderRegionsCompatible](https://docs.unity3d.com/Packages/com.unity.render-pipelines.core@17.3/api/UnityEngine.Rendering.RenderGraphModule.ExtendedFeatureFlags.html) flag for applicable raster passes.

## Enable Multiview Render Regions

If your chosen plug-in [supports Multiview Render Regions](#supported-plug-ins), to enable Multiview Render Regions:

1. Go to the **XR Plug-in Management** section of your **Project Settings** (menu: **Edit** > **Project Settings**).
2. Either:
   * Enable **Optimize Multiview Render Regions** (Oculus or OpenXR 1.14).
   * Choose the **Multiview Render Regions Optimizations** mode (OpenXR 1.15). Refer to the OpenXR [Multiview Render Regions](https://docs.unity3d.com/Packages/com.unity.xr.openxr@1.15/manual/features/multiviewrenderregions.html) for information on the available options.

## Additional resources

* [Optimize for untethered XR devices in URP](urp/xr-untethered-device-optimization)
* [Stereo rendering](SinglePassStereoRendering.md)