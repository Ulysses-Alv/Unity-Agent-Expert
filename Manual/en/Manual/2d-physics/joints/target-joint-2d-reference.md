# Target Joint 2D

This **joint** connects to a specified target, rather than another **Rigidbody** object as other joints do. This behaves in a similar way to a spring type joint.

## Properties

| **Property** | **Function** |
| --- | --- |
| **Anchor** | Define where (in terms of x, y-coordinates on the **Rigidbody 2D**) the end point of the joint connects to this **GameObject**. |
| **Target** | Define where (in terms of x, y-coordinates in world space) the other end of the joint attempts to move. |
| **Auto Configure Target** | Enable this property to automatically set the other end of the joint to the current position of the GameObject. **Note:** When this option is enabled, the target changes as you move the GameObject but the target will not change if the option is not enabled. |
| **Max Force** | Set the force that the joint can apply when attempting to move the object to the **target position**. The higher the value, the higher the maximum force. |
| ****Damping Ratio**** | Set the degree to suppress spring oscillation. In the range 0 to 1, the higher the value, the less movement. |
| **Frequency** | Set the frequency at which the spring oscillates while the GameObjects are approaching the separation distance you want (measured in cycles per second). In the range 0 to 1,000,000 - the higher the value, the stiffer the spring. **Note:** Setting **Frequency** to zero will create the stiffest spring type joint possible. |
| **Break Action** | Set the action taken when either the force or torque threshold is exceeded. |
| **Break Force** | Set the force threshold which if exceeded, will cause the joint to perform the selected **Break Action**. The default value is set to **Infinity**, which can never be exceeded and therefore the **Break Action** can never be taken while the threshold remains at this value. |