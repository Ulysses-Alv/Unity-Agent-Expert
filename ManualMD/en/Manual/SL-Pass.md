# Pass block in ShaderLab reference

To define a Pass in **ShaderLab**, use a `Pass` block. This page contains information on using `Pass` blocks. For information on how a **Shader** object works, and the relationship between **Shader objects**, SubShaders and Passes, see [Shader object fundamentals](Shaders.md).

Inside the `Pass` block, you can:

* Assign a name to the Pass, using a Name block. See [ShaderLab: assigning a name to a Pass](SL-Name.md).
* Assign key-value pairs of data to the Pass, using a Tags block. See [ShaderLab: assigning tags to a Pass](SL-PassTags.md).
* Perform operations using ShaderLab commands. See [ShaderLab: using commands](SL-Reference.md).
* Add shader code to the Pass, using a shader code block. See [ShaderLab: shader code blocks](shader-shaderlab-code-blocks.md).
* Specify package requirements using the `PackageRequirements` block. This makes Unity only run the Pass if the required packages are installed. See [ShaderLab: specifying package requirements](SL-PackageRequirements.md).

## Render pipeline compatibility

| **Feature name** | **Universal **Render Pipeline** (URP)** | **High Definition Render Pipeline (HDRP)** | **Custom SRP** | **Built-in Render Pipeline** |
| --- | --- | --- | --- | --- |
| **ShaderLab: Pass block** | Yes | Yes | Yes | Yes |

## Syntax

| **Signature** | **Function** |
| --- | --- |
| `Pass` `{`     `<optional: name>`  `<optional: tags>`     `<optional: commands>`    `<optional: shader code>` `}` | Defines a Pass. |

## Additional resources

* [Add a shader pass in a custom shader](writing-shader-create-shader-pass.md)