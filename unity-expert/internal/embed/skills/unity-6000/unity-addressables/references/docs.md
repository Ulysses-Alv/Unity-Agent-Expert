# Unity 6000.3 LTS Addressables — Reference Documentation

Source: `reference/` (local documentation)

## Files Analyzed

| File | Content Summary | Lines |
|------|-----------------|-------|
| `com.unity.addressables.md` | Addressables package overview, version info, keywords | 32 |
| `AssetBundlesIntro.md` | AssetBundle concepts, structure, building, compatibility | 97 |
| `assets-managing-introduction.md` | Runtime asset management systems comparison | 94 |
| `LoadingResourcesatRuntime.md` | Resources system, performance impact, use cases | 39 |

## Detailed Index

### com.unity.addressables.md
- Addressable Asset System description
- Async loading from any location
- Automatic dependency management
- Access: `Window->Asset Management->Addressables`
- Package version 2.9.1 for Unity 6000.3
- Samples: github.com/Unity-Technologies/Addressables-Sample

### AssetBundlesIntro.md
- **What is an AssetBundle:** Archive file grouping assets for DLC, reduced install size, optimized platform-specific assets
- **Structure:** Binary header + serialized files + resource files (textures/audio split for multithreaded loading)
- **Scene AssetBundles:** Cannot mix scenes with other assets; accessible via `AssetBundle.isStreamedSceneAssetBundle`
- **Compatibility:** Forward incompatibility — newer Unity can't load older bundles without rebuild
- **Type Trees:** Used for safe binary deserialization across Unity versions; disable with `BuildAssetBundleOptions.DisableWriteTypeTree`
- **Building:** Addressables (high-level) or `BuildPipeline.BuildAssetBundles` (low-level)
- **Conditional compilation:** `#if` directives affect what gets serialized into bundles
- **Best practice:** Build all AssetBundles in single API call to auto-manage dependencies

### assets-managing-introduction.md
- **Direct references:** Default, entire asset file loaded at scene load, no dynamic loading, can't host remotely
- **Resources system:** `Resources` folder assets bundled into serialized file, slow startup with large amounts, no remote CDN support
- **AssetBundle system:** Separate container files, on-demand download, manual dependency tracking, custom reference counting needed
- **Addressables package:** Built on AssetBundle API, auto dependency/location/memory management, Editor UI
- **ECS content management:** Separate system in Entities package using weak references
- **Build inclusion:** Scene List + direct refs + Resources folder + AssetBundle assignments + Addressable group assignments

### LoadingResourcesatRuntime.md
- **Resources class:** Find/access objects in `Resources` folders
- **Use cases:** Small projects, bootstrap content (singletons, ScriptableObject config), content needed throughout app lifetime
- **Limitations:** Large impact on startup time, build size, memory management; no CDN support; requires player rebuild for changes
- **Performance impact:** Combined into single serialized file, memory proportional to object count even if not loaded
- **Warning:** 10,000 assets = several seconds init on low-end mobile even if rarely needed

## Key Concepts Cross-Reference

| Concept | Addressables | AssetBundle | Resources |
|---------|--------------|-------------|-----------|
| Async loading | Yes (built-in) | Manual | No |
| Auto dependencies | Yes | No | N/A |
| Remote hosting | Yes | Manual | No |
| Memory management | Auto | Manual | Manual |
| Editor interface | Yes | No | No |
| Build size impact | Minimal | Configurable | High |

## Memory Behavior

- **Addressables:** Release handle when done; Auto-instantiation handles cleanup
- **AssetBundle Unload(false):** Keep loaded assets, unload bundle container only
- **AssetBundle Unload(true):** Unload bundle AND all assets loaded from it
- **Resources.UnloadUnusedAssets:** Unloads all unused assets loaded by Resources
- **GC.GetTotalMemory:** Check memory pressure

## Important Warnings

1. AssetBundles are **platform-specific** — iOS bundles ≠ Android bundles
2. Forward compatibility **NOT supported** — must rebuild bundles for newer Unity
3. Resources folder **always included** in build even if unreferenced
4. AssetBundle.LoadFromFile blocks main thread if synchronous
5. Always pair load with unload to prevent memory leaks

## Related Documentation

- [Addressables package](https://docs.unity3d.com/Packages/com.unity.addressables@2.10/manual/index.html)
- [Addressables-Sample repository](https://github.com/Unity-Technologies/Addressables-Sample)
- [BuildAssetBundleOptions.AssetBundleStripUnityVersion](https://docs.unity3d.com/ScriptReference/BuildAssetBundleOptions.AssetBundleStripUnityVersion.html)
- [AssetBundle.LoadFromFileAsync](https://docs.unity3d.com/ScriptReference/AssetBundle.LoadFromFileAsync.html)