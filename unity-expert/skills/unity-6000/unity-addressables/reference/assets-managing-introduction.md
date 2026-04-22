# Introduction to runtime asset management

How you load and manage assets in your project impacts on your application’s memory usage, load times, and build size. Unity provides several asset management systems, each designed for specific use cases and project sizes. Small applications with fixed content sizes can usually use Unity’s [default asset management system](#direct-references). Systems like [AssetBundles](#assetbundle-system) and [Addressables](#addressables-package) provide better management for complex asset layouts, or applications that need to stream assets from remote servers.

## Effect of runtime asset management on builds

By default, Unity uses the [Scene List](build-profile-scene-list.md) to determine which scenes and assets to include in a build and their load order. Unity only includes scenes in the **Scene** List in a build, and unused scenes and assets not referenced from these scenes are excluded from a build. You can also use the [`SceneManager` API](../ScriptReference/SceneManagement.SceneManager.md) to determine how to load scenes in your project.

The exception to this is if you include any scenes in a [Resources folder](#resources-system), or mark them as [Addressable](#addressables-package), then Unity handles scenes differently.

To determine what content to build, Unity uses the following information:

* Scenes included in the [Scene List](build-profile-scene-list.md).
* Any [Preferences](Preferences.md) that affect building.
* Any assets [directly referenced](#direct-references).
* If you use the [Resources system](#resources-system), Unity includes any assets in the `Resources` folder.
* If you use [AssetBundles](#assetbundle-system), Unity includes any assets and scenes assigned to an AssetBundle, plus any direct references from those assets or scenes.
* If you use [Addressables](#addressables-package), Unity includes any assets and scenes assigned to an Addressable group, plus any direct references from those assets or scenes.

## Runtime asset management options

Choosing the right approach depends on factors like your target platform, content delivery requirements, and whether you need dynamic loading capabilities. The following management systems are available:

* [Direct references](#direct-references): The default way of managing assets. Unity loads all assets referenced in a scene into memory before loading.
* [The Resources system](#resources-system): Loads and unload assets in the `Resources` folders in your project with the `Resources` API.
* [The AssetBundle system](#assetbundle-system): Organize assets into AssetBundles with the `AssetBundle` API. You can then store AssetBundles remotely and download on demand.
* [The Addressables package](#addressables-package): Built on top of the `AssetBundle` API, it automates a lot of the manual processes that the AssetBundle system has, and provides a Unity Editor interface.
* [ECS content management system](#ecs-content-management-system): If you use the [Entities package](https://docs.unity3d.com/Packages/com.unity.entities@latest) in your project, then it has its own content management system.

### Direct references

Direct reference asset management is the default way of managing assets at runtime in Unity. Whenever you drag an asset into a scene or onto a component through the **Inspector** window of the Unity Editor, you create a direct reference to that asset. When you build your application, Unity saves any referenced assets in a separate file associated with the scenes in your project.

You can use direct references with `ScriptableObject` instances to efficiently manage large asset sets and reduce build sizes. A benefit of only using direct references is that Unity only includes assets that your project directly references in the build, so your application doesn’t load unnecessary or duplicated assets.

Direct references have some limitations, as follows:

* When your application runs on the target device, Unity loads the entire asset file into memory before loading a scene. This approach isn’t dynamic and you can only load or unload a scene, which might lead to slow load times, especially on devices with less memory.
* At the time of loading, the asset must already be present in the local file system, and can’t be hosted on a content delivery network (CDN).
* You can’t make incremental content updates and need to rebuild and redeploy your entire application if you want to make changes.

For more information, refer to [Direct reference asset management](assets-direct-reference.md).

### Resources system

The [Resources system](LoadingResourcesatRuntime.md) provides a simple way of managing assets in memory. To use the Resources system, you create a folder called `Resources` in your project and add assets to it. During the build process, Unity finds the assets in the `Resources` folder and bundles them into a serialized file with metadata and indexing information. You can then use the [`Resources`](../ScriptReference/Resources.md) API to load and unload the assets in your application.

The Resources system has the following limitations:

* Assets in `Resources` folders are always included in the Player build, even if they’re not referenced by anything. This can lead to large build sizes.
* If you need to make changes to the assets in your project, you need to rebuild and redeploy your entire application, whereas AssetBundles and Addressables allow you to create incremental content updates.
* It doesn’t support content delivery from a remote server.
* Assets loaded with `Resources.Load` remain in memory until explicitly unloaded with `Resources.UnloadUnusedAssets`, which unloads all unused assets loaded by `Resources`, or until the scene that used them is unloaded.
* If you have a lot of assets in the `Resources` folder, then building and starting your application can take a long time.
* If your code loads an asset by string, but that asset is missing, then `Resources.Load` returns `null` and becomes a runtime error, so you need to manually track which assets are needed.

It’s an ideal system for prototyping, or for smaller projects which need assets throughout the lifetime of the application. For more information, refer to [Use the Resources system to load assets at runtime](assets-resources-system.md).

### AssetBundle system

Use the [AssetBundle system](AssetBundlesIntro.md) to organize assets into separate container files called AssetBundles. You can store AssetBundle files in your project files, or remotely in the cloud.

The [`AssetBundle` API](../ScriptReference/AssetBundle.md) minimizes the impact on network and system resources by allowing you to download AssetBundles on demand. You can use this on-demand behavior to add DLC and post-release content updates to your application. For example, you can deliver new content for your game without requiring users to download a new version of your game. The `AssetBundle` API then provides a way to load and unload assets from downloaded bundles.

The AssetBundle system has the following limitations:

* You can only use AssetBundles in **scripts**, and there’s no Editor interface to build AssetBundles with.
* The `AssetBundle` API doesn’t keep track of asset dependencies. For example, if you want to load a **prefab** from an AssetBundle, you need to manually load the AssetBundles, and any AssetBundle that the AssetBundle depends on, before loading the prefab. Safely unloading AssetBundles often requires writing a reference counting system. For more information, refer to [Handling dependencies between AssetBundles](AssetBundles-Dependencies.md).
* You need to manually allocate and deallocate memory, so you might create memory leaks or missing content by unloading assets from an AssetBundle that other code relies on.
* The `AssetBundle` API isn’t aware of whether an AssetBundle is hosted locally or remotely, so you need to keep track of the location of all the AssetBundles in your project.

For more information, refer to [Introduction to AssetBundles](AssetBundlesIntro.md).

### Addressables package

The [Addressables package](https://docs.unity3d.com/Packages/com.unity.addressables@latest) provides a user interface in the Unity Editor to manage and organize the assets in your project. It’s built on top of the `AssetBundle` API and automatically manages dependencies, asset locations, and memory allocation, which you otherwise have to handle manually in the AssetBundle system. The Addressables package aims to remove the limitations of the Resources and AssetBundle system to make it easier to manage assets on demand.

Once you have made an asset Addressable, you can reference it by address in your code, rather than by its file name or AssetBundle location, which means you can change the location of an asset without needing to rewrite any code. You can also use asset references to dynamically load and unload assets, so that you don’t have to keep all assets in the scene in memory at all times.

The Addressables system is designed for applications that have assets stored on a CDN, and those that are likely to have content updates over time. If your application has fixed content which is shipped in the initial download, then using direct references, or the Resources system is a better alternative to Addressables or AssetBundles.

For more information, refer to [Addressables overview](https://docs.unity3d.com/Packages/com.unity.addressables@latest/index.html?subfolder=/manual/AddressableAssetsOverview.html)

### ECS content management system

The Entities package has its own [content management system](https://docs.unity3d.com/Packages/com.unity.entities@latest/index.html?subfolder=/manual/content-management.html), where you use weak references to assets to load them at runtime. You can also create content archives to deliver content to an application. For more information, refer to [Introduction to content management](https://docs.unity3d.com/Packages/com.unity.entities@latest/index.html?subfolder=/manual/content-management-intro.html).

## Additional resources

* [Direct reference asset management](assets-direct-reference.md)
* [Use the Resources system to load assets at runtime](assets-resources-system.md)
* [Use AssetBundles to load assets at runtime](assetbundles-section.md)
* [Addressables package](https://docs.unity3d.com/Packages/com.unity.addressables@latest)
* [Entities package](https://docs.unity3d.com/Packages/com.unity.entities@latest)