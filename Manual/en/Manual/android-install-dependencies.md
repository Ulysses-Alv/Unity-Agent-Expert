# Install dependencies

Install required dependencies to develop your Android application in Unity.

To set up your Unity project for Android development, you need to install the required dependencies which are Android Build Support, SDK, NDK, and JDK. Unity distributes these dependencies as [modules](https://docs.unity3d.com/hub/manual/AddModules.html) which means you must use the [Unity Hub](https://docs.unity3d.com/hub/manual/index.html) to install them.

## Install modules

You can install the modules in two ways:

* During installation of a new Unity Editor version. For more information, refer to [Download and install the Unity Editor](https://docs.unity3d.com/hub/manual/AddEditor.html).
* By adding them to an existing Unity Editor installation. For more information, refer to [Add modules](https://docs.unity3d.com/hub/manual/AddModules.html).

The three modules to install are:

* **Android Build Support**
* **Android SDK & NDK Tools**
* **OpenJDK**

![Unity Hub displaying the three dependency modules.](../uploads/Main/android-modules.png)

Unity Hub displaying the three dependency modules.

Unity installs **Android SDK & NDK Tools** and **OpenJDK** respectively in the `SDK`, `NDK`, and `OpenJDK` folders under `/[EditorVersion]/Editor/Data/PlaybackEngines/AndroidPlayer/`.

## Additional resources

* [Customize dependencies](android-customize-dependencies.md)
* [Supported dependency versions](android-supported-dependency-versions.md)