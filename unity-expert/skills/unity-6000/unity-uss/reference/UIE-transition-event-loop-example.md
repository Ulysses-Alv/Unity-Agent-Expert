# Create looping transitions

Version: 2022.1+

This example demonstrates how to leverage the `TransitionEndEvent` to create transitions that loop.

## Example overview

This example demonstrates two looping animations:

* Yo-yo: The loop starts from state A to state B with transition, and then return back to state A with transition.
* A-to-B: The loop starts from state A to state B with transition, and then return back to state A without transition.

![Window shows a yo-yo and a A-to-B loop transition effects](../uploads/Main/uxml/loop-transition.gif)

Window shows a yo-yo and a A-to-B loop transition effects

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/master/loop-transition-example).

## Prerequisites

This guide is for developers familiar with the Unity Editor, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [UXML](UIE-UXML.md)
* [USS transition](UIE-Transitions.md)
* [Transition events](UIE-Transition-Events.md)

## Create the example

1. Create a project in Unity with any template.
2. In your **Project window**, create a folder named `loop-transition-example`.
3. Right-click in the folder, and select **Create** > **UI Toolkit** > **Editor Window**.
4. In the **UI Toolkit Editor Window Creator** window, enter `LoopingExample`.
5. Save your changes. This creates three files as `LoopingExample.cs`, `LoopingExample.uss`, and `LoopingExample.uxml`.
6. Replace `LoopingExample.cs` with the following content:

   ```
   using UnityEditor;
   using UnityEngine;
   using UnityEngine.UIElements;
   public class LoopingExample : EditorWindow
   {
       [SerializeField] private VisualTreeAsset m_VisualTreeAsset = default;
       private Label _yoyoLabel;
       private Label _a2bLabel;
       [MenuItem("Window/UI Toolkit/Transition Looping Example")]
       public static void ShowExample()
       {
           var wnd = GetWindow<LoopingExample>();
           wnd.titleContent = new GUIContent("TransitionStyle");
       }
       public void CreateGUI()
       {
           VisualElement root = rootVisualElement;
           VisualElement asset = m_VisualTreeAsset.Instantiate();
           root.Add(asset);
           SetupYoyo(root);
           SetupA2B(root);
       }
       // This method powers the yo-yo loop.
       private void SetupYoyo(VisualElement root)
       {
           _yoyoLabel = root.Q<Label>(name: "yoyo-label");
           // When the animation ends, the callback toggles a class to set the scale to 1.3 
           // or back to 1.0 when it's removed.
           _yoyoLabel.RegisterCallback<TransitionEndEvent>(evt => _yoyoLabel.ToggleInClassList("enlarge-scale-yoyo"));
           // Schedule the first transition 100 milliseconds after the root.schedule.Execute method is called.
           root.schedule.Execute(() => _yoyoLabel.ToggleInClassList("enlarge-scale-yoyo")).StartingIn(100);
       }
       // This method powers the A-to-B cycle.
       private void SetupA2B(VisualElement root)
       {
           _a2bLabel = root.Q<Label>(name:"a2b-label");
           _a2bLabel.RegisterCallback<TransitionEndEvent>(evt =>
           {
               _a2bLabel.RemoveFromClassList("enlarge-scale-a2b");
               _a2bLabel.schedule.Execute(() => _a2bLabel.AddToClassList("enlarge-scale-a2b")).StartingIn(10);
           });
           _a2bLabel.schedule.Execute(() => _a2bLabel.AddToClassList("enlarge-scale-a2b")).StartingIn(100);
       }
   }
   ```
7. Replace `LoopingExample.uxml` with the following content:

   ```
   <ui:UXML xmlns:ui="UnityEngine.UIElements" xmlns:uie="UnityEditor.UIElements" 
            xsi="http://www.w3.org/2001/XMLSchema-instance" 
            engine="UnityEngine.UIElements" editor="UnityEditor.UIElements" 
            noNamespaceSchemaLocation="../../../../UIElementsSchema/UIElements.xsd" 
            editor-extension-mode="False">
       <Style src="LoopingExample.uss" />
       <ui:VisualElement name="container">
           <ui:VisualElement>
                   <ui:Label text="Yo-yo Transition" name="yoyo-label" class="text-style" />
           </ui:VisualElement>
           <ui:VisualElement>
                   <ui:Label text="A-to-B Transition" name="a2b-label" class="text-style"/>
           </ui:VisualElement>
       </ui:VisualElement>
   </ui:UXML>
   ```
8. Replace `LoopingExample.uss` with the following content:

   ```
       #yoyo-label{
           transition-duration: 3s;
       }

       .text-style {
           font-size: 20px;
           flex-grow: 0;
           margin: 20px; 
       }

       .enlarge-scale-a2b{
           scale: 1.5 1.5;
           transition-duration: 3s;
       }
           
       .enlarge-scale-yoyo{
           scale: 1.5 1.5;
       }

       #container{
           flex-grow:1;
           justify-content: space-around;
           align-items: center;
       }
   ```
9. To test the example, from the menu, select **Window** -> **UI Toolkit** -> **Transition Looping Example**.

## Additional resources

* [USS](UIE-USS.md)
* [Handle events](UIE-Events-Handling.md)
* [Create a transition event](UIE-transition-event-example.md)
* [Create a simple transition with UI Builder and C# scripts](UIE-transition-example.md)