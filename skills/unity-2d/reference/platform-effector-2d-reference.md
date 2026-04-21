# Platform Effector 2D reference

The Platform Effector 2D applies various platform behavior such as one-way **collisions**, removal of side-friction/bounce and more.

Colliders used with the Effector are typically not set as triggers so that other **colliders** can collide with it.

## Properties

| **Property** | **Function** |
| --- | --- |
| **Use Collider Mask** | Select this option to indicate the use of the Collider Mask property. If this isn’t selected, the global collision matrix is chosen as the default for all colliders. |
| **Collider Mask** | The mask used to select specific layers allowed to interact with the Effector. |
| **Use One Way** | Select this option to indicate if one-way collision behavior is used. |
| **Use One Way Grouping** | Ensures that all contacts disabled by the one-way behavior act on all colliders. This is useful when using multiple colliders on the object passing through the platform and they all need to act together as a group. |
| **Surface Arc** | The angle of an arc centered on the local “up” defines the surface which doesn’t allow colliders to pass. Anything outside of this arc is considered for one-way collision. |
| **Use Side Friction** | Select this option to indicate if the friction is used on the platform sides. |
| **Use Side Bounce** | Select to indicate if bounce is used on the platform sides. |
| **Side Arc** | The angle of an arc that defines the sides of the platform centered on the local “left” and “right” of the Effector. Any collision normals within this arc are considered for the “side” behaviors. |

PlatformEffector2D