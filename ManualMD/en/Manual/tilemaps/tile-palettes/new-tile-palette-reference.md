# New tile palette properties reference

When you create a new tile palette, you must configure its properties before the tile palette asset can be created.

| Property | Function |
| --- | --- |
| **Name** | Provide a name for the created Tile Palette Asset. |
| **Grid** | Select the [Grid](../grid-reference.md) layout the created Tile Palette will be used to paint on. |
| **Rectangle** | Select this if creating a Palette for the default rectangular **Tilemap**. |
| **Hexagon** | Select this when creating a Palette for a [Hexagonal Tilemap](../work-with-tilemaps/hexagonal-tilemaps.md). |
| **Isometric** | Select this when creating a Palette for a [Isometric Tilemap](../work-with-tilemaps/isometric-tilemaps/isometric-tilemap-landing.md). Refer to [Creating a Tile Palette for an Isometric Tilemap](../work-with-tilemaps/isometric-tilemaps/create-tile-palette-isometric-tilemap.md) for more information. |
| **Isometric Z as Y** | Select this when creating a Palette for a [Isometric Z as Y Tilemap](../work-with-tilemaps/isometric-tilemaps/isometric-tilemap-landing.md). Refer to [Creating a Tile Palette for an Isometric Tilemap](../work-with-tilemaps/isometric-tilemaps/create-tile-palette-isometric-tilemap.md) for more information. |
| **Hexagon Type** (only available when the Hexagon Grid type is selected) | Select the type of Hexagonal Tilemap that the Palette will be used to paint on. Refer to the documentation on [Hexagonal Tilemaps](../work-with-tilemaps/hexagonal-tilemaps.md) for more information. |
| **Cell Size** | The size of a cell that the Tiles are painted on. |
| **Automatic** | The Cell Size is automatically set in **Unity units** and based on the size of the **Sprite** used to create the Tile Assets. If there are multiple Tiles, the Cell Size is adjusted to match the first Tile from the bottom left of the Palette, so that it fits exactly on a cell. |
| **Manual** | Select this option to input custom size values. |
| **Sort Mode** | Determines the [transparency sort mode](../../../ScriptReference/TransparencySortMode.md) of Renderers in the Tile Palette. |
| **Default** | The default transparency Sort Mode. This mode is determined by the **Graphics Settings** of the project. |
| **Orthographic** | Select this mode to sort Renderers based on the perpendicular distance from the **camera** to a Renderer in the Tile Palette. |
| **Perspective** | Select this mode to sort Renderers based on the direct distance from the camera to a Renderer in the Tile Palette. |
| **Custom Axis Sort** | Select this mode to sort objects based on their distance along a custom axis. |
| **Sort Axis** | Set the XYZ values for the sorting axis, if the **Sort Mode** is set to Custom Axis Sort. |