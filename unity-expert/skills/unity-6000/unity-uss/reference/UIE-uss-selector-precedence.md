# Selector precedence

When an element matches more than one selector, Unity considers several factors to determine which selector takes precedence.

How Unity determines precedence depends on whether the conflicting selectors are in the same style sheet or in different style sheets.

## Precedence for selectors in the same style sheet

When an element matches multiple selectors from the same style sheet, the selector with the highest [specificity](#selector-specificity) takes precedence.

If both selectors have the same [specificity](#selector-specificity), the selector that appears last in the USS file takes precedence.

## Precedence for selectors in different style sheets

When an element matches multiple selectors in different style sheets, Unity determines precedence according to the following factors in this order:

1. **The type of style sheet:** Selectors from user-defined style sheets takes precedence over selectors from default Unity style sheets.
2. **Selector specificity:** If both selectors are in the same type of style sheet, the selector with the highest [specificity](#selector-specificity) takes precedence
3. **The style sheets’ positions in the element hierarchy:** If both selectors have the same specificity, the selector whose style sheet is applied lowest in the element hierarchy takes precedence.
4. **The selectors’ positions in their style sheets:** If you apply both style sheets at the same level of the hierarchy, the selector closest to the end of its USS file takes precedence.

## Selector specificity

Selector specificity is a measure of relevance. The higher the specificity, the more relevant the selector is to the elements it matches.

* **Name** selectors are more specific than **Class** selectors.
* **Class** selectors are more specific than **Type** selectors.
* **Type** selectors are more specific than the **Universal** (`*`) selector.
* **Universal** selectors have the lowest specificity.

## Style overrides

You can style an element in UI Toolkit by the following:

* In a [USS](UIE-USS.md) file, write a selector that can affect the element.
* In the [UXML](UIE-UXML.md) document where the element is defined, set inline styles.
* In a C# script, [get a reference to an element and set its inline style](UIE-apply-styles-with-csharp.md).
* For an inherited property, the element gets the value from its parent element.

If you apply multiple styling methods to an element, it will undergo style overrides.

### Applied styles override inherited styles

Styles that target an element directly take precedence over styles that the element inherits, even if the inherited styles are defined in a selector with higher specificity.

### Inline styles override USS styles

Inline styles that you apply to elements in a UXML document take precedence over USS styles. They have a higher specificity than USS selectors.

**Note:** USS doesn’t support the `!important` rule used to override style declarations in CSS.

### C# styles override inline and USS styles

Styles that you set in a C# script override USS styles and inline styles set in a UXML document. They have the highest specificity.

## Additional resources

* [Apply styles in C# scripts](UIE-apply-styles-with-csharp.md)
* [Best practices for USS](UIE-USS-WritingStyleSheets.md)