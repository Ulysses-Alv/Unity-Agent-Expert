# Kinematic Body Type fundamentals

The Kinematic ****Body Type**** **Rigidbody** 2D is designed to move under simulation, but only under very explicit user control. While a [Dynamic](../dynamic/dynamic-body-type-reference.md) Rigidbody 2D is affected by gravity and forces, a Kinematic Rigidbody 2D is not. Because of this, the Kinematic Rigidbody 2D has a lower demand on system resources than a Dynamic Rigidbody 2D, allowing it to be simulated faster.

To reposition a Kinematic Rigidbody 2D, it must be repositioned explicitly via [Rigidbody2D.MovePosition](../../../../../ScriptReference/Rigidbody2D.MovePosition.md) or [Rigidbody2D.MoveRotation](../../../../../ScriptReference/Rigidbody2D.MoveRotation.md). Use physics queries to detect **collisions**, and **scripts** to decide where and how the Rigidbody 2D should move.

A Kinematic Rigidbody 2D can still move via its velocity, but the velocity is not affected by forces or gravity. A Kinematic Rigidbody 2D does not collide with other Kinematic Rigidbody 2Ds or with Static Rigidbody 2Ds and will only collide with Dynamic Rigidbody 2Ds.