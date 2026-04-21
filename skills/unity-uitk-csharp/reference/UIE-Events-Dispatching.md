# Dispatch events

The **event system** listens for events that come from the operating system or **scripts**, then uses the [EventDispatcher](../ScriptReference/UIElements.EventDispatcher.md) to dispatch those events to **visual elements**. The event dispatcher determines an appropriate dispatching strategy for each event it sends. Once determined, the dispatcher executes the strategy.

Visual elements implement default behaviors for several events. This involves the creation and execution of additional events. For example, a [`PointerMoveEvent`](../ScriptReference/UIElements.PointerMoveEvent.md) can generate an additional [`PointerEnterEvent`](../ScriptReference/UIElements.PointerEnterEvent.md) and a [`PointerLeaveEvent`](../ScriptReference/UIElements.PointerLeaveEvent.md). These events enter a queue and process after the current event. For example, the `PointerMoveEvent` finishes processing before the `PointerEnterEvent` and `PointerLeaveEvent` events.

## Dispatch behavior of event types

Each event type has its own dispatch behavior. The behavior of each event type breaks down into two stages:

* **Trickles down**: Events sent to elements during the trickle-down phase.
* **Bubbles up**: Events sent to elements during the bubble-up phase.

For a list of dispatch behavior for each event type, see the [Event reference page](UIE-Events-Reference.md).

## Event propagation

After the event dispatcher selects the event target, it computes the propagation path of the event. The propagation path is an ordered list of visual elements that receive the event. The propagation path occurs in the following order:

1. The path starts at the root of the visual element tree and descends towards the target. This is the **trickle-down phase**.
2. The event target receives the event.
3. The event then ascends the tree towards the root. This is the **bubble-up phase**.

![Propagation path](../uploads/Main/UIElementsEvents.png)

Propagation path

Most event types are sent to all elements along the propagation path. Some event types skip the bubble-up phase, and some event types are sent to the event target only.

If you hide or disable an element, it won’t receive events. Events still propagate to the ancestors and descendants of a hidden or disabled element.

## Event target

As an event travels along the propagation path, `Event.currentTarget` updates to the element handling the event. Within an event callback function, there are two properties that log the dispatch behavior:

* [`EventBase.currentTarget`](../ScriptReference/UIElements.EventBase-currentTarget.md) is the visual element on which the callback was registered.
* [`EventBase.target`](../ScriptReference/UIElements.EventBase-target.md) is the element where the event occurs, for example, the element directly under the pointer.

The target of an event depends on the event type. For pointer events, the target is most commonly the topmost pickable element, directly under the pointer. For keyboard events, the target is the element that has focus.

UI Toolkit events have a `target` property that contains a reference to the element where the event occurred. For most events that originate from the operating system, the dispatch process finds the event target automatically.

The target element is stored in `EventBase.target` and doesn’t change during the dispatch process. The property `Event.currentTarget` updates to the visual element currently handling the event.

## Picking mode and custom shapes

Most pointer events use the picking mode to decide their target. The `VisualElement` class has a `pickingMode` property that supports the following values:

* [`PickingMode.Position`](../ScriptReference/UIElements.PickingMode.Position.md) (default): performs picking based on the position rectangle.
* [`PickingMode.Ignore`](../ScriptReference/UIElements.PickingMode.Ignore.md): prevents picking as the result of a pointer event.

You can override the [`VisualElement.ContainsPoint()`](../ScriptReference/UIElements.VisualElement.ContainsPoint.md) method to perform custom intersection logic.

## Additional resources

* [Handle event callbacks and value changes](UIE-Events-Handling.md)
* [Synthesize and send events](UIE-Events-Synthesizing.md)
* [Events reference](UIE-Events-Reference.md)