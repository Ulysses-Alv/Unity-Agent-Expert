# Unity 6000.3 (6.3 LTS) - 3D Graphics Documentation Summary
## Comprehensive Reference for AI Agent Skills

---

## 1. SHADERLAB SYNTAX AND STRUCTURE

### Shader Object Structure
```shaderlab
Shader "<name>"
{
    <optional: Material properties>
    <One or more SubShader definitions>
    <optional: custom editor>
    <optional: fallback>
}
```

### SubShader Structure
```shaderlab
SubShader
{
    <optional: LOD>
    <optional: tags>
    <optional: commands>
    <One or more Pass definitions>
}
```

### Pass Structure
```shaderlab
Pass
{
    <optional: name>
    <optional: tags>
    <optional: commands>
    <optional: shader code>
}
```

### ShaderLab Languages
- **ShaderLab**: Declarative language defining Shader object structure (nested-braces syntax)
- **HLSL**: Programming language for shader programs themselves
- **NOT CG**: Despite .cginc extension, shader library files are written in HLSL

### Code Block Types
- `CGPROGRAM` / `ENDCG` - Legacy, but still used for Surface Shaders
- `HLSLPROGRAM` / `ENDHLSL` - Modern approach for vertex/fragment shaders
- Surface Shaders are NOT compatible with `HLSLPROGRAM` blocks - must use `CGPROGRAM`

### Built-in Include Files (.cginc)
| File | Purpose |
|------|---------|
| `UnityCG.cginc` | Commonly used built-in helper functions and data structures |
| `HLSLSupport.cginc` | Preprocessor macros for multi-platform development (auto-included with CGPROGRAM) |
| `UnityShaderVariables.cginc` | Built-in global variables (auto-included with CGPROGRAM) |
| `AutoLight.cginc` | Lighting and shadowing functionality |
| `Lighting.cginc` | Standard surface shader lighting models (auto-included for surface shaders) |
| `TerrainEngine.cginc` | Helper functions for terrain and vegetation shaders |
| `Tessellation.cginc` | Tessellation helper functions |
| `UnityDeferredLibrary.cginc` | Deferred shading variables |
| `UnityLightingCommon.cginc` | Common lighting variables like `_LightColor0` |
| `UnityGlobalIllumination.cginc` | GI decoding functions (`DecodeLightmap`, `ShadeSHPerPixel`) |

Location: `{unity install}/Data/Resources/CGIncludes/` (Windows) or `/Applications/Unity/Unity.app/Contents/Resources/CGIncludes/` (Mac)

---

## 2. MATERIAL PROPERTIES (Properties Block)

### Syntax
```shaderlab
Properties
{
    [optional: attribute] name("display text in Inspector", type name) = default value
}
```

### Property Types
```shaderlab
// Integer (real integer, preferred over legacy Int)
_IntegerName ("Integer display name", Integer) = 1

// Float with optional range slider
_FloatName ("Float display name", Float) = 0.5
_RangeName ("Float with range", Range(0.0, 1.0)) = 0.5

// Textures
_Texture2D ("Texture2D display name", 2D) = "white" {}   // Default: "white", "black", "gray", "bump", "red"
_Texture2DArray ("Texture2DArray display name", 2DArray) = "" {}
_Texture3D ("Texture3D", 3D) = "" {}                      // Default: gray (0.5,0.5,0.5,1)
_Cubemap ("Cubemap", Cube) = "" {}                        // Default: gray (0.5,0.5,0.5,1)
_CubemapArray ("CubemapArray", CubeArray) = "" {}

// Color and Vector
_ColorName ("Example color", Color) = (.25, .5, .5, 1)    // Maps to float4, shows color picker
_VectorName ("Example vector", Vector) = (.25, .5, .5, 1) // Maps to float4, shows 4 float fields
_VectorName2 ("Example vector", Vector, 2) = (.25, .5, .5, 1) // Shows only x,y components
```

### Property Attributes
| Attribute | Function |
|-----------|----------|
| `[Gamma]` | Indicates sRGB values for float/vector properties |
| `[HDR]` | HDR values for texture or color properties |
| `[HideInInspector]` | Hide property in Inspector |
| `[MainTexture]` | Set main texture (accessible via `Material.mainTexture`) |
| `[MainColor]` | Set main color (accessible via `Material.color`) |
| `[NoScaleOffset]` | Hide tiling/offset fields for texture |
| `[Normal]` | Indicates texture expects a normal map |
| `[PerRendererData]` | Texture comes from `MaterialPropertyBlock` |

### HLSL Variable Mapping
```hlsl
// ShaderLab Properties -> HLSL variables
fixed4 _MyColor;          // Color property
float4 _MyVector;         // Vector property
float _MyFloat;           // Float/Range property
sampler2D _MyTexture;     // 2D texture
samplerCUBE _MyCubemap;   // Cubemap
sampler3D _My3DTexture;   // 3D texture

// uniform keyword is optional
uniform float4 _MyColor;
```

### SRP Batcher Requirement
For URP, HDRP, and Custom SRP: **All per-material variables must be in the same `CBUFFER`**:
```hlsl
CBUFFER_START(UnityPerMaterial)
    float4 _MyColor;
    float _MyFloat;
    sampler2D _MyTexture;
CBUFFER_END
```

---

## 3. HLSL SEMANTICS (Vertex Input)

| Semantic | Description | Type |
|----------|-------------|------|
| `POSITION` | Vertex position | float3 or float4 |
| `NORMAL` | Vertex normal | float3 |
| `TEXCOORD0` - `TEXCOORD3` | UV coordinates | float2, float3, or float4 |
| `TANGENT` | Tangent vector for normal mapping | float4 |
| `COLOR` | Per-vertex color | float4 |
| `VPOS` | Pixel position in screen space (fragment only) | float4 (use `UNITY_VPOS_TYPE`) |
| `VFACE` | Indicates front/back facing | float |
| `SV_VertexID` | Vertex index (not in Surface Shaders) | uint |

### Prebuilt Vertex Structures (UnityCG.cginc)
```hlsl
appdata_base    // position, normal, texcoord
appdata_tan     // position, tangent, normal, texcoord
appdata_full    // position, tangent, normal, color, texcoord, texcoord1, texcoord2, texcoord3
appdata_img     // position, texcoord
```

### Custom Vertex Structure Example
```hlsl
struct custom_vertex_input {
    float4 vertex : POSITION;
    float3 normal : NORMAL;
    float2 uv : TEXCOORD0;
    float4 tangent : TANGENT;
    float4 color : COLOR;
};
```

---

## 4. BUILT-IN SHADER VARIABLES

