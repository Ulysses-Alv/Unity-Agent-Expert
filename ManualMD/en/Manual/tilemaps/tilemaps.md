# Tilemaps

Unity’s **Tilemap** system stores and handles [Tile Assets](tiles-for-tilemaps/tile-asset-reference.md) for creating 2D levels, which makes it easy to create and iterate level design cycles within Unity. The Tilemap system transfers the required information from the Tiles placed on it to other related components such as the [Tilemap Renderer](work-with-tilemaps/tilemap-renderer-reference.md) and the [Tilemap Collider 2D](work-with-tilemaps/tilemap-collider-2d-reference.md).

By default, the Tilemap package isn’t included in the Unity Editor, so you must download the **2D Tilemap** Editor package from the [Package Manager](../Packages.md).

When you create a Tilemap, the Grid component is automatically parented to the Tilemap and acts as a guide when you lay out Tiles onto the Tilemap. To create, change, or pick the Tiles for painting onto a Tilemap, use the **Tile Palette** (menu: **Window** > **2D** > **Tile Palette**) and its tools. For more information, refer to [Creating a Tile Palette](tile-palettes/create-tile-palette.md) and the [Tile Palette editor reference](tile-palettes/tile-palette-editor-reference.md).

## Render Pipeline Compatibility

Tilemaps are supported by all **render pipeline** that support 2D projects.

| **Feature name** | **Universal Render Pipeline (URP)** | **High Definition Render Pipeline (HDRP)** | **Built-in Render Pipeline** |
| --- | --- | --- | --- |
| Tilemaps | Yes | Yes | No |