# Layout events

[`GeometryChangedEvent`](../ScriptReference/UIElements.GeometryChangedEvent.md) is currently the only layout event.

The base class for `GeometryChangedEvent` is the [EventBase](../ScriptReference/UIElements.EventBase.md) class.

| **Event** | **Description** | **Trickles down** | **Bubbles up** | **Cancellable** |
| --- | --- | --- | --- | --- |
| [GeometryChangedEvent](../ScriptReference/UIElements.GeometryChangedEvent.md) | Sent when the position or the dimensions of an element changes. |  |  |  |

## Unique properties

**`oldRect`**: The former position and dimensions of the element.

**`newRect`**: The new position and dimensions of the element.

## GeometryChangedEvent

The `GeometryChangedEvent` is sent when either the position or the dimensions of an element changes. Events of this type are only sent to the event target.

It’s important to unregister from the `GeometryChangedEvent` event after you receive your callback, as additional layout changes will trigger the callback again.

**`target`**: The element with a new geometry.

## Additional resources

* [Create an aspect ratio custom control](UIE-create-aspect-ratios-custom-control.md)
* [Get resolved styles](UIE-apply-styles-with-csharp.md#get-resolved-styles)