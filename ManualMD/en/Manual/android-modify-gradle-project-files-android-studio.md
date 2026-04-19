# Modify Gradle project files with Android Studio

If you don’t want to modify **Gradle** project files within Unity, you can export your Unity project as a Gradle project and manually modify the Gradle project files within Android Studio. From Android Studio, you can view and modify the final Gradle project files that Gradle uses to build your application. To modify Gradle project files this way:

1. [Export your Android project](android-export-process.md).
2. Open the exported project in Android Studio. To do this, refer to [Migrate to Android Studio](https://developer.android.com/studio/intro/migrate.html).
3. Find the Gradle project files in the project and edit them. For information on the structure of the exported project and where all the Gradle project files are, refer to [Gradle project structure](android-gradle-overview.md#gradle-project-structure).

## Additional resources

* [Modify Gradle project files with the Android Project Configuration Manager](android-modify-gradle-project-files-agp.md)
* [Modify Gradle project files with Gradle template files](android-modify-gradle-project-files-templates.md)