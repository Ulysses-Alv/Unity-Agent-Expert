# Analyze built assets

Included in every Unity Editor installation are the low-level command-line tools WebExtract and binary2text, which you can use to inspect the contents of built Unity data. Use these tools to extract files embedded within a Unity archive, such as AssetBundles, or build outputs.

* **WebExtract**: Extracts the files in a Unity archive file to a binary format. Places the binary file in a folder named after the source file. You can use this tool to extract data such as the following:
  + [AssetBundles](AssetBundlesIntro.md) including [Addressables build](https://docs.unity3d.com/Packages/com.unity.addressables@latest) output.
  + The `data.unity3d` file from a compressed Player build.
  + Files inside the `StreamingAssets/ContentArchives` folder [for Entities builds](https://docs.unity3d.com/Packages/com.unity.entities@latest?subfolder=/manual/content-management-intro.html).
* **binary2text**: Creates a text representation of a binary `SerializedFile` (Unity’s serialized asset format). This allows you to inspect object data, references, and type information. However, this tool only contains useful information for files that contain [TypeTrees](assetbundles-optimizing.md#typetrees). By default, Player data doesn’t contain TypeTrees. Examples of binary files you can open with this tool include the following:
  + Serialized files inside archives that have been extracted using WebExtract.
  + [BuildReport](assetbundles-file-format.md#build-report) files (for example `Library/LastBuild.buildreport`).
  + Files in the `Library/Artifacts` folder, which you can inspect to understand the result of an import process.
  + Player data files including [`level` and `sharedasset` files](build-content-output.md).

## SerializedFile format

A `SerializedFile` is Unity’s serialized asset data format. It exists as either text, or binary format. You can use the [**Asset Serialization Mode** Project Setting](class-EditorManager.md#asset-serialization) to define whether Unity serializes assets in binary, text, or mixed mode (which uses the format that an asset was previously saved in, resulting in a mix of binary and text files).

The binary format is used for imported artifacts in `Library/Artifacts`, Player build outputs, AssetBundles, and other build archives. The binary2text tool can inspect this binary format.

## Use WebExtract and binary2text

WebExtract and binary2text are installed in the following locations:

| **Platform** | **Install path** |
| --- | --- |
| Windows | `C:\Program Files\Unity\Editor\Data\Tools` |
| macOS | `/Applications/Unity/Unity.app/Contents/Tools` |

### Extract archive files

To extract an archive file to binary format use the `WebExtract` tool:

1. Go to the `Tools` folder that contains `WebExtract`.
2. Run the command `WebExtract <FilePathOfArchiveFile>`, or `./WebExtract <FilePathOfArchiveFile>` if using macOS or Linux.

The tool then creates a folder that contains the binary data of the archive file.

### Create a text representation of a binary file

To create a text representation of a binary file, use the `binary2text` tool:

1. Go to the `Tools` folder that contains `binary2text`.
2. Run the command `binary2text <FilePathOfBinaryFile>`, or `./binary2text <FilePathOfBinaryFile>` if using macOS or Linux.

The tool then creates a `.txt` file in the same folder as the binary file.

**Note:** The tool has backwards compatibility, so newer versions of Unity can read files created with older versions of Unity, but older versions of Unity might not be able to read files created with newer versions.

The binary2text tool has the following optional arguments you can pass when you run it:

| **Argument** | **Description** |
| --- | --- |
| `<outputTextFile>` | Specifies a file to write the output to. If left blank, binary2text writes the output to `<BinaryFileName>.txt`. |
| `-typeinfo` | Includes the [TypeTree](assetbundles-optimizing.md#typetrees) information from the `SerializedFile` header. This includes Unity types, `MonoBehaviour`-based C# classes, and `[SerializeReference]` types, which is useful for debugging code stripping, missing types, or serialization mismatches between Unity versions. |
| `-largebinaryhashonly` | Summarizes the content with a hash instead of printing each byte. |
| `-hexfloat` | Prints floating point numbers as hexadecimal. This is more accurate than the default decimal display. |
| `-detailed` | Provides detailed output for certain data structures. |

## UnityDataTools

[UnityDataTools](https://github.com/Unity-Technologies/UnityDataTools) is a newer alternative to `WebExtract` and `Binary2Text`. It contains the following features:

* **`dump` option**: Opens an AssetBundle and produces text dumps of each `SerializedFile` it contains directly in the current working folder, without requiring a full binary extraction first. This behavior is similar to `binary2text`.
* **`analyze` option**: Populates an SQLite database with detailed information about the AssetBundle’s contents. This is faster than `binary2text` and can handle large object and asset counts that UI tools might struggle with.

For more information, refer to the [UnityDataTools repository](https://github.com/Unity-Technologies/UnityDataTools).

## Additional resources

* [Analyzing AssetBundles](assetbundles-analyze.md)
* [How Unity uses serialization](script-serialization-how-unity-uses.md)
* [Content output of a build](build-content-output.md)