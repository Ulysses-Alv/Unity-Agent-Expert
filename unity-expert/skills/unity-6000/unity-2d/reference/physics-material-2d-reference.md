# Physics Material 2D reference

A ****Physics Material** 2D** is used to adjust the friction and bounce that occur between 2D physics objects when they collide.

For more detailed information on game physics, see [3D Physics](../PhysicsSection.md).

## Use a Physics Material 2D

To create a Physics Material 2D, go to **Assets > Create > 2D > Physics Material 2D**.

| **Property** | **Function** |
| --- | --- |
| **Friction** | Set the coefficient of friction for this **collider**. The range is between 0 and 1. A value of 0 means no friction, like ice. A value of 1 means very high friction, like rubber. |
| **Bounciness** | Set the degree to which **collisions** rebound from the surface. A value of 0 indicates no bounce while a value of 1 indicates a perfect bounce with no loss of energy. |

To use your created **Physics Material 2D**, drag it onto an object with an attached 2D Collider component or drag it onto the collider in the **Inspector** window.

**Note:** The equivalent asset in 3D physics is referred to as a **Physics Material**.

PhysicsMaterial2D