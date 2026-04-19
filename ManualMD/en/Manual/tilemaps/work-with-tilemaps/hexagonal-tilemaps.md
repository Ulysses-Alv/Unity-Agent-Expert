# Hexagonal Tilemaps

In addition to regular [Tilemaps](./tilemap-reference.md), Unity provides both **Hexagonal Point Top** and **Hexagonal Flat Top** Tilemaps. Hexagonal tiles are often used in strategic tabletop games, because they have consistent distance between their centres and any point on their edge, and neighboring tiles always share edges. This makes them ideal for constructing almost any kind of large play area and allows players to make tactical decisions regarding movement and positioning.

The Hexagonal Tilemap uses an offset coordinate system, where alternative rows or columns are offset by half a cell when aligning the cells to the hexagonal grid. For Hexagonal Point Top Tilemaps, every odd row is offset to the right by half a cell’s width. For Hexagonal Flat Top Tilemaps, every odd column is offset to the top by half a cell’s height.

![Example: Hexagonal Point Top Tilemap. Offset rows are colored in yellow.](../../../uploads/Main/hex-tilemap-pointtop-offset.png)

Example: Hexagonal Point Top Tilemap. Offset rows are colored in yellow.

![Example: Hexagonal Flat Top Tilemap. Offset columns are colored in yellow.](../../../uploads/Main/hex-tilemap-flattop-offset.png)

Example: Hexagonal Flat Top Tilemap. Offset columns are colored in yellow.

## Creating a Hexagonal Tilemap

To create a **Hexagonal Tilemap**, follow the same steps to [create a regular Tilemap](./create-tilemap.md) (menu: **GameObject > 2D Object**) but choose one of the **Hexagonal** options in the **2D Object** menu.

---

* Hexagonal Tilemaps added in [2018.2](https://docs.unity3d.com/2018.2/Documentation/Manual/30_search.html?q=newin20182) NewIn20182