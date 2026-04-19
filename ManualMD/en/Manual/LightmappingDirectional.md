# Store light direction with Directional Mode

There are two **Directional Modes** for **lightmaps**: **Directional** and **Non-Directional**. Both modes are compatible with real-time lightmaps from Unity’s **Enlighten** Realtime **Global Illumination** system, and baked lightmaps from Unity’s Progressive **Lightmapper**. The default mode is Directional.

When you bake a Directional lightmap, Unity generates two lightmap textures. One texture stores information about the intensity and color of the lighting received across the target surface. This is identical to the Non-Directional lightmap. The other texture stores the dominant light direction, along with a factor describing how much of the total light received comes from that dominant direction.

![The barrels in this image have baked Non-directional lightmaps.](../uploads/Main/DirectionalLightmapping1.jpg)

The barrels in this image have baked Non-directional lightmaps.

![The barrels in this image have baked Directional lightmaps. ](../uploads/Main/DirectionalLightmapping2.jpg)

The barrels in this image have baked Directional lightmaps.

**Note:** Since lightmaps are used exclusively for diffuse lighting, this mode doesn’t preserve specular response. For more information about including baked specular effects, refer to [Reflection Probes](ReflectionProbes.md).

## Performance

Directional mode lightmaps consist of two textures, and **shaders** sample both textures during rendering. The additional texture increases video memory requirements. Generating the additional directionality texture also has an impact on baking performance.

Non-Directional mode lightmaps only include a single texture. As a result, they require less video memory and less storage than Directional maps, and are faster to decode in shaders. These optimizations reduce visual quality.

## Setting your lightmap mode

To set the mode in your [**Lighting Settings Asset**](https://docs.unity3d.com/2020.1/Documentation/Manual/class-LightingSettings.html), open the Lighting window (**Window** > **Rendering** > **Lighting**), click **Scene**, navigate to the [Lightmapping Settings](https://docs.unity3d.com/Manual/lighting-window.html#LightmappingSettings), and select **Directional Mode**.

You can set the lightmap mode for an instance of the Lighting Settings asset which can apply to one or more **Scenes**. You cannot set the lightmap mode for individual lightmaps.

## Using Directional Mode with additive loading

Unity can load Scenes additively. This means you can use [Multi-Scene editing](https://docs.unity3d.com/Manual/MultiSceneEditing.html), for example.
When you load Scenes additively, all of them must use the same Directional Mode. This includes Scenes that are not baked, such as UI elements or loading screens. Using the same [Lightmap Parameters asset](https://docs.unity3d.com/Manual/class-LightmapParameters.html) for all of the Scenes in your project can help you avoid settings conflicts.