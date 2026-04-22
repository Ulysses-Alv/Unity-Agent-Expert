# Introducing Android Library and Android Archive plug-ins

Android Archives are a compiled version of Android Libraries, and are the recommended way to format plug-ins that you want to distribute. However, while you create a **plug-in**, it’s faster to work with the Android Library format since that doesn’t require you to compile the plug-in outside of Unity and re-import the result. If you plan to modify the plug-in at all in the future, or want to iterate over it often, use an Android Library. After you finish development for the plug-in, compile it into an Android Archive.

## Android Library plug-ins

An Android Library is a directory with a specific structure that contains all the plug-in assets and the manifest.

When Unity creates the final **Gradle** project during the build or export process, it automatically includes all Android Library plug-ins in it and builds them together. Unity does this in the same way that Android Studio projects build when they have multiple subprojects.

## Android Archive plug-ins

An Android Archive (AAR) plug-in is a compiled version of an Android Library plug-in that you can use as a dependency for an Android app module. The `.aar` file itself is a `.zip` archive that contains all the compiled code, assets, and plug-in manifest. For more information on the structure of an AAR, refer to [Anatomy of an AAR file](https://developer.android.com/studio/projects/android-library.html#aar-contents).

## Providing additional Android Assets and resources

If you need to add Assets to your Unity application that should be copied as they are into the output package, include the raw assets in an Android Library plug-in or AAR. To access these assets, call the [getAssets](https://developer.android.com/reference/android/content/res/Resources.html#getAssets()) Android API from your Java code.

## Additional resources

* [Create an Android Library plug-in](android-library-plugin-create.md)
* [Import an Android Library plug-in](android-library-project-import.md)
* [Import an Android Archive plug-in](android-aar-import.md)