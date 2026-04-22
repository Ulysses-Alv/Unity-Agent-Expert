# Build assets into AssetBundles

To build assets into an AssetBundle, you must [assign assets to an AssetBundle](assetbundles-assign-assets.md), either in the Unity Editor or through a script. You can then create and use a script to build the AssetBundles. For information on the best approaches for organizing assets into AssetBundles, refer to [Preparing assets for AssetBundles](AssetBundles-Preparing.md).

**Note**: This workflow describes the creation of AssetBundles with the built-in [`BuildPipeline.BuildAssetBundles`](../ScriptReference/BuildPipeline.BuildAssetBundles.md) API. A more user-friendly alternative is to use the [Addressables](http://docs.unity3d.com/Packages/com.unity.addressables@latest/index.html) package.

## AssetBundle build script

To build AssetBundles, you must create a build script and place it a folder called `Editor` in the `Assets` folder.

The following script is an example of an AssetBundle build script. It adds a menu item at the bottom of the Assets menu called **Build AssetBundles**. When you select **Build AssetBundles** the `BuildAllAssetBundles` method is called. A progress bar appears while the build takes all the assets you labeled with an AssetBundle name and uses them to populate AssetBundles at the path that `assetBundleDirectory` defines.

```
using UnityEditor;
using System.IO;

public class CreateAssetBundles
{
    [MenuItem("Assets/Build AssetBundles")]
    static void BuildAllAssetBundles()
    {
        // Ensure the AssetBundles directory exists, and if it doesn't, create it.
        string assetBundleDirectory = "Assets/AssetBundles";
        if (!Directory.Exists(assetBundleDirectory))
            Directory.CreateDirectory(assetBundleDirectory);

        // Build all AssetBundles and place them in the specified directory.
        BuildPipeline.BuildAssetBundles(assetBundleDirectory,
                                        BuildAssetBundleOptions.None,
                                        BuildTarget.StandaloneWindows);
    }
}
```

The script has the following arguments:

* `assetBundleDirectory`: The directory to output the AssetBundles to within the current Unity project. The folder doesn’t need to be in the `Assets` folder. In the code example, it creates the folder on demand if it doesn’t exist.
* `BuildAssetBundleOptions.None`: The default value for the build options argument. You can use this argument to specify one or more flags to enable a variety of optional behaviours. For example, this argument controls the choice of [compression algorithm](assetbundles-compression-format.md). For a full list of options available refer to [the `BuildAssetBundleOptions` API documentation](../ScriptReference/BuildAssetBundleOptions.md).
* `BuildTarget.StandaloneWindows`: Defines the [target platform](../ScriptReference/BuildTarget.md) for the AssetBundles. Alternatively, you can call `EditorUserBuildSettings.activeBuildTarget`, which returns the platform profile currently set as active in the [Build Profiles](build-profiles.md) window.

## Build subsets of AssetBundles

When you give the [`BuildPipeline.BuildAssetBundles`](../ScriptReference/BuildPipeline.BuildAssetBundles.md) method no specific AssetBundle names, it builds all AssetBundles defined in the project. If you only want to build a subset of AssetBundles, you query the AssetDatabase for defined AssetBundles and then pass a filtered list to the build pipeline.

The following script demonstrates how to retrieve all AssetBundle names and their assigned assets, allowing you to filter or modify the list before building:

```
using UnityEditor;
using System.IO;
using UnityEngine;
using System.Collections.Generic;

public class BuildSubsetAssetBundles
{
    [MenuItem("Assets/Build Selected AssetBundles")]
    static void BuildSpecificAssetBundles()
    {
        string assetBundleDirectory = "Assets/AssetBundles";
        if (!Directory.Exists(assetBundleDirectory))
        {
            Directory.CreateDirectory(assetBundleDirectory);
        }

        List<AssetBundleBuild> builds = new List<AssetBundleBuild>();
        string[] allAssetBundleNames = AssetDatabase.GetAllAssetBundleNames();

        // Example: Only build AssetBundles that start with "environment"
        foreach (string bundleName in allAssetBundleNames)
        {
            if (bundleName.StartsWith("environment"))
            {
                AssetBundleBuild build = new AssetBundleBuild
                {
                    assetBundleName = bundleName,
                    assetNames = AssetDatabase.GetAssetPathsFromAssetBundle(bundleName)
                };
                builds.Add(build);
            }
        }

        if (builds.Count > 0)
        {
            BuildPipeline.BuildAssetBundles(assetBundleDirectory,
                                            builds.ToArray(),
                                            BuildAssetBundleOptions.None,
                                            BuildTarget.StandaloneWindows);
            Debug.Log($"Built {builds.Count} specific AssetBundles.");
        }
        else
        {
            Debug.Log("No AssetBundles matching criteria found to build.");
        }
    }

    [MenuItem("Assets/Log All AssetBundle Assignments")]
    static void LogAllAssetBundleAssignments()
    {
        string[] allAssetBundleNames = AssetDatabase.GetAllAssetBundleNames();
        Debug.Log($"Total AssetBundles Defined: {allAssetBundleNames.Length}");
        foreach (string bundleName in allAssetBundleNames)
        {
            string[] assetPaths = AssetDatabase.GetAssetPathsFromAssetBundle(bundleName);
            Debug.Log($"AssetBundle: {bundleName} (Assets: {assetPaths.Length})");
            foreach (string path in assetPaths)
            {
                Debug.Log($"  - {path}");
            }
        }
    }
}
```

This approach allows your script to respect existing **Inspector** assignments while giving you granular control over which AssetBundles are built.

## Perform a clean build

When you create an official AssetBundle release, perform a [clean build](build-clean-build.md) to ensure Unity rebuilds all content during the build process. To perform a clean build, pass the [`BuildAssetBundleOptions.ForceRebuildAssetBundle`](../ScriptReference/BuildAssetBundleOptions.ForceRebuildAssetBundle.md) flag as an option to `BuildPipeline.BuildAssetBundles`.

In some projects, you can delete the `Library/ShaderCache` directory to force a full recompilation of shaders, or to reclaim disk space from obsolete **shader** data. However, deleting the `ShaderCache` folder increases the time it takes for Unity to create another build.

For more information on clean builds, refer to [Create a clean build](build-clean-build.md).

## Changing target platform

The [`BuildPipeline.BuildAssetBundles`](../ScriptReference/BuildPipeline.BuildAssetBundles.md) API allows you to specify the target and sub target platform for deploying AssetBundles.

If the specified target platform differs from the one configured in [Build Profiles](build-profiles.md), Unity must recompile Editor **scripts** and reimport assets such as textures that have platform-specific representations. After the build is complete, Unity restores the original target platform settings.

This process can significantly increase build times. Additionally, the script containing the `BuildPipeline.BuildAssetBundles` call continues to execute as compiled for the current target platform, not the specified build target. This can cause issues if the build script or callback scripts rely on platform-specific code or assemblies.

To avoid this issue, ensure that any code executed during a build dynamically checks the target platform (for example, using `if` statements) rather than relying on platform-specific conditional compilation (such as, `#ifdef` statements). It’s best practice to always set the target to the desired target, then launch the script that builds AssetBundles to avoid any issues from code outside your control such as build callbacks inside packages.

For command line builds, use the `--buildTarget` or `-activeBuildProfile` [command line argument](EditorCommandLineArguments.md) to align the target platform with your build requirements. For more information, refer to [Create a build from the command line](build-command-line.md).

## Rebuilding assets incrementally

Each AssetBundle has a hash that Unity uses to determine whether it needs to be rebuilt. Unity determines how to rebuild AssetBundles [incrementally](building-introduction.md#incremental-build-pipeline) as follows:

1. If a `.manifest` file from a previous build of that AssetBundle exists, Unity compares the IncrementalBuildHash hashes from both builds.
2. If the AssetBundle hashes match, then Unity calculates and compares their type tree hashes. Unity uses `TypeTreeHash` as a secondary hash to detect whether any objects used in the AssetBundle have newer serialization formats. You can ignore this check by specifying the [`BuildAssetBundleOptions.IgnoreTypeTreeChanges`](../ScriptReference/BuildAssetBundleOptions.IgnoreTypeTreeChanges.md) flag.
3. If the AssetBundle hash and the type tree hash don’t match, then Unity rebuilds the AssetBundle unless you’ve specified `BuildAssetBundleOptions.ForceRebuildAssetBundle,` which rebuilds the AssetBundle every time.
4. Unity serializes the newly calculated hash values into the new AssetBundle’s `.manifest` file.

The AssetBundle IncrementalBuildHash considers inputs such as the target platform, included assets, dependencies, build options, and platform-specific settings like **mesh** stripping and lighting configurations. However, this hash doesn’t consider every possible build influence, which might lead to crashes or unexpected failures if the incremental build systems doesn’t detect all changes. Use incremental builds for internal development but perform a [clean build](build-clean-build.md) when creating a release build.

**Warning**: The `TypeTreeHash` is distinct from the main `AssetBundle` input hash (IncrementalBuildHash). A change in the `TypeTreeHash` can force an incremental build without changing the `AssetBundle` input hash value, so the `AssetBundle` input hash isn’t an ideal value for tracking file versions. It’s more reliable to use a hash based on the content or other version numbering scheme. For more information, refer to [Cache version hash](assetbundles-caching.md#cache-version-hash).

## Additional resources

* [Loading AssetBundles](AssetBundles-Native.md)
* [AssetBundle compression formats](assetbundles-compression-format.md)
* [Preparing assets for AssetBundles](AssetBundles-Preparing.md)