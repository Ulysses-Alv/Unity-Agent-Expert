# Enable collision detection for tiles

To enable **collision** detection for [tiles](../tiles-for-tilemaps/tile-asset-reference.md), add a **Tilemap** **Collider** 2D component to the tilemap.

The Tilemap Collider 2D component automatically generates collider shapes for each tile in the tilemap, so that **GameObjects** with a `Collider2D` component collide with the tiles.

You can choose which tiles to enable collision for.

## Add a Tilemap Collider 2D component

To add the component to a tilemap, follow these steps:

1. Select an existing tilemap GameObject in your project, or [create a tilemap](create-tilemap.md).
2. In the ****Inspector**** window, select **Add Component**.
3. Search for and add the **Tilemap Collider 2D** component from the list.

The **Scene** view displays the collider shapes as green outlines over the tilemap.

To make another GameObject collide with the tiles, refer to [Collider 2D](../../2d-physics/collider/collider-2d-landing.md).

## Disable collisions for a tile

Follow these steps:

1. In the **Project** window, select a tile asset.
2. Set **Collider Type** to **None**.

## Create a custom collider shape

To change the shape that a tile uses to register a collision, use the Custom Physics Shape window in the **Sprite** Editor.

1. In the **Project** window, select a tile asset.
2. Set the **Collider Type** to **Sprite**.
3. Select **Sprite Editor**.
4. In the **Sprite Editor** window, in the dropdown at the top-left, select **Custom Physics Shape**.

For more information about editing the shape, refer to [Create collision shapes for a sprite](../../sprite/create-collision-geometry.md).

## Reduce the number of collider shapes

To improve the performance of a tilemap with many collider shapes, add a Composite Collider 2D component. Unity merges the collider shapes of neighboring tiles together, which reduces how much Unity needs to calculate during a physics update.

Follow these steps:

1. Select a tilemap with a **Tilemap Collider 2D** component.
2. In the **Inspector** window, select **Add Component**.
3. Search for and add the **Composite Collider 2D** component from the list.

For more information, refer to [Composite Collider 2D component reference](../../2d-physics/collider/composite-collider/composite-collider-2d-reference.md).

To further configure performance and collider behavior, use the [`TilemapCollider2D` API](../../../ScriptReference/Tilemaps.TilemapCollider2D.md).

## Update collider shapes manually via API

By default, when you add or remove tiles, the Tilemap Collider 2D updates shapes during `LateUpdate` and batches multiple changes together for performance. To force immediate updates, follow these steps:

1. Call [`Tilemaps.TilemapCollider2D.ProcessTilemapChanges`](../../../ScriptReference/Tilemaps.TilemapCollider2D.ProcessTilemapChanges.md) to process changes immediately.
2. Check [`Tilemaps.TilemapCollider2D-hasTilemapChanges`](../../../ScriptReference/Tilemaps.TilemapCollider2D-hasTilemapChanges.md) to determine if processing is required.

For more information, refer to the [`TilemapCollider2D` API](../../../ScriptReference/Tilemaps.TilemapCollider2D.md).

## Additional resources

* [Set up tilemap collision (Unity Learn)](https://learn.unity.com/course/2D-adventure-robot-repair/unit/game-environment-and-physics/tutorial/set-up-tilemap-collision?version=6.5)
* [Tile asset reference](../tiles-for-tilemaps/tile-asset-reference.md)
* [Tilemap Collider 2D component reference](tilemap-collider-2d-reference.md)
* [Composite Collider 2D component reference](../../2d-physics/collider/composite-collider/composite-collider-2d-reference.md)