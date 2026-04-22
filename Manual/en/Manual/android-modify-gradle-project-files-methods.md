# Modify the Gradle project files for a Unity application

Unity provides [Player Settings](class-PlayerSettingsAndroid.md) and [Build settings](android-build-settings.md) to configure your application. When Unity builds your project, it takes these configuration options and uses them to generate [Gradle project files](android-gradle-overview.md#gradle-project-files). However, sometimes you might need more control over the **Gradle** project files.

Learn about the available methods that you can use to modify the contents of [Gradle project files](android-gradle-overview.md#gradle-project-files). Not every method is compatible with every Gradle project file. The following table shows which methods you can use to modify each Gradle project file.

| **Gradle project file** | **Gradle template** | **Android Studio** |
| --- | --- | --- |
| **Main Manifest** | Supported | Supported |
| **Android Launcher Manifest** | Supported | Supported |
| **Main Gradle** | Supported | Supported |
| **Launcher Gradle** | Supported | Supported |
| **Base Gradle** | Supported | Supported |
| **Gradle Properties Template** | Supported | Supported |
| **Proguard File** | Supported | Supported |

You can use Android Project Configuration Manager to modify custom Android Gradle project files. You cannot use this method to modify the Gradle project files built in the default `unityLibrary` and `launcher` modules.

## Gradle template files

Unity uses templates to produce the final Gradle project files. You can override these templates to control how Unity produces the final files. Gradle merges the manifests from your Android libraries into a final main manifest and makes sure that the final configuration is correct.

**Important**: If you use custom Gradle template files, be aware that if you upgrade your Unity project to a version of Unity that uses different default template files, you must rewrite your custom Gradle template files.

For information on how to use this method to modify Gradle project files, refer to [Modify Gradle project files with Gradle template files](android-modify-gradle-project-files-templates.md).

## Export to Android Studio

If you [export a Unity project for Android](android-export-process.md), Unity generates Gradle project files and places them in the exported project. If you open the exported project in [Android Studio](https://developer.android.com/studio/index.html), you can view the Gradle project files. This can be useful to verify modifications you made using Gradle templates or the Android Project Configuration Manager, and also useful to directly edit the files themselves.

**Tip**: To make sure you don’t need to re-modify the Gradle project files each time you export or build your Unity project, it’s best practice to perform the Gradle project file modifications that you want within Unity (using either Gradle template files or the Android Project Configuration Manager).

For information on how to use this method to modify Gradle project files, refer to [Modify Gradle project files with Android Studio](android-modify-gradle-project-files-android-studio.md).

## Android Project Configuration Manager

The Android Project Configuration Manager is a set of classes that you can use to set up and modify custom Gradle project files in C#. You cannot directly modify the Gradle project files that the build process creates in the default `unityLibrary` and `launcher` modules. You can create custom modules inside these modules to set up custom Gradle project files and modify them as required.

Unity applies the modifications during the build post-process, so you can check what values the Unity Editor sets and change them if you want. The entry point for the Android Project Configuration Manager is the [OnModifyAndroidProjectFiles](../ScriptReference/Android.AndroidProjectFilesModifier.OnModifyAndroidProjectFiles.md) method in the [AndroidProjectFilesModifier](../ScriptReference/Android.AndroidProjectFilesModifier.md) interface.

For information on how to use this method to modify Gradle project files, refer to [Modify Gradle project files with the Android Project Configuration Manager](android-modify-gradle-project-files-agp.md).

## Additional resources

* [Modify Gradle project files with Gradle template files](android-modify-gradle-project-files-templates.md)
* [Modify custom Gradle project files with Android Project Configuration Manager](android-modify-gradle-project-files-agp.md)
* [Modify Gradle project files with Android Studio](android-modify-gradle-project-files-android-studio.md)