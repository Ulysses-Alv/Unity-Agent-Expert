# Add a sprite mask

To hide parts of sprites, add a **sprite** mask to the **scene**. You place a **sprite mask** at a position in the scene, and the mask shape hides the sprites that overlap with it.

![A circular mask, with a crate sprite set to Visible Inside Mask. The parts of the sprite outside the mask are hidden.](../../../uploads/Main/2D_SpriteRenderer_8.png)
![The same circular mask, with a crate sprite set to Visible Outside Mask. The parts of the sprite that overlap the mask are hidden.](../../../uploads/Main/2D_SpriteRenderer_9.png)

Sprite masks only affect objects that have a [Sprite Renderer](../renderer/renderer-landing) component.

**Note:** Sprite masks are different to the mask map texture you [add as a secondary texture](../../urp/SecondaryTextures.md) to control which areas of a sprite receive light.

## Limitations

Sprite masks aren’t compatible with the following:

* [Scriptable Render Pipeline (SRP) Batcher](../../SRPBatcher.md). If you use a sprite mask, Unity falls back to using [dynamic batching](../../DrawCallBatching.md).
* GPU deformation of sprites in 2D animation. If you use a sprite mask, Unity falls back to using CPU deformation.

## Add a sprite mask

Follow these steps:

1. In your [2D renderer asset](../../urp/2DRendererData-overview.md), make sure **Depth/**Stencil Buffer**** is enabled.
2. To add a sprite mask to the scene, select **GameObject** > **2D Object** > **Sprite Mask**.

   The default mask shape is a circle.
3. In the **Hierarchy** window, select the sprite you want to mask.
4. In the ****Inspector**** window of the sprite, set **Mask Interaction** to **Visible Inside Mask** or **Visible Outside Mask**.
5. Make sure the sprite overlaps with the sprite mask.

## Change the shape of a mask

To change the shape of a mask, follow these steps:

1. Select the sprite mask **GameObject**.
2. In the **Inspector** window, select the **Sprite** picker (**⊙**).
3. Select a sprite from the list to use as a mask.

To use a custom shape, create a new sprite with the shape you want to use. Use opaque **pixels** for the mask shape, and a transparent background for outside the mask.

## Set which sprites a mask affects

To change which sprites a mask affects, assign the sprites to [sorting layers](../../2d-renderer-sorting.md), then restrict the sprite mask to a range of sorting layers using its **Custom Range** properties.

For example:

1. Open the [Tags and Layers window](../../class-TagManager.md) and create three sorting layers: **Back**, **Middle**, and **Front**.
2. Create three sprites, and set their **Sorting Layer** properties to the different layers.
3. Select your sprite mask, and in the **Inspector** window enable **Custom Range**.
4. Set the front **Sorting Layer** to **Front**, and the back **Sorting Layer** to **Back**. The sprite mask now only affects the sprites in the **Front** and **Middle** sorting layers.

You can also use [Sorting Group components](../sorting-group/sorting-group-landing.md) to prevent multiple masks from interacting with each other.

![Left: Two sprite masks affect two card sprites, so the sprites overlap incorrectly. Right: Each sprite and its mask use their own sorting group, and the sprites overlap correctly.](../../../uploads/Main/SpriteMask7.jpg)

Left: Two sprite masks affect two card sprites, so the sprites overlap incorrectly. Right: Each sprite and its mask use their own sorting group, and the sprites overlap correctly.

## Additional resources

* [Sprite Mask component reference](sprite-mask-reference.md)