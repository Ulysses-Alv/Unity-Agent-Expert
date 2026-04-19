# Child selectors

USS child selectors match elements that are the child of another element in the [visual tree](UIE-VisualTree.md).

## Syntax

A child selector consists of multiple simple selectors separated by `>`.

```
selector1 > selector2 {...}
```

You can include the wildcard selector in complex selectors. For example, the following USS rule uses the wildcard selector in a [child selector](UIE-USS-Selectors-child.md). This USS rule matches buttons that are children of elements that are children of an element with the USS class `yellow` assigned to it.

```
.yellow > * > Button{..}
```

**Note**: USS doesn’t support selecting the nth child of elements with the same name. Use [class selectors](UIE-USS-Selectors-class.md) or unique element names instead.

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

The following child selector style rule matches only the inner element. Element `#OK` with the `.yellow` class is a child of element `#container2`. `#container2` is child of element `#container1`. However, because `#OK` is not a direct descendant of `#container1`, it doesn’t match the `.yellow` selector when applied with a child selector from `#container1`.

```
#container1 > .yellow {
  background-color: yellow;
}
```

The UI looks like the following when you apply the style:

![The container2 has a yellow background color.](../uploads/Main/uss-selectors-child.png)

The container2 has a yellow background color.

## Additional resources

* [Best practices for USS](UIE-USS-WritingStyleSheets.md)