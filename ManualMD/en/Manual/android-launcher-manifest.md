# Unity Launcher Manifest

A Unity Launcher Manifest configures how the application looks and behaves before the application launches. For example, it contains the application’s icon, name, and install location. The Unity Launcher Manifest is a Unity-specific concept for Android development and you can overwrite it to integrate Unity as a component into an existing project. For more information, see [Integrating Unity into Android applications](UnityasaLibrary-Android.md).

## Settings

You can configure all the settings in the Unity Launcher Manifest from [Unity’s Player Settings](class-PlayerSettingsAndroid.md). This means, unless you want to integrate Unity as a component into an existing project, you don’t need to overwrite the Unity Launcher Manifest or manually edit any settings directly in the file.

A Unity Launcher Manifest file declares the following:

* The package’s name.
* The application’s icons.
* The application’s name.
* The application’s [starting activity](https://developer.android.com/guide/components/activities/intro-activities) and its intents.
* The application’s install location.
* The application’s supported screen sizes.