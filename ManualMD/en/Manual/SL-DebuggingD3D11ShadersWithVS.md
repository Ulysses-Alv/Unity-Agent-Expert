# Debug shaders with Visual Studio

You can use Visual Studio to debug **shaders** in a Unity application on Windows platforms using DirectX 11 or 12. This page contains information on how to do this.

**Note:** If you are using DirectX 12, Microsoft recommends that you use PIX to debug shaders instead of Visual Studio. For information on using PIX on Windows with Unity, see [Debugging shaders using PIX](DebuggingShadersWithPIX.md).

## Preparing your shaders

To debug shaders, you must compile them with debug symbols included. To do that, insert the `#pragma enable_d3d11_debug_symbols` directive into the source code of each shader that you want to debug.

**Warning:** This pragma directive can negatively affect performance. Remove it from your shader code before you make a final build. For more information on this pragma directive, see [Shader compilation: pragma directives](SL-PragmaDirectives.md).

## Creating a placeholder Visual Studio project for Windows Standalone

If you build your application for Windows Standalone, you must create a placeholder Visual Studio project. If you build your application for Universal Windows Platform, Unity generates a Visual Studio project for you.

1. Launch Visual Studio.
2. Go to **File** > **New** > **Project** > **Visual C++** > **Empty Project**.
3. Go to **Project** > **Properties** > **Configuration Properties** > **Debugging**
4. In the **Command** field, replace $(TargetPath) with the path to your Windows Standalone application (for example, *C:\MyApp\MyApp.exe*)
5. If you want to force the project to run under DirectX 11, select **Command Arguments** and type **-force-d3d11**.

## Using Visual Studio to debug shaders

For instructions on setting up Visual Studio, see the Microsoft documentation: [Install Visual Studio](https://docs.microsoft.com/en-us/visualstudio/install/install-visual-studio?view=vs-2019).

For instructions on setting up and using Visual Studio’s graphics debugging tools, see the Microsoft documentation: [Visual Studio Graphics Diagnostics](https://docs.microsoft.com/en-us/visualstudio/debugger/graphics/visual-studio-graphics-diagnostics?view=vs-2019).