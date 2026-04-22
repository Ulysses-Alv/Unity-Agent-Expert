# Introduction to camera output

A [Camera](Cameras.md) can generate a depth, depth+normals, or motion vector texture. This is a minimalistic G-buffer texture that can be used for **post-processing** effects or to implement custom lighting models.

These are mostly used by effects; for example, [post-processing effects](PostProcessingOverview.md) often use depth information.

Pixel values in the depth texture range between 0 and 1, with a non-linear distribution. Precision is usually 32 or 16 bits, depending on configuration and platform used. When reading from the Depth Texture, a high precision value in a range between 0 and 1 is returned. If you need to get distance from the Camera, or an otherwise linear 0–1 value, compute that manually using [helper macros](SL-BuiltinFunctions.md).

Depth Textures are supported on most modern hardware and graphics APIs. Special requirements are listed below:

* Direct3D 11+ (Windows), OpenGL 3+ (Mac/Linux), Metal (iOS), and popular consoles support depth textures.

The Camera’s depth Texture mode can be enabled using [Camera.depthTextureMode](../ScriptReference/Camera-depthTextureMode.md) variable from script. It is also possible to build similar textures yourself, using [Shader Replacement](SL-ShaderReplacement.md) feature.

There are three possible depth texture modes:

* **DepthTextureMode.Depth**: a depth texture.
* **DepthTextureMode.DepthNormals**: depth and view space normals packed into one texture.
* **DepthTextureMode.MotionVectors**: per-pixel screen space motion of each screen texel for the current frame. Packed into a RG16 texture.

These are flags, so it is possible to specify any combination of the above textures.

## Under the hood

Depth textures can come directly from the actual **depth buffer**, or be rendered in a separate pass, depending on the **rendering path** used and the hardware. Typically when using the **Deferred Shading** rendering path in the Built-In **Render Pipeline**, the depth textures come “for free” since they are a product of the G-buffer rendering anyway.

When enabled, the MotionVectors texture always comes from a extra render pass. Unity will render moving **GameObjects** into this buffer, and construct their motion from the last frame to the current frame.