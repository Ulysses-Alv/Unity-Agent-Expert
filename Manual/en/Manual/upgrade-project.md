# Upgrade your Unity project

Upgrading your project gives you access to the latest features, improvements, and bug fixes introduced in a newer version of the Unity Editor.

This page details some recommended best practices to upgrade your project. However, how you apply those practices to your project depends on the circumstance of your project. The experience of upgrading your project can vary widely depending on several factors such as:

* The version you’re upgrading to and from.
* The complexity and size of your project.
* The packages and assets that you use in your project.

The recommended best practice is to upgrade your project sequentially from one major version to the next. For example, if you’re upgrading from 2020 LTS to 2022, you should first upgrade to 2021 LTS, then to 2022 LTS. If you upgrade your project from a much older version straight to a newer version, you might encounter more issues than if you upgrade your project sequentially.

To upgrade your Unity project to a newer version of the Unity Editor:

* [Upgrade your Unity project](#upgrade-your-unity-project)
* [Determine if you should upgrade](#determine-if-you-should-upgrade)
* [Back up your project](#back-up-your-project)
* [Prepare your project for the upgrade](#prepare-your-project-for-the-upgrade)
* [Install a new version of the Unity Editor](#install-a-new-version-of-the-unity-editor)
* [Open your project in the new version of the Unity Editor](#open-your-project-in-the-new-version-of-the-unity-editor)
* [Resolve compatibility issues and adjust your project for changes](#resolve-compatibility-issues-and-adjust-your-project-for-changes)
* [Test your project](#test-your-project)
* [Additional resources](#additional-resources)

## Determine if you should upgrade

Before you upgrade your project, make sure that the advantages of upgrading outweigh the effort required to upgrade for your specific case. Large, complex projects with several third-party assets and packages attached might take a long time to upgrade and require significant changes.

If your project uses many third-party packages and assets, refer to their documentation to check if they’re compatible with the version of the Unity Editor you’re upgrading to. If they’re not compatible with your target version of the Unity Editor, you might need to find an alternative solution or update the assets yourself. Consider reaching out to the asset publishers for support.

For help with upgrading, refer to:

* The [What’s New in Unity](WhatsNew.md) documentation to learn about new features and improvements introduced in the version of the Unity Editor you’re upgrading to.
* The [Upgrade Guide](UpgradeGuides.md) for the version you’re upgrading to. The upgrade guides list changes in their version that might affect existing projects when you upgrade them from a previous version. Refer to the upgrade guide for breaking and behavioral changes introduced in the version they describe.
* The [Release notes](https://unity.com/releases/editor/archive) for a more exhaustive list of changes in the version you’re upgrading to.
* The documentation for the [packages](PackagesList.md) you use in your project. Check if the packages and assets in your project are compatible with the version of the Unity Editor you’re upgrading to. To find the documentation for a package, refer to [Finding package documentation](upm-docs.md).
* Release posts on the [Unity Blog](https://unity.com/blog) for more information about new features and improvements.

## Back up your project

Before you update your project, it’s crucial that you back up your project files. A backup ensures that you can revert to the previous version of your project if you have any issues during the upgrade process. You can also use a backup to compare the performance of your upgraded project against the previous version. The recommended best practice to back up your project is to use [version control](VersionControl.md).

If you’re not using version control, you can manually back up your project by copying the project folder to another location on your computer or an external drive. In most scenarios, using a version control system is better than manually backing up your project files.

**Tip**: If you’re manually backing up your project, you can exclude some auto-generated folders such as the **Library** folder to save disk space. The **Library** folder is regenerated when you open the project in the Unity Editor.

## Prepare your project for the upgrade

To prepare your project for upgrade:

* Upgrade the existing packages and assets in your project to their latest versions. This gives you an opportunity to identify and fix bugs in a more managed way before you open your project in a newer version of the Unity Editor.
* Review your project for any features or APIs that were deprecated in the version of the Unity Editor you’re upgrading to. Make a plan to update or replace them.
* Review your target build platforms and ensure that they are still supported in the version of the Unity Editor you’re upgrading to.

## Install a new version of the Unity Editor

Using the [Hub](unity-hub.md) is the recommended best practice to install the Unity Editor, create projects, and manage your Unity experience. The Hub provides a central location to manage your Editor installations, accounts, licenses, and projects. To install the newest version of Unity Editor in the Hub, refer to [Install and uninstall Editor versions](https://docs.unity3d.com/hub/manual/AddEditor.html).

Make sure you install the same modules that you used in the project you are upgrading.

To learn other ways to install the Unity Editor, refer to [Install Unity](GettingStartedInstallingUnity.md).

To manually install a specific version of the Unity Editor, you can also refer to [the Unity Download Archive](https://unity.com/releases/editor/archive).

## Open your project in the new version of the Unity Editor

To open your project in the new version of the Unity Editor:

1. Open the Unity Hub.
2. Select the **Projects** tab.
3. In the row of the project you want to upgrade, select the **Editor Version** to open the **Installs** window. The **Installs** window displays all the Unity Editor versions you have installed.
4. In the **Installs** window:
   1. Select the Editor version you want to use to open the project.
   2. If you want to change the platform that the project is built for, select one of the available platforms in the same row as the Editor version. By default, **Current platform** is selected.
   3. Select the **Open with…** button.
5. In the **Change Editor version?** window, select **Change version**.
6. In the **Opening Project in Non-Matching Editor Installation** window, select **Continue**.
7. If your project contains compilation errors, the **Enter Safe Mode?** window displays. If it does, select **Enter Safe Mode**. Safe Mode provides the best environment to resolve errors without waiting for the rest of your project to import. To learn more, refer to [Safe Mode](SafeMode.md).

**Tip**: You can also select **Install Other Editor Version** in the **Installs** window to install a different version of the Unity Editor.

## Resolve compatibility issues and adjust your project for changes

When you upgrade your project to a new version of the Unity Editor, you might encounter compatibility issues and deprecations introduced in the new version.

**Tip**: If you run into an issue that’s not covered in the release documentation, you can ask for help from the wider Unity community at [Unity Discussions](https://discussions.unity.com).

To resolve compatibility issues, open the project in [Safe Mode](SafeMode.md), and do the following:

* Review the **Console** window for any errors or warnings. Resolve any errors you find and act on deprecation notices.
* Your project might contain incompatible package versions that can cause errors. Open the Package Manager and update [packages](upm-ui-update.md) and [asset packages](upm-ui-update2.md) to their latest versions.
* Test build compatibility. If you’re targeting specific platforms, such as iOS, Android, or Web, ensure that the project builds and runs correctly on those platforms.

## Test your project

Test your project to ensure that it works as expected in the new version of the Unity Editor:

* Run your game and test all its features.
* Use your backed up project in the original Editor version to compare the performance of your project across versions. For more information, refer to [Unity Profiler](Profiler.md).
* Specifically test any functionality in your game that might be affected by a change described in release documentation like the [What’s new in Unity guide](WhatsNew.md), the [Upgrade Guide](UpgradeGuides.md), and the [release notes](https://unity.com/releases/editor/archive) for the version you’re upgrading to.
* Pay extra attention to custom **scripts**, **shaders**, or third-party assets and how they function in your project. They might require updates for compatibility.
* Check for and [upgrade outdated packages](upm-ui-update.md). Older package versions might cause performance issues in newer Editor versions.

## Additional resources

* [Upgrade Unity](UpgradeGuides.md)
* [What’s New in Unity](WhatsNew.md)
* [System requirements for Unity](system-requirements.md)
* [Install Unity](GettingStartedInstallingUnity.md)
* [API updater](APIUpdater.md)
* [Safe Mode](SafeMode.md)