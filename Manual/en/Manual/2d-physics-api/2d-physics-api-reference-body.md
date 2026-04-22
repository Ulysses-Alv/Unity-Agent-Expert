# Body definition reference for the LowLevelPhysics2D API

**Note**: This documentation is about writing C# scripts using the `LowLevelPhysics2D` API. To use 2D physics in the Unity Editor using components like the Rigidbody 2D component, refer to [2D physics](../2d-physics/2d-physics.md) instead.

Explore the properties you can use to configure a physics body in the Unity Editor. For more information, refer to [Create an object with the LowLevelPhysics2D API](2d-physics-api-physics-object.md).

| **Property** | **Description** |
| --- | --- |
| **Body Type** | Sets how the body behaves in the physics simulation. The options are:  * **Dynamic**: The body has mass. Unity applies forces, collisions, and gravity to the body. * **Kinematic**: Unity applies no forces or gravity, but allows collisions with **Dynamic** bodies. Kinematic bodies move only when you manually reposition them or set their velocity. * **Static**: Unity applies no forces or gravity, but allows collisions with **Dynamic** bodies. |
| **Body Constraints** | Places restrictions on the movement and rotation of the body. The options are:  * **None**: Allows the body to move and rotate freely. * **Position X**: Stops the body moving along the x-axis. * **Position Y**: Stops the body moving along the y-axis. * **Position**: Stops the body moving along either the x-axis or the y-axis. * **Rotation**: Stops the body rotating around the z-axis. * **All**: Stops the body moving or rotating. |
| **Transform Write Mode** | Sets the method Unity uses to copy the position of the physics body to the Transform component on the GameObject. The options are:  * **Current**: Copies the position and rotation at the end of the physics update. * **Interpolate**: Copies the position and rotation interpolated between the previous state and the current state. * **Extrapolate**: Copies the next position and rotation, extrapolated from the current position, rotation, and velocities. * **Off**: Doesn’t copy the position or rotation.  For more information, refer to [Move a GameObject](2d-physics-api-move-gameobject.md). |
| **Position** | Sets the position of the body in world space. |
| **Rotation** | Sets the rotation of the body in degrees. |
| **Linear Velocity** | Sets the initial velocity of the body in meters per second. The default is 0. |
| **Angular Velocity** | Sets the initial rotational velocity of the body in degrees per second. The default is 0. |
| **Linear Damping** | Sets how quickly the linear velocity decreases, to simulate drag, air resistance, or friction. Low values produce a slower decay rate, so that the body moves faster for longer, like a heavy object. High values produce a faster decay rate, so that the body slows down over a short amount of time, like a lightweight object. The default is 0. |
| **Angular Damping** | Sets how quickly the angular velocity decreases, to simulate drag, air resistance, or friction. The default is 0. |
| **Gravity Scale** | Scales the gravity applied to the body. Use a positive value to apply gravity in the same direction, a negative value to reverse gravity, or 0 to apply no gravity. The default is 1, which means gravity isn’t scaled up or down. |
| **Sleep Threshold** | Sets the speed below which Unity temporarily removes the body from physics calculations to save processor time. The value is in meters per second. The default is 0.05. |
| **Fast Rotation Allowed** | Removes the limit on how fast a physics body rotates. The limit avoids forces becoming too large and objects passing through each other incorrectly. Enabling this property is recommended only for circular objects. For more information, refer to [Move a GameObject](2d-physics-api-move-gameobject.md). |
| **Fast Collisions Allowed** | Force continuous **collision** detection for this body. Enable this property for a very fast-moving object that passes through other objects. Enabling this setting can reduce performance. For more information, refer to [Configure collisions between LowLevelPhysics2D API objects](2d-physics-api-collisions-enable.md). |
| **Sleeping Allowed** | Removes the body from physics calculations when it’s not moving or colliding, to save processor time. The recommended best practice is to leave this property enabled. |
| **Awake** | Sets the body as awake rather than in the sleep state. |
| **Enabled** | Removes the body from the physics simulation, but keeps the body and its handles alive. Enabling this property also removes any shapes or **joints** attached to the body from the physics simulation. |

## Additional resources

* [Create an object with the LowLevelPhysics2D API](2d-physics-api-physics-object.md)
* [Configure LowLevelPhysics2D API scenes with definitions](2d-physics-api-definitions.md)
* [Configure global LowLevelPhysics2D API settings](2d-physics-api-global-settings.md)
* [Shape definition reference for the LowLevelPhysics2D API](2d-physics-api-reference-shape.md)