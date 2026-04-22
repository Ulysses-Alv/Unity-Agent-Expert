# Control a camera in first person

Navigate through the **Scene** view while you look through a camera. In first-person, work through the lens of a camera to better frame your shots.

You can use the [Cameras overlay](cameras-overlay.md) to select and take first-person control of a **GameObject** that has a Camera or [Cinemachine camera](https://docs.unity3d.com/Packages/com.unity.cinemachine@latest/) component attached to it.

While you control a camera, you can use Editor tools as you do when you use the Scene Camera. For example, select a GameObject and press the F key to frame the camera on a specific GameObject. You can also adjust the position, orientation, and scale of a GameObject while you look through a camera to change the composition of your shot.

You can adjust the overscan of cameras you directly control with the Cameras overlay. Use overscan to intentionally see more or less of the Scene in the camera’s view than what the final output of the camera produces.

The Cameras overlay supports Timeline and Animation camera path authoring and Animated cameras. Control a camera in first person to animate cameras and generate **keyframes** for their GameObjects.

To control a camera in the first-person view:

1. Press **`** to open the [Overlays menu](overlays.md).
2. In the Overlays menu, enable the Cameras overlay.
3. In the Cameras overlay dropdown list, select a camera you want to control in first person.
4. Select **Control selected camera in first person**.
5. To adjust overscan size and opacity, select **Configure overscan settings** and do the following:
   * To select a size for the view guide, enter a value for **Overscan** or use the scroll wheel on your mouse. 1 is the default value. Enter a value greater than 1 to see more than the camera frustum. Enter a value below 1 to zoom in and see less than the camera frustum.
   * To adjust how opaque the view guide is, enter a value for **Overscan Opacity**.
6. Use the [Scene view navigation controls](SceneViewNavigation.md#tools) to move the camera through the scene.
7. To exit the camera view, in the Cameras overlay, select **Return to Scene Camera**.

## Additional resources

* [Overlays](overlays.md)
* [Cameras overlay](cameras-overlay.md)
* [Cameras](CamerasOverview.md)
* [Cinemachine](https://docs.unity3d.com/Packages/com.unity.cinemachine@latest/)
* [Timeline](https://docs.unity3d.com/Packages/com.unity.timeline@latest/)