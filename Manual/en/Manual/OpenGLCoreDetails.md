# OpenGL Core

OpenGL Core is a back end capable of supporting the latest OpenGL features on Windows, macOS, and Linux. This scales from OpenGL 3.2 to OpenGL 4.5, depending on the OpenGL driver support.

For information about how to enable OpenGL Core as your default Graphics API in the Editor or Standalone Player, refer to [Override default graphics APIs](configure-graphicsAPIs.md#override-default-gfxapi).

## OpenGL requirements

OpenGL Core has the following minimum requirements:

* macOS 10.8 (OpenGL 3.2), macOS 10.9 (OpenGL 3.2 to 4.1)
* Windows with NVIDIA since 2006 (GeForce 8), AMD since 2006 (Radeon HD 2000), Intel since 2012 (HD 4000 / IvyBridge) (OpenGL 3.2 to OpenGL 4.5)
* Linux (OpenGL 3.2 to OpenGL 4.5)

## macOS OpenGL driver limitations

The macOS OpenGL backend for the Editor and Standalone supports OpenGL 3.x and 4.x features such as tessellation and geometry **shaders**.

However, as Apple restricts the OpenGL version on OS X desktop to 4.1 at most, it does not support all DirectX 11 features (such as Unordered Access Views or Compute Shaders). This means that all shaders that are configured to target Shader Level 5.0 (with `#pragma target 50`) will fail to load on OS X.

Therefore a new shader target level is introduced: `#pragma target gl4.1`. This target level requires at least OpenGL 4.1 or DirectX 11.0 Shader Level 5 on desktop, or OpenGL ES 3.1 + Android Extension Pack on mobiles.

## OpenGL Core features

The new OpenGL back end introduces many new features (previously mostly DX11/GLES3 only):

* Compute shaders (as well as ComputeBuffers and “random write” render textures)
* Tessellation and Geometry shaders
* Indirect draw ([Graphics.DrawProcedural](../ScriptReference/Graphics.DrawProcedural.md) and [Graphics.DrawProceduralIndirect](../ScriptReference/Graphics.DrawProceduralIndirect.md))
* Advanced blend modes

## Additional resources

* [Surface Shaders with DX11 / OpenGL Core Tessellation](SL-SurfaceShaderTessellation.md)
* [Graphics API arguments](EditorCommandLineArguments.md#gfx-api-arguments)