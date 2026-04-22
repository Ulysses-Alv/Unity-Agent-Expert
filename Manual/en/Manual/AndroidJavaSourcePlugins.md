# Java and Kotlin source plug-ins

Unity can interpret individual Java and Kotlin source files as individual **plug-ins**.

Unity supports Java and Kotlin code written in source files with `.java` and `.kt` extensions. To do this, Unity interprets each source file as an individual plug-in and compiles them when it builds the Player. This type of plug-in is useful if you need to write a small amount of code for a single project. If you plan to reuse the code in multiple projects or distribute it to other people, then an [Android Library or Android Archive plug-ins](AndroidAARPlugins.md) might be more appropriate.

## Create a Java or Kotlin source plug-in

To indicate to Unity to create a plug-in from a Java (`.java`) or Kotlin (`.kt`) source file, follow these steps:

1. In the **Assets** folder, place your Java (`.java`) or Kotlin (`.kt`) source file.  
   **Tip**: It’s best practice to create a subfolder to contain your Java and Kotlin source files.
2. Select the source file and view it in the ****Inspector**** window.
3. In the **Inspector** window, under the **Select Platforms for plugin** section, enable **Android**.
4. Select **Apply**.

**Note**: You can place the source files in any folder in your project, except in special use locations such as `StreamingAssets`. If you place files in these locations, the Unity Editor doesn’t display the **Plugin Inspector**.

## Example: Create and use Kotlin source plug-in

The following example demonstrates how to create a Kotlin source plug-in and access its code from a C# script. This allows you to use Android-specific functionality in your Unity project.

1. In the **Assets** folder of your project, create a subfolder named **AndroidPlugins**, and add a Kotlin (`.kt`) file to it with the following code.

   ```
     object KotlinStringHelper {
     @JvmStatic
     fun getString(): String {
        return "Hello from Kotlin"
     }
     }
   ```
2. Select the Kotlin file. In the **Inspector** window, under the **Select Platforms for plugin** section, enable the Android platform.
3. Select **Apply**.
4. In Unity, create a C# script and add the following code. This code uses [AndroidJavaClass](../ScriptReference/AndroidJavaClass.md) to access the static Kotlin method `getString`.

   ```
     using UnityEngine;

     public class KotlinExamples : MonoBehaviour
     {
        void Start()
        {
           using (AndroidJavaClass cls = new AndroidJavaClass("KotlinStringHelper"))
           {
              string value = cls.CallStatic<string>("getString");
              Debug.Log($"KotlinStringHelper.getString returns {value}");
           }
        }
     }
   ```
5. Attach the C# Script to a **GameObject**. For more information, refer to the documentation on [Controlling a GameObject](class-MonoBehaviour.md).

## Edit Java or Kotlin files in an exported Android Studio project

By default, when you [export a Unity project for Android](android-export-process.md), Unity copies any Java/Kotlin files over to the Android Studio project. If you edit these files in Android Studio, the changes aren’t reflected in the original files in the Unity project. If you export the Unity project again, the export process overwrites your changes in Android Studio.

To resolve this, Unity provides the **Symlink Sources** [build setting](android-build-settings.md). If you select this build setting, Unity creates a [symbolic link](https://en.wikipedia.org/wiki/Symbolic_link) in the Android Studio project to Java/Kotlin files in the Unity project, instead of copying files over. This means that if you edit the files in Android Studio, the changes will reflect in the original Unity project files.

## Additional resources

* [Android plug-in types](android-plugin-types.md)
* [Calling Java and Kotlin plug-in code from C# scripts](android-plugins-java-code-from-c-sharp.md)