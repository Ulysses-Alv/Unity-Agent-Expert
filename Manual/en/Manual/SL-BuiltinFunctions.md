# Use built-in shader functions in the Built-In Render Pipeline

Unity has a number of built-in utility functions in UnityCG.cginc designed to make writing **shaders** simpler and easier.

## Transform positions

| **Function:** | **Description:** |
| --- | --- |
| `float4 UnityObjectToClipPos(float3 pos)` | Transforms a point from object space to the **camera**’s clip space in homogeneous coordinates. This is the equivalent of **mul(UNITY\_MATRIX\_MVP, float4(pos, 1.0))**, and should be used in its place. |
| `float3 UnityObjectToViewPos(float3 pos)` | Transforms a point from object space to view space. This is the equivalent of **mul(UNITY\_MATRIX\_MV, float4(pos, 1.0)).xyz**, and should be used in its place. |

### Transform other values

| **Function:** | **Description:** |
| --- | --- |
| `float3 WorldSpaceViewDir (float4 v)` | Returns world space direction (not normalized) from given object space vertex position towards the camera. |
| `float3 ObjSpaceViewDir (float4 v)` | Returns object space direction (not normalized) from given object space vertex position towards the camera. |
| `float2 ParallaxOffset (half h, half height, half3 viewDir)` | calculates UV offset for parallax normal mapping. |
| `fixed Luminance (fixed3 c)` | Converts color to luminance (grayscale). |
| `fixed3 DecodeLightmap (fixed4 color)` | Decodes color from Unity **lightmap** (RGBM or dLDR depending on platform). |
| `float4 EncodeFloatRGBA (float v)` | Encodes (0..1) range float into RGBA color, for storage in low precision render target. |
| `float DecodeFloatRGBA (float4 enc)` | Decodes RGBA color into a float. |
| `float2 EncodeFloatRG (float v)` | Encodes (0..1) range float into a float2. |
| `float DecodeFloatRG (float2 enc)` | Decodes a previously-encoded RG float. |
| `float2 EncodeViewNormalStereo (float3 n)` | Encodes view space normal into two numbers in 0..1 range. |
| `float3 DecodeViewNormalStereo (float4 enc4)` | Decodes view space normal from enc4.xy. |

## Forward rendering helper functions

These functions are only useful when using **forward rendering** (ForwardBase or ForwardAdd pass types).

| **Function:** | **Description:** |
| --- | --- |
| `float3 WorldSpaceLightDir (float4 v)` | Computes world space direction (not normalized) to light, given object space vertex position. |
| `float3 ObjSpaceLightDir (float4 v)` | Computes object space direction (not normalized) to light, given object space vertex position. |
| `float3 Shade4PointLights (...)` | Computes illumination from four point lights, with light data tightly packed into vectors. Forward rendering uses this to compute per-vertex lighting. |

## Screen-space helper functions

The following functions are helpers to compute coordinates used for sampling screen-space textures. They return `float4` where the final coordinate to sample texture with can be computed via perspective division (for example `xy/w`).

The functions also take care of [platform differences](SL-PlatformDifferences.md) in **render texture** coordinates.

| **Function:** | **Description:** |
| --- | --- |
| `float4 ComputeScreenPos (float4 clipPos)` | Computes texture coordinate for doing a screenspace-mapped texture sample. Input is clip space position. |
| `float4 ComputeGrabScreenPos (float4 clipPos)` | Computes texture coordinate for sampling a [GrabPass](SL-GrabPass.md) texure. Input is clip space position. |

## Vertex-lit helper functions

These functions are only useful when using per-vertex lit shaders (“Vertex” pass type).

| **Function:** | **Description:** |
| --- | --- |
| `float3 ShadeVertexLights (float4 vertex, float3 normal)` | Computes illumination from four per-vertex lights and ambient, given object space position & normal. |

## Texture and sampler macros

