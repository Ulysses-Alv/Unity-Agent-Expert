# Tree level of detail (LOD)

Trees from Tree Editor and trees from SpeedTree handle their level of detail (LOD) differently.

## SpeedTree

A SpeedTree **Prefab** has an LODGroup component. Refer to [LOD](LevelOfDetail.md) and [LOD Group](class-LODGroup.md) for more information about configuring LOD components.

## Tree Editor

Tree Editor trees don’t have an LODGroup component. Unity’s LOD system uses a 2D to 3D transition zone to blend 2D **billboards** with 3D tree models. This prevents a sudden popping of 2D and 3D trees, which is important for realism in **VR**.

**Note**: Billboard trees don’t receive local lighting such as Point Lights and Spot Lights. They work with directional lights, but lighting on the billboards updates only when you rotate the **camera**.

Tree Editor controls LOD at two levels:

* The tree root node.
* Individual branch groups, as a multiplier of the tree **root node**’s LOD.

The leaves and bark **shaders** also have LOD defaults. These are built-in shaders, and you can’t edit them.

Each **Terrain** tile has settings for tree drawing, such as the distance from the camera where trees switch to billboard mode. These settings can impact the gaming experience when they create transitions that are visible to the player. Refer to the [Terrain Settings reference](terrain-OtherSettings.md) for more information.

## Additional resources

* [Create trees with Tree Editor](class-Tree.md)
* [Terrain Settings reference](terrain-OtherSettings.md)
* [Import trees from SpeedTree](SpeedTree-landing.md)