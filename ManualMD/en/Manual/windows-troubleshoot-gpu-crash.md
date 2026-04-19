# Troubleshoot D3D12 GPU crashes on Windows

A GPU crash can cause the Unity Editor to close unexpectedly and return a graphics error, for example, `Unrecoverable D3D12 device error! Run with -force-d3d12-debug and see logs for more info`. Diagnosing these crashes is often difficult because the cause might not be in your project’s code. This page explains how to identify the source of a GPU crash and provides a set of steps to troubleshoot and resolve GPU crashes.

## Types of crashes

There are a number of errors that you might encounter that are referred to as GPU crashes when developing in Unity:

* **GPU timeout**: When a graphics command takes too long to execute the operation, Windows will time out the operation and reset the graphics driver, causing a crash. For example:
  + An infinite loop in a **shader** caused by stale or uninitialized data.
  + Synchronization primitives in compute shader code, for example `GroupMemoryBarrierWithGroupSync` when a compute thread misses a call.
* **GPU page fault**: Occurs when a graphics command tries to access GPU memory that’s no longer available, such as a buffer that was freed before the GPU finished using it. This action results in a crash.
* **Graphics driver crash**: An error occurs in a third-party dynamic link-library (.dll file) that Unity calls to handle tasks like shader compilation and resource allocation.
* **GPU crash**: A hardware-level failure, often caused by the GPU running out of dedicated video memory (VRAM).

## Understand the driver stack

Unity doesn’t communicate directly with your GPU. Instead, it relies on a layered driver stack between your project and the hardware. This stack contains multiple dynamic link-libraries (DLLs). Each library acts as a bridge in the rendering pipeline, handling tasks like shader compilation, resource management, and command translation.

When something goes wrong in one of these DLLs, the impact can be severe. Because they operate at the system level, traditional debugging methods are often inadequate. You can still get valuable insights from driver symbol servers, but you can’t step through the driver code like you can with your own code.

Therefore, understanding the graphics driver architecture is essential. The flow is as follows:

* The DLLs you find in crash logs are almost always user-mode drivers (UMDs). These run in your application’s process space and handle shader compilation, command buffer generation, and resource management. When they crash, your application crashes with them, which is why Unity closes suddenly.
* The kernel-mode driver (KMD) operates at a lower level and manages hardware access across all applications. KMD crashes are more severe and can cause blue screens, though modern Windows can often recover by restarting the driver.

User-mode drivers are complex as they’re responsible for compiling shaders, managing GPU memory and resources, and often working around bugs in applications.

A crash usually occurs when the driver encounters an unexpected situation, such as:

* A shader that contains an unusual pattern that triggers a compiler bug.
* An untested sequence of API calls from the application.
* Corruption in the driver’s own shader cache.
* Hardware-specific bugs that appear only under certain conditions.

## Rule out system-level issues

Before troubleshooting the driver directly, it’s important to rule out other system-level issues.

Perform the following checks:

* **Disable overlays**: Applications that display overlays such as Discord, Steam, or MSI Afterburner can interfere with Unity’s rendering process. Disable or close this software before running the Editor.
* **Reset your GPU to its default state**: If you overclock your GPU, revert it to its default factory clock speed. Additionally, open your GPU’s control panel and restore all settings to their defaults. On some laptops, restoring default settings can cause the system to use the integrated GPU instead of the dedicated GPU, especially when the laptop isn’t connected to a power source or has a low battery. Integrated GPUs usually have less processing power and might have more compatibility issues, so ensure that you configure the Unity Editor and Player to run on your dedicated GPU.
* **Check for operating system updates**: Ensure your operating system is fully up-to-date. Microsoft frequently releases patches that can resolve stability issues with core components like DirectX.
* **Update your graphics drivers**: Outdated drivers are a common source of crashes and rendering issues. It’s recommended to update your drivers to the latest version using tools such as the NVIDIA GeForce Experience or AMD Software: Adrenalin Edition.
* **Run a system file scan**: To check for and repair corrupted Windows files, open Command Prompt as an Administrator and run `sfc /scannow`, followed by `DISM /Online /Cleanup-Image /RestoreHealth`.

## Identify the type of crash

A crash is typically either a graphics driver issue or a GPU hardware issue.

To find the cause, you need to examine the crash log. For the location of crash reports on your computer, refer to the [Log files reference](log-files.md).

