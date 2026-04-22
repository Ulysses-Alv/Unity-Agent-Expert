# Navigation events

Navigation events occur at runtime when the user presses the D-pad, moves a joystick, or presses the `Escape`, `Enter` or arrow keys. They’re an indicator that the user is trying to navigate the UI, but they’re not limited to a specific input device, such as a keyboard. The difference from a [Focus](UIE-Focus-Events.md) event is that the navigation event doesn’t require the focus to move to a new UI element.

The base class for all navigation events is [NavigationEventBase](../ScriptReference/UIElements.NavigationEventBase_1.md).

All navigation events trickle down, bubble up, and are cancellable, but it’s recommended to listen to these events during the bubble-up propagation phase. This is because navigation events are triggered by input events that might also be used to interact with individual controls. For example, a button will react to a press of the `Enter` key with a button click, and it will cancel the `NavigationSubmitEvent`. Listening to these events during the bubble-up phase makes sure that they’re navigation events.

| **Event** | **Description** | **Trickles down** | **Bubbles up** | **Cancellable** |
| --- | --- | --- | --- | --- |
| [NavigationMoveEvent](../ScriptReference/UIElements.NavigationMoveEvent.md) | Sent when user makes a move input. | ✔ | ✔ | ✔ |
| [NavigationCancelEvent](../ScriptReference/UIElements.NavigationCancelEvent.md) | Sent when the user makes a cancel input. | ✔ | ✔ | ✔ |
| [NavigationSubmitEvent](../ScriptReference/UIElements.NavigationSubmitEvent.md) | Sent when the user makes a submit input. | ✔ | ✔ | ✔ |

## Event list

### NavigationMoveEvent

The `NavigationMoveEvent` is sent when the user presses the D-pad, moves a joystick, or presses the arrow keys.

Some controls will use the arrow keys for their own functionality. For example, the [ListView](UIE-uxml-element-ListView.md) allows the user to select items using the up and down arrow keys. In this case, the control will cancel the `NavigationMoveEvent` and the event won’t enter the bubble-up phase.

**`direction`**: The direction of the navigation. (`None`, `Left`, `Up`, `Right`, `Down`)

**`move`**: The move vector. If the event was triggered by an analog axis input, like a joystick, this property gives access to the directional vector.

### NavigationCancelEvent

The `NavigationCancelEvent` is triggered when the user cancels the current navigation action, such as by pressing the `Escape`` key on the keyboard. It’s important to note that this event doesn’t affect the currently focused element, which means that the UI element that had focus before the cancellation remains selected.

### NavigationSubmitEvent

The `NavigationSubmitEvent` is triggered when the user presses the submit button, such as by pressing the `Enter`` key on the keyboard.

If a control handles the event itself, it cancels the event, preventing it from entering the bubble-up phase. For example, a [TextField](UIE-uxml-element-TextField.md) that has the focus will stop the `NavigationSubmitEvent` from bubbling up. On the other hand, a focusable [Label](UIE-uxml-element-Label.md) or [Image](UIE-uxml-element-Image.md) will allow the `NavigationSubmitEvent` to bubble up to its parent elements, allowing them to handle the event if needed.

## Example

The following code example shows how to register callbacks for navigation events in a runtime UI:

```
using UnityEngine;
using UnityEngine.UIElements;

public class MyNavigationHandler: MonoBehaviour {
    void OnEnable() {
        // Get a reference to the root visual element
        var uiDocument = GetComponent < UIDocument > ();
        var rootVisualElement = uiDocument.rootVisualElement;

        // Register for navigation events
        rootVisualElement.RegisterCallback < NavigationCancelEvent > (OnNavCancelEvent);
        rootVisualElement.RegisterCallback < NavigationMoveEvent > (OnNavMoveEvent);
        rootVisualElement.RegisterCallback < NavigationSubmitEvent > (OnNavSubmitEvent);
    }

    private void OnNavSubmitEvent(NavigationSubmitEvent evt) {
        Debug.Log($"OnNavSubmitEvent {evt.propagationPhase}");
    }

    private void OnNavMoveEvent(NavigationMoveEvent evt) {
        Debug.Log($"OnNavMoveEvent {evt.propagationPhase} - move {evt.move} - direction {evt.direction}");
    }

    private void OnNavCancelEvent(NavigationCancelEvent evt) {
        Debug.Log($"OnNavCancelEvent {evt.propagationPhase}");
    }
}
```

## Additional resources

* [Dispatch events](UIE-Events-Dispatching.md)
* [Handle events](UIE-Events-Handling.md)
* [Synthesize and send events](UIE-Events-Synthesizing.md)