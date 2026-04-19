# Relative Joint 2D fundamentals

The aim of this **joint** is to maintain a relative linear and angular distance (offset) between two points. Those two points can be two **Rigidbody2D** components or a **Rigidbody2D** component and a fixed position in the world. **Note:** Connect to a fixed position in the world by setting **Connected **Rigidbody**** to None.

The joint applies a linear and angular (torque) force to both connected Rigidbody objects. It uses a simulated motor that is preconfigured to be quite powerful: It has a high **Max Force** and **Max Torque** limit. You can lower these values to make the motor less powerful motor or turn-off it off completely.

This joint has two simultaneous constraints:

* Maintain the specified linear offset between the two Rigidbody objects.
* Maintain the starting angular offset between the two Rigidbody objects.

You can use this joint to construct physical objects that need to:

* Keep a distance apart from each other, as if they are unable to move further away from each other or closer together. (You decide the distance they are apart from each other. The distance can change in real-time.)
* Rotate with respect to each other only at particular angle. (You decide the angle.)

Some uses may need the connection to be flexible, such as: A space-shooter game where the player has extra gun batteries that follow them. You can use the Relative Joint to give the trailing gun batteries a slight lag when they follow, but make them rotate with the player with no lag.

Some uses may need a configurable force, such as: A game where the **camera** follows a player using a configurable force to keep track.

## Comparing Fixed and Relative joints 2D

**FixedJoint2D** is spring type joint. **RelativeJoint2D** is a motor type joint with a maximum force and/or torque.

* The **Fixed Joint** uses a spring to maintain the relative linear and angular offsets and the Relative joint uses a motor. You can configure a joint’s spring or motor.
* The Fixed joint works with anchor points (it’s derived from script **AnchoredJoint2D**): It maintains the relative linear and angular offset between the anchors. The Relative joint doesn’t have anchor points (it’s derived directly from script **Joint2D**).
* The Relative joint can modify the relative linear and angular offsets in real time: The Fixed joint cannot.

## Additional resources

* [Joints 2D](./2d-joints-landing.md)
* [Relative Joint 2D component reference](relative-joint-2d-reference.md)