# Create a property visitor with low-level APIs

This example demonstrates how to create a property visitor with the low-level [`IPropertyBagVisitor`](../ScriptReference/Unity.Properties.IPropertyBagVisitor.md) and [`IPropertyVisitor`](../ScriptReference/Unity.Properties.IPropertyVisitor.md) interfaces. It produces the same outcome as the [Create a property visitor with the PropertyVisitor class](property-visitors-PropertyVisitor.md#data-and-utilities) example but with a visitor that implements these interfaces rather than one that derives from [`PropertyVisitor`](../ScriptReference/Unity.Properties.PropertyVisitor.md).

## Create data and utility classes

1. Create the same `Data` and `DebugUtilities` classes as in the [Create a property visitor with the PropertyVisitor class](property-visitors-PropertyVisitor.md#data-and-utilities) example.
2. Modify the `DebugUtilities` class to use the low-level visitor instead of the `PropertyVisitor`-derived visitor as follows:

   ```
   public static class DebugUtilities
   {
       private static readonly LowLevelVisitor s_Visitor = new();

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

## Create the visitor

Create the low-level visitor class as follows:

1. Create a `LowLevelVisitor` class that implements the `IPropertyVisitor` and `IPropertyBagVisitor` interfaces.
2. Add a [`StringBuilder`](https://learn.microsoft.com/en-us/dotnet/api/system.text.stringbuilder?view=net-7.0) field to the class. This can be used to build a string that represents the current state of an object.
3. Add a `Reset` method that clears the `StringBuilder` and resets the indent level.
4. Add a `GetDump` method that returns the string representation of the current state of an object.

   The `LowLevelVisitor` class now looks as follows:

   ```
   public class LowLevelVisitor
       : IPropertyBagVisitor
       , IPropertyVisitor
   {
       private const int k_InitialIndent = 0;

       private readonly StringBuilder m_Builder = new StringBuilder();
       private int m_IndentLevel = k_InitialIndent;

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

## Get the properties

To call the [`Accept`](../ScriptReference/Unity.Properties.Property_2.Accept.md) method on the property using `this`, implement the [`IPropertyVisitor`](../ScriptReference/Unity.Properties.IPropertyVisitor.md) interface. This interface allows you to specify the visitation behavior when visiting a property, similar to the [`PropertyVisitorVisitProperty`](../ScriptReference/Unity.Properties.PropertyVisitor.VisitProperty.md) method:

1. In the `LowLevelVisitor` class, override the `IPropertyBagVisitor.Visit` and `IPropertyVisitor.Visit` methods:

   ```
   void IPropertyBagVisitor.Visit<TContainer>(IPropertyBag<TContainer> propertyBag, ref TContainer container)
   {
       foreach (var property in propertyBag.GetProperties(ref container))
       {
           property.Accept(this, ref container);
       }
   }
           
   void IPropertyVisitor.Visit<TContainer, TValue>(Property<TContainer, TValue> property, ref TContainer container)
   {
       var value = property.GetValue(ref container);
       // Code goes here.
   }
   ```
2. The `IVisitPropertyAdapter` adapters used with the [`PropertyVisitor`](../ScriptReference/Unity.Properties.PropertyVisitor.md) base class require access to the internal state of the visitor, so they can’t be used outside of that class. However, you can define domain-specific adapters that have the necessary information. In the same file as the `LowLevelVisitor` class, define `IPrintValue` and `IPrintValue<in T>` interfaces for these adapters and a `PrintContext` struct to encapsulate the information the adapters need to format a property value as follows:

   ```
   // Create the following struct with methods to encapsulate the formatting of the message and display the value.
   public readonly struct PrintContext
   {
       private StringBuilder Builder { get; }
       private string Prefix { get; }
       public string PropertyName { get; }

       public void Print<T>(T value)
       {
           Builder.AppendLine($"{Prefix}- {PropertyName} = {{{TypeUtility.GetTypeDisplayName(value?.GetType() ?? typeof(T))}}} {value}");
       }
           
       public void Print(Type type, string value)
       {
           Builder.AppendLine($"{Prefix}- {PropertyName} = {{{TypeUtility.GetTypeDisplayName(type)}}} {value}");
       }

       public PrintContext(StringBuilder builder, string prefix, string propertyName)
       {
           Builder = builder;
           Prefix = prefix;
           PropertyName = propertyName;
       }
   }

   public interface IPrintValue
   {
   }

   public interface IPrintValue<in T> : IPrintValue
   {
       void PrintValue(in PrintContext context, T value);
   }
   ```
3. In the `LowLevelVisitor` class, override the adapter interface method to create type-specific adapters for [`Vector2`](../ScriptReference/Vector2.md) and [`Color`](../ScriptReference/Color.md), and update the implementation of the [`IPropertyVisitor`](../ScriptReference/Unity.Properties.IPropertyVisitor.md) to use the adapter first:

   ```
   public class LowLevelVisitor
       : IPropertyBagVisitor
       , IPropertyVisitor
       , IPrintValue<Vector2>
       , IPrintValue<Color>
   {
       public IPrintValue Adapter { get; set; }
           
       public LowLevelVisitor()
       {
           // For simplicity
           Adapter = this;
       }
       void IPropertyVisitor.Visit<TContainer, TValue>(Property<TContainer, TValue> property, ref TContainer container)
       {
           // Here, we need to manually extract the value.
           var value = property.GetValue(ref container);
           
           var propertyName = GetPropertyName(property);
           
           // We can still use adapters, but we must manually dispatch the calls. 
           if (Adapter is IPrintValue<TValue> adapter)
           {
               var context = new PrintContext(m_Builder, Indent, propertyName);
               adapter.PrintValue(context, value);
               return;
           }
               
           // Fallback behaviour here 
       }
           
       void IPrintValue<Vector2>.PrintValue(in PrintContext context, Vector2 value)
       {
           context.Print(value);
       }
       void IPrintValue<Color>.PrintValue(in PrintContext context, Color value)
       {
           const string format = "F3";
           var formatProvider = CultureInfo.InvariantCulture.NumberFormat;
           context.Print(typeof(Color), $"RGBA({value.r.ToString(format, formatProvider)}, {value.g.ToString(format, formatProvider)}, {value.b.ToString(format, formatProvider)}, {value.a.ToString(format, formatProvider)})");
       }
   }
   ```

## Complete visitor code

The complete code for the low-level visitor class and the custom adapters for `Color` and `Vector2` looks as follows:

```
using UnityEngine;
using System.Globalization;
using System.Text;
using Unity.Properties;

public readonly struct PrintContext
{
    // A context struct to hold information about how to print the property.
    private StringBuilder Builder { get; }
    private string Prefix { get; }
    public string PropertyName { get; }

    // Method to print the value of type T with its associated property name.
    public void Print<T>(T value)
    {
        Builder.AppendLine($"{Prefix}- {PropertyName} = {{{TypeUtility.GetTypeDisplayName(value?.GetType() ?? typeof(T))}}} {value}");
    }

    // Method to print the value with a specified type and its associated property name.
    public void Print(System.Type type, string value)
    {
        Builder.AppendLine($"{Prefix}- {PropertyName} = {{{TypeUtility.GetTypeDisplayName(type)}}} {value}");
    }

    // Constructor to initialize the PrintContext.
    public PrintContext(StringBuilder builder, string prefix, string propertyName)
    {
        Builder = builder;
        Prefix = prefix;
        PropertyName = propertyName;
    }
}

// Generic interface IPrintValue that acts as a marker interface for all print value adapters.
public interface IPrintValue
{
}

// Generic interface IPrintValue<T> to define how to print values of type T.
// This interface is used as an adapter for specific types (Vector2 and Color in this case).
public interface IPrintValue<in T> : IPrintValue
{
    void PrintValue(in PrintContext context, T value);
}

// LowLevelVisitor class that implements various interfaces for property visiting and value printing.
public class LowLevelVisitor : IPropertyBagVisitor, IPropertyVisitor, IPrintValue<Vector2>, IPrintValue<Color>
{
    private const int k_InitialIndent = 0;

    private readonly StringBuilder m_Builder = new StringBuilder();
    private int m_IndentLevel = k_InitialIndent;

    public IPrintValue Adapter { get; set; }

    public LowLevelVisitor()
    {
        // The Adapter property is set to this instance of LowLevelVisitor.
        // This means the current LowLevelVisitor can be used as a print value adapter for Vector2 and Color.
        Adapter = this;
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
            // If it's a collection element property, display it with brackets.
            ICollectionElementProperty => $"[{property.Name}]",
            // For other property types, display the name as it is
            _ => property.Name
        };
    }

    // This method is called when visiting a property bag (a collection of properties)
    void IPropertyBagVisitor.Visit<TContainer>(IPropertyBag<TContainer> propertyBag, ref TContainer container)
    {
        foreach (var property in propertyBag.GetProperties(ref container))
        {
            // Call the Visit method of IPropertyVisitor to handle individual properties.
            property.Accept(this, ref container);
        }
    }

    // This method is called when visiting each individual property of an object.
    // It tries to find a suitable adapter (IPrintValue<T>) for the property value type (TValue) and uses it to print the value.
    // If no suitable adapter is found, it falls back to displaying the value using its type name.
    void IPropertyVisitor.Visit<TContainer, TValue>(Property<TContainer, TValue> property, ref TContainer container)
    {
        // Here, we need to manually extract the value.
        var value = property.GetValue(ref container);

        var propertyName = GetPropertyName(property);

        // We can still use adapters, but we must manually dispatch the calls.
        // Try to find an adapter for the current property value type (TValue).
        if (Adapter is IPrintValue<TValue> adapter)
        {
            // If an adapter is found, create a print context and call the PrintValue method of the adapter.
            var context = new PrintContext(m_Builder, m_IndentLevel.ToString(), propertyName);
            adapter.PrintValue(context, value);
            return;
        }

        // Fallback behavior here - if no adapter is found, handle printing based on type information.
        var type = value?.GetType() ?? property.DeclaredValueType();
        var typeName = TypeUtility.GetTypeDisplayName(type);

        if (TypeTraits.IsContainer(type))
            m_Builder.AppendLine($"{m_IndentLevel}- {propertyName} {{{typeName}}}");
        else
            m_Builder.AppendLine($"{m_IndentLevel}- {propertyName} = {{{typeName}}} {value}");

        // Recursively visit child properties (if any).
        ++m_IndentLevel;
        if (null != value)
            PropertyContainer.Accept(this, ref value);
        --m_IndentLevel;
    }

    // Method from IPrintValue<Vector2> used to print Vector2 values.
    void IPrintValue<Vector2>.PrintValue(in PrintContext context, Vector2 value)
    {
        // Simply use the Print method of PrintContext to print the Vector2 value.
        context.Print(value);
    }

    // Method from IPrintValue<Color> used to print Color values.
    void IPrintValue<Color>.PrintValue(in PrintContext context, Color value)
    {
        const string format = "F3";
        var formatProvider = CultureInfo.InvariantCulture.NumberFormat;

        // Format and print the Color value in RGBA format.
        context.Print(typeof(Color), $"RGBA({value.r.ToString(format, formatProvider)}, {value.g.ToString(format, formatProvider)}, {value.b.ToString(format, formatProvider)}, {value.a.ToString(format, formatProvider)})");
    }
}
```

When you have the complete version of this code in your `LowLevelVisitor.cs` file, call the `DebugUtilities.PrintObjectDump` method with an instance of the `Data` class to verify that it produces the same console output as in the final stage of the [Create a property visitor with the PropertyVisitor class](property-visitors-PropertyVisitor.md#example-output) example:

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

## Additional resources

* [Property visitors](property-visitors.md)
* [Property bags](property-bags.md)
* [Property paths](property-paths.md)
* [Create a property visitor with the PropertyVisitor class](property-visitors-PropertyVisitor.md)