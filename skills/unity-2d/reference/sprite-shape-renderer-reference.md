# Sprite shape renderer reference

| Property | Function |
| --- | --- |
| **Color** | Define the vertex color of the **Sprite** Shape geometry, which tints or recolors the Sprite Shape. Use the color picker to set the vertex. See the [Color](color) section below this table for examples. |
| **Mask Interaction** | Set how the Sprite Renderer behaves when it interacts with a [Sprite Mask](../mask/mask-landing.md). Refer to examples of the different options in the [Mask Interaction](mask-interaction) section below.  * **None**: The Sprite Shape Renderer does not interact with any Sprite Masks in the Scene. This is the default option. * **Visible Inside Mask**: The Sprite Shape is visible where the Sprite Mask overlays it, but not outside of it. * **Visible Outside Mask**: The Sprite Shape is visible outside of the Sprite Mask, but not inside it. The Sprite Mask hides the sections of the Sprite Shape it overlays. |
| **Sorting Layer** | Set the [Sorting Layer](../../class-TagManager.md) of the Sprite Shape geometry, which controls its priority during rendering. Select an existing Sorting Layer from the drop-down box, or create a new Sorting Layer. |
| **Order In Layer** | Set the render priority of the Sprite Shape within its [Sorting Layer](../../class-TagManager.md). Lower numbered Sprite Shapes are rendered first, with higher numbered Sprite Shapes overlapping those below. |

SpriteShapeRenderer