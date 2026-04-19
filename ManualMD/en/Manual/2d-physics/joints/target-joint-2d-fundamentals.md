# Target Joint 2D fundamentals

Use this **joint** to connect a **Rigidbody** **GameObject** to a point in space. The aim of this joint is to keep zero linear distance between two points: An anchor point on a Rigidbody object and a world space position, called the “**Target**”. The joint applies linear force to the Rigidbody object, it does not apply torque (angular force).

The joint uses a simulated spring. You can set the spring’s stiffness and movement by adjusting its settings. For example, to set a stiff and barely moving spring:

* Set a high (1,000,000 is the highest) **Frequency** == a stiff spring.
* Set a high (1 is the highest) ****Damping Ratio**** == a barely moving spring.

To simulate a looser and more freely moving spring, you would use the following settings:

* Set a low **Frequency** == a loose spring.
* Set a low **Damping Ratio** == a moving spring.

When the spring applies its force between the Rigidbody object and target, it tends to overshoot the distance you have set between them, and then rebound repeatedly, giving in a continuous oscillation. The **Damping Ratio** sets how quickly the Rigidbody object stops moving. The **Frequency** sets how quickly the Rigidbody object oscillates either side of the distance you have specified.

This joint has one constraint:

* Maintain a zero linear distance between the anchor point on a Rigidbody object and a world space position (**Target**).

You can use this joint to construct physical objects that need to move to designated target positions and stay there until another **target position** is selected or the target is cleared. For example:

* A game where players pick up cakes, using a mouse-click, and drag them into to a plate. You can use this joint to move each cake to the plate.

You could also use the joint to allow objects to hang: If the anchor point is not the **center of mass**, then the object will rotate. Such as:

* A game where players pick up boxes. If they use a mouse-click to pick a box up by its corner and drag it, it will hang from the cursor.

## Additional resources

* [Joints 2D](./2d-joints-landing.md)
* [Target Joint 2D component reference](target-joint-2d-reference.md)

TargetJoint2D