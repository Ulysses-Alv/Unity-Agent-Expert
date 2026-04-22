# Create tile assets

This page describes how to create tile assets in Unity. You can create a tile asset directly from the Unity Editor’s **Assets** menu, or have Unity automatically generate tile assets when you drag **sprites** onto a tile palette.

Refer to the specific tasks for either method:

* [Create tile assets from the Asset menu](#asset-menu)
* [Automatically generate tile assets in the Tile Palette window](#automatic-generation)

## Prerequisites

You must install the [2D Tilemap Editor package](../../com.unity.2d.tilemap) to have the option to create tile assets in the **Assets** menu. Install the [2D Tilemap Extras](../../com.unity.2d.tilemap.extras) for additional types of tile assets.

**Note**: These packages are part of the [2D feature set](../../2DFeature.md) and are automatically installed if you select the **2D** option when creating a new project.

## Create tile assets from the Asset menu

To create empty tile assets in the Unity Editor:

1. Select **Assets** from the menu bar.
2. Go to **Create** > **2D** > **Tiles** to see the list of available tile types. If you don’t see any options, check that you have completed the prerequisites.
3. Select a type of tile asset from the available options. Unity creates the asset in the current folder in the **Project** window.

## Automatically generate tile assets in the Tile Palette window

To automatically generate tile assets in the [Tile Palette window](../tile-palettes/tile-palette-editor-reference.md):

1. To open the **Tile Palette** window, go to **Window** > **2D** > **Tile Palette**.
2. Open the **Active Palette** dropdown menu in the **Tile Palette** window. Select a tile palette from the list, or create new palette.
   ![The Tile Palette window with the active palette dropdown expanded.](../../../uploads/Main/create-tile-new-palette.png)
3. After creating or loading a tile palette, click and drag textures or sprites from the `Assets` folder directly onto the **Tile Palette** window. The editor will prompt you to choose a save location for new tile assets.
4. Select a folder as your save location, and the editor will generate new tile assets based on the imported textures or sprites, in the selected folder. The generated tiles are also automatically placed onto the palette.

## Additional resources

* [Create a tile palette](../tile-palettes/create-tile-palette.md)
* [Tile palette editor tools](../tile-palettes/tools/tile-palette-tools-landing.md)
* [Tile asset reference](./tile-asset-reference.md)