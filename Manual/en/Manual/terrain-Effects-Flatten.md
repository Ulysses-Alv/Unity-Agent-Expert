# Flatten terrain features

**Note:** This tool is only available in the [Terrain Tools package](https://docs.unity3d.com/Packages/com.unity.terrain-tools@latest/).

Use the Slope Flatten tool to flatten the details on top of a slope without changing the slope’s average angle. For example, you can use it to:

* Make a mountain top less jagged, without making the mountain less steep.
* Smooth the transition from a plain to the mountain’s side.
* Make terraces, bridges, and other artificial features look unmaintained. Refer to [Erode the terrain](terrain-Erosion.md) for other ways to do this.

## Access the tool

To flatten the terrain, in the Terrain tile’s **Inspector** window, select the **Paint Terrain** tool > tool dropdown > **Effects** > **Slope Flatten**.

**Tip**: Use a low **Brush Strength** value to limit the impact on the feature’s height. The image shows the difference in impact between Brush Strength values of `0.2` and `0.8`:

![The image shows an original mountain range, and the impact of two different Brush Strength values: 0.2 and 0.8.](../uploads/Main/terrain-Flatten-Brush.png)

The image shows an original mountain range, and the impact of two different Brush Strength values: 0.2 and 0.8.

## Tool options

The Slope Flatten tool is brush-based. To learn about working with brushes, refer to [Brushes](class-Brush.md).

## Additional resources

* [Erode the terrain](terrain-Erosion.md)
* [Sharpen peaks](terrain-Effects-Sharpen.md)
* [Add contrast to features](terrain-Effects-Contrast.md)