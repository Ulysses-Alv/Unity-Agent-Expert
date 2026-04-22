# Assembly Definition Inspector window reference

Click on an Assembly Definition Asset to set the properties for an assembly in the **Inspector** window.

![The Name and General sections of configurable properties in the Assembly Definition importer Inspector window.](../uploads/Main/name-and-general.png)

The Name and General sections of configurable properties in the Assembly Definition importer Inspector window.

## Name

| **Property** | **Description** |
| --- | --- |
| **Name** | The name for the assembly (without a file extension). Assembly names must be unique across the Project. Consider using a reverse-DNS naming style to reduce the chance of name conflicts, especially if you want to use the assembly in more than one Project. For more information on .NET requirements and recommendations for assembly naming, refer to [Assembly names](https://learn.microsoft.com/en-us/dotnet/standard/assembly/names).  **Note**: Unity uses the name you assign to the Assembly Definition asset as the default value of the Name field, but you can change the name as needed. However, if you reference an Assembly Definition by its name rather than its GUID, changing the name breaks the reference. |

## General Options

| **Property** | **Description** |
| --- | --- |
| **Allow ‘unsafe’ Code** | Enable this option if you have used the C# [`unsafe`](https://learn.microsoft.com/en-us/dotnet/csharp/language-reference/keywords/unsafe) keyword in a script within the assembly. When enabled, Unity passes the `/unsafe` option to the C# compiler when it compiles the assembly. |
| **Auto Referenced** | Specify whether this assembly is automatically referenced by Unity’s [predefined assemblies](script-compile-order-folders.md). When disabled, Unity does not automatically reference the assembly during compilation. This has no effect on whether Unity includes the assembly in the build. |
| **No Engine References** | When enabled, Unity does not add references to `UnityEditor` or `UnityEngine` assemblies when it compiles the assembly. |
| **Override References** | Enable this option to manually specify which precompiled assemblies this assembly depends upon. When enabled, the Inspector shows the Assembly References section, which you can use to specify the references.  A precompiled assembly is a library compiled outside your project. By default, assemblies you define in your project reference all the precompiled assemblies you add to the project. When you enable Override References, this assembly only references the precompiled assemblies you add under [Assembly References](#assembly-references).  **Note**: To prevent project assemblies from automatically referencing a precompiled assembly, you can disable the precompiled assembly’s Auto Referenced option. For more information, refer to [Import and configure plug-ins](plug-in-inspector.md). |
| **Root Namespace** | The default namespace for **scripts** in this assembly definition. If you use either [Rider](scripting-ide-support.md#rider) or [Visual Studio](scripting-ide-support.md#visual-studio) as your code editor, they automatically add this namespace to any new scripts you create in this assembly definition. |

For more information, refer to [Create an Assembly Definition asset](assembly-definitions-creating.md#create-asmdef).

## Assembly Definition References

![The Assembly Definition References and Assembly References sections of the Assembly Definition importer Inspector window.](../uploads/Main/assembly-ref.png)

The Assembly Definition References and Assembly References sections of the Assembly Definition importer Inspector window.

| **Property** | **Description** |
| --- | --- |
| **Assembly Definition References** | A list of assemblies to reference from the current assembly. Click the **+** button to add a new reference. Click the **-** button to remove a reference. Unity uses these references to compile the assembly and define the dependencies between assemblies. |
| **Use GUIDs** | This setting controls how Unity serializes references to other Assembly Definition assets. When you enable this property, Unity saves the reference as the asset’s GUID, instead of the Assembly Definition name. It’s good practice to use the GUID instead of the name, because it means you can make changes to the name of an Assembly Definition asset without having to update other Assembly Definition files that reference it. |

For more information, refer to [Create an Assembly Definition asset](assembly-definitions-creating.md#create-asmdef).

## Assembly References

| **Property** | **Description** |
| --- | --- |
| **Assembly References** | The Assembly References section only appears when you enable the **Override References** property in the [**General Options**](#general) section. Use this section to specify any references to precompiled assemblies on which this assembly depends. For more information, refer to [Referencing a precompiled, plugin assembly](assembly-definitions-referencing.md#reference-precompiled-assembly). |

## Platforms

![The Platforms section of the Assembly Definition importer Inspector window with Any Platform selected.](../uploads/Main/platforms.png)

The Platforms section of the Assembly Definition importer Inspector window with Any Platform selected.

| **Property** | **Description** |
| --- | --- |
| **Platforms** | The **Platforms** list defines which target platforms the assembly compiles for. If you select **Any Platform**, then Unity compiles the assembly compiles for all platforms and excludes any individual platforms you select. If you deselect **Any Platform**, then Unity compiles the assembly for no platforms by default and includes any individual platforms you select. |

For more information, refer to [Creating a platform-specific assembly](assembly-definitions-creating.md#create-platform-specific).

## Define Constraints

![A list of preprocessor symbols configured in the Define Constraints section of the Assembly Definition importer Inspector window, with one constraint highlighted as currently unmet.](../uploads/Main/define-constraints.png)

A list of preprocessor symbols configured in the Define Constraints section of the Assembly Definition importer Inspector window, with one constraint highlighted as currently unmet.

| **Property** | **Description** |
| --- | --- |
| **Define Constraints** | Define constraints specify the [scripting symbols](scripting-symbol-reference.md) that must be defined in your project for Unity to compile or reference an assembly. All the listed symbols must be defined for the assembly to compile. Constraints work like the `#if` [preprocessor directive](platform-dependent-compilation.md) in C#, but on the assembly level instead of the script level.   Prefix a symbol with an exclamation (`!`) to negate it. For example, `!ENABLE_IL2CPP` specifies that the symbol `ENABLE_IL2CPP` must not be defined for the assembly to compile. Use the `||` (OR) operator to specify that at least one of the constraints must be present in order for the constraints to be satisfied. For example, `UNITY_IOS || UNITY_EDITOR_OSX` compiles the assembly when either `UNITY_IOS` or `UNITY_EDITOR_OSX` is defined. If any constraint is currently not met, Unity marks the individual constraint and the Define Constraints section with red information icons.   You can use any of Unity’s [built-in scripting symbols](scripting-symbol-reference.md) and any [custom scripting symbols](custom-scripting-symbols.md) you’ve defined. For more information, refer to [Platform dependent compilation](platform-dependent-compilation.md).  **Note**: The **Define Constraints** apply to the currently active platform in your project [Build Profiles](build-profiles.md). To define a symbol for multiple platforms, you must switch to each platform and modify the **Define Constraints** field individually. |

For more information, refer to [Conditionally including an assembly](assembly-definition-includes.md).

## Version Defines

![The Version Defines section of the Assembly Definition importer Inspector window, with defines configured for specific versions of the Unity Test Framework and VS Code packages.](../uploads/Main/version-defines.png)

The Version Defines section of the Assembly Definition importer Inspector window, with defines configured for specific versions of the Unity Test Framework and VS Code packages.

Specify which symbols to define according to the versions of the packages and modules in a project.

| **Property** | **Description** |
| --- | --- |
| **If resource** | A package or module. |
| **version is** | An expression defining a version or range of versions. |
| **set define** | The symbol to define when an applicable version of the resource is also present in this project. |
| **Version expression outcome** | The expression evaluated as a logical statement, where `x` is the version checked. If the expression outcome displays **Invalid** then the expression is malformed. |

For more information on defining these properties and the correct syntax for version expressions, refer to [Defining symbols based on project packages](assembly-definition-includes.md#define-symbols).

## Additional resources

* [Creating assembly definitions](assembly-definitions-creating.md)
* [Referencing assemblies](assembly-definitions-referencing.md)
* [Assembly Definition Reference properties](class-AssemblyDefinitionReferenceImporter.md)
* [Assembly Definition file format reference](assembly-definition-file-format.md)