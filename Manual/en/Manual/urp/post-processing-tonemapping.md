# Tonemapping Volume Override reference for URP

Tonemapping is the process of remapping the **HDR** values of an image to a new range of values. Its most common purpose is to make an image with a low dynamic range appear to have a higher range. For more information, refer to the [Wikipedia article on Tone mapping](https://en.wikipedia.org/wiki/Tone_mapping).

## Properties

| **Property** | **Description** |
| --- | --- |
| **Mode** | Select a tonemapping algorithm to use for color grading. The options are:  * **None**: Use this option if you do not want to apply tonemapping. * **Neutral**: Use this option if you only want range-remapping with minimal impact on color hue & saturation. It is generally a good starting point for extensive color grading. * **ACES**: Use this option to apply a close approximation of the reference ACES tonemapper, for a more cinematic look. It is more contrasted than Neutral, and has an effect on actual color hue & saturation. If you use this tonemapper, Unity does all the grading operations in the ACES color spaces, for optimal precision and results. **Note**: ACES HDR tonemapping is not supported on Android devices with Adreno 300 series GPU. |

## HDR Output

When [HDR Output](post-processing/hdr-in-urp.md) is enabled, additional properties become available for the **Tonemapping** Volume Override. The available properties vary depending on the **Mode** property.

### Neutral

| Property | Description |
| --- | --- |
| **Neutral HDR Range Reduction Mode** | The curve that the Player uses for tone mapping. The options are:  * **BT2390**: The default. Defined by the [BT.2390](https://www.itu.int/pub/R-REP-BT.2390) broadcasting recommendations. * **Reinhard**: A simple Tone Mapping operator.  This option is only available when you enable [Advanced Properties](urp-asset-additional-settings.md). |
| **Hue Shift Amount** | The value determines the extent to which your content retains its original hue after you apply HDR settings. When this value is 0, the tonemapper attempts to preserve the hue of your content as much as possible by only tonemapping luminance. |
| **Detect Paper White** | Enable this property if you want URP to use the Paper White value that the display communicates to the Unity Engine. In some cases, the value the display communicates may not be accurate. Implement a calibration menu for your application so that users can display your content correctly on displays that communicate inaccurate values. |
| **Paper White** | The Paper White value of the display. If you do not enable **Detect Paper White**, you must specify a value here. |
| **Detect Brightness Limits** | Enable this property if you want URP to use the minimum and maximum nit values that the display communicates. In some cases, the value the display communicates may not be accurate. It is best practice to implement a calibration menu for your application to allow for these situations. |
| **Min Nits** | The minimum brightness value of the display. If you do not enable **Detect Brightness Limits**, you must specify a value here and in **Max Nits**. |
| **Max Nits** | The maximum brightness value of the display. If you do not enable **Detect Brightness Limits**, you must specify a value here and in **Min Nits**. |

#### Misuse of Hue Shift Amount

Creators might author some content with the intention to use **Hue Shift Amount** to produce special effects. In the illustration below, the **Hue Shift Amount** is 0 for Image A and 1 for Image B. The flames image B appear more intense because of the hue shift effect. It is preferable not to author content in this way, because settings optimized for special effects can have undesirable effects on other content in the **scene**.

![Image A: Output when Hue Shift Amount is 0. Image B: Output when Hue Shift Amount is 1.](../../uploads/urp/post-proc/hdr/HDR-Output-HueShift.png)

Image A: Output when Hue Shift Amount is 0. Image B: Output when Hue Shift Amount is 1.

### ACES

This mode has fixed presets to target 1000, 2000, and 4000 nit displays. It is best practice to implement a calibration menu for your application to ensure that the user can select the right preset.

| Property | Description |
| --- | --- |
| **ACES Preset** | The tone mapper preset to use. The options are:  * **ACES 1000 Nits**: The default. This curve targets 1000 nits displays. * **ACES 2000 Nits**: Curve that targets 2000 nits displays. * **ACES 4000 Nits**: Curve that targets 4000 nits displays. |
| **Detect Paper White** | Enable this property if you want URP to use the Paper White value that the display communicates to the Unity Engine. In some cases, the value the display communicates may not be accurate. Implement a calibration menu for your application so that users can display your content correctly on displays that communicate inaccurate values. |
| **Paper White** | The Paper White value of the display. If you do not enable **Detect Paper White**, you must specify a value here. |