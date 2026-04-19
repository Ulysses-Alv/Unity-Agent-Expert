# Manipulators

To separate your event logic from your UI code, use a manipulator to handle events. Manipulators are **state machines** that handle user interaction with UI elements. They store, register, and unregister event callbacks. A manipulator streamlines setting up that user interaction, so you don’t have to handle each callback one by one. To handle events, use or inherit from one of the manipulators that UI Toolkit supports.

To create and use a manipulator:

1. Define a dedicated class that inherits from a manipulator class supported by UI Toolkit. This class encapsulates the event-handling logic tailored to the specific user interaction you want to manage.
2. Within the class, implement methods to respond to relevant interactions, such as mouse clicks or drags. These methods capture and process the necessary information to execute the desired behavior.
3. When you have finished designing the manipulator class, instantiate it and attach it to the target UI element. This attachment lets the manipulator intercept and manage the specified events, orchestrating user interactions while maintaining clear separation from your UI code.

## Supported manipulators

The following table lists the supported manipulator classes:

| **Manipulator** | **Inherits from** | **Description** |
| --- | --- | --- |
| [**`Manipulator`**](../ScriptReference/UIElements.Manipulator.md) |  | Base class for all provided manipulators. |
| [**`KeyboardNavigationManipulator`**](../ScriptReference/UIElements.KeyboardNavigationManipulator.md) | `Manipulator` | Handles translation of device-specific input events to higher-level navigation operations with a keyboard. |
| [**`MouseManipulator`**](../ScriptReference/UIElements.MouseManipulator.md) | `Manipulator` | Handles mouse input. Has a list of activation filters. |
| [**`ContextualMenuManipulator`**](../ScriptReference/UIElements.ContextualMenuManipulator.md) | `MouseManipulator` | Displays a contextual menu when the user clicks the right mouse button or presses the menu key on the keyboard. |
| [**`PointerManipulator`**](../ScriptReference/UIElements.PointerManipulator.md) | `MouseManipulator` | Handles pointer input. Has a list of activation filters. |
| [**`Clickable`**](../ScriptReference/UIElements.Clickable.md) | `PointerManipulator` | Tracks mouse events on an element, and identifies when a click occurs, which is both a pointer press and release on the same element. |

## Examples

The following examples demonstrate how to use manipulators to handle events. They do the following:

* Inherit the manipulators from the [`PointerManipulators`](../ScriptReference/UIElements.PointerManipulator.md) class, which handles mouse input.
* Use the [`activators`](../ScriptReference/UIElements.MouseManipulator-activators.md) list to specify the conditions that activate the manipulator. For example, to activate the manipulator when the user clicks the left mouse button, add
  [`ManipulatorActivationFilter`](../ScriptReference/UIElements.ManipulatorActivationFilter.md) to the `activators` list with the `button` field set to `MouseButton.LeftMouse`.
* Use the [`target`](../ScriptReference/UIElements.Manipulator-target.md) property to access the element that the manipulator is attached to.
* Override the [`RegisterCallbacksOnTarget`](../ScriptReference/UIElements.Manipulator.RegisterCallbacksOnTarget.md) and [`UnregisterCallbacksFromTarget`](../ScriptReference/UIElements.Manipulator.UnregisterCallbacksFromTarget.md) methods to register and unregister the event callbacks.

### Create an ExampleDragger manipulator

The following example creates a manipulator that moves an element when you click and drag it:

```
using UnityEngine;
using UnityEngine.UIElements;

public class ExampleDragger : PointerManipulator
{
    private Vector3 m_Start;
    protected bool m_Active;
    private int m_PointerId;
    private Vector2 m_StartSize;

    public ExampleDragger()
    {
        m_PointerId = -1;
        activators.Add(new ManipulatorActivationFilter { button = MouseButton.LeftMouse });
        m_Active = false;
    }

    protected override void RegisterCallbacksOnTarget()
    {
        target.RegisterCallback<PointerDownEvent>(OnPointerDown);
        target.RegisterCallback<PointerMoveEvent>(OnPointerMove);
        target.RegisterCallback<PointerUpEvent>(OnPointerUp);
    }

    protected override void UnregisterCallbacksFromTarget()
    {
        target.UnregisterCallback<PointerDownEvent>(OnPointerDown);
        target.UnregisterCallback<PointerMoveEvent>(OnPointerMove);
        target.UnregisterCallback<PointerUpEvent>(OnPointerUp);
    }

    protected void OnPointerDown(PointerDownEvent e)
    {
        if (m_Active)
        {
            e.StopImmediatePropagation();
            return;
        }

        if (CanStartManipulation(e))
        {
            m_Start = e.localPosition;
            m_PointerId = e.pointerId;

            m_Active = true;
            target.CapturePointer(m_PointerId);
            e.StopPropagation();
        }
    }

    protected void OnPointerMove(PointerMoveEvent e)
    {
        if (!m_Active || !target.HasPointerCapture(m_PointerId))
            return;

        Vector2 diff = e.localPosition - m_Start;

        target.style.top = target.layout.y + diff.y;
        target.style.left = target.layout.x + diff.x;

        e.StopPropagation();
    }

    protected void OnPointerUp(PointerUpEvent e)
    {
        if (!m_Active || !target.HasPointerCapture(m_PointerId) || !CanStopManipulation(e))
            return;

        m_Active = false;
        target.ReleaseMouse();
        e.StopPropagation();
    }
}
```

