# Change events

When the value of an element changes, a `ChangeEvent` is sent. This is typically sent when a value in a field of a control changes. For example, when the user toggles a checkbox.

The `ChangeEvent` is a typed event, and contains both the previous and the new value of the **visual element**.

The event triggers after a change assigns a new value to a visual element. You can’t cancel change events to prevent a value change on a visual element.

The base class for `ChangeEvent` is the [EventBase](../ScriptReference/UIElements.EventBase.md) class.

| **Event** | **Description** | **Trickles down** | **Bubbles up** | **Cancellable** |
| --- | --- | --- | --- | --- |
| [ChangeEvent](../ScriptReference/UIElements.ChangeEvent_1.md) | A generic event sent when the value of an element changes. | ✔ | ✔ |  |

## Unique properties

**`previousValue`**: The previous value of the target control.

**`newValue`**: The new value of the target control.

## Event List

### ChangeEvent

`ChangeEvent` is a notification event that allows you to react to the value change of a visual element. For example, when you toggle a checkbox to mute music in a game, the game should stop all music.

This event applies to all controls that implement the `INotifyValueChanged<T>`, where `<T>` is the type of the `ChangeEvent`. This event is also used internally to update properties within objects linked to the UI via [Bindings](UIE-Binding.md).

It fires even when the value of a control is set by code. You can modify the value on a control without firing the `ChangeEvent` by calling `SetValueWithoutNotify` in the `INotifyValueChange<T>` interface.

You can register a callback function to receive a `ChangeEvent` in two ways:

1. Call `RegisterCallback<>()` on a visual element
2. Call `RegisterValueChangedCallback()` on a visual element deriving from `INotifyValueChange<T>`

Registering a callback via RegisterCallback works on all visual elements, regardless of whether they store internal values or not. If you want to listen to any changes that happen in the child controls of a parent element, this is the method to use.

Because `ChangeEvent` is a typed event, you must specify the type when registering the event. The code below demonstrates registering and receiving a `ChangeEvent` of the type `bool`.

```
// Registering the callback
rootVisualElement.RegisterCallback<ChangeEvent<bool>>(OnBoolChangedEvent);
```

```
// Event callback
private void OnBoolChangedEvent(ChangeEvent<bool> evt) 
{ 
    // Handling code
}
```

Elements that hold values, such as toggles and integer fields, implement the `INotifyValueChange<T>` interface. It’s possible to register a callback on these elements directly, by calling [RegisterValueChangedCallback](../ScriptReference/UIElements.INotifyValueChangedExtensions.RegisterValueChangedCallback.md). This is a more convenient way to register a callback because the type is already built-in. You can unregister event handlers again by calling `myElement.UnregisterValueChangedCallback`.

```
var newToggle = new Toggle("Test Toggle");
newToggle.RegisterValueChangedCallback(OnTestToggleChanged);
```

```
private void OnTestToggleChanged(ChangeEvent<bool> evt)
{ 
    // Handling code
}
```

**`target`**: The element where the change of state occurs.

## Examples

The following examples demonstrate the usage of `ChangeEvent` and how to set and get a control value.

To view an example, do the following:

1. Under **Assets** > **Scripts** > **Editor**, create a C# script called ChangeEventTestWindow.
2. Copy the example code into the C# script.
3. From the Editor Toolbar, select **Window** > **UI Toolkit** > **Change Events Test Window**

### Example 1: Registering callbacks to the receive change events

```
using UnityEditor;
using UnityEngine;
using UnityEngine.UIElements;

public class ChangeEventTestWindow : EditorWindow
{
    private Toggle m_MyToggle;

    [MenuItem("Window/UI Toolkit/Change Event Test Window")]
    public static void ShowExample()
    {
        ChangeEventTestWindow wnd = GetWindow<ChangeEventTestWindow>();
        wnd.titleContent = new GUIContent("Change Event Test Window");
    }

    public void CreateGUI()
    {
        // Create a toggle
        m_MyToggle = new Toggle("Test Toggle") { name = "My Toggle" };
        rootVisualElement.Add(m_MyToggle);

        // Register a callback on the toggle
        m_MyToggle.RegisterValueChangedCallback(OnTestToggleChanged);

        // Register a callback on the parent
        rootVisualElement.RegisterCallback<ChangeEvent<bool>>(OnBoolChangedEvent);
    }

    private void OnBoolChangedEvent(ChangeEvent<bool> evt)
    {
        Debug.Log($"Toggle changed. Old value: {evt.previousValue}, new value: {evt.newValue}");
    }

    private void OnTestToggleChanged(ChangeEvent<bool> evt)
    {
        Debug.Log($"A bool value changed. Old value: {evt.previousValue}, new value: {evt.newValue}");
    }
}
```

### Example 2: Setting a control value through code

```
using UnityEditor;
using UnityEngine;
using UnityEngine.UIElements;

public class ChangeEventTestWindow : EditorWindow
{
    private Toggle m_MyToggle;

    [MenuItem("Window/UI Toolkit/Change Event Test Window")]
    public static void ShowExample()
    {
        GetWindow<ChangeEventTestWindow>().titleContent = new GUIContent("Change Event Test Window");
    }

    public void CreateGUI()
    {
        // Create a toggle and register callback 
        m_MyToggle = new Toggle("Test Toggle") { name = "My Toggle" };
        m_MyToggle.RegisterValueChangedCallback((evt) => { Debug.Log("Change Event received"); });
        rootVisualElement.Add(m_MyToggle);

        // Create button to toggle the toggle's value
        Button button01 = new Button() { text = "Toggle" };
        button01.clicked += () => 
        {
            m_MyToggle.value = !m_MyToggle.value;
        };
        rootVisualElement.Add(button01);

        // Create button to toggle the toggle's value without triggering a ChangeEvent
        Button button02 = new Button() { text = "Toggle without notification" };
        button02.clicked += () =>
        {
            m_MyToggle.SetValueWithoutNotify(!m_MyToggle.value);
        };
        rootVisualElement.Add(button02);
    }
}
```