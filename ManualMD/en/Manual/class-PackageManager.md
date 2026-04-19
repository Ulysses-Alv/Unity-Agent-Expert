# Package Manager

To access the **Package Manager** settings, open the **Edit** menu, then select **Project Settings**, then select the **Package Manager** category. These settings let you manage your scoped registries and change the advanced settings for your current project.

![The Package Manager settings window. The top of the window displays pre-release package settings (A). The main area of the window (B) displays scoped registries (1) and details for the selected scoped registry (2).](../uploads/Main/class-PackageManager.png)

The Package Manager settings window. The top of the window displays pre-release package settings (A). The main area of the window (B) displays scoped registries (1) and details for the selected scoped registry (2).

**(A)** Under the **Advanced Settings** section, toggle whether [pre-release packages](#advanced_preview) are visible in the Package Manager window.

**(B)** Use the **Scoped Registries** group to manage scoped registries for your project. This group has the following areas:

**(1)** The list of scoped registries currently defined for this project.

**(2)** The details for the selected scoped registry.

## Advanced Settings

Use the **Advanced** Settings group to toggle the [Show Pre-release Package Versions](#advanced_preview) setting for your current project.

### Show Pre-release Package Versions

Enable the **Show Pre-release Package Versions** setting to display [pre-release packages](pack-preview.md) in the Package Manager window. Pre-release packages appear with the ![Pre](../uploads/Main/iconPre.png) tag in the [list panel](upm-ui-list.md) and the ![Pre-Release](../uploads/Main/iconPrerelease.png) tag in the [details panel](upm-ui-details.md). These tags serve as a reminder of which packages will release with the next long-term support (LTS) version of the Unity Editor.

**Warning**: By default, this setting is disabled because pre-release packages are still being verified as stable enough for production, but aren’t fully released yet. For more information about the lifecycle of packages, refer to [Package states and lifecycle](upm-lifecycle.md).

## Scoped Registries

The **Scoped Registries** group displays a list of scoped registries added to the current project on the left of the settings group, and the details of the selected registry on the right.

For detailed information about scoped registries, including how to use and configure them with the Unity Package Manager, refer to [Scoped Registries](upm-scoped.md).