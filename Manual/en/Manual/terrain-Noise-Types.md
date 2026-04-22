# Noise and fractal types

**Note:** This tool is only available in the [Terrain Tools package](https://docs.unity3d.com/Packages/com.unity.terrain-tools@latest/).

The Noise height sculpting tool and the Noise mask combine noise and fractal types to create a structured output. The noise type provides the structure’s general shape; the fractal type provides the structure’s details.

## Noise types

| Label in image | Type | Description | Usage example |
| --- | --- | --- | --- |
| A | **Value** | Uses a grid of points with random values. | Fine mountain details. |
| B | **Voronoi** | Breaks up a given area into cells to create patterns with a clear geometry and distinct edges. | Parched fields or tiles. You can add Perlin noise to add details within the larger Voroni pattern. For example, Voronoi noise can create the segmented appearance of a parched field, and Perlin noise can give each segment its own character. |
| C | **Perlin** | Gradient noise: Uses a grid of points with random values to simulate smoothly varying, continuous patterns. | Adds small details to an unbroken surface. |
| D | **Billow** | Similar to Perlin noise, but with more rounded features. | Treetops and hills. |
| E | **Ridge** | Similar to Perlin noise, but with sharper features. | Ridges and deep valleys. |

This image shows a mountain range sculpted with each noise type, with their default settings. The brush is the built-in circular brush at full strength. The noise type labels match the labels in the table above:

![Noise types side by side. A: Value. B: Voronoi. C: Perlin. D: Billow. E: Ridge.](../uploads/Main/Noise_Types_Side_by_Side.png)

Noise types side by side. A: Value. B: Voronoi. C: Perlin. D: Billow. E: Ridge.

## Fractal types

* Fractal Brownian Motion (**FBM**): For mountainous shapes.
* **Strata**: For natural banding like marble, tree rings, and terraces.
* **None**: Create the output only for the noise type. None of the fractal settings are available.

The images show the three fractal types applied to the Billow noise, with no other differences in settings:

![Billow noise with FBM fractal](../uploads/Main/Noise_Billow_FBM_Overrides.png)

Billow noise with FBM fractal

![Billow noise with Strata fractal](../uploads/Main/Noise_Billow_Strata_Overrides.png)

Billow noise with Strata fractal

![Billow noise with no fractal](../uploads/Main/Noise_Billow_No_Fratcal.png)

Billow noise with no fractal

## Additional resources

* [Sculpt height with the Noise tool](terrain-Sculpt-Noise-Tool.md)
* [Other ways to use noise](terrain-Noise-Other-Uses.md)
* [Noise reference](terrain-Noise-Ref.md)