# XR Origin

The ****XR** Origin** serves as the center of tracking space in an XR **scene**.

An **XR Origin** is a set of **GameObjects** and components that work together to transform XR tracking data into the scene world space.

The following topics discuss the XR Origin and how to use it in your project:

| Topic | Description |
| --- | --- |
| [XR tracking space](#xr-tracking-space) | Describes the relationship between the **XR Origin**, the device tracking space and the Unity scene. |
| [XR Origin configurations](#xr-origin-configurations) | Describes the different **XR Origin** configurations available in Unity. |

Refer to [XR Origin component](https://docs.unity3d.com/Packages/com.unity.xr.core-utils@2.5/manual/xr-origin-reference.html) for more details about the **XR Origin** and its child GameObjects and components.

## XR tracking space

XR devices choose an origin point in the real world at initialization. The positions and orientations of all tracked entities, such as the user’s headset or hand-held device, XR controllers, hands, and physical objects detected around the user, are reported relative to this chosen point.

**Note:** The criteria for choosing the initial origin varies by platform. Typically, a device chooses a point at or directly below the user’s HMD (VR) or hand-held device (AR).

If you used the tracking data directly in your scene, the user would appear to be standing at the scene origin point (0, 0 ,0). To make the user appear at a different location in the scene, you need to transform the tracking data to the desired position and orientation. Unity provides the **XR Origin** to do this automatically. Many Unity XR features, including those provided by [AR Foundation](https://docs.unity3d.com/Packages/com.unity.xr.arfoundation@6.2/) and the [XR Interaction Toolkit](https://docs.unity3d.com/Packages/com.unity.xr.interaction.toolkit@3.2/), require an **XR Origin** in the scene. You can choose from a variety of [XR Origin configurations](#xr-origin-configurations) to suit your project.

To position the XR tracking space in a Unity scene, add an **XR Origin** GameObject at the location where you want the user to appear when the scene starts. For example, to place the user at the scene origin, place the **XR Origin** GameObject there. You can rotate the **XR Origin** around its y axis to face the user in the desired starting direction.

The **XR Origin** contains GameObjects representing tracked entities as children within its hierarchy. For example, the user’s headset or hand-held device is represented by the **Camera** GameObject. Because they are children, the tracking data controlling the positions and rotations of these GameObjects is automatically transformed into world space relative to the **XR Origin**. When the user moves in the real world, these child GameObjects move relative to the **XR Origin** in the scene.

![Man wearing MR headset reaches toward a virtual object](../uploads/Main/xr_origin_diagram.webp)
  
*By using the **XR Origin** (**A**) as the parent for transforming tracking data, the tracking origin and the Unity scene origin (**B**) can be in different places and orientations.*

The parent **XR Origin** GameObject doesn’t move when the user walks around the scene. However, you can move the **XR Origin** with a script to allow the user to teleport or navigate around the scene via controller input.

## XR Origin configurations

The Unity XR packages provide several **XR Origin** configurations tailored for different types of XR applications. You can use the **GameObject > XR** menu to add an **XR Origin** to the current Scene. The available options depend on which packages you have added to your project.

**Important:** You should never have more than one active **XR Origin** in a scene. If you need different configurations of the **XR Origin** in a scene for different purposes, only enable one at the same time.

| XR type | Configuration | Menu option | Package | Notes |
| --- | --- | --- | --- | --- |
| **VR**, **MR** | XR Rig | **Convert Main Camera To XR Rig** | [XR Legacy Input Helpers](https://docs.unity3d.com/Packages/com.unity.xr.legacyinputhelpers@2.1/) (installed with the XR Plug-in Management package) | Replaces the standard Camera in a basic Unity Scene. The logic used to replace the main camera can fail in complex scenes. Unity removes this option when you install the XR Interaction Tools package. The **XR Rig** GameObject created by this menu option is slightly different than the **XR Origin** and might not be as compatible with other Unity XR features. |
| VR, MR | XR Origin | **XR Origin** | [XR Core Utils](https://docs.unity3d.com/Packages/com.unity.xr.core-utils@2.5/) (installed with the XR Interaction Toolkit package) | Does not include GameObjects for controllers. |
| VR, MR | XR Origin (VR) | **XR Origin (VR)** | [XR Interaction Toolkit](https://docs.unity3d.com/Packages/com.unity.xr.interaction.toolkit@3.2/) | Includes controller GameObjects set up for action-based input. |
| VR, MR | XR Origin (VR) | **Device-based > XR Origin (VR)** | [XR Interaction Toolkit](https://docs.unity3d.com/Packages/com.unity.xr.interaction.toolkit@3.2/) | Includes controller GameObjects set up for device-based input. |
| **AR**, MR | XR Origin (AR) | **XR Origin (AR)** | [AR Foundation](https://docs.unity3d.com/Packages/com.unity.xr.arfoundation@6.2/) | Serves as the tracking origin for hand-held AR applications. Includes controller GameObjects. This option is available if you have the XR Interaction Toolkit installed. |
| AR | XR Origin (Mobile AR) | **XR Origin (Mobile AR)** | [AR Foundation](https://docs.unity3d.com/Packages/com.unity.xr.arfoundation@6.2/) | Serves as the tracking origin for hand-held AR applications. Does not include controller GameObjects. This option is replaced with **XR Origin (AR)** if you have the XR Interaction Toolkit installed. |

**Notes:**

* In the latest versions of the Unity XR packages, the **XR Rig** has been replaced with the **XR Origin**.
* The `XROrigin` component also replaces the `CameraOffset` component, providing additional settings. Refer to [XR Origin component](https://docs.unity3d.com/Packages/com.unity.xr.core-utils@2.5/manual/xr-origin-reference.html) for more information.
* The **AR Session Origin** configuration has been replaced with **XR Origin (Mobile AR)** and **XR Origin (AR)** in [AR Foundation 5.0+](https://docs.unity3d.com/Packages/com.unity.xr.arfoundation@6.2/).
* The **XR Origin (Mobile AR)** and **XR Origin (AR)** configurations are not compatible with earlier versions of the AR Foundation package.