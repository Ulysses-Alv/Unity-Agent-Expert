# Customize UXML attributes

You can customize the UXML attribute names, the behavior of the attributes, and add decorators and custom **property drawers** in UI Builder.

## Rename the UXML attributes

You can rename a custom control’s attributes through the `name` argument of the [`UxmlAttribute`](../../ScriptReference/UIElements.UxmlAttributeAttribute.md) attribute.

The following shows an example of a custom control with renamed attributes:

```
using UnityEngine.UIElements;

[UxmlElement]
public partial class CustomAttributeNameExample : VisualElement
{
    [UxmlAttribute("character-name")]
    public string myStringValue { get; set; }

    [UxmlAttribute("credits")]
    public float myFloatValue { get; set; }

    [UxmlAttribute("level")]
    public int myIntValue { get; set; }

    [UxmlAttribute("usage")]
    public UsageHints myEnumValue { get; set; }
}
```

The following shows an example UXML document with the renamed attributes:

```
<ui:UXML xmlns:ui="UnityEngine.UIElements">
    <CustomAttributeNameExample character-name="Karl" credits="1.23" level="1" usage="DynamicColor" />
</ui:UXML>
```

If you rename an attribute and want UXML files with the old name to remain compatible, add the [`obsoleteNames`](../../ScriptReference/UIElements.UxmlAttributeAttribute-obsoleteNames.md) argument to the `UxmlAttribute` attribute.

The following example shows how to use the `obsoleteNames` argument:

```
using UnityEngine.UIElements;

[UxmlElement]
public partial class CharacterDetails : VisualElement
{
    [UxmlAttribute("character-name", "npc-name")]
    public string npcName { get; set; }

    [UxmlAttribute("character-health", "health")]
    public float health { get; set; }
}
```

The following shows an example usage in UXML:

```
<ui:UXML xmlns:ui="UnityEngine.UIElements">
    <CharacterDetails npc-name="Haydee" health="100" />
</ui:UXML>
```

## Enable runtime data binding in UI Builder

To simplify [runtime data binding](../UIE-runtime-binding.md) for the control’s properties, you can enable the **Add Binding** menu in UI Builder for a custom control’s attributes by adding the [`UxmlAttribute`](../../ScriptReference/UIElements.UxmlAttributeAttribute.md) (or [`UxmlObjectReferenceAttribute`](../../ScriptReference/UIElements.UxmlObjectReferenceAttribute.md) for a [complex data type](custom-control-attributes-complex-data-types.md)) and [`CreateProperty`](../../ScriptReference/Unity.Properties.CreatePropertyAttribute.md) attributes. For example:

```
using UnityEngine.UIElements;
using Unity.Properties;

[UxmlElement]
public partial class ExampleCustomControl : VisualElement
{
    [UxmlAttribute, CreateProperty]
    public string description { get; set; }
}
```

**Note**: Make sure to include the `using Unity.Properties;` directive in your script. For more information, refer to the [documentation](../properties.md) about `Unity Properties`.

In UI Builder, when you select the **ExampleCustomControl** in the **Hierarchy** pane and then right-click the **Description** in the ****Inspector**** pane, the **Add Binding** menu displays. You can then bind the control’s properties to data sources.

## Override the behavior of the UXML attributes

You override the get and set behavior for a UXML attribute of a custom control. This replaces the original attribute with your overridden version in the UI Builder attributes view. To override an attribute, define the new attribute and reference the original one using `[UxmlAttribute("original-attribute-name"), new-attribute]`.

You can use this attribute override to customize attributes inherited from child classes. For example, to enforce value limits in an **IntegerField**, override the value attribute and apply the [Range](../../ScriptReference/RangeAttribute.md) attribute like this:

```
using UnityEngine;
using UnityEngine.UIElements;

[UxmlElement]
public partial class MyCustomIntField : IntegerField
{
    [UxmlAttribute("value"), Range(0, 100)]
    private int valueOverride
    {
        get => this.value;
        set => this.value = value;
    }
}
```

In this example, the **MyCustomIntField** class inherits from **IntegerField** and overrides the **value** attribute. The **Range** attribute restricts the value to a range between `0` and `100`.

