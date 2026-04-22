# Introduction to building

When you create a build of your application, you create a Player. A Player is the platform-specific runtime application that Unity builds from your project. This is also known as a **project build**, which is the workflow of building a project from the Unity Editor into an application that runs on a target platform.

Building a Player for a target platform requires the platform-specific build support module for the target platform. You can add build support for a target platform when you [install Unity](GettingStartedInstallingUnity.md), or add it when you [create a Build Profile](create-build-profile.md).

Unity uses the **scenes** you define in the [Build Profiles](build-profiles.md) window or the [`BuildPipeline`](../ScriptReference/BuildPipeline.md) API to create a build of a Player. For more information, refer to [Manage scenes in a build](build-profile-scene-list.md).

## Build modes

Unity has different build modes, as follows:

* **Release** build: Includes only what’s necessary to run the application. This is the default build type.
* **Development** build: Includes scripting debug symbols and the [Profiler](Profiler.md). You can enable **development builds** in the [Build Profiles](build-profiles.md) window, which allows you to set further options such as [deep profiling support](profiler-deep-profiling.md) and script debugging. You can also use the [`BuildOptions.Development`](../ScriptReference/BuildOptions.Development.md) property to set a development build.

Both build modes provide options to build different variations of the Player application for different hardware architectures and [scripting backends](scripting-backends.md). You can customize these variations through the [build settings](build-profiles-reference.md), [Player settings](class-PlayerSettings.md), or [command-line flags](EditorCommandLineArguments.md).

## Incremental build pipeline

Unity uses an incremental build pipeline that only rebuilds the parts of your application that have changed since the last build, which helps speed up development iteration time. This build process includes build steps such as content building, code compilation, data **compression**, and signing.

By default, Unity uses the incremental build pipeline for both [development and release builds](building-introduction.md). You can use the options in the **Build Profiles** window, or use the `BuildOptions.CleanBuildCache` API to create a non-incremental build, also known as a clean build. For more information, refer to [Creating clean builds](build-clean-build.md).

**Note:** [AssetBundles](AssetBundlesIntro.md) don’t use the incremental build pipeline and have separate mechanisms for caching and reusing the results from previous builds. For more information, refer to [Build assets into an AssetBundle](AssetBundles-Building.md).

## Additional resources

* [Build Profiles](build-profiles.md)
* [Create a Build Profile](create-build-profile.md)
* [Build Profiles window reference](build-profiles-reference.md)
* [Player settings](class-PlayerSettings.md)
* [Create a clean build](build-clean-build.md)