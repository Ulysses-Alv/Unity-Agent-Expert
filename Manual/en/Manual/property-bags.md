# Property bags

A property bag is a companion object for a given .NET object type. The bag contains a collection of properties for its companion type. You can use the property bag to efficiently traverse, access, and modify data for an instance of an object of the companion type.

## Generating property bags

Unity uses one of the following methods to generate property bags for a type:

* **Reflection**: By default, Unity uses reflection to generate the property bag for a type. Reflection offers convenience and occurs [lazily](https://learn.microsoft.com/en-us/dotnet/framework/performance/lazy-initialization) only once per type when a property bag hasn’t been registered yet.
* **Code generation**: To enhance performance, you can opt-in to code generation. To generate property bags by code generation, you must:
  + Annotate the type with the [`[Unity.Properties.GeneratePropertyBag]`](../ScriptReference/Unity.Properties.GeneratePropertyBagAttribute.md) attribute.
  + Annotate the assembly with the [`[assembly: Unity.Properties.GeneratePropertyBagsForAssembly]`](../ScriptReference/Unity.Properties.GeneratePropertyBagsForAssemblyAttribute.md) attribute. Code-generated property bags are automatically registered when the domain is loaded.

## Included members

Both methods of generating property bags generate properties for the following type members:

* Public fields.
* Private or internal fields tagged with [`[SerializeField]`](../ScriptReference/SerializeField.md), [`[SerializeReference]`](../ScriptReference/SerializeReference.md), or [`[CreateProperty]`](../ScriptReference/Unity.Properties.CreatePropertyAttribute.md).
* Public, private, or internal properties tagged with [`[Unity.Properties.CreateProperty]`](../ScriptReference/Unity.Properties.CreatePropertyAttribute.md).

Adding the [`[DontCreateProperty]`](../ScriptReference/Unity.Properties.DontCreatePropertyAttribute.md) attribute to public, private, or internal fields excludes them from the property bag.

A generated property is read-only if the field is read-only or the property only has a getter. You can also use [`[Unity.Properties.CreateProperty(ReadOnly = true)]`](../ScriptReference/Unity.Properties.DontCreatePropertyAttribute.md) to make a generated property read-only.

The following example combines the Unity serialization system with the Unity Properties system:

```
using UnityEngine;
using Unity.Properties;

public class PropertyBagExample : MonoBehaviour
{
    // Serialization go through the field. 
    [SerializeField, DontCreateProperty] 
    private int m_Value;
    
    // Bindings go through the property rather than the field. 
    // This allows you to do validation, notify changes, and more.
    [CreateProperty] 
    public int value
    {
        get => m_Value;
        set => m_Value = value;
    }
    
    // This is a similar example, but for the backing field of an auto-property. 
    // Note that the name of the auto-property is postfixed with k_BackingField.
    [field: SerializeField, DontCreateProperty]
    [CreateProperty]
    public float floatValue { get; set; }
}
```

Creating properties in the property bag using serialization attributes for convenience is not always the preferred approach. [Unity’s serialization system](script-serialization.md) can only operate on fields and auto-properties, which makes it challenging to validate or propagate changes effectively.

Unlike the Unity serialization system, the properties within a property bag don’t qualify as value types with [`[SerializeField]`](../ScriptReference/SerializeField.md). Instead, struct types are recognized as value types, whereas class types are recognized as references.

In Unity serialization, although polymorphism is supported, you must use the [`[SerializeReference]`](../ScriptReference/SerializeReference.md) attribute to explicitly opt in. Otherwise, instances are serialized as value types. Note that [`UnityEngine.Object`](../ScriptReference/Object.md) types are an exception to this rule and are automatically serialized as reference types.

## Performance considerations

Generating property bags by reflection can introduce performance overhead the first time you request a property bag for a given container type. Creating properties for field members through reflection can allocate memory and increase garbage collector overhead in **IL2CPP** builds.

To avoid reflection and improve performance, [generate property bags by code generation](#generating) instead. However, be aware that this runtime optimization can come at the cost of longer compilation times. To enable the property bag to access internal and private fields and properties, make the type `partial`.

## Additional resources

* [Property visitors](property-visitors.md)
* [Property paths](property-paths.md)
* [Use `PropertyVisitor` to create a property visitor](property-visitors-PropertyVisitor.md)
* [Use low-level APIs to create a property visitor](property-visitors-low-level-api.md)