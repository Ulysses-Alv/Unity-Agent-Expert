# Writing a custom shader in ShaderLab and HLSL

Resources and techniques for writing **Shader** objects, subshaders and shader passes in **ShaderLab** and HLSL.

| **Page** | **Description** |
| --- | --- |
| [Create a shader file](class-Shader.md) | Create a customshader asset file either using the Unity Editor or manually. |
| [Add a subshader in a custom shader](writing-shader-create-subshader-object.md) | Use a `SubShader` block to add one or more sections that define different GPU settings and shader programs for different hardware, **render pipelines**, and runtime settings. |
| [Add a shader pass in a custom shader](writing-shader-create-shader-pass.md) | Use a `Pass` block to write instructions for setting the state of the GPU, and the shader programs that run on the GPU. |
| [Include a shader pass with the UsePass command](writing-shader-usepass.md) | Insert a named Pass from another **Shader object**, to avoid reduce code duplication in shader source files. |
| [Writing HLSL shader programs](writing-shader-writing-shader-programs-hlsl.md) | Resources for writing HLSL shader programs inside a `Pass` block in a custom ShaderLab shader. |
| [GLSL in Unity](SL-GLSLShaderPrograms.md) | If a platform supports OpenGL Core and OpenGL ES, you can write OpenGL Shading Language (GLSL) shader programs in Unity. |
| [Setting the render state on the GPU](writing-shader-render-state-commands.md) | Resources for using commands in a subshader or shader pass that change the render state on the GPU. |

## Additional resources

* [Shader languages reference](shaders-reference.md)
* [Writing custom shaders in URP](urp/writing-custom-shaders-urp.md)
* [Writing custom shaders in the Built-In Render Pipeline](writing-shaders-birp.md)
* [HLSL shader examples in the Built-In Render Pipeline](built-in-shader-examples.md)