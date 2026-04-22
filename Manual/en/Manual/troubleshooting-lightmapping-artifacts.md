# Troubleshooting lightmapping artifacts

Fix baked lighting artifacts such as dark spots, smudges, shadow leaks, blocky textures, fireflies, or incorrect shading when using baked **Global Illumination** (GI).

## Lightmap UV artifacts

Fix distorted or warped surfaces and shading or darkening issues.

When importing a 3D model, generating lightmaps, or setting up baked lighting in a scene, improperly configured or missing **lightmap** UVs can cause visual distortions on lightmapped surfaces.

### Symptom

Artifacts caused by incorrect lightmap UVs typically appear as:

* Distorted or warped surfaces on objects.
* Shading or darkening issues, particularly around edges.

**Example:** Severe artifacts on objects with missing lightmap UVs (left) compared to correctly authored lightmap UVs (right).

![Severe artifacts on objects with missing lightmap UVs (left) compared to correctly authored lightmap UVs (right). ](../uploads/Main/lightmap-uv-artifacts.jpg)

Severe artifacts on objects with missing lightmap UVs (left) compared to correctly authored lightmap UVs (right).

### Cause

Lightmap UVs have specific requirements distinct from [standard texture UVs](mesh-vertex-data.md#uvs):

* No overlapping UV charts.
* Adequate margin between UV charts to prevent light bleed.
* Minimal distortion in area and angles.
* Uniform scaling across UV charts.
* UV charts confined within the [0,1] UV space.

Failure to meet these criteria can result in artifacts after baking the lightmap.

### Resolution

You can resolve these issues by addressing improperly configured lightmap UVs, invalid texels, **compression** problems, **light probe** ringing, or denoising and filtering artifacts.

**Note:** Use the **UV Overlap** and **Baked Lightmap** [debug draw modes](GIVis.md) to identify lightmap UV issues.

To resolve lightmap UV problems, follow these steps.

#### Generate optimized lightmap UVs

To generate lightmap UVs, refer to [Generate lightmap UVs](LightingGiUvs-GeneratingLightmappingUVs.md).

To optimize generated lightmap UVs, in [Lightmap UVs Settings](LightingGiUvs-Reference.md), set **Pack Margin** to **Calculate**. If your lightmap fulfills the **Min Lightmap Resolution** and your object fulfills the **Min Object Scale**, Unity ensures there’s no bleeding between charts by conservatively calculating a pack margin large enough to prevent leaks.

**Note**: Large margins can cause low-texture packing.

![Comparison of a Cornell box in shaded view. Dark artifacts around the objects’ edges caused by the UV overlaps are present when Pack Margin is Manual. The artifacts arent visible when Pack Margin is Calculate.](../uploads/Main/lightmapping-pack-margin.png)

Comparison of a Cornell box in shaded view. Dark artifacts around the objects’ edges caused by the UV overlaps are present when Pack Margin is Manual. The artifacts aren’t visible when Pack Margin is Calculate.

#### Create custom lightmap UVs

For complex models, manually authoring lightmap UVs in a digital content creation (DCC) program may yield better results. Ensure the UVs are placed in the second (UV1) channel.

#### Increase lightmap resolution

Use the lowest resolution that resolves the artifacts to minimize memory usage and bake times.

## Invalid texel artifacts

Fix blocky artifacts in baked lightmaps and red-highlighted texels in the **Texel Validity** **scene view** mode.

When baking a lightmap or optimizing a scene in the Editor’s **Lighting** or **Scene** view, invalid texels in lightmaps can cause blocky artifacts.

### Symptom

Scenes with invalid texels may display:

* Blocky artifacts in baked lightmaps.
* Red-highlighted texels in the **Texel Validity** scene view mode.

**Example:** One-sided curtain **mesh** causing invalid texel artifacts after baking (left). The **Texel Validity** scene view mode highlights these invalid texels in red (right), particularly around the curtain area.

![One-sided curtain mesh causing invalid texel artifacts after baking (left)](../uploads/Main/invalid-texel-artifacts.jpg)

One-sided curtain mesh causing invalid texel artifacts after baking (left)

### Cause

Invalid texels occur when:

* Texels fail to receive light from direct, indirect, or environment sources.
* Rays originating from texels intersect back-facing surfaces.

Filtering and denoising can spread invalid texels to neighboring areas. Resolving them early minimizes reliance on **post-processing**.

### Resolution

**Note:** Use the **Texel Validity** [scene view draw mode](GIVis.md) for effective debugging.

#### Adjust mesh normals

Ensure visible polygons have outward-facing normals to avoid shading errors.

Vertex normal direction determines which side of the polygon appears as the front (visible) side. Flipping the direction of the normal renders the face invisible unless you are using double-sided **shaders**. Always ensure that the polygons, that you want to be visible, have their normals pointing outwards.

**Example:** Proper shading with outward-facing normals on the center cube. Blue polygons in the cube on the left and orange polygons in the cube on the right are flipped. Cubes on the left and right both exhibit invalid texel artifacts after baking, whereas the one in the center doesn’t.

![Proper shading with outward-facing normals on the center cube](../uploads/Main/adjust-mesh-normals.jpg)

Proper shading with outward-facing normals on the center cube

Most 3D modeling applications generate correct normals automatically.

To manually adjust normals, refer to your 3D modeling package documentation.

#### Avoid open faces

To reduce the likelihood of invalid texels with watertight meshes, follow these steps:

* Cap open faces.
* Weld overlapping vertices.
* Avoid using line or point geometry.
* Limit floating geometry to necessary cases.

Tools like xView in 3ds Max can help identify mesh issues.

#### Enable double-sided Global Illumination

For thin, one-sided meshes like curtains or foliage, enable **Double-Sided Global Illumination**.

**Example:** Enabling double-sided GI resolves invalid texels and enhances light bounce for one-sided meshes. The Texel Validity scene view no longer shows any invalid texels in red around the curtains (right). Notice that not only the artifacts are gone, but the light also bounces off curtains.

![Enabling double-sided GI resolves invalid texels and enhances light bounce for one-sided meshes](../uploads/Main/enable-double-sided-global-illumination.jpg)

Enabling double-sided GI resolves invalid texels and enhances light bounce for one-sided meshes

To enable **Double-Sided Global Illumination**, follow these steps:

* For the Built-in **Render Pipeline**, enable the **Double-Sided Global Illumination** checkbox in the [Material Inspector](StandardShaderMaterialParameters.md).

  This step also applies to custom ShaderGraph materials.
* For the Universal Render Pipeline (URP), set the **Render Face** dropdown to **Both**. You can find this property under the Surface category in the [Lit Shader](https://docs.unity3d.com/Packages/com.unity.render-pipelines.universal@latest/?subfolder=/manual/lit-shader.html).
* For the High Definition Render Pipeline (HDRP), navigate to the **Surface** category in the [Lit Shader](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest/?subfolder=/manual/Lit-Shader.html) and configure the following properties:

  + **Double-Sided**: Enable this checkbox to disable backface culling. This also enables double-sided GI.
  + **Double-Sided GI**: Use this dropdown to control double-sided GI behavior.
    - **Auto**: Enables double-sided GI if the **Double-Sided** property is enabled.
    - **On**: Forces double-sided GI to be enabled, overriding the double-sided GI behavior without making the material appear double-sided.
    - **Off**: Disables double-sided GI, overriding the double-sided GI behavior without making the material appear double-sided.

You can also control the double-sided GI property via the [kDoubleSidedEnable API](https://docs.unity3d.com/Packages/com.unity.render-pipelines.high-definition@latest/?subfolder=/manual/api/UnityEngine.Rendering.HighDefinition.HDMaterialProperties.html#UnityEngine_Rendering_HighDefinition_HDMaterialProperties_kDoubleSidedEnable).

**Example:** Double-Sided Global Illumination disabled (left) and enabled (right). Notice dark artifacts where the meshes intersect in the image on the right. In such cases, disable Double-Sided Global Illumination to achieve correct output.

![Double-Sided Global Illumination disabled (left) and enabled (right)](../uploads/Main/double-sided-global-illumination-disabled.jpg)

Double-Sided Global Illumination disabled (left) and enabled (right)

Lightmapper marks texels as invalid when they receive no illumination or intersect backfaces. This allows the dilation filter to push lighting data from valid surrounding texels into partially covered texels. Doing so helps to remove dark artifacts at geometry junctions.

Enabling double-sided GI by default for all materials would result in a darkening effect where meshes intersect or overlap.

#### Adjust the backface tolerance property

To configure the **Backface Tolerance** property, follow these steps:

1. Create a new [Lightmap Parameters asset](class-LightmapParameters.md).
2. Navigate to the **Lightmapping** section in the [Mesh Renderer](class-MeshRenderer.md) component.
3. Assign the **Lightmap Parameters** asset to the **Mesh Renderer** component.

The **Backface Tolerance** property controls the percentage of rays that must hit front-facing geometry to consider a texel valid:

* Increasing the value increases the likelihood that the **lightmapper** marks a texel as invalid when it encounters backfaces.
* Lowering this value reduces the number of invalidated texels.

**Note:** Low **Backface Tolerance** values may result in a loss of occlusion and/or bounce lighting, potentially altering the appearance of the scene.

#### Move meshes closer to the world origin

Limitations imposed by the 32-bit floating-point precision format can be the cause of invalid texels when placing lightmapped meshes far away from the world origin (0,0,0), or if they have a massive scale. This issue is not limited to lightmapped meshes; objects and animations lose precision the farther away they are from world origin.

To solve precision issues, follow these steps:

* Move meshes closer to the (0, 0, 0) origin.
* Adjust the **Push Off** value in the Lightmap Parameters asset.

## Denoising artifacts

Fix smears, smudges, or ringing around edges and black lightmaps or excessive noise.

When configuring denoising settings to bake a lightmap in the Editor’s **Lighting** or **Scene** view, machine learning denoisers may introduce artifacts.

### Symptom

Denoising artifacts may include:

* Smears, smudges, or ringing around edges.
* Black lightmaps or excessive noise when denoising fails.
* Bright **pixels** across the scene. These artifacts are sometimes called fireflies.

**Example:** Severe artifacts with low sample counts (left), while higher sample counts produce cleaner results (right). Denoising Artifacts with Open Image Denoiser enabled. Severe artifacts are visible when all sample counts are set to 32 with the Open Image Denoiser enabled (left). The image is cleaner with all sample counts set to 2K and the Open Image Denoiser enabled (right).

![Severe artifacts with low sample counts (left); higher sample counts produce cleaner results (right)](../uploads/Main/severe-artifacts-with-low-sample-counts.jpg)

Severe artifacts with low sample counts (left); higher sample counts produce cleaner results (right)

### Cause

Denoising artifacts typically occur in the following scenarios:

* The scene is too complex for the amount of data the denoiser collects.
* The denoiser lacks the necessary training data for a scene like yours.
* You have thin geometry or overlapping edges. When Unity dilates geometry during lightmapping, it might create invalid texels that the denoiser can’t resolve.
* You use high-intensity lights that require high-definition range (HDR) lightmaps.

Indirect noise often contributes significantly to poor results.

### Resolution

#### Increase the sample count

Most denoisers perform effectively with at least 256 samples. You can improve denoising performance by increasing the sample count.

#### Use a different denoiser

To select a specific denoiser based on the scene’s characteristics, follow these steps:

1. Set **Filtering** to **Advanced**.

   Refer to [Progressive GPU](progressive-lightmapper.md).
2. Specify the desired denoiser for direct, indirect, and **ambient occlusion** buffers.

#### Use Gaussian or A-Trous filters

To use Gaussian or A-Trous filters, follow these steps:

1. Set **Filtering** to **Advanced**.

   Refer to [Progressive GPU](progressive-lightmapper.md).
2. Choose the desired filter type.

You can combine these filters with machine learning denoisers for enhanced results.

## Fireflies

Fix speckled noise (fireflies) caused by a bright light source placed close to scene geometry.

### Symptom

![A Cornell box scene, with a single cube in an empty room. On the left, the image is noisy. On the right, the image is clear.](../uploads/Main/fireflies-troubleshooting-example.jpeg)
**Example:** Speckled noise (fireflies) visible (left). After baking with 16K indirect samples and enabling **Open Image Denoiser** to remove the noise (right).

Speckled noise (fireflies) can be visible in brightly lit small areas.

### Cause

Small, high luminosity areas in the lightmap cause speckled noise to appear. Such areas can be missed when spawning rays from texels. When rays intersect those high intensity areas, they skew the average luminosity estimate. This results in speckled noise which are often referred to as “fireflies”.

### Resolution

**Note:** The splat clamping feature is being developed to addressing this issue for punctual lights. Refer to the following solutions to resolve issues for now.

### Increase sample counts

This is a brute force method. To achieve acceptable results, use 4K or higher indirect sample in conjunction with filtering and denoising.

### Enable importance sampling

To reduce noise in scenes with high frequency **HDRIs**:

![A Cornell box scene, with two cuboids in an empty room, and a skylight. The left image is noisy with pixels from a skybox of a landscape at sunset, which is shown as a sphere on the right. The middle image is clear.](../uploads/Main/importance-sampling-ex1.jpeg)
**Example:** Fireflies caused by the high frequency HDRI (left). The same scene after enabling the **Importance Sampling** toggle (center). High frequency **cubemap** used to illuminate the scene (right).

![A Cornell box scene, with two cuboids in an empty room, and a skylight. The left image is clear. The middle image is noisy with pixels from a skybox of a river scene, which is shown as a sphere on the right.](../uploads/Main/importance-sampling-ex2.jpeg)
**Example:** Enabling **Importance Sampling** can introduce noise when using low frequency HDRIs. **Importance Sampling** disabled (left). **Importance Sampling** enabled (center). Low frequency cubemaps used to illuminate the scene (right).

Enable importance sampling if:

* You are using a high frequency HDRI image for environment lighting. An example of this would be an HDRI with a bright sunspot in the horizon.

Disable importance sampling if:

* You are using a low frequency HDRI image for environment lighting. An example of this would be an overcast HDRI or a solid color **skybox**.

**Warning:** Enabling importance sampling when using low frequency HDRIs can result in increased noise levels instead.

### Enable filtering and denoising

To remove the noise, enable ML denoisers in conjunction with filtering. To get clean results, use high sample counts.

### Avoid small baked emissive objects

To reduce issues when using emissive materials, reduce the number of small baked emissive objects by:
1. Set the **Global Illumination** property of the **GameObject** to None in the **Material **Inspector****.

The object will keep its self-illuminating appearance, but it won’t contribute to global illumination. Reducing the number of small baked emissive objects can help reduce noise.

### Move baked lights away from static geometry

URP and HDRP use inverse square falloff for lights by default. In the Built-in render pipeline, you can achieve the same effect by using custom falloff **scripts**. The unfortunate side effect being that such lights might lead to increased noise levels when placed very close to static geometry. Ensure that there is enough distance between the static geometry and the lights in the scene.

## Filtering artifacts

Fix Gaussian light leaks, blurs, and splotches, A-Trous speckles, and **HDR** denoise failures.

When baking a lightmap in the Editor’s Lighting or Scene view, using a Gaussian or an A-Trous filter to reduce noise can create artifacts.

### Symptom

Filtering artifacts depend on the filter type:

**Example:** Scene filtered using Gaussian filter at a 1 texel radius. With all samples set to 32, the final lightmap exhibits splotches (left). Baking with all samples set to 2K and with Gaussian filtering enabled yields better results (right).

![Scene filtered using Gaussian filter at a 1 texel radius](../uploads/Main/gaussian-filter-at-1-texel-radius.jpg)

Scene filtered using Gaussian filter at a 1 texel radius

**Example:** Scene filtered using A-Trous filter. With all samples set to 32, A-Trous edge preservation algorithm cannot distinguish between noise and edges, and thus completely fails (left). After increasing all samples to 2K, A-Trous filter yields decent results, with less blurring when compared to Gaussian (right).

![Scene filtered using A-Trous filter](../uploads/Main/a-trous-filter.jpg)

Scene filtered using A-Trous filter

### Cause

Filtering artifacts causes depend on the filter type:

* **Gaussian Filter**:

  + Light or shadow leakage through geometry.
  + Blurry lightmaps.
  + Splotchy lightmaps.
* **A-Trous Filter**:

  + Speckled noise at low sample counts.
  + Failure to denoise HDR lightmaps (noise interpreted as edges).

### Resolution

#### Increase sample count

You can reduce artifacts such as splotches or speckles by increasing sample counts.

#### Use a machine learning denoiser

Machine learning denoisers effectively remove noise with fewer biases compared to traditional filters. Combining them with filters can enhance results:

* **A-Trous**: Works well with machine learning denoisers to preserve edges.
* **Gaussian**: Performs better with a low radius setting.

#### Adjust filtering settings

To adjust the filtering strength, adjust the filter’s **radius** and **sigma. If your scene has lightmap leaking artifacts or blurring, reduce the** Gaussian\*\* filter’s **radius** value. This reduces the amount of bilateral blurring and prevents texel smearing.

Refer to [Progressive GPU](progressive-lightmapper.md).

## Compression artifacts

Fix color banding and blocky artifacts.

When optimizing lightmaps for mobile performance during lightmap baking or asset compression in the Editor’s **Lighting** or Inspector windows, artifacts can occur.

### Symptom

Color banding and blocky artifacts.

**Example:** Banding artifacts introduced by low-quality compression on Android platform (left). You can fix banding artifacts by using high-quality compression (right).

![Banding artifacts introduced by low-quality compression on Android platform (left)](../uploads/Main/banding-artifacts.jpg)

Banding artifacts introduced by low-quality compression on Android platform (left)

### Cause

Color banding and blocky artifacts are often associated with low-quality **texture compression**. Unity attempts to select automatically the best compression format for the target platform. Due to compatibility concerns, **texture formats** intended for mobile platforms are often more affected by compression artifacts.

### Resolution

#### Adjust lightmap encoding

To adjust lightmap encoding, follow these steps:

1. Open **Edit** > **Project Settings** > **Player** > **Other Settings**.
2. Modify the **Lightmap Encoding** setting for each target platform.

#### Configure HDR cubemap encoding

To configure HDR cubemap encoding for reflective objects with compression artifacts, follow these steps:

1. Open **Edit** > **Project Settings** > **Player** > **Other Settings**.
2. Modify the **HDR Cubemap Encoding** setting.

**Note:** The **HDR Cubemap Encoding** setting is available in Unity 2022.1 and later.

#### Modify lightmap compression

To control compression quality, follow these steps:

1. Go to **Window > Rendering > Lighting**.
2. In the **Lighting** window, locate the **Lightmaps** tab.
3. Find the **Lightmap Compression** dropdown under the **Lightmaps** settings section.
4. Select a compression quality.

   Higher quality presets result in fewer compression artifacts post-bake.

#### Manually override texture formats

To manually override texture formats, follow these steps:

1. Locate the lightmap textures in the **Project** window.
2. Select a lightmap texture.
3. In the **Inspector** window, enable the **Override** checkbox for the desired platform.
4. Select the desired compression format.
5. Click **Apply**.

**Note:** Custom overrides reset after rebaking lighting. Repeat these steps after each bake.

## Light probe ringing artifacts

Fix unnatural probe lighting and visible rippling artifacts.

When setting up light probes or baking lightmaps in the Editor’s **Lighting** or **Scene** view, ringing artifacts in light probes can cause incorrect shading.

### Symptom

* Objects lit by light probes exhibit unnatural lighting.
* Visible rippling or artifacts around light probes.

**Example:** Remove Ringing disabled (left) and enabled (right). Notice how negative light probe values result in wrong shading for the statue on the left. You can fix this issue by enabling the Remove ringing option, but as a result, the final output can appear very smooth.

![Remove Ringing disabled (left) and enabled (right)](../uploads/Main/remove-ringing.jpg)

Remove Ringing disabled (left) and enabled (right)

### Cause

For performance reasons, Unity encodes light probes using Spherical Harmonics. However, this lossy compression technique can’t faithfully represent environments where lighting changes quickly across the sphere of directions. In cases such as light entering through a small hole in a cave roof, Spherical Harmonics can yield brighter or dimmer illumination than expected. This is called overshooting, or ringing. For more information, refer to [Gibbs phenomenon](https://en.wikipedia.org/wiki/Gibbs_phenomenon).

### Resolution

**Note:** To make debugging easier, enable **Gizmos** in the [scene](overlays-reference-view-options.md) **viewport** and adjust the **Light Probe Visualization** property in the [workflow settings](lighting-window.md#WorkflowSettings).

#### Remove light probe ringing

To remove light probe ringing, follow these steps:

1. Select the ****Light Probe Group**** component.
2. Enable the **Remove Ringing** property.

**Note:** This option reduces contrast in probe-lit objects. Disable it if your scene does not exhibit ringing artifacts.

#### Adjust scene lighting

To mitigate ringing artifacts caused by high-intensity direct lighting, follow these steps:

* Lower the **Intensity** value in the **Light** component.
* Use [mixed lights](LightModes-introduction.md#mixed) to balance real-time direct lighting and baked indirect lighting.

## Hard edges in lightmaps

Smooth hard edges and lighting discontinuities at seams and chart boundaries.

### Symptom

When using baked lightmaps with the [Progressive Lightmapper](progressive-lightmapper.md), you may notice:

* Hard edges appearing along seams between separate mesh faces.
* Visible discontinuities in lighting at chart boundaries in the lightmap.

![Hard edges in lightmaps](../uploads/Main/stitch_off.jpg)

Hard edges in lightmaps

### Cause

When Unity bakes lightmaps, it treats mesh faces that are close together but separate as distinct in lightmap space, creating edges known as seams.

Ideally, seams are invisible, but they can sometimes appear with hard edges. This happens because the GPU cannot blend texel values between separated charts in the lightmap, treating them as distinct entities and causing visible discontinuities.

### Resolution

Seam stitching smooths hard edges on GameObjects rendered with baked lightmaps from the Progressive Lightmapper. When enabled, Unity performs additional calculations to improve the appearance of seams in the lightmap. While stitching isn’t perfect, it often significantly enhances the final result. However, this feature increases baking time, so it is disabled by default.

When active, the lightmapper identifies pairs of edges to stitch and smooths the illumination across seams. This applies only to straight edges that run horizontally or vertically along chart boundaries in the atlas and works with axis-aligned rectangles in UV space.

**Note:**

* Seam stitching works only with the [Progressive Lightmapper](progressive-lightmapper.md).
* Seam stitching applies to single GameObjects only; it cannot smooth seams between multiple GameObjects.

To smooth hard edges:

1. Select a GameObject with a Mesh Renderer component in the scene.
2. In the **Mesh Renderer Inspector**, in the **Lightmapping** section, enable **Stitch Seams**.

Alternatively, use the [MeshRenderer.stitchLightmapSeams](../ScriptReference/MeshRenderer-stitchLightmapSeams.md) API to programmatically enable seam stitching.

**Example:** A Scene with seam stitching.

![Seam stitching](../uploads/Main/stitch_on.jpg)

Seam stitching

## Light bleeding in lightmaps

Fix lightmap aliasing, UV overlaps, and red highlight warnings.

### Symptom

Light bleeding typically manifests as:

* Aliasing or pixelation in the rendered lightmap.
* Overlapping UVs detected by Unity, with warning messages in the console.
* Red highlights in the **UV Overlap** draw mode or **Baked Lightmaps Preview**.

**Example:** Red pixels indicate overlapping chart neighbourhoods.

![Light bleeding in lightmaps](../uploads/Main/ProgressiveLightmapper-UVOverlap-183-1.png)

Light bleeding in lightmaps

### Cause

Light bleeding occurs when data from one lightmap chart spills onto another due to insufficient spacing between charts, leading to graphical artifacts such as aliasing or pixelation.

Each lightmap comprises multiple charts, which Unity maps to mesh faces at runtime. These charts store lighting data used to calculate the final appearance. During rendering, the GPU samples lighting data by interpolating the values of the four nearest texels, known as the bilinear neighborhood. When charts are placed too closely, their neighborhoods overlap, causing light bleeding. Factors influencing this issue include lightmap resolution, mesh UVs, and Importer settings. Unity provides tools to detect and address these problems.

**Example:** Graphical artifacts due to chart bleeding.

![Graphical artifacts due to chart bleeding](../uploads/Main/ProgressiveLightmapper-UVOverlap-183-0.png)

Graphical artifacts due to chart bleeding

Light bleeding is caused by overlapping chart neighborhoods in the lightmap. This occurs when there is insufficient spacing between charts, often due to:

* Insufficient margins in user-provided lightmap UVs.
* Small pack margins in Unity-generated lightmap UVs.
* Low lightmap resolution.
* A high density of charts within a single lightmap.

### Resolution

To resolve light bleeding, follow these steps.

#### Adjust lightmap UV margins

* If providing UVs manually, increase the margins using your 3D modeling software.
* For Unity-generated UVs, go to the Model Import Settings and:

  + Set the **Margin Method** to **Calculate** and adjust **Min Lightmap Resolution** and **Min Object Scale**.
  + Alternatively, set the **Margin Method** to **Manual** and increase the **Pack Margin** value.

  Refer to the [Generating lightmapping UVs](LightingGiUvs-GeneratingLightmappingUVs.md) documentation for more details.

#### Increase lightmap resolution

* To increase resolution globally, adjust the **Lightmapper Settings** in the **Lighting** tab.
* To increase resolution for a specific GameObject, modify its **Lightmap Settings** in the Mesh Renderer component.

#### Use Unity’s visualization tools to identify overlaps

* Enable the **UV Overlap** draw mode in the Scene View to highlight problematic texels.

  **Example:** Scene View using UV Overlap draw mode (see dropdown in top left).

  ![UV Overlap draw mode](../uploads/Main/ProgressiveLightmapper-UVOverlap-183-2.png)

  UV Overlap draw mode
* Use the **Baked Lightmaps Preview** to view and diagnose UV overlaps in the **Lighting** window.

  **Example:** The Baked Lightmaps Preview in the Lighting window’s Baked Lightmaps tab.

  ![Baked Lightmaps Preview](../uploads/Main/ProgressiveLightmapper-UVOverlap-183-3.png)

  Baked Lightmaps Preview

By implementing these solutions, you can minimize or eliminate light bleeding artifacts.

**Example:** Same mesh as before, but without bleeding artifacts.

![Same mesh as before, but without bleeding artifacts](../uploads/Main/ProgressiveLightmapper-UVOverlap-183-4.jpg)

Same mesh as before, but without bleeding artifacts

## Additional resources

* [Double-Sided Global Illumination](../ScriptReference/Material-doubleSidedGI.md)
* [GI visualizations in the Scene View](GIVis.md)
* [Generating lightmapping UVs](LightingGiUvs-GeneratingLightmappingUVs.md)
* [Lightmap Parameters Asset](class-LightmapParameters.md)
* [MeshRenderer.stitchLightmapSeams API](../ScriptReference/MeshRenderer-stitchLightmapSeams.md)
* [Progressive Lightmapper](progressive-lightmapper.md)