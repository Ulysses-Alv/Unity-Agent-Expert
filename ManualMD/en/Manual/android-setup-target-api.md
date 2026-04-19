# Set up the Android SDK Target API

Configure the Android SDK Target API level for your application.

The Unity Hub installs the latest version of the Android SDK Target API that Google Play requires. If you need to use a more recent version, you can change it in the [Android Player Settings](class-PlayerSettingsAndroid.md).

## Configure Target API version

To configure the Android SDK Target API version in **Android **Player settings****, follow these steps:

1. Open **Android Player Settings** (menu: **Edit** > **Project Settings** > **Player**).
2. In the **Other Settings** section, change the **Target API Level**.

![Android Player Settings with the Target API Level dropdown selected.](../uploads/Main/android-sdk-target-api.png)

Android Player Settings with the Target API Level dropdown selected.

### Newer Target API version

If you select a Target API version newer than the latest installed version, the Unity Android SDK Updater can automatically download and install the new version. Unity displays a prompt and you can choose to either:

* Automatically download and install the new version of the Android SDK.
* Continue to use the highest installed version of the Android SDK.

### Older Target API version

If you select a target API version that’s not installed and is older than the latest installed version, the Unity Android SDK Updater can’t perform the update and Unity displays an error message.

In this case, to update the Android SDK Target API, you must use the Android [sdkmanager](https://developer.android.com/studio/command-line/sdkmanager) from either [Android Studio](https://developer.android.com/studio) or the [command-line tool](https://developer.android.com/studio/command-line/sdkmanager). Regardless of the method you choose, make sure to select the correct Android SDK folder for Unity in the **Android External Tools** section.

**Important**: On Windows, if you installed the Unity Editor in the default folder (`/Program Files/`), you must run the `sdkmanager` with elevated privileges (**Run as Administrator**) to perform the update.

## Additional resources

* [Install Android dependencies](android-install-dependencies.md)
* [Android External Tools reference](android-external-tools-reference.md)