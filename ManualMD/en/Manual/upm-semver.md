# Package versioning

Package versions in Unity follow the Semantic Versioning convention. By using Semantic Versioning, you can communicate the scope of changes to your users whenever you alter assets, assemblies, APIs, or other elements that affect package compatibility.

**Note**: This page uses the following terminology:

* **Consumer** refers to any code or project that depends on (and uses) your package, including:
  + Unity projects that install your package.
  + Other packages that declare your package as a dependency.
  + Assembly definitions (`.asmdef` files) in a project that reference your package’s assemblies.
  + **Scripts** that reference your public types or assets by globally unique identifiers (GUID).
* **API surface** is everything that consumers rely on. This includes assemblies and their public APIs, any assets referenced by a GUID, and package manifest constraints.

The [Semantic Versioning](https://semver.org/) convention uses the `major.minor.patch` format:

* `major`: Incrementing this element represents a breaking change. Changes that remove or hide something are typically considered breaking changes.
* `minor`: Incrementing this element represents backward-compatible additions or changes.
* `patch`: Incrementing this element represents bug fixes only.

## Basic Semantic Versioning guidelines

Start new packages at version `0.1.0` while APIs are unstable. The major version number `0` indicates that packages are in the initial development phase. During this phase, package APIs change often, frequently in a breaking manner. Keep the major version number at `0` until you consider your package stable enough and ready for use in production.

When the package is ready for production, release version `1.0.0`, then:

* Increment the major version when you make API changes that break backward compatibility. This means code that worked with the previous version might not work with the new version without modifications.
  + **Example**: Versions `1.2.3` and `2.0.0` aren’t compatible and can’t be used interchangeably without risk.
  + **Note**: When you increment the major version, always reset the minor and patch values to `0`.
* Increment the minor value any time you introduce a functionality change in a backward-compatible way.
  + **Example**: You can use Version `1.3.0` to fulfill a dependency on `1.2.0`, because `1.3.0` is backward-compatible. However, you can’t use `1.2.0` to fulfill a dependency on `1.3.0`.
  + **Note**: When you increment the minor version, always reset the patch version to `0`.
* Increment the patch value any time you introduce a bug fix or any small backward-compatible change that doesn’t change the API.
  + **Example**: Versions `1.3.0` and `1.3.1` are interchangeable because they have the same API, even though `1.3.1` contains a bug fix not present in `1.3.0`.

Unity’s Package Manager relies on these versioning rules to resolve dependencies and select compatible package versions automatically.

## Change the version number

When you have decided which version number to increment, you can use either of the following methods to set the version number:

* Update the **Version** field in the ****Inspector**** window.
* Update the `version` property in the package manifest file.

For more information, refer to [Edit the package manifest](cus-edit-manifest.md).

## Summary of versioning rules

The following table summarizes the Semantic Versioning rules for incrementing parts of a version number. Later sections describe these changes in detail.

| **Area** | **Change type or context** | **Minimum version to increment** |
| --- | --- | --- |
| Assets | • Remove an asset visible to the Asset Database.  • Change an asset’s GUID. | Major |
| Assets | • Add assets or content that introduce new features. | Minor |
| Assets | • Adjust existing assets or content without changing their GUID. | Patch |
| Assemblies | • Remove or rename an assembly.  • Move public APIs between assemblies. | Major |
| Assemblies | • Remove platform support.  • Add a define constraint.  • Toggle the **Auto Referenced** setting (refer to [Automatic referencing](#automatic-referencing-in-assemblies).) | Major |
| Assemblies | • Add an assembly.  • Add platforms.  • Remove a define constraint. | Minor |
| Assemblies | • Change assembly references. | Patch |
| Assemblies | • Toggle **Allow ‘unsafe’ Code** or **Override References**. | Refer to [patch changes in assemblies](#no-api-changes-in-assemblies) |
| Package manifest | • Rename the package (`name` property). | Not supported |
| Package manifest | • Add or remove or change a dependency (internal only, no API exposure or behavior change). | Patch |
| Package manifest | • Dependency introduces a new behavior or exposes new types through your API. | Minor |
| Package manifest | • Dependency change alters behavior or API in a way that’s not backward compatible. | Major |
| Package manifest | • Increase the value in the `unity` or `unityRelease` property. | Minor or major |
| Scripting | • Rename or remove public types, methods, properties, fields, or events.  • Change signatures (parameter types or count or order), return types, or member visibility (public to internal or private).  • Rename namespaces or move public types or namespaces between assemblies.  • Change enumeration (enum) values or underlying types.  • Remove enumeration (enum) members.  • Change public field or property types used in Unity serialization (including fields shown in the **Inspector** window.)  • Remove or rename serialized members. | Major |
| Scripting | • Add new public types, members, or events.  • Add overloads or optional parameters to existing methods. | Minor |
| Scripting | • Internal or private refactors.  • Performance improvements.  • Bug fixes that don’t change the public API surface.  • Documentation, comments, or formatting. | Patch |
| Policy | • Mark APIs as obsolete. | Minor |
| Policy | • Remove obsolete APIs. | Major |

For more information about these rules, refer to the relevant details sections for [assets](#details-assets), [assemblies](#details-assemblies), [package manifest](#details-manifest), and [policies](#details-policy).

## Details for versioning changes involving assets

The Asset Database tracks assets by referencing the GUID in the assets’ `.meta` files.

The following breaking changes require an increment of the major version:

* Removing an asset that projects or packages might reference.
* Changing an asset’s GUID. If you change an asset’s GUID, the Asset Database understands this as removing the original asset and then adding a new (identical) asset. This results in a broken reference, because the original GUID no longer points to the asset, so the Asset Database can’t resolve the reference.

The following non-breaking changes require an increment of the minor version:

* Adding new assets as a backward-compatible feature.
* Updating asset content without changing GUIDs, if the addition is a feature.

The following non-breaking changes require an increment of the patch version:

* Updating asset content without changing GUIDs, if the addition is a bug fix.

## Details for versioning changes involving assemblies

Assembly definitions (`.asmdef`) define a group of scripts that the Unity Editor’s compilation pipeline uses to produce managed assemblies (`.dll` files). The following guidance also applies to precompiled assemblies.

### Automatic referencing in assemblies

The **Auto Referenced** property in the [Assembly Definition Inspector window](class-AssemblyDefinitionImporter.md#general) controls whether Unity implicitly references an assembly. You can also control this behavior using the `autoReferenced` property in the [assembly definition file](assembly-definition-file-format.md) (`.asmdef`) for your package.

Increment the major or minor version, depending on the scenario:

* Adding an assembly that has its **Auto Referenced** property enabled can cause compilation conflicts in consumers. In this scenario, increment the major version.
* Adding an assembly that has its **Auto Referenced** property disabled is backward compatible. In this scenario, increment the minor version.

Avoid bundling third‑party dynamic link library (DLL) files in unrelated packages, which can create unexpected implicit references.

### Breaking changes in assemblies

Increment the major version when you make the following breaking changes to an assembly:

* Remove an assembly definition or precompiled assembly.
* Rename an assembly (assembly name or DLL file).
* Move public APIs from one assembly to another.
* Remove platform support. Refer to [Assembly Definition Inspector window](class-AssemblyDefinitionImporter.md#platforms) (**Platforms** group) or [Assembly Definition file format reference](assembly-definition-file-format.md) (`includePlatforms` or `excludePlatforms` properties).
* Add a [define constraint](class-AssemblyDefinitionImporter.md#define-constraints) that can exclude the assembly from compilation. For information about managing constraints in a JSON format, refer to the `defineConstraints` property in [Assembly Definition file format reference](assembly-definition-file-format.md).
* Change the state of the **Auto Referenced** property in ways that alter availability to consumers. Refer to [Automatic referencing in assemblies](#automatic-referencing-in-assemblies).
* Mark an assembly as a [test assembly](test-framework/workflow-create-test-assembly.md). Unity doesn’t normally include test assemblies in builds (or compile them in some cases). When this happens, any assembly referencing the missing assembly fails to locate it, unless it’s also a test assembly.

### Backward compatible API changes in assemblies

Increment the minor version when you make non-breaking changes to an assembly. Treat the following changes as non-breaking changes only when you add or change the assembly while the [auto referenced property](#automatic-referencing-in-assemblies) is disabled:

* Add a new assembly (new `.asmdef` file with new scripts).
* Add platform support. Refer to [Assembly Definition Inspector window](class-AssemblyDefinitionImporter.md#platforms) (**Platforms** group) or [Assembly Definition file format reference](assembly-definition-file-format.md) (`includePlatforms` or `excludePlatforms` properties).
* Remove a [define constraint](class-AssemblyDefinitionImporter.md#define-constraints) so the assembly is always compiled. For information about managing constraints in a JSON format, refer to the `defineConstraints` property in [Assembly Definition file format reference](assembly-definition-file-format.md).
* Demote a test assembly to a regular assembly.

**Note**: If you make these assembly changes while the [auto referenced property](#automatic-referencing-in-assemblies) is enabled, then the changes become breaking changes. In such cases, you must increment the major version instead of the minor version.

### No API changes in assemblies

Increment the patch version when you make the following non-breaking changes to an assembly:

* Change the list of references in an assembly definition (`.asmdef`).
* Toggle **Allow ‘unsafe’ Code** or **Override References**. Refer to [Assembly Definition Inspector window](class-AssemblyDefinitionImporter.md#general) (**General Options** group) or [Assembly Definition file format reference](assembly-definition-file-format.md) (`allowUnsafeCode` or `overrideReferences` property).

These changes don’t alter the public API surface for consumers.

## Details for versioning changes involving the package manifest

The [package manifest](upm-manifestPkg.md) (`package.json`) defines a package’s identity, dependencies, and metadata.

### Name changes

Changing the `name` property isn’t supported. Publish the package as a new package with a different name.

### Dependency changes

Depending on the change, updating dependencies might require incrementing different parts of the version number.

| **Change** | **Change type or context** | **Minimum version** |
| --- | --- | --- |
| Add a dependency | • Changes to internal API only.  • Changes with no API exposure and no changes to behavior. | • Patch ([Auto Referenced](#automatic-referencing-in-assemblies) disabled)  • Major ([Auto Referenced](#automatic-referencing-in-assemblies) enabled) |
| Add a dependency | • Changes that introduce new behavior.  • Changes that expose new types in your API. | • Minor ([Auto Referenced](#automatic-referencing-in-assemblies) disabled)  • Major ([Auto Referenced](#automatic-referencing-in-assemblies) enabled) |
| Add a dependency | • Changes that alter existing behavior in a way that’s not backward-compatible.  • Changes to existing APIs. | Major |
| Remove a dependency | • Changes to internal API only without changing behavior. | Patch |
| Remove a dependency | • Changes that break behavior  • Removing APIs that expose their types. | Major |
| Change a dependency version | • Changes to internal API only without changing behavior. | • Patch ([Auto Referenced](#automatic-referencing-in-assemblies) disabled)  • Major ([Auto Referenced](#automatic-referencing-in-assemblies) enabled) |
| Change a dependency version | • Changes that introduce new behavior.  • Changes that expose new types in your API. | • Minor ([Auto Referenced](#automatic-referencing-in-assemblies) disabled)  • Major ([Auto Referenced](#automatic-referencing-in-assemblies) enabled) |
| Change a dependency version | • Changes that alter behavior or API that aren’t backward-compatible. | Major |

### Other manifest changes

You can change the `description`, `category`, `keywords`, and `displayName` properties in any release. Consider incrementing the minor version if the changes accompany new features.

Changing the `unity` or `unityRelease` properties requires at least minor increments because these changes alter supported Unity Editor versions. Increasing the Unity version excludes a package version from working on previous Unity editors and might break a dependent project or package. Decreasing the Unity version makes the package available to older Unity editors.

## Details for versioning changes involving policy

Policy changes include deprecation and removal of packages.

When you remove functionality from your package:

1. Increment the minor version and release the package, marking the affected APIs as obsolete. Include a message and guidance.
2. Increment the major version and release the package, removing those APIs.

## Additional information

* [Assembly Definitions](assembly-definition-files.md)
* [Package manifest](upm-manifestPkg.md)
* [Semantic Versioning](http://semver.org) (external website)