# SubShader block in ShaderLab reference

To define a SubShader in **ShaderLab**, you use a `SubShader` block. This page contains information on using SubShader blocks.

Inside the `SubShader` block, you can:

* Assign a **LOD** (level of detail) value to the SubShader, using the `LOD` block. See [assigning a LOD value to a SubShader](SL-ShaderLOD.md).
* Assign key-value pairs of data to the SubShader, using the `Tags` block. See [ShaderLab: assigning tags to a SubShader](SL-SubShaderTags.md).
* Add GPU instructions or **shader** code to the SubShader, using ShaderLab commands. See [ShaderLab: using commands](SL-Reference.md).
* Define one or more Passes, using the `Pass` block. See [ShaderLab: defining a Pass](SL-Pass.md).
* Specify package requirements using the `PackageRequirements` block. This makes Unity only run the SubShader if the required packages are installed. See [ShaderLab: specifying package requirements](SL-PackageRequirements.md).

## Render pipeline compatibility

| **Feature name** | **Universal **Render Pipeline** (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** | **Built-in Render Pipeline** |
| --- | --- | --- | --- | --- |
| **ShaderLab: SubShader block** | Yes | Yes | Yes | Yes |

## Syntax

| **Signature** | **Function** |
| --- | --- |
| `SubShader` `{`     `<optional: LOD>`     `<optional: tags>`     `<optional: commands>`     `<One or more Pass definitions>` `}` | Defines a SubShader.  You can define as many Passes as you like within a SubShader. |

## Additional resources

* [Add a subshader in a custom shader](writing-shader-create-subshader-object.md)
* [Add a shader pass in a custom shader](writing-shader-create-shader-pass.md)