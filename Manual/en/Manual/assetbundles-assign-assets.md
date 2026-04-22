# Assign assets to an AssetBundle

To build assets into an AssetBundle, you must first assign assets to an AssetBundle, either in the Unity Editor or through a script. You can then [create and use a script](AssetBundles-Building.md) to build the AssetBundles. For information on the best approaches for organizing assets into AssetBundles, refer to [Organizing assets into AssetBundles](AssetBundles-Preparing.md).

**Note**: This workflow describes the creation of AssetBundles with the built-in [`BuildPipeline.BuildAssetBundles`](../ScriptReference/BuildPipeline.BuildAssetBundles.md) API. An alternative is to use the [Addressables](http://docs.unity3d.com/Packages/com.unity.addressables@latest/index.html) package, which is built on top of AssetBundles and provides a UI to organize the assets in your project.

## Assign assets to AssetBundles in the Editor

To assign a given asset to an AssetBundle in the Unity Editor, perform the following steps:

1. Select the asset in the **Project window** and view it in the [Inspector](UsingTheInspector.md).
2. Use the **AssetBundle** left-hand dropdown menu at the bottom of the Inspector window to assign or create an AssetBundle:
   * To create a new AssetBundle, select the left-hand dropdown and select **New**, or choose an existing AssetBundle from the list. **Tip:** To organize AssetBundles with subfolders, use the `/` character. For example, use the AssetBundle name `environment/forest` to create an AssetBundle named `forest` under an `environment` subfolder.
3. Optionally use the right-hand menu to assign or create an [AssetBundle variant](AssetBundles-Preparing.md#assetbundle-variants):
   * To create a new variant, select the right-hand dropdown and select **New**, or choose an existing variant from the list.

### Assigning multiple assets to an AssetBundle

You can assign an AssetBundle to a folder in your project. By default, all assets in that folder are assigned to the same AssetBundle as the folder. However, the AssetBundle assignments for individual assets takes precedence. To assign an AssetBundle to a folder:

1. Select the folder in its containing parent folder the Project window.
2. Use the **AssetBundle** dropdown to assign a new or existing AssetBundle to the folder.

You can also select multiple assets and assign them to an AssetBundle. However, assigning assets to an AssetBundle in this way overrides any existing AssetBundle assignments for those assets.

## Assign assets to AssetBundles with a script

To assign assets to AssetBundles in your code, use the [`BuildPipeline.BuildAssetBundles`](../ScriptReference/BuildPipeline.BuildAssetBundles.md) method with an array of [`AssetBundleBuild`](../ScriptReference/AssetBundleBuild.md) structures. This approach overrides any AssetBundle assignments made in the Inspector.

If you want your script to respect AssetBundle assignments made in the Inspector, use [`AssetDatabase.GetAllAssetBundleNames`](../ScriptReference/AssetDatabase.GetAllAssetBundleNames.md) and [`AssetDatabase.GetAssetPathsFromAssetBundle`](../ScriptReference/AssetDatabase.GetAssetPathsFromAssetBundle.md) to retrieve the necessary information and populate the `AssetBundleBuild` array.

## Additional resources

* [Organizing assets into AssetBundles](AssetBundles-Preparing.md)
* [Build assets into AssetBundles](AssetBundles-Building.md)