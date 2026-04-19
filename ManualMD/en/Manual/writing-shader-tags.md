# Configure when Unity uses a shader

Resources for using `#pragma` and `Tag` blocks in **ShaderLab** to configure when Unity compiles and runs a **shader**.

| **Page** | **Description** |
| --- | --- |
| [Set the rendering pipeline and render pass for a shader](writing-shader-tags-introduction.md) | Use a `Tag` statement to set the **render pipeline** and render order for a shader. |
| [Set a shader to require a package](writing-shader-tags-require-package.md) | Use a `PackageRequirements` block to add package requirements to a subshader or a pass. |
| [Set a shader to require a shader model or GPU feature](SL-ShaderCompileTargets.md) | Use a `#pragma target` or `#pragma require` statement to set the minimum shader model or GPU feature for a shader. |
| [Set the graphics API or platform for a shader](SL-ShaderCompilationAPIs.md) | Use a `#pragma only_renderers` or `#pragma exclude_renderers` statement to set the graphics API or platform for a shader. |

## Additional resources

* [Replace shaders at runtime in the Built-In Render Pipeline](SL-ShaderReplacement.md)