# Check the consistency of the import process

Asset importers, including any [scripted importers](ScriptedImporters.md) that you write, should produce consistent, or [deterministic](https://learn.microsoft.com/en-us/dotnet/csharp/language-reference/compiler-options/code-generation#deterministic) results. A deterministic importer always produces the same output from the same input and set of dependencies.

To verify if an importer is deterministic, the [Asset Database](AssetDatabase.md) has the following ways that you can check the consistency of the import results of the assets currently in your project:

* [Manually reimport](#manually-reimport-assets) one or more assets in the Unity Editor.
* [Use the `-consistencyCheck` command line argument](#run-consistency-checks-with-the-command-line) to open the Editor.

## Manually reimport assets

Manually reimporting an asset causes Unity to check whether a new import result matches the previous cached import results.

To start a manual reimport, select an asset and perform one of the following actions:

* Right-click on the asset and select **Reimport** from the context menu.
* Go to **Assets** > **Reimport**.

Unity then checks the resulting import for consistency. If it detects inconsistent results, Unity prints a warning in the [Console window](Console.md) with details about which importer caused the issue, such as:

`Importer(<name of importer>) generated inconsistent result for asset(guid:<guid of the asset>) "<name of asset>"`

When you reimport a root asset, Unity also reimports and performs the consistency check on its child assets.

## Run consistency checks with the command-line

If you open Unity with the command-line, you can use the `-consistencyCheck` command line argument to perform a check on all assets and importers in the project. Unity then logs any inconsistencies found during startup to the Console, and the Editor log file.

You can also use the `-consistencyCheckSourceMode` argument to check against a different source, or a [Unity Accelerator cache server](UnityAccelerator.md). **Important:** If you use Unity **Accelerator** as the source, Unity disables uploads and downloads from the cache server to test the import process. For more information, refer to [Unity Editor command line arguments](EditorCommandLineArguments.md#configuration-arguments).

The consistency checker only checks asset imports that can be cached. For example, if you disable caching for a `ScriptedImporter` instance, then the consistency check for that import is also disabled. For more information, refer to [`AssetImporters.ScriptedImporterAttribute`](../ScriptReference/AssetImporters.ScriptedImporterAttribute.md) for details on how to disable caching for a scripted importer.

### Editor log output

When Unity detects importer inconsistencies, you can find more information about the issue in the Editor log. This includes information such as:

* Which asset caused the failed consistency check and its GUID.
* The content hash for the asset import result.
* A comparison of the previous revision of the asset.

The following is an example of an Editor log that shows importer inconsistencies:

```
ConsistencyChecker - guid: a1945cd7aab67441ba89015f97494624, dependenciesHash.value: fb8cfb407bba82d4daded6031688ba9b, artifactid: 07078a054d3f597b4c2cc47e8e4c0bde, producedFiles[0].extension: , producedFiles[0].contentHash: 8490a5ed35a4361d679e6055a386969e

ConsistencyChecker - guid: a1945cd7aab67441ba89015f97494624, dependenciesHash.value: fb8cfb407bba82d4daded6031688ba9b, artifactid: 07078a054d3f597b4c2cc47e8e4c0bde, producedFiles[1].extension: .info, producedFiles[1].contentHash: 4bd9140e19d2e44782f1131172e514ba

ConsistencyChecker - guid: a1945cd7aab67441ba89015f97494624, dependenciesHash.value: fb8cfb407bba82d4daded6031688ba9b, artifactid: 646d6432767729ea7d288c636183de97, producedFiles[0].extension: , producedFiles[0].contentHash: 5ca760170f85012ce16aa8c22e8d9ea1

ConsistencyChecker - guid: a1945cd7aab67441ba89015f97494624, dependenciesHash.value: fb8cfb407bba82d4daded6031688ba9b, artifactid: 646d6432767729ea7d288c636183de97, producedFiles[1].extension: .info, producedFiles[1].contentHash: 4bd9140e19d2e44782f1131172e514ba

Importer(ScriptedImporter:Assembly-CSharp::RandomImporter) generated inconsistent result for asset(guid:a1945cd7aab67441ba89015f97494624) "Assets/first.rand"
```

If [Unity Accelerator](UnityAccelerator.md) isn’t available, the consistency checker prints out warnings in the Editor log. Additionally, if you haven’t set up Unity Accelerator, Unity outputs the following message:

`ConsistencyChecker - Cacheserver is not enabled`

If Unity can’t connect to the cache server, it outputs the following message:

`ConsistencyChecker - Not connected to accelerator/cacheserver`

## Debugging an importer

You can use the `binary2text` tool to examine the contents of the `Library` folder to find out exactly what the importer generated. The `binary2text` tool is located in the `Data/Tools` folder in your Unity Editor installation. On macOS, this tool is located at `Unity.app/Contents/Helpers/`.

For example:

`./Data/Tools/binary2text ./myProject/Library/Artifacts/84/8490a5ed35a4361d679e6055a386969e output_file`

The hash value is the asset’s content hash. To find the reported inconsistency and what content was generated during the check, refer to the [Editor log](log-files.md).

## Additional resources

* [Introduction to Unity Accelerator](UnityAccelerator.md)
* [Log files reference](log-files.md)
* [Command line arguments](CommandLineArguments.md)