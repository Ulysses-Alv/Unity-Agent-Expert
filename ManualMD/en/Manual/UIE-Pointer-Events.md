# Pointer events

Pointer events fire for UI interactions with a pointing device. Similar to mouse events, pointer events provide additional information about the used input device, such as pen pressure or tilt angle.

Pointer events always precede mouse events in UI Toolkit.

Pointer events don’t have a persistent position. They also don’t have a set position when they’re released from the touch device.

Some pointer events, such as `PointerStationaryEvent` and `PointerCancelEvent` events, have conditions triggered by the operating system (OS) of the input device.

The base class for all pointer events is [PointerEventBase](../ScriptReference/UIElements.PointerEventBase_1.md).

| **Event** | **Description** | **Trickles down** | **Bubbles up** | **Cancellable** |
| --- | --- | --- | --- | --- |
| [PointerDownEvent](../ScriptReference/UIElements.PointerDownEvent.md) | Sent when you press a pointer. | Yes | Yes | Yes |
| [PointerUpEvent](../ScriptReference/UIElements.PointerUpEvent.md) | Sent when you release a pointer. | Yes | Yes | Yes |
| [PointerMoveEvent](../ScriptReference/UIElements.PointerMoveEvent.md) | Sent when the pointer changes state. | Yes | Yes | Yes |
| [PointerEnterEvent](../ScriptReference/UIElements.PointerEnterEvent.md) | Sent when the pointer enters a **visual element**, or one of its descendants. | Yes |  | Yes |
| [PointerLeaveEvent](../ScriptReference/UIElements.PointerLeaveEvent.md) | Sent when the pointer leaves a visual element and all of its descendants. | Yes |  | Yes |
| [PointerOverEvent](../ScriptReference/UIElements.PointerOverEvent.md) | Sent when the pointer enters a visual element. | Yes | Yes | Yes |
| [PointerOutEvent](../ScriptReference/UIElements.PointerOutEvent.md) | Sent when the pointer leaves a visual element. | Yes | Yes | Yes |
| [PointerStationaryEvent](../ScriptReference/UIElements.PointerStationaryEvent.md) | Sent when a pointer type (like a stylus or finger) doesn’t change for a set amount of time determined by the operating system. | Yes | Yes | Yes |
| [PointerCancelEvent](../ScriptReference/UIElements.PointerCancelEvent.md) | Sent when a pointer action is cancelled by the operating system. | Yes | Yes | Yes |

## Unique properties

**`altitudeAngle`**: The altitudeAngle contains the angle of the stylus relative to the surface, in radians. A value of 0 indicates that the stylus is parallel to the surface. A value of pi/2 indicates that it’s perpendicular to the surface.

**`azimuthAngle`**: The azimuthAngle contains the angle of the stylus relative x-axis, in radians. A value of 0 indicates that the stylus points along the x-axis of the device.

**`button`**: The button property returns an integer that identifies the mouse button pressed to trigger the event. The following table lists the integer and associated mouse button:

| **Integer** | **Button** |
| --- | --- |
| 0 | Left button |
| 1 | Right button |
| 2 | Middle button |

**`clickCount`**: The clickCount property contains the number of times the button is pressed.

**`deltaPosition`**: The `deltaPosition` property contains the difference between the pointer’s position during the previous mouse event and its position during the current mouse event.

**`deltaTime`**: The `deltaTime` property contains the amount of time that has passed since the last recorded change in pointer values, in seconds.

**`localPosition`**: The `localPosition` property returns the pointer position relative to the target visual element.

