# Respond to events with custom controls

If you’re implementing [custom controls](UIE-custom-controls.md), you can respond to UI Toolkit events in the following ways:

* [Register an event callback](UIE-Events-Handling.md#register-an-event-callback)
* [Override the HandleEventBubbleUp or HandleEventTrickleDown virtual methods](#override-a-handle-event-virtual-method)

Your response to events depends on the situation. The following are the differences between callbacks and virtual method overrides:

* Callbacks must register on instances of the class. Virtual methods require modifying the class itself.
* Virtual method overrides have a slight performance advantage because they don’t require a lookup in the callback registry during the propagation phase.

## Override the HandleEventBubbleUp or HandleEventTrickleDown virtual methods

A virtual method override applies to all instances of the class. For a class that overrides [`HandleEventBubbleUp`](../ScriptReference/UIElements.CallbackEventHandler.HandleEventBubbleUp.md) or [`HandleEventTrickleDown`](../ScriptReference/UIElements.CallbackEventHandler.HandleEventTrickleDown.md), you can also register callbacks on its instances.

To override the `HandleEventBubbleUp` or `HandleEventTrickleDown` methods, or both, derive a new subclass of `VisualElement`.

`HandleEventBubbleUp` and `HandleEventTrickleDown` execute on each instance of a **visual element** subclass when that instance receives an event.

The following example shows how to customize those virtual methods:

```
override void HandleEventBubbleUp(EventBase evt)
{
    // Call the base function.
    base.HandleEventBubbleUp(evt);

    if (evt.eventTypeId == PointerDownEvent.TypeId())
    {
        // ...
    }
    else if (evt.eventTypeId == MouseUpEvent.TypeId())
    {
        // ...
    }
    // More event types
}
```

For a given class instance, executing custom code in the following cases have the same results:

* In a callback set for the BubbleUp phase
* In the HandleEventBubbleUp method override

In either case, if you stop the propagation of the event, it prevents reactions to the event after you have executed the current target callbacks and method overrides.

## Best practices

The following are best practices for handing events with custom controls.

### Implement behaviors

In general, to implement behaviors from your element, use a `HandleEventBubbleUp` method override.

Given that the BubbleUp is the default propagation phase for callbacks, you can move any code from callbacks to a `HandleEventBubbleUp` method without any concerns about changing the timing of code execution.

The benefits of implementing behaviors as method overrides include:

* Method overrides don’t require a lookup in the callback registry.
* Instances without callbacks don’t enter the propagation process.

### Stop event propagation

When handling an event inside a callback or a virtual method override, you can stop further event propagation by calling one of the StopPropagation methods on the event. For example, a parent panel might stop propagation during the trickle-down phase to prevent its children from receiving events.

You can’t prevent the execution of the [`EventBase.PreDispatch()`](../ScriptReference/UIElements.EventBase.PreDispatch.md) and [`EventBase.PostDispatch()`](../ScriptReference/UIElements.EventBase.PostDispatch.md) methods inside the event class itself.

The following methods affect event propagation:

* [`StopImmediatePropagation()`](../ScriptReference/UIElements.EventBase.StopImmediatePropagation.md): Stops the event propagation process immediately to prevent any subsequent callbacks from executing for the event.
* [`StopPropagation()`](../ScriptReference/UIElements.EventBase.StopPropagation.md): Stops the event propagation process after the last callback on the current element. This ensures that all callbacks execute on the current element, while preventing any further elements from responding to the event.

## Additional resources

* [HandleEventBubbleUp](../ScriptReference/UIElements.CallbackEventHandler.HandleEventBubbleUp.md)
* [HandleEventTrickleDown](../ScriptReference/UIElements.CallbackEventHandler.HandleEventTrickleDown.md)
* [EventBase.PreDispatch()](../ScriptReference/UIElements.EventBase.PreDispatch.md)
* [EventBase.PostDispatch()](../ScriptReference/UIElements.EventBase.PostDispatch.md)
* [StopImmediatePropagation()](../ScriptReference/UIElements.EventBase.StopImmediatePropagation.md)
* [StopPropagation()](../ScriptReference/UIElements.EventBase.StopPropagation.md)