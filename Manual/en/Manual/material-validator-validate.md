# Validate and correct materials in the Built-In Render Pipeline

## Validate albedo

Unity’s [Material charts](StandardShaderMaterialCharts.md) define the standard luminance range as 50–243 sRGB for non-metals, and 186–255 sRGB for metals. **Validate Albedo** mode colors any **pixels** outside of these ranges with different colors to indicate that the value is too low or too high.

In the example below, the first texture is below the minimum luminance value, and therefore too dark. The fourth texture is above the maximum luminance value, and therefore too bright.

![A Scene (without the Material Validator enabled) in which the first and fourth Materials have incorrect albedo values](../uploads/Main/materialValidator4.jpg)

A Scene (without the Material Validator enabled) in which the first and fourth Materials have incorrect albedo values

![The same Scene with the Material Validator enabled and set to Validate Albedo. The texture that is below the minimum luminance value is red. The texture that is above the minimum luminance value is blue](../uploads/Main/materialValidator5.jpg)

The same Scene with the Material Validator enabled and set to Validate Albedo. The texture that is below the minimum luminance value is red. The texture that is above the minimum luminance value is blue

The [material charts](StandardShaderMaterialCharts.md) provide albedo values for common Materials. The brightness of albedo values has a dramatic impact on the amount of diffuse bounce light generated, so it is important for [Global Illumination](LightingInUnity.md#globalIllumination) baking to make sure that your different Material types are within the correct luminance ranges, in proportion with each other. To help you get these values right, you can select from the presets in the Luminance Validation drop-down, which provides common Material albedo values to verify the luminance ranges of particular Material types.

## Validate specular

Unity’s [Material charts](StandardShaderMaterialCharts.md) define two separate **specular color** ranges:

* **Non-metallic materials**: 40–75 sRGB
* **Metallic materials**: 155 - 255 sRGB

In Unity, all non-metallic Materials have a constant specular color that always falls within the correct range. However, it is common for metallic Materials to have specular values that are too low. To help you identify metallic Materials with this issue, the Material Validator’s **Validate Metal Specular** mode colors all pixels that have a specular color value that is too low. This includes all non-metallic materials by definition.

In the example below, the left material is below the minimum specular value, and therefore too dark. This also applies to the **Scene**’s background. The right material has specular values with in the valid range.

![A Scene with two metallic Materials. The left has incorrect metallic specular values](../uploads/Main/materialValidator7.jpg)

A Scene with two metallic Materials. The left has incorrect metallic specular values

![The same Scene with the Material Validator enabled and set to Validate Metal Specular](../uploads/Main/materialValidator8.png)

The same Scene with the Material Validator enabled and set to **Validate Metal Specular**

## Validate pure metals

Unity defines physically-based shading materials with a specular color greater than 155 sRGB as metallic. For Unity to define a metallic Material as a **pure metal**

If a non-metallic surface has a specular color value that is too high, but has a non-zero albedo value, this is often due to an authoring error. The Material Validator also has an option called **Check Pure Metals**. When you enable this option, the Material Validator colors in yellow any Material that Unity defines as metallic but which has a non-zero albedo value. An example of this can be seen in the images below. It shows three materials, the left and right materials are pure metals, but the middle material is not, so the Material Validator colors it yellow:

![A Scene with three metallic Materials. The middle Material is not a pure metal (it has a non-zero albedo value)](../uploads/Main/materialValidator9.jpg)

A Scene with three metallic Materials. The middle Material is not a pure metal (it has a non-zero albedo value)

![The same Scene with the Material Validator enabled and set to Validate Metal Specular, with Check Pure Metals enabled)](../uploads/Main/materialValidator10.png)

The same Scene with the **Material Validator** enabled and set to **Validate Metal Specular**, with Check Pure Metals enabled)

In the second image above, the background is red because the Materials in the background are below the minimum specular value for the Material Validator’s **Validate Metal Specular** mode.

For complex materials that combine metallic and non-metallic properties, the pure metal checker is likely to pick up some invalid pixels, but if a Material is totally invalid, it’s usually a sign of an authoring error.

## Override the default luminance values

Depending on the art style of your project, you might want the luminance values of Materials to differ from the preset luminance ranges. In this case, you can override the built-in albedo values used by the Material Validator with your own values. To override the preset luminance ranges, assign an array of [AlbedoSwatchInfo](../ScriptReference/Rendering.AlbedoSwatchInfo.md) values for each desired Material type to the property [EditorGraphicsSettings.albedoSwatches](../ScriptReference/Rendering.EditorGraphicsSettings-albedoSwatches.md).