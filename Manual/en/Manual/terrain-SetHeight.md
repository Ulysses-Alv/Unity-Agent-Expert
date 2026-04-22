# Set the height of an area or tile

The Set Height tool has two main uses:

* Set the height of an entire tile, or multiple tiles, to a specific value. This is important because you can’t work below the **Terrain** tile’s (GameObject) height. To sculpt sunken features, such as lakes or rivers, use the Set Height tool to move the starting height for sculpting above the **GameObject**’s initial height. You can also raise some tiles’ starting point well above other tiles, to create elevation zones in the world. Use sculpting tools such as Bridge, Raise or Lower Terrain, and Slope Flatten to smooth the transition between tiles.
* Use a brush to raise or lower parts of the terrain. Use this to create flat areas that look human-made, such as roads. You can also create cliffs by using Set Height and then adding details with the Raise or Lower Terrain tool.

  **Tip:** Sometimes, the Bridge tool is a better option for creating flat features. For more information, refer to [Build a bridge between two points](terrain-Sculpt-Bridge.md).

You can also use Set Height to reset a tile if you want to restart sculpting.

## Access the tool

To set a height, in the Terrain tile’s **Inspector** window, select the **Paint Terrain** tool > tool dropdown > **Set Height**.

## Tool options

The Set Height tool is brush-based, unless you use the **Flatten Tile** or **Flatten All** options. To learn about working with brushes, refer to [Brushes](class-Brush.md).

## Set Height Controls

The **Height Control** tool has the following properties:

| **Property** | **Function** |
| --- | --- |
| **Space** | Set the starting point for the height value you enter in the **Height** field. Note that the **Set Height** tool can’t lower a tile below its Transform Position y-axis value, even if you enter a **Height** value lower than the y-axis value. The tool also can’t raise the tile above its maximum height property (**Terrain Height** in the [Terrain Settings](terrain-OtherSettings.md)). If you select a height that, when added to the tile’s y-axis value, exceeds the tile’s maximum height property, the value is adjusted down when you select **Flatten Tile** or **Flatten All**, or use a brush. There are two options:  * **World:** The tile’s height is relative to a y-axis value of `0`. * **Local:** The tile’s height is relative to its current y-axis value. |
| **Height** | Enter a height value relative to either the world or local y-axis value. You can also sample a height on the terrain: press **Shift** (macOS: **Control**) and click on the Terrain to sample the height at the cursor position. If you’re using the Terrain Tools package, use press **Ctrl** instead of **Shift**. |
| **Flatten Tile** | Change the entire tile to the specified height, rather than use a brush-based workflow. Note that this flattens all features on the tile. |
| **Flatten All** | Change all **scene** tiles to the specified height, rather than use a brush-based workflow. Note that this flattens all features on the tile. |

## Additional resources

* [Primary sculpting tools](terrain-Basic-Sculpting.md)
* [Advanced sculpting tools](terrain-Sculpt.md)
* [Change existing height with effects](terrain-Effects.md)
* [Transform terrain features](terrain-Transform.md)
* [Erode the terrain](terrain-Erosion.md)