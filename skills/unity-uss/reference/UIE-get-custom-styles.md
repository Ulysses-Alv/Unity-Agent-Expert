# Get custom styles in C# scripts

You can use the [`VisualElement.customStyle`](../ScriptReference/UIElements.VisualElement-customStyle.md) property to get the value of a [custom style property (variables)](UIE-USS-CustomProperties.md) applied to an element. However, you can’t directly query it as you would with [`VisualElement.style`](../ScriptReference/UIElements.VisualElement-style.md) or [`VisualElement.resolvedStyle`](../ScriptReference/UIElements.VisualElement-resolvedStyle.md). Instead, do the following:

1. Register to the [`CustomStyleResolvedEvent`](../ScriptReference/UIElements.CustomStyleResolvedEvent.md) event.
2. Call the [`TryGetValues`](../ScriptReference/UIElements.ICustomStyle.TryGetValue.md) method to query the returned object of the `element.customStyle` property.

Assume you have defined a custom style property `--my-custom-color` in USS like this:

```
.my-selector
{
    --my-custom-color: red;
}
```

The following example class shows how to get the value of `--my-custom-color` applied to an element:

```
public class HasCustomStyleElement : VisualElement
{
    // Custom style property definition from code indicating the type and the name of the property.
    private static readonly CustomStyleProperty<Color> s_CustomColor = new ("--my-custom-color");

    private Color customColor { get; set; }

    public HasCustomStyleElement()
    {
        RegisterCallback<CustomStyleResolvedEvent>(OnCustomStyleResolved);
    }

    private void OnCustomStyleResolved(CustomStyleResolvedEvent evt)
    {
        // If the custom style property is resolved for this element, you can query its value through the `customStyle` accessor.
        if (evt.customStyle.TryGetValue(s_CustomColor, out var value))
        {
            customColor = value;
        }
        // Otherwise, put some default value.
        else
        {
            customColor = new Color();
        }
    }
}
```

## Additional resources

* [Manage UI asset references from C# scripts](UIE-manage-asset-reference.md)
* [Apply styles with C#](UIE-apply-styles-with-csharp.md)
* [`customStyle`](../ScriptReference/UIElements.VisualElement-customStyle.md)
* [`CustomStyleResolvedEvent`](../ScriptReference/UIElements.CustomStyleResolvedEvent.md)
* [`TryGetValues`](../ScriptReference/UIElements.ICustomStyle.TryGetValue.md)