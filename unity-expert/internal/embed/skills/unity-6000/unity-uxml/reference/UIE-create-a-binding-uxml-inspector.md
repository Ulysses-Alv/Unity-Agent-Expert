# Create a binding with the Inspector

**Version**: 2021.3+

The example demonstrates how to create bindings among a custom **Inspector**, a custom Editor, and serialized objects.

## Example overview

This example creates the following:

* A custom Inspector UI component with two fields that bind to a **GameObject**’s name and the scale of a USS transform.
* A custom Editor window that creates the same bindings inside an `InspectorElement`.

![Displays a Unity Editor window and Inspector with a custom field to change the name of a GameObject, and another field to adjust its size.](../uploads/Main/uie_bind_inspector.png)

Displays a Unity Editor window and Inspector with a custom field to change the name of a GameObject, and another field to adjust its size.

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/master/create-a-binding-inspector).

## Prerequisites

This guide is for developers familiar with the Unity Editor, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [Visual Tree](UIE-VisualTree.md)
* [UXML](UIE-UXML.md)
* [USS](UIE-USS.md)
* [`CreateInspectorGUI`](../ScriptReference/Editor.CreateInspectorGUI.md)
* [`InspectorElement`](../ScriptReference/UIElements.InspectorElement.md)

## Define the custom Inspector component

Define a custom Inspector UI class called `TankScript`, style it with USS, and set the **visual elements** and the binding path in UXML.

1. In your **Project window**, create a folder named `create-a-binding-inspector` to store all your files.
2. Create a C# script named `TankScript.cs` and replace its content with the following:

   ```
   using UnityEngine;

   [ExecuteInEditMode]
   public class TankScript : MonoBehaviour
   {
       public string tankName = "Tank";
       public float tankSize = 1;

       private void Update()
       {
           gameObject.name = tankName;
           gameObject.transform.localScale = new Vector3(tankSize, tankSize, tankSize);
       }
   }
   ```
3. In the **create-a-binding-inspector** folder, create a folder named `Editor`.
4. In the **Editor** folder, create a USS file named `tank_inspector_styles.uss` and replace its contents with the following:

   ```
   .container {
       background-color: rgb(80, 80, 80);
       flex-direction: column;
   }
   Label {
       background-color: rgb(80, 80, 80);
   }
   TextField:hover {
       background-color: rgb(255, 255, 0);
   }
   FloatField {
       background-color: rgb(0, 0, 255);
   }
   ```
5. Create a UXML file named `tank_inspector_uxml.uxml` and replace its contents with the following:

   ```
   <UXML xmlns="UnityEngine.UIElements" xmlns:ue="UnityEditor.UIElements">
       <Style src="tank_inspector_styles.uss" />
       <VisualElement name="row" class="container">
           <Label text="Tank Script - Custom Inspector" />
           <ue:PropertyField binding-path="tankName" name="tank-name-field" />
           <ue:PropertyField binding-path="tankSize" name="tank-size-field" />
       </VisualElement>
   </UXML>
   ```

## Create the Inspector UI with binding

Create a C# script to register the custom Inspector for `TankScript`. You don’t need to call the `Bind()` method as binding in the Inspector window is implicit.

1. Create a C# script named `TankEditor.cs` and replace its contents with the following:

   ```
   using UnityEditor;
   using UnityEngine;
   using UnityEngine.UIElements;

   [CustomEditor(typeof(TankScript))]
   public class TankEditor : Editor
   {
       [SerializeField]
       VisualTreeAsset visualTree;

       public override VisualElement CreateInspectorGUI()
       {
           var uxmlVE = visualTree.CloneTree();
           return uxmlVE;
       }
   }
   ```
2. In the Project window, select `TankEditor.cs`.
3. Drag `tank_inspector_uxml.uxml` to **Visual Tree** in the Inspector.
4. Drag `tank_inspector_styles.uss` to **Style Sheet** in the Inspector.

## Test the first binding

1. In Unity, add an empty GameObject to a **scene**.
2. Select the GameObject and drag `TankScript.cs` to **Add Component** in the Inspector. This generates the **Tank Script** component in the Inspector.
3. In **Tank Script**, you can change the name of the Tank in the **Tank Name** box, and change the Transform scale value in the **Tank Size** box.

## Bind with `InspectorElement`

Create a custom Editor window with two fields that bind to the `TankScript` object. Create a C# script that calls the [`InspectorElement` constructor](../ScriptReference/UIElements.InspectorElement-ctor.md). The [`InspectorElement`](../ScriptReference/UIElements.InspectorElement.md) creates the UI and automatically binds the UI to objects.

1. Create a C# file named `SimpleBindingExampleInspectorElement.cs` and replace its contents with the following:

   ```
   using UnityEditor;
   using UnityEngine;
   using UnityEditor.UIElements;
       
   namespace UIToolkitExamples
   {
       public class SimpleBindingExampleInspectorElement : EditorWindow
       {
           [MenuItem("Window/UIToolkitExamples/Simple Binding Example Inspector Element")]
           public static void ShowDefaultWindow()
           {
               var wnd = GetWindow<SimpleBindingExampleInspectorElement>();
               wnd.titleContent = new GUIContent("Simple Binding with Inspector Element");
           }
       
           TankScript m_Tank;
           public void OnEnable()
           {
               m_Tank = FindObjectOfType<TankScript>();
               if (m_Tank == null)
                   return;
       
               var inspector = new InspectorElement(m_Tank);
               rootVisualElement.Add(inspector);
           }
       }
   }
   ```

## Test the second binding

1. In Unity, select **Window** > **UIToolkitExamples** > **Simple Binding Example Inspector Element**.
2. In the Editor window, if you change the values in the fields, the values in the Inspector change, and vice versa.

## Additional resources

* [SerializedObject data binding](UIE-Binding.md)
* [Bindable elements](UIE-bindable-elements.md)
* [Binding data type conversion](UIE-binding-data-type-conversion.md)
* [Implementation details](UIE-binding-implementation-details.md)
* [Binding examples](UIE-binding-examples.md)