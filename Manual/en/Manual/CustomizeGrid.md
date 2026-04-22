# Customize the grid

This section provides information on how to customize the grid in the **Scene** view window.

## Display and hide grid lines

To toggle whether the grid is visible on all axes, select **Toggle grid visibility** (![Toggle grid visibility icon](../uploads/Main/SceneGrids-Vis-icon-inline.png)) in the [Grid and Snap](GridAndSnapOverlay).

## Change the axis that the grid appears on

If you are in [orthographic mode](SceneViewNavigation.md) (Iso) and align your view with an axis, the Editor chooses to display the grid axis orthogonal to that axis.

To change which axis the grid appears on:

1. In the **Grid and Snap** overlays **toolbar**, select the downward facing arrow next to **Toggle grid visibility** (![Toggle grid visibility icon](../uploads/Main/SceneGrids-Vis-icon-inline.png)) to open the **Grid Visual** menu. .
2. In the **Grid Plane** section, select the axis you want to display.

## Resize the grid

You can set the size of the grid lines that display in the **Scene view** window. The size of the grid affects how your grid looks and how your **GameObjects** [automatically snap to the grid](GridSnap.md) and how your GameObjects [move, rotate, or scale using snap increments](SnapIncrements.md).

If you set a size for all axes at once, you display a uniform, square-based, grid. However, you can also use different values on any of the three axes to display a non-uniform, rectangular-based, grid. By default, the grid is set to a uniform size of 1 on all axes.

You can also use the **Grid size field** in the [Grid and Snap](GridAndSnapOverlay) to set a uniform grid size without opening the **Snapping** menu.

To resize the grid:

1. In the **Grid and Snap** overlay, select the arrow next to the **Grid Snapping** button to open the **Snapping** menu.
2. To enter a uniform value for the grid size and make all grid lines the same length:
   1. Select the link icon in the **Size** property.
   2. Enter a value for the **X** property to set the size of the grid lines on all axes.
3. To enter non-uniform values for the grid size and specify the size of each grid axis separately:
   1. Make sure that the link icon in the **Size** property is not selected.
   2. Enter a value for the **X**, **Y**, and **Z** properties to set the size of the grid lines on each axis.

**Tip**: You can also use [these keyboard shortcuts](GridShortcuts.md) to increase or decrease the size of the grid.

## Change the default color of the grid lines

To change the color of the visible grid lines in the **Scene view** window:

1. Open the Unity [Preferences](Preferences.md) window.
2. Select the **Colors** tab in the Preferences window.
3. Use the color picker to set a new color.

## Change the opacity of the grid

To adjust the brightness of the grid lines:

1. From the **Grid and Snap** overlays toolbar, select the downward facing arrow next to **Toggle grid visibility** (![Toggle grid visibility icon](../uploads/Main/SceneGrids-Vis-icon-inline.png)) to open the **Grid Visual** menu.
2. Use the **Opacity** slider to make the grid more or less visible in the Scene view.

## Move the grid to the handle of a GameObject

To move the grid to the handle of a GameObject:

1. Select a GameObject.
2. In the **Grid and Snap** overlay, select the downward facing arrow next to **Toggle grid visibility** (![Toggle grid visibility icon](../uploads/Main/SceneGrids-Vis-icon-inline.png)) to open the **Grid Visual** menu.
3. In the **Grid Visual** menu, in the **Move to** property, select **Handle** to move the grid to the handle of the selected GameObject.
4. To move the grid back to its default position, in the **Move to** property, select **Origin**.

## Reset grid values and settings to default

To reset the grid axis and opacity settings to their defaults:

1. In the **Grid and Snap** overlay, select the downward facing arrow next to **Toggle grid visibility** to open the **Grid Visual** menu (![Toggle grid visibility icon](../uploads/Main/SceneGrids-Vis-icon-inline.png)).
2. Select the **More** menu (⋮).
3. Select **Reset**.

To reset the size of the grid and the increment snap values to their defaults:

1. In the **Grid and Snap** overlays toolbar, select the downward facing arrow next to **Toggle grid snapping** (![Toggle grid snapping icon](../uploads/Main/SceneGrids-Mode-icon-inline.png)) to open the **Snapping** menu.
2. Select the **More** menu (⋮).
3. Select **Reset**.

## Additional resources

* [Scene view grid snapping](GridSnapping.md)
* [Grid and Snap overlay](GridAndSnapOverlay)
* [Overlays](overlays.md)
* [The Scene view](UsingTheSceneView.md)
* [Position GameObjects](PositioningGameObjects.md)