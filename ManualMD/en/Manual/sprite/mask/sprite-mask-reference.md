# Sprite Mask component reference

Explore the properties you use to hide and reveal parts of sprites with a **sprite** mask. For more information about using **sprite masks**, refer to [Masking sprites](../mask/mask-landing.md).

| **Property** | **Description** |
| --- | --- |
| **Mask Source** | Selects what to use as the source of the mask texture. In the mask texture, use opaque pixels for the mask shape, and a transparent background for outside the mask. The options are:  * **Sprite**: Uses the texture from a sprite. * **Supported Renderer**: Uses the texture from the sprite on a renderer component attached to the GameObject. |
| **Sprite** | Sets the sprite to use as the source of the mask texture. This property is available only when you set **Mask Source** to **Sprite**. |
| **Supported Renderer** | Sets the renderer to use as the source of the mask texture. This option is available only when you set **Mask Source** to **Supported Renderer**, and the GameObject has one of the following renderer components attached:  * Sprite Renderer * Sprite Shape Renderer * Tilemap Renderer |
| **Sprite Sort Point** | Selects the point Unity uses as the distance of the mask from the camera, for example in sorting calculations. This option is available only when you set **Mask Source** to **Sprite**. The options are:  * **Center**: Uses the center of the sprite. * **Pivot**: Uses the **Pivot** property of the sprite. For more information about setting the pivot point, refer to [Sprite Editor](../sprite-editor/open-sprite-editor). |
| **Alpha Cutoff** | Sets the lowest alpha value Unity uses as a mask **pixel**. Lower values mean Unity includes more transparent pixels in the mask shape. |
| **Custom Range** | Sets the range of sorting layers that the mask affects. For more information, refer to the [Custom Range properties](#custom-range-properties) section. |
| **Rendering **Layer Mask**** | Sets which rendering layers the mask affects. For more information, refer to [Rendering layers](../../urp/features/rendering-layers.md). |

## Custom Range properties

Sets the range of layers that the mask affects. These properties are only available when you enable **Custom Range**.

For example, if you set the **Back** **Sorting Layer** property to sorting layer 1 and the **Front** **Sorting Layer** property to sorting layer 2, the mask affects only sorting layer 2.

| **Property** | **Sub-Property** | **Description** |
| --- | --- | --- |
| **Front** | Sets the highest layer and sublayer the mask affects. |
| N/A | **Sorting Layer** | Sets the highest layer the mask affects, so that layers in front of this layer aren’t affected by the mask. |
| N/A | **Order in Layer** | Sets the highest sublayer in the **Sorting Layer** that the mask affects. |
| **Back** | Sets the lowest layer and sublayer not affected by the mask. |
| N/A | **Sorting Layer** | Sets the lowest layer not affected by the mask, so that both this layer and the layers behind aren’t affected by the mask. |
| N/A | **Order in Layer** | Sets the lowest sublayer not affected by the mask. |

## Additional resources

* [Sprite Editor](../sprite-editor/open-sprite-editor)
* [2D renderer sorting](../../2d-renderer-sorting.md)