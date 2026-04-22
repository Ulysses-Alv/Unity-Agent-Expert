# Tooltip event

The Tooltip event is sent to check if a **visual element** underneath the pointer is able to display a tooltip. This is an Editor-only event.

Tooltips are usually set using the `tooltip` property. You can also respond to the Tooltip event to set tooltips.

You can handle the Tooltip event in two ways:

1. Set a callback to the [`TooltipEvent`](../ScriptReference/UIElements.TooltipEvent.md). This adds a tooltip to a visual element that doesn’t have one set. This can also override the tooltip set to a visual element.
2. Declare a custom VisualElement (such as declaring a class that extends VisualElement), and override the `ExecuteDefaultAction` method.

If you set the callback or implement a custom visual element to declare tooltips, don’t set the value for the `tooltip` property via code or UXML.

When you set a `tooltip` property, the visual element under the mouse cursor automatically registers a callback to handle the `TooltipEvent`. This callback also stops further propagation of the event.

If you register a custom callback to handle the `TooltipEvent`, you must stop the propagation of the event, or the tooltip can be overridden later in the propagation phase.

The base class for Tooltip events is the [EventBase](../ScriptReference/UIElements.EventBase.md) class.

| **Event** | **Description** | **Trickles down** | **Bubbles up** | **Cancellable** |
| --- | --- | --- | --- | --- |
| [TooltipEvent](../ScriptReference/UIElements.TooltipEvent.md) | Sent just before Unity displays a tooltip. | Yes | Yes | Yes |

## Unique properties

`rect`: Rectangle of the hovered visual element in the panel coordinate system.

`tooltip`: The `tooltip` property is a text string to display inside the tooltip box during the `tooltip` event. The following callback event sets the tooltip property during the event:

```
   evt.tooltip = "Tooltip set by parent!";
```

## Event list

### TooltipEvent

The [`TooltipEvent`](../ScriptReference/UIElements.TooltipEvent.md) is sent just before the Unity Editor displays a tooltip. The handler should set the `TooltipEvent.tooltip` string and the `TooltipEvent.rect`.

**`target`**: The visual element under the mouse.

## Examples

The following examples display the behavior of the `ToolTipEvent`.

To view an example:

1. Under **Assets > Scripts > Editor**, create a C# script called SampleWindow.
2. Copy one of the following examples into the C# script.
3. From the Editor Toolbar, select **Window** > **UI Toolkit** > **SampleWindow**.

### Example 1: Registering a callback to the TooltipEvent on the parent visual element

```
using UnityEditor;
using UnityEngine;
using UnityEngine.UIElements;

public class SampleWindow : EditorWindow
{
   [MenuItem("Window/UI Toolkit/SampleWindow")]
   public static void ShowExample()
   {
       SampleWindow wnd = GetWindow<SampleWindow>();
       wnd.titleContent = new GUIContent("SampleWindow");
   }

   public void CreateGUI()
   {
       VisualElement label = new Label("Hello World! This is a UI Toolkit Label.");
       rootVisualElement.Add(label);

       label.tooltip = "And this is a tooltip";

       // If you comment out the registration of the callback, the tooltip that displays for the label is "And this is a tooltip".
       // If you keep the registration of the callback, the tooltip that displays for the label (and any other child of rootVisualElement)
       // is "Tooltip set by parent!".
       rootVisualElement.RegisterCallback<TooltipEvent>(evt =>
       {
           evt.tooltip = "Tooltip set by parent!";
           evt.rect = (evt.target as VisualElement).worldBound;
           evt.StopPropagation();
       }, TrickleDown.TrickleDown); // Pass the TrickleDown.TrickleDown parameter to intercept the event before it reaches the label.
   }
}
```

### Example 2: Declaring a custom visual element and overriding ExecuteDefaultAction

```
using UnityEditor;
using UnityEngine;
using UnityEngine.UIElements;

public class SampleWindow : EditorWindow
{
   [MenuItem("Window/UI Toolkit/SampleWindow")]
   public static void ShowExample()
   {
       SampleWindow wnd = GetWindow<SampleWindow>();
       wnd.titleContent = new GUIContent("SampleWindow");
   }

   private void CreateGUI()
   {
       CustomLabel custom1 = new CustomLabel("custom 1");
       rootVisualElement.Add(custom1);

       CustomLabel custom2 = new CustomLabel("custom 2");
       rootVisualElement.Add(custom2);
   }
}

public class CustomLabel : Label
{
   private static int m_InstanceCounter = 0;

   private int m_CurrentCounter;

   public CustomLabel(string labelText) : base(labelText)
   {
       m_CurrentCounter = m_InstanceCounter++;
   }

   protected override void ExecuteDefaultAction(EventBase evt)
   {
       // Other events need to be handled as usual.
       base.ExecuteDefaultAction(evt);

       if (evt.eventTypeId == TooltipEvent.TypeId())
       {
           TooltipEvent e = (TooltipEvent)evt;

           // Apply an offset to the tooltip position.
           var tooltipRect = new Rect(worldBound);
           tooltipRect.x += 10;
           tooltipRect.y += 10;
           e.rect = tooltipRect;

           // Set a custom/dynamic tooltip. 
           e.tooltip = $"This is instance # {m_CurrentCounter + 1} of my CustomLabel";

           // Stop propagation avoids other instances of handling of the event that may override the values set here.
           e.StopPropagation();
       }
   }
}
```