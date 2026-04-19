# Set up your environment for Embedded Linux

Set up your environment to develop with Embedded Linux.

To create a Unity application for Embedded Linux, you first need to set up your Unity project to support Embedded Linux. To support Embedded Linux, a Unity project requires certain packages and dependencies.

**Note:** You must install the Embedded Linux platform package before you can set up your environment. For more information, refer to [Install the platform package for Embedded Linux](embedded-linux-install-editor.md).

## Install the toolchain and SDK packages

After you create a new project with Unity, you must install the following toolchain and SDK packages for your operating system and target architecture. To install a Unity package for Embedded Linux toolchain and SDK, refer to [Install a UPM package by name](upm-ui-quick.md).

**Note**: When you install Sysroot and toolchain pacakges, the base sysroot package `com.unity.sysroot.base` is installed as a dependency. The base sysroot package contains common code shared between dependent sysroot and toolchain packages.

### Toolchain packages

| **Operating system** | **Package** |
| --- | --- |
| **Windows** | `com.unity.toolchain.win-x86_64-linux` |
| **Windows(Arm64)** | `com.unity.toolchain.win-arm64-linux` |
| **macOS** | `com.unity.toolchain.macos-x86_64-linux` |
| **macOS(Arm64)** | `com.unity.toolchain.macos-arm64-linux` |
| **Linux** | `com.unity.toolchain.linux-x86_64-linux` |

When you install a toolchain package, ensure to install the appropriate sysroot package based on the intended target platform. For example, when you install the `com.unity.toolchain.linux-x86_64-linux` toolchain package, ensure to install the corresponding sysroot package `com.unity.sdk.linux-x86_64`.

### SDK packages

| **Target architecture** | **Package** |
| --- | --- |
| **aarch64** | `com.unity.sdk.linux-arm64` |
| **x86\_64** | `com.unity.sdk.linux-x86_64` |

When you install a sysroot package, ensure to install the appropriate toolchain package based on the intended target platform. For example, when you install the `com.unity.sdk.linux-x86_64` sysroot package, ensure to install the corresponding `com.unity.toolchain.linux-x86_64-linux` toolchain package.

After you install the packages, the Package Manager window displays the toolchain and SDK packages that are installed for Embedded Linux.

![Package Manager window with Embedded Linux toolchain package](../uploads/Main/embedded-linux-proj-setup-1.png)

Package Manager window with Embedded Linux toolchain package

### Automated packages installation

You can also install the required Embedded Linux packages automatically from the ****Build Profiles**** window.

To initiate the automatic installation of packages:

1. Go to **File** > **Build Profiles** from Unity’s main menu.
2. In the **Build Profiles** window, select **Add Build Profile** to open the **Platform Browser** window.
3. In the **Platform Browser** window, select **Embedded Linux** and then click **Add Build Profile**.
4. Go to **Edit** > **Project Settings** from Unity’s main menu.

If the Unity Editor has the **Install toolchain packages automatically** option enabled by default, then the required packages such as the toolchain and SDKs for installed architectures are installed automatically. After the package installation process is complete, you can check which packages are installed in the **Toolchain Management (Embedded Linux)** tab in **Project Settings**.

![Installed packages in Toolchain Management (Embedded Linux)](../uploads/Main/embedded-linux-proj-setup-3.png)

Installed packages in Toolchain Management (Embedded Linux)

If the option to automatically install packages is disabled, you can switch to the **Toolchain Management (Embedded Linux)** tab and click **Install sdk and toolchain packages**.

After you’ve installed the packages, the Package Manager window displays the list of all the toolchain packages that are installed for Embedded Linux.

![Install sdk and toolchain packages button in Toolchain Management (Embedded Linux)](../uploads/Main/embedded-linux-proj-setup-2.png)

Install sdk and toolchain packages button in Toolchain Management (Embedded Linux)

## Dependencies

Embedded Linux must provide direct and **indirect dependencies** so Unity can run correctly.

### Direct dependencies

Direct dependencies load at the application startup.

* `libm.so.6`
* `libgcc_s.so.1`
* `libpthread.so.0`
* `libc.so.6` (glibc 2.35 or later)
* `libdl.so.2`
* `librt.so.1`

### Indirect dependencies

Indirect dependencies load when needed during the application runtime as a shared library.

| **Type** | **Dependencies** |
| --- | --- |
| **Audio** | `libpulse-simple.so.0` `libpulse.so.0` `libesd.so.0` `libasound.so or libasound.so.2` |
| **Wayland** | `libwayland-client.so.0` `libwayland-egl.so.1` `libwayland-cursor.so.0` `libxkbcommon.so.0` |
| **X11** | `libX11.so.6` `libXext.so.6` `libXcursor.so.1` `libXinerama.so.1` `libXi.so.6` `libXrandr.so.2` `libXss.so.1` `libXxf86vm.so.1` |
| **OpenGL ES** | `libEGL.so.1` `libGLESv2.so.2` |
| **Vulkan** | `libvulkan.so.1` |
| **Devices** | `libudev.so.1` |

## Additional resources

* [System requirements](system-requirements.md)
* [Develop for Embedded Linux](embedded-linux-develop.md)