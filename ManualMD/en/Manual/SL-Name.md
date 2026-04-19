# Name directive in ShaderLab reference

This page contains information on using a `Name` block in your **ShaderLab** code to assign a name to a Pass. For information on defining Passes, see [ShaderLab: defining a Pass](SL-Pass.md). For information on how a **Shader** object works, and the relationship between **Shader objects**, SubShaders and Passes, see [Shader object fundamentals](Shaders.md).

## Render pipeline compatibility

| **Feature name** | **Universal **Render Pipeline** (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** | **Built-in Render Pipeline** |
| --- | --- | --- | --- | --- |
| **ShaderLab: Name block** | Yes | Yes | Yes | Yes |

## Syntax

| **Signature** | **Function** |
| --- | --- |
| `Name "<name>"` | Sets the name of the Pass. |

## Additional resources

* [Add a shader pass in a custom shader](writing-shader-create-shader-pass.md)