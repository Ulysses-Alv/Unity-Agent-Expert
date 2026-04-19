# Set language direction

You can use the `language-direction` ([`LanguageDirection`](../../ScriptReference/UIElements.LanguageDirection.md)) attribute to set the text direction for **visual elements**, specifically whether it flows from left to right (LTR) or right to left (RTL). This is important for languages such as Arabic and Hebrew, which users read from right to left. This attribute corresponds to the [dir](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/dir) property in HTML and cascades to child elements.

Language direction can impact the text layout behaviors, such as the position of ellipses and punctuation.

Language direction supports the following values:

* Inherited (default): The element inherits the text direction from its parent.
* LTR (Left-to-Right): Forces the text within the element to flow from left to right.
* RTL (Right-to-Left): Forces the text within the element to flow from right to left. Selecting this option can only get basic RTL support like text reversal. To get more comprehensive RTL support, such as line breaking, word wrapping, or text shaping, enable [Advanced Text Generator](../UIE-advanced-text-generator.md).

You can set the language direction in UI Builder, UXML, or C# **scripts**.

## In UI Builder

1. Select the visual element you want to modify.
2. In the ****Inspector**** panel, expand **Attributes**.
3. From the **Language Direction** dropdown, select the desired direction.

## In UXML

To set the language direction in UXML, use the `language-direction` attribute. For example:

```
<Label text="Hello" language-direction="RTL" />
```

## In C# scripts

To set the language direction in C# scripts, use the [`LanguageDirection`](../../ScriptReference/UIElements.LanguageDirection.md) property. For example:

```
new TextElement() {languageDirection = LanguageDirection.RTL};
```

## Additional resources

* [Advanced Text Generator](../UIE-advanced-text-generator.md)
* [Get started with text](../UIE-get-started-with-text.md)
* [Font assets](../UIE-font-asset-landing.md)
* [Text effects](../UIE-text-effects.md)