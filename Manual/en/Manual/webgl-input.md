# Input in Web

Unity Web supports various types of input from different devices, including gamepads, joysticks, touchscreens, keyboards, and movement sensors.

* [Gamepad and Joystick support](#gamepad-and-joystick-support)
* [Touch support](#touch-support)
* [Keyboard input and focus handling](#keyboard-input-and-focus-handling)
* [Mobile sensor support](#mobile-sensor-support)
* [Lock your cursor](#cursor-locking-support)
* [Full-screen mode](#full-screen-mode)

## Gamepad and Joystick support

Web supports the following inputs for gamepads:

* [Input](../ScriptReference/Input.md)
* [InputSystem](https://docs.unity3d.com/Packages/com.unity.inputsystem@latest)

Web also supports joysticks for browsers that support the [HTML5 Gamepad API](https://developer.mozilla.org/docs/Web/API/Gamepad_API).

Some browsers allow access to input devices only after the user interacts with the device while the application is in focus. This type of security measure prevents the user from using connected devices for browser fingerprinting purposes. Therefore, make sure your application instructs the user to press a button on their gamepad/joystick before you call [Input.GetJoystickNames()](../ScriptReference/Input.GetJoystickNames.md) to check for connected devices.

### Map the controller

The Web **game controller** mapping for the old input system aligns with the [W3 spec](https://www.w3.org/TR/gamepad/#dfn-standard-gamepad), where the button mapping layout is laid out as follows:

| Buttons | Description |
| --- | --- |
| buttons[0] | Bottom button in right cluster = CROSS (X) |
| buttons[1] | Right button in right cluster = CIRCLE |
| buttons[2] | Left button in right cluster = SQUARE |
| buttons[3] | Top button in right cluster = TRIANGLE |

**Note:** As Web follows the W3 spec, it might not be consistent with other platforms, and it requires a different approach if you target multiple platforms simultaneously, such as Windows and Web. For more information, refer to the W3 documentation on [Remapping](https://www.w3.org/TR/gamepad/#dfn-standard-gamepad).

## Touch support

Unity Web implements [Input.touches](../ScriptReference/Input-touches.md) and other related APIs in browsers and devices with touch support. By default, mobile devices display a soft keyboard on the touch screen for the user to enter text into UI input fields. To disable this behavior, use the [WebGLInput.mobileKeyboardSupport](../ScriptReference/WebGLInput-mobileKeyboardSupport.md) property.

## Keyboard input and focus handling

By default, Unity Web processes all keyboard input the web page receives, regardless of whether the Web canvas has focus or not. This allows the user to use a keyboard-based application without the need to click on the Web canvas.

The following notes apply to specific keyboard inputs as described below:

| Keyboard input behavior | Description |
| --- | --- |
| HTML elements (such as text fields) | If you introduce HTML elements (such as text fields) in a web page that’s meant to receive keyboard inputs, it can cause errors. Unity consumes the input events before the rest of the page can receive them.   To make HTML elements receive a keyboard input, set `WebGLInput.captureAllKeyboardInput` to `false`. This way, the application receives input only if the Web canvas has focus. |
| **Esc** key | Some browsers like Safari block the ability to use the **Esc** key to switch to full-screen mode because the **Esc** key is protected in the web environment. Other browsers like Google Chrome can behave unpredictably. When the browser is in full-screen mode, the user can use only the **Esc** key to exit full-screen mode. Depending on your browser type, clicking the **Esc** key might not always trigger the event to forward from the browser to the application. |
| `HideMobileInput` | The `HideMobileInput` option in [TextMeshPro](https://docs.unity3d.com/Packages/com.unity.ugui@2.0/manual/TextMeshPro/index.html) (also known as TMP) input fields has no effect on the Web platform. This option has no effect because you need a **text input field** to trigger the virtual keyboard in browsers. The Web platform instantiates a text input field to trigger the virtual keyboard when the Unity player expects text entry. The text input field appears above the virtual keyboard much like other platforms. Because the virtual keyboard relies on the mobile text input field, it can’t appear hidden on the Web platform. |

## Mobile sensor support

For browsers and mobile devices with touch support, Unity Web includes support for the following sensors:

* Accelerometer with [Input.acceleration](../ScriptReference/Input-acceleration.md)
* LinearAcceleration with [Gyroscope.userAcceleration](../ScriptReference/Gyroscope-userAcceleration.md)
* Gyroscope with [Gyroscope.rotationRate](../ScriptReference/Gyroscope-rotationRate.md)
* Gravity with [Gyroscope.gravity](../ScriptReference/Gyroscope-gravity.md)
* Attitude with [Gyroscope.attitude](../ScriptReference/Gyroscope-attitude.md)
* The [Input System](https://docs.unity3d.com/Packages/com.unity.inputsystem@latest) package also supports these sensors.

**Important**: Browsers allow sensor input only in secure contexts, which means, you must serve the page over [HTTPS](https://en.wikipedia.org/wiki/HTTPS). The single exception is `http://localhost`, which you can use during development.

## Cursor lock support

The Unity Web platform supports cursor locking, which uses the HTML5 API `Element.requestPointerLock`. Use cursor lock to lock the mouse cursor to the center of the game window. When the cursor is locked, it appears hidden in Unity and doesn’t move when the mouse moves. This is particularly helpful for first-person games, where the mouse cursor is typically used to control the orientation of the player’s angle.

**Note:** As browser support for cursor locking varies, refer to the Mozilla documentation on the [Element: requestPointerLock() method](https://developer.mozilla.org/en-US/docs/Web/API/Element/requestPointerLock).

### Activate cursor locking

To lock the cursor, use the [`Cursor.lockState`](../ScriptReference/Cursor-lockState.md) property. For example, the following code switches the cursor into the locked state when the user clicks the left mouse button:

```
void Update()
{
    if (Input.GetMouseButtonDown(0))
    {
        Cursor.lockState = CursorLockMode.Locked;
    }
}
```

Cursor locking needs to be activated by user interaction. For more information, refer to [Additional considerations for full-screen mode and cursor locking](#additional-considerations).

### Deactivate cursor locking

Press the **Esc** key to unlock the cursor.

### Sticky cursor lock

The `stickyCursorLock` property is commonly used in first-person games, because it’s useful to maintain the cursor lock mode regardless of the browser behavior.

Use [`stickyCursorLock`](../ScriptReference/WebGLInput-stickyCursorLock.md) to ensure that the state of `Cursor.lockState` is persistent, even if the browser releases the cursor lock from the Unity canvas (typically using the **Esc** key), in which case the cursor is locked again the next time the canvas is in focus.

Therefore, if you set `WebGLInput._stickyCursorLock` to `true`, the `Cursor.lockState` remains in `CursorLockMode.Locked` state
even if the Unity canvas HTML element unlocks the cursor.

The following occurs if you set `WebGLInput._stickyCursorLock` to `false`:

* `Cursor.lockState` remains synchronized with the browser’s cursor lock state.
* If the user presses the **Esc** key to cancel the canvas cursor lock, `Cursor.lockState` changes to `CursorLockMode.None`.

**Note:** In Web, `stickyCursorLock` is set to `true` by default.

## Full-screen mode

Use full-screen mode in your game to do the following:

* Use the entire screen for your game.
* Hide the browser user interface (UI) elements such as the address bar and tabs.
* Hide the Unity player UI elements such as the title bar and **toolbar**.

The Unity Web platform supports full-screen mode which uses the HTML5 API, `Element.requestFullscreen`.

**Note:** As browser support for full-screen mode varies, refer to the Mozilla documentation on the [Element: requestFullscreen() method](https://developer.mozilla.org/en-US/docs/Web/API/Element/requestFullscreen).

### Activate full-screen mode in Web

To enable full-screen mode, use the `Screen.fullScreen` property. For example, the following code switches the game to full-screen mode when the user presses the **F** key:

```
void Update()
{
    if (Input.GetKeyDown(KeyCode.F))
    {
        Screen.fullScreen = true;
    }
}
```

**Note:** The `Screen.fullScreen` property is set to `false` by default.

Full-screen mode needs to be activated by user interaction. For more information, refer to [Additional considerations for full-screen mode and cursor locking](#additional-considerations).

### Exit full-screen mode

To exit full-screen mode, press the **Esc** key again, or hover the mouse pointer to the top of the screen to access the address bar and tabs.

## Additional considerations for full-screen mode and cursor locking

Due to security concerns, browsers only allow you to lock your cursor and enable full-screen mode after a user-initiated event such as a mouse-click or key press.

Because Unity doesn’t support separate events and rendering loops, it defers event handling until the browser no longer acknowledges full-screen or cursor lock requests issued from Unity scripting as a direct response to the event which triggered it.

Therefore, Unity triggers the request on the next user-initiated event, instead of the event that triggered the cursor lock or full-screen request.

To enable cursor lock or full-screen modes with better results, use mouse/key down events to trigger responses instead of mouse/key up events. This way you can ensure that the deferred request is guaranteed to be triggered by the corresponding mouse/key up event if not by a user-initiated event earlier.

You can also use Unity’s [UI.Button](https://docs.unity3d.com/Packages/com.unity.ugui@1.0/api/UnityEngine.UI.Button.html) component to achieve the desired behavior by creating a subclass of `Button`, which overrides the `OnPointerDown` method.

**Note:** Some browsers might display a notification message or prompt the user to grant permission before they can enter full-screen mode or lock the cursor.

## Additional resources

* [WebGLInput.mobileKeyboardSupport](../ScriptReference/WebGLInput-mobileKeyboardSupport.md)
* [WebGLInput.captureAllKeyboardInput](../ScriptReference/WebGLInput-captureAllKeyboardInput.md)