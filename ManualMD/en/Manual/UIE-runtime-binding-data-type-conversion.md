# Convert data types

You can use type converters to convert data types between the data source and the UI. This allows you to:

* Bind data types that are different from the UI, such as binding to a `Texture2D` property from an `int` value.
* Convert data types that are the same as the UI when you want a different representation of the data, such as displaying an angle in degrees instead of radians.

There are two categories of type converters: global converters and per-binding converters.

## Global converters

Global converters are statically registered and are available to all binding instances, regardless of the binding modes. The binding system provides several pre-registered global converters for convenience, such as converters between style values and their underlying data representation. You can use them to bind a regular `float` or a `StyleKeyword` to a `StyleFloat` property, instead of using `StyleFloat` in the data source.

To register a global converter, use the [`ConverterGroups.RegisterGlobalConverter`](../ScriptReference/UIElements.ConverterGroups.RegisterGlobalConverter.md) method. This method requires a delegate to convert a source type to a destination type. The delegate must be a `ref` delegate to enable bidirectional usage. Upon registration, global converters automatically apply to any binding instances requiring that specific type of conversion, without any additional configuration.

Register global converters during application startup to ensure they’re available when needed. You can register global converters in static constructors, assembly initialization, or startup methods. The following example registers a global converter that converts between a `Texture2D` and a `TextureHandle`:

```
static class RegisterGlobalConverters
{
#if UNITY_EDITOR
    [InitializeOnLoadMethod]
#endif
    [RuntimeInitializeOnLoadMethod]
    public static void Register()
    {
        // Register global converters here
        ConverterGroups.RegisterGlobalConverter((ref TextureHandle handle) => TextureHandle.ResolveTexture(handle));
        ConverterGroups.RegisterGlobalConverter((ref Texture2D texture) => TextureHandle.FromTexture(texture));
    }
}

public struct TextureHandle
{
    public static Texture2D ResolveTexture(TextureHandle handle)
    {
        return /* Actual texture */; 
    } 
    
   public static TextureHandle FromTexture(Texture2D texture)
   {
        return new TextureHandle { handle = /* Compute handle */ }; 
   }

    public int handle;
}
```

## Per-binding converters

Per-binding converters apply to a specific binding instance. You can register an individual or a group of per-binding converters. You register per-binding converters where you create the binding instance. This ensures that the converters are available when the binding is used.

### Individual converter

To register an individual converter, use the [`DataBinding.sourceToUiConverters.AddConverter`](../ScriptReference/UIElements.DataBinding-sourceToUiConverters.md) or [`DataBinding.uiToSourceConverters.AddConverter`](../ScriptReference/UIElements.DataBinding-uiToSourceConverters.md) method. These methods require a delegate to convert a source type to a destination type. The delegate must be a `ref` delegate to enable bidirectional usage.

The following example registers and applies individual converters to convert between radians and degrees for a binding instance:

```
var binding = new DataBinding();
binding.sourceToUiConverters.AddConverter((ref float radian) => Mathf.RadToDeg * radian);
binding.uiToSourceConverters.AddConverter((ref float degree) => Mathf.DegToRad * degree);
```

### Group converter

To register a group of converters and apply them to a `DataBinding` instance, use the [`ConverterGroup`](../ScriptReference/UIElements.ConverterGroup.md) class. Register converter groups once during initialization, similar to global converters. You can register converter groups in static constructors, assembly initialization, or startup methods.

The following example shows how to register a group of converters:

```
static class RegisterConverterGroup
{
#if UNITY_EDITOR
    [InitializeOnLoadMethod]
#endif
    [RuntimeInitializeOnLoadMethod]
    public static void Register()
    {
        // Create a converter group.
        var group = new ConverterGroup("Inverters");

        // Add converters to the converter group.
        group.AddConverter((ref int v) => -v);
        group.AddConverter((ref float v) => -v);
        group.AddConverter((ref double v) => -v);
        // Register the converter group
        ConverterGroups.RegisterConverterGroup(group);

        // Add a converter to an existing converter group.
        if (ConverterGroups.TryGetConverterGroup("Inverters", out var group))
        {
            group.AddConverter((ref short v) => -v);
        }
    }
}
```

### Apply a converter group

After you have registered a converter group, you can apply it to a binding instance. You can apply a converter group to a binding instance in C#, UI Builder, or UXML.

To apply a converter group in C#, use the [`DataBinding.ApplyConverterGroupToUI`](../ScriptReference/UIElements.DataBinding.ApplyConverterGroupToUI.md) or [`DataBinding.ApplyConverterGroupToSource`](../ScriptReference/UIElements.DataBinding.ApplyConverterGroupToSource.md) method. These methods take a converter group name as a parameter.

The following example applies the converter group `Inverters` to a binding instance in C#:

```
var binding = new DataBinding();
if (ConverterGroups.TryGetConverterGroup("Inverters", out var group))
{
    binding.ApplyConverterGroupToUI(group);
    binding.ApplyConverterGroupToSource(group);
}
```

**Note**: When you apply a converter group onto another one, it operates in a “fire-and-forget” manner. This means that when you apply the converter group, it independently performs its intended function without additional ongoing monitoring or management from you.

The following example applies the converter group `Inverters` to a binding in UXML:

```
<ui:DataBinding source-to-ui-converters="Inverters" ui-to-source-converters="Inverters" />
```

For details about how to apply a converter group in UI Builder, refer to [Get started with runtime binding](UIE-get-started-runtime-binding.md).

## Make converters available in UI Builder

For converters to appear and be selectable in UI Builder:

* **Register converter groups globally**: Only globally registered converter groups are available in UI Builder’s dropdown.
* **Unique names**: Ensure each converter group has a unique name to avoid conflicts in UI Builder.
* **Register early**: Ensure converter groups are registered before UI Builder loads, typically in static constructors with appropriate attributes.

The UI Builder automatically discovers and lists all registered converter groups in its interface, allowing you to select them from dropdown menus when configuring data bindings.

## Best practices

Follow these tips and best practices to optimize performance:

* **Minimize memory allocation**: Whenever possible, avoid writing conversion delegates that allocate memory, particularly when dealing with `enum` types. Allocating memory during conversions can introduce unnecessary overhead and impact performance. Instead, opt for efficient and memory-friendly conversion approaches.
* **Keep converters efficient**: It’s important to keep converters focused on performing quick and efficient type conversions. Avoid performing extensive tasks or complex operations within converters, as this can lead to decreased performance and potentially introduce unnecessary complexity.
* **Integrate type conversion in the data source**: Instead of relying solely on per-binding conversions, consider incorporating the type conversion directly into your data source. By handling the conversion within the data source itself, you can streamline the process and potentially improve overall performance. This approach can also reduce the complexity associated with managing conversions on a per-binding basis.

## Known limitations

Source and destination types must match perfectly unless they’re of type `UnityEngine.Object`. For example, you can’t convert an `int` to a `float`, or a `float` to an `int`. This can be inconvenient, especially with `enum` types. This limitation will be addressed in a future release.

## Additional resources

* [Runtime data binding](UIE-runtime-binding.md)
* [Get started with runtime binding](UIE-get-started-runtime-binding.md)
* [Create runtime data binding](UIE-runtime-binding-types.md)
* [Create custom binding types](UIE-runtime-binding-custom-types.md)
* [Define a data source](UIE-runtime-binding-define-data-source.md)
* [Define logging levels](UIE-runtime-binding-logging-levels.md)