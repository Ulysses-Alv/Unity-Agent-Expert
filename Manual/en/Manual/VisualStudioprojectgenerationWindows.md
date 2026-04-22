# Visual Studio project generation for Windows

The **Create Visual Studio Solution** Build Setting makes Unity generate a Visual Studio [Solution](https://docs.microsoft.com/en-us/visualstudio/ide/solutions-and-projects-in-visual-studio?view=vs-2019#solutions) instead of build a Player.
Generating a Visual Studio Solution enables you to change your build process. For example, you can:

* Modify your application’s manifest.
* Add C++ code.
* Modify embedded resources.
* Launch your application with the debugger attached.

**Note**: This option is only available when you are using the Windows build target.

## Create a Visual Studio Solution

**Warning:** Make sure you save and back up any changes you make in Visual Studio before rebuilding your Unity project. If you do not back up your changes, Unity overwrites them during the build process.

1. Go to **File** > **Build Profiles**.
2. From the list of platforms in the **Platforms** panel, select **Windows** or [create a build profile](create-build-profile.md) for the **Windows** platform.
3. In the **Platform Settings** section, enable **Create Visual Studio Solution**.
4. Select **Build** to generate the solution.

By default, Unity stores the Visual Studio Solution you generate in the same directory as your built project.

### Scripting backends

Depending on which [scripting backend](scripting-backends.md) your Unity project uses, the Visual Studio Solution that Unity generates includes a different number of projects. For the [Mono](./scripting-backends-mono.md) scripting backend, the solution includes three projects. For the [IL2CPP](./scripting-backends-il2cpp.md) scripting backend, the solution includes four projects.

**Note**: You can only build Windows IL2CPP on a Windows machine.

## Project retarget wizard

Unity generates a Visual Studio Solution that targets Visual Studio 2015. If you are using a newer version of Visual Studio, you might see the **Retarget Projects** dialog box when you open your Visual Studio Solution.

The **Retarget Projects** dialog prompts you to upgrade your Visual Studio Solution to match the Visual Studio version you have installed. Select **OK** to accept this suggestion.

## Solution contents

Below is a list of the projects Unity generates for your Visual Studio Solution.

| **Project** | **Description** |
| --- | --- |
| **projectName**  (represents your project name) | This is your main project. Visual Studio builds this project into the final application executable. You are most likely to make changes here. |
| **UnityData**  (stored inside the projectName project) | This project contains all the Unity-specific files you need to build your project, such as assets. |
| **UnityPlayerStub** | This is a stub library for `UnityPlayer.dll`. Unity uses this to link your executable to the `UnityPlayer.dll` and expose the available `UnityPlayer.dll` API. |
| **Il2CppOutputProject**   (IL2CPP scripting backend only) | This project contains:  • The generated C++ code which Unity converts from managed assemblies. • The source code for IL2CPP runtime. • The IL2CPP garbage collector.   You can debug this code from inside Visual Studio. |

## Build configuration options in Visual Studio

Visual Studio provides a variety of build configuration options. All optimization and debug references in this section refer to the Visual Studio Solution only. To select your options, click on **Release** and choose from the drop-down menu.

### Types of build configurations

| **Configuration** | **Use** | **Description** |
| --- | --- | --- |
| **Debug** | Use the **Debug** configuration when debugging C++ code. | • Disables all optimization. • Preserves all debugging information in the code. • Results in code that runs slowly. |
| **Release** | Use the **Release** configuration to profile your game. | • Enables code optimizations. |
| **Master** | Use the **Master** configuration for game submission and final testing. | • Disables the profiler. • Results in the same build time as the **Release** configuration. |
| **MasterWithLTCG** (Only available when the IL2CPP scripting backend is enabled.) | If you are using the IL2CPP scripting backend, you can choose either **Master** or **MasterWithLTCG** for game submission and final testing. | • Enables link time code generation for generated C++ code, IL2CPP runtime and IL2CPP garbage collection. • Results in much longer build times compared to the **Master** configuration. • Resulting application executes faster. |