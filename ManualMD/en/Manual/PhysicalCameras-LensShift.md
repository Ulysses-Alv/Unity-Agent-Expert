# Widen the view with Lens Shift

**Lens Shift** offsets the **camera**’s lens from its sensor horizontally and vertically. This allows you to change the focal center, and reposition a subject in the rendered frame, with little or no distortion.

This technique is common in architectural photography. For example, if you want to capture a tall building, you could rotate the camera. But that distorts the image, making parallel lines appear to converge.

![Rotating the camera up to capture the top of the building causes vertical lines to converge](../uploads/Main/LensShift_VRot.png)

Rotating the camera up to capture the top of the building causes vertical lines to converge

If you shift the lens up instead of rotating the camera, you can change the composition of the image to include the top of the building, but parallel lines stay straight.

![Shifting the lens along the Y axis moves the focal center, but keeps vertical lines straight](../uploads/Main/LensShift_VShift.png)

Shifting the lens along the Y axis moves the focal center, but keeps vertical lines straight

Similarly, you can use a horizontal lens shift to capture wide objects without the distortion you might get by rotating the camera.

![Rotating the camera to frame the building causes horizontal lines to converge. ](../uploads/Main/LensShift_HRot.png)

Rotating the camera to frame the building causes horizontal lines to converge.

![Shifting the lens horizontally reframes the image, but keeps the horizontal lines straight.](../uploads/Main/LensShift_HShift.png)

Shifting the lens horizontally reframes the image, but keeps the horizontal lines straight.

### Lens shifts and frustum obliqueness

One side effect of a lens shift is that it makes the camera’s [view frustum](UnderstandingFrustum.md) oblique. That means the angle between the camera’s center line and its frustum is smaller on one side than on the other.

![The images above show the camera frustuim before (left) and after (right) a Y-axis lens shift. Shifting the lens up makes the frustum oblique.](../uploads/Main/ObliqueFrustum_LensShift.png)

The images above show the camera frustuim before (left) and after (right) a Y-axis lens shift. Shifting the lens up makes the frustum oblique.

You can use this to create visual effects based on perspective. For example, in a racing game, you might want to keep the perspective low to the ground. A lens shift is a way of achieving an oblique frustum without scripting.

For further information, see documentation on [Using an Oblique Frustum](ObliqueFrustum.md).