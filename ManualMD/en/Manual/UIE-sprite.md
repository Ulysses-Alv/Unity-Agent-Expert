# Include sprites in text

To use sprites in your rich text tags, such as emojis, you need a **sprite** asset. You create sprite assets from atlas [textures](Textures.md) that contain a set of sprites.

![An example atlas texture](../uploads/Main/font/sprite-asset.png)

An example atlas texture

You can use as many texture atlases and assets as you want. However, using more than one atlas for a single text element increases draw calls and uses more system resources. As a general rule, if you’re importing multiple sprites, pack them into one atlas to keep draw calls low. Make sure the atlas resolution is high enough for your target platform.

Use the [`<sprite>` rich text tag](UIE-supported-tags.md#sprite) to include sprites in your text.

Refer to [Sprite asset properties reference](UIE-sprite-asset-properties.md) for more information.

## Create a sprite asset

You create sprite assets from atlas textures. Although sprite assets and their source textures are separate entities, you must keep the source textures in the project after you’ve created the sprite assets.

To create a sprite asset:

1. [Import the texture](ImportingTextures.md).
2. Select the texture atlas and change the following texture options in the **Inspector** window:

   * Set the **Texture Type** to **Sprite (2D and UI)**.
   * Set the **Sprite Mode** to **Multiple**.
3. Select **Sprite Editor** and divide the texture into individual sprites.
4. Right-click the sprite and select **Create** > **Text Core** > **Sprite Asset**. This creates a new sprite asset.
5. From the Inspector window, you can further customize the appearance and names of each glyph. Refer to [Sprite Asset properties](UIE-sprite-asset-properties.md) for more information.
6. Place the sprite asset to the path set in the [UITK Text Settings](UIE-text-setting-asset.md).

Once you have created the sprite asset, you can revert the atlas texture’s **Texture Type** to its original setting.

## Use a sprite asset

To use sprites in the rich text tag, reference the sprite asset name and the sprite name as `<sprite="assetName" name="spriteName">` or by index as `<sprite="assetName" index=1>`.

You can add the `tint=1` attribute to the tag to tint the sprite with the text object’s vertex color. You can also choose a different color by adding a color attribute to the tag, for example: `<sprite="assetName" index=1 color=#55FF55FF>`.

For runtime UI, if you have set a sprite asset as the default in the [UITK Text Settings](UIE-text-setting-asset.md), you can omit the asset name as `<sprite index=1>` (or shorthand `<sprite=1>`), or `<sprite name="spriteName">`.

## Assign and use Unicode for a sprite

You can assign a Unicode to a sprite and use the Unicode directly in your text object instead of the `<sprite>` tag.

For example, the Unicode for a smiling face emoji is `U+1F60A`. To assign it to a sprite in your sprite asset:

1. In the sprite asset’s Inspector window, find the glyph in the **Sprite Character Table**. You can browse or search by index or name.
2. Click the glyph to enable the edit mode.
3. In the Unicode box, enter `+1F60A`.
4. Click the **Unicode** label to save your changes. The Unicode changes to `0xF1F60A`.

To use the smiling face emoji in your text object, enter `\U00F1F60A`.

## Additional resources

* [Sprite asset properties](UIE-sprite-asset-properties.md)
* [Rich text tags](UIE-rich-text-tags.md)