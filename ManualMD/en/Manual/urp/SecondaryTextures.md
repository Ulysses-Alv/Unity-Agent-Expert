# Add a normal map or a mask map to a sprite in URP

To add surface detail such as bumps, grooves, and scratches to a 2D **sprite** in the Universal **Render Pipeline** (URP), add the following texture maps:

* A **normal map** to make details catch the light from a [2D light](2d-index.md) as if they represent real geometry. For more information, refer to [Introduction to normal maps](../StandardShaderMaterialParameterNormalMap.md).
* A mask map to control which areas of the sprite receive light, to create highlights, rim lighting, and shadowing.

In 2D projects, normal maps and mask maps are called secondary textures of the sprite. You add them using the **Secondary Textures** overlay of the **Sprite Editor** window.

![Left: A rock with Normal Map set to disabled. The light is bright and flat. Right: The same rock with Normal Map set to enabled. The light highlights cracks and crevices.](../../uploads/urp/2D/image_11.png)

Left: A rock with Normal Map set to disabled. The light is bright and flat. Right: The same rock with Normal Map set to enabled. The light highlights cracks and crevices.

![Left: A mask map texture for a smooth rock. Center: The sprite without the mask map. The rock receives light uniformly. Right: The sprite with the mask map. The rock receives less light where the mask map is less red.](../../uploads/urp/2D/maskmap.png)

Left: A mask map texture for a smooth rock. Center: The sprite without the mask map. The rock receives light uniformly. Right: The sprite with the mask map. The rock receives less light where the mask map is less red.

**Note:** The mask map texture you add as a secondary texture is different to the [sprite mask](../sprite/mask/mask-landing.md) you use to hide or reveal parts of a sprite.

## Import a texture map

To import a texture map, follow these steps:

1. Drag the texture map into the **Project** window to import it into your project.
2. In the ****Inspector**** window, set the Texture Type to **Normal Map** for a normal map, or **Default** for a mask map.
3. Disable **sRGB (Color Texture)**.

To make sure the lighting works correctly, make sure your texture map lines up with the sprite texture. Unity samples texture maps using the same uv coordinates it uses to sample the sprite texture.

![A sprite texture of a rock, and its normal map with the same position and alignment.](../../uploads/urp/2D/ST_Align.png)

A sprite texture of a rock, and its normal map with the same position and alignment.

## Add the texture map to a sprite

Add the normal map or mask map as a **Secondary Texture** to the sprite. Follow these steps:

1. Make sure the **Sprite Renderer** component of the sprite uses a material that supports texture maps, for example the default `Sprite-Lit-Default` material. For more information, refer to [Prepare and upgrade sprites for 2D lighting in URP](PrepShader.md).
2. Select the original texture of the sprite in the **Project** window, then select **Open Sprite Editor**.
3. Select the **Secondary Textures** tab from the dropdown.

   Unity displays the **Secondary Textures** overlay in the bottom-right corner of the **Sprite Editor** window.
4. In the **Secondary Textures** overlay, select the **Add** (**+**) button to add a new secondary texture.
5. Select the **Name** dropdown, then select **\_NormalMap** for a normal map, or **\_MaskTex** for a mask map.

   **Note:** The dropdown might also include names used by other Unity packages you installed, even if you have since uninstalled the packages.
6. To select your normal map or mask map, drag it from the **Project** window to the **Texture** field, or select the picker (**⊙**).
7. Select **Apply**.

To display the secondary texture in the **Sprite Editor** window, select it in the **Secondary Textures** overlay. To display the sprite texture again, select anywhere outside of the overlay.

## Configure a 2D light to use the texture map

To check the effect of the texture map on the sprite texture under a 2D light, follow these steps:

1. [Add a 2D light](LightTypes.md) to the **scene**.
2. Select the light in the **Hierarchy** window to open its **Inspector** window.
3. To enable the normal map, in the **Normal Maps** section, set **Use Normal Map** to **Fast** or **Accurate**.
4. To enable the mask map, in the **Blending** section, set **Blend Style** to **Multiply with Mask (R)** or **Additive with Mask (R)**.

## Additional resources

* [Secondary textures - Lit sprites and 2D VFX tutorial](https://www.youtube.com/watch?v=InNZsUWNb8k) on the Unity YouTube channel
* [2D lighting for pixel art](https://learn.unity.com/course/2d-lighting-for-pixel-art) on the Unity Learn website