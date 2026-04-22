# Scalers reference

Explore the built-in scalers provided by Adaptive Performance to automatically manage your project’s settings at runtime.

## General performance scalers

This section details the general-purpose scalers and their default settings:

| **General scalers** | **Min Scale** | **Max Scale** | **Max Level** | **Visual Impact** | **Target** | **Setting scaled** |
| --- | --- | --- | --- | --- | --- | --- |
| **AdaptiveLOD** | 0.4 | 1 | 3 | High | GPU | [`QualitySettings.lodBias`](../../ScriptReference/QualitySettings-lodBias.md) |
| **AdaptiveResolution** | 0.5 | 1 | 9 | Low | GPU/FillRate | [`AdaptivePerformanceRenderSettings.RenderScaleMultiplier`](../../ScriptReference/AdaptivePerformance.AdaptivePerformanceRenderSettings.RenderScaleMultiplier.md)/[`ScalableBuffers`](../../ScriptReference/ScalableBufferManager.md) |
| **AdaptiveFramerate** | 15 | 60 | 45 | High | CPU/GPU/FillRate | [`Application.targetFrameRate`](../../ScriptReference/Application-targetFrameRate.md) |
| **AdaptiveViewDistance** | 50 | 1,000 | 40 | High | GPU | [`Camera.main.farClipPlane`](../../ScriptReference/Camera-farClipPlane.md) |
| **AdaptivePhysics** | 0.5 | 1 | 5 | Low | CPU | [`Time.fixedDeltaTime`](../../ScriptReference/Time-fixedDeltaTime.md) |
| **AdaptiveLayerCulling** | 0.01 | 1 | 40 | Medium | CPU | [`Camera.main.layerCullDistances`](../../ScriptReference/Camera-layerCullDistances.md) |

## Universal Render Pipeline (URP) scalers

Adaptive Performance provides the following scalers to control URP-specific rendering features:

| **URP scalers** | **Min Scale** | **Max Scale** | **Max Level** | **Visual Impact** | **Target** | **Setting scaled** |
| --- | --- | --- | --- | --- | --- | --- |
| **AdaptiveBatching** | 0 | 1 | 1 | Medium | CPU | [`AdaptivePerformanceRenderSettings.SkipDynamicBatching`](../../ScriptReference/AdaptivePerformance.AdaptivePerformanceRenderSettings.SkipDynamicBatching.md) |
| **AdaptiveLUT** | 0 | 1 | 1 | Medium | CPU/GPU | [`AdaptivePerformanceRenderSettings.LutBias`](../../ScriptReference/AdaptivePerformance.AdaptivePerformanceRenderSettings.LutBias.md) |
| **AdaptiveMSAA** | 0 | 1 | 2 | Medium | GPU/FillRate | [`AdaptivePerformanceRenderSettings.AntiAliasingQualityBias`](../../ScriptReference/AdaptivePerformance.AdaptivePerformanceRenderSettings.AntiAliasingQualityBias.md) |
| **AdaptiveShadowCascade** | 0 | 1 | 2 | Medium | CPU/GPU | [`AdaptivePerformanceRenderSettings.MainLightShadowCascadesCountBias`](../../ScriptReference/AdaptivePerformance.AdaptivePerformanceRenderSettings.MainLightShadowCascadesCountBias.md) |
| **AdaptiveShadowDistance** | 0.15 | 1 | 3 | Low | GPU | [`AdaptivePerformanceRenderSettings.MaxShadowDistanceMultiplier`](../../ScriptReference/AdaptivePerformance.AdaptivePerformanceRenderSettings.MaxShadowDistanceMultiplier.md) |
| **AdaptiveShadowQuality** | 0 | 1 | 3 | High | CPU/GPU | [`AdaptivePerformanceRenderSettings.ShadowQualityBias`](../../ScriptReference/AdaptivePerformance.AdaptivePerformanceRenderSettings.ShadowQualityBias.md) |
| **AdaptiveShadowmapResolution** | 0.15 | 1 | 3 | Low | GPU | [`AdaptivePerformanceRenderSettings.MainLightShadowmapResolutionMultiplier`](../../ScriptReference/AdaptivePerformance.AdaptivePerformanceRenderSettings.MainLightShadowmapResolutionMultiplier.md) |
| **AdaptiveSorting** | 0 | 1 | 1 | Medium | CPU | [`AdaptivePerformanceRenderSettings.SkipFrontToBackSorting`](../../ScriptReference/AdaptivePerformance.AdaptivePerformanceRenderSettings.SkipFrontToBackSorting.md) |
| **AdaptiveTransparency** | 0 | 1 | 1 | High | GPU | [`AdaptivePerformanceRenderSettings.SkipTransparentObjects`](../../ScriptReference/AdaptivePerformance.AdaptivePerformanceRenderSettings.SkipTransparentObjects.md) |
| **AdaptiveDecals** | 0.01 | 1 | 20 | Medium | GPU | [`AdaptivePerformanceRenderSettings.DecalsDrawDistance`](../../ScriptReference/AdaptivePerformance.AdaptivePerformanceRenderSettings.DecalsDrawDistance.md) |

