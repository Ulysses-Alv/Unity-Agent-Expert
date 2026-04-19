# Trail Renderer component reference

[Switch to Scripting](../ScriptReference/TrailRenderer.md "Go to TrailRenderer page in the Scripting Reference")

This section contains the following sub-sections:

* [Trail settings](#trail-settings)
* [Materials](#materials)
* [Lighting](#lighting)
* [Probes](#probes)
* [Additional Settings](#additional-settings)

## Trail settings

| **Property** | **Function** |
| --- | --- |
| **Width** | Define a width value, and a curve value to control the width of your trail along its length.  The curve is sampled at each vertex, so its accuracy is limited by the number of vertices in your trail. The overall width of the trail is controlled by the width value. |
| **Time** | Define the lifetime of a point in the trail, in seconds. |
| **Min Vertex Distance** | The minimum distance between points in the trail, in world units. |
| **AutoDestruct** | Enable this to destroy the **GameObject** to which the **Trail Renderer** component is attached after it has not moved for **Time** seconds. |
| **Emitting** | When this is enabled, Unity adds new points in the trail. When this is disabled, Unity does not add new points to the trail. Use this to pause and unpause trail generation. |
| **Color** | Define a gradient to control the color of the trail along its length.  Unity samples colors from the Color gradient at each vertex. Between each vertex, Unity applies linear interpolation to colors. Adding more vertices to your trail might give a closer approximation of a detailed gradient. |
| **Corner Vertices** | This property dictates how many extra vertices are used when drawing corners in a trail. Increase this value to make the trail corners appear rounder. |
| **End Cap Vertices** | This property dictates how many extra vertices are used to create end caps on the trail. Increase this value to make the trail caps appear rounder. |
| **Alignment** | Set the direction that the trail faces. The options are:  * **View**: The trail faces the Camera. * **TransformZ**: The trail faces the Z axis of its Transform component. |
| **Texture Mode** | Control how the Texture is applied to the trail. The options are:  * **Stretch**: Map the texture once along the entire length of the trail. * **Tile**: Repeat the texture along the trail, based on its length in world units. To set the tiling rate, use [Material.SetTextureScale](../ScriptReference/Material.SetTextureScale.md). * **DistributePerSegment**: Map the texture once along the entire length of the trail, assuming all vertices are evenly spaced. * **RepeatPerSegment**: Repeat the texture along the trail, repeating at a rate of once per trail segment. To adjust the tiling rate, use [Material.SetTextureScale](../ScriptReference/Material.SetTextureScale.md). |
| **Shadow Bias** | Set the amount to move shadows away from the Light to remove shadowing artifacts. |
| **Generate Lighting Data** | If enabled, Unity builds the trail geometry with normals and tangents included. This allows it to use Materials that use the **Scene** lighting. |

## Materials

The **Materials** section lists all the [materials](class-Material.md) that this component uses.

| **Property** | **Description** |
| --- | --- |
| **Size** | The number of elements in the material list.  If you decrease the number of elements, Unity deletes the elements at the end of the list. If you increase the number of elements, Unity adds new elements to the end of the list. Unity populates new elements with the same material that the element at the end of the list uses. |
| **Element** | The materials in the list. You can assign a material asset to each element.  By default, Unity orders the list alphabetically based on the name of the materials. This list is reorderable, and Unity updates the number of the elements automatically as you change their order. |

## Lighting

The **Lighting** section contains properties that relate to lighting.

| **Property** | **Description** |
| --- | --- |
| **Cast Shadows** | Specify if and how this Renderer casts shadows when a suitable [Light](class-Light.md) shines on it.  This property corresponds to the [Renderer.shadowCastingMode](../ScriptReference/Renderer-shadowCastingMode.md) API.  The options are:  * **On**: This Renderer casts a shadow when a shadow-casting Light shines on it. * **Off**: This Renderer does not cast shadows. * **Two-sided**: This Renderer casts two-sided shadows. This means that single-sided objects like a plane or a quad can cast shadows, even if the light source is behind the mesh. For [Baked Global Illumination](class-LightingSettings.md#MixedLighting) or Enlighten [Realtime Global Illumination](class-LightingSettings.md#RealtimeLighting) to support two-sided shadows, the material must support [Double Sided Global Illumination](https://docs.unity3d.com/Manual/class-Material.html). * **Shadows Only**: This Renderer casts shadows, but the Renderer itself isn’t visible. |
| **Receive Shadows** | Specify if Unity displays shadows cast onto this Renderer.  This property only has an effect if you enable [Baked Global Illumination](class-LightingSettings.md#MixedLighting) or Enlighten [Realtime Global Illumination](class-LightingSettings.md#RealtimeLighting) for this scene.  This property corresponds to the [Renderer.receiveShadows](../ScriptReference/Renderer-receiveShadows.md) API.  If your project uses the Universal Render Pipeline (URP), this property appears in the **Inspector** window of materials instead. For more information, refer to [Shader material Inspector window reference for URP](urp/shaders-in-universalrp-reference.md). |
| **Contribute Global Illumination** | Include this Renderer in global illumination calculations, which take place at bake time.  This property only has an effect if you enable [Baked Global Illumination](class-LightingSettings.md#MixedLighting) or Enlighten [Realtime Global Illumination](class-LightingSettings.md#RealtimeLighting) for this scene.  Enabling this property enables the Contribute GI flag in the GameObject’s [Static Editor Flags](StaticObjects.md). It corresponds to the [StaticEditorFlags.ContributeGI](../ScriptReference/StaticEditorFlags.ContributeGI.md) API. |
| **Receive Global Illumination** | Whether Unity provides global illumination data to this Renderer from baked lightmaps, or from runtime Light Probes.  This property is only editable if you enable **Contribute Global Illumination**. It only has an effect if you enable [Baked Global Illumination](class-LightingSettings.md#MixedLighting) or Enlighten [Realtime Global Illumination](class-LightingSettings.md#RealtimeLighting) for this scene.  This property corresponds to the [MeshRenderer.receiveGI](../ScriptReference/MeshRenderer-receiveGI.md) API.  The options are:  * **Lightmaps**: Unity provides global illumination data to this Renderer from lightmaps. * **Light Probes**: Unity provides global illumination data to this Renderer from [Light Probes](LightProbes.md) in the scene. |
| **Prioritize Illumination** | Enable this property to always include this Renderer in Enlighten Realtime Global Illumination calculations. This ensures that the Renderer is affected by distant emissives, even those which are normally excluded from Global Illumination calculations for performance reasons.  This property is visible only if **Contribute GI** is enabled in the GameObject’s [Static Editor Flags](StaticObjects.md), your project uses the Built-in Render Pipeline, and Enlighten [Realtime Global Illumination](class-LightingSettings.md#RealtimeLighting) is enabled in your scene. |
| ****Mesh** **LOD** for **Global Illumination**** | The Mesh LOD index to use for global illumination and when baking lights. A less detailed model lowers global illumination quality, but decreases the bake memory usage and increases the bake performance. This setting is only visible if Mesh LODs have been generated for the mesh and the renderer contributes to global illumination. |

## Probes

The **Probes** section contains properties relating to [Light Probe](LightProbes.md) and [Reflection Probes](class-ReflectionProbe.md).

| **Property** | **Description** |
| --- | --- |
| **Light Probes** | Set how this Renderer receives light from the [Light Probes](LightProbes.md) system.  This property corresponds to the [Renderer.lightProbeUsage](../ScriptReference/Renderer-lightProbeUsage.md) API.  The options are:  * **Off**: The Renderer doesn’t use any interpolated Light Probes. * **Blend Probes**: The Renderer uses one interpolated Light Probe. This is the default value. * **Use Proxy Volume**: The Renderer uses a 3D grid of interpolated Light Probes. * **Custom Provided**: The Renderer extracts Light Probe shader uniform values from the [MaterialPropertyBlock](../ScriptReference/MaterialPropertyBlock.md). |
| **Proxy Volume Override** | Set a reference to another GameObject that has a Light Probe Proxy Volume component.  This property is only visible when **Light Probes** is set to **Use Proxy Volume**. |
| **Reflection Probes** | Set how the Renderer receives reflections from the [Reflection Probe](class-ReflectionProbe.md) system. This property is available only if you use the Built-In Render Pipeline.  This property corresponds to the [Renderer.probeAnchor](../ScriptReference/Renderer-probeAnchor.md) API.  The options are:  * **Off**: Disables Reflection Probes. Unity uses a skybox for reflection. * **Blend Probes**: Enables Reflection Probes. Blending occurs only between Reflection Probes. This is useful in indoor environments where the character may transition between areas with different lighting settings. * **Blend Probes and Skybox**: Enables Reflection Probes. Blending occurs between Reflection Probes, or between Reflection Probes and the default reflection. This is useful for outdoor environments. * **Simple**: Enables Reflection Probes, but no blending occurs between Reflection Probes when there are two overlapping volumes. |
| **Anchor Override** | Set the Transform that Unity uses to determine the interpolation position when using the [Light Probe](LightProbes.md) or [Reflection Probe](class-ReflectionProbe.md) systems. By default, this is the centre of the bounding box of the Renderer’s geometry.  This property corresponds to the [Renderer.probeAnchor](../ScriptReference/Renderer-probeAnchor.md) API. |

## Additional Settings

The **Additional Settings** section contains additional properties.

| **Property** | **Description** |
| --- | --- |
| **Motion Vectors** | Set whether to use motion vectors to track this Renderer’s per-pixel, screen-space motion from one frame to the next. You can use this information to apply post-processing effects such as motion blur.  **Note:** not all platforms support motion vectors. See [SystemInfo.supportsMotionVectors](../ScriptReference/SystemInfo-supportsMotionVectors.md) for more information.  This property corresponds to the [Renderer.motionVectorGenerationMode](../ScriptReference/Renderer-motionVectorGenerationMode.md) API.  The options are:  * **Camera Motion Only**: Use only Camera movement to track motion. * **Per Object Motion**: Use a specific pass to track motion for this Renderer. * **Force No Motion**: Do not track motion. |
| **Dynamic Occlusion** | When **Dynamic Occlusion** is enabled, Unity’s [occlusion culling](OcclusionCulling.md) system culls this Renderer when it is blocked from a Camera’s view by a Static Occluder. Otherwise, the system does not cull this Renderer when it is blocked from a Camera’s view by a Static Occluder.  Dynamic Occlusion is enabled by default. Disable it for effects such as drawing the outline of a character behind a wall. |
| **Sorting Layer** | The name of this Renderer’s [Sorting Layer](class-TagManager.md). |
| **Order in Layer** | This Renderer’s order within a [Sorting Layer](class-TagManager.md). |

TrailRenderer