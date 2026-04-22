# Wheel collider friction

The Wheel **collider** calculates wheel friction separately from the rest of the **physics engine**, and ignores standard [Physics Material](class-PhysicsMaterial.md) settings. Instead, it uses a slip-based friction model, which provides more realistic behavior.

In most cases, the **Wheel collider**’s default friction settings are sufficient to create a working vehicle. To work with the wheel friction settings, you should have an understanding of forward slip (also called longitudinal slip) and sideways slip (also called lateral slip) in the context of real-world vehicle dynamics. This understanding is necessary if you want to build an extremely realistic vehicle or you want to have finer control of the friction and slip behavior of your wheels.

On real-world vehicles, wheels can exert high forces and high friction at low slip because the rubber stretches to compensate for the slip. When the slip gets too high, the wheel starts to slide or spin, which reduces the amount of force they exert. Unity uses a wheel friction curve to define and describe this behavior.

## Wheel friction curve properties

The Wheel collider has two groups of properties for wheel friction: **Forward Friction** and **Sideways Friction**. Each group has the same four settings (refer to the [Wheel collider component reference](class-WheelCollider.md) for detailed information on each property):

* **Extremum Slip** ([`WheelFrictionCurve.extremumSlip`](../ScriptReference/WheelFrictionCurve-extremumSlip.md))
* **Extremum Value** ([`WheelFrictionCurve.extremumValue`](../ScriptReference/WheelFrictionCurve-extremumValue.md))
* **Asymptote Slip** ([`WheelFrictionCurve.asymptoteSlip`](../ScriptReference/WheelFrictionCurve-asymptoteSlip.md))
* **Asymptote Value** ([`WheelFrictionCurve.asymptoteValue`](../ScriptReference/WheelFrictionCurve-asymptoteValue.md))

These [`WheelFrictionCurve`](../ScriptReference/WheelFrictionCurve.md) properties describe the curve that demonstrates the relationship between the slip and the wheel’s force in a typical wheel friction setup. There is one curve for forward friction, and one for sideways friction.

![Typical shape of a wheel friction curve.](../uploads/Main/WheelFrictionCurve.png)

Typical shape of a wheel friction curve.

The curve starts at 0 slip and 0 force. When the slip increases, the force also increases until it reaches a maximum force that the wheel can maintain (the **Extremum** point). The coordinates for this point are (`ExtremumSlip` , `ExtremumValue`).

When the slip increases further, the wheel starts to slide or spin, and can no longer maintain the maximum force. As a result, the force decreases until it reaches a point where it can remain steady and consistent as the slip continues to increase (the **Asymptote** point). The coordinates for this point are (`AsymptoteSlip` , `AsymptoteValue`).