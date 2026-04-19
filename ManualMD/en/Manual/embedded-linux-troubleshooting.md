# Troubleshooting the Embedded Linux Unity Editor

This page lists the common problems that might occur when using the Embedded Linux Unity Editor.

## The Player build fails with an error

When exporting the Unity project, the following error appears: `No EmbeddedLinux Burst Support on X86/Arm32 architecture`. Disabling Burst in Settings and removing it from the project doesn’t fix the issue because another package has a dependency on it.

### Solution

Burst isn’t supported on 32-bit Embedded Linux platforms. To prevent facing the error, you can disable Burst by starting the Unity Editor with the `--burst-disable-compilation` argument.

## Wayand is not available error

Starting the player on the target device fails with: SDL `Error: wayland,x11 not available` although Wayland is available.

### Solution

This error might occur due to one of the following reasons:

* The Wayland library isn’t located by SDL2. Make sure that the Wayland library is locatable by `dlopen` from the player application.
* Unity requires at least Wayland version 1.18. Because some systems have only version 1.16 or lower available, ensure to check that at least Wayland version 1.18 is supported on the target device.
* The connection to the Wayland display fails. Make sure the Wayland environment is properly set up.

## Rendering issues with MSAA on i.MX8 Embedded Linux devices

Multisample Anti-aliasing (MSAA) is known to cause rendering issues on i.MX8 Embedded Linux devices.

### Solution

To avoid rendering issues on i.MX8 devices, disable MSAA in your project’s quality settings. Go to **Edit** > **Project Settings** > **Quality**, then set **Anti Aliasing** to **Disabled**.

## Additional resources

* [Troubleshooting the Linux Editor issues](https://docs.unity3d.com/Documentation/Manual/linux-editor-troubleshooting.html)