### Create an ExampleResizer manipulator

The following example creates a manipulator that resizes an element when you drag it:

```
using UnityEngine;
using UnityEngine.UIElements;

public class ExampleResizer : PointerManipulator
{
    private Vector3 m_Start;
    protected bool m_Active;
    private int m_PointerId;
    private Vector2 m_StartSize;
    public ExampleResizer()
    {
        m_PointerId = -1;
        activators.Add(new ManipulatorActivationFilter { button = MouseButton.LeftMouse });
        m_Active = false;
    }

    protected override void RegisterCallbacksOnTarget()
    {
        target.RegisterCallback<PointerDownEvent>(OnPointerDown);
        target.RegisterCallback<PointerMoveEvent>(OnPointerMove);
        target.RegisterCallback<PointerUpEvent>(OnPointerUp);
    }

    protected override void UnregisterCallbacksFromTarget()
    {
        target.UnregisterCallback<PointerDownEvent>(OnPointerDown);
        target.UnregisterCallback<PointerMoveEvent>(OnPointerMove);
        target.UnregisterCallback<PointerUpEvent>(OnPointerUp);
    }

    protected void OnPointerDown(PointerDownEvent e)
    {
        if (m_Active)
        {
            e.StopImmediatePropagation();
            return;
        }

        if (CanStartManipulation(e))
        {
            m_Start = e.localPosition;
            m_StartSize = target.layout.size;
            m_PointerId = e.pointerId;
            m_Active = true;
            target.CapturePointer(m_PointerId);
            e.StopPropagation();
        }
    }

    protected void OnPointerMove(PointerMoveEvent e)
    {
        if (!m_Active || !target.HasPointerCapture(m_PointerId))
            return;

        Vector2 diff = e.localPosition - m_Start;

        target.style.height = m_StartSize.y + diff.y;
        target.style.width = m_StartSize.x + diff.x;

        e.StopPropagation();
    }

    protected void OnPointerUp(PointerUpEvent e)
    {
        if (!m_Active || !target.HasPointerCapture(m_PointerId) || !CanStopManipulation(e))
            return;

        m_Active = false;
        target.ReleasePointer(m_PointerId);
        m_PointerId = -1;
        e.StopPropagation();
    }
}
```

### Add or remove a manipulator

To add a manipulator to an element, use the [`AddManipulator`](../ScriptReference/UIElements.VisualElementExtensions.AddManipulator.md) method. To remove a manipulator from an element, use the [`RemoveManipulator`](../ScriptReference/UIElements.VisualElementExtensions.AddManipulator.md) method.

The following example adds and removes the `ExampleDragger` to a `VisualElement`:

```
// Create a VisualElement.
var myElement = new VisualElement();
var dragger = new ExampleDragger();

// Add manipulators to the VisualElement.
myElement.AddManipulator(dragger);

// Remove manipulators from the VisualElement.
myElement.RemoveManipulator(dragger);
```

The following example adds the `ExampleResizer` to a `VisualElement`:

```
var box = new VisualElement()
{
    style =
    {
        left = 100,
        top = 100,
        width = 100,
        height = 100,
        backgroundColor = Color.red
    },
    pickingMode = PickingMode.Position,
};

box.AddManipulator(new ExampleResizer());
```

## Additional resources

* [Handle events](UIE-Events-Handling.md)
* [Create a drag-and-drop UI inside a custom Editor window](UIE-create-drag-and-drop-ui.md)
* [Create a drag-and-drop UI to drag between Editor windows](UIE-drag-across-windows.md)
* 📺 **Video**: [UI Toolkit Tutorial Series: Drag and Drop Manipulator](https://www.youtube.com/watch?v=HvsCvq0L6I4)