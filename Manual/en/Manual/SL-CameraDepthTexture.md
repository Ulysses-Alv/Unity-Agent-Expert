# Output a depth texture from a camera

Use `DepthTextureMode` to output a depth texture or a depth-normals texture from a **camera**.

## DepthTextureMode.Depth texture

This builds a screen-sized depth texture.

Depth texture is rendered using the same **shader** passes as used for shadow caster rendering (`ShadowCaster` pass type). So by extension, if a shader does not support shadow casting (i.e. there’s no shadow caster pass in the shader or any of the fallbacks), then objects using that shader will not show up in the depth texture.

* Make your shader [fallback](SL-Fallback.md) to some other shader that has a shadow casting pass, or
* If you’re using [surface shaders](SL-SurfaceShaders.md), adding an `addshadow` directive will make them generate a shadow pass too.

Note that only “opaque” objects (that which have their materials and shaders setup to use [render queue](SL-SubShaderTags.md) <= 2500) are rendered into the depth texture.

## DepthTextureMode.DepthNormals texture

This builds a screen-sized 32 bit (8 bit/channel) texture, where view space normals are encoded into R&G channels, and depth is encoded in B&A channels. Normals are encoded using Stereographic projection, and depth is 16 bit value packed into two 8 bit channels.

[`UnityCG.cginc` include file](SL-BuiltinIncludes.md) has a helper function `DecodeDepthNormal` to decode depth and normal from the encoded **pixel** value. Returned depth is in 0..1 range.

For examples on how to use the depth and normals texture, please refer to [Replacing shaders at runtime](SL-ShaderReplacement.md) or **Ambient Occlusion** in [Post-processing and full-screen effects](PostProcessingOverview.md).

## Additional resources

* [Camera Inspector window reference for the Built-In Render Pipeline](class-Camera.md)