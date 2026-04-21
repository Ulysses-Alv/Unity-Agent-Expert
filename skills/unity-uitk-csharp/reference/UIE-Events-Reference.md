# Event reference

UI Toolkit raises an event when a user interacts with and changes the state of elements from UI Toolkit.
The event design is similar to the [Event interface](https://developer.mozilla.org/en-US/docs/Web/API/Event) for HTML elements.

Event types fit into a hierarchy based on the [EventBase](../ScriptReference/UIElements.EventBase.md) class. Each event family implements an interface that defines the common characteristics of all events of the same family. For example, `BlurEvent` and `FocusEvent` use the [FocusEventBase](../ScriptReference/UIElements.FocusEventBase_1.md) class.

Select any of the event types listed below for more information on the event, its parent class, and links to the API documentation.

| Topic | Description |
| --- | --- |
| [Capture events](UIE-Capture-Events.md) | Events that capture the user’s interaction with the UI. |
| [Change events](UIE-Change-Events.md) | Events that occur when the user changes the state of an element. |
| [Click events](UIE-Click-Events.md) | Events that occur when the user clicks an element. |
| [Command events](UIE-Command-Events.md) | Events that occur when the user invokes a command. |
| [Drag and drop events](UIE-Drag-Events.md) | Events that occur when the user drags and drops an element. |
| [Layout events](UIE-Layout-Events.md) | Events that occur when the layout engine changes the layout of an element. |
| [Focus events](UIE-Focus-Events.md) | Events that occur when the user focuses on an element. |
| [Input events](UIE-Input-Events.md) | Events that occur when the user inputs text. |
| [Keyboard events](UIE-Keyboard-Events.md) | Events that occur when the user presses a key. |
| [Mouse events](UIE-Mouse-Events.md) | Events that occur when the user moves the mouse. |
| [Navigation events](UIE-Navigation-Events.md) | Events that occur when the user navigates through the UI. |
| [Panel events](UIE-Panel-Events.md) | Events that occur when the user interacts with a panel. |
| [Pointer events](UIE-Pointer-Events.md) | Events that occur when the user interacts with a pointer device. |
| [Tooltip events](UIE-Tooltip-Events.md) | Events that occur when the user interacts with a tooltip. |
| [Transition events](UIE-Transition-Events.md) | Events that inform you of changes to the state of transition. |
| [ContextualMenu events](UIE-contextual-menus.md) | Events that occur when the user interacts with a contextual menu. |
| [IMGUI events](UIE-IMGUI-Events.md) | Events that occur when the user interacts with an IMGUI element. |

## Additional resources

* [Dispatch events](UIE-Events-Dispatching.md)
* [Handle events](UIE-Events-Handling.md)
* [Synthesize and send events](UIE-Events-Synthesizing.md)