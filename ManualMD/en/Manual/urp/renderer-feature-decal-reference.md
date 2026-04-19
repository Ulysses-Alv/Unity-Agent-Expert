# Decal Renderer Feature reference

This section describes the properties of the Decal Renderer Feature.

![Decal Renderer Feature, Inspector view.](../../uploads/urp/decal/decal-rf-inspector.png)  
*Decal Renderer Feature, Inspector view.*

### Technique

Select the rendering technique for the Renderer Feature.

This section describes the options in this property.

#### Automatic

Unity selects the rendering technique automatically based on the build platform. The [Accurate G-buffer normals](rendering/accurate-g-buffer-normals.md) option is also taken into account, as it prevents normal blending from working correctly without the D-Buffer technique.

#### DBuffer

Unity renders decals into the Decal buffer (DBuffer). Unity overlays the content of the DBuffer on top of the opaque objects during the opaque rendering.

Selecting this technique reveals the **Surface Data** property. The Surface Data property lets you specify which surface properties of decals Unity blends with the underlying meshes. The Surface Data property has the following options:

* **Albedo**: decals affect the base color and the emission color.
* **Albedo Normal**: decals affect the base color, the emission color, and the normals.
* **Albedo Normal MAOS**: decals affect the base color, the emission color, the normals, the metallic values, the smoothness values, and the **ambient occlusion** values.

**Limitations:**

* This technique does not support the OpenGL and OpenGL ES API.
* This technique requires the DepthNormal prepass, which makes the technique less efficient on GPUs that implement tile-based rendering.
* This technique does not work on **particles** and **terrain** details.

#### Screen Space

Unity renders decals after the opaque objects using normals that Unity reconstructs from the depth texture, or from the G-Buffer when using the Deferred **rendering path**. Unity renders decals as meshes on top of the opaque meshes. This technique supports only the normal blending. When using the Deferred rendering path with [Accurate G-buffer normals](rendering/accurate-g-buffer-normals.md), blending of normals is not supported, and will yield incorrect results.

Screen space decals are recommended for mobile platforms that use tile-based rendering, because URP doesn’t create a DepthNormal prepass unless you enable **Use Rendering Layers**.

Selecting this technique reveals the following properties.

| **Property** | **Description** |
| --- | --- |
| **Normal Blend** | The options in this property (Low, Medium, and High) determine the number of samples of the depth texture that Unity takes when reconstructing the normal vector from the depth texture. The higher the quality, the more accurate the reconstructed normals are, and the higher the performance impact is. |
| **Low** | Unity takes one depth sample when reconstructing normals. |
| **Medium** | Unity takes three depth samples when reconstructing normals. |
| **High** | Unity takes five depth samples when reconstructing normals. |

### Max Draw Distance

The maximum distance from the **Camera** at which Unity renders decals.

### Use Rendering Layers

Select this check box to enable the [Rendering Layers](features/rendering-layers.md) functionality.

If you enable **Use Rendering Layers**, URP creates a DepthNormal prepass. This makes decals less efficient on GPUs that implement tile-based rendering.