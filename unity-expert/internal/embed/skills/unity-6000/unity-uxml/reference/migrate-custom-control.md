# Migrate custom controls from an earlier version to Unity 6

Unity 6 introduces improvements that simplify the creation of custom controls. This guide explains how to migrate custom controls from previous versions to Unity 6’s revamped workflow.

## Old system

Previously, creating a custom control required using the `UxmlTraits` class to define attributes. The following example demonstrates this approach for a `ProgressBar` control.

The examples focus on attribute authoring and don’t include the behavior of the control. It creates a `UxmlTraits` class that defines the attributes of the `ProgressBar` control. For each attribute, the class specifies the following:

* Its type
* Default value
* Name

Then, it gets the value for each attribute and assigns it to the element. It also adds a `UxmlFactory` class that creates instances of the element.

```
using UnityEngine.UIElements;

public class ProgressBar : VisualElement
{
    public new class UxmlFactory : UxmlFactory<ProgressBar, UxmlTraits> { }

    public new class UxmlTraits : BindableElement.UxmlTraits
    {
        UxmlFloatAttributeDescription m_LowValue = new UxmlFloatAttributeDescription { name = "low-value", defaultValue = 0 };
        UxmlFloatAttributeDescription m_HighValue = new UxmlFloatAttributeDescription { name = "high-value", defaultValue = 100 };
        UxmlFloatAttributeDescription m_Value = new UxmlFloatAttributeDescription { name = "value", defaultValue = 0 };
        UxmlStringAttributeDescription m_Title = new UxmlStringAttributeDescription() { name = "title", defaultValue = string.Empty };

        public override void Init(VisualElement ve, IUxmlAttributes bag, CreationContext cc)
        {
            base.Init(ve, bag, cc);
            var bar = ve as ProgressBar;
            bar.lowValue = m_LowValue.GetValueFromBag(bag, cc);
            bar.highValue = m_HighValue.GetValueFromBag(bag, cc);
            bar.value = m_Value.GetValueFromBag(bag, cc);
            bar.title = m_Title.GetValueFromBag(bag, cc);
        }
    }

    public string title { get; set; }

    public float lowValue { get; set; }

    public float highValue { get; set; }

    public float value { get; set; }
}
```

## Unity 6 system

Unity 6 simplifies custom control creation with the [`UxmlElement`](../../ScriptReference/UIElements.UxmlElementAttribute.md) and [`UxmlAttribute`](../../ScriptReference/UIElements.UxmlAttributeAttribute.md) attributes, eliminating the need for `UxmlTraits` and `UxmlFactory`. The new system streamlines the process by automatically handling the conversion of values to and from UXML attribute strings.

The following shows the updated the `ProgressBar` example:

```
using UnityEngine.UIElements;

[UxmlElement]
public partial class ProgressBar : VisualElement
{
    [UxmlAttribute]
    public string title { get; set; }

    [UxmlAttribute]
    public float lowValue { get; set; }

    [UxmlAttribute]
    public float highValue { get; set; } = 100;

    [UxmlAttribute]
    public float value { get; set; }
}
```

## Key changes

Here’s the key changes:

* Use `UxmlElement` to [declare a custom control](../UIE-create-custom-controls.md).
* Mark the class as partial.
* Use `UxmlAttribute` to [define attributes](custom-control-attributes-built-in-types.md), removing the need for `UxmlTraits`.
* The UXML attribute name is automatically derived from the property name, but you can [customize it](custom-control-customize-uxml-attributes.md).

## UXML usage

The UXML usage remains the same for both systems:

```
<ui:UXML xmlns:ui="UnityEngine.UIElements">
    <ProgressBar title="My Progress bar" low-value="0" high-value="1" value="0.5" />
</ui:UXML>
```

## Important considerations

If an element lacks the `UxmlElement` attribute, Unity defaults to the `UxmlTraits` and `UxmlFactory` systems for serialization. Ensure consistency by using a single serialization approach per **visual element**. While both systems can coexist in a UXML file, don’t mix them with a single element.

After migrating to the new system, reimport UXML assets to match the updated code. While this process happens automatically, but any previously built Asset Bundles containing UXML files need to be rebuilt to function correctly.

## Additional resources

* [Troubleshooting custom control library compilation](../UIE-troubleshooting-custom-control-library-compilation.md)
* [Create a custom control](../UIE-create-custom-controls.md)
* [Customize the custom control UXML tag name](custom-control-customize-uxml-tag-names.md)
* [Define UXML attributes for built-in data types](custom-control-attributes-built-in-types.md)
* [Define UXML attributes for complex data types](custom-control-attributes-complex-data-types.md)
* [Customize UXML attributes](custom-control-customize-uxml-attributes.md)