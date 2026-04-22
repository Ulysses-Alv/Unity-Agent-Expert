# AssetBundle and Addressables determinism

[AssetBundles](AssetBundlesIntro.md) and [Addressables](https://docs.unity3d.com/Packages/com.unity.addressables@latest) have related workflows for ensuring deterministic builds.

## AssetBundle determinism best practices

When building AssetBundles, every referenced asset and dependency must remain identical. Avoid timestamps or build-date suffixes in AssetBundle names and use static identifiers or semantic versions.

By design, you must build AssetBundles deterministically, independent of whether [incremental builds](building-introduction.md#incremental-build-pipeline) are enabled.

If clearing the cache breaks AssetBundle determinism, your build process has non-deterministic elements that need to be fixed. A common way to trace this issue back is as follows:

* Use a **version control** system to identify assets that might have been unintentionally modified.
* Use tools like [UnityDataTools](https://github.com/Unity-Technologies/UnityDataTools/tree/main/UnityDataTool) to find differences in the serialized data contained within the AssetBundles.

For more information about building AssetBundles, refer to [Build assets into an AssetBundle](AssetBundles-Building.md).

## Addressable determinism

Because the Addressables system is built on top of the `AssetBundle` API, all the determinism best practices from the previous AssetBundle section apply to Addressables. Addressable builds are deterministic when the build environment and inputs are identical.

The deterministic behavior of Addressables also depends on both the Scriptable Build Pipeline (SBP) and the [Addressable hash calculation](https://docs.unity3d.com/Packages/com.unity.addressables@latest/index.html?subfolder=/manual/build-content-catalogs.html).

Because SBP calculates the content-based hash from the build outputs, it’s more sensitive to changes compared to AssetBundle hashing, which Unity computes from the build inputs.

In AssetBundle workflows, a non-deterministic build might produce the same cache identifier, such as an input-based hash or a filename, despite generating different binary content. This can lead to reusing outdated cached AssetBundles. SBP hashes the build output to avoid this issue.

Like AssetBundles, both the Addressable assets and all their referenced dependencies must be identical across builds to ensure determinism.

Addressables uses a content hash that SBP generates for caching. The SBP hash calculation isn’t limited to the binary content of the AssetBundle. It also includes the names of all dependent AssetBundles referenced by the target AssetBundle. For more information, refer to [Addressable asset dependencies](https://docs.unity3d.com/Packages/com.unity.addressables@latest/index.html?subfolder=/manual/AssetDependencies.html).

### Editor automation scripts

If your project uses Editor automation, such as custom **scripts** that automatically assign assets to Addressable groups, ensure the following:

* The order of assigned Addressables and Addressable groups are deterministic.
* Each build environment has identical [Addressable Asset Settings](https://docs.unity3d.com/Packages/com.unity.addressables@latest/index.html?subfolder=/manual/AddressableAssetSettings.html).

### Identifying non-deterministic Addressable builds

Every Addressables build automatically creates a build layout report, which includes information about all the built assets in your project. For more information, refer to [Create a build report](https://docs.unity3d.com/Packages/com.unity.addressables@latest/index.html?subfolder=/manual/BuildLayoutReport.html).

For binary-level inspection, use the [UnityDataTools](https://github.com/Unity-Technologies/UnityDataTools) project which can dump or export binaries into readable databases or text files, allowing detailed comparison of serialized data down to individual component fields.

## Additional resources

* [AssetBundles introduction](AssetBundlesIntro.md)
* [Build assets into an AssetBundle](AssetBundles-Building.md)
* [Addressables documentation](https://docs.unity3d.com/Packages/com.unity.addressables@latest)
* [UnityDataTools](https://github.com/Unity-Technologies/UnityDataTools)