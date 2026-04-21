# Style text with USS

You can style text with USS text properties inline in [UXML](UIE-UXML.md), a [USS](UIE-USS.md) file, or directly in [UI Builder](UIBuilder.md). To learn more about USS text properties, refer to [Text properties](UIE-USS-SupportedProperties.md#unity-text).

## Style text in USS and UXML

[Text properties](UIE-USS-SupportedProperties.md#unity-text) are regular USS style properties. You can set text style properties on any element. Unlike most USS style properties, text style properties propagate to child elements.

The following USS example styles the `Label` text to bold and italic, and has a font size of `39px`:

```
Label {
    -unity-font-style: bold-and-italic; 
    font-size: 39px;
}
```

The following UXML inline style example applies the same style to the `Label` text:

```
<ui:UXML xmlns:ui="UnityEngine.UIElements" xmlns:uie="UnityEditor.UIElements">
    <ui:VisualElement>
        <ui:Label text="Label" style="-unity-font-style: bold-and-italic; font-size: 39px;" />
    </ui:VisualElement>
</ui:UXML>
```

## Style text in UI builder

To style text in UI Builder, you can use the **Text** section in a UI control’s **Inspector** window to style text.

If the UI control is a text element that inherits from [TextElement](UIE-uxml-element-TextElement.md), such as [Label](UIE-uxml-element-Label.md) or [Button](UIE-uxml-element-Button.md), you can also set the following text styles directly in the **Canvas** on selected text elements:

* Horizontal text align
* Vertical text align
* Text wrap

![Text styles are exposed as toggles in the Canvas on selected elements](../uploads/Main/UIBuilder/CanvasTextToggles.png)

Text styles are exposed as toggles in the Canvas on selected elements

## Additional resources

* [Get started with text](UIE-get-started-with-text.md)
* [Text effects](UIE-text-effects.md)
* [Font assets](UIE-font-asset.md)
* [Rich text tags](UIE-rich-text-tags.md)
* [Style sheet assets](UIE-style-sheet.md)
* [Include sprites in text](UIE-sprite.md)