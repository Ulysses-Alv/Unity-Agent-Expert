# Wheel Joint 2D fundamentals

Use this **joint** to simulate wheels and suspension. The aim of the joint is to keep the position of two points on a line that extends to infinity, whilst at the same time making them overlap. Those two points can be two **Rigidbody2D** components or a **Rigidbody2D** component and a fixed position in the world. (Connect to a fixed position in the world by setting **Connected Rigidbody** to None).

Wheel Joint 2D acts like a combination of a [Slider Joint 2D](./slider-joint-2d-reference.md) (without its motor or limit constraints) and a [Hinge Joint 2D](./hinge-joint-2d-reference.md) (without its limit constraint).

The joint applies a linear force to both connected rigid body objects to keep them on the line, an angular motor to rotate the objects on the line, and a spring to simulate wheel suspension.

Set the **Maximum Motor Speed** and **Maximum Motor Force** (torque, in this joint) to control the angular motor speed, and make the two rigid body objects rotate.

You can set the wheel suspension stiffness and movement in order to simulate different degrees of suspension. For example, to simulate a stiff, barely moving suspension:

* Set a high (1,000,000 is the highest) **Frequency** == stiff suspension.
* Set a high (1 is the highest) ****Damping Ratio**** == barely moving suspension.

To simulate a looser and more freely moving suspension, you would use the following settings:

* Set a low **Frequency** == loose suspension.
* Set a low **Damping Ratio** == moving suspension.

It has two simultaneous constraints:

* Maintain a zero relative linear distance away from a specified line between two anchor points on two rigid body objects.
* Maintain an angular speed between two anchor points on two rigid body objects. (Set the speed via the **Maximum Motor Speed** option and maximum torque via **Maximum Motor Force**.)

You can use this joint to construct physical objects that need to react as if they are connected with a rotational pivot but cannot move away from a specified line. Such as:

* Simulating wheels with a motor to drive the wheels and a line defining the movement allowed for the suspension.

## Behavior difference to the Wheel Collider

Unlike the [Wheel Collider](../../class-WheelCollider.md) used with 3D physics, the Wheel Joint 2D uses a separate ****Rigidbody**** object for the wheel, which rotates when the force is applied. (The Wheel Collider, by contrast, simulates the suspension using a raycast and the wheel’s rotation is purely a graphical effect). The wheel object will typically be a [Circle Collider 2D](../collider/circle-collider-2d-reference.md) with a [Physics Material 2D](../physics-material-2d-reference.md) that gives the right amount of traction for your gameplay.

## Additional resources

* [Joints 2D](./2d-joints-landing.md)
* [Wheel Joint 2D component reference](wheel-joint-2d-reference.md)

WheelJoint2D