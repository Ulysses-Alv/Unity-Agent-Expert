# Platform-specific texture overrides panel reference

To set the size and format of the imported texture, use the platform-specific overrides panel.

![The Texture Import Settings window with the panel of tabs at the bottom highlighted.](../uploads/Main/texture-import-settings-bottom.png)

The Texture Import Settings window with the panel of tabs at the bottom highlighted.

## Default tab

The **Default** tab lets you set the final size and format of the imported texture on all platforms.

| **Property** | **Description** |
| --- | --- |
| **Max Size** | Sets the maximum dimensions of the imported texture in **pixels**. |
| **Resize Algorithm** | Sets the algorithm Unity uses to downscale the texture, if the texture file has larger dimensions than **Max Size**. The options are:  * **Mitchell**: Resizes the texture using the Mitchell algorithm. This is the default. * **Bilinear**: Resizes the texture using bilinear interpolation. This setting might preserve more small, sharp details than **Mitchell**. |
| **Format** | Sets the number of channels and the data type of the final texture, except for platforms you override in the platform-specific override tabs. The default is **Automatic**, where Unity selects a value based on the number of channels in the texture, and the **Compression** setting. For more information, refer to the [Default formats](#default-formats) section. |
| **Compression** | Sets the compression of the final texture. Unity uses this setting to select an appropriate texture format. This setting is only available if you set **Format** to **Automatic**. For more information, refer to the [Default formats](#default-formats) section. The options are:  * **None**: Don’t compress the texture. * **Low Quality**: Compresses the texture using a low-quality texture format. The compressed texture might use less memory than **Normal Quality**. * **Normal Quality**: Compresses the texture using a standard texture format. * **High Quality**: Compresses the texture using a high-quality texture format. The compressed texture might use more memory than **Normal Quality**. |
| **Use Crunch Compression** | Compresses the texture using the Crunch compression library, which helps Unity use the lowest possible amount of space. This setting is only available if you set **Format** to **Automatic**. Unity decompresses the texture to DXT or ETC format on the CPU, then uploads them to the GPU at runtime. If you enable this setting, textures might take a long time to compress, but decompression at runtime is fast. |
| **Compressor Quality** | Sets the image quality of the compressed texture, if you enable **Use Crunch Compression**. A higher value might use more memory and increase compression time. |

## Platform-specific overrides tabs

The platform-specific override tabs let you override the settings in the **Default** tab for specific platforms. For more information about build platforms, refer to [Build Profiles window reference](build-profiles-reference.md).

| **Property** | **Description** |
| --- | --- |
| **Max Size** | Sets the maximum dimensions of the imported texture in pixels. |
| **Resize Algorithm** | Sets the algorithm Unity uses to downscale the texture, if the texture file has larger dimensions than **Max Size**. The options are:  * **Mitchell**: Resizes the texture using the Mitchell algorithm. This is the default. * **Bilinear**: Resizes the texture using bilinear interpolation. This setting might preserve more small, sharp details than **Mitchell**. |
| **Format** | Sets the final **texture format**. The available texture formats depend on the platform, and the file type of the texture. The default is selected by Unity based on the platform and the settings in the **Default** tab. For more information, refer to the [Default formats](#default-formats) section. |
| **Compressor quality** | Sets the quality of the final texture that compression produces. This setting isn’t available for all texture formats. The options are:  * **Fast**: Compresses the texture to create a low-quality image that might use less memory than **Normal**. * **Normal**: Compresses the texture using the standard compression. * **Best**: Compresses the texture and retains a high-quality image that might use more memory than **Normal**. * **Slider**: Sets the quality of the compressed texture. This slider is only available if you select a texture that uses Crunch compression. A higher value might use more memory and increase compression time. |
| **Split Alpha Channel** | Allows alpha splitting for this Texture on these platforms: **tvOS**, **iOS**, and **Android**. For more information, refer to [Notes on Android for Texture compression formats](texture-choose-format-by-platform.md#android). |
| **Override ETC2 fallback** | ETC2 texture decompression fallback override on Android devices that don’t support ETC2.  Allows to choose which texture format to decompress the texture to on Android devices that have no ETC2 texture format support. For more information, refer to [Notes on Android for Texture compression formats](texture-choose-format-by-platform.md#android). |

## Default formats

The following table lists the default **Format** Unity selects in the platform-specific override tab, based on the **Format** and **Compression** settings in the **Default** tab.

For more information about each format, refer to [GPU texture formats reference](texture-formats-reference.md).

| **Platform** | **Texture channels** | **Format for no compression** | **Format for Normal Quality compression** | **Format for High Quality compression** | **Format for Low Quality compression** |
| --- | --- | --- | --- | --- | --- |
| **Windows, Linux, macOS** | RGB | RGB 24 bit | RGB Compressed DXT1 | RGB(A) Compressed BC7 | RGB Compressed DXT1 |
| **Windows, Linux, macOS** | RGBA | RGBA 32 bit | RGBA Compressed DXT5 | RGB(A) Compressed BC7 | RGBA Compressed DXT5 |
| **Windows, Linux, macOS** | **HDR** | RGBA Half | RGB Compressed BC6H | RGB Compressed BC6H | RGB Compressed BC6H |
| ****WebGL**** ([configurable](webgl-texture-compression.md)) | RGB | RGB 24 bit | RGB Compressed DXT1 | RGB Compressed DXT1 | RGB Compressed DXT1 |
| **WebGL** | RGBA | RGBA 32 bit | RGBA Compressed DXT5 | RGBA Compressed DXT5 | RGBA Compressed DXT5 |
| **Android** ([configurable](android-requirements-and-compatibility.md#texture-compression)) | RGB | RGB 24 bit | RGBA Compressed ASTC 6x6 block RGB Compressed ETC2 RGB Compressed ETC | RGBA Compressed ASTC 4x4 block RGB Compressed ETC2 RGB Compressed ETC | RGBA Compressed ASTC 8x8 block RGB Compressed ETC2 RGB Compressed ETC |
| **Android** | RGBA | RGBA 32 bit | RGBA Compressed ASTC 6x6 block RGBA Compressed ETC2 | RGBA Compressed ASTC 4x4 block RGBA Compressed ETC2 | RGBA Compressed ASTC 8x8 block RGBA Compressed ETC2 |
| **iOS** ([configurable](iphone.md)) | RGB | RGB 24 bit | RGBA Compressed ASTC 6x6 block | RGBA Compressed ASTC 4x4 block | RGBA Compressed ASTC 8x8 block |
| **iOS** | RGBA | RGBA 32 bit | RGBA Compressed ASTC 6x6 block | RGBA Compressed ASTC 4x4 block | RGBA Compressed ASTC 8x8 block |
| **tvOS** | RGB | RGB 24 bit | RGBA Compressed ASTC 6x6 block | RGBA Compressed ASTC 4x4 block | RGBA Compressed ASTC 8x8 block |
| **tvOS** | RGBA | RGBA 32 bit | RGBA Compressed ASTC 6x6 block | RGBA Compressed ASTC 4x4 block | RGBA Compressed ASTC 8x8 block |
| **Default** | RGBA | RGBA 32 bit | RGBA 16 bit | RGBA 16 bit | RGBA 16 bit |

## Additional resources

* [GPU texture formats reference](texture-formats-reference.md)
* [Texture type and shape settings reference](class-TextureImporter-type-and-shape.md)