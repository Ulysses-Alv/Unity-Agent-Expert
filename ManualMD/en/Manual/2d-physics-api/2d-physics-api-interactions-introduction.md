# Introduction to collision in the LowLevelPhysics2D API

**Note**: This documentation is about writing C# scripts using the `LowLevelPhysics2D` API. To use 2D physics in the Unity Editor using components like the Rigidbody 2D component, refer to [2D physics](../2d-physics/2d-physics.md) instead.

In Unity, a **collision** happens when two objects occupy the same physical space.

To detect collision between `LowLevelPhysics2D` API objects, Unity uses the `PhysicsShape` objects attached to the `PhysicsBody`. The shape defines the physics body for the purposes of physical collisions. The shape is similar to the **Collider** 2D component in the regular [2D physics](../2d-physics/2d-physics.md) system.

Shapes don’t need to be the same shape as the body, and you can add multiple shapes to one body.

For more information about adding a shape, refer to [Create a physics object](2d-physics-api-physics-object.md).

## Collision types

In the `LowLevelPhysics2D` API, collisions are enabled by default.

How a shape reacts to collisions depends on its type:

* Collider: The shape can’t enter other objects. When the shape makes contact with another shape, it reports a contact event. This is the default.
* Trigger: The shape goes through other objects. When the shape enters another shape, it reports a trigger event.

For more information, refer to [Configure collisions between LowLevelPhysics2D API objects](2d-physics-api-collisions-enable.md).

To detect collisions and overlaps using callbacks or events, refer to [Detect collisions between objects](2d-physics-api-collision-handle.md).

## Collision layers

Unity uses layers to determine which objects collide with each other.

By default the following applies:

* Objects use [GameObject layers](../Layers.md), and all objects are on the built-in layer called **Default**.
* Objects collide with objects on all other layers.

To change this behavior, change the layers that objects are on, and the layers they collide with. For more information, refer to [Configure collisions between 2D physics objects](2d-physics-api-collisions-enable.md).

In the `LowLevelPhysics2D` API, you can use a different set of 64 layers, instead of the standard set of 32 **GameObject** layers. For more information, refer to [Use up to 64 layers](2d-physics-api-collisions-enable.md#use-up-to-64-layers).

## Additional resources

* [Check LowLevelPhysics2D API interactions by casting rays](2d-physics-api-raycasting.md)
* [Optimize the LowLevelPhysics2D API with multithreading](2d-physics-api-multithreading.md)