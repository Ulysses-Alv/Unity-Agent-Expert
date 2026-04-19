# Build a Linux application

To build your Unity application on the Linux platform, use the following steps:

1. Open the [Build Profiles window](BuildSettings.md) from **File** > **Build Profiles**.
2. Select **Add Build Profile** to open the **Platform Browser** window.
3. Select **Linux** from the list of available platforms and set the required build settings. If **Linux** isn’t an option, select **Install with Unity Hub** and follow the installation instructions. For information on how to install modules, refer to [Add modules](https://docs.unity3d.com/hub/manual/AddModules.html).
4. Select **Switch Profile** to set the new build profile as the active profile.
5. Select **Build** or **Build And Run**. For more information on these options, refer to [Build your application](build-profiles-reference.md#build-your-application).
6. In the Linux file chooser window, select the destination for Unity to place the build.
7. In the **Name** field, enter an appropriate name for the build.
8. Select **Save**. This starts the build process.

## Linux Player build binaries

When you build a Unity application on the Linux platform, Unity produces the following files, where `ProjectName` is the name of your application:

* `ProjectName.x86_64`: This is the project executable file for your application. It contains the program entry point that initiates the Unity engine when launched.
* `UnityPlayer.so`: This `.so` file contains all the native Unity engine code. It’s signed with the Unity Technologies certificate allowing you to verify that no malicious entities have tampered with your engine code.
* `*.pdb` files: These are symbol files you can use to debug managed (C#) code. Unity copies these files to the build directory if you enable ****Development Build**** in the **Platform Settings** section of the [Build Profiles](Buildsettings-linux.md) window.
* `*_s.debug` files: These are symbol files you can use to debug native (C/C++) code. Unity copies these files to the build directory if you enable **Development Build** in the **Platform Settings** section of the [Build Profiles](Buildsettings-linux.md) window.
* `ProjectName_Data` folder: This folder contains all the data needed to run your application.
* `libdecor-0.so.0`: This `.so` file is a Linux shared library used by Wayland clients to manage client-side window decorations. It ensures consistent window behavior and appearance across various compositors.
* `libdecor-cairo.so`: This `.so` file is a Linux shared library that integrates Cairo graphics with `libdecor`. It renders client-side decorations in Wayland to enhance visual consistency and performance.

If you’re using the **IL2CPP** **scripting backend**, your application Player build also includes the following file and folder:

* `GameAssembly.so`: This `.so` file contains all managed (C#) game logic and **scripts** converted into native code (C/C++) for enhanced performance.
* `ProjectName_BackUpThisFolder_ButDontShipItWithYourGame` folder: This folder contains intermediate files generated during IL2CPP builds that are useful for debugging rather than distribution.

## Additional resources

* [Linux build settings reference](Buildsettings-linux.md)