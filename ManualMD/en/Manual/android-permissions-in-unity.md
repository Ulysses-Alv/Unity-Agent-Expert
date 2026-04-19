# Android permissions in Unity

To get permission to access device features or data outside of your Unity application’s sandbox, there are two stages:

1. At build time, declare the permission in the application’s [Android App Manifest](android-manifest.md).
2. At runtime, request permission from the user.

For some permissions, Unity automatically handles both the build-time Android App Manifest entries and runtime permission requests. For more information, see [Unity-handled permissions](#unity-handled-permissions).

## Declare permissions

An Android application can only request the user for permission to use device features or data that it declares in its Android App Manifest. This is the build-time stage of Android permission handling. For information on how to declare permissions, see [Declare permissions for an application](android-permissions-declare.md).

**Note**: If a **plug-in** declares a permission in its Android App Manifest, **Gradle** automatically adds the permission to the final Android App Manifest during the Gradle merge stage

## Runtime permissions

You can request runtime permissions at any time while the application is running. It’s best practice to request permission for restricted data or a device feature when the application requires the data/feature for the first time. You should also show the user a message that explains why the application requires the data/feature before you send the request. This is important to give users context so they are less likely to deny the permission out of confusion.

For information on how to request permission at runtime, see [Request runtime permissions](android-RequestingPermissions.md).

**Important**: If you try to request the user for a permission that the Android App Manifest doesn’t declare, the device doesn’t show the permission dialogue.

## Unity-handled permissions

Depending on the [Player Settings](class-PlayerSettingsAndroid.md) and Unity APIs that the application uses, Unity automatically adds some required permissions to the Android App Manifest and sends permission requests at runtime. This also applies to any plug-ins that the application uses.

To stop Unity from automatically sending permission requests at runtime, add `unityplayer.SkipPermissionsDialog` to the Android App Manifest. This is useful if you want to provide information about why the application requires each permission before the user sees the permission request. In this case, you need to send the permission requests manually. For more information, see [Request runtime permissions](android-RequestingPermissions.md).

The following list explains which permissions Unity handles automatically:

* If the application uses vibration such as [Handheld.Vibrate](../ScriptReference/Handheld.Vibrate.md), Unity adds the `VIBRATE` permission to the Android App Manifest.
* If the application uses the [Application.internetReachability](../ScriptReference/Application-internetReachability.md) property, Unity adds the `ACCESS_NETWORK_STATE` permission to the Android App Manifest.
* If the application uses a location API such as [LocationService](../ScriptReference/LocationService.md), Unity adds the `ACCESS_FINE_LOCATION` permission to the Android App Manifest and requests permission from the user the first time that the application uses location.
* If the application uses the [WebCamTexture](../ScriptReference/WebCamTexture.md) class Unity adds the `CAMERA` permission to the Android App Manifest and requests permission from the user the first time that the application uses `WebCamTexture`.
* If the application uses the [Microphone](../ScriptReference/Microphone.md) class adds the `RECORD_AUDIO` permission to the Android App Manifest and requests permission from the user the first time that the application uses `Microphone`.

## Additional resources

* [Declare permissions for an application](android-permissions-declare.md)
* [Request runtime permissions](android-RequestingPermissions.md)