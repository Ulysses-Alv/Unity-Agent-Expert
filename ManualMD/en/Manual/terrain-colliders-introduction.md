# Introduction to Terrain colliders

[Terrain colliders](class-TerrainCollider.md) match the shape of a [Terrain](script-Terrain.md) exactly, for extremely accurate Terrain **collision** simulation. The Terrain **collider** is the most accurate collider type for Terrains.

![A Terrain GameObject. Right: A Terrain collider, which matches the geometry of the Terrain GameObject exactly.](../uploads/Main/terrain-colliders-terrain.png)

A Terrain GameObject. Right: A Terrain collider, which matches the geometry of the Terrain GameObject exactly.

![A Terrain collider, which matches the geometry of the Terrain GameObject exactly.](../uploads/Main/terrain-colliders-terrain-collider.png)

A Terrain collider, which matches the geometry of the Terrain GameObject exactly.

## Construction of a Terrain collider

Unlike [primitive colliders](primitive-colliders.md) and [Mesh colliders](mesh-colliders.md), Terrain colliders are heightmap-based. PhysX generates Terrain colliders from the **heightmap** data of the corresponding Terrain, rather than from a pre-defined shape or a **Mesh**.

For detailed information on heightmap-based colliders in PhysX, see the Nvidia PhysX documentation [Geometry: Height Fields](https://gameworksdocs.nvidia.com/PhysX/4.1/documentation/physxguide/Manual/Geometry.html#height-Fields).

The Terrain collider builds its collision geometry to match the heightmap data of the assigned [TerrainData](../ScriptReference/TerrainData.md) asset, including its shape, position and scale. The benefit of this is that you can make the shape of the collider exactly the same as the shape of the visible Terrain, which creates more precise and realistic collisions.

## Colliders for Terrain Trees and Details

A Terrain has two **brush** types that add **Prefabs** to the Terrain: **Trees** and **Details**. These brushes allow you to quickly add trees, bushes, rocks, and other terrain details. For more detailed information, see Terrain documentation on [Brushes](class-Brush.md), [Trees](terrain-Trees-Landing.md), and [Grass and other details](terrain-Grass.md).

When you create a Prefab, you add any colliders you want to that Prefab. When you use a Tree Brush to apply that Prefab to the Terrain, you need to use the Terrain collider to enable or disable those colliders. On the Terrain collider, use **Enable Tree Colliders** to toggle Tree Prefab colliders.

![A Terrain and four trees.](../uploads/Main/terrain-colliders-terrain-trees.png)

A Terrain and four trees.

![The collision geometry of the Terrain. Note the rectangular colliders around the bottom of each tree trunk. These are present because the original tree model Prefab contains a Box collider.](../uploads/Main/terrain-colliders-terrain-trees-collider.png)

The collision geometry of the Terrain. Note the rectangular colliders around the bottom of each tree trunk. These are present because the original tree model Prefab contains a Box collider.

The Terrain collider only generates colliders for Prefabs that you add with the Tree Brush. It does not generate colliders for Prefabs that you add with the Details Brush, because the Details Brush only renders a Prefab’s Mesh, and not its colliders. If you want something on the Terrain **GameObject** to have a collider, use the Trees brush to add it.

Unity only generates Tree Prefab colliders at runtime, so they only appear in the [Physics Debug](PhysicsDebugVisualization.md) collider Geometry view at runtime.

## Optimize Terrain colliders

The Terrain collider is the correct collider choice for a Terrain in almost all cases. Terrain colliders accurately represent the shape of the Terrain, enabling precise **collision detection** between the Terrain and other colliders. Terrain colliders also require minimal setup and configuration time, because Unity automatically generates them when you create a Terrain.

If you are looking for opportunities to improve performance, you can try an alternative approach to Terrain collisions only if the Terrain has large flat areas, or if the physics collision does not need to be very accurate. The following sections describe two alternative approaches.

### Use a lower resolution Terrain collider

You can create a TerrainData with a lower resolution, and apply it to the Terrain collider.

Duplicate the Terrain’s existing TerrainData, and reduce the resolution.
Set the TerrainData field of the Terrain collider to the reduced resolution TerrainData.

Note that Unity displays a warning on the Terrain **Inspector** that the collider and terrains have different TerrainDatas, but you can ignore this as long as you’ve done it intentionally.

### Create a Mesh collider from the Terrain:

You can export the Terrain’s heightmap and use external software to turn it into a simplified [Mesh collider](mesh-colliders.md).

Install the [Terrain Tools](https://docs.unity3d.com/Packages/com.unity.terrain-tools@latest) package.
Use the [Export Heightmap](https://docs.unity3d.com/Packages/com.unity.terrain-tools@latest?subfolder=/manual/toolbox-export-heightmaps.html) tool to export the Terrain heightmap.
Use 3D modeling software to create a mesh from the heightmap.
Use 3D modeling software to run a mesh simplification routine on the mesh.
Import the mesh back into Unity to use as the source for a mesh collider.