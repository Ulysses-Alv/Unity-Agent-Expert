# OnCollision events

Collision events occur when two non-trigger colliders make contact.

Example uses for collision events include:

* When a projectile hits a target, destroy both the projectile and the enemy.
* When a player character touches a door, trigger an animation to open the door.
* When a player character touches a power-up, increase the player’s size.

Working with collision events primarily involves the following API functions:

* [`Collider.OnCollisionEnter`](../ScriptReference/Collider.OnCollisionEnter.md): Unity calls this function on each collider when two colliders first make contact.
* [`Collider.OnCollisionStay`](../ScriptReference/Collider.OnCollisionStay.md): Unity calls this function on each collider once per physics update while two colliders are in contact.
* [`Collider.OnCollisionExit`](../ScriptReference/Collider.OnCollisionExit.md): Unity calls this function on each collider when two colliders cease contact.

For collision events, at least one of the objects involved must have a dynamic physics body (that is, a Rigidbody or ArticulationBody that has **Is Kinematic** disabled). If both **GameObjects** in a collision are kinematic physics bodies, the collision does not call `OnCollision` functions.

The following example prints a message to the console when Unity calls each function.

```
using UnityEngine;
using System.Collections;

public class DoorObject : MonoBehaviour
{
    // “other” refers to the collider that is touching this collider
    void OnCollisionEnter (Collision other)
    {
        Debug.Log ("A collider has made contact with the DoorObject Collider");
    }

    void OnCollisionStay (Collision other)
    {
        Debug.Log ("A collider is in contact with the DoorObject Collider");
    }
    
    void OnCollisionExit (Collision other)
    {
        Debug.Log ("A collider has ceased contact with the DoorObject Collider");
    }
}
```

For examples of practical applications for `OnCollision` events, refer to [example scripts for collider events](collider-interactions-example-scripts.md).

## Additional resources

* [Collider interactions](collider-interactions.md)
* [Collision](collision-section.md)
* [Collider](../ScriptReference/Collider.md)