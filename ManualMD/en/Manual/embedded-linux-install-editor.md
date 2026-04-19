# Install the platform package for Embedded Linux

Install the platform package for Embedded Linux on Windows, Linux, and macOS.

**Note:** The Unity Editor for Embedded Linux is available under separate terms. For licensing and download information, contact your Account Manager or the [Unity Sales](https://create.unity.com/unity-for-industries) team.

## Requirements and compatibility

Your system must meet the [System requirements for the Unity Editor](system-requirements.md#editor).

## Installation

To install the platform package for Embedded Linux on your system, follow these steps:

1. On the [Unity Embedded Linux releases page](https://share.unity.com/receive/?packageCode=2Ss9L8RG3gTRMl006iljdTBnHIGSjvEW5tapO7oKdwE#keycode=h60EllPSliiFduyHnNJx0dq2n82rIfAK8kldTFbV7o0), go to **EmbeddedLinux builds**.  
   **Note:** If you’re unable to access the Embedded Linux releases page, contact your Account Manager or the [Unity Sales](https://create.unity.com/unity-for-industries) team.
2. Select the folder for the specific Unity version you want to work with.
3. Download the setup installer for your operating system.

   | **Operating system** | **Installer file extension** |
   | --- | --- |
   | **Windows** | .exe |
   | **Linux** | .xz |
   | **macOS** | .pkg |
4. Run the installer to install to a folder of your choice.
5. On the [Unity download archive](https://unity.com/releases/editor/archive), install the corresponding Unity 6 Editor version in the same folder through the [Unity Hub](unity-hub.md).

**Note:** The Editor version must match the setup installer version for the Editor to work correctly.

The Editor for Embedded Linux is now installed for your Windows or Linux systems. Next, set up your environment for [Embedded Linux](embedded-linux-environment-setup.md). However, if you’re using macOS, follow the steps in [Additional instructions for macOS installation](#macOS).

### Additional instructions for macOS installation

On macOS, the Unity Editor and platform package for Embedded Linux require additional steps because the default path where Unity installs the package might not work.

1. Move the `/Applications/Unity/Hub/Editor/<version>/` folder to `/Applications/Unity/Unity-<version>/`. For example, if using Unity Editor version 6000.0.58f2, move the contents of `/Applications/Unity/Hub/Editor/6000.0.58f2/` to `/Applications/Unity/Unity-6000.0.58f2/`.
2. Open the `UnitySetup-Embedded-Linux-Support-for-Editor-<version>` setup installer `.pkg` file and follow the installation prompts.
3. In Unity Hub, go to the **Installs** tab.
4. Select the **Manage** option for the corresponding Unity version and choose **Remove from Hub**.
5. From the **Installs** tab, select **Locate** and navigate to the new location in the `/Applications/Unity/` folder.

The Unity Editor and platform package for Embedded Linux now works correctly on macOS. Next, [set up your environment for Embedded Linux](embedded-linux-environment-setup.md).

## Additional resources

* [Build for Embedded Linux from the command line](embedded-linux-build-command-line.md)