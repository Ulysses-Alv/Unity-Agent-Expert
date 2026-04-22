# Assembly Definition file format reference

Assembly Definition and Assembly Definition Reference assets are JSON files. You can edit the asset files inside the Unity Editor using the **Inspector** window, but you can also modify the JSON content directly with an external tool.

## Assembly Definition JSON

An Assembly Definition is a JSON object with the following fields:

| Key | Type | Required | Description |
| --- | --- | --- | --- |
| `allowUnsafeCode` | bool | Optional | False by default. For more information on the meaning of this option, refer to **Allow ‘unsafe’ Code** in [General Options](class-AssemblyDefinitionImporter.md#general). |
| `autoReferenced` | bool | Optional | True by default. For more information on the meaning of this option, refer to **Auto Referenced** in [General Options](class-AssemblyDefinitionImporter.md#general). |
| `defineConstraints` | string[] | Optional | Can be empty. For more information, refer to [Define Constraints](class-AssemblyDefinitionImporter.md#define-constraints). |
| `excludePlatforms` | string[] | Optional | The platform name strings to exclude or an empty array. The `excludePlatforms` array must be empty if `includePlatforms` contains values. Valid platform name strings are the [`AssemblyDefinitionPlatform.Name`](Scriptref:Compilation.AssemblyDefinitionPlatform.Name) properties of the array of [`AssemblyDefinitionPlatform`](Scriptref:Compilation.AssemblyDefinitionPlatform) objects returned by [`CompilationPipeline.GetAssemblyDefinitionPlatforms`](../ScriptReference/CompilationPipeline.GetAssemblyDefinitionPlatforms.md). Only platforms with build support installed for the current Editor are valid. For more information, refer to [Platforms](class-AssemblyDefinitionImporter.md#platforms). |
| `includePlatforms` | string[] | Optional | The platform name strings to include or an empty array. The `includePlatforms` array must be empty if `excludePlatforms` contains values. Valid platform name strings are the [`AssemblyDefinitionPlatform.Name`](Scriptref:Compilation.AssemblyDefinitionPlatform.Name) properties of the array of [`AssemblyDefinitionPlatform`](Scriptref:Compilation.AssemblyDefinitionPlatform) objects returned by [`CompilationPipeline.GetAssemblyDefinitionPlatforms`](../ScriptReference/CompilationPipeline.GetAssemblyDefinitionPlatforms.md). Only platforms with build support installed for the current Editor are valid. For more information, refer to [Platforms](class-AssemblyDefinitionImporter.md#platforms). |
| `name` | string | Required | For more information, refer to [Name](class-AssemblyDefinitionImporter.md#name). |
| `noEngineReferences` | bool | Optional | False by default. For more information on the meaning of this option, refer to **No Engine References** in [General Options](class-AssemblyDefinitionImporter.md#general). |
| `overrideReferences` | bool | Optional | False by default. Set to true if `precompiledReferences` contains values. For more information on the meaning of this option, refer to **Override References** in [General Options](class-AssemblyDefinitionImporter.md#general). |
| `precompiledReferences` | string[] | Optional | The file names of referenced DLL libraries including extension, but without other path elements. Can be empty. This array is ignored unless you set `overrideReferences` to true. For more information, refer to [Referencing assemblies](assembly-definitions-referencing.md). |
| `references` | string[] | Optional | References to other assemblies created with Assembly Definition assets. You can use either the GUID of the Assembly Definition asset file or the name of the assembly as defined by the `name` field of the Assembly Definition. You must use the same form for all references in the list. Can be empty.  You can use the [`AssetDatabase.AssetPathToGUID`](../ScriptReference/AssetDatabase.AssetPathToGUID.md) function to retrieve the GUID of an asset. The GUID is also part of the metadata associated with every asset.   Note that the Editor displays a **Use GUIDs** option in the Assembly Definition Inspector. This option is not serialized in the associated JSON file. Instead, the choice is inferred from the form of reference found in the file. For more information, refer to [Referencing assemblies](assembly-definitions-referencing.md). |
| `versionDefines` | object[] | Optional | Contains an object for each version define. This object has three fields:  * `name`:string – The name of the resource. * `expression`:string – The expression defining the version or range of versions of the resource. * `define`:string – The symbol to define.  For more information, refer to [Version Defines](class-AssemblyDefinitionImporter.md#version-defines). |

### Example with assembly names and included platforms

The following example Assembly Defintion JSON uses assembly names for references to other Assembly Definitions and the `includePlatforms` key to specify an array of included platforms:

```
{
    "name": "BeeAssembly",
    "references": [
        "Unity.CollabProxy.Editor",
        "AssemblyB",
        "UnityEngine.UI",
        "UnityEngine.TestRunner",
        "UnityEditor.TestRunner"
    ],
    "includePlatforms": [
        "Android",
        "LinuxStandalone64",
        "WebGL"
    ],
    "excludePlatforms": [],
    "overrideReferences": true,
    "precompiledReferences": [
        "Newtonsoft.Json.dll",
        "nunit.framework.dll"
    ],
    "autoReferenced": false,
    "defineConstraints": [
        "UNITY_2019",
        "UNITY_INCLUDE_TESTS"
    ],
    "versionDefines": [
        {
            "name": "com.unity.ide.vscode",
            "expression": "[1.7,2.4.1]",
            "define": "MY_SYMBOL"
        },
        {
            "name": "com.unity.test-framework",
            "expression": "[2.7.2-preview.8]",
            "define": "TESTS"
        }
    ],
    "noEngineReferences": false
}
```

### Example with GUIDs and excluded platforms

The following example Assembly Defintion JSON uses assembly GUIDs for references to other Assembly Definitions and the `excludePlatforms` key to specify an array of excluded platforms:

```
{
    "name": "BeeAssembly",
    "references": [
        "GUID:17b36165d09634a48bf5a0e4bb27f4bd",
        "GUID:b470eee7144904e59a1064b70fa1b086",
        "GUID:2bafac87e7f4b9b418d9448d219b01ab",
        "GUID:27619889b8ba8c24980f49ee34dbb44a",
        "GUID:0acc523941302664db1f4e527237feb3"
    ],
    "includePlatforms": [],
    "excludePlatforms": [
        "iOS",
        "macOSStandalone",
        "tvOS"
    ],
    "allowUnsafeCode": false,
    "overrideReferences": true,
    "precompiledReferences": [
        "Newtonsoft.Json.dll",
        "nunit.framework.dll"
    ],
    "autoReferenced": false,
    "defineConstraints": [
        "UNITY_2019",
        "UNITY_INCLUDE_TESTS"
    ],
    "versionDefines": [
        {
            "name": "com.unity.ide.vscode",
            "expression": "[1.7,2.4.1]",
            "define": "MY_SYMBOL"
        },
        {
            "name": "com.unity.test-framework",
            "expression": "[2.7.2-preview.8]",
            "define": "TESTS"
        }
    ],
    "noEngineReferences": false
}
```

## Assembly Definition Reference JSON

An Assembly Definition Reference is a JSON object with only one required field: a string called `reference`. It specifies the assembly to reference, either by assembly name or by Assembly Definition asset GUID. You can use the [`AssetDatabase.AssetPathToGUID`](../ScriptReference/AssetDatabase.AssetPathToGUID.md) function to retrieve the GUID of an asset. The GUID is also part of the metadata associated with every asset.

The following Assembly Definition Reference example references another asset by name:

```
{
    "reference": "AssemblyA"
}
```

The following Assembly Definition Reference example references another asset by asset GUID:

```
{
    "reference": "GUID:f4de40948f4904ecb94b59dd38aab8a1"
}
```

## Additional resources

* [Create an Assembly Definition](assembly-definitions-creating.md#create-asmdef)
* [Create an Assembly Definition Reference](assembly-definitions-creating.md#create-asmref)