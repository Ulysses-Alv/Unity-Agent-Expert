# Screen configuration

Unity provides features that you can use to configure the screen when in the Editor and at runtime.

## Screen orientation

You can control the screen orientation of your application on Android devices. Detecting a change in orientation or forcing a specific orientation is useful for creating behaviors that depend on how the user holds the device.

To retrieve the current application orientation, access [Screen.orientation](../ScriptReference/Screen-orientation.md) property. The available orientations are:

| **Orientation** | **Behavior** |
| --- | --- |
| **Portrait** | The application is in portrait mode and expects the device to be upright and the home button at the bottom. |
| **PortraitUpsideDown** | The application is in portrait mode but upside down and expects the device to be upright and the home button at the top. |
| **LandscapeLeft** | The application is in landscape mode and expects the device to be upright and the home button on the right side. |
| **LandscapeRight** | The application is in landscape mode and expects the device to be upright and the home button on the left side. |

To manually control the screen orientation, set [Screen.orientation](../ScriptReference/Screen-orientation.md) to one of the above orientations, or use [ScreenOrientation.AutoRotation](../ScriptReference/ScreenOrientation.AutoRotation.md). When you enable autorotation, you can still disable some orientation on a case-by-case basis.

The following properties control autorotation:

* [Screen.autorotateToPortrait](../ScriptReference/Screen-autorotateToPortrait.md)
* [Screen.autorotateToPortraitUpsideDown](../ScriptReference/Screen-autorotateToPortraitUpsideDown.md)
* [Screen.autorotateToLandscapeLeft](../ScriptReference/Screen-autorotateToLandscapeLeft.md)
* [Screen.autorotateToLandscapeRight](../ScriptReference/Screen-autorotateToLandscapeRight.md)

The screen orientation of your application may vary depending on the device orientation set by the user, regardless of your application’s default orientation setting.

The following table shows the default screen orientation set for your application and the actual orientation the application loads in based on the device orientation.

|  | **Default screen orientation: Autorotate** | **Default screen orientation: Portrait** | **Default screen orientation: Landscape** |
| --- | --- | --- | --- |
| **Device orientation: Autorotate** | The application screen loads in portrait, and can rotate between portrait, portrait upside down, landscape right, and landscape left orientations. **Note**: On a tablet, the application screen loads in portrait or landscape orientation based on the device orientation. | The application screen loads in portrait and can rotate between portrait and portrait upside down orientations. | The application screen loads in the same orientation as the device orientation. If the device orientation is portrait, the application screen loads in portrait, and can rotate to landscape orientation. In landscape orientation, the screen can only rotate between landscape right or landscape left orientations. **Note**: On a tablet, if the user turns off the device autorotation setting, and then reopens the application, the application screen orientation remains unchanged. |
| **Device orientation: Portrait lock** | The application screen loads and remains locked in portrait orientation.  **Note**: Whilst in portrait orientation, if you set the [ScreenOrientation.AutoRotation](../ScriptReference/ScreenOrientation.AutoRotation.md) property for portrait to false, the application screen remains in portrait until you rotate the device to landscape orientation. Once in landscape, the screen can’t rotate back to portrait orientation. If you now set the [ScreenOrientation.AutoRotation](../ScriptReference/ScreenOrientation.AutoRotation.md) property for landscape to false, the screen orientation remains unchanged. | The application screen loads and remains locked in portrait orientation. | The application screen loads and remains locked in portrait orientation even if you set the portrait option to false. To allow screen rotation, assign the required orientation to [Screen.orientation](../ScriptReference/ScreenOrientation.md) property. You can enable autorotation with [ScreenOrientation.AutoRotation](../ScriptReference/ScreenOrientation.AutoRotation.md). |
| **Device orientation: Landscape lock** | The application loads and remains locked in landscape orientation. | The application loads in landscape and remains locked in landscape left or landscape right orientation depending on the device orientation even if you set either option to false.  To change the orientation, assign the required orientation to [Screen.orientation](../ScriptReference/ScreenOrientation.md) property. You can enable autorotation with [ScreenOrientation.AutoRotation](../ScriptReference/ScreenOrientation.AutoRotation.md). | The application loads in landscape and remains locked in landscape left or landscape right orientation depending on the device orientation. |

## Multi-window mode

Android’s [multi-window mode](https://developer.android.com/guide/topics/large-screens/multi-window-support) supports Unity applications. Users can resize the windows that will contain your Unity applications so it’s best practice to make your user interface scale to non-standard **aspect ratios** and resolutions.