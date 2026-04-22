# Hinge Joint 2D fundamentals

The Hinge **Joint** 2D’s is used to have a joint that allows a **GameObject** to rotate around a particular point, for example a door hinge, wheels, or pendulums.

You can use this joint to make two points overlap. Those two points can be two ****Rigidbody** 2D** components, or a **Rigidbody 2D** component and a fixed position in the world space. Connect the **Hinge Joint** 2D to a fixed position in the world by setting **Connected Rigidbody** to **None**. The joint applies a linear force to both connected Rigidbody 2D GameObjects.

The Hinge Joint 2D has a simulated rotational motor which you can turn on or off. Set the **Maximum Motor Speed** and **Maximum Motor Force** to control the angular speed (**Torque**) and make the two Rigidbody 2D GameObjects rotate in an arc relative to each other. Set the limits of the arc using **Lower Angle** and **Upper Angle**.

## Constraints

Hinge Joint 2D has three simultaneous constraints. All are optional:

* Maintain a relative linear distance between two anchor points on two Rigidbody 2D GameObjects.
* Maintain an angular speed between two anchor points on two Rigidbody 2D GameObjects (limited with a maximum torque in **Maximum Motor Force**).
* Maintain an angle within a specified arc.

You can use this joint to construct physical GameObjects that need to behave as if they are connected with a rotational pivot. For example:

* A see-saw pivot where the horizontal section is connected to the base. Use the joint’s **Angle Limits** to simulate the highest and lowest point of the see-saw’s movement.
* A pair of scissors connected together with a hinge pivot. Use the joint’s **Angle Limits** to simulate the closing and maximum opening of the scissors.
* A simple wheel connected to the body of a car with the pivot connecting the wheel at its center to the car. In this example you can use the Hinge Joint 2D’s motor to rotate the wheel.

## Additional resources

* [Joints 2D](./2d-joints-landing.md)
* [Hinge Joint 2D component reference](hinge-joint-2d-reference.md)