# Introduction to assets in Unity

An asset is any item that you use in your Unity project to create your application, such as textures, 3D models, or sound files. Assets can include:

* ****Visual elements****: 3D models, textures, or **sprites**.
* **Audio elements**: Sound effects or music.
* **Abstract items**: Color gradients, animation masks, arbitrary text, or numeric data.

## Importing assets

To use assets in Unity, you must [import them into your project](import-assets.md). You can [add assets to the `Assets` folder](ImportingAssets.md) of your project, or [use scripts](ScriptedImporters.md) to import assets automatically.

**Important:** Using cloud-based storage methods to store your project is an unsupported workflow. It can cause synchronization issues which corrupt your project. Use [version control](VersionControl.md) to manage your project.

Unity supports a wide range of asset formats. For more information, refer to [Supported asset type reference](assets-supported-types.md).

If you’re working on a complex project with a large team of people, you can use the [Unity Accelerator](UnityAccelerator.md) cache server to speed up asset management.

## How Unity manages assets

Unity uses the [Asset Database](AssetDatabase.md) to store the assets in your project and maintain consistency between the original source files and their imported versions used by your application at runtime. You can use the [Import Activity window](ImportActivityWindow.md) to inspect how Unity imports the assets in your project.

## Grouping assets together

You can use [AssetBundles](assetbundles-section.md) to group together assets in an archive file format, which you can then use to update assets remotely, or provide DLC content for your application.

You can also use [asset packages](AssetPackages.md) to package assets together to share between other Unity projects.

## Managing assets through scripts

You can perform many of the loading, importing, and unloading operations that Unity does with the [Asset Database APIs](AssetDatabaseCustomizingWorkflow.md).

An alternative method of managing loading assets is with the [Resources system](LoadingResourcesatRuntime.md), but it can impact on the performance of your application.

The [Addressables package](https://docs.unity3d.com/Packages/com.unity.addressables@latest) provides a streamlined workflow for managing asset loading at runtime and is the recommended system for organizing assets in Unity projects.

## Additional resources

* [Supported asset type reference](assets-supported-types.md)
* [Introduction to Unity Accelerator](UnityAccelerator.md)
* [Programming with the Asset Database](AssetDatabaseCustomizingWorkflow.md)
* [Introduction to AssetBundles](AssetBundlesIntro.md)