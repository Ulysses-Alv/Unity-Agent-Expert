# Build a terrace

**Note:** This tool is only available in the [Terrain Tools package](https://docs.unity3d.com/Packages/com.unity.terrain-tools@latest/).

The Terrace tool creates a series of flat areas up a slope, forming steps. Terraces are associated with agriculture, salt production, and other human activities.

## Access the tool

To sculpt a terrace, in the Terrain tile’s **Inspector** window, select the **Paint Terrain** tool > tool dropdown > **Sculpt** > **Terrace**.

## Tool options

The Terrace tool is brush-based. To learn about working with brushes, refer to [Brushes](class-Brush.md).

## Terrace controls

| **Property** | **Description** |
| --- | --- |
| **Terrace Count** | The number of steps in the terrace. The Unity Editor divides the slope under the brush into equal slices to create the number of steps you specified; if the **Terrace Count** value is high but the painted area is low to the ground, the steps can be so narrow they’re hard to see. |
| **Jitter** | Vary the terrace count. The range is from 0 to 1. The higher the value, the greater the difference in the number of steps created with each brush click. |
| **Interior Corner Weight** | Controls the angle of the drop between the steps. The range is from 0 to 1. A value of 0 creates a sheer drop from step to step. A value of `1` creates a sloped drop, giving an eroded look and less distinct steps. |

## Additional resources

* [Erode the terrain](terrain-Erosion.md)
* [Build a bridge between two points](terrain-Sculpt-Bridge.md)