# Rigidbody component reference

[Switch to Scripting](../ScriptReference/Rigidbody.md "Go to Rigidbody page in the Scripting Reference")

Use the **Rigidbody** component to apply a Rigidbody to your **GameObject**. A Rigidbody provides a physics-based way to control the movement and position of a GameObject. Instead of the [Transform](../ScriptReference/Transform.md) properties, you can use simulated physics forces and torque to move the GameObject, and let the **physics engine** calculate the results. For more information, see [Introduction to Rigidbody Physics](RigidbodiesOverview.md).

To monitor the performance of a Rigidbody, use the [Physics Debug Visualization](PhysicsDebugVisualization.md) tool.

## Properties

| **Property** | **Function** |
| --- | --- |
| **Mass** | Define the mass of the GameObject (in kilograms). **Mass** is set to 1 by default. As in real life, mass does not affect how quickly an item falls under gravity. To simulate resistance forces that slow down movement, use **Drag**. |
| **Linear Damping** | Define the decay rate of a Rigidbody’s linear velocity, to simulate drag, air resistance, or friction. Low values produce a slower decay rate, so that the GameObject moves faster for longer (this is useful for simulating heavy real-world objects). High values produce a faster decay rate, so that the GameObject slows down over a short amount of time (this is useful for simulating lightweight real-world objects). |
| **Angular Damping** | Define the decay rate of a Rigidbody’s rotational velocity, to simulate drag, air resistance, or friction. Low values produce a slower decay rate, so that the GameObject moves faster for longer (this is useful for simulating heavy real-world objects). High values produce a faster decay rate, so that the GameObject slows down over a short amount of time (this is useful for simulating lightweight real-world objects). Note that you cannot make the GameObject stop rotating just by setting its **Angular Damping** to infinity. **Angular Damping** is set to 0.05 by default. |
| **Automatic **Center Of Mass**** | Enable **Automatic Center Of Mass** to use the physics system’s predicted center of mass for the Rigidbody, based on its shape and scale. Disable to set your own X, Y and Z coordinates for the center of mass. |
| **Automatic Tensor** | Enable **Automatic Tensor** to use the physics system’s predicted tensor and tensor rotation for the Rigidbody, based on all connected **colliders**. Like mass, an inertia tensor defines how much force or torque is required to make a Rigidbody move; however, while mass affects linear movement, inertia tensor affects rotational movement. Disable to set your own X, Y and Z coordinates for the tensor instead (see properties below). |
| - **Inertia Tensor** | Define the inertia tensor of this Rigidbody. The higher the Inertia Tensor value is, the more torque is required to make the Rigidbody rotate on its axis. |
| - **Inertia Tensor Rotation** | Define the rotation of the inertia tensor. |
| **Use Gravity** | Toggle the effects of gravity on the Rigidbody. If enabled, the physics system applies a force to move the GameObject in the direction of simulated gravity (by default, down the y axis). **Use Gravity** is enabled by default. |
| **Is Kinematic** | Toggle between physics-based and kinematic movement for the GameObject. When **Is Kinematic** is enabled, the physics system cannot apply forces to move or rotate the GameObject, instead, Unity can only move and rotate it via its **Transform**. See [Introduction to Rigidbody physics: Rigidbody without physics-based movement](RigidbodiesOverview.md) for details. **Is Kinematic** is disabled by default. |
| **Interpolate** | The **Interpolate** setting on a Rigidbody provides two options to smooth the appearance of a Rigidbody’s motion if it appears jittery at run time. These options are **Interpolate** and **Extrapolate**.  Both interpolation and extrapolation calculate the pose of the Rigidbody (that is, the position and rotation) between physics updates. Which one you should choose depends on which option produces the best visual outcome for your use case.  For detailed information on the Interpolate property, see [Apply interpolation to a Rigidbody](rigidbody-interpolation.md). |
| - **None** | Apply no interpolation or extrapolation. This is the default option. |
| - **Interpolate** | Use the pose and velocity of the Rigidbody from the previous two physics updates to calculate and apply the pose of the Rigidbody in the current frame.  Interpolate is more accurate than Extrapolate, but it has a time lag of one physics update. It’s usually best for situations where accuracy is important; for example, if the Rigidbody’s velocity varies, or if there are other physics elements that influence the Rigidbody’s movement. |
| - **Extrapolate** | Use the pose and velocity of the Rigidbody from the previous physics update, and predict the pose of the Rigidbody in the next physics update, to calculate and predict the pose in the current frame.  Extrapolate makes the Rigidbody appear to move slightly ahead of where it should be, and can be slightly inaccurate. It’s usually best for situations where accuracy is not important; for example, if the Rigidbody moves at a constant velocity, and there are no other physics elements that influence the Rigidbody’s movement. |
| **Collision Detection** | Define how the physics system detects collisions between this Rigidbody’s collider and other colliders in the scene. Unity generates one collision per pair of colliders, and determines the method of collision detection based on this **Collision Detection** property.   **Collision Detection** is set to **Discrete** by default. For more information on each collision detection type, refer to [Continuous collision detection](ContinuousCollisionDetection.md). |
| - **Discrete** | The physics system uses discrete **collision** detection to calculate collisions for this Rigidbody’s collider. Select **Discrete** if this Rigidbody is not involved in any fast-moving collisions. **Discrete** **collision detection** is not very computationally intensive. |
| - **Continuous** | The physics system uses sweep-based CCD to calculate collisions between this Rigidbody’s collider and any static colliders (that is, colliders without an associated Rigidbody). Select **Continuous** if this Rigidbody is involved in fast-moving collisions with static colliders. Sweep-based CCD is more computationally intensive than **Discrete** or **Continuous Speculative**. |
| - **Continuous Dynamic** | The physics system uses sweep-based CCD to calculate collisions between this Rigidbody’s collider and all other colliders, except for those that are set to **Discrete** collision detection. Select **Continuous Dynamic** if this Rigidbody is involved in fast-moving collisions with any colliders. Sweep-based CCD is more computationally intensive than **Discrete** or **Continuous Speculative**. |
| - **Continuous Speculative** | The physics system uses **speculative continuous collision detection** to calculate collisions between this Rigidbody’s collider and all other colliders. Select **Continuous Speculative** if collision accuracy is not important for this Rigidbody. Speculative collision detection is more computationally intensive than **Discrete**, but less computationally intensive than **Continuous** or **Continuous Dynamic**. |
| **Constraints** | Place restrictions on the Rigidbody’s movement. |
| - **Freeze Position** | Stops the Rigidbody moving in the world X, Y and Z axes selectively. |
| - **Freeze Rotation** | Stops the Rigidbody rotating around the local X, Y and Z axes selectively. |

## Layer overrides

The Layer Overrides section provides properties that allow you to override the project-wide [Layer-based collision detection](LayerBasedCollision.md) settings for all colliders attached to this Rigidbody.

| **Property** | **Description** |
| --- | --- |
| **Layer Override Priority** | Define the priority of the collider override for colliders attached to this Rigidbody. When colliders have conflicting overrides, Unity uses the settings of the collider with the higher value priority.  For example, if a collider with a **Layer Override Priority** of 1 collides with a Collider with a **Layer Override Priority** of 2, the physics system uses the settings for the Collider with the **Layer Override Priority** of 2. |
| **Include Layers** | Choose which Layers to include in collisions with colliders attached to this Rigidbody. |
| **Exclude Layers** | Choose which Layers to exclude in collisions with colliders attached to this Rigidbody. |

Rigidbody