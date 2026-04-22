# Integrated development environment (IDE) support

An integrated development environment (IDE) is an application that combines a range of tools for developing software, typically including a code editor, code completion, code analysis and diagnostics, running tests, and debugging. Unity supports the following C# IDEs:

* [Visual Studio](#visual-studio)
* [Visual Studio Code](#vs-code)
* [JetBrains Rider](#rider)

New installations of the Unity Editor on Windows and macOS include the default IDE for the platform by default. You can also add an IDE as a new module to an existing Unity installation. For more information, refer to [Add modules to a Unity Editor installation](https://docs.unity3d.com/hub/manual/AddModules.html).

## External script editor preference

The **External Script Editor** setting in the [Preferences window](Preferences.md) determines which IDE Unity C# script files open in for editing. Unity automatically assigns the IDE included in an installation as the default external script Editor. To change this setting, go to **Edit** > **Preferences** (macOS: **Unity** > **Settings**) > **External Tools** > **External Script Editor**.

This setting is enough if all you want to do is open, view, or make simple edits to C# source files. A full IDE experience including code analysis and debugging typically requires additional plugins or extensions, which are detailed in the following sections.

## Visual Studio (Windows)

[Visual Studio](https://visualstudio.microsoft.com/downloads/) is the default IDE for Unity on Windows. It’s recommended to always use the latest version where possible. New installations of the Unity Editor on Windows include Visual Studio Community and the other debugging requirements listed in this section by default. If you’re using a pre-existing installation of Visual Studio or the Unity Editor, you might need to install or configure some of the items manually.

### Visual Studio configuration for debugging

In addition to your installation of Visual Studio, the full IDE experience including debugging Unity C# code requires:

For your Visual Studio IDE:

* [Visual Studio Tools for Unity](https://docs.microsoft.com/en-us/visualstudio/gamedev/unity/get-started/visual-studio-tools-for-unity)

For your Unity Editor:

* Unity [Visual Studio Editor package](https://docs.unity3d.com/Packages/com.unity.ide.visualstudio@latest/) (Included in a Unity Editor installation as part of the [Engineering feature set](DeveloperToolsFeature.md))
* Visual Studio set as the [external script editor](#script-editor).

For more information on the debugging features of the Unity Editor, refer to [Debug C# code in Unity](managed-code-debugging.md).

For more information on using Visual Studio with Unity, refer to [Using Visual Studio Tools for Unity](https://docs.microsoft.com/en-us/visualstudio/gamedev/unity/get-started/using-visual-studio-tools-for-unity?pivots=windows)

### Code analyzers

Unity uses Visual Studio’s C# compiler to compile **scripts**. When you use the Visual Studio Editor package with Visual Studio, both Unity and Visual Studio display details of any errors in your scripts.

Due to differences in the way Unity and Visual Studio compiles user code, `Microsoft.Unity.Analyzers.dll` isn’t configured automatically in the Unity Editor. To configure it, follow the instructions in [Install and use an existing analyzer or source generator](install-existing-analyzer.md) using the [Microsoft.Unity.Analyzers nuget repository](https://www.nuget.org/packages/Microsoft.Unity.Analyzers) as your source.

### Project and solution files

Unity automatically creates and maintains a Visual Studio .sln and .csproj file. You can control whether Unity generates .csproj files for additional packages in your project. For more information, refer to [Using the Visual Studio Editor package](https://docs.unity3d.com/Packages/com.unity.ide.visualstudio@latest/index.html?subfolder=/manual/using-visual-studio-editor.html).

## Visual Studio Code (Windows, macOS, Linux)

[Visual Studio Code](https://code.visualstudio.com/download) is the default IDE for Unity projects on macOS. It’s recommended to always use the latest version where possible. New installations of the Unity Editor on macOS include Visual Studio Code and the other debugging requirements listed in this section by default. If you’re using a pre-existing installation of VS Code or the Unity Editor, you might need to install or configure some of the items manually.

### VS Code configuration for debugging

In addition to your installation of Visual Studio Code itself, the full IDE experience including debugging Unity C# code requires:

For your Visual Studio Code IDE:

* [C# For Visual Studio Code Extension](https://marketplace.visualstudio.com/items?itemName=ms-dotnettools.csharp)
* [C# Dev Kit for Visual Studio Code Extension](https://marketplace.visualstudio.com/items?itemName=ms-dotnettools.csdevkit)
* [Unity for Visual Studio Code Extension](https://marketplace.visualstudio.com/items?itemName=visualstudiotoolsforunity.vstuc)

For your Unity Editor:

* [Unity Visual Studio Editor package](https://docs.unity3d.com/Packages/com.unity.ide.visualstudio@latest/) 2.0.20+ (Included in a Unity Editor installation as part of the [Engineering feature set](DeveloperToolsFeature.md))
* Visual Studio Code set as the [external script editor](#script-editor).

**Note**: The Unity **Visual Studio Code** Editor package [com.unity.ide.vscode](https://docs.unity3d.com/Packages/com.unity.ide.vscode@latest) is no longer supported and should not be used. The **Visual Studio Editor** package [com.unity.ide.visualstudio](https://docs.unity3d.com/Packages/com.unity.ide.visualstudio@latest/) now supports Visual Studio Code in addition to Visual Studio.

For more information on the debugging features of the Unity Editor, refer to [Debug C# code in Unity](managed-code-debugging.md).

For information on using VS Code with Unity, refer to the VS Code documentation on [Unity Development with VS Code](https://code.visualstudio.com/docs/other/unity).

## JetBrains Rider (Windows, macOS, Linux)

[JetBrains Rider](https://www.jetbrains.com/rider/) is a C# IDE with comprehensive C# language and debugging support. It’s recommended to always use latest version where possible.

### Rider configuration for debugging

In addition to your installation of Rider itself, the full IDE experience including debugging Unity C# code requires:

For your Unity Editor:

* [Unity JetBrains Rider package](https://docs.unity3d.com/Packages/com.unity.ide.rider@latest/) (Included in a Unity Editor installation as part of the [Engineering feature set](DeveloperToolsFeature.md)).
* Rider set as the [external script editor](#script-editor).

For more information on the debugging features of the Unity Editor, refer to [Debug C# code in Unity](managed-code-debugging.md).

For more information on using JetBrains Rider with Unity, refer to [Rider for Unity](https://www.jetbrains.com/lp/dotnet-unity/).

## Additional resources

* [Debug C# code in Unity](managed-code-debugging.md)