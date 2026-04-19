# GameActivity requirements and compatibility

Refer to the following system requirements and compatibility information before using the GameActivity application entry point in your Unity Android project.

## Unity and GameActivity library version compatibility

The following table lists the recommended GameActivity library version for each Unity version.

| **Unity version** | **GameActivity library version** |
| --- | --- |
| 6000.3 | 3.0.5 |
| 6000.2 | 3.0.5 |
| 6000.1 | 3.0.5 |
| 6000.0.22f1 and later | 3.0.5 |
| 6000.0.0b16 - 6000.0.21f1 | 3.0.2 |
| 6000.0.0b14 - 6000.0.0b15 | 3.0.0 |
| 6000.0.0a1 - 6000.0.0b13 | 2.0.0 |

## Dependencies

GameActivity requires the following dependencies:

* CMake build system
* AndroidX

### CMake

GameActivity uses [CMake](https://developer.android.com/ndk/guides/cmake) to produce the bridge code (`libgame.so`) during the build process.

**Note**: If you provide a custom Android SDK, be sure the SDK has CMake 3.22.1 included.

### AndroidX

GameActivity requires the following AndroidX **Gradle** dependencies:

* `androidx.appcompat:appcompat`
* `androidx.games:games-activity`
* `androidx.core:core`
* `Androidx.constraintlayout`

Gradle installs AndroidX and these dependencies automatically.

## Plug-in compatibility

If you use GameActivity, your application player loop runs on a native thread rather than a Java thread. This means that calling Java APIs like [myLooper](https://developer.android.com/reference/android/os/Looper#myLooper()) from [plug-ins](PluginsForAndroid.md) will fail. In the case of `myLooper` it’s because there’s no Java looper present on the native thread. This also means that any API that uses APIs such as `myLooper` will also fail. For example, [registerInputDeviceListener](https://developer.android.com/reference/android/hardware/input/InputManager#registerInputDeviceListener(android.hardware.input.InputManager.InputDeviceListener,%20android.os.Handler)) will fail if the handler is null. It’s important to understand this limitation when you create [Android plug-ins](PluginsForAndroid.md).

## Choreographer

If you use GameActivity, Unity tries to use the [NDK choreographer](https://developer.android.com/ndk/reference/group/choreographer) to synchronize frame times. If the [Device API Level](../ScriptReference/AndroidSdkVersions.md) is lower than 24, or your application uses a 32-bit Player and the Device API Level is lower than 29, Unity uses the [Java choreographer](https://developer.android.com/reference/android/view/Choreographer).

## Additional resources

* [The Activity application entry point](android-application-entries-activity.md)