# Google Play delivery considerations

Consider the following key points and best practices before you publish an application to Google Play.

## Best practice checklist

To help launch an Android application successfully, Android provides best practices to follow. You can find these best practices in Android’s [Launch checklist](https://developer.android.com/distribute/best-practices/launch/launch-checklist) documentation.

### Public symbols

If your application crashes on a device, Google can use an [Android symbols](android-symbols.md) package to make a native stack trace human-readable on the [Android Vitals](https://developer.android.com/topic/performance/vitals) dashboard. It’s best practice to generate a [Public symbols](android-symbols.md#public-symbols) package for your application and upload it to Google Play. For information on how to do this, refer to [Generating a symbols package](android-symbols.md#generating-a-symbols-package).

### Deobfuscation file

Similar to symbol files, Unity can produce a deobfuscation file if you apply minification to your application build. For more information on applying minification, refer to [Android Player Settings](class-PlayerSettingsAndroid.md#minify). A deobfuscation file is automatically generated as a mapping file in the same location as your application build.

If you apply minification, it’s best practice to upload the deobfuscation file when publishing your application on Google Play. A deobfuscation file deciphers the method names in the stack trace, allowing you to identify and resolve the exact cause of the application crashes. For more information, refer to Google’s documentation on [Deobfuscate or symbolicate crash stack traces](https://support.google.com/googleplay/android-developer/answer/9848633)

## Additional resources

* [Google Play delivery requirements](android-distribution-google-play.md)