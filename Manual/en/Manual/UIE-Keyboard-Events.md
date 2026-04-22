# Keyboard events

Keyboard events occur when you press or release keys on the keyboard. Each event includes information about the [modifier](../ScriptReference/UIElements.KeyboardEventBase_1-modifiers.md), [text character](../ScriptReference/UIElements.KeyboardEventBase_1-character.md), and related [key code](../ScriptReference/UIElements.KeyboardEventBase_1-keyCode.md) for the event.

Many standard controls use the `KeyDownEvent` to encode shortcuts or accessibility behaviors. The following examples all use keyboard events:

* The `Toggle` and `Button` classes listen for `Enter` and `Spacebar` key presses as replacement actions for mouse clicks.
* The **ScrollView** and Slider controls use directional arrow key presses to modulate their values.
* The **TextField control** looks at both the `keyCode` property and the character property to execute special actions or to accept text.

The base class for all keyboard events is [KeyboardEventBase](../ScriptReference/UIElements.KeyboardEventBase_1.md).

| **Event** | **Description** | **Trickles down** | **Bubbles up** | **Cancellable** |
| --- | --- | --- | --- | --- |
| [KeyDownEvent](../ScriptReference/UIElements.KeyDownEvent.md) | Sent when the user presses a key on the keyboard. | Yes | Yes | Yes |
| [KeyUpEvent](../ScriptReference/UIElements.KeyUpEvent.md) | Sent when the user releases a key on the keyboard. | Yes | Yes | Yes |

## Unique properties

**`keyCode`**: The `keyCode` property returns a character key that corresponds directly to a physical key on an input device, such as a keyboard or joystick. The difference between the `character` property and the `keyCode` property is that `keyCode` represents a physical key, while `character` represents the entry of a specific character. For example, both `a` and `A` return `keyCode=KeyCode.A` during a `keyDownEvent`.

**`character`**: The `character` property returns a character code during a `keyDownEvent`.

**`modifiers`**: The `modifiers` property returns which modifier key is held down. Some examples of modifier keys are the `Shift`, `Ctrl`, or `Alt` keys.

For more information, see the [Modifier keys section](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key/Key_Values#modifier_keys) of the MDN documentation.

## Event List

The following list provides the name, description, and target of each event in the event family.

By default, a **visual element** doesn’t receive keyboard events. Only elements that are focusable and currently in focus are targeted for keyboard events. This is because keyboard events trickle down and bubble up, allowing parent elements to receive them as well.

In summary, to begin receiving keyboard events, you must mark the element as `focusable=true` and explicitly give it focus using `element.Focus()`. This ensures that the element is eligible to receive keyboard events.

### KeyDownEvent

A [KeyDownEvent](../ScriptReference/UIElements.KeyDownEvent.md) is sent each time you press a key on the keyboard. The key pressed contains the `keyCode` property for that event. If that key press has text input associated with it, additional events are sent for each character of text input. The `character` property contains the character for those events.

When you press and release `a`, UI Toolkit sends the following events:

```
KeyDownEvent { keyCode=KeyCode.A }
KeyDownEvent { character=’a’ }
KeyUpEvent { keyCode=KeyCode.A }
```

When you press and release `Ctrl+a`, UI Toolkit sends the following events:

```
KeyDownEvent { keyCode=KeyCode.LeftControl, modifiers=EventModifiers.Control }
KeyDownEvent { keyCode=KeyCode.A, modifiers=EventModifiers.Control }
KeyUpEvent { keyCode=KeyCode.A, modifiers=EventModifiers.Control }
KeyUpEvent { keyCode=KeyCode.LeftControl }
```

**`target`**: The visual element that has focus. If no element has focus, the root visual element of the panel.

### KeyUpEvent

A [KeyUpEvent](../ScriptReference/UIElements.KeyUpEvent.md) is sent when you release a key on the keyboard. The keyCode property for that event contains the key being released. `KeyDownEvent` has additional events sent when a keystroke has an associated text input.

When you press and release `a`, UI Toolkit sends the following events:

```
KeyDownEvent { keyCode=KeyCode.A }
KeyDownEvent { character=’a’ }
KeyUpEvent { keyCode=KeyCode.A }
```

When you press and release `Ctrl+a`, UI Toolkit sends the following events:

```
KeyDownEvent { keyCode=KeyCode.LeftControl, modifiers=EventModifiers.Control }
KeyDownEvent { keyCode=KeyCode.A, modifiers=EventModifiers.Control }
KeyUpEvent { keyCode=KeyCode.A, modifiers=EventModifiers.Control }
KeyUpEvent { keyCode=KeyCode.LeftControl }
```

**`target`**: The visual element that has focus. If no element has focus, the root visual element of the panel.

## Examples

The following code example prints a message to the console when you press a key in a TextField. This code sample highlights the event firing of both `KeyUpEvent` and `KeyDownEvent`.

1. Create a Unity project with any template.
2. In the SampleScene, select **GameObject** > **UI Toolkit** > **UI Document**.
3. Create a C# script named `KeyboardEventTest.cs` with the following contents:

   ```
   using UnityEngine;
   using UnityEngine.UIElements;

   // Add KeyboardEventTest to a GameObject with a valid UIDocument.
   // When the user presses a key, it will print the keyboard event properties to the console.
   [RequireComponent(typeof(UIDocument))]
   public class KeyboardEventTest : MonoBehaviour
   {
       void OnEnable()
       {
           var root = GetComponent<UIDocument>().rootVisualElement;
           root.Add(new Label("Press any key to see the keyDown properties"));
           root.Add(new TextField());
           root.Q<TextField>().Focus();
           root.RegisterCallback<KeyDownEvent>(OnKeyDown, TrickleDown.TrickleDown);
           root.RegisterCallback<KeyUpEvent>(OnKeyUp, TrickleDown.TrickleDown);
       }
       void OnKeyDown(KeyDownEvent ev)
       {
           Debug.Log("KeyDown:" + ev.keyCode);
           Debug.Log("KeyDown:" + ev.character);
           Debug.Log("KeyDown:" + ev.modifiers);
       }

       void OnKeyUp(KeyUpEvent ev)
       {
           Debug.Log("KeyUp:" + ev.keyCode);
           Debug.Log("KeyUp:" + ev.character);
           Debug.Log("KeyUp:" + ev.modifiers);
       }
   }
   ```
4. Select the UIDocument **GameObject** in the Hierarchy window.
5. Drag `KeyboardEventTest.cs` to **Add Component** in the **Inspector** window.
6. Enter Play mode and type in the TextField.

## Additional resources

* [UI Toolkit events](UIE-Events.md)
* [UI Toolkit built-in controls reference](UIE-ElementRef.md#built-in-controls)
* [UI Toolkit controls](UIE-Controls.md)