### Transformation Matrices (float4x4, column major)
```hlsl
UNITY_MATRIX_MVP        // Model * View * Projection
UNITY_MATRIX_MV         // Model * View
UNITY_MATRIX_V          // View
UNITY_MATRIX_P          // Projection
UNITY_MATRIX_VP         // View * Projection
UNITY_MATRIX_T_MV       // Transpose of Model * View
UNITY_MATRIX_IT_MV      // Inverse transpose of Model * View
unity_ObjectToWorld     // Model matrix
unity_WorldToObject     // Inverse world matrix
```

### Camera and Screen
```hlsl
_WorldSpaceCameraPos    // float3 - World space camera position
_ProjectionParams       // float4 - x:1/-1, y:near, z:far, w:1/far
_ScreenParams           // float4 - x:width, y:height, z:1+1/width, w:1+1/height
_ZBufferParams          // float4 - For Z buffer linearization
unity_OrthoParams       // float4 - Orthographic camera params
```

### Time
```hlsl
_Time               // float4 - (t/20, t, t*2, t*3)
_SinTime            // float4 - (t/8, t/4, t/2, t)
_CosTime            // float4 - (t/8, t/4, t/2, t)
unity_DeltaTime     // float4 - (dt, 1/dt, smoothDt, 1/smoothDt)
```

### Lighting (Forward Rendering)
```hlsl
_LightColor0            // fixed4 - Light color (from UnityLightingCommon.cginc)
_WorldSpaceLightPos0    // float4 - Directional: (dir, 0), Other: (pos, 1)
unity_WorldToLight      // float4x4 - World-to-light matrix
unity_WorldToShadow     // float4x4[4] - World-to-shadow matrices
```

### LOD
```hlsl
unity_LODFade   // float4 - x:fade(0..1), y:fade quantized to 16 levels
```

---

## 5. PRAGMA DIRECTIVES

### Shader Stages
```hlsl
#pragma vertex vert       // Required - vertex shader function
#pragma fragment frag     // Required - fragment shader function
#pragma geometry geom     // Optional - geometry shader (not on Metal)
#pragma hull hull         // Optional - DirectX 11 hull shader
#pragma domain domain     // Optional - DirectX 11 domain shader
```

### Surface Shaders
```hlsl
#pragma surface surf Lambert                              // Basic
#pragma surface surf BlinnPhong addshadow fullforwardshadows vertex:disp tessellate:tessFixed nolightmap
#pragma surface surf Lambert vertex:dispNone tessellate:tessEdge tessphong:_Phong nolightmap
```

### Shader Variants and Keywords
```hlsl
#pragma multi_compile _ KEYWORD_ON         // Declare keywords (all included in build)
#pragma multi_compile_local _ KEYWORD      // Local scope
#pragma shader_feature _ KEYWORD           // Unused keywords stripped from build
#pragma shader_feature_local _ KEYWORD     // Local scope
#pragma hardware_tier_variants <values>    // Built-in Render Pipeline only - graphics tiers
#pragma skip_variants KEYWORD1 KEYWORD2    // Strip specified keywords
```

### GPU Requirements
```hlsl
#pragma target 4.0            // Minimum shader model
#pragma require integers mrt8 // Specific GPU features
#pragma require geometry      // Auto-added when using #pragma geometry
#pragma require tessellation  // Auto-added when using #pragma hull or #pragma domain
```

### Graphics APIs
```hlsl
#pragma only_renderers metal vulkan      // Compile only for these APIs
#pragma exclude_renderers gles3 webgpu   // Exclude these APIs
```

Supported renderer values: `d3d11`, `glcore`, `gles3`, `metal`, `ps4`, `ps5`, `playstation`, `switch`, `vulkan`, `webgpu`, `xboxseries`

### GPU Instancing
```hlsl
#pragma instancing_options assumeuniformscaling
#pragma instancing_options forcemaxcount:128
#pragma instancing_options nomatrices
#pragma instancing_options procedural:SetupFunction
```

### Other
```hlsl
#pragma once                              // Include file only once (requires caching preprocessor)
#pragma enable_d3d11_debug_symbols        // Debug symbols (reduces performance!)
#pragma skip_optimizations <apis>         // Force optimizations off
#pragma disable_fastmath                  // Precise IEEE 754 (Metal only)
#pragma editor_sync_compilation           // Force sync compilation (Editor only)
#pragma enable_cbuffer                    // Emit cbuffer(name) even on platforms without CBUFFER support
```

---

## 6. GPU RENDER STATE COMMANDS

### Cull
```shaderlab
Cull Back       // Default - cull back-facing polygons
Cull Front      // Cull front-facing (inside-out geometry)
Cull Off        // No culling (transparent objects, double-sided walls)
```

### ZTest (Depth Testing)
```shaderlab
ZTest Disabled  // Disable depth test
ZTest Never     // Draw nothing
ZTest Less      // Draw if in front (not equal)
ZTest LEqual    // Draw if in front or equal (DEFAULT)
ZTest Equal     // Draw only if equal distance
ZTest Greater   // Draw if behind
ZTest NotEqual  // Draw if not equal
ZTest GEqual    // Draw if behind or equal
ZTest Always    // No depth testing
```

### ZWrite (Depth Buffer Writing)
```shaderlab
ZWrite On   // Enable (DEFAULT for opaque)
ZWrite Off  // Disable (for semi-transparent)
```

### ZClip
```shaderlab
ZClip Off   // Disable depth clipping
ZClip On    // Enable (DEFAULT)
```

### Blend
```shaderlab
Blend Off                                           // Disable blending (DEFAULT)
Blend One Zero                                      // No blending (copy source)
Blend SrcAlpha OneMinusSrcAlpha                     // Standard alpha blending
Blend One One                                       // Additive blending
Blend OneMinusDstColor One                          // Soft additive
Blend DstColor Zero                                 // Multiplicative
Blend DstColor SrcColor                             // 2x multiplicative
Blend One OneMinusSrcColor                          // Screen blend

// Separate RGB and alpha blending
Blend One Zero, Zero One                            // RGB: copy, Alpha: zero

// Blend equation: finalValue = sourceFactor * sourceValue operation destinationFactor * destinationValue
```

### Blend Factors
`One`, `Zero`, `SrcColor`, `SrcAlpha`, `SrcAlphaSaturate`, `DstColor`, `DstAlpha`, `OneMinusSrcColor`, `OneMinusSrcAlpha`, `OneMinusDstColor`, `OneMinusDstAlpha`

### BlendOp
```shaderlab
BlendOp Add         // Add (DEFAULT)
BlendOp Sub         // Subtract destination from source
BlendOp RevSub      // Subtract source from destination
BlendOp Min         // Minimum
BlendOp Max         // Maximum
// Plus logical operations (DX 11.1+ or Vulkan) and advanced OpenGL operations
```

### ColorMask
```shaderlab
ColorMask RGB       // Write to R, G, B channels only
ColorMask A         // Write to alpha only
ColorMask 0         // Disable all color writes
ColorMask R G B A   // Any combination without spaces
```

