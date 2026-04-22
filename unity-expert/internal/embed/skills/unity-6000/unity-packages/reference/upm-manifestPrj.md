# Project manifest file

When Unity loads a project, the Unity Package Manager reads the **project manifest** so that it can compute a list of which packages to retrieve and load. When a user installs or uninstalls a package through the [Package Manager window](upm-ui.md), the Package Manager stores those changes in the project manifest file. The project manifest file manages the list of packages through the [dependencies](#dependencies) object.

In addition, the project manifest serves as a configuration file for the Package Manager, which uses the manifest to customize the registry URL and register custom registries.

You can find the project manifest file, called `manifest.json`, in the `Packages` folder under the root folder of your Unity project. Like the [package manifest file](upm-manifestPkg.md), the project manifest file uses JSON (JavaScript Object Notation) syntax.

## Properties

All properties are optional. However, if your project manifest file does not contain any values, the Package Manager window doesn’t load, and the Package Manager doesn’t load any packages.

| **Key** | **JSON Type** | **Description** |
| --- | --- | --- |
| **dependencies** | Object | Collection of packages required for your project. This includes only direct dependencies (indirect dependencies are listed in [package manifests](upm-manifestPkg.md)). Each entry maps the [package name](upm-manifestPkg.md#name) to the minimum [version](upm-semver.md) required for the project:   `{`   `"dependencies": {`     `"com.my-package": "2.3.1",`     `"com.my-other-package": "1.0.1-preview.1",`       `etc.`   `}` `}`  Specifying a version number indicates that you want the Package Manager to download the package from the package registry (that is, the [source](upm-concepts.md#Sources) of the package is the registry). However, in addition to using [versions](upm-manifestPkg.md#version), you can also specify a [path to a local folder or tarball file](upm-localpath.md), or a [Git URL](upm-git.md).   **Note**: You do not need to specify [embedded](upm-embed.md) packages here because the Package Manager finds them inside your project’s `Packages` folder and loads them automatically. The Package Manager ignores any entry if there is an embedded package with the same [name](upm-manifestPkg.md#name) in its own package manifest. |
| **enableLockFile** | Boolean | Enables a lock file to ensure that dependencies are resolved in a deterministic manner. This is set to `true` by default. For more information, see [Using lock files](upm-conflicts-auto.md). |
| **resolutionStrategy** | String | Upgrades [indirect dependencies](upm-dependencies.md) based on Semantic Versioning rules. This is set to `lowest` by default. For more information, see [Setting a resolution strategy](#strategize) below. |
| **scopedRegistries** | Array of Objects | Specify custom registries in addition to the default registry. This allows you to host your own packages.   Refer to [Scoped Registries](upm-scoped-use.md#scopes) for more details. |
| **testables** | Array of Strings | Lists the names of packages whose tests you want to load in the Unity [Test Framework](https://docs.unity3d.com/Packages/com.unity.test-framework@latest). For more information, see [Adding tests to a package](cus-tests.md).  **Note**: You do not need to specify [embedded](upm-embed.md) packages here because the Unity Test Framework assumes they are testable by default. |
| **pinnedPackages** | Array of Strings | Lists the names of packages that you want to lock at the version specified in the `dependencies` object. Items in this list must exist in the `dependencies` object, otherwise an error is generated.  Adding package names to this list makes Package Manager use the exact package version specified in the `dependencies` object, even if a different package version is more compatible with your version of the Unity Editor. |

## Example

```
{
  "scopedRegistries": [{
    "name": "My internal registry",
    "url": "https://my.internal.registry.com",
    "scopes": [
      "com.company"
    ]
  }],
  "dependencies": {
    "com.unity.package-1": "1.0.0",
    "com.unity.package-2": "2.0.0",
    "com.company.my-package": "3.0.0",
    "com.unity.my-local-package": "file:<path>/my_package_folder",
    "com.unity.my-local-tarball": "file:<path>/my_package_tarball.tgz",
    "com.unity.my-git-package": "https://my.repository/my-package.git#v1.2.3"
  },
  "enableLockFile": true,
  "resolutionStrategy": "highestMinor",
  "testables": [ "com.unity.package-1", "com.unity.package-2" ],
  "pinnedPackages": ["com.unity.package-1"]
}
```

## Setting a resolution strategy

While you can force Unity’s package dependency resolution to use higher versions of indirect dependencies by adding them explicitly to the project manifest, this isn’t a good strategy, for two reasons:

* It puts more responsibility on the project owner to maintain dependency versions.
* Over time, you might have dependencies that are not required by the project.

A better approach is to customize how the Package Manager selects indirect dependencies based on Semantic Versioning rules by setting the **resolutionStrategy** property:

| **Value:** | **Description:** |
| --- | --- |
| **lowest** | Do not upgrade indirect dependencies. Instead, it uses exactly the requested version. This is the default mode. |
| **highestPatch** | Upgrade to the highest version with the same Major and Minor components. For example, with the requested version 1.2.3, this strategy selects the highest version in the range `[1.2.3, 1.3.0)` (that is, `>= 1.2.3` and `< 1.3.0`). |
| **highestMinor** | Upgrade to the highest version with the same Major component. For example, with the requested version 1.2.3, this strategy selects the highest version in the range `[1.2.3, 2.0.0)` (that is, `>= 1.2.3` and `< 2.0.0`).  **Note**: Version `1.0.0` marks the first stable, production-ready version. Below that, versions `0.X.Y` indicate that their API is not yet stable, and successive minor versions might introduce breaking changes. This part of the SemVer specification allows releasing early versions of a package without hampering rapid development. Because of this, when the target version is `0.X.Y`, **highestMinor** behaves like **highestPatch** in order to ensure choosing a backward-compatible version. For example, with the requested version `0.1.3`, this strategy selects the highest version in the range `[0.1.3,0.2.0)`. |
| **highest** | Upgrade to the highest version. For example, with the requested version 1.2.3, this strategy selects the highest version in the range `[1.2.3,)` (that is, `>= 1.2.3` with no upper bound) |

**Note**: These ranges never allow a dependency to jump from a stable release to an [experimental](pack-exp.md) or [pre-release](pack-preview.md) package.

## Additional resources

* [Package manifest file](upm-manifestPkg.md)