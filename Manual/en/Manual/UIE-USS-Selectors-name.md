# Name selectors

USS name selectors match elements based on the name of an element. USS Name selectors are analogous to CSS ID selectors that match elements with a specific `id` attribute.

To set the name of an element:

* In C# script, use `VisualElement.name`.
* In UXML, use the `name` attribute. For example: `<VisualElement name="my-nameName">`.

To avoid unexpected matches, make element names unique within a panel.

## Syntax

A name selector consists of an element’s assigned name prefixed with a number sign (`#`).

```
#ElementName { ... }
```

**Note**: Only use the number sign (`#`) when you write the selector in a USS file. Don’t use it when you assign the name to an element in a UXML or C# file. An element name that includes the number sign is invalid. For example `<Button name="#OK" />` is invalid.

## Example

To demonstrate how simple selectors match elements, here is an example UI Document.

```
<UXML xmlns="UnityEngine.UIElements">
  <VisualElement name="container1">
    <VisualElement name="container2" class="yellow">
      <Button name="OK" class="yellow" text="OK" />
      <Button name="Cancel" text="Cancel" />
    </VisualElement>
  </VisualElement>
</UXML>
```

With no styles applied, the UI looks like the following:

![Example buttons with margins and thin blue borders.](../uploads/Main/uss-selectors-nostyle.png)

Example buttons with margins and thin blue borders.

The following name selector style rule matches the second `Button` element.

```
#Cancel {
    border-width: 2px;
    border-color: DarkRed;
    background-color: pink;
}
```

The UI looks like the following when you apply the style:

![The Cancel button has a dark red border and a pink background.](../uploads/Main/uss-selectors-name.png)

The Cancel button has a dark red border and a pink background.

## Additional resources

* [Best practices for USS](UIE-USS-WritingStyleSheets.md)