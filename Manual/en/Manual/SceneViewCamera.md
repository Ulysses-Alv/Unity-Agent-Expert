# Scene Camera reference

The Camera settings menu contains options for configuring the **Scene** Camera. These adjustments do not affect the settings on **GameObjects** with Camera components.

To open the Scene Camera settings menu, click the Camera icon in the View options overlay in the **Scene view**.

You can also configure the Scene Camera in script with the [SceneView.CameraSetting](../ScriptReference/SceneView.CameraSettings.md) API.

**Tip**: To reset the properties to their default values, click the cog icon in the top right corner of the Scene Camera settings menu and select **Reset**.

![The Camera settings menu in context of the Scene view toolbar](../uploads/Main/SceneViewCameraSettings.png)

The Camera settings menu in context of the Scene view toolbar

| **Property** | **Description** |
| --- | --- |
| **Field of View** | Change the height of the Scene Camera’s view angle. |
| **Dynamic Clipping** | Enable to have the Editor calculate the Camera’s near and **far clipping planes** relative to the **viewport** size of the Scene. |
| **Clipping Planes** | Set the distances from the Camera where Unity starts and stops rendering GameObjects in the Scene. These distances apply to both [perspective and orthographic (isometric) projection modes](SceneViewNavigation.md#orientation-overlay). The **Near** and **Far** properties are modifiable only when **Dynamic Clipping** is disabled. |
|  | **Near:** Set the closest point to the Camera that the Editor renders GameObjects. |
|  | **Far:** Set the furthest point from the Camera that the Editor renders GameObjects. |
| **Occlusion Culling** | Enable occlusion culling in the Scene view. This prevents the Editor from rendering GameObjects that the Camera cannot see because they are hidden behind other GameObjects. |
| **Camera Easing** | Make the Camera ease in and out of motion in the Scene view. This makes the Camera ease into motion when it starts moving instead of starting at full speed, and ease out when it stops. You can set the [duration](../ScriptReference/SceneView.CameraSettings-easingDuration.md) in the API. |
| **Camera Acceleration** | Enable acceleration when moving the camera. When enabled, the camera initially moves at a speed based on the speed value, and continuously increases speed until movement stops. When disabled, the camera accelerates to a constant speed based on the **Camera Speed**. |
| **Camera Speed** | Set the current speed the Scene camera uses in the Scene view. In [Flythrough mode](SceneViewNavigation.md#flythrough), you can change the speed of the Camera while moving. To do this, use the mouse scroll wheel or drag two fingers on a trackpad. |
|  | **Min**: Set the minimum speed of the Camera in the Scene view. Valid values are between 0.01 and 98. |
|  | **Max**: Set the maximum speed of the Camera in the Scene view. Valid values are between 0.02 and 99. |

## Additional resources

* [Cameras overlay](cameras-overlay.md)
* [Control a camera in first person](control-camera.md)
* [Cameras](CamerasOverview.md)
* [Cinemachine](https://docs.unity3d.com/Packages/com.unity.cinemachine@latest/)
* [Timeline](https://docs.unity3d.com/Packages/com.unity.timeline@latest/)