# Create a clean build

By default, Unity creates builds incrementally, however the incremental pipeline can cause caching issues, or incomplete builds. For example, Unity [serializes](script-serialization.md) **scenes** and assets for the target platform during the build process, and either reuses all content from the previous build, or rebuilds everything. However, this detection might fail when global settings affect build output. You can create a clean build to resolve this issue.

When Unity creates a clean build, it discards some cached build data but reuses previously imported assets and cached **shaders** to rebuild all content and code. Use a clean build when preparing release builds or troubleshooting issues caused by corrupted or outdated build caches.

To create a clean build:

1. Open the **Build Profiles** window (**File** > **Build Profiles**).
2. Next to the **Build** button, select the drop-down.
3. Select **Clean Build**.

You can also create a clean build from a [build script](build-script-build.md) by passing [`BuildOptions.CleanBuildCache`](../ScriptReference/BuildOptions.CleanBuildCache.md) in the call to [`BuildPipeline.BuildPlayer`](../ScriptReference/BuildPipeline.BuildPlayer.md).

If you don’t want Unity to reuse imported assets and cached shaders, you can delete the `Library` folder to force Unity to repeat all imports and shader compilation. This approach increases build time but helps avoid issues that might occur at earlier stages of the build process.

To create a clean build without using cached shaders and assets:

1. Close the Editor.
2. Delete the `Library` folder
3. Use the previous instructions to create a clean build.

## Additional resources

* [Create a build from the Editor](BuildSettings.md)
* [Content output of a build](build-content-output.md)