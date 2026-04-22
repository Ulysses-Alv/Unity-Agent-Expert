# Bind with binding path in C# script

**Version**: 2021.3+

This example demonstrates how to bind with the binding path in a C# script.

## Example overview

This examples create a custom Editor window to change the name of a **GameObject**.

![Displays a Unity Editor window with a custom field to change the name of a GameObject.](../uploads/Main/uie_simple_binding.png)

Displays a Unity Editor window with a custom field to change the name of a GameObject.

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/master/bind-with-binding-path).

## Prerequisites

This guide is for developers familiar with the Unity Editor, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [`bindingPath`](../ScriptReference/UIElements.BindableElement-bindingPath.md)
* [`Bind()`](../ScriptReference/UIElements.BindingExtensions.Bind.md)

## Bind with the binding path

Create a custom Editor window in C# with a `TextField`. Set the binding path to the name property of a GameObject and make an explicit call to the `Bind()` method.

1. Create a project in Unity with any template.
2. In your **Project window**, create a folder named `bind-with-binding-path` folder to store your file.
3. In the **bind-with-binding-path** folder, create a folder named `Editor`.
4. In the **Editor** folder, create a C# script named `SimpleBindingExample.cs` and replace its contents with the following:

   ```
   using UnityEditor;
   using UnityEngine;
   using UnityEditor.UIElements;
   using UnityEngine.UIElements;

   namespace UIToolkitExamples
   {
       public class SimpleBindingExample : EditorWindow
       {
           TextField m_ObjectNameBinding;

           [MenuItem("Window/UIToolkitExamples/Simple Binding Example")]
           public static void ShowDefaultWindow()
           {
               var wnd = GetWindow<SimpleBindingExample>();
               wnd.titleContent = new GUIContent("Simple Binding");
           }

           public void CreateGUI()
           {
               m_ObjectNameBinding = new TextField("Object Name Binding");
               // Note: the "name" property of a GameObject is "m_Name" in serialization.
               m_ObjectNameBinding.bindingPath = "m_Name";
               rootVisualElement.Add(m_ObjectNameBinding);
               OnSelectionChange();
           }

           public void OnSelectionChange()
           {
               GameObject selectedObject = Selection.activeObject as GameObject;
               if (selectedObject != null)
               {
                   // Create the SerializedObject from the current selection
                   SerializedObject so = new SerializedObject(selectedObject);
                   // Bind it to the root of the hierarchy. It will find the right object to bind to.
                   rootVisualElement.Bind(so);

                   // Alternatively you can instead bind it to the TextField itself.
                   // m_ObjectNameBinding.Bind(so);
               }
               else
               {
                   // Unbind the object from the actual visual element that was bound.
                   rootVisualElement.Unbind();
                   // If you bound the TextField itself, you'd do this instead:
                   // m_ObjectNameBinding.Unbind();

                   // Clear the TextField after the binding is removed
                   m_ObjectNameBinding.value = "";
               }
           }
       }
   }
   ```

## Test the binding

1. In Unity, select **Window** > **UIToolkitExamples** > **Simple Binding Example**. A custom Editor window appears with a text field.
2. Select any GameObject in your **scene**. The name of the GameObject appears in your Editor window’s text field. If you change the name of the GameObject in the text field, the name of the GameObject changes.

## Additional resources

* [SerializedObject data binding](UIE-Binding.md)
* [Bindable elements](UIE-bindable-elements.md)
* [Binding data type conversion](UIE-binding-data-type-conversion.md)
* [Implementation details](UIE-binding-implementation-details.md)
* [Binding examples](UIE-binding-examples.md)