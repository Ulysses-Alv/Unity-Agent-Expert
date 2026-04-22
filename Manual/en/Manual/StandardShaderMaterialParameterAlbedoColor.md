# Set the color of a material in the Standard Shader

The **Albedo** property controls the base color and transparency of the material.

## Set a single color

To use a single color across a material, do either of the following:

* Select the swatch to open the **Color** window.
* Select the dropper to select a new color from anywhere on your screen.

In the **Color** window, use the **Alpha** (**A**) slider to set transparency. Refer to [Transparency](StandardShaderTransparency.md) for more information.

Refer to [Color](graphics-color.md) for more information about color.

## Set a texture

To assign a texture to a material, drag the texture asset from the **Project window** to the **Albedo** property.

Unity uses the swatch color to tint the texture.

The alpha channel of the texture controls the transparency of the material. Refer to [Transparency](StandardShaderTransparency.md) for more information.

Refer to [Textures](Textures.md) for more information about textures.

## Additional resources

* [Material.Color](../ScriptReference/Material-color.md)