### Stencil
```shaderlab
Stencil
{
    Ref 2                       // Reference value (0-255)
    ReadMask 255                // Read mask (default: 255)
    WriteMask 255               // Write mask (default: 255)
    Comp equal                  // Comparison: Never, Less, Equal, LEqual, Greater, NotEqual, GEqual, Always
    Pass keep                   // On pass: Keep, Zero, Replace, IncrSat, DecrSat, Invert, IncrWrap, DecrWrap
    Fail keep                   // On fail
    ZFail decrWrap              // On depth fail
    // Per-face operations (override general ones):
    CompBack always
    PassBack replace
    CompFront always
    PassFront replace
}
```

### Offset (Depth Bias)
```shaderlab
Offset Factor, Units
// Factor: Scales maximum Z slope
// Units: Scales minimum resolvable depth buffer value
// Example: Offset -1, -1 to prevent z-fighting on coplanar geometry
```

### AlphaToMask
```shaderlab
AlphaToMask On      // Enable alpha-to-coverage (MSAA)
AlphaToMask Off     // Disable (DEFAULT)
```

---

## 7. SURFACE SHADERS (Built-in Render Pipeline ONLY)

### Basic Surface Shader
```shaderlab
Shader "Example/Diffuse Simple" {
  SubShader {
    Tags { "RenderType" = "Opaque" }
    CGPROGRAM
    #pragma surface surf Lambert
    
    struct Input {
        float4 color : COLOR;
    };
    
    void surf (Input IN, inout SurfaceOutput o) {
        o.Albedo = 1;
    }
    ENDCG
  }
  Fallback "Diffuse"
}
```

### Surface Shader with Texture
```shaderlab
Shader "Example/Diffuse Texture" {
    Properties {
      _MainTex ("Texture", 2D) = "white" {}
    }
    SubShader {
      Tags { "RenderType" = "Opaque" }
      CGPROGRAM
      #pragma surface surf Lambert
      
      struct Input {
          float2 uv_MainTex;
      };
      sampler2D _MainTex;
      
      void surf (Input IN, inout SurfaceOutput o) {
          o.Albedo = tex2D (_MainTex, IN.uv_MainTex).rgb;
      }
      ENDCG
    }
    Fallback "Diffuse"
}
```

### Lighting Models
- **Lambert** - Diffuse lighting
- **BlinnPhong** - Specular lighting
- Custom: `half4 Lighting<Name>(SurfaceOutput s, UnityGI gi)`

### Surface Shader Optimization Directives
| Directive | Effect |
|-----------|--------|
| `halfasview` | Faster specular - half-vector computed per vertex |
| `noforwardadd` | One directional light only, rest as vertex/SH - smaller shader |
| `noambient` | Disable ambient lighting and SH lights |
| `nolightmap` | Disable lightmap support |
| `addshadow` | Generate shadow caster pass |
| `fullforwardshadows` | Support all shadow types in forward rendering |
| `approxview` | View direction normalized per vertex |

### Tessellation Surface Shaders
```shaderlab
#pragma target 4.6
#pragma surface surf BlinnPhong vertex:disp tessellate:tessFixed nolightmap
// or
#pragma surface surf BlinnPhong vertex:disp tessellate:tessDistance nolightmap
// or
#pragma surface surf BlinnPhong vertex:disp tessellate:tessEdge nolightmap
// or with Phong smoothing
#pragma surface surf Lambert vertex:dispNone tessellate:tessEdge tessphong:_Phong nolightmap
```

### Tessellation Functions
```hlsl
#include "Tessellation.cginc"

// Fixed tessellation
float4 tessFixed() { return _Tess; }

// Distance-based
float4 tessDistance(appdata v0, appdata v1, appdata v2) {
    return UnityDistanceBasedTess(v0.vertex, v1.vertex, v2.vertex, minDist, maxDist, _Tess);
}

// Edge-length based (recommended)
float4 tessEdge(appdata v0, appdata v1, appdata v2) {
    return UnityEdgeLengthBasedTess(v0.vertex, v1.vertex, v2.vertex, _EdgeLength);
}

// With frustum culling (best performance)
float4 tessEdge(appdata v0, appdata v1, appdata v2) {
    return UnityEdgeLengthBasedTessCull(v0.vertex, v1.vertex, v2.vertex, _EdgeLength, maxDisplacement);
}
```

### SurfaceOutput Structure
```hlsl
struct SurfaceOutput {
    half3 Albedo;       // Diffuse color
    half3 Normal;       // Normal in tangent space
    half3 Emission;     // Emissive color
    half Specular;      // Specular power (0-1)
    half Gloss;         // Specular intensity
    half Alpha;         // Alpha
};
```

---

## 8. VERTEX AND FRAGMENT SHADERS (Modern Approach)

### Basic Vertex/Fragment Shader
```shaderlab
Shader "VertexInputSimple" {
    SubShader {
        Pass {
            HLSLPROGRAM
            #pragma vertex vert
            #pragma fragment frag
            #include "UnityCG.cginc"
            
            struct v2f {
                float4 position : SV_POSITION;
                half4 color : COLOR;
            };
            
            v2f vert (appdata_base v) {
                v2f o;
                o.position = UnityObjectToClipPos(v.vertex);
                o.color.xyz = v.normal * 0.5 + 0.5;
                o.color.w = 1.0;
                return o;
            }
            
            half4 frag (v2f i) : SV_Target {
                return i.color;
            }
            ENDHLSL
        }
    }
}
```

### Built-in Transform Functions
```hlsl
float4 UnityObjectToClipPos(float3 pos)   // Object -> Clip space (use instead of mul(UNITY_MATRIX_MVP, pos))
float3 UnityObjectToViewPos(float3 pos)   // Object -> View space
float3 WorldSpaceViewDir(float4 v)        // World space direction to camera (not normalized)
float3 ObjSpaceViewDir(float4 v)          // Object space direction to camera
float3 WorldSpaceLightDir(float4 v)       // World space direction to light (forward rendering)
float3 ObjSpaceLightDir(float4 v)         // Object space direction to light
```

### Texture Sampling Macros
```hlsl
UNITY_DECLARE_TEX2D(name)           // Declare texture + sampler
UNITY_DECLARE_TEX2D_NOSAMPLER(name) // Declare texture without sampler
UNITY_DECLARE_TEX2DARRAY(name)      // Declare texture array
UNITY_SAMPLE_TEX2D(name, uv)        // Sample texture
UNITY_SAMPLE_TEX2D_SAMPLER(name, samplername, uv)  // Sample with another texture's sampler
UNITY_SAMPLE_TEX2DARRAY(name, uv)   // Sample texture array (uv.z = array index)
UNITY_SAMPLE_TEX2DARRAY_LOD(name, uv, lod)  // Sample with explicit mipmap level
```

