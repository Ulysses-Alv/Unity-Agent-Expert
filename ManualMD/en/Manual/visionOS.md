# visionOS

visionOS is the operating system of the Vision Pro, Apple’s latest spatial computing device. Unity developers can leverage existing 3D **scenes** and assets to build games or applications for visionOS. For more information about visionOS, refer to Apple’s [visionOS Overview](https://developer.apple.com/visionos/learn).

visionOS provides a few different modes in which apps can be displayed: Windows, Volumes or Spaces. Users can use Windows to present 2D or 3D content (without stereo), or use Volumes to present 3D content and objects. With Volumes, users have the flexibility to walk around and interact with 3D content from any angle.

Depending on application type, visionOS apps can run in either a **Shared Space** or a **Full Space**. The Shared Space works as a multitasking environment similar to the desktop of a personal computer; in this mode, users can see and interact with Windows and Volumes from multiple applications simultaneously. For more immersive experiences, you can run your application in a Full Space, which only displays content from one app at time. Windowed apps always run in a Shared Space. Fully immersive (VR) content always run in a Full Space. Immersive (MR) content can run in either a Shared Space or a Full Space, and switch between them, depending on the **Mode** setting of its Volume **Camera** Window Configuration.

## Requirements

Fully Immersive (VR) and Immersive (MR) apps are only available to Unity Pro, Unity Enterprise, and Unity Industry users. [Learn more about these plans](https://unity.com/pricing).

## Getting Started

All visionOS application types require that you install the visionOS module to the Unity Editor via the [Unity Hub](https://docs.unity3d.com/hub/manual/AddModules.html). The visionOS module is available for Unity 2022.3.5f1+.

In addition, Fully Immersive (VR) and Immersive (MR) apps require that you install extra packages. Refer to [Install PolySpatial and visionOS support](https://docs.unity3d.com/Packages/com.unity.polyspatial.visionos@1.2/manual/Install.html) for instructions.

For Fully Immersive (VR) apps, you need:

* `com.unity.xr.visionos`

For Immersive (MR) apps, you need:

* `com.unity.xr.visionos`
* `com.unity.polyspatial.visionos`
* `com.unity.polyspatial.xr`

If you do not install these packages, you can only build Windowed apps.

Once you have installed the necessary modules and packages, navigate to **Project Settings > XR Plug-in Manager > Apple visionOS > App Mode** and select between **Mixed Reality - Volume or Immersive Space** and **Virtual Reality - Fully Immersive Space** to configure your build for the desired application mode.

## Windowed apps

Unity describes applications that live within Windows as Windowed apps. By default, your Unity content is built as a Windowed app on the visionOS platform if you don’t install and enable the visionOS **XR** provider **plug-in** in the **XR Plug-in Management** section of your ****Project Settings****. Refer to [Choose XR provider plug-ins](XR.md) for more information about enabling provider plug-ins.

To detect interactions on the Unity application within a window component on visionOS, use the Input System package’s ([com.unity.inputsystem](https://docs.unity3d.com/Packages/com.unity.inputsystem@1.6/manual/index.html)) touch support. To learn more about designing for the window component on visionOS, visit [Apple’s Human Interface Guidelines for visionOS](https://developer.apple.com/design/human-interface-guidelines/windows#visionOS).

## Mixed Reality (Immersive) and Virtual Reality (Fully Immersive) Apps

For more information about developing fully immersive **virtual reality** (VR) apps and immersive **mixed reality** (MR) apps for visionOS, refer to Unity’s [visionOS-specific documentation](https://docs.unity3d.com/Packages/com.unity.polyspatial.visionos@latest).