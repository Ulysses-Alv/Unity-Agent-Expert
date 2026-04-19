# How collider surface values combine

When two colliders are in contact, the physics system uses the surface properties of each **collider** to calculate the total friction and bounce between the two surfaces.

In Unity, you use the [Physics Material](class-PhysicsMaterial.md) asset to control these parameters. The Physic Material asset is represented in the API by the [`PhysicsMaterial`](../ScriptReference/PhysicsMaterial.md) class.

The Physics Material asset provides two properties: **Friction Combine** ([`PhysicsMaterial.frictionCombine`](../ScriptReference/PhysicsMaterial-frictionCombine.md)) and **Bounce Combine** ([`PhysicsMaterial.bounceCombine`](../ScriptReference/PhysicsMaterial-bounceCombine.md)). These properties each provide four options to control how the physics system calculates the total friction and bounce between two colliders:

| **Priority** | **Property** | **Description** |
| --- | --- | --- |
| 1 | **Maximum** | Use the largest of the two values. |
| 2 | **Multiply** | Use the sum of one value multiplied by the other. |
| 3 | **Minimum** | Use the smallest of the two values. |
| 4 | **Average** | Use the mean average of the two values; that is, the sum of both values, divided by two. |

**Friction Combine** applies to both ****Dynamic Friction**** and **Static Friction**.

The properties in the table are in priority order. Unity takes priority into consideration when the colliders in a collider pair have Physic Material assets with different combine settings. For example, if one Physic Material asset’s **Friction Combine** is set to **Average**, and the other Physic Material asset’s **Friction Combine** is set to **Maximum**, Unity uses the **Maximum** calculation.