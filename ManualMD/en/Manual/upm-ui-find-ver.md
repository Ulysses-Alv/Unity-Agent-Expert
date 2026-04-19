# Find a specific version of a package

To find a specific package version:

1. In Unity, open the **Package Manager** window. Refer to [Access the Package Manager window](upm-ui-access.md).
2. If you’re looking for a [pre-release package](pack-preview.md), follow these steps:

   * Select **Project Settings** under the [Advanced](upm-ui.md#Advanced) ![](../uploads/Main/iconSettings.png) menu.

     ![Project Settings opens the Package Manager project settings window](../uploads/Main/upm-settings.png)

     Project Settings opens the Package Manager project settings window

     The Package Manager [Project Settings](class-PackageManager.md) window appears.
   * Under the **Advanced Settings** group, select **Show Pre-release Package Versions**.
   * Close the **Project Settings** window. Any [pre-release packages](pack-preview.md) now appear in the list with the **Pre** label.
3. You can either browse the [list of packages](upm-ui-list.md) or find a specific package using one of these methods:

   * [Search for a package](upm-ui-search.md) by name or keyword. **Note:** The context you selected in the navigation panel limits the search scope. For example, if you’re searching for an **Asset Store** package, make sure you select the **My Assets** context before starting your search. When you switch contexts, the Package Manager window clears the search box. But if you return to the previous context, the Package Manager window restores your search term.
   * [Sort the list](upm-ui-sort.md) (for example, by **Recently updated**, to find the most recently updated packages).
   * [Apply filters](upm-ui-filter2.md) to the list.
4. Select a package from the [list of packages](upm-ui-list.md).
5. For packages from the Unity Registry, select the **Version History** tab in the [details panel](upm-ui-details.md). If multiple versions are available, select **See other versions** and expand the entries to display information specific to each version.

   ![Version numbers for installed packages display in multiple places and the recommended version appears in the Manage dropdown](../uploads/Main/upm-ui-update.png)

   Version numbers for installed packages display in multiple places and the recommended version appears in the **Manage** dropdown
6. When the **Version History** tab displays, you can perform a variety of actions when you expand a specific version:

   * You can view the changelog for that version and click **Changelog** to view the full list of changes across versions.
   * If you don’t already have this package installed in your project, you can [install](upm-ui-install.md) this version.
   * If another version of this package is already installed in your project, you can [update](upm-ui-update.md) the package to this version.
   * If you installed this package in your project but don’t want it anymore, you can [remove](upm-ui-remove.md) the package from your project.

## Additional resources

* [Find packages and feature sets](upm-ui-find.md)