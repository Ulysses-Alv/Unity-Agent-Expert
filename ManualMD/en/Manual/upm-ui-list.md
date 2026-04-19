# List panel reference

The Package Manager window displays the list of feature sets, packages, or **Asset Store** packages that meet your criteria.

You set the criteria by selecting a context in the [navigation panel](upm-ui-nav.md), and by optionally setting additional filter and search criteria.

For important information about the way search works, refer to [Search box](upm-ui-search.md).

![The image on the left displays features sets and packages installed in your project, and the image on the right displays packages from My Assets](../uploads/Main/upm-ui-list.png)

The image on the left displays features sets and packages installed in your project, and the image on the right displays packages from **My Assets**

**(A)** [Feature sets](FeatureSets.md) appear at the top of the list panel when you select the **Unity Registry** or **In Project** contexts. Feature sets are indicated with the feature set icon (![Icon showing layers of items](../uploads/Main/iconPkg.png)). Toggle the expander icon (![Right-pointing triangle](../uploads/Main/iconCollapsed.png)) to expand or collapse (![Down-pointing triangle](../uploads/Main/iconExpanded.png)) the list of feature sets. When you select a feature set from the list, [details about it appear on the right](fs-details.md). These details include a brief description, a list of included packages, and a link to the **QuickStart** guide for that feature set.

**(B)** Asset Store packages appear in the **In Project** list for all packages you added to your project from the **My Assets** list.

**(C)** Registry groups organize the list of packages installed in your project or available for installation. The list has separate groups for packages that come from Unity’s registry and other scoped registries you installed in your project. Select the expander on the left to expand (![Right-pointing triangle](../uploads/Main/iconCollapsed.png)) or collapse (![Down-pointing triangle](../uploads/Main/iconExpanded.png)) the list of packages for each registry.

**(D)** The installed or imported version of the package. If the package isn’t installed or imported yet, the listed version varies by package format:

* For **UPM packages**, the list view displays the recommended version. If a recommended version doesn’t exist, the list view displays the latest version.
* For **asset packages** (`.unitypackage` format), the list view displays the latest of version of the package, unless you downloaded the package but didn’t import it, in which case the downloaded version displays.

The following labels appear beside the version number for packages whose [state](upm-lifecycle.md) isn’t [Released](pack-safe.md):

* ![Pre-release label](../uploads/Main/iconPre.png) [Pre-release](pack-preview.md)
* ![Experimental label](../uploads/Main/iconExp.png) [Experimental](pack-exp.md)
* ![Custom label](../uploads/Main/iconCustom.png) [Custom](cus-location.md) (embedded)

For packages from the Asset Store, the version that appears is either the version you already downloaded or the version that’s available for download from the Asset Store.

**(E)** These icons display the status of the package:

| **Icon** | **Description** |
| --- | --- |
| Check mark | A check mark indicates that the package or feature set is already [installed](upm-ui-install.md), [enabled](upm-ui-disable.md), or [imported](upm-ui-import.md). |
| Upward pointing arrow | The update icon indicates that the package has an available update. To update your package, follow the appropriate instructions based on the package type:  - For Unity Package Manager (UPM) packages, refer to [Switch to another version of a UPM package](upm-ui-update.md).  - For asset packages, refer to [Update an asset package](upm-ui-update2.md). - You can’t update feature set versions because they’re fixed to the Editor version. |
| Lock icon | The lock icon indicates a package that’s installed as part of a feature set. |
| Link icon | The link icon indicates a package that’s installed as a dependency. |
| Folder with plus sign | The import icon indicates that there’s an asset package available to import. |
| Downward pointing arrow | The download icon indicates that there’s an asset package available to download. |
| Warning icon | A warning icon indicates an issue with the package, such as [lifecycle deprecation](pack-deprecated.md). |
| Stop icon | An error icon indicates [package version deprecation](pack-deprecated.md) or an issue that occurred during installation or loading. For information about resolving errors, refer to [Package Manager troubleshooting](upm-errors.md). |

**(F)** The **My Assets** context displays a counter showing the number of packages from the Asset Store that are available but not shown in the list. To load more packages from the Asset Store , click the **Load** link.

**Note**: If you select the **My Assets** context but the Package Manager window doesn’t list any packages, refer to [Error messages in the Package Manager window](upm-errors.md#Errors) for an explanation and solution.

**(G)** The [status bar](upm-ui.md#StatusBar) displays messages about the package load status and network warnings.

**(H)** Click the **reload** ![Reload button](../uploads/Main/iconReload.png) button to force the Package Manager to reload your packages and feature sets.

By default, the Package Manager window displays the list of all packages and feature sets available in the selected context, but you can [filter](upm-ui-find.md) the list to display packages and feature sets that meet your criteria.

You can also [include pre-release packages](class-PackageManager.md#advanced_preview) in the list and [search](upm-ui-search.md) for a specific package or feature set by [name (ID)](upm-manifestPkg.md#name) or [display name](upm-manifestPkg.md#displayName).

## Additional resources

* [Package types](upm-package-types.md)