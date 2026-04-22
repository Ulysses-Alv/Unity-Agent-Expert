# Noise reference

**Note:** This tool is only available in the [Terrain Tools package](https://docs.unity3d.com/Packages/com.unity.terrain-tools@latest/).

The noise feature, both as part of the Noise height sculpting tool and as a mask on other brush-based tools, always uses the same property set.

## Noise Field Preview

The Noise Field Preview is interactive. It both shows and changes the noise output.

### Preview colors

The Noise Field Preview uses color coding to display the status of each point in the noise output. The values of a correct output are between 0 (black) and 1 (white); if your values are under `0` or over `1`, you need to adjust your settings.

| **Color** | **Output value** |
| --- | --- |
| **Cyan** | Under `0`: change settings to correct this. |
| **Black**, **Gray**, or **White** | From `0` to `1`. |
| **Red** | Over `1`: change settings to correct this. |

### Change the output

Zooming in and out, or panning the texture, changes the noise output:

* Translation and Panning: Click on the preview to pan the noise field, and drag the cursor to change the translation of the noise field.
* Scale and Zoom: To change the scale of the noise field, place the cursor on the preview, and scroll to zoom in or out.

You can use the **Transform Settings** fields instead of the preview; they perform the same operations, and include a rotation option.

### Export

To export the noise output as a texture, select **Export**. Refer to [Other way to use noise](terrain-Noise-Other-Uses.md) for more information.

## Coordinate Space

There are two ways to use the noise output:

* **Brush** fits the noise output into the brush. The smaller the brush, the more detail you get with each brush stroke.
* **World** maintains the same output regardless of the brush size.

## Noise Settings Asset

Save your noise settings to reuse or share them. For more information, refer to [Other way to use noise](terrain-Noise-Other-Uses.md).

## Transform Settings

The Unity Editor uses input points to calculate the noise output. The **Transform Settings** change the position, rotation, and scale of these points, which changes the output.

**Tip:** You can also change the noise output by panning and zooming in the Noise Field Preview.

| **Property** | **Description** |
| --- | --- |
| **Translation** | Move the center point of the noise. |
| **Rotation** | Rotate all inputs around the center point of the noise. |
| **Scale** | The size of the field of input points; changing the scale zooms the output in or out. |
| **Flip Scale** | Inverts the noise output.] |

## Domain Settings

The properties of the Domain Settings are calculated in order. For example, **Lacunarity** is calculated after **Frequency**.

| **Property** |  | **Description** |
| --- | --- | --- |
| **Noise Type** |  | For details of noise types, refer to [Noise and fractal types](terrain-Noise-Types.md). |
| **Fractal Type** |  | For details of fractal types, refer to [Noise and fractal types](terrain-Noise-Types.md) . |
| **Strata** |  | These options are only available if you select **Strata** in **Fractal Type**. They default to 0, which gives a plain field. To generate detailed noise, change the values. |
|  | **Offset** | Move each point in the output higher (positive values) of lower (negative values). |
|  | **Scale** | The size of the field of input points; changing the scale zooms the output in or out. |
| **Octaves** |  | The number of noise layers. Each octave has a higher frequency and a lower amplitude than the previous one, and applies on top of all previous octaves. More octaves generate a more detailed noise. |
| **Amplitude** |  | For each octave: the height of the peaks and the depth of the valleys in the generated noise. A higher amplitude gives more variation, and a lower amplitude gives a smoother noise. The values of **Amplitude** and **Persistence** are multiplied for each octave. |
| **Persistence** |  | For each octave: how much it impacts the noise output. A higher value gives each octave a greater impact on the final pattern, which creates a more detailed noise. A lower value gives a smoother noise, because the higher octaves have less influence. The values of **Amplitude** and **Persistence** are multiplied for each octave. |
| **Frequency** |  | How often a noise pattern repeats. This determines the scale and **level of detail** of the generated noise. A higher frequency gives a more tightly packed and smaller-scaled pattern, while a lower frequency leads to a larger-scaled and more spread-out pattern. The values of **Frequency** and **Lacunarity** are multiplied for each octave. |
| **Lacunarity** |  | Adds a layer, with a pattern that’s a multiplication of the output after **Frequency** is calculated. A value between `-1` and `1` decreases the scale of the noise at each octave while a value under `-1` or over `1` increases the scale of the pattern. A value of `1` has no impact. |

### Warp Settings

Warp Settings apply turbulence to the noise.

| **Property** | **Description** |
| --- | --- |
| **Iterations** | How many turbulence runs to perform. Higher values creates a more detailed noise. The value range is from 0 to 8. 0 applies no noise, even if the **Strength**\*\* value is greater than 0. |
| **Strength** | How much impact each turbulence iteration has. The value range is from –2 to 2. 0 applies no warp, even if the **Iteration** value is greater than 0. The higher and lower values create more movement, in opposite directions to each other. |
| **Offset** | Move the output in 4D to move through variations without changing the **Iteration** and **Strength** properties. The **W** value is time. |

The image below display the noise output for a value noise with and without warp.

![Value noise without warp](../uploads/Main/Noise_Value_no_Warp.png)

Value noise without warp

![Value noise with warp](../uploads/Main/Noise_Value_With_Warp.png)

Value noise with warp

## Additional resources

* [Other way to use noise](terrain-Noise-Other-Uses.md)
* [Noise and fractal types](terrain-Noise-Types.md)
* [Sculpt height with the noise tool](terrain-Sculpt-Noise-Landing.md)