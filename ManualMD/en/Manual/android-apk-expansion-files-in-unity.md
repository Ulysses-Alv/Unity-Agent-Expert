# APK expansion files in Unity

This page describes how [APK expansion files](android-OBBsupport.md) work in the context of a Unity Android application.

Unity can automatically split an application into a main **APK** and a single APK expansion file that uses the `.obb` file extension. In Unity, this process is called splitting the application binary. The APK expansion file this process generates is the main APK expansion file and if you want to create the patch APK expansion, you must do that manually.

## Create the main APK expansion file

To create the main APK expansion file for your application, indicate to Unity to split the application. For information on how to do this, refer to [Splitting the application binary](class-PlayerSettingsAndroid.md#splitapplicationbinary).

Now when you [Build the application](android-BuildProcess.md), Unity generates the APK and the main APK expansion file then copies them both to the output directory. Unity uses the name of the application followed by `.main` for the name of the APK expansion file. For example, if the application is called `my-app`, the APK will be `my-app.apk` and the APK expansion file will be `my-app.main.obb`.

If the application starts and it can’t find and load the main APK expansion file, only the first ****Scene**** is available. In this case, try to download the APK expansion file. For more information, refer to [Download process](https://developer.android.com/google/play/expansion-files#DownloadProcess).

## Create the patch APK expansion file

Unity only automatically creates the main APK expansion file, but you can manually create the patch APK expansion file if the application requires more storage space. For information on how to create the patch APK expansion file, refer to Android’s [Development checklist](https://developer.android.com/google/play/expansion-files#Checklist) documentation.

### APK compatibility

When you create the patch expansion file, you must include a `unity_obb_guid` file within it so Android recognizes the APK expansion as being compatible with the APK. To do this:

1. Find the main APK expansion file that Unity generates. It’s a zip archive.
2. Extract the files from the APK expansion zip archive.
3. In the list of extracted files, find the `unity_obb_guid` file.
4. Copy this file over to the patch expansion file.

## Additional resources

* [Manually install an APK expansion file](android-apk-expansion-files-install.md)
* [Host APK expansion files](android-apk-expansion-files-host.md)