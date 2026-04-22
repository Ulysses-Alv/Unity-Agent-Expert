# Introduction to the Resources system

The [`Resources`](../ScriptReference/Resources.md) class allows you to find and access objects that are in `Resources` folders in your project.

You can use the Resources system to make an asset available to a project without loading it in as part of a **scene**. When you add content to a folder named `Resources` in your project, Unity makes it available to load when needed, independent of the scenes that you build.

However, having the content constantly available at runtime has a [significant performance impact on your project](#impact-of-large-asset-sets-on-the-resources-system), especially if it contains a large amount of assets. The [AssetBundles](AssetBundlesIntro.md) system, or the [Addressables](https://docs.unity3d.com/Packages/com.unity.addressables@latest) package provide a better way of managing complex sets of assets. For more information on the different asset management options, refer to [Introduction to asset management](assets-managing-introduction.md).

## Using the Resources system

The `Resources` system can be helpful in the following situations:

* In smaller projects if the content meets the following criteria:
  + It’s needed throughout the project’s lifetime.
  + It’s not memory-intensive.
  + It doesn’t need patching, or doesn’t vary across platforms and devices.
  + It’s used for minimal bootstrapping.

Examples of the latter include MonoBehaviour singletons used to host **prefabs**, or ScriptableObject instances containing third-party configuration data, such as a Facebook App ID.

## Impact of large asset sets on the Resources system

It’s best practice to avoid putting a lot of content into the Resources system for the following reasons:

* It makes memory management more difficult.
* Placing a lot of content in the `Resources` folder slows down application startup and the length of builds.
* The `Resources` system makes it harder to deliver custom content to specific platforms and prevents incremental content upgrades.
* Making changes to assets in the `Resources` folder requires a player rebuild and redeployment, whereas AssetBundles are better suited for incremental content updates.

If you need to manage content via a CDN, or do regular patch updates, then [AssetBundles](AssetBundlesIntro.md) and the Addressables package are the recommended alternative. You can also use [direct references](assets-direct-reference.md) if you don’t need the extra control that the Resources system provides.

Unity combines all objects referenced from assets in the `Resources` folder into a single serialized file when you build a project, similar to building one large [AssetBundle](AssetBundlesIntro.md). Unity takes time to construct the data structures that index all the objects in the file, and uses memory proportional to the number of objects, even if they aren’t loaded. This has an impact on performance if the number of individual assets in Resource folders are large, or if there are prefabs that contain a lot of objects. For example, initializing a Resources system containing 10 thousand assets takes several seconds on low-end mobile devices, even though most of the objects contained in Resources folders are rarely needed in an application’s first scene.

## Additional resources

* [Load and unload assets with the Resources system](assets-resources-system-load.md)
* [Addressables](https://docs.unity3d.com/Packages/com.unity.addressables@latest)
* [Optimizing assets](assets-optimizing.md)
* [AssetBundles](AssetBundlesIntro.md)