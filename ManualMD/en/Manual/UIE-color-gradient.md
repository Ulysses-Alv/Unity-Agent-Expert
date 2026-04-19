# Color gradients

You can apply gradients of up to four colors to text. When you add a color gradient, TextCore applies gradient colors to each character in a text string.

Use the [`<gradient>` rich text tags](UIE-supported-tags.md#gradient) to apply color gradient presets to text objects. A gradient preset overrides the text’s local gradient type and colors.

## Create a color gradient preset

To create a color gradient with one or more colors, and place it in a specific folder, follow these steps:

1. From the menu, select **Assets** > **create** > **Text Core** > **Color Gradient** to add a new color gradient preset.
2. In the Color Gradient’s **Inspector** window, select a color mode from the dropdown menu:

   * **[Single](#single-color)**: Applies a single color.
   * **[Horizontal Gradient](#horizontal-gradients)**: Applies two colors and produces a side-to-side transition between them on each character.
   * **[Vertical Gradient](#vertical-gradients)**: Applies two colors and produces an up-and-down transition between the two on each character.
   * **[Four Corners Gradient](#four-corner-gradients):** Applies four colors. Each color radiates from its assigned corner of each character.
3. Set the gradient colors. The number of colors available in the **Colors** field depends on the gradient mode you choose. Each swatch corresponds to the color’s origin on a character **sprite**.
4. Place the color gradient asset into the path set for the **Color Gradient Presets** in the [Panel Text Setting asset](UIE-text-setting-asset.md#color-gradients-presets).

## Apply color gradient presets

To apply a color gradient preset to text, reference the color gradient by name as `<gradient="name-of-the-color-gradient">Your text</gradient>`. If you apply the color gradient like this, the color of the text is multiplied by the object’s current vertex colors.

To apply the pure gradient to a selection of text, use a `<color>` tag to reset the color to white. For example:

```
<color=white><gradient="Light to Dark Green - Vertical">Light to Dark Green gradient</gradient></color>
```

## Color gradient mode examples

The following shows color gradient examples for each mode.

### Single color

A single-color gradient.

![Text with each character in yellow.](../uploads/Main/font/ColorGradient_Single-Y_half.png)

Text with each character in yellow.

### Horizontal gradients

A side-to-side gradient with two colors.

![Text with color transition from yellow to black side-to-side for each character.](../uploads/Main/font/ColorGradient_Horiz-YB_half.png)

Text with color transition from yellow to black side-to-side for each character.

![Text with color transition from black to yellow side-to-side for each character.](../uploads/Main/font/ColorGradient_Horiz-BY_half.png)

Text with color transition from black to yellow side-to-side for each character.

### Vertical gradients

An up-and-down gradient with two colors.

![Text with color transition from black to yellow up-and-down for each character.](../uploads/Main/font/ColorGradient_Vert-BY_half.png)

Text with color transition from black to yellow up-and-down for each character.

![Text with color transition from yellow to black up-and-down for each character.](../uploads/Main/font/ColorGradient_Vert-YB_half.png)

Text with color transition from yellow to black up-and-down for each character.

### Four corner gradients

A gradient with four colors. Each color radiates from one corner. This is the most versatile gradient type. You can vary some colors and keep others identical to create different kinds of gradients.

![Text with four-color gradient for each character with yellow on the up-left, black on the up-right, red on the down-left, and green on the down-right. ](../uploads/Main/font/ColorGradient_4-Corner-YBRG_half.png)

Text with four-color gradient for each character with yellow on the up-left, black on the up-right, red on the down-left, and green on the down-right.

This example shows three corners with one color and the fourth with a different color.

![Text with four-color gradient for each character with black on the up-left and yellow for the other corners.](../uploads/Main/font/ColorGradient_1-Corner-BYYY_half.png)

Text with four-color gradient for each character with black on the up-left and yellow for the other corners.

This example shows pairs of adjacent corners with the same color to create horizontal or vertical gradients.

![Text with four-color gradient for each character with black on the left up and down, and yellow on the right up and down.](../uploads/Main/font/ColorGradient_2-Corner-BYBY_half.png)

Text with four-color gradient for each character with black on the left up and down, and yellow on the right up and down.

![Text with four-color gradient for each character with black on the up left and right, and yellow on the down left and right.](../uploads/Main/font/ColorGradient_2-Corner-BBYY_half.png)

Text with four-color gradient for each character with black on the up left and right, and yellow on the down left and right.

This example shows pairs of diagonally opposite corners the same color to create diagonal gradients.

![Text with four-color gradient for each character with black on the up-left and down-right, and yellow on the up-right and down-left.](../uploads/Main/font/ColorGradient_2-Corner-BYYB_half.png)

Text with four-color gradient for each character with black on the up-left and down-right, and yellow on the up-right and down-left.

![Text with four-color gradient for each character with yellow on the up-left and down-right, and black on the up-right and down-left.](../uploads/Main/font/ColorGradient_2-Corner-YBBY_half.png)

Text with four-color gradient for each character with yellow on the up-left and down-right, and black on the up-right and down-left.

This example creates horizontal and vertical three-color gradients with a dominant color at one end and a transition between two colors at the other.

![Text with four-color gradient for each character with yellow on the left up and down, red on the up-right, and black on the down-right.](../uploads/Main/font/ColorGradient_3-Corner-YRYB_half.png)

Text with four-color gradient for each character with yellow on the left up and down, red on the up-right, and black on the down-right.

![Text with four-color gradient for each character with yellow on the up left and right, red on the down-left, and black on the down-right.](../uploads/Main/font/ColorGradient_3-Corner-YYRB_half.png)

Text with four-color gradient for each character with yellow on the up left and right, red on the down-left, and black on the down-right.

This example gives two diagonally opposite corners the same color and the other two corners different colors to create a diagonal stripe three-color gradient.

![Text with four-color gradient for each character with red on the up-left, yellow on the up-right and down-left, and black on the down-right.](../uploads/Main/font/ColorGradient_3-Corner-RYYB_half.png)

Text with four-color gradient for each character with red on the up-left, yellow on the up-right and down-left, and black on the down-right.

![Text with four-color gradient for each character with yellow on the up-left and down-right, black on the up-right, and red on the down-left.](../uploads/Main/font/ColorGradient_3-Corner-YBRY_half.png)

Text with four-color gradient for each character with yellow on the up-left and down-right, black on the up-right, and red on the down-left.

## Additional resources

* [Rich text tags](UIE-rich-text-tags.md)
* [Supported rich text tags](UIE-supported-tags.md)
* [Style sheets assets](UIE-style-sheet.md)
* [Sprites assets](UIE-sprite.md)