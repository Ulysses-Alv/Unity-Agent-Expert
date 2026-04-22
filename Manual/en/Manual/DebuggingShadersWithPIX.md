# Debug shaders with PIX

PIX is a performance tuning and debugging tool by Microsoft, for Windows developers. It offers a range of modes for analyzing an application’s performance, and includes the ability to capture frames of DirectX projects from an application for debugging.

Use PIX to investigate issues in Windows 64-bit (x86\_64) Standalone or Universal Windows Platform applications.

To install PIX, [download and run the Microsoft PIX installer](https://blogs.msdn.microsoft.com/pix/download/) and follow the instructions.

For more information about PIX, see Microsoft’s [PIX Introduction](https://blogs.msdn.microsoft.com/pix/introduction/) and [PIX Documentation](https://blogs.msdn.microsoft.com/pix/documentation/).

## Debugging DirectX shaders with PIX

You should use a built version of your Unity application to capture frames, rather than a version running in the Unity Editor. This is because you need to launch the target application from within PIX to capture GPU frames.

Using a **development build** adds additional information to PIX, which makes navigating the **scene** capture easier.

### Create a project with a debug-enabled Shader

To debug the shader with source code in PIX, you need to insert the following pragma into the shader code: `#pragma enable_d3d11_debug_symbols`

### Example

The following walkthrough uses a basic example to demonstrate the entire process.

#### Create a basic project:

1. Create a new Unity project (see the Hub documentation on [Projects](https://docs.unity3d.com/hub/manual/Projects.html)).
2. In the top menu, go to **Assets** > **Create** > **Shader** > **Standard Surface Shader**. This creates a new shader file in your **Project** folder.
3. Select the shader file, and in the **Inspector** window, click **Open**. This opens the shader file in your scripting editor. Insert `#pragma enable_d3d11_debug_symbols` into the shader code, underneath the other `#pragma` lines.
4. Create a new Material (menu: **Assets** > **Create** > **Material**).
5. In the Material’s Inspector window, select the **Shader** dropdown, go to **Custom**, and select the shader you just created.
6. Create a 3D cube GameObject (menu: **GameObject** > **3D Object** > **Cube**).
7. Assign your new Material to your new GameObject. To do this, drag the Material from the **Project** window to the 3D cube.

### Capture a frame from a Windows Standalone application:

1. Go to **File** > **Build Profiles**, and under **Platforms**, select **Windows** or [create a build profile](create-build-profile.md) for the **Windows** platform. Set the **Architecture** to **Intel 64-bit**, and enable the **Development Build**.
2. Click **Build**.
3. Launch **PIX.**
4. Click on **Home**, then **Connect**
5. Select Computer **localhost** to use your PC for capturing, and click **connect**.
6. In the **Select Target Process** box, select the **Launch Win32** tab and use the **Browse** button to select your application’s executable file. Note that here, “Win32” means a non-UWP application; your application file must be a 64-bit binary file.
7. Enable **Launch for GPU Capture**, then use the **Launch** button to start the application.
8. Use your application as normal until you are ready to capture a frame. To capture a frame, press **Print Screen** on your keyboard, or click the **camera** icon on the GPU Capture Panel. A thumbnail of the capture appears in the panel. To open the capture, click the thumbnail.
9. To start analysis on the capture, click the highlighted text or the small **Play** icon on the menu bar.
10. Select the **Pipeline** tab and use the\_\_ Events\_\_ window to navigate to a draw call you are interested in.
11. In the lower half of the **Pipeline** tab, select a render target from the **OM** (Output Merger) list to view the output of draw call. Select a **pixel** on the object you want to debug. Note that you can right-click a pixel to view the draw call history, as a way of finding draw calls you are interested in.
12. Select **Debug Pixel** on the **Pixel Details** panel.
13. On the debug panel, use the Shader Options to select which shader stage to debug.
14. Use the **toolbar** or keyboard shortcuts to step through the code.

For more information on debugging shaders with PIX, see Microsoft’s video series [PIX on Windows](https://www.youtube.com/playlist?list=PLeHvwXyqearWuPPxh6T03iwX-McPG5LkB), particularly [Part 5 - Debug Tab](https://www.youtube.com/watch?v=jRflMYmXv2g).

For more information on GPU capture in PIX, see Microsoft’s documentation on [GPU Captures](https://blogs.msdn.microsoft.com/pix/gpu-captures/).

* 2018–09–17 Page published