# AssetBundle compression formats

AssetBundle files are an archive format that contains the following structures:

* A small header data structure. This is never compressed.
* A content section containing its virtual files, which you can optionally compress.

You can compress AssetBundle files in the following formats:

* [Full file LZMA compression](#lzma-compression).
* [Chunk-based LZ4 compression](#lz4-compression).
* [Uncompressed format](#uncompressed-format).

By default, Unity compresses the content section with LZMA and caches AssetBundles with LZ4.

Because different data compresses with different size savings, it’s best practice to rebuild your project with each supported option and measure the size difference.

If you download and store data with a custom caching solution you can use [`AssetBundle.RecompressAssetBundleAsync`](../ScriptReference/AssetBundle.RecompressAssetBundleAsync.md) to change **compression**, for example to convert LZMA format AssetBundles to uncompressed or LZ4 format after download.

**Note:** If you use the Addressables system, you can set the compression format in the Unity Editor, in the group settings **Inspector**. For more information, refer to the [Addressables documentation](https://docs.unity3d.com/Packages/com.unity.addressables@latest?subfolder=/manual/ContentPackingAndLoadingSchema.html).

## LZMA compression

LZMA compresses the entire content section of the AssetBundle file as a single stream. This full content approach results in smaller file sizes than chunk-based LZ4 compression. LZMA is the preferred format for AssetBundles downloaded from a Content Delivery Network (CDN).

However, you must decompress the entire file into RAM to read an asset from archives compressed with LZMA. This approach works best when an AssetBundle contains assets that you want to load all at once. For example, packaging all assets for a character or **scene** in one AssetBundle.

LZMA is the format used when calling `BuildPipeline.BuildAssetBundles` with no specific compression specified, for example `BuildAssetBundleOptions.None`.

**Note:** The Web platform doesn’t support LZMA compression. Use LZ4 compression instead. For more information, refer to [AssetBundles in Web](webgl-assetbundles.md).

## LZ4 compression

LZ4 uses a chunk based algorithm which decompresses the AssetBundle in chunks. While writing the AssetBundle, Unity individually compresses each 128 KB chunk of the content before saving it. This approach means the total AssetBundle file size is larger than LZMA compression. However, you can selectively retrieve and load just the chunks needed for a requested object, rather than decompressing the entire AssetBundle.

LZ4 has comparable loading times to uncompressed bundles with the added benefit of reduced size on disk. This is the format preferred by the [AssetBundle cache](assetbundles-caching.md), and it can also be a good choice for AssetBundles that you distribute as part of your installation or in other cases where size isn’t important.

To use this format, specify the flag [`BuildAssetBundleOptions.ChunkBasedCompression`](../ScriptReference/BuildAssetBundleOptions.ChunkBasedCompression.md) when calling `BuildPipeline.BuildAssetBundles`.

## Uncompressed format

You can build AssetBundles so that the data is completely uncompressed. This creates a larger file download size, but faster load times once downloaded because no decompression is required. Uncompressed AssetBundles are helpful if only a few objects are loaded out of a larger AssetBundle.

Uncompressed AssetBundles are 16-byte aligned. To build uncompressed AssetBundles, specify the flag [`BuildAssetBundleOptions.UncompressedAssetBundle`](../ScriptReference/BuildAssetBundleOptions.UncompressedAssetBundle.md) when calling `BuildPipeline.BuildAssetBundles`.

## Recompression when caching

When an AssetBundle is distributed, it’s typically in LZMA format for optimal file size. If the AssetBundle is downloaded via Unity’s built-in caching support ([`UnityWebRequestAssetBundle`](../ScriptReference/Networking.UnityWebRequestAssetBundle.md)), Unity dynamically recompresses it to LZ4 when it’s downloaded and written to the cache. This conversion optimizes for faster decompression and load times on the client.

You can force AssetBundles to be cached in an uncompressed format by setting [`Caching.compressionEnabled`](../ScriptReference/Caching.compressionEnabled.md) to `false`.

## Additional resources

* [Optimizing AssetBundles](assetbundles-optimizing.md)
* [Downloading AssetBundles](AssetBundles-Integrity.md)