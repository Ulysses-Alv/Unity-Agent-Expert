# Change peak sharpness

**Note:** This tool is only available in the [Terrain Tools package](https://docs.unity3d.com/Packages/com.unity.terrain-tools@latest/).

Use the Sharpen Peaks tool to sharpen an existing peak by:

* Changing the existing peak’s slope.
* Removing height and detail from the area around the peak.

![Side by side comparison of a peak and its sharpened version.](../uploads/Main/terrain-Effects-Sharpen.png)

Side by side comparison of a peak and its sharpened version.

## Access the tool

To sharpen the terrain, in the Terrain tile’s **Inspector** window, select the **Paint Terrain** tool > tool dropdown > **Effects** > **Shparen Peaks**.

## Tool options

The Sharpen Peaks tool is brush-based. To learn about working with brushes, refer to [Brushes](class-Brush.md).

## Tool controls

To control how sharp the peak becomes, use the **Sharpness** slider. The range is from 0 to 1, where 1 is sharp and 0 is smooth.

The **Sharpness** controls only the angles that the brush creates. Use the **Brush Strength** control to change the impact of those angles.

## Additional resources

* [Flatten terrain features](terrain-Effects-Flatten.md)
* [Add contrast to features](terrain-Effects-Contrast.md)