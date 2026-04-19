# Import and configure plug-ins

If you have a [managed plug-in](plug-ins-managed.md) or a [native plug-in](plug-ins-native.md), you can import it into Unity, then configure it. Within the Editor, a **plug-in** is treated as an asset, similar to a script, and you can configure it in an Inspector window.

You can use the plug-in configurations to specify where a plug-in runs; which platforms and which platform configurations, as well as under which conditions

## Import a plug-in

The simplest way to import a plug-in to your project is to click and drag the plug-in to the Assets folder or one of its subfolders. Unity recognizes specific file and folder types as plug-ins. It can also apply default settings that match the plug-in’s intended platform.

**Note**: When you import a plug-in, Unity loads it into memory. A native plug-in can’t be unloaded; it remains loaded in a Unity session even after you change its settings. To unload the plug-in, you must restart Unity.

### Supported plug-in file and folder types

Unity treats files with the following extensions as a plug-in:

* .a
* .aar
* .bc
* .c
* .cc
* .cpp
* .dll
* .def
* .dylib
* .h
* .jar
* .jslib
* .jspre
* .m
* .mm
* .prx
* .rpl
* .so
* .sprx
* .suprx
* .swift
* .winmd
* .xcframework
* .xex
* .xib

Unity also treats certain folders as bundled plug-ins. Unity doesn’t look for additional plug-in files within these folders, so everything within the folder is considered a single plug-in. Unity treats folders with the following extensions as a bundled plug-in:

* .androidlib
* .bundle
* .framework
* .plugin

### Plug-in default settings

Unity automatically applies platform-specific default settings to the plug-in if the plug-in’s path within the Assets folder matches a platform-specific pattern. If the path doesn’t match any pattern, Unity applies the Editor platform default settings to the plug-in.

The following table shows the path patterns Unity recognizes. Portions of the path that appear in brackets are optional. When the path includes double dots, it can include more folders.

| **Folder path patterns** | **Default settings** |
| --- | --- |
| **Assets/../Editor/(x86** or **x86\_64 or x64)** | **Platform**: Editor only  **CPU** (optional): Based on the subfolder, if present. |
| **Assets/../Plugins/(x86\_64** or **x86 or x64)** | **Platform**: Windows, Linux and macOS  **CPU** (optional): Based on the subfolder, if present. |
| **Assets/Plugins/iOS** | **Platform**: iOS |
| **Assets/Plugins/WSA/(SDK80** or **SDK81** or **PhoneSDK81)/(x86** or **ARM)** | **Platform**: Universal Windows Platform  **SDK** (optional): Based on the subfolder, if present. For compatibility reasons, SDK81 is Win81, PhoneSDK81 is WindowsPhone81.  **CPU** (optional): Based on the subfolder, if present  **Note**: You can use the keyword **Metro** instead of **WSA**. |

## Change plug-in settings

In Unity, plug-ins are either **managed** or **native**. The following table shows which settings are relevant for each type of plug-in:

| **Setting** | **Managed** | **Native** |
| --- | --- | --- |
| **Select platforms for plugin** | x | x |
| **Platform settings** | x | x |
| **Asset Labels** | x | x |
| **Asset Bundles** | x | x |
| **General** | x |  |
| **Define Constraints** | x |  |
| **Plugin load settings** |  | x |

To view and change plug-in settings in the **Inspector**, select the plug-in file in the **Project** window.

![Settings for a managed plug-in (left) and a native plug-in (right).](../uploads/Main/PlugInsCombined.png)

Settings for a managed plug-in (left) and a native plug-in (right).

## Common plug-in settings

**Select platform for plugin** and **Platform settings** specify in which builds Unity includes the plug-in.

### Select platform for plugin

| **Platform** | **Description** |
| --- | --- |
| **Any Platform** | Allows the plug-in to work with any platforms, including those that aren’t installed yet. If selected, this option changes the following individual platform selections from include to exclude logic. |
| **Editor** | For Play mode and for any **scripts** that run in Edit mode. |
| **Standalone** | Windows, Linux and macOS. |
| Any installed platform. | Any platform included in your Unity installation, such as Android, iOS and Web. |

### Platform settings

For each selected platform, you can specify additional platform-specific settings such as CPU architecture and dependencies. Unity shows only settings that apply to your platforms and, where possible, to the specific plug-in type on that platform. For example, a native plug-in file with a .dll extension can only run on Windows, so Unity shows only Windows settings.

| **Platform** | **Options** | **Description** |
| --- | --- | --- |
| Editor | * CPU architecture * OS | Most managed plug-ins are compatible with any CPU and OS.  Most native plug-ins are only compatible with a single OS and, depending on how they were compiled, might be compatible with only a single CPU architecture. |
| Windows, Linux and macOS | * CPU architecture * OS | Managed libraries are typically compatible with any OS and CPU architecture, unless they access specific system APIs.  Native libraries are only compatible with a single OS, but can be compatible with the 32-bit, the 64-bit, or both CPU architectures. |
| Universal Windows Platform | * SDK * CPU architecture * Scripting backend | For more information, refer to [Use UWP plug-ins with IL2CPP](uwp-il2cpp-plugins.md). |
| Android | CPU architecture | The CPU architecture must match the architecture that the library was compiled for. Unity does not validate your settings.  See also: [AAR plug-ins and Android Libraries](AndroidAARPlugins.md). |
| iOS and tvOS | * Framework dependencies * CPU architecture * Add to Embedded Binaries * Compile flags | Specify the CPU architecture your plug-in is compatible with. Unity provides ARM64 and X64 Simulator architectures for testing purposes.   When you select **Add to Embedded Binaries** option, Unity sets the Xcode project options to copy the plug-in file into the final application package. Do this for: • Dynamically loaded libraries. • Bundles and frameworks containing dynamically loaded libraries. • Any assets and resources that need to be loaded at run time.  In the **Compile Flags** field, set the compile flags for plug-in source code files that Unity must compile as part of the build. |

