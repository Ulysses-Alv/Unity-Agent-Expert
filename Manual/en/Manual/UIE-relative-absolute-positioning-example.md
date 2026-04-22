# Relative and absolute positioning

This example demonstrates the difference between relative and absolute positioning. This example also demonstrates how to use C# and UXML/USS to add and style UI controls.

## Example overview

The example uses the automatic layout system to add boxes to an Editor and a runtime UI. One box demonstrates a relative offset of `25 px`, while another box demonstrates the absolute position of `25 px, 25 px`.

The example structures the Editor UI with C# script and the runtime UI with UXML and CSS.

![Visual element positioning](../uploads/Main/UIEcoordinatesExample.png)

Visual element positioning

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/master/relative-and-absolute-position).

## Prerequisites

This guide is for developers familiar with the Unity Editor, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [The visual tree](UIE-VisualTree.md)
* [Coordinate and position systems](UIE-coordinate-and-position-system.md)
* [UXML](UIE-UXML.md)
* [USS](UIE-USS.md)

## Create the example for the Editor UI

Create a custom Editor window and add all the boxes with a C# script: four boxes with gray backgrounds for comparison purposes; one box with a black background set up using absolute position placement; one box with a purple background set up using relative position placement.

1. Create a Unity project with any template.
2. Right-click in the Project window, then select **Create** > **UI Toolkit** > **Editor Window**.
3. In the **C#** box of the **UI Toolkit Editor Window Creator** window, enter `PositioningTestWindow`.
4. Clear the **UXML** and **USS** checkboxes.
5. Select **Confirm**. This creates a C# file called `PositioningTestWindow.cs`.
6. Replace `PositioningTestWindow.cs` with the following content:

   ```
   using UnityEditor;
   using UnityEngine;
   using UnityEngine.UIElements;

   public class PositioningTestWindow : EditorWindow
   {
       [MenuItem("Window/UI Toolkit/Positioning Test Window")]
       public static void ShowExample()
       {
           var wnd = GetWindow<PositioningTestWindow>();
           wnd.titleContent = new GUIContent("Positioning Test Window");
       }

       public void CreateGUI()
       {
           for (int i = 0; i < 2; i++)
           {
               var temp = new VisualElement();
               temp.style.width = 70;
               temp.style.height = 70;
               temp.style.marginBottom = 2;
               temp.style.backgroundColor = Color.gray;
               this.rootVisualElement.Add(temp);
           }

           // Relative positioning
           var relative = new Label("Relative\nPos\n25, 0");
           relative.style.width = 70;
           relative.style.height = 70;
           relative.style.left = 25;
           relative.style.marginBottom = 2;
           relative.style.backgroundColor = new Color(0.2165094f, 0, 0.254717f);
           this.rootVisualElement.Add(relative);

           for (int i = 0; i < 2; i++)
           {
               var temp = new VisualElement();
               temp.style.width = 70;
               temp.style.height = 70;
               temp.style.marginBottom = 2;
               temp.style.backgroundColor = Color.gray;
               this.rootVisualElement.Add(temp);
           }

           // Absolute positioning
           var absolutePositionElement = new Label("Absolute\nPos\n25, 25");
           absolutePositionElement.style.position = Position.Absolute;
           absolutePositionElement.style.top = 25;
           absolutePositionElement.style.left = 25;
           absolutePositionElement.style.width = 70;
           absolutePositionElement.style.height = 70;
           absolutePositionElement.style.backgroundColor = Color.black;
           this.rootVisualElement.Add(absolutePositionElement);
       }
   }
   ```
7. To see the example, from the menu, select **Window** > **UI Toolkit** > **Positioning Test Window**.

## Create the example for the runtime UI

1. Create a USS file named `PositioningTest.uss` with the following content:

   ```
   .box {
       height: 70px;
       width: 70px;
       margin-bottom: 2px;
       background-color: gray;
   }
   #relative{
       width: 70px;
       height: 70px;
       background-color: purple;
       left: 25px;
       margin-bottom: 2px;
       position:relative;
   }
   #absolutePositionElement{
       left: 25px;
       top: 25px;
       width: 70px;
       height: 70px;
       background-color: black;
       position: absolute;
   }
   ```
2. Create a UXML document named `PositioningTest.uxml` with the following content:

   ```
   <ui:UXML xmlns:ui="UnityEngine.UIElements" xmlns:uie="UnityEditor.UIElements"
   xsi="http://www.w3.org/2001/XMLSchema-instance" engine="UnityEngine.UIElements"
   editor="UnityEditor.UIElements" noNamespaceSchemaLocation="../UIElementsSchema/UIElements.xsd"
   editor-extension-mode="False">
       <Style src="PositioningTest.uss"/>
       <ui:VisualElement class="box"/>
       <ui:VisualElement  class="box"/>
       <ui:Label text="Relative\nPos\n25, 0" name="relative" />
       <ui:VisualElement  class="box"/>
       <ui:VisualElement  class="box"/>
       <ui:Label text="Absolute\nPos\n25, 25" name="absolutePositionElement" />
   </ui:UXML>
   ```
3. Create a C# script named `PositioningTestRuntime.cs` with the following content:

   ```
   using UnityEngine;
   using UnityEngine.UIElements;

   public class PostionTestRuntime : MonoBehaviour
   {
       void OnEnable()
       {
           GetComponent<UIDocument>();
       }
      }
   ```
4. Right-click in the Hierarchy window, and then select **UI Toolkit** > **UI Document**.
5. In the Inspector window of the **UI Document**, select **UI Document** > **Source Asset** > **PositioningTest**.
6. In the Inspector window of the **UI Document**, select **Add Component** > **Positioning Test Runtime**.
7. Enter Play mode and adjust the resolution as necessary to see the result.

## Additional resources

* [Panels](UIE-panels.md)
* [Draw order](UIE-draw-order.md)
* [The Layout Engine](UIE-LayoutEngine.md)