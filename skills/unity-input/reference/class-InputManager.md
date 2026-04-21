# Input Manager

|  |
| --- |
| **Important**: **Input Manager** is a legacy feature and not recommended for new projects. **For new projects you should use the [Input System Package](Input.md).** |

---

The **Input Manager** window allows you to define input axes and their associated actions for your Project. To access it, from Unity’s main menu, go to **Edit > Project Settings**, then select **Input Manager** from the navigation on the right.

The Input Manager uses the following types of controls:

* **Key** refers to any key on a physical keyboard, such as W, Shift, or the space bar.
* **Button** refers to any button on a physical controller (for example, gamepads), such as the X button on a remote control.
* A **virtual axis** (plural: **axes**) is mapped to a control, such as a button or a key. When the user activates the control, the axis receives a value in the range of [–1..1]. You can use this value in your **scripts**.

## Physical keys

The Physical keys option allows you to map key codes to the physical keyboard layout, rather than to the language-specific layout that may vary between users in different regions.

For example, on some keyboards the first row of letters reads “QWERTY”, and on others it reads “AZERTY”. This means if you scripted specific controls to use the well known “WASD” keys for movement, they would not be in the correct physical arrangement (like the arrow-key arrangement) on an AZERTY-layout keyboard.

With Physical Keys enabled, Unity uses a generic ANSI/ISO “Qwerty” layout to represent the **physical location** of the keys regardless of the user’s actual layout. This means if you specify the “Q” key, it will always be the left-most letter on the first row of letter keys, even if the user’s keyboard has a different letter in that position.

Note, you should not read key input for in-game text input, because this will not allow users to enter non-Latin characters. Instead, use [`Input.compositionString`](../ScriptReference/Input-compositionString.md).

## Virtual axes

Every Project you create has a number of input axes created by default. These axes enable you to use keyboard, mouse, and joystick input in your Project straight away.

To see more about these axes, open the **Input Manager** window, and click the arrow next to any axis name to expand its properties.

Each input axis has the following properties:

| **Property** | **Function** |
| --- | --- |
| **Name** | Axis name. You can use this to access the axis from scripts. |
| **Descriptive Name, Descriptive Negative Name** | These values are deprecated and do not work. Previously, they displayed for the user on the Rebind Controls screen at startup, but this screen has also been deprecated. |
| **Negative Button, Positive Button** | The controls to push the axis in the negative and positive direction respectively. These can be keys on a keyboard, or buttons on a joystick or mouse. |
| **Alt Negative Button, Alt Positive Button** | Alternative controls to push the axis in the negative and positive direction respectively. |
| **Gravity** | Speed in units per second that the axis falls toward neutral when no input is present. |
| **Dead** | How far the user needs to move an analog stick before your application registers the movement. At runtime, input from all analog devices that falls within this range will be considered null. |
| **Sensitivity** | Speed in units per second that the axis will move toward the target value. This is for digital devices only. |
| **Snap** | If enabled, the axis value will reset to zero when pressing a button that corresponds to the opposite direction. |
| **Type** | The type of input that controls the axis. Select from these values:  - Key or Mouse button - Mouse Movement - Joystick Axis |
| **Axis** | The axis of a connected device that controls this axis. |
| **JoyNum** | The connected Joystick that controls this axis. You can select a specific joystick, or query input from all joysticks. |

Axis values can be:

* Between –1 and 1 for joystick and keyboard input. The neutral position for these axes is 0. Some types of controls, such as buttons on a keyboard, aren’t sensitive to input intensity, so they can’t produce values other than –1, 0, or 1.
* Mouse delta (how much the mouse has moved during the last frame) for mouse input. The values for mouse input axes can be larger than 1 or smaller than –1 when the user moves the mouse quickly.

### Adding, removing, and copying virtual axes

To add a virtual axis, increase the number in the **Size** field. This creates a new axis at the bottom of the list. The new axis copies the properties of the previous axis in the list.

To remove a virtual axis, you can either:

* Decrease the number in the **Size** field. This removes the last axis in the list.
* Right-click any axis, and select **Delete Array Element**.  
  **Note:** You can’t undo this action.

To copy a virtual axis, right-click it and select **Duplicate Array Element**.

### Mapping virtual axes to controls

To map a key or button to an axis, enter its name in the **Positive Button** or **Negative Button** property in the Input Manager.

Key names follow these naming conventions:

| **Key family** | **Naming convention** |
| --- | --- |
| Letter keys | `a`, `b`, `c`… |
| Number keys | `1`, `2`, `3`… |
| Arrow keys | `up`, `down`, `left`, `right` |
| Numpad keys | `[1]`, `[2]`, `[3]`, `[+]`, `[equals]`… |
| Modifier keys | `right shift`, `left shift`, `right ctrl`, `left ctrl`, `right alt`, `left alt`, `right cmd`, `left cmd` |
| Special keys | `backspace`, `tab`, `return`, `escape`, `space`, `delete`, `enter`, `insert`, `home`, `end`, `page up`, `page down` |
| Function keys | `f1`, `f2`, `f3`… |

Mouse buttons are named `mouse 0, mouse 1, mouse 2,` and so on.

Joystick buttons follow these naming conventions:

| **Button origin** | **Naming convention** |
| --- | --- |
| A specific button on any joystick | `joystick button 0`, `joystick button 1`, `joystick button 2`… |
| A specific button on a specific joystick | `joystick 1 button 0`, `joystick 1 button 1`, `joystick 2 button 0`… |

You can also query input for a specific key or button with [`Input.GetKey`](../ScriptReference/Input.GetKey.md) and the naming conventions specified above. For example:

```
Input.GetKey("a");
```

Another way to access keys is to use the [`KeyCode`](../ScriptReference/KeyCode.md) enumeration.

### Using virtual axes in scripts

To access virtual axes from scripts, you can use the axis name.

For example, to query the current value of the Horizontal axis and store it in a variable, you can use [`Input.GetAxis`](../ScriptReference/Input.GetAxis.md) like this:

```
float horizontalInput = Input.GetAxis ("Horizontal");
```

For axes that describe an event rather than a movement (for example, firing a weapon in a game), use [`Input.GetButtonDown`](../ScriptReference/Input.GetButtonDown.md) instead.

If two or more axes have the same name, the query returns the axis with the largest absolute value. This makes it possible to assign more than one input device to an axis name.

For example, you can create two axes named Horizontal and assign one to keyboard input and the other to joystick input. If the user is using the joystick, input comes from the joystick and keyboard input is null. Otherwise, input comes from the keyboard and joystick input is null. This enables you to write a single script that covers input from multiple controllers.

**Example**

You can use input from the **Horizontal** and **Vertical** axes and the [`transform.Translate`](../ScriptReference/Transform.Translate.md) method to move a **GameObject** in XZ space (forward, back, left, or right). Add the following code to the `update()` method on a script attached to the GameObject you want to move:

```
float moveSpeed = 10;
//Define the speed at which the object moves.

float horizontalInput = Input.GetAxis("Horizontal");
//Get the value of the Horizontal input axis.

float verticalInput = Input.GetAxis("Vertical");
//Get the value of the Vertical input axis.

transform.Translate(new Vector3(horizontalInput, verticalInput, 0) * moveSpeed * Time.deltaTime);
//Move the object to XYZ coordinates defined as horizontalInput, 0, and verticalInput respectively.
```

[`Time.deltaTime`](../ScriptReference/Time-deltaTime.md) represents the time that passed since the last frame. Multiplying the `moveSpeed` variable by `Time.deltaTime` ensures that the GameObject moves at a constant speed every frame.

InputManager