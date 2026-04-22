# Bind with UXML and C# script

**Version**: 2021.3+

This example demonstrates how to create a binding and set the binding path with UXML, and call the `Bind()` method with a C# script.

## Example overview

This example creates a custom Editor window with a TextField that binds to the name property of any **GameObject** in a **scene**.

![Displays a Unity Editor window with a custom field to change the name of a GameObject.](../uploads/Main/uie_uxml_binding.png)

Displays a Unity Editor window with a custom field to change the name of a GameObject.

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/master/bind-with-uxml-csharp).

## Prerequisites

This guide is for developers familiar with the Unity Editor, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [Visual Tree](UIE-VisualTree.md)
* [UXML](UIE-UXML.md)
* [`bindingPath`](../ScriptReference/UIElements.BindableElement-bindingPath.md)
* [`Bind()`](../ScriptReference/UIElements.BindingExtensions.Bind.md)

## Define the binding path in UXML

Define the **visual elements** and the binding path in UXML.

1. Create a project in Unity with any template.
2. In your **Project window**, create a folder named `bind-with-uxml-csharp` folder to store all your files.
3. Create a UI Document named `binding_example.uxml` and replace its content with the following:

   ```
   <UXML xmlns:ui="UnityEngine.UIElements">
       <ui:VisualElement name="top-element">
           <ui:Label name="top-label" text="UXML-Defined Simple Binding"/>
           <ui:TextField name="GameObjectName" label="Name" text="" binding-path="m_Name"/>
       </ui:VisualElement>
   </UXML>
   ```

## Create the binding in C#

Create the binding in a C# script and make an explicit call to the `Bind()` method.

1. Create a folder named `Editor`.
2. In the **Editor** folder, create a C# script file named `SimpleBindingExampleUXML.cs`.
3. Replace the contents of `SimpleBindingExampleUXML.cs` with the following:

   ```
   using UnityEditor;
   using UnityEngine;
   using UnityEditor.UIElements;
   using UnityEngine.UIElements;

   namespace UIToolkitExamples
   {
       public class SimpleBindingExampleUXML : EditorWindow
       {
           [SerializeField]
           VisualTreeAsset visualTree;

           [MenuItem("Window/UIToolkitExamples/Simple Binding Example UXML")]
           public static void ShowDefaultWindow()
           {
               var wnd = GetWindow<SimpleBindingExampleUXML>();
               wnd.titleContent = new GUIContent("Simple Binding UXML");
           }

           public void CreateGUI()
           {
               visualTree.CloneTree(rootVisualElement);
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
               }
               else
               {
                   // Unbind the object from the actual visual element
                   rootVisualElement.Unbind();

                   // Clear the TextField after the binding is removed
                   var textField = rootVisualElement.Q<TextField>("GameObjectName");
                   if (textField != null)
                   {
                       textField.value = string.Empty;
                   }
               }
           }
       }
   }
   ```
4. In the Project window, select `SimpleBindingExampleUXML.cs` and drag `binding_example.uxml` to the **Visual Tree** field in the **Inspector**.

## Test the binding

1. In Unity, select **Window** > **UIToolkitExamples** > **Simple Binding Example UXML**. A custom Editor window appears with a text field.
2. Select any GameObject in your scene. The name of the GameObject appears in your Editor window’s text field.
3. If you change the name of the GameObject in the text field, the name of the GameObject changes.

## Additional resources

* [SerializedObject data binding](UIE-Binding.md)
* [Bindable elements](UIE-bindable-elements.md)
* [Binding data type conversion](UIE-binding-data-type-conversion.md)
* [Implementation details](UIE-binding-implementation-details.md)
* [Binding examples](UIE-binding-examples.md)