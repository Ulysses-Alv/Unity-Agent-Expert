# Renderer module reference

The Renderer module’s settings determine how a **particle**’s image or [Mesh](class-Mesh.md) is transformed, shaded, and overdrawn by other particles.

## Properties

For some properties in this section, you can use different modes to set their value. For information on the modes you can use, refer to [Vary Particle System properties over time](varying-particle-system-properties-over-time.md).

| **Property:** | **Function:** |
| --- | --- |
| **Render Mode** | How Unity produces the rendered image from the graphic image (or Mesh). For more information, refer to [Particle rendering and shading](particle-rendering-shading.md).  * **Billboard**: Unity renders the particles as billboards and they face the direction you specify in **Render Alignment**. * **Stretched Billboard**: The particles face the **Camera** with various possible scaling options applied. For more information, refer to [Stretched Billboard dropdown](#stretched-billboard). * **Horizontal Billboard**: The particle plane is parallel to the XZ “floor” plane. * **Vertical Billboard**: The particle is upright on the world Y-axis, but turns to face the Camera. * **Mesh**: Unity renders the particle from a 3D Mesh instead of a Billboard. For more information, refer to [Mesh dropdown](#mesh) and [Particle rendering and shading](particle-rendering-shading.md). * **None**: Unity doesn’t render any particles. This can be useful alongside the [Trails](PartSysTrailsModule.md) module, if you want to only render the trails and hide any default particle rendering. |
| **Normal Direction** | Specifies how to calculate lighting for the billboard. A value of 0 means Unity calculates lighting as though the billboard was a sphere. This results in the billboard looking more like a sphere. A value of 1 means Unity calculates lighting for the billboard as a flat quad.  This property is only available when using one of the Billboard rendering modes: **Billboard**, **Stretched Billboard**, **Horizontal Billboard** or **Vertical Billboard**. |
| **Material** | The material Unity uses to render the particles. |
| **Trail Material** | The material Unity uses to render particle trails.  This option is only available when the **Trails** module is enabled. |
| **Sort Mode** | The order in which Unity draws and overlays particles with a Particle System.  * **None**: When this setting is enabled, Unity doesn’t sort particles. * **By Distance**: Sorts particles in the system based on distance to the active Camera. Unity renders particles closer to the Camera on top of those that are further away. The order of particles doesn’t change when you rotate the camera with the setting. * **Oldest in Front**: Unity renders particles that have existed the longest at the front of the Particle System. * **Youngest in Front**: Unity renders particles that have existed for the shortest time at the front of the Particle System. * **By Depth**: Unity renders particles based on their distance from the camera’s near plane. The order of particles can change when you rotate the camera with this setting. |
| **Sorting Fudge** | The bias of the **Particle System** sort ordering. Lower values increase the relative chance that Unity draws Particle Systems over other transparent **GameObjects**, including other Particle Systems. This setting only affects Particle Systems as a whole that appear in the **scene** - it does not perform sorting on individual particles within a system. |
| **Min Particle Size** | The smallest particle size (regardless of other settings), expressed as a fraction of viewport size.  This property is only available when using one of the Billboard rendering modes: **Billboard**, **Stretched Billboard**, **Horizontal Billboard** or **Vertical Billboard**. |
| **Max Particle Size** | The largest particle size (regardless of other settings), expressed as a fraction of viewport size.   This property is only available when using one of the Billboard rendering modes: **Billboard**, **Stretched Billboard**, **Horizontal Billboard** or **Vertical Billboard**. |
| **Render Alignment** | This property determines the direction that particle billboards face.  * **View**: Particles face the Camera plane. * **World**: Particles align with the world axes. * **Local**: Particles align to the **Transform component** of their GameObject. * **Facing**: Particles face the direct position defined by the Transform component in the GameObject of the active Camera. * **Velocity**: Particles face in the same direction as their velocity vector. |
| **Enable Mesh GPU Instancing** | This property is only available when using the Mesh render mode.  This property controls whether Unity renders the Particle System using GPU Instancing. This requires the use of a compatible shader. For more information, refer to [Particle Mesh GPU Instancing](PartSysInstancing.md). |
| **Flip** | Mirror a proportion of the particles across the specified axes. A higher value flips more particles. |
| **Allow Roll** | Controls whether camera-facing particles can rotate around the Z-axis of the **camera**. Disabling this can be particularly useful for **VR** applications, where HMD (Head-Mounted Display) roll can cause undesirable particle rotation for Particle Systems. |
| **Pivot** | Modify the central pivot point for rotating particles. The value is a multiplier of the particle size. |
| **Visualize Pivot** | Preview the particle pivot points in the **Scene** View. |
| **Masking** | Set how the particles rendered by the Particle System behave when they interact with a **Sprite Mask**.  * **No Masking**: The Particle System does not interact with any Sprite Masks in the Scene. This is the default option. * **Visible Inside Mask**: The particles are visible where the Sprite Mask overlays them, but not outside of it. * **Visible Outside Mask**: The particles are visible outside of the Sprite Mask, but not inside it. The Sprite Mask hides the sections of the particles it overlays. |
| **Apply Active Color Space** | When rendering in Linear Color Space, the system converts particle colors from Gamma Space before it uploads them to the GPU. |
| **Custom Vertex Streams** | Configure which particle properties are available in the Vertex **Shader** of your Material. For more information, refer to [Particle System vertex streams and Standard Shader support](PartSysVertexStreams.md). |
| **Cast Shadows** | If this property is enabled, the Particle System creates shadows when a shadow-casting Light shines on it.  * **On**: Enables shadows for this Particle System. * **Off**: Disables shadows for this Particle System. * **Two-Sided**: Select **Two Sided** to allow shadows to be cast from either side of the Mesh. Backface culling is not taken into account when this property is enabled. * **Shadows Only**: Select **Shadows Only** to make it so that the shadows are visible, but the Mesh itself is not. |
| **Shadow Bias** | Move the shadows along the direction of the light. This removes shadowing artifacts caused by approximating volumes with **billboards**. |
| **Motion Vectors** | Set whether to use motion vectors to track the per-pixel, screen-space motion of this Particle System’s Transform component from one frame to the next.  * **Camera Motion Only**: Use only Camera movement to track motion. * **Per Object Motion**: Use a specific pass to track motion for this Renderer. * **Force No Motion**: Do not track motion.  **Note:** Not all platforms support motion vectors. For more information, refer to [SystemInfo.supportsMotionVectors](../ScriptReference/SystemInfo-supportsMotionVectors.md). |
| **Receive Shadows** | Dictates whether particles in this system can receive shadows from other sources. Only opaque materials can receive shadows. |
| **Sorting Layer ID** | The name of the Renderer’s sorting layer. |
| **Order in Layer** | This Renderer’s order within a sorting layer. |
| **Light Probes** | Probe-based lighting interpolation mode. |
| **Reflection Probes** | If enabled, and if reflection probes are present in the Scene, Unity assigns a reflection texture from the nearest reflection probe to this GameObject and sets the texture as a built-in Shader uniform variable. |
| **Anchor Override** | A Transform that determines the interpolation position when you use the [Light Probe](LightProbes.md) or [Reflection Probe](ReflectionProbes.md) systems. |

### Stretched Billboard dropdown

The Stretched Billboard dropdown contains the following settings.

| **Value** | **Description** |
| --- | --- |
| **Camera Scale** | Stretches particles according to Camera movement. Set this to zero to disable Camera movement stretching. |
| **Velocity Scale** | Stretches particles proportionally to their speed. Set this to zero to disable stretching based on speed. |
| **Length Scale** | Stretches particles proportionally to their current size along the direction of their velocity. Setting this to zero makes the particles disappear, having effectively zero length. |
| **Freeform Stretching** | Indicates whether particles should use freeform stretching. With this stretching behavior, particles don’t become thin when you view them head-on. |
| **Rotate With Stretch** | Indicates whether to rotate particles based on the direction they stretch in. This is added on top of other particle rotation. This property only has an effect when you enable **Freeform Stretching**. If you disable **Freeform Stretching**, particles always rotate based on the direction they stretch in, even if **Rotate With Stretch** is disabled. |

### Mesh dropdown

The Mesh dropdown contains the following settings.

| **Value** | **Description** |
| --- | --- |
| **Mesh Distribution** | Specifies the method that Unity uses to randomly assign meshes to particles.  * **Uniform Random**: Unity randomly assigns meshes to particles with an even weighting. The Particle System as a whole should contain a roughly equal number of each possible mesh at any given moment. * **Non-uniform Random**: Unity randomly assigns meshes to particles with user-defined weights for each mesh. When this setting is enabled, the Renderer module Inspector window displays a Meshes list and a Mesh Weightings field for each mesh in the list. You can use the Mesh Weightings field to control how often Unity assigns each mesh to a particle. |
| **Mesh Weightings** | Controls how likely Unity is to assign this mesh to a particle. The weights work relative to each other; Unity is twice as likely to assign a mesh with double the weight of another mesh, regardless of their absolute value. For more information, refer to [Particle rendering and shading](particle-rendering-shading.md).  This setting is only available when the **Mesh Distribution** property is set to **Non-uniform Random**. |

## Additional resources

* [Particle rendering and shading](particle-rendering-shading.md)