When you add the **MyCustomIntField** in the UI Builder, the overridden version replaces the **value** attribute. This allows you to set the value within the specified range.

## Add attribute decorators in UI Builder

You can add the following decorator attributes on the UXML attribute fields. When you add a decorator attribute, the corresponding UI controls, such as a slider for [`Range`](../../ScriptReference/RangeAttribute.md), are displayed in the UI Builder’s Inspector panel:

* [`TextArea`](../../ScriptReference/TextAreaAttribute.md)
* [`Tooltip`](../../ScriptReference/TooltipAttribute.md)
* [`Range`](../../ScriptReference/RangeAttribute.md)
* [`Header`](../../ScriptReference/HeaderAttribute.md)
* [`Min`](../../ScriptReference/MinAttribute.md)
* [`Multiline`](../../ScriptReference/MultilineAttribute.md)
* [`Space`](../../ScriptReference/SpaceAttribute.md)
* [`Delayed`](../../ScriptReference/DelayedAttribute.md)

The following example creates a custom control with decorators on its attribute fields:

```
using UnityEngine.UIElements;
using UnityEngine;

[UxmlElement]
public partial class ExampleText : VisualElement
{
    [TextArea, UxmlAttribute]
    public string myText;

    [Header("My Header")]
    [Range(0, 100)]
    [UxmlAttribute]
    public int rangedInt;

    [Tooltip("My custom tooltip")]
    [Min(10)]
    [UxmlAttribute]
    public int minValue = 100;
}
```

The UI Builder displays the attributes with the decorators:

![Attributes with the decorators](../../StaticFiles/ScriptRefImages/UIB-decorators.png)

Attributes with the decorators

## Add custom property drawer in UI Builder

You can add [custom property drawers](../../ScriptReference/PropertyDrawer.md) on fields in the UI Builder.

The following example creates a custom control with a custom property drawer:

```
using UnityEngine;
using UnityEngine.UIElements;

public class MyDrawerAttribute : PropertyAttribute { }

[UxmlElement]
public partial class MyDrawerExample : Button
{
    [UxmlAttribute]
    public Color myColor;

    [MyDrawer, UxmlAttribute]
    public string myText;
}
```

To access other serialized properties, prepend the name with `serializedData`. The following code example uses `serializedData.myColor` to find the `myColor` attribute:

```
using UnityEditor;
using UnityEditor.UIElements;
using UnityEngine.UIElements;

[CustomPropertyDrawer(typeof(MyDrawerAttribute))]
public class MyDrawerAttributePropertyDrawer : PropertyDrawer
{
    public override VisualElement CreatePropertyGUI(SerializedProperty property)
    {
        var row = new VisualElement { style = { flexDirection = FlexDirection.Row } };
        var textField = new TextField("My Text") { style = { flexGrow = 1 } };
        var button = new Button { text = ":" };
        button.clicked += () => textField.value = "RESET";

        // Get the parent property
        var parentPropertyPath = property.propertyPath.Substring(0, property.propertyPath.LastIndexOf('.'));
        var parent = property.serializedObject.FindProperty(parentPropertyPath);

        var colorProp = parent.FindPropertyRelative("myColor");
        textField.TrackPropertyValue(colorProp, p =>
        {
            row.style.backgroundColor = p.colorValue;
        });

        row.style.backgroundColor = colorProp.colorValue;
        row.Add(textField);
        row.Add(button);
        textField.BindProperty(property);

        return row;
    }
}
```

The custom property drawer looks like the following:

![Custom property drawer with a color picker](../../StaticFiles/ScriptRefImages/UIB-propertydrawer.gif)

Custom property drawer with a color picker

## Additional resources

* [create a custom inventory property drawer](example-create-custom-inventory-property-drawer.md)
* [`PropertyDrawer`](../../ScriptReference/PropertyDrawer.md)
* [`DecoratorDrawer`](../../ScriptReference/DecoratorDrawer.md)
* [`UxmlObjectReferenceAttribute`](../../ScriptReference/UIElements.UxmlObjectReferenceAttribute.md)
* [Troubleshooting custom control library compilation](../UIE-troubleshooting-custom-control-library-compilation.md)