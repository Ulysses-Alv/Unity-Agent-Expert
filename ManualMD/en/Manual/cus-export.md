# Export and sign your UPM package

Export and sign your Unity Package Manager (UPM) package so you can share it with others.

Starting with Unity 6.3, publishers can sign their package directly from the Package Manager window or by using command line arguments.

When you sign a package, you must associate it with one of your organizations, so package consumers can quickly determine who created and owns it.

The only packages you can sign are packages that have the **Custom** or **Local** label in the Package Manager window:

* A **Custom** label appears on packages you’re developing or customizing in the `Packages` folder of your project.
* A **Local** label appears on packages you installed with the Package Manager’s [Install package from disk](upm-ui-local.md) method, or package folders you [reference as dependencies](upm-localpath.md) in your [project manifest file](upm-manifestPrj.md).

Use the following procedures to export and sign your **UPM package** with the Package Manager window, [command line arguments](#export-cli), or [API](#export-api).

### Sign a package from the Package Manager window

To sign a package from the Package Manager window:

1. Make sure you’re signed in to the Unity Hub.
2. Open the **Package Manager** Window.
3. Select the package you want to sign.
4. Select **Export**.

   ![The Package Manager window with the Export button for an editable package](../uploads/Main/upm-ui-export.png)

   The Package Manager window with the Export button for an editable package

   **Note**: Another way to access the export dialog is with the menus. Go to the **Project** window and select the base folder for your package in the `Packages` folder. Open the **Assets** menu (or right-click your package folder) and select **Export As UPM Package**.
5. In the **Export Package** dialog that appears, open the **Authoring Org** menu and select the organization you want to associate this package with.

   **Note**: For large projects whose contributors span multiple organizations, be sure to select the wider organization (or company-wide organization). If that organization doesn’t exist yet, refer to [Considerations for companies with multiple organizations](upm-signature.md#multi-org).
6. Select a location to store the signed package and select **Select Folder** (Windows) or **Choose** (macOS).

   **Note**: If the package already exists in that location, a warning message prompts you to confirm overwriting the file.
7. When the export process completes, your file management application opens at the location you specified, showing you the newly created file. A confirmation message also displays in the **Console** window.

The export operation creates a tarball file (`.tgz`), which is a compressed archive file. This tarball file contains an encrypted file (`.attestation.p7m`), which contains the package signature.

Refer to [Share your package](cus-share.md) for information about distributing this tarball file to others.

### Sign a package using command line arguments

If you publish packages using continuous integration (CI), you can sign your package from the command line.

To gather your organization’s **Organization ID** and sign your package from the command line:

1. If the project that uses the package you want to sign is open, close the Unity Editor.
2. Go to the [Unity Cloud Dashboard](https://cloud.unity.com/account/my-organizations) and select the organization you want to use to sign your package. You can also select the link for your account in the top right corner of the dashboard and select **Switch organization** for the intended organization, then open the menu again and select **Manage organization**.

   **Note**: For large projects whose contributors span multiple organizations, be sure to select the wider organization (or company-wide organization). If that organization doesn’t exist yet, refer to [Considerations for companies with multiple organizations](../Manual/upm-signature.md#multi-org).
3. From the **My organizations** page (or the **Organization Settings** page), locate the **Organization ID** field.
4. Copy the **Organization ID** value.
5. Open a command prompt window.
6. Change directories to the location of your Unity Editor. For more information, refer to [Unity Editor command line arguments](EditorCommandLineArguments.md) and [Locate the Editor program file](https://docs.unity3d.com/hub/manual/AddEditor.html#add-a-manually-downloaded-editor-to-the-hub).
7. Input the following command, replacing the placeholder values represented by angled brackets:

##### Windows

```
Unity.exe -batchmode -username <email_address> -password <your_password> \
    -upmPack <path_to_package_folder> <path_to_tarball> \
    -cloudOrganization <your_organization_id> -logfile -
```

##### macOS

```
Unity.app/Contents/MacOS/Unity -batchmode -username <email_address> -password <your_password> \
    -upmPack <path_to_package_folder> <path_to_tarball> \
    -cloudOrganization <your_organization_id> -logfile -
```

**Note**: Although optional, specifying `-logfile -` is helpful because it sends the command results to the command window.

| **Parameter to replace** | **Description** |
| --- | --- |
| `<email_address>` | The email address you use to sign in to Unity products and services. |
| `<your_password>` | The password you use to sign in to Unity products and services. |
| `<path_to_package_folder>` | The fully qualified path to the folder that contains the `package.json` file for the package you want to sign. **Note**: Don’t include `package.json` in this parameter value. |
| `<path_to_tarball>` | The output path where you want to save the signed tarball file (`.tgz`). **Note**: If the folder doesn’t exist, the command creates it for you. |
| `<your_organization_id>` | The Organization ID you copied from the Unity Cloud Dashboard. |

The output tarball file contains an encrypted file (`.attestation.p7m`), which contains the package signature.

Refer to [Share your package](cus-share.md) for information about distributing this tarball file to others.

### Sign a package using the API

You can use the `Client.Pack` method to sign a package. One of this method’s parameters is the ID of the organization that owns the package.

Follow these steps to get the organization ID and call the `Client.Pack` method:

1. Go to the [Unity Cloud Dashboard](https://cloud.unity.com/account/my-organizations) and select the organization you want to use to sign your package. You can also select the link for your account in the top right corner of the dashboard and select **Switch organization** for the intended organization, then open the menu again and select **Manage organization**.

   **Note**: For large projects whose contributors span multiple organizations, be sure to select the wider organization (or company-wide organization). If that organization doesn’t exist yet, refer to [Considerations for companies with multiple organizations](../Manual/upm-signature.md#multi-org).
2. From the **My organizations** page (or the **Organization Settings** page), locate the **Organization ID** field.
3. Copy the **Organization ID** value.
4. Make sure you’re signed into the Unity Editor as a member of the organization from the previous step.
5. Call the [Client.Pack](../ScriptReference/PackageManager.Client.Pack.md) method and include organization ID that you copied in a previous step.

## Additional resources

* [Package signatures](upm-signature.md)
* [Share your package](cus-share.md)
* [Create and export asset packages](AssetPackagesCreate.md)