### Observe symptoms and analyze the Editor log

Common symptoms of a GPU-related failure include:

* A high GPU load before the crash.
* Visual glitches or freezes in the Editor.
* Sudden Editor exits.
* Windows resetting the display driver.

If Unity intercepts the crash, you can examine the stack trace at the end of the `Editor.log` file. A graphics-related failure typically shows a specific stack trace that involves Unity’s rendering thread and the graphics API. Look for these key markers:

* **A rendering thread failure**: The stack trace includes `GfxDeviceWorker::RunGfxDeviceWorker`. This indicates the engine failed while dispatching instructions to the GPU rather than during standard game logic.
* **Graphics API calls**: The stack trace mentions the graphics API layer, such as `D3D12Fence::Wait`.
* **Device error codes**: The log contains errors like `Device failed error (887a0006)` or `Device removal`.
* **GPU memory status**: The log might indicate the memory status, for example: `d3d12: GfxDevice was not out of Local memory`.

A call stack might look like the following example:

```
========== OUTPUTTING STACK TRACE ==================

0x00007FFF7F3580DA (KERNELBASE) RaiseException
0x00007FFDF037F596 (Unity) LaunchBugReporter
0x00007FFDF0BC03A7 (Unity) EditorMonoConsole::LogToConsoleImplementation
0x00007FFDF0BC0EE3 (Unity) EditorMonoConsole::LogToConsoleImplementation
0x00007FFDF2AA2ED8 (Unity) DebugStringToFilePostprocessedStacktrace
0x00007FFDF2AA26E1 (Unity) DebugStringToFile
0x00007FFDF21608F5 (Unity) CheckDeviceStatus
0x00007FFDF21CB7A2 (Unity) CheckCompletedValue
0x00007FFDF21D1E70 (Unity) D3D12Fence::Wait
0x00007FFDF2195721 (Unity) GfxDeviceD3D12::QueuePresent
0x00007FFDF21C20DE (Unity) D3D12Window::EndRendering
0x00007FFDF1D077CD (Unity) GfxDeviceWorker::RunCommand
0x00007FFDF1D09FCC (Unity) GfxDeviceWorker::RunExt
0x00007FFDF1D0A0FF (Unity) GfxDeviceWorker::RunGfxDeviceWorker
0x00007FFDF2AC12FC (Unity) Thread::RunThreadWrapper
0x00007FFF80CBE8D7 (KERNEL32) BaseThreadInitThunk
0x00007FFF81F2C53C (ntdll) RtlUserThreadStart

========== END OF STACKTRACE ===========
```

The log might also point to the exact script or asset that caused the GPU to enter a fatal state. For example, a reference to `CreateResources` in the log might trace back to a specific script and line number:

```
CreateResources
...
GPUCrash:CreateResources () (at Assets/GPUCrash.cs:25)
GPUCrash:Update () (at Assets/GPUCrash.cs:47)
```

### Detailed crash logs

To view detailed logs, check the following prerequisites:

