# Rigidbody 2D Simulated property

The Simulated property is common to all available **Body Types**. Use this property to start (enabled) or stop (disabled) a **Rigidbody** 2D and any attached **Collider** 2Ds and **Joint** 2Ds from interacting with the 2D physics simulation. Changing this property is more memory and processor efficient than enabling or disabling individual Collider 2D and Joint 2D components.

When you enable the Simulated property, the following occurs:

* The Rigidbody 2D moves via the simulation (gravity and physics forces are applied).
* Any attached Collider 2Ds continue creating new contacts and continuously reevaluate contacts.
* Any attached Joint 2Ds are simulated and constrain the attached Rigidbody 2D.
* All internal physics objects for Rigidbody 2D, Collider 2D, and Joint 2D stay in memory.

When you disable the Simulated property, the following occurs:

* The Rigidbody 2D isn’t moved by the simulation (gravity and physics forces aren’t applied).
* The Rigidbody 2D doesn’t create new contacts, and any attached Collider 2D contacts are destroyed.
* Any attached Joint 2Ds aren’t simulated, and don’t constrain any attached Rigidbody 2Ds.
* All internal physics objects for Rigidbody 2D, Collider 2D, and Joint 2D remain in memory.

## Improve efficiency with the Simulated property

You can stop and start individual elements of the 2D physics simulation by enabling and disabling physics related components individually on both Collider 2D and Joint 2D components. However, enabling and disabling individual elements of the physics simulations means internal **GameObjects** and physics-based components are constantly created and destroyed, which can cost high memory usage and processor power. Therefore, it’s more efficient to disable the physics simulation entirely rather than disabling the individual components.

**Note**: When you disable a Rigidbody 2D’s Simulated option, any attached Collider 2D is effectively ‘invisible’ and can’t be detected by any physics queries, such as with `Physics.Raycast`.