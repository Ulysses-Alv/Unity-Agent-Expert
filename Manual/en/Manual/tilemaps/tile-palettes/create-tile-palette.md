# Create a Tile Palette

Place a selection of Tiles onto **Tile Palettes** so that you can pick Tiles from the Palette to paint on **Tilemaps**.

## Prerequisites

To create a tilemap you must first ensure you have done the following:

* Download the **2D Tilemap Editor** package from the [Package Manager](../../Packages.md).

## Create the tile palette asset

To create the tile palette asset, do the following:

1. To create a **Tile Palette**, open the **Tile Palette** window by going to **Window** > **2D** > **Tile Palette**.
2. Select the **New Palette** drop-down menu to open a list of Tile Palettes available in the Project, or for the option to create a new Palette. Select the **Create New Palette** option to create a new Palette.

   ![The Tile Palette window with Create New Palette selected in the Tiles dropdown.](../../../uploads/Main/2d-tile-palette-create-new.png)

   The Tile Palette window with Create New Palette selected in the Tiles dropdown.
3. After selecting the option to create a new Tile Palette, the **Create New Palette** dialog box becomes available. It contains the different property settings and options available when creating a new Palette. Name the newly created Palette and select the desired settings, then select the Create button. Select the folder to save the Palette Asset file into when prompted. The newly created Palette is automatically loaded in the Tile Palette window.

   ![A blank Tile Palette.](../../../uploads/Main/2d-tile-palette-blank.png)

   A blank Tile Palette.
4. Drag and drop Textures or **Sprites** from the Assets folder onto the **Tile Palette**, and choose where to save the new [Tile Assets](../tiles-for-tilemaps/tile-asset-reference.md) when prompted. New Tile Assets are generated in the selected save location, and the Tiles are placed on the grid of the active Tile Palette window.

   ![Drag and drop directly onto the Tile Palette window.](../../../uploads/Main/2d-tile-palette-drag-drop.png)

   Drag and drop directly onto the Tile Palette window.
5. Use **inspector** to change the current active Brush and its properties. For more information on these properties, refer to the [Brush Inspector window reference](brushes/brush-inspector-reference.md).

## Editing the Tile Palette

The tools for picking and painting with Tiles can also be used to edit the Tile Palette directly, allowing you to move and manipulate the Tiles currently placed on the **Tile Palette**. Select the Palette you want to edit from the **Palette** dropdown menu (the default Palette is named ‘New Palette’), then select **Edit** to unlock the Palette for editing.

![The Tile Palette Edit toggle.](../../../uploads/Main/2d-tile-palette-toggle-edit.png)

The Tile Palette Edit toggle.

Refer to [Tile Palette editor reference](./tile-palette-editor-reference.md) for the shortcuts and functions of the Palette tools, which can also be used to edit the Palette.

## Creating Palette Assets from existing Grid Prefabs

You can convert an existing [Prefab](../../Prefabs.md) to a Palette Asset, so that you can use it in the Tile Palette window. To do this, the Prefab must not already be a Palette Asset, and it must have a Grid component on its topmost ****GameObject****.

![The Tile Palette toolbar.](../../../uploads/Main/2d-tile-palette-active-brush-toolbar.png)

The Tile Palette toolbar.

To convert a Prefab, drag and drop it onto the Tile Palette **toolbar** (highlighted in the image above). The Editor automatically converts it to a Palette Asset, and adds a Grid Palette Asset. The new Palette Asset has the same name as its source, and it becomes available in the Palette dropdown menu.

## Use a Tile Set as a Tile Palette

You can create a Tile Palette pre-populated with Tiles from the Sprites of your chosen Textures by creating a [Tile Set](tile-set-landing.md). The Tile Set automatically creates and updates the Tiles from your chosen Textures whenever the Textures change.

For more information on how to create a Tile Set, refer to [Create a Tile Set](tile-set-create.md).

## Tile Palette Grid visibility toggle

![The Tile Palette Grid toggle.](../../../uploads/Main/2d-tile-palette-toggle-grid.png)

The Tile Palette Grid toggle.

Switch the visibility of the Grid on the Tile Palette on or off by selecting the toggle highlighted above.

## Tile Palette Gizmos visibility toggle

The Tile Palette can display [Gizmos](../../GizmosMenu.md) over the currently selected Palette Asset, to help you visualize specific criteria. For example, you can add a Gizmo that displays a special icon for Tiles that contain no Sprites.

![The Tile Palette Gizmos toggle.](../../../uploads/Main/2d-tile-palette-gizmos.png)

The Tile Palette Gizmos toggle.

To display the default Unity and the Palette Asset’s Gizmos on the Tile Palette, enable the Gizmos toggle (highlighted above). The Tile Palette immediately displays any component with [MonoBehaviour.OnDrawGizmos()](../../../ScriptReference/MonoBehaviour.OnDrawGizmos.md) in the Palette Asset.

To add your own custom gizmos to a Palette Asset, add a component with `DrawGizmo` to the Palette Asset:

1. Select the Palette Asset in the ****Project window****.
2. Open the Palette Asset in [prefab editing mode](../../EditingInPrefabMode.md).
3. Add the component while in prefab editing mode.
4. Save the asset while in prefab editing mode.
5. Exit prefab editing mode.