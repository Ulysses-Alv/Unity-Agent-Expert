# Editor script determinism

Scripts that modify assets during import or **post-processing** can unintentionally break build determinism. Before building, ensure that all generated assets are version-controlled and stable.

To prevent Unity Editor **scripts** from causing non-deterministic builds, follow these guidelines:

* Sort collected assets deterministically by GUID or asset path before assigning them to serialized fields.
* [`Directory.GetFiles`](https://learn.microsoft.com/en-us/dotnet/api/system.io.directory.getfiles?view=net-9.0) and [`AssetDatabase.FindAssets`](../ScriptReference/AssetDatabase.FindAssets.md) don’t guarantee sorted results.
* Avoid deleting and re-importing assets before builds, because Unity generates new GUIDs for re-imported assets.

Avoid non-deterministic coding patterns such as the following:

* Postprocessor scripts that mark assets dirty and re-save them during import.
* AssetPostprocessor scripts that use random seeds or machine-specific paths.
* Random or time-based values (`Random.value`, `Guid.NewGuid`, `DateTime.Now`, noise sampling) in asset generation.
* Creating new ScriptableObject assets at build time, which generate different GUIDs per machine.
* Modifying asset serialized data during the build process, because changes can be difficult to detect.

## Pre-serialized runtime data determinism

In some workflows, tools generate pre-serialized runtime data such as Protobuf message binaries, JSON blobs, or other binary payloads during the build process and embed them into the build as assets or resources.

Even if the generated code is deterministic, the serialized data might not be. Review the official documentation or source of the serialization library to verify whether it produces deterministic output under identical inputs.

Avoid regenerating serialized data during every build. Instead, generate once, validate, and commit the data to **version control**.

## Additional resources

* [`AssetDatabase` API reference](../ScriptReference/AssetDatabase.md)
* [`AssetPostprocessor` API reference](../ScriptReference/AssetPostprocessor.md)
* [ScriptableObject](class-ScriptableObject.md)
* [Customize the build pipeline](build-customize-build-pipeline.md)