# Create and export asset packages

To share any assets from your project’s `Assets` folder, you can export them as an **asset package** in the `.unitypackage` format. Asset packages are a convenient way to copy several assets or an entire **scene** from one project to another quickly, with little additional overhead.

If you want to distribute assets in a more formal way with long-term support, a defined structure, and semantic versioning, an alternative option is to [create your own Unity package](CustomPackages.md).

**Note**: If the assets in your asset package are high quality and other users might find them useful, refer to [Publishing to the Asset Store](AssetStorePublishing.md) for information on how to create a package draft and upload it to the **Asset Store**.

To create and export an asset package:

1. Open the project you want to export assets from.
2. Go to the **[Project](ProjectView.md)** window and select one or more items or folders. These items become the starting list for your export. You can also select the `Assets` folder to include all assets as your starting point.
3. Choose **Assets** > **Export Assets** from the menu to open the **Exporting package** dialog.
4. In the dialog, select the checkboxes for the assets you want to include in the package.
5. Enable the **Include dependencies** checkbox to automatically select any assets (including scripts) used by the assets you select.
6. Enable the **Include all **scripts**** checkbox to export all scripts from your project.

   If you disable **Include all scripts** but enable **Include dependencies**, Unity only exports scripts that your selected items directly depend on. Unity can’t trace the full chain of dependencies for your scripts, so enabling **Include all scripts** reduces the likelihood of compilation errors when using the exported package in another project.
7. Click **Export** to bring up the file explorer and choose where you want to store your package file.
8. Name and save the package.

## Re-exporting asset packages

To create a new version of your asset package with updated contents, perform the export procedure again and select the changed and unchanged files to include in the new version.

### Naming strategies

Name your updated package using incremental names suffixed with a version identifier, for example `MyAssetPackageVer1`, `MyAssetPackageVer2`, and so on. Unity recognizes these as updates of the same package. Use a naming convention that’s clear for you and anyone you share it with.

**Warning:** Don’t remove files from asset packages and then add different files with the same name. If you remove a file and later replace it, use a unique name for the replacement file. Attempting to reuse previously used file names can confuse Unity’s system for tracking assets and produce warnings on import. For more information, refer to [Asset metadata](AssetMetadata.md).

## Additional resources

* [Create custom packages](CustomPackages.md)