# Introduction to programming in Unity

Unity is customizable and extensible by design and almost everything is scriptable to some extent. Many items you can configure through the various [Editor views](unity-editor.md) have a corresponding public C# class representation that you can interact with in code.

You can use Editor APIs to customize and extend the Editor authoring tools to improve your development workflows. You can use Engine APIs to define the runtime functionality of your application, including graphics, physics, character behavior, and responses to user input.

The [Scripting API reference](../ScriptReference/index.md) provides the complete and authoritative reference for all public Unity APIs. The Manual provides additional context and guidance.

## The Unity scripting environment

Unity supports scripting in the C# programming language. C# (pronounced C-sharp) is a managed, object-oriented programming language, which is part of the .NET platform and runs in the cross-platform .NET runtime. Other .NET languages can be used with Unity if they can compile a compatible DLL, refer to [Managed plugins](plug-ins-managed.md) for further details.

The scripting environment refers to both:

* Your own local environment or context in which you’re writing code. This includes your code editor (IDE) and integrated source control solution and operating system.
* The C# scripting enviroment Unity provides. A given version of Unity supports given versions of the .NET platform, which determines the .NET libraries you can use in your code.

For more information on the scripting environment and tools, refer to [Environment and tools](environment-and-tools.md).

## How scripting in Unity works

C# scripts (files with a `.cs` file extension) are [assets](AssetWorkflow.md) in your project, stored in the `Assets` folder and saved as part of the [asset database](AssetDatabase.md). You can create template scripts that derive from the common [built-in Unity types](fundamental-unity-types.md) through the **Scripting** submenu of the **Assets > Create** menu.

You configure a default [External Script Editor](Preferences.md#external-tools), which is the program Unity opens your script assets in for editing. Usually this will be one of the [supported IDEs](scripting-ide-support.md) for Unity development.

You can create your own regular C# types and logic to use in your game, as long as the code you write is compatible with the active [.NET profile](dotnet-profile-support.md). But your scripted types gain additional functionality in Unity when they inherit from a built-in Unity type.

If your custom types inherit from [UnityEngine.Object](class-Object.md), they’ll be assignable to fields in the **Inspector** window. Inheriting from [MonoBehaviour](class-MonoBehaviour.md) allows a script to be attached to a **GameObject** as a component to control the behaviour of a GameObject in a **scene**.

For more information on fundamental Unity types you can inherit from, refer to [Fundamental Unity types](fundamental-unity-types.md).

For more information on viewing **scripts** and editing script components in the Inspector, refer to [Inspecting scripts](inspecting-scripts.md).

## Editor and runtime scripts

There are two distinct contexts in which your code might run:

* In the Unity Editor at authoring time, usually as part of custom editing tools and windows that support your own development process.
* In your application at runtime, as part of your user’s experience.

Source files that contain Editor-only code are commonly referred to as **Editor scripts** and can be excluded from Player builds either by placing them in folders named `Editor` or by using [assembly definition files](assembly-definition-files.md) to define Editor-only assemblies. For more information, refer to [Editor folder](assembly-definition-intro#special-folders).

Regions of code can also be conditionally compiled for Editor-only use with the `#if UNITY_EDITOR` preprocessor directive. For more information, refer to [Conditional compilation](conditional-compilation.md).

The majority of Unity’s core public APIs are in either the `UnityEditor` or `UnityEngine` namespaces. APIs in the `UnityEditor` namespace are only available in the Editor and attempting to use them in runtime code produces compilation errors.

While working in the Editor, you might have some code that you want to run in both Edit mode and Play mode. You might call some `UnityEditor` APIs while in Play mode for the purposes of testing, visualization, or triggering asset operations. Likewise, you might want to call `UnityEngine` APIs in Edit mode by using [`[ExecuteInEditMode]`](../ScriptReference/ExecuteInEditMode.md) or [`[ExecuteAlways]`](../ScriptReference/ExecuteAlways.md) attributes on those sections of code.

However, always remember to exclude Editor scripts and code that uses `UnityEditor` APIs from your runtime Player builds.

## Compilation and code reload

[Compilation](script-compilation.md) transforms the C# code you write into code that runs on a given target platform. Some aspects of compilation are under your control and others aren’t. By [organizing your scripts into assemblies](assembly-definition-files.md) you can reduce unnecessary recompilation and manage your dependencies effectively. With [conditional compilation](conditional-compilation.md) you can selectively include or exclude sections of your code from compilation.

Depending on your settings, Unity [recompiles and reloads your code](compilation-and-code-reload.md) in various contexts. Reloading code is important for changes to take effect or to preserve state when transitioning between Edit mode and Play mode, but it also impacts performance and iteration times. It’s important to understand these costs and how you can configure Unity’s code reload behavior to mitigate them.

## Additional resources

* [Creating scripts](creating-scripts.md)
* [Naming scripts](naming-scripts.md)
* [Scripts in the Inspector window](inspecting-scripts.md)
* [Fundamental Unity types](fundamental-unity-types.md)