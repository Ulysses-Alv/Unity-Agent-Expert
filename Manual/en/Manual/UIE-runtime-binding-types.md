# Create a runtime binding in C# scripts

To bind a property of a **visual element** to a data source in C#, create an instance of [`DataBinding`](../ScriptReference/UIElements.DataBinding.md). With this binding type, you can define a [`dataSource`](../ScriptReference/UIElements.DataBinding-dataSource.md) and a [`dataSourcePath`](../ScriptReference/UIElements.DataBinding-dataSourcePath.md) directly on the binding instance.

## Basic workflow

To create a runtime binding in C#, follow these steps:

1. Create a binding. Bindings are objects that you can create, register, or unregister to a visual element through a unique ID.
2. [Define the data source](UIE-runtime-binding-define-data-source.md) and the data source path for the binding object. The data source is the object that contains the property you want to bind to. The data source path is a relative path from the data source to the property you want to bind to.
3. [Define the binding mode and update triggers](UIE-runtime-binding-mode-update.md) for the binding object. The binding mode defines how changes are replicated between the data source and the UI. The update trigger defines when to update the binding object.
4. Register the binding object to the visual element.
5. If necessary, [add type converters to convert data types between the data source and the UI](UIE-runtime-binding-data-type-conversion.md).

Assume you have a data source object named `ExampleMultiPropertiesObject` that contains a `Vector3` property named `vector3Value` and a `float` property named `sumOfVector3Properties`, which is the sum of the `x`, `y`, and `z` components of the `vector3Value` property. For more information about this data source object, refer to the [Create a data source asset](UIE-bind-to-multiple-properties-with-runtime-binding.md#create-a-data-source-asset) section of the [Bind to multiple properties with runtime binding](UIE-bind-to-multiple-properties-with-runtime-binding.md) example.

You can then create a binding object and register it to a visual element as follows:

```
using Unity.Properties;
using UnityEngine;
using UnityEngine.UIElements;

public class RuntimeBindingExample
{
    public VisualElement CreateBindingExample()
    {
        var dataSource = ScriptableObject.CreateInstance<ExampleMultiPropertiesObject>();

        var root = new VisualElement
        {
            name = "root",
            dataSource = dataSource
        };

        var vector3Field = new Vector3Field("Vec3 Field");

        vector3Field.SetBinding("value", new DataBinding
        {
            dataSourcePath = new PropertyPath(nameof(ExampleMultiPropertiesObject.vector3Value))
        });

        root.Add(vector3Field);

        var floatField = new FloatField("Float Field") { value = 42.2f };

        floatField.SetBinding("value", new DataBinding
        {
            dataSourcePath = new PropertyPath(nameof(ExampleMultiPropertiesObject.sumOfVector3Properties))
        });

        root.Add(floatField);

        return root;
    }
}
```

It’s equivalent to the following UXML:

```
<engine:UXML xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:engine="UnityEngine.UIElements" xmlns:editor="UnityEditor.UIElements" noNamespaceSchemaLocation="../../UIElementsSchema/UIElements.xsd" editor-extension-mode="False">
    <engine:VisualElement data-source="MulPropertyObject.asset" name="VisualElement" style="flex-grow: 1;">
        <engine:Vector3Field label="Vec3 Field">
            <Bindings>
                <engine:DataBinding property="value" data-source-path="vector3Value" binding-mode="ToSource"/>
            </Bindings>
        </engine:Vector3Field>
        <engine:FloatField label="Float Field" name="FloatField">
            <Bindings>
                <engine:DataBinding property="value" data-source-path="sumOfVector3Properties" binding-mode="ToTarget"/>
            </Bindings>
        </engine:FloatField>
    </engine:VisualElement>
</engine:UXML>
```

## Register and unregister binding objects

You can use the following methods to manage binding objects:

* [`SetBinding`](../ScriptReference/UIElements.VisualElement.SetBinding.md)
* [`GetBinding`](../ScriptReference/UIElements.VisualElement.GetBinding.md)
* [`TryGetBinding`](../ScriptReference/UIElements.VisualElement.TryGetBinding.md)
* [`GetBindingInfos`](../ScriptReference/UIElements.VisualElement.GetBindingInfos.md)
* [`HasBinding`](../ScriptReference/UIElements.VisualElement.HasBinding.md)
* [`ClearBinding`](../ScriptReference/UIElements.VisualElement.ClearBinding.md)
* [`ClearBindings`](../ScriptReference/UIElements.VisualElement.ClearBindings.md)

## Report a change

You can create the bindable properties in the same way as other data sources, which means that you can also use `VisualElement` types as [data sources](UIE-runtime-binding-define-data-source.md). The main difference between a `VisualElement` type and other data sources is that `VisualElement` types come with built-in versioning. You must use the built-in versioning of a `VisualElement` type to propagate changes.

To report a change, call the [`NotifyPropertyChanged`](../ScriptReference/UIElements.CallbackEventHandler.NotifyPropertyChanged.md) method. This method takes a [`BindingId`](../ScriptReference/UIElements.BindingId.md) that identifies the property that changed. `BindingId` is typically the name of the target property of the UI element.

The following example shows how to report a change for a `VisualElement` type:

```
public static readonly BindingId intValueProperty = nameof(intValue);

private int m_IntValue;

[CreateProperty]
public int intValue
{
    get => m_IntValue;
    set
    {
        if (m_IntValue == value)
            return;
        m_IntValue = value;
        
        // This instructs the binding system that a change occured.
        NotifyPropertyChanged(intValueProperty);
    }
}
```

## Best practices

Follow these tips and best practices to optimize performance:

* **Use correct binding IDs**: The binding system uses the binding ID to identify the binding object and the target property of the element. The binding ID must be the target property of the element. For example, if you want to bind to the `value` property of a `Vector3Field`, the binding ID must be `Vector3Field.valueProperty` or `value`. In a nutshell, the binding ID must match the name or path of the field or property.
* **Avoid internal data updates with bindings**: Don’t use bindings to update the internal data of a visual element. For example, don’t use bindings to synchronize the `x`, `y`, and `z` sub-elements for a `Vector3Field`. Instead, use a binding to synchronize the `value` property of the `Vector3Field` with a `Vector3` property of a data source.

## Known limitations

UI Toolkit doesn’t report changes in `element.style` and `element.resolvedStyle`. Therefore, you can use binding instances to target the resolved style of an element but can’t track changes to them.

## Additional resources

* [Runtime data binding](UIE-runtime-binding.md)
* [Create custom binding types](UIE-runtime-binding-custom-types.md)
* [Define logging levels](UIE-runtime-binding-logging-levels.md)