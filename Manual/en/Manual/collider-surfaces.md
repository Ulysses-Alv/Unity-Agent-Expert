# Collider surfaces

In real-world physics, objects that can collide have different surface textures and properties that affect how they collide with each other, and how they interact with each other.

To control how objects collide with each other in the physics simulation, you can adjust the friction and bounciness of your **colliders**. In Unity, you use the [Physics Material](class-PhysicsMaterial.md) asset to control these parameters. The Physics Material asset is represented in the API by the [`PhysicsMaterial`](../ScriptReference/PhysicsMaterial.md) class.

For more detailed information on how PhysX applies friction and bounce, see the Nvidia PhysX documentation [Rigid Body Dynamics: Friction and Restitution](https://docs.nvidia.com/gameworks/content/gameworkslibrary/physx/guide/Manual/RigidBodyDynamics.html#friction-and-restitution).

| **Topic** | **Description** |
| --- | --- |
| [Collider surface friction](collider-surface-friction.md) | How Unity handles friction on collider surfaces, and how to configure friction properties. |
| [Collider surface bounciness](collider-surface-bounce.md) | How Unity handles bounciness on collider surfaces, and how to configure bounce properties. |
| [How collider surface values combine](collider-surfaces-combine.md) | How Unity combines the values of surface properties in a collider pair; for example, how it calculates the friction between two colliders that have different friction values. |
| [Create and apply a Physics Material](create-apply-physics-material.md) | Create and configure a Physics Material to define a collider’s surface properties, and apply it to a collider. |
| [Physics Material component reference](class-PhysicsMaterial.md) | Reference page for the Physics Material asset. |