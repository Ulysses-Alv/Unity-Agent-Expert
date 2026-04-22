# Lightmap Parameters Asset Inspector window reference

[Switch to Scripting](../ScriptReference/LightmapParameters.md "Go to LightmapParameters page in the Scripting Reference")

## Properties

When you click on a **Lightmap** Parameters Asset in the **Project window**, the **Inspector** window displays the values defined in that Asset. The parameters and their descriptions are listed in the table below.

### Realtime GI

These parameters configure [Enlighten Realtime Global Illumination](realtime-gi-using-enlighten.md).

See [render pipeline feature comparison](render-pipelines-feature-comparison.md) for more information about support for **Enlighten** Realtime **Global Illumination** across **render pipelines**.

| **Property** | **Function** |
| --- | --- |
| **Resolution** | This value scales the **Realtime Resolution** value in the Scene tab of the Lighting Window (menu: **Window** > **Rendering** > **Lighting** > **Scene**) to give the final resolution of the lightmap in texels per unit of distance. |
| **Cluster Resolution** | The ratio of the cluster resolution (the resolution at which the light bounces are calculated internally) to the final lightmap resolution. See documentation on [GI Visualizations in the Scene view](GIVis.md) for more information. |
| **Irradiance Budget** | This value determines the precision of the incoming light data used to light each texel in the lightmap. Each texel’s lighting is obtained by sampling a “view” of the Scene from the texel’s position. Lower values of irradiance budget result in a more blurred sample. Higher values increase the sharpness of the sample. A higher irradiance budget improves the lighting, but this increases run-time memory usage and might increase CPU usage. |
| **Irradiance Quality** | Use the slider to define the number of rays that are cast and used to compute which clusters affect a given output lightmap texel. Higher values offer visual improvements in the lightmap, but increase precomputing time in the Unity Editor. The value does not affect runtime performance. |
| **Modelling Tolerance** | This value controls the minimum size of gaps in **Mesh** geometry that allows light to pass through. Make this value lower to allow light to pass through smaller gaps in your environment. |
| **Edge Stitching** | If enabled, this property indicates that UV charts in the lightmap should be joined together seamlessly, to avoid unwanted visual artifacts. |
| **Is Transparent** | If enabled, the object appears transparent during the Global Illumination calculations. Back-faces do not contribute to these calculations, and light travels through the surface. This is useful for invisible emissive surfaces. |
| **System Tag** | A group of objects whose lightmap Textures are combined in the same lightmap atlas is known as a “system”. The Unity Editor automatically defines additional systems and their accompanying atlases if all the objects can’t be fitted into a single atlas. However, it is sometimes useful to define separate systems yourself (for example, to ensure that objects inside different rooms are grouped into one system per room). Change the **System Tag** number to force new system and lightmap creation. The exact numeric sequence values of the tag are not significant. |

### Baked GI

These parameters configure [lightmapping](Lightmappers.md).

See [render pipeline feature comparison](render-pipelines-feature-comparison.md) for more information about support for lightmapping across render pipelines.

| **Property** | **Function** |
| --- | --- |
| **Anti-aliasing Samples** | Determines the number of sub-texel positions to use when sampling a lightmap texel. Set the value above 1 to use supersampling to improve lightmap quality and reduce artifacts related to aliasing. •A value of 1 disables super sampling. •A value of 4 provides 2x2 supersampling. This is the default value which removes most aliasing artifacts. •A value of 16 provides 4x4 supersampling. Use these values to remove artifacts, like jagged edges in direct lighting. This can happen when you use baked shadows.  **Note**: A higher anti-aliasing sample value uses more memory. This means that if you use a high sample number in a large Scene with large lightmap texture sizes, the lightmap bake might not complete. |
| **Backface Tolerance** | Specifies the percentage of front-facing geometry sampling ray hits a texel must have for Unity to consider it valid. This makes it possible for Unity to invalidate a texel if too many of the rays cast from it hit backfaces (e.g. if the texel is inside geometry). A value of 1.0 means that Unity considers a texel invalid when any of its rays hits a backface, for example. When a texel is invalid, Unity clones valid values from surrounding texels to prevent artifacts. Lower this value to solve lighting problems caused by backface samples. Use the Texel Validity [Scene View Draw Mode](GIVis.md) to adjust this value. |
| **Pushoff** | Pushes ray origins away from geometry along the normal based on the value you specify in modelling units. Unity applies this value to all baked lightmaps. It affects direct light, indirect light, and baked **ambient occlusion**. Adjust this setting to reduce self-occlusion and self-shadowing artifacts. |
| **Baked Tag** | Groups specific sets of objects in atlases. As with the System Tag, the exact numeric value is not significant. Unity never puts **GameObjects** with different **Baked Tag** values in the same atlas. However, there is no guarantee that objects with the same tag end up in the same atlas because there may be too many objects with that tag to fit into one atlas. See image A (below) for an example of this. You do not need to set this value when you use the multi-scene bake API because Unity groups automatically in that case. You can use **Baked Tag** to replicate some of the behavior of the **Lock Atlas** option. See [Baked Tags: Details](Lightmapping-baked-tags.md) for more information. |
| **Limit Lightmap Count** | Applies a limit to the number of lightmaps that Unity can use to pack together GameObjects with the same Baked Global Illumination settings. When you enable this limit, a related setting appears called **Max Lightmaps**. That setting determines the lightmap limit. The [Lightmapping Settings](class-LightingSettings.md) define the size of these lightmaps.  Unity considers GameObjects to have the same Baked Global Illumination settings if they have equal values for **Anti-aliasing Samples**, **Pushoff**, **Baked Tag**, and **Backface Tolerance**. This means that Unity might pack together GameObjects associated with different Lightmap Parameter Assets.   To pack GameObjects into a set number of lightmaps, Unity scales down UV layouts until all fit within the specified number of lightmaps. This process may reduce lightmap resolution. |

## Baked AO

These parameters configure Baked Ambient Occlusion.

| **Property** | **Function** |
| --- | --- |
| **Quality** | The number of rays cast when evaluating baked ambient occlusion (AO). A higher numbers of rays increases AO quality but also increases bake time. |
| **Anti-aliasing Samples** | The number of samples to take when doing anti-aliasing of AO. A higher number of samples increases the AO quality but also increases the bake time. |

## General GI

| **Property** | **Function** |
| --- | --- |
| **Backface Tolerance** | The percentage of rays shot from an output texel that must hit front faces for the lighting system to consider them usable. This allows Unity to invalidate a texel if too many of the rays cast from it hit back faces ( (for example, if the texel is inside some geometry). The lighting system clones valid values from the surrounding texels to avoid unintended artifacts.  If **Backface Tolerance** is 0.0, the lighting system rejects the texel only if it sees nothing but backfaces. If it is 1.0, the lighting system rejects the ray origin if it has even one ray that hits a backface. |

LightmapParameters