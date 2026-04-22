# Mesh collider component reference

[Switch to Scripting](../ScriptReference/MeshCollider.md "Go to MeshCollider page in the Scripting Reference")

The **Mesh** **Collider** component takes a [mesh asset](class-Mesh.md) and builds a collider that matches the geometry of that mesh.

Mesh colliders are a convenient way to represent complex shapes and can provide more accurate **collision detection** for complex geometry than primitives. However, for moving or dynamic objects, primitive colliders can be more efficient because they offer more stability and better-performing physics simulation.

Mesh colliders that are marked as **Convex** can collide with other mesh colliders.

| **Property** | **Description** |
| --- | --- |
| **Convex** | Make the mesh collider collide with other mesh colliders. **Convex** mesh colliders are limited to 255 triangles. |
| **Is Trigger** | Use the collider as a trigger for events. When **Is Trigger** is enabled, other colliders pass through this collider, and trigger the messages [`OnTriggerEnter`](../ScriptReference/Collider.OnTriggerEnter.md), [`OnTriggerStay`](../ScriptReference/Collider.OnTriggerStay.md), and [`OnTriggerExit`](../ScriptReference/Collider.OnTriggerExit.md). |
| **Provides Contacts** | Generate contact information for this collider at all times. Usually, a collider only generates contact data if there is something to send it to; in this case, the messages [`OnCollisionEnter`](../ScriptReference/Collider.OnCollisionEnter.md), [`OnCollisionStay`](../ScriptReference/Collider.OnCollisionStay.md), or [`OnCollisionExit`](../ScriptReference/Collider.OnCollisionExit.md). When **Provides Contacts** is enabled, the collider generates contact data for the physics system at all times. Contact generation is resource-intensive, so **Provides Contacts** is disabled by default. |
| **Cooking Options** | Enable or disable the mesh cooking options that affect how the physics system processes meshes. When you set the **Cooking Options** to any other value than the default settings (that is, everything enabled except **None**), the mesh collider must use a mesh that has an [`isReadable`](../ScriptReference/Mesh-isReadable.md) value of `true`. For details on mesh cooking, refer to [Prepare a mesh for mesh colliders](prepare-mesh-for-mesh-collider.md). Select from the following options:  * **None**: Disables all **Cooking Options**. This is disabled by default. * **Everything**: Enables all **Cooking Options**. This is enabled by default. * **Cook for Faster Simulation**: The cooking process preprocesses the mesh data and stores it in memory to speed up runtime calculations. This is useful for complex meshes in the scene. When this setting is disabled, the physics system uses a faster cooking time, but retrieves the mesh data more slowly at runtime. This is enabled by default. * **Enable Mesh Cleaning**: The cooking process attempts to clear the mesh of degenerate triangles on the mesh (that is, triangles in which all three points are collinear, and do not form a triangle shape) and other geometrical artifacts. This results in a mesh that is better suited for use in collision detection, and tends to produce more accurate contact points. When this setting is disabled, the physics system uses a faster cooking time but implements less optimization. This is enabled by default. * **Weld Colocated Vertices**: The cooking process combines vertices that have the same position. This results in a mesh that is better suited for use in collision detection, and tends to produce more accurate contact points. When this setting is disabled, the physics system uses a faster cooking time but implements less optimization. This is enabled by default. * **Use Fast Midphase**: The cooking process uses the fastest mid-phase acceleration structure and algorithm available for your output platform. The fastest algorithm doesn’t require any R-Trees for spatial access. If you encounter midphase issues at runtime, disable this option. Instead, Unity then uses the slower legacy midphase algorithm. This is enabled by default. |
| **Material** | Reference to the [Physics Material](class-PhysicsMaterial.md) that determines how this collider interacts with others. |
| **Mesh** | Reference to the mesh to use for collisions. |

## Layer overrides

The Layer Overrides section provides properties that allow you to override the project-wide [Layer-based collision detection](LayerBasedCollision.md) settings for this collider.

| **Property** | **Description** |
| --- | --- |
| **Layer Override Priority** | Define the priority of this collider override. When two colliders have conflicting overrides, the settings of the collider with the higher value priority are taken.  For example, if a collider with a **Layer Override Priority** of 1 collides with a Collider with a **Layer Override Priority** of 2, the physics system uses the settings for the Collider with the **Layer Override Priority** of 2. |
| **Include Layers** | Choose which Layers to include in collisions with this collider. |
| **Exclude Layers** | Choose which Layers to exclude in collisions with this collider. |

## Additional resources

* [Collision](collision-section.md)
* [Mesh colliders](mesh-colliders.md)
* [Physics materials](class-PhysicsMaterial.md)
* [Prepare a mesh for mesh colliders](prepare-mesh-for-mesh-collider.md)