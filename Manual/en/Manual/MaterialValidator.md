# Material Validator window reference for the Built-In Render Pipeline

## Validate Albedo mode

![The PBR Validation Settings when in Validate Albedo mode, which appear in the scene view](../uploads/Main/materialValidator3.png)

The PBR Validation Settings when in Validate Albedo mode, which appear in the scene view

The PBR Validation Settings that appear in the **Scene** view when you set Material Validation to Validate Albedo.

| **Property** | **Description** |
| --- | --- |
| **Check Pure Metals** | Enable this checkbox if you want the Material Validator to highlight in yellow any **pixels** it finds which Unity defines as metallic, but which have a non-zero albedo value. See [Pure Metals](material-validator-validate.md#pureMetals) for more details. By default, this is not enabled. |
| **Luminance Validation** | Use the drop-down to select a preset configuration for the Material Validator. If you select any option other than **Default Luminance**, you can also adjust the Hue Tolerance and Saturation Tolerance. The color bar underneath the name of the property indicates the albedo color of the configuration. The Luminance value underneath the drop-down indicates the minimum and maximum luminance value. The Material Validator highlights any pixels with a luminance value outside of these values. This is set to **Default Luminance** by default. |
| **Hue Tolerance** | When checking the albedo color values of your material, this slider allows you to control the amount of error allowed between the hue of your material, and the hue in the validation configuration. |
| **Saturation Tolerance** | When checking the albedo color values of your material, this slider allows you to control the amount of error allowed between the saturation of your material, and the saturation in the validation configuration. |
| **Color Legend** | These colors correspond to the colours that the Material Validator displays in the Scene view when the pixels for that Material are outside the defined values.  This property has the following options:  * **(Red) Below Minimum Luminance Value**: The Material Validator highlights in red any pixels which are below the minimum luminance value defined in **Luminance Validation** (meaning that they are too dark). * **(Blue) Above Maximum Luminance Value**: The Material Validator highlights in blue any pixels which are above the maximum luminance value defined in **Luminance Validation** (meaning that they are too bright). * **(Yellow) Not A Pure Metal**: If you have Check Pure Metals enabled, the Material Validator highlights in yellow any pixels which Unity defines as metallic, but which have a non-zero albedo value. |

## Validate Metal Specular mode

![The PBR validations settings, when in Validate Metal Specular mode](../uploads/Main/materialValidator6.png)

The PBR validations settings, when in Validate Metal Specular mode

The PBR Validation Settings that appear in the **Scene view** when you set **Material Validation** to **Validate Metal Specular**.

| **Property** | **Description** |
| --- | --- |
| **Check Pure Metals** | Enable this checkbox if you want the Material Validator to highlight in yellow any pixels it finds which Unity defines as metallic, but which have a non-zero albedo value. See Pure Metals, below, for more details. By default, this is not enabled. |
| **Color Legend** | These colors correspond to the colours that the Material Validator displays in the Scene view when the pixels for that Material are invalid - meaning their specular value falls outside the valid range for that type of material (metallic or non-metallic).   This color legend has the following valid ranges:  * **(Red) Below Minimum Luminance Value**: 40 for non-metallic, or 155 for metallic. * **(Blue) Above Maximum Luminance Value**: 75 for non-metallic, or 255 for metallic. * **(Yellow) Not A Pure Metal**: If you have **Check Pure Metals** enabled, the Material Validator highlights in yellow any pixels which Unity defines as metallic, but which have a non-zero albedo value. |

## Additional resources

* [Material Validator](material-validator-introduction.md)