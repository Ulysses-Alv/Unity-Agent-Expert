# Use collisions to trigger other events

When two **colliders** make contact, they call functions that you can use to trigger other events in your project. You can place any code you like in these functions to respond to the **collision** event.

Collider events require configuration via C# script; you cannot configure them using only the user interface.

There are two collider event types:

* Collision events: Collision events occur when two colliders make contact, and neither collider has **Is Trigger** enabled. The most common collision functions are [`Collider.OnCollisionEnter`](../ScriptReference/Collider.OnCollisionEnter.md), [`Collider.OnCollisionStay`](../ScriptReference/Collider.OnCollisionStay.md), and [`Collider.OnCollisionExit`](../ScriptReference/Collider.OnCollisionExit.md).
* Trigger events: Trigger events occur when two colliders make contact, at least one collider has **Is Trigger** enabled, and at least one collider has a **Rigidbody** or ArticulationBody. The most common trigger functions are [`Collider.OnTriggerEnter`](../ScriptReference/Collider.OnTriggerEnter.md), [`Collider.OnTriggerStay`](../ScriptReference/Collider.OnTriggerStay.md), and [`Collider.OnTriggerExit`](../ScriptReference/Collider.OnTriggerExit.md).

A collider that has **Is Trigger** enabled is called a **trigger collider**. Trigger colliders do not physically collide with other colliders; instead, they create a space that sends an event when other colliders pass through it.

Note: The 2D physics system has equivalent functions with **2D** appended to the name (for example, `OnCollisionEnter2D`). For details of these 2D functions, refer to the [`MonoBehaviour`](../ScriptReference/MonoBehaviour.md) API class documentation.