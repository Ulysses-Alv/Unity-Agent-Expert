# Create trees and leaves from meshes

You can base a Tree Editor tree or its leaves on meshes that you have imported into Unity.

Working with an imported **mesh** isn’t the same as working with SpeedTree. For more information about importing from SpeedTree, refer to [Import trees from SpeedTree](SpeedTree-landing.md).

## Import a mesh

To learn about exporting models from a 3D modeling application and importing them into Unity, refer to [Models](models.md).

The mesh you want to use needs to be in the `Ambient-Occlusion` folder the trees are in. For more information, refer to [Shaders and the Ambient-Occlusion folder](terrain-Trees-Mat-Shaders.md).

## Use a mesh for the tree

A Tree Editor tree has a **Mesh Filter** component that can reference your imported mesh.

To assign your imported mesh to the tree:

1. In the **Inspector** window > **Tree Hierarchy** view, select either the root node or one of the branch groups.
2. In the **Mesh (Mesh Filter)** component, select the mesh picker (circle).
3. In the **Select** Mesh window, select the mesh you want to use.

When you use an imported mesh, you can’t add leaves or branches to the tree.

## Use a mesh for the leaves

You can use the leaf group’s mesh geometry mode to use an imported mesh. It can be flowers, fruit, or any other object you want to attach to the tree.

Note that the mesh’s transform must be `0,0,0`. Any other transform is added as a distance from the tree, so the object floats by the tree instead of being attached to the branches. Hide the original mesh from the **scene** to avoid seeing it at the `0,0,0` location.

## Additional resources

* [Tree level of detail (LOD)](terrain-Tree-LOD.md)
* [Shaders and the Ambient-Occlusion folder](terrain-Trees-Mat-Shaders.md)
* [Tree Hierarchy view UI reference](terrain-Tree-Hierarchy-UI.md)
* [Root node reference](tree-Root-Node.md)
* [Branch group reference](tree-Branches.md)
* [Leaf group reference](tree-Leaves.md)
* [Terrain Settings reference](terrain-OtherSettings.md)