# Constraint components

Constraint components allow you to control the position, rotation, and scale of **GameObjects** based on other objects in your **scene**. They provide a way to create relationships between objects without using parent-child hierarchies or writing custom **scripts**.

Constraints are useful for character **rigging**, **camera** control, and creating relationships between objects. You can use multiple constraints on a single GameObject and blend between multiple source objects to achieve complex behaviors.

| **Topic** | **Description** |
| --- | --- |
| **[Constraint components introduction](Constraints.md)** | Understand Constraint components and how they control GameObject transforms based on other objects. |
| **[Aim Constraint component](class-AimConstraint.md)** | Constrain a GameObject to aim at one or more target objects. |
| **[Look At Constraint component](class-LookAtConstraint.md)** | Rotate a GameObject to look at a target object. |
| **[Parent Constraint component](class-ParentConstraint.md)** | Control a GameObject’s transform as if it were a child of one or more parent objects. |
| **[Position Constraint component](class-PositionConstraint.md)** | Constrain a GameObject’s position to follow one or more source objects. |
| **[Rotation Constraint component](class-RotationConstraint.md)** | Constrain a GameObject’s rotation to match one or more source objects. |
| **[Scale Constraint component](class-ScaleConstraint.md)** | Constrain a GameObject’s scale to match one or more source objects. |

## Additional resources

* [Add components to GameObjects](unity-components.md)
* [Manage components and their values](InspectorManageComponents.md)