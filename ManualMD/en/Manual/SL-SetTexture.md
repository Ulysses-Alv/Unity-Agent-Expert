# ShaderLab: legacy texture combining

**Note**: The **ShaderLab** functionality on this page is legacy, and is documented for backwards compatibility only. If your **shader** source file includes HLSL code, Unity ignores these commands completely. If your shader source file does not include HLSL code, Unity compiles these commands into regular shader programs on import.

## Render pipeline compatibility

| **Feature name** | **Built-in **Render Pipeline**** | **Universal Render Pipeline (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** |
| --- | --- | --- | --- | --- |
| **Legacy texture combining** | Yes | No | No | No |

# Overview

After the basic vertex lighting has been calculated, textures are applied. In ShaderLab this is done using **SetTexture** command.

Fixed function texturing is the place to do old-style combiner effects. You can have multiple SetTexture commands inside a pass - all textures are applied in sequence, like layers in a painting program. SetTexture commands must be placed at the end of a [Pass](SL-Pass.md).

## Syntax

```
SetTexture [TextureName] {Texture Block}
```

Assigns a texture. *TextureName* must be defined as a texture property. How to apply the texture is defined inside the *TextureBlock*.

The texture block controls how the texture is applied. Inside the texture block can be up to two commands: `combine` and `constantColor`.

## Texture block `combine` command

`combine` *src1* \* *src2*: Multiplies src1 and src2 together. The result will be darker than either input.

`combine` *src1* + *src2*: Adds src1 and src2 together. The result will be lighter than either input.

`combine` *src1* - *src2*: Subtracts src2 from src1.

`combine` *src1* `lerp` (*src2*) *src3*: Interpolates between src3 and src1, using the alpha of src2. Note that the interpolation is opposite direction: src1 is used when alpha is one, and src3 is used when alpha is zero.

`combine` *src1* \* *src2* + *src3*: Multiplies src1 with the alpha component of src2, then adds src3.

All the **src** properties can be either one of *previous*, *constant*, *primary* or *texture*.

* **Previous** is the the result of the previous SetTexture.
* **Primary** is the color from the [lighting calculation](SL-Material.md) or the vertex color if it is [bound](SL-BindChannels.md).
* **Texture** is the color of the texture specified by *TextureName* in the SetTexture (see above).
* **Constant** is the color specified in **ConstantColor**.

Modifiers:

* The formulas specified above can optionally be followed by the keywords **Double** or ****Quad**** to make the resulting color 2x or 4x as bright.
* All the **src** properties, except `lerp` argument, can optionally be preceded by **one -** to make the resulting color negated.
* All the **src** properties can be followed by **alpha** to take only the alpha channel.

## Texture block `constantColor` command

**ConstantColor color:** Defines a constant color that can be used in the combine command.

## Functionality removed in Unity 5.0

Unity versions before 5.0 did support texture coordinate transformations with a `matrix` command inside a texture block. If you need this functionality now, consider rewriting your shader as a [programmable shader](writing-shader-writing-shader-programs-hlsl.md) instead, and do the UV transformation in the **vertex shader**.

Similarly, 5.0 removed signed add (`a+-b`), multiply signed add (`a*b+-c`), multiply subtract (`a*b-c`) and dot product (`dot3`, `dot3rgba`) texture combine modes. If you need them, do the math in the **pixel** shader instead.

## Details

Before [fragment programs](writing-shader-writing-shader-programs-hlsl.md) existed, older graphics cards used a layered approach to textures. The textures are applied one after each other, modifying the color that will be written to the screen. For each texture, the texture is typically combined with the result of the previous operation. These days it is advisable to use actual fragment programs.

Note that each texture stage may or might not be clamped to 0..1 range, depending on the platform. This might affect SetTexture stages that can produce values higher than 1.0.

## Separate Alpha & Color computation

By default, the combiner formula is used for calculating both the RGB and alpha component of the color. Optionally, you can specify a separate formula for the alpha calculation. This looks like this:

```
SetTexture [_MainTex] { combine previous * texture, previous + texture }
```

Here, we multiply the RGB colors and add the alpha.

## Specular highlights

By default the **primary** color is the sum of the diffuse, ambient and **specular colors** (as defined in the [Lighting calculation](SL-Material.md)). If you specify **SeparateSpecular On** in the pass options, the specular color will be added in *after* the combiner calculation, rather than before. This is the default behavior of the built-in VertexLit shader.

## Graphics hardware support

Modern graphics cards with [fragment shader](writing-shader-writing-shader-programs-hlsl.md) support (“shader model 2.0” on desktop) support all **SetTexture** modes and at least 4 texture stages (many of them support 8). If you’re running on really old hardware (made before 2003 on PC, or before iPhone3GS on mobile), you might have as low as two texture stages. The shader author should write separate [SubShaders](SL-SubShader.md) for the cards they want to support.

## Examples

### Alpha Blending Two Textures

This small examples takes two textures. First it sets the first combiner to just take the **\_MainTex**, then is uses the alpha channel of **\_BlendTex** to fade in the RGB colors of **\_BlendTex**

```
Shader "Examples/2 Alpha Blended Textures" {
    Properties {
        _MainTex ("Base (RGB)", 2D) = "white" {}
        _BlendTex ("Alpha Blended (RGBA) ", 2D) = "white" {}
    }
    SubShader {
        Pass {
            // Apply base texture
            SetTexture [_MainTex] {
                combine texture
            }
            // Blend in the alpha texture using the lerp operator
            SetTexture [_BlendTex] {
                combine texture lerp (texture) previous
            }
        }
    }
}
```

### Alpha Controlled Self-illumination

This shader uses the alpha component of the **\_MainTex** to decide where to apply lighting. It does this by applying the texture to two stages; In the first stage, the alpha value of the texture is used to blend between the vertex color and solid white. In the second stage, the RGB values of the texture are multiplied in.

```
Shader "Examples/Self-Illumination" {
    Properties {
        _MainTex ("Base (RGB) Self-Illumination (A)", 2D) = "white" {}
    }
    SubShader {
        Pass {
            // Set up basic white vertex lighting
            Material {
                Diffuse (1,1,1,1)
                Ambient (1,1,1,1)
            }
            Lighting On

            // Use texture alpha to blend up to white (= full illumination)
            SetTexture [_MainTex] {
                constantColor (1,1,1,1)
                combine constant lerp(texture) previous
            }
            // Multiply in texture
            SetTexture [_MainTex] {
                combine previous * texture
            }
        }
    }
}
```

We can do something else for free here, though; instead of blending to solid white, we can add a self-illumination color and blend to that. Note the use of **ConstantColor** to get a \_SolidColor from the properties into the texture blending.

```
Shader "Examples/Self-Illumination 2" {
    Properties {
        _IlluminCol ("Self-Illumination color (RGB)", Color) = (1,1,1,1)
        _MainTex ("Base (RGB) Self-Illumination (A)", 2D) = "white" {}
    }
    SubShader {
        Pass {
            // Set up basic white vertex lighting
            Material {
                Diffuse (1,1,1,1)
                Ambient (1,1,1,1)
            }
            Lighting On

            // Use texture alpha to blend up to white (= full illumination)
            SetTexture [_MainTex] {
                // Pull the color property into this blender
                constantColor [_IlluminCol]
                // And use the texture's alpha to blend between it and
                // vertex color
                combine constant lerp(texture) previous
            }
            // Multiply in texture
            SetTexture [_MainTex] {
                combine previous * texture
            }
        }
    }
}
```

And finally, we take all the lighting properties of the vertexlit shader and pull that in:

```
Shader "Examples/Self-Illumination 3" {
    Properties {
        _IlluminCol ("Self-Illumination color (RGB)", Color) = (1,1,1,1)
        _Color ("Main Color", Color) = (1,1,1,0)
        _SpecColor ("Spec Color", Color) = (1,1,1,1)
        _Emission ("Emmisive Color", Color) = (0,0,0,0)
        _Shininess ("Shininess", Range (0.01, 1)) = 0.7
        _MainTex ("Base (RGB)", 2D) = "white" {}
    }

    SubShader {
        Pass {
            // Set up basic vertex lighting
            Material {
                Diffuse [_Color]
                Ambient [_Color]
                Shininess [_Shininess]
                Specular [_SpecColor]
                Emission [_Emission]
            }
            Lighting On

            // Use texture alpha to blend up to white (= full illumination)
            SetTexture [_MainTex] {
                constantColor [_IlluminCol]
                combine constant lerp(texture) previous
            }
            // Multiply in texture
            SetTexture [_MainTex] {
                combine previous * texture
            }
        }
    }
}
```