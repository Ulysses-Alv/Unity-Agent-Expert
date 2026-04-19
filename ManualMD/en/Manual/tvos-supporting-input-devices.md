# Supporting input devices on tvOS

While tvOS builds on the foundation of iOS, it does create new challenges such as adapting content to function with tvOS inputs, and for display on a bigger screen.

There are two main inputs for tvOS:

* The Apple TV Remote
* Made For iOS (MFi) controllers

## Apple TV Remote

The Apple TV Remote (Siri Remote) is a multi-purpose input device that works as a traditional menu navigation controller, app controller, gyroscope, acceleration sensor, and as a touch gesture device. Unity routes Apple TV Remote input to corresponding Unity APIs, but performs no other processing on that input. Your application might need some adjustments to its input scheme to leverage the Apple TV Remote’s specific input features. For instance, your application can treat it as a traditional application controller, with one analog axis and an extra action button, or your application can use the accelerometer for interactions such as steering. You can experiment with various schemes when porting an app to tvOS.

## Made for iOS (MFi)

Unity provides Made For iOS (MFi), which is a standardized controller support for iOS and tvOS. MFi controllers offer out of the box input mappings, and you can set up custom action mappings from **Edit** > **Project Settings** > **Input Manager**. For more information, refer to [Handle Game Controller input](ios-handle-game-controller-input.md) and [Game Controllers](https://developer.apple.com/design/human-interface-guidelines/tvos/remote-and-controllers/game-controllers/).

Two further wireless MFi controllers can be connected to an Apple TV device, which effectively turns it into a console. Your application can use the controllers in the same way as iOS MFi controllers, but you must ensure its usability with the Apple TV Remote alone. The tvOS system limits the number of additional controllers to two.

Here are some technical details on accessing specific TV Remote features:

| **Feature** | **Description** |
| --- | --- |
| **Touch area** | Maps to both [Input.touches](../ScriptReference/Input-touches.md) ([Touch.type](../ScriptReference/Touch-type.md) is set to [Indirect](../ScriptReference/TouchType.Indirect.md) and is ignored by the Unity GUI), and the Joystick Input API (for example, [Input.GetAxis(“Horizontal”)](../ScriptReference/Input.GetAxis.md)). |
| **Touch area click** | Maps to button A, which then maps to [joystick button 14](ios-handle-game-controller-input.md#InputMapping) |
| **Gyroscope** | Maps to [Input.gyro](../ScriptReference/Input-gyro.md). [Input.gyro.attitude](../ScriptReference/Gyroscope-attitude.md) derives from the gravity vector, and as such it doesn’t rotate around the axis parallel to the gravity vector. The same applies for [Input.gyro.rotationRate](../ScriptReference/Gyroscope-rotationRate.md). |
| **Acceleration** | Maps to [Input.acceleration](../ScriptReference/Input-acceleration.md).   **Note**: [Input.acceleration](../ScriptReference/Input-acceleration.md) derives from the gyroscope API and might have some instabilities. The tvOS SDK doesn’t have a dedicated accelerometer API. |
| **Pause/Play button** | Maps to button X, which then maps to [joystick button 15](ios-handle-game-controller-input.md#InputMapping) |
| **Menu button** | A long press calls the tvOS task switcher. You can’t override this behavior.  Your app can process short taps one of two ways:  **a)** Return to the tvOS system home screen, if [UnityEngine.tvOS.Remote.allowExitToHome](../ScriptReference/tvOS.Remote-allowExitToHome.md) is true.  **b)** Let your app respond to taps (which maps to the Pause button/[joystick button 0](ios-handle-game-controller-input.md#InputMapping) when [UnityEngine.tvOS.Remote.allowExitToHome](../ScriptReference/tvOS.Remote-allowExitToHome.md) is false. This is the default behavior.  Your app should switch between **a)** and **b)** based on its current state:  - If the user is currently interacting with the top menu, enable behavior **a)**.  - If the user is interacting with the app in real time, enable behavior **b)** and call the in-app pause menu when they press this button. |
| **Swipe to the edge of the remote** | Generates directional pad (D-pad) up/down/left/right button presses.  For a list of mappings, refer to [Game Controller input mapping](ios-handle-game-controller-input.md#InputMapping). |

You can control the Apple TV Remote operational modes via a dedicated API as follows:

* [UnityEngine.tvOS.Remote.allowExitToHome](../ScriptReference/tvOS.Remote-allowExitToHome.md)
* [UnityEngine.tvOS.Remote.allowRemoteRotation](../ScriptReference/tvOS.Remote-allowRemoteRotation.md)
* [UnityEngine.tvOS.Remote.reportAbsoluteDpadValues](../ScriptReference/tvOS.Remote-reportAbsoluteDpadValues.md)
* [UnityEngine.tvOS.Remote.touchesEnabled](../ScriptReference/tvOS.Remote-touchesEnabled.md)

**Note**: When [UnityEngine.tvOS.Remote.allowExitToHome](../ScriptReference/tvOS.Remote-allowExitToHome.md) is false, the Menu button maps to [joystick button 0](ios-handle-game-controller-input.md#InputMapping). This causes a conflict with the default [Input](class-InputManager.md) window, because it also uses [joystick button 0](ios-handle-game-controller-input.md#InputMapping) to map the Submit virtual button. This results in the Menu button triggering actions on UI elements. To fix this issue, remove or modify the Submit virtual button bindings in the [Input](class-InputManager.md) window (menu: **Edit** > **Project Settings**, then select the **Input** category).