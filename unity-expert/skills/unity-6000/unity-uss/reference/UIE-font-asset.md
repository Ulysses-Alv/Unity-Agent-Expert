# Introduction to font assets

Font assets act as containers for fonts. With font assets, you can import fonts into your projects and create variations of a font from one file and not change the original font. Another reason to create font assets is to deal with languages with massive character sets, such as Chinese. You can include only the characters you need, rather than the full set of characters.

You can convert TrueType (`.ttf`), OpenType (`.otf`), or TrueType Collection(`.ttc`) fonts to font assets. A font asset includes a font atlas texture file that contains all of the characters. There are several [atlas populate modes](#atlas-population-modes) and [render modes](#atlas-render-modes) to choose from when you create a font asset.

## Atlas population modes

TextCore supports the **Static**, **Dynamic**, and **Dynamic OS** atlas population modes.

### Static font asset

Static font assets pre-bake characters into the atlas texture during conversion. Use static font assets to contain known text in your project, such as labels and section titles. When you create a static font asset, the font atlas is empty by default. You must manually generate the font atlas and include all the characters required for the project. The project build doesn’t require you to include the font source file, so it’s fast and efficient.

Refer to instructions on how to [create a Static font asset](UIE-get-started-with-text.md#create-a-static-font-asset).

### Dynamic font asset

Dynamic font assets start with an empty atlas texture. Dynamic font assets look for characters in the source font file and adds the characters dynamically as you use them in a UI. Use Dynamic font assets for unknown text in your project, such as any **text input fields**. Dynamic font assets are flexible, but they have additional performance overhead because the project build requires to include all source font files.

See instructions on how to [create a Dynamic font asset](UIE-get-started-with-text.md#create-a-dynamic-font-asset).

### Dynamic OS font asset

Dynamic OS font assets are a dynamic font asset that references the font file on the operating system instead of the source font file. Compared to the Dynamic font asset, dynamic OS font assets have less memory overhead. The project build doesn’t require to include the font source files. Make sure that the fonts are on the operating system of the targeted platforms. Dynamic OS font assets are great candidates for fallback font assets.

To create a dynamic OS font asset, import the font from the operating system to your project, and use it to create a dynamic font asset and change the **Atlas Population Mode** to **Dynamic OS**.

## Atlas render modes

TextCore supports atlas render modes for bitmap and Signed Distance Field (SDF).

### Bitmap fonts

Bitmap rendering generates font assets that perfectly align each **pixel** of the font with the **scene** and GUI pixels. Use the following render modes for pixel art:

* **SMOOTH** or **SMOOTH\_HINTED** to render text to an antialiased bitmap. **SMOOTH\_HINTED** aligns character pixels with texture pixels for a sharper image.
* **RASTER** or **RASTER\_HINTE** to render text to a non-antialiased bitmap. **RASTER\_HINTE** aligns character pixels with texture pixels for a sharper image.
* **COLOR** or **COLOR\_HINTED** to render text to a color-image bitmap. **COLOR\_HINTED** aligns character pixels with texture pixels for a sharper image.

### SDF fonts

SDF rendering generates font assets that look crisp when you transform or magnify them. SDF rendering supports effects such as [outlines and drop shadows](UIE-text-effects.md).

Unlike bitmap font textures, SDF font assets contain contour distance information. In font atlases, this information looks like grayscale gradients that run from the middle of each glyph to a point past its edge. The gradient’s mid-point corresponds to the edge of the glyph.

The images below show bitmap and SDF font assets and the rendered text they produce. The bitmap fonts produce text whose edges are more or less jagged and blurry, depending on how far the text is from the **camera**, and how it’s transformed and distorted. The SDF font produces text with smooth edges regardless of the distance from the camera.

![A bitmap font, atlas texture and rendered result](../uploads/Main/font/FontAsssetBMPRaster.png)

A bitmap font, atlas texture and rendered result

![A smoothed bitmap, atlas texture and rendered result](../uploads/Main/font/FontAsssetBMPSmooth.png)

A smoothed bitmap, atlas texture and rendered result

![An SDF font, atlas texture and rendered result](../uploads/Main/font/FontAsssetSDF.png)

An SDF font, atlas texture and rendered result

Use the following render mode for SDF fonts:

* **SDFAA**: Use this mode to produce font atlases for most use cases except when you render large text such as titles that exceed 90 point size on screen with a large outline. **SDFAA** is the default mode for Dynamic or Dynamic OS font assets for fast generation. **SDFAA** is a faster but less accurate SDF generation mode.
* **SDFAA\_HINTED** Use this mode to align character pixels with texture pixels for a crisper result.
* **SDF32**: Default mode for Static font assets. Use this mode for fonts with complex or small characters. This is a slower and more accurate SDF generation mode with 32x oversampling.
* **SDF16**: Use this mode for Static font assets and mostly large text, such as titles that are 72 points in size or larger. It’s a slower and more accurate SDF generation mode with 16x oversampling.
* **SDF8**: Use this mode for a slower and more accurate SDF generation with 8x oversampling.
* **SDF**: Use this mode for a slower and more accurate SDF generation without oversampling.

## Font asset variant

If you want to make changes only to a font asset’s metrics, instead of creating a new font asset, create a font asset variant based on another font asset. A font asset variant can have styling that distinguishes it from the original font asset, without consuming extra space for textures. The variant stores its own [Face Info, Character, Glyph, and Adjustment Pair data](UIE-font-asset-properties.md), but still refers to the original atlas.

To create a font asset variant, right-click the font asset and select **Create** > **Text**> **Font Asset Variant**.

## Padding

Characters in the font texture need padding between them so they can be rendered separately. This padding is specified in pixels. Padding also creates room for the SDF gradient. The larger the padding, the smoother the transition, which allows for higher-quality rendering and larger effects, like thick outlines. Use a padding of `5` for a `512x512` texture. A good rule of thumb is to have a sampling-to-padding ratio of `1:10`.

## Enable Multi Atlas Textures for large texture

For most fonts, use a `512x512` texture resolution if you include all ASCII characters. To support thousands of characters, use large textures and enable **Multi Atlas Textures** in the font asset’s **Inspector** window. This automatically creates new textures of the same size once the main texture is full.

You can also enable **Multi Atlas Textures** if you design for mobile devices, where an imposed maximum texture size prevents an entire set of glyphs from fitting in a single atlas of sufficient quality.

## Use a font asset

To use a font asset, in UI Builder, select the font from **Font Asset** in the Inspector window.

**Note**: The **Font** field is for backward compatibility. If you want to use a font from the **Font** list for your legacy projects, you must select **None** from **Font Asset**. Otherwise, the font you selected from **Font** won’t take effect.

You can apply a font asset to a text element in USS with the following syntax:

```
-unity-font-definition: <resource> | <url>
```

For example:

```
Label {
-unity-font-definition: url("/Assets/UI Toolkit/Resources/Fonts & Materials/LiberationSans SDF.asset");
}
```

For more information, refer to [USS text properties](UIE-USS-SupportedProperties.md#unity-font).

## Best practices

In summary, consider the following when you create a font asset:

* **User input text**: Use Dynamic or Dynamic OS population mode, and SDFAA mode for fast generation. You might need to add fallbacks for localization.
* **Generic labels**: Use the Static population mode and SDF16 mode with a smaller sampling point size.
* **Titles**: Use the Static population mode and SDF32 mode with a higher sampling point size.
* **Pixel Art**: Use any of the bitmap font render modes.
* **Padding**: Use a sampling-to-padding ratio of `1:10`.
* Enable **Multi Atlas Textures** for large texture.

## Additional resources

* [Font Asset properties reference](UIE-font-asset-properties.md)
* [Font Asset Creator properties reference](UIE-font-creator-properties.md)
* [Text effects](UIE-text-effects.md)