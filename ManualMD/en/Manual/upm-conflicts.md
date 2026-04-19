# Package version conflict and resolution

When you add a package to a **project manifest**, Unity considers that package a [dependency](upm-dependencies.md) of the project (a direct dependency). However, a package can also have dependencies on other packages, which create **indirect dependencies** in any projects that require that package.

Since most projects require more than one package in order to develop games and apps, the Package Manager has to evaluate all the requested package versions to retrieve from the registry (whether direct or indirect), and decide which among those package versions to install. To do this, it computes the set of packages that satisfies all [direct and indirect dependencies](upm-dependencies.md) in the project, starting with the project dependencies and recursively exploring each indirect dependency, collecting all the dependency information, then picking a set of packages that satisfies the dependency requirements without any conflict. For example, this dependency graph represents a project with four **direct dependencies** and all of their indirect dependencies:

![A graph of direct and indirect package dependencies for a project](../uploads/Main/upm-conflicts.svg)

A graph of direct and indirect package dependencies for a project

In this example:

* The light blue nodes represent the project’s direct dependencies.
* The dark blue nodes show the same package and version as indirect dependencies in this project.
* The red nodes show two different versions of the same package, which is a conflict.

**Note**: When resolving packages, the Package Manager selects versions using the following general priority:

1. Packages that you installed from other [sources](upm-concepts.md#Sources), such as [embedded packages](upm-embed.md), and dependencies declared with [local paths](upm-localpath.md), [Git URLs](upm-git.md), and [built-in](pack-build.md) packages over version-based dependencies.
2. Direct dependencies that you defined as [pinned packages](upm-manifestPrj.md#enableLockFile).
3. Direct dependencies that are declared with versions.

## Choosing the best solution

Depending on the set of packages defined in the project manifest, it could take a long time to evaluate all possible package combinations: a project could potentially depend on hundreds of packages, each of which depend on hundreds of other packages, most requiring different versions.

### Lock files and resolutionStrategy

To provide the most efficient solution, the Package Manager prioritizes package versions that it previously used by tracking them in a [lock file](upm-conflicts-auto.md). This guarantees that subsequent dependency resolution using the same inputs results in the same outputs. It also minimizes time-consuming operations such as downloading, extracting, or copying packages.

When the Package Manager can’t find a solution that uses only locked packages, it chooses the safest upgrades by selecting patch versions first, then minor versions, then major versions. However, you can change this behavior using the [resolutionStrategy](upm-manifestPrj.md#resolutionStrategy) property.

## Example

In this example, there are multiple versions of the following packages requested:

* `burst@1.2.2` (twice) and `burst@1.3.0-preview.3`
* `collections@0.5.1-preview.11` and `collections@0.5.2-preview.8`
* `jobs@0.2.4-preview.11` (twice) and `jobs@0.2.5-preview.20`

Using the set of [direct and indirect dependencies](upm-dependencies.md), the Package Manager selects the highest version of the burst package (`burst@1.3.0-preview.3`), which satisfies the `collections@0.5.2-preview.8` package’s dependency:

![In the dependency graph, the blue nodes indicate which versions the Package Manager selected](../uploads/Main/upm-resolution.svg)

In the dependency graph, the blue nodes indicate which versions the Package Manager selected

## Additional resources

* [Package dependency and resolution](upm-dependencies.md)