# Cache behavior in Web

In the Unity Web platform, the **Cache API** lets you store the asset data cached in .data files and [AssetBundles](AssetBundlesIntro.md) within the browser cache. Storage limits for the browser cache such as maximum file size, maximum overall cache size, and eviction criteria are dependent on the browser and platform that you’re using. For more information, refer to [Browser storage limits and eviction criteria](https://developer.mozilla.org/en-US/docs/Web/API/IndexedDB_API/Browser_storage_limits_and_eviction_criteria).

## Data Caching

To access **Data Caching**, open the **Publishing Settings** for Web from **File** > **Build Profiles** > **Player Settings**. This enables the browser to cache the main data files into the IndexedDB database.

Using the default browser HTTP cache doesn’t guarantee that the browser caches a particular response. This is because the browser HTTP cache has limited space, and the browser might not be able to cache files that are too large.

To improve your loading speed, `IndexedDB` allows you to cache files above the browser limit. When you cache more files, you increase the chance that downloaded content is available on the user’s machine during the next run of the build.

To allow the browser to cache the main data files (.data) in the `IndexedDB` cache, enable **Data Caching** in **Player** settings (menu: **Edit** > **Project Settings** > **Player**). By default, the standard Web loader configuration also caches [AssetBundles](AssetBundlesIntro.md) and revalidates them. You don’t need to override the cache behavior unless you want to change this default (for example, to stop caching bundles or to cache other file types).

If you do want to customize cache control via the instance directly, make sure `Module.cacheControl(url)` returns `must-revalidate` for the requested AssetBundle URL to keep caching enabled. For example, you can override the `unityInstance.Module.cacheControl()` function in the fulfillment callback of the Promise that `createUnityInstance()` returns.

For further information on `createUnityInstance()`, refer to [Compressed builds and server configuration](webgl-deploying.md).

## Customize Web Cache behavior

By default, the cache configuration stores the asset data file .data and AssetBundle files .bundle, and revalidates them before loading them from the cache. To change this behavior, add a new [Web template](webgl-templates.md) that changes the `UnityLoader` configuration.

The following example shows how to add a custom `cacheControl` function to the `UnityLoader` configuration within the `index.html` file.

```
var config = {
   // ...
#if USE_DATA_CACHING
   cacheControl: function (url) {
     // Caching enabled for .data and .bundle files.
     // Revalidate if file is up to date before loading from cache
     if (url.match(/\.data/) || url.match(/\.bundle/)) {
         return "must-revalidate";
     }

     // Caching enabled for .mp4 and .custom files
     // Load file from cache without revalidation.
     if (url.match(/\.mp4/) || url.match(/\.custom/)) {
         return "immutable";
     }

     // Disable explicit caching for all other files.
     // Note: the default browser cache may cache them anyway.
     return "no-store";
   },
#endif
   // ...
}
```

The `cacheControl` function takes the url of a request as a parameter and returns one of the following:

* `must-revalidate` - If the function returns [must-revalidate](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control#must-revalidate), the cache returns to an enabled state and the file is revalidated before being loaded from the cache.
* `immutable` - If the function returns [immutable](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control#immutable), the cache is enabled and the file is loaded from the cache without revalidation.
* `no-store` - If the function returns [no-store](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control#no-store), the cache is disabled.

The browser automatically stores (caches) certain file types such as .html, .js, .css, .json, .jpg, .png, so they don’t need to be explicitly stored in the Web Cache. Typical candidates for the Web cache include large files and files that use a custom file format.