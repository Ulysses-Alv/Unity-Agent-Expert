# Introduction to SerializedObject data binding

You can use the SerializedObject data binding system to bind to [serialized properties](../ScriptReference/SerializedProperty.md). This means you can bind **visual elements** to the following objects that are compatible with the [Serialization system](script-serialization.md):

* User-defined [`ScriptableObject`](../ScriptReference/ScriptableObject.md) classes
* User-defined [`MonoBehaviour`](../ScriptReference/MonoBehaviour.md) classes
* Native Unity component types
* Native Unity asset types
* Primitive C# types such as `int`, `bool`, or `float`.
* Native Unity types such as `Vector3`, `Color`, or `Object`.

## Value binding

You can only bind the `value` property of visual elements that implement the [`INotifyValueChanged`](../ScriptReference/UIElements.INotifyValueChanged_1.md) interface. For example, you can bind [`TextField.value`](../ScriptReference/UIElements.TextField-value.md) to a `string`, but you can’t bind [`TextField.name`](../ScriptReference/UIElements.VisualElement-name.md) to a `string`.

You can bind between an object and any visual element that either derives from [`BindableElement`](../ScriptReference/UIElements.BindableElement.md) or implements the [`IBindable`](../ScriptReference/UIElements.IBindable.md) interface.

## Create a binding

