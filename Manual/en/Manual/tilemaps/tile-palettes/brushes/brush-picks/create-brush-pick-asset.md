# Create a Brush Pick asset

By default, Unity stores the [Brush Picks data asset](../../tile-palette-preferences-reference.md) in the `GridBrush` folder located in the `Library` of your project. You can store this data directly in your project instead, which helps to keep it in source control and consistent with all users of your project.

To change the Brush Picks asset used by the project, go to **Edit > Preferences… > 2D > Tile Palette**, and select and edit the [Brush Picks asset](../../tile-palette-preferences-reference.md) settings in the Brush Picks section of Tile Palette Preferences.

## To create a new Brush Picks asset

* Select **New** to create a new a Brush Picks asset for your project. Unity automatically sets this new asset as the active Brush Picks asset.

## To clone an existing Brush Picks asset

* Select **Clone** to create a new Brush Picks asset that copies the contents of the active Brush Picks asset. If no active Brush Picks asset is present, this clones the Brush Picks asset contained in the project’s `Library\Gridbrush` folder.

**Important:** Unity saves Brush Picks to the same location as the current active Brush Picks asset, or the `Library` folder if no active Brush Picks asset is set. This means that the names of the Brush Picks must not have the same name as another asset in the same location to prevent conflicts.

## Additional resources

* [Tile Palette Brush Picks overlay reference](./brush-picks-overlay-reference.md)
* [Using a Brush Pick](./work-with-brush-picks.md)
* [Tile Palette preferences reference](../../tile-palette-preferences-reference.md)