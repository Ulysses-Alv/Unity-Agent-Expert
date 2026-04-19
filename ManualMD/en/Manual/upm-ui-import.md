# Download and import an asset package

You can import an asset package that you got from the **Asset Store** and download it to your project by using the **Package Manager** window.

**Notes:**

* This task applies only to asset packages in the `.unitypackage` format. For information about installing packages from the Asset Store that are in the UPM format, refer to [Install a UPM package from Asset Store](upm-ui-install2.md). For information about the differences between the package types, refer to [Package types](upm-package-types.md).
* Follow this task to import an asset package that isn’t in your project yet. For information about updating an asset package that’s already in your project, refer to [Update an asset package](upm-ui-update2.md).

## Prerequisites

Before you begin:

* Make sure you’ve acquired the package from the Asset Store. Refer to [Purchase or download a package from the Asset Store](AssetPackagesPurchase.md).
* Make sure the package appears in the [List view](upm-ui-list.md) of the Package Manager window when you select the **My Assets** context.

## Procedure

1. Open the **Package Manager** window and select **My Assets** from the [navigation panel](upm-ui-nav.md).

   ![Switch the context to My Assets](../uploads/Main/upm-ui-myassets.png)

   Switch the context to **My Assets**

   The list panel displays any packages that you acquired from Asset Store using the Unity ID you’re logged in with.
2. If many asset packages appear in the **My Assets** list, you can [search for a specific package](upm-ui-search.md) by name.
3. Select the asset package you want to import from the [list of packages](upm-ui-list.md). The asset package information appears in the [details panel](upm-ui-details.md).

   If you haven’t downloaded the asset package to this computer or device before, the **Download** button appears. Otherwise, the **Import #.# to project** button appears instead and you can skip to step 5 to import the package right away.
4. Click the **Download** button to download the asset package if it appears in the details panel.

   ![Download button in the corner of the Details panel](../uploads/Main/upm-ui-import.png)

   **Download** button in the corner of the Details panel

   If the download button appears as **Download update #.#\***, the asterisk (\*) means that you have an older version of that package with the same version number in your package cache. The publisher has subsequently published an updated version of the package without changing its version number. Click that download button to get the latest version of the package.
     
     
   While the package downloads to your computer, a progress bar appears and **Pause** and **Cancel** buttons replace the **Download** button you clicked.

   ![Download progress, and buttons to control the progress](../uploads/Main/upm-ui-download.png)

   Download progress, and buttons to control the progress

   These buttons give you control over the download progress:

   * Click the **Pause** button to pause downloading. You can click the **Play** button to continue.
   * Click the **Cancel** button to stop the download entirely.

   When the download finishes, the Package Manager replaces **Download** with a menu containing actions to **Import #.# to project** and **Re-download #.#**.
5. Click **Import #.# to project** to import the selected asset package.

   If the asset package is a complete project, a confirmation dialog appears. Refer to [Importing a complete project from the Asset Store](#complete-proj).

   For all other asset packages, an import window displays, with all items selected and ready to import.

   ![The Import window](../uploads/Main/upm-ui-import-assets.png)

   The **Import** window
6. Clear any items you don’t want to import and click **Import**.

   The Package Manager puts the contents of the imported package into the `Assets` folder, so that you can access them from your **Project** window.

When the import operation completes, the imported package remains in the **My Assets** list, but you can also view it in the **In Project** list. Whenever an update is available for the package, you can view it in the same lists, plus the **Updates** list, which is nested below **In Project**.

**Tip**: You can download multiple packages with one click by using the multiple select feature. For more information, refer to [Perform an action on multiple packages](upm-ui-multi.md).

## Viewing imported assets

After you import an asset package, or a subset of assets from an asset package, you can view a list of the imported assets.

To view imported assets:

1. Select **My Assets** in the [navigation panel](upm-ui-nav.md).
2. Select the package in the [list panel](upm-ui-list.md).
3. Select the **Imported Assets** tab in the [details panel](upm-ui-details.md).

## Importing a complete project from the Asset Store

A complete project has assets and **project settings**. The process for importing a complete project is similar to other asset packages, except there’s an additional step for project settings.

The project settings in the complete project might conflict with settings in your project, so when you import a complete project from the Asset Store, the Package Manager opens a dialog:

![Dialog for importing a complete project](../uploads/Main/upm-ui-import-project.png)

Dialog for importing a complete project

Choose **Import** if you’re sure that you want to add the package’s assets to your project. An **Import Unity Package** window appears, like the one pictured below. Clear any asset content you don’t want to import and click **Next**. The second step of the window lists the package’s project settings. If you have any of those project settings in your project, Package Manager identifies them with an **Override** flag. Those settings will be overwritten with the project settings from the package unless you clear the selection. Clear any project settings you don’t want to import and click **Import**.

Choose **Switch Project** if you want to explore the package’s assets and project settings in a safe environment. Choosing **Switch Project** creates a new, temporary project that has only that package’s assets and project settings. After exploring the temporary project, you can close it without saving. If you want to continue to use that package in your main project, return to your main project and import the package again. Choose **Import** on the warning dialog that appears.

![Importing asset content (left) and project settings (right) from a complete package](../uploads/Main/upm-ui-import-assets2.png)

Importing asset content (left) and project settings (right) from a complete package

## Additional resources

* [Asset packages](AssetPackages.md)
* [Package types](upm-package-types.md)
* [Install a UPM package from Asset Store](upm-ui-install2.md)
* [Install a UPM package from a registry](upm-ui-install.md)