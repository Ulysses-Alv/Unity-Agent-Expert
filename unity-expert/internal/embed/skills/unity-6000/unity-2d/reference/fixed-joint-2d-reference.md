# Fixed Joint 2D component reference

Use the **Fixed **Joint** 2D** to connect two **GameObjects** controlled by **Rigidbody** 2D physics to keep them in a position relative to each other, so the GameObjects are always offset at a given position and angle. It is a spring-type 2D joint for which you don’t need to set maximum forces. You can set the spring to be rigid or soft.

Refer to [Fixed Joint 2D and Relative Joint 2D](fixed-joint-2d-fundamentals.md#fixed-relative) for the differences between ****Fixed Joint** 2D** and ****Relative Joint 2D****.

## Properties

| **Property** | **Function** |
| --- | --- |
| **Enable **Collision**** | Enable this property to enable collisions between the two connected GameObjects. |
| **Connected Rigid Body** | Specify the other object this joint connects to. Leave this as **None** to have the other end of the joint fixed at a point in space defined by the **Connected Anchor** property. Select the circle icon to the right to view a list of GameObjects to connect to. |
| **Auto Configure Connected Anchor** | Enable this property to automatically set the anchor location for the other object this joint connects to. You do not need to enter coordinates for the **Connected Anchor** property if you enable this property. |
| **Anchor** | Define where (in terms of x, y-coordinates on the **Rigidbody 2D**) the end point of the joint connects to this GameObject. |
| **Connected Anchor** | Define where (in terms of x, y-coordinates on the **Rigidbody 2D**) the end point of the joint connects to the other GameObject. |
| ****Damping Ratio**** | Set the degree to suppress spring oscillation. In the range 0 to 1, the higher the value, the less movement. |
| **Frequency** | Set the frequency at which the spring oscillates while the GameObjects are approaching the separation distance you want (measured in cycles per second). In the range 0 to 1,000,000 - the higher the value, the stiffer the spring. **Note:** Setting **Frequency** to zero will create the stiffest spring type joint possible. |
| **Break Action** | Set the action taken when either the force or torque threshold is exceeded. |
| **Break Force** | Set the force threshold which if exceeded, will cause the joint to perform the selected **Break Action**. The default value is set to **Infinity**, which can never be exceeded and therefore the **Break Action** can never be taken while the threshold remains at this value. |
| **Break Torque** | Set the torque threshold which if exceeded, will cause the joint to perform the selected **Break Action**. The default value is set to **Infinity**, which can never be exceeded and therefore the **Break Action** can never be taken while the threshold remains at this value. |

FixedJoint2D