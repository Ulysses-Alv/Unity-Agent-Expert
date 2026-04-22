# Gradle template variables

You can use the following variables in custom **Gradle** template files:

| **Variable** | **Description** |
| --- | --- |
| **ABIFILTERS** | Specifies the Application Binary Interfaces (ABIs) that your application should support. For example, `armeabi-v7a`, `arm64-v8a`. Gradle creates application builds for the specified `ABIFILTERS` values only. |
| **APIVERSION** | The API version to build for. Unity sets the **APIVERSION** and **TARGETSDK** to the same value (**Target API Level** in the [Android Player Settings](class-PlayerSettingsAndroid.md)). |
| **APPLICATIONID** | Android Application ID. For example, `com.mycompany.myapp`. |
| **APPLY\_PLUGINS** | Specifies a list of Gradle **plug-ins** to use during the build process. |
| **BUILDTOOLS** | The SDK build tools to use. |
| **BUILD\_SCRIPT\_DEPS** | Specifies a list of dependencies and repositories required during the build process. |
| **BUILTIN\_NOCOMPRESS** | Specifies a default list of file extensions to exclude from **compression**. The file extensions are: `.unity3d`, `.ress`, `.resource`, `.obb`, `.bundle`, `.unityexp`. |
| **DEBUGSYMBOLLEVEL** | Indicates the type of symbols package that contains debug metadata required for debugging your application. You can set the values to **none** for no debug metadata, **symbol\_table** for symbols package with only symbol tables, or **full** for symbols package with symbol tables and debugging information. |
| **DEFAULT\_CONFIG\_SETUP** | Includes additional configuration components for `android.defaultConfig`. |
| **DEPS** | The list of project dependencies. This is the list of libraries that the project uses. |
| **DIR\_GRADLEPROJECT** | The directory where Unity creates the Gradle project. |
| **DIR\_UNITYPROJECT** | The directory of your Unity project. |
| **EXTERNAL\_SOURCES** | The build script required to produce build artifacts such as, GameActivity and Swappy native libraries. |
| **GOOGLE\_PLAY\_DEPENDENCIES** | Specifies the Google Play services your application requires. For example, `com.google.android.gms:play-services-ads:23.2.0`. |
| **IL\_CPP\_BUILD\_SETUP** | The build script required to produce build artifacts related to **IL2CPP**, such as `libil2cpp.so`. |
| **LAUNCHER\_SOURCE\_BUILD\_SETUP** | Unity’s internal build script. |
| **LIBSDKTARGET** | The target API level the Unity library supports. |
| **MINIFY\_DEBUG** | Indicates whether to minify debug builds. |
| **MINIFY\_RELEASE** | Indicates whether to minify release builds. |
| **MINSDK** | The minimum API version that supports the application. |
| **NAMESPACE** | The application namespace. For example, `com.MyCompany.MyApp`. |
| **NDKPATH** | Specifies the Android Native Development Kit (NDK) installation folder path set in the **Android** section of **External Tools**, menu: **Edit** > **Preferences** > **External Tools** (macOS: **Unity** > **Settings** > **External Tools**). |
| **NDKVERSION** | The Android NDK (Native Development Kit) version Unity is using. For example, `ndkVersion "27.2.12479018"`. |
| **PACKAGING** | Specifies the required **APK** packaging options. |
| **PLAY\_ASSET\_PACKS** | Specifies the [asset packs](play-asset-delivery.md) to include in your application. |
| **SIGN** | Complete the `signingConfigs` section if this build is signed. |
| **SIGNCONFIG** | Indicates whether the build is signed. If this property is set to `signingConfig.release`, the build is signed. |
| **SOURCE\_BUILD\_SETUP** | Unity’s internal build script. |
| **SPLITS** | Indicates whether your application supports multiple APK builds. |
| **SPLITS\_VERSION\_CODE** | Build script that sets version code for split APKs. |
| **TARGETSDK** | The API version to target. Unity sets the and **APIVERSION** to the same value (**Target API Level** in the [Android Player Settings](class-PlayerSettingsAndroid.md)). |
| **USER\_PROGUARD** | Specifies a custom ProGuard file to use for minification. |
| **VERSIONCODE** | The internal version number for the application. It’s used to indicate how recent the application version is, where higher number denotes a more recent version. |
| **VERSIONNAME** | The application version number expressed as a string. This version number is visible to the users. |

Custom `settingsTemplate.gradle` files can also contain the following variables:

| **Variable** | **Description** |
| --- | --- |
| **INCLUDES** | The list of Android Library plug-ins to include in the Gradle project. |
| **ARTIFACTORYREPOSITORY** | Adds a reference to Unity’s maven repository only for internal use. Unity deletes this during the build process. |

Custom `gradleTemplate.properties` files can also contain the following variables:

| **Variable** | **Description** |
| --- | --- |
| **ADDITIONAL\_PROPERTIES** | Contains additional properties for the application. This includes:  • The Gradle template version. • The path to the Unity project. • If the application uses the [Android App Bundle](android-BuildProcess.md#publishing-format) publishing format, a flag that indicates to keep native libraries compressed. • If the application uses the [GameActivity](android-application-entries-game-activity.md) application entry point, a flag that indicates to use AndroidX. |
| **JVM\_HEAP\_SIZE** | The maximum size of the [Java virtual machine (JVM) heap](https://www.ibm.com/docs/en/integration-bus/10.0?topic=development-jvm-heap-sizing). |
| **STREAMING\_ASSETS** | The list of files in the [Steaming Assets](StreamingAssets.md) folder that Gradle shouldn’t compress. |

## Additional resources

* [Modify Gradle project files with Gradle template files](android-modify-gradle-project-files-templates.md)