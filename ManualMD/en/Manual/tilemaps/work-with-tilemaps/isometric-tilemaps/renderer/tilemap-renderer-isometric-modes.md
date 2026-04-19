# Tilemap Renderer isometric modes

The **Tilemap Renderer** component renders the [Tilemap](../../tilemap-reference.md) in the **Scene**. Unity creates Tilemaps with the Tilemap Renderer attached by default. The Tilemap Renderer can:

* [Render in batches with **Chunk Mode**](#batch)
* [Render individually with **Individual Mode**](#individual)

The Render Mode affects the sorting of Tilemap **Sprites** when rendered.

## Rendering in batches with Chunk Mode

**Chunk Mode** is the default rendering mode of the Tilemap Renderer.

When set to **Chunk Mode**, the Tilemap Renderer handles Sprites on a Tilemap in batches and renders them together. They’re treated as a single sort item when sorted in the [2D Transparent Queue](../../../../2d-renderer-sorting.md). While this reduces draw calls, other renderers can’t render between any part of the Tilemap, preventing other rendered Sprites from interweaving with Tilemap Sprites.

In **Chunk Mode**, the Tilemap Renderer isn’t able to sort Tiles from multiple textures individually and doesn’t render the tile sprites consistently. For information on how to work with this, refer to [Use the Sprite Atlas for sorting Sprites](sort-isometric-sprites-sprite-atlas.md).

## Rendering individually with Individual Mode

In **Individual Mode**, the Tilemap Renderer sorts and renders the Sprites on a Tilemap with consideration of other Renderers in the Scene, such as the [Sprite Renderers](../../../../sprite/renderer/sprite-renderer-reference.md) and [Mesh Renderers](../../../../class-MeshRenderer.md). Use this mode if other Renderers interact with Sprites and objects on the Tilemap.

In this mode, the Tilemap Renderer sorts sprites based on their position on the Tilemap and the sorting properties set in the Tilemap Renderer. For example, this allows a character sprite to go in-between obstacle sprites (refer to the example below).

![In Individual Mode, the character can be rendered behind the tower Sprites, and also appear above the ground Sprites.](../../../../../uploads/Main/2D_IsoTilemap_14.png)

In Individual Mode, the character can be rendered behind the tower Sprites, and also appear above the ground Sprites.

Using the same example in **Chunk Mode**, character sprites might get hidden behind ground sprites:

![The same Scene rendered in Chunk Mode.](../../../../../uploads/Main/2D_IsoTilemap_15.png)

The same Scene rendered in **Chunk Mode**.

Using **Individual Mode** might reduce performance as there is more overhead when rendering each sprite individually on the Tilemap.

To sort and render tile sprites on an **Isometric Z as Y** Tilemap, you must set the **Transparency Sort Axis** to a **Custom Axis**. For more information on how to do this, refer to [Custom Sorting Axis for Tilemaps in Individual Mode](sort-sprites-custom-sorting-axis.md).