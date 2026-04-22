# Define logging levels

During the update of the binding, errors might occur where binding objects try to access invalid properties, encounter `null` values along a property path, or encounter missing type converters. If you set the binding system to log all errors to the Console, it can impact performance.

To control the console output, you can define logging levels for the binding system. The following are the available logging levels:

* **All logs**: Errors are consistently logged to the console.
* **One log per result**: Errors are logged to the console only the first time they occur.
* **No logs**: Error logging is disabled.

## In Unity Editor windows

You can set logging levels from the More (⋮) menu in a Unity Editor window that supports data binding, such as UI Builder, **Scene** view, Game view, and the PanelSettings’s **Inspector** windows.

To set logging levels:

1. Open the More (⋮) menu in the Unity Editor window.
2. Select **Binding Console Logs**.
3. Choose one of the logging levels from the list.

**Note:**:

* The default logging level in UI Builder is set to **No logs**.
* Each panel maintains independent logging level settings. PanelSettings assets and views (Scene, Game) operate on separate panels with their own logging configurations.

## In C# scripts

You can set the global and per-panel configurations to customize logging behavior in C# **scripts**.

The following example sets the global log level of all panels or windows.

```
Binding.SetGlobalLogLevel(BindingLogLevel.Once);
```

The following example sets the log level per panel:

```
Binding.SetPanelLogLevel(myElement.panel, BindingLogLevel.None);
```

**Note:** The per-panel or the per-window logging level settings override the global logging level settings.

## Troubleshooting

When debugging binding issues, verify that Unity’s property system can access your object’s properties. The binding system relies on `Unity.Properties` to detect and access bindable properties. If properties aren’t properly exposed or serialized, bindings might fail silently.

## Additional resources

* [Runtime data binding](UIE-runtime-binding.md)
* [Create runtime data binding](UIE-runtime-binding-types.md)
* [Define a data source](UIE-runtime-binding-define-data-source.md)
* [Define binding modes and update triggers](UIE-runtime-binding-mode-update.md)
* [Convert data types](UIE-runtime-binding-data-type-conversion.md)