# RenderDoc integration

The Editor supports integrated launching and capture of the [RenderDoc](https://github.com/baldurk/renderdoc) graphics debugger, for detailed frame introspection and debugging.

The integration is only supported for RenderDoc versions 0.26 or later, so if an earlier version is currently installed it is required that you update to at least version 0.26.

**Note:** While the integration is only available in the Editor, it is quite possible to use RenderDoc as normal with no extra setup in standalone player builds.

**Note:** Frames can only be captured if Unity is running on a platform and API that RenderDoc supports. If another API is in use, the RenderDoc integration will be temporarily disabled until a supported API is enabled. Refer to the [RenderDoc documentation](https://renderdoc.org/docs/getting_started/features.html#) for more information on supported platforms and APIs.

## Loading RenderDoc

If a RenderDoc installation is detected, then at any time after loading the Editor you can right click on the tab for the [Game View](GameView.md) or [Scene View](UsingTheSceneView.md) and click the ‘Load RenderDoc’ option. This will reload the graphics device so you must save any changes, but afterwards RenderDoc will be ready to capture without having to restart the editor or build a standalone player.

**Note:** You can also launch the Editor via RenderDoc as normal, or pass the -load-renderdoc command line option to load RenderDoc from startup.

## Capturing a frame with RenderDoc

When a compatible version of RenderDoc is detected as loaded into the Editor, a new button will appear on the right side of the **toolbar** on the Game and **Scene** Views.

![Capturing a frame with RenderDoc](../uploads/Main/RenderDocCaptureButton.png)

Capturing a frame with RenderDoc

Pressing this button will trigger a capture of the next frame of rendering for the view. If the RenderDoc tool UI has not been opened, a new instance will be launched to show the capture, and if it is already running the newest capture will automatically appear there. From there you can open the capture and debug using the tool.

![List of frame captures in RenderDoc](../uploads/Main/RenderDocCaptureList.jpg)

List of frame captures in RenderDoc

## Including shader debug information

By default to optimise the size of DirectX11 shaders, debugging information is stripped out. This means that constants and resources will have no names, and the **shader** source will not be available. To include this debugging information in your shader, include **#pragma enable\_d3d11\_debug\_symbols** in your shader’s `HLSLPROGRAM` block.

## Alternative graphics debugging techniques

If you build a standalone player using D3D11, you can capture a frame and debug using the [Visual Studio graphics debugger](SL-DebuggingD3D11ShadersWithVS.md).

## Configuring Android SDK in RenderDoc

When using RenderDoc to develop Android applications, ensure to configure the Android SDK correctly using one of the following ways:

### Set RenderDoc to use the same SDK as Unity (recommended)

1. In the RenderDoc window, go to **Tools** > **Settings** > **Android**.
2. Under **Android SDK root path**, enter the path that Unity uses for Android SDK. You can find this path in the Unity Editor under **Edit** > **Preferences** > **Android SDK Tools Installed with Unity**.

### Align ADB versions and paths between RenderDoc and Unity

If the **Android SDK root path** in RenderDoc settings isn’t set, RenderDoc uses the path set for **ADB** in the `PATH` environment variable.

1. Find the ADB path set in the `PATH` variable.

   * Windows PowerShell: `Get-Command adb | Select-Object Path`
   * Windows command prompt: `where adb`
2. Check the ADB version in Windows terminal: `C:\[adb path from command output]\adb.exe version`.
3. Find the ADB path that Unity uses. In the Unity Editor, go to **Edit** > **Preferences** > **Android SDK Tools Installed with Unity**.
4. Check the corresponding ADB version in Windows terminal: `C:\[path to Android sdk platform-tools]\adb.exe version`.
5. Compare the ADB versions that RenderDoc and Unity use. If they differ, choose one of the ADB locations and configure both tools to use the same path.

**Note**: Make sure platform-tools from the selected SDK are [up to date](https://developer.android.com/studio/intro/update) by using the Android SDK manager in Android Studio.