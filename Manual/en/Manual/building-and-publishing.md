# Building and publishing

When you create a build of your application, you create a Player. A Player is the platform-specific runtime application that Unity builds from your project. You can control how Unity creates a build through Unity’s build and player settings.

| **Topic** | **Description** |
| --- | --- |
| **[Introduction to building](building-introduction.md)** | Understand the fundamentals of Unity’s build process, including build modes, incremental building, and platform compatibility. |
| **[Content output of a build](build-content-output.md)** | Information about the files that Unity creates when you make a build of your project. |
| **[Create a build from the Editor](BuildSettings.md)** | Use **build profiles** to build your application for different platforms with unique build configurations. |
| **[Create a clean build](build-clean-build.md)** | Rebuild all content from scratch without using cached results to resolve build cache issues. |
| **[Create a scripts-only build](build-scripts-only.md)** | Build only the scripting assemblies without processing **scenes** and assets for faster iteration during development. |
| **[Customize the build pipeline](build-customize-build-pipeline.md)** | Create custom **scripts** and use callbacks to customize the build pipeline and run it from the command line. |
| **[Include additional files in a build](StreamingAssets.md)** | Use the StreamingAssets folder to include additional files in a build. |
| **[Reducing the file size of a build](ReducingFilesize.md)** | Tips to reduce the size of the build. |
| **[Deterministic builds](build-deterministic-builds.md)** | Create consistent, reproducible builds by controlling factors that can cause build variations. |
| **[Build cache location reference](build-cache-location-reference.md)** | Reference for where Unity stores build cache files on different operating systems. |

## Additional resources

* [`BuildPipeline` API reference](../ScriptReference/BuildPipeline.md)
* [`BuildReport` API reference](../ScriptReference/Build.Reporting.BuildReport.md)
* [Player settings](class-PlayerSettings.md)
* [Platform-specific Player settings](PlatformSpecific.md)
* [Command line arguments](CommandLineArguments.md)
* [Managing assets at runtime](assets-managing-runtime.md)
* [Unity Build Automation](https://unity.com/solutions/ci-cd)