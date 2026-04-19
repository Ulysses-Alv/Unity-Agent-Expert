# Contents of the Asset Database

Unity’s Asset Database keeps track of assets in your project to ensure the original asset source file and its imported counterpart remain in sync.

## Source assets and artifacts

On import, Unity converts most asset types from their original format to one optimized for use at runtime. The original file is referred to as the source asset and its imported counterpart is called an artifact. Unity converts source asset files to artifacts to make them easier to author within the Editor and easier for hardware to consume at runtime.

Artifacts have a serializable C# class representation that inherits from [`UnityEngine.Object`](../ScriptReference/UnityEngine.Object.md), so you can access their properties and work with them in code. For example, imported PNG files are converted to serialized instances of the [Texture](../ScriptReference/Texture.md) class, which inherits from `UnityEngine.Object`. This is the object you access when you interact with the Texture in code.

The difference between the source asset file and the artifact is bigger for some asset types than others. Some types of asset files Unity creates, such as **.**prefab****, **.**scene****, **.asset**, and **.mat**, already contain serialized data in their source file, so the artifact file Unity produces is similar to the source file. These file types are human-readable if [Asset Serialization Mode](class-EditorManager.md) is set to the default setting of **Force Text**, whereas binary files such as textures or audio are usually not.

Each artifact file name is a GUID with no file extension. Unity separates these files into subfolders, and each subfolder has a name that matches the first two characters of the artifact filename. These artifact files contain binary data, and are not designed to be human-readable. While it’s useful to understand that these files contain the data used by the Asset Database, you don’t need to view, edit, or use these files directly.

The [metadata file](AssetMetadata.md) which Unity creates during the import process, stored next to the original asset file, contains the asset’s import settings, and contains a GUID which allows Unity to connect the original asset file with the artifact in the asset database.

You can view the artifacts produced for the assets in your project in the [Import Activity window](ImportActivityWindow.md).

## Asset Database structure

The Asset Database consists of the following database files in the `Library` folder:

* The source Asset Database (`Library\SourceAssetDB`): Contains meta-information about your source asset files which Unity uses to determine whether the file has been modified, and therefore whether it should reimport the files. This includes information such as last modified date, a hash of the file’s contents, GUIDs, and other meta-information.
* The Artifact database (`Library\ArtifactDB`): Contains information about the import results of each source asset. Each artifact contains the import dependency information, artifact meta-information and a list of artifact files.

**Note:** Make sure to exclude these database files from **version control** systems.

## Asset caching and dependencies

The Asset Database keeps track of all the dependencies for each asset and caches artifacts. Artifacts are treated as cached data because Unity can always regenerate them from the source assets, using the import settings and **project settings** saved in your project.

An asset’s import dependencies include all the data that might influence the imported data. If an imported asset’s source file or any of its dependencies change, then Unity reimports the file and updates the imported version of the data.

For more information on what constitutes an asset’s dependencies and the refresh process, refer to [Refreshing the Asset Database](AssetDatabaseRefreshing.md).

If you’re working in a team, every team member’s copy of Unity performs the import process if an asset or dependency changes, which can be time consuming. Unity provides a solution to this called Unity **Accelerator**. For more information, refer to [Caching assets](importing-caching-assets.md).

## Asset reimport on platform change

Changing the active [target platform](build-profiles.md) can cause Unity to reimport your assets. This usually happens when the way the asset is imported differs between platforms, which is often the case. For example, different platforms have different **texture formats**, so textures are imported differently for each platform.

The results for importing your assets on different platforms are stored as separate pieces of cached data. The platform is part of the hash the Asset Database uses to store the import results for Unity’s [built-in importers](BuiltInImporters).

When you change to a different platform, any assets in your project that haven’t already been imported for the new platform are reimported. However, the new reimported data does not overwrite the old cached import data for the previous platform. When you switch back to a platform you’ve already imported assets for, those asset import results are already cached and ready to use, making the switch back faster.

## Additional resources

* [Refreshing the Asset Database](AssetDatabaseRefreshing.md)
* [Programming with the Asset Database](AssetDatabaseCustomizingWorkflow.md)