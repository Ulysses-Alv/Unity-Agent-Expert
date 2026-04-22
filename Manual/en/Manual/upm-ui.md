# The Package Manager window

Access the Package Manager window from the Unity Editor’s **Window** > **Package Management** menu.

Use the Package Manager window to:

* View which [packages and feature sets](PackagesList.md) are available for installation or already installed in your project.
* Check [which package versions are available](upm-ui-find-ver.md).
* Install, update, or remove [UPM packages](upm-ui-actions.md) or [feature sets](fs-install.md).
* Download and import, update, or remove [asset packages](upm-ui-actions-ap.md).
* [Disable built-in packages](upm-ui-disable.md).

![The Package Manager window](../uploads/Main/upm-ui.png)

The Package Manager window

The Package Manager window displays:

**(A)** The experimental package indicator, which warns you if your project has [experimental packages](pack-exp.md).

**(B)** The **install** ![The install button](../uploads/Main/iconAdd.png) button, which you can click to [install a package](upm-ui-actions.md) directly into your project by entering a git URL, a local path, or a package name.

**(C)** The [navigation panel](upm-ui-nav.md), which you can use to select a context to change what appears in the list panel **(H)**.

**(D)** The [Sort](upm-ui-sort.md) menu, which you can use to sort the list of packages and feature sets by name or date.

**(E)** The [Filter](upm-ui-filter2.md) menu, which you can use to narrow down which packages appear in the list panel **(H)**. The **Filters** menu and the **Clear Filters** button are disabled for the **Built-in** list. They’re also disabled for the **In Project** context (unless you have subscription-based packages), because that context in the navigation panel has a nested item for **Updates**.

![The Filters menu and the Clear Filters button](../uploads/Main/upm-ui-asset-filters.png)

The **Filters** menu and the **Clear Filters** button

**(F)** The [search box](upm-ui-search.md), which you can use to look for packages and feature sets by name.

**(G)** The **Advanced** menu ![The Advanced menu](../uploads/Main/iconSettings.png), which you can use to access the **project settings** for the Package Manager, preferences, and more. Refer to [Advanced settings](#Advanced) for details.

**(H)** The [list panel](upm-ui-list.md), which displays packages for the type you selected in the navigation panel, limited by any filter and search parameters you specified.

**(I)** The [details panel](upm-ui-details.md), which displays information specific to the [package](upm-ui-details.md) or [feature set](fs-details.md) selected in the list panel.

**(J)** Buttons to perform any of the following actions at the project level:

* [Install](upm-ui-install.md) a UPM package.
* [Locate](cus-edit-manifest.md#locate-manifest) the **package manifest** file in the **Project** window.
* A **Manage** dropdown, which has entries for managing a UPM package, such as:
  + [Update](upm-ui-update.md) or [remove](upm-ui-remove.md) UPM packages.
  + [Customize](upm-embed.md#embed-cached) a UPM package, which copies the package to your project folder.
  + [Open the package manifest](cus-edit-manifest.md#open-manifest) for **immutable** packages.
  + [Edit the package manifest](cus-edit-manifest.md#edit-manifest) for **mutable** packages.
* [Download and import](upm-ui-import.md), [update](upm-ui-update2.md), or [remove](upm-ui-remove-asset.md) asset packages.
* [Install or remove feature sets](fs-install.md).
* [Disable or enable](upm-ui-disable.md) **built-in packages**.
* Install or remove [services](UnityServices.md).

**(K)** The package details tabs, which display further information about the selected package or feature set. The tabs are dynamic, based on the selected item. For information about these tabs, refer to [Details panel](upm-ui-details.md).

**(L)** The [status bar](#StatusBar), which displays information when the Package Manager loads packages and feature sets. This information includes errors and warning messages, the number of **Asset Store** packages available, and a link to load more packages from the Asset Store.

**(M)** The **Refresh list** ![The Refresh list button](../uploads/Main/iconReload.png) button lets you refresh the list of packages displayed. In the **My Assets** context, **Refresh list** is a menu, which has a **Check for updates** option. You can use **Check for updates** to check for updates to all packages on your computer, not just the ones that are visible in the **My Assets** context.

## Advanced settings

The advanced settings ![The advanced settings menu](../uploads/Main/iconSettings.png) menu allows you to perform these actions:

| **Menu item** | **Action results** |
| --- | --- |
| **Project Settings** | Select this item to open the [Package Manager project settings](class-PackageManager.md), where you can:  - List [pre-release packages](pack-preview.md) when browsing the Unity Registry. - Add, edit, and remove [scoped registries](upm-scoped.md) in your project. |
| **Preferences** | Select this item to view and set [Preferences](Preferences.md) for the Unity Editor and related windows and tools. |
| **Manual resolve** | Select this item to force the Package Manager to resolve the project’s packages. If needed, it re-installs altered or missing packages and removes extraneous packages. |

## Status bar

The Package Manager displays messages in the status bar at the bottom of the Package Manager window.

There are typically four status messages that the Package Manager might display:

* The first time you open the Package Manager window in a new project, the **Refreshing list** message appears briefly:

  ![Message for refreshing packages and features](../uploads/Main/upm-ui-loading.png)

  Message for refreshing packages and features

  This message also appears when you click **Refresh list** ![The Refresh list button](../uploads/Main/iconReload.png)
* When you select the **My Assets** context in the navigation panel, the load bar appears above the date. It displays the number of packages from the Asset Store. Select **Load** to load more packages.

  ![On the left, the load bar displays the number of My Assets packages loaded and the total number available.](../uploads/Main/upm-ui-assets-loadbar.png)

  On the left, the load bar displays the number of **My Assets** packages loaded and the total number available.
* Most of the time, the status bar displays the date and time of when the Package Manager window last refreshed its information. However, if the Package Manager [detects a problem](https://docs.unity3d.com/Manual/upm-errors.html), such as a network issue, the Package Manager prompts you to sign in:

  ![Network error message](../uploads/Main/upm-ui-errors.png)

  Network error message
* If your network connection is working, but you’re not signed into your [Unity account](https://id.unity.com/), the Package Manager doesn’t display any packages from the Asset Store. When you select **My Assets** in the navigation panel, the Package Manager prompts you to sign in:

  ![Logged out of Unity account](../uploads/Main/upm-ui-unityid.png)

  Logged out of Unity account

  In the list panel, click **Sign in** to sign in to your Unity account through the [Unity Hub](unity-hub.md).

For information on how to resolve these errors and more, refer to [Package Manager troubleshooting](upm-errors.md).

## Additional resources

* [Get started with packages](Packages.md)
* [Packages and package management](PackagesList.md)
* [Finding package documentation](upm-docs.md)