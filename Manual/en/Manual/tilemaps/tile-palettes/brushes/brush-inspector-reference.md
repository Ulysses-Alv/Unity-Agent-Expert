# Brush Inspector window reference

You can use the Brush **Inspector** to change the current [active brush](./active-brush.md) and its properties.

The **Brush Inspector** window is at the bottom of the [Tile Palette editor](../tile-palette-editor-reference.md) window. To minimize or expand the **Brush Inspector**, select the **Brush Inspector** visibility toggle at the upper right corner of the window, or drag the edge of the **Brush Inspector** window.

Select the brush type for the active brush from the dropdown menu from one of the available brushes provided by the [2D Tilemap Extras package](https://docs.unity3d.com/Packages/com.unity.2d.tilemap.extras@latest).

Refer to the [Scriptable Brushes](https://docs.unity3d.com/Packages/com.unity.2d.tilemap.extras@latest/index.html?subfolder=/manual/Brushes.html) documentation for more information on their specific properties and usage.

| Brush Type | Description |
| --- | --- |
| **Brush type dropdown** | Select the [active brush](./active-brush.md)’s brush type from the available options. |
| **[Default Brush](#default-brush-properties)** | The default brush with the default properties. |
| **[GameObject Brush](#gameobject-brush-properties)** | Select this brush to instance, place and manipulate **GameObjects** onto the ****Scene**** view. |
| **[Group Brush](#group-brush-properties)** | Select this brush to pick tiles which are grouped together according to their position and set properties. |
| **[Random Brush](#random-brush-properties)** | Select this brush to place random tiles onto a **tilemap** from a [Tile Set](https://docs.unity3d.com/Packages/com.unity.2d.tilemap.extras@latest/index.html?subfolder=/manual/RandomBrush.html%23usage) you define. For more information about the brush and how to define a Tile Set, refer to the [Random Brush](https://docs.unity3d.com/Packages/com.unity.2d.tilemap.extras@latest/index.html?subfolder=/manual/RandomBrush.html) documentation in the [2D Tilemap Extras package](https://docs.unity3d.com/Packages/com.unity.2d.tilemap.extras@latest). |
| **[Line Brush](#line-brush-properties)** | Select this brush to draw a line of tiles onto the tilemap. |

## Default Brush properties

The following properties are visible by default when you select the **Default Brush**. Some properties are common to all brush types.

| Property | Function |
| --- | --- |
| **Script** | Displays the currently assigned script Asset that provides a fixed set of APIs for painting on Tilemaps. The default is the [GridBrush](../../../../ScriptReference/GridBrushBase.md). Users may use or create their own [Scriptable Brushes](./create-scriptable-brush.md) which become available from the dropdown menu. The Script property updates to reflect the current active Brush. |
| **Flood Fill Contiguous Only** | Enable this property to have the [Flood Fill tool](../tools/fill-area-with-flood-fill-tool.md) only affect Tiles on a Tilemap which are both the same as the targeted Tile and are contiguous to each other from the targeted position. When disabled, Flood Fill will change all Tiles which are the same as the targeted Tile on a Tilemap regardless of their position. This only affects the Default Brush. |
| **Lock Z Position** | Enable this property to change the z-position of the active Brush. Disable to prevent any changes to the current z-position of the active Brush. |
| **Z Position** | Only available when **Lock Z Position** is disabled. Enter the desired z-axis value (only whole numbers) for this Brush when painting Tiles, which also adjusts the relative heights of Tiles on a [Z as Y Isometric Tilemap](../../work-with-tilemaps/isometric-tilemaps/create-isometric-tilemap.md). Refer to [Adjusting the Tile height in the Palette](../../work-with-tilemaps/isometric-tilemaps/create-tile-palette-isometric-tilemap.md) for more information. |
| **Reset** | Select to reset the z-position value back to zero. |

## GameObject Brush properties

The following properties are visible only when you select the **GameObject Brush**. For more information about the brush and how to use it, refer to the [GameObject Brush](https://docs.unity3d.com/Packages/com.unity.2d.tilemap.extras@latest/index.html?subfolder=/manual/GameObjectBrush.html) documentation in the [2D Tilemap Extras package](https://docs.unity3d.com/Packages/com.unity.2d.tilemap.extras@latest).

| Property | Description |
| --- | --- |
| **Cell** | Displays the positions and details of the element(s) contained in this brush in a linear array. This is non-interactable and you can’t change or add to the properties in this section. |
| **Size** | Set the brush size by specifying the **X**, **Y**, and **Z** values in cells for each axis. |
| **Pivot** | Set the pivot coordinates to define the brush’s rotation point. |
| **Anchor** | Set the anchor coordinates of the cells that decide where in the cells the brush places GameObjects when you paint with it. |

## Group Brush properties

The following properties are visible only when you select the **Group Brush**. For more information about the brush and how to use it, refer to the [Group Brush](https://docs.unity3d.com/Packages/com.unity.2d.tilemap.extras@latest/index.html?subfolder=/manual/GroupBrush.html) documentation in the [2D Tilemap Extras package](https://docs.unity3d.com/Packages/com.unity.2d.tilemap.extras@latest).

| Property | Description |
| --- | --- |
| **Gap** | Set the brush’s search radius. The value you set specifies the number of cells along each axis the brush search will scan. The brush then repeats the scan from each selected tile until the **Limit** value is reached. |
| **Limit** | Set the brush’s total range. The value you set specifies the number of cells around the initial picked position. The brush value represents the maximum number of cells from the point of origin that the brush search will scan. |

## Random Brush properties

The following properties are visible only when you select the **Random Brush**. For more information about the brush and how to use and define [Tile Sets](https://docs.unity3d.com/Packages/com.unity.2d.tilemap.extras@latest/index.html?subfolder=/manual/RandomBrush.html%23usage), refer to the [Random Brush](https://docs.unity3d.com/Packages/com.unity.2d.tilemap.extras@latest/index.html?subfolder=/manual/RandomBrush.html) documentation in the [2D Tilemap Extras package](https://docs.unity3d.com/Packages/com.unity.2d.tilemap.extras@latest).

| Property | Description |
| --- | --- |
| **Pick Random Tiles** | Enable this property to pick the tiles from the current selection as a random Tile Set. |
| **Add To Random Tiles** | Enable this property to add the picked Tile Sets to existing Tile Sets instead of replacing them. |
| **Tile Set Size** | Set the size of the Tile Set that this brush paints. |
| **Number of Tiles** | The number of Tile Sets that this brush paints. |

## Line Brush properties

The following properties are visible only when you select the **Line Brush**. For more information about the brush and how to use it, refer to the [Line Brush](https://docs.unity3d.com/Packages/com.unity.2d.tilemap.extras@latest/index.html?subfolder=/manual/LineBrush.html) documentation in the [2D Tilemap Extras package](https://docs.unity3d.com/Packages/com.unity.2d.tilemap.extras@latest).

| Property | Description |
| --- | --- |
| **Fill Gaps** | Enable to have Unity create orthogonal connections between all tiles that connect the start and end of the line painted. |
| **Line Start** | Set the coordinates of the starting point of the line. |

## Additional resources

* [2D Tilemap Extras package](https://docs.unity3d.com/Packages/com.unity.2d.tilemap.extras@latest)
* [Tile Palette editor reference](../tile-palette-editor-reference.md)