# Mobile Keyboard

In most cases, Unity will handle keyboard input automatically for GUI elements but it is also easy to show the keyboard on demand from a script.

## GUI Elements

The keyboard will appear automatically when a user taps on editable GUI elements. Currently, [GUI.TextField](../ScriptReference/GUI.TextField.md), [GUI.TextArea](../ScriptReference/GUI.TextArea.md) and [GUI.PasswordField](../ScriptReference/GUI.PasswordField.md) will display the keyboard; see the [GUI class](../ScriptReference/GUI.md) documentation for further details.

## Manual Keyboard Handling

**Note**: The `TouchScreenKeyboard` API is only supported on native iPhone, Android, and Windows Store Apps. Attempting to access its properties on unsupported platforms throws an exception. If you’re targeting multiple platforms, make sure your code checks `TouchScreenKeyboard.isSupported` before attempting to use it.

Use the **[TouchScreenKeyboard.Open()](../ScriptReference/TouchScreenKeyboard.Open.md)** function to open the keyboard. Please see the [TouchScreenKeyboard](../ScriptReference/TouchScreenKeyboard.md) scripting reference for the parameters that this function takes.

## Keyboard Layout Options

The Keyboard supports the following options:-

| ***Property:*** | ***Function:*** |
| --- | --- |
| **[TouchScreenKeyboardType.Default](../ScriptReference/TouchScreenKeyboardType.Default.md)** | Letters. Can be switched to keyboard with numbers and punctuation. |
| **[TouchScreenKeyboardType.ASCIICapable](../ScriptReference/TouchScreenKeyboardType.ASCIICapable.md)** | Letters. Can be switched to keyboard with numbers and punctuation. |
| **[TouchScreenKeyboardType.NumbersAndPunctuation](../ScriptReference/TouchScreenKeyboardType.NumbersAndPunctuation.md)** | Numbers and punctuation. Can be switched to keyboard with letters. |
| **[TouchScreenKeyboardType.URL](../ScriptReference/TouchScreenKeyboardType.URL.md)** | Letters with slash and .com buttons. Can be switched to keyboard with numbers and punctuation. |
| **[TouchScreenKeyboardType.NumberPad](../ScriptReference/TouchScreenKeyboardType.NumberPad.md)** | Only numbers from 0 to 9. |
| **[TouchScreenKeyboardType.PhonePad](../ScriptReference/TouchScreenKeyboardType.PhonePad.md)** | Keyboard used to enter phone numbers. |
| **[TouchScreenKeyboardType.NamePhonePad](../ScriptReference/TouchScreenKeyboardType.NamePhonePad.md)** | Letters. Can be switched to phone keyboard. |
| **[TouchScreenKeyboardType.EmailAddress](../ScriptReference/TouchScreenKeyboardType.EmailAddress.md)** | Letters with @ sign. Can be switched to keyboard with numbers and punctuation. |

## Text Preview

By default, an edit box will be created and placed on top of the keyboard after it appears. This works as preview of the text that user is typing, so the text is always visible for the user. However, you can disable text preview by setting **TouchScreenKeyboard.hideInput** to true. Note that this works only for certain keyboard types and input modes. For example, it will not work for phone keypads and multi-line text input. In such cases, the edit box will always appear. **TouchScreenKeyboard.hideInput** is a global variable and will affect all keyboards.

## Visibility and Keyboard Size

There are three keyboard properties in [TouchScreenKeyboard](../ScriptReference/TouchScreenKeyboard.md) that determine keyboard visibility status and size on the screen.

| ***Property:*** | ***Function:*** |
| --- | --- |
| **visible** | Returns **true** if the keyboard is fully visible on the screen and can be used to enter characters. |
| **area** | Returns the position and dimensions of the keyboard. |
| **active** | Returns **true** if the keyboard is activated. This property is not static property. You must have a keyboard instance to use this property. |

Note that **TouchScreenKeyboard.area** will return a Rect with position and size set to 0 until the keyboard is fully visible on the screen. You should not query this value immediately after **TouchScreenKeyboard.Open()**. The sequence of keyboard events is as follows:

* **TouchScreenKeyboard.Open()** is called. **TouchScreenKeyboard.active** returns true. **TouchScreenKeyboard.visible** returns false. **TouchScreenKeyboard.area** returns (0, 0, 0, 0).
* Keyboard slides out into the screen. All properties remain the same.
* Keyboard stops sliding. **TouchScreenKeyboard.active** returns true. **TouchScreenKeyboard.visible** returns true. **TouchScreenKeyboard.area** returns real position and size of the keyboard.

## Secure Text Input

It is possible to configure the keyboard to hide symbols when typing. This is useful when users are required to enter sensitive information (such as passwords). To manually open keyboard with secure text input enabled, use the following code:

```
TouchScreenKeyboard.Open("", TouchScreenKeyboardType.Default, false, false, true);
```

![Hiding text while typing](../uploads/Main/KeyboardSecure.png)

Hiding text while typing

## Alert keyboard

To display the keyboard with a black semi-transparent background instead of the classic opaque, call **TouchScreenKeyboard.Open()** as follows:

```
TouchScreenKeyboard.Open("", TouchScreenKeyboardType.Default, false, false, true, true);
```

![Alert keyboard](../uploads/Main/KeyboardAlert.png)

Alert keyboard