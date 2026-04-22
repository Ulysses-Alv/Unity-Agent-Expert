# Generate your Visual Studio C++ solution

When you [build a project from Unity for Universal Windows Platform](windowsstore-buildsettings.md) (UWP), Unity automatically generates a Visual Studio solution.

## Create your Visual Studio C++ solution

Unity generates a Visual Studio solution containing the following projects:

| **Project** | **Description** |
| --- | --- |
| **projectName** | Contains your main project code. Visual Studio builds this project into an application package, which you can deploy to a device or uploaded to the Microsoft Store.   **Note:** Unity doesn’t overwrite this project when you build on top of it. |
| **Unity Data** | Contains all the Unity-specific files you need to build your project, such as assets. |
| **Il2CppOutputProject** | Contains the generated C++ code which Unity converts from managed assemblies.   **Note:** This project is overwritten every time you build over it. |

## Visual Studio build configurations

Unity provides the following [build configuration options in Visual Studio](https://learn.microsoft.com/en-us/visualstudio/ide/understanding-build-configurations):

| **Configuration** | **Description** |
| --- | --- |
| **Debug** | Use **Debug** to debug your code. This configuration:   - Disables all optimization.   - Preserves all debugging information in the code.   - Results in code that runs slowly.   - Results in the fastest build time. |
| **Release** | Use **Release** to profile your game. This configuration:   - Enables code optimizations. |
| **Master** | Use **Master** for game submission and final testing. This configuration:   - Disables the profiler.   - Results in the same build time as the Release configuration.   - Results in the same build time as the **Release** configuration. |
| **MasterWithLTCG** | Use **MasterWithLTCG** for game submission and final testing. This configuration:   - Enables link time code generation for generated C++ code, IL2CPP runtime, and IL2CPP garbage collection.   - Results in much longer build times compared to the **Master** configuration.   - Results in an application that executes faster than the **Master** configuration. |