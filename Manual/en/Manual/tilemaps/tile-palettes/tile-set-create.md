# Create a Tile Set

A Tile Set is an asset which helps you create a [Tile Palette](tile-palette-landing.md) and other supporting assets to help you paint Tiles onto a **Tilemap**.

Unlike a Tile Palette, the Tile Set keeps the references of the textures that create Tiles, and automatically updates both itself and the Tiles if the texture changes.

Use a Tile Set asset if your Tiles are directly correlated with the textures you use, and you don’t need a lot of customization for your Tiles and Tile Palette.

## Create a TileSet Asset

To create a [Tile Set asset](tile-set-properties.md), from the main menu, select **Assets** > **Create** > **2D** > **Tile Palette** > **New Tile Set**.

## Add sprite textures

To populate the Tile Set with your **sprite** textures, follow these steps:

1. Prepare your sprite textures using the Sprite Editor. For more information, refer to [Create sprites from a texture](../../sprite/sprite-editor/use-editor.md).
2. Select the Tile Set asset in the **Project** window.
3. In the ****Inspector**** window, open the **Texture Sources** dropdown.
4. To add a texture, select the **Add** (**+**) button.
5. Open the **Element** dropdown, then select the **Texture** picker (**⊙**) to choose a texture.

To customize generating tiles from each texture, select the **Tile Template** picker (**⊙**) to pair the Texture with a [Tile Template Asset](tile-template-asset.md).

## Adjust the grid

To adjust the grid of the Tile Palette, use the settings under the **Palette Grid** dropdown. You can set properties such as the cell layout or the cell size of the Tile Palette.

## Atlas the textures

To atlas the Textures used in this Tile Set, follow these steps:

1. Check the **Create Atlas** toggle.
2. Change the Sprite Atlas settings as needed. For more information, refer to [Sprite Atlas](../../sprite/atlas/atlas-landing.md).
3. The sprite textures added in the Tile Set will be automatically added as “Objects for Packing” in the Sprite Atlas on **Apply**.

Using a Sprite Atlas to atlas the Textures can help improve the performance of Tilemaps painted with this Tile Set by reducing draw calls, and reducing visual artifacts from the edges of Sprites from the Textures used.

## Apply changes

To apply changes you make to the Tile Set, select **Apply**.

## Use the Tile Set as a Tile Palette

When you’ve finished making changes to the Tile Set asset, follow these steps to use the Tile Set as a Tile Palette:

1. Select **Open Tile Palette Window**.
2. Set the **Active Palette** dropdown to the Tile Set.

You can now start painting Tiles with your Tile Set.

## Additional resources

* [Work with Tile Palettes](tile-palette-landing.md)
* [Tile Palette window reference](tile-palette-editor-reference.md)