### Inline Sampler Suffixes (DX11, DX12, Metal, Vulkan only)
```hlsl
SamplerState sampler_linear_repeat;     // Linear filtering, repeat wrapping
SamplerState sampler_point_clamp;       // Point filtering, clamp wrapping
SamplerState sampler_trilinear_mirror;  // Trilinear filtering, mirror wrapping
SamplerState sampler_linear_repeat_aniso8;  // + 8x anisotropic filtering
// Suffixes: point, linear, trilinear, clamp, repeat, mirror, mirroronce, compare
// Anisotropic: aniso2, aniso4, aniso8, aniso16
// Axis-specific: clampu, repeatv, mirrorw
```

### Depth Texture Macros
```hlsl
COMPUTE_EYEDEPTH(o)         // Compute eye space depth in vertex program
DECODE_EYEDEPTH(i)          // Decode depth texture value to eye space depth
LinearEyeDepth(i)           // Same as DECODE_EYEDEPTH
Linear01Depth(i)            // Linear depth in 0..1 range
// Note: UNITY_REVERSED_Z is defined on DX11/12 and Metal (Z range 1-0)
```

### Shadow Macros
```hlsl
UNITY_DECLARE_SHADOWMAP(tex)          // Declare shadowmap
UNITY_SAMPLE_SHADOW(tex, uv)          // Sample shadow (uv.xy = location, uv.z = depth compare)
UNITY_SAMPLE_SHADOW_PROJ(tex, uv)     // Projective shadowmap read
```

### Constant Buffers (CBUFFER)
```hlsl
CBUFFER_START(UnityPerMaterial)
    float4 _MyColor;
    float _MyFloat;
CBUFFER_END

CBUFFER_START(RarelyUpdatedVariables)
    float4 lightPosition;
CBUFFER_END

CBUFFER_START(FrequentlyUpdatedVariables)
    float4 colorTint;
CBUFFER_END
```

---

## 9. SHADER KEYWORDS AND VARIANTS

### Declaring Keywords
```hlsl
// multi_compile: ALL variants included in build
#pragma multi_compile _ FEATURE_ON FEATURE_OFF
#pragma multi_compile_local _ LOCAL_FEATURE

// shader_feature: Unused variants stripped from build
#pragma shader_feature _ USE_TEXTURE
#pragma shader_feature_local _ LOCAL_OPTION

// Hardware tier variants (Built-in Render Pipeline only)
#pragma hardware_tier_variants d3d11
```

### Conditional Code
```hlsl
#if defined(FEATURE_ON)
    // Feature enabled code
#else
    // Feature disabled code
#endif

// Or with keywords
#if USE_TEXTURE
    half4 tex = UNITY_SAMPLE_TEX2D(_MainTex, uv);
#endif
```

### Toggling Keywords (C#)
```csharp
// On Material
material.EnableKeyword("FEATURE_ON");
material.DisableKeyword("FEATURE_ON");
bool isEnabled = material.IsKeywordEnabled("FEATURE_ON");

// On MaterialPropertyBlock
propertyBlock.EnableKeyword("FEATURE_ON");
propertyBlock.DisableKeyword("FEATURE_ON");

// Global keywords
Shader.EnableKeyword("GLOBAL_FEATURE");
Shader.DisableKeyword("GLOBAL_FEATURE");
```

### Built-in Keyword Sets (Shortcuts)
- `#pragma multi_compile_fwdbase` - Forward base pass lighting/shadow keywords
- `#pragma multi_compile_fwdadd` - Forward additive pass keywords
- `#pragma multi_compile_fwdadd_fullshadows` - Forward add with full shadows
- `#pragma multi_compile_shadowcaster` - Shadow caster pass keywords
- `#pragma multi_compile_instancing` - GPU instancing keywords
- `#pragma multi_compile_fog` - Fog keywords

---

## 10. MATERIALS AND PROPERTIES (C# API)

### Accessing Material Properties
```csharp
// Access via Material
Material mat = renderer.material;  // Creates instance (use sharedMaterial for shared)
mat.color = Color.red;             // Sets _Color
mat.mainTexture = someTexture;     // Sets _MainTex
mat.SetFloat("_MyFloat", 0.5f);
mat.SetColor("_MyColor", new Color(1, 0, 0, 1));
mat.SetTexture("_MainTex", texture);
mat.SetVector("_MyVector", new Vector4(1, 2, 3, 4));

// Access shared material (doesn't create instance)
Material sharedMat = renderer.sharedMaterial;

// MaterialPropertyBlock for per-instance properties without material duplication
MaterialPropertyBlock block = new MaterialPropertyBlock();
block.SetColor("_MyColor", Color.red);
block.SetTexture("_MainTex", texture);
renderer.SetPropertyBlock(block);
```

### Creating Materials
```csharp
Material mat = new Material(Shader.Find("Standard"));
mat = new Material(shaderInstance);
```

---

## 11. STANDARD SHADER PARAMETERS

### Standard Shader Properties
| Property | Type | Description |
|----------|------|-------------|
| `_Color` | Color | Albedo color (tint) |
| `_MainTex` | 2D | Albedo texture (RGB = color, A = transparency for cutout/fade) |
| `_Metallic` | Range(0-1) | Metallic value (0 = dielectric, 1 = metallic) |
| `_MetallicGlossMap` | 2D | Metallic (R) + Smoothness (A) texture |
| `_Glossiness` | Range(0-1) | Smoothness value |
| `_BumpMap` | 2D | Normal map |
| `_BumpScale` | Float | Normal map intensity |
| `_ParallaxMap` | 2D | Height map for parallax |
| `_Parallax` | Range(0.005-0.08) | Height scale |
| `_OcclusionMap` | 2D | Ambient occlusion (R channel) |
| `_OcclusionStrength` | Float | AO intensity |
| `_EmissionColor` | Color | Emission color |
| `_EmissionMap` | 2D | Emission texture |
| `_DetailMask` | 2D | Detail mask (A channel controls detail blending) |
| `_DetailAlbedoMap` | 2D | Detail albedo (multiplied with main albedo) |
| `_DetailNormalMap` | 2D | Detail normal map |
| `_DetailNormalMapScale` | Float | Detail normal intensity |
| `_UVSec` | Integer | UV set for detail maps (0 or 1) |

### Rendering Modes
| Mode | Blend | ZWrite | Description |
|------|-------|--------|-------------|
| Opaque | Off | On | Default, best performance |
| Cutout | Off | On | Alpha tested transparency (hard edges) |
| Fade | SrcAlpha OneMinusSrcAlpha | Off | Soft transparency (fades to invisible) |
| Transparent | SrcAlpha OneMinusSrcAlpha | Off | Physical transparency (glass-like) |

