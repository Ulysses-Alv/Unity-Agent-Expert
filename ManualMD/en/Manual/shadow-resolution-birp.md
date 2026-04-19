# Configure shadow resolution in the Built-In Render Pipeline

To calculate the resolution of a shadow map, Unity:

1. Determines the area of the screen view that the Light can illuminate. For directional lights, the whole screen can be illuminated. For Spot Lights and Point Lights, the area is the onscreen projection of the shape of the light’s extent: a sphere for point lights, or a cone for Spot Lights. The projected shape has a width and height in pixels on the screen; the larger of those two values is then taken. This value is called the Light’s **pixel** size.
2. Determines the shadow quality multiplier. Unity uses the **Shadow Resolution** setting for this, which is set in the [Quality Settings](class-QualitySettings.md) window). The quality settings correspond to the following values:
   * Very High: 1.0
   * High: 0.5
   * Medium: 0.25
   * Low: 0.125
3. Performs the following calculation, and then clamps the result to the maximum size:

| **Light type:** | **Formula:** | **Maximum resolution, in pixels:** |
| --- | --- | --- |
| Directional | [NextPowerOfTwo](../ScriptReference/Mathf.NextPowerOfTwo.md) (pixel size \* shadow quality multiplier \* 3.8) | 4096 x 4096 when Shadow Resolution is Very High Quality and/or if the GPU has 512MB or more of RAM, 2048 x 2048 otherwise. |
| Spot Lights | [NextPowerOfTwo](../ScriptReference/Mathf.NextPowerOfTwo.md) (pixel size \* shadow quality multiplier \* 2.0) | 2048 x 2048 if the GPU has 512MB or more of RAM, 1024 x 1024 otherwise. |
| Point Lights | [NextPowerOfTwo](../ScriptReference/Mathf.NextPowerOfTwo.md) (pixel size \* shadow quality multiplier \* 1.0) | 1024 x 1024 if the GPU has 512MB or more of RAM, 512 x 512 otherwise. |

## Override shadow map resolution

In the Built-in **Render Pipeline**, you can set the resolution of a Light’s shadow map by setting the [Light.shadowCustomResolution](../ScriptReference/Light-shadowCustomResolution.md) property to a value greater than 0. When this value is greater than 0, Unity performs the following calculation for all Light types:

[NextPowerOfTwo](../ScriptReference/Mathf.NextPowerOfTwo.md) (shadowCustomResolution)

It then clamps the maximum resolution based on Light type and hardware, as shown in the table above.

## Additional resources

* [Shadows](Shadows.md)