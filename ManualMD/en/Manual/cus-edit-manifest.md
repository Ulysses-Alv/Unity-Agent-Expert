# Edit the package manifest

The package manifest (`package.json`) defines the metadata for a Unity Package Manager (UPM) package, including its registered name, version, dependencies, Unity Editor compatibility, and other configuration details. The Unity Editor uses this file to identify, load, and manage the package within the project.

You can view or edit the package manifest depending on the context:

* View the contents of the package manifest for **UPM packages** you added to your project from a registry, as a tarball, or from a Git URL.
* Edit the contents of the package manifest for packages you’re developing or customizing.

## View the package manifest

For all packages, including **immutable** packages, you can view most of the package manifest’s properties directly in the Unity Editor:

1. Open the **Package Manager** window.
2. Select a package from the [list panel](upm-ui-list.md).
3. Use one of the following methods to access the ****Inspector**** window:

   * From the [details panel](upm-ui-details.md), select **Manage** > **Open Manifest**.
   * Manually open or go to the [Inspector window](InspectorOptions.md).

The **Inspector** window displays information about the package, which it reads from the package manifest file (`package.json`).

To view all fields, including ones not shown in the **Inspector** window, open the `package.json` file with one of the following methods:

* From the [details panel](upm-ui-details.md), select **Manage** > **Open Manifest Externally**, to open `package.json` in the script editor defined in [Preferences](Preferences.md) > **External Tools**.
* [Locate the manifest file](#locate-manifest) manually and open it in a script editor.

## Locate the manifest file

You can locate the `package.json` file to view the complete contents of the package manifest. If the package is **mutable**, you can also edit the file.

You can use several methods to locate the manifest file.

### Locate the manifest file using the Package Manager window

1. Open the **Package Manager** window.
2. Select a package from the [list panel](upm-ui-list.md).
3. From the [details panel](upm-ui-details.md), select **Locate**. Focus shifts to the **Project** window, with the package manifest file (`package.json`) highlighted.
4. To open the package manifest file, right-click `package.json` and select **Open**. You can also select **Show in Explorer** (Windows) or **Reveal in Finder** (macOS), then open the file manually.

### Locate the manifest file manually

1. Open your operating system’s file management application.
2. Go to the folder for your Unity project or the folder where you’re developing a **local package**.
3. If you’re in the Unity project folder, open the `Packages` subfolder, then go to your package’s subfolder
4. Open the `package.json` file with your preferred script editor.

With the `package.json` file open, refer to the [package manifest reference](upm-manifestPkg.md) to identify the different fields and their usage.

## Edit the manifest file for mutable packages

If you’re developing a UPM package or customizing a registry package you copied into your project, follow these steps when you need to make changes to the manifest file.

You can use several methods to edit the manifest file.

### Edit the manifest file within the Unity Editor

1. Open the **Package Manager** window.
2. Select your package from the [list panel](upm-ui-list.md).
3. From the [details panel](upm-ui-details.md), select the **Manage** dropdown button and select **Edit Manifest**.

   **Note**: You can also manually open or change to the [Inspector window](InspectorOptions.md) and select **Edit**.
4. Make your changes while referring to the [package manifest reference](upm-manifestPkg.md).

   **Note**: If you change the **Name** field in the **Inspector** window, changes replace only the final identifier in the `name` property of the [package manifest](upm-manifestPkg.md) file.
5. Select **Apply** to save your changes.

If fields you need to edit aren’t available in the **Inspector** window, use the manual method for [editing the manifest file in a script editor](#edit-man-editor).

### Edit the manifest file in a script editor

1. Open the **Package Manager** window.
2. Select your package from the [list panel](upm-ui-list.md).
3. From the [details panel](upm-ui-details.md), select the **Manage** dropdown button and select **Edit Manifest Externally**. The `package.json` file opens in the script editor defined in [Preferences](Preferences.md) > **External Tools**.
4. Make your changes while referring to the [package manifest reference](upm-manifestPkg.md).
5. Save the file.

## Additional resources

* [Package types](upm-package-types.md)
* [Package creation](cus-pkg-lp.md)
* [Package manifest reference](upm-manifestPkg.md)