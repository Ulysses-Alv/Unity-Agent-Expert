# Refreshing the Asset Database

Unity refreshes the Asset Database in the following situations:

* When the Unity Editor regains focus (if you have enabled [Auto-Refresh](Preferences.md) in the Preferences window)
* When you select **Assets > Refresh** from the menu.
* When you call [`AssetDatabase.Refresh`](../ScriptReference/AssetDatabase.Refresh.md) from code. Some other `AssetDatabase` APIs trigger a refresh, but only for the assets you specify.

## Refresh process

Unity performs the following steps during an Asset Database refresh:

1. Check if any files in the `Assets` and `Packages` folders have been added, modified, or deleted since the last check and update the Asset Database accordingly.
2. Import and compile code-related files such as .dll, .asmdef, .asmref, .rsp, and .cs files.
3. Reload the [scripting domain](domain-reloading.md), if [Refresh](../ScriptReference/AssetDatabase.Refresh.md) was not invoked from a script.
4. Post-process all of the assets for the imported code-related files.
5. Import non-code-related assets and post-process all remaining imported assets.
6. [Hot reload](#hotreloading) the assets.

### Restarting the refresh process

Unity restarts the Asset Database refresh process if:

* The import of an asset fails.
* After the import, a file that the importer used has changed on disk.
* In [`OnPostProcessAllAssets`](../ScriptReference/AssetPostprocessor.OnPostprocessAllAssets.md), you call any of the following:
  + [`ForceReserializeAssets`](../ScriptReference/AssetDatabase.ForceReserializeAssets.md)
  + [`AssetImporter.SaveAndReimport`](../ScriptReference/AssetImporter.SaveAndReimport.md)
  + Any `AssetDatabase` API that queues an additional refresh, such as MoveAsset, CreateAsset and ImportAsset.
* The last modified timestamp of the file being imported changes while it’s being imported. This can happen if you start pulling files from a **version control** system while the Editor has focus.
* An importer creates a file in the middle of an import. For example, FBX models can restart a refresh by extracting their Textures from the model.
* An assembly reload happens after compiling **scripts**. If you generate a C# file during the refresh process, that new file must then be compiled.
* You save an asset as **Text only** but the asset must be serialized as binary. For example, **Scenes** with Terrains in them must be serialized as Binary, since the **terrain** data would be unwieldy if viewed as an array of characters in a text file.

## Dependency tracking

The Asset Database keeps track of two types of Asset dependencies: **static dependencies** and **dynamic dependencies**. If any dependency of an asset changes, then Unity reimports the asset.

### Static dependencies

A static dependency is a value, setting, or property that an importer depends on. Static dependencies are known before the asset is imported, and are not affected by the behavior of the importer during the import process. If any of an asset’s static dependencies change, then Unity re-imports the asset.

Static dependencies include:

* The name of the asset.
* ID of the importer associated with the asset.
* The version of the importer.
* The currently selected build target platform.

### Dynamic dependencies

Dynamic dependencies are defined by the content of the source asset and are usually discovered during the import process. For example, a **shader** might reference another shader, and a **prefab** might depend on other prefabs.

The importer might also use a global state conditionally based on the content of the source asset, in which case it also becomes a dynamic dependency. Examples of this are the target platform, the project’s color space, the graphics API, the scripting runtime version, or the Texture **compression** state.

Unity stores dynamic dependencies of an asset in an [`AssetImportContext`](../ScriptReference/AssetImporters.AssetImportContext.md).

## Import and compile code-related files

In the list of changed or added files, Unity gathers the ones that relate to code, and sends them to the script compilation pipeline. The compiler generates assemblies from the script files and assembly definition files in your project. For more information on this step, refer to [Organizing scripts into assemblies](assembly-definition-files.md).

## Import non-code-related assets

Each asset’s importer processes that type of asset, and identifies files to import based on the file extensions. For example, the TextureImporter is responsible for importing .jpg, .png and .psd files, among others.

There are two types of importers:

* Native importers: Unity’s [built-in importers](assets-supported-types.md#asset-importers).
* [Scripted importers](ScriptedImporters.md): Custom importers that extend Unity’s import capabilities. These are written in C#.

Unity processes all native importers first, and then all scripted importers in a separate phase.

## Preprocessing and postprocessing

The `AssetPostprocessor` class provides a range of callbacks which allow you to perform work before (preprocessing) and after (postprocessing) import for different asset types. For the details of these callbacks and usage examples, refer to the [`AssetPostprocessor`](../ScriptReference/AssetPostprocessor.OnPreprocessAsset.md) API reference.

## Hot reloading

Hot reloading is the process of applying changes to scripts and assets while the Editor is open. This might happen while the Editor is in Play mode or Edit mode and requires no restart of the Editor for changes to take effect.

When you change and save a script, Unity hot-reloads all of the project’s script data. It first stores the values of [serializable variables](script-serialization-rules.md) in all loaded scripts, reloads the scripts, then restores the values. Data stored in non-serializable variables is lost.

This affects all Editor windows and all MonoBehaviour scripts in the project. Unlike other cases of serialization, Unity serializes private fields by default when reloading, even if they don’t have the [SerializeField](../ScriptReference/SerializeField.md) attribute.

## Additional resources

* [Contents of the Asset Database](asset-database-contents.md)
* [Programming with the Asset Database](AssetDatabaseCustomizingWorkflow.md)