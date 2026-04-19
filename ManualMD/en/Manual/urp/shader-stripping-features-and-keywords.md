# Settings and keywords reference for stripping keywords in URP

Explore the settings and **shader** keywords you can use to speed up build time in the Universal **Render Pipeline** (URP) by stripping shader variants. For more information, refer to [Strip shader variants](shader-stripping-features.md).

**Important**: When you remove built-in shader keywords, Unity might strip shader variants your built project needs. To check how many shader variants your project has, enable strict shader variant matching, otherwise Unity might replace missing shader variants with others. For more information, refer to [Check how many shader variants you have](../shader-how-many-variants.md#highlight-missing-shaders).

## Decals

| **Feature** | **Location in the Editor** | **Keyword set for Shader Build Settings** |
| --- | --- | --- |
| Decals that use **DBuffer** mode | [Universal Renderer](urp-universal-renderer.md). In the Decals Renderer Feature, set **Technique** to **Screen Space**. | `_ _DBUFFER_MRT1 _DBUFFER_MRT2 _DBUFFER_MRT3` |
| Decals that use **Screen Space** mode | [Universal Renderer](urp-universal-renderer.md). In the Decals Renderer Feature, set **Technique** to **DBuffer**. | `_DECAL_NORMAL_BLEND_LOW` `_DECAL_NORMAL_BLEND_MEDIUM` `_DECAL_NORMAL_BLEND_HIGH` |
| Rendering layers for decals | [Universal Renderer](urp-universal-renderer.md). In the Decals Renderer Feature, disable **Use Rendering Layers**. | `_ _DECAL_LAYERS` |

## Fog

| **Feature** | **Location in the Editor** | **Keyword set for Shader Build Settings** |
| --- | --- | --- |
| Fog | [Lighting window](../lighting-window.md). In the **Other Settings** section, disable **Fog**. | `_ _FOG_LINEAR FOG_EXP FOG_EXP2` |

To keep fog but reduce shader variants by using dynamic branching, refer to [Enable dynamic branching in prebuilt shaders](shader-stripping-fog.md).

## Level of detail (LOD)

| **Feature** | **Location in the Editor** | **Keyword set for Shader Build Settings** |
| --- | --- | --- |
| **LOD** transitions | [URP Asset](universalrp-asset.md). In the **Quality** section, disable **LOD Cross Fade**. | `_ LOD_FADE_CROSSFADE` |

For other ways to reduce LOD shader variants, refer to [Strip LOD shader variants](shader-stripping-lod).

## Lighting

| **Feature** | **Location in the Editor** | **Keyword set for Shader Build Settings** |
| --- | --- | --- |
| Additional lights (Forward rendering paths only) | [URP Asset](universalrp-asset.md). In the **Lighting section**, set **Additional Lights** to **Disabled**. | `_ _ADDITIONAL_LIGHTS_VERTEX _ADDITIONAL_LIGHTS` |
| Light cookies | [Light components](light-component.md). In the **Emission** section, remove cookie textures from all lights in your project. | `_ _LIGHT_COOKIES` |
| Rendering layers for lights | [URP Asset](universalrp-asset.md). For more information, refer to [Rendering Layers](features/rendering-layers.md). | `_ _LIGHT_LAYERS` |

## Lightmapping

| **Feature** | **Location in the Editor** | **Keyword set for Shader Build Settings** |
| --- | --- | --- |
| **Lightmaps** | [Lightmapping settings](../Lightmaps-reference.md). Disable **Realtime **Global Illumination**** and **Baked Global Illumination**. | `_ LIGHTMAP_ON` |
| Directional lightmaps | [Lightmapping settings](../Lightmaps-reference.md). Set **Directional Mode** to **Non-Directional**. | `_ DIRLIGHTMAP_COMBINED` |
| Bicubic sampling of lightmaps | [Graphics settings](urp-global-settings.md). In the **Lighting** section, disable **Use Bicubic Lightmap Sampling**. | `_ LIGHTMAP_BICUBIC_SAMPLING` |

## Materials

| **Feature** | **Location in the Editor** | **Keyword set for Shader Build Settings** |
| --- | --- | --- |
| Alpha clipping | [Material properties](shaders-in-universalrp-reference.md). In the **Surface Options** section, disable **Alpha Clipping**. | `_ _ALPHATEST_ON` |
| Clear coat | [Material properties](shaders-in-universalrp-reference.md). In the **Detail Inputs** section, disable **Clear coat** for all Complex Lit materials. | `_ _CLEARCOAT _CLEARCOATMAP` |
| Detail mapping | [Material properties](shaders-in-universalrp-reference.md). In the **Detail Inputs** section, remove the textures for all materials. | `_ _DETAIL_MULX2 _DETAIL_SCALED` |
| Emission | [Material properties](shaders-in-universalrp-reference.md). In the **Detail Inputs** section, disable **Emission** for all materials. | `_ _EMISSION` |
| Environment reflections | [Material properties](shaders-in-universalrp-reference.md). In the **Advanced Options** section, disable **Environment Reflections**. | `_ _ENVIRONMENTREFLECTIONS_OFF`. Removing this keyword enables environment reflections. |
| Height maps | [Material properties](shaders-in-universalrp-reference.md). In the **Surface Inputs** section, remove the **Height Map** texture for all materials. | `_ _PARALLAXMAP` |
| Metallic or specular maps | [Material properties](shaders-in-universalrp-reference.md). In the **Surface Inputs** section, remove the **Metallic Map** or **Specular Map** texture for all materials. | `_ _METALLICSPECGLOSSMAP` |
| **Normal maps** | [Material properties](shaders-in-universalrp-reference.md). In the **Surface Inputs** section, remove the **Normal Map** texture for all materials. | `_ _NORMALMAP` |
| Occlusion maps | [Material properties](shaders-in-universalrp-reference.md). In the **Surface Inputs** section, remove the **Occlusion Map** texture for all materials. | `_ _OCCLUSIONMAP` |
| Shadows | [Material properties](shaders-in-universalrp-reference.md). In the **Surface Options** section, disable **Receive shadows**. | `_ _RECEIVE_SHADOWS_OFF`. Set **Override type** to `shader_feature` or `multi_compile`. Removing this keyword enables **GameObjects** receiving shadows. |
| Smoothness from **Base Map** alpha | [Material properties](shaders-in-universalrp-reference.md). In the **Surface Inputs** section, under **Metallic Map** or **Specular Map**, set **Source** to **Metallic Alpha** or **Specular Alpha**. | `_ _SMOOTHNESS_TEXTURE_ALBEDO_CHANNEL_A` |
| Specular highlights | [Material properties](shaders-in-universalrp-reference.md). In the **Advanced Options** section, disable **Specular Highlights**. | `_ _SPECULARHIGHLIGHTS_OFF`. Removing this keyword enables specular highlights. |
| Specular highlights on transparent materials | [Material properties](shaders-in-universalrp-reference.md). In the **Surface Options** section on all transparent materials, disable **Preserve Specular Lighting** and set **Blending Mode** to either **Alpha**, **Premultiply**, or **Additive**. | `_ _ALPHAPREMULTIPLY_ON _ALPHAMODULATE_ON` |
| Specular setup | [Material properties](shaders-in-universalrp-reference.md). In the **Surface Options** section, set **Workflow mode** to **Metallic** for all materials. | `_ _SPECULAR_SETUP` |
| Transparency | In the **Surface Options** section of the Material properties, set **Surface Type** to **Opaque** for all materials. | `_SURFACE_TYPE_TRANSPARENT` |
| **Terrain** holes | [URP Asset](universalrp-asset.md). In the **Rendering** section. Disable **Terrain Holes**. | N/A |

## Post-processing

| **Feature** | **Location in the Editor** | **Keyword set for Shader Build Settings** |
| --- | --- | --- |
| Bloom | [Volume Overrides](VolumeOverrides.md). Remove the [Bloom](post-processing-bloom.md) Volume Override. | `_ _BLOOM_LQ _BLOOM_HQ _BLOOM_LQ_DIRT _BLOOM_HQ_DIRT` |
| Chromatic Aberration | [Volume Overrides](VolumeOverrides.md). Remove the [Chromatic Aberration](post-processing-chromatic-aberration.md) Volume Override. | `_ _CHROMATIC_ABERRATION` |
| Depth of Field - high quality sampling | [Volume Overrides](VolumeOverrides.md). In the [Depth of Field](depth-of-field-urp.md) Volume Override, disable **High Quality Sampling**. | `_ _HIGH_QUALITY_SAMPLING` |
| Dithering | [Camera properties](camera-component-reference.md). In the **Rendering** section, disable **Dithering**. | `_ _DITHERING` |
| Film Grain | [Volume Overrides](VolumeOverrides.md). Remove the [Film Grain](post-processing-film-grain.md) Volume Override. | `_ _FILM_GRAIN` |
| Lens Distortion | [Volume Overrides](VolumeOverrides.md). Remove the [Lens Distortion](post-processing-lens-distortion.md) Volume Override. | `_ _DISTORTION` |
| Screen Space **Ambient Occlusion** (SSAO) | [Universal Renderer](urp-universal-renderer.md). Remove the [Screen Space Ambient Occlusion Renderer Feature](post-processing-ssao.md). | `_ _SCREEN_SPACE_OCCLUSION` |
| Tonemapping | [Volume Overrides](VolumeOverrides.md). Remove [Tonemapping](post-processing-tonemapping.md) Volume Override. | `_ _HDR_GRADING _TONEMAP_ACES _TONEMAP_NEUTRAL` |
| Alpha processing | [URP Asset](universalrp-asset.md). In the ****Post-processing**** section, disable **Alpha Output**. | `_ _ENABLE_ALPHA_OUTPUT` |

## Rendering paths

| **Feature** | **Location in the Editor** | **Keyword set for Shader Build Settings** |
| --- | --- | --- |
| Accurate G-buffer normals (Deferred rendering paths only) | [Universal Renderer](urp-universal-renderer.md). In the **Rendering** section, disable **Accurate G-buffer normals**. This has no effect on platforms that use the Vulkan graphics API. | `_ _GBUFFER_NORMALS_OCT` |
| Fast sRGB to linear conversion | [URP Asset](universalrp-asset.md). In the **Post-processing** section, disable **Fast sRGB/Linear conversions**. | `_ _USE_FAST_SRGB_LINEAR_CONVERSION` |
| Forward+ **rendering path** | [Universal Renderer](urp-universal-renderer.md). In the **Rendering** section, set **Rendering Path** to **Forward+** or **Deferred+**. | `_ _CLUSTER_LIGHT_LOOP` |

## Reflections

| **Feature** | **Location in the Editor** | **Keyword set for Shader Build Settings** |
| --- | --- | --- |
| **Reflection Probe** blending | [Reflection Probe properties](../class-ReflectionProbe.md). Disable **Probe Blending** for all Reflection Probes. | `_ _REFLECTION_PROBE_BLENDING` |
| Reflection Probe box projection | [Reflection Probe properties](../class-ReflectionProbe.md). Disable **Box Projection** for all Reflection Probes. | `_ _REFLECTION_PROBE_BOX_PROJECTION` |
| Reflection Probe rotation | [Graphics settings](urp-global-settings.md). In the **Lighting** section, disable **Use Reflection Probe Rotation**. | `_ REFLECTION_PROBE_ROTATION` |

## Shadows

| **Feature** | **Location in the Editor** | **Keyword set for Shader Build Settings** |
| --- | --- | --- |
| Main lights cast shadows | [URP Asset](universalrp-asset.md). In the **Lighting** section, under **Main Light**, disable **Cast Shadows**. | `_ _MAIN_LIGHT_SHADOWS _MAIN_LIGHT_SHADOWS_CASCADE _MAIN_LIGHT_SHADOWS_SCREEN` |
| Additional lights cast shadows | [URP Asset](universalrp-asset.md). In the **Lighting** section, under **Additional Lights**, disable **Cast Shadows**. | `_ _ADDITIONAL_LIGHT_SHADOWS` |
| Point Lights cast shadows | [Light components](light-component.md). In the **Shadows** section, set **Shadow Type** to **No Shadows** for all Point Lights. | `_ _CASTING_PUNCTUAL_LIGHT_SHADOW` |
| Soft shadows | [URP Asset](universalrp-asset.md). In the **Shadows** section, disable **Soft Shadows**. | `_ _SHADOWS_SOFT _SHADOWS_SOFT_LOW _SHADOWS_SOFT_MEDIUM _SHADOWS_SOFT_HIGH` |
| Shadows from the Subtractive Lighting Mode | [URP Asset](universalrp-asset.md). In the **Lighting** section, disable **Mixed Lighting**. | `_ LIGHTMAP_SHADOW_MIXING` |
| Shadows from the **Shadowmask** Lighting Mode | [Lightmapping settings](../Lightmaps-reference.md). Set **Lighting Mode** to **Baked Indirect** or **Subtractive**. | `_ SHADOWS_SHADOWMASK` |

## Additional resources

* [Reducing shader variants](../shader-variants-landing.md)