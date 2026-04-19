# Shader Import Settings window reference

[Switch to Scripting](../ScriptReference/ShaderImporter.md "Go to ShaderImporter page in the Scripting Reference")

View and edit properties and settings for a [shader](materials-and-shaders.md).

## ShaderLab shader properties

The following properties are available only for [ShaderLab shaders](SL-landing.md).

| **Property** | **Description** |
| --- | --- |
| **Default Maps** | Displays the textures that appear as material properties. |
| **NonModifiable Maps** | Displays the textures that have the `[NonModifiableTextureData]` attribute in the shader. These textures don’t appear as material properties. |

## Shader graph properties

The following properties are available only for [shader graphs](shader-graph.md).

| **Property** | **Description** |
| --- | --- |
| **Open Shader Editor** | Opens the **Shader Graph** window so you can edit the shader graph. |
| **View Generated Shader** | Opens your text editor with the **ShaderLab** shader code that Shader Graph generates. |
| **Copy Shader** | Copies the ShaderLab shader code that Shader Graph generates to the clipboard. |

## Imported object

| **Property** | **Description** |
| --- | --- |
| **Surface Shader** | Opens your text editor with the shader Unity generates, if the shader is a [surface shader](writing-surface-shaders.md). |
| **Fixed Function** | Opens your text editor with the shader Unity generates, if the shader is a fixed function shader. |
| **Preprocess only** | Sets **Compile and show code** to generate only the preprocessed code, instead of the final code. For more information, refer to [Shader compilation](shader-compilation.md). |
| **Strip #line directives** | Removes `#line` statements that display how the preprocessed code maps to the original ShaderLab and HLSL code. This property is available only if you enable **Preprocess only**. |
| **Compiled code** | Select **Compile and show code** to open your text editor with the shader code Unity compiles. Use the dropdown to select which graphics APIs and platforms Unity compiles the code for. For more information, refer to [Compiled code dropdown](#compiled-code-dropdown).  This property is only available for code shaders. For shader graphs, select **View Generated Shader**. |
| **Cast shadows** | Displays whether geometry casts shadows when it uses this shader. |
| **Render queue** | Displays which render queue Unity uses for geometry that uses this shader. For more information, refer to [Set the render pipeline and render order for a shader](writing-shader-tags-introduction.md). |
| ****LOD**** | Displays the shader LOD value of the shader, which Unity uses to prioritize shaders. For more information, refer to [Optimize shaders](SL-ShaderPerformance.md). |
| **Ignore projector** | Displays whether geometry that uses this shader receives effects from projector components. |
| **Disable batching** | Displays whether the shader prevents Unity from applying **dynamic batching** to geometry that uses the shader. For more information, refer to [Batch meshes at runtime](static-batching-enable.md). |
| **Keywords** | Displays the **Overridable** and **Not Overridable** keywords the shader defines. For more information, refer to [shader keywords](shader-keywords.md). |
| **SRP Batcher** | Displays whether the shader is compatible with the [Scriptable Render Pipeline (SRP) Batcher](SRPBatcher-landing.md). |
| **Properties** | Displays the properties in the `Properties` block of the shader code. For more information, refer to [Properties block reference in ShaderLab](SL-Properties.md). |

### Compiled code dropdown

For more information, refer to [Shader compilation](shader-compilation.md).

| **Property** | **Description** |
| --- | --- |
| **Current graphics device** | Compiles the shader code for the graphics device on your current machine. |
| **Current build platform** | Compiles the shader code for the current build platform. For more information, refer to [Build profiles](BuildSettings.md) |
| **All platforms** | Compiles the shader code for all platforms and graphics APIs. |
| **Custom** | Compiles the shader code for the graphics APIs you select. The options are:  * **D3D**: Compiles shader code into DirectX Bytecode (DXBC) for DirectX 11 or DirectX 12. * **GLES3x**: Compiles shader code into OpenGL Shading Language (GLSL). * **Metal**: Compiles shader code into Metal Shading Language (MSL). * **OpenGLCore**: Compiles shader code into OpenGL Shading Language (GLSL). * **Vulkan**: Compiles shader code into SPIR-V. * **WebGPU**: Compiles shader code into WebGPU Shading Language (WGSL). |
| **Skip unused shader\_features** | Strips shader variants that materials don’t use. For more information, refer to [Changing how shaders work via branching and keywords](SL-MultipleProgramVariants.md). |

The bottom of the **Compiled code** dropdown displays the number of shader variants Unity compiles. Select **Show** to open your text editor with a list of the shader variants.

## Additional resources

* [Introduction to shaders](shader-introduction.md)
* [Assets and media](assets-and-media.md)