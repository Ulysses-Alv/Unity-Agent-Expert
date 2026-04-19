# Assign a channel-packed texture to a URP material

In the [Lit](lit-shader.md) and [Complex Lit](shader-complex-lit.md) **shaders** in the Universal **Render Pipeline** (URP), you can use a single RGBA texture for the metallic, smoothness, and occlusion properties. This technique is called channel packing or texture packing.

Follow these steps:

1. Create a texture with the following channels:

   | **Channel** | **Property** |
   | --- | --- |
   | Red | Metallic |
   | Green | Occlusion |
   | Blue | Not used |
   | Alpha | Smoothness |
2. [Import the texture](../ImportingTextures.md).
3. In the **Inspector** window for the texture, disable **sRGB (Color Texture)**.
4. Assign the texture to the **Metallic** and **Occlusion** properties of the material. Unity automatically uses the correct channels from the texture.

You can use [Shader Graph](../shader-graph.md) instead if you pack channels in textures differently.

## Additional resources

* [Materials](../Materials.md)
* [Custom textures](../Textures-landing.md)