# Switch to another version of a UPM package

Use the information on this page to update **UPM packages** that you installed from the [Unity Registry](upm-ui-install.md) or a [scoped registry](upm-scoped.md), or from a [local source](upm-ui-local.md). The information on this page also applies to packages you [installed from the Asset Store](upm-ui-install2.md), if they’re in the UPM format.

You can also use the information on this page to switch to a registry version of a package, if your project currently uses a nonregistry version of the same package. Nonregistry versions include packages installed from disk (local) or from Git or as a tarball.

If you want to install a specific version of a package, refer to [install the package by name](upm-ui-quick.md) and follow its optional step to specify a version.

If you want to update packages that you installed from a Git URL, you can use any of the following methods:

* Locate the package in the **Package Manager** window, select it, then click the **Update** button.
* [Reinstall the package as a Git dependency](upm-ui-giturl.md) using a new revision. For more information about how to specify revisions with **Git dependencies**, refer to [Targeting a specific revision](upm-git.md#revision).
* Reinstall the package from the Unity Registry.

To change to a different version of a UPM package:

1. Open the **Package Manager** window and select **In Project**, **Unity Registry**, or **My Registries** from the [navigation panel](upm-ui-nav.md). You can also select the **Updates** entry, which lists all packages in your project that have updates available. An arrow icon appears next to packages that have updates available.
2. Select the installed package you want to update from the [list of packages](upm-ui-list.md). The package information appears in the [details panel](upm-ui-details.md).
3. If the package is part of a **feature set** (indicated by the lock icon near the package name) you must unlock the package before you switch versions. To unlock the package and select another version, select **Unlock**. Starting with Unity 6.3, **Unlock** is in the **Manage** dropdown.

   **Note**: The package is temporarily unlocked. If you select a different context in the [navigation panel](upm-ui-nav.md), or close either the **Package Manager** window or the Unity Editor, the package reverts to a locked state. However, if you change versions when the package is unlocked (for example, with the [Install a package from a registry by name](upm-ui-quick.md) method), the package stays unlocked.
4. Select a package in the [list of packages](upm-ui-list.md).
5. In the [details panel](upm-ui-details.md), select the **Version History** tab. If multiple versions are available, expand the entries to view information specific to each version.
6. Click the **Update to #.#** button, or click the **Update** button beside the version listed in the **Version History** tab.

   When the progress bar finishes, any new functionality is immediately available.

**Notes:**

* If you switch from a nonregistry version of a package to a registry version, the Package Manager removes the nonregistry version from its version history, but the **local package** still exists on disk, with any customization intact.
* If you switch to an older version of a package, you might have to run the [API Updater](APIUpdater.md) on the package contents.
* You can switch versions for multiple packages with one click by using the multiple select feature. For more information, refer to [Perform an action on multiple packages or feature sets](upm-ui-multi.md).

## Additional resources

* [Package types](upm-package-types.md)
* [Add and remove UPM packages or feature sets](upm-ui-actions.md)