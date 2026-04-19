# Grid Selection properties reference

Use the [Select tool](./select-tiles-with-select-tool.md) to select one or more grid cells. The Grid Selection ****Inspector**** window displays the contents and properties of the tiles at the selected location.

| **Property** | **Function** |
| --- | --- |
| **Tile** | Displays the tile at the selected cell location. If you select multiple cells, and they display the same tile, then this property displays that tiles’ name. If you select multiple cells with different tiles, this property is blank. |
| ****Sprite**** | Displays the sprite assigned to the tile in the **Tile** property above. If you select multiple cells with the same tile, then this displays the same sprite. If you select multiple cells with different tiles, this property is blank. This is grayed out by default and can’t be edited. |
| **Color** | The vertex color of the sprite. This is grayed out if [Lock Color](#LockColor) is enabled, so you can’t edit it. |
| ****Collider** Type** | The collider type of the tiles at the selected location. This is grayed out and can’t be edited. |
| **Position** | Enter the offset (in cells) for each axis to shift the tile sprites along the respective axis. This changes where Unity displays the tile sprites on the **tilemap**, but doesn’t change the tiles’ actual cell positions on the tilemap. |
| **Rotation** | This rotates one or more tile sprites at the selected location. Enter the rotation (in degrees) for each axis to rotate the tile sprites around the respective axis. This changes how Unity displays the tile sprites on the tilemap, but doesn’t change the tiles’ actual cell positions on the tilemap. |
| **Scale** | Scales the size of one or more tile sprites at the selected location. Enter the factor for each axis to scale the tile sprite by along the respective axis. This changes how Unity displays the tile sprites on the tilemap, but doesn’t change the tiles’ actual cell positions on the tilemap. |
| **Lock Color** | Select this to prevent changes to the [Color](#Color) of the tile, and clear this to enable the **Color** property. When this property is grayed out, its state remains fixed. Refer to the [Tilemaps.TileFlags](../../../../../ScriptReference/Tilemaps.TileFlags.md) set in the [Tile Asset](../../../tiles-for-tilemaps/tile-asset-reference.md) to change this property. |
| **Lock Transform** | Select this to prevent changes to the Transforms of the tile, and clear this to enable the **Transform** properties. When this property is grayed out, its state remains fixed. Refer to the [Tilemaps.TileFlags](../../../../../ScriptReference/Tilemaps.TileFlags.md) set in the [Tile Asset](../../../tiles-for-tilemaps/tile-asset-reference.md) to change this property. |
| **Delete Selection** | Select this to delete the contents of selected cells in the tilemap. |

## Additional resources

* [Modify Tilemap reference](./modify-tilemap-reference.md)
* [Tile Palette editor reference](../../tile-palette-editor-reference.md)