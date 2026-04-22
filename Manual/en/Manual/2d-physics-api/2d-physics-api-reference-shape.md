# Shape definition reference for the LowLevelPhysics2D API

**Note**: This documentation is about writing C# scripts using the `LowLevelPhysics2D` API. To use 2D physics in the Unity Editor using components like the Rigidbody 2D component, refer to [2D physics](../2d-physics/2d-physics.md) instead.

Explore the properties you can use to configure a physics shape in the Unity Editor. For more information, refer to [Create an object with the LowLevelPhysics2D API](2d-physics-api-physics-object.md).

## Surface Material

The properties in the **Surface Material** section determine how the shape interacts with other shapes.

| Property | Description |
| --- | --- |
| **Friction** | Sets the coefficient of friction for this **collider**. A value of 0 means no friction, like ice. A value of 1 means very high friction, like rubber. The default is 0.6. |
| **Bounciness** | Sets how bouncy the surface is, and how much other colliders bounce off it. A value of 0 means the surface is not at all bouncy, like soft clay. A value of 1 means the surface is very bouncy, like rubber. The default is 0. |
| **Friction Mixing** | Determines the method Unity uses to mix the friction of two objects when they make contact. The options are:  * **Average**: Uses the average of the two values. * **Mean**: Uses the geometric mean of the two values. The geometric mean multiplies the two values then returns the square root. * **Multiply**: Multiplies the two values together. * **Minimum**: Uses the smaller value. * **Maximum**: Uses the larger value. |
| **Bounciness Mixing** | Determines the method Unity uses to mix the bounciness of two objects when they make contact. The options are:  * **Average**: Uses the average of the two values. * **Mean**: Uses the geometric mean of the two values. The geometric mean multiplies the two values then returns the square root. * **Multiply**: Multiplies the two values together. * **Minimum**: Uses the smaller value. * **Maximum**: Uses the larger value. |
| **Friction Priority** | Determines which shape contributes its **Friction Mixing** mode when two shapes come into contact. Unity uses the **Friction Mixing** mode from the shape with the highest **Friction Priority** value. If the two shapes have the same **Friction Priority** value, Unity uses the highest `SurfaceMaterial.MixingMode` enumeration value from the two shapes. |
| **Bounciness Priority** | Determines which shape contributes its **Bounciness Mixing** mode when two shapes come into contact. Unity uses the **Bounciness Mixing** mode from the shape with the highest **Bounciness Priority** value. If the two shapes have the same **Bounciness Priority** value, Unity uses the highest `SurfaceMaterial.MixingMode` enumeration value from the two shapes. |
| **Rolling Resistance** | Sets how resistant the shape is to rolling. The range of values is 0 to 1, where 0 means no rolling resistance and 1 means full rolling resistance. |
| **Tangent Speed** | Sets the speed the surface moves other objects that come into contact with it, in meters per second. For example, if you set **Tangent Speed** to 5, the surface acts like a conveyor belt that moves objects along the surface at 5 meters per second. You can use a positive or negative value to set the direction of movement. |
| **Custom Color** | Overrides the color Unity uses to draw the debug visualization of the shape. The alpha value is ignored. The default is [`Color.Clear`](../../ScriptReference/Color.Clear.md), which is black with an alpha of 0, and means Unity uses different colors to represent the current state of the shape. For more information, refer to [Draw a debug visualization of objects](2d-physics-api-debug-drawing.md). |

## Contact Filter

The properties in the **Contact Filter** section determine which other shapes the shape collides with. For more information, refer to [Configure collisions between objects](2d-physics-api-collisions-enable.md).

| Property | Description |
| --- | --- |
| **Categories** | Sets the layers this shape belongs to. |
| **Contacts** | Sets the layers this shape collides with. |
| **Group Index** | Assigns the shape to a group, and overrides the **Categories** and **Contacts** properties. Use the following values:  * 0: Assigns no group. The shape uses the **Categories** and **Contacts** properties to determine collisions. * Positive value: The shape always collides with other shapes that have the same **Group Index**. * Negative value: The shape doesn’t collide with other shapes that have the same **Group Index**. |

## Other properties

| Property | Description |
| --- | --- |
| **Density** | Sets the density in kg/m2. The default is 1. Use this property to create a hollow shape for example. |
| **Is Trigger** | Enables the shape going through other shapes without generating a **collision** response. Instead, the shape generates a trigger event when another shape overlaps. For more information, refer to [Introduction to collision in the LowLevelPhysics2D API](2d-physics-api-interactions-introduction.md). |
| **Trigger Events** | Produces a trigger event when the shape starts and ends overlapping another shape. This property works whether you enable or disable **Is Trigger** on this shape. To fetch events, refer to [`PhysicsWorld.triggerBeginEvents`](../../ScriptReference/LowLevelPhysics2D.PhysicsWorld.triggerBeginEvents.md). |
| **Contact Events** | Produces a contact event when the shape starts and ends contact with another shape. This property has no effect if the shape is attached to a body set to **Static**. To fetch events, refer to [`PhysicsWorld.contactBeginEvents`](../../ScriptReference/LowLevelPhysics2D.PhysicsWorld.contactBeginEvents.md). |
| **Hit Events** | Produces a hit event when the shape collides with another shape. This property has no effect if the shape is attached to a body set to **Static**. To fetch events, refer to [`PhysicsWorld.shapeHitEvents`](../../ScriptReference/LowLevelPhysics2D.PhysicsWorld.shapeHitEvents.md). |
| **Contact Filter Callbacks** | Calls the `IContactFilterCallback` method of the shape if it makes contact with another shape, so you can run your own code that determines whether the shapes should collide. |
| **Pre Solve Callbacks** | Calls the `OnPreSolve2D` method of the dynamic body before Unity calculates the collision response, so you can run your own code. |
| **Start Static Contacts** | Produces a contact event if the shape is attached to a body set to **Static**, and the shape makes contact with another shape. Enabling this property can reduce performance if you have many static shapes. |
| **Start Mass Update** | Updates the mass of the body when the shape is created. Disabling this property can improve performance if you add multiple shapes to the same body, but you must call [ApplyMassFromShapes](../../ScriptReference/LowLevelPhysics2D.PhysicsBody.ApplyMassFromShapes.md) to recalculate after you add a shape. |

## Additional resources

* [Create an object with the LowLevelPhysics2D API](2d-physics-api-physics-object.md)
* [Configure LowLevelPhysics2D API scenes with definitions](2d-physics-api-definitions.md)
* [Configure global LowLevelPhysics2D API settings](2d-physics-api-global-settings.md)