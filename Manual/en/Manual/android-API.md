# Android mobile scripting

For cross-platform Projects, use the `UNITY_ANDROID` define directive to conditionally compile Android-specific C# code. For more information, refer to [Platform dependent compilation](platform-dependent-compilation.md).

## Access device-specific features and properties

Applications can access most features of an Android device through the [Input](../ScriptReference/Input.md) and [Handheld](../ScriptReference/Handheld.md) classes. For more information, see:

* [Mobile device input](MobileInput.md)
* [Mobile keyboard](MobileKeyboard.md)

### Vibration support

To trigger a vibration, call [Handheld.Vibrate](../ScriptReference/Handheld.Vibrate.md). Devices without vibration hardware ignore this call.

### Activity indicator

Mobile operating systems have built-in activity indicators your application can use during slow operations. For more information, refer to [Handheld.StartActivityIndicator](../ScriptReference/Handheld.StartActivityIndicator.md).

To access device-specific properties, use these APIs:

| **Script** | **Device property** |
| --- | --- |
| [SystemInfo.deviceUniqueIdentifier](../ScriptReference/SystemInfo-deviceUniqueIdentifier.md) | Always returns the md5 of `ANDROID_ID`. For more information, see Android developer documentation on [ANDROID\_ID](https://developer.android.com/reference/android/provider/Settings.Secure.html#ANDROID_ID). |
| [SystemInfo.deviceName](../ScriptReference/SystemInfo-deviceName.md) | Returns the device name. For Android devices, Unity tries to read `device_name` and `bluetooth_name` from secure system settings. If these strings have no values, Unity returns `<unknown>`. |
| [SystemInfo.deviceModel](../ScriptReference/SystemInfo-deviceModel.md) | Returns the device model. This often includes the manufacturer name and model number (for example, “LGE Nexus 5 or ”SAMSUNG-SM-G900A"). |
| [SystemInfo.operatingSystem](../ScriptReference/SystemInfo-operatingSystem.md) | Returns the operating system name and version. |