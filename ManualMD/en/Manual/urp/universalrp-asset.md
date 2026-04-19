# Universal Render Pipeline asset reference for URP

In the URP Asset, you can configure settings for:

* [**Rendering**](#rendering)
* [**Quality**](#quality)
* [**Lighting**](#lighting)
* [**Shadows**](#shadows)
* [**Post-processing**](#post-processing)
* [**Volumes**](#volumes)
* [**Adaptive Performance**](#adaptive-performance)

**Note:** If you have the experimental 2D Renderer enabled (menu: **Graphics Settings** > add the 2D Renderer Asset under **Scriptable Render Pipeline Settings**), some of the options related to 3D rendering in the URP Asset don’t have any impact on your final app or game.

## Rendering

The **Rendering** settings control the core part of the pipeline rendered frame.

| Property | Description |
| --- | --- |
| **Depth Texture** | Enables URP to create a `_CameraDepthTexture`. URP then uses this [depth texture](https://docs.unity3d.com/Manual/SL-DepthTextures.html) by default for all **Cameras** in your **scene**. You can override this for individual cameras in the [Camera Inspector](camera-component-reference.md). |
| **Opaque Texture** | Enable this to create a `_CameraOpaqueTexture` as default for all cameras in your scene. This works like the [GrabPass](https://docs.unity3d.com/Manual/SL-GrabPass.html) in the built-in render pipeline.  The **Opaque Texture** provides a snapshot of the scene right before URP renders any transparent meshes. You can use this in transparent Shaders to create effects like frosted glass, water refraction, or heat waves. You can override this for individual cameras in the [Camera Inspector](camera-component-reference.md). |
| **Opaque Downsampling** | Set the sampling mode on the opaque texture to one of the following:  * **None**: Produces a copy of the opaque pass in the same resolution as the camera. * **2x Bilinear**: Produces a half-resolution image with bilinear filtering. * **4x Box**: Produces a quarter-resolution image with box filtering. This produces a softly blurred copy. * **4x Bilinear**: Produces a quarter-resolution image with bi-linear filtering. |
| ****Terrain** Holes** | If you disable this option, the URP removes all Terrain hole **Shader** variants when you build for the Unity Player, which decreases build time. |
| **GPU Resident Drawer** | The GPU Resident Drawer automatically uses the [`BatchRendererGroup`](https://docs.unity3d.com/Manual/batch-renderer-group.html) API to draw GameObjects with GPU instancing. For more information, refer to [Use the GPU Resident Drawer](gpu-resident-drawer.md).  Available options:  * **Disabled**: Unity doesn’t automatically draw GameObjects with GPU instancing. * **Instanced Drawing**: Unity automatically draws GameObjects with GPU instancing. |
| **Small-Mesh Screen-Percentage** | Set the screen percentage Unity uses to cull small GameObjects, to speed up rendering. Unity culls GameObjects that fill less of the screen than this value.  This setting might not work if you use your own [Level of Detail (LOD) meshes](https://docs.unity3d.com/Manual/LevelOfDetail.html).  Set the value to 0 to stop Unity culling small GameObjects.  To prevent Unity culling an individual GameObject that covers less screen space than this value, go to the **Inspector** window for the GameObject and add a **Disallow Small Mesh Culling** component. |
| **GPU **Occlusion Culling**** | Enable Unity to use the GPU instead of the CPU to exclude **GameObjects** from rendering when they’re hidden behind other GameObjects. Refer to [Use GPU occlusion culling](gpu-culling.md) for more information. |
| **SRP Batcher** | Enable the SRP Batcher. This is useful if you have many different Materials that use the same Shader. The SRP Batcher is an inner loop that speeds up CPU rendering without affecting GPU performance. When you use the SRP Batcher, it replaces the SRP rendering code inner loop.  If both **SRP Batcher** and **Dynamic Batching** are enabled, SRP Batcher will take precedence over dynamic batching as long as the shader is SRP Batcher compatible.  **Note**: If assets or shaders in a project are not optimized for use with the SRP Batcher, low performance devices might be more performant when you disable the SRP Batcher. |
| **Dynamic Batching** | Enable [Dynamic Batching](https://docs.unity3d.com/Manual/DrawCallBatching.html), to make the render pipeline automatically batch small dynamic objects that share the same Material. This is useful for platforms and graphics APIs that do not support GPU instancing.  If your targeted hardware does support GPU instancing, disable **Dynamic Batching**. You can change this at run time. |
| **Debug Level** | Set the level of debug information that the render pipeline generates.  Available options:  * **Disabled**: Debugging is disabled. This is the default. * **Profiling**: Makes the render pipeline provide detailed information tags, which you can find in the FrameDebugger. |
| **Shader Variant Log Level** | Set the level of information about Shader Stripping and Shader Variants you want to display when Unity finishes a build.  Available options:  * **Disabled**: Unity doesn’t log anything. * **Only Universal**: Unity logs information for all of the [URP Shaders](shaders-in-universalrp.md). * **All**: Unity logs information for all Shaders in your build.   You can check the information in Console panel when your build has finished. |
| **Store Actions** | Defines if Unity discards or stores the render targets of the DrawObjects Passes.  Available options:  * **Auto**: Unity uses the **Discard** option by default, and falls back to the **Store** option if it detects any injected Passes. * **Discard**: Unity discards the render targets of render Passes that are not reused later (lower memory bandwidth). * **Store**: Unity stores all render targets of each Pass. **Store** significantly increases the memory bandwidth on mobile and tile-based GPUs. |

## Quality

These settings control the quality level of the URP. This is where you can make performance better on lower-end hardware or make graphics look better on higher-end hardware.

**Tip:** If you want to have different settings for different hardware, you can configure these settings across multiple Universal **Render Pipeline** assets, and switch them out as needed.

| **Property** | **Subproperty** | **Description** |
| --- | --- | --- |
| **HDR** | N/A | Enable this to allow rendering in High Dynamic Range (HDR) by default for every camera in your scene. With HDR, the brightest part of the image can be greater than 1.  This gives you a wider range of light intensities, so your lighting looks more realistic, such as being able to pick out details and experience less saturation even with bright light. This is useful if you want a wide range of lighting or to use [bloom](https://docs.unity3d.com/Manual/PostProcessing-Bloom.html) effects.   If you’re targeting lower-end hardware, you can disable this to skip HDR calculations and get better performance. You can override this for individual cameras in the Camera Inspector. |
| N/A | **HDR Precision** | The precision of the Camera color buffer in HDR rendering. The 64 bit precision lets you avoid banding artifacts, but requires higher bandwidth and might make sampling slower. Default value: 32 bit.  This property is available only when you enable [Advanced Properties](urp-asset-additional-settings.md). |
| **Anti Aliasing (MSAA)** | N/A | Use [Multisample Anti-aliasing](anti-aliasing.md#msaa) by default for every Camera in your scene while rendering. This softens edges of your geometry, so they’re not jagged or flickering. In the drop-down menu, select how many samples to use per pixel: **2x**, **4x**, or **8x**. The more samples you choose, the smoother your object edges are.  If you want to skip MSAA calculations, or you don’t need them in a 2D game, select **Disabled**. You can override this for individual cameras in the Camera Inspector.  **Note**: On mobile platforms that do not support the [StoreAndResolve](https://docs.unity3d.com/ScriptReference/Rendering.RenderBufferStoreAction.StoreAndResolve.html) store action, if **Opaque Texture** is selected in the URP asset, Unity ignores the **Anti Aliasing (MSAA)** property at runtime (as if Anti Aliasing (MSAA) is set to Disabled). |
| **Render Scale** | N/A | This slider scales the render target resolution (not the resolution of your current device). Use this when you want to render at a smaller resolution for performance reasons or to upscale rendering to improve quality.  **Note**: This only scales the game rendering. UI rendering is left at the native resolution for the device. |
| **Upscaling Filter** | N/A | Select which image filter Unity uses when performing the upscaling. Unity performs upscaling when the Render Scale value is less than 1.0. |
| N/A | **Automatic** | Unity selects one of the filtering options based on the Render Scale value and the current screen resolution. If integer scaling is possible, Unity selects the Nearest-Neighbor option, otherwise Unity selects the Bilinear option. |
| N/A | **Bilinear** | Unity uses the bilinear or linear filtering provided by the graphics API. |
| N/A | **Nearest-Neighbor** | Unity uses the nearest-neighbor or point sampling filtering provided by the graphics API.  **Note**: The **Nearest-Neighbour** filter doesn’t work if **Post-processing** is enabled. |
| N/A | **FidelityFX Super Resolution 1.0** | Unity uses the AMD FidelityFX Super Resolution 1.0 (FSR) technique to perform upscaling.  Unlike most other Upscaling Filter options, this filter remains active even at a Render Scale value of 1.0. This filter can still improve image quality even when no scaling is occurring. This also makes the transition between scale values 0.99 and 1.0 less noticeable in cases where dynamic resolution scaling is active.  **Note**: This filter is only supported on devices that support Unity shader model 4.5 or higher. On devices that do not support Unity shader model 4.5, Unity uses the **Automatic** option instead. |
| N/A | **Override FSR Sharpness** | Unity shows this check box when you select the FSR filter. Selecting this check box lets you specify the intensity of the FSR sharpening pass. |
| N/A | **FSR Sharpness** | Specify the intensity of the FSR sharpening pass. A value of 0.0 provides no sharpening, a value of 1.0 provides maximum sharpness.  **Note**: This option has no effect when FSR is not the active upscaling filter. |
| N/A | **Spatial-Temporal Post-Processing (STP) 1.0** | Uses the Spatial Temporal Post-Processing (STP) technique to perform upscaling. Selecting this option forces the **Anti-Aliasing** setting in the [Camera Inspector window](camera-component-reference.md) to **Temporal Anti-aliasing (TAA)**.  This setting improves image quality even without scaling, so it remains active when **Render Scale** is set to 1.0.  For more information on STP, refer to [Spatial-Temporal Post-processing](stp/stp-upscaler.md).  **Note**: This setting is supported only on non-GLES devices that support compute shaders. On unsupported devices, Unity uses **Automatic** instead. |
| ****LOD** Cross Fade** | N/A | Use this property to enable or disable the LOD cross-fade. If you disable this option, URP removes all LOD cross-fade shader variants when you build the Unity Player, which decreases the build time. |
| **LOD Cross Fade Dithering Type** | N/A | When a [LOD group](https://docs.unity3d.com/Manual/class-LODGroup.html) has **Fade Mode** set to **Cross Fade**, Unity renders the Renderer’s LOD meshes with cross-fade blending between them using either alpha testing or stencil testing. This property defines the dithering pattern of LOD cross-fade.  Available options:  * **Bayer Matrix**: better performance than the Blue Noise option, but has a repetitive pattern. * **Blue Noise**: uses a precomputed blue noise texture and provides a better look than the Bayer Matrix option, but has a slightly higher performance cost. * **2x2 Stencil**: Unity uses bits 4 and 8 in the stencil buffer to apply a 2 × 2 dithering pattern, instead of alpha testing. This setting significantly decreases the number of shader variants. For more information, refer to [Reducing shader variants in URP](shader-stripping-landing.md). |

## Lighting

These settings affect the lights in your scene.

If you disable some of these settings, the relevant [keywords](https://docs.unity3d.com/Manual/shader-keywords) are [stripped from the shader variables](shader-stripping-landing.md). If there are settings that you know for certain you won’t use in your game or app, you can disable them to improve performance and reduce build time.

| **Property** | **Subproperty** | **Description** |
| --- | --- | --- |
| **Main Light** | N/A | These settings affect the main [Directional Light](https://docs.unity3d.com/Manual/Lighting.html) in your scene. You can select this by assigning it as a [Sun Source](https://docs.unity3d.com/Manual/GlobalIllumination.html) in the Lighting Inspector. If you don’t assign a sun source, the URP treats the brightest directional light in the scene as the main light.  You can choose between [Pixel Lighting](https://docs.unity3d.com/Manual/LightPerformance.html) and **None**. If you choose None, URP doesn’t render a main light, even if you’ve set a sun source. |
| N/A | **Cast Shadows** | Check this box to make the main light cast shadows in your scene. |
| N/A | **Shadow Resolution** | This controls how large the shadow map texture for the main light is. High resolutions give sharper, more detailed shadows. If memory or rendering time is an issue, try a lower resolution. |
| **Light Probe System** | N/A | Select the light probe system this URP Asset uses.  Available options:  * **Light Probe Groups (Legacy)**: Use the same [Light Probe Group system](https://docs.unity3d.com/Manual/class-LightProbeGroup.html) as the Built-In Render Pipeline. * **Adaptive Probe Volumes**: Use [Adaptive Probe Volumes](probevolumes.md). |
| **Memory Budget** | N/A | Limits the width and height of the textures that store baked Global Illumination data, which determines the amount of memory Unity sets aside to store baked Adaptive Probe Volume data. These textures have a fixed depth.  Available options:  * **Memory Budget Low** * **Memory Budget Medium** * **Memory Budget High** |
| **SH Bands** | N/A | Determines the [spherical harmonics (SH) bands](https://docs.unity3d.com/Manual/LightProbes-TechnicalInformation.html) Unity uses to store probe data. L2 provides more precise results, but uses more system resources.  Available options:  * **Spherical Harmonics L1** * **Spherical Harmonics L2** |
| **Enable Streaming** | N/A | Enable to stream Adaptive Probe Volume data from CPU memory to GPU memory at runtime. Refer to [Optimize loading Adaptive Probe Volume data](probevolumes-streaming.md) for more information. |
| **Estimated GPU Memory Cost** | N/A | Indicates the amount of texture data used by Adaptive Probe Volumes in your project. |
| **Additional Lights** | N/A | Here, you can choose to have additional lights to supplement your main light. Choose between [Per Vertex](https://docs.unity3d.com/Manual/LightPerformance.html), [Per Pixel](https://docs.unity3d.com/Manual/LightPerformance.html), or **Disabled**. |
| N/A | **Per Object Limit** | This slider sets the limit for how many additional lights can affect each GameObject. Unity ignores this setting if you select the Forward+ **rendering path**. |
| N/A | **Cast Shadows** | Check this box to make the additional lights cast shadows in your scene. |
| N/A | **Shadow Atlas Resolution** | This controls the size of the textures that cast directional shadows for the additional lights.  This is a sprite atlas that packs up to 16 shadow maps. High resolutions give sharper, more detailed shadows. If memory or rendering time is an issue, try a lower resolution. |
| N/A | **Shadow Resolution Tiers** | Set the resolution of the shadows cast by additional lights at various tiers.  Resolutions must have a value of 128 or greater, and are rounded to the next power of two.  **Note**: This property is only visible when the **Cast Shadows** property is enabled for Additional Lights. |
| N/A | **Cookie Atlas Resolution** | The size of the cookie atlas the additional lights use. All additional lights are packed into a single cookie atlas.  This property is only visible when the **Light Cookies** property is enabled. |
| N/A | **Cookie Atlas Format** | The format of the cookie atlas for additional lights. All additional lights are packed into a single cookie atlas.  Available options:  * **Grayscale Low** * **Grayscale High** * **Color Low** * **Color High** * **Color HDR**  This property is only visible when the **Light Cookies** property is enabled. |
| **Reflection Probes** | N/A | Use these properties to control **reflection probe** settings. |
| N/A | **Probe Blending** | Smooth the transitions between Reflection Probes. For more information, refer to [Reflection Probe Blending](../blend-reflection-probes-birp.md). |
| N/A | **Probe Atlas Blending** | If enabled, reflection probes will be added to the Forward Plus data grid and combined into a single atlas texture. The atlas is used by default when both Forward Plus and the GPU Resident Drawer are used. This property is only applicable if you select the Forward+ rendering path. This property is only visible when the **Probe Blending** property is enabled. |
| N/A | **Box Projection** | Create reflections on objects based on their position within the probe’s box, while still using a single probe as the reflection source. For more information, refer to [Advanced Reflection Probe features](../AdvancedRefProbe.md). |
| **Mixed Lighting** | N/A | Enable [Mixed Lighting](https://docs.unity3d.com/Manual/LightMode-Mixed.html) to configure the pipeline to include mixed lighting shader variants in the build. |
| **Use Rendering Layers** | N/A | With this option selected, you can configure certain Lights to affect only specific GameObjects. For more information on Rendering Layers and how to use them, refer to the documentation on [Rendering Layers](features/rendering-layers.md). |
| **Light Cookies** | N/A | Enables [light cookies](https://docs.unity3d.com/Manual/Cookies.html). This property enables **Cookie Atlas Resolution** and **Cookie Atlas Format** for additional lights. |
| **SH Evaluation Mode** | N/A | Defines the spherical harmonic (SH) lighting evaluation type.  Available options:  * **Auto**: Unity selects a mode automatically. * **Per Vertex**: Evaluate lighting per vertex. * **Mixed**: Evaluate lighting partially per vertex, partially per pixel. * **Per Pixel**: Evaluate lighting per pixel. |

## Shadows

These settings let you configure how shadows look and behave, and find a good balance between the visual quality and performance.

The **Shadows** section has the following properties.

| **Property** | **Subproperty** | **Description** |
| --- | --- | --- |
| **Max Distance** | N/A | The maximum distance from the Camera at which Unity renders the shadows. Unity does not render shadows farther than this distance.  **Note**: This property is in metric units regardless of the value in the **Working Unit** property. |
| **Working Unit** | N/A | The unit in which Unity measures the shadow cascade distances. |
| **Cascade Count** | N/A | The number of [shadow cascades](https://docs.unity3d.com/Manual/shadow-cascades.html). With shadow cascades, you can avoid crude shadows close to the Camera and keep the Shadow Resolution reasonably low.  For more information, refer to the documentation on [shadow cascades](../shadow-cascades.md). Increasing the number of cascades reduces the performance. Cascade settings only affects the main light. |
| N/A | **Split** **1** | The distance where cascade 1 ends and cascade 2 starts. |
| N/A | **Split** **2** | The distance where cascade 2 ends and cascade 3 starts. |
| N/A | **Split** **3** | The distance where cascade 3 ends and cascade 4 starts. |
| N/A | **Last Border** | The size of the area where Unity fades out the shadows. Unity starts fading out shadows at the distance **Max Distance** - **Last Border**, at **Max Distance** the shadows fade to zero. |
| N/A | **Depth Bias** | Use this setting to reduce [shadow acne](https://docs.unity3d.com/Manual/ShadowPerformance.html). |
| N/A | **Normal Bias** | Use this setting to reduce [shadow acne](https://docs.unity3d.com/Manual/ShadowPerformance.html). |
| N/A | **Soft Shadows** | Select this check box to enable extra processing of the shadow maps to give them a smoother look. **Performance impact**: High impact on platforms that use tile-based rendering, such as mobile platforms and untethered XR platforms. When this option is disabled, Unity samples the shadow map once with the default hardware filtering. |
| N/A | **Quality** | Select the quality level of soft shadow processing.  Available options:  * **Low**: good balance of quality and performance for mobile platforms. Filtering method: 4 PCF taps. * **Medium**: good balance of quality and performance for desktop platforms. Filtering method: 5x5 tent filter. This is the default value. * **High**: best quality, higher performance impact. Filtering method: 7x7 tent filter. |
| **Conservative Enclosing Sphere** | N/A | Enable this option to improve shadow frustum culling and prevent Unity from excessively culling shadows in the corners of the shadow cascades.  Disable this option only for compatibility purposes of existing projects created in previous Unity versions.  If you enable this option in an existing project, you might need to adjust the shadows cascade distances because the shadow culling enclosing spheres change their size and position.  **Performance impact**: Enabling this option is likely to improve performance, because the option minimizes the overlap of shadow cascades, which reduces the number of redundant static shadow casters. |

## Post-processing

This section allows you to fine-tune global **post-processing** settings.

| **Property** | **Description** |
| --- | --- |
| **Grading Mode** | Select the [color grading](https://docs.unity3d.com/Manual/PostProcessing-ColorGrading.html) mode to use for the Project.  * **High Dynamic Range**: This mode works best for high precision grading similar to movie production workflows. Unity applies color grading before tonemapping. * **Low Dynamic Range**: This mode follows a more classic workflow. Unity applies a limited range of color grading after tonemapping. |
| **LUT Size** | Set the size of the internal and external [look-up textures (LUTs)](https://docs.unity3d.com/Manual/PostProcessing-ColorGrading.html) that the Universal Render Pipeline uses for color grading. Higher sizes provide more precision, but have a potential cost of performance and memory use. You cannot mix and match LUT sizes, so decide on a size before you start the color grading process.  The default value, **32**, provides a good balance of speed and quality. |
| **Alpha Processing** | With this setting enabled, URP post-processing effects output properly processed alpha values. With this setting disabled, URP discards the alpha channel by replacing alpha values with 1. The render target requires a format with the alpha channel. If you use HDR rendering, set the **HDR Precision** property to 64 bit, because the 32 bit format does not have the alpha channel.  **Rendering into a render texture** If you are rendering the output with the alpha channel into a render texture, ensure that the **Color Format** property of the render texture has the alpha channel. On the camera that renders output with the alpha values, set the **Background Type** property in the **Environment** section to **Solid Color**. This lets you identify and process the alpha values in a shader.  **Limitations**  * In a setup with camera stacking, post-processing effects on an overlay camera still affect all cameras below it. This setting lets you configure different post-processing effects for separate cameras (not in the same camera stack) and use render textures and a compositing pass to combine them. * When applying post-processing effects, this feature preserves the alpha values as they were before applying an effect. As a consequence, pre-built post-processing effects that draw pixels beyond the original borders of objects (for example, bloom or depth of field effects) might render with sharp edges around objects they are affecting. This does not apply to effects that distort geometry, like the panini projection or the lens distortion effect. Those effects also distort the alpha channel. |
| **Fast sRGB/Linear Conversions** | Select this option to use faster, but less accurate approximation functions when converting between the sRGB and Linear color spaces. |
| **Data Driven Lens Flare** | Allocate the shader variants and memory URP needs for [lens flares](shared/lens-flare/lens-flare-srp-reference.md) effect. |
| **Screen Space Lens Flare** | Allocate the shader variants and memory URP needs for [screen space lens flares](shared/lens-flare/reference-screen-space-lens-flare.md). |

## Volumes

| Property | Description |
| --- | --- |
| **Volume Update Mode** | Select how Unity updates Volumes at run time.  * **Every Frame**: Unity updates volumes every frame. * **Via Scripting**: Unity updates volumes when triggered via scripting.  In the Editor, Unity updates Volumes every frame when not in Play mode. |
| **Volume Profile** | Set the [Volume Profile](Volume-Profile.md) that a scene uses by default.  Refer to [Understand volumes](Volumes.md) for more information. |

The list of Volume Overrides that the Volume Profile contains appears below **Volume Profile**. You can add, remove, disable, and enable Volume Overrides, and edit their properties. Refer to [Volume Overrides](VolumeOverrides.md) for more information.

## Adaptive Performance

This section is available if [Adaptive Performance](../adaptive-performance/adaptive-performance.md) is enabled in the project. The **Use Adaptive Performance** property lets you enable the Adaptive Performance functionality.

| Property | Description |
| --- | --- |
| **Use Adaptive Performance** | Select this check box to enable the Adaptive Performance functionality, which adjusts the rendering quality at runtime. |

## Additional resources

[Create a URP Asset](urp-asset-create.md)