* Your device must be running Windows 11.
* You must enable [Microsoft’s Graphics diagnostics tools](https://learn.microsoft.com/en-us/windows/uwp/gaming/use-the-directx-runtime-and-visual-studio-graphics-diagnostic-features)
* Start the Editor with the following arguments:
  + `-force-d3d12-debug`
  + `-force-d3d12-editor-pix-events`

**Note**: You can also use the argument `-force-d3d12-debug-gbv` to return crash information, however, it’s recommended only as a secondary step in the debug process. It enables GPU-based validation (GBV), which can add a large amount of debug information to crash logs. Analyze the log without this flag first to avoid the production of large files.

These arguments enable [Microsoft’s Device Removed Extended Data (DRED) API](https://learn.microsoft.com/en-us/windows/win32/direct3d12/use-dred), which can provide more insight into the GPU state at the time of the crash.

The log displays errors in the following formats:

* Validation errors: `d3d12 validation error: <message>`
* DXGI errors: `d3d12 DXGI error: <message>`

### DLL reference

For most crashes, the name of the .dll file in your crash log is the key to identifying the source. The vast majority of these happen in a User-Mode Driver (UMD), which translates commands from graphics APIs like DirectX into instructions your GPU can understand.

| **DLL file** | **Manufacturer** | **Type** | **Description** |
| --- | --- | --- | --- |
| `amdxc64.dll` | AMD | UMD | Detects whether DirectX, OpenGL or Vulkan graphics API is in use. |
| `atidxx64.dll` | AMD | UMD | Translates DirectX API calls, shaders, and resources into GPU-specific structures. |
| `igc64.dll` | Intel | UMD | Compiles HLSL or GLSL into optimized machine code for Intel GPU architectures. Crashes here often indicate shader compilation issues. For example, this might be caused by a shader that uses features in an unexpected way or hits compiler bugs. |
| `igd10um64.dll` | Intel | UMD | Translates DirectX API calls, shaders, and resources into GPU-specific structure for older Intel GPUs, for example Intel HD/UHD Graphics pre–11th gen. |
| `igd10um64xe.dll` | Intel | UMD | Translates DirectX API calls, shaders, and resources into GPU-specific structures for modern Xe architecture (11th Gen Core and newer, including Arc GPUs). |
| `nvwgf2umx.dll` | NVIDIA | UMD | Translates DirectX API calls, shaders, and resources into GPU-specific structures. |
| `nvoglv64.dll` | NVIDIA | UMD | If you target OpenGL instead of DirectX, this translates OpenGL API calls, compiles GLSL shaders, and manages the entire OpenGL rendering pipeline. |

## Fix a graphics driver crash

After identifying the driver, follow these steps to resolve the issue.

### 1. Perform a clean driver installation

Updating a driver isn’t always enough to fix a crash. A clean installation removes all old and potentially corrupt files before installing a fresh version.

* **AMD**:

  1. Use the official [**AMD Cleanup Utility**](https://www.amd.com/en/resources/support-articles/faqs/GPU-601.html) or a third-party tool like Display Driver Uninstaller (DDU) to completely remove all existing driver files.
  2. Download and install the latest stable driver from the official AMD website. If crashes began after an update, consider rolling back to a previously stable version.
* **NVIDIA**:

  1. Download the latest Game Ready or Studio driver from the NVIDIA website.
  2. Run the installer and select **Custom (Advanced)**.
  3. On the next screen, check the **Perform a clean installation** box and proceed. For persistent issues, use DDU before reinstalling.
* **Intel**:

  1. Use a clean up tool like DDU to cleanly remove all exisiting driver files.
  2. Download and install the latest graphics drivers from the official Intel Download Center or your computer manufacturer’s website.

If the crash persists, try installing an older version of a driver.

### 2. Clear the shader cache

Corrupted shader caches maintained by Unity and the graphics driver can cause recurring compilation crashes. To clear the caches completely, delete the contents of the following folders:

| **Type of cache** | **Location** | **Description** |
| --- | --- | --- |
| **Unity’s shader cache** | `Library/ShaderCache` | Delete the contents of this folder to clear the project Library shader cache for the Unity Editor. |
| **Driver shader caches** | `%LOCALAPPDATA%\AMD\DxCache`   `%LOCALAPPDATA%\AMD\DxcCache`   `%LOCALAPPDATA%\AMD\VkCache`   `%LOCALAPPDATA%\AMD\GLCache`   `%LOCALAPPDATA%\AMD\DX9Cache`   `%LOCALAPPDATA%\Intel\ShaderCache`   `%LOCALAPPDATA%\NVIDIA\DXCache`   `%LOCALAPPDATA%\D3DSCache`   `%ProgramFiles(x86)%\Steam\steamapps\shadercache` | Delete the contents of these folders to clear the driver shader cache for the built Player. |

After clearing these caches, restart your computer to ensure no cached shaders remain in memory.

### 3. Use a fallback graphics API

If the crash is specific to DX12, you can temporarily revert to DX11 while you investigate the root cause. Use the [D3D12 Device Filtering Asset](introduction-d3d12-device-filtering-asset.md) to create a fallback.

## Report a crash to Unity

If you have tried the solutions in this document and the crash still persists, you can submit a bug report. When you submit your report, include as much detail as possible and list the steps to reproduce the issue.

For more information on reporting bugs, refer to [How to report bugs](https://unity.com/releases/editor/qa#bug-reporting).

## Additional resources

* [Log files reference](log-files.md)
* [Introduction to D3D12 Device Filtering Asset](introduction-d3d12-device-filtering-asset.md)
* [How to report bugs](https://unity.com/releases/editor/qa#bug-reporting)