# Introduction to the camera view

What a **camera** sees is defined by its transform and its camera component. The transform position defines the viewpoint, its forward (Z) axis points in the view direction, its upward (Y) axis points toward the top of the screen, and its rightward (X) axis points toward the right side of the screen. Settings on the camera component define the size and shape of the region that falls within the view. With these parameters set up, the camera can display what it currently “sees” to the screen. As the **GameObject** moves and rotates, the displayed view moves and rotates accordingly.

## The shape of the viewed region

Both perspective and orthographic cameras have a limit on how far they can “see” from their current position. The limit is defined by a plane that is perpendicular to the camera’s forward (Z) direction. This is known as the *far clipping plane* since objects at a greater distance from the camera are “clipped” (ie, excluded from rendering). There is also a corresponding *near clipping plane* close to the camera - the viewable range of distance is that between the two planes.

Without perspective, objects appear the same size regardless of their distance. This means that the viewing volume of an orthographic camera is defined by a rectangular box extending between the two **clipping planes**.

When perspective is used, objects appear to diminish in size as the distance from camera increases. This means that the width and height of the viewable part of the **scene** grows with increasing distance. The viewing volume of a perspective camera, then, is not a box but a pyramidal shape with the apex at the camera’s position and the base at the **far clipping plane**. The shape is not exactly a pyramid, however, because the top is cut off by the **near clipping plane**; this kind of truncated pyramid shape is known as a *frustum*. Since its height is not constant, the frustum is defined by the ratio of its width to its height (known as the *aspect ratio*) and the angle between the top and bottom at the apex (known as the *field of view* or *FOV*).

## The view frustum

The word **frustum** refers to a solid shape that looks like a pyramid with the top cut off parallel to the base. This is the shape of the region that can be seen and rendered by a perspective camera.

**Note:** By default, cameras perform frustum culling, which excludes all Renderers that don’t fall within the view frustum of the camera.

The following thought experiment should help to explain why this is the case.

Imagine holding a straight rod (a broom handle or a pencil, for example) end-on to a camera and then taking a picture. If the rod were held in the centre of the picture, perpendicular to the camera lens, then only its end would be visible as a circle on the picture; all other parts of it would be obscured. If you moved it upward, the lower side would start to become visible but you could hide it again by angling the rod upward. If you continued moving the rod up and angling it further upward, the circular end would eventually reach the top edge of the picture. At this point, any object above the line traced by the rod in world space would not be visible on the picture.

![Frustum view example](../uploads/Main/Rods.png)

Frustum view example

The rod could just as easily be moved and rotated left, right, or down or any combination of horizontal and vertical. The angle of the “hidden” rod simply depends on its distance from the centre of the screen in both axes.

The meaning of this thought experiment is that any point in a camera’s image actually corresponds to a line in world space and only a single point along that line is visible in the image. Everything behind that position on the line is obscured.

The outer edges of the image are defined by the diverging lines that correspond to the corners of the image. If those lines were traced backwards towards the camera, they would all eventually converge at a single point. In Unity, this point is located exactly at the camera’s transform position and is known as the centre of perspective. The angle subtended by the lines converging from the top and bottom centres of the screen at the centre of perspective is called the field of view (often abbreviated to FOV).

As stated above, anything that falls outside the diverging lines at the edges of the image will not be visible to the camera, but there are also two other restrictions on what it will render. The near and far clipping planes are parallel to the camera’s XY plane and each set at a certain distance along its centre line. Anything closer to the camera than the near clipping plane and anything farther away than the far clipping plane will not be rendered.

![The near and far clipping planes.](../uploads/Main/ViewFrustum.png)

The near and far clipping planes.

The diverging corner lines of the image along with the two clipping planes define a truncated pyramid - the view frustum.