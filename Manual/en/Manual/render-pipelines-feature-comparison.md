# Render pipeline feature comparison

This page contains information on feature support in the Built-in **Render Pipeline**, the Universal Render Pipeline (URP), and the High Definition Render Pipeline (HDRP). It also contains suggested alternatives for certain features.

For ease of reference, features are broken down into the following categories:

* [Platform Support](#platform-support)
* [Lighting](#lighting)
* [Color](#color)
* [Camera](#camera)
* [Shaders](#shaders)
* [World building](#world-building)
* [Visual effects](#visual-effects)
* [2D and UI](#2d-and-ui)
* [XR](#xr)
* [Scripting](#scripting)
* [Optimizations](#optimizations)
* [Debug](#debug)

## Platform Support

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Windows** | Yes | Yes  DirectX 11, DirectX 12, Vulkan | Yes  DirectX 11, DirectX 12, Vulkan |
| **MacOS** | Yes | Yes | Yes |
| **Linux** | Yes | Yes | Yes  Vulkan |
| **XBox One** | Yes | Yes | Yes |
| **XBox Series** | Yes | Yes | Yes |
| **PlayStation 4** | Yes | Yes | Yes |
| **PlayStation 5** | Yes | Yes | Yes |
| **Nintendo Switch** | Yes | Yes | No |
| **iOS** | Yes | Yes | No |
| **Android** | Yes | Yes | No |
| **Desktop **VR**** | Yes | Yes | Yes |
| **Mobile VR** | Yes | Yes | No |
| **Hololens** | Yes | Yes | No |
| ****WebGL**** | Yes | Yes | No |
| **WebGPU** | No | Yes  Experimental | No |

## Lighting

### Lights

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Directional** | Yes | Yes | Yes |
| **Spot** | Yes  Supported shapes: Cone | Yes  Supported shapes: Cone | Yes  Supported shapes: Cone, Pyramid, Box |
| **Point** | Yes | Yes | Yes |
| **Area** | Yes  Supported shapes: Rectangle (only with Enlighten Realtime Global Illumination), Disc (baked only) | Yes  Supported shapes: Rectangle (baked only), Disc (baked only) | Yes  Supported shapes: Rectangle (real-time and baked), Tube (real-time only), Disc (baked only) |
| **Inner Spot angle** | No | Yes | Yes |
| **Shading technique** | Yes  Multiple passes | Yes  Single pass only | Yes  Hybrid tile and cluster |
| **Culling Per Object** | Yes | Yes | No |
| **Culling Per Tile** | No | No | Yes |
| **Culling Per Layer** | Yes  Based on GameObject layers | Yes  Based on dedicated Rendering layers (32 layers available) | Yes  Based on dedicated Rendering layers (only the first 16 are used) |
| **Multiple directional lights** | Yes | Yes  Supports shadowing for one directional light at a time. | Yes  Supports shadowing for one directional light at a time. To simulate more, use Spot light with a Box shape. |
| **Number of real-time lights per object** | Unlimited | Forward: 1 Directional light + up to 8 lights per-object depending on the settings in the Render Pipeline asset. Deferred and Deferred+: Unlimited lights for opaque objects. For transparents or objects that shade in forward mode the same limit as above applies. Forward+: Not Applicable. Forward+ shade lights per-tile, see “Number of real-time lights per Camera” below.  For more information check: Render Path Comparison | Unlimited (HDRP doesn’t classify lights per object, see bellow the “Number of real-time lights per Camera” section). |
| **Number of real-time lights per Camera** | Unlimited | On mobile platforms: 32. On other platforms: 256 (configurable to improve build time and performance adding the package Packages/com.unity.render-pipelines.universal-config, and changing the value of MAX\_VISIBLE\_LIGHT\_COUNT\_DESKTOP in Runtime/ShaderConfig.cs.hlsl and  k\_MaxVisibleLightCountDesktop in Runtime/ShaderConfig.cs). | In Deferred: 63 per 16x16 pixel tile.  In Forward: 63 per 32x32 pixel cluster.  You can disable the cap for the deferred rendering path in Frame Settings > Light Loop Debug > Deferred Tile. This will affect performance. |
| **Light attenuation type** | Yes  Legacy | Yes  InverseSquared | Yes  InverseSquared |
| **Vertex Lights** | Yes | Yes | No |
| **SH Lights** | Yes | No | No |
| **Light Cookies** | Yes  Only shape, no color (Alpha channel). | Yes  Colored cookie (RGB) | Yes  Colored cookie (RGB) |
| **IES Lights** | No | No | Yes |
| **Light Distance Fade** | No | No | Yes |
| **Physical light units** | No | No | Yes |
| **Light anchor tool** | Yes | Yes | Yes |

### Shadows

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Directional Light Shadows** | Yes | Yes | Yes |
| **Multiple Directional Light Shadows** | Yes | No | No |
| **Directional Light Cascade Shadows** | Yes  1, 2, or 4 cascades, control by percentage only. | Yes  1 to 4 cascades. Control by percentage, or switch to meters.  Settings are in URP asset. | Yes  1 to 4 cascades. Control by percentage or distance.   Settings are in Volumes. |
| **Shadow Cascade Blending** | No | No | Yes |
| **Spot Light Shadows** | Yes | Yes | Yes |
| **Point Light Shadows** | Yes | Yes | Yes |
| **Area Light Shadows** | No | No | Yes  Only for rectangle shape |
| **Stable Fit Shadow Projection** | Yes | Yes | Yes |
| **Close Fit Shadow Projection** | Yes | No | No  No support for orthographic projection. Use wide field of view to emulate. |
| **Shadow Screen Space Pass** | Yes | Yes | Yes  Mostly used for Raytracing. |
| **Shadow bias** | Yes  Supported types: Constant clip space offset. Normal bias | Yes  Supported types: Offsets shadowmap texels in the light direction, Normal bias | Yes  Supported types: Slope bias, Normal bias |
| **PCF filtering (Percentage Closer Filtering)** | No | Yes  PCF Tent 5x5 | Yes  For Point and Spot lights in “Low” (PCF 3x3 4 taps) and “Medium” (PCF 5x5 9 taps) quality settings, and for Directional lights in “Low” (PCF 5x5 9 taps) and “Medium” (PCF Tent 5x5 9 taps) quality settings. |
| **PCSS filtering (Percentage Closer Soft Shadows)** | No | No | Yes  For Point and Spot lights, Directional lights and Area lights when in “High” quality in the HDRP settings for each kind of light. |
| **EVSM filtering (Exponential Variance Shadow Mapping)** | No | No | Yes  For area lights only in all quality settings except “High”. |
| **Shadow update modes** | No | No | Yes  Every frame, On Enable, On Demand. Each cascade can be updated independently using the API (RequestSubShadowMapRendering) |
| **Resolution settings** | Yes  Configured in Quality Settings. | Yes  Configured in URP Lighting Settings. | Yes  Configure shadowmap resolution or budget per Light, and configure shadow atlas size in HDRP Asset. |
| **Dynamic rescale** | No | No | Yes |
| **Directional light shadow caching** | No | No | Yes |
| **Punctual and Area lights shadow caching** | No | No | Yes |
| **Dynamic Shadow Casters** | No | No | Yes |
| **Static Shadow Casters** | No | No | Yes |
| **Contact Shadows** | No | No | Yes |
| **Micro Shadows** | No | No | Yes |
| **Shadow Cascade Volume Controller** | No | No | Yes |
| **Shadow Distance Fade** | Yes | Yes | Yes  Configure in Shadows Volume. |
| **Shadow Mask** | Yes  Configure per scene. | Yes  Configure per scene. | Yes  Configure per Light. |
| **Distance Shadow Mask** | Yes  Configure per scene. | Yes  Forward Rendering Path only. | Yes  Configure per Light. |
| **Transparent objects casting shadows** | Yes  Use the Transparent Cutout shaders for objects with “gaps” such as fences, vegetation, etc. or custom pixel-lit shaders using the Geometry render queue. Opacity can affect shadow intensity using dithering. | No  Alpha clipping only. | Yes  Raytracing only (incl support of colored shadows). For raster, alpha clipping only (use Alpha clipping > Shadow threshold option). |
| **Transparent objects receiving shadows** | No | Yes | Yes |
| **Shadow tint** | No | No | Yes  Shadow tint or penumbra tint options |
| **Shadow matte** | No | No | Yes  Use Unlit master node and “Enable Shadow Matte” |
| **Shadow Layers** | No | No | Yes |

### Global Illumination

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Baked **Global Illumination**: Progressive CPU **Lightmapper**** | Yes | Yes | Yes |
| **Baked Global Illumination: Progressive GPU Lightmapper** | Yes | Yes | Yes |
| **Baked Global Illumination: Enlighten Lightmapper** | No  Deprecated | No  Deprecated | No  Deprecated |
| **Baked Global Illumination: Mixed Lighting modes** | Yes  Subtractive, Baked indirect, Shadow Mask, Distance Shadow mask | Yes  Subtractive, Baked indirect, Shadow Mask, Distance Shadow mask | Yes  Baked indirect, Shadow Mask (per light setting), Distance Shadow mask (per light setting) |
| **Baked Global Illumination: Double Sided GI** | Yes | Yes | Yes |
| **Realtime Global Illumination: Pre-computed (Enlighten)** | Yes  Support ends after Unity 6 | Yes  Support ends after Unity 6 | Yes  Support ends after Unity 6 |
| **Realtime Global Illumination: Screen space GI** | No | No | Yes |
| **Screen Space Reflections** | Yes | No | Yes |
| **Planar Reflections** | No | No | Yes |
| **Raytraced Reflections** | No | No | Yes |
| **Screen Space Refractions** | No | No | Yes |
| **Non directional **lightmap**** | Yes | Yes | Yes |
| **Directional lightmap** | Yes | Yes | Yes |

### Light Probes

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| ****Light Probes**: Blending** | Yes | Yes | Yes |
| **Light Probes: Custom provided** | Yes | Yes | Yes |
| **Light Probes: Occlusion Probes** | Yes | Yes | Yes |
| **Light Probes: Proxy volumes (LPPV)** | Yes | No  Deprecated | Yes |

### Adaptive Probe Volumes

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Probe Volumes: Per **pixel** lighting** | No | Yes | Yes |
| **Probe Volumes: Per vertex lighting** | No | Yes | No |
| **Probe Volumes: Blending** | No | Yes | Yes |
| **Probe Volumes: Touchup volumes** | No | Yes | Yes |
| **Probe Volumes: Stream from disk** | No | Yes  Incl. non compute pass for lower end devices. | Yes |
| **Probe Volumes: Sky occlusion** | No | Yes | Yes |
| **Probe Volumes: **Reflection Probes** Normalization** | No | No | Yes |

### Reflection Probes

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Reflection Probes: Baked** | Yes | Yes | Yes |
| **Reflection Probes: Real-time** | Yes | Yes | Yes |
| **Reflection Probes: Real-time time sliced all faces at once** | Yes | Yes | No |
| **Reflection Probes: Real-time time sliced individual faces** | Yes | Yes | Yes |
| **Reflection Probes: On-demand API** | Yes | Yes | Yes |
| **Reflection Probes: Proxy volumes** | No | No | Yes |
| **Reflection Probes: Simple sampling** | Yes  Choose per-object how to sample reflection probes. | Yes  Choose in the URP asset whether URP always performs simple sampling from reflection probes, or if it blends. | No  See Reflection Hierarchy. |
| **Reflection Probes: Blend probes sampling** | Yes  Choose per-object how to sample reflection probes.  Blend up to 2 reflection probes. | Yes  Choose in the URP asset whether URP always performs simple sampling from reflection probes, or if it blends.  If blending is enabled, URP always blends sky + 1 local probe, or 2 probes if there is no sky. | Yes  No blending limit. All reflections including planar blend between each other. |
| **Reflection Probes: Blend probes and skybox sampling** | Yes | Yes | Yes  See Reflection Hierarchy. |
| **Reflection Probes: Distanced based roughness** | No | No | Yes |
| **Reflection Probes: Box Projection** | Yes | Yes | Yes |
| **Reflection Probes: Oriented Box Projection** | No | No | Yes |
| **Reflection Probes: Sphere Projection** | Yes | No | Yes |
| **Reflection Probes: Per probe resolution** | Yes | No | Yes |

### Raytracing

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Ray-traced **Ambient Occlusion**** | No | No | Yes |
| **Ray-traced Contact Shadows** | No | No | Yes |
| **Ray-traced Global Illumination** | No | No | Yes |
| **Ray-traced Reflections** | No | No | Yes  Incl. decals |
| **Ray-traced Shadows** | No | No | Yes |
| **Ray-traced Recursive Rendering** | No | No | Yes |
| **AxF (automotive material)** | No | No | Yes |
| **Terrain** | No | No | Yes  Heightmap only. No support for terrain details and terrain trees. For vegetation on terrain, we recommend using mixed rendering modes for ray traced GI and reflections which allow to support foliage/vegetation & deformation, and for shadows using raster shadows. |
| **VFX Graph **particles**** | No | No | Yes |
| **Windows** | No | No | Yes  Requires DX12, and platforms with hardware supporting RTX. |
| **Linux** | No | No | No |
| **PS5** | No | No | Yes |
| **XBox Series** | No | No | Yes |

### Pathtracing

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Lit (Pathtracing support)** | No | No | Yes |
| **Hair (Pathtracing support)** | No | No | Yes |
| **Eye (Pathtracing support)** | No | No | No |
| **Fabric (Pathtracing support)** | No | No | Yes |
| **AxF (Pathtracing support)** | No | No | Yes |
| **Decals (Pathtracing support)** | No | No | Yes  Except emissive decals. |
| **Punctual lights (point, spot, directional)** | No | No | Yes  Incl. box spot light |
| **Area Lights** | No | No | Yes  Incl. volumetric fog control |
| **Tube and Disc Lights** | No | No | Yes |
| ****Depth of field** (Pathtracing support)** | No | No | Yes |
| ****Exponential Fog** (Pathtracing support)** | No | No | Yes |
| **Anisotropic fog (Pathtracing support)** | No | No | Yes |
| **Local Volumetric fog (Pathtracing support)** | No | No | No |
| ****HDRI** sampling** | No | No | Yes |
| **Denoiser** | No | No | Yes  NVidia Optix™ AI accelerated denoiser or Intel® Open Image Denoise. |
| **Recorder integration** | No | No | Yes |
| **VFX Graph particles** | No | No | No |

### Environment lighting

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Skybox** | Yes  Per project. Set in the Lighting window. | Yes  Per project. Set in the Lighting window. | Yes  Configure in Visual Environment Volume. Choose from HDRI sky, physical sky or custom sky. |
| **Gradient** | Yes  Per project. Set in the Lighting window. | Yes  Per project. Set in the Lighting window. | Yes  Configure in Visual Environment Volume. |
| **Color** | Yes  Per project. Set in the Lighting window. | Yes  Per project. Set in the Lighting window. | No  Alternative: use Gradient sky. |
| **Real-time Ambient Mode** | Yes  By script. | Yes  By script. | Yes  Configure in Visual Environment Volume. |
| **Baked Ambient Mode** | Yes  By default | Yes  By default | Yes  Configure in Visual Environment Volume.  Uses the environment stored in Lighting > Environment settings. |
| **Indirect Lighting Controller** | No | No | Yes |

## Color

### HDR

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| ****HDR** rendering** | Yes | Yes | Yes |
| **HDR output** | No | Yes  Only certain devices are supported, see URP Documentation for more information. | Yes  Use the HDROutputSettings API to configure HDR output options and HDRP’s tonemapping HDR options. See HDRP HDR documentation for more information. |

### Color space

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Linear** | Yes  Requires OpenGLES3.0 | Yes | Yes |
| **Gamma** | Yes | Yes | No |

## Camera

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Multi display** | Yes | Yes | Yes |
| **Hardware Dynamic resolution** | Yes | Yes  Supported on Vulkan, DirectX 12, Metal, and compatible consoles. | Yes  Supported on Vulkan, DirectX 12, Metal, and compatible consoles. |
| **Software Dynamic resolution** | No | No  Except for XR direct-to-eye rendering. | Yes |
| **Nvidia Deep learning super sampling (DLSS)** | No | No | Yes |
| **AMD FidelityFX™ Super Resolution 1.0** | No | Yes | Yes |
| **AMD FidelityFX™ Super Resolution 2.0** | No | No | Windows platforms only. |
| **TAA upsampler** | No | No | Yes |
| **Unity Spatial Temporal Post Process (STP)** | No | Yes | Yes |
| **Mip bias** | No | Yes | Yes |
| **Compositing multiple cameras** | Yes  Use multiple cameras | Yes  Use Camera stacking | Yes  Use HDRP Compositor |
| **Physical camera** | Yes  Affects field of view only | Yes  Affects field of view only | Yes  Affects field of view, exposure, bloom, and depth of field |
| **Multi-sample anti-aliasing (MSAA)** | Yes  Forward rendering path only | Yes  Forward rendering path only | Yes  Forward rendering path only |
| **Fast approximate anti-aliasing (FXAA)** | Yes | Yes | Yes |
| **Subpixel morphological anti-aliasing (SMAA)** | Yes | Yes | Yes |
| **Temporal anti-aliasing (TAA)** | Yes  Additional pass to render all object’s motion vectors. No override support for transparents. | Yes  Additional pass to render all object’s motion vectors. No override support for transparents. No custom motion vectors in ShaderGraph. Not compatible with Dynamic Resolution, MSAA nor Camera Stacking. 16-bits. | Yes  Uses per-shader motion vectors and the stencil buffer to only render motion vectors where needed. Allows custom motion vectors in shaders, for example using the ShaderGraph or Alembic packages. 32-bits (float). |
| **Depth Texture** | Yes | Yes | Yes |
| **Depth + Normal Texture** | Yes | Yes | Yes |
| **Depth + Normal Texture (affected by decals)** | No | No | Yes  Except if MSAA is enabled. |
| **Color Texture** | No | Yes | Yes |
| **Motion Vectors** | Yes | Yes | Yes |
| **Thickness Pass** | No | No | Yes  Optional |

## Shaders

### Authoring workflows

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Material Variants** | Yes | Yes | Yes |
| ****Shader** Graph** | Yes | Yes | Yes |
| **ShaderLab** | Yes | Yes  Hand-coded shaders are supported, but not recommended. Shader Graph is recommended (more easily upgradable). | Yes  Hand-coded shaders are supported, but not recommended. Shader Graph is recommended (more easily upgradable). |
| ****ShaderLab** Command: UsePass** | Yes | Yes | Yes |
| **ShaderLab Command: GrabPass** | Yes | No  Alternative: Scene color node in ShaderGraph / create a Renderer Feature | No  Alternative: Scene color node in ShaderGraph / create a Custom Pass |
| **Surface Shaders** | Yes | No. For more information about alternatives, refer to [Surface shaders](SL-SurfaceShaders.md) | No. For more information about alternatives, refer to [Surface shaders](SL-SurfaceShaders.md). |
| **Diffusion Profiles** | No | No | Yes |

### Prebuilt shaders

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Camera-relative rendering** | No | No | Yes |
| **Material type: Metallic / Standard** | Yes | Yes | Yes |
| **Material type: Specular** | Yes | Yes | Yes |
| **Surface Type and Blend Mode: Opaque** | Yes | Yes | Yes |
| **Surface Type and Blend Mode: Faded (Alpha Blend)** | Yes | Yes | Yes  Also supports premultiplied alpha. |
| **Surface Type and Blend Mode: Transparent** | Yes | Yes | Yes |
| **Surface Type and Blend Mode: Cutout** | Yes | Yes | Yes |
| **Surface Type and Blend Mode: Additive** | No | Yes | Yes |
| **Surface Type and Blend Mode: Multiply** | No | Yes | No  Not compatible with half res transparent optimization and offscreen UI blending. |
| **Surface Type and Blend Mode: Translucence** | No | No | Yes  Translucent Material Type in the Lit shaders |
| **Surface Inputs: Albedo (Base Map)** | Yes | Yes | Yes |
| **Surface Inputs: Specular** | Yes | Yes | Yes  Base map when using Material Type Specular |
| **Surface Inputs: Metallic** | Yes | Yes | Yes  Stored in the R channel of the Mask Map. |
| **Surface Inputs: Ambient Occlusion** | Yes | Yes | Yes  Stored in the G channel of the Mask Map. |
| **Surface Inputs: Smoothness** | Yes | Yes | Yes  Stored in the A channel of the Mask Map. |
| **Surface Inputs: **Normal Map**** | Yes | Yes | Yes |
| **Surface Inputs: Bent Normal Map** | No | No | Yes |
| **Surface Inputs: Detail Map** | Yes | Yes | Yes  Albedo stored in the R channel of the Detail Map. Smoothness stored in the B channel. |
| **Surface Inputs: Detail Normal Map** | Yes | Yes | Yes  Y axis stored in the G channel of the Detail Map. X axis stored in the A channel. |
| **Surface Inputs: Heightmap** | Yes  Pixel displacement only. | Yes  Pixel displacement only. | Yes  Pixel or vertex displacement. |
| **Surface Inputs: Detail Mask** | Yes | Yes | Yes  Stored in the B channel of the Mask Map. |
| **Surface Inputs: Emissive Map** | Yes | Yes | Yes |
| **Surface Inputs: Transmission Mask** | No | No | Yes  Also available in Shader Graph. |
| **Surface Inputs: Specular Occlusion** | No | No | Yes |
| **Surface Inputs: Planar & Triplanar mapping** | No  Alternative: Use ShaderGraph Triplanar node | No  Alternative: Use ShaderGraph Triplanar node | Yes |
| **Included Lighting models: ClearCoat** | No | Yes | Yes |
| **Included Lighting models: SubSurface Scattering** | No | No | Yes  Includes Dual Lobe multipliers and diffuse shading power for skin. |
| **Included Lighting models: Colored Translucent** | No | No | Yes |
| **Included Lighting models: Anisotropy** | No | No | Yes |
| **Included Lighting models: Irridescence** | No | No | Yes |
| **Included Lighting models: Fabric** | No | No | Yes |
| **Included Lighting models: Eye** | No | No | Yes  Eye caustics available in High quality mode. |
| **Included Lighting models: Hair (Approximate)** | No | No | Yes  Based on Kajiya Kay parametrization |
| **Included Lighting models: Hair (Physical)** | No | No | Yes  Based on Marshner-Disney parametrization. Cinematic option available to trade performance vs quality. |
| **Add custom Lighting model as a plugin** | Yes | No | Yes  Pluggable Master Node (undocumented) |
| **Features: Light Cookies** | Yes  Supports grayscale textures | Yes  Supports RGB textures | Yes  Supports RGB textures |
| **Features: Parallax Mapping** | Yes | Yes | Yes |
| **Features: Light Distance Fade** | No | No | Yes |
| **Features: Shadow Distance Fade** | Yes | Yes | Yes  Configure in Shadows Volume. |
| **Features: Shadow Cascade Blending** | No | No | Yes |
| **Features: GPU Instancing** | Yes | Yes | Yes |
| **Features: Double Sided GI** | Yes | No | Yes |
| **Features: Two Sided** | No | Yes  Configure with the RenderFace property. | Yes  Configure with the Double Side toggle. |
| **Features: Material sorting priority** | No | Yes | Yes |
| **Features: Renderer sorting priority** | No | No  Not exposed (use debug mode of the inspector). | Yes |
| **Features: GPU Tesselation** | No | No | Yes  Use the LitTesselation Shader. |
| **Features: Specular Occlusion** | No | No | Yes  Use only the ambient occlusion map, or both bent and ambient occlusion maps. |
| **Features: Geometric Specular Anti-Aliasing** | No | No | Yes |
| **Features: Add precomputed velocity** | No | Yes  Requires per vertex velocity information (for example, Alembic, or using ShaderGraph custom velocity option) | Yes  Requires per vertex velocity information (for example, Alembic, or using ShaderGraph custom velocity option) |

### Shader Graph

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **UI** | Yes | Yes | Yes |
| **Full screen master node** | No | Yes | Yes |
| **Custom **post-processing**** | No | Yes | Yes |
| **Custom **render texture**** | Yes | Yes | Yes |
| **Text **Mesh** Pro** | No | Yes | No |
| **Decals** | No | Yes  Affect opaque objects only. | Yes  Affect opaque and transparent objects. |
| **Fog Volume** | No | No | Yes |
| ****Terrain**** | No | No | No |
| **Physically Based Sky** | No | No | Yes |

## World building

### Environment

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Terrain system: Physically Based Shader** | No | Yes | Yes |
| **Terrain system: Simple Lit (Blinn-Phong) Shader** | Yes | No | No |
| **Terrain system: SpeedTree Shader** | Yes | Yes | Yes |
| **Terrain system: Paint Trees** | Yes | Yes | Yes |
| **Terrain system: Vegetation** | Yes | Yes | Yes |
| **Terrain system: Terrain Details** | Yes | Yes | Yes |
| **Terrain system: number of Terrain Layers** | Yes  Unlimited | Yes  With height-based blending: 4 With alpha blending: unlimited | Yes  8 |
| **Terrain system: GPU Instanced Rendering** | Yes | Yes | Yes |
| **Terrain system: Terrain Holes** | Yes | Yes | Yes |
| **Water system: Physically Based Shader** | No | No | Yes |
| **Water system: Compute based wave simulation** | No | No | Yes |
| **Water system: CPU wave simulation for gameplay integration** | No | No | Yes  Simulation on CPU (no latency) or on GPU with GPU readback (low latency but high performance). |
| **Water system: Local currents** | No | No | Yes |
| **Water system: Local foam** | No | No | Yes |
| **Water system: Surface deformer** | No | No | Yes |
| **Water system: Water excluder** | No | No | Yes |
| **Water system: Support for decals (eg: Foam)** | No | No | Yes |
| **Water system: Underwater Caustics** | No | No | Yes |
| **Water system: Under water rendering** | No | No | Yes  Includes Water line rendering. |
| **Water system: Under water volumetrics** | No | No | Yes |
| **SpeedTree 8** | Yes | Yes  Both Shader Graph and ShaderLab shaders available. | Yes  Shader Graph shader only. Support for subsurface map and motion vectors. |
| **SpeedTree 9** | Yes | Yes  Shader Graph shaders only. | Yes  Shader Graph shaders only. |
| ****Wind Zone**** | Yes | Yes | Yes |

### Sky

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Procedural Skybox** | Yes | Yes | No  This sky type is deprecated, but you can still use if if you install the Procedural Sky Sample. |
| **6 sided Skybox** | Yes | Yes | Yes  HDRI Sky supports cubemaps. Unity can build cubemaps from this type on import. |
| **Cubemap** | Yes | Yes | Yes  HDRI Sky supports cubemaps. |
| **Panoramic** | Yes | Yes | Yes  HDRI Sky supports cubemaps. Unity can build cubemaps from this type on import. |
| **Physical Sky** | No | No | Yes |
| **Gradient Sky** | No | No | Yes |
| **Sky distortion** | No | No | Yes  Procedural or Flowmap |
| **Cloud layers** | No | No | Yes |
| **Atmospheric Scattering** | No | No | Yes |

### Volumetrics

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Volumetric Clouds** | No | No | Yes |
| **Linear Fog** | Yes  Configure in Graphics Settings. | Yes  Configure in Graphics Settings. | No |
| **Exponential Fog** | Yes  Configure in Graphics Settings. | Yes  Configure in Graphics Settings. | Yes  Configure with the Fog Override. |
| **Exponential Squared** | Yes  Configure in Graphics Settings. | Yes  Configure in Graphics Settings. | No |
| **Local Volumetrics** | No | No | Yes |
| **3D Render texture for local volumetric fog** | No | No | Yes |
| **Volumetric Material (use ShaderGraph for local volumetric fog)** | No | No | Yes |
| **Fog Scattering & Atmospheric scattering** | No | No | Yes |

## Visual effects

### Post-processing

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Feature support** | Yes  Uses separate package: Post-Processing V2 | Yes  Uses integrated post-processing solution | Yes  Uses integrated post-processing solution |
| **Per-project default effects** | No | Yes  A default volume can be applied to the project using HDRP default settings. | Yes  A default volume can be applied to the project using HDRP default settings. |
| **Ambient Occlusion** | Yes  Supported types: Multi-scale ambient occlusion | Yes  Supported types: Screen Space Ambient Occlusion (Depth or normal) | Yes  Supported types: Screen Space Ground Truth Ambient Occlusion (Normal + temporal, tinted multi bounced, specular occlusion), Ray-traced ambient occlusion |
| **Exposure** | Yes  Supported types: Fixed, Automatic | Yes  Supported types: Fixed | Yes  Supported types: Fixed, Automatic (Eye adaptation), Curve Mapping, Physical Camera settings and Histogram. |
| **Bloom** | Yes | Yes | Yes |
| **Chromatic Aberration** | Yes | Yes | Yes |
| ****Tonemapping** (Neutral, ACES, Custom)** | Yes | Yes | Yes |
| **White balance** | No | Yes | Yes |
| **Color Channel Mixer** | No | Yes | Yes |
| **Color Adjustments** | Yes | Yes | Yes |
| **Color Curves** | Yes | Yes | Yes |
| **Lift, Gamma, Gain** | Yes | Yes | Yes |
| **Shadows, Midtones, Highlights** | Yes | Yes | Yes |
| **Split Toning** | No | No | Yes |
| **Depth of Field** | Yes  Supported types: Bokeh | Yes  Supported types: Bokeh and Gaussian | Yes  Supported types: Bokeh |
| **Physical Depth of Field** | No | No | Yes |
| **Film Grain** | Yes | Yes | Yes |
| **Lens Distortion** | Yes | Yes | Yes |
| ****Camera** Motion Blur** | Yes | Yes | Yes |
| **Object Motion Blur** | Yes | Yes | Yes |
| **Screen Space Reflections** | Yes | No | Yes |
| **Screen Space Refractions** | No | No | Yes |
| **Vignette** | Yes | Yes | Yes |
| **Panini Projection** | No | Yes | Yes |
| **Screen Space **Lens Flares**** | No | Yes | Yes |
| **Custom post-processing** | Yes | Yes  Available with Full Screen Pass Renderer Feature. | Yes  Either with script or ShaderGraph. |

### CPU Particles (Shuriken)

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Feature support** | Yes | Yes | Yes  GPU instancing is not supported. Distortion and flipbook blending can be done with Shader Graph. No standard particle shader available, but example particles ShaderGraph shaders can be found in the HDRP package samples. |
| **Lit Particles** | Yes | Yes | Yes |
| **Simple lit particles** | Yes  Uses Blinn Phong. | Yes | Yes |
| **Unlit particles** | Yes | Yes | Yes |
| **Soft Particles** | Yes | Yes | Yes  Available via Shader Graph. Examples available in the particles system shader samples in the HDRP package sample. |
| **Distortion** | Yes | Yes | Yes  Available via Shader Graph. Examples available in the particles system shader samples in the HDRP package sample. |
| **Flipbook blending** | Yes | Yes | Yes  Available via Shader Graph. Examples available in the particles system shader samples in the HDRP package sample. |
| **Trail** | Yes | Yes | Yes |
| **GPU Instancing** | Yes | Yes | No |

### GPU Particles (VFX Graph)

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Feature support** | No | Yes  Compute capable hardware only. OpenGL ES not supported. | Yes  Compute capable hardware only. |
| **Supported Shader graph shaders** | No | Yes  VFX Shader Graph, Lit, Unlit, 2D Sprite Lit, 2D Custom Lit | Yes  VFX Shader Graph, Lit, Unlit, Hair, Fabric (no support for tessellation, nor decal) |
| **Renders with 2D** | No | Yes  3D particles and support for sorting layers and sprites. No support for 2D Physics, 2D sprites emitters, nor 2D Lights | No |
| **Lit Particles** | No | Yes | Yes |
| **Simple lit particles** | No | No | Yes |
| **Unlit particles** | No | Yes | Yes |
| ****Soft Particles**** | No | Yes | Yes |
| **Distortion** | No | No | Yes |
| **Flipbook blending** | No | Yes  Linear interpolation or flipbook motion vectors | Yes  Linear interpolation or flipbook motion vectors |
| **Trail** | No | Yes  Experimental | Yes  Experimental |
| **Half-resolution** | No | No | Yes |
| **Decals** | No | Yes | Yes |
| **Motion Vectors** | No | Yes | Yes |
| **Camera Buffer** | No | Yes  Depth collision, Color buffer | Yes  Depth collision, Color buffer |
| **Instancing** | No | Yes | Yes |
| **Skinned Mesh Sampling** | No | Yes | Yes |
| **6-way lighting** | No | Yes | Yes |
| **Volumetric Fog Output** | No | No | Yes |
| **Custom HLSL Blocks** | No | Yes | Yes |

### Decals

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Decals** | No  Suggested alternative: Projector component. | Yes  Projector decals (Forward or Forward+, using a Decal buffer) or Screen Space Decals (Deferred or Deferred+), decal layers, no emissive decals | Yes  Projector decals (Forward or Deferred, using a Decal Buffer) or Mesh decals, support on transparents (Cluster decals), support decal layers, surface gradients, emissive decals |
| **Decal layers** | No | No | Yes |

### Other visual effects

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Projector component** | Yes | No  Use the Decal Renderer Feature instead. | No  Use the Decal Projector instead. |
| ****Line Renderer** component** | Yes | Yes | Yes |
| **Trail Renderer component** | Yes | Yes  You can also use VFX Graph to create a custom trail effect. | Yes  You can also use VFX Graph to create a custom trail effect. |
| **Billboard Renderer component** | Yes | Yes | No  Suggested alternative: Use VFX Graph. |
| **Halo component** | Yes | No  Suggested alternative: Use a Lens Flare (SRP) Data asset and a Lens Flare (SRP) component. | No  Suggested alternative: Use a Lens Flare (SRP) Data asset and a Lens Flare (SRP) component, or use a celestial body in a directional light. |
| **Lens flares** | Yes  Use a Flare asset and either a Light component or a Lens Flare component. | Yes  Use a Lens Flare (SRP) Data asset and a Lens Flare (SRP) component. Also check Screen Space Lens Flares (Post processing) that offer an alternative or complementary solution for lens flares. | Yes  Use a Lens Flare (SRP) Data asset and a Lens Flare (SRP) component. Also check Screen Space Lens Flares (Post processing) that offer an alternative or complementary solution for lens flares. |

## 2D and UI

### 2D

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| ****Sprite**** | Yes | Yes | No |
| ****Tilemap**** | Yes | Yes | No |
| **Sprite Shape** | Yes | Yes | No |
| **Pixel-Perfect** | Yes  Using the “2D Pixel Perfect” package. | Yes  Using the “2D Pixel Perfect” package. | No |
| **2D Lights** | No | Yes  Available when Using 2D Renderer | No |
| **2D Shadows** | No | Yes  Available when Using 2D Renderer | No |
| **VFX Graph** | No | Yes  Supports 3D particles, sorting layers, ShaderGraph Sprite Lit and Custom Lit targets, Alpha Clipping with Shadergraph Sprite Lit, Unlit and Custom Lit targets.  Does not support 2D Physics, 2D sprites emitters, or 2D Lights. | No |

### UI

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Screen Space - Overlay** | Yes | Yes | Yes |
| **Screen Space - Camera** | Yes | Yes | Yes |
| **World Space** | Yes | Yes | Yes |
| ****Text Mesh** Pro** | Yes | Yes | Yes |

## XR

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Multipass** | Yes | Yes | Yes |
| **Single Pass (double-wide)** | Yes | No  Deprecated | No  Deprecated |
| **Single Pass Instanced** | Yes | Yes | Yes  Windows and PSVR only |
| **Multiview** | Yes | Yes | No |
| ****AR** Foundation** | Yes | Yes | No |
| **Foveated Rendering** | No | Yes  Compatible with Forward on Oculus Quest (Fixed foveated rendering), PSVR2 and VisionPro (eye tracking). | No |
| **HDR** | Yes | Yes | Yes |

## Scripting

### Render Pipeline Hooks

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Camera.RenderWithShader** | Yes | No  Suggested alternative: ScriptableRenderPass | No  Suggested alternative: CustomPass with material override |
| **Camera.AddCommandBuffer (Camera.Remove[All]CommandBuffer)** | Yes | No | No  Alternative: Custom Passes |
| **Camera.Render** | Yes | Yes | Yes |
| **Light.AddCommandBuffer (LightRemove[All]CommandBuffer)** | Yes | No | No |
| **OnPreCull** | Yes | No  Suggested alternative: RenderPipelineManager.beginCameraRendering or ScriptableRenderPass | No |
| **OnPreRender** | Yes | No  Suggested alternative: RenderPipelineManager.beginCameraRendering or ScriptableRenderPass | No |
| **OnPostRender** | Yes | No | No |
| **OnRenderImage** | Yes | No  Suggested alternative: ScriptableRenderPass | No  Suggested alternative: Fullscreen CustomPass |
| **OnRenderObject** | Yes | Yes  RenderPipelineManager.endCameraRendering or ScriptableRenderPass API | No  RenderPipelineManager.endCameraRendering or Custom Pass |
| **OnWillRenderObject** | Yes | Yes | Yes |
| **OnBecameVisible** | Yes | Yes | Yes |
| **OnBecameInvisible** | Yes | Yes | Yes |
| **Camera Replacement Material** | No | No  Suggested alternative: use Render Feature to replace materials per-Pass. | No  Suggested alternative: Custom Pass with Shader override |
| **RenderPipeline.BeginFrameRendering** | No | Yes | Yes |
| **RenderPipeline.EndFrameRendering** | No | Yes | Yes |
| **RenderPipeline.BeginCameraRendering** | No | Yes | Yes |
| **RenderPIpeline.EndCameraRendering** | No | Yes | Yes |
| **UniversalRenderPipeline.RenderSingleCamera** | No | Yes | No |
| **ScriptableRenderPass** | No | Yes  This is similar to the CustomPass feature in HDRP. | No |
| **Custom Renderers** | No | Yes | No |
| **CustomPass** | No | No | Yes  This is similar to the ScriptableRenderPass feature in URP. |
| **Custom Post Process Pass** | No | No | Yes |
| **Arbitrary Output Variables (AOV) API** | No | No | Yes |
| **Recording API** | No | No | Yes |

## Optimizations

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| **Batching By Shader** | No | Yes | Yes |
| **Batching By Material** | Yes | Yes | Yes |
| **Batching by Material Property Blocks** | Yes | Yes  Batched separately from the SRP Batcher | Yes  Batched separately from the SRP Batcher |
| ****Dynamic batching** (without SRP Batcher)** | Yes | Yes | Yes |
| **Real-time batching with SRP Batcher** | No | Yes  Requires OpenGLES 3.1. | Yes |
| **Dynamic Shadows Batching** | Yes | No  Deprecated. | Yes |
| **GPU Resident Drawer** | No | Yes  Disabled by default. Supported on Forward+ and Deferred+. Supported on Vulkan, DirectX 11, DirectX 12, Metal, and consoles. Disabled on other platforms. | Yes  Enabled by default. |
| **GPU Occlusion Culling** | No | Yes  Supported on Forward+ and Deferred+. Supported on Vulkan, DirectX 11, DirectX 12, Metal, and consoles. Disabled on other platforms. | Yes |
| **GPU Instancing** | Yes | Yes  Disabled by default when SRP Batcher is enabled. | Yes  Disabled by default when SRP Batcher is enabled. |
| **Render Graph** | No | Yes  Optimized for tile-based GPUs with raster pass merging. | Yes |
| **Dynamic Render Pass Culling** | No | Yes | Yes  Disabled by default. |
| **Variable Rate Shading API** | Yes  API only. | Yes  API, and integration with custom Scriptable Renderer Features. | Yes  API, and integration with custom passes. |

## Debug

| **Feature** | **Built-in Render Pipeline** | **URP** | **HDRP** |
| --- | --- | --- | --- |
| ****Scene** View Modes** | Yes | Yes | Yes |
| **Volume Explorer** | No | Yes | Yes |
| **Rendering debugger** | No | Yes | Yes |
| **Runtime GPU profiler** | No | Yes | Yes |
| **Render Graph viewer** | No | Yes | Yes |
| **Probe volume debugger** | No | Yes | Yes |
| **Color Monitors** | No  Use post processing stack v2 | No | Yes  Waveform, Parade, Vectorscope |
| **Shader Graph Heatmap** | Yes | Yes | Yes |
| **VFX Graph profiler** | No | Yes | Yes |