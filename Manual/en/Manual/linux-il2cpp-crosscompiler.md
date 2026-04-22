# Linux IL2CPP cross-compiler

The Linux **IL2CPP** cross-compiler is a set of sysroot and toolchain packages that allow you to build Linux IL2CPP Players on any Standalone platform without needing to use the Linux Unity Editor or rely on Mono.

If you meet the prerequisites, Unity automatically installs these packages for you when you choose the Linux build target. If you want to opt out of this process and use your own sysroot and toolchain packages, go to **Edit** > **Project Settings** > **Toolchain Management** and disable the **Install Toolchain package automatically** checkbox. If you already have these installed, you also need to remove them from the Package Manager.

## Prerequisites

Unity needs the following to install the IL2CPP cross-compiler packages:

* Enough available disk space for your chosen Linux toolchain package. For further information, refer to the [Required disk space for Linux toolchain packages](#requiredDiskSpace).
* Scripting back end set to IL2CPP. To set the scripting back end to IL2CPP: Go to **Edit** > **Project Setting** > **Player Settings for Linux** > **Other Settings** > **Configuration**. Set the **Scripting Backend** to **IL2CPP**.
* IL2CPP module. For information on how to install the IL2CPP module, follow the steps documented on [Adding modules](https://docs.unity3d.com/hub/manual/AddModules.html).

## Linux sysroot package

A Linux sysroot package is a directory which includes all the headers and libraries you need to build for Linux.

Every operating system (OS) has its own build systems which vary from one to another. If you build using the headers and libraries of a particular OS, the built Player might not run on other operating systems. To address this, Unity provides a sysroot to build with which works on all supported Linux platforms.

## Linux toolchain packages

Unity provides toolchain packages for macOS, Windows, and Linux. Each of these platforms builds for Linux in a unique way.

A Linux toolchain package is a set of tools (including the compiler and linker) that Unity needs to build for Linux from each of these operating systems.

## Required disk space for Linux toolchain packages

Make sure you have enough disk space to account for the package download, decompression, and use.

In the rare instances where you are unsure whether you have enough space, define a `UNITY_SYSROOT_CACHE` environment variable and use it to store the uncompressed sysroots and toolchain packages. An environment variable is a variable that you set outside of Unity which is available for Unity to reference. In this case, you set a cache that Unity can reference when decompressing the sysroot and toolchain packages. Environment variables are specific to your operating system, so you need to follow your system’s guidelines to set them.

The table below shows the total disk space requirements for each toolchain package:

| Toolchain package | Required disk space |
| --- | --- |
| `com.unity.toolchain.linux-x86_64` | 462 MB |
| `com.unity.toolchain.macos-x86_64-linux-x86_64` | 2 GB |
| `com.unity.toolchain.win-x86_64-linux-x86_64` | 2 GB |

## Using the Linux IL2CPP cross-compiler

If you meet all the prerequisites on this page, you can build your project as a Linux Player. Unity automatically uses the Linux IL2CPP cross-compiler at build time.

To build a Linux Player, follow these steps:

1. Open the **Build Profiles** window (menu: **File** > **Build Profiles**).
2. From the list of platforms in the **Platforms** panel, select **Linux** or [create a build profile](create-build-profile.md) for the **Linux** platform.
3. Select **Switch Platform**.
4. Select **Build** or **Build And Run**.

## Additional resources

* [How to add modules in the Unity Editor](https://docs.unity3d.com/hub/manual/AddModules.html)
* [Linux build settings reference](Buildsettings-linux.md) platform docs.
* [Linux Player settings](PlayerSettings-linux.md)
* [IL2CPP additional arguments reference](handling-il2cpp-additional-args.md)