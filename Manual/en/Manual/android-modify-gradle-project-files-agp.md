# Modify Gradle project files with the Android Project Configuration Manager

You can use Android Project Configuration Manager to modify the contents of custom **gradle** project files in your Android project. You cannot modify the contents of the default gradle project files `unityLibrary` and `launcher` modules with Android Project Configuration Manager. However, you can create custom modules to set up custom Gradle project files and modify them as required.

The entry point for the Android Project Configuration Manager is the [OnModifyAndroidProjectFiles](../ScriptReference/Android.AndroidProjectFilesModifier.OnModifyAndroidProjectFiles.md) method in the [AndroidProjectFilesModifier](../ScriptReference/Android.AndroidProjectFilesModifier.md) interface. This means to use the Android Project Configuration Manager, first create a class that implements `AndroidProjectFilesModifier` and declares a body for `OnModifyAndroidProjectFiles`. The following code example shows how to do this.

```
using UnityEditor.Android;

public class ModifyProjectScript : AndroidProjectFilesModifier
{
    public override void OnModifyAndroidProjectFiles(AndroidProjectFiles projectFiles)
    {
    }
}
```

## Additional resources

* [Export an Android project](android-export-process.md)
* [Modify Gradle project files with Gradle template files](android-modify-gradle-project-files-templates.md)
* [Modify Gradle project files with Android Studio](android-modify-gradle-project-files-android-studio.md)