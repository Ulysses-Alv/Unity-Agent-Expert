# Manually install an APK expansion file

Main and patch expansion files must be in a particular location on the device for the application to read from them. If you [Build and Run](android-BuildProcess.md) your application, Unity installs both the **APK** and the main APK expansion file on the device.

If you instead want to install the application manually, or want to install a patch expansion file, you must use the ****ADB**** utility to copy APK expansion files into the correct location on your device. For information on how to do this, refer to [Testing file reads](https://developer.android.com/google/play/expansion-files.html#TestingReading).

**Note**: The APK expansion file name must correspond to the format that Google requires. For more information, refer to [expansion files](https://developer.android.com/google/play/expansion-files.html).

## Additional resources

* [Host APK expansion files](android-apk-expansion-files-host.md)