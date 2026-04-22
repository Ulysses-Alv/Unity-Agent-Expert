# Shadow mapping

Unity performs shadow mapping to render real-time shadows.

## How shadow mapping works

Shadow mapping is the process of creating shadow textures called shadow maps. Unity generates a shadow map from the perspective of a light in a similar way to how a **camera** generates a depth texture. If you consider a camera at the same location as the light, the areas of the **scene** that the camera can’t see are the same areas of the scene that rays from the light cannot reach. Those areas are in shadow.

Unity populates the shadow map with information about how far rays from the light travel before they hit a surface, and then samples the shadow map to calculate real-time shadows for **GameObjects** that the light hits.

## Shadow map resolution

The larger the shadow map resolution, the better Unity can capture small details in the shadow-casting geometry, and the more precise the shadows can be. Larger shadow map resolution also requires more memory bandwidth.

For information on configuring the shadow map resolution, refer to the following pages:

* [Configure shadow resolution in URP](urp/shadow-resolution-urp.md)
* [Configure shadow resolution in the Built-In Render Pipeline](shadow-resolution-birp.md)

## Shadow mapping and performance

To create a shadow map for a point light, Unity captures the scene in six directions. The performance impact is similar to rendering shadows from six spot lights, and increases the number of draw calls significantly.

On mobile platforms, this process uses a significant amount of the available resources per frame. Reduce the number of point lights within the camera view as much as possible.

Real-time shadows from spot lights are significantly faster to render than real-time shadows from point lights.

## Additional resources

* [Configure shadow resolution in URP](urp/shadow-resolution-urp.md)
* [Universal Render Pipeline Asset](urp/universalrp-asset.md)
* [Configure shadow resolution in the Built-In Render Pipeline](shadow-resolution-birp.md)