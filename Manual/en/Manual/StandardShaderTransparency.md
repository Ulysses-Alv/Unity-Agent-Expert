# Make a material transparent

To make a material transparent if it uses the Standard **Shader**, follow these steps in the ****Inspector**** window of the material:

1. Set **Rendering Mode** to **Transparent** or **Fade**.
2. Select the **Albedo** swatch to open the **Color** window.
3. Set the **Alpha** (**A**) slider. A value of 0 means fully transparent, and a value of 1 means fully opaque.

![A range of alpha values from 0 to 1](../uploads/Main/StandardShaderTransparencyGraduationTable.jpg)

A range of alpha values from 0 to 1

## Use a texture to control transparency

If you use a texture for **Albedo** instead of a color, by default the alpha channel of the texture controls the transparency of the material. A value of 0 in the alpha channel means fully transparent, and a value of 1 means fully opaque. You can use different values to make different areas more or less transparent.

Unity combines the alpha channel of the texture with the alpha value you set in the **Color** window. For example, if you set the alpha value in the **Color** window to 0.1, opaque texture **pixels** become almost fully transparent.

To check the alpha channel of a texture, follow these steps:

1. Select the texture in the **Project window**.
2. Go to the texture preview section and select the **Alpha** (**A**) button. The preview displays black for 0 and white for 1, with shades of gray between.

![An imported texture. On the left, RGB is selected and the texture preview section displays all the texture channels. On the right, A is selected and the texture preview section displays the alpha channel only.](../uploads/Main/StandardShaderTransparencyMapRGBAlphaToggle.png)

An imported texture. On the left, **RGB** is selected and the texture preview section displays all the texture channels. On the right, **A** is selected and the texture preview section displays the alpha channel only.

![The window uses a single texture to create a fully opaque window frame, partially transparent windows, and fully transparent gaps.](../uploads/Main/StandardShaderTransparencyMapBrokenWindow.jpg)

The window uses a single texture to create a fully opaque window frame, partially transparent windows, and fully transparent gaps.

Refer to [Default Import Settings reference](texture-type-default.md) for more information about texture transparency settings.

## Additional resources

* [Material.Color](../ScriptReference/Material-color.md)