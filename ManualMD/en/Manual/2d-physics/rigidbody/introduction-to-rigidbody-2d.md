# Introduction to Rigidbody 2D

You can attach a Rigidbody 2D component to a **GameObject** to control it with the physics system. The Rigidbody 2D shares similar properties with its standard [Rigidbody](../../class-Rigidbody.md) counterpart, but it’s adapted to 2D development. For example, GameObjects that have a Rigidbody 2D component attached to them can only move along the XY plane and can only rotate on an axis perpendicular to that plane.

## How a Rigidbody 2D works

The Unity Editor’s [Transform](../../class-Transform.md) component defines how to position, rotate, and scale a GameObject (and its child GameObjects) within the **Scene**. When you change this component, it updates other components which can affect where they render or the position of other **colliders**. Unity’s 2D physics system can move colliders and make them interact with each other, so Unity requires a method for the physics system to communicate this movement of colliders back to the Transform components. This movement and connection with colliders is what a Rigidbody 2D component is for. The Rigidbody 2D component overrides the **Transform component** and updates it to the position and/or rotation it defines instead.

**Note:** You can override the Rigidbody 2D by directly modifying the Transform component yourself (because Unity exposes all properties on all components). However, this will cause issues such as unpredictable movement or GameObjects passing through and into each other.

## Collider 2D and Rigidbody 2D interaction

Any Collider 2D component added to the same GameObject or child GameObject is implicitly attached to that Rigidbody 2D GameObject, causing the Collider 2D to move with the Rigidbody 2D. When attached, you should never move the Collider 2D directly using the Transform or any collider offset; move the Rigidbody 2D instead. Moving the Rigidbody 2D provides the best performance and ensures correct **collision** detection. Collider 2Ds attached to the same Rigidbody 2D won’t collide with each other. This means you can create a set of colliders that act effectively as a single compound collider, all moving and rotating in sync with the Rigidbody 2D.

Adding a Rigidbody 2D moves a **sprite** in a physically convincing way by applying forces from the scripting API. When you attach the appropriate collider component to the sprite GameObject, it’s affected by collisions with other moving GameObjects. Using the Unity physics system can simplify many common gameplay mechanics and portray realistic behavior with minimal coding.

**Note:** Although Rigidbody 2Ds are often described as colliding with each other, it’s the Collider 2Ds attached to each of those bodies which collide. Rigidbody 2Ds can’t collide with each other without Colliders.

## Additional resources

* [API documentation on Rigidbody2D](../../../ScriptReference/Rigidbody2D.md)
* [API documentation on Collider2D](../../../ScriptReference/Collider2D.md)

Rigidbody2D