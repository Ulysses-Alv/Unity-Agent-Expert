# Dedicated Server AssetBundles

You can apply the Dedicated Server optimizations to AssetBundles starting in Unity Editor version **2023.1.0a21**. You can build an AssetBundle through scripting. Refer to the section on [AssetBundle](AssetBundlesIntro.md) for more information on building AssetBundles in general.

To build an AssetBundle to undergo the same Dedicated Server optimizations as discussed for the Player, specify the `subtarget` field in the `BuildAssetBundlesParameters` struct to be `StandaloneBuildSubtarget.Server` when calling the `BuildAssetBundle` method. Refer to the following example:

```
BuildAssetBundlesParameters serverAssetBundleParameters =
{
    outputPath = /*some example asset path here, not entirely relevant*/,
    options = BuildAssetBundleOptions.None,
    targetPlatform = BuildTarget.StandaloneWindows64,  //alternately, the MacOS or Linux build targets, any desktop platform
    subtarget = StandaloneBuildSubtarget.Server
};
BuildPipeline.BuildAssetBundles(serverAssetBundleParameters);
```

After you build the AssetBundle, you can load it by a Player at runtime. Refer to [Using AssetBundles Natively](AssetBundles-Native.md).

**Warning**: While the AssetBundle loading process checks if the AssetBundle target platform matches the target platform of the player, it doesn’t check the AssetBundle subtarget. Make sure to not load an AssetBundle that was built for a non-server standalone player. Don’t try to load an AssetBundle that targets the Dedicated Server subtarget (or vice-versa).