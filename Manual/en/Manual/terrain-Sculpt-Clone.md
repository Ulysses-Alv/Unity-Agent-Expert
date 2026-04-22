# Clone a part of the terrain

**Note:** This tool is only available in the [Terrain Tools package](https://docs.unity3d.com/Packages/com.unity.terrain-tools@latest/).

You can clone a sample of your **terrain**’s height, textures, or both.

## Use the tool

To clone a part of the terrain:

1. In the Terrain tile’s **Inspector** window, select the **Paint Terrain** tool > tool dropdown > **Sculpt** > **Clone**.
2. Hold **Ctrl** (macOS: hold **Option**) and click to set the sample area.
3. Click to apply the sample area anywhere else on the terrain.
4. To change the sample, hold **Ctrl** (macOS: hold **Option**) again.

**Note:** If you paint onto the sample area, the sample changes to match the new design.

## Tool options

The Clone tool is brush-based. To learn about working with brushes, refer to [Brushes](class-Brush.md).

You can change brushes or brush behaviours after you sample, and the sample matches the new brush.

## Clone Movement Behaviours

The clone movement behaviour determines how the sample area moves when you paint.

There are four clone movement behaviours:

* **Follow Always**: The sample area follows the brush center when you move the brush, even if you’re not painting at the same time.
* **Follow On Paint**: The sample area only follows the brush center if you’re painting as you move the brush.
* **Snap**: Like **Follow On Paint**, but the sample area return to its starting point when you stop painting.
* **Fixed**: The sample area stays where you set it.

![Use Follow Always to keep the sample and the brush at a fixed distance](../uploads/Main/terrain-Clone-Moving-Stamp.gif)

Use **Follow Always** to keep the sample and the brush at a fixed distance

![Use Snap for clones that always start from the same spot](../uploads/Main/terrain-Clone-Snap-Stamp.gif)

Use **Snap** for clones that always start from the same spot

## Clone controls

| **Property** | **Description** |
| --- | --- |
| **Terrain sources to clone** | Clone the heightmap, the textures, or both. You can change this behaviour after you sample, and the sample adjusts.   **Note:** If you’re cloning textures to a tile that doesn’t have a base layer, the sample tile’s base layer becomes the clone tile’s base layer. For more information about textures and base layers, refer to [Textures and layers](terrain-Textures-Landing.md). |
| **Clone Movement Behaviour** | Refer to the explanation above. You can change this behaviour after you sample, and the sample adjusts. |
| **Height Offset** | If you’re cloning heightmaps, you can use a positive or negative offset to get a clone that’s a different height from the sample. The range is the negative and positive values of the **Terrain Height** setting. The value `0` means there’s no offset, and the clone’s height is the same as the sample’s height. Note that you can’t use this tool to dig below the terrain tile’s minimum height, even if the negative offset value is large enough. For example, if your offset value is `-600`, but the sample is at `200`, the clone will stop at `0`; it won’t go down to `-400`.   You can change this behaviour after you sample, and the sample adjusts. |

## Additional resources

* [Erode the terrain](terrain-Erosion.md)
* [Add noise to the terrain](terrain-Sculpt-Noise-Landing.md)