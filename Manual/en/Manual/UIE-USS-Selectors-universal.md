# Universal selectors

A USS universal selector, also called the wildcard selector, matches any element. USS universal selectors are analogous to CSS universal selectors.

## Syntax

The following is the syntax for a USS universal selector:

```
* { ... }
```

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

The following universal class selector style rule matches every element, and changes its background color to yellow. This includes the main area of the window, because the stylesheet is applied to the window’s root element.

```
* {
    background-color: yellow;
}
```

The UI looks like the following when you apply the style:

![Every element has a yellow background color.](../uploads/Main/uss-selectors-universal.png)

Every element has a yellow background color.

## Additional resources

* [Best practices for USS](UIE-USS-WritingStyleSheets.md)