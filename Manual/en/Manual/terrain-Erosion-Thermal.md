# Thermal erosion

**Note:** This tool is only available in the [Terrain Tools package](https://docs.unity3d.com/Packages/com.unity.terrain-tools@latest/).

Thermal erosion limits a **terrain**’s maximum slope based on the material it’s made of. This is a natural phenomenon that depends not on wind and water, but on the properties of the soil itself. Different soil types have a different angle of repose, also known as resting angle. This is the steepest angle at which a sloping surface formed of a particular loose material is stable; at a steeper angle, the material starts to slide downhill.

Thermal erosion limits both the terrain’s ability to form large formations such as cliffs and mountains, and small details within those formations. For example, wet clay smooths out fine details and doesn’t form steep mountains, whereas granite scree can form cliffs with detailed surfaces.

In nature, thermal erosion is a short-distance process; the sediment moves only until it reaches an area of the slope where it can rest again. This is in contrast to wind and hydraulic erosion, which can carry sediment over long distances. The Thermal Erosion tool simulates this by moving sediment only until the sediment reaches its angle of repose; after that, the tool has no more impact.

**Note:** Thermal erosion doesn’t increase angles, only decreases them. If you want to create steeper slopes, you can use the **Raise or Lower Terrain** tool.

## Physical Material Presets and custom soil types

The Unity Editor has built-in soil types, each with its own resting angle. These presets are based on the real values of the soil types they represent.

If you change the value of the **Resting Angle** property of any preset, the **Physical Material Presets** changes to **Custom**. In other words, you can’t change the resting angle of a preset soil type, only use the preset type as a starting point for a custom type.

## Access the tool

To erode the terrain, in the Terrain tile’s **Inspector** window, select the **Paint Terrain** tool > tool dropdown > **Erosion** > **Thermal**.

## Tool options

The Erosion tool is brush-based. To learn about working with brushes, refer to [Brushes](class-Brush.md).

## Thermal Erosion Controls

The built-in soil types change only the resting angle. They are identical in every other way.

| **Property** | **Function** |
| --- | --- |
| **Physical Material Presets** | Choose a soil type to use realistic resting angles. Refer to Physical Material Presets above for more information. |
| **Iterations** | The number of times the erosion is applied for each brush stroke. A low values gives each iteration a small impact. |
| **Time Interval (dt)** | How much time passes during each iteration of the simulation. A low value gives each iteration a small impact. |
| **Resting Angle** | The maximum slope angle for the material. Change this value to create new soil types. |
| **Resting Angle Jiiter** | The amount of randomness in the resting angle applied with each brush stroke. A higher value can give a more realistic result, because materials don’t slide evenly. For example, plant roots can hold soil at a steeper angle than the material’s resting angle. |

## Additional resources

* [Considerations for working with the Erosion tool](terrain-Erosion-Considerations.md)
* [Hydraulic erosion](terrain-Erosion-Hydraulic.md)
* [Wind erosion](terrain-Erosion-Wind.md)