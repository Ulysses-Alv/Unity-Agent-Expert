# Create a tilemap

A **tilemap** is a **GameObject** which acts as a grid that you place your selected tiles on.

There are multiple types of tilemap, these are:

* Rectangular
* Hexagonal Point Top
* Hexagonal Flat Top
* Isometric
* Isometric Z as Y

The default tilemap is Rectangular. Refer to the respective pages for [Hexagonal](./hexagonal-tilemaps.md) and [Isometric](isometric-tilemaps/isometric-tilemap-landing.md) tilemaps for more information on their specific features and uses.

## Create a tilemap asset

To create a tilemap asset, do the following:

1. Right-click in Hierarchy window and select **2D Object** > **Tilemap**.
2. Select the type of tilemap you want to create from the options available.

After creating a tilemap, Unity creates a new [Grid](../grid-reference.md) GameObject with a child tilemap GameObject in the **scene**. The Grid GameObject determines the layout of its child tilemaps. The child tilemap consists of the [Tilemap component](./tilemap-reference.md) and [Tilemap Renderer](./tilemap-renderer-reference.md) component. Unity paints tiles onto the tilemap GameObject.

**Note**: If you don’t have these options in the menu bar, you may not have the **2D Tilemap Editor** package installed. In this scenario, download the **2D Tilemap Editor** package from the [Package Manager](../../Packages.md).

### Create a tilemap asset in the Tile Palette window

You can also create a new tilemap from the **Tile Palette** window. To do this, do the following:

1. Open the Tile Palette window.
2. If you have a tile palette you want to create a tilemap asset for, open that tile palette in the Tile Palette window.
3. In the **Active Target** dropdown menu, select the **Create New Tilemap** option.
4. Select the type of tilemap you want to create. If you have an active tile palette, select **From Tile Palette** to create a new tilemap with the same settings as the tile palette.

## Create additional tilemaps

You can create additional tilemaps for a Grid with the following steps:

1. Select the Grid in the Hierarchy window.
2. Right-click on the selected GameObject and go to **2D Object** > **Tilemap** and select the type of tilemap you want.

If the type of tilemap you select does not match the type of grid, you may encounter a dialog with a warning. For more information on how to handle this, refer to [Troubleshooting mismatched Cell Layouts](troubleshoot-mismatched-cell-layouts.md).

## Update tilemap asset properties

After creating the tilemaps, you can adjust the properties of the [Grid](../grid-reference.md) GameObject to adjust the properties of its child tilemaps. This ensures consistency across all the child tilemaps of a Grid. These changes also affect attached components such as the [Tilemap Renderer](./tilemap-renderer-reference.md) and [Tilemap Collider 2D](./tilemap-collider-2d-reference.md) components.

## Additional resources

* [Tilemap class](../../../ScriptReference/Tilemaps.Tilemap.md) (Scripting API)