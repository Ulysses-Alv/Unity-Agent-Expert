# Gradle templates

Gradle templates configure how to build an Android application using [Gradle](android-gradle-overview.md). Each Gradle template represents a single Gradle project. Gradle projects can include, and depend on, other Gradle projects.

## Gradle template files

A Gradle template consists of the following files:

| File | Location | Contains |
| --- | --- | --- |
| `baseProjectTemplate.gradle` | In the exported project, `root/build.gradle` folder | Configuration information that affects all modules in the final Gradle project. It specifies which Android Gradle Plugin version to use and locations of java plugins. The locations are a combination of online repositories and java plugins inside of this project. |
| `launcherTemplate.gradle` | In the exported project, `root/launcher/build.gradle` folder | Instructions on how to build the Android application. This includes bundling, signing, and whether to split the **apk**. It depends on the unityLibrary project and outputs either an .apk file or an app bundle. |
| `mainTemplate.gradle` | In the exported project, `root/unityLibrary/build.gradle` folder | Instructions on how to build Unity as a Library. This outputs an .aar file. You can override the Unity template with a custom template in the Unity Editor. Refer to the Providing a custom Gradle build template section on this page for more details. |
| `libTemplate.gradle` | Varies | If an [Android Library](AndroidAARPlugins.md) **plug-in** doesn’t include a `build.gradle` file, Unity uses the `libTemplate.gradle` file as a template to generate one. After Unity generates the `build.gradle` file, or if one already exists in the plug-in’s directory, Unity copies the plug-in into the Gradle project. |
| `settingsTemplate.gradle` | In the exported project, `root/settings.gradle` file | Specifies the names of modules that the Gradle build system should include when it builds the project. You can override the Unity template with a custom template in the Unity Editor. Refer to the Providing a custom Gradle build template section on this page for more details. |
| `gradleTemplate.properties` | In the exported project, `root/gradle.properties` file | Configures the Gradle build system and specifies properties such as the size of the [Java virtual machine (JVM) heap](https://www.ibm.com/docs/en/integration-bus/10.0?topic=development-jvm-heap-sizing) |

To have more control over the Gradle project files that Unity produces, you can override Unity’s default Gradle template files. For information on how to do this, refer to [Modify Gradle project files with Gradle template files](android-modify-gradle-project-files-templates.md).

## Modifying the exported Gradle project using C#

To modify the Gradle project after Unity assembles it, create a class that inherits from [IPostGenerateGradleAndroidProject](../ScriptReference/Android.IPostGenerateGradleAndroidProject.md) and override the [OnPostGenerateGradleAndroidProject](../ScriptReference/Android.IPostGenerateGradleAndroidProject.OnPostGenerateGradleAndroidProject.md) function. This function receives the path to the unityLibrary module as a parameter and you can use it to reach the application’s manifest and resources through C# scripting.