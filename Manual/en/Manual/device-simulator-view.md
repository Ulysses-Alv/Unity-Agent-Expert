# The Simulator view

The Simulator view displays your application on a simulated mobile device. Use it to see how your application appears with the screen shape, resolution, and orientation of that device.

![A screenshot of the Simulator view](../uploads/Main/device-simulator-view.png)

A screenshot of the Simulator view

## Using the Simulator view

To open the Simulator view, do one of the following:

* In the [Game view](GameView.md), in the top left corner, use the drop-down menu to switch between the Game view and the Simulator view.
* In the menu: **Window** > **General** > **Device Simulator**

Unity simulates the device in the Simulator view. To control the simulation, use the [toolbar](#toolbar) and the [Control Panel](#control-panel).

### Toolbar

The toolbar is at the top of the Simulator view and contains options that, along with the [Control Panel](#control-panel), control the simulation.

![The toolbar at the top of the Simulator view.](../uploads/Main/device-simulator-toolbar.png)

The toolbar at the top of the Simulator view.

| **Control** | **Description** |
| --- | --- |
| **Game/Simulator view** | Use this drop-down menu to switch between the Simulator view and the Game view. |
| **Device selection drop-down menu** | Choose the device to simulate from the available [device definitions](device-simulator-adding-a-device.md). |
| **Scale** | Zoom in or out on the simulated screen. |
| **Fit to Screen** | Scale the display to fit inside the window. |
| **Rotate** | Simulates the physical rotation of the device. The image on the device screen rotates together with the device if you enable auto-rotation and if the device supports rotation. Otherwise, if you rotate the device it causes the image to be sideways or upside down. **Note**: The Device Simulator doesn’t support gyroscope simulation. |
| **Safe Area** | Indicates whether the view draws a line around the [SafeArea](../ScriptReference/Screen-safeArea.md) of the current device screen. |
| **Play Mode behavior** | This section describes the play mode behavior based on your selection.   This property has the following options:  * **Focused**: Shifts focus on the selected Game view while the Editor is in Play mode. Only one Game view can be in focus when you enter Play mode. Using Maximized or Fullscreen on Display mode implies focus on the Maximized Game view. Enabling **Focused** on a Game view disables it on other Game views. * **VSync (Game view only)**: Allows syncing, which is useful when recording a video, for example. Unity attempts to render the Game view at the monitor refresh rate, though this is not guaranteed. When this option is enabled, it is still useful to maximize the Game view in Play mode to hide other views and reduce the number of views that Unity renders. |
| **Enter Play Mode:** | Choose from the options to determine the settings for the Editor when it enters the Play mode.   This property has the following options:  * **Normally:** Views the Game view without forcing focus or maximizing any views to full screen. * **Maximized**: Runs the Play mode with the Simulator view maximized to 100% of the Editor window. **Note**: This option doesn’t create a Simulator view if you disable the **Create Game View On Play** Editor preference. |
| **Control Panel** | Displays and hides the [Control Panel](#control-panel). |

### Control Panel

To open the Control Panel, click **Control Panel** in the top-right corner of the [Toolbar](#toolbar). By default, the Control Panel contains settings for the [Application simulated class](device-simulator-simulated-classes.md).

![The Device Simulator Control Panel.](../uploads/Main/device-simulator-control-panel.png)

The Device Simulator Control Panel.

| **Property** | **Description** |
| --- | --- |
| **System Language** | Specifies the value to receive from [Device.Application.systemLanguage](../ScriptReference/Device.Application-systemLanguage.md). |
| **Internet Reachability** | Specifies the value to receive from [Device.Application.internetReachability](../ScriptReference/Device.Application-internetReachability.md). |
| **On Low Memory** | Calls the [lowMemory](../ScriptReference/Device.Application-lowMemory.md) event. |

Device Simulator plugins can change the appearance of the control panel and add content and controls. Some packages and features, like [Adaptive Performance](adaptive-performance/adaptive-performance.md), contain Device Simulator plugins. Their UI appears in the Control Panel if you install the package or enable the feature. For more information, refer to [Device Simulator plugins](device-simulator-plugins.md).