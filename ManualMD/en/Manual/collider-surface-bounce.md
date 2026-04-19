# Collider surface bounciness

In physics, the bounciness of a surface is called the “coefficient of restitution”. The coefficient of restitution is represented as a number between 0 and 1, which defines how much speed an object retains in the opposite direction along the line of impact after a **collision**.

In PhysX, the Physic Material asset’s **Bounce** property has a value of 0–1, which represents the coefficient of restitution.

0 represents a completely inelastic collision (no bounce)
1 represents a perfectly elastic collision (full bounce)
A value between 0 and 1 indicates a partially elastic collision with varying degrees of bounciness.

The default **Bounce** value in Unity is 0.

## Velocity threshold for bounce behavior

In physics simulation, you can also define a threshold at which colliders no longer bounce off each other. If two colliding objects have a relative velocity below the defined value, they do not bounce off each other. To configure the bounce threshold, go to the Physics Settings (**Edit** > **Project Settings** > **Physics**) and set the **Bounce Threshold** to the minimum velocity at which colliders should bounce. This value is set to 2 by default.