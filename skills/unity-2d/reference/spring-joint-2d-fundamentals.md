# Spring Joint 2D fundamentals

This **joint** behaves like a spring, while keeping a linear distance between two points. You set this via the **Distance** setting. Those two points can be two **Rigidbody2D** components or a **Rigidbody2D** component and a fixed position in the world. (Connect to a fixed position in the world by setting **Connected Rigidbody** to None). The joint applies a linear force to both rigid bodies. It doesn’t apply torque (an angle force).

The joint uses a simulated spring. You can set the spring’s stiffness and movement:

A stiff, barely moving spring…

* A high (1,000,000 is the highest) **Frequency** == a stiff spring.
* A high (1 is the highest) ****Damping Ratio**** == a barely moving spring.

A loose, moving spring…

* A low **Frequency** == a loose spring.
* A low **Damping Ratio** == a moving spring.

When the spring applies its force between the objects, it tends to overshoot the distance you have set between them, and then rebound repeatedly, giving in a continuous oscillation. The **Damping Ratio** sets how quickly the objects stop moving. The **Frequency** sets how quickly the objects oscillate either side of the target distance.

This joint has one constraint:

* Maintain a zero linear distance between two anchor points on two **Rigidbody** objects.

You can use this joint to construct physical objects that need to react as if they are connected together using a spring or a connection which allows rotation. For example:

* A character whose body is composed of multiple objects that act as if they are semi-rigid. Use the **Spring Joint** to connect the character’s body parts together, allowing them to flex to and from each other. You can specify whether the body parts are held together loosely or tightly.

**Note:** Spring Joint 2D uses a Box 2D spring-joint, which the Distance Joint 2D also uses with its frequency set to zero.

## Additional resources

* [Joints 2D](./2d-joints-landing.md)
* [Spring Joint 2D component reference](spring-joint-2d-reference.md)