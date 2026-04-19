# Remove local asset packages

When you import [local asset packages](AssetPackages.md), the Unity Editor places them in the `Assets` directory in your project.

You can manually remove assets from a project if you know the assets aren’t in use. You might consider this action to unclutter your project directory or to free up space on your local hard drive.

**Warning**: Make sure your project isn’t using any of the assets you remove. Unlike installed packages, the Package Manager doesn’t track assets that you imported from local **asset packages**, so it can’t detect when you remove dependent assets.

## Important: before you begin

* Only use this procedure to remove assets added with the [Import local asset packages](AssetPackagesImport.md) procedure.
  + To remove **UPM packages** installed from a registry, refer instead to [Remove a UPM package from a project](upm-ui-remove.md).
  + To remove asset packages downloaded and imported via the Package Manager window, refer instead to [Remove imported assets from a project](upm-ui-remove-asset.md).
* This procedure removes assets from the current project. It doesn’t remove the same assets that might exist in other projects.

## Procedure

To remove one or more local asset package items from your project:

1. Open the project you want to remove assets from.
2. Open the **Package Manager** window and [select a context](upm-ui-nav.md#contexts) from the [navigation panel](upm-ui-nav.md).
3. Select the package you want to remove from your project and take note of the package name **(A)** and the publisher name **(B)**. These names might help you identify the asset directory in a later step.

   ![The My Assets context shows the package name (A) and the publisher name (B)](../uploads/Main/upm-ui-myassets-details.png)

   The **My Assets** context shows the package name (A) and the publisher name (B)
4. Open the **Project window**.
5. Expand the `Assets` directory and locate the subdirectory for the package you identified in the **Package Manager** window in an earlier step.
6. Explore the directory structure that the publisher created **(C)**, confirm it’s the correct package, and identify the assets you want to delete **(D)**. **Note**: Unity doesn’t impose directory names or structures on publishers, so their assets might not be in an easily identifiable directory. The directory structure might be simple or complex.

   ![The Project window with an assets folder selected (C) and its contents (D)](../uploads/Main/upm-proj-assets.png)

   The Project window with an assets folder selected (C) and its contents (D)
7. Select the asset or assets you want to delete. Your selection can be a single asset, multiple assets, all assets in a subdirectory, or more.
8. Right-click the selected items, and select **Delete**.

**Important**: **Scenes** that depend on deleted assets aren’t reported as errors in the **Console window**. If you suspect the removal caused issues, import the package again. Refer to [Download and import an asset package](upm-ui-import.md).

## Additional resources

* [Remove a UPM package from a project](upm-ui-remove.md)
* [Remove imported assets from a project](upm-ui-remove-asset.md)