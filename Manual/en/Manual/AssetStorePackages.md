# Manage Asset Store packages in the Editor

Asset Store packages are collections of files and data from Unity projects, or elements of projects. An **Asset Store** package type is either [a UPM package](asset-store-upm.md) or an **asset package** (`.unitypackage` format). When you add an Asset Store package to your project, the [Unity Package Manager](upm-ui.md) unpacks the package and maintains its directory structure and metadata about assets. This metadata includes information such as import settings and links to other assets.

## Using Asset Store packages

To use an Asset Store package, perform the following steps:

1. Open the [Unity Package Manager window](upm-ui.md)
2. Select the [My Assets](upm-ui-nav.md#contexts) context to view the [list of available Asset Store packages](upm-ui-list.md). You can also [search by name](upm-ui-search.md) for Asset Store packages.
3. The method to add the package to your project depends on the type you’ve downloaded:
   * For asset packages (`.unitypackage`), refer to [Download and import an asset package](upm-ui-import.md).
   * For **UPM packages**, refer to [Install a UPM package from Asset Store](upm-ui-install2.md).

If the Asset Store package has a newer version available, you can also update it directly in the Package Manager window. To update an asset package (`.unitypackage`), refer to [Updating an asset package](upm-ui-update2.md), and to update a UPM package, refer to [Switch to another version of a UPM package](upm-ui-update.md).

## Location of downloaded asset package files

**Note**: This information applies only to asset packages (`.unitypackage`) you get from the Asset Store. It doesn’t apply to UPM packages you get from the Asset Store.

You can use the Package Manager window to manage your Asset Store packages within your project. However, if you need to access files from an asset package directly, you can find them in the cache for asset packages. This cache, which is separate from the [global cache](upm-cache.md), makes reusing and sharing asset packages more efficient. You can find the cache for asset packages in the following paths (which might be within hidden folders depending on your computer settings):

* macOS: `~/Library/Unity/Asset Store-5.x`
* Windows: `C:\Users\accountName\AppData\Roaming\Unity\Asset Store-5.x`
* Linux: `~/.local/share/unity3d/Asset Store-5.x`

These folders contain subfolders that correspond to particular Asset Store vendors and are available after download and import. The Package Manager stores asset files inside the subfolders using a structure defined by the Asset Store package publisher.

You can override the default location of the cache for asset packages. For information, refer to [Customize the asset package cache location](upm-config-cache-as.md).

## Additional resources

* [Package types](upm-package-types.md)
* [Manage packages in the Asset Store](AssetPackagesOrganize.md)
* [Publishing Asset Store packages](asset-store-publishing.md)