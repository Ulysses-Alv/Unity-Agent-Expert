# Creating a scene with the LowLevelPhysics2D API

Create and debug a **scene** that uses 2D physics with the `LowLevelPhysics2D` API.

**Note**: This documentation is about writing C# scripts using the `LowLevelPhysics2D` API. To use 2D physics in the Unity Editor using components like the Rigidbody 2D component, refer to [2D physics](../2d-physics/2d-physics.md) instead.

| **Topic** | **Description** |
| --- | --- |
| [Create a physics world](2d-physics-api-world.md) | Create a new world to add 2D physics objects to, or fetch the default world Unity automatically creates. |
| [Create a physics object](2d-physics-api-physics-object.md) | Create a 2D physics body, then attach shapes to the body. |
| [Add a sprite](2d-physics-api-add-sprite.md) | Add a SpriteRenderer component to a **GameObject** that has a physics script attached. |
| [Move a GameObject](2d-physics-api-move-gameobject.md) | Configure a 2D physics script to update the **Transform component** of a GameObject. |
| [Draw a debug visualization of objects](2d-physics-api-debug-drawing.md) | Configure how Unity draws 2D physics objects in the Unity Editor or at runtime, and draw your own shapes. |
| [Combine shapes](2d-physics-api-connect-combine-shapes.md) | To combine physics shapes into a compound shape, create a composer using a `PhysicsComposer` object. |
| [Connect objects with joints](2d-physics-api-joints.md) | Create a connection or constraint between two physics objects. |
| [Destroy objects and manage memory](2d-physics-api-destroy.md) | Destroy a world, body, and shape when you no longer need them, and free up memory when necessary. |

## Additional resources

* [Configure LowLevelPhysics2D API scenes](2d-physics-api-properties-landing.md)
* [PhysicsCore2D repository](https://github.com/Unity-Technologies/PhysicsExamples2D/tree/master/PhysicsCore2D) on GitHub