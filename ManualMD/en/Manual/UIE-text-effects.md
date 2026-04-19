# Text effects

You can apply text effects to enhance the visual appearance of the text. However, these effects are only supported for font assets with [SDF rendering mode](UIE-font-asset.md#sdf-fonts) and not for bitmap. The range of the text effects is limited by the padding defined for the font asset. Therefore, to increase the size of the effect, you must increase the padding.

You can use text effects to:

* Apply text shadows
* Apply text outlines

## Apply text shadows

To apply shadows to the text, set the `text-shadow` property in a USS file, inline in UXML, or a C# script.

`text-shadow` is a shorthand property that sets the following properties:

* `text-shadow-offset-x`: Horizontal shadow displacement. Positive values move the shadow to the right, while negative values move it to the left.
* `text-shadow-offset-y`: Vertical shadow displacement. Positive values move the shadow down, while negative values move it up.
* `text-shadow-blur-radius`: The blur intensity of the shadow. Higher values result in more blur shadows, while 0 creates a sharp shadow.
* `text-shadow-color`: The color of the shadow, either as a hex code or in RGBA format.

The following example applies a shadow to a `Label` element:

![Text shadow example](../uploads/Main/uitk/text-shadow.png)

Text shadow example

```
Label {
    text-shadow: 6.1px -2.4px 0px rgb(144, 31, 32);
}
```

## Apply text outlines

To apply outlines to text, set the `unity-text-outline` property in a USS file, inline in UXML, or a C# script.

`unity-text-outline` is a shorthand property that sets the following properties:

* `-unity-text-outline-width`: The width of the outline.
* `-unity-text-outline-color`: The color of the outline, either as a hex code or in RGBA format.

The following example applies an outline to a `Label` element:

![Text outline example](../uploads/Main/uitk/text-outline.png)

Text outline example

```
Label {
    text-outline: 6px rgb(144, 31, 32);
}
```

## Additional resources

* [Get started with text](UIE-get-started-with-text.md)
* [Font assets](UIE-font-asset.md)
* [Text properties](UIE-USS-SupportedProperties.md#unity-text)
* [Style text with USS](UIB-styling-ui-text.md)