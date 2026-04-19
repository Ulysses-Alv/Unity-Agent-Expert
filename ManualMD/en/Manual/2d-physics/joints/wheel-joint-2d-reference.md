# Wheel Joint 2D

Use the **Wheel **Joint** 2D** to simulate a rolling wheel, on which an object can move. You can apply motor power to the joint. The wheel uses a suspension spring to maintain its distance from the main body of the vehicle.

## Properties

| **Property** | **Function** |
| --- | --- |
| **Enable **Collision**** | Enable this property to enable collisions between the two connected **GameObjects**. |
| **Connected Rigid Body** | Specify the other GameObject this joint connects to. If you leave this as **None**, the other end of the joint is fixed to a point in space defined by the **Connected Anchor** setting. Select the circle icon to the right to view a list of GameObjects to connect to. |
| **Auto Configure Connected Anchor** | Enable this property to automatically set the anchor location for the other GameObject this **Hinge Joint** 2D connects to. You do not need to enter coordinates for the **Connected Anchor** property if you enable this property. |
| **Anchor** | Define where (in terms of x, y-coordinates on the **Rigidbody 2D**) the end point of the joint connects to this GameObject. |
| **Connected Anchor** | Define where (in terms of x, y-coordinates on the **Rigidbody 2D**) the end point of the joint connects to the other GameObject. |
| **Suspension** | Select this to expand this property’s settings. |
| ****Damping Ratio**** | Set the degree to suppress spring oscillation. In the range 0 to 1, the higher the value, the less movement. |
| **Frequency** | Set the frequency at which the spring oscillates while the GameObjects are approaching the separation distance you want (measured in cycles per second). In the range 0 to 1,000,000 - the higher the value, the stiffer the spring. **Note:** Setting **Frequency** to zero will create the stiffest spring type joint possible. |
| **Angle** | Set the world movement angle for the suspension. |
| **Use Motor** | Enable this to apply motor force to the joint. |
| **Motor** | Select this to expand this property’s settings. |
| **Motor Speed** | Target speed (degrees per second) for the motor to reach. |
| **Maximum Motor Force** | Set the maximum torque (or rotation) the motor can apply when attempting to reach the target speed. |
| **Break Action** | Set the action taken when either the force or torque threshold is exceeded. |
| **Break Force** | Set the force threshold which if exceeded, will cause the joint to perform the selected **Break Action**. The default value is set to **Infinity**, which can never be exceeded and therefore the **Break Action** can never be taken while the threshold remains at this value. |
| **Break Torque** | Set the torque threshold which if exceeded, will cause the joint to perform the selected **Break Action**. The default value is set to **Infinity**, which can never be exceeded and therefore the **Break Action** can never be taken while the threshold remains at this value. |