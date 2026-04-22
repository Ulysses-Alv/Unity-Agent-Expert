# Organizing assets into AssetBundles

When you create an AssetBundle, there are some [limitations](#assetbundle-limitations) and organizational strategies to be aware of.

There are some limitations to what assets you can put in an AssetBundle, and how you set them up. The following table contains limitations of AssetBundles:

| **Limitation** | **Description** |
| --- | --- |
| **File types** | * You can’t combine scenes and assets into one AssetBundle. You must store scenes in a separate AssetBundle to one which contains assets. * You can’t include script assets in AssetBundles. * You can’t include files in the StreamingAssets folder. * You can’t assign assets or scenes to more than one AssetBundle. |
| **Naming** | The name of the AssetBundle must differ from the output folder. |
| **Platform support** | AssetBundles can only be loaded on the specific platform you build them for. The Editor can load any AssetBundle, regardless of the platform defined in the [Build Profile](build-profiles.md), but some assets might not load properly if they use a platform-specific format not supported by the operating system the Editor is installed on. |

## AssetBundle naming conventions

Unity automatically puts any AssetBundle names in lowercase during the build process. For example, an AssetBundle named `Foo` will result in a file named `foo`. To avoid any conflicts or issues, use lowercase names for your AssetBundles.

There’s no established or expected file extension for AssetBundles. You can provide a custom extension when defining an AssetBundle name, and Unity respects it. However, this can conflict with the AssetBundle variant feature, which expects variants to be distinguished by different extensions with the same root file name.

You can use the [`BuildAssetBundleOptions.AppendHashToAssetBundleName`](../ScriptReference/BuildAssetBundleOptions.AppendHashToAssetBundleName.md) flag to include an AssetBundle’s content hash as part of its file name. This can be useful for web hosting, allowing different versions of the same AssetBundle to exist together.

However, this flag has downsides:

* The built-in AssetBundle caching doesn’t automatically replace older versions, so multiple versions of the same AssetBundle might stay in the cache.
* It can affect file-level patching systems on console platforms, because small changes are treated as an entirely new file rather than a small patch.

## AssetBundle variants

You can use AssetBundle variants to create multiple versions of an AssetBundle optimized for different situations or configurations, such as varying graphics settings. For example, you can create variants for low, medium, and high resolution textures to support devices with varying performance capabilities or user-selected graphics settings.

When using AssetBundle variants, Unity allows you to specify a variant name for each AssetBundle. The combination of AssetBundle name and variant name uniquely identifies each AssetBundle variant. For example, an AssetBundle named `environment` might have variants named `lowQuality`, `mediumQuality`, and `highQuality`.

The asset file names within the AssetBundles must match, but the content can differ based on the purpose of the variant.

You can create a variant in the **Inspector**, using the variant dropdown next to the AssetBundle, or use [`AssetImporter.assetBundleVariant`](../ScriptReference/AssetImporter-assetBundleVariant.md).

## Strategies for organizing AssetBundles

You can organize your assets in the following ways:

* **[By function](#organize-assets-by-function)**: Organize assets based on functional portions of your project. Typical categories might include user interface, characters, environments, and other elements frequently used throughout your application.
* **[By type](#organize-assets-by-type)**: Type Grouping focuses on bundling assets of the same type together, such as audio tracks or localization files.
* **[By use at runtime](#organize-assets-by-use-at-runtime)**: Group assets that are loaded and used simultaneously. This strategy is often used for level-based projects.

You can mix these strategies within your project to get the best efficiency for your application. For example, you might group UI elements for different platforms into one AssetBundle, but group its interactive content by level or **scene**.

Whatever organizational approach you take, it’s best practice to follow these general guidelines:

* Separate frequently updated objects from rarely changed ones into different AssetBundles.
* Create separate AssetBundles for sets of objects unlikely to be loaded at the same time, such as standard and high definition assets. You can use [AssetBundle variants](#assetbundle-variants) to organize this.
* Group objects together that are commonly unloaded together. For example, a model, its textures, and its animations.
* Split an AssetBundle if less than half of its content is ever loaded at the same time.
* Combine small AssetBundles if the content is frequently loaded simultaneously.
* If multiple objects in one AssetBundle depend on a single asset from another AssetBundle, consider moving the dependency to its own AssetBundle. Similarly, if several AssetBundles reference the same group of assets, consider putting the dependencies into a shared AssetBundle.

### Organize assets by function

Organizing assets by how they functionally work together is useful in situations like the following:

* Bundling all the textures and layout data for a UI screen.
* Bundling all the models and animations for a character or set of characters.
* Bundling the textures and models for pieces of the scenery shared across multiple levels.

This approach is ideal for downloadable content (DLC), because you can make small changes to your project without having to distribute large amounts of unchanged assets to your users. However, when using this approach, you must be familiar with where and when each asset is used in the project.

### Organize assets by type

You can organize assets of similar type, such as audio tracks or language localization files, into a single AssetBundle, which can be useful in AssetBundles that change rarely.

Grouping AssetBundles this way can result in fewer AssetBundles changing and requiring distribution when creating an incremental build. However, more AssetBundles might need to be downloaded and loaded to assemble all dependent objects together at runtime.

### Organize assets by use at runtime

You can group together assets that Unity loads and uses at the same time, which is useful if you want to load AssetBundles based on scenes. For example, you can use this approach for a level-based game where each level contains unique assets, and an AssetBundle contains all a scene’s dependencies.

However, this approach means that an asset in one AssetBundle can only be used at the same time that the rest of the assets are going to be used, or it increases load times.

**Important** An AssetBundle containing a scene automatically includes all assets referenced in that scene unless they are explicitly assigned to a separate AssetBundle. This might lead to duplicated assets if any other scenes use the referenced assets in other AssetBundles.

## Additional resources

* [Managing assets with the Asset Database](AssetDatabase.md)
* [Build assets into an AssetBundle](AssetBundles-Building.md)
* [Avoiding asset duplication](assets-avoid-duplication.md)