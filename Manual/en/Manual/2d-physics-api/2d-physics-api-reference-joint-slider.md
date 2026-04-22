# Slider joint definition reference for the LowLevelPhysics2D API

**Note**: This documentation is about writing C# scripts using the `LowLevelPhysics2D` API. To use 2D physics in the Unity Editor using components like the Rigidbody 2D component, refer to [2D physics](../2d-physics/2d-physics.md) instead.

Explore the properties you can use to configure a `LowLevelPhysics2D` API slider **joint** in the Unity Editor.

A slider joint allows one body to restrict a second body so it can only slide along one axis. The axis is defined by the rotation of the first object.

| **Property** | **Sub-property** | **Description** |
| --- | --- | --- |
| **Local Anchor A** | N/A | Sets where one end of the joint connects to the first body. The rotation of this body determines the axis the second object can slide along. |
| N/A | **Position** | Sets the position where the end of the joint connects, relative to the position of the first body. |
| N/A | **Rotation** | Sets the rotation of the anchor in degrees. |
| **Local Anchor B** | N/A | Sets where the other end of the joint connects to the second body. |
| N/A | **Position** | Sets the position where the end of the joint connects, relative to the position of the second body. |
| N/A | **Rotation** | Sets the rotation of the anchor in degrees. |
| **Enable Spring** | N/A | Enables **Spring Target Translation**, **Spring Frequency**, and **Spring Damping**. The spring connects the two bodies, and stretches and compresses to try to bring the bodies together. The strength of the pull is proportional to the current distance between the bodies. |
| **Spring Target Translation** | N/A | Sets the target speed of the spring pulling and pushing. |
| **Spring Frequency** | N/A | Sets the stiffness of the spring in Hertz (cycles per second). A higher value makes the spring less stiff. The default is 0. |
| **Spring Damping** | N/A | Sets how quickly the spring stops pulling or pushing. A higher value makes the spring settle more quickly. The default is 0. |
| **Enable Motor** | N/A | **Enables Motor Speed** and **Max Motor Force**, which make the joint act as a motor to pull or push the two connected bodies. |
| **Motor Speed** | N/A | Sets the target motor speed, in meters per second. |
| **Max Motor Force** | N/A | Sets the maximum force that the motor applies to reach the target speed, in newtons. |
| **Enable Limit** | N/A | Enables **Min Distance Limit** and **Max Distance Limit**, which restrict the position of the second body. |
| **Lower Translation Limit** | N/A | Sets the minimum position for the second object in meters. |
| **Upper Translation Limit** | N/A | Sets the maximum position for the second object in meters. |
| **Force Threshold** | N/A | Sets the amount of linear force that creates an `OnJointThreshold2D` event. |
| **Torque Threshold** | N/A | Sets the amount of rotational force that creates an `OnJointThreshold2D` event. |
| **Tuning Frequency** | N/A | Sets the overall stiffness of the joint in Hertz (cycles per second). A higher value makes the joint less stiff. |
| **Tuning Damping** | N/A | Sets how quickly overall the joint stops pulling or pushing. A higher value makes the joint settle more quickly. |
| **Draw Scale** | N/A | Scales the joint Unity draws as a debug visualization. For more information, refer to [Draw a debug visualization of objects](2d-physics-api-debug-drawing.md). |
| **Collide Connected** | N/A | Enables the shapes creating **collision** or trigger events when they come into contact with each other. For more information, refer to [Collisions and interactions in the LowLevelPhysics2D API](2d-physics-api-interactions-landing.md). |

## Additional resources

* [`PhysicsSliderJointDefinition`](../../ScriptReference/LowLevelPhysics2D.PhysicsSliderJointDefinition.md)
* [Configure LowLevelPhysics2D API scenes with definitions](2d-physics-api-definitions.md)
* [Configure global LowLevelPhysics2D API settings](2d-physics-api-global-settings.md)