# Bindable data types and fields

This section describes the data type conversions and fields supported by [`PropertyField`](../ScriptReference/UIElements.PropertyField.md).

## Data type conversions

Most built-in controls in UI Toolkit implement the [`INotifyValueChanged`](../ScriptReference/UIElements.INotifyValueChanged_1.md) interface for a specific data type. For example, `DoubleField` implements `INotifyValueChanged<Double>`, which means you can bind the `DoubleField` control to a property of type `double`. You can bind more data types with the binding system. For example, you can bind a property of type `int` to a `DoubleField`.

The following table lists the source and target data types that you can bind:

| **Source data type** | **Target data type of `INotifyValueChanged`** |
| --- | --- |
| `long` | * `long` * `int` * `string` * `float` |
| `int` | * `int` * `string` * `float` |
| `double` | * `double` * `float` |
| `float` | * `float` * `double` |
| `char` | * `char` * `string` |

**Note**: To prevent data loss, use a [`PropertyField`](../ScriptReference/UIElements.PropertyField.md) or use the appropriate data types as shown in the [Fields supported by PropertyField table](#fields) below.

## Fields supported by `PropertyField`

When the **Inspector** window is populated, if a custom Editor doesn’t exist for a type, the binding system calls [`InspectorElement.FillDefaultInspector()`](../ScriptReference/UIElements.InspectorElement.FillDefaultInspector.md). This creates a [`PropertyField`](../ScriptReference/UIElements.PropertyField.md) for each [`SerializedProperty`](../ScriptReference/SerializedProperty.md) on the `SerializedObject`.

Each [`PropertyField`](../ScriptReference/UIElements.PropertyField.md) has a matching control under it, such as an `IntegerField` to represent an `int`, or a `Toggle` to represent a `bool`. If the property has sub-properties, the [`PropertyField`](../ScriptReference/UIElements.PropertyField.md) has a `Foldout` control. That `Foldout` has the appropriate default controls under each sub-property. If you created a custom `PropertyDrawer` for a Property, the **visual tree** returned from the `PropertyDrawer`’s `CreatePropertyGUI()` method is used instead.

**Note**:

When you create a custom UI, don’t use a [`PropertyField`](../ScriptReference/UIElements.PropertyField.md) unless you can’t find a specific control for it. For example, to create a control in your visual tree and bind it to a `Color`, use a [`ColorField`](../ScriptReference/UIElements.ColorField.md) rather than a [`PropertyField`](../ScriptReference/UIElements.PropertyField.md). This makes visual trees more lightweight and operations faster.

The use of `PropertyField` incurs additional overhead because it resolves the concrete field to use after binding. This extra step can affect performance. However, `PropertyField` is beneficial for some uses. For example, it supports custom drawers, which you can use to handle specific properties in a personalized way. This eliminates the need to guess the field type for drawing, which can reduce errors.

The following table lists fields supported by [`PropertyField`](../ScriptReference/UIElements.PropertyField.md) and their data type:

| **Field** | **Data type** |
| --- | --- |
| [`BoundsField`](../ScriptReference/UIElements.BoundsField.md) | Bounds |
| [`BoundsIntField`](../ScriptReference/UIElements.BoundsIntField.md) | BoundsInt |
| [`ColorField`](../ScriptReference/UIElements.ColorField.md) | Color |
| [`CurveField`](../ScriptReference/UIElements.CurveField.md) | AnimationCurve |
| [`FloatField`](../ScriptReference/UIElements.FloatField.md) | float |
| [`GradientField`](../ScriptReference/UIElements.GradientField.md) | Gradient |
| [`Hash128Field`](../ScriptReference/UIElements.Hash128Field.md) | Hash128 |
| [`IntegerField`](../ScriptReference/UIElements.IntegerField.md) | int |
| [`IntegerField`](../ScriptReference/UIElements.IntegerField.md) | int, for the ArraySize |
| [`LayerMaskField`](../ScriptReference/UIElements.LayerMaskField.md) | unit |
| [`LongField`](../ScriptReference/UIElements.LongField.md) | long |
| [`ObjectField`](../ScriptReference/UIElements.ObjectField.md) | UnityEngine.Object |
| [`PopupField`](../ScriptReference/UIElements.PopupField_1.md)<string> | Enum |
| [`RectField`](../ScriptReference/UIElements.RectField.md) | Rect |
| [`RectIntField`](../ScriptReference/UIElements.RectIntField.md) | RectInt |
| [`TextField`](../ScriptReference/UIElements.TextField.md) | string |
| [`TextField`](../ScriptReference/UIElements.TextField.md), with a maxLength=1 | char |
| [`Toggle`](../ScriptReference/UIElements.Toggle.md) | bool |
| [`Vector2Field`](../ScriptReference/UIElements.Vector2Field.md) | Vector2 |
| [`Vector2IntField`](../ScriptReference/UIElements.Vector2IntField.md) | Vector2Int |
| [`Vector3Field`](../ScriptReference/UIElements.Vector3Field.md) | Vector3 |
| [`Vector3IntField`](../ScriptReference/UIElements.Vector3IntField.md) | Vector3Int |
| [`Vector4Field`](../ScriptReference/UIElements.Vector4Field.md) | Vector4 |

## Additional resources

* [SerializedObject data binding](UIE-Binding.md)
* [Bindable elements](UIE-bindable-elements.md)
* [Implementation details](UIE-binding-implementation-details.md)
* [Binding examples](UIE-binding-examples.md)