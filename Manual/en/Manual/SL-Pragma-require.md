# HLSL pragma require command reference

Here are all the valid values for the `#pragma require` directive.

| **Value** | **Description** |
| --- | --- |
| `interpolators10` | At least 10 vertex-to-fragment interpolators (“varyings”) are available. |
| `interpolators15` | At least 15 vertex-to-fragment interpolators (“varyings”) are available.  **Note:** Internally, this also automatically adds `integers` to the list of requirements. |
| `interpolators32` | At least 32 vertex-to-fragment interpolators (“varyings”) are available. |
| `integers` | Integers are a supported data type, including bit/shift operations.  **Note:** Internally, this also automatically adds `interpolators15` to the list of requirements. |
| `mrt4` | At least 4 render targets are supported. |
| `mrt8` | At least 8 render targets are supported. |
| `derivatives` | **Pixel** **shader** derivative instructions (ddx/ddy) are supported. |
| `samplelod` | Explicit texture **LOD** sampling (tex2Dlod / SampleLevel) is supported. |
| `fragcoord` | Pixel location (XY on screen, ZW depth in clip space) input in pixel shader is supported. |
| `2darray` | 2D texture arrays are a supported data type. |
| `cubearray` | **Cubemap** arrays are a supported data type. |
| `instancing` | SV\_InstanceID input system value is supported. |
| `geometry` | Geometry shader stages are supported. |
| `compute` | Compute shaders, structured buffers, and atomic operations are supported. |
| `randomwrite` or `uav` | “Random write” (UAV) textures are supported. |
| `tesshw` | Hardware tessellation is supported, but not necessarily tessellation (hull/domain) shader stages. For example, Metal supports tessellation, but not hull or domain stages. |
| `tessellation` | Tessellation (hull/domain) shader stages are supported. |
| `msaatex` | The ability to access multi-sampled textures (Texture2DMS in HLSL) is supported. |
| `sparsetex` | Sparse textures with residency info (“Tier2” support in DirectX terms; `CheckAccessFullyMapped` HLSL function). |
| `framebufferfetch` or `fbfetch` | The ability to get the current framebuffer from GPU memory is supported. This process is sometimes called framebuffer fetch. For more information, refer to [Writing shaders for different graphics APIs](SL-PlatformDifferences.md). |
| `setrtarrayindexfromanyshader` | Setting the render target array index from any shader stage (not just the geometry shader stage) is supported. |
| `inlineraytracing` | Inline **ray tracing** is supported, so you can generate ray queries in the **rasterization** and compute stages of a shader. Refer to [SystemInfo.supportsInlineRayTracing](../ScriptReference/SystemInfo-supportsInlineRayTracing.md) for more information. |