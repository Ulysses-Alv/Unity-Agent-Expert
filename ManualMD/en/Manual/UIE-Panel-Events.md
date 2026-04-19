# Panel events

A panel represents a visible instance of a UI hierarchy. It handles element behavior event dispatching within the hierarchy of the **visual tree**. It holds a reference to the root **visual element** of the hierarchy. For runtime UI, it’s comparable to the Canvas in UGUI.

You must attach an instance of a visual element to a panel for it to render or receive events.

Panel events fire on a visual element when its relationship with a panel changes. For example, when you add the visual element to a panel (`AttachToPanelEvent`) or remove it from a panel (`DetachFromPanelEvent`).

Panel events are only sent to visual elements and their descendants within a hierarchy that are directly affected when panel changes occur. Parent elements don’t receive events when descendant visual elements attach or detach from the panel.

For example, in the UXML code below, when you add the `parent` visual element to a hierarchy that’s already attached to a panel, `parent`, `child`, and `grandchild` all receive the same event. If you remove `parent` from the same UXML hierarchy, all visual elements would receive `DetachFromPanel` events.

```
<ui:UXML xmlns:ui="UnityEngine.UIElements" xmlns:uie="UnityEditor.UIElements">
   <ui:VisualElement name="parent">
       <ui:VisualElement name="child">
           <ui:VisualElement name="grandchild" />
       </ui:VisualElement>
   </ui:VisualElement>
</ui:UXML>
```

The base class for all panel events is [PanelChangedEventBase](../ScriptReference/UIElements.PanelChangedEventBase_1.md).

| **Event** | **Description** | **Trickles down** | **Bubbles up** | **Cancellable** |
| --- | --- | --- | --- | --- |
| [AttachToPanelEvent](../ScriptReference/UIElements.AttachToPanelEvent.md) | Sent right after the element (or one of its parents) is attached to a panel. |  |  |  |
| [DetachFromPanelEvent](../ScriptReference/UIElements.DetachFromPanelEvent.md) | Sent just before an element (or one of its parents) is detached from a panel. |  |  |  |

## Unique properties

**`originPanel`**: The `originPanel` contains data specific to the `DetachFromPanelEvent`. It contains the source panel that the visual element detaches from during a panel change.

**`destinationPanel`**: The `destinationPanel` contains data specific to the `AttachToPanelEvent`. It provides the panel where the visual element is now attached.

## Event list

The following list provides the name, description, and target of each event in the event family.

### AttachToPanelEvent

The [`AttachToPanelEvent`](../ScriptReference/UIElements.AttachToPanelEvent.md) fires after a visual element attaches to a panel. It also occurs when you add a visual element to a hierarchy that’s attached to a panel.

**`target`**: The visual element that’s attached to a panel.

### DetachFromPanelEvent

The [`DetachFromPanelEvent`](../ScriptReference/UIElements.DetachFromPanelEvent.md) triggers before you remove a visual element from a panel. It also occurs when you remove a visual element from a hierarchy that’s attached to a panel.

**`target`**:The visual element that’s detaching from a panel.

## Examples

The following example creates an Editor window with a button that will add additional labels to the window. Clicking the labels will remove them again.

This example implements a custom label class that prints a message to the console whenever the instance of the VisualElement is attached or detached from a panel. It highlights the behavior of the AttachToPanelEvent and DetachFromPanelEvent events and how to use the originPanel and destinationPanel properties.

To see the example in action, do the following:

1. Create a Unity project with any template.
2. In the **Project window**, create a folder named `Editor`.
3. In the `Editor` window, create a C# script named `PanelEventsTestWindow.cs` with the following content:

   ```
   using UnityEditor;
   using UnityEngine;
   using UnityEngine.UIElements;
       
   public class PanelEventsTestWindow : EditorWindow
   {
       [MenuItem("Window/UI Toolkit/Panel Events Test Window")]
       public static void ShowExample()
       {
           PanelEventsTestWindow wnd = GetWindow<PanelEventsTestWindow>();
           wnd.titleContent = new GUIContent("Panel Events Test Window");
       }
       
       public void CreateGUI()
       {
           // Set a name for the panel
           rootVisualElement.panel.visualTree.name = "Our Window Root Visual Element";
       
           // Add a button which will add new instances of our custom labels to the window
           rootVisualElement.Add(new Button(() => rootVisualElement.Add(new CustomLabel())) { text = "Add New Label" });
       }
   }
       
   // Custom label class which prints out a console message when it is attached or detached.
   public class CustomLabel : Label
   {
       private static int m_InstanceCounter = 0;
       private int m_LabelNumber;
       
       public CustomLabel() : base()
       {
           m_LabelNumber = ++m_InstanceCounter;
           text = $"Label #{m_LabelNumber} - click me to detach";
           RegisterCallback<AttachToPanelEvent>(evt =>
           {
               Debug.Log($"I am label {m_LabelNumber} and I " +
                   $"just got attached to panel '{evt.destinationPanel.visualTree.name}'");
           });
           RegisterCallback<DetachFromPanelEvent>(evt =>
           {
               Debug.Log($"I am label {m_LabelNumber} and I " +
                   $"just got detached from panel '{evt.originPanel.visualTree.name}'");
           });
           // Register a pointer down callback that removes this element from the hierarchy
           RegisterCallback<PointerDownEvent>(evt => this.RemoveFromHierarchy());
       }
   }
   ```
4. To see the example live, from the menu, select **Window** > **UI Toolkit** > **Panel Events Test Window**.