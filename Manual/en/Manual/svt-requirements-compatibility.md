# Streaming Virtual Texturing requirements and compatibility

## Render pipeline support

In the [High Definition Render Pipeline (HDRP)](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest/) you can use Streaming Virtual Texturing (SVT) in **Shaders** you create with [Shader Graph](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest).

## System requirements

### GPU requirements

* GPU compute
* Texture2DArrays
* AsyncReadBack

### Supported platforms

* Windows
* Mac
* Linux

**Note:** Streaming Virtual Texturing is not supported on platforms that use ARM architecture.

Some additional console platforms support this functionality. Consult the documentation for your target platform to determine if it supports Streaming Virtual Texturing.

### Supported graphics APIs

* DirectX 11
* DirectX 12
* Metal
* Vulkan

## Texture compatibility

* SVT is built on top of [Texture2D](../ScriptReference/Texture2D.md). SVT Textures go through the same importer, which limits the maximum texture size. SVT doesn’t support textures larger than 16K x 16K. It also does not support UDIM Textures, or merging tiled image files into one large texture.
* SVT doesn’t support all Texture graphics formats. The supported formats are:

  + `GraphicsFormat::kFormatRGBA_DXT1_SRGB`
  + `GraphicsFormat::kFormatRGBA_DXT1_UNorm`
  + `GraphicsFormat::kFormatRGBA_DXT3_SRGB`
  + `GraphicsFormat::kFormatRGBA_DXT3_UNorm`
  + `GraphicsFormat::kFormatR_BC4_UNorm`
  + `GraphicsFormat::kFormatRG_BC5_UNorm`
  + `GraphicsFormat::kFormatRGB_BC6H_SFloat`
  + `GraphicsFormat::kFormatRGB_BC6H_UFloat`
  + `GraphicsFormat::kFormatRGBA_BC7_SRGB`
  + `GraphicsFormat::kFormatRGBA_BC7_UNorm`
  + `GraphicsFormat::kFormatR8_SRGB`
  + `GraphicsFormat::kFormatR8_UNorm`
  + `GraphicsFormat::kFormatR8G8_SRGB`
  + `GraphicsFormat::kFormatR8G8_UNorm`
  + `GraphicsFormat::kFormatR8G8B8_SRGB`
  + `GraphicsFormat::kFormatR8G8B8_UNorm`
  + `GraphicsFormat::kFormatR8G8B8A8_SRGB`
  + `GraphicsFormat::kFormatR8G8B8A8_UNorm`
  + `GraphicsFormat::kFormatR16_SFloat`
  + `GraphicsFormat::kFormatR16_UNorm`
  + `GraphicsFormat::kFormatR16G16_SFloat`
  + `GraphicsFormat::kFormatR16G16_UNorm`
  + `GraphicsFormat::kFormatR16G16B16A16_SFloat`
  + `GraphicsFormat::kFormatR16G16B16A16_UNorm`
  + `GraphicsFormat::kFormatR32_SFloat`
  + `GraphicsFormat::kFormatR32G32_SFloat`
  + `GraphicsFormat::kFormatR32G32B32A32_SFloat`
  + `GraphicsFormat::kFormatA2B10G10R10_UNormPack32`
* SVT does not support the following:

  + Mirror wrap mode.
  + Per-axis clamping (for example, repeating horizontally while clamping vertically).
  + Textures with resolutions that aren’t a power of two.
* Textures in SVT don’t have mipmaps smaller than the tile size (128 x 128 pixels). SVT clamps the sampling to this mipmap level, so you can see aliasing for certain content at a distance.
* SVT only supports trilinear filtering in the Player, and not in the Editor.
* SVT only supports streaming from disk in the Player, and not in the Editor.
* The maximum anisotropic filtering level is 8. This is because each tile has an 8-pixel border.
* You cannot set Textures that are part of a stack dynamically in the Player. In the Player, you cannot call `Material.SetTexture` or `Renderer.SetPropertyBlock` to modify the Texture Stack.

## Unity feature compatibility

* SVT does not support the following:

  + AssetBundles, including Addressables and Live Link.
  + Scriptable Build Pipeline (SBP) and Subscenes.
  + Crunch **compression**.
* The **aspect ratios** of stack layers must match.