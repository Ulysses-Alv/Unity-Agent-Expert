# Synthesize and send events

Before you synthesize and send custom events, understand [how the UI Toolkit event system allocates and sends operating system events](UIE-Events-Dispatching.md).

UI Toolkit sends events to **visual elements** through the panel. If an event has no target, it’s sent to the root element of the panel. To have a propagation path, an element must have a target, and the sender must set that target in advance. Some event types don’t need a target. For example, keyboard events are sent to the focused element, and pointer events are sent to the element under the pointer.

The event system uses a pool of events to avoid allocating event objects repeatedly.

To synthesize and send your own events:

1. Get an event object from the pool of events.
2. Fill in the event properties.
3. Enclose the event in a `using` block to ensure it’s returned to the event pool.
4. Pass the event to `panel.visualTree.SendEvent()`.

You can send operating system events, such as keyboard and pointer events. To do so, use a `UnityEngine.Event` to initialize the UI Toolkit event.

The following example demonstrates how to synthesize and send events:

```
void SynthesizeAndSendKeyDownEvent(IPanel panel, KeyCode code,
     char character = '\0', EventModifiers modifiers = EventModifiers.None)
{
    // Create a UnityEngine.Event to hold initialization data.
    var evt = new Event() {
        type = EventType.KeyDownEvent,
        keyCode = code,
        character = character,
        modifiers = modifiers
    };

    using (KeyDownEvent keyDownEvent = KeyDownEvent.GetPooled(evt))
    {
        panel.visualTree.SendEvent(keyDownEvent);
    }
}
```

**Important**: Don’t send events that are from outside the operating system or aren’t present in the `UnityEngine.Event` types. UI Toolkit sends some events as a reaction to internal state changes. External processes must not send those events. For example, if you send [`PointerCaptureEvent`](../ScriptReference/UIElements.PointerCaptureEvent.md), visual elements assume that the underlying conditions for that event are met and won’t set pointer capture for them. This might break the internal configurations of the visual element and cause undefined behaviors.

## Additional resources

* [Events reference](UIE-Events-Reference.md)
* [Handle event callbacks and value changes](UIE-Events-Handling.md)
* [Dispatch events](UIE-Events-Dispatching.md)
* [Manipulators](UIE-manipulators.md)