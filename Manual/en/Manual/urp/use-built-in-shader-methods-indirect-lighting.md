# Use indirect lighting in a custom URP shader

To use indirect lighting in a custom Universal **Render Pipeline** (URP) **shader**, follow these steps:

1. Add an `#include` line inside the `HLSLPROGRAM` block in your shader file. Refer to the sections on this page for specific include file paths.
2. Use any of the methods from the following sections.

## Use reflection probes in a custom URP shader

To use **reflection probes** or the **skybox**, use the following `#include` line:

```
#include "Packages/com.unity.render-pipelines.universal/ShaderLibrary/GlobalIllumination.hlsl"
```

| **Method** | **Syntax** | **Description** |
| --- | --- | --- |
| `GlossyEnvironmentReflection` | `half3 GlossyEnvironmentReflection(half3 reflectVector, float3 positionWS, half perceptualRoughness, half occlusion, float2 normalizedScreenSpaceUV)` | Samples reflection probes if the object is within a reflection probe volume. Samples the skybox otherwise. The function supports [reflection probe blending](lighting/reflection-probes-introduction.md). The `perceptualRoughness` argument selects the mip level of the sampled texture. The higher the value, the more blurred the reflections look. The `occlusion` argument is the reflection probe attenuation factor. The `normalizedScreenSpaceUV` argument is the normalized screen space position of the fragment. |

The `GlossyEnvironmentReflection` function has the following simpler overload, which does not support blending:

```
half3 GlossyEnvironmentReflection(half3 reflectVector, half perceptualRoughness, half occlusion)
```

## Sample Adaptive Probe Volumes

To sample Adaptive Probe Volumes (APV), use the following `#include` line:

```
#include "Packages/com.unity.render-pipelines.core/Runtime/Lighting/ProbeVolume/ProbeVolume.hlsl"
```

| **Method** | **Syntax** | **Description** |
| --- | --- | --- |
| `EvaluateAdaptiveProbeVolume` | `void EvaluateAdaptiveProbeVolume(in float3 posWS, in float3 normalWS, in float3 viewDir, in float2 positionSS, in uint renderingLayer, out float3 bakeDiffuseLighting)` | Samples lighting from Adaptive Probe Volumes. The `positionSS` argument is the screen space position of a fragment that is being rendered. The `renderingLayer` argument is the [rendering layer mask](probevolumes-troubleshoot-light-leaks.md#layers). |

## Sample light probes

To sample light probes, use the following `#include` line:

```
#include "Packages/com.unity.render-pipelines.core/ShaderLibrary/AmbientProbe.hlsl"
```

| **Method** | **Syntax** | **Description** |
| --- | --- | --- |
| `SampleSH` | `real3 SampleSH(real3 normalWS)` | Sample [light probes](../LightProbes-landing.md). |

## Example

The following URP shader code computes diffuse and specular **ambient light** using the `GlossyEnvironmentReflection` function.

```
float3 CustomLightingAmbient(float3 BaseColor, 
                            float3 NormalWS, 
                            float Metallic, 
                            float Smoothness, 
                            float Reflectance, 
                            float AmbientOcclusion, 
                            float PositionWS, 
                            float2 screenspaceUV, 
                            float3 ViewDirWS)
{
    // Diffuse ambient lighting.
    // Sample the sky probe with high roughness and use a normal vector instead of
    // a reflection vector to approximate ambient light
    float3 DiffuseAmbient = GlossyEnvironmentReflection(NormalWS, PositionWS, 1, 1, screenspaceUV);
    // If the surface is metallic, set ambient to zero, otherwise multiply by BaseColor
    DiffuseAmbient *= lerp(BaseColor, float3(0,0,0), Metallic);

    // Specular ambient lighting - reflections
    // Surfaces are more reflective at glancing angles. Compute the fresnel value for reflectance.
    float Fresnel = pow((1 - dot(NormalWS, ViewDirWS)), 5);
    // Reflectance is 100% at the glancing angle
    Reflectance = lerp(Reflectance, 1, Fresnel);
    // If the object is metal, use the base color for reflectance instead.
    float3 ReflectionColor = lerp(float3(Reflectance, Reflectance, Reflectance), BaseColor, Metallic);
    // Compute the reflection vector and use it to sample the sky probe for reflections.
    float3 ReflectionVector = reflect(-ViewDirWS, NormalWS);
    float3 SpecularAmbient = GlossyEnvironmentReflection(ReflectionVector, PositionWS, 1 - Smoothness, 1, screenspaceUV);
    
    return (DiffuseAmbient + SpecularAmbient) * AmbientOcclusion;
}
```

## Shader graph custom function node example

[Shader Graph Feature Examples](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Shader-Graph-Sample-Feature-Examples.html) contain examples of custom lighting models.

To learn how to use the `GlossyEnvironmentReflection` function in a custom function node:

1. In the example **scene**, open the **CustomLightingSimple** shader graph.
2. Navigate to the **SampleReflectionProbes** sub-graph.
3. Inspect the **URPReflectionProbe** node. The node contains the HLSL code that uses the `GlossyEnvironmentReflection` function:

   ```
   #ifdef SHADERGRAPH_PREVIEW
       reflection = float3(0,0,0);
   #else
       reflection = GlossyEnvironmentReflection(reflectVector, positionWS, roughness, 1.0h, screenspaceUV);
   #endif
   ```

## Additional resources

* [Writing custom shaders](writing-custom-shaders-urp.md)
* [Upgrade custom shaders for URP compatibility](urp-shaders/birp-urp-custom-shader-upgrade-guide.md)
* [Reflection probes in URP](lighting/reflection-probes-introduction.md)
* [HLSL in Unity](../writing-shader-writing-shader-programs-hlsl.md)
* [Adaptive Probe Volumes (APV) in URP](probevolumes.md)
* [Light probes](../LightProbes-landing.md)
* [Shader Graph Feature Examples](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Shader-Graph-Sample-Feature-Examples.html)