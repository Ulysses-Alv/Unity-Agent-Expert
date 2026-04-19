# Shadow Caster 2D component reference

Explore the properties you can use to add and configure shadows from 2D lights in the Universal **Render Pipeline** (URP).

For more information, refer to [Add shadows from 2D lights in URP](2DShadows.md).

| **Property** | **Description** |
| --- | --- |
| **Casting Source** | Sets the source for the shape of the shadows. Unity displays the shape as a white outline in the Scene view. The options are the following:  * **Sprite Skin**: Uses the bounding box of the bones defined in the **Sprite Skin** component attached to the GameObject. This option is available only if you add a Sprite Skin component to the GameObject. * **Sprite Renderer**: Uses the shape defined in the [Sprite Editor](../sprite/sprite-editor/sprite-editor-landing) window. Uses the shape in the **Skinning Editor** tab if the GameObject has a **Sprite Skin** component, otherwise uses the shape in the **Custom Outline** tab. * **Collider 2D**: Uses the shape defined in the collider component attached to the GameObject. This option is available only if you add a collider component to the GameObject, for example a Box Collider 2D component. * **Shape Editor**: Uses the shape defined in the **Edit Shape** property. * **None**: Casts no shadows. |
| **Casting Option** | Sets how the GameObject casts shadows. The options are the following:  * **Self Shadow**: Casts shadows only onto itself. * **Cast Shadow**: Casts shadows only onto other objects. * **Cast and Self Shadow**: Casts shadows onto itself and other objects. * **None**: Casts no shadows.  The default is **Cast Shadow**. |
| **Target Sorting Layers** | Sets the [2D sorting layers](../2d-renderer-sorting.md#sortlayer) that receive shadows from this **GameObject**. The default is **Everything**, which means other GameObjects on any sorting layer receive shadows. |
| **Trim Edge** | Increases or decreases the size of the **Casting Source**. A **Trim Edge** value of 1 means the **Casting Source** is its original size. Lower values decrease the size. This property is available only if you set **Casting Source** to ****Sprite** Renderer**. The default is 0.2. |
| **Alpha Cutoff** | Sets the minimum alpha value a **pixel** must have to cast a shadow. The default is 0.1. This property is available only if you set **Casting Source** to ****Sprite Renderer****. |
| **Edit Shape** | If **Casting Source** is set to **Shape Editor**, select this button to edit the shape of the shadow. Unity displays the shape as a white outline in the **Scene** view. Click and drag the white circles to move the points, or click a line to add a point. |
| **Position** | Sets the position of a point on the shape. This property is only available if you select **Edit Shape**. |

## Additional resources

* [Add shadows from 2D lights in URP](2DShadows.md)