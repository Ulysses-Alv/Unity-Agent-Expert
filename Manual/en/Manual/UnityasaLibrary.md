# Using Unity as a Library in other applications

Unity as a Library is intended for specialist users who use native platform technologies such as Java/Android, Objective C/iOS, or Windows Win32/UWP, and want to include Unity-powered features in their games or applications.

This documentation assumes that you have experience with developing for native platform technologies such as Java/Android, Objective C/iOS, or Windows Win32/UWP, and that you are familiar with the structure of the project, language features and specific platform configuration options (for example, user permissions).

You can use Unity as a Library in other applications by integrating your content and the Unity runtime components in a native platform project. This enables you to embed content that uses 3D or 2D real-time rendering, like **AR** experiences, interaction with 3D models, and 2D mini-games. The Unity Runtime Library exposes ways to manage loading, activating, and unloading within the native application.

The following platforms currently support Unity as a Library:

* [Android](UnityasaLibrary-Android.md)
* [iOS](UnityasaLibrary-iOS.md)
* [Windows](UnityasaLibrary-Windows.md) and [Universal Windows Platform](UnityasaLibrary-UWP.md)

To determine platform versions and other dependencies, see the [system requirements](system-requirements.md) page.

## Limitations

When hosted by another application, Unity doesn’t control the runtime lifecycle, so it might not work in all scenarios. Known limitations include:

* On Android and iOS:
  + Only full-screen rendering is supported. However, if you are a Unity Industry customer, the limitations and features might differ.
  + When Unity is in an unloaded state (after calling [`Application.Unload`](../ScriptReference/Application.Unload.md)), it retains some amount of memory (between 80–180Mb) to be able to instantly switch back and run again in the same process. The amount of memory that’s not released largely depends on the device’s graphics resolution.
* On iOS, if the Unity runtime quits entirely (after calling [`Application.Quit`](../ScriptReference/Application.Quit.md)), it’s not possible to reload Unity again in the same app session.
* You can’t load more than one instance of the Unity runtime, or integrate more than one Unity runtime.
* You might need to adapt your [native](plug-ins-native.md) and [managed](plug-ins-managed.md) **plug-ins** to work properly.