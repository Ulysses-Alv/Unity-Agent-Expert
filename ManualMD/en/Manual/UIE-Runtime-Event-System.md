# Runtime UI event system and input handling

UI Toolkit uses an **event system** to handle input and send events to all active panels. This system processes input events and sends them to the appropriate elements in the UI.

When you enter Play mode, UI Toolkit creates a default event system that isn’t part of any **Scene**, and
provides basic support for most input devices. That default event system is used if there is not other active event system to replace it.

Refer to [Use UI Toolkit and uGUI with different input systems](#ugui-input) for more information on how to use the **EventSystem** component to replace UI Toolkit’s default event system.

## Use UI Toolkit with different input systems

Unity provides two input handling systems: the legacy [Input Manager](class-InputManager.md) and the new [Input System package](https://docs.unity3d.com/Packages/com.unity.inputsystem@1.8/manual/index.html). The Input Manager is part of the core Unity platform and serves as the default system unless you install the Input System package. The new Input System package provides greater flexibility and broader devices and platforms support.

UI Toolkit’s default event system can receive events from both input systems. It automatically detects the active input handling system and processes events accordingly.

### Set the active input handling system

To set the active input handling system in your project:

1. Select **Edit** > **Project Settings** > **Player**.
2. In the **Player** window, under **Other Settings** > **Configuration**, set **Active Input Handling** to one of the following options:

   * **Input Manager (Old)**: Use the legacy Input Manager.
   * **Input System Package (New)**: Use the Input System package.
   * **Both**: Use the Input System package if available, otherwise fall back to the legacy Input Manager.

If the Input System package is active in your project, UI Toolkit automatically derives its events from actions defined in the input system. To [configure those actions](https://docs.unity3d.com/Packages/com.unity.inputsystem@1.8/manual/ActionsEditor.html), go to **Edit** > **Project Settings** > **Input System Package**.

### Set up input handling with the Input Manager

To set up input handling with Input Manager, go to **Edit** > **Project Settings** > **Input Manager**.

![Input Manager settings](../uploads/Main/uitk/input-manager-nav.png)

Input Manager settings

You can configure the **Horizontal** and **Vertical** axes to influence how [`NavigationMoveEvents`](../ScriptReference/UIElements.NavigationMoveEvent.md) are generated in UI Toolkit. You can also modify the **Submit** and **Cancel** actions to generate [`NavigationSubmitEvent`](../ScriptReference/UIElements.NavigationSubmitEvent.md) and [`NavigationCancelEvent`](../ScriptReference/UIElements.NavigationCancelEvent.md) in UI Toolkit.

### Set up input handling with the Input System package

To set up input handling with the Input System package, go to **Edit** > **Project Settings** > **Input System Package**.

![Input System settings](../uploads/Main/uitk/input-system-nav.png)

Input System settings

The Input System package offers enhanced configurability compared to Input Manager. You can use the project-wide actions asset to set up how `NavigationMoveEvents`, [`PointerMoveEvent`](../ScriptReference/UIElements.PointerMoveEvent.md) (“UI / Point” action), [`PointerDownEvent`](../ScriptReference/UIElements.PointerDownEvent.md), [`PointerUpEvent`](../ScriptReference/UIElements.PointerUpEvent.md) (“UI / Click”), [`WheelEvent`](../ScriptReference/UIElements.WheelEvent.md) (“UI / ScrollWheel”), `NavigationSubmitEvent`, and `NavigationCancelEvent` are generated for UI Toolkit.

Refer to the Input System package’s [UI Support](https://docs.unity3d.com/Packages/com.unity.inputsystem@1.11/manual/UISupport.html) documentation for more information about each individual action.

## Use UI Toolkit and uGUI with different input systems

You can use UI Toolkit’s UI Documents and uGUI components at the same time.
To do so, add an **EventSystem** component to the scene by any of the following:

* In the **Hierarchy** view, select **UI** > EventSystem.
* Create any uGUI component.

When you use UI Toolkit with an **EventSystem** component, you need to choose an appropriate input module
for the input system that’s active in your project.

The following table outlines the required components and settings for each input system usage:

| **Usage** | **Required components** | **Active Input Handling** |
| --- | --- | --- |
| **UI Toolkit elements with legacy Input Manager** | Default event system (No Scene component required) | **Input Manager (Old)** |
| **UI Toolkit elements with Input System package** | Default event system (No Scene component required) | **Input System Package (New)** or **Both** |
| **UI Toolkit elements and uGUI components with legacy Input Manager** | A **Standalone Input Module** and an **EventSystem** component | **Input Manager (Old)** or **Both** |
| **UI Toolkit elements and uGUI components with Input System package** | An **Input System UI Input Module** and an **EventSystem** component | **Input System Package (New)** or **Both** |

When you add your first uGUI element to the Scene, an [**EventSystem**](https://docs.unity3d.com/Packages/com.unity.ugui@3.0/manual/script-EventSystem.html) and a [**Standalone Input Module**](https://docs.unity3d.com/Packages/com.unity.ugui@3.0/manual/script-StandaloneInputModule.html) are automatically added to the Scene.

The **EventSystem** belongs to uGUI. It’s responsible for uGUI events, derived from either legacy Input Manager or the Input System package, through an interchangeable Input Module component.

The **Standalone Input Module** dispatches events to UI Toolkit elements.

When you activate the **Input System package** in your project, an [**Input System UI Input Module**](https://docs.unity3d.com/Packages/com.unity.inputsystem@1.8/manual/UISupport.html) is added instead of the **Standalone Input Module**. The **Input System UI Input Module** and its accompanying **EventSystem** ensure that the events from both UI Toolkit and uGUI elements are properly dispatched.

The **EventSystem** is responsible for reading the Scene and executing events, while the **Input System UI Input Module** processes input and initiates event execution. You can change the **Standalone Input Module** or **Input System UI Input Module** with other input modules, which can alter the type of input consumed. Regardless of the input module used, all events are executed through the **EventSystem**.

If you add and enable a uGUI **EventSystem** in the Scene, UI Toolkit detects it and creates two uGUI-compabible components for each UI Toolkit panel: [`PanelRaycaster`](https://docs.unity3d.com/Packages/com.unity.ugui@3.0/api/UnityEngine.UIElements.PanelRaycaster.html) and [`PanelEventHandler`](https://docs.unity3d.com/Packages/com.unity.ugui@3.0/api/UnityEngine.UIElements.PanelEventHandler.html). These components serve as intermediaries between uGUI events and UI Toolkit events. The presence of these components deactivates the UI Toolkit’s automated, built-in input processing. This means that the UI Toolkit relies on these components to handle input events when they are present.

If the Scene uses more than one Panel Settings asset, the event system dispatches pointer events to their panels based on their sorting order. UI Toolkit determines the recipient of pointer events by comparing the sorting order of its panel with that of uGUI canvases and other valid raycast targets. This process decides whether a UI Toolkit element, a uGUI object, or another object in the Scene receives the event. Similarly, UI Toolkit uses `currentSelectedGameObject` of the **EventSystem** to manage the focus. When a UI Toolkit panel wants to get focus, it removes the focus from other uGUI objects, and when a uGUI object becomes selected, UI Toolkit panels automatically lose their focus.

A pointer event propagates through the panels until a panel responds to it. The first panel that uses an event to affect the focused element becomes the focused panel for the event system. This panel continues to receive keyboard events until another event causes a different panel to become the focused panel.

**Note**: Stopping an event’s propagation and giving an element focus are distinct actions. For example, when you click a button, it stops the propagation and allows only the button to respond to the press, but it doesn’t prevent other default click actions, such as giving focus to the button or any clicked focusable element.

In some cases, you might need to add a one-frame delay after input is processed before requesting focus for an element, especially if that input affects focus through other code paths, as shown in the example below:

```
public class FocusOnNextFrameExample : MonoBehaviour
{
    void OnEnable()
    {
        var root = GetComponent<UIDocument>().rootVisualElement;
        root.Q<Button>("my-button").clicked += () =>
        {
            root.schedule.Execute(() => root.Q<TextField>("my-text-field").Focus());
        };
    }
}
```

## Additional resources

* [Input System package](https://docs.unity3d.com/Packages/com.unity.inputsystem@1.8/manual/index.html)
* [Get started with runtime UI](UIE-get-started-with-runtime-ui.md)
* [Render UI in the Game view](UIE-render-runtime-ui.md)
* [Panel Settings properties](UIE-Runtime-Panel-Settings.md)
* [FAQ for input and event systems with UI Toolkit](UIE-faq-event-and-input-system.md)