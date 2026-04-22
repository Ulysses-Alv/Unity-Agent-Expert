# Safe Mode

## Overview

Unity’s Safe Mode is a mode that the Unity Editor can enter when you open a project that has script compilation errors. Safe Mode is designed to provide the best environment for resolving compilation errors, so that you can quickly return your project to a functional state.

In Safe Mode, Unity provides a minimal version of the Editor user interface, with limited functions. It only imports script-related assets, and prevents the import of non-script assets (such as models, materials, textures and prefabs). This is because Safe Mode is not meant for content production, only for resolving compilation errors.

Safe Mode never allows managed code to run from your project, or its packages. This means your own **scripts** such as [Editor scripts](ExtendingTheEditor.md), [asset post-processors](../ScriptReference/AssetPostprocessor.md), and [scripted importers](ScriptedImporters.md) do not run. Safe Mode also disables [assembly overrides](class-AssemblyDefinitionImporter.md#assembly-references), and the [Burst](com.unity.burst.md) and [Roslyn](roslyn-analyzers.md) Analysers. This ensures the Editor in Safe Mode is always fully functional and reliable, even when it opens a project in a very broken state.

Unity automatically exits Safe Mode when it detects there are no more compilation errors. When it exits Safe Mode, Unity imports your project in full, and the Editor restores its normal full functionality.

## Common causes of compilation errors

Some common scenarios that might result in you encountering compilation errors occur when you:

* Upgrade a project from an older version of Unity to a newer version.
* Open a project in a different version of Unity from the one it was created in.
* Open a project with a missing [package](PackagesList.md), or an incorrect version of a package.
* Open a project which has errors in your own scripts.
* Open a project under [version control](VersionControl.md), where the latest changes you pulled contained compilation errors.

Opening a project that has compilation errors without Safe Mode can cause many kinds of problems. For example, [packages](PackagesList.md) in your project might not load or function properly, and your assets might be imported incorrectly, resulting in incorrect cached artifacts in your [Library](ImportingAssets.md#asset-processing) or in your [Cache Server](UnityAccelerator.md).

In these situations, you usually don’t want to wait for the rest of the project to import before you can resolve the errors. Safe Mode provides you with the tools to resolve these script-related problems yourself, or to use version control to update your project to a newer version that resolves the errors, without waiting for a full import of the project.

## Entering Safe Mode

When you open a project which has compilation errors, the Editor displays a dialog to ask whether you want to enter Safe Mode:

![The Enter Safe Mode? dialog prompts you to enter Safe Mode when you open a project with compilation errors](../uploads/Main/SafeModeDialog.png)

The Enter Safe Mode? dialog prompts you to enter Safe Mode when you open a project with compilation errors

At this point you have three choices:

* **Enter Safe Mode**
* **Ignore** the errors and open your project
* **Quit** Unity

In most cases, you should select **Enter Safe Mode** to resolve the errors in your project (or, if you’re working with version control, to pull changes which contain fixes to the errors). Safe Mode provides the best environment for resolving compilation errors, so that you can quickly return your project to a functional state before Unity imports the rest of your project.

However, there are some cases where you might not want to enter Safe Mode, in which case you can **Quit** Unity, or **Ignore** the errors.

**Note**: *You can disable this dialog in ****Edit > Preferences > Asset Pipeline > Show Enter Safe Mode Dialog****. If you disable the dialog, Unity automatically enters Safe Mode when it opens a project with compilation errors.*

### Quitting without entering Safe Mode

Safe Mode is specifically designed for fixing compilation errors. If you are working on a Unity project as part of a team, but you are not responsible for the scripts that are causing errors, and don’t know what to do, you should select **Quit** in the dialog, and contact the programmers on your team for advice.

### Ignoring the errors and continuing import

There are some situations where you don’t need the project to be in a usable state (for example, if you are opening an old project to copy some parts out, or to simply inspect how it is configured). In this case, you can ignore the errors and open your project in a broken state anyway.

If you select **Ignore**, and later want to open the project in Safe Mode, you can close and re-open Unity to access the “Enter Safe Mode” dialog again.

### Implications of ignoring compilation errors

If you choose to ignore the errors, Unity continues to import the rest of your assets and opens your project fully. Potential implications here include:

* Your project might not be in a usable state. You might not be able to enter Play Mode, or create builds of your project, until the errors are resolved. In addition, any packages in your project might not load correctly, or at all.
* Unity might need to import your assets twice; once at launch, and again after you resolve the project’s compilation errors. This increases the amount of time it takes to load your project in a usable state.
* If your project uses a [Scriptable Render Pipeline](scriptable-render-pipeline-introduction.md), your **render pipeline** might not load, resulting in visual problems such as the [error shader](shader-error.md).
* Script compilation errors can cause secondary errors within your project. For example, if a [scripted importer](ScriptedImporters.md) in your project is unable to load because of compilation errors, your assets might import in an incorrect state.

Safe Mode is designed to help you avoid all of these problems.

## The Editor in Safe Mode

In Safe Mode, Unity provides a minimal Editor interface with limited functions.

![The Unity Editor in Safe Mode](../uploads/Main/SafeModeEditor.png)

The Unity Editor in Safe Mode

The Unity Editor displays a Safe Mode banner in the **toolbar** section at the top of the Editor, which replaces the standard Editor toolbar. The toolbar indicates that you are in Safe Mode, and provides an **Exit Safe Mode** button, which allows you to ignore the remaining errors and exit Safe Mode. The banner also indicates whether your project uses [preview packages](pack-preview.md).

The Unity Editor retains its [integration with code editors](scripting-ide-support.md) in Safe Mode, so you can double-click on script assets or [console error](Console.md)[s](Console.md) to open their associated scripts, and you can open the C# project via the **Assets** menu. It also retains its integration with [Version Control systems](VersionControl.md).

### The restricted set of windows in Safe Mode

The Editor displays a limited selection of windows in Safe Mode. These are:

* The Console window
* The Project window
* The **Inspector** window
* The Package Manager window

These are the only windows available because they relate to fixing compilation errors. No other windows are available in Safe Mode.

### The Editor Menu in Safe Mode

The options available in the Editor’s main menu are restricted to a limited set while in Safe Mode. You can only see and select the menu options which relate to working with scripts; the regular options for creating and working with other content are not available. For example, you cannot create or open **Scenes**, or create non-script assets such as primitive shapes, lights, or **cameras**. The ****GameObject**** and **Component** menus are not present, and the **Window** menu only offers a restricted set of windows.

![The restricted Editor Menu as it appears in Safe Mode, displaying the available options for creating assets.](../uploads/Main/SafeModeLimitedMenu.png)

The restricted Editor Menu as it appears in Safe Mode, displaying the available options for creating assets.

### The Project window in Safe Mode

In Safe Mode, the **Project window** has some functional differences compared with its normal behavior outside of Safe Mode.

The primary difference is that you can only select compilation-related assets. Other asset types are not selectable. They remain visible as grayed-out entries in the Project window, but you cannot select or edit them.

Specifically, the compilation-related asset types that you can interact with are:

* C# files (.cs)
* DLL files (.dll)
* Assembly definitions (.asm)
* Response files (.rsp)
* Ruleset files (.ruleset)

In addition, the icons for the non-selectable assets do not show a preview of the content of the asset. Instead, they show a generic icon to represent the asset’s type.

![The Project window displays generic icons for non-selectable assets.](../uploads/Main/SafeModeProjectWindow.png)

The Project window displays generic icons for non-selectable assets.

The **create (+)** menu button is disabled, and the Project window context menu has a reduced set of options available.

![The create (+) menu button in the project window, disabled in Safe Mode](../uploads/Main/SafeModeCreateButton.png)

The ****create (+)**** menu button in the project window, disabled in Safe Mode

## Exiting Safe Mode

Unity automatically exits Safe Mode when you have resolved all the compilation errors. Unity then continues to open your project and import your assets.

To exit Safe Mode while there are still compilation errors remaining, select the **Exit Safe Mode** button in the Safe Mode toolbar. This is not recommended (see [Implications of ignoring compilation errors](SafeMode.md#IgnoringErrors)), and Unity prompts you with a dialog to confirm your decision.

![The Exit Safe Mode button in the Safe Mode toolbar](../uploads/Main/SafeModeExitButton.png)

The **Exit Safe Mode** button in the Safe Mode toolbar

If you exit Safe Mode while errors remain in your project, and later want to return to Safe Mode, you can close and re-open Unity to access the “Enter Safe Mode” dialog again..

## Safe Mode in Batch Mode

In Batch Mode, Unity automatically quits if there are compilation errors in your project, unless you use the `-ignoreCompilerErrors` [command line argument](CommandLineArguments.md).