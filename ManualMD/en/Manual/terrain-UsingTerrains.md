# Create and edit Terrains

To add a Terrain GameObject to your Scene, select **GameObject > 3D Object > Terrain** from the menu. This also adds a corresponding Terrain Asset to the Project view. When you do this, the landscape is initially a large, flat plane. The Terrain’s Inspector window provides a number of tools to create detailed landscape features.

![Terrain editing tools in the Inspector](../uploads/Main/1.1-TerrainInspector.png)

Terrain editing tools in the Inspector

Terrain editing tools in the Inspector

The **toolbar** provides five options to adjust your Terrain:

* [Create adjacent Terrain tiles](terrain-CreateNeighborTerrains.md).
* [Sculpt](terrain-Height-Landing.md) and [paint](terrain-Textures-Landing.md) your Terrain.
* Add [trees](terrain-Trees-Landing.md).
* Add [details such as grass, flowers, and rocks](terrain-Grass.md).
* Change [settings for the selected Terrain](terrain-OtherSettings.md).

Select the paintbrush icon to access painting tools, which allow you to modify the Terrain. Use the cursor to sculpt the height of the Terrain, or paint texture onto the Terrain. Choose from several built-in Brush shapes, or define your own Brush using a texture. You can also change the size and opacity (the strength of the applied effect) of the Brush. Once you’ve defined the properties, your cursor takes the shape of the selected Brush. Click or drag on the Terrain to create different shapes and textures.

Similar to how you paint with a Brush on the Terrain, you can add textures, trees, and details like grass, flowers, and rocks. You can also create additional connected Terrain tiles, change the height of an entire tile, and even write custom Brushes with complex effects.

## Terrain Overlay Windows

Terrain tools can be used with overlays. When a Terrain **GameObject** is selected, overlays will appear in the **Scene** View. These overlays can be docked, dragged, and collapsed.

The **Terrain Toolbar** overlay contains the terrain tools, divided into categories. From left to right, the menu buttons are **Sculpt Mode**, **Materials Mode**, **Foliage Mode**, and **Neighbor Terrains Mode**. If there are custom tools which inherit from [TerrainAPI.TerrainPaintToolWithOverlays](../ScriptReference/TerrainTools.TerrainPaintToolWithOverlays_1.md), they will appear in a fifth menu button for **Custom Brushes Mode**.

![Terrain Toolbar](../uploads/Main/terrainOverlays-TerrainToolbar.png)

Terrain Toolbar

The **Tool Settings** overlay can be found under the overlay with this icon ![Tool settings icon](../uploads/Main/terrainOverlays-toolSettingsIcon.png). Some tools have additional settings that can be edited. For example, here is what the tool settings overlay looks like for the Set Height Tool.

![Set Height Tool Settings](../uploads/Main/terrainOverlays-ToolSettingsSetHeight.png)

Set Height Tool Settings

Certain tools may have brush masks. When selecting a tool that has access to brushes, the **Brush Masks** overlay will be visible.

![Brush Masks in Overlays](../uploads/Main/terrainOverlays-BrushMasks.png)

Brush Masks in Overlays

Certain tools which have access to brushes may also have access to the **Brush Attributes** overlay, in which users can change tool settings like brush opacity and size.

![Brush Attributes in Overlays](../uploads/Main/terrainOverlays-BrushAttributes.png)

Brush Attributes in Overlays

The Paint Details tool has access to Target Strength in this overlay.

![Brush Attributes Target Strength in Overlays](../uploads/Main/terrainOverlays-BrushAttributesTargetStrength.png)

Brush Attributes Target Strength in Overlays

## Terrain keyboard shortcuts

The Terrain **Inspector** has the following keyboard shortcuts enabled by default.

| Keys | Function |
| --- | --- |
| Comma ( , ) and Period ( . ) | Cycle through the available Brushes |
| Shift-Comma ( < ) and Shift-Period ( > ) | Cycle through the available objects for trees, textures, and details |
| Open Bracket ( [ ) and Close Bracket ( ] ) | Decrease and increase the Brush Size |
| Minus ( - ) and Equal ( = ) | Decrease and increase Brush Opacity |

To set your own custom shortcuts, use the **Shortcuts Manager**.

* On Windows and Linux, select **Edit > Shortcuts**.
* On macOS, select **Unity > Shortcuts**.

Under **Category**, select **Terrain** to display Terrain-related shortcut assignments. For more information about setting and modifying shortcuts, refer to [Keyboard shortcuts](ShortcutsManager.md).

![List of Terrain shortcuts](../uploads/Main/1.1-TerrainShortcuts.png)

List of Terrain shortcuts

Additionally, the standard F keystroke works slightly differently for Terrain. As Terrains are typically very large, when you press the F key, the **Scene view** focuses on the area where you position the cursor. This provides a very quick and intuitive way to jump to the Terrain area you want to edit. If you press F while the cursor is not over a Terrain tile, it uses the standard framing behavior, framing the selection around the whole GameObject when you place the cursor over the Scene view.

## Additional resources

* [Scripting Reference for custom Terrain tools](../ScriptReference/TerrainTools.TerrainPaintTool_1.md)