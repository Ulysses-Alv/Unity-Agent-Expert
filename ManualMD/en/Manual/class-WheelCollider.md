# Wheel collider component reference

[Switch to Scripting](../ScriptReference/WheelCollider.md "Go to WheelCollider page in the Scripting Reference")

The **Wheel **collider**** is a collider for ground vehicles. It has built-in **collision** detection, wheel physics, and a slip-based tire friction model.

For guidance on using the **Wheel collider**, see the [Unity Wheel collider tutorial](WheelColliderTutorial.md).

| **Property** | **Description** |
| --- | --- |
| **Mass** | Set the mass of the Wheel collider (in kilograms). The default value is 20. |
| **Radius** | Set the distance (in meters) from the center of the Wheel collider to the edge. Use this property to adjust the size of the Wheel collider. The default value is 0.5. |
| **Wheel Damping Rate** | Set the rate at which the wheel’s rotational movement slows down when no forces are present (for example, when there is no acceleration). Higher values apply more damping, and cause the wheel to slow down more quickly. Lower values apply less damping, and the wheel takes longer to come to a stop. Use this property to fine-tune the responsiveness of wheels on a simulated vehicle. The default value is 0.25. |
| **Suspension Distance** | Set the maximum distance along the vertical Y axis that the Wheel collider can move from its target position (the position when equal or no forces are present). When the wheel encounters an uneven surface or an obstacle, it can move up or down within this defined range of vertical movement, simulating a suspension system compressing or extending.   The default value is 0.3. A larger value provides more room for the wheel to move, allowing it to handle larger bumps or obstacles. A smaller value restricts the wheel’s travel, making it less capable of dealing with rough terrain or obstacles with significant height differences. |
| **Force App Point Distance** | Set the point on the wheel where Unity should apply wheel forces. The value expresses the point as distance in meters along the vertical Y axis from the base of the wheel at rest. The default value is 0, which places the point at the wheel’s base. For vehicle simulation, the ideal value is one which applies forces slightly below the vehicle’s **center of mass**. |
| **Center** | Position the center of the wheel relative to the **GameObject**’s Transform. The default value for each axis is 0. |

## Suspension Spring

The Suspension Spring properties simulate the behavior of a vehicle’s suspension system. To do this, the Wheel collider sets a **target position** on the vertical Y axis that it always tries to return to, and has **Spring** and **Damper** properties that apply forces to return it to that position.

For more details on the Suspension Spring properties, refer to [Wheel collider suspension](wheel-colliders-suspension.md).

| **Property** | **Description** |
| --- | --- |
| **Spring** | Set the stiffness of the simulated spring (in newtons per meter). The default value is 35000 N/m, which simulates a normal vehicle. |
| **Damper** | Set the strength of the simulated damper or shock absorber (in newton-seconds per square meter). The default value is 4500 N-s/m2, which simulates a normal vehicle. |
| **Target Position** | Define the Wheel collider’s rest point (that is, the location of the center of the Wheel collider when there are no forces or equal forces acting upon it) along the Suspension Distance. A value of 0 indicates the point of fully extended suspension (the bottom of the suspension distance). A value of 1 indicates the point of fully compressed suspension (the top of the suspension distance). By default the target position is 0.5. For most vehicle simulations, a typical value is between 0.3 and 0.7. |

## Forward Friction and Sideways Friction

The Forward Friction and Sideways Friction properties correspond to the points in the following wheel curve diagram:

![Typical shape of a wheel friction curve.](../uploads/Main/WheelFrictionCurve.png)

Typical shape of a wheel friction curve.

These points are represented in the WheelCurve

See [Wheel collider Friction](wheel-colliders-friction.md) for more details.

| **Property** | **Description** |
| --- | --- |
| **Extremum Slip** | The value of the Extremum point on the Slip axis. The default is 0.4. |
| **Extremum Value** | The value of the Extremum point on the Force axis. The default is 1. |
| **Asymptote Slip** | The value of the Asymptote point on the Slip axis. The default is 0.8. |
| **Asymptote Value** | The value of the Asymptote point on the Force axis. The default is 0.5. |
| **Stiffness** | The multiplier for the **Extremum Value** and **Asymptote Value** values. The default is 1. The stiffness changes the stiffness factor of the friction. Setting this to zero completely disables all friction from the wheel. |

## Layer overrides

The Layer Overrides section provides properties that allow you to override the project-wide [Layer-based collision detection](LayerBasedCollision.md) settings for this collider.

| **Property** | **Description** |
| --- | --- |
| **Layer Override Priority** | Define the priority of this collider override. When two colliders have conflicting overrides, the settings of the collider with the higher value priority are taken.  For example, if a collider with a **Layer Override Priority** of 1 collides with a Collider with a **Layer Override Priority** of 2, the physics system uses the settings for the Collider with the **Layer Override Priority** of 2. |
| **Include Layers** | Choose which Layers to include in collisions with this collider. |
| **Exclude Layers** | Choose which Layers to exclude in collisions with this collider. |