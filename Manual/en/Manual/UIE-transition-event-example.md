## Create a transition event

**Version**: 2021.3+

This example demonstrates the lifecycle of a transition event.

## Example overview

The example creates a custom Editor window with a button and color palette. If you click the button, the following appears:

* A transition starts to change the color palette from blue to green.
* The transition events that are sent during the transition.
* The duration of the events.

You can find the completed files that this example creates in this [GitHub repository](https://github.com/Unity-Technologies/ui-toolkit-manual-code-examples/tree/master/transition-events-example).

## Prerequisites

This guide is for developers familiar with the Unity Editor, UI Toolkit, and C# scripting. Before you start, get familiar with the following:

* [UXML](UIE-UXML.md)
* [USS transition](UIE-Transitions.md)
* [Transition events](UIE-Transition-Events.md)

## Create the example

1. Create a project in Unity with any template.
2. From the menu, select **Assets** > **Create** > **UI Toolkit** > **Editor Window**.
3. In the **UI Toolkit Editor Window Creator** window, enter `TransitionExample`.
4. Save your changes. This creates three files as `TransitionExample.cs`, `TransitionExample.uss`, and `TransitionExample.uxml`.
5. Replace `TransitionExample.cs` with the following content:

   ```
   using System;
   using UnityEditor;
   using UnityEngine;
   using UnityEngine.UIElements;

   public class TransitionExample : EditorWindow
   {
       [SerializeField] private VisualTreeAsset m_VisualTreeAsset = default;

       private Button clickMeButton;
       private VisualElement colorChanger;
       private Label eventLabel;
       private Label timeLabel;

       private DateTime lastEvent;
       private static readonly TimeSpan NearlyInstantaneousThreshold = TimeSpan.FromMilliseconds(10);

       private static readonly string ClickMeButtonClass = "click-me";
       private static readonly string ColorChangerClass = "color-changer";
       private static readonly string ColorChangerTransitionClass = "color-transition";
       private static readonly string EventLabelName = "eventLabel";
       private static readonly string TimeLabelName = "timeLabel";
       private static readonly string TimeBelowThresholdText = "Almost instantaneous.";

       [MenuItem("Window/UI Toolkit/Transition Example")]
       public static void ShowExample()
       {
           TransitionExample wnd = GetWindow<TransitionExample>();
           wnd.titleContent = new GUIContent("TransitionExample");
           wnd.minSize = new Vector2(500f, 400f);
       }

       public void CreateGUI()
       {
           lastEvent = DateTime.Now;

           // Each editor window contains a root VisualElement object
           VisualElement root = rootVisualElement;

           // Instantiate UXML
           VisualElement uxmlAsset = m_VisualTreeAsset.Instantiate();
           root.Add(uxmlAsset);

           // Get the relevant elements by querying the root element
           clickMeButton = root.Q<Button>(className: ClickMeButtonClass);

           colorChanger = root.Q<VisualElement>(className: ColorChangerClass);

           eventLabel = root.Q<Label>(name: EventLabelName);

           timeLabel = root.Q<Label>(name: TimeLabelName);

           // Add callbacks for clicking on the button and monitoring the color changing element.
           clickMeButton.RegisterCallback<ClickEvent>(OnClickEvent);

           colorChanger.RegisterCallback<TransitionRunEvent>(OnTransitionRun);
           colorChanger.RegisterCallback<TransitionStartEvent>(OnTransitionStart);
           colorChanger.RegisterCallback<TransitionEndEvent>(OnTransitionEnd);
           colorChanger.RegisterCallback<TransitionCancelEvent>(OnTransitionCancel);
       }

       private void OnDisable()
       {
           clickMeButton.UnregisterCallback<ClickEvent>(OnClickEvent);

           colorChanger.UnregisterCallback<TransitionRunEvent>(OnTransitionRun);
           colorChanger.UnregisterCallback<TransitionStartEvent>(OnTransitionStart);
           colorChanger.UnregisterCallback<TransitionEndEvent>(OnTransitionEnd);
           colorChanger.UnregisterCallback<TransitionCancelEvent>(OnTransitionCancel);
       }

       private void OnClickEvent(ClickEvent evt)
       {
           colorChanger.ToggleInClassList(ColorChangerTransitionClass);
       }

       private void OnTransitionRun(TransitionRunEvent evt)
       {
           DisplayLatestEvent("TransitionRunEvent", DateTime.Now);
       }

       private void OnTransitionStart(TransitionStartEvent evt)
       {
           DisplayLatestEvent("TransitionStartEvent", DateTime.Now);
       }

       private void OnTransitionEnd(TransitionEndEvent evt)
       {
           DisplayLatestEvent("TransitionEndEvent", DateTime.Now);
       }

       private void OnTransitionCancel(TransitionCancelEvent evt)
       {
           DisplayLatestEvent("TransitionCancelEvent", DateTime.Now);
       }

       private void DisplayLatestEvent(string eventType, DateTime timestamp)
       {
           // If two events are sent too close together, add both to the Latest event line.
           // This happens if the delay is set to 0 and the TransitionRun and TransitionStart
           // are sent at the same time, or if the button was pressed before the transition
           // was over, thus sending TransitionCancel and TransitionRun (and potentially
           // TransitionStart) events close together.
           var elapsed = timestamp - lastEvent;
           if (elapsed <= NearlyInstantaneousThreshold)
           {
               timeLabel.text = TimeBelowThresholdText;
               eventLabel.text += eventType;
           }
           else
           {
               timeLabel.text = $"{elapsed:s\\.ff} s";
               eventLabel.text = eventType;
           }

           lastEvent = timestamp;
       }
   }
   ```
6. Replace `TransitionExample.uxml` with the following content:

   ```
   <ui:UXML xmlns:ui="UnityEngine.UIElements" xmlns:uie="UnityEditor.UIElements"
           xsi="http://www.w3.org/2001/XMLSchema-instance" engine="UnityEngine.UIElements" editor="UnityEditor.UIElements"
           noNamespaceSchemaLocation="../../UIElementsSchema/UIElements.xsd" editor-extension-mode="False">
       <Style src="TransitionExample.uss"/>
       <ui:VisualElement class="main-container">
           <ui:Button text="Click Me!" class="click-me"/>
           <ui:VisualElement class="color-changer"/>
           <ui:VisualElement class="label-section">
               <ui:VisualElement class="label-line">
                   <ui:Label text="Latest event(s) :  "/>
                   <ui:Label name="eventLabel"/>
               </ui:VisualElement>
               <ui:VisualElement class="label-line">
                   <ui:Label text="Time since last event :  "/>
                   <ui:Label name="timeLabel"/>
               </ui:VisualElement>
           </ui:VisualElement>
       </ui:VisualElement>
   </ui:UXML>
   ```
7. Replace `TransitionExample.uss` with the following content:

   ```
   .click-me {
       width: 250px;
       height: 50px;
       font-size: 40px;
       -unity-font-style: bold-and-italic;
       margin: 30px;
   }

   .color-changer {
       margin: 10px;
       width: 150px;
       height: 150px;
       border-width: 10px;
       border-radius: 75px;
       background-color: rgb(0, 31, 138);
       transition: background-color 3s ease-in-out 1s;
   }

   .main-container {
       align-items: center;
       justify-content: space-between;
       flex-grow: 1;
       background-color: rgb(60, 60, 60);
   }

   .label-section {
       margin: 10px;
       border-width: 2px;
       width: 95%;
       align-items: center;
   }

   .label-line {
       flex-direction: row;
       margin: 5px;
       flex-grow: 1;
       align-items: center;
       width: 90%;
       height: 25px;
       font-size: 14px;
       padding: 0;
   }

   .color-transition {
       background-color: rgb(177, 221, 111);
   }
   ```
8. Select `TransitionExample.cs` and drag `TransitionExample.uxml` to the ****Visual Tree** Asset** field in the **Inspector** if it’s not already there.
9. To try the example, from the menu, select **Window** > **UI Toolkit** > **Transition Example**.
10. Click the button to view the color transition and the description of the events sent. There is a long delay between the `TransitionRunEvent` and the `TransitionStartEvent`. During the long duration, you can click the button again to interrupt the transition.

## Additional resources

* [USS](UIE-USS.md)
* [Handle events](UIE-Events-Handling.md)
* [Create looping transitions](UIE-transition-event-loop-example.md)
* [Create a simple transition with UI Builder and C# scripts](UIE-transition-example.md)