# Unity Player command line arguments reference

You can launch Unity Players from the command line and pass in arguments to change how the Player executes.

These arguments work on all standalone platforms except for the [Web platform](webgl.md). The following table specifies any additional platform requirements.

For more information on launching Unity from the command line, refer to [Run Unity from a command-line interface](command-line-run-unity.md).

| **Command** | **Description** |
| --- | --- |
| `-batchmode` | Run the application in headless mode. In this mode, the application doesn’t display anything or accept user input. |
| `-disable-gpu-skinning` | Disables Graphics Processing Unit (GPU) **skinning** at startup. |
| `-dontConnectAcceleratorEvent` (UWP only) | Disable connecting to AcceleratorKeyEvent. This might help if you have issues with input in XAML elements.  **Note:** Unity can’t handle some keyboard keys, such as F10, Ctrl, Alt, and Tab. |
| `-force-clamped` | Use this together with `-force-glcoreXY` to prevent checks for additional OpenGL extensions, allowing the application to run between platforms with the same code paths. |
| `-force-d3d11` (Windows only) | Force the application to use Direct3D 11 for rendering. |
| `-force-d3d11-bitblt-model` (Windows only) | Force the application to use DXGI BitBlt model swapchain when using Direct3D 11. For more information, refer to [`PlayerSettings.useFlipModelSwapchain`](../ScriptReference/PlayerSettings-useFlipModelSwapchain.md). |
| `-force-d3d11-flip-model` (Windows only) | Force the application to use DXGI flip model swapchain when using Direct3D 11. For more information, refer to [`PlayerSettings.useFlipModelSwapchain`](../ScriptReference/PlayerSettings-useFlipModelSwapchain.md). |
| `-force-d3d11-no-singlethreaded` (Windows and UWP only) | Force DirectX 11.0 to be created without a `D3D11_CREATE_DEVICE_SINGLETHREADED` flag. |
| `-force-d3d11-singlethreaded` (Windows and UWP only) | Force DirectX 11.0 to be created with a `D3D11_CREATE_DEVICE_SINGLETHREADED` flag. |
| `-force-d3d12` (Windows only) | Force the application to use Direct3D 12 for rendering. |
| `-force-device-index` | Make the Standalone Player use a specific GPU device by passing it the index of that GPU. This option is supported for D3D11, D3D12, Metal, and Vulkan graphics APIs, but isn’t supported for OpenGL. |
| `-force-driver-type-warp` (Windows and UWP only) | Force the DirectX 11.0 driver type WARP device. For more information, refer to Microsoft’s documentation on [Windows Advanced Rasterization Platform](https://docs.microsoft.com/en-gb/windows/win32/direct3darticles/directx-warp?redirectedfrom=MSDN). |
| `-force-feature-level-10-0` (Windows and UWP only) | Force DirectX 11.0 feature level 10.0. |
| `-force-feature-level-10-1` (Windows and UWP only) | Force DirectX 11.0 feature level 10.1. |
| `-force-feature-level-11-0` (Windows and UWP only) | Force DirectX 11.0 feature level 11.0. |
| `-force-gfx-direct` | Force single threaded rendering. |
| `-force-glcore` | Force the application to use the OpenGL core profile for rendering. The Editor tries to use the most recent OpenGL version available, and all OpenGL extensions exposed by the OpenGL drivers. Unity uses Direct3D if the platform doesn’t support OpenGL. |
| `-force-glcoreXY` | Similar to `-force-glcore`, but requests a specific OpenGL context version. Accepted values for XY: 32, 33, 40, 41, 42, 43, 44, or 45. |
| `-force-low-power-device` (macOS only) | Make the Standalone Player use a low power device. |
| `-force-metal` (macOS only) | Make the Standalone Player use Metal as the default graphics API. |
| `-forceTextBoxBasedKeyboard` (UWP only) | Use TextBox-based implementation for TouchScreenKeyboard. This implementation allows switching to different implementations, in case there are issues with the default.   **Note:** This implementation has an effect only on UWP XAML applications. |
| `-force-vulkan` | Force the application to use Vulkan for rendering. |
| `-force-wayland` (Linux only) | Activate experimental Wayland support when running a Linux player. |
| `-logFile <pathname>` | Specify where Unity writes the standalone Player log file. To output to the console, specify `-` for the path name. On Windows, use `-logfile` to direct the output to `stdout`, which by default isn’t the console. |
| `-log-memory-performance-stats` | Adds a detailed memory report to the Player log file when closing the Player. |
| `-max-async-pso-job-count` | Set the number of parallel threads Unity uses to create pipeline state objects (PSOs) when you use [`Experimental.Rendering.ShaderWarmup`](../ScriptReference/Experimental.Rendering.ShaderWarmup.md) to prewarm **shader** variants. This only has an effect if your application runs on a platform that uses the DirectX 12, Metal, or Vulkan graphics API. |
| `-monitor N` | Run Standalone Player on the specified monitor, indicated by a 1-based index number. |
| `-nographics` | When you use this argument in batch mode, Unity doesn’t initialize a graphics device. This makes it possible to run your automated workflows on machines that don’t have a GPU.   **Note**: Output logs are turned off in this mode. |
| `-nolog` | Disables Player logging. For more information on the content and location of Player logs, refer to [Log files reference](log-files.md) folder. |
| `-no-stereo-rendering` | Turn off stereo rendering. |
| `-parentHWND <HWND> delayed` (Windows only) | Embed the Windows Standalone application into another application. When you use this argument, you need to pass the parent application’s window handle (‘HWND’) to the Windows Standalone application.  When you pass `-parentHWND 'HWND' delayed`, the Unity application is hidden while it runs. You must also call [SetParent](https://docs.microsoft.com/en-gb/windows/win32/api/winuser/nf-winuser-setparent?redirectedfrom=MSDN) from the [Microsoft Developer library](https://docs.microsoft.com/en-gb/) for Unity in the application. Microsoft’s `SetParent` embeds the Unity window. When it creates Unity processes, the Unity window respects the position and size provided as part of Microsoft’s [STARTUPINFO](https://docs.microsoft.com/en-gb/windows/win32/api/processthreadsapi/ns-processthreadsapi-startupinfoa?redirectedfrom=MSDN) structure.  To resize the Unity window, check its [GWLP\_USERDATA](https://docs.microsoft.com/en-gb/windows/win32/api/winuser/nf-winuser-getwindowlongptra?redirectedfrom=MSDN) in Microsoft’s `GetWindowLongPtr` function. Its lowest bit is set to 1 when the graphics initialize and it’s safe to resize. Its second lowest bit is set to 1 after the Unity splash screen finishes displaying.  For more information, see this downloadable example: [EmbeddedWindow.zip](../uploads/Examples/EmbeddedWindow.zip) |
| `-popupwindow` | Create the window as a dialog window, without a frame. This command isn’t supported on macOS. |
| `-screen-fullscreen` | Override the default full-screen state. This must be 0 or 1. |
| `-screen-height` | Override the default screen height. This must be an integer from a supported resolution. |
| `-screen-quality` | Override the default screen quality. Example usage would be: `/path/to/myGame -screen-quality Beautiful`. The supported options match the [Quality Settings](class-QualitySettings.md) names. |
| `-screen-width` | Override the default screen width. This width value must be an integer from a supported resolution. |
| `-single-instance` (Linux and Windows only) | Run only one instance of the application at the time. If another instance is already running then launching the application again with `-single-instance` focuses the existing one. |
| `-systemallocator` | Forces the platform to use the system allocator. This can be useful if you want to use tools like address sanitizers to debug memory issues. You should only use this option for debugging purposes. |
| `-timestamps` | Prefixes every [Player log](log-files.md) message with the current timestamp and thread ID. |
| `-wait-for-managed-debugger` | Makes the Player wait for a managed debugger to attach before launching. Generates a dialog that displays the Player’s process ID. For more information on managed debugging, refer to [Debug C# code in Unity](managed-code-debugging.md#DebugInEditor). |
| `-wait-for-native-debugger` | Makes the Player wait for a native debugger to attach before launching. Generates a dialog that displays the Player’s process ID. |
| `-window-mode` (Windows only) | Override full screen windowed mode. Accepted values are `exclusive` or `borderless`. For more information, refer to [Player settings](class-PlayerSettings.md). |

## Additional resources

* [Create a build from the command line](build-command-line.md)
* [Android Player command-line arguments](android-custom-activity-command-line.md)
* [Use Unity Accelerator on the command line](accelerator-command-line.md)
* [Build for Embedded Linux from the command line](embedded-linux-build-command-line.md)
* [Build for QNX from the command line](qnx-build-command-line.md)
* [Test Framework command-line arguments reference](test-framework/reference-command-line.md)
* [Profiler command-line arguments reference](profiler-command-line-arguments.md)
* [Run Project Auditor from the command line](project-auditor/run-from-command-line)