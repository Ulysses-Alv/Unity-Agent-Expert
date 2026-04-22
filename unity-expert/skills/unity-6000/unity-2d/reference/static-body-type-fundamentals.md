# Static Body Type fundamentals

A Static **Rigidbody** 2D is designed to not move under simulation at all. If anything collides with it, a Static Rigidbody 2D behaves like an immovable object (as though it has infinite mass). It is also the least resource intensive ****Body Type****. A Static body only collides with Dynamic Rigidbody 2Ds.

**Note**: Having two Static Rigidbody 2Ds collide is not supported, since they are not designed to move.

## Use a large number of Static Collider 2D

Aside from setting the Rigidbody 2D to the **Static** Body Type, there is another scenario where a Static Rigidbody 2D is created. This is when a **GameObject** with a **Collider** 2D component does not have a Rigidbody 2D component at all. All Collider 2Ds without a Rigidbody 2D component are internally considered to be attached to a single hidden **Static** Rigidbody 2D component.

This means that you are able to create a large number of Static Collider 2Ds as you do not have to add a Rigidbody 2D component for each individual GameObject. Both methods of creating Static Collider 2Ds have their advantages, depending on the scenario.

If an individual Static Collider 2D needs to be moved or reconfigured at runtime, then add a Rigidbody 2D component and set it to the **Static** Body Type, as it is faster to simulate the Collider 2D when it has its own Rigidbody 2D. If a group of Collider 2Ds needs to be moved or reconfigured at runtime, it is faster to have them all be children of the single hidden parent Rigidbody 2D than to move each GameObject individually.

**Note**: As stated above, Static Rigidbody 2Ds are designed not to move, and **collisions** between two Static Rigidbody **2D objects** that intersect are not registered. However, Static Rigidbody 2Ds and Kinematic Rigidbody 2Ds will interact with each other if one of their Collider 2Ds is set to be a **trigger**. There is also a feature that changes what a Kinematic body will interact with (see [Use Full Kinematic Contacts](../kinematic/kinematic-body-type-reference.md) for more information).