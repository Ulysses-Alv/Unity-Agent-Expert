# Remove a UPM package from a project

When you “remove” a **UPM package** from your project, the Package Manager is actually removing the project’s **direct dependency** from your [project manifest](upm-manifestPrj.md). The result of removing the direct dependency varies, based on the dependencies for the package you are removing:

* If there are no other packages or **feature sets** that have a dependency on this package, any Editor or runtime functionality that it implemented is no longer available in your project. For more information about direct and **indirect dependencies**, refer to [Package dependency and resolution](upm-dependencies.md).
* If another installed package or an installed feature set depends on the package you are trying to remove, this procedure removes only the dependency from your project manifest. The package itself and all its functionality is still installed in your project, and appears in the [list panel](upm-ui-list.md) with the dependency icon ![Dependency icon](../uploads/Main/iconDependency.png).

## Before you begin

Make sure you understand these important notes before you begin:

* Use this procedure to remove a package only if you added it to the current project by *installing* it, such as (but not limited to) [Install a feature set](fs-install.md), [Install a UPM package from a registry](upm-ui-install.md), [Install a UPM package from Asset Store](upm-ui-install2.md), and installing custom packages. Don’t use this procedure to try to:

  + Remove **asset packages** that you *downloaded* *and* *imported* to your project. For information about removing asset packages that you downloaded and imported, refer to [Remove imported assets from a project](upm-ui-remove-asset.md).
  + Remove local asset packages that you *imported* to your project. For information about removing local asset packages that you imported, refer to [Removing local asset packages](upm-ui-remove-local.md).
* If you use this procedure to remove a UPM package that you [installed from a registry](upm-ui-install.md) or [installed from the Asset Store](upm-ui-install2.md), the operation removes the package from the current project. It doesn’t remove the same package that might exist in other projects. It also doesn’t remove the package from the global cache; this action isn’t supported by the Package Manager, and manually manipulating the global cache is discouraged.
* If you use this procedure to remove a package that you [embedded in your project](upm-concepts.md#Embedded), the Package Manager deletes the entire package folder from your computer. However, removing packages installed from any other source (including [local](upm-concepts.md#Local) packages) removes only the reference to the package in the manifest but leaves the package itself and its contents intact.

## Procedure

To remove an installed package:

1. Open the **Package Manager** window and select **In Project** from the [navigation panel](upm-ui-nav.md).
2. Select the package you want to remove from the [list of packages](upm-ui-list.md). The [details panel](upm-ui-details.md) now displays that package’s information.
3. Click **Remove**.

   If this button isn’t displayed, you might be viewing the **My Assets** list. Refer to [Before you begin](#prereqs).

   If this button is disabled, you can’t remove this package. Hover over the button to find out why you can’t remove the package. For more information, refer to [Locked and non-removable packages](#noremovals).

   When the progress bar finishes, the package disappears from the list.
4. If you want to restore a removed UPM package, follow the instructions to [install a UPM package from a registry](upm-ui-install.md) or [install a UPM package from the Asset Store](upm-ui-install2.md).

**Note**: You can remove multiple packages with one click by using the multiple select feature. For more information, refer to [Perform an action on multiple packages or feature sets](upm-ui-multi.md).

## Locked and non-removable packages

You can remove packages only if they’re not required by another package or a feature set. The Package Manager enforces this by disabling the **Remove** button for all required packages.

| **Required by** | **Description** |
| --- | --- |
| [A feature set](FeatureSets.md) | If a feature set requires the package, it displays a lock icon in both the [list panel](upm-ui-list.md) and in the [details panel](upm-ui-details.md). The details panel also displays the name of the feature set that requires the package below the lock icon in the details panel.   However, even if you click the **Unlock** button, you still can’t remove the package from your project until you remove all feature sets that require it. Unlocking the package lets you [request a different version](upm-ui-update.md) for your project, but it still doesn’t let you remove it. |
| [Another package](upm-dependencies.md) | If one or more packages require the selected package, the **Remove** button is disabled. You can find the name of the package that has the dependency from the **Dependencies** tab in the [details panel](upm-ui-details.md). If you don’t need the other packages, you can remove them and the Package Manager automatically removes this package too. |

**Note**: You can unlock multiple packages with one click by using the multiple select feature. For more information, refer to [Perform an action on multiple packages or feature sets](upm-ui-multi.md).

## Additional resources

* [Package types](upm-package-types.md)
* [Add and remove UPM packages or feature sets](upm-ui-actions.md)
* [Add and remove asset packages](upm-ui-actions-ap.md)