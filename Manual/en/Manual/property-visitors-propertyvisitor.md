# Create a property visitor with the PropertyVisitor class

This example demonstrates how to use the [`PropertyVisitor`](../ScriptReference/Unity.Properties.PropertyVisitor.md) base class to create a property visitor. For an equivalent example that uses the [`IPropertyBagVisitor`](../ScriptReference/Unity.Properties.IPropertyBagVisitor.md) and [`IPropertyVisitor`](../ScriptReference/Unity.Properties.IPropertyVisitor.md) interfaces, refer to [Use low-level APIs to create a property visitor](property-visitors-low-level-api.md).

## Create data and utility classes

Create a simple data class and a utility class that uses the visitor to visit and print its properties:

1. Create the following C# class called `Data` with properties that can be visited:

   ```
   using System.Collections.Generic;
   using UnityEngine;

   public class Data
   {
       public string Name = "Henry";
       public Vector2 Vec2 = Vector2.one;
       public List<Color> Colors = new List<Color> { Color.green, Color.red };
       public Dictionary<int, string> Dict = new Dictionary<int, string> { { 5, "zero" } };
   }
   ```
2. Create the following utility class called `DebugUtilities`:

   ```
   using UnityEngine;

   public static class DebugUtilities
   {
       public static void PrintObjectDump<T>(T value)
       {
       // Magic goes here.
       }
   }
   ```

## Create the visitor

Create the visitor class as follows:

