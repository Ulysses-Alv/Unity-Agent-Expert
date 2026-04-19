# Navigation panel reference

Use the Package Manager’s navigation panel to select which subset of packages you want to view.

![The navigation panel (A), the list panel (B), and the details panel (C)](../uploads/Main/upm-ui-nav.png)

The navigation panel **(A)**, the list panel **(B)**, and the details panel **(C)**

Select a context from the navigation panel **(A)** to list those types of packages in the list panel **(B)**. When you select a package from the list panel, its details display in the details view **(C)**.

## Context list

Select the context you want from the navigation panel. You can select from these options:

| **Context** | **Description** |
| --- | --- |
| **In Project** | Displays all feature sets and packages installed in your project, including [local](upm-localpath.md), [git](upm-git.md), and [embedded](upm-embed.md) packages, and packages installed from any registry. The list also includes any packages from the Asset Store that you added from the **My Assets** context.  **In Project** has nested entries that appear when the proper conditions are met: - **Updates** lists all packages in your project that you can update, including Asset Store packages. - **Errors and Warnings** lists all packages in your project that require attention. |
| **Unity Registry** | Displays all **feature sets** and packages on the [Unity package registry](upm-concepts.md#Registry), regardless of whether they’re already installed in your project. This doesn’t include packages installed from any other location or from any [scoped registry](upm-scoped.md). |
| **My Assets** | Displays all [Asset Store packages](AssetStorePackages.md) that you have purchased using the Unity ID you’re currently logged in with. |
| **Built-in** | Displays only built-in Unity packages, which represent core Unity features. You can use these packages to [turn Unity modules on and off](upm-ui-disable.md).  **Tip**: You can find out more about what each built-in package (module) implements in the [Unity Scripting API](upm-api.md). Each module assembly page lists which APIs the built-in package implements. |
| **Services** | Displays integrated [services](UnityServices.md) to help you to engage, retain, and monetize audiences. |
| **My Registries** | Displays any packages available from any [scoped registry](upm-scoped.md) installed in this project.  This context appears only if you installed a scoped registry in this project.  If you installed a scoped registry but it’s not listed in the **My Registries** context or the **My Registries** context isn’t available at all, it might be because the package registry server you added doesn’t implement either of the `/-/v1/search` or `/-/all` endpoints, which means that it’s not compatible with Unity’s Package Manager. |

**Note**: If you entered text in the [search box](upm-ui-search.md) or set [filters](upm-ui-filter2.md), the list panel displays only feature sets and packages which match the context, search criteria, and active filters.

## Additional resources

* [List panel](upm-ui-list.md)
* [Details panel](upm-ui-details.md)