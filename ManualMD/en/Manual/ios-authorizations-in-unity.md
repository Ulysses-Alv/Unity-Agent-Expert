# iOS authorizations in Unity

Apple’s operating system requires applications to request authorization before accessing sensitive information or device features. If your application requires access to features such as the device’s camera, microphone, or location, the device user must grant access to your application.

For more information on requesting authorizations on iOS devices, refer to the [Apple documentation](https://developer.apple.com/documentation/uikit/protecting_the_user_s_privacy/requesting_access_to_protected_resources).

To access the device features that your application requires, you must:

1. Provide a feature usage description in the **info.plist** file.
2. Send a request for authorization.

## Provide feature usage description

A feature usage description or purpose string is a message that the iOS system displays when your application tries to access the device feature. It’s best practice to add a clear description stating the reason for application access.

You can enter the feature usage description in **Configuration** section of [Player settings](class-PlayerSettingsiOS.md) (menu: **Edit** > **Project Settings** > **Player** > **Other Settings**). When you add the usage description to Player settings, Unity also adds it as a key-value pair to the **info.plist** file. For example, if your application wants to access the device camera, add Camera usage description in the Player settings. The description you enter automatically appears as **[NSCameraUsageDescription](https://developer.apple.com/documentation/bundleresources/information_property_list/nscamerausagedescription/) key** value in the **info.plist** file. For more information, refer to [Apple-specific iOS Player Settings](class-PlayerSettingsiOS.md#Configuration).

Alternatively, you can directly add the required feature usage description key-value pairs in the **info.plist** file.

## Send a request for authorization

After you add the feature usage description, send a request for camera or microphone access using [Application.RequestUserAuthorization](../ScriptReference/Application.RequestUserAuthorization.md) method. The system displays a dialog describing the request for authorization.

For a code example that shows how to use this API, refer to [Application.RequestUserAuthorization](../ScriptReference/Application.RequestUserAuthorization.md).

For information about accessing device location, refer to [LocationService](../ScriptReference/LocationService.md) API.

You can check the authorization status of your request using [Application.HasUserAuthorization](../ScriptReference/Application.HasUserAuthorization.md) method.

**Note**: The device user can modify the feature authorization through device settings at any given time. You can use [Application.HasUserAuthorization](../ScriptReference/Application.HasUserAuthorization.md) method to verify the authorization status of a feature before accessing it.