# Forward and Forward+ rendering paths in URP

The Universal **Render Pipeline** (URP) has the following forward **rendering paths**:

* Forward
* Forward+

## Forward rendering path

The **Forward rendering** path is the default rendering path in URP. Unity lights each **GameObject** in turn, and there’s a limit to the number of lights that can affect each GameObject.

## Forward+ rendering path

The Forward+ rendering path is similar to the Forward rendering path, but there’s no limit to the number of lights that can affect each GameObject. There’s still a limit on the number of lights visible per-camera.

Using the Forward+ rendering path reduces the number of lights Unity calculates for each GameObject. Unity divides the screen into tiles, then identifies which lights affect which tiles. When Unity calculates the lighting for a GameObject, it uses only the lights that affect the tile the GameObject is in.

![An example of the Lighting Complexity Debug Draw Mode using the Forward+ rendering path. Each grid square is a tile, and each value represents the number of lights affecting the tile.](../../../uploads/urp/lighting-complexity.png)

An example of the Lighting Complexity [Debug Draw Mode](../../GIVis.md) using the Forward+ rendering path. Each grid square is a tile, and each value represents the number of lights affecting the tile.

Unity ignores the following settings if you select the Forward+ rendering path:

* **Additional Lights** in the URP Asset.
* **Main Light** in the URP Asset.
* **Additional Lights** > **Per Object Limit** in the URP Asset.

## Additional resources

* [Light limits in URP](../lighting/light-limits-in-urp.md)
* [Introduction to rendering paths in URP](../rendering-paths-introduction-urp.md)