### How scalers map to URP settings

This table maps each `AdaptivePerformanceRenderSettings` property to the specific URP feature it modifies at runtime.

This list covers every internal Unity Editor setting affected by the Adaptive Performance built-in scalers. Most scalers apply their calculated scale directly, while others adjust settings in more specialized ways when required.

| **Adaptive Performance scaler setting** | **URP setting** |
| --- | --- |
| `AdaptivePerformanceRenderSettings.SkipDynamicBatching` | [`RenderingData.supportsDynamicBatching`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@latest?subfolder=/api/UnityEngine.Rendering.Universal.RenderingData.html#UnityEngine_Rendering_Universal_RenderingData_supportsDynamicBatching) |
| `AdaptivePerformanceRenderSettings.LutBias` | [`UniversalPostProcessingData.lutSize`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@latest?subfolder=/api/UnityEngine.Rendering.Universal.UniversalPostProcessingData.html?q=UniversalPostProcessingData) |
| `AdaptivePerformanceRenderSettings.AntiAliasingQualityBias` | [`CameraData.antialiasingQuality`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@latest?subfolder=/api/UnityEngine.Rendering.Universal.CameraData.html#UnityEngine_Rendering_Universal_CameraData_antialiasingQuality) |
| `AdaptivePerformanceRenderSettings.MainLightShadowCascadesCountBias` | [`ShadowData.mainLightShadowCascadesCount`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@latest?subfolder=/api/UnityEngine.Rendering.Universal.ShadowData.html#UnityEngine_Rendering_Universal_ShadowData_mainLightShadowCascadesCount) |
| `AdaptivePerformanceRenderSettings.MaxShadowDistanceMultiplier` | [`CameraData.maxShadowDistance`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@latest?subfolder=/api/UnityEngine.Rendering.Universal.CameraData.html#UnityEngine_Rendering_Universal_CameraData_maxShadowDistance) |
| `AdaptivePerformanceRenderSettings.ShadowQualityBias` | [`ShadowData.supportsSoftShadows ShadowData.supportsAdditionalLightShadows ShadowData.supportsMainLightShadows`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@latest?subfolder=/api/UnityEngine.Rendering.Universal.ShadowData.html) |
| `AdaptivePerformanceRenderSettings.MainLightShadowmapResolutionMultiplier` | [`ShadowData.mainLightShadowmapWidth ShadowData.mainLightShadowmapHeight`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@latest?subfolder=/api/UnityEngine.Rendering.Universal.ShadowData.html#UnityEngine_Rendering_Universal_ShadowData_mainLightShadowmapWidth) |
| `AdaptivePerformanceRenderSettings.SkipFrontToBackSorting` | [`cameraData.defaultOpaqueSortFlags`](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@latest?subfolder=/api/UnityEngine.Rendering.Universal.CameraData.html#UnityEngine_Rendering_Universal_CameraData_defaultOpaqueSortFlags) |
| `AdaptivePerformanceRenderSettings.SkipTransparentObjects` | Disables the forward transparency pass in the URP, which means transparent objects in the **scene** don’t display. |
| `AdaptivePerformanceRenderSettings.DecalsDrawDistance` | Scales [Max Draw Distance](../urp/renderer-feature-decal-reference.md#max-draw-distance) |

## Additional resources

* [Modifying asset quality with scalers](scalers-introduction.md)
* [Create custom scalers](create-custom-scalers.md)
* [Use scaler profiles](scaler-profiles.md)