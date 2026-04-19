# Customize a scriptable Brush’s thumbnail

Unity determines the preview image displayed in a Brush Pick’s thumbnail by the [Asset Preview](../../../../../ScriptReference/AssetPreview.md) utility. To customize the thumbnail image of your [scriptable brush](../create-scriptable-brush.md), you will need to implement the [RenderStaticPreview](../../../../../ScriptReference/Editor.RenderStaticPreview.md) method for the [editor](../../../../../ScriptReference/Editor.md) of your scriptable brush.

![Example of Brush Pick thumbnail.](../../../../../uploads/Main/brush-pick-thumbnail.png)

Example of Brush Pick thumbnail.

The icon in the lower right of the thumbnail preview displays the [Brush type](../brush-inspector-reference.md) of the Brush Pick. The icon is determined by the `icon` property of the editor for the Brush. Override this property to provide your own icon for your Brush type.

![This message appears when a compatible Brush isnt present.](../../../../../uploads/Main/tile-palette-brush-pick-invalid.png)

This message appears when a compatible Brush isn’t present.

Unity stores Brush Picks as serialized values based on the active Brush. If you change the script or remove your custom scriptable brush, Unity might not be able to load saved Brush Picks as the Brush has changed. A message appears if Unity isn’t able to find a compatible Brush for the Brush Pick.

## Additional resources

* [Scriptable Brushes](../create-scriptable-brush.md)
* [Scriptable Tiles](../../../tiles-for-tilemaps/scriptable-tiles/scriptable-tiles.md)