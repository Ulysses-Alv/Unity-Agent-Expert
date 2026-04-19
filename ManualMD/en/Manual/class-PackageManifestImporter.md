# Package Manifest window reference

To view or edit a **package manifest**, select the `package.json` file from a `Packages` subfolder in the **Project** window. The Unity Editor displays the manifest’s contents in a dedicated ****Inspector**** window.

Starting with Unity 6.3, you can also access this dedicated **Inspector** window from within the **Package Manager** window. For more information, refer to [preview the package manifest](cus-edit-manifest.md#open-manifest).

This window displays read-only fields if the package is **immutable**, such as a package you installed from a registry, as a tarball, or from a Git URL. The fields are editable if the package is **mutable**, such as a package you’re developing locally, or a package you copied into the `Packages` folder of your project.

![The Package Manifest window for a mutable package (left) and an immutable package (right)](../uploads/Main/class-PackageManifestImporter.png)

The Package Manifest window for a mutable package (left) and an immutable package (right)

**(A)** You can select **Importer**, **Edit**, or **Open** when the package is mutable. Select **Edit** to make the read-only fields editable. Select **Open** if you prefer to load this package manifest in your default script editor.

* The **Edit** button appears when you access a mutable package from the **Package Manager** window then change to the **Inspector** window.
* The **Open** button appears when you select **Edit** or access the a mutable package’s `package.json` file from the **Project** window. Access this file by opening it in your file management application or by selecting **Locate** in the **Package Manager** window.
* To view the JSON directly within the **Inspector** window, open the **Importer** menu and select **UnityEditor.TextScriptImporter**.

**(B)** Select **View in Package Manager** to open the Package Manager window and load this package in its [details panel](upm-ui-details.md).

**(C)** The [Information](#Info) section has details about this specific package version.

**(D)** The **Brief description** text box shows the `description` field of the manifest. This text also appears in the [details panel](upm-ui-details.md) of the Package Manager window. For more information, refer to the documentation for the [description](upm-manifestPkg.md#description) property.

**(E)** The **Dependencies** section shows the list of packages that this package depends on. Refer to [Dependencies](#Depend) for information about how to manage the dependencies of a package you’re editing.

**(F)** If the package is mutable, **Visibility in Editor** determines whether the **Project** window hides the package’s assets and omits them from the results when you use the [Object Picker](InspectorReferences.md) in the **Inspector window**. This field maps to the [hideInEditor](upm-manifestPkg.md#hideInEditor) property of the package manifest file (`package.json`).

**(G)** If the package is mutable and you made changes, select **Revert** to discard the changes. Select **Apply** to save any changes you have made to the manifest.

## Information section

| **Property** | **Description** |
| --- | --- |
| **Name** | A substring of the package’s technical (or official) name. This field maps to the [name](https://docs.unity3d.com/6000.2/Documentation/Manual/upm-manifestPkg.html#name) property of the package manifest file (`package.json`) with the leading identifiers removed. |
| **Organization name** | The identifier of the [Unity Organization](https://docs.unity.com/cloud/en-us/organizations) that created this package. |
| **Display name** | The user-facing name on display in the **Project** window and the **Package Manager** window. This field maps to the [displayName](upm-manifestPkg.md#displayName) property of the package manifest file (`package.json`). |
| **Version** | The package version number. For more information, refer to the documentation for the [version](upm-manifestPkg.md#version) property. |
| **Minimum Unity version** | For mutable packages, you can enable this option to specify the lowest Unity version this package is compatible with. When you enable this option, the **Major**, **Minor**, and **Release** properties appear.   If this package is compatible with all Unity versions, clear this checkbox and remove the **Major**, **Minor**, and **Release** properties.   For more information, refer to the documentation for the [unity](upm-manifestPkg.md#unity) property. |
| **Major** | Specify the MAJOR part of the minimum Unity version. For more information, refer to the documentation for the [unity](upm-manifestPkg.md#unity) property. |
| **Minor** | Specify the MINOR part of the minimum Unity version. For more information, refer to the documentation for the [unity](upm-manifestPkg.md#unity) property. |
| **Release** | Specify the UPDATE and RELEASE part of the minimum Unity version. For more information, refer to the documentation for the [unityRelease](upm-manifestPkg.md#unityRelease) property. |

Depending on the window’s state (editable or read-only), additional information displays for [optional properties](upm-manifestPkg.md#optional), such as **Documentation URL**, **Licenses URL**, **Changelog URL**, and **Author**.

## Dependencies section

Lists the other packages that are dependencies for this package. Each entry consists of the technical name for the package (for example, `com.unity.probuilder`) and its version number.

If you’re viewing a mutable package, you can add a dependency by following these steps:

1. Select the **Add** (**+**) button. A new row appears in the list.
2. Enter the package name on the left and the version on the right. For more information, refer to the documentation for the [dependencies](upm-manifestPkg.md#dependencies) property.

If you’re viewing a mutable package, you can remove a dependency by following these steps:

1. Click the selector ![Selector](../uploads/Main/iconSelect.png) to the left of the package you want to remove.
2. Select the **Remove** (**−**) button. The row disappears from the list.

## Additional resources

* [Package manifest file](upm-manifestPkg.md)