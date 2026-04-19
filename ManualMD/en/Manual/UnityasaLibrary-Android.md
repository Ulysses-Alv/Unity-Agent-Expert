# Integrating Unity into Android applications

This page describes how to integrate the Unity Runtime Library into Android applications using the Unity as a Library feature.

You can use this feature to include Unity-powered features, such as 3D/2D Real-Time Rendering, **AR** Experience, 3D model interaction, or 2D mini-games, into your application. The Unity Runtime Library exposes controls to manage when and how to load, activate, and unload content within the application.

**Important**: The introduction of Unity as a Library in your project might require you to adapt [native](plug-ins-native.md) and [managed](plug-ins-managed.md) **plug-ins** to work properly for Android. Plug-ins that make changes to **Gradle** manifests need to use the Gradle changes outlined in [Using Unity as a library in native iOS/Android apps](https://discussions.unity.com/t/using-unity-as-a-library-in-native-ios-android-apps/744882).

## How it works

You don’t need to do anything different when you build your Gradle project from Unity.

Every [Android Gradle project](android-gradle-overview.md) that Unity generates has the following structure:

* A library part in the **unityLibrary** module that you can integrate into any other Gradle project. This contains the Unity runtime and Player data.
* A thin launcher part in the **launcher** module that contains the application name and its icons. This is a simple Android application that launches Unity. You can replace this module with your own application.

To integrate Unity into another Android Gradle project, you must include the **unityLibrary** module of the generated Android Gradle project in your Android Unity Project through the *settings.gradle* file.

This [repository](https://github.com/Unity-Technologies/uaal-example/blob/master/docs/android.md) contains example Projects and plug-ins that demonstrate how to integrate Unity into an Android app, along with further documentation.

To control a Player, relay an Intent to launch Unity activity and extend it if needed. For more information, see Android developer documentation on [Intents and Intent Filters](https://developer.android.com/guide/components/intents-filters). You can also use the UnityPlayer Java API.

## IUnityPlayerLifecycleEvents

IUnityPlayerLifecycleEvents provides a way to interact with two important lifecycle events of the Unity Player:

* **Unload** - The application calls `IUnityPlayerLifecycleEvents.onUnityPlayerUnloaded` when `Application.Unload` or `UnityPlayer.unload()` unloads the Unity Player. This puts the Unity Player in a paused state where it unloads all **Scenes**, but keeps everything else loaded in the memory.
* **Quit** - The application calls `IUnityPlayerLifecycleEvents.onUnityPlayerQuitted` when the Unity Player quits. The process that was running Unity ends after this call.

You can pass an instance of `IUnityPlayerLifecycleEvents` to the UnityPlayer constructor, or to override methods in subclasses of `UnityPlayer` and `UnityPlayerActivity`.

## Limitations

Unity doesn’t control the runtime lifecycle, so Unity as a Library might not work for all possible use cases. Known limitations include:

* Unity as a Library only supports full-screen rendering. However, if you are a Unity Industry customer, the limitations and features might differ.
* You can’t load or integrate more than one instance of the Unity runtime.
* You might need to adapt third-party plug-ins (both [native](plug-ins-native.md) and [managed](plug-ins-managed.md)) to work with the Unity runtime.
* Unity as a Library isn’t compatible with the [Xamarin app platform](https://dotnet.microsoft.com/en-us/apps/xamarin).
* You can’t integrate Unity Runtime Library as a dynamic module with [Play Feature Delivery](https://developer.android.com/guide/playcore/feature-delivery).

## Additional resources

* [Using Unity as a Library in other applications](UnityasaLibrary.md)