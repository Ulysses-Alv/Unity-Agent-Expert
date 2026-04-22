# AssetBundle caching

Unity’s built-in disk-based caching system prevents redundant downloads by storing AssetBundles downloaded via [`UnityWebRequestAssetBundle`](../ScriptReference/Networking.UnityWebRequestAssetBundle.md).

Unity converts AssetBundles in the cache to LZ4 for optimal performance. If [`Caching.compressionEnabled`](../ScriptReference/Caching.compressionEnabled.md) is `false`, AssetBundles are stored uncompressed. Use the [`Caching`](../ScriptReference/Caching.md) class to manage cached AssetBundles, for example to clear or check cache presence.

If you use the Web platform, refer to [Web caching](webgl-caching.md) for more information on how to handle caching.

If you use [`UnityWebRequestAssetBundle.GetAssetBundle`](../ScriptReference/Networking.UnityWebRequestAssetBundle.GetAssetBundle.md), only those that accept a `Hash128` or `uint version` argument use the built-in cache.

When downloading without caching, Unity downloads the AssetBundle into a memory file then loads it. If the downloaded file is in LZMA format then Unity recompresses it into a memory file in LZ4 format and loads that instead.

## Location of the AssetBundle cache

The local AssetBundle cache exists at a root folder (for example, `Application.temporaryCachePath`), with a structure like the following:

```
RootCacheFolder
- BundleName1 
  - Hash1 __data (AssetBundle content) 
  - Hash1__info (cache metadata)
- BundleName2
  - Hash2 __data 
  - Hash2 __info
```

During download, Unity uses a temporary directory for the incoming data, and then establishes the recompressed AssetBundle in this structure.

By default, the last component of the download URL is used as the AssetBundle name for the cache. To change the name, you can specify it in the `CachedAssetBundle.name` structure, which can be passed to some `UnityWebRequestAssetBundle.GetAssetBundle` signatures. Note that the cache is based on the AssetBundle name, not the full URL.

## Cache version hash

The AssetBundle cache uses hash values as the version of an AssetBundle. You can provide a 128-bit hash (`Hash128`) or 32-bit `uint` value to distinguish versions. This can be a build-generated hash, a numeric version, or a custom value. It’s crucial to update this version when a new AssetBundle build is released so devices download the correct new version rather than using a stale cached one.

**Important**: Unity doesn’t automatically perform hash calculations on AssetBundles during download and load to verify content against a provided version hash. The version hash is only an identifier. [Cyclic redundancy checks (CRC)](AssetBundles-Integrity.md#cyclic-redundancy-checks) are available for content validation.

You can also use [`BuildAssetBundleOptions.ForceRebuildAssetBundle`](../ScriptReference/BuildAssetBundleOptions.ForceRebuildAssetBundle.md) to perform a full rebuild, and [`BuildAssetBundleOptions.AssetBundleStripUnityVersion`](../ScriptReference/BuildAssetBundleOptions.AssetBundleStripUnityVersion.md) to prevent changes to the hash between Unity versions.

## Cache clearing

Unity performs cache clearing automatically on desktop systems. It checks the cache on startup, and clears any AssetBundles that haven’t been accessed in [`Cache.expirationDelay`](../ScriptReference/Cache.expirationDelay.md) in the last 150 days.

The platform providers for mobile devices maintain their caches. To explicitly clear the cache on mobile platforms use [`Cache.ClearCache`](../ScriptReference/Cache.ClearCache.md).

## Additional resources

* [Loading assets from AssetBundles](AssetBundles-Native.md)
* [Verifying downloaded AssetBundles](AssetBundles-Integrity.md)
* [`UnityWebRequestAssetBundle` API documentation](../ScriptReference/Networking.UnityWebRequestAssetBundle.md)