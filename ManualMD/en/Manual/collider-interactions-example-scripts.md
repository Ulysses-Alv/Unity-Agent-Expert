# Example scripts for collider events

The following examples demonstrate ways to call events from **collision** functions. They use `OnCollisionEnter` and `OnTriggerEnter` respectively, but the concepts apply to all `OnCollision` and `OnTrigger` functions.

## Example: Different events for different GameObject properties

You can configure your **scripts** to trigger different events based on the properties of the other **collider**’s associated **GameObject**, such as its name or tag. This is useful if, for example, you want to allow some colliders to produce an event, but not others.

The following example prints a different message depending on whether the other collider that has touched this collider has a tag of “Player” or “Enemy”.

```
using UnityEngine;
using System.Collections;

public class DoorObject : Monobehaviour
{
    private void OnCollisionEnter(Collision other)
    {
        if (other.CompareTag("Player"))
        {
            Debug.Log ("The player character has touched the door.")
        }

        if (other.CompareTag("Enemy"))
        {
            Debug.Log ("An enemy character has touched the door!")
        }
    }
}
```

## Example: Send an event message every physics update

The following example uses a trigger collider to produce a hoverpad. The trigger collider is positioned directly on top of a hoverpad GameObject, and applies a constant upward force to any GameObject within its trigger.

```
using UnityEngine;
using System.Collections;

public class HoverPad : MonoBehaviour
{
    // define a value for the upward force calculation
    public float hoverForce = 12f;

    // whenever another collider is in contact with this trigger collider…
    void OnTriggerStay (Collider other)
    {
        // …add an upward force to the Rigidbody of the other collider.
        other.rigidbody.AddForce(Vector3.up * hoverForce, ForceMode.Acceleration) 
    }
}
```