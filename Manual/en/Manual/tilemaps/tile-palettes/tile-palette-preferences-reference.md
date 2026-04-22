# Tile palette preferences reference

This page documents the different Preference settings which affect the Tile Palette window’s features and functions. Adjust these settings to customize the default behavior and default tools of the [Tile Palette](./create-tile-palette.md) window while editing.

To access the Tile Palette preferences settings, go to **Edit > Preferences > 2D > Tile Palette** (macOS: **Unity > Settings > 2D > Tile Palette**).

| Setting | Description |
| --- | --- |
| **Target Edit Mode** | This sets the behavior for editing a target when a **Prefab** instance is selected in the **Active Target** list. Editing a Prefab instance in the **Scene** isn’t recommended due to performance issues. **Note**: This affects the **Active Tilemap** list in pre–2021.2 Editor versions. |
| **Enable Dialog** | Displays a dialog for you to choose between [editing the target in Unity’s prefab editing mode](../../EditingInPrefabMode.md) or in the **Scene view**. |
| **Edit In Prefab Mode** | Select this to edit only the target Tilemap in the prefab editing mode. |
| **Edit In Scene** | Select this to edit only the target Tilemap in the Scene view. |
| **Mouse Grid Position At Z** | Enable this setting to have Unity display the Brush marquee at the current z position rather than at the default z = 0 position. This also affects how you position your mouse cursor to draw tiles at different z positions ([refer to examples below](#gridz)). |
| **Active Targets Sorting Mode** | Sets the method for sorting the **Active Targets** (affects **Active Tilemaps** in pre–2021.2 Editor versions) list in the Tile Palette. |
| **Alphabetical** | Sorts the list in alphabetical order, with the first item at the top of the list. |
| **ReverseAlphabetical** | Sorts the list in reverse alphabetical order, with the first item at the top of the list. |
| **Restore Edit Mode Active Target** | Enable this to restore the targeted **Active Target** (**Active Tilemap** in pre–2021.2 Editor versions) in the Tile Palette window after returning to Edit mode from Play mode. |
| **Create Tile Method** | Use this to specify the method Unity uses create [Tile Assets](../tiles-for-tilemaps/tile-asset-reference.md) when you drag a texture or a **sprite** onto a Tile Palette. The default option is **DefaultTile**, which creates a tile type Tile Asset. |
| **Show Open Tile Palette In Scene View** | Enable this to display the **Open Tile Palette** overlay  in the Scene view when selecting an object that interacts with the Tile Palette. |
| **Prefab Mode Grid Cell Size** | Sets the default **Cell Size** when opening a tilemap prefab in the prefab editing mode when there is no valid Grid in the prefab. |
| **Prefab Mode Grid Cell Gap** | Sets the default **Cell Gap** when opening a tilemap prefab in the prefab editing mode when there is no valid Grid in the prefab. |
| **Prefab Mode Grid Cell Layout** | Sets the default **Cell Layout** when opening a tilemap prefab in the prefab editing mode when there is no valid Grid in the prefab. |
| **Prefab Mode Grid Cell Swizzle** | Sets the default **Cell Swizzle** when opening a tilemap prefab in the prefab editing mode when there is no valid Grid in the prefab. |
| **Default Tile Palette Tools** | This determines the list of tools which are available in the Tile Palette window to use with Grid Brushes. The left panel stores the list of tools usable in the Tile Palette. The right panel stores the list of unused tools, which you can add to the available tools. To add a tool to the usable list, select it in the right panel and then select **Add (+)** to add it to the left panel. To remove a tool from the usable list, select it in the left panel and then select **Remove (-)** to remove it from the left panel. |
| **Save** | Save changes made to the Default Tile Palette Tools. |
| **Revert** | Revert changes made to the Default Tile Palette Tools. |
| **Reset** | Resets the Default Tile Palette Tools back to the system default. |

## Brush Picks

| Setting | Description |
| --- | --- |
| **Brush Picks Asset** | Select the asset where Unity stores the Brush Picks in your project. If there is no active asset (**None**), Unity stores the Brush Picks in the project’s Library folder by default. |
| **New** | Creates a new Brush Picks asset in the project and sets this new asset as the active Brush Picks asset. |
| **Clone** | Clones a new Brush Picks asset from the active Brush Picks asset. If there is no active Brush Picks asset, it will clone the Brush Picks present in the Library folder. |

### Effect of Mouse Grid Position At Z

The following examples display the differences between not enabling the setting (the default option) and enabling it. When the setting isn’t enabled, Unity displays the Brush marquee at the default z = 0 position. To draw the following tiles, you position the mouse cursor within the area marked by the white marquee as shown below.

![Example 1: Position of mouse cursor in relation to tiles when **Mouse Grid Position At Z ** is disabled.](../../../uploads/Main/tilemap-grid-z-off.png)

Example 1: Position of mouse cursor in relation to tiles when \*\*Mouse Grid Position At Z \*\* is disabled.

Enable **Mouse Grid Position At Z** to display the Brush marquee at the current z position. To draw tiles in the same position shown in Example 1, position your mouse cursor within the blue marquee area instead.

![Example 2: Position of mouse cursor in relation to tiles when Mouse Grid Position At Z is enabled.](../../../uploads/Main/tilemap-grid-z-on.png)

Example 2: Position of mouse cursor in relation to tiles when **Mouse Grid Position At Z** is enabled.

## Grid Brush preferences reference

Use the following settings to customize the default behavior of the Grid Paint tool when you use it to paint in the Tile Palette window. To access these settings, go to **Edit > Preferences… > 2D > Grid Brush**.

| Setting | Description |
| --- | --- |
| **Show Flood Fill Preview** | Enable this setting to display a preview when you use the Flood Fill tool and hover over the Active Target in the Scene view. If the preview displays many Tiles at once, performance can be negatively affected. You can disable this setting to improve performance. |
| **Flood Fill Preview Fill Extents** | This setting is only available if you enabled **Show Flood Fill Preview**. Set the limit of the extent from the current cursor position that the Flood Fill preview displays. This value is the number of tiles extended from the current cursor position along the x and y-axis that the Flood Fill preview displays. **Important:** Setting this value to 0 causes the Flood Fill preview to have an unlimited extent, which can cause performance issues as the preview includes more tiles. You can improve performance by setting a lower value to limit the number of tiles included in the preview. |
| **Flood Fill Preview Erase Extents** | This setting is only available if you enabled **Show Flood Fill Preview**. Set the limit of the extent from the current cursor position that the Flood Fill preview displays, when the Flood Fill tool erases tiles. This value is the number of tiles extended from the current cursor position along the x and y-axis that the Flood Fill preview displays. **Important:** Setting this value to 0 causes the Flood Fill preview to have an unlimited extent, which can cause performance issues as the preview includes more tiles. You can improve performance by setting a lower value to limit the number of tiles included in the preview. |

## Additional resources

* [Project settings](../../comp-ManagerGroup.md)
* [Tilemap](../work-with-tilemaps/tilemap-reference.md)