# Details panel reference (feature sets)

In the Package Manager window, when you select a [feature set](FeatureSets.md) from the list panel, the details panel on the right displays details for the selected feature set. The details panel presents the contents of the feature set as a kind of miniature Package Manager window:

![For a feature set, the details panel shows a brief description, a link to the QuickStart guide, and a list of included packages](../uploads/Main/fs-details.png)

For a feature set, the details panel shows a brief description, a link to the QuickStart guide, and a list of included packages

**(A)** When you select a feature set from the [list panel](upm-ui-list.md), its information appears in the details panel.

**(B)** The display name of the feature set.

**(C)** A link to the **QuickStart** guide for this feature set, containing details of how you can use this set of packages together.

**(D)** A button to **Install** or **Remove** the feature set.

**(E)** The **Details** tab, which contains a description of the feature set and its technical name (not pictured).

**(F)** The **Packages Included** tab, which displays the following information:

1. The list of included packages.
2. The details of the selected package. The information shown includes the display name of the package, the recommended or installed version for the feature set, and its description.
3. A shortcut to load the selected package in the Package Manager window. Selecting this shortcut replaces the feature set in the list panel and details panel. When you access details from the package directly, the Package Manager window provides more information than when you access them from inside the feature set. For example, the package details view shows dependency information and any samples the package has.

## Package version overrides

Feature sets are collections of packages that work well together for a specific version of Unity, which means the Package Manager installs specific package versions that your feature set requires. However, there are a couple of reasons why the Package Manager might actually install a different version (override the requested version):

* Another package or feature set requires a different version of the same package and the [Package Manager resolved](upm-conflicts.md) the package version by overriding it.
* You requested a different package version and it didn’t conflict with the version that the feature set requires. In this case, a **Reset** button displays, which you can click to return to the version that the feature set requires (recommended).
  + **Note**: The **Reset** button displays only when the major or minor number in the package version changes. The **Reset** button doesn’t display when the patch number in the package version changes. For more information on semantic version schemes, refer to [Versioning](upm-semver.md).

If the Package Manager installs a version other than the one you requested, an information message explains the reason for the change.

## Additional resources

* [Details panel reference](upm-ui-details.md)