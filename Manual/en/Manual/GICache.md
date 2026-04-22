# Global Illumination (GI) cache

The **Global Illumination** (GI) cache is an internal data cache that the Unity Editor uses to store intermediate files when it precomputes lighting data for Baked Global Illumination and **Enlighten** Realtime Global Illumination. By storing this data in a cache, Unity can speed up subsequent precomputations.

All Unity projects on your computer share the cache, so projects with similar content and the same lightmapping backend can use the same cached files. You can also copy the `GiCache` folder between different machines. This can make your lighting build faster, because Unity fetches files from the `GiCache` folder instead of recomputing them.

You can manage the size, location, and **compression** of the cache using the **GI cache** preferences. For more information, refer to [Preferences](Preferences.md).

You can also set a custom location for the GI cache folder. For more information, refer to [Unity Editor special command line arguments](EditorCommandLineArguments.md).

If you try to bake with more than one instance of the Editor running on the same computer, the Editor displays the following warning message: “The GI Cache is using increasing amounts of space on your hard drive to support concurrent **lightmap** generation. To prevent failed bakes, close all other instances of the Unity Editor.”

## Clear the GI cache

For details on how to clear the GI cache, refer to the [GI Cache Preferences](Preferences.md) documentation. You can’t clear the GI cache by selecting **Clear Baked Data** in the Lighting window.

You should only clear the GI cache as a last resort. If your project has an issue that forces you to clear the GI cache, [report it as a bug](https://unity.com/releases/editor/qa/bug-reporting).