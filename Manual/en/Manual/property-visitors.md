# Property visitors

Property visitors define the logic for visiting the property bag of a given type to perform operations on its properties. If a type has been instrumented to expose its properties through a property bag, then this allows you to create additional functionality for it without further modifying the type directly.

You can create highly generic visitors to define both the visitation algorithm itself and the visitation process. This differs from the classic implementation of the visitor pattern by allowing you to handle any type rather than only specific ones. It enables features like serialization and Inspector-like UI generation for types you discover at runtime.

The following is the fundamental flow of visitation, which happens on the property bags and their companion objects:

1. An instance of a type accepts a visitor.
2. The visitor visits the property bag of the instance.
3. The property bag can loop through its properties and accept the visitor.

## Create a property visitor to get properties

You can use the following approaches to create your visitors to get properties:

* Derive from the [`Unity.Properties.PropertyVisitor`](../ScriptReference/Unity.Properties.PropertyVisitor.md) base class. For an example, refer to [Use `PropertyVisitor` to create a property visitor](property-visitors-PropertyVisitor.md).
* Implement the [`IPropertyBagVisitor`](../ScriptReference/Unity.Properties.IPropertyBagVisitor.md) and [`IPropertyVisitor`](../ScriptReference/Unity.Properties.IPropertyVisitor.md) interfaces. For an example, refer to [Use low-level APIs to create a property visitor](property-visitors-low-level-api.md).

The first approach is the easiest way to get started. However, for more extensive customization of the visitation behavior for both the property bags and the properties, use the second approach, which offers greater flexibility and the potential for performance improvements.

The following example uses the [`PropertyVisitor`](../ScriptReference/Unity.Properties.PropertyVisitor.md) class to create a simple visitor that gets the properties of a given type that are tagged with a certain attribute:

```
public class BindableAttribute
    : Attribute
{
}

public class GatherBindablePropertiesVisitor
    : PropertyVisitor
{
    public List<PropertyPath> BindableProperties { get; set; }

    protected override void VisitProperty<TContainer, TValue>(Property<TContainer, TValue> property, ref TContainer container, ref TValue value)
    {
        if (property.HasAttribute<BindableAttribute>())
            BindableProperties.Add(PropertyPath.AppendProperty(default, property));
    }
}
```

The following is the equivalent example that uses the `IPropertyBagVisitor` interface to create the visitor:

```
public class BindableAttribute
    : Attribute
{
}

public class GatherBindablePropertiesVisitor
    : IPropertyBagVisitor
{
    public List<PropertyPath> BindableProperties { get; set; }

    void IPropertyBagVisitor.Visit<TContainer>(IPropertyBag<TContainer> propertyBag, ref TContainer container)
    {
        // Loop through the properties of the container object.
        foreach (var property in propertyBag.GetProperties(ref container))
        {
            if (property.HasAttribute<BindableAttribute>())
                BindableProperties.Add(PropertyPath.AppendProperty(default, property));
        }
    }
}
```

The low-level visitor is more performant because it doesn’t need to loop through all the properties of the property bag and extract their value. You can also use low-level visitors to visit properties that aren’t part of a property bag.

## Performance considerations

Property bags, properties, and visitors are all implemented using generic types to keep code as strongly-typed as possible and to avoid boxing allocations during visitation. The trade-off of using generic types is that the JIT compiler generates the [intermediate language (IL)](https://learn.microsoft.com/en-us/dotnet/standard/managed-code) for a given method the first time it’s called. This can result in a slower execution the first time a visitor is accepted on an object.

## Additional resources

* [Property bags](property-bags.md)
* [Property paths](property-paths.md)
* [Use `PropertyVisitor` to create a property visitor](property-visitors-PropertyVisitor.md)
* [Use low-level APIs to create a property visitor](property-visitors-low-level-api.md)