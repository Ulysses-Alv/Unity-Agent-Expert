# Extend the default Unity activity

The `UnityPlayerActivity` of a Unity Android application is responsible for basic interactions between the Android operating system and the application. You can use **plug-ins** to create your own [Activity](https://developer.android.com/reference/android/app/Activity.html) that extends and overrides Unity’s default `UnityPlayerActivity`.

**Note**: If you’re creating a custom activity with GameActivity application entry point, you need to extend the `UnityPlayerGameActivity` class.

| **Topic** | **Description** |
| --- | --- |
| [Create a custom activity](android-custom-activity.md) | Extend the default Unity activity to control interactions between Unity and Android. |
| [Specify Android Player command-line arguments](android-custom-activity-command-line.md) | Specify startup command-line arguments to pass to Unity. |

## Additional resources

* [Android plug-ins](PluginsForAndroid.md)