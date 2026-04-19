# How the Mesh LOD generator works

Mesh **LOD** provides a LOD generator that runs at model import time.

The LOD generator analyzes the **mesh** and determines the edges to collapse and the triangles to keep, while preserving the original shape and avoiding visual artifacts. After selecting the triangles to keep for a LOD, the generator stores indices of the triangles in the index buffer of the original mesh.

The original mesh must have at least 256 triangles for the simplification process to start.

In the first iteration, the generator runs on the original mesh. In each subsequent iteration, it runs on the LOD mesh from the previous iteration, and appends a set of indices to the index buffer. The generator doesn’t add any new vertices. All LODs reuse the same vertex buffer.

![LOD generation iterative process](../../uploads/Main/lod/mesh-lod-generation-process.png)

LOD generation iterative process

Each subsequent LOD contains approximately half the number of indices as the previous one.

The generation process stops when one of the following conditions is met:

* The generator reaches a LOD index specified in the [**Limit LODs**](../FBXImporter-Model.md#meshlodprops) setting.
* There are not enough triangles to continue mesh simplification. In most cases the minimum number is around 64.
* The generator can’t omit any more vertices.

The part of the index buffer that a specific LOD uses is defined by the [MeshLodRange](../../ScriptReference/UnityEngine.MeshLodRange.md) struct. The index range is defined relative to the sub-mesh range.

The LOD generator works best for organic shapes or relatively flat surfaces with a high vertex density, for example: living creatures, rocks, sculptures, and so on. The generator might produce undesired results on meshes with a high curvature or meshes that consist of many disconnected pieces.

### Mesh LOD memory footprint

The Mesh LOD feature stores all data about LODs in the index buffer of the original mesh. Each subsequent LOD contains approximately half the number of indices as the previous one. The maximum increase in size of the original index buffer is 100%. The increase is smaller if you use the [**Maximum Levels**](../FBXImporter-Model.md#meshlodprops) or the [**Discard Odd Levels**](../FBXImporter-Model.md#meshlodprops) property.

The index buffer of a mesh is relatively small compared to the total size of the mesh. This makes the memory footprint of the Mesh LOD feature smaller than the footprint of the **LOD Group** feature. With LOD Group, Unity has to store a complete mesh with both the vertex and the index buffer for each LOD as well as extra **GameObjects** and components.

## Additional resources

* [Introduction to level of detail](../LevelOfDetail.md)