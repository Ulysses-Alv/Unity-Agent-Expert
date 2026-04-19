# Handling dependencies between AssetBundles

AssetBundles become dependent on others when their objects reference objects in another AssetBundle. If a referenced object isn’t assigned to any AssetBundle, Unity embeds it in the dependent AssetBundles during the build. If multiple AssetBundles reference the same unassigned object, each AssetBundle contains its own copy, increasing memory usage.

To load dependent AssetBundles, ensure dependencies are loaded into memory before accessing their dependent AssetBundles. For example, if **Bundle 1** contains a material that reference a texture in **Bundle 2**, load **Bundle 2** into memory before accessing the material from **Bundle 1**. Unity doesn’t automatically resolve dependencies. To manage dependencies at runtime, you can use the `AssetBundleManifest`. For more information, refer to [Loading assets from AssetBundles](AssetBundles-Native.md).

## How Unity tracks references

Unity tracks references between AssetBundles in the following ways:

* **Manifest-level dependencies**: Dependencies between AssetBundles are in the build folder’s root `.manifest` file and accessible via [`AssetBundleManifest.GetAllDependencies`](../ScriptReference/AssetBundleManifest.GetAllDependencies.md).
* **Low-level SerializedFile references**: Direct object references between AssetBundles appear in `SerializedFile` text dumps as external references. These low-level references don’t record the name of the AssetBundle and only record the names of the serialized files inside the AssetBundles. They only work when the corresponding AssetBundles have already been loaded. For example:
  + A `Material` in `bundle0` references a `Shader` in `bundle1`.
  + The `SerializedFile` inside `bundle0` has an external reference table in its header. This table has an entry that points to the `SerializedFile` path of `bundle1`.
  + The `m_Shader.m_FileID` of the `Material` object records the index of external reference table corresponding to the SerializedFile inside `bundle1`.
* **AssetBundle object**: The `AssetBundle` object consolidates dependency information in its `m_Container` and `m_PreloadTable`. When a load is requested for an asset, the `m_PreloadTable` ensures all necessary objects, including those from other AssetBundles, are identified for loading. This can involve very large data structures for AssetBundles with many assets and dependencies.

### Script type representation in AssetBundles

The `MonoScript` object in Unity represents a specific `MonoBehaviour`-derived class. This also covers classes derived from ScriptableObject. The `MonoScript` object records the assembly, namespace and class as strings that uniquely identify the type.

When Unity serializes a `MonoBehaviour` object it records the GUID and LFID of the `MonoScript` class, which then directly records the name of the class.

When building AssetBundles Unity includes `MonoScript` objects for each MonoBehaviour-derived class that’s part of the build. These `MonoScript` objects might be inside the same SerializedFile as the MonoBehavior (a local reference), or in an external Serialized file. In both cases MonoScripts are referenced with exactly the same mechanism as other [direct references](assets-direct-reference.md) between objects.

The following actions can result in changes to the MonoScript data:

* Moving the script file to a location that comes under another [assembly definition file](ScriptCompilationAssemblyDefinitionFiles).
* Changing the name of the assembly definition file containing the class.
* Adding or changing the class namespace.
* Changing the class name.

When those changes happen, rebuild the AssetBundles in your project.

## Avoid duplicated information across AssetBundles

By default, Unity doesn’t optimize duplicate data across AssetBundles. For example, if two AssetBundles each contain a **prefab** referencing the same unassigned material, Unity embeds a copy of the material in both AssetBundles. This increases install size, runtime memory usage, and affects batching, because Unity treats each copy as unique.

Assign shared assets to a common AssetBundle to avoid duplication. During the build, Unity automatically includes dependencies in the assigned AssetBundle. This significantly reduces the size of other AssetBundles. For example:

* Extract a shared Material and its dependencies into a `modulesmaterials` AssetBundle.
* Prefab AssetBundles then only reference the `modulesmaterials` AssetBundle, reducing their size.

For more information, refer to [Avoiding asset duplication](assets-avoid-duplication.md).

## Runtime loading

When using a common AssetBundle for shared assets, load it into memory before loading AssetBundles that depend on it. In the following example, the shared Material is loaded correctly because its AssetBundle (`materialsAB`) is loaded first:

```
using System.IO;
using UnityEngine;

public class InstantiateAssetBundles : MonoBehaviour
{
    void Start()
    {
        // Load the AssetBundles
        AssetBundle materialsAB = AssetBundle.LoadFromFile(Path.Combine(Application.dataPath, Path.Combine("AssetBundles", "modulesmaterials")));
        AssetBundle moduleAB = AssetBundle.LoadFromFile(Path.Combine(Application.dataPath, Path.Combine("AssetBundles", "example-prefab")));

        // Check for errors
        if (materialsAB == null || moduleAB == null)
        {
            Debug.Log("Failed to load AssetBundle!");
            return;
        }
        
        GameObject prefab = moduleAB.LoadAsset<GameObject>("example-prefab");
        // Instantiate the prefab
        Instantiate(prefab);
    }
}
```

## AssetBundle unloading

Properly manage dependencies when unloading AssetBundles to prevent crashes or undefined behavior. Dependent AssetBundles must not remain loaded after their dependencies are unloaded. Reloading dependencies separately can also cause issues. When an AssetBundle is loaded, it establishes data structures that specifically point to objects inside dependent AssetBundles. If a referenced AssetBundle is unloaded and then reloaded, its objects are assigned new InstanceIDs and are not reconnected to the dependent AssetBundle, which can lead to crashes or serialization errors.