**`modifiers`**: The modifiers property returns the modifier key currently held down. Some examples of modifiers are the `Shift`, `Ctrl`, or `Alt` keys. For more information, see the [Modifier keys section](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values#modifier_keys) of the MDN documentation.

**`pointerId`**: The pointerId property returns an integer that identifies the pointer that sends the event.

**`pointerType`**: The pointerType property returns a string that defines the type of pointer that creates the event.

**`position`**: The position property returns the pointer position in the screen or world coordinate system.

**`pressedButtons`**: The pressedButton property returns an integer that identifies which combination of mouse buttons are currently pressed.

The number is the sum of the individual buttons’ integer value (see table below). For example, holding the right mouse button and the middle mouse button pressed at the same time will result in pressedButton having a value of 6.

| **Integer** | **Button** |
| --- | --- |
| 1 | Left button |
| 2 | Right button |
| 4 | Middle button |

**`pressure`**: The pressure property returns the amount of pressure currently applied by a touch. If the device doesn’t report pressure, the value of this property is 1.0f.

**`radius`**: The radius property returns an estimate of the radius of a touch. Add radiusVariance to get the maximum touch radius, subtract it to get the minimum touch radius.

**`radiusVariance`**: The radiusVariance property value determines the accuracy of the touch radius. Add this value to the radius to get the maximum touch radius, subtract it to get the minimum touch radius.

**`tangentialPressure`**: The tangentialPressure property returns a float value representing the pressure applied to an additional pressure-sensitive control on the stylus.

**`twist`**: The twist property returns the rotation of the stylus around its axis, in radians.

## Event list

The following list provides the name, description, and target of each event in the event family. For more information on the event, see the [UI Toolkit API](../ScriptReference/UIElements.MouseEventBase_1.md).

### PointerDownEvent

The [PointerDownEvent](../ScriptReference/UIElements.PointerDownEvent.md) is sent when you press a pointer inside a visual element.

**`target`**: The visual element that receives the pointer capture. Otherwise, it’s the topmost selectable element under the cursor.

### PointerUpEvent

A [PointerUpEvent](../ScriptReference/UIElements.PointerUpEvent.md) triggers when you release a pointer within a visual element.

When the `PointerUpEvent` event triggers, it also removes the pointer coordinates. It also clears the cache of the pointer, so there’s no record of the pointer location.

**`target`**: The visual element that receives the pointer capture. Otherwise, it’s the topmost selectable element under the cursor.

### PointerMoveEvent

The [PointerMoveEvent](../ScriptReference/UIElements.PointerMoveEvent.md) occurs when the pointer changes state.

**`target`**: The visual element that receives the pointer capture. Otherwise, it’s the topmost selectable element under the cursor.

### PointerEnterEvent

The [PointerEnterEvent](../ScriptReference/UIElements.PointerEnterEvent.md) is sent when the pointer enters a visual element, or one of its descendants.

**`target`**: The visual element (or one of its descendants) that the pointer exits.

### PointerLeaveEvent

A [PointerLeaveEvent](../ScriptReference/UIElements.PointerLeaveEvent.md) is sent when the pointer leaves a visual element and all its descendants. For example if a visual element contains a child, then the parent element will receive this event when the pointer is no longer over either the parent or the child. The parent element won’t receive the [PointerLeaveEvent](../ScriptReference/UIElements.PointerLeaveEvent.md) while the pointer is still over one of its child elements, even if it’s no longer the topmost element underneath the pointer. It will receive the [PointerOverEvent](../ScriptReference/UIElements.PointerOverEvent.md) instead.

**`target`**: The visual element (or one of its descendants) that the pointer exits.

### PointerOverEvent

The [PointerOverEvent](../ScriptReference/UIElements.PointerOverEvent.md) is sent when the pointer enters a visual element.

**`target`**: The visual element under the pointer.

### PointerOutEvent

The [PointerOutEvent](../ScriptReference/UIElements.PointerOutEvent.md) is sent when the pointer leaves a visual element.

**`target`**: The visual element that the pointer exits.

### PointerStationaryEvent

The [PointerStationaryEvent](../ScriptReference/UIElements.PointerStationaryEvent.md) is sent when a pointer type (like a stylus or finger) doesn’t change for a set amount of time determined by the operating system.

**`target`**: The visual element that captures the pointer, or the topmost selectable element under the pointer.

### PointerCancelEvent

The [PointerCancelEvent](../ScriptReference/UIElements.PointerCancelEvent.md) is sent when a pointer action is cancelled by the operating system.

**`target`**: The visual element that captures the pointer, or the topmost selectable element under the pointer.

## Examples

The following code sample creates an Editor window with a red box containing a yellow box. It prints messages to the console when the pointer leaves a visual element or its child. It demonstrates the behavior of the [PointerOutEvent](../ScriptReference/UIElements.PointerOutEvent.md) and the [PointerLeaveEvent](../ScriptReference/UIElements.PointerLeaveEvent.md).

To see the example in action, do the following:

1. Under **Assets > Scripts > Editor**, create a new UXML file called PointerEventsTestWindow.uxml
2. Copy the UXML code from below into it
3. Under **Assets > Scripts > Editor**, create a new C# file called PointerEventsTestWindow.cs
4. Copy a code sample into the C# script.
5. From the Editor Toolbar, select **Window** > **UI Toolkit** > **Pointer Events Test Window**.

### UXML code

```
<ui:UXML xmlns:ui="UnityEngine.UIElements" xmlns:uie="UnityEditor.UIElements" xsi="http://www.w3.org/2001/XMLSchema-instance" engine="UnityEngine.UIElements" editor="UnityEditor.UIElements" noNamespaceSchemaLocation="../../UIElementsSchema/UIElements.xsd" editor-extension-mode="False">
    <ui:VisualElement style="flex-grow: 1; justify-content: center; align-items: center;">
        <ui:VisualElement name="Red_Box" style="background-color: rgb(183, 34, 46); width: 50%; height: 50%; align-items: center; justify-content: center;">
            <ui:VisualElement name="Yellow_Box" style="width: 175%; height: 50%; background-color: rgb(197, 163, 0);" />
        </ui:VisualElement>
    </ui:VisualElement>
</ui:UXML>
```

### C# code

```
using UnityEditor;
using UnityEngine;
using UnityEngine.UIElements;
using UnityEditor.UIElements;

public class PointerEventsTestWindow : EditorWindow
{
    [MenuItem("Window/UI Toolkit/Pointer Events Test Window")]
    public static void ShowExample()
    {
        PointerEventsTestWindow wnd = GetWindow<PointerEventsTestWindow>();
        wnd.titleContent = new GUIContent("Pointer Events Test Window");
    }

    public void CreateGUI()
    {
        // Import UXML
        VisualTreeAsset visualTree = AssetDatabase.LoadAssetAtPath<VisualTreeAsset>("Assets/Scripts/Editor/PointerEventsTestWindow.uxml");
        visualTree.CloneTree(rootVisualElement);

        // Get the red box and register pointer event callbacks
        VisualElement redBox = rootVisualElement.Q("Red_Box");
        redBox.RegisterCallback<PointerOutEvent>(OnPointerOutEvent, TrickleDown.TrickleDown);
        redBox.RegisterCallback<PointerLeaveEvent>(OnPointerLeaveEvent, TrickleDown.TrickleDown);
    }

    private void OnPointerLeaveEvent(PointerLeaveEvent evt)
    {
        Debug.Log($"Pointer LEAVE Event. Target: {(evt.target as VisualElement).name}");
    }

    private void OnPointerOutEvent(PointerOutEvent evt)
    {
        Debug.Log($"Pointer OUT Event. Target: {(evt.target as VisualElement).name}");
    }
}
```

## Additional resources

* [Create a drag-and-drop UI](UIE-create-drag-and-drop-ui.md)