# Terrain Settings reference

Each **Terrain** tile has its own settings. If you’re using the built-in Terrain system, you need to change settings for each individual tile. If you’re using the [Terrain Tools package](https://docs.unity3d.com/Packages/com.unity.terrain-tools@5.1/manual/terrain-toolbox.html), you can change all tiles at once.

**Tip:** When you create a new tile from an existing one, the new tile inherits most of the settings from its parent tile.

To change the settings for a single terrain tile, in both the built-in Terrain system and the Terran Tools package:

1. In the **Hierarchy** window, select the tile.
2. In the ****Inspector**** window, select **Terrain Settings** (gear icon).

## Basic Settings

| **Property** | **Description** |
| --- | --- |
| **Grouping ID** | The ID the tile uses for **Auto connect**. Only tiles with the same grouping ID can connect. |
| **Auto Connect** | Automatically connect to neighboring tiles that have the same **Grouping ID**. |
| **Reconnect** | If you change the **Grouping ID** of a tile, or disable **Auto Connect** for one or more tiles, you might lose connections between tiles. To reconnect tiles, click the **Reconnect** button. **Reconnect** only connects two adjacent tiles if they have the same **Grouping ID** and if both tiles have **Auto Connect** enabled. This option isn’t available in the **Terrain Toolbox** window, which handles multiple tiles at once. |
| **Draw** | Enable rendering the terrain. Drawing is enabled by default; disable it when you want to focus on other tiles while painting or in Play mode. The only element that isn’t drawn is the terrain’s height; trees, grass, and other details are still drawn. Note that disabling drawing doesn’t prevent actions on the tile. For example, height sculpting over the tile’s area still changes its height. |
| **Draw Instanced** | Enable instanced rendering. For more information, refer to [Optimizing draw calls](optimizing-draw-calls.md). |
| **Enable **Ray Tracing** Support** | Use ray tracing to draw the terrain. Ray tracing renders a more realistic terrain, but has an impact on performance. This option isn’t available in the **Terrain Toolbox** window, which handles multiple tiles at once. |
| ****Pixel** Error** | Simplifies the generated terrain to optimize rendering. The lower the value, the more faithful the rendering to the original maps, such as the **heightmaps** and textures, and the higher the performance needs. |
| **Minimum Detail Limit** | Higher levels increase the heightmap’s complexity. `0` means no limit on how simple the heightmap can be. **Minimum Detail Limit** and **Maximum Complexity Limit** have a combined value that is limited by **Heightmap Resolution** (refer to Texture Resolution below). Because of the limit of their combined value, they limit each other. This option isn’t available in the **Terrain Toolbox** window, which handles multiple tiles at once. |
| **Maximum Complexity Limit** | Higher levels simplify the heightmap. `0` means no limit on how complex the heightmap can be. **Maximum Complexity Limit** and **Minimum Detail Limit** have a combined value that is limited by **Heightmap Resolution** (refer to Texture Resolution below). Because of the limit of their combined value, they limit each other. This option isn’t available in the **Terrain Toolbox** window, which handles multiple tiles at once. |
| **Base Map Dist.** | The maximum distance to **render textures** in full resolution. Beyond this distance, a lower resolution composite image is used. |

## Cast shadows

Defines how the terrain casts shadows onto other objects. Shadow interactions with other objects rely on [Rendering.ShadowCastingMode](../ScriptReference/Rendering.ShadowCastingMode.md).

| **Property** | **Description** |
| --- | --- |
| **Off** | The terrain doesn’t cast shadows. |
| **On** | The terrain casts one sided shadows. |
| **Two Sided** | The terrain casts a shadow from either side of a surface, ignoring backface culling. |
| **Shadows Only** | Render only the terrain’s shadows, not the terrain itself. |

## Reflection probes

To control how terrain reflects its environment, set the blend options for adjacent **reflection probes** and the **skybox**. Reflection probes create more realistic reflections, but decrease performance. Global environment reflections are controlled from the [Lighting window](lighting-window.md). You can override these reflections for a specific Terrain tile with the **Reflection Probes** setting. This setting only has an effect when you use the default material or a custom material with a **shader** that supports rendering with reflections. For more information, refer to [Reflection Surface Shader examples in the Built-In Render Pipeline](SL-SurfaceShaderExamples-Reflection.md).

Note that HDRP doesn’t support skyboxes for environment reflections.

| **Property** | **Description** |
| --- | --- |
| **Off** | Disable reflection probe blending and use only the skybox for reflection. |
| **Blend Probes** | Blends only adjacent probes and ignores the skybox. Use this for indoors or overhung areas, such as caves, to prevent sky reflections. |
| **Blend Probes And Skybox** | Works like **Blend Probes**, but also blends the skybox. Use this for open air areas, where the sky is always visible. |
| **Simple** | Use reflection probes, but don’t blend adjacent ones. |

## Material

Replace the terrain’s default material. You can select a material you’ve created earlier, or duplicate the default material. Changes done to any material, including the default material, impact all Terrain tiles using that material.

| **Property** | **Description** |
| --- | --- |
| **Create** | Creates a new copy of the default material and applies it to the current tile. The new material opens in the Inspector window. The **Create** button is only available once for each tile; it’s removed from the interface after use. |
| **Select Material** | Use a material you’ve created earlier. In the **Material** field, select the texture picker (circle). |

## Tree and Detail Objects

| **Property** | **Description** |
| --- | --- |
| **Draw** | Enable rendering tress, grass, and details. Drawing is enabled by default; disable it when you want to focus on the topography. |
| **Bake Light Probes For Trees** | Create internal [light probes](LightProbes.md) at the position of each tree, and use their information to render light. These probes don’t affect other renderers in the scene.  If you don’t check this box, trees are still affected by [Light Probe Groups](class-LightProbeGroup.md). This option is only effective for trees whose prototype has the **Mesh Renderer** > **Receive Global Illumination** > **Light Probe** property enabled (refer to [Mesh Renderer component](class-MeshRenderer.md)). |
| **Remove Light Probe Ringing** | Ringing is an unwanted light probe behaviour that can happen when there are different light intensities in the **scene**. It causes light to render where it shouldn’t, such as the dark side of an object. **Remove Light Probe Ringing** removes the ringing, but it also makes the probe less exact, and reduces light contrast. For more information, refer to [Light Probe Groups: Ringing](light-probes-troubleshooting.md#Ringing). This option is only available when you use **Bake Light Probes For Trees**. |
| **Preserve Tree Prototype Layers** | For [ray casting](use-layers.md), let tree instances use the layer value of their prototype, instead of the Terrain **GameObject**’s layer value. |
| **Detail Distance** | How far, in meters, an object ([Detail Mesh or Grass Texture](terrain-Grass.md)) needs to be from the **camera** to be culled. If the distance is too short, a user moving through the scene might see the change from culling to visible as objects popping in or out of the frame. |
| **Detail Density Scale** | A factor applied to the density of all details with the **Affected by Density Scaling** property enabled. `0` means no details; `1` means the full density defined for each detail. For more information, refer to [Grass and other details](terrain-Grass.md). This option isn’t available in the **Terrain Toolbox** window, which handles multiple tiles at once. |
| **Tree Distance** | How far, in meters, a tree needs to be from the camera to be culled. If the distance is too short, a user moving through the scene might see the change from culling to visible as objects popping in or out of the frame. For SpeedTree trees, this property has no impact; use the tree’s **LOD** group. |
| ****Billboard** Start** | How far, in meters, a tree needs to be from the camera to change from a **3D object** to a billboard. If the distance is too short, a user moving through the scene might see the trees change shape. For SpeedTree trees, this property has no impact; use the tree’s **LOD group**. |
| **Fade Length** | The distance over which trees transition between 3D objects and billboards. A longer distance creates a slower transition, and a user moving through the scene might see the branches move from facing the camera (billboard) to their final 3D positions. For SpeedTree trees, this property has no impact; use the tree’s LOD group. |
| **Max **Mesh** Trees** | The maximum number of visible trees that can render as 3D meshes. Beyond this limit, trees are all billboards, even if they’re close enough to the camera to be 3D. The Editor prioritises the trees nearest the camera to be 3D. For SpeedTree trees, this property has no impact; use the tree’s LOD group. |

## Tree Motion Vectors

For the animation of trees moving in the wind: Set whether to use motion vectors. Motion vectors track each renderer’s per-pixel, screen-space motion from one frame to the next. You can use this information to apply post-prcessing effects.   
  
**Note:** Not all platforms support motion vectors. Refer to [SystemInfo.supportsMotionVectors](../ScriptReference/SystemInfo-supportsMotionVectors.md) for more information.  
  
This property corresponds to the [Terrain.treeMotionVectorModeOverride](../ScriptReference/Terrain-treeMotionVectorModeOverride.md) method. This option isn’t available in the **Terrain Toolbox** window, which handles multiple tiles at once.

| **Property** | **Description** |
| --- | --- |
| **Camera Motion Only** | Use only camera movement to track motion. |
| **Per Object Motion** | Track motion when the camera moves, the GameObject moves, or both move. |
| **Force No Motion** | Don’t track motion. |
| **Inherit From Prototype** | Use the **Motion Vector** property setting of the tree prototype. |

## Detail Scatter Mode

How to cover an area with details (mesh or grass). This option isn’t available in the **Terrain Toolbox** window, which handles multiple tiles at once.

| **Property** | **Description** |
| --- | --- |
| **Coverage** | Paints based on the detail’s **Detail Density** value (refer to [Grass and other details](terrain-Grass.md)). |
| **Instance Count** | Paints based on the **Detail Resolution** and **Detail Resolution Per Patch** values in the **Mesh Resolution** properties below. |

## Wind Settings for Grass

| **Property** | **Description** |
| --- | --- |
| **Speed** | How quickly the wind moves the grass from side to side. |
| **Size** | The size of ripples on grassy areas. Lower values create a uniform movement; higher values create waves of motion in different directions. |
| **Bending** | How much the grass bends in the wind. `0` stops all movement. `1` moves the grass halfway towards the ground, but doesn’t lay it flat. |
| **Grass Tint** | The color tint applied to all grass objects. The final color for each grass is the **Grass Tint** multiplied by the grass’s **Healthy Color** and **Dry Color** values, which you can set for each grass type individually (refer to [Grass and other details](terrain-Grass.md)). |

**Wind Settings for Grass** don’t impact trees. To animate tree movement, refer to [Animate trees with wind](class-WindZone.md).

## Mesh Resolution

With all resolutions, higher values provide more detail but require more memory and processing power.

| **Property** | **Description** |
| --- | --- |
| **Terrain Width** | The terrain tile’s x-axis value, in world units. The value range is from 1 to 100,000. Changing this value changes the size of anything already on the tile, such as sculpted hills. |
| **Terrain Length** | The terrain tile’s z-axis value, in world units. The value range is from 1 to 100,000. Changing this value changes the size of anything already on the tile, such as sculpted hills. |
| **Terrain Height** | The possible difference between the lowest and highest points in the heightmap (y-axis), in world units. The value range is from 1 to 10,000. You can’t use negative values; if you need to sculpt below the surface level of your world, raise the terrain’s starting point (refer to [Sculpt the terrain with a brush](terrain-RaiseLowerTerrain.md)). Changing this value changes the size of anything already on the tile, such as sculpted hills. |
| **Detail Resolution Per Patch** | The size in pixels of each detail patch, which is a part of the **Detail Resolution** map. A high value reduces draw calls, but might increase triangle count because detail patches are culled per batch. The recommended value is 16. If your **Detail Distance** is high and your grass is sparse, use a higher value. The range is from `8` to `128`, and the value is squared to form a grid. |
| **Detail Resolution** | The number of pixels in the detail resolution map. A higher value gives more exact detail painting. The range is `0` to `4048`, and the value is squared to form a grid. |

## Holes Settings

**Compress Holes Texture**: Compress the texture used to paint holes to the `DXT1` format in the Player during runtime. This format is widely supported on PC and console platforms. For more information, refer to [TextureFormat.DXT1](https://docs.unity3d.com/ScriptReference/TextureFormat.DXT1.html) and [Paint holes in the terrain](terrain-PaintHoles.md).

This option isn’t available in the **Terrain Toolbox** window.

## Texture Resolutions

With all resolutions, higher values provide more detail but require more memory and processing power.

| **Property** | **Description** |
| --- | --- |
| **Heightmap Resolution** | The pixel resolution of the terrain tile’s heightmap. The higher the resolution, the more detail the terrain contours can include. This value must be a power of two plus one, for example, 513, which is 512 + 1. Changing this value changes the possible values of **Minimum Detail Limit** and **Maximum Complexity Limit**: their combined value is the power of the resolution minus 4. If **Heightmap Resolution** is `33 x 33`, **Minimum Detail Limit** and **Maximum Complexity Limit** is locked at `0`. |
| **Control Texture Resolution** | The resolution of the splatmap that controls the blending of the different Terrain Textures. |
| **Base Texture Resolution** | The resolution of the composite texture the terrain uses when you view it from a distance greater than the [Basemap Distance](https://docs.unity3d.com/ScriptReference/Terrain-basemapDistance.html) (you can set a basemap distance for all terrain tiles in **Project Settings** > **Quality** > **Terrain Setting Overrides**). |

The **Require resampling on change** banner reminds you that when you change properties under **Texture Resolutions**, the Unity Editor resizes the tile’s content to the new size you specify. This can affect the quality of your content.

### Heightmap Import and Export

You can import heightmaps to use in your project, or export your existing heightmap. This option isn’t available in the **Terrain Toolbox** window, because you can’t export multiple tiles at once.

The heightmap is an image file in the RAW grayscale format.

For more information, refer to [Work with heightmaps](terrain-Heightmaps.md).

## Lighting

These options aren’t available in the **Terrain Toolbox** window, which handles multiple tiles at once.

| **Property** | **Description** |
| --- | --- |
| **Contribute **Global Illumination**** | Use the Terrain in Global Illumination computations. When you enable this property, **Lightmapping** properties become available. |
| **Receive Global Illumination** | Use **lightmaps** or light probes as illumination sources. If you disable **Contribute Global Illumination**, select **Light Probes** in this list. If you enable **Contribute Global Illumination**, you can select **Lightmaps**. The Lightmapping options are available only if you enable **Contribute Global Illumination** in the **Lighting** section, and select **Lightmaps** from the **Receive Global Illumination** dropdown. |

## Lightmapping

These options aren’t available in the **Terrain Toolbox** window, which handles multiple tiles at once.

These options are available only if you enable **Contribute Global Illumination** in the **Lighting** section, and select **Lightmaps** from the **Receive Global Illumination** dropdown.

| **Property** | **Description** |
| --- | --- |
| **Scale in Lightmap** | Specify the relative size of an object’s [UV tile within a lightmap](LightingGiUvs.md). A value greater than `1.0` increases the number of pixels (the lightmap resolution) used, and creates more exact lighting for important or detailed areas. A value smalle than `1.0` decreases the number of pixels and is better for areas with fewer details. Set the value to `0` to let the object contribute lighting to other objects in the scene without being lightmapped itself. |
| **Lightmap Parameters** | Choose or create a Lightmap Parameters asset for the tile. Refer to [Lightmap Parameters](class-LightmapParameters.md) for more information. |
| **Rendering **Layer Mask**** | For the scriptable **render pipelines**, set the rendering layer for this Terrain tile. For more information, refer to [Rendering Layers in URP](urp/features/rendering-layers.md) or [Renderings Layer in HDRP](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@16.0/manual/Rendering-Layers). |

## Quality Settings

**Ignore Quality Settings**: Set the tile to ignore the Terrain Override Settings set in the [Quality settings](class-QualitySettings.md#Terrain).

This option isn’t available in the **Terrain Toolbox** window, which handles multiple tiles at once.

## Terrain Collider

The Terrain **Collider** component manages **collisions** on the terrain. For details about the **Terrain Collider** component properties, refer to [Terrain Collider component reference](class-TerrainCollider.md).

You can’t set a collider in the **Terrain Toolbox** window. You have to set them for each tile.

## Additional resources

* [Terrain Collider component reference](class-TerrainCollider.md)
* [Work with heightmaps](terrain-Heightmaps.md)
* [Mesh Renderer component](class-MeshRenderer.md)