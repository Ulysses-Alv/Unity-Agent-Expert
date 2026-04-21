# Bind custom controls to data

You bind custom controls to serialized properties to synchronize values between the control and the property. You can create a bindable custom control derived from the [`BaseField`](../ScriptReference/UIElements.BaseField_1.md) generic base class instead of [`BindableElement`](../ScriptReference/UIElements.BindableElement.md). This provides the following advantages:

* Implements the `INotifyValueChanged` interface for the generic parameter type that you specify.
* Makes the control focusable by default.
* Provides a horizontal layout with a label element on the left and input on the right.

![FloatField is a built-in UI Toolkit control which inherits from BaseField.<br/>A. The label element.<br/>B. The input element.](../uploads/Main/base-field-example.png)

FloatField is a built-in UI Toolkit control which inherits from `BaseField`.  
A. The label element.  
B. The input element.

**Note**: It’s possible to create custom controls derived from built-in UI controls if you understand their internal hierarchy and existing USS classes. Unity discourages this practice because your custom controls might be dependent on their internal structure, which is subject to change in the future.

To [bind](UIE-Binding.md) custom controls to data:

* Implement the [`INotifyValueChanged`](../ScriptReference/UIElements.INotifyValueChanged_1.md) interface and listen for [`ChangeEvent`](../ScriptReference/UIElements.ChangeEvent_1.md)s as needed.
* Inherit from the [`BindableElement`](../ScriptReference/UIElements.BindableElement.md) class or implement the [`IBindable`](../ScriptReference/UIElements.IBindable.md) interface.

Refer to [SerializedObject data binding](UIE-Binding.md) for more details.

For a bindable custom control example, refer to [Create a bindable custom control](UIE-create-bindable-custom-control.md).

## Additional resources

* [Create custom controls](UIE-create-custom-controls.md)
* [SerializedObject data binding](UIE-Binding.md)
* [Create a bindable custom control](UIE-create-bindable-custom-control.md)