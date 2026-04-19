# Error and loading shaders

Sometimes, Unity can’t render objects with regular **shaders**. When this happens, Unity renders the objects with special shaders:

* [The default error shader](#default-error-shader)
* [The loading shader](#loading-shader)
* [The Streaming Virtual Texturing error material](#svt-error-material)

The special shader that Unity uses depends on the reason that Unity can’t use the original shader.

## The default error shader

Unity renders an object with the default error shader when there’s a problem with that object’s material or shader; for example, if no material is assigned, if the shader doesn’t compile, or if the shader isn’t supported.

Unity uses the default error shader in the Unity Editor, and in builds.

The default error shader is magenta (bright pink).

![The magenta error shader.](../uploads/Main/shader-error.png)

The magenta error shader.

When you use the [BatchRendererGroup](../ScriptReference/Rendering.BatchRendererGroup.md) API, Unity doesn’t display the default error shader. Use [BatchRendererGroup.SetErrorMaterial](../ScriptReference/Rendering.BatchRendererGroup.SetErrorMaterial.md) to set a material to use instead.

If your project uses the Universal Rendering Pipeline (URP), Unity might render an object using the default error shader if the object uses a shader from the Built-In **Render Pipeline**. Refer to [Converting your shaders](urp/upgrading-your-shaders.md) for more information.

## The loading shader

Unity renders an object with the loading shader to indicate that Unity is compiling the [shader variant](shader-variants.md) needed to display that object.

Unity shows the loading shader in the Unity Editor when [asynchronous shader compilation](AsynchronousShaderCompilation.md) is enabled, or in a **development build** when [Shader Live Link support](../ScriptReference/BuildOptions.ShaderLivelinkSupport.md) is enabled.

The loading shader is cyan (bright blue).

![The cyan loading shader.](../uploads/Main/shader-loading.png)

The cyan loading shader.

When you use the [BatchRendererGroup](../ScriptReference/Rendering.BatchRendererGroup.md) API, Unity doesn’t display the loading shader. Use [BatchRendererGroup.SetLoadingMaterial](../ScriptReference/Rendering.BatchRendererGroup.SetLoadingMaterial.md) to set a material to use instead.

## The Virtual Texturing error material

If your project uses [Streaming Virtual Texturing](svt-streaming-virtual-texturing.md)  (SVT), Unity uses a special material to indicate problems in your SVT setup. For more information, see [Virtual Texturing error material](svt-error-material.md).