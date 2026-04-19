# Change the render type of a camera in URP

Use a **Camera**’s **Render Type** property to make it a Base Camera or an Overlay Camera.

To change the type of a Camera in the Unity Editor:

1. Create or select a Camera in your **scene**.
2. In the Camera **Inspector**, use the **Render Type** drop-down menu to select a different type of Camera. Select either:

   * **Base** to change the Camera to a Base Camera
   * **Overlay** to change the Camera to an Overlay Camera

You can change a Camera’s type in a script, by setting the `renderType` property of the Camera’s [Universal Additional Camera Data](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@latest/index.html?subfolder=/api/UnityEngine.Rendering.Universal.UniversalAdditionalCameraData.html) component, like this:

```
var cameraData = camera.GetUniversalAdditionalCameraData();
cameraData.renderType = CameraRenderType.Base;
```