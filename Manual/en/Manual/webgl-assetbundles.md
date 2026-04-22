# AssetBundles in Web

**Note**: Unity’s [Addressables system](https://docs.unity3d.com/Packages/com.unity.addressables@latest?subfolder=/manual/index.html) is the recommended approach for asset loading. The following content applies only if you use AssetBundles directly.

As all asset data must be pre-downloaded before your content starts, consider moving assets out of your main data files and into [AssetBundles](AssetBundlesIntro.md). Doing so creates a smaller loader **scene** that loads quickly, while AssetBundles dynamically load assets on-demand as the user proceeds through your content. AssetBundles also help with [Asset data memory](webgl-memory.md) management: You can unload asset data from memory for assets that you don’t need any more by calling [AssetBundle.Unload](../ScriptReference/AssetBundle.Unload.md).

The following considerations apply when using AssetBundles on the Web platform:

* When you use class types in your AssetBundle that aren’t used in your main build, Unity might strip the code for those classes from the build. This can cause a fail when trying to load assets from the AssetBundle. Use [BuildPlayerOptions.assetBundleManifestPath](../ScriptReference/BuildPlayerOptions-assetBundleManifestPath.md) to fix that issue, and refer to [Distribution size and code stripping](webgl-distributionsize-codestripping.md) for other options.
* The **WebGL** graphics API doesn’t support threading. As HTTP downloads become available only after they’re downloaded, Unity Web platform builds need to decompress AssetBundle data on the main thread after the download is complete, blocking the main thread. To avoid this interruption, [LZMA AssetBundle compression](assetbundles-compression-format.md) isn’t available for AssetBundles on Web. AssetBundles are compressed using LZ4 instead, which is de-compressed efficiently on-demand. If you need smaller **compression** sizes than LZ4 delivers, configure your web server to use gzip or Brotli compression (on top of LZ4 compression) on your AssetBundles. Refer to [Deploying compressed builds](webgl-deploying.md) for more information on how to do this.
* The Web platform doesn’t support file-system-based AssetBundle caching with [UnityWebRequestAssetBundle.GetAssetBundle](../ScriptReference/Networking.UnityWebRequestAssetBundle.GetAssetBundle.md). Instead the browser caches the WebRequest responses; refer to [Cache behavior in Web](webgl-caching.md) for details. This means that the entire AssetBundle file is held in memory when an AssetBundle loads, and you must promptly unload unused AssetBundles, especially when they’re large.

## Additional resources

* [Addressables package](https://docs.unity3d.com/Packages/com.unity.addressables@latest?subfolder=/manual/index.html)
* [Distribution size and code stripping](webgl-distributionsize-codestripping.md)
* [Optimize your Web build](web-optimization.md)