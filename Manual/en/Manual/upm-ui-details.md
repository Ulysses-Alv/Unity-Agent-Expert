# Details panel reference

The right panel of the Package Manager window displays details for the selected package.

![Package details for a Unity Package Manager (UPM) package (left) and an asset package (right)](../uploads/Main/upm-ui-details.png)

Package details for a Unity Package Manager (UPM) package (left) and an asset package (right)

These details include the following information:

**(A)** The display name.

**(B)** The [package version](upm-ui-find-ver.md), the date the package was published to the registry or the Asset Store, and any [source or asset labels](upm-ui-details.md#Tags). Other information might display, if applicable:

* An information tooltip (![Information tooltip](../uploads/Main/iconInfo.png)) might also appear before the version. You can click the tooltip for more information about the package. This tooltip can appear, for example, if the package version [you requested doesn’t match](upm-dependencies.md#overridden-version) the version installed.
* For Unity Package Manager (UPM) packages, a lock icon (![Lock icon](../uploads/Main/iconLock.png)) appears before the version number when an installed **feature set** requires the selected package. Locking the package prevents you from accidentally [changing the version](upm-ui-update.md) of the package so the feature set continues to work effectively.

**(C)** The registry that hosts the package and the package author (UPM packages), or the package publisher (Asset Store).

**(D)** Links to additional resources:

* For **UPM packages**: links to [open the package documentation](upm-docs.md) page, the package changelog (if available), and the license information.
* For **asset packages**: a link to open the package’s official page on the Asset Store, links to the publisher’s support page and website (if available), and a link to review the package on the Asset Store.

**(E)** A [locate](cus-edit-manifest.md#locate-manifest) button to show the package’s manifest file in the **Project** window, and a **Manage** dropdown. The entries in the **Manage** dropdown are dynamic depending on various conditions, but include actions for managing a UPM package, such as [unlock](upm-ui-update.md#fs-lock), [install](upm-ui-install.md), [update](upm-ui-update.md), [remove](upm-ui-remove.md), or [disable](upm-ui-disable.md), and more.

**(F)** Controls to [download and import](upm-ui-import.md), [remove](upm-ui-remove-asset.md), and [update](upm-ui-update2.md) an asset package. Sometimes these controls are buttons. Other times, the control is a menu with the most common action displayed first. The **Package Manager** window displays **In Project** (and a check mark) when your project and cache are up-to-date with the latest version of an asset package. Specifically, it means at least one item from that package is at the latest level. However, you can still select other actions from the menu when **In Project** is displayed.

**(G)** Additional information tabs, which vary by package type:

* For UPM packages, including UPM packages from Asset Store:
  + **Details**: Details such as package name and Editor compatibility, and a description of the package.
  + **Version History**: Package Manager always lists the installed or recommended package versions, with action buttons to [install](upm-ui-install.md), [update](upm-ui-update.md), or [remove](upm-ui-remove.md). If other supported versions are available and compatible with your version of the Unity Editor, the Package Manager lists those versions as well. Expand each entry to check the release date, changelog for that version, and a link to the full changelog.
  + **Dependencies**: Dependency information for the package. This section lists dependencies in two directions:
    - **Is using**: Whether this package depends on another package, and if so, which version. Installed packages display the status of the dependencies after the version number.
    - **Used by**: Whether another package depends on this one, and if so, which version.
  + **Samples**: This tab displays only if the package author provided samples. The tab displays a description and an **Import** button next to the sample.
  + **Images** (UPM packages from Asset Store only): Thumbnails of the marketing images.
* For asset packages:
  + **Overview**: Information about the package:
    - Any [custom labels](AssetPackagesLabels) you assigned to this package on the Asset Store.
    - The supported versions of Unity.
    - The disk space required.
    - The date you acquired this package from the Asset Store.
    - A brief description of the package.
  + **Releases**: Release information, including the release date of the original package and the release date of the current package.
  + **Imported Assets**: Lists the items from the asset package that you imported into your project. This tab appears only after you download a package and import it into your project.
  + **Images**: Thumbnails of the marketing images, audio, and video available on the Asset Store.

**(H)** The technical name for the package (UPM packages only), which maps to the [name](upm-manifestPkg.md#required) field in the **package manifest** file.

## Labels

Some packages display labels next to the package name or version number. These labels offer information about the source or state of the package:

* [Source](upm-concepts.md#Sources) label types indicate where the package originates from (for example, whether it comes from a local folder or if you downloaded it from a package registry).
* [State](upm-lifecycle.md) label types indicate the package’s stage of the development cycle. For example, whether it’s an experimental package or on the release track for a specific Unity version. If the Package Manager window displays a package without any state-based labels, it means the package is in the Released state.
* [Asset Store](AssetStore.md) labels indicate that the package originated from the Asset Store. These labels might also indicate a special status (for example, deprecated).

Some source labels imply state labels and vice versa. For example, if you embedded a package in your project, then Unity automatically assumes it’s a custom package in development, so only the **custom** label appears in the [details panel](upm-ui-details.md).

The Package Manager window displays the following values:

| **Label** | **Type** | **Meaning** |
| --- | --- | --- |
| [Pre-Release](pack-preview.md) | state | This version of the package is at an earlier stage of development, but Unity guarantees to release it by the end of the Long Term Support (LTS) release cycle. It might not yet have complete documentation or be fully validated by either the development team or Unity’s Quality Assurance team. |
| [Experimental](pack-exp.md) | state | These packages are either new packages or contain experimental modifications. Unity doesn’t support Experimental packages because they’re in the early stages of development. |
| [Deprecated](pack-deprecated.md) | state | These packages are no longer under development for that Unity Editor version, or a specific version of a package has a critical issue. Unity doesn’t support Deprecated packages. |
| [Custom](cus-location.md) | state, source | This package is [embedded](cus-location.md) in the `Packages` folder of your project. Most custom package developers start by embedding a new package in their project, which is why the **Custom** label appears. |
| [Local](cus-location.md) | source | The Package Manager installed this package from a folder or tarball file on your local disk outside the `Packages` folder of your project. |
| [git](upm-git.md) | source | The Package Manager installed this package in your project directly from a Git repository. |
| [asset store](AssetPackagesPurchase.md) | asset | This is an asset package that you purchased or downloaded from the Asset Store. |
| [deprecated](AssetStoreAdmin#pkg-status) | asset | This package is no longer available on the Asset Store, unless you have downloaded or purchased it before. In other words, it’s no longer discoverable by new customers. |

## Additional resources

* [Package types](upm-package-types.md)
* [Details panel reference (feature sets)](fs-details.md)