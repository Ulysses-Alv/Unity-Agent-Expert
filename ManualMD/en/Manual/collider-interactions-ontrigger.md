# OnTrigger events

Trigger **colliders** don’t cause **collisions**. Instead, they detect other colliders that pass through them, and call functions that you can use to initiate events.

Example uses for triggers include:

* When the player reaches a specific area at the end of a corridor, activate a cinematic cutscene.
* When the player character walks within a space in front of a sliding door, trigger an animation to open the door.
* When projectiles pass through a trigger collider in the far distance, disable or destroy the projectile.

Working with trigger colliders primarily involves the following API functions:

* [`Collider.OnTriggerEnter`](../ScriptReference/Collider.OnTriggerEnter.md): Unity calls this function on a trigger collider when it first makes contact with another collider.
* [`Collider.OnTriggerStay`](../ScriptReference/Collider.OnTriggerStay.md): Unity calls this function on a trigger collider once per frame if it detects another Collider inside the trigger collider.
* [`Collider.OnTriggerExit`](../ScriptReference/Collider.OnTriggerExit.md): Unity calls this function on a trigger collider when it ceases contact with another collider.

The following example prints a message to the Console when Unity calls each function.

```
using UnityEngine;
using System.Collections;

public class DoorObject : MonoBehaviour
{
    // “other” refers to the collider on the GameObject inside this trigger
    void OnTriggerEnter (Collider other)
    {
        Debug.Log ("A collider has entered the DoorObject trigger");
    }

    void OnTriggerStay (Collider other)
    {
        Debug.Log ("A collider is inside the DoorObject trigger");
    }
    
    void OnTriggerExit (Collider other)
    {
        Debug.Log ("A collider has exited the DoorObject trigger");
    }
}
```

For examples of practical applications for `OnTrigger` events, see [Example scripts for collider events](collider-interactions-example-scripts.md).