To avoid this, track dependencies and never unload an AssetBundle if it is referenced by another AssetBundle, unless you also unload that referencing AssetBundle. Implementing a reference-counting system is a safe way to manage AssetBundle unloading.

### Example unload strategy: reference counting

Implement a reference-counting system to track and safely unload AssetBundles only when they’re no longer in use.

The following example tracks dependencies and safely unloads unused AssetBundles:

```
using System.Collections.Generic;
using System.IO;
using UnityEngine;

public class AssetBundleManager
{
    // Path to the directory containing AssetBundles
    private string assetBundlesDirectory;
    // The AssetBundleManifest containing dependency information
    private AssetBundleManifest assetBundleManifest;
    // Reference counts for loaded AssetBundles
    private Dictionary<string, int> assetBundleReferenceCounts = new Dictionary<string, int>();
    // Loaded AssetBundles cache
    private Dictionary<string, AssetBundle> loadedAssetBundles = new Dictionary<string, AssetBundle>();

    // Initializes the AssetBundleManager with the manifest and directory path
    public void Initialize(string manifestBundlePath, string assetBundlesDirectory)
    {
        this.assetBundlesDirectory = assetBundlesDirectory;
        AssetBundle manifestBundle = AssetBundle.LoadFromFile(manifestBundlePath);
        assetBundleManifest = manifestBundle.LoadAsset<AssetBundleManifest>("AssetBundleManifest");
        manifestBundle.Unload(false);
    }

    // Loads an AssetBundle and its dependencies, incrementing reference counts
    public AssetBundle LoadBundle(string bundlePath)
    {
        AssetBundle bundle = LoadAssetBundleIfNotLoaded(bundlePath);
        IncrementReferenceCount(bundle.name);

        string[] dependencyBundleNames = assetBundleManifest.GetAllDependencies(bundle.name);
        foreach (string dependency in dependencyBundleNames)
        {
            string dependencyBundlePath = GetAssetBundlePathFromName(dependency);
            LoadAssetBundleIfNotLoaded(dependencyBundlePath);
            IncrementReferenceCount(dependency);
        }

        return bundle;
    }

    // Loads an AssetBundle if it is not already loaded
    private AssetBundle LoadAssetBundleIfNotLoaded(string bundlePath)
    {
        if (!loadedAssetBundles.TryGetValue(bundlePath, out AssetBundle bundle))
        {
            // For simplicity, this example only shows the case of synchronous loading, but support for
            // LoadFromFileAsync() and the other load methods could also be added with similar code.
            bundle = AssetBundle.LoadFromFile(bundlePath);
            
            if (bundle == null)
            {
                throw new System.Exception($"Failed to load AssetBundle at path {bundlePath}");
            }
            loadedAssetBundles.Add(bundlePath, bundle);
        }

        return bundle;
    }

    // Unloads an AssetBundle and its dependencies if their reference counts reach zero
    public void UnloadBundle(AssetBundle bundle)
    {
        string[] dependencyBundleNames = assetBundleManifest.GetAllDependencies(bundle.name);

        DecrementReferenceCount(bundle.name);
        foreach (string dependency in dependencyBundleNames)
        {
            DecrementReferenceCount(dependency);
        }

        List<string> bundlesToUnload = new List<string>();
        foreach (KeyValuePair<string, AssetBundle> loadedBundleEntry in loadedAssetBundles)
        {
            if (assetBundleReferenceCounts[loadedBundleEntry.Value.name] <= 0)
            {
                bundlesToUnload.Add(loadedBundleEntry.Key);
            }
        }

        foreach (string bundlePath in bundlesToUnload)
        {
            loadedAssetBundles[bundlePath].Unload(true);
            loadedAssetBundles.Remove(bundlePath);
        }
    }

    // Gets the full path of an AssetBundle given its name
    private string GetAssetBundlePathFromName(string name)
    {
        return Path.Combine(assetBundlesDirectory, name);
    }

    // Increments the reference count for a given AssetBundle
    private void IncrementReferenceCount(string bundleName)
    {
        if (assetBundleReferenceCounts.ContainsKey(bundleName))
        {
            assetBundleReferenceCounts[bundleName]++;
        }
        else
        {
            assetBundleReferenceCounts[bundleName] = 1;
        }
    }

    // Decrements the reference count for a given AssetBundle
    private void DecrementReferenceCount(string bundleName)
    {
        if (assetBundleReferenceCounts.ContainsKey(bundleName))
        {
            assetBundleReferenceCounts[bundleName]--;
        }
        else 
        {
            string errorMessage = $"Attempted to decrement reference count for non-existent bundle: {bundleName}";
            throw new KeyNotFoundException(errorMessage);
        }
    }
}
```

**Note:** When using [LZ4 compressed and uncompressed](assetbundles-compression-format.md) AssetBundles, [`AssetBundle.LoadFromFile`](../ScriptReference/AssetBundle.LoadFromFile.md) only loads the catalog of its content in memory, but not the content itself. To check if this is happening, use the [Memory Profiler](https://docs.unity3d.com/Packages/com.unity.memoryprofiler@latest) package to inspect memory usage.

## Additional resources

* [`AssetBundleManifest`](../ScriptReference/AssetBundleManifest.md)
* [Loading assets from AssetBundles](AssetBundles-Native.md)