# Sort isometric Sprites with the Sprite Atlas

In **Chunk Mode**, the **Tilemap** Renderer isn’t able to sort Tiles from multiple textures individually and doesn’t render the tile **sprites** consistently.

![The Scene view with an example of isometric tile rendering issues.](../../../../../uploads/Main/2D_IsoTilemap_10.png)

The Scene view with an example of isometric tile rendering issues.

Pack all the individual Sprites that make up the Tilemap into a single [Sprite Atlas](https://docs.unity3d.com/Manual/SpriteAtlas.html) to solve this issue. To do this:

1. Create a **Sprite Atlas** from the Assets menu (go to: **Atlas > Create > Sprite Atlas**).
2. Add the Sprites to the Sprite Atlas by dragging them to the **Objects for Packing** list in the Atlas’ **Inspector** window.

   ![Objects for packing list](../../../../../uploads/Main/2D_IsoTilemap_11.png)

   **Objects for packing** list
3. Click **Pack Preview**. Unity packs the Sprites into the Sprite Atlas during Play mode, and correctly sorts and renders them. This is only visible in the Editor during Play mode.

   ![Isometric tilemap](../../../../../uploads/Main/2D_IsoTilemap_12.png)

   Isometric tilemap