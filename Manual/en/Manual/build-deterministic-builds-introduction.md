# Deterministic builds introduction

Build determinism is important if you want to be able to repeat a build process and get the same set of binary files. Creating a deterministic build enables regulatory compliance for industries requiring binary verification, and is a useful way to ensure that users don’t download unchanged content. Deterministic builds make unauthorized modifications easier to detect, and you can also track any bugs to their exact code and asset states.

## Understanding non-deterministic builds

Non-deterministic builds can happen in the following situations:

* **During asset import and conversion:** Non-deterministic builds can happen when Unity imports source assets and converts them to an [internal format](assets-supported-types.md#native-asset-importers), or when Unity converts imported assets into packaged assets such as [AssetBundles](AssetBundlesIntro.md).
* **Editor **scripts** that collect or modify assets at build time:** Non-deterministic builds can happen if, for example, tools or custom postprocessors automatically gather project assets and assign them to serialized array fields inside a ScriptableObject. If the order of objects in that array isn’t explicitly sorted, the serialized data might differ between machines or consecutive builds.

## Build environment consistency

The most common cause of non-deterministic builds are mismatched environments. To ensure a consistent environment consider the following:

* **Consistent build machine environments:** Keep operating system and CPU architecture consistent across all build machines. Floating-point rounding behavior differs between architectures, which can change serialized float values especially in AnimationClips or imported **Mesh** data.
* **Consistent toolchains:** Use the same Unity Editor version, Xcode version, and toolchains across builds. Small serialization changes between Unity versions can change AssetBundle byte layouts.
* **Disable automatic new line conversion in Git:** Windows and macOS encode line endings differently, and automatic line conversion modifies file hashes and breaks determinism. To disable automatic conversion, run the following git command:

  ```
  git config --global core.autocrlf false
  ```

### Use a cache server to preserve determinism

To achieve deterministic builds in Unity, the most effective method is to force all machines to share identical imported assets via [Unity Accelerator](UnityAccelerator.md).

When multiple developers connect to the same remote **Accelerator**, each machine reuses the same imported artifacts instead of re-importing locally.

Unity Accelerator has the following limitations:

* It only caches imported assets.
* It doesn’t cache built AssetBundles or Addressable build artifacts.
* The determinism benefit applies to import-time consistency, not AssetBundle reuse.

## Additional resources

* [Introduction to Unity Accelerator](UnityAccelerator.md)
* [Managing assets at runtime](assets-managing-runtime.md)
* [AssetBundles introduction](AssetBundlesIntro.md)
* [Customize the build pipeline](build-customize-build-pipeline.md)