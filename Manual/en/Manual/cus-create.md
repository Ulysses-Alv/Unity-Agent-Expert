# Create the package structure

Create the structure for your own Unity Package Manager (UPM) package by using the Package Manager window.

This process creates the layout and mandatory elements for your package. You can provide other details and elements of your package in later tasks.

This process creates a **UPM package** inside your project folder with the intention of using this package in the current project. However, these instructions also explain what to do if you want to use this package in other (or multiple) projects.

## Create a UPM package

1. Make sure you’re signed in to the Unity Editor using your Unity ID.
2. Open the **Package Manager** window, if it’s not already open. Refer to [Access the Package Manager window](upm-ui-access.md).
3. Open the install menu (**+**).
4. Select **Create package**.
5. Type the **Package display name** for the package, which is the name that will appear at the top of the [details panel](upm-ui-details.md). Display names can contain alphanumeric characters and special characters. For more information, refer to [Package name guidelines](cus-naming.md).
6. Select **Create**.

After the domain reloads, the **Package Manager** window refreshes and shows your package in the [list panel](upm-ui-list.md) and [details panel](upm-ui-details.md).

You can also view your new package in the **Project** window, under the `Packages` folder.

The package creation operation performs the following actions:

* Creates a **Technical Name** for your package, in the format `<top-level-domain>.<org-name>.displayname`. This **Technical Name** maps to the `name` property of the [package manifest file](upm-manifestPkg.md#required) (`package.json`).
* Sets the package version to an initial value of `0.1.0`.
* Creates the required files, including the [package manifest file](upm-manifestPkg.md#required) (`package.json`) and populates them with default and placeholder values.
* Creates a set of files and folders recommended for the UPM package format. You can view these items in the **Project** window or in your file management application. For information about the full layout of folders and files created by this operation, refer to [Package layout](cus-layout.md).

## Use your package in other projects

If you plan to use this package in another project, or across multiple projects, it’s preferable to move the package out of your project folder.

For information about the differences between placing your package inside or outside your project folder, refer to [Comparison of package creation locations](cus-location.md).

To move your package out of the current project:

1. Open or go to the **Project** window in the Unity Editor.
2. Expand the `Packages` folder.
3. Right-click your package’s base folder and select **Show in Explorer** (Windows) or **Reveal in Finder** (macOS).
4. Use your file management application to move your package to a location outside of your project folder.

To use this package in the current project or any other project, you must establish a reference to this package in your **project manifest** file. You need to establish this reference only once. Although you can edit the [project manifest file](upm-manifestPrj.md) (`manifest.json`) manually to include a `dependencies` value to a [local folder](upm-localpath.md), using the **Package Manager** window is less error prone.

To establish a reference to this package in your project manifest file with the Package Manager window:

1. Go to the **Package Manager** window.
2. Open the install menu (**+**).
3. Select **Install package from disk**.
4. Go to the location where you moved your package folder and select that folder.
5. Select the `package.json` file in that folder.
6. Select **Open**.

The package appears in the **Package Manager** window with the **Local** label.

## Next steps

With the structure of your UPM package complete, you can proceed with other activities in the [package development workflow](CustomPackages.md). Typical activities at this stage of package development include:

* [Add code to your package](cus-add-code.md)
* [Edit the package manifest](cus-edit-manifest.md)
* [Edit the assembly definition](cus-asmdef.md)

## Additional resources

* [Package creation](cus-pkg-lp.md)
* [Package development workflow](CustomPackages.md)
* [Create and export asset packages](AssetPackagesCreate.md)