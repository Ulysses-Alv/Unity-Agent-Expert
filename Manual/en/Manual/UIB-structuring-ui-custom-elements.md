# Create a custom control with two attributes

**Version**: 2023.2+

This example demonstrates how to create a simple custom control with two attributes.

## Example overview

This example creates a custom control called `MyElement` with two attributes and exposes it to UXML and UI Builder. This example also shows how to add a custom control to a UI in the UI Builder.

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/2023.2/create-custom-control-with-custom-attributes).

## Prerequisites

This guide is for developers who are familiar with Unity, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [UXML](UIE-UXML.md)
* [UI Builder](UIBuilder.md)

## Create the example

To create a new custom control class in C#, inherit it from the `VisualElement` class. This allows you to create and use this element in C#, but won’t automatically expose it in UXML and UI Builder. To expose it, add the [`UxmlElement`](../ScriptReference/UIElements.UxmlElementAttribute.md) attribute. To expose the attributes, add the [`UxmlAttribute`](../ScriptReference/UIElements.UxmlAttributeAttribute.md) attribute to each property that you want to be visible in UXML and the UI Builder.

1. Create a Unity project with any template.
2. In the `Assets` folder, create a C# script named `MyElement.cs` with the following content:

```
using UnityEngine;
using UnityEngine.UIElements;

[UxmlElement]
partial class MyElement : VisualElement
{
    string _myString = "default_value";
    int _myInt = 2;

    [UxmlAttribute]
    public string myString
    {
        get => _myString;
        set
        {
            _myString = value;
            Debug.Log($"myString set to: {_myString}");
            // React here, e.g. update UI
        }
    }

    [UxmlAttribute]
    public int myInt
    {
        get => _myInt;
        set
        {
            _myInt = value;
            Debug.Log($"myInt set to: {_myInt}");
            // React here, e.g. update UI
        }
    }
}
```

## Create a UXML to see the attribute

1. Create a UXML file with any name you want.
2. Double-click the UXML file to open it in the UI Builder.
3. In the **Library** section of the UI Builder, select **Project** > **Custom Controls (C#)** > **MyElement**.
4. Drag **MyElement** to the Hierarchy window.
5. To see the custom attributes for **MyElement**, go to the ****Inspector**** tab of **MyElement**:

![Custom attributes for a custom control](../uploads/Main/UIBuilder/CustomElementAttributes.png)

Custom attributes for a custom control

## Additional resources

* [Create custom controls](UIE-create-custom-controls.md)