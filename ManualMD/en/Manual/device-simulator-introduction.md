# Device Simulator introduction

The Device Simulator is a Unity Editor feature that simulates how your application appears and behaves on a mobile device.

The Device Simulator consists of:

* The Simulator view: Views your application on a simulated mobile device.
* Simulated classes: Tests code that responds to device-specific behaviors.
* Device definitions: Describes the device to simulate.
* Device Simulator plugins: Configures the UI of the Simulator view.

## Controls in the Simulator view

The Simulator view simulates many common features of mobile devices, including:

* Auto-rotation
* Screen safe area
* Touch input

## Player Settings

The Device Simulator reacts to the following **Player Settings** in the same way that a real device does:

* **Fullscreen Mode**
* **Resolution Scaling**
* **Default Orientation**
* **Graphics API**
* **Render outside safe area**

## Simulated touch input

If you click on the simulated device screen with the mouse cursor, the device simulator creates touch events in the active input solution (either the [Input Manager](class-InputManager.md), the [Input System](https://docs.unity3d.com/Packages/com.unity.inputsystem@latest), or both, depending on your project settings).

**Note**: The Device Simulator only simulates input when the Editor is in Play mode. The Device Simulator doesn’t support multitouch; it can only simulate one finger touch.

## Limitations

The main purpose of the Device Simulator is to view the layout of your application on a target device and test basic interactions. It doesn’t provide an accurate representation of how your application runs on the device.

The Simulator view does not simulate the following:

* The performance characteristics of a device, such as a device’s processor speed or available memory.
* The rendering capabilities of a device.
* Native plugins that don’t work in the Editor.
* Platform #define directives for the simulated device, like UNITY\_IOS.
* Gyroscope rotation.

Only one Simulator view can simulate at one time. This is the active Simulator view.

* If you have only one Simulator view open and no Game views open, the one Simulator view is active regardless of whether it’s visible or not.
* If you have multiple Simulator views open and no Game views open, the last Simulator view that had focus is active.
* If you have a mix of Simulator views and Game views open, if you focus a Game view, Unity disables all simulators and if you focus a Simulator view, the Simulator view remains active while it has focus.

The Device Simulator doesn’t simulate all APIs in the simulated classes. For more information, see [Simulated classes](device-simulator-simulated-classes.md).