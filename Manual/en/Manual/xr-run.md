# Run an XR application

Running an **XR** app developed in Unity works the same as running any type of app on a device. In most cases, you can use the Unity **Build and Run** function to both build and run an app on a connected device.

**Note:** The method you use to install an already built app on a device varies by platform, and you should refer to the platform’s documentation for detailed information.

## Build and Run

To build your game or application and run it on a device:

1. Connect the target device to your build computer.
2. Open the **Build Profiles** window (menu: **File > Build Profiles**).
3. If necessary, choose the target platform and click **Switch platform**.
4. Choose the connected XR device from the **Run device** options in the build settings. (Not required on all platforms.)
5. Ensure the other build settings are correct for your chosen platform. Refer to [Build profiles](BuildSettings.md) for more information.
6. Click **Build And Run**.

**Tip:** After you have configured your build settings, you can build and run your project directly with the **File > Build And Run** menu command.

Refer to [Publishing profiles](BuildSettings.md) for more information about building Unity projects.

## Play mode

On the Windows platform, some XR provider **plug-ins** support a “hybrid” play mode in which the project runs on a connected XR device when you enter Play mode. The Unity **Game view** mirrors the headset display. You have the choice of the following options for the Game view:

* **Left Eye**: shows the left eye only
* **Right Eye**: shows the right eye only
* **Both Eyes**: shows both eyes, side by side
* **Occlusion **mesh****: shows the both eyes, side by side, mapped onto a mesh similar to that used by headsets to compensate for lens distortion.

**Tip:** When developing for the Quest 2 or Pro, you can switch to the Windows platform and use the Quest Link to take advantage of the faster iteration turn around provided by Play mode compared to building and uploading your project to the device.

### Mock HMD

The Mock HMD package provides a simulated HMD in the Unity **Game view** in Play mode. The Mock HMD is a provider plug-in that simulates the presence of a head-mounted device. Enable the Mock HMD in the [XR Plug-in Management](xr-plugin-management.md) settings section. Input and tracking are not simulated.

Refer to the [Mock HMD](https://docs.unity3d.com/Packages/com.unity.xr.mock-hmd@1.3/manual/index.html) documentation for more information.

**Tip:** The XR Interaction Toolkit package provides an [XR Device Simulator](https://docs.unity3d.com/Packages/com.unity.xr.interaction.toolkit@3.2/manual/xr-device-simulator.html) that translates keyboard and mouse input into movement and interaction.

### XR Simulation

The **AR** Foundation package provides an XR Simulation feature that you can use to test AR apps in the Editor. XR Simulation provides pre-built test environments that you can use in Play Mode to simulate how your app would behave in different physical settings. The simulation reports detected planes and other AR features as if you were navigating within a real environment. You can modify the provided environments and create your own.

See [XR Simulation](https://docs.unity3d.com/Packages/com.unity.xr.arfoundation@6.2/manual/xr-simulation/simulation.html) for more information.

## Additional resources

Most XR devices use one of the existing OS platforms. For additional information about running projects on these platforms, see:

* [Android](android.md)
* [iOS](iphone.md)
* [visionOS](visionOS.md)
* [Windows](Windows.md)