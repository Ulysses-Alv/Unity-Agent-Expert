# Sphere collider component reference

[Switch to Scripting](../ScriptReference/SphereCollider.md "Go to SphereCollider page in the Scripting Reference")

The Sphere **collider** is a built-in sphere-shaped collider. It is useful for in-application items such as balls and marbles, or as a simple collider shape that you can stretch and scale to make marbles, projectiles, planets, and other spherical objects. **Sphere colliders** are also useful for items that need to roll and tumble, such as falling boulders.

The Sphere collider has relatively low resource requirements.

| **Property** | **Description** |
| --- | --- |
| **Edit Collider** | Enable the **Edit Collider** button to display the collider’s contact points in the **Scene** view. You can click and drag these contact points to modify the size and shape of the collider. Alternatively, use the **Center** and **Size** properties. |
| **Is Trigger** | If enabled, this collider is used for triggering events, and is ignored by the **physics engine**. |
| **Provides Contacts** | Enable **Provides Contacts** to generate contact information for this collider at all times. Usually, a collider only generates contact data if there is something to send it to; in this case, the messages [`OnCollisionEnter`](../ScriptReference/Collider.OnCollisionEnter.md), [`OnCollisionStay`](../ScriptReference/Collider.OnCollisionStay.md), or [`OnCollisionExit`](../ScriptReference/Collider.OnCollisionExit.md). When **Provides Contacts** is enabled, the collider generates contact data for the physics system at all times. Contact generation is resource-intensive, so **Provides Contacts** is disabled by default. |
| **Material** | Reference to the [Physics Material](class-PhysicsMaterial.md) that determines how this collider interacts with others. |
| **Center** | The position of the collider in the object’s local space. |
| **Radius** | The size of the collider. This is the only available property for resizing the collider; it is not possible to resize along a specific axis (for example, to flatten the sphere into a non-spherical shape). |

## Layer overrides

The Layer Overrides section provides properties that allow you to override the project-wide [Layer-based collision detection](LayerBasedCollision.md) settings for this collider.

| **Property** | **Description** |
| --- | --- |
| **Layer Override Priority** | Define the priority of this collider override. When two colliders have conflicting overrides, the settings of the collider with the higher value priority are taken.  For example, if a collider with a **Layer Override Priority** of 1 collides with a Collider with a **Layer Override Priority** of 2, the physics system uses the settings for the Collider with the **Layer Override Priority** of 2. |
| **Include Layers** | Choose which Layers to include in **collisions** with this collider. |
| **Exclude Layers** | Choose which Layers to exclude in collisions with this collider. |