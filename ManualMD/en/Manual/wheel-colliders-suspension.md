# Wheel collider suspension

The [Wheel collider component](class-WheelCollider.md) has a set of properties that simulate a vehicle’s suspension system.

To simulate suspension, the Wheel **collider** sets a ****Target Position**** on the **Suspension Distance** line that it always tries to return the center to, and has **Spring** and **Damper** properties that affect how it moves away from that position and how it returns to that position.

## Target Position

The **Target Position** of the **Wheel collider** is the point along the **Suspension Distance** line that the center of the Wheel collider always returns to when there are no forces (or equal forces) acting upon it.

To set the **Target Position**, specify a coordinate between 0 and 1 along the **Suspension Distance** line.

* A value of 0 sets the target position to the location of the wheel at fully extended suspension (the point furthest from the vehicle body, at the bottom of the **Suspension Distance** line).
* A value of 1 sets the target position to the location of the wheel at fully compressed suspension (the point closest to the vehicle body, at the top of the **Suspension Distance** line).

By default the target position is 0.5, exactly midway between the points of maximum suspension spring extension and **compression**. For most vehicle simulations, a typical value is between 0.3 and 0.7.

## Spring

In a real-world suspension system, the spring connects the wheel and axle to the vehicle’s frame and body, and sustains the weight of the vehicle body. The spring extends and compresses in response to changes in the **terrain**, and absorbs some of the upward force from the ground, so that the vehicle’s body doesn’t respond to every bump in the ground’s surface. How well the spring can absorb energy depends on its stiffness.

In the Unity PhysX simulation, the Wheel collider simulates the spring by moving up and down the **Suspension Distance** line on the vertical Y axis, away from the **Target Position**. The value of the **Spring** property represents the spring’s stiffness (in newtons per meter).

A low value simulates a soft, flexible suspension spring that extends and compresses easily, without requiring very much force. In soft suspension, the flexible spring absorbs more bumps and jolts, so the vehicle’s body movement is smoother.

A high value simulates a stiff suspension spring that has more resistance to extension or compression, and therefore requires more force to move. In hard suspension, more bumps and jolts transfer to the vehicle’s body, but the vehicle overall has more responsive handling.

## Damper

In a real-world suspension system, the damper opposes the suspension spring’s movement and dissipates its stored energy. The strength of the damper’s impact defines how quickly the spring slows down and stops bouncing after being compressed or extended. The damper is often called the shock absorber.

In the Unity PhysX simulation, the Wheel collider simulates the damper or shock absorber by reducing the energy of the suspension **Spring**. The value of the **Damper** property represents the rate at which the energy dissipates (in newton seconds per meter).

A high value simulates a hard damper that dissipates the **Spring**’s energy quickly. Hard damping quickly reduces bounciness and returns the Wheel collider to a steady state at the **Target Position**.

A lower value simulates a soft damper that dissipates the **Spring**’s energy slowly. Soft damping allows more bounciness before the Wheel collider settles back into the **Target Position**.

## Mass and suspension values

By default, the Wheel collider’s **Spring** value is 35000, and the **Damper** value is 4500. These default values assume that the car’s total mass is 1500 kilograms.

To set the vehicle’s mass, add a [Rigidbody](class-Rigidbody.md) component to the vehicle’s root **GameObject**. For a vehicle to work well with the default suspension settings, set the Rigidbody’s **Mass** to the recommended value of 1500. You can then test and iterate based on your specific vehicle setup.

PhysX calculates masses and forces proportionally, meaning they depend on the relative distances between each value. If you want to use a lower mass value for your vehicle (to match other Rigidbody masses on different GameObjects in the scene), you must also decrease the Wheel collider’s **Spring** and **Damper** values in the same proportion. For example, if you set the vehicle’s **Mass** to 15, you should also adjust the **Spring** and **Damper** to 350 and 45, respectively, instead of 35000 and 4500. This ensures consistent and realistic behavior in your vehicle simulation.