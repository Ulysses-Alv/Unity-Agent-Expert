# Writing HLSL shader programs

Write High-Level **Shader** Language (HLSL) vertex and fragment shader programs inside a `Pass` block in a custom **ShaderLab** shader.

| **Page** | **Description** |
| --- | --- |
| [Introduction to HLSL shader programs](writing-shader-programs-introduction.md) | Learn about writing shader programs in HLSL in Unity, and vertex and fragment shaders. |
| [Add an HLSL shader program](writing-shader-add-shader-program.md) | Use the `HLSLPROGRAM` directive to add a shader program to a shader pass. |
| [Reuse HLSL code](writing-shader-include-shader-program.md) | Include or duplicate HLSL in different places, or share `#pragma` directives across different shaders. |
| [Input data into HLSL](SL-VertexProgramInputs.md) | Create a `struct` that declares shader variables and connects them to the vertex data of the **mesh**. |
| [Sample a texture in HLSL](SL-SamplerStates.md) | Declare and sample a texture in HLSL using a sampler. |
| [Writing HLSL for different graphics APIs](SL-PlatformDifferences.md) | Write HLSL to take into account how graphics rendering differs between different graphics APIs. |
| [Use 16-bit precision in shaders](SL-Use16BitPrecisionInShaders.md) | Use 16-bit precision in GPU calculations, so shaders use less memory, bandwidth, and power, and calculations are faster. |
| [Pass information to the shader compiler](writing-shader-programs-pragma-directives.md) | Use the `#define`, `#pragma`, and `#define_for_platform_compiler` directives to pass information to the shader compiler. |

## Additional resources

* [Shader languages reference](shaders-reference.md)
* [HLSL pragma directives reference](SL-PragmaDirectives.md)
* [HLSL shader examples in the Built-In Render Pipeline](built-in-shader-examples.md)
* [Writing custom shaders in the Universal Render Pipeline (URP)](urp/writing-custom-shaders-urp.md)