# Configure driving forces on a Configurable Joint

Use the Configurable **Joint** to create custom targets and driving forces for a simulated motor.

To configure driving forces on a **Configurable Joint**:
1. Choose whether you want to [apply a target position](#apply-target-position) or [apply a target velocity](#apply-target-velocity) for the driving force to reach. Targets can be linear (directional) or angular (rotational).
1. [Configure driving forces](#configure-driving-forces) for the respective axes, using properties that simulate a spring and damper system.

**Note**: You can apply both a linear force and a rotational force to a joint. However, you should not apply more than one linear force or more than one rotational force. If you apply more than one, the physics system attempts to solve both at once and produces unwanted results. For example, if you apply a rotational force **Target Rotation** and a linear force ****Target Velocity****, you would create a joint that stays at a specific rotation and moves at a specific linear velocity. However, if you apply two rotational forces, a **Target Rotation** and a **Target Angular Velocity**, the physics system attempts to keep the object at a specific rotation while also rotating it at a specific velocity.

## Apply a target position

To apply a spring-like behavior that always tries to return the object to a specified position or rotation, set a **target position**. You can set a linear position, a rotational position, or both.

* To set a position in the space defined by **Axis**, use the **Target Position**’s **X**, **Y** and **Z** values.
* To set a rotation the space defined by **Axis**, use the **Target Rotation**’s **X**, **Y** and **Z** values.

You must also set the respective linear or angular axis drives, as described on [Driving forces with Configurable Joints](configurable-joints-driving-forces.md).

**Important**: For a Configurable Joint to apply the driving force necessary to reach the target position, the **Position Spring** value on any affected axis drive must be something other than `0`.

## Apply a target velocity

To apply a specific continuous velocity, set a target velocity. You can set a linear motion velocity, a rotational motion velocity, or both.

To set a target linear velocity, use the **Target Velocity**’s **X**, **Y** and **Z** values. These values configure each axis individually, in meters per second.

To set a target angular velocity, there are two options:

* To configure the driving force for each axis individually, set **Rotational Drive Mode** to **X and YZ**, and use the **Target Angular Velocity**’s **X**, **Y** and **Z** values. These values configure each axis individually, in radians per second.
* To configure the rotation without setting specific driving forces, set **Rotational Drive Mode** to **Slerp Drive**. The **Slerp Drive** uses the **quaternion**’s spherical interpolation or “slerp” functionality to reorient the joint. Rather than isolating individual axes, the slerp process finds the minimum total rotation that takes the object from the current orientation to the target, and applies it on all axes as necessary.

You must also set the respective linear or angular axis drives, as described on [Driving forces with Configurable Joints](configurable-joints-driving-forces.md).

**Important**: For a Configurable Joint to apply the driving force necessary to reach the target velocity, the **Position Damper** value on any affected axis drive must be something other than `0`.

## Configure driving forces

For each axis drive you want to apply force on, use **Position Spring** to configure the spring force to apply, and **Position Damper** to configure the damping effect on that spring force.

Apply trial and error on the spring and damper values to achieve the results you need, and test at run time to check that the forces you apply are enough to counter any other forces applied to the object.

**Note**: You can apply both a linear force and a rotational force to a joint. However, you should not apply more than one directional force or more than one rotational force. If you apply more than one, the physics system attempts to solve both at once and produces unwanted results. For example, if you apply a rotational force **Target Rotation** and a linear force **Target Velocity**, you would create a joint that stays at a specific rotation and moves at a specific linear velocity. However, if you apply two rotational forces, a **Target Rotation** and a **Target Angular Velocity**, the physics system attempts to keep the object at a specific rotation while also rotating it at a specific velocity.