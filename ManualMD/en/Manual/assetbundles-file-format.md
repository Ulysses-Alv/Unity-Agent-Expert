# AssetBundle file format reference

When you use `BuildPipeline.BuildAssetBundles` to [build AssetBundles](AssetBundles-Building.md), Unity writes the following files to a specified output directory:

* [The AssetBundle files](#assetbundle-files).
* [A `.manifest` file for each AssetBundle file](#manifest-files).
* [A manifest bundle](#manifest-bundles).

**Note:** If you use the [Addressables package](http://docs.unity3d.com/Packages/com.unity.addressables@latest) to build AssetBundles, it produces an AssetBundle file but no manifest files or manifest bundles.

## Unity archive file format

The Unity archive file format is a generic packaging format that can store any type of file, similar to a .zip file. Archive files are [mounted](https://en.wikipedia.org/wiki/Mount_(computing)) into Unity’s [virtual file system](https://en.wikipedia.org/wiki/Virtual_file_system) (VFS), which allows them to be accessed in a uniform way across different platforms.

Unity uses the archive format for AssetBundles, where archive files are created as a final stage of the [AssetBundle build process](AssetBundles-Building.md). They are mounted into the Unity virtual file system when the AssetBundle is [loaded](AssetBundles-Native.md). The archive format is also used for Players built with LZ4 **compression**, in which case the archive is mounted automatically when the Player runs.

Usually you don’t need to interact with archives at a low level. However, Unity provides the API [`ArchiveFileInterface`](../ScriptReference/Unity.IO.Archive.ArchiveFileInterface.md) to manage archive files directly with low-level APIs as needed.

## AssetBundle files

The AssetBundle file is an [archive file](#archive-format) that contains multiple files which loads in assets at runtime. The following diagram contains an example of an AssetBundle file layout:

![Internal structure of a typical AssetBundle file depicted as a series of containers. ](../uploads/Main/AssetBundles-Building-0.png)

Internal structure of a typical AssetBundle file depicted as a series of containers.

* It contains a series of nested containers, where the outermost is an instance of the archive file system with the same name as the AssetBundle.
* In this example, the `ArchiveFileSystem` archive file contains two files:
  + The main AssetBundle file, in Unity’s serialized file format. This file contains the [`AssetBundle`](../ScriptReference/AssetBundle.md) object and the objects from all the assets included in the AssetBundle.
  + One or more **audio clips** stored inside a file with the `.resource` extension.

The way Unity organizes the content inside an AssetBundle is called the build layout. The build layout depends on whether the AssetBundle contains assets or **scenes**:

* **Asset AssetBundles**: Contains all the objects from the assets in a single serialized file inside the Unity archive file. It also contains all the objects referenced by those assets, unless another AssetBundle exposes the referenced asset. The same referenced object might be duplicated in multiple AssetBundles.
* **Scene AssetBundles**: Has a similar layout for a [Player build](build-content-output.md). Each scene has a serialized file with the objects from that scene, with a file name like `PlayerBuild-SceneFileName`. This corresponds to the `level0` files in a Player build. Unity stores assets referenced from a scene file in a `sharedasset` file, for example `PlayerBuild-SceneFileName.sharedasset`, unless another AssetBundle exposes the referenced asset. Because the `sharedasset` calculation only covers scenes that are in the same AssetBundle, if scenes are stored in individual AssetBundles rather than storing them together in a single AssetBundle, then Unity might create a significant amount of duplicated objects. However, there can be performance and distribution advantages to storing scenes in separate AssetBundles.

The naming convention for the serialized file inside an archive is as follows:

* Asset AssetBundles name the serialized file as `CAB-`, followed by the MD4 hash of the AssetBundle name. For example, `CAB-cc6c60ef8808e0fc6663136604321554`.
* Scene AssetBundles created with [`BuildPipeline.BuildAssetBundles`](../ScriptReference/BuildPipeline.BuildAssetBundles.md) use a name that begins with `PlayerBuild-` followed by the scene’s file name (without path and extension). For example, `Assets/Scene1.unity` becomes `PlayerBuild-Scene1`. Because of this convention give each scene file a unique file name to avoid conflicts.
* Scene AssetBundles created by [Addressables](https://docs.unity3d.com/Packages/com.unity.addressables@latest) name each scene serialized file as `CAB-`, followed by the MD4 hash of the project relative path of the scene. The accompanying `sharedasset` file has the same file name, followed by the `.sharedasset` extension.

As with Player builds, each serialized file inside an AssetBundle can be accompanied by a `.resS` and a `.resource` file to store large binary data. All audio or video is stored in `.resource` files, and textures and meshes are stored in the `.resS` format. These files are named after the serialized file they accompany, for example `CAB-cc6c60ef8808e0fc6663136604321554.resource` or `PlayerBuild-Scene1.sharedasset.resS`.

## Examining AssetBundle file contents

The Unity Editor installation includes the `WebExtract` and `Binary2Text` executables which you can use to extract the files nested inside the AssetBundle and produce a text format dump of the contents of a binary `SerializedFile`. Its output is similar to the YAML format used by Unity.

For more information about these tools, refer to [Analyzing AssetBundles](assetbundles-analyze.md).

## Manifest files

For every AssetBundle generated, Unity generates an associated manifest file. The manifest file has the extension `.manifest` and you can open it with any text editor.

It includes the following content:

* The hashes required for incremental build calculations and content verification. For more information, refer to [`BuildPipeline.GetHashForAssetBundle`](../ScriptReference/BuildPipeline.GetHashForAssetBundle.md).
* The cyclic redundancy check (CRC). You can get this information with [`BuildPipeline.GetCRCForAssetBundle`](../ScriptReference/BuildPipeline.GetCRCForAssetBundle.md).
* The list of scenes or assets in the AssetBundle.
* The list of any AssetBundles that this AssetBundle depends on, as absolute paths.
* Type usage information that is used for type stripping. This includes usage of Unity objects, MonoBehaviour and ScriptableObject-derived classes, and [`SerializeReference`](../ScriptReference/SerializeReference.md) types. For more information, refer to [`BuildPlayerOptions.assetBundleManifestPath`](../ScriptReference/BuildPlayerOptions-assetBundleManifestPath.md).

The following is an example of the contents of an AssetBundle manifest file:

```
ManifestFileVersion: 0
UnityVersion: 6000.2.0a6
CRC: 4208470199
Compression: None
Hashes:
  AssetFileHash:
    serializedVersion: 2
    Hash: 81197c4674c1f389b3568a0aa1b41119
  TypeTreeHash:
    serializedVersion: 2
    Hash: 3c2131fb3360d17991621f547033218e
  IncrementalBuildHash:
    serializedVersion: 2
    Hash: 489e266cfc1b361a94c3efc39afecb54
HashAppended: 0
ClassTypes:
- Class: 1
  Script: {instanceID: 0}
- Class: 4
  Script: {instanceID: 0}
SerializeReferenceClassIdentifiers: []
Assets:
- Assets/Scenes/Scene2.unity
- Assets/Scenes/SampleScene.unity
Dependencies:
- C:/MyBuild/audio.bundle
- C:/MyBuild/sprites.bundle
```

Unity uses the `.manifest` files for its [incremental build pipeline](building-introduction.md). When performing a build, Unity checks the existing AssetBundles and `.manifest` files to determine if the AssetBundle needs to be rebuilt or can be reused. If you delete the `.manifest` files then Unity always rebuilds the AssetBundles from scratch.

Manifest files aren’t required to load the AssetBundles, so you don’t need to distribute them. If `BuildAssetBundleOptions.AppendHashToAssetBundleName` is used, the hash is appended to the AssetBundle file name, but this hash isn’t included in the `.manifest` file name.

### Root manifest file

In addition to the .manifest file for each AssetBundle, Unity generates a root `.manifest`, named after the build folder itself (for example, `MyBuildFolder.manifest`). This file lists the generated AssetBundles and their dependencies, with paths relative to the build directory. It also contains the CRC of the manifest AssetBundle. This root manifest is crucial for code stripping during Player builds, because its path can be passed to `BuildPipeline.BuildPlayer` to prevent script types and Unity modules required by AssetBundles from being stripped.

The following is an example of a root manifest file:

```
ManifestFileVersion: 0
CRC: 2309754985
AssetBundleManifest:
  AssetBundleInfos:
    Info_0:
      Name: bundle_prefab
      Dependencies:
        Dependency_0: bundle_sobject
    Info_1:
      Name: bundle_sobject
      Dependencies: {}
```

## Manifest bundles

Unity additionally generates a manifest bundle, which is an AssetBundle file named after the directory it’s located in. It contains the [`AssetBundleManifest`](../ScriptReference/AssetBundleManifest.md) object which Unity uses to determine which bundle dependencies to load at runtime.

There are two additional files generated by the build.

The first is a small AssetBundle that is named after the directory that it’s located in (where the AssetBundles were built to). This file is called the `Manifest Bundle` and it contains the [AssetBundleManifest](../ScriptReference/AssetBundleManifest.md) object which will be useful for figuring out which bundle dependencies to load at runtime. For information on how to use this bundle, refer to [Loading assets from bundles](AssetBundles-Native.md).

The manifest bundle also includes its own `.manifest` file. The following is an example of a manifest bundle’s manifest file:

```
ManifestFileVersion: 0
AssetBundleManifest:
  AssetBundleInfos:
    Info_0:
      Name: scene1assetbundle
      Dependencies: {}
```

The `.manifest` file for the manifest bundle records how AssetBundles relate, and what their dependencies are. This information is similar to the information recorded by the `AssetBundleManifest` object inside the manifest bundle.

The manifest file is important to prevent code stripping for unused types that you use in your AssetBundle. If you have enabled [code stripping](managed-code-stripping.md) in your project, set [`BuildPlayerOptions.assetBundleManifestPath`](../ScriptReference/BuildPlayerOptions-assetBundleManifestPath.md) to pass the path to this manifest when performing player builds.

## Build Report

AssetBundle builds additionally generate a `BuildReport` file which is a Unity `SerializedFile` written to `Library/LastBuild.buildreport` inside the project directory. This file is useful to view a summary of the timings for build steps and a detailed view of the contents of the AssetBundle. You can use the [`BuildReport`](../ScriptReference/Build.Reporting.BuildReport.md) API to read information from the `BuildReport` file.

You can also use the following tools to view the content of the `BuildReport` file:

* [Project Auditor](https://docs.unity3d.com/Packages/com.unity.project-auditor@latest)
* [Build Report Inspector](https://docs.unity3d.com/Packages/com.unity.build-report-inspector@latest)

## Additional resources

* [Loading assets from AssetBundles](AssetBundles-Native.md)
* [Handling dependencies between AssetBundles](AssetBundles-Dependencies.md)
* [`AssetBundleManifest` API documentation](../ScriptReference/AssetBundleManifest.md)