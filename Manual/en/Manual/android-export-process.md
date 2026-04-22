# Export an Android project

If you need more control over the build pipeline, you can export a Unity project as a **Gradle** project and import that into Android Studio. This is useful if you want more control over the build pipeline, want to see or modify the [Android App Manifest](android-manifest.md) that Unity generates for your application, or [integrate Unity-powered features into another Android application](UnityasaLibrary-Android.md).

## Exporting

To export a Unity project for Android Studio:

1. Open the **Build Profiles** window. (menu: **File** > **Build Profiles**).
2. From the list of platforms in the **Platforms** panel, select **Android** or [create a build profile](create-build-profile.md) for the **Android** platform.  
   **Note**: If Android is grayed out, [set up your project for Android development](android-sdksetup.md).
3. Enable **Export Project**.
4. Click **Export**.
5. Select the destination folder and click **Select Folder** to start the export process.

After Unity exports the Gradle project, you can import the Gradle project into Android Studio. For information on how to do this, refer to [Migrate to Android Studio](https://developer.android.com/studio/intro/migrate.html). For information on the file structure of the exported Gradle project, refer to [Gradle project structure](android-gradle-overview.md#gradle-project-structure).

## Additional resources

* [Gradle for Android](android-gradle-overview.md)
* [Build your application for Android](android-BuildProcess.md)