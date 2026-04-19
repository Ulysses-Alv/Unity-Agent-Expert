# Comparison of the binding systems

The following table compares the SerializedObject data binding and runtime binding:

| **Comparison** | **Runtime binding** | **SerializedObject data binding** |
| --- | --- | --- |
| UI | Runtime UI, and Editor UI without serialized data. | Editor UI |
| Data types | Supports any plain C# `object` as a data source. | Supports only the data types supported by `SerializedProperty`. |
| Binding target | Can target multiple properties of the same control. | Targets only the `value` property of a `INotifyValueChanged<T>` control. |
| Property paths for a list or an array | Uses the syntax of `Path.To.List[2]`. | Uses the syntax of `Path.To.List.Array.data[2]`. |
| Extensibility | You can create your own binding types and attributes. | Not extensible. |

## Additional resources

* [Runtime data binding](UIE-runtime-binding.md)
* [SerializedObject data binding](UIE-editor-binding.md)