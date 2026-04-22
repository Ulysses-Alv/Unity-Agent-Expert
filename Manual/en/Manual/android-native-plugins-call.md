# Call native plug-in for Android code

The process to call code in native **plug-ins** for Android is the same as standard [native plug-ins](plug-ins-native.md).

**Note**: If you use individual C/C++ source files as plug-ins, use `__Internal` as the plug-in name in the [DllImport](https://docs.microsoft.com/en-us/dotnet/api/system.runtime.interopservices.dllimportattribute) attribute.

It’s best practice to wrap all native plug-in method calls with an additional C# code layer that:

* Checks [Application.platform](../ScriptReference/Application-platform.md) and only calls native methods when the application is running on Android devices using the architecture that you compiled the native plug-in for. On other platforms and architectures, the additional C# code layer should return dummy values.
* Uses [platform defines](platform-dependent-compilation.md) to control platform dependent code compilation and only compile code that uses the plug-in on platforms that have the plug-in available.

## Sample package

The [AndroidNativePlugin.unitypackage](../uploads/Examples/AndroidNativePlugin.zip) zip file contains a simple example of a native code plug-in distributed as a Unity package.

The sample shows how to invoke C++ code from a Unity application. The package includes a **scene** which displays the sum of two values as calculated by a native plug-in. To compile the plug-in, use [Android NDK](https://developer.android.com/ndk/index.html). For information on how to install Android NDK via the Unity Hub, see [Android environment setup](android-sdksetup.md).

To install the sample:

1. Download the zip file.
2. Extract the `AndroidNativePlugin.unitypackage` file.
3. In a Unity project, click **Assets** > **Import Package** > **Custom Package**.
4. In the **Import Package** file dialog, find and select the extracted `AndroidNativePlugin.unitypackage` file.