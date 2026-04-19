# Lightmap data format

Unity stores **lightmaps** with different compressions and encoding schemes, depending on the target platform and the compression setting in the Lighting Window.

## Encoding schemes

Unity projects can use two techniques to encode baked light intensity ranges into low dynamic range textures when this is needed:

* **RGBM encoding**. RGBM encoding stores a color in the **RGB** channels and a multiplier (**M**) in the alpha channel. The range of RGBM lightmaps goes from 0 to 34.49(52.2) in linear space, and from 0 to 5 in gamma space.
* **Double Low Dynamic Range (dLDR) encoding**. dLDR encoding is used on mobile platforms by simply mapping a range of [0, 2] to [0, 1]. Baked light intensities that are above a value of 2 will be clamped. The decoding value is computed by multiplying the value from the lightmap texture by 2 when gamma space is used, or 4.59482(22.2) when linear space is used. Some platforms store lightmaps as dLDR because their hardware compression produces poor-looking artifacts when using RGBM.

When Linear Color Space is used, the lightmap texture is marked as sRGB and the final value used by the **shaders** (after sampling and decoding) will be in Linear Color Space. When Gamma Color Space is used, the final value will be in Gamma Color Space.

**Note**: When encoding is used, the values stored in the lightmaps (GPU texture memory) are always in Gamma Color Space.

The **Decode Lightmap** shader function from the *UnityCG.cginc* shader include file handles the decoding of lightmap values after the value is read from the lightmap texture in a shader.

## HDR lightmap support

You can use HDR lightmaps on Windows, Mac, Linux, iOS, tvOS, and Android. To control the encoding/compression of the lightmaps for these platforms, go to **Edit** > **Project Settings** > **Player** > **Other Settings** > **Lightmap Encoding**.

Choosing **High Quality** will enable **HDR** lightmap support, whereas **Normal Quality** will switch to using **RGBM** encoding.
**Low Quality** will switch to **dLDR** encoding on mobile platforms, on other platforms it’s equivalent to **Normal Quality**.

When lightmap **Compression** is enabled in the Lighting Window, the lightmaps will be compressed using the **BC6H** compression format on desktop and console platforms. For mobile platforms, Unity selects the HDR format according to the table below.

## Advantages of High Quality (BC6H) lightmaps

* HDR lightmaps don’t use any encoding scheme to encode lightmap values, so the supported range is only limited by the 16-bit floating point **texture format** that goes from 0 to 65504.
* BC6H format quality is superior to DXT5 + RGBM format encoding, and it doesn’t produce any of the color banding artifacts that RGBM encoding has.
* Shaders that need to sample HDR lightmaps are a few ALU instructions shorter because there is no need to decode the sampled values.
* BC6H format has the same GPU memory requirements as DXT5.

Here is the list of encoding schemes and their **texture compression** formats per target platform:

| **Target platform** | **Encoding** | **Compression - size (bits per pixel)** |
| --- | --- | --- |
| Standalone(PC, Mac, Linux) | RGBM / HDR | DXT5 / BC6H - 8 bpp |
| **WebGL** 2.0 | RGBM / RGBM / HDR | DXT5 - 8 bpp |
| iOS ASTC (1) | dLDR / RGBM / HDR | ASTC - 3.56 bpp / ASTC - 3.56 bpp / RGB9E5 - 32 bpp |
| tvOS | dLDR / RGBM / HDR | ASTC - 3.56 bpp / ASTC - 3.56 bpp / RGB9E5 - 32 bpp |
| Android ASTC (2) | dLDR / RGBM / HDR | ASTC - 3.56 bpp / ASTC - 3.56 bpp / ASTC HDR - 3.56 bpp |
| Android ETC2 | dLDR / RGBM / HDR | ETC2 RGB - 4 bpp / ETC2 RGBA - 8 bpp / ASTC HDR - 3.56 bpp |
| Android ETC | dLDR / RGBM / HDR | ETC1 RGB - 4 bpp / ETC2 RGBA - 8 bpp / ASTC HDR - 3.56 bpp |

[1] The texture compression format used for lightmaps on **iOS** depends on the **Texture compression format** setting in the [Player Settings](class-PlayerSettingsiOS.md).

[2] The texture compression format used for lightmaps on **Android** depends on [Player settings](class-PlayerSettingsAndroid.md#Rendering) and [build settings](android-build-settings.md). Refer to [Texture compression settings](android-requirements-and-compatibility.md#texture-compression) for more details on how these settings interact.

## Precomputed real-time Global Illumination (GI)

The inputs to the GI system have a different range and encoding to the output. Surface albedo is 8-bit unsigned integer RGB in gamma space and emission is 16-bit floating point RGB in linear space. For advice on providing custom inputs using a meta pass, refer to the documentation on [Lightmapping and shaders](MetaPass.md).

The irradiance output texture is stored using the RGB9E5 shared exponent floating point format if the graphics hardware supports it, or RGBM with a range of 5 as fallback. The range of RGB9E5 lightmaps is [0, 65408]. For details on the RGB9E5 format, refer to [Khronos.org: EXT\_texture\_shared\_exponent](https://www.khronos.org/registry/OpenGL/extensions/EXT/EXT_texture_shared_exponent.txt).

## Additional resources

* [Texture Importer Override](texture-choose-format-by-platform.md)
* [Texture Types](class-TextureImporter.md)
* [Global Illumination](lighting-window.md)

---

* 2019–04–26 Page amended
* Baked lightmaps added in Unity [2017.2](https://docs.unity3d.com/2017.2/Documentation/Manual/30_search.html?q=newin20172) NewIn20172
* HDR lightmap support added in Unity [2017.3](https://docs.unity3d.com/2017.3/Documentation/Manual/30_search.html?q=newin20173) NewIn20173
* Lightmap encoding settings for mobile platform added in Unity [2019.1](https://docs.unity3d.com/2019.1/Documentation/Manual/30_search.html?q=newin20191) NewIn20191
* Update lightmap texture compression formats for Android and iOS