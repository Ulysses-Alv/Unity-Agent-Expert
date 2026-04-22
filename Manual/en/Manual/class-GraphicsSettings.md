# Graphics

Use the **Graphics** settings (main menu: **Edit** > **Project Settings**, then select the **Graphics** category) to apply global settings for Graphics.

**Note**: You can access **Graphics** settings from the **Build Profiles** window (menu: **File** > **Build Profiles**). With [build profiles](build-profiles.md), you can customize the Graphics settings per build profile to set different values for each platform. For more information, refer to [Override settings with build profiles](build-profiles-override-settings.md).

When you build your project, the values of these settings become static. You can’t change the settings at runtime.

This section provides documentation on the following groups of properties:

* [Set Default Render Pipeline Asset](#SRLoop)
* [Shader Build Settings](#shader-build-settings)
* [Shader stripping](#stripping)
* [Shader loading](#shader-loading)
* [Culling settings](#culling-settings)
* [Pipeline Specific Settings](#pipeline-specific-settings)
  + [Camera Settings](#Camera)
  + [Tier Settings](#Tier)
  + [Built-in shader settings](#Built-in)
  + [Always-included Shaders](#Always)

## Set Default Render Pipeline Asset

Use the **Default **Render Pipeline**** field to set the default render pipeline Unity uses to render your project. For example, set **Default Render Pipeline** to a Universal Render Pipeline (URP) Asset to use URP to render your project.

If **Default Render Pipeline** is set to **None**, Unity uses the Built-In Render Pipeline.

For more information, refer to the following:

* [Pipeline Specific Settings](#pipeline-specific-settings)
* [How to get, set, and configure the active render pipeline](srp-setting-render-pipeline-asset.md)

## Shader Build Settings

These properties allow you to [reduce shader variants](shader-variant-stripping.md) in your build.

**Important**: When you remove built-in shader keywords, Unity might strip shader variants your built project needs. To check how many shader variants your project has, enable strict shader variant matching, otherwise Unity might replace missing shader variants with others. For more information, refer to [Check how many shader variants you have](shader-how-many-variants.md#highlight-missing-shaders).

Open the **Keyword Declaration Overrides** dropdown to add the keyword sets you want to set to a different keyword type or remove.

| **Property** | **Description** |
| --- | --- |
| **Keyword Set** | The keyword set you want to set to a different keyword type or remove. Use spaces to separate keywords in the set, for example `FOG_LINEAR FOG_EXP FOG_EXP2`. Enter all the keywords in the set, including `_` if it exists.  To change the keyword type, use the **Type Override** property. To exclude a keyword in the build, select the foldout (triangle) then disable the checkbox for the keyword.  For more information, refer to [reduce shader variants](shader-variant-stripping.md). |
| **Type override** | The keyword type to use for the keyword set. The options are:  * **Default**: Uses the keyword type the shader code uses. * **shader\_feature**: Compiles shader variants for keyword combinations that materials in your build use, and removes other shader variants. Use this option if the keywords are material properties. * **multi\_compile**: Compiles shader variants regardless of whether they’re used by materials in your build. Use this option if you need to use a C# script to change the state of the keywords at runtime. * **dynamic\_branch**: Uses dynamic branching for conditional code that uses these keywords, and doesn’t compile shader variants. Use this option if your shaders run on a fast GPU, and don’t have asymmetric code branches where one branch is longer or more complex than the other.  For more information, refer to [How Unity compiles branching shaders](shader-conditionals-choose-a-type.md). |
| **Selected keywords** | The number of keywords you select from the keyword set to be included in the build, out of the total number of keywords in the set. |
| **Revert** | Cancel your changes. |
| **Apply** | Apply your changes. |

## Shader stripping

These properties allow you to configure [shader variant stripping](shader-variant-stripping.md) in your build.

By default, Unity examines the **scenes** in the build and automatically strips [shader variants](shader-variants.md) that are not used in those scenes. However, this can cause problems if you want to use shaders or variants at runtime that would not otherwise be included in the build. For example, if you use AssetBundles or Addressables that rely on those shaders or variants, or if you use [shader keywords](shader-keywords.md) to change variants at runtime.

| **Property** | **Function** |
| --- | --- |
| **Lightmap Modes** | Determine the shader variant stripping behavior for lightmap-related shaders. The options are:  * **Automatic**: Unity examines the scenes in the build and automatically strips variants that are not used in those scenes. This is the default setting. * **Custom**: Select this option to manually include or exclude variants for the following lightmap modes:   + **Baked Non-Directional**   + **Baked Directional**.   + **Realtime Non-Directional**   + **Realtime Directional**   + **Baked Shadowmask**   + **Baked Subtractive** |
| **Fog Modes** | Determine the shader variant stripping behavior for shaders that relate to Unity’s built-in fog effect. The options are:  * **Automatic**: Unity examines the scenes in the build and automatically strips variants that are not used in those scenes. This is the default setting. * **Custom**: Select this option to manually include or exclude variants for the following fog modes:   + **Linear**   + **Exponential**   + **Exponential Squared** |
| **Instancing Variants** | Determine the shader variant stripping behavior for shaders that relate to GPU instancing. The options are:  * **Strip Unused**: Unity only includes variants for GPU instancing for a given shader if at least one material that uses that shader has **Enable instancing** enabled. This is the default setting. * **Strip All**: Strip all variants for GPU instancing, even if they are being used. * **Keep All**: Include all variants for GPU instancing, even if they are not being used. |

## Shader loading

These properties allow you to configure how Unity prewarms [shader variant collections](shader-variant-collections.md) in your application.

For information on prewarming, including important information about graphics API support, refer to [Shader loading: Prewarming shader variants](shader-prewarm.md).

| **Property** | **Function** |
| --- | --- |
| **Renderer Light Probe Selection** | Choose the type of probe Unity uses when a Renderer receives global illumination from Light Probes but is not within the volume of influence (the tetrahedron) of any group of Light Probes. Options:  * **Find closest Light Probe**: the default option. With this option selected, Unity searches for the nearest Light Probe, which consumes significant system resources. * **Use Ambient Probe**: Select this option to make Unity fall back to the [Ambient Probe](../ScriptReference/RenderSettings-ambientProbe.md) in situations described above. This option uses less system resources. |
| **Preloaded shaders** | The shader variant collections to prewarm on application start. |
| **Preload shaders after showing first scene** | This property determines when Unity prewarms the shader variants specified in **Preloaded shaders**.  If enabled, Unity loads and prewarms the variants after the first scene has fully loaded. Otherwise, Unity loads and prewarms them before showing the first scene. |
| **Preloaded time limit per frame (ms)** | Determines how Unity prewarms the shader variants specified in **Preloaded shaders** if **Preload shaders after showing first scene** is enabled.  If the value is 0, Unity preloads all shader variants on the next frame after it displays the first scene.  Otherwise, Unity preloads new shader variants over multiple frames after it displays the first scene. In a given frame, Unity prewarms shader variants until it reaches the time limit set in this property, then doesn’t begin prewarming new shader variants until the next frame. |

### Tracked shader variants

The Unity Editor can track which [shader variants](shader-variants.md) your application uses when it runs. You can use this information to build [shader variant collections](shader-variant-collections.md).

| **Control** | **Function** |
| --- | --- |
| **Create asset** | Creates a new shader variant collection asset using the currently tracked shader variants. |
| **Clear** | Clear tracked shader variants. |

## Culling settings

**Camera-Relative Culling** determines whether Unity uses the **camera** position as the reference point for culling.

| **Property** | **Function** |
| --- | --- |
| **Lights** | Use the camera position as the reference point to cull lights instead of the world space origin. Enable **Lights** to reduce flickering if lights are far away from the camera. Refer to [Understanding the View Frustum](UnderstandingFrustum.md) for more information. |
| **Shadows** | Use the camera position as the reference point to cull shadows instead of the world space origin. Enable **Shadows** to reduce flickering if shadows are far away from the camera. Refer to [Understanding the View Frustum](UnderstandingFrustum.md) for more information. |

## Pipeline-specific settings

The **Graphics** settings window always contains the Built-in Render Pipeline settings below.

However if your project uses a Scriptable Render Pipeline (SRP), for example the Universal Render Pipeline (URP) or the High Definition Render Pipeline (HDRP), the window contains a tab for each SRP in your project. Refer to the following pages for more information:

* [URP Graphics Settings](urp/urp-global-settings.md)
* [HDRP Graphics Settings window reference](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@17.0/manual/Default-Settings-Window.html)

### Camera Settings

These properties control various rendering settings.

| **Property** | **Function** |
| --- | --- |
| **Transparency Sort Mode** | Define the order for rendering objects by their distance along a specific axis. Renderers in Unity are sorted by several criteria, such as their layer number or their distance from the camera. This is generally only useful in 2D development: for example, sorting Sprites by height or along the Y-axis. The options are:  * **Default**: Sort objects based on the Camera mode. * **Perspective**: Sort objects based on perspective view. * **Orthographic**: Sort objects based on orthographic view. * **Custom Axis**: Sort objects based on the sort mode defined with the **Transparency Sort Axis**. |
| **Transparency Sort Axis** | Define a custom **Transparency Sort Mode**. |

### Tier Settings

In the Built-in Render Pipeline, you can use **Tier settings** to change rendering and shader compilation settings for different types of hardware. For more information, see [Graphics tiers](graphics-tiers.md).

| **Property** | **Function** |
| --- | --- |
| **Standard Shader Quality** | Set the quality of the [Standard Shader](shader-StandardShader.md) to *High*, *Medium*, or *Low*. |
| **Reflection Probes Box Projection** | Enable projection for reflection UV mappings on [Reflection Probes](class-ReflectionProbe.md). |
| **Reflection Probes Blending** | Enable [blending on Reflection Probes](UsingReflectionProbes.md). |
| **Detail Normal Map** | Enable [Detail Normal Map](StandardShaderMaterialParameterDetail.md) sampling, if assigned. |
| **Enable Semitransparent Shadows** | Enable Semitransparent Shadows. This adds or removes the UNITY\_USE\_DITHER\_MASK\_FOR\_ALPHABLENDED\_SHADOWS shader compiler define. |
| **Enable Light Probe Proxy Volume** | Enable rendering a 3D grid of interpolated [Light Probes](class-LightProbeProxyVolume.md). |
| **Cascaded Shadows** | Enable using cascaded shadow maps. This adds or removes the UNITY\_NO\_SCREENSPACE\_SHADOWS shader compiler define. |
| **Prefer 32 bit shadow maps** | Enable 32-bit float shadow map when you are targeting platforms that use DX11 or DX12. Most platforms have a fixed shadow map format that you can’t adjust. These vary in format, and can be 16-bit, 24-bit, or 32-bit, and can also be either float- or integer-based. 32-bit shadow maps give higher quality shadows than 16-bit, but use increased memory and bandwidth on the GPU. **Note:** To use 32-bit shadow maps, make sure the **depth buffer** is also set to 32-bit. |
| **Use HDR** | Enable [High Dynamic Range rendering](hdr-landing.md) for this tier. |
| **HDR Mode** | Select the format to use for the HDR buffer when **HDR** is enabled for the current Graphics Tier. By default, this is set to *FP16*. The options are:  * **FP16**: Color render texture format, 16-bit floating point per channel. * **R11G11B10**: Color render texture format. R and G channels are 11-bit floating point, B channel is 10-bit floating point. |
| **Rendering Path** | Choose how Unity should render graphics. Different rendering paths affect the performance of your game, and how lighting and shading are calculated. Some paths are more suited to different platforms and hardware than others. *Deferred* rendering is not supported when using Orthographic projection. If the camera’s projection mode is set to Orthographic, these values are overridden, and the camera always uses *Forward* rendering. The options are:  * **Forward** * **Deferred** * **Legacy Vertex Lit**  For more information, refer to [Rendering paths](rendering-paths-introduction.md). |
| **Realtime Global Illumination CPU Usage** | The CPU budget you allow Enlighten Realtime Global Illumination to use for lighting calculations at runtime. Increasing this makes the system react faster to changes in lighting at a cost of using more CPU time. **Note:** Some platforms allow all CPUs to be occupied by worker threads whereas some enforce maximums. For example, some gaming consoles allow a maximum of 4 CPU cores. For Android devices, if it is a bigLittle architecture, only the little CPUs are used; otherwise the maximum is one less than the total number of CPUs. If you use the URP or HDRP render pipelines, you can configure this property in the [quality settings](class-QualitySettings.md). The options are:  * **Low**: 25% of the allowed CPU threads are used as worker threads. * **Medium**: 50% of the allowed CPU threads are used as worker threads. * **High**: 75% of the allowed CPU threads are used as worker threads. * **Unlimited**: 100% of the allowed CPU threads are used as worker threads. |

### Built-in shader settings

Use these settings to specify which shader to use for each of the listed features in the Built-in Render Pipeline.

| **Rendering path** | **Shader to use** |
| --- | --- |
| **Deferred** | Use with [Deferred shading](RenderTech-DeferredShading.md). |
| **Deferred Reflection** | Use with [Reflection Probes](class-ReflectionProbe.md) in deferred shading. |
| **Screen Space shadows** | Use with cascaded shadow maps for directional lights on PC/console platforms. |
| **Motion vectors** | Use for object-based motion vector calculations. |
| **Lens Flare** | Use with [Lens Flares](class-Flare.md). |
| **Light Halo** | Use with [Light Halos](class-Halo.md). |

For each of these features, you can choose which type of Shader to use:

* **No Support** disables this calculation. Use this setting if you are not using deferred shading. This will save some space in the built game data files.
* **Built-in Shader** uses Unity’s built-in Shaders to do the calculation. This is the default.
* **Custom Shader** uses your own compatible Shader to do the calculation. This enables you to do deep customization of deferred rendering.

When you choose **Custom shader**, a **Shader** reference property appears below the feature property where you can set a reference to the Shader you want to use.

### Always Included Shaders

This is a list of shaders for which Unity includes all possible [variants](shader-variants.md) in every build. This can be useful if you want to use shaders or variants at runtime that would not otherwise be included in the build; for example, if you use AssetBundles or Addressables that rely on those shaders or variants, or if you use [shader keywords](shader-keywords.md) to change variants at runtime.

**Warning:** This feature is not recommended for shaders that have a large number of variants, such as the [Standard Shader](shader-StandardShader.md); it can lead to significant runtime and build time performance problems. You should instead create [shader variant collections](shader-variant-collections.md) that contain only the variants you need, and include those in your build.

**Note:** This setting overrides [shader keyword declaration settings](shader-keywords.md#declaring-keywords-definition-type). For each shader in the list, Unity includes all sets of all keywords, even if you use the “shader feature” declaration type.

To add a shader to the list, increase the value in the **Size** property. To remove the last shader in the list, decrease the **Size** property. To remove a shader which is not the last one in the list, you can set the value to **None**.

## Additional resources

* [Create and manage build profiles](create-build-profile.md)

GraphicsSettings