# Hinge Joint 2D component reference

This **joint** allows a **GameObject** controlled by **Rigidbody** 2D physics to be attached to a point in space around which it can rotate. The rotation can be left to happen passively (for example, in response to a collision) or can be actively powered by a motor torque provided by the Joint 2D itself. You can set limits to prevent the hinge from making a full rotation, or make more than a single rotation.

## Properties

| **Property** | **Function** |
| --- | --- |
| **Enable **Collision**** | Enable this property to enable collisions between the two connected GameObjects. |
| **Connected Rigidbody** | Specify the other GameObject this joint connects to. If you leave this as **None**, the other end of the joint is fixed to a point in space defined by the **Connected Anchor** setting. Select the circle icon to the right to view a list of GameObjects to connect to. |
| **Auto Configure Connected Anchor** | Enable this property to automatically set the anchor location for the other GameObject this **Hinge Joint** 2D connects to. You do not need to enter coordinates for the **Connected Anchor** property if you enable this property. |
| **Anchor** | Define where (in terms of x, y-coordinates on the **Rigidbody 2D**) the end point of the joint connects to this GameObject. |
| **Connected Anchor** | Define where (in terms of x, y-coordinates on the **Rigidbody 2D**) the end point of the joint connects to the other GameObject. |
| **Use Motor** | Enable this to apply motor force to the joint. |
| **Motor** | Select this to expand this property’s settings. |
| **Motor Speed** | Set the target motor speed (in degrees per second). |
| **Maximum Motor Force** | Set the maximum torque (or rotation) the motor can apply when attempting to reach the target speed. |
| **Use Limits** | Enable this to limit the rotation angle. |
| **Angle Limits** | Select this to expand the Angle limits settings. Set the limits used when **User Limits** is enabled. |
| **Lower Angle** | Set the lower end of the rotation arc allowed by the limit. |
| **Upper Angle** | Set the upper end of the rotation arc allowed by the limit. |
| **Break Action** | Set the action taken when either the force or torque threshold is exceeded. |
| **Break Force** | Set the force threshold which if exceeded, will cause the joint to perform the selected **Break Action**. The default value is set to **Infinity**, which can never be exceeded and therefore the **Break Action** can never be taken while the threshold remains at this value. |
| **Break Torque** | Set the torque threshold which if exceeded, will cause the joint to perform the selected **Break Action**. The default value is set to **Infinity**, which can never be exceeded and therefore the **Break Action** can never be taken while the threshold remains at this value. |

HingeJoint2D