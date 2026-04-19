# Introduction to customizing the build pipeline

A build pipeline is the automated processing and tooling that transforms your project from source assets and code into a Player for one or more target platforms. Unity includes a build pipeline for creating Player builds for many target platforms, and you can also create content-only builds, for example when using [Addressables](https://docs.unity3d.com/Packages/com.unity.addressables@latest).

You can automate and customize the behavior of the built-in build pipeline to meet the specific needs of your project and development workflow.

## Editor-based customizations

You can create Editor-based custom **scripts** and callbacks as follows:

* **[Custom build scripts](build-script-build.md)**: Use [`BuildPipeline.BuildPlayer`](../ScriptReference/BuildPipeline.BuildPlayer.md) to build a Player. These scripts can also include the following:
  + **Content-only build scripts**: Scripts that perform content-only builds, for example using [`BuildPipeline.BuildAssetBundles`](../ScriptReference/BuildPipeline.BuildAssetBundles.md).
  + **Combined build scripts**: Scripts that build one or more content-only builds and then build a Player build.
* **[Build callbacks](build-callbacks.md)**: Callbacks hook into a stage of a Player or content-only build to perform extra steps during a build.

These customizations can perform various tasks. For example, you can import external assets, validate a project’s configuration, or adjust **scene** content during builds. You can also analyze build results with the [`BuildReport`](../ScriptReference/Build.Reporting.BuildReport.md) API or upload builds to servers.

## External build pipeline customizations

You can use scripts or continuous integration (CI) tools that run on a build machine or cloud service (for example [Unity Build Automation](https://unity.com/solutions/ci-cd)) to customize the build pipeline. These tools perform one or more builds [from the command line](build-command-line.md).

External scripts can perform actions that don’t depend on the Unity API and can happen before or after the Unity Editor runs. These include pulling source control branches, synchronizing assets from content creation systems, processing build output with platform-specific tools, analyzing results with tools like [UnityDataTools](https://github.com/Unity-Technologies/UnityDataTools), and publishing builds with notifications.

## Build determinism

[Build determinism](build-deterministic-builds.md) is important if you want to be able to repeat a build process and get the same results. When designing your build pipeline customizations, make them work in a way that’s repeatable and always produces the same results when given the same inputs. For example, introducing timestamps or randomized data during a build callback breaks the ability to repeat the same build and get identical results. For more information, refer to [Introduction to deterministic builds](build-deterministic-builds-introduction.md).

## Additional resources

* [`BuildPipeline.BuildPlayer` API reference](../ScriptReference/BuildPipeline.BuildPlayer.md)
* [`BuildReport` API reference](../ScriptReference/Build.Reporting.BuildReport.md)
* [Build callbacks](build-callbacks.md)
* [Create a custom build script](build-script-build.md)
* [Build a player from the command line](build-command-line.md)
* [Content output of a build](build-content-output.md)