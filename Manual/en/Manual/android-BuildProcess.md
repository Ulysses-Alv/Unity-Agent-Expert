# Build your application for Android

Refer to the following instructions on how to build your Unity application for Android and considerations to be aware of when you do. For information on the build process for Android and the tools Unity uses, refer to [How Unity builds Android applications](how-unity-builds-android-applications.md).

Instead of building your application, you can also export the Unity project as a Gradle project and import that into Android Studio. This is useful if you want more control over the build pipeline, want to see or modify the [Android App Manifest](android-manifest.md) that Unity generates for your application, or [integrate Unity-powered features into another Android application](UnityasaLibrary-Android.md). For more information, refer to [Exporting an Android project](android-export-process.md).

Some digital distribution services that host Android applications have particular requirements that can change the build process. For example, [Google Play](android-distribution-google-play.md) requires your application to be an [Android App Bundle](https://developer.android.com/guide/app-bundle) (AAB) and not an **APK**. If you are targeting a specific digital distribution service with your build, refer to the documentation for that [Digital distribution service](android-distribution.md) first to check if the requirements differ.

Some digital distribution services have a limit on the initial install size of your application. Unity includes multiple methods that you can use to optimize the install size. For more information, refer to [Optimize distribution size](android-optimize-distribution-size.md).

If you want to build your application for debugging purposes, application patching can help you reduce the time it takes to build your application. For more information, refer to [Application patching](android-AppPatching.md).

## Configure the build

Before you create a build, configure your project’s settings so that Unity builds the application with the runtime settings and build system properties you want. There are two sets of settings that configure a Unity build:

* [Player settings](class-PlayerSettingsAndroid.md): Configure runtime and build settings for the application.
* [Build settings](android-build-settings.md): Configure build system settings.

### Publishing format

Unity can build Android applications in the following publishing formats:

* APK
* [Android App Bundle](https://developer.android.com/guide/app-bundle) (AAB)

By default, Unity builds Android applications in the APK publishing format. To make Unity build the Android application as an AAB:

1. Select **File** > **Build Profiles**.
2. From the list of platforms in the **Platform** pane, select **Android**.
3. Enable **Build App Bundle (Google Play)**. This setting is visible only when **Export Project** is disabled. If you want to [export the project](android-export-process.md) and build it in Android Studio, enable **Export Project**, then enable **Export for App Bundle**.

## Build the application

To build your Unity application for Android:

1. Select **File** > **Build Profiles**.
2. Select **Add **Build Profile**** to open the Platform Browser window.
3. From the list of platforms in the **Platform** pane, select **Android**.  
    If **Android** isn’t an option, select **Install with Unity Hub** and follow the installation instructions. Refer to [set up your project for Android development](android-sdksetup.md).
4. Select **Add Build Profile**.
5. Select **Switch Profile** to set the new build profile as the active profile.
6. Disable **Export Project**. If you want to export your project for Android Studio instead of building it within Unity, refer to [Exporting your Android project](android-export-process.md).
7. If you want to use the **Build and Run** option to immediately run the build on a target device, set **Run Device** to the device you want to run the application on. For information on how to connect your target device to Unity, refer to [Debug on Android devices](android-debugging-on-an-android-device.md).
8. Click either **Build** or **Build and Run**.
9. Select the destination for Unity to place the application. If you selected **Build and Run**, Unity also installs the application on the **Run Device**.
10. Click **Save**. This starts the build.

If you selected **Build and Run**, when Unity creates the build:

* If the **Split Application Binary** Player Setting is enabled and the **Build App Bundle (Google Play)** Build Setting is disabled, Unity builds [Android expansion files](android-OBBsupport.md) (OBB) for the APK and places them in the correct location on your device.
* If the ****Development Build**** Build Setting is enabled, Unity also sets up a Profiler tunnel and enables **CheckJNI** before it launches your application.

**Tip**: After you specify the output path the first time, you can use Ctrl+B (macOS: Cmd+B) keyboard shortcut to build and run the application.

## Application signing

Android applications must be digitally signed to run on an Android device. There are two types of application signing:

* Debug signing: The default signing method for a new Unity Project. Applications that use debug signing are able to run on an Android device, but you can’t publish them.
* Custom signing: The signing method that [Gradle](android-gradle-overview.md) uses when you provide custom signing information. Applications that use custom signing are able to run on an Android device and you can publish them.

To provide custom signing information, [create a keystore](android-keystore-create.md) and [load it](android-keystore-load.md) into Publishing Settings.

When you provide custom signing information, Unity doesn’t store keystores and key passwords on disk for security reasons. This means that you need to re-enter key passwords each time you restart the Unity Editor. If you don’t provide the passwords and attempt to build the application, the build process fails. To avoid entering passwords each time you open the Unity Editor, it’s best practice to only provide custom signing information when you want to build the application to publish. To create a build for testing on a device, don’t provide custom signing information and use debug signing instead.

## Additional resources

* [Sign your app](https://developer.android.com/studio/publish/app-signing).