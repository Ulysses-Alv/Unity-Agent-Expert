# Grass and other details

A **Terrain** might have grass clumps and other small objects (such as rocks) covering its surface. Unity renders these objects using textured **quads** or full meshes, depending on the **level of detail** and performance you require.

Instancing details work with all **render pipelines**, including the [High Definition Render Pipeline](https://docs.unity3d.com/Manual/high-definition-render-pipeline.html) (HDRP).

![Terrain with grass](../uploads/Main/TerrainWithGrass.png)

Terrain with grass

Terrain details are available in several modes, where each mode has a different use during game development.

| **Mode** | **Description** |
| --- | --- |
| Instanced mesh | Unity recommends you use this mode because it’s best suited for most scenarios that require the placement of any number of arbitrary meshes on Terrain.  To add an instanced mesh, select **Add Detail Mesh** and enable the **Use GPU Instancing** option. |
| Vertex Lit mesh | Doesn’t use GPU instancing for rendering. Instead, it combines all detail instances into one mesh, which results in simple shading and limitations to the number of instances you can achieve.  To add a Vertex Lit mesh, select **Add Detail Mesh** and set the **Render Mode** to **Vertex Lit**. |
| Grass mesh | Works similar to the vertex list mesh, but Unity treats these meshes as grass so they always have up-facing normals and move in the wind.  To add a grass mesh, select **Add Detail Mesh** and set the **Render Mode** to **Grass**. |
| Grass Texture | Lets you create grass quad meshes directly from Texture, which move in the wind. You don’t need to specify any custom meshes, and there’s also an option you can enable to make grass quads always face the camera for a billboard effect.  Select **Add Grass Texture** for this option. |

## Paint Details

To enable grass and detail painting, select the **Paint Details** button on the Terrain **Inspector** **toolbar**.

![Paint Details in the Terrain Inspector](../uploads/Main/TerrainDetailsPaintTool.png)

Paint Details in the Terrain Inspector

To access **Paint Details** from an overlay:

1. In the **Terrain Tools** overlay, select **Foliage Mode** ![Foliage Mode Menu](../uploads/Main/terrainOverlays-FoliageModeMenuButton.png). Foliage Mode tools display at the end of the **Terrain Tools** overlay.
2. From the available Foliage Mode tools on the **Terrain Tools** overlay, select **Paint Details** ![Paint Details](../uploads/Main/terrainOverlays-PaintDetailsButton.png).

Initially, a Terrain has no grass or details available. In the Inspector, click the **Edit Details** button to display a menu with the **Add Detail **Mesh**** and **Add Grass Texture** options. Click either option to bring up a window that lets you choose **Assets** to add to the Terrain for painting.

### Add Detail Mesh

![The Add Detail Mesh window](../uploads/Main/TerrainDetailAddWindow.png)

The Add Detail Mesh window

Use the **Detail **Prefab**** field to select a Prefab from your Project. To create an authentic look, a random factor affects each mesh to determine the size and color. Unity uses the Perlin noise algorithm to generate such random factors.

The **Min Width**, **Max Width**, **Min Height**, and **Max Height** values specify the upper and lower scalings along the X and Y axes with which the meshes are randomly generated.

The **Align to Ground (%)** value specifies how much the detail axis aligns to the terrain normal. 0 is unaligned, and 100 is aligned to the normal.

The **Position Jitter (%)** controls the randomness of the detail distribution from ordered to random. This is only available when legacy distribution is turned off in [Quality settings](class-QualitySettings.md#Terrain).

The **Noise Seed** setting sets the random number generator seed. Each seed value represents a unique generation. If you set the same seed for two types of details, the generated instances become identical. The **Noise Spread** value refers to the scaling of the noise pattern over the X-Z plane, with higher values indicating more variations within a given area.

The **Hole Edge Padding (%)** setting controls how far away detail objects are from the edge of the hole area. For more information, see [DetailPrototype.holeEdgePadding](https://docs.unity3d.com/2021.2/Documentation/ScriptReference/DetailPrototype-holeEdgePadding.html).

The **Detail Density** value controls the detail’s density relative to its size. You can enable this in “Coverage” detail scatter mode, set in [Terrain Settings](terrain-OtherSettings.md).

**Healthy Color** and **Dry Color** settings control the color variation. Unity performs linear interpolation between the healthy and dry colors, and selects a color based on the size of the mesh. For example, the system considers bigger-sized meshes to be healthier. The linear interpolation results for such meshes lean more towards the **Healthy Color** tint. These color settings disappear when you select **Use GPU Instancing** because they have no effect on instanced meshes.

The **Affected by Density Scale** checkbox determines whether the Detail Density Scale setting set in the [Terrain Settings](terrain-OtherSettings.md) affects this detail.

You can set the **Render Mode** to **Vertex Lit** or **Grass**.

| **Render Mode** | **Description** |
| --- | --- |
| **Vertex Lit** | Renders detail meshes as solid, vertex lit **GameObjects** in the **scene**, which don’t move in the wind. |
| **Grass** | Renders detail meshes in the scene with simplified lighting similar to Grass Textures, and do move in the wind. |

**Note**: To apply a texture on a detail using Vertex Lit or Grass render mode, assign a texture to the **MainTex** material property (default property name for Base Color Map in built-in/URP). No other material properties are considered for non-instanced details.

#### Use GPU Instancing

You can specify if you want to use GPU instancing to render your detail meshes. With GPU instancing, Unity uses the material and the **shader** specified on the prefab for rendering. This lets you perform more customization in the shader and use [Shader Graph](https://docs.unity3d.com/Packages/com.unity.shadergraph@latest).

Unity still applies the Width and Height noise to each instance, but disables the **Healthy Color**/**Dry Color** noise. You can use your own color variation technique, along with other randomizations, in the shader.

Under the hood, some Unity [GPU instancing](GPUInstancing.md) limitations do affect rendering. For example:

* Your target platform needs to support GPU instancing.
* Your shader must support GPU instancing. (Most shaders do support it.)
* Objects are rendered in batches of 1,023 or fewer.
* Instanced **light probe** lighting or **lightmap** lighting aren’t supported.

Unlike regular instancing where Unity dynamically builds up the instancing constant buffer every frame, instancing details have persistent constant buffers for better CPU and GPU performance when the target hardware allows for it, at the cost of moderately increased GPU memory usage.

If you select **Use GPU Instancing**, the **Render Mode** setting becomes grayed out and unavailable.

### Add Grass Texture

![The Add Grass Texture window](../uploads/Main/TerrainGrassAddWindow.png)

The Add Grass Texture window

The Add Grass Texture window

**Detail Texture** is the Texture that represents the grass. You can download Textures from the **Asset Store**, or create your own Textures. A Texture is a small image with alpha set to zero for the empty areas. Note that “Grass” is a generic term; it’s possible for a Texture to represent flowers or man-made objects such as barbed wire coils.

The **Min Width**, **Max Width**, **Min Height**, **Max Height**, **Noise Seed**, **Noise Spread**, **Hole Edge Padding (%)**, **Align to Ground (%)**, **Position Jitter (%)**, **Detail Density**, **Affected by Density Scale**, **Healthy Color**, and **Dry Color** settings work the same as they do for meshes. The Add Detail Mesh section above provides more details.

When you enable the ****Billboard**** option, the grass images rotate so that they always face the **Camera**. This is useful when you want to display a dense field of grass because clumps are two-dimensional, and not visible from the side. However, with sparse grass, the rotations of individual clumps might become visible, which can create a strange visual effect.

## Remove details

To remove a detail, select the detail to remove, click the **Edit Details** button, then select **Remove**. When you remove a detail, Unity clears the data that represents the Terrain area on which you scattered that detail.

---

* 2022–07–20 Page amended
* Added information about detail density and distribution.