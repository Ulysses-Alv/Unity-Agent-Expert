# Install a UPM package from a registry

**Note**: When you install a **UPM package** using Package Manager, the Package Manager evaluates other packages and their dependencies to determine if there are version conflicts with the version you selected. If it detects any conflicts or restrictions, it installs the version that resolves these issues. For more information, refer to [Package dependency and resolution](upm-dependencies.md).

Use the same process for installing a package from either the Unity registry or any [scoped registry](upm-scoped.md) defined in your project.

1. Open the **Package Manager** window and select **Unity Registry** from the [navigation panel](upm-ui-nav.md). If you set up a [scoped registry](upm-scoped.md) and you want to install a package from a scoped registry, choose **My Registries** instead.

   **Note**: If you haven’t defined any scoped registries for this project, **My Registries** doesn’t appear in the navigation panel.

   ![Select Unity Registry or My Registries](../uploads/Main/upm-ui-unityregistry.png)

   Select **Unity Registry** or **My Registries**
2. Select the package you want to install from the [list of packages](upm-ui-list.md). The package information appears in the [details panel](upm-ui-details.md).

   **Note:** If the Package Manager doesn’t list the package you want to install in the list of packages, it might be an undiscoverable package, such as an [experimental](pack-exp.md) or [pre-release](pack-preview.md) package. The Package Manager doesn’t display experimental packages unless they’re already installed, but pre-release packages appear in the Package Manager when you enable the [Show Pre-release Package Versions](class-PackageManager.md#advanced_preview) project setting.
3. Optional: If multiple versions are available, select the version to install. For more information on available versions, refer to [Find a specific version of a package](upm-ui-find-ver.md).
4. Click **Install**.

![Install button in the corner of the details panel](../uploads/Main/upm-ui-install.png)

When the progress bar finishes, the new package is ready to use.

**Note**: You can install multiple packages with one click by using the multiple select feature. For more information, refer to [Perform an action on multiple packages or feature sets](upm-ui-multi.md).

## Additional resources

* [Package types](upm-package-types.md)
* [Install a UPM package from Asset Store](upm-ui-install2.md)
* [Download and import an asset package](upm-ui-import.md)