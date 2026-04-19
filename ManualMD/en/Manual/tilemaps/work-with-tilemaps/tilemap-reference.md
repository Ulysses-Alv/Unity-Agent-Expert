# Tilemap component reference

The ****Tilemap**** component stores and manages [Tile Assets](../tiles-for-tilemaps/tile-asset-reference.md) for creating 2D levels. It transfers the required information from the tiles placed on it to other related components such as the [Tilemap Renderer](./tilemap-renderer-reference.md) and the [Tilemap Collider 2D](./tilemap-collider-2d-reference.md).

The 2D Tilemap Editor package is automatically installed when you create a project with the [2D template](https://docs.unity3d.com/hub/manual/project-create.html). You can install the 2D Tilemap Editor directly from the Unity registry via the [Package Manager](../../upm-ui-install.md).

To create, edit, and pick the tiles for [painting onto a tilemap](../tile-palettes/tile-palette-editor-reference.md), refer to the [Tile Palette](../tile-palettes/create-tile-palette.md) (menu: **Window** > **2D** > **Tile Palette**) documentation for more information on its features and tools.

## Properties

| Property | Description |
| --- | --- |
| **Animation Frame Rate** | Set the frame rate at which tile animations play. Increasing or decreasing this value changes the frame rate of the tile animations. For example, if set to 2, tile animations play at double the base frame rate. If set to 3, tile animations play at triple the base frame rate. |
| **Color** | Select a color to apply as a tint to the tiles on this tilemap. Set to white (default color) to have Unity render the tiles without tint. |
| **Tile Anchor** | Enter the amount (in cells) along the xyz axes to offset tile anchor positions on the tilemap. |
| **Orientation** | Select the orientation of tiles on the tilemap. Each of the following options performs the same function by orienting the tiles along the selected plane. |
| **XY** | Unity orients tiles along the XY plane. |
| **XZ** | Unity orients tiles along the XZ plane. |
| **YX** | Unity orients tiles along the YX plane. |
| **YZ** | Unity orients tiles along the YZ plane. |
| **ZX** | Unity orients tiles along the ZX plane. |
| **ZY** | Unity orients tiles along the ZY plane. |
| **Custom** | Select this option to enable the custom orientation settings below. |
| **Offset** | Set the position offset of the custom orientation. This option is only available when you set the tilemap’s **Orientation** to **Custom**. |
| **Rotation** | Set the rotation of the custom orientation. This option is only available when you set the tilemap’s **Orientation** to **Custom**. |
| **Scale** | Set the scale of the custom orientation. This option is only available when you set the tilemap’s **Orientation** to **Custom**. |
| **Info** | Expand this to view the assets present in the tilemap. |
| **Tiles** | Displays a list of [Tile Assets](../tiles-for-tilemaps/tile-asset-reference.md) present in the tilemap. |
| ****Sprites**** | Displays a list of sprites present in the tilemap. |

## Additional resources

* [Tilemap Renderer component reference](./tilemap-renderer-reference.md)
* [Tile Palette preferences reference](../tile-palettes/tile-palette-preferences-reference.md)

Tilemap