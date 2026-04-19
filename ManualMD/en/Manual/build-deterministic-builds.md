# Deterministic builds

Deterministic builds produce identical output when given the same source code, assets, and build settings. This consistency is essential for reliable deployment pipelines, debugging, and collaborative development.

In Unity, many factors can introduce non-deterministic behavior, which causes builds to vary.

| **Topic** | **Description** |
| --- | --- |
| **[Introduction to deterministic builds](build-deterministic-builds-introduction.md)** | Understand deterministic builds and common factors that can cause build variations. |
| **[Asset import determinism](build-deterministic-assets.md)** | Create consistent asset import results across different machines and build environments. |
| **[Editor scripts determinism](build-deterministic-editor-scripts.md)** | Create reliable build **scripts** that produce consistent results. |
| **[AssetBundle and Addressables determinism](build-deterministic-assetbundles-addressables.md)** | Configure AssetBundle and Addressables builds for deterministic output. |

## Additional resources

* [Customize the build pipeline](build-customize-build-pipeline.md)
* [Introduction to Unity Accelerator](UnityAccelerator.md)
* [Use AssetBundles to load assets at runtime](assetbundles-section.md)
* [Managing assets at runtime](assets-managing-runtime.md)
* [Import Activity window](ImportActivityWindow.md)