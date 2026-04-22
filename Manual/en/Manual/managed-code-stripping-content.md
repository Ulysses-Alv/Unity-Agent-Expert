# How code stripping affects content

When you build a Player in Unity, the [UnityLinker](unity-linker.md) process runs to reduce the size of the assemblies by removing unused types. This process can strip the following:

* Entire C# Unity objects such as `MonoBehaviour` and `ScriptableObject` instances. Unity can strip these from core Unity assemblies like `UnityEngine.dll` and `Unity.dll`, but preserves them in custom assemblies once those assemblies are marked as used.
* Regular C# structures and classes.
* Portions of classes (for example, methods that aren’t used).
* Entire Unity modules, such as AI or Physics, if no type from that module is referenced.

The affect that code stripping has on content appears in the Player as null reference exceptions, missing type errors, or crashes. Additionally, [AssetBundle](AssetBundlesIntro.md) or Addressables content-only build can’t use types that have been stripped out of a Player build.

You can use the [Build Report Inspector’s](https://github.com/Unity-Technologies/BuildReportInspector) Stripping tab to inspect which modules and types have been preserved in a Player build.

You can prevent the linker from stripping code in the following circumstances:

* [Prevent code stripping of types used in AssetBundles](#prevent-code-stripping-assetbundles)
* [Prevent code stripping of types used in Addressables](#prevent-code-stripping-addressables)
* [Prevent module code stripping](#prevent-code-stripping-modules)

## Prevent code stripping of types used in AssetBundles

Unity might strip a Unity object type that’s present in AssetBundles, but not in any **scene** or asset included in a Player build. This means that content from the AssetBundle might work when testing in Play mode, or on builds with minimal code stripping, but errors or crashes occur in a Player built with code stripping enabled.

To resolve this, you can create a custom build script that directly includes the AssetBundle’s manifest file, as follows:

1. [Build the AssetBundle](AssetBundles-Building.md) before building the Player. The output folder includes one `.manifest` file for each AssetBundle, and a root `.manifest` file with the same name as the output folder.
2. Pass the path of the root `.manifest` file into the Player build process with the [`BuildPipeline.BuildPlayer`](../ScriptReference/BuildPipeline.BuildPlayer.md) API.

For more information, refer to [AssetBundle file format reference](assetbundles-file-format.md) and [`BuildPlayerOptions.assetBundleManifestPath`](../ScriptReference/BuildPlayerOptions-assetBundleManifestPath.md).

You can also use the [Preserve attribute](managed-code-stripping-preserving.md), or manually author a [`link.xml`](managed-code-stripping-xml-formatting.md) file in your project, to further influence the code stripping behavior.

## Preventing stripping of types used in Addressables

If a type of Unity object is present in content built with [Addressables](https://docs.unity3d.com/Packages/com.unity.addressables@latest), but not in any scene or asset built into the Player, then Unity might strip that type.

Addressables automatically generates a [`link.xml`](managed-code-stripping-xml-formatting.md) file inside the `Library/com.unity.addressables` folder describing types used in the content. The Player build takes this `link.xml` file into account, provided the Addressables content has been built before, or at the same time as, the Player build.

You can also use the [Preserve attribute](managed-code-stripping-preserving.md), or manually author a [`link.xml`](managed-code-stripping-xml-formatting.md) file in your project, to further influence the code stripping behavior.

## Preventing module stripping

You can’t use a `link.xml` file to prevent entire Unity modules (for example AI or Physics) from being stripped.

If you need to use types from a module in an Addressables build, make sure the Player build includes some content that references that module. For example, to keep the Physics module, add a component such as a **Rigidbody** or **Collider** to a **GameObject** in a scene included in your build. This creates a dependency that the Player build process recognizes, and preserves the module.

## Script-only builds

Player builds record the used types in a `ScriptsOnlyCache.yaml` file inside the `Library/PlayerDataCache` folder, which you can look at to diagnose any code stripping issues. This file is used to support the [Incremental build pipeline](building-introduction.md#incremental-build-pipeline) which can reuse the content from a previous Player build. In that case the Player build passes the cached type usage from that file to the Unity linker.

**Note:** The `ScriptsOnlyCache.yaml` file uses an internal format that might change in future Unity versions. To control code stripping, use `link.xml` files in your project instead of modifying this file directly.

## Additional resources

* [Link XML formatting reference](managed-code-stripping-xml-formatting.md)
* [Addressables package](https://docs.unity3d.com/Packages/com.unity.addressables@latest)
* [Build assets into an AssetBundle](AssetBundles-Building.md)