# Add trees to the terrain

Use the Paint Trees tool to place trees on the **terrain** and customize their painting. You can paint trees with a brush or mass place trees across the entire terrain tile.

As an example, use the [free Unity Terrain - HDRP Demo Scene](https://assetstore.unity.com/packages/3d/environments/unity-terrain-hdrp-demo-scene-213198), which has six SpeedTree models. You can also create trees with the [Tree Editor](class-Tree.md) (for the Built-In Render Pipeline).

## The Paint Trees tool

To access tree painting:

1. In the **Hierarchy** window, select the terrain.
2. In the terrain’s ****Inspector**** window, from the **Terrain** **toolbar**, select **Paint Trees**.

### Select trees

To add trees to the terrain, add a tree prototype to the Paint Tree tool:

1. Select **Edit Trees** > **Add Tree**.
2. Select a tree.

The **Add Tree** window has different options based on the tree you’re adding:

* **Bend Factor**: Adjust the tree’s responsiveness to the wind in the **scene**. Trees from the SpeedTree Modeler don’t have a **Bend Factor**; only Tree Editor trees do. Refer to [Animate trees with Wind Zones](class-WindZone.md) for more information.
* ****NavMesh** **LOD** Index**: SpeedTree uses level of detail (LOD) groups. The Unity Editor manages the transitions between groups, but if you’re placing trees that are never going to be near the **camera**, or if you need to limit the display quality of trees without editing the model, you can select a specific **LOD group**. The options are:
  + **First**: The highest LOD, for trees that are close to the player.
  + **Last**: The lowest LOD, for trees that are far away from the player.
  + **Custom**: To select a different level.

  Refer to [Tree level of detail (LOD)](terrain-Tree-LOD.md) for more information.
* **Cast Shadows**: For Tree Editor trees. This is not editable; it lists the option selected in the tree’s **Mesh Renderer** > **Lighting** > **Cast Shadows** setting.

### Customize trees for painting

Customize tree placement and characteristics. Except for **Brush Size** and **Tree Density**, the settings apply to both mass placed trees and brush painted trees.

| **Property** | **Function** |
| --- | --- |
| **Brush Size** | The size of the area each brush stroke covers with trees. |
| **Tree Density** | How far apart trees must be. This limits the number of trees a single brush stroke adds. Note that the limit is by brush stroke, not area, so repeated brush strokes in the same area create higher densities. Note also that the larger the brush, the farther apart the trees, even if the density is 100. |
| **Tree Height** | A range for tree height randomization. Use a wide range for a varied look, and a narrow range for a uniform look. To specify a value instead of a range, disable **Random**. The possible values are a scale of `0.01` to `2` of the tree’s original height. |
| **Lock Width to Height** | By default, a tree’s width is locked to its height so that trees are always scaled uniformly. To specify a width, disable **Lock Width to Height**. |
| **Tree Width** | A range for tree width randomization. Use a wide range for a varied look, and a narrow range for a uniform look. To specify a value instead of a range, disable **Random**. The possible values are a scale of `0.01` to `2` of the tree’s original width. |
| **Random Tree Rotation** | If your tree has a level of detail (LOD) Group (for example, a tree imported from SpeedTree), use the **Random Tree Rotation** setting to help create the impression of a random, natural-looking forest. Disable this option if you want to place trees with a fixed rotation. |
| **Color Variation** | The amount of random shading applied to trees. This works only if the **shader** reads the `_TreeInstanceColor` property. All trees created with Tree Editor read the `_TreeInstanceColor` property, because they all use built-in shaders. Some SpeedTree trees might not read it. |
| **Tree Contribute Global Illumination** | Non-editable; indicates whether the tree asset’s **Lighting** > **Contribute Global Illumination** option is enabled. |

### Mass place trees

You can place trees across the entire terrain tile in a single action. Refer to Customize trees for painting (above) for information about customizing trees before painting.

To place trees on an entire terrain tile:

1. Select tree assets as explained above. Note that mass placing uses all the trees in the **Trees** list, not just the highlighted tree.
2. Select **Mass Place Trees**. The **Place Trees** window opens.
3. Enter the number of trees you want to place.
4. To keep the trees you already have, check the **Keep Existing Trees** box. If you don’t check this box, Unity removes all existing trees and replaces them with the new trees.

You can use the brush tools to add or remove trees that you mass placed.

### Paint trees with a brush

You can paint trees onto the terrain with a brush.

To paint trees:

1. Select trees as explained above.
2. Refer to Customize trees for painting (above) for information about the brush settings and customizing trees for painting.
3. Paint trees on the terrain using the brush.
4. To remove:
   * All trees from an area: hold **Shift** as you paint.
   * Only the tree selected in the **Trees** list: hold **Ctrl** as you paint.

## Edit trees

You can edit the tree’s properties before you place it, or between painting strokes. For example, for SpeedTree trees, you can change the LOD group you use when painting different parts of the terrain.

To edit a tree:

1. In the **Trees** list, select the tree you want to edit.
2. Select **Edit Trees** > **Edit Tree**. The **Edit Tree** window opens. It has the same options as the **Add Tree** window.

## Update imported trees

If you change an imported tree in a separate 3D modeling application, you need to refresh the tree in the Paint Tool to update its appearance on the terrain.

To update the trees, in the Terrain tile’s **Inspector** window, go to **Paint Trees** > **Trees** > **Refresh**.

**Note:** Trees from SpeedTree use their own shader to animate wind movement. The Unity Editor adds this shader only to trees that SpeedTree exports with the file type `spm` or `st8`. If you edit the trees using a different 3D modeling software and export them as `fbx` or `obj` files, the Unity Editor assigns them built-in shaders, and the wind effect is lost. To preserve the wind animation, always use SpeedTree to edit and re-export SpeedTree trees.

## Additional resources

* [Create trees with Tree Editor](class-Tree.md)
* [Add collision to trees](terrain-Tree-Colliders.md)
* [Tree level of detail (LOD)](terrain-Tree-LOD.md)
* [Animate trees with Wind Zones](class-WindZone.md)
* [Terrain Settings reference](terrain-OtherSettings.md)