# AssetBundle platform considerations

AssetBundles are platform specific, meaning at runtime you can only load AssetBundles built for the target platform. This is because AssetBundles contain platform-specific asset formats, which leads to unexpected behavior if loaded on other platforms. This restriction prevents the accidental delivery of incorrect content to the wrong platform.

## Android specifics

To distribute AssetBundle content for the Android platform, you have the following options:

* Build AssetBundles into the [`StreamingAssets` folder](StreamingAssets.md).
* Package AssetBundles into custom asset packs.
* Deploy AssetBundles through a CDN.

The option you choose depends on the requirements of your project.

### Build AssetBundles into the StreamingAssets folder

AssetBundles placed in the [`StreamingAssets` folder](StreamingAssets.md) are automatically packaged as part of your application’s **APK** or OBB file.

You can’t load these files directly from the APK or OBB file with the standard file path APIs. Instead, you must load them with the [`UnityWebRequestAssetBundle`](../ScriptReference/Networking.UnityWebRequestAssetBundle.md) API with a `file:///` URL scheme. For example:

```
var assetBundlePath = "file://" + Application.streamingAssetsPath + "/" + assetbundleToLoad;
var loadOp = UnityWebRequestAssetBundle.GetAssetBundle(assetBundlePath);
yield return loadOp.SendWebRequest();
var assetBundle = DownloadHandlerAssetBundle.GetContent(loadOp);
```

Avoid using LZMA **compression** for AssetBundles in the `StreamingAssets` folder because it’s inefficient to decompress LZMA files to a temporary memory file. Use LZ4 compression or leave them uncompressed. For more information, refer to [AssetBundles compression format](assetbundles-compression-format.md).

### Package AssetBundles into custom asset packs

You can place AssetBundles into Android’s asset pack format and deliver them to the Google Play Store. For more information, refer to [Asset packs in Unity](android-asset-packs-in-unity.md).

### Deploy AssetBundles through a CDN

You can host AssetBundles on a web server and download them as needed with [`UnityWebRequestAssetBundle`](../ScriptReference/Networking.UnityWebRequestAssetBundle.md).

You can use Unity’s [Cloud Content Delivery (CCD)](https://docs.unity.com/ugs/manual/ccd/manual/UnityCCD) service to simplify deployment, especially when integrated with Addressables. This approach is suitable for applications with large, dynamic content. Downloaded AssetBundles are typically saved to the application cache directory or `Application.persistentDataPath`.

## Console downloads and CRC

Typically, AssetBundles for consoles aren’t downloaded directly from the web using in-game code. Instead, they’re usually delivered via the platform’s own DLC (Downloadable Content) mechanisms. These DLCs are securely purchased, downloaded, and stored locally by the platform’s system, external to the game’s direct control. Once installed, the game can detect and load these new files.

When loading an AssetBundle on console platforms, don’t perform CRC. Consoles often have weaker CPUs that can take significant time to validate the CRC for a newly opened AssetBundle. Even on platforms with fast storage, a CPU-bound CRC check can slow down loading considerably.

## Additional resources

* [Downloading AssetBundles](AssetBundles-Integrity.md)
* [AssetBundle compression formats](assetbundles-compression-format.md)
* [AssetBundles in Web](webgl-assetbundles.md)