# Pass tags in ShaderLab reference

This page contains information on using a `Tags` block in your **ShaderLab** code to assign tags to a Pass. It also contains information on using the `LightMode` tag.

For information on how a **Shader** object works, and the relationship between **Shader objects**, SubShaders and Passes, see [Shader object fundamentals](Shaders.md).

## Render pipeline compatibility

| **Feature name** | **Universal **Render Pipeline** (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** | **Built-in Render Pipeline** |
| --- | --- | --- | --- | --- |
| **ShaderLab: Pass Tags block** | Yes | Yes | Yes | Yes |
| **ShaderLab: LightMode Pass tag** | Yes | Yes | Yes | Yes |

## Syntax

| **Signature** | **Function** |
| --- | --- |
| `Tags {"<name1>" = "<value1>" "<name2>" = "<value2>"}` | Applies the given tags to the Pass.  You can define as many tags as you like. |

### LightMode tag

| **Signature** | **Function** |
| --- | --- |
| “LightMode” = “[value]” | Sets the LightMode value for this Pass. |

Valid values for this tag depend on the render pipeline.

#### LightMode tag valid values

These are the valid values for the `LightMode` Pass tag in the Built-in Render Pipeline. For more information on the LightMode tag, see [ShaderLab: using Pass tags](SL-PassTags.md).

| **Value** | **Function** | **Supported in Built-In Render Pipeline** | **Supported in HDRP** | **Supported in URP** |
| --- | --- | --- | --- | --- |
| `Always` | Always rendered; does not apply any lighting. This is the default value. | Yes | No | No |
| `ForwardBase` | Used in **Forward rendering**; applies ambient, main directional light, vertex/SH lights and **lightmaps**. | Yes | No | No |
| `ForwardAdd` | Used in Forward rendering; applies additive per-pixel lights, one Pass per light. | Yes | No | No |
| `Deferred` | Used in **Deferred Shading**; renders G-buffer. | Yes | No | No |
| `ShadowCaster` | Renders object depth into the shadowmap or a depth texture. | Yes | Yes | Yes |
| `MotionVectors` | Used to calculate per-object motion vectors. | Yes | Yes | Yes |
| `Vertex` | Used in legacy Vertex Lit rendering when the object is not lightmapped; applies all vertex lights. | Yes | No | No |
| `VertexLMRGBM` | Used in legacy Vertex Lit rendering when the object is lightmapped, and on platforms where the lightmap is RGBM encoded (PC & console). | Yes | No | No |
| `VertexLM` | Used in legacy Vertex Lit rendering when the object is lightmapped, and on platforms where lightmap is double-LDR encoded (mobile platforms). | Yes | No | No |
| `Meta` | This Pass is not used during regular rendering, only for lightmap baking or **Enlighten** Realtime **Global Illumination**. For more information, see [Lightmapping and shaders](MetaPass.md). | Yes | Yes | Yes |

### PassFlags tag

In the Built-in Render Pipeline, use the `PassFlags` Pass tag to specify what data Unity provides to the Pass.

| **Value** | **Function** |
| --- | --- |
| OnlyDirectional | Valid only in the Built-in Render Pipeline, when the rendering path is set to Forward, in a Pass with a `LightMode` tag value of `ForwardBase`.  Unity provides only the main directional light and ambient/light probe data to this Pass. This means that data of non-important lights is not passed into vertex-light or spherical harmonics shader variables. See [Forward rendering path](RenderTech-ForwardRendering.md) for details. |

#### Example

```
Shader "Examples/ExamplePassFlag"
{
    SubShader
    {
        Pass
        {    
              Tags { "LightMode" = "ForwardBase" "PassFlags" = "OnlyDirectional" }
            
              // The rest of the code that defines the Pass goes here.
        }
    }
}
```

### RequireOptions tag

In the Built-in Render Pipeline, the `RequireOptions` Pass tag enables or disables a Pass based on **project settings**.

| **Value** | **Function** |
| --- | --- |
| `SoftVegetation` | Render this Pass only if [QualitySettings-softVegetation](../ScriptReference/QualitySettings-softVegetation.md) is enabled. |

## Additional resources

* [Configure if or when Unity uses a shader](writing-shader-tags.md)