1. Create a class named `ObjectVisitor` that inherits from [`PropertyVisitor`](../ScriptReference/Unity.Properties.PropertyVisitor.md).
2. Add a [`StringBuilder`](https://learn.microsoft.com/en-us/dotnet/api/system.text.stringbuilder?view=net-7.0) field to the class. This can be used to build a string that represents the current state of an object.
3. Add a `Reset` method that clears the `StringBuilder` and resets the indent level.
4. Add a `GetDump` method that returns the string representation of the current state of an object.

   The `ObjectVisitor` class now looks as follows:

   ```
   // `PropertyVisitor` is an abstract class that you must derive from. 
   public class ObjectVisitor: PropertyVisitor
   {
       private const int k_InitialIndent = 0;
       private readonly StringBuilder m_Builder = new StringBuilder();
           
       private int m_IndentLevel = k_InitialIndent;
           
       private string Indent => new (' ', m_IndentLevel * 2);
           
       public void Reset()
       {
           m_Builder.Clear();
           m_IndentLevel = k_InitialIndent;
       }

       public string GetDump()
       {
           return m_Builder.ToString();
       }
   }
   ```

## Get the top-level properties

Inside the `ObjectVisitor` class, override the [`VisitProperty`](../ScriptReference/Unity.Properties.PropertyVisitor.VisitProperty.md) method to visit each property of an object and log the property name. The [`PropertyVisitor`](../ScriptReference/Unity.Properties.PropertyVisitor.md) doesn’t require any members to implement and by default, it simply visits every property and performs no action.

1. Inside the `ObjectVisitor` class, add the following override `VisitProperty` method:

   ```
   protected override void VisitProperty<TContainer, TValue>(Property<TContainer, TValue> property, ref TContainer container, ref TValue value)
   {
       m_Builder.AppendLine($"- {property.Name}");
   }
   ```
2. Now that you have a minimal visitor, you can implement the utility method. Update the `PrintObjectDump` method in the `DebugUtilities` class to create a new `ObjectVisitor` instance and use it to visit the properties of the given object:

   ```
   public static class DebugUtilities
   {
       private static readonly ObjectVisitor s_Visitor = new();
           
       public static void PrintObjectDump<T>(T value)
       {
           s_Visitor.Reset();
               
           // This is the main entry point to run a visitor.
           PropertyContainer.Accept(s_Visitor, ref value);
           Debug.Log(s_Visitor.GetDump());
       }
   }
   ```
3. From an appropriate place in your code, call the `PrintObjectDump` method and pass a `Data` object to it. For example, from a simple MonoBehaviour component this might look as follows:

   ```
   using UnityEngine;

   public class MyMonoBehaviour : MonoBehaviour
   {
       
       void Start()
       {
       DebugUtilities.PrintObjectDump(new Data());
       }

   }
   ```

   This prints the following output to the console:

   ```
   - Name
   - Vec2
   - Colors
   - Dict
   ```

## Get the sub-properties

As the output from the previous section shows, when you override the [`VisitProperty`](../ScriptReference/Unity.Properties.PropertyVisitor.VisitProperty.md) method, it doesn’t automatically visit the sub-properties of an object. To get the sub-properties, use the [`PropertyContainer.Accept`](../ScriptReference/Unity.Properties.PropertyContainer.Accept.md) method to apply the visitor on each of the values recursively.

In the `ObjectVisitor` class, update the `VisitProperty` method to apply the visitor recursively on the value to nest in:

```
protected override void VisitProperty<TContainer, TValue>(Property<TContainer, TValue> property, ref TContainer container, ref TValue value)
{
    m_Builder.AppendLine($"{Indent}- {property.Name}");
        
    ++m_IndentLevel;
    // Apply this visitor recursively on the value to nest in.
    if (null != value)
        PropertyContainer.Accept(this, ref value);
    --m_IndentLevel;
}
```

The console output is now as follows:

```
- Name
- Vec2
- x
- y
- Colors
- 0
    - r
    - g
    - b
    - a
- 1
    - r
    - g
    - b
    - a 
- Dict
- 5
    - Key
    - Value
```

## Display more information about each property

To further enhance the printed output, get the property name of collection items, and the type and value of each property. Certain properties have special names, especially when dealing with collection items. The conventions for property names are as follows:

* For list items, the name corresponds to the index.
* For dictionaries, the name is derived from the string version of the key value.
* For sets, the name is based on the string version of the value.

To make this distinction more explicit, enclose the property name in square brackets.

1. Inside the `ObjectVisitor` class, add the following method:

   ```
   private static string GetPropertyName(IProperty property)
   {
       return property switch
       {
           // You can also treat `IListElementProperty`, `IDictionaryElementProperty`, and `ISetElementProperty` separately.
           ICollectionElementProperty => $"[{property.Name}]",
           _ => property.Name
       };
   }
   ```
2. Update the `VisitProperty` method to use [`TypeUtility.GetTypeDisplayName`](../ScriptReference/Unity.Properties.TypeUtility.GetTypeDisplayName.md) to retrieve the display name of a given type.

   ```
   protected override void VisitProperty<TContainer, TValue>(Property<TContainer, TValue> property, ref TContainer container, ref TValue value)
   {
       var propertyName = GetPropertyName(property);
           
       // Get the concrete type of the property or its declared type if value is null.
       var typeName = TypeUtility.GetTypeDisplayName(value?.GetType() ?? property.DeclaredValueType());
           
       m_Builder.AppendLine($"{Indent}- {propertyName} = {{{typeName}}} {value}");
           
       ++m_IndentLevel;
       if (null != value)
           PropertyContainer.Accept(this, ref value);
       --m_IndentLevel;
   }
   ```

   The console output is now as follows:

   ```
   - Name = {string} Henry
   - Vec2 = {Vector2} (1.00, 1.00)
   - x = {float} 1
   - y = {float} 1
   - Colors = {List<Color>} System.Collections.Generic.List`1[UnityEngine.Color]
   - [1] = {Color} RGBA(0.000, 1.000, 0.000, 1.000)
       - r = {float} 0
       - g = {float} 1
       - b = {float} 0
       - a = {float} 1
   - [1] = {Color} RGBA(1.000, 0.000, 0.000, 1.000)
       - r = {float} 1
       - g = {float} 0
       - b = {float} 0
       - a = {float} 1
   - Dict = {Dictionary<int, string>} System.Collections.Generic.Dictionary`2[System.Int32,System.String]
   - [5] = {KeyValuePair<int, string>} [5, five]
       - Key = {int} 5
       - Value = {string} five
   ```

## Display less information for collection types

Because `List<T>` doesn’t override the `ToString()` method, the list value is displayed as `System.Collections.Generic.List1[UnityEngine.Color]`. To reduce the amount of information displayed, update the `VisitProperty` to use the [`TypeTraits.IsContainer`](../ScriptReference/Unity.Properties.TypeTraits.IsContainer.md) utility method to only display the value for types that don’t contain sub-properties, such as primitives, enums, and strings.

In the `ObjectVisitor` class, update the `VisitProperty` method to use [`TypeTraits.IsContainer`](../ScriptReference/Unity.Properties.TypeTraits.IsContainer.md) to determine whether the value is a container type. If it is, display the type name without the value. Otherwise, display the type name and the value:

```
protected override void VisitProperty<TContainer, TValue>(Property<TContainer, TValue> property, ref TContainer container, ref TValue value)
{
    var propertyName = GetPropertyName(property);

    var type = value?.GetType() ?? property.DeclaredValueType();
    var typeName = TypeUtility.GetTypeDisplayName(type);
    
    // Only display the values for primitives, enums and strings.
    if (TypeTraits.IsContainer(type))
        m_Builder.AppendLine($"{Indent}- {propertyName} {{{typeName}}}");
    else
        m_Builder.AppendLine($"{Indent}- {propertyName} = {{{typeName}}} {value}");
    
    ++m_IndentLevel;
    if (null != value)
        PropertyContainer.Accept(this, ref value);
    --m_IndentLevel;
}
```

The console output is now as follows:

```
- Name = {string} Henry
- Vec2 {Vector2}
- x = {float} 1
- y = {float} 1
- Colors {List<Color>}
- [0] {Color}
    - r = {float} 0
    - g = {float} 1
    - b = {float} 0
    - a = {float} 1
- [1] {Color}
    - r = {float} 1
    - g = {float} 0
    - b = {float} 0
    - a = {float} 1
- Dict {Dictionary<int, string>}
- [5] {KeyValuePair<int, string>}
    - Key = {int} 5
    - Value = {string} five
```

**Tip**: To reduce the amount of information displayed, you can also use the following methods to override a `Visit` specialization for the collection types:

* [`PropertyVisitor.VisitCollection`](../ScriptReference/Unity.Properties.PropertyVisitor.VisitCollection.md)
* [`PropertyVisitor.VisitList`](../ScriptReference/Unity.Properties.PropertyVisitor.VisitList.md)
* [`PropertyVisitor.VisitDictionary`](../ScriptReference/Unity.Properties.PropertyVisitor.VisitDictionary.md)
* [`PropertyVisitor.VisitSet`](../ScriptReference/Unity.Properties.PropertyVisitor.VisitSet.md)

These are similar to the [`VisitProperty`](../ScriptReference/Unity.Properties.PropertyVisitor.VisitProperty.md) method, but they expose the generic parameters of their respective collections types.

## Add type-specific overrides

Add type-specific overrides to display [`Vector2`](../ScriptReference/UnityEngine.Vector2.md) and [`Color`](../ScriptReference/UnityEngine.Color.md) types in a more compact manner.

Use [`PropertyVisitor`](../ScriptReference/Unity.Properties.PropertyVisitor.md) along with [`IVisitPropertyAdapter`](../ScriptReference/Unity.Properties.IVisitPropertyAdapter.md). When an adapter is registered for a specific type, if the targeted type is encountered during the visitation, the adapter is called instead of the [`VisitProperty`](../ScriptReference/Unity.Properties.PropertyVisitor.VisitProperty.md) method.

In the `ObjectVisitor` class, add the `IVisitPropertyAdapter` for `Vector2` and `Color`:

```
public class DumpObjectVisitor
    : PropertyVisitor
    , IVisitPropertyAdapter<Vector2>
    , IVisitPropertyAdapter<Color>
{
    public DumpObjectVisitor()
    {
        AddAdapter(this);
    }
    
    void IVisitPropertyAdapter<Vector2>.Visit<TContainer>(in VisitContext<TContainer, Vector2> context, ref TContainer container, ref Vector2 value)
    {
        var propertyName = GetPropertyName(context.Property);
        m_Builder.AppendLine($"{Indent}- {propertyName} = {{{nameof(Vector2)}}} {value}");
    }

    void IVisitPropertyAdapter<Color>.Visit<TContainer>(in VisitContext<TContainer, Color> context, ref TContainer container, ref Color value)
    {
        var propertyName = GetPropertyName(context.Property);
        m_Builder.AppendLine($"{Indent}- {propertyName} = {{{nameof(Color)}}} {value}");
    }
}
```

## Start visitation on sub-properties

When you run a visitor on data, by default, it starts the visitation on the top-level object. For any property visitor, to start the visitation on sub-properties of an object, pass a [`PropertyPath`](../ScriptReference/Unity.Properties.PropertyPath.md) to the [`PropertyContainer.Accept`](../ScriptReference/Unity.Properties.PropertyContainer.Accept.md) method.

1. Update the `DebugUtilities` method to take an optional `PropertyPath`:

   ```
   public static class DebugUtilities
   {
       private static readonly ObjectVisitor s_Visitor = new();

       public static void PrintObjectDump<T>(T value, PropertyPath path = default)
       {
           s_Visitor.Reset();
           if (path.IsEmpty)
               PropertyContainer.Accept(s_Visitor, ref value);
           else
               PropertyContainer.Accept(s_Visitor, ref value, path);
           Debug.Log(s_Visitor.GetDump());
       }
   }
   ```
2. Call the `PrintObjectDump` method with the `Data` object. The console output is now as follows:

   ```
   - Name {string} = Henry
   - Vec2 {Vector2} = (1.00, 1.00)
   - Colors {List<Color>}
     - [0] = {Color} RGBA(0.000, 1.000, 0.000, 1.000)
     - [1] = {Color} RGBA(1.000, 0.000, 0.000, 1.000)
   - Dict {Dictionary<int, string>}
     - [5] {KeyValuePair<int, string>}
       - Key {int} = 5
       - Value {string} = five
   ```

## Complete visitor code

The complete code for the visitor class now looks as follows:

```
using System.Text;
using Unity.Properties;
using UnityEngine;

public class ObjectVisitor
: PropertyVisitor
, IVisitPropertyAdapter<Vector2>
, IVisitPropertyAdapter<Color>
{
    private const int k_InitialIndent = 0;

    // StringBuilder to store the dumped object's properties and values.
    private readonly StringBuilder m_Builder = new StringBuilder();
    private int m_IndentLevel = k_InitialIndent;

    // Helper property to get the current indentation.
    private string Indent => new(' ', m_IndentLevel * 2);

    public ObjectVisitor()
    {
        // Constructor, it initializes the ObjectVisitor and adds itself as an adapter
        // to handle properties of type Vector2 and Color.
        AddAdapter(this);
    }

    // Reset the visitor, clearing the StringBuilder and setting indentation to initial level.
    public void Reset()
    {
        m_Builder.Clear();
        m_IndentLevel = k_InitialIndent;
    }

    // Get the string representation of the dumped object.
    public string GetDump()
    {
        return m_Builder.ToString();
    }

    // Helper method to get the property name, handling collections and other property types.
    private static string GetPropertyName(IProperty property)
    {
        return property switch
        {
            // If it's a collection element property, display it with brackets
            ICollectionElementProperty => $"[{property.Name}]",
            // For other property types, display the name as it is
            _ => property.Name
        };
    }

    // This method is called when visiting each property of an object.
    // It determines the type of the value and formats it accordingly for display.
    protected override void VisitProperty<TContainer, TValue>(Property<TContainer, TValue> property, ref TContainer container, ref TValue value)
    {
        var propertyName = GetPropertyName(property);

        // Get the type of the value or property.
        var type = value?.GetType() ?? property.DeclaredValueType();
        var typeName = TypeUtility.GetTypeDisplayName(type);

        // Only display the values for primitives, enums, and strings, and treat other types as containers.
        if (TypeTraits.IsContainer(type))
            m_Builder.AppendLine($"{Indent}- {propertyName} {{{typeName}}}");
        else
            m_Builder.AppendLine($"{Indent}- {propertyName} = {{{typeName}}} {value}");

        // Increase indentation level before visiting child properties (if any).
        ++m_IndentLevel;
        if (null != value)
            PropertyContainer.Accept(this, ref value);
        // Decrease indentation level after visiting child properties.
        --m_IndentLevel;
    }
    
    // This method is a specialized override for Vector2 properties.
    // It displays the property name and its value as a Vector2.
    void IVisitPropertyAdapter<Vector2>.Visit<TContainer>(in VisitContext<TContainer, Vector2> context, ref TContainer container, ref Vector2 value)
    {
        var propertyName = GetPropertyName(context.Property);
        m_Builder.AppendLine($"{Indent}- {propertyName} = {{{nameof(Vector2)}}} {value}");
    }

    // This method is a specialized override for Color properties.
    // It displays the property name and its value as a Color.
    void IVisitPropertyAdapter<Color>.Visit<TContainer>(in VisitContext<TContainer, Color> context, ref TContainer container, ref Color value)
    {
        var propertyName = GetPropertyName(context.Property);
        m_Builder.AppendLine($"{Indent}- {propertyName} = {{{nameof(Color)}}} {value}");
    }
   
}
```

## Additional resources

* [Property visitors](property-visitors.md)
* [Property bags](property-bags.md)
* [Property paths](property-paths.md)
* [Create a property visitor with low-level APIs](property-visitors-low-level-api.md)