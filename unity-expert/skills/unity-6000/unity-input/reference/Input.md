# Input

**Input** in Unity refers to users sending signals from the outside world to your game or app using a physical device.

Unity supports input from many types of device, such as gamepads, mouse, keyboard, touchscreen, joystick, movement-sensors like accelerometers or gyroscopes, and **VR** and **AR** controllers. Almost every project requires input of some kind, whether it is to allow users to navigate a UI, control a character in a game, or move around and interact with objects in **virtual reality**.

Unity has two methods of implementing Input:

* The **Input System Package**, which is newer, more flexible, and better supported.
* The legacy ****Input Manager**** (which includes the built-in `Input` class).

## The Input System Package

By default, the [Input System Package](https://docs.unity3d.com/Packages/com.unity.inputsystem@latest) provides input support in Unity. If needed, you can also install it through the [Package Manager](Packages.md).

This is the recommended solution for most projects. It provides an intuitive interface in the Editor to configure input, and provides a variety of workflows to suit your project and coding style.

![The Actions Editor window, with some default actions provided by the Input System Package](../uploads/Main/InputManagerActionsEditor.png)

The Actions Editor window, with some default actions provided by the Input System Package

Get started with the [Input System Package](https://docs.unity3d.com/Packages/com.unity.inputsystem@latest).

## Legacy projects

If you are supporting a project that needs to use the older Input Manager, refer to the [legacy Input Manager documentation](InputLegacy.md).

## Additional resources and examples

* 👥 **Community**: [Join the conversation on Unity Discussions](https://discussions.unity.com/lists/input)
* 📺 **Video**: [Unity Input System in Unity 6 (1/7): Input Action Editor](https://youtu.be/TiTKAseu17A?list=PLX2vGYjWbI0RpLvO3B7aH-ObfcOifMD20)
* 📺 **Video**: [Unity Input System in Unity 6 (2/7): Input System Scripting](https://youtu.be/Cd2Erk_bsRY?list=PLX2vGYjWbI0RpLvO3B7aH-ObfcOifMD20)
* 📺 **Video**: [Unity Input System in Unity 6 (3/7): Input System Mobile controls](https://youtu.be/aI-r7ILNDug?list=PLX2vGYjWbI0RpLvO3B7aH-ObfcOifMD20)
* 📺 **Video**: [Unity Input System in Unity 6 (4/7): Input System and UI toolkit](https://youtu.be/GdjP5pggaHw?list=PLX2vGYjWbI0RpLvO3B7aH-ObfcOifMD20)
* 📺 **Video**: [Unity Input System in Unity 6 (5/7): Rebinding Input System controls](https://youtu.be/JfuqMaOiNPs?list=PLX2vGYjWbI0RpLvO3B7aH-ObfcOifMD20)
* 📺 **Video**: [Unity Input System in Unity 6 (6/7): Player Input Component](https://youtu.be/beDfIBLfx4c?list=PLX2vGYjWbI0RpLvO3B7aH-ObfcOifMD20)
* 📺 **Video**: [Unity Input System in Unity 6 (7/7): Player Input Manager and local multiplayer](https://youtu.be/lGxXQzE5Vu8?list=PLX2vGYjWbI0RpLvO3B7aH-ObfcOifMD20)