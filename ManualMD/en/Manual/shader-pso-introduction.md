# Introduction to PSO tracing and warming

Understand what PSO tracing and warming is, and the project requirements to use the `GraphicsStateCollection` API.

**Important**: `GraphicsStateCollection` is an experimental API and might be changed or removed in the future.

When Unity uses a **shader** variant for the first time in your built application, the **scene** might stutter. This is because Unity and the graphics driver need to compile the shader variant, and create a pipeline state object (PSO) with the compiled shader code and its related GPU state.

To reduce stuttering, compile shaders and create PSOs before they’re first needed, so that the graphics driver caches them to disk. This is called warming up, prewarming, or precooking graphics states.

## Project requirements

The [GraphicsStateCollection](../ScriptReference/Experimental.Rendering.GraphicsStateCollection.md) API is supported on Unity 6.0 or newer.

## Supported graphics APIs

You can use the `GraphicsStateCollection` API on the following graphics APIs:

* DirectX 12
* Vulkan
* Metal
* DirectX 11
* OpenGL
* OpenGLES
* WebGL

**Note**: On DirectX 11, OpenGL, OpenGLES, or **WebGL**, the `GraphicsStateCollection` API automatically falls back to shader warm up as outlined in [Other methods to warm up shaders](shader-prewarm-other.md).

## Shader compatibility

The following shader types aren’t compatible with PSO tracing:

* Compute shaders
* **Ray tracing** shaders

## Additional resources

* [Graphics API support](GraphicsAPIs.md)
* [GraphicsStateCollection](../ScriptReference/Experimental.Rendering.GraphicsStateCollection.md) API documentation
* [Trace and manage PSO data collections](shader-pso-trace.md)
* [PSO tracing and warming example](shader-pso-example.md)