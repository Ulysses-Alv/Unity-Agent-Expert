# Collecting performance data on an Android device

Use the [Profiler](profiler-introduction.md) to collect performance data about your application. You can [collect performance data while in Play mode](profiling-play-mode.md) in the Unity Editor. However, to get the most accurate data about your application, you can connect the Profiler directly to an Android device that’s on your network.

## Prerequisites

* If you’re using a firewall, open ports `54998` to `55511` in your firewall’s outbound rules. These are the ports Unity uses for remote profiling.
* Disable mobile data on the device
* Set the same subnet for the Android device and the host computer that’s running the Unity Editor for device detection to work.

## Enabling remote profiling

To enable remote profiling, follow these steps:

1. Connect the device to your WiFi network. The Profiler uses a local WiFi network to send profiling data from your device to the Unity Editor.
2. Attach your device to your computer via cable.
3. Open the **Build Profiles** window (menu: **File > Build Profiles**).
4. Enable the ****Development Build**** setting.
5. Enable the **Autoconnect Profiler** setting.
6. Select **Build & Run**.
7. When the application launches on the device, open the Profiler window in the Unity Editor (menu: **Window > Analysis > Profiler**).

After you open the Profiler window, it populates with data from your application. If the Editor doesn’t connect to the device automatically, select the Target Selection dropdown menu in the Profiler window and choose the appropriate device to start the Profiler connection manually.

You can also plug the target device directly into your computer to avoid network or connection issues.

## Profiling with Android Debug Bridge

Android devices support profiling through [Android Debug Bridge (adb)](https://developer.android.com/studio/command-line/adb). To profile with Android Debug Bridge (adb), follow these steps:

1. Put the device in Development mode and enable the USB debugging setting.
2. Attach the device to your computer via cable and make sure that it’s displayed in the **adb** devices list.
3. Open the Build Profiles window (menu: **File > Build Profiles**).
4. Enable the **Development Build** setting.
5. Select **Build & Run**.
6. When the application launches on the device, open the Profiler window (menu: **Window** > **Analysis** > **Profiler**).
7. From the Target Selection dropdown menu, select `AndroidProfiler(ADB@127.0.0.1:34999)`. The entry in the dropdown menu is only visible when the selected target is Android.

### Configuring Android Debug Bridge manually

The Editor automatically creates an adb tunnel for your application when you select **Build & Run**. You can configure this tunnel manually if you want to profile another application, or you restart the adb server.

To configure the tunnel manually:

1. Open a Terminal window or Command prompt.
2. Enter the following:
3. Required when Editor-to-Android connection is established via USB cable:  
   `adb forward tcp:34999 localabstract:Unity-{insert bundle identifier here}`
4. Required when Android-to-Editor connection is established via USB cable  
   `adb reverse tcp:34998 tcp:34999`

## Additional resources

* [Profiler introduction](profiler-introduction.md)
* [Using the Profiler window](ProfilerWindow.md)
* [Debug on an Android device](android-debugging-on-an-android-device.md)
* [Collect performance data in Play mode](profiling-play-mode.md)