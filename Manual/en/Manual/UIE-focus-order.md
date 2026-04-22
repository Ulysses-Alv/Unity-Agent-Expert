# Focus system in UI Toolkit

The focus system in UI Toolkit manages which element currently receives keyboard or navigation input.
It controls how focus moves between elements, and raises focus events when elements gain or lose focus.

## Key Concepts

The focus system consists of several key components:

* [`focusable`](../ScriptReference/UIElements.Focusable-focusable.md): Can this element receive focus?
* [`canGrabFocus`](../ScriptReference/UIElements.Focusable-canGrabFocus.md): Is this element currently able to be focused (visible, enabled, etc.)?
* [`tabIndex`](../ScriptReference/UIElements.Focusable-tabIndex.md): Order in which elements are focused when navigating.
* [`delegatesFocus`](../ScriptReference/UIElements.Focusable-delegatesFocus.md): If true, passes focus to a suitable child element.
* [`Focus()`](../ScriptReference/UIElements.Focusable.Focus.md): Method to request focus on an element.
* [`Blur()`](../ScriptReference/UIElements.Focusable.Blur.md): Method to remove focus from an element.
* [`FocusController`](../ScriptReference/UIElements.Focusable-focusController.md): Manages focus state and navigation for a panel.
* [`IFocusRing`](../ScriptReference/UIElements.IFocusRing.md): Interface for focus navigation logic.
* [`FocusChangeDirection`](../ScriptReference/UIElements.IFocusRing.GetFocusChangeDirection.md): Describes navigation direction, such as Next, Previous, Up, and Down.

## Focus workflow

The focus system operates as follows:

1. User input can trigger focus changes in two ways: directly, through pointer interactions (such as mouse clicks or touch), or indirectly, by triggering navigation events (such as Tab, Shift+Tab, or arrow keys), which then result in focus events.
2. The [`EventDispatcher`](../ScriptReference/UIElements.EventDispatcher.md) processes input and, when appropriate, asks the `FocusController` to change focus.
3. The `FocusController` consults the `IFocusRing` to determine the next focusable element.
4. The `FocusController` generates focus change events (such as `FocusEvent` and `BlurEvent`) and adds them to the `EventDispatcher`’s event queue. The `EventDispatcher` dispatches these events after it finishes processing the current user input.
5. After the `EventDispatcher` dispatches the focus events, the newly focused element receives focus. Each element along the event propagation path can handle these events using event callbacks in code.
6. If an element has `delegatesFocus = true`, focus is automatically given to a child element that can be focused.

The following diagram illustrates the focus system components and their interactions:

![Focus system diagram](../uploads/Main/uitk/focus-workflow.png)

Focus system diagram

## Focus order of elements

Each panel has a focus ring that defines the focus order of elements. By default, a depth-first search (DFS) on the **visual tree** defines the focus order of elements. For example, the focus order for the tree depicted in the following image is **F**, **B**, **A**, **D**, **C**, **E**, **G**, **I**, **H**.

![Focus order](../uploads/Main/focus-order.png)

Focus order

Some events use the focus order to define which element holds the focus. For example, the target for a keyboard event is the element in focus.

You can use the [`tabIndex`](../ScriptReference/UIElements.Focusable-tabIndex.md) property to customize the focus order as follows:

* If the `tabIndex` is zero, the element keeps its default tab order, as determined by the focus ring algorithm.
* If the `tabIndex` is positive, the element is placed in front of other elements that either have a zero `tabIndex` (`tabIndex = 0`) or a `tabIndex` value smaller than its own.
* If the `tabIndex` is negative, the element is removed from the tab navigation.

## Change focus

[Focus events](UIE-Focus-Events.md) occur when an element gains or loses focus. You can use the `element.Focus()` method to change focus. This method sends a `FocusEvent` that is processed in the event queue. The actual focus change takes effect after the current event callback has finished executing.

The following example sets focus on a button and tracks focus events:

```
using UnityEngine;
using UnityEngine.UIElements;

[RequireComponent(typeof(UIDocument))]
public class FocusExample : MonoBehaviour
{
    void OnEnable()
    {
        var startGameButton = GetComponent<UIDocument>().rootVisualElement.Q<Button>("StartGame");

        startGameButton.RegisterCallback<FocusInEvent>(ev =>
        {
            Debug.Log("FocusInEvent: button is focused and can be activated with keyboard or gamepad input.");
        });

        Debug.Log("Calling startGameButton.Focus().");
        startGameButton.Focus();

        var isFocused = startGameButton.focusController.focusedElement == startGameButton;
        Debug.Log("Immediately after startGameButton.Focus(): isFocused=" + isFocused);

        // Expected output:
        // > Calling startGameButton.Focus().
        // > FocusInEvent: button is focused and can be activated with keyboard or gamepad input.
        // > Immediately after startGameButton.Focus(): isFocused=true
    }
}
```

## Additional resources

* [Focus events](UIE-Focus-Events.md)