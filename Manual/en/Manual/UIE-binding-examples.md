# Binding examples

Learn Editor data binding from examples.

## Basic Editor binding examples

| **Topics** | **Description** |
| --- | --- |
| [Bind with binding path in C# script](UIE-create-a-binding-csharp.md) | Use `bindingPath` to create a binding that changes the name of a **GameObject** in a custom Editor window. |
| [Bind without the binding path](UIE-bind-without-bindpath.md) | Use `BindProperty()` to create a binding that changes the name of a GameObject in a custom Editor window. |
| [Bind with UXML and C#](UIE-create-a-binding-uxml-bind.md) | Create a binding and set the binding path in UXML, and bind with `Bind()` in C#. |
| [Create a binding with the Inspector](UIE-create-a-binding-uxml-inspector.md) | Create a binding that binds among a custom **Inspector**, a custom Editor, and serialized objects. |
| [Bind to nested properties](UIE-bind-to-nested-properties.md) | Use the `binding-path` attribute of a BindableElement in UXML to bind fields to nested properties of a SerializedObject |
| [Bind to a UXML template](UIE-bind-uxml-template.md) | Create a binding and set binding paths with UXML templates. |
| [Receive callbacks when a bound property changes](UIE-create-a-binding-callback.md) | Creates a custom Editor window with a TextField that binds to the name of a GameObject in a **scene**. |
| [Receive callbacks when any bound properties change](UIE-create-a-binding-callback-any-properties.md) | Create a custom Inspector with two fields that warns the user if the values of the fields fall outside certain ranges. |

## Advanced Editor binding examples

| **Topics** | **Description** |
| --- | --- |
| [Bind to a list with ListView](UIE-bind-to-list.md) | Create a list of toggles and binds the list to an underlying list of objects. |
| [Bind to a list without ListView](UIE-bind-to-list-without-listview.md) | Create a binding that binds to a list with array instead of ListView. |
| [Bind a custom control](UIE-bind-custom-control.md) | Create a custom control and bind it to a native Unity type. |
| [Bind a custom control to custom data type](UIE-bind-to-custom-data-type.md) | Create a custom control and bind it to a custom data type. |

## Additional resources

* [Create a bindable custom control](UIE-create-bindable-custom-control.md)