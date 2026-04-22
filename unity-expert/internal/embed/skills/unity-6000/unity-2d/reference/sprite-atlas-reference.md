# Sprite Atlas Inspector window reference

Explore the properties and settings you can use to customize a **sprite** atlas. For more information, refer to [Create a sprite atlas](create-sprite-atlas.md).

| **Property** | **Description** |
| --- | --- |
| **Type** | Sets whether the sprite atlas is a lower-resolution version of another sprite atlas. The options are:  * **Master**: Indicates the sprite atlas can be used as the parent of a **Variant** sprite atlas. This is the default. * **Variant**: Indicates that the sprite atlas is a lower-resolution version of another sprite atlas.  For more information, refer to [Create lower resolution versions of sprite atlases](master-variant/master-variant-sprite-atlases.md). |
| **Master Atlas** | Sets the parent **sprite atlas**. This property is available only if you set **Type** to **Variant**. The parent sprite atlas must have a **Type** of **Master**. For more information, refer to [Create lower resolution versions of sprite atlases](master-variant/master-variant-sprite-atlases.md). |
| **Include in Build** | Enables Unity loading the sprite atlas and attaching it to your sprites when your project starts. Disable this property to load the sprite atlas yourself at runtime instead. For more information, refer to [Load sprite atlases manually at runtime](distribution/load-sprite-atlas-spriteatlasmanageratlasrequested.md). |

## Variant

| **Property** | **Description** |
| --- | --- |
| **Scale** | Sets the scale of the variant sprite atlas as a multiple of its parent sprite atlas. For example, set **Scale** to 0.5 to create a sprite atlas texture half the resolution of the parent sprite atlas. The maximum value is 1. This property is available only if you set **Type** to **Variant**. For more information, refer to [Create lower resolution versions of sprite atlases](master-variant/master-variant-sprite-atlases.md). |

## Packing

These properties are available only if you set **Type** to **Master**.

| **Property** | **Description** |
| --- | --- |
| **Allow Rotation** | Rotates sprites in the sprite atlas texture to find a better fit. Disable this property if you pack [Canvas UI](UICanvas) elements, to avoid UI elements rotating. |
| **Tight Packing** | Packs each sprite based on its **mesh** shape from the Custom Outline tab of the Sprite Editor window, instead of the default rectangle. |
| **Alpha Dilation** | Expands the colors of the edge of the sprite into adjacent transparent **pixels**, to prevent visible edges. |
| **Padding** | Sets how many pixels Unity uses between individual sprites, to prevent pixels from one sprite appearing in an adjacent sprite. The default value is 4. |

## Texture

| **Property** | **Description** |
| --- | --- |
| **Read/Write** | Allows you to access the data of the sprite atlas in C# **scripts** using methods such as [Texture2D.SetPixels](../../../ScriptReference/Texture2D.SetPixels.md). If you enable this property, Unity creates a duplicate copy of the texture data, which can reduce performance. This property only has an effect if the sprite atlas is uncompressed or uses DXT **compression**. |
| **Generate Mip Maps** | Creates mipmap levels for the sprite atlas texture. For more information, refer to [Mipmaps](../../texture-mipmaps-introduction.md). |
| **sRGB** | Stores the sprite atlas colors in gamma space. For more information, refer to [Color spaces](../../color-spaces-landing.md). |
| **Filter Mode** | Specifies how Unity filters the sprite atlas when it stretches. The options are:  * **Point**: Uses the nearest pixel. This makes the sprite atlas appear pixelated. * **Bilinear**: Uses a weighted average of the four nearest texels. This makes the sprite atlas appear blurry when you magnify it. * **Trilinear**: Uses a weighted average of the two nearest mipmap levels, which Unity filters using bilinear filtering. This creates a soft transition between mipmap levels, but the sprite atlas is slightly more blurry. |
| ****Aniso Level**** | Sets the anisotropic filtering level of the sprite atlas. This increases texture quality when you view the sprite atlas at a steep angle. This property is available only if you enable **Generate Mip Maps** and set **Filter Mode** to **Bilinear** or **Trilinear**. |

## Show Platform Settings

The properties in this section set the size and format of the sprite atlas texture.

| **Property** | **Description** |
| --- | --- |
| **Show Platform Settings For** | Selects which texture the properties in the **Default** tab and the platform-specific tabs display. Select either the main texture of the sprite atlas, or a secondary texture. For more information about secondary textures, refer to [Add normal map and mask map textures to a sprite](../../urp/SecondaryTextures.md). |

