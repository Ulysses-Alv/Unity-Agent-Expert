# ShaderLab: legacy lighting

**Note**: The **ShaderLab** functionality on this page is legacy, and is documented for backwards compatibility only. If your **shader** source file includes HLSL code, Unity ignores these commands completely. If your shader source file does not include HLSL code, Unity compiles these commands into regular shader programs on import.

## Render pipeline compatibility

| **Feature name** | **Built-in **Render Pipeline**** | **Universal Render Pipeline (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** |
| --- | --- | --- | --- | --- |
| **Legacy lighting** | Yes | No | No | No |

## Overview

The material and lighting parameters are used to control the built-in vertex lighting. Vertex lighting is the standard Direct3D/OpenGL lighting model that is computed for each vertex. **Lighting on** turns it on. Lighting is affected by **Material** block, **ColorMaterial** and **SeparateSpecular** commands.

Vertex Coloring & Lighting is the first effect to be calculated for any rendered geometry. It operates on the vertex level, and calculates the base color that is used before textures are applied.

## Syntax

The top level commands control whether to use fixed function lighting or not, and some configuration options. The main setup is in the **Material Block**, detailed further below.

### Color

```
    Color color
```

Sets the object to a solid color. A color is either four RGBA values in parenthesis, or a color property name in square brackets.

### Material

```
    Material {Material Block}
```

The Material block is used to define the material properties of the object.

### Lighting

```
    Lighting On | Off
```

For the settings defined in the Material block to have any effect, you must enable Lighting with the *Lighting On* command. If lighting is off instead, the color is taken straight from the *Color* command.

### SeparateSpecular

```
    SeparateSpecular On | Off
```

This command makes specular lighting be added to the end of the shader pass, so specular lighting is unaffected by texturing. Only has effect when *Lighting On* is used.

### ColorMaterial

```
    ColorMaterial AmbientAndDiffuse | Emission
```

Makes per-vertex color be used instead of the colors set in the material. **AmbientAndDiffuse** replaces Ambient and Diffuse values of the material; **Emission** replaces Emission value of the material.

### Material Block

This contains settings for how the material reacts to the light. Any of these properties can be left out, in which case they default to black (i.e. have no effect).

**Diffuse color:** The diffuse color component. This is an object’s base color.

**Ambient color:** The ambient color component. This is the color the object has when it’s hit by the **ambient light** set in the [Lighting Window](lighting-window.md).

****Specular color**:** The color of the object’s specular highlight.

**Shininess number:** The sharpness of the highlight, between 0 and 1. At 0 you get a huge highlight that looks a lot like diffuse lighting, at 1 you get a tiny speck.

**Emission color:** The color of the object when it is not hit by any light.

The full color of lights hitting the object is:

**Ambient** \* [Lighting Window’s Ambient Intensity setting](lighting-window.md) +
(Light Color \* **Diffuse** + Light Color \* **Specular**) + **Emission**

The light parts of the equation (within parenthesis) is repeated for all lights that hit the object.

Typically you want to keep the Diffuse and Ambient colors the same (all built-in shaders do this).

## Examples

Always render object in pure red:

```
Shader "Solid Red" {
    SubShader {
        Pass { Color (1,0,0,0) }
    }
}
```

Basic Shader that colors the object white and applies vertex lighting:

```
Shader "VertexLit White" {
    SubShader {
        Pass {
            Material {
                Diffuse (1,1,1,1)
                Ambient (1,1,1,1)
            }
            Lighting On
        }
    }
}
```

An extended version that adds material color as a property visible in Material **Inspector**:

```
Shader "VertexLit Simple" {
    Properties {
        _Color ("Main Color", COLOR) = (1,1,1,1)
    }
    SubShader {
        Pass {
            Material {
                Diffuse [_Color]
                Ambient [_Color]
            }
            Lighting On
        }
    }
}
```

And finally, a full fledged vertex-lit shader (see also [SetTexture](SL-SetTexture.md) reference page):

```
Shader "VertexLit" {
    Properties {
        _Color ("Main Color", Color) = (1,1,1,0)
        _SpecColor ("Spec Color", Color) = (1,1,1,1)
        _Emission ("Emmisive Color", Color) = (0,0,0,0)
        _Shininess ("Shininess", Range (0.01, 1)) = 0.7
        _MainTex ("Base (RGB)", 2D) = "white" {}
    }
    SubShader {
        Pass {
            Material {
                Diffuse [_Color]
                Ambient [_Color]
                Shininess [_Shininess]
                Specular [_SpecColor]
                Emission [_Emission]
            }
            Lighting On
            SeparateSpecular On
            SetTexture [_MainTex] {
                Combine texture * primary DOUBLE, texture * primary
            }
        }
    }
}
```