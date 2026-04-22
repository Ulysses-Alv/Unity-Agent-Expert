# Application patching

Building an application for Android can take a significant amount of time. For faster iterations during development, you can patch the application package instead of rebuilding it. When you patch an application package, Unity only processes files you made changes to since the last patch and sends them to the connected Android device.

## Patching an application

To patch an application, you can use the Unity Editor or, if you implement your own build pipeline, the scripting API.

The first time you patch an application, Unity sets up the patch application environment on the device. This means that the first patch takes longer than future patches.

### From the Editor

To patch an application from the Unity Editor:

1. Open the **Build Profiles** window (menu: **File** > **Build Profiles**).
2. From the list of platforms in the **Platforms** panel, select **Android** or [create a build profile](create-build-profile.md) for the **Android** platform.
3. In the **Platform Settings** section, enable ****Development Build****.
4. In the **Build to Device** setting, select **Patch** or, if you want to run the patch after Unity finishes building it, **Patch And Run**.

### Using the scripting API

If you implement your own build pipeline, you can use the scripting API to patch your application. To do this, pass the [BuildOptions.BuildScriptsOnly](../ScriptReference/BuildOptions.BuildScriptsOnly.md) and [BuildOptions.PatchPackage](../ScriptReference/BuildOptions.PatchPackage.md) options to the [BuildPipeline.BuildPlayer](../ScriptReference/BuildPipeline.BuildPlayer.md) method.

For example:

```
BuildPlayerOptions buildPlayerOptions = new BuildPlayerOptions();
buildPlayerOptions.scenes = new[] { "Assets/Scene1.unity"};
buildPlayerOptions.target = BuildTarget.Android;
// Use Patch & Run for all builds for optimal turn-around times.
// (In prior versions, an initial regular Build & Run was necessary.)
buildPlayerOptions.options = BuildOptions.PatchPackage | BuildOptions.AutoRunPlayer | BuildOptions.Development;
BuildPipeline.BuildPlayer(buildPlayerOptions);
```

## How application patching works

When you patch an application, Unity:

* Creates a minimal **APK** that contains all compiled Java sources.
* Takes files that have changed since the last patch and stores them in a directory inside the application’s cache folder at: `/storage/emulated/0/Android/data/<PackageName>/pram-shadow-files/`.
* Stores native library files, including scripts compiled using [Il2cpp](./scripting-backends-il2cpp.md), in the application’s data directory at: `/data/data/<PackageName>/pram-shadow-files/lib/`.

When the application starts, it checks for a marker file inside the APK. If the marker exists, the application redirects its file reads to the `pram-shadow-files` directory.

A **Build & Run** doesn’t clear any of the patch files. Instead, it replaces the APK install. This disables the redirect to the `pram-shadow-files` directory. This means that subsequent patch builds can re-use unchanged files.

## Clearing patch files

You can use the Android storage settings to clear the application’s cache. Clearing the cache removes the patch files installed by the application patching build process.

The steps to clear an application’s cache are different depending on the Android device. To find the steps for your device:

1. Go to [Get help from your device manufacturer](https://support.google.com/android/answer/3094742).
2. Find your device manufacturer and follow the link to its support website.
3. Search the support website for steps on clearing an application’s cache.