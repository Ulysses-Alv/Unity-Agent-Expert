# Grid and Snap overlay reference

Use the **Grid and Snap** overlay to modify how the **Scene** grid looks, how **GameObjects** snap to the grid, and how GameObjects move incrementally.

The settings on the **Grid and Snap** overlay are global to all **Scene views**.

From the **Grid and Snap** overlay, you can select the dropdown arrows next to the **Toggle grid visibility** and **Toggle grid snapping** buttons to open the [Grid Visual](#grid-visual-menu-reference) and [Snapping](#snapping-menu-reference) menus respectively.

## Grid and Snap overlay

| **Property** | **Description** |
| --- | --- |
| **Grid size field** | Enter a value to use as the current grid size value. This value sets the size of the grid homogeneously on all axes. |
| **Toggle grid visibility** | Select **Toggle grid visibility** to display or hide the grid in the Scene view. |
| **Open the Grid Visual menu** | Select the downward facing arrow next to **Toggle grid visibility** to open the [Grid Visual](#grid-visual-menu-reference) menu. |
| **Toggle grid snapping** | Enable [automatic snapping to the grid](GridSnap.md). The icon is blue when automatic snapping to the grid is enabled. |
| **Open the Snapping menu** | Select the downward facing arrow next to **Toggle grid snapping** to open the [Snapping](#snapping-menu-reference) menu. |

## Grid Visual menu reference

| **Property** | **Description** |
| --- | --- |
| **Grid Plane** | Select the axes you want to display the grid on. |
| **Opacity** | Use the opacity slider to [change the opacity of the grid](CustomizeGrid.md#change-the-opacity-of-the-grid). |
| **Move to** | Select **Handle** to move the grid to the handle of the selected GameObject. Select **Origin** to move the grid back to its default position of (0, 0, 0). |

## Snapping menu reference

| **Property** | **Description** |
| --- | --- |
| **Grid Size** | Set the [size of the grid visuals and the increment snap value to use for the Move tool](CustomizeGrid.md#resize-the-grid). |
| **Snap to Grid** | Enable snapping GameObjects to the absolute position on the grid. This option only functions when your handle rotation is set to **Global**. |
| **Rotate** | Enter a [increment snap value](SnapIncrements.md) for the **Rotate** tool to use. |
| **Scale** | Enter a [increment snap value](SnapIncrements.md) for the **Scale** tool to use. |
| **Align Selected** | Select the axes you want to [align the selected GameObject to](GridAlign.md). |

## Additional resources

* [Scene view grid snapping](GridSnapping.md)
* [Customize the grid](CustomizeGrid.md)
* [Overlays](overlays.md)