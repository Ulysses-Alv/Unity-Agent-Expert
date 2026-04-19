# Introduction to HLSL shader programs

In Unity, you usually write **shader** programs in High-Level Shader Language (HLSL). To add HLSL code to your [shader asset](SL-Shader.md), add the code in a [shader code block](shader-shaderlab-code-blocks.md).

**Note:** Unity also supports writing shader programs in other languages like [OpenGL Shading Language (GLSL)](SL-GLSLShaderPrograms.md), but this isn’t generally needed or recommended.

There are two main types of shader programs you write in HLSL:

* **Vertex shader** programs, also known as vertex shaders. These programs usually transform the position of the **mesh** from object space to clip space, and pass data such as UV coordinates to the fragment shader program.
* Fragment shader programs, also known as fragment shaders or **pixel** shaders. These programs usually calculate and output the color of each pixel on-screen.

To make it easier to write shaders across different HLSL versions, graphics APIs, and platforms, Unity provides shader libraries with preprocessor macros. For more information, refer to [Shader methods in the Universal Render Pipeline (URP)](urp/use-built-in-shader-methods.md) or [Shader methods in the Built-In Render Pipeline](use-built-in-shader-methods-birp.md).

For general information on writing HLSL, refer to the [HLSL documentation](https://docs.microsoft.com/en-us/windows/win32/direct3dhlsl/dx-graphics-hlsl) on the Microsoft website.

## Additional resources

* [Optimize shaders](SL-ShaderPerformance.md)
* [Writing shaders for different graphics APIs](SL-PlatformDifferences.md)
* [Use 16-bit precision in shaders](SL-Use16BitPrecisionInShaders.md)
* [Custom Function Node](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest/index.html?subfolder=/manual/Custom-Function-Node.html) in Shader Graph