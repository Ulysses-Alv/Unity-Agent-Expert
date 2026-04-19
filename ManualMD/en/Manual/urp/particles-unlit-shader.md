# Particles Unlit shader material Inspector window reference for URP

Explore the properties and settings you can use to customize the default **Particles** Unlit **shader** in the Universal **Render Pipeline** (URP).

**Note:** URP implements certain Lit shader material features as shader variants, which require more [**SRP Batcher**](../SRPBatcher.md) batches. For optimal performance, disable unused features.

## Surface Options

The **Surface Options** section controls how URP renders the material on-screen.

| **Property** | **Sub-property** | **Description** |
| --- | --- | --- |
| **Surface Type** |  | Sets whether the surface is opaque or transparent. The options are the following:  * **Opaque**: The surface is fully opaque and doesn’t use transparency blending. URP renders opaque surfaces first. * **Transparent**: Blends the surface with the background. URP renders transparent surfaces in a separate render pass after opaque surfaces.  This property affects batching performance. |
| **Blending Mode** |  | Sets how the shader blends the color of a transparent material with the background. This property is available only when you set **Surface Type** to **Transparent**. The options are:  * **Alpha**: Calculates transparency using the alpha value of the material. * **Premultiply**: Calculates transparency using the alpha value of the material, but multiplies the RGB values by the alpha value first. * **Additive**: Adds the color of the material to the background color. * **Multiply**: Multiplies the color of the material with the background color.  For more information about blending modes, refer to [Blending modes](blending-modes.md). This property affects batching performance. |
| **Render Face** |  | Specifies which faces of the mesh the shader renders. The options are:  * **Front**: Renders only front‑facing triangles. This is the default. * **Back**: Renders only back‑facing triangles. * **Both**: Renders both front and back faces. |
| **Alpha Clipping** |  | Discards **pixels** if their alpha value is lower than the **Threshold** value. The default value is 0.5. Use this property to create hard edges between opaque and transparent areas, for example to create blades of grass. This property affects batching performance. |
|  | **Threshold** | Sets the minimum alpha value for a pixel to be visible. The default value is 0.5. This property is only available if you enable **Alpha Clipping**. |
| **Color Mode** |  | Sets how the shader blends the material color and the particle color. For more information, refer to [Color Mode dropdown](#color-mode-dropdown). |

### Color Mode dropdown

| **Option** | **Description** |
| --- | --- |
| **Multiply** | Creates a darker final color by multiplying the two colors. |
| **Additive** | Creates a brighter final color by adding the two colors together. |
| **Subtractive** | Creates a dark effect by subtracting the particle color from the material color. |
| **Overlay** | Blends the particle color over the material color. This option creates a brighter color at values over 0.5, and darker colors at values under 0.5. |
| **Color** | Colorizes the material color with the particle color, while keeping the value and saturation of the material color. Use this option to add splashes of color to monochrome **scenes**. |
| **Difference** | Returns the difference between the color values. Use this option to blend particle and material colors that are similar to each other. |

## Surface Inputs

The **Surface Inputs** describe the surface itself. For example, you can use these properties to make your surface look wet, dry, rough, or smooth.

| **Property** | **Sub-property** | **Description** |
| --- | --- | --- |
| **Base Map** |  | Sets the surface color, also known as diffuse. To assign a texture, select the picker (**⊙**). To assign a color, either as the main color or a tint on top of the texture, select the swatch or the dropper.  If you set **Surface Type** to **Transparent** or enable **Alpha Clipping**, Unity uses the alpha channel to calculate the transparency of the surface. |
| ****Normal Map**** |  | Adds details like bumps, scratches, and grooves. To assign a normal map, select the picker (**⊙**). Use the slider to adjust the intensity of the effect. The range is 0 to 1, where lower values create less visible details. For more information, refer to [Introduction to normal maps (bump mapping)](../StandardShaderMaterialParameterNormalMap.md). This property affects batching performance. |
| **Emission** |  | Enables the surface to emit a color as a light source. This property affects batching performance. |
|  | **Emission Map** | Sets the color the material emits as a light source. The default is black, which means no emission. To assign a texture, select the picker (**⊙**). To assign a color, either as the main color or a tint on top of the texture, select the swatch or the dropper. For effects like lava that glow brighter than white while still being another color, increase the **Intensity** value in the color picker. This property affects batching performance. |
|  | **Global Illumination** | Sets whether Unity bakes the emissive light cast onto other GameObjects, or calculates it in real-time. The options are:  * **Realtime**: Unity calculates the emission in real-time. This property only has an effect if you [Create lightmaps at runtime with Enlighten Realtime Global Illumination](../realtime-gi-using-enlighten-landing.md). * **Baked**: Unity bakes the emission into lightmaps. This property only has an effect if you [bake lightmaps before runtime](../Lightmapping-baking-before-runtime.md). * **None**: Disables other GameObjects receiving the emissive light. |

## Advanced Options

The **Advanced Options** section controls the rendering calculations Unity uses.

| **Property** | **Sub-property** | **Description** |
| --- | --- | --- |
| **Flip-Book Blending** |  | Blends between the frames of a flipbook texture. Enable this property to make animations smoother for flipbook textures with limited frames. This property might reduce performance. |
| **Soft Particles** |  | Indicates whether to fade particles as they approach the edges of opaque **GameObjects**. This property is only available if you set **Surface Type** to **Transparent**, and works only if you enable **Depth Texture** in the [URP Asset](universalrp-asset.md). For more information, refer to [Soft particles](../particle-color.md#soft-particles). |
|  | **Surface Fade** | Sets when particles fade, in world units. This property is available only if you enable **Soft Particles**. The properties are:  * **Near**: Sets the distance from the opaque GameObject where the particle fades out completely. * **Far**: Sets the distance from the opaque GameObject where the particle is fully opaque. |
| ****Camera** Fading** |  | Indicates whether to fade particles as they approach the camera. This property is only available if you set **Surface Type** to **Transparent**, and works only if you enable **Depth Texture** in the [URP Asset](universalrp-asset.md). |
|  | **Distance** | Sets when particles fade, in world units. This property is available only if you enable **Camera Fading**. The properties are:  * **Near**: Sets the distance from the camera where a particle fades out completely. * **Far**: Sets the distance from the camera where a particle is fully opaque. |
| **Distortion** |  | Distorts the background behind the particles, for example to create a heat shimmer effect. This property is available only if you set **Surface Type** to **Transparent**, and works only if you enable **Depth Texture** in the [URP Asset](universalrp-asset.md). |
|  | **Strength** | Sets how much a particle distorts the background. Negative values have the opposite effect of positive values. For example, if something is offset to the right with a positive value, the equal negative value offsets it to the left. |
|  | **Blend** | Sets the visibility of the **distortion effect**. 0 means no visible distortion. 1 means only the distortion effect is visible. |
| **Vertex Streams** |  | Lists the vertex streams that the material requires in order to work properly. If the vertex streams aren’t correctly assigned, select the **Fix Now** button to apply the correct setup of vertex streams to the **particle system**. |
| **Sorting Priority** |  | Determines when Unity renders the material. Unity renders materials with lower values first. Use this property to prevent Unity from rendering pixels twice by rendering materials behind other materials. This property works similarly to the [render queue](../../ScriptReference/Material-renderQueue.md) in the Built-In Render Pipeline. |