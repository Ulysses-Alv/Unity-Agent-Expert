# Troubleshooting lightmapping quality issues

If you use baked [Global Illumination (GI)](global-illumination-configure.md), troubleshoot baked lighting quality issues such as flat normal maps, visible texture coordinate (UV) seams, blurry **lightmaps**, and aliasing. These issues make surfaces appear flat, disrupt texture continuity, or cause shadows to lose detail.

## Flat normal maps

This issue occurs if you use non-directional lightmaps with baked lighting. Non-directional lightmaps don’t capture the directionality of incoming light, which affects the appearance of surface detail when a **GameObject** has a [normal map](StandardShaderMaterialParameterNormalMap.md).

For more information about directional lightmaps, refer to [Store light direction with Directional Mode](LightmappingDirectional.md).

**Note:** Non-directional lightmaps store the intensity of light at various points but don’t account for the direction from which the light is coming, which means they don’t work with normal maps. As a result, details that help give an illusion of relief might be missing.

### Symptoms

* Objects lose their depth effect, and the lighting appears flat after you bake if **Directional Mode** is disabled.
* Normal maps don’t provide a clear representation of relief.

![Non-directional lightmaps with a baked light - left. Directional lightmaps with a baked light - right.](../uploads/Main/flat-normal-maps.jpg)

Non-directional lightmaps with a baked light - left. Directional lightmaps with a baked light - right.

### Cause

When you bake in non-directional mode, Unity doesn’t generate a secondary texture to store the dominant light direction.

### Resolution

To troubleshoot flat normal maps, do the following.

* Enable [Directional Mode](LightmappingDirectional.md), so Unity stores light direction.
* In the ****Inspector**** window for each light, set **Light Mode** to **Mixed**, so Unity bakes indirect diffuse lighting. For more information, refer to [Light Modes](LightModes-introduction.md).

## Lightmap UV seams

Visible seams can appear in lightmaps due to GPU limitations when blending color values between separate UV charts, especially with high filtering or low lightmap resolution. Seams between objects are more pronounced in modular meshes.

### Symptoms

* Visible seams in lightmaps after baking.
* More noticeable seams when using modular meshes.
* Increased seam visibility if you use high filtering or low lightmap resolution.

![Left: The lightmap UV split is visible on the hand of the model after baking. Right: After you enable Stitch Seams and bake again, the lightmap UV split is no longer visible.](../uploads/Main/lightmap-uv-seams.jpg)

Left: The lightmap UV split is visible on the hand of the model after baking. Right: After you enable **Stitch Seams** and bake again, the lightmap UV split is no longer visible.

### Cause

By default, GPUs don’t blend color values between separate UV charts, leading to visible seams in baked lightmaps. Filtering and low lightmap resolution can further exacerbate the issue by bleeding texel color values into adjacent UV charts.

### Resolution

To troubleshoot lightmap UV seams, do the following.

#### Enable Stitch Seams

