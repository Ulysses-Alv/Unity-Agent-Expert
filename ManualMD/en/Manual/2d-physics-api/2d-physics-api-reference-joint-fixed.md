# Fixed joint definition reference for the LowLevelPhysics2D API

**Note**: This documentation is about writing C# scripts using the `LowLevelPhysics2D` API. To use 2D physics in the Unity Editor using components like the Rigidbody 2D component, refer to [2D physics](../2d-physics/2d-physics.md) instead.

Explore the properties you can use to configure a `LowLevelPhysics2D` API fixed **joint** in the Unity Editor.

A **fixed joint** fixes two bodies to the same position and rotation with a spring.

| **Property** | **Sub-property** | **Description** |
| --- | --- | --- |
| **Local Anchor A** | N/A | Sets where one end of the joint connects to the first body. |
| N/A | **Position** | Sets the position where the end of the joint connects, relative to the position of the first body. |
| N/A | **Rotation** | Sets the rotation of the anchor in degrees. |
| **Local Anchor B** | N/A | Sets where the other end of the joint connects to the second body. |
| N/A | **Position** | Sets the position where the end of the joint connects, relative to the position of the second body. |
| N/A | **Rotation** | Sets the rotation of the anchor in degrees. |
| **Linear Frequency** | N/A | Sets the stiffness of the spring that tries to keep the positions of the two bodies the same, in Hertz (cycles per second). A higher value makes the joint less stiff. The default is 0. |
| **Linear Damping** | N/A | Sets how quickly the spring stops pulling or pushing. A higher value makes the joint settle more quickly. The default is 0. |
| **Angular Frequency** | N/A | Sets the stiffness of the spring that tries to keep the rotation of the two bodies the same, in Hertz (cycles per second). A higher value makes the joint’s rotation less stiff. The default is 0. |
| **Angular Damping** | N/A | Sets how quickly the spring stops rotating. A higher value makes the rotation of the joint settle more quickly. The default is 0. |
| **Force Threshold** | N/A | Sets the amount of linear force that creates an `OnJointThreshold2D` event. |
| **Torque Threshold** | N/A | Sets the amount of rotational force that creates an `OnJointThreshold2D` event. |
| **Tuning Frequency** | N/A | Sets the overall stiffness of the joint in Hertz (cycles per second). A higher value makes the joint less stiff. |
| **Tuning Damping** | N/A | Sets how quickly overall the joint stops pulling or pushing. A higher value makes the joint settle more quickly. |
| **Draw Scale** | N/A | Scales the joint Unity draws as a debug visualization. For more information, refer to [Draw a debug visualization of objects](2d-physics-api-debug-drawing.md). |
| **Collide Connected** | N/A | Enables the shapes creating **collision** or trigger events when they come into contact with each other. For more information, refer to [Collisions and interactions in the LowLevelPhysics2D API](2d-physics-api-interactions-landing.md). |

## Additional resources

* [`PhysicsFixedJointDefinition`](../../ScriptReference/LowLevelPhysics2D.PhysicsFixedJointDefinition.md)
* [Configure LowLevelPhysics2D API scenes with definitions](2d-physics-api-definitions.md)
* [Configure global LowLevelPhysics2D API settings](2d-physics-api-global-settings.md)