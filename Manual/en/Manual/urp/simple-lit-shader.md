# Simple Lit shader material Inspector window reference for URP

Explore the properties and settings you can use to customize the default Simple Lit **shader** in the Universal **Render Pipeline** (URP).

**Note:** URP implements certain Lit shader material features as shader variants, which require more [**SRP Batcher**](../SRPBatcher.md) batches. For optimal performance, disable unused features.

## Surface Options

The **Surface Options** section controls how URP renders the material on-screen.

| **Property** | **Sub-property** | **Description** |
| --- | --- | --- |
| **Surface Type** |  | Sets whether the surface is opaque or transparent. The options are the following:  * **Opaque**: The surface is fully opaque and doesn’t use transparency blending. URP renders opaque surfaces first. * **Transparent**: Blends the surface with the background. URP renders transparent surfaces in a separate render pass after opaque surfaces.  This property affects batching performance. |
| **Blending Mode** |  | Sets how the shader blends the color of a transparent material with the background. This property is available only when you set **Surface Type** to **Transparent**. The options are:  * **Alpha**: Calculates transparency using the alpha value of the material. * **Premultiply**: Calculates transparency using the alpha value of the material, but multiplies the RGB values by the alpha value first. * **Additive**: Adds the color of the material to the background color. * **Multiply**: Multiplies the color of the material with the background color.  For more information about blending modes, refer to [Blending modes](blending-modes.md). This property affects batching performance. |
| **Preserve Specular Lighting** |  | Keeps specular highlights on transparent surfaces by not applying the alpha value. The specular highlights appear even if the material is fully transparent. This property is available only if you set **Surface Type** to **Transparent** and **Blending Mode** to either **Alpha** or **Additive**. |
| **Render Face** |  | Specifies which faces of the mesh the shader renders. The options are:  * **Front**: Renders only front‑facing triangles. This is the default. * **Back**: Renders only back‑facing triangles. * **Both**: Renders both front and back faces. |
| **Alpha Clipping** |  | Discards **pixels** if their alpha value is lower than the **Threshold** value. The default value is 0.5. Use this property to create hard edges between opaque and transparent areas, for example to create blades of grass. This property affects batching performance. |
|  | **Threshold** | Sets the minimum alpha value for a pixel to be visible. The default value is 0.5. This property is only available if you enable **Alpha Clipping**. |
| **Receive Shadows** |  | Determines whether the material receives and displays shadows cast from other objects. This property affects batching performance. |

## Surface Inputs

The **Surface Inputs** describe the surface itself. For example, you can use these properties to make your surface look wet, dry, rough, or smooth.

| **Property** | **Sub-property** | **Description** |
| --- | --- | --- |
| **Base Map** |  | Sets the surface color, also known as diffuse. To assign a texture, select the picker (**⊙**). To assign a color, either as the main color or a tint on top of the texture, select the swatch or the dropper.  If you set **Surface Type** to **Transparent** or enable **Alpha Clipping**, Unity uses the alpha channel to calculate the transparency of the surface. |
| **Specular Map** |  | Sets the color of reflections. This property is available only if you set **Workflow Mode** to **Specular**. To assign a texture, select the picker (**⊙**). To assign a color, select the swatch or the dropper. |
|  | **Smoothness** | Sets the smoothness of the surface, which affects the clarity of reflections. The range is 0 to 1, where 0 provides a wide, rough highlight, and 1 provides a small, sharp highlight like glass. For more information, refer to [Configure smoothness](../StandardShaderMaterialParameterSmoothness). |
|  | **Source** | Selects a source to use as a smoothness map. The options are the following:  * **Metallic Alpha**: Uses the alpha channel from the **Metallic Map** texture. This is the default option. * **Albedo Alpha**: Uses the alpha channel from the **Base Map** texture.  Unity multiplies the texture values by the **Smoothness** value. This property affects batching performance. |
| ****Normal Map**** |  | Adds details like bumps, scratches, and grooves. To assign a normal map, select the picker (**⊙**). Use the slider to adjust the intensity of the effect. The range is 0 to 1, where lower values create less visible details. For more information, refer to [Introduction to normal maps (bump mapping)](../StandardShaderMaterialParameterNormalMap.md). This property affects batching performance. |
| **Emission** |  | Enables the surface to emit a color as a light source. This property affects batching performance. |
|  | **Emission Map** | Sets the color the material emits as a light source. The default is black, which means no emission. To assign a texture, select the picker (**⊙**). To assign a color, either as the main color or a tint on top of the texture, select the swatch or the dropper. For effects like lava that glow brighter than white while still being another color, increase the **Intensity** value in the color picker. This property affects batching performance. |
| **Tiling** |  | Scales the surface input textures along their UV axes. The default value is 1, which means no scaling. Set a higher value to make the textures repeat across your **mesh**. Set a lower value to stretch the textures. |
| **Offset** |  | Positions the surface input textures on the mesh. |

## Advanced Options

The **Advanced Options** section controls the rendering calculations Unity uses.

| Property | Description |
| --- | --- |
| **Specular Highlights** | Adds specular highlights from direct lighting, for example [Directional, Point, and Spot lights](../Lighting.md). This means that your material reflects the shine from these light sources. Disable this property to make shaders render faster. This property affects batching performance. |
| **Sorting Priority** | Determines when Unity renders the material. Unity renders materials with lower values first. Use this property to prevent Unity from rendering pixels twice by rendering materials behind other materials. This property works similarly to the [render queue](../../ScriptReference/Material-renderQueue.md) in the Built-In Render Pipeline. |
| **Enable GPU Instancing** | Renders meshes with the same geometry and material in a single batch when possible. This makes rendering faster. URP can’t render meshes in one batch if they have different materials or if the hardware doesn’t support GPU instancing. For more information, refer to [GPU instancing](../GPUInstancing-landing.md). |
| **Alembic Motion Vectors** | Enables motion vector support for meshes streamed through [Alembic](../com.unity.formats.alembic) files. For more information, refer to [Built-in shader support for motion vectors](features/motion-vectors-shader-support.md). |
| ****XR** Motion Vectors Pass (SpaceWarp)** | Enables URP running a motion vector render pass in XR to improve performance. This property is only available if your project meets the prerequisites for [URP Application SpaceWarp](https://docs.unity3d.com/Packages/com.unity.xr.openxr@latest?subfolder=/manual/features/spacewarp.html). |