# Twist the terrain

**Note:** This tool is only available in the [Terrain Tools package](https://docs.unity3d.com/Packages/com.unity.terrain-tools@latest/).

To rotate elevated parts of the **terrain** around a central point, use the Twist tool.

![Use Twist to rotate elevated parts of the terrain around a central point](../uploads/Main/terrain-Transform-Twist.png)

Use **Twist** to rotate elevated parts of the terrain around a central point

## Access the tool

To twist the terrain, in the Terrain tile’s **Inspector** window, select the **Paint Terrain** tool > tool dropdown > **Transform** > **Twist**.

## Tool options

The Twist tool is brush-based. To learn about working with brushes, refer to [Brushes](class-Brush.md).

**Tip:** To invert the twist direction, hold **Ctrl** (macOS: **Cmd**).

## Twist controls

| **Property** | **Function** |
| --- | --- |
| **Targets** | Impact only the material, only the height, or both. |
| **Smudge Amount** | Positive values rotate the terrain clockwise. Negative values rotate the terrain counterclockwise. The range is from –100 to 100. |

**Tip**: The change in the materials is only visible if there are at least two materials in the smudge area.

## Additional resources

* [Pinch the terrain](terrain-Transform-Pinch.md)
* [Smudge the terrain](terrain-Transform-Smudge.md)