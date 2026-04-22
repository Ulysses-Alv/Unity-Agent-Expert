# Pinch or spread out the terrain

**Note:** This tool is only available in the [Terrain Tools package](https://docs.unity3d.com/Packages/com.unity.terrain-tools@latest/).

To pinch elevated parts of the **terrain** to a central point, or to spread them away from that point, use the **Pinch** tool.

The images show how you can create narrow crevices by pulling two ridge lines closer together, and push away the terrain to smooth an area into a paved road.

![An example of the Pinch tool closing the gap between two mountains to create a narrow pass](../uploads/Main/terrain-Pinch-In.gif)

An example of the Pinch tool closing the gap between two mountains to create a narrow pass

![Using the Pinch tool to push away the terrain can help a path look like it was dug out of the ground rather than laid down on the surface](../uploads/Main/terrain-Pinch-Out.gif)

Using the Pinch tool to push away the terrain can help a path look like it was dug out of the ground rather than laid down on the surface

## Access the tool

To pinch the terrain, in the Terrain tile’s **Inspector** window, select the **Paint Terrain** tool > tool dropdown > **Transform** > **Pinch**.

## Tool options

The Pinch tool is brush-based. To learn about working with brushes, refer to [Brushes](class-Brush.md).

**Tip:** To spread rather than pinch, hold **Ctrl** (macOS: **Cmd**).

## Pinch controls

| **Property** | **Function** |
| --- | --- |
| **Targets** | Impact only the material, only the height, or both. |
| **Pinch Amount** | Negative values push the terrain away from the center of the brush. Positive values pull the terrain toward the center of the brush. The range is from –100 to 100. |

**Tip**: The change in the materials is only visible if there are at least two materials in the pinch area.

## Additional resources

* [Smudge the terrain](terrain-Transform-Smudge.md)
* [Twist the terrain](terrain-Transform-Twist.md)