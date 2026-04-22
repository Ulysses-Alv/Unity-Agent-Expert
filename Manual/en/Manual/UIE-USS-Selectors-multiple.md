# Multiple selectors

A multiple selector is an combination of multiple simple selectors. It selects any elements that match all the simple selectors.

## Syntax

A multiple selector consists of multiple simple selector without anything to separate them:

```
selector1selector2 {...}
```

The USS parser can’t parse a multiple selector if it can’t distinguish each selector in the combination.

For example, the following USS rule combines two type selectors: `ListView`, and `Button`.

```
ListViewButton{...}
```

The USS parser can’t separate the two element types, it interprets them as a single class called ListViewButton, which might not be the desired result.

You can combine [USS class selectors](UIE-USS-Selectors-class.md) and [name selectors](UIE-USS-Selectors-name.md) into multiple selectors. Because they’re are with the period (.) and number sign (#) respectively, the parser can clearly identify them. Type selectors have no identifying character, so you can only use one of them in a multiple selector, and it must be the first selector in the combination. For example:

```
ListView.yellow#vertical-list{...}
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

The following name selector style rule matches the first button.

```
Button.yellow {
  background-color: yellow;
}
```

The UI looks like the following when you apply the style:

![The OK button has a yellow background color.](../uploads/Main/uss-selectors-multiple.png)

The OK button has a yellow background color.

## Additional resources

* [Best practices for USS](UIE-USS-WritingStyleSheets.md)