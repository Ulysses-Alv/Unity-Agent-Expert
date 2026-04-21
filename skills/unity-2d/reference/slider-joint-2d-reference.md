# Slider Joint 2D

This **joint** allows a **GameObject** controlled by **Rigidbody** physics to slide along a line in space. The object can freely move anywhere along the line in response to **collisions** or forces. Alternatively, it can be also be moved along by a motor force, with limits applied to keep its position within a certain section of the line.

## Properties

| **Property** | **Function** |
| --- | --- |
| **Enable Collision** | Enable this property to enable collisions between the two connected GameObjects. |
| **Connected Rigidbody** | Specify the other object this joint connects to. Leave this as **None** to have the other end of the joint fixed at a point in space defined by the **Connected Anchor** property. Select the circle icon to the right to view a list of GameObjects to connect to. |
| **Auto Configure Connected Anchor** | Enable this property to automatically set the anchor location for the other object this joint connects to. You do not need to enter coordinates for the **Connected Anchor** property if you enable this property. |
| **Anchor** | Define where (in terms of x, y-coordinates on the **Rigidbody 2D**) the end point of the joint connects to this GameObject. |
| **Connected Anchor** | Define where (in terms of x, y-coordinates on the **Rigidbody 2D**) the end point of the joint connects to the other GameObject. |
| **Auto Configure Angle** | Enable this property to automatically detect between the two GameObjects. The joint then keeps the same angle between the two GameObjects. You do not need to manually specify the angle when you enable this property. |
| **Angle** | Enter the angle that the joint keeps between the two objects. |
| **Use Motor** | Use the sliding motor? Check the box for yes. |
| **Motor** | Expand for motor-related property settings. |
| **Motor Speed** | Set the target motor speed (meters/sec). |
| **Maximum Motor Force** | Set the maximum force the motor can apply while attempting to reach the target speed. |
| **Use Limits** | Enable this property to set limits to the linear force. |
| **Translation Limits** | Expand to set the limited distance that the translation can travel. |
| **Lower Translation** | Set the minimum distance the GameObject can be from the connected anchor point. |
| **Upper Translation** | Set the maximum distance the GameObject can be from the connected anchor point. |
| **Break Action** | Set the action taken when either the force or torque threshold is exceeded. |
| **Break Force** | Set the force threshold which if exceeded, will cause the joint to perform the selected **Break Action**. The default value is set to **Infinity**, which can never be exceeded and therefore the **Break Action** can never be taken while the threshold remains at this value. |
| **Break Torque** | Set the torque threshold which if exceeded, will cause the joint to perform the selected **Break Action**. The default value is set to **Infinity**, which can never be exceeded and therefore the **Break Action** can never be taken while the threshold remains at this value. |