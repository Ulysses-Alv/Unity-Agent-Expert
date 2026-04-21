# LOD directive in ShaderLab reference

This page contains information on using a `LOD` block in your **ShaderLab** code to assign a **LOD** (level of detail) value to a SubShader.

## Render pipeline compatibility

| **Feature name** | **Universal **Render Pipeline** (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** | **Built-in Render Pipeline** |
| --- | --- | --- | --- | --- |
| **ShaderLab: SubShader LOD block** | Yes | Yes | Yes | Yes |

## Syntax

In ShaderLab, you assign a LOD value to a SubShader by placing a `LOD` block inside a `SubShader` block.

| **Signature** | **Function** |
| --- | --- |
| LOD [value] | Assigns the given LOD value to the SubShader. |

## LOD values for Unity’s built-in shaders

In the Built-in Render Pipeline, Unity’s built-in **shaders** have the following LOD values:

| **LOD value** | **Shader name** |
| --- | --- |
| 100 | Unlit/Texture Unlit/Color Unlit/Transparent Unlit/Transparent Cutout |
| 300 | Standard Standard (Specular Setup) Autodesk Interactive |

## LOD values for legacy shaders

In the Built-in Render Pipeline, Unity’s built-in [legacy shaders](Built-inShaderGuide.md) have the following LOD values:

| **LOD value** | **Shader name** |
| --- | --- |
| 100 | VertexLit |
| 150 | Decal Reflective VertexLit |
| 200 | Diffuse |
| 250 | Diffuse Detail Reflective Bumped Unlit Reflective Bumped VertexLit |
| 300 | Bumped Specular |
| 400 | Bumped Specular |
| 500 | Parallax |
| 600 | Parallax Specular |

## Additional resources

* [Optimize shaders](SL-ShaderPerformance.md)