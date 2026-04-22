# Event functions

Event functions are a set of predefined callbacks that all [MonoBehaviour](class-MonoBehaviour.md) script components can potentially receive. The callbacks are triggered by various Unity Editor and Engine events, including:

* [Regular frame and physics updates](#regular-updates)
* **Scene** and object [lifecycle events](#initialization-events), such as initialization and destruction of objects in a scene.
* [UI events](#gui-events)
* [Input events](#input-events)
* [Physics events](#physics-events)

Implement the appropriate method signature in your `MonoBehaviour`-derived class to allow your game objects to react to the source events.

For a full list of the available callbacks, refer to the [`MonoBehaviour`](../ScriptReference/MonoBehaviour.md) API reference where they are listed under **Messages**. The rest of this section gives an overview of some of the key groups of event functions.

## Regular update events

A playing game continuously renders a series of frames. A key concept in games programming is making changes to position, state, and behavior of objects just before each frame is rendered. The [`Update`](../ScriptReference/MonoBehaviour.Update.md) function is the main place for this kind of code in Unity. `Update` is called before the frame is rendered and also before animations are calculated.

The physics system also updates in discrete time steps in a similar way to the frame rendering. A separate event function called [`FixedUpdate`](../ScriptReference/MonoBehaviour.FixedUpdate.md) is called just before each physics update. Since the physics updates and frame updates don’t occur with the same frequency, you can get more accurate results from physics code if you place it in the `FixedUpdate` function rather than `Update`.

It’s also sometimes useful to make additional changes at a point after the `Update` and `FixedUpdate` functions have been called for all objects in the scene, and after all animations have been calculated. Some examples of this scenario are:

* When a **camera** should remain trained on a target object. The adjustment to the camera’s orientation must be made after the target object has moved.
* When the script code should override the effect of an animation. For example, to make a character’s head look towards a target object in the scene.

The [`LateUpdate`](../ScriptReference/MonoBehaviour.LateUpdate.md) function can be used for these kinds of situations.

For an overview of how frame and physics updates fit into the overall event function update sequence, refer to [Event function execution order](execution-order.md).

## Initialization events

It’s often useful to be able to call initialization code in advance of any updates that occur during gameplay. The [`Start`](../ScriptReference/MonoBehaviour.Start.md) function is called before the first frame or physics update on an object. The [`Awake`](../ScriptReference/MonoBehaviour.Awake.md) function is called for each object in the scene at the time when the scene loads. Note that although the various objects’ `Start` and `Awake` functions are called in arbitrary order, all instances of `Awake` will have finished before the first `Start` is called. This means that code in a `Start` function can make use of other initializations previously carried out in the `Awake` phase.

For an overview of how initialization functions fit into the overall event function update sequence, refer to [Event function execution order](execution-order.md).

## GUI events

For projects that use the legacy [IMGUI UI system](gui-Basics.md), the [`MonoBehavior.OnGUI`](../ScriptReference/MonoBehaviour.OnGUI.md) callback is called periodically to draw and manage UI elements.

**Note**: Adding an `OnGUI` callback, even if empty, to any object in a scene also adds IMGUI processing to the frame handling process, which creates extra overhead. Only use `OnGUI` if your project uses the [IMGUI UI system](gui-Basics.md), which is no longer recommended.

For more information, refer to [UI systems](UIToolkits.md).

## Input events

The `MonoBehaviour` class also has a set of event functions named with the prefix `OnMouse` (e.g., [`OnMouseOver`](../ScriptReference/MonoBehaviour.OnMouseOver.md), [`OnMouseDown`](../ScriptReference/MonoBehaviour.OnMouseDown.md)) for reacting to user actions with the mouse. These callbacks are only supported for projects using the legacy [Input Manager](class-InputManager.md), which is no longer recommended.

For more information, refer to [Input](Input.md).

## Physics events

The physics system reports **collisions** against an object by calling event functions on that object’s script component.

The [`OnCollisionEnter`](../ScriptReference/MonoBehaviour.OnCollisionEnter.md), [`OnCollisionStay`](../ScriptReference/MonoBehaviour.OnCollisionStay.md) and [`OnCollisionExit`](../ScriptReference/MonoBehaviour.OnCollisionExit.md) functions are called as contact is made, held and broken. The corresponding [`OnTriggerEnter`](../ScriptReference/MonoBehaviour.OnTriggerEnter.md), [`OnTriggerStay`](../ScriptReference/MonoBehaviour.OnTriggerStay.md) and [`OnTriggerExit`](../ScriptReference/MonoBehaviour.OnTriggerExit.md) functions are called when the object’s **collider** is configured as a trigger collider.

For more information on physics event functions, refer to [Collider interactions](collider-interactions.md).

## Additional resources

* [MonoBehaviour](class-MonoBehaviour.md)
* [Event function execution order](execution-order.md)
* [`MonoBehaviour` API reference](../ScriptReference/MonoBehaviour.md)