### Fresnel Effect
The Standard Shader includes a built-in Fresnel effect for reflections. Metallic surfaces have stronger Fresnel at grazing angles.

---

## 12. LIGHTING

### Light Types
- **Directional Light** - Infinite parallel rays (sun)
- **Point Light** - Omnidirectional from a point
- **Spot Light** - Cone-shaped light
- **Area Light** - Rectangular light (baked only in Built-in RP)

### Light Components
```csharp
Light light = GetComponent<Light>();
light.type = LightType.Directional;  // Directional, Point, Spot, Area, Rectangle
light.color = Color.white;
light.intensity = 1.0f;
light.range = 10.0f;          // For Point/Spot lights
light.spotAngle = 60.0f;      // For Spot lights
light.shadows = LightShadows.Soft;  // None, Hard, Soft
light.cookie = cookieTexture; // Cookie/cutter texture
```

### Rendering Paths (Built-in RP)
| Path | Description | Use Case |
|------|-------------|----------|
| Forward | One pass per light per object | Most common, good for few lights |
| Deferred | G-buffer based, lighting is screen-space pass | Many lights, desktop/console |
| Vertex Lit | Per-vertex lighting (legacy) | Very low-end hardware |

### Forward Rendering Pass Types
- `ForwardBase` - Ambient, lightmaps, main directional, vertex/SH lights
- `ForwardAdd` - Additive per-pixel lights (one invocation per light)
- `ShadowCaster` - Renders depth for shadows
- `Deferred` - For deferred shading path
- `Vertex`, `VertexLMRGBM`, `VertexLM` - Legacy vertex lit

### Direct vs Indirect Lighting
- **Direct**: Light that comes straight from light sources
- **Indirect**: Light that has bounced off surfaces (GI)

### HDR (High Dynamic Range)
- HDR allows color values > 1.0 for realistic lighting
- Use `[HDR]` attribute on color properties
- HDR color picker in Inspector
- Required for bloom, tonemapping, and other post-processing effects

---

## 13. SHADOWS

### Shadow Configuration
```csharp
// Quality Settings
QualitySettings.shadowResolution = ShadowResolution.High;
QualitySettings.shadowDistance = 150.0f;       // Max shadow cast distance
QualitySettings.shadowCascades = 4;            // 0, 2, 4 cascades
QualitySettings.shadowNearPlaneOffset = 2.0f;  // Prevent shadow acne

// Per Light
light.shadowStrength = 1.0f;       // 0 = no shadow, 1 = full shadow
light.shadowBias = 0.05f;          // Shadow map bias
light.shadowNormalBias = 0.4f;     // Normal bias
light.shadowNearPlane = 0.1f;      // Near plane for shadow camera
```

### Shadow Cascades (Directional Lights)
- **0 cascades**: Single shadow map (cheapest)
- **2 cascades**: Split into near/far regions
- **4 cascades**: Split into 4 regions for better quality at all distances
- Cascades improve quality near camera while maintaining coverage at distance

### Shadow Types
- **Hard Shadows**: Sharp edges, faster
- **Soft Shadows**: Filtered edges, more expensive

### Shadow Optimization
- Reduce `shadowDistance` to limit shadow rendering range
- Use fewer cascades for performance
- Lower shadow resolution
- Use `ShadowDistance` to cull shadows at distance
- Not important lights can use vertex shadows instead of pixel shadows

---

## 14. LIGHTMAPPING AND GI

### Lightmap Modes
| Mode | Description |
|------|-------------|
| Realtime | Light calculated at runtime |
| Mixed | Direct light realtime, indirect baked |
| Baked | All lighting precomputed |

### Lightmapping Configuration
```csharp
// Light component
light.lightmapBakeType = LightmapBakeType.Baked;  // Baked, Mixed, Realtime
light.bounceIntensity = 1.0f;                      // GI bounce contribution
light.shadowType = LightShadowType.Soft;

// Lighting Settings
Lightmapping.bakeGI = true;
Lightmapping.indirectResolution = 1.0f;            // Texels per unit
Lightmapping.bounceIntensity = 1.0f;
```

### Lightmap Parameters
- **Resolution**: Texels per world unit
- **Padding**: Space between UV charts (prevent bleeding)
- **Compressed**: Use compressed lightmaps (smaller, slightly lower quality)
- **Ambient Occlusion**: Enable baked AO
- **Bounces**: Number of GI bounces

### Light Probes
- Sample baked GI at arbitrary points in space
- Used by dynamic objects to receive indirect lighting
- Placed in tetrahedral groups (4 probes per group)
- Dynamic objects with MeshRenderer automatically use light probes

```csharp
// Light Probe Proxy Volume (LPPV) - for large objects
// Provides interpolated light probe data across large volumes
// Configure via LightProbeProxyVolume component
```

### Light Probe Placement
- Place probes where dynamic objects will be
- More probes in areas with complex lighting
- Fewer probes in uniform lighting areas
- Use Density Resolution to control probe density

---

## 15. REFLECTION PROBES

### Types
| Type | Update Mode | Use Case |
|------|-------------|----------|
| Realtime | Every frame | Dynamic environments |
| Baked | Bake time | Static environments |
| Custom | Via script | Controlled updates |

### Configuration
```csharp
ReflectionProbe probe = GetComponent<ReflectionProbe>();
probe.mode = ReflectionProbeMode.Realtime;    // Realtime, Baked, Custom
probe.refreshMode = ReflectionProbeRefreshMode.EveryFrame;  // EveryFrame, ViaScripting, Awake
probe.resolution = 128;                       // Cubemap resolution
probe.hdr = true;                             // HDR cubemap
probe.boxProjection = true;                   // Box projection for interior spaces
probe.intensity = 1.0f;
probe.blendDistance = 1.0f;                   // Blend zone between probes
probe.nearClip = 0.3f;
probe.farClip = 1000.0f;
probe.cullingMask = -1;                       // Layer mask
```

### Best Practices
- Use baked probes for static environments (best performance)
- Use realtime probes sparingly (expensive)
- Use box projection for interior spaces
- Set appropriate resolution (64-256 typically)
- Use blend distance for smooth transitions between probes

---

## 16. GPU INSTANCING

### Shader Modifications for GPU Instancing
```hlsl
// Add instancing pragma
#pragma multi_compile_instancing

// Use UNITY_INSTANCING_BUFFER_START/END for per-instance properties
UNITY_INSTANCING_BUFFER_START(Props)
    UNITY_DEFINE_INSTANCED_PROP(float4, _Color)
    UNITY_DEFINE_INSTANCED_PROP(float, _Scale)
UNITY_INSTANCING_BUFFER_END(Props)

// In vertex shader
void vert(appdata v, out v2f o)
{
    UNITY_SETUP_INSTANCE_ID(v);
    UNITY_TRANSFER_INSTANCE_ID(v, o);
    
    float4 color = UNITY_ACCESS_INSTANCED_PROP(Props, _Color);
    // ...
}

// In fragment shader
half4 frag(v2f i) : SV_Target
{
    UNITY_SETUP_INSTANCE_ID(i);
    float4 color = UNITY_ACCESS_INSTANCED_PROP(Props, _Color);
    // ...
}
```

