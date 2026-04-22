# 9-slicing

9-slicing is a 2D technique which allows you to reuse an image at various sizes without needing to prepare multiple [Assets](../../AssetWorkflow.md). It involves splitting the image into nine portions, so that when you re-size the [Sprite](../sprite-landing.md), the different portions scale or tile (that is, repeat in a grid formation) in different ways to keep the Sprite in proportion. This is useful when creating patterns or [Textures](../../class-TextureImporter.md), such as walls or floors in a 2D environment.

This is an example of a 9-sliced Sprite, split into nine sections. Each section is labelled with a letter from A to I.

![Example of a 9-sliced sprite.](../../../uploads/Main/9-slice-0.png)

Example of a 9-sliced sprite.

The following points describe what happens when you change the dimensions of the image:

* The four corners (A, C, G and I) do not change in size.
* The B and H sections stretch or tile horizontally.
* The D and F sections stretch or tile vertically.
* The E section stretches or tiles both horizontally and vertically.

This section describes how to set 9-slicing up, and which settings to apply depending on whether you want to stretch or tile the areas shown above.

### Limitations and known issues

* The only two Collider2Ds that support 9-slicing are BoxCollider2D and PolygonCollider2D.
* You cannot edit BoxCollider2D or PolygonCollider2D when the **Sprite Renderer**’s **Draw Mode** is set to **Sliced** or **Tiled**. Editing in the **Inspector** window is disabled, and a warning appears to notify you that the Collider2D cannot be edited because it is driven by the Sprite Renderer component’s tiling properties.
* When the shape is regenerated in **Auto Tiling**, additional edges might appear within the Collider2D’s shape. This may have an effect on **collisions**.