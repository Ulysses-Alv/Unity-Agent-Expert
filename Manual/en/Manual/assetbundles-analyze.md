# Analyzing AssetBundles

Unity provides several tools to inspect the internal structure and contents of AssetBundles to help with troubleshooting.

## WebExtract and binary2text

WebExtract and binary2text are low-level command-line tools included with every Unity Editor installation.

* `WebExtract`: Extracts the files embedded within an AssetBundle, similar to unzipping an archive. It creates a folder (named after the AssetBundle) containing its internal files in the same directory as the AssetBundle.
* `binary2text`: Converts a binary `SerializedFile` (the core Unity data format) into a human-readable text format, similar to Unity’s YAML format.

You can also use [UnityDataTools](https://github.com/Unity-Technologies/UnityDataTools), which is an alternative to the WebExtract and binary2text tools.

For more information, refer to [Analyze built assets](assets-analyze-built-assets.md).

## BuildReport

The `BuildReport` file (located at `Library/LastBuild.buildreport` in your project directory) is a Unity `SerializedFile` that records the content of each AssetBundle. This `SerializedFile` contains detailed information about the build, including timings for build steps and a granular view of AssetBundle contents. You can view this file in the following ways:

* **Build Report **Inspector****: Use the [Build Report Inspector](https://docs.unity3d.com/Packages/com.unity.build-report-inspector@latest) package to view the `BuildReport` contents within the Editor.
* **Project Auditor**: [Project Auditor](project-auditor/project-auditor) can view `BuildReports`, but might struggle with large builds that involve many objects.
* **Programmatic Access**: Use [`BuildReport`](../ScriptReference/Build.Reporting.BuildReport.md) to access to the data within the `BuildReport` file.
* **Text Format**: Enable **Preferences > Diagnostics > BuildReportingEditor > SerializeBuildReportAsText** to output the `BuildReport` directly as YAML text.

## Additional resources

* [UnityDataTools repository](https://github.com/Unity-Technologies/UnityDataTools)
* [Optimizing AssetBundles](assetbundles-optimizing.md)
* [AssetBundle file format reference](assetbundles-file-format.md)
* [Build assets into AssetBundles](AssetBundles-Building.md)
* [`BuildReport` API documentation](../ScriptReference/Build.Reporting.BuildReport.md)
* [`AssetBundleManifest` API documentation](../ScriptReference/AssetBundleManifest.md)