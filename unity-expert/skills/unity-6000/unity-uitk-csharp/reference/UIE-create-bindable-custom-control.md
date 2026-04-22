# Create a bindable custom control

**Version**: 2023.2+

This example demonstrates how to create a bindable custom control in a custom Editor window.

## Example overview

This example creates a custom control bound to a property with the double data type. You can adapt this example to bind to properties with other data types such as a string or an integer.

![A script in the Unity Editor with a custom control.](../uploads/Main/custom-control-binding-example.png)

A script in the Unity Editor with a custom control.

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/2023.2/create-bindable-custom-control).

## Prerequisites

This guide is for developers familiar with the Unity Editor, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [UI Builder](UIBuilder.md)
* [Visual Tree](UIE-VisualTree.md)
* [UXML](UIE-UXML.md)
* [SerializedObject data binding](UIE-Binding.md)

## Create the custom control class

Create a C# class to define the custom control.

1. Create a Unity project with any template.
2. Create a folder named `ExampleField` to store your files.
3. In the `ExampleField` folder, create a C# script named `ExampleField.cs` and replace its content with the following:

```
using UnityEngine.UIElements;

namespace UIToolkitExamples
{
    // ExampleField inherits from BaseField with the double type. ExampleField's underlying value, then, is a double.
    [UxmlElement]
    public partial class ExampleField : BaseField<double>
    {
        Label m_Input;

        // Default constructor is required for compatibility with UXML factory
        public ExampleField() : this(null)
        {

        }

        // Main constructor accepts label parameter to mimic BaseField constructor.
        // Second argument to base constructor is the input element, the one that displays the value this field is
        // bound to.
        public ExampleField(string label) : base(label, new Label() { })
        {
            // This is the input element instantiated for the base constructor.
            m_Input = this.Q<Label>(className: inputUssClassName);
        }

        // SetValueWithoutNotify needs to be overridden by calling the base version and then making a change to the
        // underlying value be reflected in the input element.
        public override void SetValueWithoutNotify(double newValue)
        {
            base.SetValueWithoutNotify(newValue);

            m_Input.text = value.ToString("N");
        }
    }
}
```

## Use the custom control in a UXML file

1. In the `ExampleField` folder, create a UI Document named `ExampleField.uxml`.
2. Open `ExampleField.uxml` in a text editor and replace its contents with the following:

```
<ui:UXML xmlns:ui="UnityEngine.UIElements" xmlns:example="UIToolkitExamples">
    <example:ExampleField label="Binding Target" binding-path="m_Value" />
</ui:UXML>
```

## Create the class for the binding target

1. In Unity, double-click `ExampleField.uxml` to open it in the UI Builder. The ExampleField displays in the Hierarchy window and is visualized in the **Viewport**. If you select the ExampleField in the Hierarchy window, the **Inspector** window shows the values assigned to the **Binding Path** and **Label** boxes.
2. In the `ExampleField` folder, create a C# script named `ExampleFieldComponent.cs` and replace its contents with the following:

```
using UnityEngine;

namespace UIToolkitExamples
{
    public class ExampleFieldComponent : MonoBehaviour
    {
        [SerializeField]
        double m_Value;
    }
}
```

## Create the custom Editor for the binding target

1. In the `ExampleField` folder, create a folder named `Editor`.
2. In the `Editor` folder, create a C# script named `ExampleFieldCustomEditor.cs` and replace its contents with the following:

```
using UnityEditor;
using UnityEngine.UIElements;
using UnityEngine;

namespace UIToolkitExamples
{
    [CustomEditor(typeof(ExampleFieldComponent))]
    public class ExampleFieldCustomEditor : Editor
    {
        [SerializeField]
        VisualTreeAsset m_Uxml;

        public override VisualElement CreateInspectorGUI()
        {
            var parent = new VisualElement();

            m_Uxml?.CloneTree(parent);

            return parent;
        }
    }
}
```

## Test the example

1. In Unity, select `ExampleFieldCustomEditor.cs` in the **Project window**.
2. Drag `ExampleField.uxml` into the **Uxml** box in the Inspector window.
3. Add an empty **GameObject** to a **scene**.
4. Add the `ExampleFieldComponent` component to the GameObject. The custom control appears in the Inspector with the default value of `0` for the Binding Target. If you change the value of the underlying double property, the UI reflects that change.

## Additional resources

* [Custom controls](UIE-create-custom-controls.md)