# Introduction to AssetBundles

An AssetBundle is an archive file that you can use to group together assets to create downloadable content (DLC), or reduce the initial installation size of your application. You can also use AssetBundles to load optimized platform-specific assets, or lower memory usage at runtime.

An AssetBundle can contain platform-specific non-code assets, such as models, textures, **prefabs**, **audio clips**, or entire **scenes** that Unity then loads at runtime. AssetBundles are platform-specific because Unity builds asset data into formats based on the [`BuildTarget`](../ScriptReference/BuildTarget.md) you set when you [create a build](building-introduction.md). For example, an AssetBundle built for iOS isn’t compatible for Android.

You can also [compress AssetBundles](assetbundles-compression-format.md) using LZMA or LZ4 to efficiently distribute the archives.

To build and define AssetBundles, you can use the high-level [Addressables](http://docs.unity3d.com/Packages/com.unity.addressables@latest) package, which provides a way of defining and building AssetBundles from the Unity Editor. If you prefer low-level API control, you can use the [`BuildPipeline.BuildAssetBundles`](../ScriptReference/BuildPipeline.BuildAssetBundles.md), [`AssetBundle`](../ScriptReference/AssetBundle.md), and [`UnityWebRequestAssetBundle`](../ScriptReference/Networking.UnityWebRequestAssetBundle.md) native APIs.

## Reasons to use AssetBundles

Using AssetBundles can help with content distribution and optimizing your application’s performance. The following are benefits of using the AssetBundle system:

* **Dynamic content delivery**: You can use AssetBundles to load assets on demand, which is especially useful for games with downloadable content (DLC), episodic updates, or live service models. It also helps efficiently manage memory, ensuring only the assets that your application needs are loaded into memory.
* **Reduced build size**: Moving assets to AssetBundles can reduce the initial application build size, which is important for mobile games or platforms with strict size limitations.
* **Platform compatibility**: You can create AssetBundles for different platforms, reducing the need to include platform-specific assets in your application’s build.

AssetBundles are useful if you want to optimize asset loading, such as only streaming content near a character’s location, only loading relevant localized content, or loading assets in the background. However, the AssetBundle system provides low-level asset management, so you might want to consider using the [Addressables](http://docs.unity3d.com/Packages/com.unity.addressables@latest) package, which provides a higher level way of managing AssetBundles in your project.

If you’re prototyping, or have a particularly small project, you might want to consider using the [resources system](LoadingResourcesatRuntime.md).

## Structure of an AssetBundle

The AssetBundle is a container file format, similar to a zip file. It contains the following file types in a binary-format header:

* **Serialized files**: Contains serialized Unity objects. This is the same binary file format used in Player builds. The output depends on what the AssetBundle contains:
  + Only assets: Unity creates a single serialized file.
  + Only scenes: Unity creates two serialized files per scene. One file contains the objects from the scene hierarchy, and the second contains any referenced objects.
* **Resource files:** Contains chunks of binary data stored separately for certain assets such as textures and audio. This separation allows Unity to use multithreaded code to efficiently load the assets from disk.

The AssetBundle file always includes a serialized [`AssetBundle`](../ScriptReference/AssetBundle.md) object that acts like a directory for the contents of the AssetBundle. You can use the `AssetBundle` instance to load assets from a specific AssetBundle archive in your code. For a more detailed look at the internal file format, refer to [AssetBundle file format reference](assetbundles-file-format.md).

## Scene AssetBundles

An AssetBundle not containing scenes is built based on a list of assets. Unity supports assigning scene files to an AssetBundle, but you can’t mix scenes with other assets inside a single AssetBundle. In the API, this type of AssetBundle is called a streaming scene AssetBundle, accessible via [`AssetBundle.isStreamedSceneAssetBundle`](../ScriptReference/AssetBundle-isStreamedSceneAssetBundle.md).

The process of building scenes into AssetBundles is similar to what happens during a Player build, reusing much of the same code.

## AssetBundle support between Unity versions

Any AssetBundles you create with older versions of Unity are usually compatible with newer versions of Unity. However, if there are large changes between versions, Unity might not be able to load the data, and you must rebuild the AssetBundle with a newer version of Unity. Unity doesn’t support forward-compatibility of AssetBundles, so you can’t load an AssetBundle built with a newer version of Unity into an older version of Unity.

If the serialization format for an object has changed, AssetBundle load code reads that object using the safe binary read deserialization method. This method uses type trees to match fields in the serialized data with the current object serialization layout, which affects performance. If AssetBundles are built without type trees (`BuildAssetBundleOptions.DisableWriteTypeTree`), any serialization changes in the newer Unity version prevent a successful load and might lead to a crash. Additionally, safe binary reads are slow, and you can avoid this by rebuilding AssetBundles to match the current Player build.

**Tip**: By default, Unity includes the version of the Unity Editor used to build the AssetBundle inside the AssetBundle header. This information can result in AssetBundles unnecessarily being rebuilt. To avoid this, exclude the Editor version from the header. For more information, refer to [`BuildAssetBundleOptions.AssetBundleStripUnityVersion`](../ScriptReference/BuildAssetBundleOptions.AssetBundleStripUnityVersion.md).

Unity serializes AssetBundles based on the Unity version and the C# types present during their creation. Unity stores the information in type tree structures and uses this information when it loads the objects from different versions of the Unity Editor. For more information on TypeTrees, refer to [Optimizing AssetBundles](assetbundles-optimizing.md#typetrees)

## Script support

AssetBundles can’t contain assemblies, so you can’t use them to distribute new C# classes or changes to existing classes. However, you can use an AssetBundle to distribute serialized object instances, such `ScriptableObject` assets.

Unity matches objects based on their assembly, namespace, and class names. It then creates an object that’s an instance of that class, using the serialized values to set the fields of the object. It uses the information stored in the type tree to adjust field mappings for different versions of Unity.

Unity uses the [conditional compilation](platform-dependent-compilation.md) information in your code to determine the fields to include in an AssetBundle. If a field’s compilation has been made conditional with an `#if` directive and the associated symbol isn’t defined at AssetBundle build time, then Unity doesn’t include that field in the AssetBundle.

For example, in the following code snippet the `always` field is unconditionally included in the AssetBundle, but the `experimental` field is only included if the `EXPERIMENTAL_FEATURE` symbol is defined at build time:

```
public class MyData : ScriptableObject
{
    public int always;
#if EXPERIMENTAL_FEATURE
    public int experimental;
#endif

public int ConditionalDataValue()
{
#if EXPERIMENTAL_FEATURE
    return experimental;
#else  
    return always;
#endif 
}

}
```

## Building AssetBundles

You can use the following method to build AssetBundles:

* **Addressables**: The [Addressables](http://docs.unity3d.com/Packages/com.unity.addressables@latest) package is a user-friendly way of defining and building AssetBundles from the Unity Editor. It simplifies AssetBundle creation and management with a high-level API.
* **Native APIs**: `BuildPipeline.BuildAssetBundles`, `AssetBundle`, and `UnityWebRequestAssetBundle` are native APIs you can use to build AssetBundles, but they require manual dependency management, and you must write your own build script to use them.

For more information, refer to [Build assets into an AssetBundle](AssetBundles-Building.md).

### Building multiple AssetBundles

When you build or rebuild AssetBundles, it’s best practice to use a single [`AssetBundle`](../ScriptReference/AssetBundle.md) API call to build all the project’s AssetBundles together. AssetBundles can reference and have dependencies on other AssetBundles. For example, a material in one AssetBundle can reference a texture in another AssetBundle. When you build AssetBundles together, Unity automatically manages the references and dependencies between them.

## Additional resources

* [Addressables package](http://docs.unity3d.com/Packages/com.unity.addressables@latest)
* [`BuildPipeline.BuildAssetBundles` API documentation](../ScriptReference/BuildPipeline.BuildAssetBundles.md)
* [Organizing assets into AssetBundles](AssetBundles-Preparing.md)