# Install a UPM package from a local folder

The Package Manager can load a **UPM package** from anywhere on your computer even if you saved it outside your Unity project folder (for example, if you have a package called **com.unity.my-local-package** and you save it on the `Desktop` but your Unity project is under the `Documents` folder).

You can also use a folder inside your project folder, provided that it is not one of the [reserved project sub-folders](#PkgLocation).

To install a UPM package from your local disk:

1. Open the [Package Manager window](upm-ui-access.md), if it’s not already open.
2. Click the **install** ![](../uploads/Main/iconAdd.png) button in the toolbar. The options for installing packages appear.

   ![Install package from disk button](../uploads/Main/upm-ui-local.png)

   Install package from disk button
3. Select **Install package from disk** from the install menu to bring up a file browser.
4. Navigate to the folder root of your **local package**.
5. Double-click the `package.json` file in the file browser.

The file browser closes, and the package now appears in the [package list](upm-ui-list.md) with the ![](../uploads/Main/iconLocal.png) label.

![A package installed from a local folder, with the option to update to a higher version](../uploads/Main/upm-ui-local-ver.png)

A package installed from a local folder, with the option to update to a higher version

Remember that if you updated to the registry version but you made changes locally to your project, the registry version will overwrite your local changes.

## Local packages inside your project

You can place a local package anywhere inside your project except under these folders:

| **Project folder:** | **Reason:** |
| --- | --- |
| `Assets` | If you place a package inside this folder, the Asset Database imports any assets under this folder twice: once as assets and once as package contents. |
| `Library` | Do not modify the contents of this folder. |
| `ProjectSettings` | This folder is for [settings](comp-ManagerGroup.md) assets only. |
| `Packages` | If you place a package under this folder, the Package Manager automatically interprets it as an **Embedded package**, regardless of the reference in the **project manifest**. |

## Additional resources

* [Package types](upm-package-types.md)
* [Add and remove UPM packages or feature sets](upm-ui-actions.md)
* [Add and remove asset packages](upm-ui-actions-ap.md)