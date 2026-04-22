# Area Effector 2D reference

The Area Effector 2D applies forces within an area defined by the attached **Collider** 2Ds when another (target) Collider 2D comes into contact with the Effector 2D. You can configure the force at any angle with a specific magnitude and random variation on that magnitude. You can also apply both linear and angular drag forces to slow down **Rigidbody** 2Ds.

Collider 2Ds that you use with the Area Effector 2D would typically be set as triggers, so that other Collider 2Ds can overlap with it to have forces applied. Non-triggers will still work, but forces will only be applied when Collider 2Ds come into contact with them.

## Properties

| **Property** | **Function** |
| --- | --- |
| **Use Collider Mask** | Check to enable use of the **Collider Mask** property? If this is not enabled, the Global **Collision** Matrix will be used as the default for all Collider 2Ds. |
| **Collider Mask** | The mask used to select specific Layers allowed to interact with the Area Effector 2D. |
| **Use Global Angle** | Check this to define the **Force Angle** as a global (world-space) angle. If this is not checked, the **Force Angle** is considered a local angle by the **physics engine**. |
| **Force Angle** | The angle of the force to be applied. |
| **Force Magnitude** | The magnitude of the force to be applied. |
| **Force Variation** | The variation of the magnitude of the force to be applied. |
| **Drag** | The linear drag to apply to Rigidbody 2Ds. |
| **Angular Drag** | The angular drag to apply to Rigidbody 2Ds. The options are:  * **Collider**: Defines the target point as the current position of the Collider 2D. Applying force here can generate torque (rotation) if the Collider 2D isn’t positioned at the center of mass. * **Rigidbody**: Defines the target point as the current center of mass of the Rigidbody 2D. Applying force here will never generate torque (rotation). |

AreaEffector2D