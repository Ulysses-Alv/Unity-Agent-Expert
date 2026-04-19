# Supported dependency versions

Learn which versions of the Android SDK, NDK, and JDK each version of Unity supports.

Refer to the following tables to understand the Android dependency versions supported by Unity. Each version of Unity requires a specific version of the Android NDK and Android JDK, but there are no exact version requirements for the Android SDK.

## SDK

Unity relies on tools that the Android SDK provides and different versions of the SDK usually have the same tools available. This means you can use any recent version of the SDK as they all contain the build tools that Unity requires.

The following table shows the supported versions of Android SDK tools installed with each Unity version:

| **Unity version** | **SDK Build tools version** | **SDK Command-line tools version** | **SDK Platform tools version** |
| --- | --- | --- | --- |
| 6000.2 - 6000.3+ | 36.0.0 | 16 | 36.0.0 |
| 6000.1 | 34.0.0 | 16 | 34.0.5 |
| 6000.0 | 36.0.0 | 16 | 36.0.0 |
| 2022.3 LTS | 34.0.0 | 6 | 32.0.0 |
| 2021.3 LTS | 34.0.0 | 6 | 32.0.0 |

## NDK

The following table shows the NDK version that each Unity version supports:

| **Unity version** | **NDK version** |
| --- | --- |
| 6000.0+ | r27c (27.2.12479018) |
| 2022.3 LTS | r23b (23.1.7779620) |
| 2021.3 LTS | r21d (21.3.6528147) |

## JDK

The following table shows the JDK version that each Unity version supports:

| **Unity version** | **JDK version** |
| --- | --- |
| 6000.0+ | 17 (OpenJDK version 17) |
| 2022.3 LTS | 11 (OpenJDK version 11) |
| 2021.3 LTS | 11 (OpenJDK version 11) |

**Note**: Unity versions 6 and later support CMake version 3.22.1, while earlier Unity versions don’t support CMake.

## Additional resources

* [Install Android dependencies](android-install-dependencies.md)
* [Customize dependencies](android-customize-dependencies.md)
* [CMake](https://developer.android.com/ndk/guides/cmake)