### Complete GPU Instancing Shader Example
```shaderlab
Shader "Custom/InstancedShader" {
    Properties {
        _Color ("Color", Color) = (1,1,1,1)
        _MainTex ("Albedo", 2D) = "white" {}
    }
    SubShader {
        Pass {
            Tags { "LightMode" = "ForwardBase" }
            
            CGPROGRAM
            #pragma vertex vert
            #pragma fragment frag
            #pragma multi_compile_instancing
            #include "UnityCG.cginc"
            
            struct appdata {
                float4 vertex : POSITION;
                float2 uv : TEXCOORD0;
                UNITY_VERTEX_INPUT_INSTANCE_ID
            };
            
            struct v2f {
                float2 uv : TEXCOORD0;
                float4 vertex : SV_POSITION;
                UNITY_VERTEX_INPUT_INSTANCE_ID
            };
            
            UNITY_INSTANCING_BUFFER_START(Props)
                UNITY_DEFINE_INSTANCED_PROP(float4, _Color)
            UNITY_INSTANCING_BUFFER_END(Props)
            
            sampler2D _MainTex;
            float4 _MainTex_ST;
            
            v2f vert (appdata v) {
                v2f o;
                UNITY_SETUP_INSTANCE_ID(v);
                UNITY_TRANSFER_INSTANCE_ID(v, o);
                o.vertex = UnityObjectToClipPos(v.vertex);
                o.uv = TRANSFORM_TEX(v.uv, _MainTex);
                return o;
            }
            
            fixed4 frag (v2f i) : SV_Target {
                UNITY_SETUP_INSTANCE_ID(i);
                fixed4 col = UNITY_ACCESS_INSTANCED_PROP(Props, _Color);
                col *= tex2D(_MainTex, i.uv);
                return col;
            }
            ENDCG
        }
    }
}
```

### Enabling GPU Instancing (C#)
```csharp
// Enable on Material
material.enableInstancing = true;

// Use MaterialPropertyBlock for per-instance data
MaterialPropertyBlock block = new MaterialPropertyBlock();
block.SetColor("_Color", Color.red);
Graphics.DrawMeshInstanced(mesh, 0, material, matrices, count, block);

// Or use DrawMeshInstancedIndirect for GPU-driven instancing
Graphics.DrawMeshInstancedIndirect(mesh, 0, material, bounds, argsBuffer);
```

### GPU Instancing Options
```hlsl
#pragma instancing_options assumeuniformscaling  // All instances have uniform scale
#pragma instancing_options forcemaxcount:128     // Force max instances per draw
#pragma instancing_options nomatrices            // Don't pass unity_ObjectToWorld
#pragma instancing_options procedural:setup      // Custom setup function
```

---

## 17. SRP BATCHER

### What is SRP Batcher?
- Ultra-fast rendering path for SRPs (URP, HDRP)
- Reduces CPU overhead by caching material property locations
- Works with any shader that has properties in a single `CBUFFER`

### Shader Requirements for SRP Batcher
```hlsl
// ALL per-material properties must be in ONE CBUFFER named UnityPerMaterial
CBUFFER_START(UnityPerMaterial)
    float4 _MainTex_ST;
    float4 _Color;
    float _Metallic;
    float _Glossiness;
CBUFFER_END

// Textures and samplers are NOT in the CBUFFER
sampler2D _MainTex;
```

### SRP Batcher Incompatible Patterns
- Multiple CBUFFER blocks for per-material properties
- Properties outside any CBUFFER
- Properties in CBUFFERs with different names
- Legacy fixed-function shaders

### Enabling SRP Batcher
```csharp
// Enabled by default in URP and HDRP
// Verify in Frame Debugger
// Profile with Profiler module "SRP Batcher"
```

---

## 18. DRAW CALL BATCHING

### Static Batching
- Combines static GameObjects at build time
- Requires objects to share the same material
- Mark objects as Static in Inspector
- Increases memory usage (combined mesh stored)

### Dynamic Batching
- Combines small meshes at runtime
- Limited to meshes with < 900 vertex attributes
- Objects must share the same material
- Disabled in some render pipelines

### SRP Batcher
- Different from traditional batching
- Doesn't combine meshes
- Reduces CPU overhead by caching property locations
- Works with different materials (as long as they use the same shader)

### Enabling Batching
```csharp
// Static batching - mark GameObjects as Static
// Dynamic batching - enabled in Player Settings
// SRP Batcher - enabled by default in URP/HDRP
```

### Reducing Draw Calls
- Use atlasing (combine textures)
- Use GPU instancing for repeated objects
- Use static batching for static objects
- Minimize material count
- Use LOD groups to reduce complexity at distance

---

## 19. RENDER PIPELINES

### Built-in Render Pipeline (BIRP)
- Default pipeline in Unity
- Supports Surface Shaders
- Forward and Deferred rendering paths
- Fixed lighting model

### Universal Render Pipeline (URP)
- Optimized for performance across platforms
- Single-pass forward rendering
- Uses Shader Graph or custom HLSL shaders
- SRP Batcher enabled by default
- Does NOT support Surface Shaders
- Does NOT use UnityCG.cginc (use URP shader library instead)

### High Definition Render Pipeline (HDRP)
- High-fidelity graphics for PC/console
- Deferred/Forward+ rendering
- Uses Shader Graph or custom HLSL shaders
- Physical lighting model
- Does NOT support Surface Shaders

### Render Pipeline Comparison
| Feature | BIRP | URP | HDRP |
|---------|------|-----|------|
| Surface Shaders | Yes | No | No |
| Shader Graph | No | Yes | Yes |
| SRP Batcher | No | Yes | Yes |
| Forward Rendering | Yes | Yes | Yes |
| Deferred Rendering | Yes | No | Yes |
| GPU Resident Drawer | No | Yes | Yes |

### Setting Render Pipeline
- Create Render Pipeline Asset (URP/HDRP)
- Assign in Project Settings > Graphics > Scriptable Render Pipeline Settings
- Upgrade materials using Material Upgrade tool

---

## 20. COLOR SPACES

### Gamma Color Space
- Default for legacy projects
- Non-linear color representation
- sRGB encoding
- Simpler but less accurate

### Linear Color Space
- Recommended for new projects
- Physically accurate lighting calculations
- Better color blending and lighting
- Requires sRGB textures and framebuffers

