# Work with text

Text objects are defined by the following attributes of a UI control:

* The `text` attribute of some UI controls, such as [Label](UIE-uxml-element-Label.md) or [TextElement](UIE-uxml-element-TextElement.md), that sets the display text.
* The `value` attribute of the [TextField](UIE-uxml-element-TextField.md) that accepts the input text, and the `label` attribute that sets the display text for [TextField](UIE-uxml-element-TextField.md).

You can use [USS](UIE-USS.md) [text properties](UIE-USS-SupportedProperties.md#unity-text) to style text, such as set the font size and color, etc.

You can also add a new font to style text. Convert fonts to font assets before you use them in your project. In addition to USS styling, you can use rich text tags to style certain words in a text string.

| **Topic** | **Description** |
| --- | --- |
| [Get started with text](UIE-get-started-with-text.md) | Learn how to style text, create fonts, and style with rich text tags and style sheets by examples. |
| [Advanced Text Generator](UIE-advanced-text-generator.md) | Add comprehensive Unicode support and text shaping capabilities to your project. |
| [Style text with USS](UIB-styling-ui-text.md) | Style text with USS text properties inline in UXML, a USS file, or directly in UI Builder. |
| [Style text with rich text tags](UIE-rich-text-tags.md) | Style words between tags in a text string. |
| [Font assets](UIE-font-asset-landing.md) | Understand different font assets and all their properties. |
| [Text effects](UIE-text-effects.md) | Apply text effects to text elements to enhance the visual appearance of the text. |
| [Style sheet assets](UIE-style-sheet.md) | Create custom text styles to extend the rich text tags. |
| [Include sprites in text](UIE-sprite.md) | Create **sprite** assets to interpret emoji characters and include them in text. |
| [Color gradients](UIE-color-gradient.md) | Create color gradients to apply up to four colors for each character in a text string. |
| [Color emojis](UIE-color-emojis.md) | Include color glyphs and emojis in text. |
| [Language direction](ui-systems/language-direction.md) | Set the text directionality of a text element to support right-to-left (RTL) languages. |
| [UITK Text Settings assets](UIE-text-setting-asset.md) | Referenced by a Panel Settings asset and controls the default values for all text objects used within that Panel. |
| [Fallback font](UIE-fallback-font.md) | Add fallback font for missing character in a font asset. |
| [Create custom text animation](ui-systems/create-custom-text-animation.md) | Use the `TextElement.PostProcessTextVertices` API to create custom text animation |

## Additional resources

* [MeshGenerationContext](../ScriptReference/UIElements.MeshGenerationContext.md)
* [UI Renderer](UIE-ui-renderer.md)
* 📖 **E-Book**: [UI Toolkit for advanced Unity developers - Graphic and font assets preparation](best-practice-guides/ui-toolkit-for-advanced-unity-developers/graphic-and-font-assets-preparation.md)
* 📖 **E-Book**: [UI Toolkit for advanced Unity developers - Text](best-practice-guides/ui-toolkit-for-advanced-unity-developers/text.md)