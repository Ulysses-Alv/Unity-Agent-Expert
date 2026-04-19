# Box collider component reference

[Switch to Scripting](../ScriptReference/BoxCollider.md "Go to BoxCollider page in the Scripting Reference")

The Box **collider** is a built-in cube-shaped collider. It is useful for in-application items such as boxes and crates, or as a simple collider shape that you can stretch and flatten to make ledges, steps, and any other cuboid objects.

The **Box collider** is a low-poly shape, with low resource requirements.

| **Property** | **Description** |
| --- | --- |
| **Edit Collider** | Enable the **Edit Collider** button to display the collider’s contact points in the **Scene** view. You can click and drag these contact points to modify the size and shape of the collider. Alternatively, use the **Center** and **Size** properties. |
| **Is Trigger** | Enable **Is Trigger** to use the collider as a trigger for events. When **Is Trigger** is enabled, other colliders pass through this collider, and trigger the messages [`OnTriggerEnter`](../ScriptReference/Collider.OnTriggerEnter.md), [`OnTriggerStay`](../ScriptReference/Collider.OnTriggerStay.md), and [`OnTriggerExit`](../ScriptReference/Collider.OnTriggerExit.md). |
| **Provides Contacts** | Enable **Provides Contacts** to generate contact information for this collider at all times. Usually, a collider only generates contact data if there is something to send it to; in this case, the messages [`OnCollisionEnter`](../ScriptReference/Collider.OnCollisionEnter.md), [`OnCollisionStay`](../ScriptReference/Collider.OnCollisionStay.md), or [`OnCollisionExit`](../ScriptReference/Collider.OnCollisionExit.md). When **Provides Contacts** is enabled, the collider generates contact data for the physics system at all times. Contact generation is resource-intensive, so **Provides Contacts** is disabled by default. |
| **Material** | Add the [Physics Material component](class-PhysicsMaterial.md) that determines the friction and bounciness of this collider. |
| **Center** | Define the position of the collider on each axis in the **GameObject**’s local space. By default, this is set to (0, 0, 0). |
| **Size** | Define the size of the collider on each axis, in **Unity units**. By default, this is set to (1, 1, 1). |

## Layer overrides

The Layer Overrides section provides properties that allow you to override the project-wide [Layer-based collision detection](LayerBasedCollision.md) settings for this collider.

| **Property** | **Description** |
| --- | --- |
| **Layer Override Priority** | Define the priority of this collider override. When two colliders have conflicting overrides, the settings of the collider with the higher value priority are taken.  For example, if a collider with a **Layer Override Priority** of 1 collides with a Collider with a **Layer Override Priority** of 2, the physics system uses the settings for the Collider with the **Layer Override Priority** of 2. |
| **Include Layers** | Choose which Layers to include in **collisions** with this collider. |
| **Exclude Layers** | Choose which Layers to exclude in collisions with this collider. |