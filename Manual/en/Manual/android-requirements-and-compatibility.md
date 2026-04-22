# Android requirements and compatibility

Before you begin to develop an Android application in Unity, check Unity’s requirements and compatibility information for Android to make sure you’re aware of any limitations for developing a Unity application for this platform.

## Android support

Unity supports Android 7.1 “Nougat” (API level 25) and above. For more information, refer to [AndroidSdkVersions](../ScriptReference/AndroidSdkVersions.md).

You can now target Android API levels 35 and 36 in your Unity projects.

## Graphics API support

Android devices support [Vulkan](https://developer.android.com/ndk/guides/graphics) and [OpenGL ES](https://developer.android.com/guide/topics/graphics/opengl). This section contains information about the graphics APIs Unity supports for Android.

| **Graphics API** | **Support** |
| --- | --- |
| **Vulkan** | Yes |
| **OpenGL ES 1.0** | No |
| **OpenGL ES 1.1** | No |
| **OpenGL ES 2.0** | No |
| **OpenGL ES 3.0** | Yes |
| **OpenGL ES 3.1** | Yes |
| **OpenGL ES 3.2** | Yes |

## Render pipeline compatibility

Not every [render pipeline](render-pipelines.md) is compatible with Android due to hardware and graphics API limitations.

| **Feature** | **Built-in Render Pipeline** | **Universal Render Pipeline** | **High Definition Render Pipeline** | **Custom Scriptable Render Pipeline** |
| --- | --- | --- | --- | --- |
| **Android** | Yes | Yes | No | Yes |

## Manifest element attributes

This section contains compatibility information on [Android App Manifest element](https://developer.android.com/guide/topics/manifest/manifest-intro#reference) attributes.

* For the [<Activity>](https://developer.android.com/guide/topics/manifest/activity-element.html) element, Unity only supports the `singleTask` [launchMode](https://developer.android.com/guide/topics/manifest/activity-element.html#lmode).

## Emulator compatibility

Unity doesn’t support Android emulators. To test your application, you can:

* [Test on an Android device](android-debugging-on-an-android-device.md).
* If you only need to test mobile input for your application, use [Unity Remote](UnityRemote5.md).
* If you only need to test the appearance of an Android device, use the [Device Simulator](android-device-simulator.md).

## Texture compression

The standard texture **compression** formats on Android are [Ericsson Texture Compression (ETC)](https://en.wikipedia.org/wiki/Ericsson_Texture_Compression) and [Adaptable Scalable Texture Compression (ASTC)](https://www.khronos.org/opengl/wiki/ASTC_Texture_Compression). To target the widest range of Android devices, use one of these **texture compression** formats. Unity’s default texture compression format is ASTC. If an Android device doesn’t support the texture compression format you use for a texture, Unity decompresses the texture at runtime. This increases memory usage and decreases rendering speed.

A subset of Android devices supports the DXT texture compression format. This format supports textures with an alpha channel as well as high compression rates or high image quality. For digital distribution services that filter content based on texture compression format, it is best practice to create separate builds of your application for each texture compression format.

There are two ways to change the default texture compression format for your application:

* In [Android Player Settings](class-PlayerSettingsAndroid.md) with the **Texture compression formats** setting. For more information, refer to [texture compression targeting](android-distribution-google-play.md#texture-compression-targeting).
* In [Android build settings](android-build-settings.md) with the **Texture Compression** setting. The default value for this is **Use **Player Settings****.

The value you set in build settings has priority over the one you set in Player Settings. Use it to change the texture compression format for a particular build.

You can also customize the texture compression format for individual textures. The value you set for an individual **texture overrides** the default texture compression format value. For information on how to change the **texture format** of individual textures, see [Texture Importer](class-TextureImporter.md).

## Playing video files

This section provides additional information for playing video files on Android:

* To play video files on Android, use the [Video Player](VideoPlayer.md) component. If your application tries to play a video file that the device doesn’t support, Unity doesn’t play the video.
* You can use any resolution or number of audio channels so long as the target device supports them. **Note**: Not all devices support resolutions greater than 640 × 360.
* Unity supports playback from uncompressed asset bundles. For [Android Pie](https://en.wikipedia.org/wiki/Android_Pie) and above, Unity supports playback from compressed asset bundles.
* Unity doesn’t support native webM/VP8 transparency. To play VP8-encoded webM clips with transparency, transcode the clips to a supported format.
* In Android versions prior to `6.0.1`, videos with transparency that have a higher resolution than the device support render **pixels** outside the supported resolution as white.
* Unity reports format compatibility issues in the `adb logcat` output and prefixes them with `AndroidVideoMedia`. This file might display other device-specific error messages near the video format issues Unity reports. These device-specific errors aren’t visible to Unity and often explain what the compatibility issue is.

## 16 KB memory page size support

Unity supports devices with 16 KB memory page sizes, a feature introduced in Android 15. For more information about the feature, refer to the Android documentation on [Support 16 KB page sizes](https://developer.android.com/guide/practices/page-sizes?_gl=1*86rcxf*_up*MQ..*_ga*MTMxMTA4MTU5Ny4xNzM3OTk1NjU1*_ga_6HH9YJMN9M*MTczNzk5NTY1NC4xLjAuMTczNzk5NTY1NC4wLjAuMTA3NjE4ODIy).

Unity applications built for devices with 4 KB memory page size might not work on devices with 16 KB memory page size. To ensure your application is compatible with the devices that use 16 KB memory page size, consider the following key points:

* Update Unity version to the most recent patch. Update the native **plug-ins**, rebuild, and resubmit your application to Google Play. The same application executable will be able to support devices with both 4 KB and 16 KB memory page sizes. For more information, refer to Android documentation on [Build your app with support for 16 KB devices](https://developer.android.com/guide/practices/page-sizes#build).
* If your project contains a plug-in with a `.so` file that’s aligned to 4 KB instead of 16 KB, the Unity Editor displays a warning during the build process.
* If your application uses third-party plug-ins or SDKs that link to native libraries, contact the plug-in creators to confirm that their native libraries are compatible with 16 KB page sizes. Update the relevant SDKs and resubmit your application.
* For guidelines on testing your application on devices with 16 KB memory page sizes, refer to the Android documentation on [Enable 16 KB mode on a device using developer options](https://developer.android.com/guide/practices/page-sizes#developer-option).