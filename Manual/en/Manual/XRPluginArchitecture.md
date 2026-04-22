# XR architecture

Unity supports **XR** development through its **plug-in** framework and a set of feature and tool packages. Go to the [XR Plug-in Management](xr-plugin-management.md) category in ****Project Settings**** to enable XR support in a Unity project and to choose the plug-ins for the XR platforms that your project supports. Use the Unity Package Manager to install the additional feature packages.

The following diagram illustrates the current Unity XR plug-in framework structure and how it works with platform provider implementations:

![The Unity XR plug-in framework structure](../uploads/Main/unity-xr-tech-stack.png)

XR subsystems define a common interface for XR features. XR plug-ins implement these subsystem interfaces to provide data to the subsystem at runtime. Your XR application can access data for an XR feature through Unity Engine and package APIs.

## XR provider plug-in framework

An XR provider plug-in is a Unity plug-in that supports one or more XR device platforms. For example, the ARCore plugin supports the Android **AR** platform on hand-held Android devices, while the OpenXR plug-in supports several XR devices on multiple operating systems.

An XR provider plug-in implements interfaces called *subsystems*. A plug-in that implements one or more subsystems is called a *provider* plug-in. Typically, a provider plug-in uses the device platform’s native libraries to implement the Unity interfaces for their devices.

Unity uses subsystem interfaces to communicate with providers for various platforms, powering the XR features of your application. Because of these interfaces, you can reuse the same feature code in your application across all XR devices that have a provider for that feature.

### Subsystems

XR subsystems give you access to XR features in your Unity app. Unity defines a common interface for XR subsystems so that all provider plug-ins implementing a feature generally work the same way in your app. Often you can change the active provider and rebuild your app to run on a different XR platform, as long as the platforms are largely similar.

The Unity Engine defines a set of fundamental XR subsystems. Unity packages can provide additional subsystems. For example, the AR Subsystems package contains many of the AR-specific subsystem interfaces.

The subsystems defined in the Unity Engine include:

| Subsystem | Description |
| --- | --- |
| [Display](../ScriptReference/XR.XRDisplaySubsystem.md) | Stereo XR display. |
| [Input](../ScriptReference/XR.XRInputSubsystem.md) | Spatial tracking and controller input. |
| [Meshing](../ScriptReference/XR.XRMeshSubsystem.md) | Generate 3D meshes from environment scans. |

**Note:** Unity applications typically do not interact with subsystems directly. Instead, the features provided by a subsystem are exposed to the application through an XR plug-in or package. For example, the [ARMeshManager](https://docs.unity3d.com/Packages/com.unity.xr.arfoundation@6.2/api/UnityEngine.XR.ARFoundation.ARMeshManager.html) component in the AR Foundation package lets you add the meshes created by the Meshing subsystem to a **scene**.