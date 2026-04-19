# Other methods to warm up shaders

Warm up **shaders** if your build targets the DirectX 11, OpenGL, OpenGLES, or **WebGL** graphics APIs.

If your project uses the DirectX 11, OpenGL, OpenGLES, or WebGL graphics APIs, Unity provides multiple methods to warm up shaders in your project.

## Use the GraphicsStateCollection API (recommended)

The [GraphicsStateCollection](../ScriptReference/Experimental.Rendering.GraphicsStateCollection.md) API is compatible with DirectX 11, OpenGL, OpenGLES, or WebGL. On these APIs, `GraphicsStateCollection` automatically fallbacks to shader warm up. The `GraphcsStateCollection` API is the recommended method of warming up shaders. To learn more about warming up shaders with this API, refer to [Warm up PSOs](shader-prewarm.md).

**Note**: `GraphicsStateCollection` automatic fallback is available in Unity 6.1 and later. If your project uses an older version of Unity, use an alternative method of shader warmup outlined in the following section.

## Other methods to warm up shaders

On DirectX 11, OpenGL, OpenGLES, or WebGL graphics APIs, you can also use the following methods to warm up shaders:

* To warm up a **Shader object**, use the [Experimental.Rendering.ShaderWarmup](../ScriptReference/Experimental.Rendering.ShaderWarmup.md) API.
* To warm up all variants of all Shader objects currently in memory, use the [Shader.WarmupAllShaders](../ScriptReference/Shader.WarmupAllShaders.md) API.

To warm up a shader variant collection, do any of the following:

* Use the [Experimental.Rendering.ShaderWarmup](../ScriptReference/Experimental.Rendering.ShaderWarmup.md) API.
* Use the [ShaderVariantCollection.WarmUp](../ScriptReference/ShaderVariantCollection.WarmUp.md) API.
* Add the shader variant collection to the [Preloaded shaders section of the Graphics Settings window](class-GraphicsSettings.md#shader-loading). Unity then uses the `ShaderVariantCollection.WarmUp` to warm up the shader variants when your built application starts.

**Note**: These methods aren’t recommended on Vulkan, DirectX 12, and Metal because they can create the wrong PSOs due to missing graphics states.

## Additional resources

* [Introduction to PSO tracing and warming](shader-pso-introduction.md)