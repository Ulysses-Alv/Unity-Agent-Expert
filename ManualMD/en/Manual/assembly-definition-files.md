# Organizing scripts into assemblies

Assemblies are individual units of compiled code that group types and resources together. Organizing **scripts** into assemblies has important advantages, especially as your codebase grows.

Assemblies help you think clearly about the architecture of your code and about managing dependencies. By exercising fine-grained control over references, you can reduce unnecessary recompilation time and make your code easier to debug.

| **Topic** | **Description** |
| --- | --- |
| [Introduction to assemblies in Unity](assembly-definitions-intro.md) | Understand the fundamentals of how assemblies work in Unity and why using them to organize your scripts is beneficial. |
| [Creating assembly assets](assembly-definitions-creating.md) | Create various kinds of assembly assets to customize your assemblies. |
| [Referencing assemblies](assembly-definitions-referencing.md) | Set up references between assemblies, override the default references and understand the limitations Unity places on references. |
| [Conditionally including assemblies](assembly-definition-includes.md) | Use scripting symbols to conditionally include or exclude assemblies from compilation. |
| [Assembly metadata and compilation details](assembly-definition-metadata.md) | Define metadata for your assemblies. |
| [Assembly Definition Inspector window reference](class-AssemblyDefinitionImporter.md) | Inspector-editable properties of assembly defintion assets and their meaning. |
| [Assembly Definition Reference Inspector window reference](class-AssemblyDefinitionReferenceImporter.md) | Inspector-editable properties of assembly defintion reference assets and their meaning. |
| [Assembly Definition file format reference](assembly-definition-file-format.md) | Assembly definition file format reference. |
| [Predefined assemblies reference](script-compile-order-folders.md) | Unity’s predefined assemblies and the order in which Unity compiles them. |

## Additional resources

* [Special folders and script compilation order](script-compile-order-folders.md)
* [Scripting backends](scripting-backends.md)