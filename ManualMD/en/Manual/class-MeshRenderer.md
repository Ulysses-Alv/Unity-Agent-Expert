# Mesh Renderer component reference

[Switch to Scripting](../ScriptReference/MeshRenderer.md "Go to MeshRenderer page in the Scripting Reference")

A **Mesh** Renderer component renders a mesh. It works with a [Mesh Filter](class-MeshFilter.md) component on the same **GameObject**; the **Mesh Renderer** renders the mesh that the Mesh Filter references.

To render a deformable mesh, use a [Skinned Mesh Renderer](https://docs.unity3d.com/Manual/class-SkinnedMeshRenderer.html) instead.

In C# code, the [MeshRenderer](../ScriptReference/MeshRenderer.md) class represents a Mesh Renderer component. The `MeshRenderer` class inherits much of its functionality from the [Renderer](../ScriptReference/Renderer.md) class. As such, this component has a lot in common with other components that inherit from `Renderer`, such as [Line Renderer](class-LineRenderer.md), and [Trail Renderer](class-TrailRenderer.md).

## Mesh Renderer Inspector reference

* [Materials](#materials)
* [Lighting](#lighting)
* [Lightmapping](#lightmapping)
* [Probes](#probes)
* [Ray Tracing](#ray-tracing)
* [Mesh LOD](#mesh-lod)
* [Additional Settings](#additional-settings)

## Materials

The **Materials** section lists all the [materials](class-Material.md) that this component uses.

| **Property** | **Description** |
| --- | --- |
| **Size** | The number of elements in the material list.  If you decrease the number of elements, Unity deletes the elements at the end of the list. If you increase the number of elements, Unity adds new elements to the end of the list. Unity populates new elements with the same material that the element at the end of the list uses. |
| **Element** | The materials in the list. You can assign a material asset to each element.  By default, Unity orders the list alphabetically based on the name of the materials. This list is reorderable, and Unity updates the number of the elements automatically as you change their order. |

**Note:** If there are more materials than there are sub-meshes, Unity renders the last sub-mesh with each of the remaining materials, one on top of the next. If the materials are not fully opaque, you can layer different materials and create interesting visual effects. However, fully opaque materials overwrite previous layers, so any additional opaque materials that Unity applies to the last sub-mesh negatively affect performance and produce no benefit.

## Lighting

The **Lighting** section contains properties that relate to lighting.

| **Property** | **Description** |
| --- | --- |
| **Cast Shadows** | Specify if and how this Renderer casts shadows when a suitable [Light](class-Light.md) shines on it.  This property corresponds to the [Renderer.shadowCastingMode](../ScriptReference/Renderer-shadowCastingMode.md) API.  The options are:  * **On**: This Renderer casts a shadow when a shadow-casting Light shines on it. * **Off**: This Renderer does not cast shadows. * **Two-sided**: This Renderer casts two-sided shadows. This means that single-sided objects like a plane or a quad can cast shadows, even if the light source is behind the mesh. For [Baked Global Illumination](class-LightingSettings.md#MixedLighting) or Enlighten [Realtime Global Illumination](class-LightingSettings.md#RealtimeLighting) to support two-sided shadows, the material must support [Double Sided Global Illumination](https://docs.unity3d.com/Manual/class-Material.html). * **Shadows Only**: This Renderer casts shadows, but the Renderer itself isn’t visible. |
| **Receive Shadows** | Specify if Unity displays shadows cast onto this Renderer.  This property only has an effect if you enable [Baked Global Illumination](class-LightingSettings.md#MixedLighting) or Enlighten [Realtime Global Illumination](class-LightingSettings.md#RealtimeLighting) for this scene.  This property corresponds to the [Renderer.receiveShadows](../ScriptReference/Renderer-receiveShadows.md) API.  If your project uses the Universal Render Pipeline (URP), this property appears in the **Inspector** window of materials instead. For more information, refer to [Shader material Inspector window reference for URP](urp/shaders-in-universalrp-reference.md). |
| **Contribute Global Illumination** | Include this Renderer in global illumination calculations, which take place at bake time.  This property only has an effect if you enable [Baked Global Illumination](class-LightingSettings.md#MixedLighting) or Enlighten [Realtime Global Illumination](class-LightingSettings.md#RealtimeLighting) for this scene.  Enabling this property enables the Contribute GI flag in the GameObject’s [Static Editor Flags](StaticObjects.md). It corresponds to the [StaticEditorFlags.ContributeGI](../ScriptReference/StaticEditorFlags.ContributeGI.md) API. |
| **Receive Global Illumination** | Whether Unity provides global illumination data to this Renderer from baked lightmaps, or from runtime Light Probes.  This property is only editable if you enable **Contribute Global Illumination**. It only has an effect if you enable [Baked Global Illumination](class-LightingSettings.md#MixedLighting) or Enlighten [Realtime Global Illumination](class-LightingSettings.md#RealtimeLighting) for this scene.  This property corresponds to the [MeshRenderer.receiveGI](../ScriptReference/MeshRenderer-receiveGI.md) API.  The options are:  * **Lightmaps**: Unity provides global illumination data to this Renderer from lightmaps. * **Light Probes**: Unity provides global illumination data to this Renderer from [Light Probes](LightProbes.md) in the scene. |
| **Prioritize Illumination** | Enable this property to always include this Renderer in Enlighten Realtime Global Illumination calculations. This ensures that the Renderer is affected by distant emissives, even those which are normally excluded from Global Illumination calculations for performance reasons.  This property is visible only if **Contribute GI** is enabled in the GameObject’s [Static Editor Flags](StaticObjects.md), your project uses the Built-in Render Pipeline, and Enlighten [Realtime Global Illumination](class-LightingSettings.md#RealtimeLighting) is enabled in your scene. |
| **Mesh **LOD** for **Global Illumination**** | The Mesh LOD index to use for global illumination and when baking lights. A less detailed model lowers global illumination quality, but decreases the bake memory usage and increases the bake performance. This setting is only visible if Mesh LODs have been generated for the mesh and the renderer contributes to global illumination. |

## Lightmapping

The **Lightmapping** section contains properties relating to baked and real-time **lightmaps**.

This section is visible only if only if **Receive Global Illumination** is set to **Lightmaps**.

| **Property:** | **Function:** |
| --- | --- |
| **Scale in Lightmap** | Specify the relative size of the Renderer’s UVs within a lightmap.  A value of 0 results in the Renderer not being lightmapped, but it still contributes to lighting other Renderers in the Scene. A value greater than 1.0 increases the number of pixels (the lightmap resolution) used for this Renderer, while a value less than 1.0 decreases it.  You can use this property to optimize lightmaps so that important and detailed areas are more accurately lit. For example, an isolated building with flat, dark walls uses a low lightmap scale (less than 1.0) while a collection of colourful motorcycles displayed close together requires a high scale value.  This property corresponds to the [MeshRenderer.scaleInLightmap](../ScriptReference/MeshRenderer-scaleInLightmap.md) API. |
| **Stitch Seams** | Enable [Lightmap seam stitching](troubleshooting-lightmapping-artifacts.md#Lightmapping-SeamStitching) for this Renderer.  This property corresponds to the [MeshRenderer.stitchLightmapSeams](../ScriptReference/MeshRenderer-stitchLightmapSeams.md) API. |
| **Lightmap Parameters** | Choose or create a [Lightmap Parameters](class-LightmapParameters.md) asset for this Renderer. |
| **Optimize Realtime UVs** | Determine whether Unity [merges the real-time lightmap UVs](LightingGiUvs.md) for this Renderer.  This property is visible only if your Project uses the Built-in Render Pipeline, and Enlighten [Realtime Global Illumination](class-LightingSettings.md#RealtimeLighting) is enabled in your scene. |
| **Max Distance** | Specify the maximum worldspace distance for UV chart simplification in real-time lightmaps. If charts are within this distance, Unity simplifies them.  This property is visible only if your Project uses the Built-in Render Pipeline, and Enlighten [Realtime Global Illumination](class-LightingSettings.md#RealtimeLighting) is enabled in your scene. |
| **Max Angle** | Specify the maximum angle in degrees between faces sharing a UV edge in real-time lightmaps. If the angle between the faces is below this value, Unity simplifies the UV charts.  This property is visible only if your Project uses the Built-in Render Pipeline, and Enlighten [Realtime Global Illumination](class-LightingSettings.md#RealtimeLighting) is enabled in your scene. |
| **Ignore Normals** | Enable this option to prevent Unity from splitting the UV charts in real-time lightmaps.  This property is visible only if your Project uses the Built-in Render Pipeline, and Enlighten [Realtime Global Illumination](class-LightingSettings.md#RealtimeLighting) is enabled in your scene. |
| **Min Chart Size** | Specify the minimum texel size used for a UV chart in real-time lightmaps. If stitching is required, a value of 4 creates a chart of 4x4 texels to store lighting and directionality. If stitching is not required, a value of 2 reduces the texel density and provides better lighting build times and game performance.  This property is visible only if your Project uses the Built-in Render Pipeline, and Enlighten [Realtime Global Illumination](class-LightingSettings.md#RealtimeLighting) is enabled in your scene. |

When you’ve baked your lighting data (menu: **Window** > **Rendering** > **Lighting** > **Generate Lighting** ), this section also shows the baked lightmaps and real-time lightmaps in the current scene that this Renderer uses.

## Probes

The **Probes** section contains properties relating to [Light Probe](LightProbes.md) and [Reflection Probes](class-ReflectionProbe.md).

| **Property** | **Description** |
| --- | --- |
| **Light Probes** | Set how this Renderer receives light from the [Light Probes](LightProbes.md) system.  This property corresponds to the [Renderer.lightProbeUsage](../ScriptReference/Renderer-lightProbeUsage.md) API.  The options are:  * **Off**: The Renderer doesn’t use any interpolated Light Probes. * **Blend Probes**: The Renderer uses one interpolated Light Probe. This is the default value. * **Use Proxy Volume**: The Renderer uses a 3D grid of interpolated Light Probes. * **Custom Provided**: The Renderer extracts Light Probe shader uniform values from the [MaterialPropertyBlock](../ScriptReference/MaterialPropertyBlock.md). |
| **Proxy Volume Override** | Set a reference to another GameObject that has a Light Probe Proxy Volume component.  This property is only visible when **Light Probes** is set to **Use Proxy Volume**. |
| **Reflection Probes** | Set how the Renderer receives reflections from the [Reflection Probe](class-ReflectionProbe.md) system. This property is available only if you use the Built-In Render Pipeline.  This property corresponds to the [Renderer.probeAnchor](../ScriptReference/Renderer-probeAnchor.md) API.  The options are:  * **Off**: Disables Reflection Probes. Unity uses a skybox for reflection. * **Blend Probes**: Enables Reflection Probes. Blending occurs only between Reflection Probes. This is useful in indoor environments where the character may transition between areas with different lighting settings. * **Blend Probes and Skybox**: Enables Reflection Probes. Blending occurs between Reflection Probes, or between Reflection Probes and the default reflection. This is useful for outdoor environments. * **Simple**: Enables Reflection Probes, but no blending occurs between Reflection Probes when there are two overlapping volumes. |
| **Anchor Override** | Set the Transform that Unity uses to determine the interpolation position when using the [Light Probe](LightProbes.md) or [Reflection Probe](class-ReflectionProbe.md) systems. By default, this is the centre of the bounding box of the Renderer’s geometry.  This property corresponds to the [Renderer.probeAnchor](../ScriptReference/Renderer-probeAnchor.md) API. |

### Ray Tracing

The Ray Tracing section contains properties that relate to ray tracing effects.

This section only appears if your graphics card supports ray tracing and you [upgrade your project to use DirectX 12 in the Player settings window](GraphicsAPIs.md).

These properties have no effect unless you configure your project in one of the following ways:

* You use the High Definition **Render Pipeline** (HDRP) and [enable ray tracing](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@14.0/manual/Ray-Tracing-Getting-Started.html).
* You write your own C# script to implement ray tracing effects with [RayTracingShader](../ScriptReference/Rendering.RayTracingShader.md) and [RayTracingAccelerationStructure](../ScriptReference/Rendering.RayTracingAccelerationStructure.md).

| **Property** | **Description** |
| --- | --- |
| **Ray Tracing Mode** | Specify if and how often Unity updates this Renderer during ray tracing. More frequent updates increase rendering time. The options are:  * **Off**: Unity excludes this Renderer from ray tracing calculations. This option causes the Renderer to not appear in ray-traced reflections, for example. * **Static**: Unity doesn’t update the [ray tracing acceleration structure](../ScriptReference/Rendering.RayTracingAccelerationStructure.md) if the GameObject’s **Transform** or mesh geometry changes. This means ray tracing effects don’t change at runtime. * **Dynamic Transform**: Unity updates the ray tracing acceleration structure if the GameObject’s **Transform** changes, but not if the mesh geometry changes. This is the default value for a Mesh Renderer. If your project uses HDRP, this setting is the same as **Dynamic Geometry** because HDRP rebuilds the acceleration structure every frame. * **Dynamic Geometry**: Unity updates the ray tracing acceleration structure every frame with the GameObject’s **Transform** and mesh geometry. This is the default value for a Skinned Mesh Renderer. |
| **Procedural Geometry** | Causes Unity to render this Renderer procedurally with an intersection **shader** rather than with the mesh in the **Mesh Filter** component. |
| **Acceleration Structure Build Flags** | Specifies whether this Renderer overrides the default build flags you specify when you create a [ray tracing acceleration structure](../ScriptReference/Rendering.RayTracingAccelerationStructure.md). |

### Mesh LOD

Settings for configuring the [Mesh LOD](lod/mesh-lod-introduction.md) behavior.

| **Property** | **Description** |
| --- | --- |
| **LOD Override** | Disables automatic LOD selection and sets the LOD index to the value in the **Override Level** property. |
| **Override Level** | Sets the LOD index to the value in this property when **LOD Override** is enabled. |
| **LOD Selection Bias** | Sets the value that Unity adds to the calculated LOD index. Increasing this value results in Unity selecting less detailed LODs, reducing the value - in more detailed LODs. Both the calculated index and the **LOD Selection Bias** value can be a decimal. Unity floors the sum of the values to calculate the LOD index. For example: if Unity calculated an index value of 4.3, and the bias value is 0.9, the resulting LOD index is 5. |

## Additional Settings

The **Additional Settings** section contains additional properties.

| **Property** | **Description** |
| --- | --- |
| **Motion Vectors** | Set whether to use motion vectors to track this Renderer’s per-pixel, screen-space motion from one frame to the next. You can use this information to apply post-processing effects such as motion blur.  **Note:** not all platforms support motion vectors. See [SystemInfo.supportsMotionVectors](../ScriptReference/SystemInfo-supportsMotionVectors.md) for more information.  This property corresponds to the [Renderer.motionVectorGenerationMode](../ScriptReference/Renderer-motionVectorGenerationMode.md) API.  The options are:  * **Camera Motion Only**: Use only Camera movement to track motion. * **Per Object Motion**: Use a specific pass to track motion for this Renderer. * **Force No Motion**: Do not track motion. |
| **Dynamic Occlusion** | When **Dynamic Occlusion** is enabled, Unity’s [occlusion culling](OcclusionCulling.md) system culls this Renderer when it is blocked from a Camera’s view by a Static Occluder. Otherwise, the system does not cull this Renderer when it is blocked from a Camera’s view by a Static Occluder.  Dynamic Occlusion is enabled by default. Disable it for effects such as drawing the outline of a character behind a wall. |

MeshRenderer