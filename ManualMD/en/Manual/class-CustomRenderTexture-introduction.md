# Introduction to Custom Render Textures

**Custom Render Textures** are a special type of texture that allow you to update a texture with a **shader**. They are an extension to [Render Textures](class-RenderTexture.md). You can use Custom Render Textures to create complex simulations like caustics, ripple simulation for rain effects, and liquid splatters.

The Custom Render Textures feature provides a scripting and Shader framework to help with complicated configuration like varying update frequency, partial or multi-pass updates.

To use this framework you need to assign a Material to the Custom Render Texture asset. Custom Render Textures require a compatible Material. For more information, see [Writing a shader for a Custom Render Texture](class-CustomRenderTexture-write-shader.md).

## Render Pipeline Compatibility

The following table describes the compatibility between the **Custom Render Textures** feature and each **render pipeline**:

| Feature | Universal Render Pipeline (URP) | High Definition Render Pipeline (HDRP) | Custom Scriptable Render Pipeline (SRP) | Built-in Render Pipeline |
| --- | --- | --- | --- | --- |
| Custom Render Textures | Yes (1) | Yes (1) | Yes (1) | Yes (1) |

**Note:**

1. To create Materials that update and initialize **Custom Render Textures** in [Shader Graph](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest), selecting **Custom Render Texture** as the target of the shader graph.