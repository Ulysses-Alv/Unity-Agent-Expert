# Delete an asset package from the cache

When you download an asset package (`.unitypackage` file), the Unity Package Manager stores it in a cache for asset packages. Each time you import a downloaded package or a custom package, the Unity Package Manager stores it in the `Assets` directory in your project.

You might want to delete a package from the asset package cache to free up space on your local hard drive. If drive space is your main reason for deleting packages from the cache, consider [changing the asset package cache location](upm-config-cache-as.md).

**Important**:

* This procedure applies only to [asset packages](AssetPackages.md) that you [downloaded and imported](upm-ui-import.md) into your project. This procedure doesn’t apply to **UPM packages** that you installed from a registry. Some UPM packages also come from the **Asset Store**. There is no equivalent procedure for removing UPM packages from the global cache.
* Deleting an asset package from the asset package cache doesn’t remove the same package from any projects where you use that package. If you also want to remove the same package from a project, refer to [Remove imported assets from a project](upm-ui-remove-asset.md).

To delete a package from the asset package cache:

1. Open the **Package Manager** window.
2. Select the **My Assets** context from the [navigation panel](upm-ui-nav.md).
3. Select the asset package you want to delete from the [list panel](upm-ui-list.md).
4. Take note of two important values, which you need for a later step:
   * The publisher’s name, located below the display name in the [details panel](upm-ui-details.md) of the **Package Manager** window.
   * The **Display name** value, located in the ****Inspector**** window.

     ![The publisher name in the Package Manager window (left) and the Display name in the Inspector window (right)](../uploads/Main/upm-details-and-inspector.png)

     The publisher name in the Package Manager window (left) and the **Display name** in the Inspector window (right)
5. Go to the root of the asset package cache:
   * For Unity Editor 2022.1 or later:

     a) Open the Unity Editor’s [Preferences](Preferences.md) window.

     b) Select the **Package Manager** category.

     c) Click the folder icon beside **My Assets** > **Cache Location**.

     d) Select **Show in Explorer** (Windows) or **Reveal in Finder** (macOS).
   * For Unity Editor 2021.3 or earlier:

     a) Refer to [Location of downloaded asset package files](AssetStorePackages.md#asset-location) to determine the cache location for your operating system.

     b) Use your operating system’s file manager application to go to that directory.
6. Select the subdirectory that corresponds to the publisher’s name, which you recorded in an earlier step.

   ![Select the folder for the publisher name to find the .unitypackage file](../uploads/Main/upm-as-cache.png)

   Select the folder for the publisher name to find the .unitypackage file
7. The publisher’s directory has one or more subdirectories. The publisher determines the names of these subdirectories:
   * If there is one subdirectory, go into it.
   * If there is more than one subdirectory, explore the subdirectories until you find the one that contains the package you want to delete.
8. Select the `<name>.unitypackage` file, where `<name>` corresponds to the **Inspector** window’s **Display name** value that you identified in an earlier step.
9. Delete the file.

This package is now removed from the asset package cache, but any imported assets remain in any projects where they’re in use.

## Adding the deleted package back to the asset package cache

If the package that you deleted was an asset package, and you want to add this package back to your cache, refer to [Download and import an asset package](upm-ui-import.md).

## Additional resources

* [Package types](upm-package-types.md)
* [Asset Store packages](AssetStorePackages.md)
* [Remove imported assets from a project](upm-ui-remove-asset.md)
* [Download and import an asset package](upm-ui-import.md)