### Setting Color Space
```csharp
// Project Settings > Player > Other Settings > Color Space
// Or via script:
PlayerSettings.colorSpace = ColorSpace.Linear;
```

### Linear vs Gamma Differences
- Lighting calculations are physically correct in Linear
- Color blending is correct in Linear
- HDR values work properly in Linear
- Post-processing effects require Linear
- Textures must be marked as sRGB for proper conversion

### sRGB Texture Sampling
- Textures with color data should be imported as sRGB
- Normal maps, height maps, etc. should NOT be sRGB
- In Linear space, sRGB textures are automatically converted to linear when sampled

---

## 21. TEXTURES AND COMPRESSION

### Texture Types
| Type | Description |
|------|-------------|
| Default | Standard 2D texture |
| Normal Map | Tangent space normal map |
| Editor GUI | UI elements in Editor |
| Sprite | 2D sprite graphics |
| Cursor | Mouse cursor |
| Cookie | Light cookie/cutter |
| Lightmap | Baked lightmap |
| Single Channel | Single channel data (height, AO, etc.) |

### Texture Compression Formats
| Format | Platform | Quality | Size |
|--------|----------|---------|------|
| ASTC | Mobile, modern GPUs | Excellent | Small |
| ETC2 | OpenGL ES 3.0 | Good | Small |
| DXT/BC | Desktop (D3D) | Good | Small |
| PVRTC | iOS (older) | Fair | Very small |
| RGBA Crunched DXT | Desktop | Good | Very small |
| RGBA Crunched ETC2 | Mobile | Good | Very small |
| Uncompressed | All | Best | Large |

### Mip Maps
- Pre-calculated smaller versions of textures
- Reduce aliasing at distance
- Improve cache performance
- Enable for most textures (disable for UI/cookies)

### Texture Import Settings
```csharp
// Via TextureImporter
TextureImporter importer = AssetImporter.GetAtPath(path) as TextureImporter;
importer.textureType = TextureImporterType.Default;
importer.sRGBTexture = true;          // For color textures
importer.mipmapEnabled = true;
importer.textureCompression = TextureImporterCompression.Compressed;
importer.maxTextureSize = 2048;
importer.textureShape = TextureImporterShape.Texture2D;
```

### Texture Arrays
```hlsl
// Declare
UNITY_DECLARE_TEX2DARRAY(_TextureArray);

// Sample (uv.z = array slice index)
half4 color = UNITY_SAMPLE_TEX2DARRAY(_TextureArray, float3(uv, sliceIndex));
```

### 3D Textures
```hlsl
// Declare
sampler3D _Volume;

// Sample
half4 color = tex3D(_Volume, float3(u, v, w));
```

### Cubemaps
```hlsl
// Declare
samplerCUBE _Cube;
// or
UNITY_DECLARE_TEXCUBE(_Cube);

// Sample
half4 color = texCUBE(_Cube, direction);
```

---

## 22. CAMERAS AND RENDERING

### Camera Types
- **Perspective**: Realistic 3D perspective
- **Orthographic**: No perspective distortion (2D/technical)

### Physical Camera Properties
```csharp
Camera cam = GetComponent<Camera>();
cam.fieldOfView = 60.0f;           // Field of view in degrees
cam.nearClipPlane = 0.3f;          // Near clipping plane
cam.farClipPlane = 1000.0f;        // Far clipping plane
cam.orthographic = false;          // Perspective vs orthographic
cam.orthographicSize = 5.0f;       // Orthographic half-size

// Physical camera
cam.usePhysicalProperties = true;
cam.focalLength = 35.0f;           // mm
cam.sensorSize = new Vector2(36, 24);  // mm (full frame)
cam.gateFit = Camera.GateFitMode.Horizontal;  // Horizontal, Vertical, Fill, Overflow, None
cam.lensShift = new Vector2(0, 0); // Lens shift for perspective correction
```

### Camera Output
```csharp
// Render to texture
cam.targetTexture = renderTexture;

// Multiple cameras
// Use Camera.depth to control rendering order
// Use Camera.rect for viewport rects (split screen)
// Use Camera.cullingMask for layer-based rendering
```

### Camera Rays
```csharp
// Screen point to ray
Ray ray = cam.ScreenPointToRay(Input.mousePosition);

// Viewport point to ray
Ray ray = cam.ViewportPointToRay(new Vector3(0.5f, 0.5f, 0));

// Screen to world point
Vector3 worldPos = cam.ScreenToWorldPoint(new Vector3(x, y, distance));
```

---

## 23. LEVEL OF DETAIL (LOD)

### LOD Group
```csharp
LODGroup lodGroup = GetComponent<LODGroup>();
LOD[] lods = new LOD[3];
lods[0] = new LOD(0.6f, new Renderer[] { highDetailRenderer });
lods[1] = new LOD(0.3f, new Renderer[] { mediumDetailRenderer });
lods[2] = new LOD(0.1f, new Renderer[] { lowDetailRenderer });
lodGroup.SetLODs(lods);
lodGroup.RecalculateBounds();
```

### Shader LOD
```shaderlab
SubShader {
    LOD 300  // Higher = more detailed
    // ...
}

// Set maximum LOD from script
Shader.globalMaximumLOD = 500;
material.shaderMaximumLOD = 300;
```

### LOD Fade
```hlsl
// unity_LODFade.x = fade amount (0..1)
// unity_LODFade.y = fade quantized to 16 levels
half fade = unity_LODFade.x;
```

---

## 24. POST-PROCESSING

### Built-in Post-Processing
- Post-processing stack (legacy)
- Volume framework (URP/HDRP)

### Common Effects
- Bloom
- Color Grading
- Vignette
- Depth of Field
- Motion Blur
- Ambient Occlusion
- Chromatic Aberration
- Film Grain

### URP Post-Processing
- Uses Volume component
- Profile-based configuration
- Global or local volumes
- Blending between volumes

---

## 25. GRAPHICS APIs AND CONFIGURATION

### Supported Graphics APIs
| API | Platforms |
|-----|-----------|
| Direct3D 11 | Windows |
| Direct3D 12 | Windows (experimental) |
| Metal | macOS, iOS, tvOS |
| Vulkan | Android, Windows, Linux |
| OpenGL Core | Windows, macOS, Linux |
| OpenGL ES 3.0+ | Android, iOS |
| WebGL 2.0 | Web |
| WebGPU | Web (experimental) |
| PlayStation 4/5 | Consoles |
| Xbox Series | Console |
| Nintendo Switch | Console |

### Graphics Jobs
- Offloads rendering work to worker threads
- Can improve CPU performance
- Configure in Player Settings
- D3D12 and Metal support graphics jobs

### Device Filtering
- D3D12 and Vulkan support device filtering
- Create device filtering assets to select specific GPUs
- Useful for multi-GPU systems

---

