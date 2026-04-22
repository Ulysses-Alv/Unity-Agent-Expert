# Applying 2D billboards for low LOD meshes

Billboards are a level-of-detail (LOD) method for drawing complicated 3D Meshes in a simpler way when they are far away from the **Camera**. When a **Mesh** is far away from the Camera, its size on screen means there is no need to draw it in full detail. Instead, you can replace the complex 3D Mesh with a 2D **billboard** representation.

The [Billboard Renderer](class-BillboardRenderer.md) renders [Billboard assets](class-BillboardAsset.md).

A Billboard asset is a collection of pre-rendered images of a mesh. Use it with the [Billboard Renderer](class-BillboardRenderer.md) to an object that is distant from the Camera at a low [level of detail (LOD)](LevelOfDetail.md).

The most common way to generate a Billboard Asset is to create files in [SpeedTree Modeler](https://unity.com/products/speedtree), and then import them into Unity.

It is also possible to create your own Billboard Assets from script. For more information, see the API reference for [BillboardAsset](../ScriptReference/BillboardAsset.md).

Certain features, such as SpeedTree, export Billboard Assets, but you can also create them yourself. For information on how to create a Billboard Asset, see the [BillboardAssets](class-BillboardAsset.md) Manual page and the [BillboardAsset](../ScriptReference/BillboardAsset.md) Script reference page.