To create a binding, either [call `Bind()`](#call-bind) or [`BindProperty()`](#call-bindproperty).

### Call `Bind()`

You can call [`Bind()`](../ScriptReference/UIElements.BindingExtensions.Bind.md) to bind an element to a [SerializedObject](../ScriptReference/SerializedObject.md). Before you bind an element, you must [set the binding path](#set-binding-path) and create a SerializedObject.

Use this method if you don’t have easy access to the `SerializedProperty` for the binding. Refer to [Create a binding with a C# script](UIE-create-a-binding-csharp.md) for an example.

The [`Bind()`](../ScriptReference/UIElements.BindingExtensions.Bind.md) extension method sets up an entire hierarchy of visual elements with specified [`bindingPath`](../ScriptReference/UIElements.IBindable-bindingPath.md) properties. You can call the [`Bind()`](../ScriptReference/UIElements.BindingExtensions.Bind.md) method on a single element or the parent of the hierarchy that you want to bind. For example, you can call [`Bind()`](../ScriptReference/UIElements.BindingExtensions.Bind.md) on the `rootVisualElement` of an Editor window. This binds all child elements with specified [`bindingPath`](../ScriptReference/UIElements.IBindable-bindingPath.md) properties.

Don’t call [`Bind`](../ScriptReference/UIElements.BindingExtensions.Bind.md) from the `Editor.CreateInspectorGUI` or `PropertyDrawer.CreatePropertyGUI` override. [`Bind`](../ScriptReference/UIElements.BindingExtensions.Bind.md) is called automatically on the visual elements that these methods return.

### Call `Unbind()`

The [`Unbind()`](../ScriptReference/UIElements.BindingExtensions.Unbind.md) method stops the value tracking for the element and all its direct and indirect child elements. In general, you don’t need to call [`Unbind()`](../ScriptReference/UIElements.BindingExtensions.Unbind.md) because tracking stops when a user closes the Inspector or Editor window. Call [`Unbind()`](../ScriptReference/UIElements.BindingExtensions.Unbind.md) if you must bind elements to different targets in their lifetimes.

If you construct an [`InspectorElement`](../ScriptReference/UIElements.InspectorElement.md) in C# by calling its constructor, binding occurs during the constructor call. If you want to rebind an [`InspectorElement`](../ScriptReference/UIElements.InspectorElement.md) after it has been constructed, you must call [`Unbind()`](../ScriptReference/UIElements.BindingExtensions.Unbind.md) and then either call [`Bind()`](../ScriptReference/UIElements.BindingExtensions.Bind.md) explicitly or let a bind operation from a parent create a binding.

### Set binding path

If you call [`Bind()`](../ScriptReference/UIElements.BindingExtensions.Bind.md) to create the binding, you must set the visual element’s binding path to the property name of the object that you want to bind to.

For example:

* If you have the following component script:

  ```
  using UnityEngine;

  public class MyComp : MonoBehaviour
  {
      [SerializeField]
      int m_Count;
  }
  ```

  To bind your visual element to `m_Count`, set the binding path to `m_Count`.
* If you want to bind a visual element to a **GameObject**’s name property, which is `m_Name`, set the binding path to `m_Name`.

You can set the binding path in UI Builder, UXML, or with a C# script:

* In UI Builder, enter the binding path in the **Binding Path** field for a visual element in the [Inspector](UIB-interface-overview.md#uibuilder-inspector).
* In UXML, set the `binding-path` attribute for a visual element. Refer to [Define the binding path in UXML](UIE-create-a-binding-uxml-bind.md#binding-path-uxml) for an example.
* In C#, set [`bindingPath`](../ScriptReference/UIElements.IBindable-bindingPath.md) from the [`IBindable`](../ScriptReference/UIElements.IBindable.md) interface. Refer to [Bind with the binding path](UIE-create-a-binding-csharp.md#binding-path-csharp) for an example.

### Call `BindProperty()`

You can call [`BindProperty()`](../ScriptReference/UIElements.BindingExtensions.BindProperty.md) to bind an element to a [`SerializedProperty`](../ScriptReference/SerializedProperty.md) directly.

Use this method if you already have a `SerializedProperty` object, and especially if you traverse the properties of a [`SerializedObject`](../ScriptReference/SerializedObject.md) to build a UI dynamically. Refer to [Bind without the binding path](UIE-bind-without-bindpath.md) for an example.

### Bind elements to nested properties

You can bind a visual element to nested properties in the source object. To do so, combine the binding path of an element with the binding path of the first ancestor. Use this method with the following elements:

* [`BindableElement`](../ScriptReference/UIElements.BindableElement.md)
* [`TemplateContainer`](../ScriptReference/UIElements.TemplateContainer.md) (corresponds to the `<Instance>` tag in UXML)
* [`GroupBox`](../ScriptReference/UIElements.GroupBox.md)

Refer to [Bind to nested properties](UIE-bind-to-nested-properties.md) for an example.

### Receive callbacks when values change

You can create a binding to receive a callback when a bound serialized property changes. To do so, leverage the [`TrackPropertyValue()`](../ScriptReference/UIElements.BindingExtensions.TrackPropertyValue.md) extension method, which is available to any `VisualElement`. This registers a callback that executes when the provided `SerializedProperty` changes. Refer to [Receive callbacks when a serialized property changes](UIE-create-a-binding-callback.md) for an example.

You can also create a binding to receive a callback when any properties of the bound serialized object change. To do so, leverage the [`TrackSerializedObjectValue()`](../ScriptReference/UIElements.BindingExtensions.TrackSerializedObjectValue.md) extension method, which is available to any `VisualElement`. This registers a callback that executes when the provided `SerializedObject` changes. Refer to [Receive callbacks when any properties change](UIE-create-a-binding-callback-any-properties.md) for an example.

### Bind custom elements

You can create custom elements and bind them to serialized properties through the [value binding](#value-binding) system.

To create bindable custom elements:

1. [Declare a custom element](UIE-create-custom-controls.md).
2. Inherit the element from `BindableElement` or implement the `IBindable` interface.
3. Implement the `INotifyValueChanged` interface.
4. Implement the `SetValueWithoutNotify()` method of the `INotifyValueChanged` interface.
5. Implement the `value` property accessors of the `INotifyValueChanged` interface.

Refer to [Create and style a custom control](UIE-bind-to-list-without-listview.md#custom-control) for an example.

**Note**: You can’t bind custom data types directly to custom elements because the binding system only allows you to bind an element to a type supported by SerializedPropertyType of `enum`. This means you can’t define a class or structure, for example, a struct called `MyStruct`, and bind it to an element that implements `INotifyValueChanged<MyStruct>`. However, you can bind to serializable nested properties of custom data types. This includes polymorphic serialization (when a field uses the [`[SerializeReference]` attribute](../ScriptReference/SerializeReference.md)). Refer to [Bind a custom control to custom data type](UIE-bind-to-custom-data-type.md) for an example.

## Bind time

Based on the type of UI you create, binding occurs at various times. This is called bind time.

The following table describes the bind time of a control:

| **Condition** | **Automatic bind time (assuming binding path was set)** |
| --- | --- |
| **An `InspectorElement` constructed in C#** | During the [constructor](../ScriptReference/UIElements.InspectorElement-ctor.md) call |
| **A child element that is under the return value of [`CreateInspectorGUI()`](../ScriptReference/Editor.CreateInspectorGUI.md) or [`CreatePropertyGUI()`](../ScriptReference/PropertyDrawer.CreatePropertyGUI.md) when those methods return** | After [`CreateInspectorGUI()`](../ScriptReference/Editor.CreateInspectorGUI.md) or [`CreatePropertyGUI()`](../ScriptReference/PropertyDrawer.CreatePropertyGUI.md) returns |
| **A child element that is under an element when [`Bind()`](../ScriptReference/UIElements.BindingExtensions.Bind.md) or [`BindProperty()`](../ScriptReference/UIElements.BindingExtensions.BindProperty.md) is called on the parent element** | During the [`Bind()`](../ScriptReference/UIElements.BindingExtensions.Bind.md) or [`BindProperty()`](../ScriptReference/UIElements.BindingExtensions.BindProperty.md) call |
| **Other** | No automatic binding; you must bind the element or one of its parents manually |

The following are best practices when creating a binding regarding bind time:

* If you create a custom [`Editor`](../ScriptReference/Editor.md) or custom [`PropertyDrawer`](../ScriptReference/PropertyDrawer.md), set the elements’ binding paths instead of calling [`Bind()`](../ScriptReference/UIElements.BindingExtensions.Bind.md) or [`BindProperty()`](../ScriptReference/UIElements.BindingExtensions.BindProperty.md) on any visual elements that are in the [visual tree](UIE-VisualTree.md) by the end of the body of [`CreateInspectorGUI()`](../ScriptReference/Editor.CreateInspectorGUI.md) or `CreatePropertyGUI()`. These elements are bound automatically after [`CreateInspectorGUI()`](../ScriptReference/Editor.CreateInspectorGUI.md) or [`CreatePropertyGUI()`](../ScriptReference/PropertyDrawer.CreatePropertyGUI.md) returns. However, if you add any elements to the visual tree after that point, call [`Bind()`](../ScriptReference/UIElements.BindingExtensions.Bind.md) or [`BindProperty()`](../ScriptReference/UIElements.BindingExtensions.BindProperty.md) to bind them.
* If you create any other type of UI, call [`Bind()`](../ScriptReference/UIElements.BindingExtensions.Bind.md) or [`BindProperty()`](../ScriptReference/UIElements.BindingExtensions.BindProperty.md) regardless of the time at which the elements get added to the visual tree. If you call [`Bind()`](../ScriptReference/UIElements.BindingExtensions.Bind.md) or [`BindProperty()`](../ScriptReference/UIElements.BindingExtensions.BindProperty.md) and bind multiple controls at the same time, set the binding path of each control and then call [`Bind()`](../ScriptReference/UIElements.BindingExtensions.Bind.md) on the lowest-level parent element that encompasses all the controls. [`Bind()`](../ScriptReference/UIElements.BindingExtensions.Bind.md) binds the element on which it’s called if it has a binding path and recursively binds all its child elements if they have binding paths. To prevent a negative performance impact, don’t bind a visual element with the [`Bind()`](../ScriptReference/UIElements.BindingExtensions.Bind.md) method more than once.

## Bind to a serialized property backing field

When you use an auto-property, the compiler automatically generates a backing field with a name as `<PropertyName>k__BackingField`. This field is not explicitly visible in your code but can be referenced if necessary, as in binding scenarios.

For example, the following example defines an auto-property `SomeProp` and serializes it:

```
[field: SerializeField] 
public float SomeProp 
{ 
    get; 
    private set; 
}
```

The compiler generates the following backing field:

```
[SerializedField]
private float <SomeProp>k__BackingField;

public float SomeProp
{
    get => <SomeProp>k__BackingField;
    set => <SomeProp>k__BackingField = value;
}
```

To bind to `<SomeProp>k__BackingField` in UXML, you must escape `<` and `>` because they’re reserved for tags. For example, set the `binding-path` as follows:

```
<editor:PropertyField name="some-prop" binding-path="&lt;SomeProp&gt;k__BackingField"/>
```

## Binding examples

Try the following examples to learn how to code with data binding:

* [Bind with binding path in C# script](UIE-create-a-binding-csharp.md)
* [Bind without the binding path](UIE-bind-without-bindpath.md)
* [Bind with UXML and C#](UIE-create-a-binding-uxml-bind.md)
* [Create a binding with the Inspector](UIE-create-a-binding-uxml-inspector.md)
* [Bind to nested properties](UIE-bind-to-nested-properties.md)
* [Bind to a UXML template](UIE-bind-uxml-template.md)
* [Receive callbacks when a bound property changes](UIE-create-a-binding-callback.md)
* [Receive callbacks when any bound properties change](UIE-create-a-binding-callback-any-properties.md)
* [Bind to a list with ListView](UIE-bind-to-list.md)
* [Bind to a list without ListView](UIE-bind-to-list-without-listview.md)
* [Bind a custom control](UIE-bind-custom-control.md)
* [Bind a custom control to custom data type](UIE-bind-to-custom-data-type.md)

## Additional resources

* [Binding data type conversion](UIE-binding-data-type-conversion.md)
* [Implementation details](UIE-binding-implementation-details.md)
* [Binding examples](UIE-binding-examples.md)