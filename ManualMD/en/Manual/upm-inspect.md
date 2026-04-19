# Inspect packages

The **Project window** displays the list of packages currently installed in your project from all sources. This list includes **immutable** packages that you installed from a package registry and **mutable** packages (such as embedded and local packages).

![Registry package (immutable) on the left and an embedded package (mutable) on the right](../uploads/Main/upm-inspect.png)

Registry package (immutable) on the left and an embedded package (mutable) on the right

You can inspect the contents of any package that appears in the **Project** window. You can also inspect the **package manifest** through a dedicated **Inspector**.

To inspect a package manifest, open the package’s `package.json` file from the **Project** window.

![Inspecting a package manifest](../uploads/Main/upm-project-view.png)

Inspecting a package manifest

For packages with a [custom or local label](cus-location.md), you can change the package contents and [edit the package manifest](class-PackageManifestImporter.md).

## Additional resources

* [Package manifest reference](upm-manifestPkg.md)