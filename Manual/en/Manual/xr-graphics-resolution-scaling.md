# Resolution control in XR projects

Understand how to control the resolution in your untethered **XR** project.

You can implement resolution scaling to reduce GPU workload and optimize performance of your XR project. Some methods of resolution control on XR devices implement dynamic resolution. Dynamic resolution adjusts the resolution automatically to maintain a stable frame rate.

The way you control the resolution in your XR project depends on the **render pipeline** your project uses.

## Resolution control reference

The following table lists the methods you can use to scale resolution in your XR project and which render pipeline each method is compatible with:

| **Method** | **Render pipeline compatibility** |
| --- | --- |
| [Control resolution with XRSettings.renderViewportScale](#render-viewport-scale) | URP |
| [Change render scale in the URP Asset](#urp-asset) | URP |
| [Enable automatic viewport dynamic resolution](#openxr) | URP |
| [Dynamic resolution system](#hdrp) | HDRP |
| [Change the XRSettings.eyeTextureResolutionScale](#hdrp) | HDRP |

## Control resolution with URP

If your project uses the universal render pipeline (URP), Unity recommends that you [Control resolution with XRSettings.renderViewportScale](#render-viewport-scale). If your project uses post-processing, you can use the `XRSettings.renderViewportScale` with additional set up as outlined in [Use XRSettings.renderViewPortScale with post-processing enabled](#render-viewport-scale-post-processing).

You can also [Change render scale in the URP Asset](#urp-asset), but this is an expensive method.

If your project uses Unity 6.3 or newer, and [OpenXR](https://docs.unity3d.com/Packages/com.unity.xr.openxr@1.16/manual/index.html) 1.16 or newer, you can [Enable automatic viewport dynamic resolution](#openxr) to automatically change the resolution on OpenXR runtimes.

### Control resolution with XRSettings.renderViewportScale

URP renders directly to the eye texture when you don’t enable [post-processing](urp/integration-with-post-processing.md) methods, or when you enable **HDR**. When URP renders directly to the eye texture, you can use [XRSettings.renderViewportScale](../ScriptReference/XR.XRSettings-renderViewportScale.md) directly. `XRSettings.renderViewportScale` controls how much of the allocated eye texture URP uses for rendering. You can change this value at runtime to dynamically adjust the eye render resolution.

Adjusting the `renderViewPortScale` can be the best way to take advantage of dynamic resolution. You can change the **viewport** scale every frame without the performance penalty of reallocating the eye textures.

To enable dynamic resolution, you can use [XRSettings.renderViewportScale](../ScriptReference/XR.XRSettings-renderViewportScale.md) to control the portion of the allocated eye texture to which the rendering pipeline renders. For example, if you set a value of 0.5, then URP renders into one quarter of the texture (half the height and half the width). The XR device then scales this area to fill the final display. You can change the `renderViewportScale` value at runtime to dynamically adjust the eye render resolution.

`XRSettings.renderViewportScale` is compatible with [fixed foveated rendering](xr-foveated-rendering.md).

**Note**: `XRSettings.renderViewportScale` is a developer hint to the XR runtime for the desired viewport scale. However, the runtime might choose to use this value or ignore it based on its internal logic and performance considerations. You can query [XRSettings.appliedRenderViewportScale](../ScriptReference/XR.XRSettings-appliedRenderViewportScale.md) to monitor what scale the device is using and make informed decisions about how to adjust `XRSettings.renderViewportScale` for subsequent frames.

#### Use XRSettings.renderViewPortScale with post-processing enabled

If you enable post-processing methods in your URP XR project, URP renders to intermediate textures rather than directly to the eye texture.

To enable dynamic resolution with intermediate textures:

1. Open your Main Camera in the **Inspector**. Under **Output**, enable **URP Dynamic Resolution**.
2. Use [XRSettings.renderViewportScale](../ScriptReference/XR.XRSettings-renderViewportScale.md) to adjust the scale of the final area of the viewport sent to the device (the last blit viewport). URP internally uses the [ScalableBufferManager](../ScriptReference/ScalableBufferManager.md) to scale the intermediate textures.

**Note:** You can’t use this approach with [fixed foveated rendering](xr-foveated-rendering.md), or if you have enabled [Temporal Anti Aliasing](urp/anti-aliasing.md).

### Change render scale in the URP Asset

Change the URP render scale directly in your URP Asset to scale the eye texture directly. Reducing the value of the URP render scale can reduce the resolution. This method is an alternative to [XRSettings.eyeTextureResolutionScale](../ScriptReference/XR.XRSettings-eyeTextureResolutionScale.md), which isn’t supported in URP but is supported in HDRP.

**Important**: Changing the render scale reallocates the eye texture, which is an expensive operation. You should not perform this operation every frame.

To change the render scale of your URP Asset:

1. Open your URP Asset in the ****Inspector**** window.
2. Expand the **Quality** section.
3. Set desired the **Render Scale** value.

You can also change this value with the [UniversalRenderPipeline.asset.renderScale](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@17.3/api/UnityEngine.Rendering.Universal.UniversalRenderPipelineAsset.html) API.

### Enable automatic viewport dynamic resolution (OpenXR)

In Unity 6.3 and newer with OpenXR 1.16 and newer, you can use the Automatic viewport dynamic resolution feature to automatically control resolution on OpenXR runtimes. To learn more, refer to [Automatic viewport dynamic resolution](https://docs.unity3d.com/Packages/com.unity.xr.openxr@1.16/manual/features/automaticdynamicresolution.html) (OpenXR package documentation).

## Control resolution with HDRP

If your project uses HDRP, you can use the dynamic resolution system or change the `XRSettings.eyeTextureResolutionScale` as outlined in [Configure HDRP for virtual reality](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.3/manual/configure-hdrp-for-virtual-reality.html).

**Note**: `XRSettings.renderViewportScale` isn’t supported on HDRP.

## Additional resources

* [Dynamic resolution](DynamicResolution-landing.md)
* [Optimize for untethered XR devices in URP](xr-untethered-device-optimization.md)