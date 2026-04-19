# Add contrast to features

**Note:** This tool is only available in the [Terrain Tools package](https://docs.unity3d.com/Packages/com.unity.terrain-tools@latest/).

Use the Contrast tool to make low features lower and high features higher. This is distinct from the [Sharpen Peaks](terrain-Effects-Sharpen.md) and [Slope Flatten](terrain-Effects-Flatten.md) tools, which can keep the average height of the feature.

![Contrast added to a hill side, creating more distinct gullies](../uploads/Main/terrain-Contrast-Result.png)

Contrast added to a hill side, creating more distinct gullies

## Access the tool

To add contrast to the terrain, in the Terrain tile’s **Inspector** window, select the **Paint Terrain** tool > tool dropdown > **Effects** > **Contrast**.

## Tool options

The Contrast tool is brush-based. To learn about working with brushes, refer to [Brushes](class-Brush.md).

## Tool controls

The **Feature Size** property controls how large a feature needs to be to have contrast applied to it. The range is from 1 to 100. A low value adds details to smaller features. A high value adds details only to large features.

![Showing the difference between the brush strokes for Feature Size values of 15 and 95](../uploads/Main/terrain-Contrast-Feature-Size.png)

Showing the difference between the brush strokes for Feature Size values of 15 and 95

## Additional resources

* [Flatten terrain features](terrain-Effects-Flatten.md)
* [Change peak sharpness](terrain-Effects-Sharpen.md)