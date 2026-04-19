# Active brush

The active brush is the current brush you are using to paint on a **tilemap**. The active brush controls the behavior and properties of picked tiles as you paint on a tilemap.

You use the **Default Brush** by default when you open the [Tile Palette editor](../tile-palette-editor-reference.md). You can select other types of brushes from the [brush type](../tile-palette-editor-reference.md) dropdown menu in the [Brush Inspector window](./brush-inspector-reference.md). Each of these brush types have different behaviors and uses. Refer to the [2D Tilemap Extras package](https://docs.unity3d.com/Packages/com.unity.2d.tilemap.extras@latest) documentation for more information and API examples.

After selecting tiles from the tile palette, you can then paint on the tilemap with the active brush. The **Tile Palette** editor displays a [brush preview](#brushpreview) at your cursor in the ****Scene**** view, which displays how the tile will appear before you paint it on the tilemap.

## Brush preview

This brush preview helps you visualize the placement and orientation of your tiles before you paint them onto the tilemap. You can use several shortcuts to change the orientation or position of the active brush, which changes the orientation or position of the tiles you paint with that brush. For a list of these shortcuts, refer to the [Active brush shortcuts reference](./active-brush-shortcuts-reference.md).

**Note:** When painting with multiple tiles selected, the brush preview can change as the cursor hovers over neighboring rows or columns. This effect is more pronounced when painting on [Hexagonal Tilemaps](../../work-with-tilemaps/hexagonal-tilemaps.md).

![Unity displays a preview of the picked tile at the cursor in the Scene view.](../../../../uploads/Main/2d-palette-brush-preview.png)

Unity displays a preview of the picked tile at the cursor in the **Scene** view.

## Additional resources

* [Tile Palette editor reference](../tile-palette-editor-reference.md)
* [Scriptable Brushes](./create-scriptable-brush.md)
* [2D Tilemap Extras package](https://docs.unity3d.com/Packages/com.unity.2d.tilemap.extras@latest)