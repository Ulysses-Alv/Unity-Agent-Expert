# Performance tips for trees

To improve the performance of trees:

* Performance depends on the polygon count of your tree model. Test your trees on your platform, and create simpler trees if performance is low.
* Don’t create too many leaves and branches, because this can reduce performance.
* Each **Terrain** tile has settings for tree drawing, such as the distance from the **camera** where trees switch to **billboard** mode. These settings can impact the gaming experience when they create transitions that are visible to the player. Refer to the [Terrain Settings reference](terrain-OtherSettings.md) for more information.

## Additional resources

* [Tree level of detail (LOD)](terrain-Tree-LOD.md)
* [Shaders and the Ambient-Occlusion folder](terrain-Trees-Mat-Shaders.md)
* [Tree Hierarchy view UI reference](terrain-Tree-Hierarchy-UI.md)
* [Root node reference](tree-Root-Node.md)
* [Branch group reference](tree-Branches.md)
* [Leaf group reference](tree-Leaves.md)
* [Terrain Settings reference](terrain-OtherSettings.md)