# Billboard Renderer component reference

[Switch to Scripting](../ScriptReference/BillboardRenderer.md "Go to BillboardRenderer page in the Scripting Reference")

Explore the properties and settings in the **Billboard** Renderer component to customize how Unity renders billboards and their interaction with lighting.

## Properties

Properties on this component are split into the following sections:

* [General](#General)
* [Lighting](#Lighting)
* [Probes](#Probes)
* [Additional Settings](#AdditionalSettings)

### General

This section contains general properties in the root of the component.

| **Property:** | **Function:** |
| --- | --- |
| **Billboard** | Specifies the [Billboard Asset](class-BillboardAsset.md) this component renders. |

### Lighting

The **Lighting** section contains properties that specify how this Billboard Renderer interacts with lighting in Unity.

| **Property:** | **Function:** |
| --- | --- |
| **Cast Shadows** | Specify if and how the Mesh casts shadows when a suitable [Light](class-Light.md) shines on it.   This property has the following options:  * **On**: The Mesh casts a shadow when a shadow-casting Light shines on it. * **Off**: The Mesh does not cast shadows. * **Two Sided**: The Mesh casts two-sided shadows from either side. Enlighten Realtime Global Illumination and the Progressive Lightmapper do not support two-sided shadows. * **Shadows Only**: Shadows from the Mesh are visible, but not the Mesh itself. |
| **Receive Shadows** | Enable this option to make the **Mesh** display any shadows that are cast upon it. This is only supported when using the [Progressive Lightmapper](progressive-lightmapper.md). |

### Probes

The **Probes** section contains properties relating to Light Probes and Reflection Probes.

| **Property** | **Function** |
| --- | --- |
| **Light Probes** | Set how this Renderer receives light from the [Light Probe](LightProbes.md) system. For more information, see [Light Probes](LightProbes.md).  This property has the following options:  * **Off**: The Renderer doesn’t use any interpolated Light Probes. * **Blend Probes**: The Renderer uses one interpolated Light Probe. This is the default value. * **Use Proxy Volume**: The Renderer uses a 3D grid of interpolated Light Probes. * **Custom Provided**: The Renderer extracts Light Probe shader uniform values from the [MaterialPropertyBlock](../ScriptReference/MaterialPropertyBlock.md). |
| **Proxy Volume Override** | Set a reference to another GameObject that has a Light Probe Proxy Volume component.  This property is only visible when **Light Probes** is set to **Use Proxy Volume**. |
| **Reflection Probes** | Set how the Renderer receives reflections from the [Reflection Probe](class-ReflectionProbe.md) system.  This property has the following options:  * **Off**: Disables Reflection Probes. Unity uses a skybox for reflection. * **Blend Probes**: Enables Reflection Probes. Blending occurs only between Reflection Probes. This is useful in indoor environments where the character may transition between areas with different lighting settings. * **Blend Probes and Skybox**: Enables Reflection Probes. Blending occurs between Reflection Probes, or between Reflection Probes and the default reflection. This is useful for outdoor environments. * **Simple**: Enables Reflection Probes, but no blending occurs between Reflection Probes when there are two overlapping volumes. |

### Additional Settings

This section contains additional rendering properties.

| **Property** | **Function** |
| --- | --- |
| **Motion Vectors** | Set whether to use motion vectors to track this Renderer’s per-pixel, screen-space motion from one frame to the next. You can use this information to apply post-processing effects such as motion blur. **Note**: Not all platforms support motion vectors. Refer to [SystemInfo.supportsMotionVectors](../ScriptReference/SystemInfo-supportsMotionVectors.md) for more information.   This property has the following options:  * **Camera Motion Only**: Use only Camera movement to track motion. * **Per Object Motion**: Use a specific pass to track motion for this Renderer. * **Force No Motion**: Do not track motion. |
| **Dynamic Occlusion** | When **Dynamic Occlusion** is enabled, Unity culls this Renderer when it is blocked from a Camera’s view by a Static Occluder. Dynamic Occlusion is enabled by default.  When **Dynamic Occlusion** is disabled, Unity does not cull this Renderer when it is blocked from a Camera’s view by a Static Occluder. Disable Dynamic Occlusion to achieve effects such as drawing the outline of a character behind a wall.  See documentation on [occlusion culling](OcclusionCulling.md) for more information. |

BillboardRenderer