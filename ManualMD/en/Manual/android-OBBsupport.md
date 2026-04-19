# APK expansion files

APK expansion files are the asset splitting solution for the **APK** publishing format. They enable an application to split its assets into:

* Core assets like **scripts**, **plug-ins**, and assets that the application requires for the first **scene**.
* Additional assets like [Streaming Assets](StreamingAssets.md) and assets that the application requires for additional scenes.

The core assets go in the main APK file and the additional assets go in APK expansion files.

[Digital distribution services](android-distribution.md) often have an application size limit. This makes it necessary for APK applications that are larger than the size limit to use APK expansion files. For example, Google Play requires applications that are larger than 100MB to use APK expansion files. It allows you to use two APK expansion files, the main APK expansion file, and the patch APK expansion file, which can be up to 2GB each. For more information, see [APK Expansion Files](https://developer.android.com/google/play/expansion-files.html).

Refer to the following sections to learn about APK expansion files and how to work with them in Unity.

| **Topic** | **Description** |
| --- | --- |
| [APK expansion files in Unity](android-apk-expansion-files-in-unity.md) | Learn how APK expansion files work in Unity and how you can create main and optional patch APK expansion files for your application. |
| [Create the main APK expansion file](android-apk-expansion-files-in-unity.md#CreateMainAPK) | Split your application into the APK and the main APK expansion file. |
| [Create the patch APK expansion file](android-apk-expansion-files-in-unity.md#CreatePatchAPK) | Create the optional patch APK expansion file. |
| [Manually install an APK expansion file](android-apk-expansion-files-install.md) | Manually install an APK expansion file on a device for local testing. |
| [Host APK expansion files](android-apk-expansion-files-host.md) | Learn how to host APK expansion files for your application. |