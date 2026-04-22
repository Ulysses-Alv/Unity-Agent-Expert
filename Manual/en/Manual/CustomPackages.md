# Package development workflow

Create your own Unity Package Manager (UPM) package to extend the functionality of the Unity Editor.

You can use packages that you create for one of your projects or for multiple projects. You can also share your package with others in your organization or studio so they can use it in any project. You can even distribute your package through official channels, for the wider community to use.

Follow this workflow to create entirely new **UPM packages**. For information about modifying copies of existing packages, refer to [Copy a Unity package to your project folder](upm-embed.md#embed-cached).

To create your own UPM package, complete the following tasks:

1. [Create the package structure](#createPkg).
2. [Add code to your package](#addCode).
3. [Edit the assembly definition](#edit-asmdef).
4. [Edit the package manifest](#editMan).
5. [Create samples for your package](#createSample).
6. [Add tests to your package](#addTests).
7. [Set the package version number](#serVersion).
8. [Update the changelog for your package](#updateLog).
9. [Add legal requirements to your package](#addLegal).
10. [Document your package](#docPkg).

## Prerequisites

Before you begin, determine whether you want to create your package inside or outside your project folder. For more information, refer to [Comparison of package creation locations](cus-location.md).

## Create the package structure

Use Package Manager’s **(+)** > **Create package** feature to create the folder structure and starter files. For more information, refer to [Create the package structure](cus-create.md).

## Add code to your package

Add the code, libraries, and any required assets that make your UPM package unique. For more information, refer to [Add code to your package](cus-add-code.md).

## Edit the assembly definition

Any code in your package must be organized into one or more assemblies using assembly definition files (`.asmdef` assets). Unity creates a separate assembly for each assembly definition, which includes all code in its folder and subfolders.

Package Manager’s **Create package** operation creates these assembly definitions for you, but you might need to modify them as your package evolves.

For more information, refer to [Edit the assembly definition](cus-asmdef.md).

## Edit the package manifest

The **package manifest** (`package.json`) defines the metadata for a UPM package, including its name, version, dependencies, and other configuration details.

Package Manager’s **Create package** operation creates this file for you, but it contains the minimum amount of required information. As your package evolves, you usually need to add more information to the package manifest.

For more information, refer to [Edit the package manifest](cus-edit-manifest.md).

## Create samples for your package

Although including samples in your package is optional, they’re a valuable tool for helping users learn how to use your package.

For more information, refer to [Create samples for your package](cus-samples.md).

## Add tests to your package

Including tests in your package is optional, but they’re a useful way to verify that your code works as intended and to help others validate changes safely.

The **Create package** operation creates example test files in the proper subfolder for you to modify.

For more information, refer to [Add tests to your package](cus-tests.md).

## Set the package version number

Your package version number must follow the principles of semantic versioning.

For information about semantic versioning and how to set your package’s version number, refer to [package versioning](upm-semver.md).

## Update the changelog for your package

Update the `CHANGELOG.md` file every time you publish a new version of your package.

For more information, refer to [Update the package changelog](cus-changelog.md).

## Add legal requirements to your package

Your package might have legal requirements or third party elements, which you need to communicate to those who consume it.

For information, refer to [Add required legal files to your package](cus-legal.md).

## Document your package

Including documentation in your package is optional, but offering clear guidance gives users a better experience and makes your package easier to adopt.

For information, refer to [Document your package](cus-document.md).

## Additional resources

* [Package creation](cus-pkg-lp.md)
* [Create and export asset packages](AssetPackagesCreate.md)