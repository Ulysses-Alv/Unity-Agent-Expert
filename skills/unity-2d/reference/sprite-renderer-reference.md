# Sprite Renderer component reference

Adjust the color of a **sprite**, orientation of the sprite texture, and other properties to affect how a sprite is rendered.

| Property | Description |
| --- | --- |
| **Sprite** | Set the texture the sprite **GameObject** renders in the **scene**. To change the set sprite, select the **Sprite** picker (**⊙**) to open a window that displays available assets. |
| **Open Sprite Editor** | Select this button to open the currently set sprite texture in the **Sprite Editor** window. |
| **Color** | Select a color for the sprite, which tints the sprite’s texture with the selected color. |
| **Flip** | Activate an axis to flip the sprite texture along that axis. This won’t move or mirror the GameObject’s location. |
| **Draw Mode** | Select how Unity renders and scales the sprite by selecting from the following options:  * **Simple**: This is the default mode. This scales the entire sprite uniformly. * **Sliced**: Select this mode if the sprite is [9-sliced](../9-slice/9-slice-landing.md), and have it scale according to 9-slice scaling behavior. Use the **Size** property to enter the sprite’s new **Width** and **Height** dimensions in Unity units. * **Tiled**: Select this mode if the sprite is 9-sliced. This mode causes the middle of the 9-Sliced sprite to tile instead of scale when its dimensions change. Refer to this mode [options](#draw-mode-tiled) to adjust how this mode renders the sprite. |
| **Mask Interaction** | Set how the Sprite Renderer behaves when interacting with a [Sprite Mask](../mask/mask-landing.md).  * **None:** Renders without interacting with Sprite masks. * **Visible Inside Mask:** Renders only the parts of the mesh that overlap the Sprite mask. * **Visible Outside Mask:** Renders only the parts of the mesh that don’t overlap the Sprite mask. |
| **Sprite Sort Point** | Choose which point on your sprite Unity should use to calculate its distance from the camera. This determines the sprite’s [render order during sorting](../../2d-renderer-sorting.md#sortpoint). Refer to [Sprite Sort Point](set-sort-point-sprite.md) for more details.  * **Center:** Unity uses the center of the sprite to calculate the distance between it and the camera. * **Pivot:** Unity uses the set Pivot position on the sprite to calculate the distance between it and the camera. You can edit the sprite’s Pivot position in the Sprite Editor. For more information, refer to [Sprite Editor tab reference for the Sprite Editor window](../sprite-editor/sprite-editor-window-reference.md). |
| **Material** | Select the material for the sprite. The default material is `Sprite-Lit-Default`. Select the **Material** picker (⊙) to open a window that displays available materials. |

### Draw Mode - Tiled properties

The following properties are available when you set the **Draw Mode** to **Tiled**.

| Properties | Description |
| --- | --- |
| **Size** | Enter the sprite’s new **Width** and **Height** dimensions in **Unity units**. |
| **Tile Mode** | Select the tiling behavior of the sprite:  * **Continuous**: This is the default behavior. Select this to have the tiles tile evenly when the sprite dimensions change. * **Adaptive**: Select this mode to have the sprite texture stretch when its dimensions change, similar to Simple mode. When the scale of the changed dimensions meets the Stretch Value, the middle slice of the sprite begins to tile. Use the **Stretch Value** slider to set the value between 0 and 1. The max value is 1, which is double the original sprite’s scale. |

## Additional Settings

| Property | Description |
| --- | --- |
| **Sorting Layer** | Set the [Sorting Layer](../../class-TagManager.md) of the sprite, which controls its priority during rendering. Select an existing Sorting Layer from the dropdown box, or create a new Sorting Layer.  * **Default**: The default layer the sprite is on. * **Add Sorting Layer**: Select this to create a new sorting layer for the selected sprite. |
| **Order in Layer** | Set the render priority of the sprite within its **Sorting Layer**. Lower-numbered sprites are rendered first, with higher-numbered sprites overlapping those below. |
| **Rendering **Layer Mask**** | Selects which rendering layers this GameObject belongs to. |

## Additional resources

* [Sprite Mask Interaction](../mask/hide-reveal-parts-sprite-mask.md)
* [Set Sort Point for a Sprite](set-sort-point-sprite.md)
* [9-slicing](../9-slice/9-slicing.md)