| **Macro:** | **Use:** |
| --- | --- |
| `UNITY_DECLARE_TEX2D(name)` | Declares a Texture and Sampler pair. |
| `UNITY_DECLARE_TEX2D_NOSAMPLER(name)` | Declares a Texture without a Sampler. |
| `UNITY_DECLARE_TEX2DARRAY(name)` | Declares a Texture array Sampler variable. |
| `UNITY_SAMPLE_TEX2D(name,uv)` | Sample from a Texture and Sampler pair, using given Texture coordinate. |
| `UNITY_SAMPLE_TEX2D_SAMPLER( name,samplername,uv)` | Sample from Texture (name), using a Sampler from another Texture (samplername). |
| `UNITY_SAMPLE_TEX2DARRAY(name,uv)` | Sample from a Texture array with a float3 UV; the z component of the coordinate is array element index. |
| `UNITY_SAMPLE_TEX2DARRAY_LOD(name,uv,lod)` | Sample from a Texture array with an explicit mipmap level. |

## Shadow mapping macros

Declaring and sampling shadow maps can be very different depending on the platform. Unity has several macros to help with this:

| **Macro:** | **Use:** |
| --- | --- |
| `UNITY_DECLARE_SHADOWMAP(tex)` | Declares a shadowmap Texture variable with name “tex”. |
| `UNITY_SAMPLE_SHADOW(tex,uv)` | Samples shadowmap Texture “tex” at given “uv” coordinate (XY components are Texture location, Z component is depth to compare with). Returns single float value with the shadow term in 0..1 range. |
| `UNITY_SAMPLE_SHADOW_PROJ(tex,uv)` | Similar to above, but does a projective shadowmap read. “uv” is a float4, all other components are divided by .w for doing the lookup. |

The format of `tex` must be [RenderTextureFormat.Shadowmap](../ScriptReference/RenderTextureFormat.Shadowmap.md).

**NOTE:** Not all graphics cards support shadowmaps. Use [SystemInfo.SupportsRenderTextureFormat](../ScriptReference/SystemInfo.SupportsRenderTextureFormat.md) to check for support.

## Depth Texture helper macros

Most of the time, Depth Texture are used to render Depth from the Camera. The [UnityCG.cginc include file](SL-BuiltinIncludes.md) contains some macros to deal with the above complexity in this case:

* **COMPUTE\_EYEDEPTH(i)**: computes eye space depth of the vertex and outputs it in **o**. Use it in a vertex program when **not** rendering into a depth texture.
* **DECODE\_EYEDEPTH(i)/LinearEyeDepth(i)**: given high precision value from depth texture **i**, returns corresponding eye space depth.
* **Linear01Depth(i)**: given high precision value from depth texture **i**, returns corresponding linear depth in range between 0 and 1.
* **UNITY\_TRANSFER\_DEPTH(o)**: Deprecated. Computes the eye space depth of the vertex and outputs it in **o** (which must be a float2). On platforms with native depth textures this macro does nothing, because the Z buffer value is rendered implicitly.
* **UNITY\_OUTPUT\_DEPTH(i)**: Deprecated. Returns eye space depth from **i** (which must be a float2). On platforms with native depth textures this macro always returns zero, because the Z buffer value is rendered implicitly.

**Note:** On DX11/12 and Metal, the Z buffer range is 1–0 and UNITY\_REVERSED\_Z is defined. On other platforms, the range is 0–1.

For example, this shader would render depth of its **GameObjects**:

```
Shader "Render Depth" {
    SubShader {
        Tags { "RenderType"="Opaque" }
        Pass {
            HLSLPROGRAM

            #pragma vertex vert
            #pragma fragment frag
            #include "UnityCG.cginc"

            struct v2f {
                float4 pos : SV_POSITION;
                float2 depth : TEXCOORD0;
            };

            v2f vert (appdata_base v) {
                v2f o;
                o.pos = UnityObjectToClipPos(v.vertex);
                UNITY_TRANSFER_DEPTH(o.depth);
                return o;
            }

            half4 frag(v2f i) : SV_Target {
                UNITY_OUTPUT_DEPTH(i.depth);
            }
            ENDHLSL
        }
    }
}
```