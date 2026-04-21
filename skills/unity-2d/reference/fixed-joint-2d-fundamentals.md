# Fixed Joint 2D fundamentals

The aim of this **joint** is to maintain a relative linear and angular offset between two points. Those two points can be two ****Rigidbody** 2D** components or a **Rigidbody 2D** component and a fixed position in the world. (Connect to a fixed position in the world by setting **Connected Rigidbody** to None).

The linear and angular offsets are based upon the relative positions and orientations of the two connected points, so you change the offsets by moving the connected **GameObjects** in your **Scene** view.

The joint applies both linear and torque forces to connected Rigidbody 2D GameObjects. It uses a simulated spring that is pre-configured to be as stiff as the simulation can provide. You can change the spring’s value to make it weaker using the **Frequency** setting.

When the spring applies its force between the GameObjects, it tends to overshoot the desired distance between them and then rebound repeatedly, resulting in a continuous oscillation. The **damping ratio** determines how quickly the oscillation reduces and brings the GameObjects to rest. The frequency is the rate at which it oscillates either side of the target distance; the higher the frequency, the stiffer the spring.

Fixed Joint 2D has two simultaneous constraints:

* Maintain the linear offset between two anchor points on two Rigidbody 2D GameObjects.
* Maintain the angular offset between two anchor points on two Rigidbody 2D GameObjects.

You can use this joint to construct physical GameObjects that need to react as if they are rigidly connected. They can’t move away from each other, they can’t move closer together, and they can’t rotate with respect to each other, such as a bridge made of sections which hold rigidly together.

You can also use this joint to create a less rigid connection that flexes - for example, a bridge made of sections which are slightly flexible.

## Comparing Fixed and Relative joints 2D

It is important to know the major differences between **Fixed Joint** 2D and **Relative Joint 2D**:

* ****Fixed Joint 2D**** is a spring-type joint. **Relative Joint 2D** is a motor-type joint with a maximum force and/or torque.
* **Fixed Joint 2D** uses a spring to maintain the relative linear and angular offsets. **Relative Joint 2D** uses a motor. You can configure a joint’s spring or motor.
* **Fixed Joint 2D** works with anchor points (it’s derived from script **Anchored Joint 2D**); it maintains the relative linear and angular offset between the anchors. **Relative Joint 2D** doesn’t have anchor points (it’s derived directly from script **Joint 2D**).
* **Fixed Joint 2D** cannot modify the relative linear and angular offsets in real time. **Relative Joint 2D** can.

## Additional resources

* [Joints 2D](./2d-joints-landing.md)
* [Fixed Joint 2D component reference](fixed-joint-2d-reference.md)