# Install a UPM package by name

If you know the exact name of a **UPM package** you want to install, you can use the **Install package by name** option to install it. This is a quick way to install a UPM package that comes from a registry.

This method works for any UPM package and version that’s currently hosted on the Unity package registry or any [scoped package registry](upm-scoped.md) you have set up for the current project. It also applies to any UPM package in the Unity registry that came from the **Asset Store**.

Specifying the version is optional. If you don’t know which version to install, or want to install the latest compatible version, enter only the package name.

**Note:** The latest compatible version might not be the latest published package. If there is a [released](pack-safe.md) package version and a newer [pre-release](pack-preview.md) or [experimental](pack-exp.md) version, Package Manager selects the released package version, unless you explicitly input a value in the optional **Version** field.

## Before you begin

Make sure you know the package’s name. The package name is a unique identifier, not the display name used on user interfaces and documentation.

For packages in the Unity registry, the name’s unique identifier uses reverse domain name notation. Examples include `com.unity.example` and `com.meta.example`. For packages in a scoped registry, the name might not follow the same pattern.

For any package in the UPM format, if you can view the package in Package Manager, select it and view its details in the ****Inspector**** window. The **Name** value is the package name.

Other ways of determining a package name vary, depending on the registry that hosts it:

* For a package in the Unity registry:
  + Use the lists in [Released packages](pack-safe.md) and [Pre-release packages](pack-preview.md) to identify a package by its display name, such as `2D Animation`. The hyperlink value is the name of that package; in this case, `com.unity.2d.animation`.
  + The [package documentation](upm-docs.md) might have installation instructions that explicitly provide the `name` value.
  + The [package documentation](upm-docs.md) URL often nests the package name in its address. Using the **AR** Foundation
    package as an example, its package documentation URL is `https://docs.unity3d.com/Packages/com.unity.xr.arfoundation@6.2/manual/index.html`; the `com.unity.xr.arfoundation` component is that package’s name.
* For a package in a scoped registry, if you can’t determine its name by using the Package Manager and Inspector windows, contact the package creator and request the package name. The package creator recorded this value as a required property in the [package manifest](upm-manifestPkg.md) file (`package.json`).

## Procedure

To install a UPM package by name, follow these steps:

1. Open the [Package Manager window](upm-ui-access.md), if it’s not already open.
2. Click the **install** ![Install button](../uploads/Main/iconAdd.png) button in the toolbar. The options for installing packages appear.
3. Select **Install package by name** from the install menu. Two text boxes and an **Install** button appear.
4. Enter the package **Name**, as determined in the [Before you begin](#prereqs) section.

   **Note**: If you enter an invalid package name or version, Package Manager warns you that it can’t find that name or version. Verify that the package name and version are correct and try again.
5. (Optional) If you know which version you want to install, enter the [full package version](upm-manifestPkg.md#version), such as `1.3.0-pre.2`, in the **Version (optional)** box.
6. Click **Install**. If Unity was able to install the package successfully, the package now appears in the [package list](upm-ui-list.md) like any other installed UPM package.

## Additional resources

* [Package types](upm-package-types.md)
* [Add and remove UPM packages or feature sets](upm-ui-actions.md)
* [Add and remove asset packages](upm-ui-actions-ap.md)