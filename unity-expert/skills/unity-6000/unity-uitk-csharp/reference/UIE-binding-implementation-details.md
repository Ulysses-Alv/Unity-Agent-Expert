# Binding system implementation details

The following is an explanation of how the UI Toolkit binding system works at the implementation level.

## Binding creation

When you call the `Bind()` method, it synchronizes and tracks the first value asynchronously. This means that the `value` property of the fields can’t update immediately. This allows you to set up binding for hierarchies that aren’t yet added to any UI.

## Change detection

The binding system relies upon the serialized data of Unity objects to implement change detection.

After you create a binding between a `SerializedObject` and one or more elements, the system polls this object for changes at a regular interval, but not every frame, in two steps:

1. The system serializes and polls the `SerializedObject` in native code to detect any change to serialized bytes. If there are no changes, it stops.
2. If there are changes, the system compares the property and the displayed value on the bound element. If needed, the system updates the displayed value. It does this for each specific property-to-element binding for the given object.

## Operation throttling

Some binding operations might be too long to handle in a single frame. If the binding operations are handled in a single frame, the UI might become unresponsive. To prevent this, these binding operations might happen across several frames. For example, it might take several updates to detect changes, depending on the number of objects polled.

## Additional resources

* [SerializedObject data binding](UIE-Binding.md)
* [Bindable elements](UIE-bindable-elements.md)
* [Binding data type conversion](UIE-binding-data-type-conversion.md)
* [Binding examples](UIE-binding-examples.md)