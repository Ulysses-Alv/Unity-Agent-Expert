# Before you start

Before you start working on height, consider the following options and settings.

## Maximum height and depth

Each Terrain tile has a maximum height. When you reach that height, the top of your mountain flattens. You can’t raise it anymore or add any detail. To go higher, change the value of the tile’s **Terrain Height** property in **Terrain Settings** > **Mesh Resolution**. For more information, refer to [Terrain Settings reference](terrain-OtherSettings.md).

You can’t lower the **terrain** below the Terrain tile’s surface (the GameObject’s y-axis position). If you want to sculpt sunken features, such as lakes or rivers, use the Set Height tool to set a starting height for sculpting. For more information, refer to [Set the height of an area or tile](terrain-SetHeight.md).

## Heightmap resolution

If your terrain’s sculpting doesn’t look detailed enough, you can raise the **heightmap**’s resolution. For more information, refer to [Terrain Settings reference](terrain-OtherSettings.md).

## Start sculpting with an imported heightmap

If you have a heightmap from a 3D modeling program, you can import it for use on one or more tiles. For more information, refer to [**Work with heightmaps**](terrain-Heightmaps.md). You can then sculpt the heightmap to add detail.

## Additional resources

* [Create and edit Terrains](terrain-UsingTerrains.md)
* [The Terrain Tools package](https://docs.unity3d.com/Packages/com.unity.terrain-tools@latest/)
* [Terrain colliders](terrain-colliders.md)