Enable [Stitch Seams](troubleshooting-lightmapping-artifacts.md#Lightmapping-SeamStitching) to blend color values between UV charts that share a stitchable common edge.

**Note:** **Stitch Seams** only works for single GameObjects and doesn’t support multiple GameObjects.

#### Combine meshes in a digital content creation tool

To join overlapping vertices, do the following:

* Use a 3D modeling tool to merge modular meshes. Join overlapping vertices before you export, to prevent automatic lightmap unwrapping issues in Unity.
* Join overlapping vertices in the [import Settings for a model](FBXImporter-Model.md).

![Cornell Box scene assembled using modular meshes with visible seams after lightmap baking - left. Meshes combined in a digital content creation (DCC) tool before exporting, with no lightmap seams visible - right.](../uploads/Main/combine-meshes-in-a-dcc-tool.jpg)

Cornell Box scene assembled using modular meshes with visible seams after lightmap baking - left. Meshes combined in a digital content creation (DCC) tool before exporting, with no lightmap seams visible - right.

#### Other possible solutions

* Use [Mesh.CombineMeshes](https://docs.unity3d.com/ScriptReference/Mesh.CombineMeshes.html) to merge multiple meshes into one, achieving a similar result as in a DCC tool.
* Convert meshes to ProBuilder using the [ProBuilderize](https://docs.unity3d.com/Packages/com.unity.probuilder@latest/?subfolder=/manual/Object_ProBuilderize.html) option. Use the [Merge Objects](https://docs.unity3d.com/Packages/com.unity.probuilder@latest/?subfolder=/manual/Object_Merge.html?q=merge) option to combine multiple objects.
* Disable filtering and denoising. This prevents blurring and dilation issues but can require increasing sample counts for cleaner results.
* Adjust **Lightmap Padding** to increase spacing between UV layouts in the UV atlas, reducing texel bleeding.
* For low-end platforms, use [Light Probe Groups](class-LightProbeGroup.md).
* For per-pixel lighting with streamlined placement, use [Adaptive Probe Volumes (APV)](urp/probevolumes-use.md).
* If possible, use trim meshes to hide seams between objects.

## Blurry lightmaps

Low lightmap resolution results in blurry lighting effects due to insufficient texel density.

### Symptoms

* Shadows appear soft or lack detail.
* Lightmaps are stretched or distorted.
* Light leaks through geometry.

### Cause

Lightmaps require a sufficient resolution to accurately capture lighting information. Too few texels per unit can result in a loss of detail, while non-uniform UV scaling can distort the baked lighting output.

### Resolution

![Cornell box scene baked at 8 texels per unit resolution - left. Same scene baked at 32 texels per unit resolution - right.](../uploads/Main/blurry-lightmaps.jpg)

Cornell box scene baked at 8 texels per unit resolution - left. Same scene baked at 32 texels per unit resolution - right.

To troubleshoot blurry lightmaps, do the following:

* To adjust lightmap resolution globally, adjust the **Lightmap Resolution** property in the [Lighting](lighting-window.md) window. This changes the resolution for all objects in the **scene**.
* To adjust lightmap resolution per object, modify the **Scale in Lightmap** parameter in the [Mesh Renderer](class-MeshRenderer.md) component. This acts as a multiplier of the **Lightmap Resolution** setting.

**Note:** Increasing lightmap resolution significantly increases baking times and memory usage.

Recommended lightmap resolution settings:

* Use **Low Resolution** for:
  + Surfaces receiving uniform or indirect lighting.
  + Small objects.
  + Surfaces not visible to the player.
* Use **High Resolution** for:
  + Areas of high importance.
  + Surfaces receiving baked direct lighting and shadows.

#### Adjust the Max Lightmap Size property

The [Max Lightmap Size](Lightmaps-reference.md) property limits the final resolution of the lightmap. Increase this setting if:

* Objects exceed the maximum space in the lightmap atlas.
* Unity generates a high number of low-resolution lightmaps even after you increase the lightmap resolution.

#### Maintain uniform lightmap UV scale

To prevent stretched or distorted lightmaps, do the following:

* When you create custom lightmap UVs, scale UV charts uniformly across all axes.
* Avoid scaling individual UV charts disproportionately.
* When you scale GameObjects in the editor, ensure uniform scaling across all axes.
* Use the baked lightmap [Scene View Debug Draw Mode](GIVis.md) to inspect lightmap uniformity, and enable **Show Lightmap Resolution** to visualize texel distribution.

![Lightmap texels stretched horizontally resulting in distorted output - left. Uniformly scaled lightmap texels - right.](../uploads/Main/maintain-uniform-lightmap-uv-scale.jpg)

Lightmap texels stretched horizontally resulting in distorted output - left. Uniformly scaled lightmap texels - right.

## Lightmap aliasing

A stair-stepping effect can appear around edges, especially in high-contrast areas with baked direct lighting.

To address lightmap aliasing, do the following:

* In the **Lightmap Parameters Asset**, increase **Anti-aliasing Samples**. For more information, refer to [Lightmap Parameters Asset](class-LightmapParameters.md).
* In the **Inspector** window for a light, adjust the **Baked Shadow Angle** property. For more information, refer to [Light component Inspector window reference](urp/light-component.md).
* In the **Lighting** window, set **Filtering** to **Advanced**, then under **Direct Filter** adjust **Radius**.

## Area distortion

Uneven lighting and artifacts can appear due to improper scaling of UV charts, leading to lower resolution in certain areas.

### Symptoms

* Uneven lighting across the surface.
* Blurry or pixelated lightmap artifacts in some areas.
* Lower resolution in some regions despite adequate lightmap settings.

**Example:**

In the image below, two spotlights with the same parameters light the sides of a cylinder. The right-hand side has a higher **Area Error value**, in the ****Mesh** Import** and **Generate Lightmap UVs** settings, which distorts the triangles and leads to a lower resolution, creating artifacts in the light.

![Cylinder with distortions.](../uploads/Main/LightingGiUvs-GeneratingLightmappingUVs-4.png)
![Cylinder with distortions.](../uploads/Main/LightingGiUvs-GeneratingLightmappingUVs-5.jpg)

### Cause

Area distortion occurs when UV charts are scaled improperly, leading to uneven texel distribution. This can happen when:

* The **Area Error** value is too high, causing variation in triangle sizes.
* UV charts are disproportionate to the object’s geometry, causing stretched or compressed areas.
* The automatic UV unwrapping process fails to balance texel distribution.

### Resolution

To minimize area distortion, do the following:

* Reduce the **Area Error** value in the UV settings.
* Inspect the UV layout to ensure charts are proportionally scaled.
* Manually adjust problematic UV charts in a 3D modeling application if necessary.
* Test different lightmap resolutions for optimal balance between texel density and memory usage.
* Re-bake the lightmap and check for improvements in resolution consistency.

## Angle distortion

Lighting artifacts can appear due to improperly shaped UV charts, leading to skewed or stretched textures in the baked lightmap.

### Symptoms

* Warped or stretched lighting artifacts in baked lightmaps.
* Inconsistent shading or lighting across different areas of the same mesh.
* Skewed or distorted UV charts in the layout.

**Example:**

The first image has a high **Angle Error**, resulting in artifacts. The second image has the default **Angle Error** (8%). Meshes with more triangles can experience significant angle distortion.

![High angle error.](../uploads/Main/LightingGiUvs-GeneratingLightmappingUVs-2.png)
![Default angle error (8%).](../uploads/Main/LightingGiUvs-GeneratingLightmappingUVs-3.jpg)

### Cause

Angle distortion occurs when the automatic UV unwrapping process fails to maintain correct angles between triangles. This can happen when:

* The **Angle Error** value, in the **Mesh Import** and **Generate Lightmap UVs** settings, is too high, causing excessive distortion in UV charts.
* The lightmap UV generation prioritizes minimizing seams over preserving accurate angles.
* The mesh has complex geometry, making it difficult to create an ideal UV layout.

### Resolution

To reduce angle distortion, do the following:

* Lower the **Angle Error** value in the UV settings.
* Check the UV charts to ensure they accurately represent the mesh geometry.
* Manually adjust the UV layout in a 3D modeling tool if necessary.
* Re-bake the lightmap and inspect for improvements in lighting accuracy.

## Additional resources

* [Generate lightmap UVs](LightingGiUvs-GeneratingLightmappingUVs.md)
* [Debug Draw Modes for lighting reference](GIVis.md)
* [Lightmap Data Format](Lightmaps-TechnicalInformation.md)