# Xcode frame debugger Unity integration

The [Xcode](https://developer.apple.com/xcode/)[frame debugger tool](https://developer.apple.com/documentation/metal/frame_capture_debugging_tools) lets you capture a frame of your application to see the commands that the GPU performed during that frame, examine data in GPU memory, and identify bottlenecks in your **shaders**. This lets you analyze GPU performance in fine detail.

Unity integrates with the Xcode frame debugger in the following ways:

* On macOS, iOS, and tvOS, you can use Xcode frame debugger to analyze frames from your application while it’s running on the target device.
* On macOS, you can use Xcode frame debugger to analyze frames from the Unity Editor.

**Important:**

* Frame debugging only works if the application is running on a platform and graphics API that Xcode supports.
* Xcode only supports macOS with Metal graphics.
* If Unity uses another API, the Xcode integration is disabled until you select a supported graphics API.

## Capture a frame from your application

To use the Xcode frame debugger to capture a frame from do one of the following:

* Launch your application from Xcode and either use the Xcode UI or the [FrameCapture](../ScriptReference/Apple.FrameCapture.md) API to request frame captures. You can analyze the frame captures in Xcode immediately, or save them to disk. This workflow is supported on all platforms that use Metal.
* Launch your application directly from the command line without an Xcode project, and use the [FrameCapture](../ScriptReference/Apple.FrameCapture.md) API to save frame captures to disk. **Note:** iOS doesn’t support this workflow.

### Capture a frame from your application with Xcode

This section describes how to launch your application and perform a frame capture using the Xcode UI or the FrameCapture API. This workflow is supported on all platforms that use Metal.

**1. Create an Xcode project.**

You can either create an Xcode project from the Unity Editor, or use any other Xcode project to launch macOS applications.

To build an Xcode project from the Unity Editor:

1. Open the **Build Profiles** window (**File > Build Profiles**).
2. macOS only: Enable **Create Xcode Project**.
3. Click Build.

To use another Xcode project to launch macOS applications:

1. In Xcode, create a new, empty macOS project, or open an existing macOS project.
2. Go to Product > Scheme > Edit scheme, and open the Info tab.
3. Set **Executable** to your built Unity application.

**2. Edit the Xcode project scheme so that you can perform frame captures.**

Either, use the Xcode GUI. To do this, follow the [Enabling Frame Capture guide](https://developer.apple.com/documentation/metal/frame_capture_debugging_tools/enabling_frame_capture) in the Xcode documentation to set your project scheme’s GPU **Frame Capture** setting to Metal.

You can also use the XcScheme API to configure an Xcode project scheme, which is useful for automated builds. For more information, see the [XcScheme API documentation](../ScriptReference/iOS.Xcode.XcScheme.md).

**3. Launch your project from Xcode, and perform a frame capture.**

In Xcode, press the Frame Capture button (camera icon) to capture the next frame of data.

You can also use the [FrameCapture](../ScriptReference/Apple.FrameCapture.md) API to perform a frame capture from a script. For information on analyzing frame capture data in Xcode, see the [Xcode frame debugger documentation](https://developer.apple.com/documentation/metal/frame_capture_debugging_tools).

### Capture a frame from your application with the command line

This section describes how to launch your application from the command line, perform a frame capture using the [FrameCapture](../ScriptReference/Apple.FrameCapture.md) API, and save the results to disk. **Important:** This workflow isn’t supported on iOS. To perform frame captures on iOS, you must always launch your application from Xcode.

1. Add calls to [FrameCapture.BeginCaptureToFile](../ScriptReference/Apple.FrameCapture.BeginCaptureToFile.md) and [EndCapture](../ScriptReference/Apple.FrameCapture.EndCapture.md) to your code, so that you can perform frame captures as required.
2. Launch your application from the command line, with the following flag: `-enable-metal-capture.` Xcode performs frame captures when your code requests them, and saves the results to disk.

For information on how to analyze this data in Xcode, see Apple’s [Xcode frame debugger documentation](https://developer.apple.com/documentation/metal/frame_capture_debugging_tools).

## Capture a frame from the Unity Editor

On macOS, you can use Xcode frame debugger to analyze frames from the Unity Editor. If you use Xcode to launch the Unity Editor, you can request frame captures from the Unity Editor UI.

### Capture a frame from Unity Editor with Xcode

Follow these steps to launch the Unity Editor from Xcode and perform a frame capture using the Unity Editor UI, the Xcode UI, or the FrameCapture API. You can analyze the frame capture immediately, or save the results to disk. This workflow is supported on macOS.

**Preqrequisite:** If the Unity Editor is open, close it.

1. In Xcode, create a new, empty macOS project, or open an existing macOS project.
2. Go to **Product > Scheme > Edit scheme**, and open the **Info** tab.
3. Set **Executable** to Unity Editor.
4. Follow the [Enabling Frame Capture guide in the Xcode documentation](https://developer.apple.com/documentation/metal/frame_capture_debugging_tools/enabling_frame_capture) to set your project scheme’s **GPU Frame Capture** setting to **Metal**.
5. **Optional - If you have the Unity Hub installed:** Go to the **Arguments** tab in the Xcode Scheme settings window and pass the path to your Unity project as an argument using `-projectPath`. This is to prevent the Unity Hub from opening when Xcode launches the Editor and instead directly debug the Editor with your project.
   ![The FrameDebug window displaying the project path](../uploads/Main/xcode-frame-debugger-cmd.png)
6. Run the Xcode project to launch the Unity Editor.
7. Perform a frame capture using one of the following methods:
   * **Unity Editor:** Use the **Xcode Capture** button to the right of the **Scene** view or Game view to perform a frame capture.
   * **Xcode:** Select the **Frame Capture** button (camera icon) to capture the next frame of data.
   * **Script:** Use the `FrameCapture` API. For more information, see the [FrameCapture API](https://docs.unity3d.com/ScriptReference/Apple.FrameCapture.html) documentation.
     For information on analyzing frame capture data in Xcode, see Apple’s [Xcode frame debugger documentation](https://developer.apple.com/documentation/metal/frame_capture_debugging_tools).

### Capture a frame from the Unity Editor with the command line

This workflow describes how to launch the Unity Editor from the command line, use the [FrameCapture API](../ScriptReference/Apple.FrameCapture.md) to perform a frame capture, and save the frame capture to disk. This workflow is supported on macOS.

1. Add calls to [FrameCapture.BeginCaptureToFile](../ScriptReference/Apple.FrameCapture.BeginCaptureToFile.md) and [EndCapture](../ScriptReference/Apple.FrameCapture.EndCapture.md) to your code, so that you can perform frame captures as required.
2. [Launch the Unity Editor from the command line](CommandLineArguments.md), with the following flag: `-enable-metal-capture`. Xcode performs frame captures when your code requests them, and saves the results to disk.

For information on how to analyze this data in Xcode, see the [Xcode frame debugger documentation](https://developer.apple.com/documentation/metal/frame_capture_debugging_tools).

## Additional resources:

* [XcScheme API documentation](../ScriptReference/iOS.Xcode.XcScheme.md)
* [FrameCapture API](../ScriptReference/Apple.FrameCapture.md)
* [Launch the Unity Editor from the command line](CommandLineArguments.md)