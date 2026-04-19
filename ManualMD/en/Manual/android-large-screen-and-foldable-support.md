# Large screen and foldable device support

On large screen and foldable devices running Android 12 or newer, your application can run simultaneously with other applications in multi-window mode, irrespective of its default configuration. If your application is designed for a specific orientation or **aspect ratio**, or is non-resizable, the Android OS automatically switches it into [compatibility mode](https://developer.android.com/guide/topics/large-screens/large-screen-compatibility-mode). In this mode, the Android OS scales your application to fit the available screen size. However, in some cases, the application might not be optimized for a specific screen size, resulting in a suboptimal user experience.

**Note**: For devices running Android 12 or newer versions, [multi-window mode](https://developer.android.com/guide/topics/large-screens/multi-window-support) is the default behavior. Starting with Android 16, for applications targeting Android 16 and newer versions, the OS introduces additional changes for large screen devices. For more information, refer to the documentation on [Android 16 behavior changes for large screen devices](#Android16).

To ensure consistent user experience, consider designing your application to adapt to various screen sizes, including large screen and foldable devices. You can use the following settings and APIs.

## Resizeable Activity Player setting

The [**Resizeable Activity**](class-PlayerSettingsAndroid.md#resizeable) Player setting enables multi-window mode in your application and allows the Android OS to consider it resizable. For new projects, Unity enables this setting by default and sets `android.resizeableActivity="true"` in the Android manifest file. In this case, the aspect ratio restrictions are ignored.

If you choose to set `android.resizeableActivity="false"`, the Android OS behaves in the following way:

| **Android OS version** | **Behavior** |
| --- | --- |
| **Older than Android 12** | Displays your application as full-screen and disables multi-window mode. |
| **Android 12 and newer** | Enables multi-window mode by default and activates the compatibility mode for your application to adjust it to the available screen dimensions.   **Note**: For applications targeting Android 16, the Android OS ignores the `resizeableActivity` manifest attribute and adjusts your application to fit the available screen space. For more information, refer to the documentation on [Android 16 behavior changes for large screen devices](#Android16). |

In the compatibility mode, the application might not display as expected. In scenarios such as transitioning between screens on large screens and foldable devices, the OS might prompt the user to restart the application. Restarting the application causes loss of activity and the overall application state.

**Note**: On small-screen devices running Android 12 or newer versions, a resizable application supports multi-window mode based on the [minimum width and minimum height](class-PlayerSettingsAndroid.md#minWidthHeight) of the activity. A non-resizable application doesn’t support multi-window mode on small-screen devices.

## Application and Configuration APIs

The [AndroidApplication API](../ScriptReference/Android.AndroidApplication.md) provides information on your Android application while it runs on a device. You can use this API to perform the following actions:

* Access the current activity’s Java instance.
* Invoke events on the UI thread or main thread of your application.
* Detect and handle configuration changes while your application runs on a device.

Using the [AndroidApplication.onConfigurationChanged](../ScriptReference/Android.AndroidApplication-onConfigurationChanged.md) method, you can detect any device configuration changes. You can retrieve device information, such as screen layout, screen size, and orientation. The method also retrieves additional configuration details, such as whether the screen is folded or unfolded, the type of keyboard in use, and user preferences for language and region. For more information on which device properties you can retrieve, refer to [AndroidConfiguration API](../ScriptReference/Android.AndroidConfiguration.md). Use the [AndroidApplication.onConfigurationChanged](../ScriptReference/Android.AndroidApplication-onConfigurationChanged.md) method to make necessary adjustments, such as rearranging the user interface, to ensure your application adapts seamlessly to different device configurations.

## Android 16 behavior changes for large screen devices

Starting with Android 16, applications targeting Android 16 or later versions are required to adapt their content dynamically on various screen sizes for optimal user experience. For large screen devices, the Android OS will ignore any fixed settings for orientation, resizability, or aspect ratio defined in manifest attributes or through APIs. Instead, the OS will try to adjust your application’s content to fit the available screen space.

For more information about which attributes are ignored, refer to the [Android documentation](https://developer.android.com/about/versions/16/behavior-changes-16#implementation-details). Ensure that your application supports these changes to avoid any layout issues.

### Exemptions

The Android 16 behavior changes do not apply in the following cases:

* Applications categorized as games
* Applications running on devices with the smallest screen width less than 600dp
* Applications running on devices where user manually selects application’s default behavior in the device’s aspect ratio settings

To enable exemption for your game application, Unity provides the **App Category** Player setting. This setting replaces the [**androidIsGame**](../ScriptReference/PlayerSettings.Android-androidIsGame.md) API, which is no longer supported.

### App Category Player setting

The **App Category** setting allows you to specify the category for your application, such as a game or any other app type. For applications categorized as games, enabling this setting ensures an exemption from the Android 16 behavior changes. This means the Android OS will not ignore any fixed orientation, resizability, or aspect ratio settings defined for your application.

The **App Category** must be set to **Game**, if you’re using [Adaptive Performance](adaptive-performance/adaptive-performance.md) or [GameMode API](../ScriptReference/Android.AndroidGame.GameMode.md). To do so, follow these steps:

1. Open **Android Player settings** > **Other Settings** > **Configuration**.
2. Enable **App Category**.
3. In the category dropdown, select **Game**.

After the steps are complete, Unity automatically adds `android:appCategory=game` in the application manifest file during the build process.

**Notes**:

* For new projects, the **App Category** is enabled and set to **Game** by default.
* For existing projects with [**androidIsGame**](../ScriptReference/PlayerSettings.Android-androidIsGame.md) set to false, the **App Category** remains disabled.

## Additional resources

* [Android documentation on resizeableActivity](https://developer.android.com/guide/topics/manifest/application-element#resizeableActivity)
* [Behavior changes: Apps targeting Android 16 or higher](https://developer.android.com/about/versions/16/behavior-changes-16)