### Default tab

The **Default** tab lets you set the final size and format of the sprite atlas texture on all platforms.

| **Property** | **Description** |
| --- | --- |
| **Max Texture Size** | Sets the maximum dimensions of the sprite atlas in pixels. If any of the sprites or textures are larger than the **Max Texture Size**, Unity ignores the **Max Texture Size** setting and uses the minimum size that contains the sprites and textures at their original dimensions.  **Important**: Don’t use **Max Texture Size** values smaller than 0.25 for [variant sprite atlases](master-variant/master-variant-sprite-atlases.md), otherwise Unity compresses the sprites and textures twice. |
| **Format** | Sets the number of channels and the data type of the sprite atlas, except for platforms you override in the platform-specific override tabs. The default is **Automatic**, where Unity selects a value based on the number of channels in the sprites, and the **Compression** setting. For more information, refer to [Default formats](../../class-TextureImporter-type-specific.md#default-formats). |
| **Compression** | Sets the compression of the sprite atlas. Unity uses this setting to select an appropriate texture format. This setting is only available if you set **Format** to **Automatic**. For more information, refer to [Default formats](../../class-TextureImporter-type-specific.md#default-formats). The options are:  * **None**: Don’t compress the sprite atlas. * **Low Quality**: Compresses the sprite atlas using a low-quality texture format. The compressed texture might use less memory than **Normal Quality**. * **Normal Quality**: Compresses the sprite atlas using a standard texture format. * **High Quality**: Compresses the sprite atlas using a high-quality texture format. The compressed texture might use more memory than **Normal Quality**. |
| **Use Crunch Compression** | Compresses the sprite atlas using the Crunch compression library, which helps Unity use the lowest possible amount of space. This setting is only available if you set **Format** to **Automatic**. Unity decompresses the sprite atlas to DXT or ETC format on the CPU, then uploads it to the GPU at runtime. If you enable this setting, textures might take a long time to compress, but decompression at runtime is fast. |
| **Compressor Quality** | Sets the image quality of the compressed texture, if you enable **Use Crunch Compression**. A higher value might use more memory and increase compression time. |

### Platform-specific overrides tabs

The platform-specific override tabs let you override the settings in the **Default** tab for specific platforms. For more information about build platforms, refer to [Build Profiles window reference](../../build-profiles-reference.md).

| **Property** | **Description** |
| --- | --- |
| **Max Texture Size** | Sets the maximum dimensions of the imported texture in pixels. If any of the sprites or textures are larger than the **Max Texture Size**, Unity ignores the **Max Texture Size** setting and uses the minimum size that contains the sprites and textures at their original dimensions.  **Important**: Don’t use **Max Texture Size** values smaller than 0.25 for [variant sprite atlases](master-variant/master-variant-sprite-atlases.md), otherwise Unity compresses the sprites and textures twice. |
| **Format** | Sets the sprite atlas format. The available **texture formats** depend on the platform. Unity selects the default format based on the platform and the settings in the **Default** tab. For more information, refer to [Default formats](../../class-TextureImporter-type-specific.md#default-formats). |
| **Compressor quality** | Sets the quality of the sprite atlas that compression produces. This setting isn’t available for all texture formats. |

## Objects for Packing

| **Property** | **Description** |
| --- | --- |
| **Packables** | The list of sprites, textures, and folders to pack into the sprite atlas.  To add an object for packing, either drag it from the **Project** window onto the **Objects for Packing** label, or select the **Add** (**+**) button.  To remove an object, select it in the list, then select the **Remove** (**-**) button.  For more information, refer to [Create a sprite atlas](create-sprite-atlas.md). |
| **Pack Preview** | Combines the **Packables** objects into a sprite atlas, then displays a preview in the **Preview** window at the bottom of the ****Inspector**** window. |

## Preview window

If your sprites have **normal map** or mask map secondary textures, open the dropdown at the top-right of the **Preview** window to display the different sprite atlas textures.

## Additional resources

* [Sprite atlas best practices in Unity 6](https://www.youtube.com/watch?v=hXlpnwD-TgY) on the Unity YouTube channel
* [2D game art, animation, and lighting for artists](https://unity.com/resources/2d-game-art-animation-lighting-unity-6-3-lts?isGated=false)