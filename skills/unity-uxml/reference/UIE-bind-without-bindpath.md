# Bind without the binding path

**Version**: 2021.3+

You can call [`BindProperty()`](../ScriptReference/UIElements.BindingExtensions.BindProperty.md) to bind an element to a `SerializedProperty` object directly, instead of with the binding path. This example demonstrates how to bind with [`BindProperty()`](../ScriptReference/UIElements.BindingExtensions.BindProperty.md).

## Example overview

This example creates a custom Editor window to change the name of a **GameObject**.

![A custom Unity Editor window which allows you to change a GameObjects name.](../uploads/Main/uie_simple_binding_property.png)

A custom Unity Editor window which allows you to change a GameObject’s name.

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/master/bind-without-binding-path).

## Prerequisites

This guide is for developers familiar with the Unity Editor, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [`BindProperty`](../ScriptReference/UIElements.BindingExtensions.BindProperty.md)

## Bind with `BindProperty()`

Create a custom Editor window in C# with a TextField. Find the name property of a GameObject and bind to the property directly with the [`BindProperty()`](../ScriptReference/UIElements.BindingExtensions.BindProperty.md) method.

1. In your **Project window**, create a folder named `bind-without-binding-path` to store your file.
2. In the **bind-without-binding-path** folder, create a folder named `Editor`.
3. In the **Editor** folder, create a C# script named `SimpleBindingPropertyExample.cs` and replace its contents with the following:

   ```
   using UnityEditor;
   using UnityEngine;
   using UnityEditor.UIElements;
   using UnityEngine.UIElements;

   namespace UIToolkitExamples
   {
       public class SimpleBindingPropertyExample : EditorWindow
       {
           TextField m_ObjectNameBinding;

           [MenuItem("Window/UIToolkitExamples/Simple Binding Property Example")]
           public static void ShowDefaultWindow()
           {
               var wnd = GetWindow<SimpleBindingPropertyExample>();
               wnd.titleContent = new GUIContent("Simple Binding Property");
           }
               
           public void CreateGUI()
           {
               m_ObjectNameBinding = new TextField("Object Name Binding");
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

                   // Note: the "name" property of a GameObject is actually named "m_Name" in serialization.
                   SerializedProperty property = so.FindProperty("m_Name");
                   // Bind the property to the field directly
                   m_ObjectNameBinding.BindProperty(property);
               }
               else
               {
                   // Unbind any binding from the field
                   m_ObjectNameBinding.Unbind();
               }
           }
       }
   }
   ```

## Test the binding

1. In Unity, select **Window** > **UIToolkitExamples** > **Simple Binding Property Example**. A custom Editor window appears with a text field.
2. Select a GameObject in your **scene**. The name of the GameObject appears in your Editor window’s text field. If you change the name of the GameObject in the text field, the name of the GameObject changes.

## Additional resources

* [SerializedObject data binding](UIE-Binding.md)
* [Bindable elements](UIE-bindable-elements.md)
* [Binding data type conversion](UIE-binding-data-type-conversion.md)
* [Implementation details](UIE-binding-implementation-details.md)
* [Binding examples](UIE-binding-examples.md)