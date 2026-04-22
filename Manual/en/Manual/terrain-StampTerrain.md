# Stamp a shape onto the terrain

Use **Stamp **Terrain**** to stamp a shape onto the terrain. You can use either:

* A brush, either built-in or custom. You can set the brush’s height to create a full-sized feature in one click, rather than build up the height from multiple brush uses (like the Raise or Lower Terrain tool).
* A **mesh**, such as an imported 3D model. This option is available only in the Terrain Tools package.

## Access the tool

To use a stamp, in the Terrain tile’s **Inspector** window, select the **Paint Terrain** tool > tool dropdown > **Stamp Terrain**.

## Tool options

The Stamp Terrain tool is brush-based. To learn about working with brushes, refer to [Brushes](class-Brush.md).

## Stamp Controls in the built-in Terrain system

| **Property** | **Function** |
| --- | --- |
| **Stamp Height** | How much to increase the height when using the stamp. The range is from `0` to the Terrain tile’s maximum height (refer to **Terrain Height** in the [Terrain Settings](terrain-OtherSettings.md)). Note that **Stamp Height** is multiplied by the **Opacity** property. For example, if **Stamp Height** is `200` and **Opacity** is `0.5`, the stamp’s final height is `100`. |
| **Subtract** | Remove the stamp from the terrain. Note that **Subtract** works only if the **Max – Add** property is greater than `0`, and can’t dig below the Terrain tile’s **GameObject** y-axis coordinate. |
| **Max - Add** | If the value is `0`, the final height is the higher between the stamp and the existing terrain. If the value is `1`, the final height is the combined heights of the stamp and terrain. |

## Stamp Controls in the Terrain Tools package

### General Stamp Settings

| **Property** |  | **Function** |
| --- | --- | --- |
| **Stamp Mode** |  | Stamp a brush or a mesh. |
| **Behaviour** |  | Which parts of the Terrain to stamp. |
|  | **Min** | Applies the stamp only where it lowers the Terrain. Areas under the stamp that are lower than the **Stamp Height** value remain unchanged. |
|  | **Set** | Applies the stamp without referencing the height of the underlying Terrain. |
|  | **Max** | Applies the stamp only where it raises the Terrain. Areas under the stamp that are higher than the **Stamp Height** value remain unchanged. |
| **Stamp Height** |  | The range is between the negative and positive values of the Terrain tile’s maximum height (refer to **Terrain Height** in the [Terrain Settings](terrain-OtherSettings.md)). Use negative values to subtract the stamp from the height, and positive values to add to the height. Note that **Stamp Height** is multiplied by the **Brush Strength** property. For example, if **Stamp Height** is `200` and **Brush Strength** is `0.5`, the stamp’s final height is `100`. |
| **Blend Amount** |  | How much to override existing Terrain features when using the stamp. The default value is `0`, which fully overrides existing terrain features. To keep those features, set the value to `1`; this makes the stamp shape indistinct. **Tip:** Use a solid smooth brush, such as the built-in circular brush, to raise the terrain without changing its details. For example, you can raise a hill without losing contours or details you created. Note that the higher the value of **Stamp Height**, the more distinct the stamp is, even for high values of **Blend Amount**. |

### Mesh Stamp Settings

Only available in the Terrain Tools package, when the **Stamp Mode** is **Mesh**.

| **Property** | **Function** |
| --- | --- |
| **Mesh** | The mesh to use as a stamp. |
| **Scale** | Make the mesh larger or smaller along any of its axes. |
| **Rotation** | Rotate the mesh around any of its axes. |

## Additional resources

* [Primary sculpting tools](terrain-Basic-Sculpting.md)
* [Advanced sculpting tools](terrain-Sculpt.md)
* [Change existing height with effects](terrain-Effects.md)
* [Transform terrain features](terrain-Transform.md)
* [Erode the terrain](terrain-Erosion.md)