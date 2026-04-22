# Import Sprites for an Isometric Tilemap

[Import](../../../sprite/import-images-sprites/import-images-sprites-landing.md) the individual Tiles or Tilesheet images for your Isometric **Tilemap** into your Unity Project by placing the Textures into the Assets folder. Select the imported images to view their [Texture Importer](../../../sprite/import-images-sprites/import-images-sprites-landing.md) settings in the **Inspector** window.

When importing **Sprites** for use in an Isometric Tilemap, use the following recommended settings. For further information about each setting, refer to [Sprite (2D and UI) texture Import Settings window reference](../../../texture-type-sprite.md).

1. **Texture Type** - Set this to **Sprite (2D and UI)**. Other Texture types are not supported for Tilemaps.
2. **Sprite Mode** - Set this to **Single** if the Texture contains only a single Sprite. Set to **Multiple** if it contains multiple Sprite Textures, for example a Sprite sheet with multiple Tile Textures.
3. ****Pixels** Per Unit (PPU)** - This value is the number of pixels that make up one **Unity unit** for the selected Sprite. This determines the size of the Sprite when it is rendered on the Tilemap. This is also affected by the **Cell Size** setting of the Grid that contains the Tilemap, which determines how many Unity units make up a single Cell. Refer to the example below to see how PPU values and the [Grid’s](../../grid-reference.md) **Cell Size** settings interact.
4. ****Mesh** Type** - Set to **Tight** to ensure the Tile Meshes follow the outline of the imported Sprites, and the Tiles are drawn flush together on the Tilemap. Due to the general diamond shape of most Isometric Tiles, setting this to **Full Rect** may result in drawing of wasted transparent spaces at the corners of an Isometric Tile, and is not recommended.
5. **Generate Physics Shape** - If the Tiles do not need to interact with [Physics2D](../../../2d-physics/2d-physics.md), then clear this option. Leave this option enabled to generate a Physics Shape based on the shape of the Tile Sprite, for use with the [Tilemap Collider](../../work-with-tilemaps/tilemap-collider-2d.md) component. To make the generated Physics Shape match the cell of the Tilemap instead, select the Tile Asset and set its [Collider Type](../../../tilemaps/tiles-for-tilemaps/tiles-landing.md) property to **Grid**.

In the example below, the imported Sprite is a 256x128 image, and the Isometric Tilemap has a **Cell Size** of (XYZ: 1, 0.5, 1) Unity units. To make the Sprite fit exactly on a single Cell of the Tilemap, set its PPU value to 256. Its entire width then corresponds to one Unity unit, which is equal to the width (X value: 1) of a single Cell.

![Left: Sprite set to 256 PPU. Right: The same Sprite set to 128 PPU.](../../../../uploads/Main/2D_IsoTilemap_1.png)

**Left:** Sprite set to 256 PPU. **Right:** The same Sprite set to 128 PPU.

If the Sprite is set to a PPU value of 128, then it becomes 2 (256px/128) Unity units in width. This causes the Sprite to visually appear to cover 2 Cells in width when painted on the Tilemap. However, the original Cell position of the Tile remains unchanged.

After you import the Sprites, refine the outlines of the Sprites by opening the Sprite Editor for each of them and editing their outlines. For Sprites in an Isometric Tilemap, you should set the **Pivot** of the Sprite so that the ‘ground’ is relative to the Sprite. For more information, refer to [Sprite Editor tab reference for the Sprite Editor window](../../../sprite/sprite-editor/sprite-editor-window-reference.md)

If the Texture is imported with **Sprite Mode** set to **Multiple** and contains multiple Sprites, then edit the outline of each Sprite in the Sprite Editor.

---

* Isometric Tilemaps added in [2018.3](https://docs.unity3d.com/2018.3/Documentation/Manual/30_search.html?q=newin20183) NewIn20183