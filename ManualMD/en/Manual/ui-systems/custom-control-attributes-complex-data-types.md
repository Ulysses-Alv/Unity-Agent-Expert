# Define custom control attributes for complex data types

You can use attribute converters and UxmlObject to support complex data types attributes for your custom controls.

## Attribute converters

[Attribute converters](../../ScriptReference/UIElements.UxmlAttributeConverter_1.md) convert a [UxmlAttribute](../../ScriptReference/UIElements.UxmlAttributeAttribute.md) type to and from a string. It provides a set of [built-in converters](../../ScriptReference/UIElements.UxmlAttributeConverter_1.md) for common types, such as `int`, `float`, `bool`, `string`, and `Vector2`. For complex data types, such as lists or custom classes, you need to create custom converters.

The following C# code example creates a custom class called `Person` and a custom control class called `Department`. The `Department` class contains a `Person` field for the manager and a list of `Person` objects for the employees.

```
using System;
using System.Collections.Generic;
using UnityEngine.UIElements;

namespace AttributeConverterExample
{
    [Serializable]
    public class Person
    {
        public string name;
        public int age;
        public string nationality;
    }

    [UxmlElement]
    public partial class Department : VisualElement
    {
        [UxmlAttribute]
        public Person manager;

        [UxmlAttribute]
        public List<Person> employees;
    }
}
```

If you edit the custom control in UI Builder, you get the following error message:

```
[UxmlElement] 'Department' has a [UxmlAttribute] 'manager' of an unknown type 'Person'. To fix this error, define a custom UxmlAttributeConverter<Person>.
```

To address this, create an attribute converter for the `Person` class that inherits from `UxmlAttributeConverter<Person>` and implements the `ToString` and `FromString` methods to convert the `Person` object to and from a string. The following example creates a converter for the `Person` class:

```
using UnityEditor.UIElements;

namespace AttributeConverterExample
{
    public class PersonConverter : UxmlAttributeConverter<Person>
    {
        const char k_Separator = ':';

        public override string ToString(Person value)
        {
            return $"{value.name}{k_Separator}{value.age}{k_Separator}{value.nationality}";
        }

        public override Person FromString(string value)
        {
            var person = new Person();
            var split = value.Split(k_Separator);
            if (split.Length == 3)
            {
                person.name = split[0];
                person.age = int.Parse(split[1]);
                person.nationality = split[2];
            }
            return person;
        }
    }
}
```

The converter combines the three values into a single string format, with the pattern of `[name]:[age]:[nationality]`. The `FromString` method splits the string into the three values and assigns them to the `Person` object.

The following shows an example usage of the `Department` element in a UXML file:

```
<ui:UXML xmlns:ui="UnityEngine.UIElements">
    <AttributeConverterExample.Department name="Dunder Mifflin" manager="Michael Scott:58:USA" employees="Dwight Schrute:53:USA,Jim Halpert:44:USA" />
</ui:UXML>
```

## UxmlObjects

When using attribute converters, as data types become more intricate or when dealing with extensive lists, the generated strings can quickly grow convoluted and unwieldy. To simplify this, you can use `UxmlObjects`.

[`UxmlObjects`](../../ScriptReference/UIElements.UxmlObjectAttribute.md) are classes instantiated from UXML that contain UXML attributes.

To understand how to use `UxmlObjects`, take the previous `Person` and `Department` classes example, convert them into `UxmlObjects` as shown in the following example:

```
using System;
using System.Collections.Generic;
using UnityEngine.UIElements;

namespace UxmlObjectExample
{
    [UxmlObject]
    public partial class Person
    {
        [UxmlAttribute]
        public string name;

        [UxmlAttribute]
        public int age;

        [UxmlAttribute]
        public string nationality;
    }

    [UxmlElement]
    public partial class Department : VisualElement
    {
        [UxmlObjectReference("manager")]
        public Person manager;

        [UxmlObjectReference("employees")]
        public List<Person> employees;
    }
}
```

The `Person` class now resembles a custom element, and uses `UxmlAttribute` for attribute declaration and introducing the [UxmlObject](../../ScriptReference/UIElements.UxmlObjectAttribute.md) attribute to signify its status as a `UxmlObject`.

Instead of using `UxmlAttribute` for `UxmlObject` fields, this example uses [UxmlObjectReference](ScriptRefe:UIElements.UxmlObjectReferenceAttribute) to assign a [`name`](../../ScriptReference/UIElements.UxmlObjectReferenceAttribute-name.md). The name specifies the parent element of the `UxmlObject`.

Previously, all `UxmlObjects` were stored as direct children of the element, which caused scalability issues when multiple `UxmlObject` fields were present. It was difficult to distinguish which `UxmlObjects` belonged to which field.

The `UxmlObjectReference` attribute resolves this issue by specifying a name for the `UxmlObject`. This name corresponds to the UXML tag name. In the previous example, the `manager` and `employees` fields belong to the `Person` class. The `UxmlObjectReference` attribute ensures that the `Person` class is correctly assigned to the `manager` and `employees` fields.

The following shows an example usage in a UXML file:

```
<ui:UXML xmlns:ui="UnityEngine.UIElements">
    <UxmlObjectExample.Department name="Department One">
        <manager>
            <UxmlObjectExample.Person name="Manager One" age="58" nationality="Canadian" />
        </manager>
        <employees>
            <UxmlObjectExample.Person name="Person Two" age="53" nationality="Canadian" />
            <UxmlObjectExample.Person name="Person Three" age="44" nationality="Canadian" />
        </employees>
    </UxmlObjectExample.Department>
</ui:UXML>
```

**Note:** `UxmlObjectReferenceFields` also works with interfaces. However, you can only assign `UxmlObjects` that implement the specified interface. For an example of using interfaces, refer to [UxmlObjectReferenceAttribute](../../ScriptReference/UIElements.UxmlObjectReferenceAttribute.md).

## Additional resources

* [Define custom control attributes for built-in types](custom-control-attributes-built-in-types.md)