**Tip**: For information on the other common settings, see [Asset Bundles](AssetBundlesIntro.md) and [Searching in the Editor](Searching.md).

### Managed plug-in settings

Managed plug-ins can be third-party libraries or user-compiled assemblies that you want to include in the project.

### General - Auto Reference

The **Auto Reference** setting controls how assembly definitions in the project reference a plug-in file. When you enable Auto Reference, all predefined assemblies and assembly definitions automatically reference the plug-in file.

**Auto Reference** is enabled by default.

To limit the scope in which a plug-in can be referenced, disable Auto Reference. You then need to explicitly declare all references to that plug-in. You might want to do this if:

* You want to prevent scripts from mistakenly using the plug-in.
* You are iterating on the plug-in and want to reduce compilation time. If you explicitly declare a plug-in, Unity recompiles only the dependent assemblies, not your entire project.
* You want to prevent plug-ins used in an [Asset Store package](AssetStorePackages.md) from conflicting with other code in a project into which the package is imported.

When you disable **Auto Reference** Unity cannot reference a plug-in from the predefined assemblies it creates for your project. These predefined assemblies contain all the scripts in your project that you have not assigned to another assembly using an assembly definition file. To reference classes, functions, or other resources from a plug-in that has the **Auto Reference** property disabled, the referencing code must be in an assembly created with an assembly definition file. For example, if a set of scripts in your project uses a plug-in, you must create an assembly definition file for those scripts, and add an explicit reference to the plug-in in the definition file.

More than one assembly can use the plug-in, but all assemblies must explicitly declare the dependency. To learn more about assembly definitions in Unity, see [Assembly definitions](assembly-definition-files.md).

**Note**: The **Auto Reference** option has no effect on whether a file is included in the build. To control build settings for plug-ins use [Platform settings](#platform).

### General - Validate References

Unity can check that your plug-in’s references are available in the project. If you don’t perform this validation, users can encounter runtime errors when the application tries to use a missing reference.

Enable the **Validate References** option to check:

* Whether the plug-in’s references exist. For example, if your plug-in references the plug-in Newtonsoft.Json.dll, and Unity cannot find Newtonsoft.Json.dll, Unity displays an error.
* Whether strong named references will be able to load. This is important because strong named references need to match the version. For example, if you compile a plug-in with a reference to “b.dll” version 2.0.0, that version must be in the project.

If you don’t want to check strong named references, but still want to check that references exist:

1. In the **Plugin** inspector, enable **Validate References**.
2. In **Project Settings** > **Player** > **Other Settings** > disable **Assembly Version Validation**.

### Define Constraints

You can specify conditions under which Unity loads the plug-in to memory and references it. These conditions are symbols that must be satisfied, which means either defined or undefined.

Constraints work like the #if preprocessor directive in C#, but on the assembly level instead of the script level. You can learn more about constraints in [Assembly Definition properties](class-AssemblyDefinitionImporter.md).

You can use any of Unity’s built-in define symbols, or add symbols in **Project Settings** > **Player** > **Other Settings** > **Script Compilation** > **Scripting Define Symbols**. The symbols you add are platform-specific, so you need to define them for each relevant platform. See [Platform dependent compilation](platform-dependent-compilation.md) for more information, including a list of the built-in symbols.

**Tip**: To specify that a symbol must be undefined, prefix it with a negating ! (exclamation mark) symbol.

In the following example, we want Unity to load and reference the plug-in only on non-IL2CPP scripting runtimes for Unity 2018.3 or newer. We define two constraints, and both must be met:

* `ENABLE_IL2CPP` is not defined
* `UNITY_2018_3_OR_NEWER` is defined

![Define Constraints for a specific runtime and Unity version](../uploads/Main/Define_Constraints.png)

Define Constraints for a specific runtime and Unity version

## Plugin load settings - Load on startup

You can start executing native code that is independent of graphics initialization, scripts, asset loading, **scenes**, and so on. This is different to the default way a player loads a native plug-in, which is to wait until the first call to one of the plug-in’s functions, usually performed by a script.

To load plug-ins before the application executes any scripts:

1. Implement `UnityPluginLoad` in the plug-in. See [Low-level native plug-in interface](native-plugin-interface.md).
2. In the Editor, select **Plugin load settings** > **Load on startup**.

**Tip**: For an example of a C# script calling plug-in functions, see [Manual: Native plug-ins](plug-ins-native.md).

## Additional resources

* [Managed plug-ins](plug-ins-managed.md)
* [Native plug-ins](plug-ins-native.md)
* [Plugins for desktop](plug-ins-for-desktop.md)
* [Native plug-in interface](native-plugin-interface.md)