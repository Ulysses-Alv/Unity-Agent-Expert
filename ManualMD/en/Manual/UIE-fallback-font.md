# Fallback font

Each font asset has a limited number of characters. When you use a character that isn’t in the current font asset, TextCore searches the **Fallback List** until it finds a font asset with that character. The text element then uses that font asset to render the character.

You can set a list of fallback fonts to distribute fonts over multiple textures, or use different fonts for specific characters. It requires extra computing resources to search the list for missing characters and additional draw calls to use additional fonts.

You can use the same character for multiple font assets in the fallback list. Match the style of the characters in the fallback list to style of the main font asset as possible as you can.

## Fallback font usage

In general, use fallback font assets to:

* Work with languages that have large alphabets, such as Chinese, Korean, and Japanese. Use fallback fonts to distribute an alphabet across several assets.
* Include special characters from other alphabets in your text.

**Tips**: Dynamic OS font assets are great candidates for fallback font assets.

## Fallback font chain

You can create local and global fallback font asset lists. Set local font asset lists in the [Font Assets property](UIE-font-asset-properties.md) and set global font asset lists in [UITK Text Settings](UIE-text-setting-asset.md). In addition to the fallback fonts, TextCore searches other assets, such as the default **sprite** asset, for missing glyphs. Together, these assets form the fallback chain.

The following is the asset order in the fallback chain:

1. Local fallback font assets list
2. Global fallback font assets list
3. Default sprite asset
4. Dynamic OS fallback
5. Missing glyphs character

## Additional resources

* [Font assets](UIE-font-asset.md)
* [Include sprites in text](UIE-sprite.md)