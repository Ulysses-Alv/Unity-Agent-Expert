# Asset import determinism

Before you build assets, confirm whether the imported assets are deterministic.

You can use the following tools to verify this:

* **Use the consistency check argument**: Open the Unity Editor with the following launch argument, which enables the ConsistencyChecker to record results in the Editor log:

  ```
  -consistencyCheck
  ```
* **Check the log and reimport assets**: If the log contains a message such as the following, the asset import process is nondeterministic.:

  ```
  Importer(<ImporterName>) generated inconsistent result for asset (guid:<GUID>)
  ```

## Keeping import history

Use the [Import Activity window](ImportActivityWindow.md) to inspect re-import reasons and timestamps.

You can disable automatic artifact cleanup when analyzing imports as follows:

```
EditorUserSettings.artifactGarbageCollection = false
```

This prevents Unity from deleting previous import artifacts. Deleting the `Library` folder clears all historical import data.

## Non-deterministic asset types

Some Unity Editor-generated assets are non-deterministic and can produce different binary outputs when regenerated, even with identical source data and settings.

To maintain deterministic builds, avoid regenerating these assets unless the source content changes, and keep previously baked or generated data under **version control**. Only rebuild **lightmaps**, or other baked assets when the source content changes, and not as part of every build routine.

| **Asset type** | **Behavior** |
| --- | --- |
| **Lightmaps** | Re-baking lightmaps for every build produces different binary data, even with the same lighting setup. This is because the baking system involves not only floating-point rounding but also sampling noise that can make lightmaps non-reproducible. |
| ****NavMesh**** | Build NavMeshes under identical input and environmental conditions. However, NavMeshes are sensitive to floating-point operations because Unity calculates the NavMesh based on the **scene**’s meshes. Therefore, the NavMesh might bake differently on other architectures. |
| **Generated assets based on noise sampling** | Creating assets based on sampling noise textures or other randomness can generate assets differently each time. |

## Additional resources

* [Import Activity window](ImportActivityWindow.md)
* [Managing assets at runtime](assets-managing-runtime.md)
* [Command line arguments](CommandLineArguments.md)