## 26. PERFORMANCE OPTIMIZATION

### Shader Compilation
- **Asynchronous compilation**: Shaders compile in background (prevents stuttering)
- **Shader prewarming**: Compile shaders before they're needed
- **Shader variant stripping**: Remove unused variants from builds
- **Shader caching**: Cache compiled shaders between sessions

### Reducing Shader Variants
```hlsl
// Use shader_feature instead of multi_compile when possible
#pragma shader_feature _ USE_TEXTURE  // Strips unused variants

// Skip specific variants
#pragma skip_variants KEYWORD1 KEYWORD2

// Limit instancing variants
#pragma instancing_options forcemaxcount:64
```

### PSO (Pipeline State Object) Caching
- PSOs are compiled shader + render state combinations
- Warm up PSOs at load time to prevent runtime stuttering
- Use `Shader.WarmupAllShaders()` for prewarming

### Draw Call Optimization
1. **SRP Batcher**: Fastest for SRP projects
2. **GPU Instancing**: For many copies of same mesh
3. **Static Batching**: For static objects sharing materials
4. **Dynamic Batching**: For small dynamic meshes (limited)

### Memory Optimization
- Use compressed textures
- Use appropriate texture sizes
- Enable mipmaps for 3D textures
- Use texture atlases
- Reduce shader variant count

### Profiling Tools
- **Frame Debugger**: Step through rendering events
- **Profiler**: CPU/GPU timing, memory, rendering module
- **GPU Profiler**: Platform-specific GPU profiling
- **RenderDoc**: External GPU debugging tool
- **PIX**: DirectX debugging tool

---

## 27. COMMON GOTCHAS AND PITFALLS

### Shader Compilation
- Surface Shaders only work in Built-in Render Pipeline
- `UnityCG.cginc` is Built-in RP specific - don't use in URP
- `#pragma target` affects which hardware can run the shader
- Geometry shaders don't work on Metal
- Tessellation requires shader model 4.6+

### Materials
- `renderer.material` creates a copy - use `sharedMaterial` for shared
- `MaterialPropertyBlock` is more efficient than material copies for per-instance data
- SRP Batcher requires ALL per-material properties in single `CBUFFER_START(UnityPerMaterial)`

### Textures
- sRGB flag must be correct: ON for colors, OFF for data (normals, height, etc.)
- Mipmaps increase memory by ~33% but improve quality and performance
- Texture arrays require all slices to be the same size

### Rendering
- Blending disables Early-Z/Hi-Z optimizations on GPU
- Transparent objects must be sorted back-to-front
- ZWrite Off for transparent objects to avoid depth artifacts
- Use `ZTest LEqual` (default) for most cases

### Color Spaces
- Linear color space is recommended for new projects
- All color textures must be marked sRGB for correct conversion
- Normal maps and data textures should NOT be sRGB

### GPU Instancing
- Only works with `Graphics.DrawMeshInstanced` or compatible renderers
- All instanced objects must use the same material
- Per-instance properties require `UNITY_INSTANCING_BUFFER_START/END`
- SkinnedMeshRenderer instancing requires additional setup

---

## 28. UNITY 6000-SPECIFIC FEATURES

### Unity 6.3 LTS (6000.3) Key Points
- **WebGPU** support added as experimental graphics API
- **DXC compiler** support for DX12, Vulkan, Metal with 16-bit and wave operations
- **Device feature detection** keywords for runtime GPU feature selection:
  - `UNITY_DEVICE_SUPPORTS_NATIVE_16BIT` - 16-bit data type support
  - `UNITY_DEVICE_SUPPORTS_WAVE_*` - Wave operation support (8/16/32/64/128)
- **PackageRequirements** block in ShaderLab for SubShader/Pass-level package dependencies
- **Integer** property type (backed by real integer, not float like legacy `Int`)
- **Vector** property with component count: `Vector, 2` shows only x,y fields
- **Caching Shader Preprocessor** supports `#pragma once`
- **Render Graph** architecture in Unity 6 for improved rendering performance
- **GPU Resident Drawer** for efficient rendering of many objects (URP/HDRP)

### New Shader Compilation Features
```hlsl
// Runtime GPU feature detection (DX12, Vulkan, Metal)
#pragma multi_compile _ UNITY_DEVICE_SUPPORTS_NATIVE_16BIT
#pragma multi_compile _ UNITY_DEVICE_SUPPORTS_WAVE_32

#if defined(UNITY_DEVICE_SUPPORTS_NATIVE_16BIT)
    // Use 16-bit types
#else
    // Fallback to 32-bit
#endif
```

---

## 29. QUICK REFERENCE CHEATSHEET

### Minimal Unlit Shader
```shaderlab
Shader "Custom/Unlit" {
    Properties { _MainTex("Texture", 2D) = "white" {} }
    SubShader {
        Pass {
            HLSLPROGRAM
            #pragma vertex vert
            #pragma fragment frag
            #include "UnityCG.cginc"
            
            struct appdata { float4 vertex : POSITION; float2 uv : TEXCOORD0; };
            struct v2f { float2 uv : TEXCOORD0; float4 vertex : SV_POSITION; };
            
            sampler2D _MainTex;
            v2f vert(appdata v) {
                v2f o;
                o.vertex = UnityObjectToClipPos(v.vertex);
                o.uv = v.uv;
                return o;
            }
            fixed4 frag(v2f i) : SV_Target { return tex2D(_MainTex, i.uv); }
            ENDHLSL
        }
    }
}
```

### Minimal Lit Surface Shader
```shaderlab
Shader "Custom/Lit" {
    Properties { _Color("Color", Color) = (1,1,1,1) _MainTex("Texture", 2D) = "white" {} }
    SubShader {
        Tags { "RenderType"="Opaque" }
        CGPROGRAM
        #pragma surface surf Lambert
        struct Input { float2 uv_MainTex; };
        fixed4 _Color;
        sampler2D _MainTex;
        void surf(Input IN, inout SurfaceOutput o) {
            o.Albedo = tex2D(_MainTex, IN.uv_MainTex).rgb * _Color.rgb;
        }
        ENDCG
    }
}
```

### Standard Transparency Setup
```shaderlab
// Opaque (default)
// No blending, ZWrite On

// Cutout (alpha tested)
Tags { "RenderType"="TransparentCutout" }
// clip(alpha - cutoff) in fragment shader

// Fade (soft transparency)
Tags { "RenderType"="Transparent" }
Blend SrcAlpha OneMinusSrcAlpha
ZWrite Off

// Transparent (physical)
Tags { "RenderType"="Transparent" }
Blend One OneMinusSrcAlpha
ZWrite Off
```

---

*This summary was extracted from Unity 6000.3.9f1 (Unity 6.3 LTS) local documentation, built on 2026-03-29, job ID 65670558.*
