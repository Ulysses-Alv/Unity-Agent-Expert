# Scriptable tiles

Scriptable tiles are tiles that you can assign behavior **scripts** to and you can paint with the scriptable tiles on a **Tilemap** component.

These C# scripts allow you to customize how the tile interacts with other tiles or other behaviours defined by the [TileBase](../../../../ScriptReference/Tilemaps.TileBase.md) class.

All tiles added to a Tilemap component must inherit from `TileBase`. `TileBase` provides a tilemap with a fixed set of APIs to communicate its rendering properties. For most cases of the APIs, the location of the tile and the instance of the tilemap the tile is placed on is passed in as arguments of the API. You can use this to find any required attributes for setting the tile information.

The most common methods to override are:

* `RefreshTile` determines which Tiles in the vicinity are updated when this Tile is added to the Tilemap.
* `GetTileData` determines what the Tile looks like on the Tilemap.